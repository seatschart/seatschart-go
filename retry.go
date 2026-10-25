// Retry and idempotency helpers for the SeatsChart Go client.
// Copied into the generated package by scripts/generate-server-sdks.sh.
//
// Throttled (429) and unavailable (503) responses are retried with exponential backoff and full
// jitter, honouring Retry-After. Only requests that are safe to repeat are retried: idempotent
// methods (GET, HEAD, OPTIONS, PUT, DELETE) and requests carrying an Idempotency-Key header, which
// is resent unchanged so the API replays the stored outcome instead of applying it twice.
//
//	cfg := seatschart.NewConfiguration()
//	cfg.HTTPClient = seatschart.NewRetryingHTTPClient(nil, seatschart.DefaultRetryPolicy())
//	client := seatschart.NewAPIClient(cfg)
package seatschart

import (
	"crypto/rand"
	"fmt"
	"io"
	mathrand "math/rand"
	"net/http"
	"strconv"
	"strings"
	"time"
)

// IdempotencyKeyHeader carries the idempotency key of transactional calls.
const IdempotencyKeyHeader = "Idempotency-Key"

// NewIdempotencyKey returns a random UUIDv4 for one logical operation; reuse it for every retry.
func NewIdempotencyKey() string {
	var b [16]byte
	if _, err := rand.Read(b[:]); err != nil {
		panic(err)
	}
	b[6] = (b[6] & 0x0f) | 0x40
	b[8] = (b[8] & 0x3f) | 0x80
	return fmt.Sprintf("%x-%x-%x-%x-%x", b[0:4], b[4:6], b[6:8], b[8:10], b[10:16])
}

// RetryPolicy bounds the retries of 429/503 responses.
type RetryPolicy struct {
	MaxAttempts   int           // total attempts including the first (>= 1)
	BaseDelay     time.Duration // attempt n waits up to BaseDelay * 2^(n-1)
	MaxDelay      time.Duration // cap of a computed backoff wait
	MaxRetryAfter time.Duration // a longer Retry-After is not waited for
	Statuses      map[int]bool  // retryable statuses
	Sleep         func(time.Duration)
	Random        func() float64
}

// DefaultRetryPolicy: 4 attempts, 250 ms base, 8 s cap, Retry-After up to 30 s, statuses 429 and 503.
func DefaultRetryPolicy() RetryPolicy {
	return RetryPolicy{MaxAttempts: 4, BaseDelay: 250 * time.Millisecond, MaxDelay: 8 * time.Second,
		MaxRetryAfter: 30 * time.Second, Statuses: map[int]bool{429: true, 503: true}}
}

// IsRetrySafe reports whether repeating the request cannot apply its effect twice.
func IsRetrySafe(method string, header http.Header) bool {
	switch strings.ToUpper(method) {
	case http.MethodGet, http.MethodHead, http.MethodOptions, http.MethodPut, http.MethodDelete:
		return true
	}
	return strings.TrimSpace(header.Get(IdempotencyKeyHeader)) != ""
}

// ParseRetryAfter returns the Retry-After wait (delta-seconds or HTTP date) and whether it was valid.
func ParseRetryAfter(value string, now time.Time) (time.Duration, bool) {
	raw := strings.TrimSpace(value)
	if raw == "" {
		return 0, false
	}
	if secs, err := strconv.ParseFloat(raw, 64); err == nil {
		if secs < 0 {
			return 0, false
		}
		return time.Duration(secs * float64(time.Second)), true
	}
	at, err := http.ParseTime(raw)
	if err != nil {
		return 0, false
	}
	if d := at.Sub(now); d > 0 {
		return d, true
	}
	return 0, true
}

// Delay is the wait before retry number attempt (1-based); ok=false when Retry-After is too long.
func (p RetryPolicy) Delay(attempt int, retryAfter time.Duration, hasRetryAfter bool) (time.Duration, bool) {
	if hasRetryAfter {
		return retryAfter, retryAfter <= p.MaxRetryAfter
	}
	shift := attempt - 1
	if shift < 0 {
		shift = 0
	}
	if shift > 30 {
		shift = 30
	}
	ceiling := p.BaseDelay << uint(shift)
	if ceiling > p.MaxDelay || ceiling <= 0 {
		ceiling = p.MaxDelay
	}
	random := p.Random
	if random == nil {
		random = mathrand.Float64
	}
	return time.Duration(random() * float64(ceiling)), true
}

// RetryTransport is an http.RoundTripper applying a RetryPolicy.
type RetryTransport struct {
	Base   http.RoundTripper
	Policy RetryPolicy
}

func (t *RetryTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	base := t.Base
	if base == nil {
		base = http.DefaultTransport
	}
	resp, err := base.RoundTrip(req)
	if err != nil || !IsRetrySafe(req.Method, req.Header) {
		return resp, err
	}
	sleep := t.Policy.Sleep
	if sleep == nil {
		sleep = time.Sleep
	}
	for attempt := 1; attempt < t.Policy.MaxAttempts && t.Policy.Statuses[resp.StatusCode]; attempt++ {
		retryAfter, has := ParseRetryAfter(resp.Header.Get("Retry-After"), time.Now())
		wait, ok := t.Policy.Delay(attempt, retryAfter, has)
		if !ok {
			break
		}
		next := req.Clone(req.Context())
		if req.Body != nil && req.Body != http.NoBody {
			if req.GetBody == nil {
				break // body cannot be replayed
			}
			body, gerr := req.GetBody()
			if gerr != nil {
				break
			}
			next.Body = body
		}
		_, _ = io.Copy(io.Discard, resp.Body)
		_ = resp.Body.Close()
		sleep(wait)
		resp, err = base.RoundTrip(next)
		if err != nil {
			return resp, err
		}
	}
	return resp, err
}

// NewRetryingHTTPClient wraps base (http.DefaultClient when nil) with policy.
func NewRetryingHTTPClient(base *http.Client, policy RetryPolicy) *http.Client {
	if base == nil {
		base = &http.Client{}
	}
	c := *base
	c.Transport = &RetryTransport{Base: base.Transport, Policy: policy}
	return &c
}
