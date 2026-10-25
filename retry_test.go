package seatschart

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"sync"
	"testing"
	"time"
)

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
