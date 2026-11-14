package seatschart

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"sync"
	"testing"
)

type seenRequest struct {
	method, path, query, workspace, idempotency, apiKey string
}

// recordingServer answers every call with answer(path, query) (nil: 404 not_found) and records each request.
func recordingServer(t *testing.T, answer func(path, query string) any) (*APIClient, *[]seenRequest) {
	t.Helper()
	var mu sync.Mutex
	seen := &[]seenRequest{}
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		mu.Lock()
		*seen = append(*seen, seenRequest{r.Method, r.URL.Path, r.URL.RawQuery, r.Header.Get(WorkspaceKeyHeader),
			r.Header.Get(IdempotencyKeyHeader), r.Header.Get("x-api-key")})
		mu.Unlock()
		w.Header().Set("Content-Type", "application/json")
		body := answer(r.URL.Path, r.URL.RawQuery)
		if body == nil {
			w.WriteHeader(http.StatusNotFound)
			body = map[string]string{"error": "not_found", "message": "not found"}
		}
		_ = json.NewEncoder(w).Encode(body)
	}))
	t.Cleanup(server.Close)
	cfg := NewConfiguration()
	cfg.Servers = ServerConfigurations{{URL: server.URL}}
	cfg.AddDefaultHeader("x-api-key", "sk_live_platform")
	return NewAPIClient(cfg), seen
}

func overviewItem(id string) map[string]any {
	return map[string]any{"workspaceId": id, "displayName": id, "externalId": nil, "status": "active", "createdByApi": true,
		"createdAt": "2026-11-14T00:00:00.000Z", "pricingSource": "partner", "charts": 0, "events": 0,
		"seats": map[string]int{"sold": 0, "held": 0, "available": 0}, "revenue": map[string]any{"known": false, "cents": nil},
		"usage": map[string]any{"month": "2026-11", "usedObjects": 0, "usedSeats": 0, "projectedSeats": 0}}
}

func overviewPage(next any, ids ...string) map[string]any {
	items := []map[string]any{}
	for _, id := range ids {
		items = append(items, overviewItem(id))
	}
	return map[string]any{"month": "2026-11", "items": items, "nextCursor": next, "seats": nil, "account": nil,
		"totals": map[string]any{"workspaces": 3, "active": 3, "charts": 0, "events": 0,
			"seats": map[string]int{"sold": 0, "held": 0, "available": 0}, "revenue": map[string]any{"cents": 0, "complete": false},
			"usedObjects": 0}}
}

func TestForWorkspaceScopesEveryCallAndLeavesTheOriginal(t *testing.T) {
	client, seen := recordingServer(t, func(string, string) any { return map[string]bool{"ok": true} })
	org, err := client.ForWorkspace("  org-1042 ")
	if err != nil {
		t.Fatal(err)
	}
	ctx := context.Background()
	if _, _, err := org.Workspaces().Members.Remove(ctx, "org-1042", "u1"); err != nil {
		t.Fatal(err)
	}
	if _, _, err := client.Workspaces().Members.Remove(ctx, "org-1042", "u1"); err != nil {
		t.Fatal(err)
	}
	if got := (*seen)[0]; got.workspace != "org-1042" || got.apiKey != "sk_live_platform" {
		t.Fatalf("scoped call: %+v", got)
	}
	if got := (*seen)[1]; got.workspace != "" {
		t.Fatalf("the original client must not send %s: %+v", WorkspaceKeyHeader, got)
	}
	again, _ := org.ForWorkspace("org-7")
	if again.cfg.DefaultHeader[WorkspaceKeyHeader] != "org-7" || org.cfg.DefaultHeader[WorkspaceKeyHeader] != "org-1042" {
		t.Fatal("re-scoping must replace the workspace on a new client only")
	}
	if org.cfg.HTTPClient != client.cfg.HTTPClient {
		t.Fatal("a scoped client shares the HTTP client (and its retry policy)")
	}
	if _, err := client.ForWorkspace("   "); err == nil {
		t.Fatal("an empty workspace must be refused")
	}
}

func TestCreateAlwaysSendsAnIdempotencyKey(t *testing.T) {
	client, seen := recordingServer(t, func(string, string) any { return nil })
	ws := client.Workspaces()
	body := *NewPostApiWorkspacesRequest("Teatro Aurora")
	body.SetExternalId("org-1042")
	_, _, _ = ws.Create(context.Background(), body, "")
	_, _, _ = ws.Create(context.Background(), body, "provision-org-1042")
	if got := (*seen)[0]; got.method != "POST" || got.path != "/api/workspaces" || len(got.idempotency) != 36 {
		t.Fatalf("generated key: %+v", got)
	}
	if got := (*seen)[1].idempotency; got != "provision-org-1042" {
		t.Fatalf("caller key: %q", got)
	}
}

func TestWorkspaceCallsReachTheirRoutes(t *testing.T) {
	client, seen := recordingServer(t, func(string, string) any { return nil })
	ws, ctx := client.Workspaces(), context.Background()
	_, _, _ = ws.Get(ctx, "ws_1")
	_, _, _ = ws.GetByExternalID(ctx, "org-1042")
	_, _, _ = ws.Keys.List(ctx, "org-1042")
	_, _, _ = ws.Keys.Create(ctx, "org-1042", *NewPostApiWorkspacesByWorkspaceIdKeysRequest("Box office"))
	_, _, _ = ws.Keys.Update(ctx, "org-1042", "k1", *NewPatchApiWorkspacesByWorkspaceIdKeysByKeyIdRequest())
	_, _, _ = ws.Keys.Rotate(ctx, "org-1042", "k1")
	_, _, _ = ws.Keys.Revoke(ctx, "org-1042", "k1")
	_, _, _ = ws.Members.List(ctx, "org-1042")
	_, _, _ = ws.Invitations.List(ctx, "org-1042")
	_, _, _ = ws.Invitations.Create(ctx, "org-1042", *NewPostApiWorkspacesByWorkspaceIdInvitationsRequest("a@example.test", "viewer"))
	_, _, _ = ws.Invitations.Revoke(ctx, "org-1042", "i1")
	want := []string{"GET /api/workspaces/ws_1", "GET /api/workspaces/by-external-id/org-1042", "GET /api/workspaces/org-1042/keys",
		"POST /api/workspaces/org-1042/keys", "PATCH /api/workspaces/org-1042/keys/k1",
		"POST /api/workspaces/org-1042/keys/k1/actions/rotate", "DELETE /api/workspaces/org-1042/keys/k1",
		"GET /api/workspaces/org-1042/members", "GET /api/workspaces/org-1042/invitations",
		"POST /api/workspaces/org-1042/invitations", "DELETE /api/workspaces/org-1042/invitations/i1"}
	var got []string
	for _, r := range *seen {
		got = append(got, r.method+" "+r.path)
	}
	if strings.Join(got, "\n") != strings.Join(want, "\n") {
		t.Fatalf("routes:\n%s\nwant:\n%s", strings.Join(got, "\n"), strings.Join(want, "\n"))
	}
}

func TestListAllOverviewWalksEveryPageWithTheFilters(t *testing.T) {
	client, seen := recordingServer(t, func(_ string, query string) any {
		if strings.Contains(query, "cursor=c2") {
			return overviewPage(nil, "ws_3")
		}
		return overviewPage("c2", "ws_1", "ws_2")
	})
	items, err := client.Workspaces().ListAllOverview(context.Background(),
		func(r ApiGetApiWorkspacesOverviewRequest) ApiGetApiWorkspacesOverviewRequest { return r.CreatedByApi(true) })
	if err != nil {
		t.Fatal(err)
	}
	if len(items) != 3 || items[2].GetWorkspaceId() != "ws_3" {
		t.Fatalf("items: %+v", items)
	}
	if len(*seen) != 2 || !strings.Contains((*seen)[0].query, "createdByApi=true") || !strings.Contains((*seen)[1].query, "createdByApi=true") {
		t.Fatalf("pages: %+v", *seen)
	}
}
