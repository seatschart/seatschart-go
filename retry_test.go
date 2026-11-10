package seatschart

import (
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"sync"
	"sync/atomic"
	"testing"
	"time"
)

// slowServer answers after delay for the first `slow` calls (or until the client gives up), then at once.
func slowServer(slow int32, delay time.Duration, keys *[]string, mu *sync.Mutex) *httptest.Server {
	var calls int32
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		mu.Lock()
		*keys = append(*keys, r.Header.Get(IdempotencyKeyHeader))
		mu.Unlock()
		if atomic.AddInt32(&calls, 1) <= slow {
			select {
			case <-r.Context().Done():
				return
			case <-time.After(delay):
			}
		}
		w.WriteHeader(200)
		_, _ = w.Write([]byte(`{"ok":true}`))
	}))
}

func timeoutPolicy(timeout time.Duration) RetryPolicy {
	p := DefaultRetryPolicy()
	p.Timeout = timeout
	p.Sleep = func(time.Duration) {}
	return p
}

func TestDefaultTimeouts(t *testing.T) {
	p := DefaultRetryPolicy()
	if p.Timeout != 30*time.Second || p.ConnectTimeout != 10*time.Second || !p.RetryOnTimeout {
		t.Fatalf("defaults %+v", p)
	}
	rt, ok := NewRetryingHTTPClient(nil, p).Transport.(*RetryTransport)
	if !ok {
		t.Fatal("not a RetryTransport")
	}
	if tr, ok := rt.Base.(*http.Transport); !ok || tr == http.DefaultTransport {
		t.Fatalf("connect timeout transport not built: %T", rt.Base)
	}
}

func TestTimedOutGetIsRetriedThenFailsWithTimeoutError(t *testing.T) {
	var keys []string
	var mu sync.Mutex
	srv := slowServer(10, 2*time.Second, &keys, &mu)
	defer srv.Close()
	p := timeoutPolicy(50 * time.Millisecond)
	p.MaxAttempts = 3
	_, err := NewRetryingHTTPClient(nil, p).Get(srv.URL + "/api/events")
	var te *TimeoutError
	if !IsTimeout(err) || !errors.As(err, &te) {
		t.Fatalf("want *TimeoutError, got %v", err)
	}
	if te.Code() != "timeout" || te.Attempts != 3 || te.Path != "/api/events" || te.Limit != 50*time.Millisecond {
		t.Fatalf("error %+v", te)
	}
	mu.Lock()
	defer mu.Unlock()
	if len(keys) != 3 {
		t.Fatalf("calls %d", len(keys))
	}
}

func TestTimedOutKeyedPostIsRetriedWithTheSameKey(t *testing.T) {
	var keys []string
	var mu sync.Mutex
	srv := slowServer(1, 2*time.Second, &keys, &mu)
	defer srv.Close()
	req, _ := http.NewRequest(http.MethodPost, srv.URL+"/api/seats/reserve", strings.NewReader(`{"eventId":"e"}`))
	req.Header.Set(IdempotencyKeyHeader, "order-11")
	resp, err := NewRetryingHTTPClient(nil, timeoutPolicy(50*time.Millisecond)).Do(req)
	if err != nil || resp.StatusCode != 200 {
		t.Fatalf("resp %v err %v", resp, err)
	}
	_ = resp.Body.Close()
	mu.Lock()
	defer mu.Unlock()
	if len(keys) != 2 || keys[0] != "order-11" || keys[1] != "order-11" {
		t.Fatalf("keys %v", keys)
	}
}

func TestUnkeyedPostTimeoutIsNotResent(t *testing.T) {
	var keys []string
	var mu sync.Mutex
	srv := slowServer(5, 2*time.Second, &keys, &mu)
	defer srv.Close()
	_, err := NewRetryingHTTPClient(nil, timeoutPolicy(50*time.Millisecond)).Post(srv.URL, "application/json", strings.NewReader("{}"))
	var te *TimeoutError
	if !errors.As(err, &te) || te.Attempts != 1 {
		t.Fatalf("err %v", err)
	}
	mu.Lock()
	defer mu.Unlock()
	if len(keys) != 1 {
		t.Fatalf("calls %d", len(keys))
	}
}

func TestTimeoutCoversTheBody(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		_, _ = w.Write([]byte(`{"partial":`))
		w.(http.Flusher).Flush()
		select {
		case <-r.Context().Done():
		case <-time.After(2 * time.Second):
		}
	}))
	defer srv.Close()
	resp, err := NewRetryingHTTPClient(nil, timeoutPolicy(80*time.Millisecond)).Get(srv.URL)
	if err != nil {
		t.Fatalf("headers should arrive: %v", err)
	}
	defer resp.Body.Close()
	_, err = io.ReadAll(resp.Body)
	if !IsTimeout(err) {
		t.Fatalf("body read: want timeout, got %v", err)
	}
}

func TestZeroTimeoutWaitsAndCallerCancelIsNotATimeout(t *testing.T) {
	var keys []string
	var mu sync.Mutex
	srv := slowServer(1, 120*time.Millisecond, &keys, &mu)
	defer srv.Close()
	resp, err := NewRetryingHTTPClient(nil, timeoutPolicy(0)).Get(srv.URL)
	if err != nil || resp.StatusCode != 200 {
		t.Fatalf("no timeout: %v", err)
	}
	_ = resp.Body.Close()
	slow := slowServer(5, 2*time.Second, &keys, &mu)
	defer slow.Close()
	client := NewRetryingHTTPClient(&http.Client{Timeout: 50 * time.Millisecond}, timeoutPolicy(time.Minute))
	_, err = client.Get(slow.URL)
	var te *TimeoutError
	if err == nil || errors.As(err, &te) {
		t.Fatalf("the caller's own deadline must not become a retried *TimeoutError: %v", err)
	}
}

type script struct {
	mu       sync.Mutex
	statuses []int
	headers  []map[string]string
	keys     []string
	bodies   []string
}

func (s *script) handler(w http.ResponseWriter, r *http.Request) {
	s.mu.Lock()
	defer s.mu.Unlock()
	buf := new(strings.Builder)
	if r.Body != nil {
		b := make([]byte, 1024)
		n, _ := r.Body.Read(b)
		buf.Write(b[:n])
	}
	s.keys = append(s.keys, r.Header.Get(IdempotencyKeyHeader))
	s.bodies = append(s.bodies, buf.String())
	i := len(s.keys) - 1
	if i < len(s.headers) {
		for k, v := range s.headers[i] {
			w.Header().Set(k, v)
		}
	}
	status := 200
	if i < len(s.statuses) {
		status = s.statuses[i]
	}
	w.WriteHeader(status)
}

func testPolicy(waits *[]time.Duration) RetryPolicy {
	p := DefaultRetryPolicy()
	p.Sleep = func(d time.Duration) { *waits = append(*waits, d) }
	p.Random = func() float64 { return 0.5 }
	return p
}

func TestKeyedPostRetriesWithSameKeyAndBody(t *testing.T) {
	s := &script{statuses: []int{429, 200}, headers: []map[string]string{{"Retry-After": "1"}}}
	srv := httptest.NewServer(http.HandlerFunc(s.handler))
	defer srv.Close()
	var waits []time.Duration
	client := NewRetryingHTTPClient(nil, testPolicy(&waits))
	req, _ := http.NewRequest(http.MethodPost, srv.URL+"/api/seats/reserve", strings.NewReader(`{"eventId":"e"}`))
	req.Header.Set(IdempotencyKeyHeader, "order-9")
	resp, err := client.Do(req)
	if err != nil || resp.StatusCode != 200 {
		t.Fatalf("status %v err %v", resp, err)
	}
	if len(s.keys) != 2 || s.keys[0] != "order-9" || s.keys[1] != "order-9" {
		t.Fatalf("keys %v", s.keys)
	}
	if s.bodies[1] != `{"eventId":"e"}` {
		t.Fatalf("body not replayed: %q", s.bodies[1])
	}
	if len(waits) != 1 || waits[0] != time.Second {
		t.Fatalf("waits %v", waits)
	}
}

func TestUnkeyedPostIsNotRetried(t *testing.T) {
	s := &script{statuses: []int{503, 200}}
	srv := httptest.NewServer(http.HandlerFunc(s.handler))
	defer srv.Close()
	var waits []time.Duration
	resp, _ := NewRetryingHTTPClient(nil, testPolicy(&waits)).Post(srv.URL, "application/json", strings.NewReader("{}"))
	if resp.StatusCode != 503 || len(s.keys) != 1 || len(waits) != 0 {
		t.Fatalf("status %d calls %d waits %v", resp.StatusCode, len(s.keys), waits)
	}
}

func TestGetGivesUpAfterMaxAttemptsWithJitteredBackoff(t *testing.T) {
	s := &script{statuses: []int{503, 503, 503, 503, 200}}
	srv := httptest.NewServer(http.HandlerFunc(s.handler))
	defer srv.Close()
	var waits []time.Duration
	resp, _ := NewRetryingHTTPClient(nil, testPolicy(&waits)).Get(srv.URL)
	if resp.StatusCode != 503 || len(s.keys) != 4 {
		t.Fatalf("status %d calls %d", resp.StatusCode, len(s.keys))
	}
	want := []time.Duration{125 * time.Millisecond, 250 * time.Millisecond, 500 * time.Millisecond}
	for i := range want {
		if waits[i] != want[i] {
			t.Fatalf("waits %v", waits)
		}
	}
}

func TestOtherStatusesAndLongRetryAfterReturnImmediately(t *testing.T) {
	s := &script{statuses: []int{409, 429, 200}, headers: []map[string]string{{}, {"Retry-After": "3600"}}}
	srv := httptest.NewServer(http.HandlerFunc(s.handler))
	defer srv.Close()
	var waits []time.Duration
	client := NewRetryingHTTPClient(nil, testPolicy(&waits))
	if resp, _ := client.Get(srv.URL); resp.StatusCode != 409 {
		t.Fatalf("got %d", resp.StatusCode)
	}
	if resp, _ := client.Get(srv.URL); resp.StatusCode != 429 {
		t.Fatalf("got %d", resp.StatusCode)
	}
	if len(waits) != 0 {
		t.Fatalf("waits %v", waits)
	}
}

func TestParseRetryAfterAndDelay(t *testing.T) {
	now := time.Date(2026, 10, 23, 12, 0, 0, 0, time.UTC)
	if d, ok := ParseRetryAfter("2", now); !ok || d != 2*time.Second {
		t.Fatal("seconds")
	}
	if d, ok := ParseRetryAfter("Fri, 23 Oct 2026 12:00:03 GMT", now); !ok || d != 3*time.Second {
		t.Fatalf("date %v", d)
	}
	if _, ok := ParseRetryAfter("soon", now); ok {
		t.Fatal("invalid accepted")
	}
	p := DefaultRetryPolicy()
	p.Random = func() float64 { return 1 }
	if d, _ := p.Delay(30, 0, false); d != 8*time.Second {
		t.Fatalf("cap %v", d)
	}
	if _, ok := p.Delay(1, time.Hour, true); ok {
		t.Fatal("long Retry-After accepted")
	}
	if NewIdempotencyKey() == NewIdempotencyKey() || len(NewIdempotencyKey()) != 36 {
		t.Fatal("keys")
	}
}
