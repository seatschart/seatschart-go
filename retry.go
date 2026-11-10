// Retry, timeout and idempotency helpers for the SeatsChart Go client.
// Copied into the generated package by scripts/generate-server-sdks.sh.
//
// Throttled (429) and unavailable (503) responses are retried with exponential backoff and full
// jitter, honouring Retry-After. Only requests that are safe to repeat are retried: idempotent
// methods (GET, HEAD, OPTIONS, PUT, DELETE) and requests carrying an Idempotency-Key header, which
// is resent unchanged so the API replays the stored outcome instead of applying it twice.
//
// Every attempt is bounded (as in seats.io's SDKs): 10 s to connect (DefaultConnectTimeout, on the
// transport NewRetryingHTTPClient builds) and 30 s from sending the request until the response body
// is read (DefaultRequestTimeout, RetryPolicy.Timeout; 0 disables it). An attempt that times out is
// retried like a 503 when it is safe to repeat; the last one fails with a *TimeoutError (IsTimeout).
//
//	cfg := seatschart.NewConfiguration()
//	cfg.HTTPClient = seatschart.NewRetryingHTTPClient(nil, seatschart.DefaultRetryPolicy())
//	client := seatschart.NewAPIClient(cfg)
package seatschart

import (
	"context"
	"crypto/rand"
	"errors"
	"fmt"
	"io"
	mathrand "math/rand"
	"net"
	"net/http"
	"strconv"
	"strings"
	"time"
)

// IdempotencyKeyHeader carries the idempotency key of transactional calls.
const IdempotencyKeyHeader = "Idempotency-Key"

// Default timeouts of one attempt.
const (
	DefaultConnectTimeout = 10 * time.Second
	DefaultRequestTimeout = 30 * time.Second
)

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

// RetryPolicy bounds the retries of 429/503 responses and timed-out attempts.
type RetryPolicy struct {
	MaxAttempts    int           // total attempts including the first (>= 1)
	BaseDelay      time.Duration // attempt n waits up to BaseDelay * 2^(n-1)
	MaxDelay       time.Duration // cap of a computed backoff wait
	MaxRetryAfter  time.Duration // a longer Retry-After is not waited for
	Statuses       map[int]bool  // retryable statuses
	Timeout        time.Duration // per attempt, until the body is read (0 = no timeout)
	ConnectTimeout time.Duration // dial timeout of the transport NewRetryingHTTPClient builds (0 = Go's default)
	RetryOnTimeout bool          // retry a timed-out attempt when the request is safe to repeat
	Sleep          func(time.Duration)
	Random         func() float64
}

// DefaultRetryPolicy: 4 attempts, 250 ms base, 8 s cap, Retry-After up to 30 s, statuses 429 and 503,
// 30 s per attempt, 10 s to connect, timed-out attempts retried.
func DefaultRetryPolicy() RetryPolicy {
	return RetryPolicy{MaxAttempts: 4, BaseDelay: 250 * time.Millisecond, MaxDelay: 8 * time.Second,
		MaxRetryAfter: 30 * time.Second, Statuses: map[int]bool{429: true, 503: true},
		Timeout: DefaultRequestTimeout, ConnectTimeout: DefaultConnectTimeout, RetryOnTimeout: true}
}

// TimeoutError reports a request whose last attempt got no complete response within the timeout
// (Code "timeout"). The outcome on the server is unknown: retry writes with the same Idempotency-Key.
// It satisfies net.Error (Timeout() is true) and unwraps to the transport error.
type TimeoutError struct {
	Method   string
	Path     string
	Limit    time.Duration // the per-attempt timeout that elapsed
	Attempts int
	Err      error
}

func (e *TimeoutError) Error() string {
	return fmt.Sprintf("seatschart: %s %s timed out after %s (%d attempt(s)): %v", e.Method, e.Path, e.Limit, e.Attempts, e.Err)
}

// Unwrap returns the transport error.
func (e *TimeoutError) Unwrap() error { return e.Err }

// Timeout is always true (net.Error).
func (e *TimeoutError) Timeout() bool { return true }

// Temporary is always true (net.Error).
func (e *TimeoutError) Temporary() bool { return true }

// Code is the stable error code "timeout".
func (e *TimeoutError) Code() string { return "timeout" }

// IsTimeout reports whether err (as returned by the generated client) is a request timeout.
func IsTimeout(err error) bool {
	var te *TimeoutError
	if errors.As(err, &te) {
		return true
	}
	var ne net.Error
	return errors.As(err, &ne) && ne.Timeout()
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
	sleep := t.Policy.Sleep
	if sleep == nil {
		sleep = time.Sleep
	}
	safe := IsRetrySafe(req.Method, req.Header)
	current := req
	for attempt := 1; ; attempt++ {
		resp, timedOut, err := t.send(base, current)
		canRetry := safe && attempt < t.Policy.MaxAttempts
		if timedOut {
			timeoutErr := &TimeoutError{Method: req.Method, Path: req.URL.Path, Limit: t.Policy.Timeout, Attempts: attempt, Err: err}
			if !canRetry || !t.Policy.RetryOnTimeout {
				return nil, timeoutErr
			}
			next, ok := replay(req)
			if !ok {
				return nil, timeoutErr
			}
			wait, _ := t.Policy.Delay(attempt, 0, false)
			sleep(wait)
			current = next
			continue
		}
		if err != nil || !canRetry || !t.Policy.Statuses[resp.StatusCode] {
			return resp, err
		}
		retryAfter, has := ParseRetryAfter(resp.Header.Get("Retry-After"), time.Now())
		wait, ok := t.Policy.Delay(attempt, retryAfter, has)
		if !ok {
			return resp, nil
		}
		next, ok := replay(req)
		if !ok {
			return resp, nil // body cannot be replayed
		}
		_, _ = io.Copy(io.Discard, resp.Body)
		_ = resp.Body.Close()
		sleep(wait)
		current = next
	}
}

// replay clones req with a fresh body for another attempt; false when the body cannot be replayed.
func replay(req *http.Request) (*http.Request, bool) {
	next := req.Clone(req.Context())
	if req.Body != nil && req.Body != http.NoBody {
		if req.GetBody == nil {
			return nil, false
		}
		body, err := req.GetBody()
		if err != nil {
			return nil, false
		}
		next.Body = body
	}
	return next, true
}

// send runs one attempt under Policy.Timeout. The deadline keeps running while the caller reads the body
// (it is released when the body is closed); timedOut reports that this deadline, not the caller's, fired.
func (t *RetryTransport) send(base http.RoundTripper, req *http.Request) (*http.Response, bool, error) {
	if t.Policy.Timeout <= 0 {
		resp, err := base.RoundTrip(req)
		return resp, err != nil && req.Context().Err() == nil && IsTimeout(err), err
	}
	parent := req.Context()
	ctx, cancel := context.WithTimeout(parent, t.Policy.Timeout)
	resp, err := base.RoundTrip(req.WithContext(ctx))
	if err != nil {
		timedOut := parent.Err() == nil && (errors.Is(ctx.Err(), context.DeadlineExceeded) || IsTimeout(err))
		cancel()
		return nil, timedOut, err
	}
	resp.Body = &deadlineBody{ReadCloser: resp.Body, ctx: ctx, parent: parent, cancel: cancel,
		method: req.Method, path: req.URL.Path, timeout: t.Policy.Timeout}
	return resp, false, nil
}

// deadlineBody keeps the attempt's deadline over the body and turns its expiry into a *TimeoutError.
type deadlineBody struct {
	io.ReadCloser
	ctx     context.Context
	parent  context.Context
	cancel  context.CancelFunc
	method  string
	path    string
	timeout time.Duration
}

func (b *deadlineBody) Read(p []byte) (int, error) {
	n, err := b.ReadCloser.Read(p)
	if err != nil && err != io.EOF && b.parent.Err() == nil && errors.Is(b.ctx.Err(), context.DeadlineExceeded) {
		return n, &TimeoutError{Method: b.method, Path: b.path, Limit: b.timeout, Attempts: 1, Err: err}
	}
	return n, err
}

func (b *deadlineBody) Close() error {
	err := b.ReadCloser.Close()
	b.cancel()
	return err
}

// NewTransport is http.DefaultTransport with the given dial (connect) timeout; 0 keeps Go's default.
func NewTransport(connectTimeout time.Duration) http.RoundTripper {
	def, ok := http.DefaultTransport.(*http.Transport)
	if !ok || connectTimeout <= 0 {
		return http.DefaultTransport
	}
	tr := def.Clone()
	tr.DialContext = (&net.Dialer{Timeout: connectTimeout, KeepAlive: 30 * time.Second}).DialContext
	return tr
}

// NewRetryingHTTPClient wraps base (a new client when nil) with policy. A base without its own
// Transport gets NewTransport(policy.ConnectTimeout).
func NewRetryingHTTPClient(base *http.Client, policy RetryPolicy) *http.Client {
	if base == nil {
		base = &http.Client{}
	}
	c := *base
	inner := base.Transport
	if inner == nil {
		inner = NewTransport(policy.ConnectTimeout)
	}
	c.Transport = &RetryTransport{Base: inner, Policy: policy}
	return &c
}
