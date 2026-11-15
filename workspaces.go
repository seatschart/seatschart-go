// One workspace per organizer (ticketing platforms) for the SeatsChart Go client.
// Copied into the generated package by scripts/generate-server-sdks.sh. It uses the generated types, so it lives
// in this "_generated" directory, which the go tool ignores in the standalone runtime module (sdks/runtime/go); its
// tests run inside the generated package (scripts/publish-server-sdks.mjs).
//
// A platform operates every organizer workspace it administers with one workspace admin key:
//
//	org, err := client.ForWorkspace("org-1042") // every call sends X-Workspace-Key: org-1042
//	ws := client.Workspaces()
//	created, _, err := ws.Create(ctx, body, "provision-org-1042")
//	found, _, err := ws.GetByExternalID(ctx, "org-1042")
//	items, err := ws.ListAllOverview(ctx, func(r ApiGetApiWorkspacesOverviewRequest) ApiGetApiWorkspacesOverviewRequest {
//		return r.CreatedByApi(true)
//	})
//
// Every workspaceID argument also takes the external id you gave the workspace.
package seatschart

import (
	"context"
	"errors"
	"net/http"
	"strings"
)

// WorkspaceKeyHeader names the workspace a workspace admin key acts in (its id or external id).
const WorkspaceKeyHeader = "X-Workspace-Key"

// ForWorkspace returns a client scoped to one workspace your workspace admin key administers: every call carries
// X-Workspace-Key and acts exactly as that workspace's own full-access key would (its rate limits, usage and audit
// trail). The receiver is unchanged; both clients share the HTTP client (and its retry policy). Ordinary keys get
// 403 ERR_WORKSPACE_HEADER_NOT_ALLOWED; a workspace of another account answers 404.
func (c *APIClient) ForWorkspace(idOrExternalID string) (*APIClient, error) {
	workspace := strings.TrimSpace(idOrExternalID)
	if workspace == "" {
		return nil, errors.New("seatschart: ForWorkspace needs a workspace id or external id")
	}
	cfg := *c.cfg
	cfg.DefaultHeader = make(map[string]string, len(c.cfg.DefaultHeader)+1)
	for name, value := range c.cfg.DefaultHeader {
		cfg.DefaultHeader[name] = value
	}
	cfg.DefaultHeader[WorkspaceKeyHeader] = workspace
	return NewAPIClient(&cfg), nil
}

// Workspaces names the organizer workspace calls like the TypeScript SDK (workspaces.create, getByExternalId, keys,
// members, invitations, overview). The generated WorkspaceAPI stays available for every other workspace route.
type Workspaces struct {
	api         *WorkspaceAPIService
	Keys        WorkspaceKeys
	Members     WorkspaceMembers
	Invitations WorkspaceInvitations
}

// Workspaces returns the organizer workspace calls of this client.
func (c *APIClient) Workspaces() *Workspaces {
	return &Workspaces{api: c.WorkspaceAPI, Keys: WorkspaceKeys{c.WorkspaceAPI}, Members: WorkspaceMembers{c.WorkspaceAPI},
		Invitations: WorkspaceInvitations{c.WorkspaceAPI}}
}

// Create provisions a workspace (POST /api/workspaces): DisplayName and, for an organizer, ExternalId,
// TemplateWorkspaceId and IssueKey. It is always sent with an Idempotency-Key (yours, or a fresh one the retries
// reuse): a replay returns the same workspace with IdempotentReplay true and every secret null.
func (w *Workspaces) Create(ctx context.Context, body PostApiWorkspacesRequest, idempotencyKey string) (*PostApiWorkspaces200Response, *http.Response, error) {
	if strings.TrimSpace(idempotencyKey) == "" {
		idempotencyKey = NewIdempotencyKey()
	}
	return w.api.PostApiWorkspaces(ctx).PostApiWorkspacesRequest(body).IdempotencyKey(idempotencyKey).Execute()
}

// Get returns one of your workspaces.
func (w *Workspaces) Get(ctx context.Context, workspaceID string) (*GetApiWorkspaces200ResponseItemsInner, *http.Response, error) {
	return w.api.GetApiWorkspacesByWorkspaceId(ctx, workspaceID).Execute()
}

// GetByExternalID returns the workspace with this external id among the workspaces you belong to (404 otherwise).
func (w *Workspaces) GetByExternalID(ctx context.Context, externalID string) (*GetApiWorkspaces200ResponseItemsInner, *http.Response, error) {
	return w.api.GetApiWorkspacesByExternalIdByExternalId(ctx, externalID).Execute()
}

// Overview is the consolidated overview request (per workspace and in total); chain its filters (ExternalId, Status,
// CreatedByApi, CreatedFrom, CreatedTo, Month, Limit, Cursor) and Execute it.
func (w *Workspaces) Overview(ctx context.Context) ApiGetApiWorkspacesOverviewRequest {
	return w.api.GetApiWorkspacesOverview(ctx)
}

// ListAllOverview returns every overview item matching the filters set by filter (nil: none), page after page.
func (w *Workspaces) ListAllOverview(ctx context.Context,
	filter func(ApiGetApiWorkspacesOverviewRequest) ApiGetApiWorkspacesOverviewRequest) ([]GetApiWorkspacesOverview200ResponseItemsInner, error) {
	var items []GetApiWorkspacesOverview200ResponseItemsInner
	cursor := ""
	for {
		request := w.Overview(ctx)
		if filter != nil {
			request = filter(request)
		}
		if cursor != "" {
			request = request.Cursor(cursor)
		}
		page, _, err := request.Execute()
		if err != nil {
			return items, err
		}
		items = append(items, page.GetItems()...)
		cursor = page.GetNextCursor()
		if cursor == "" {
			return items, nil
		}
	}
}

// WorkspaceKeys manages the API keys of a workspace you administer (workspace admin keys are read only here).
type WorkspaceKeys struct{ api *WorkspaceAPIService }

// List returns the live keys (prefix only).
func (k WorkspaceKeys) List(ctx context.Context, workspaceID string) (*GetApiWorkspacesByWorkspaceIdKeys200Response, *http.Response, error) {
	return k.api.GetApiWorkspacesByWorkspaceIdKeys(ctx, workspaceID).Execute()
}

// Create mints an ordinary key (full access by default, or Scopes); its Secret is returned once.
func (k WorkspaceKeys) Create(ctx context.Context, workspaceID string, body PostApiWorkspacesByWorkspaceIdKeysRequest) (*PostApiWorkspacesByWorkspaceIdKeys201Response, *http.Response, error) {
	return k.api.PostApiWorkspacesByWorkspaceIdKeys(ctx, workspaceID).PostApiWorkspacesByWorkspaceIdKeysRequest(body).Execute()
}

// Update renames or rescopes a key (next request).
func (k WorkspaceKeys) Update(ctx context.Context, workspaceID, keyID string, body PatchApiWorkspacesByWorkspaceIdKeysByKeyIdRequest) (*PatchApiWorkspacesByWorkspaceIdKeysByKeyId200Response, *http.Response, error) {
	return k.api.PatchApiWorkspacesByWorkspaceIdKeysByKeyId(ctx, workspaceID, keyID).PatchApiWorkspacesByWorkspaceIdKeysByKeyIdRequest(body).Execute()
}

// Rotate issues a new secret for the same key (returned once); the old one stops at once.
func (k WorkspaceKeys) Rotate(ctx context.Context, workspaceID, keyID string) (*PostApiWorkspacesByWorkspaceIdKeys201Response, *http.Response, error) {
	return k.api.PostApiWorkspacesByWorkspaceIdKeysByKeyIdActionsRotate(ctx, workspaceID, keyID).Execute()
}

// Revoke stops the key at once (AlreadyRevoked on a repeat).
func (k WorkspaceKeys) Revoke(ctx context.Context, workspaceID, keyID string) (*DeleteApiWorkspacesByWorkspaceIdKeysByKeyId200Response, *http.Response, error) {
	return k.api.DeleteApiWorkspacesByWorkspaceIdKeysByKeyId(ctx, workspaceID, keyID).Execute()
}

// WorkspaceMembers lists and removes the dashboard members of a workspace you administer.
type WorkspaceMembers struct{ api *WorkspaceAPIService }

// List returns the members with their role.
func (m WorkspaceMembers) List(ctx context.Context, workspaceID string) (*GetApiWorkspacesByWorkspaceIdMembers200Response, *http.Response, error) {
	return m.api.GetApiWorkspacesByWorkspaceIdMembers(ctx, workspaceID).Execute()
}

// Remove removes a member (not the last admin, not yourself).
func (m WorkspaceMembers) Remove(ctx context.Context, workspaceID, userID string) (*DeleteApiWorkspacesByWorkspaceIdMembersByUserId200Response, *http.Response, error) {
	return m.api.DeleteApiWorkspacesByWorkspaceIdMembersByUserId(ctx, workspaceID, userID).Execute()
}

// WorkspaceInvitations invites an organizer's team to the dashboard of its workspace.
type WorkspaceInvitations struct{ api *WorkspaceAPIService }

// List returns the pending invitations.
func (i WorkspaceInvitations) List(ctx context.Context, workspaceID string) (*GetApiWorkspacesByWorkspaceIdInvitations200Response, *http.Response, error) {
	return i.api.GetApiWorkspacesByWorkspaceIdInvitations(ctx, workspaceID).Execute()
}

// Create invites Email with Role (admin, map_editor or viewer); InviteToken is returned once.
func (i WorkspaceInvitations) Create(ctx context.Context, workspaceID string, body PostApiWorkspacesByWorkspaceIdInvitationsRequest) (*PostApiWorkspacesByWorkspaceIdInvitations201Response, *http.Response, error) {
	return i.api.PostApiWorkspacesByWorkspaceIdInvitations(ctx, workspaceID).PostApiWorkspacesByWorkspaceIdInvitationsRequest(body).Execute()
}

// Revoke deletes a pending invitation.
func (i WorkspaceInvitations) Revoke(ctx context.Context, workspaceID, invitationID string) (*DeleteApiWorkspacesByWorkspaceIdMembersByUserId200Response, *http.Response, error) {
	return i.api.DeleteApiWorkspacesByWorkspaceIdInvitationsByInvitationId(ctx, workspaceID, invitationID).Execute()
}
