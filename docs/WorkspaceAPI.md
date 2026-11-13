# \WorkspaceAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteApiWorkspacesByWorkspaceId**](WorkspaceAPI.md#DeleteApiWorkspacesByWorkspaceId) | **Delete** /api/workspaces/{workspaceId} | Deactivate a workspace (legacy DELETE)
[**DeleteApiWorkspacesByWorkspaceIdInvitationsByInvitationId**](WorkspaceAPI.md#DeleteApiWorkspacesByWorkspaceIdInvitationsByInvitationId) | **Delete** /api/workspaces/{workspaceId}/invitations/{invitationId} | Revoke a pending invitation of a workspace you administer
[**DeleteApiWorkspacesByWorkspaceIdKeysByKeyId**](WorkspaceAPI.md#DeleteApiWorkspacesByWorkspaceIdKeysByKeyId) | **Delete** /api/workspaces/{workspaceId}/keys/{keyId} | Revoke a key of a workspace you administer
[**DeleteApiWorkspacesByWorkspaceIdMembersByUserId**](WorkspaceAPI.md#DeleteApiWorkspacesByWorkspaceIdMembersByUserId) | **Delete** /api/workspaces/{workspaceId}/members/{userId} | Remove a member from a workspace you administer
[**GetApiWorkspace**](WorkspaceAPI.md#GetApiWorkspace) | **Get** /api/workspace | Get workspace profile for the authenticated tenant
[**GetApiWorkspaceBuyerExtraData**](WorkspaceAPI.md#GetApiWorkspaceBuyerExtraData) | **Get** /api/workspace/buyer-extra-data | Get the buyer extraData publication (opt-in, allow-listed keys)
[**GetApiWorkspaceChartValidation**](WorkspaceAPI.md#GetApiWorkspaceChartValidation) | **Get** /api/workspace/chart-validation | Get chart validation severities
[**GetApiWorkspaceDefaults**](WorkspaceAPI.md#GetApiWorkspaceDefaults) | **Get** /api/workspace/defaults | Get the workspace defaults (renderer settings, draft chart drawings)
[**GetApiWorkspaceIntegration**](WorkspaceAPI.md#GetApiWorkspaceIntegration) | **Get** /api/workspace/integration | Get the workspace integration mode (pricing source)
[**GetApiWorkspaceSettings**](WorkspaceAPI.md#GetApiWorkspaceSettings) | **Get** /api/workspace/settings | Get workspace settings (effective value and source)
[**GetApiWorkspaces**](WorkspaceAPI.md#GetApiWorkspaces) | **Get** /api/workspaces | List my workspaces
[**GetApiWorkspacesByExternalIdByExternalId**](WorkspaceAPI.md#GetApiWorkspacesByExternalIdByExternalId) | **Get** /api/workspaces/by-external-id/{externalId} | Get one of my workspaces by its external id
[**GetApiWorkspacesByWorkspaceId**](WorkspaceAPI.md#GetApiWorkspacesByWorkspaceId) | **Get** /api/workspaces/{workspaceId} | Get one of my workspaces
[**GetApiWorkspacesByWorkspaceIdInvitations**](WorkspaceAPI.md#GetApiWorkspacesByWorkspaceIdInvitations) | **Get** /api/workspaces/{workspaceId}/invitations | List the pending invitations of a workspace you administer
[**GetApiWorkspacesByWorkspaceIdKeys**](WorkspaceAPI.md#GetApiWorkspacesByWorkspaceIdKeys) | **Get** /api/workspaces/{workspaceId}/keys | List the API keys of a workspace you administer
[**GetApiWorkspacesByWorkspaceIdMembers**](WorkspaceAPI.md#GetApiWorkspacesByWorkspaceIdMembers) | **Get** /api/workspaces/{workspaceId}/members | List the members of a workspace you administer
[**GetApiWorkspacesOverview**](WorkspaceAPI.md#GetApiWorkspacesOverview) | **Get** /api/workspaces/overview | Consolidated overview of the organizer workspaces you administer
[**PatchApiWorkspace**](WorkspaceAPI.md#PatchApiWorkspace) | **Patch** /api/workspace | Update workspace display name
[**PatchApiWorkspaceChartValidation**](WorkspaceAPI.md#PatchApiWorkspaceChartValidation) | **Patch** /api/workspace/chart-validation | Update chart validation severities
[**PatchApiWorkspaceDefaults**](WorkspaceAPI.md#PatchApiWorkspaceDefaults) | **Patch** /api/workspace/defaults | Change the workspace defaults (renderer settings, draft chart drawings)
[**PatchApiWorkspaceIntegration**](WorkspaceAPI.md#PatchApiWorkspaceIntegration) | **Patch** /api/workspace/integration | Change the workspace integration mode
[**PatchApiWorkspaceSettings**](WorkspaceAPI.md#PatchApiWorkspaceSettings) | **Patch** /api/workspace/settings | Override workspace settings or return them to the default
[**PatchApiWorkspacesByWorkspaceId**](WorkspaceAPI.md#PatchApiWorkspacesByWorkspaceId) | **Patch** /api/workspaces/{workspaceId} | Rename one of my workspaces or change its external id
[**PatchApiWorkspacesByWorkspaceIdKeysByKeyId**](WorkspaceAPI.md#PatchApiWorkspacesByWorkspaceIdKeysByKeyId) | **Patch** /api/workspaces/{workspaceId}/keys/{keyId} | Rename a key of a workspace you administer or change its scopes
[**PostApiWorkspacePublicKeyRotate**](WorkspaceAPI.md#PostApiWorkspacePublicKeyRotate) | **Post** /api/workspace/public-key/rotate | Rotate the public workspace key
[**PostApiWorkspaces**](WorkspaceAPI.md#PostApiWorkspaces) | **Post** /api/workspaces | Create a workspace (provision an organizer)
[**PostApiWorkspacesByWorkspaceIdActionsActivate**](WorkspaceAPI.md#PostApiWorkspacesByWorkspaceIdActionsActivate) | **Post** /api/workspaces/{workspaceId}/actions/activate | Reactivate a workspace
[**PostApiWorkspacesByWorkspaceIdActionsDeactivate**](WorkspaceAPI.md#PostApiWorkspacesByWorkspaceIdActionsDeactivate) | **Post** /api/workspaces/{workspaceId}/actions/deactivate | Deactivate a workspace
[**PostApiWorkspacesByWorkspaceIdActionsDelete**](WorkspaceAPI.md#PostApiWorkspacesByWorkspaceIdActionsDelete) | **Post** /api/workspaces/{workspaceId}/actions/delete | Delete a workspace
[**PostApiWorkspacesByWorkspaceIdActionsSetDefault**](WorkspaceAPI.md#PostApiWorkspacesByWorkspaceIdActionsSetDefault) | **Post** /api/workspaces/{workspaceId}/actions/set-default | Set my default workspace
[**PostApiWorkspacesByWorkspaceIdInvitations**](WorkspaceAPI.md#PostApiWorkspacesByWorkspaceIdInvitations) | **Post** /api/workspaces/{workspaceId}/invitations | Invite a person to a workspace you administer
[**PostApiWorkspacesByWorkspaceIdKeys**](WorkspaceAPI.md#PostApiWorkspacesByWorkspaceIdKeys) | **Post** /api/workspaces/{workspaceId}/keys | Mint an API key in a workspace you administer
[**PostApiWorkspacesByWorkspaceIdKeysByKeyIdActionsRotate**](WorkspaceAPI.md#PostApiWorkspacesByWorkspaceIdKeysByKeyIdActionsRotate) | **Post** /api/workspaces/{workspaceId}/keys/{keyId}/actions/rotate | Rotate the secret of a key of a workspace you administer
[**PutApiWorkspaceBuyerExtraData**](WorkspaceAPI.md#PutApiWorkspaceBuyerExtraData) | **Put** /api/workspace/buyer-extra-data | Publish selected extraData keys to buyers (or stop)



## DeleteApiWorkspacesByWorkspaceId

> DeleteApiWorkspacesByWorkspaceId200Response DeleteApiWorkspacesByWorkspaceId(ctx, workspaceId).IdempotencyKey(idempotencyKey).Execute()

Deactivate a workspace (legacy DELETE)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/seatschart/seatschart-go"
)

func main() {
	workspaceId := "workspaceId_example" // string | 
	idempotencyKey := "idempotencyKey_example" // string | Required on transactional routes. Reuse the same key for every retry of one logical operation; the API replays the stored outcome. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WorkspaceAPI.DeleteApiWorkspacesByWorkspaceId(context.Background(), workspaceId).IdempotencyKey(idempotencyKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorkspaceAPI.DeleteApiWorkspacesByWorkspaceId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteApiWorkspacesByWorkspaceId`: DeleteApiWorkspacesByWorkspaceId200Response
	fmt.Fprintf(os.Stdout, "Response from `WorkspaceAPI.DeleteApiWorkspacesByWorkspaceId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workspaceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteApiWorkspacesByWorkspaceIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **idempotencyKey** | **string** | Required on transactional routes. Reuse the same key for every retry of one logical operation; the API replays the stored outcome. | 

### Return type

[**DeleteApiWorkspacesByWorkspaceId200Response**](DeleteApiWorkspacesByWorkspaceId200Response.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, multipart/form-data, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteApiWorkspacesByWorkspaceIdInvitationsByInvitationId

> DeleteApiWorkspacesByWorkspaceIdInvitationsByInvitationId(ctx, workspaceId, invitationId).IdempotencyKey(idempotencyKey).Execute()

Revoke a pending invitation of a workspace you administer



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/seatschart/seatschart-go"
)

func main() {
	workspaceId := "workspaceId_example" // string | Workspace id or its external id.
	invitationId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	idempotencyKey := "idempotencyKey_example" // string | Required on transactional routes. Reuse the same key for every retry of one logical operation; the API replays the stored outcome. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.WorkspaceAPI.DeleteApiWorkspacesByWorkspaceIdInvitationsByInvitationId(context.Background(), workspaceId, invitationId).IdempotencyKey(idempotencyKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorkspaceAPI.DeleteApiWorkspacesByWorkspaceIdInvitationsByInvitationId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workspaceId** | **string** | Workspace id or its external id. | 
**invitationId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteApiWorkspacesByWorkspaceIdInvitationsByInvitationIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **idempotencyKey** | **string** | Required on transactional routes. Reuse the same key for every retry of one logical operation; the API replays the stored outcome. | 

### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteApiWorkspacesByWorkspaceIdKeysByKeyId

> DeleteApiWorkspacesByWorkspaceIdKeysByKeyId200Response DeleteApiWorkspacesByWorkspaceIdKeysByKeyId(ctx, workspaceId, keyId).IdempotencyKey(idempotencyKey).Execute()

Revoke a key of a workspace you administer



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/seatschart/seatschart-go"
)

func main() {
	workspaceId := "workspaceId_example" // string | Workspace id or its external id.
	keyId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	idempotencyKey := "idempotencyKey_example" // string | Required on transactional routes. Reuse the same key for every retry of one logical operation; the API replays the stored outcome. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WorkspaceAPI.DeleteApiWorkspacesByWorkspaceIdKeysByKeyId(context.Background(), workspaceId, keyId).IdempotencyKey(idempotencyKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorkspaceAPI.DeleteApiWorkspacesByWorkspaceIdKeysByKeyId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteApiWorkspacesByWorkspaceIdKeysByKeyId`: DeleteApiWorkspacesByWorkspaceIdKeysByKeyId200Response
	fmt.Fprintf(os.Stdout, "Response from `WorkspaceAPI.DeleteApiWorkspacesByWorkspaceIdKeysByKeyId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workspaceId** | **string** | Workspace id or its external id. | 
**keyId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteApiWorkspacesByWorkspaceIdKeysByKeyIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **idempotencyKey** | **string** | Required on transactional routes. Reuse the same key for every retry of one logical operation; the API replays the stored outcome. | 

### Return type

[**DeleteApiWorkspacesByWorkspaceIdKeysByKeyId200Response**](DeleteApiWorkspacesByWorkspaceIdKeysByKeyId200Response.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, multipart/form-data, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteApiWorkspacesByWorkspaceIdMembersByUserId

> DeleteApiWorkspacesByWorkspaceIdMembersByUserId(ctx, workspaceId, userId).IdempotencyKey(idempotencyKey).Execute()

Remove a member from a workspace you administer



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/seatschart/seatschart-go"
)

func main() {
	workspaceId := "workspaceId_example" // string | Workspace id or its external id.
	userId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	idempotencyKey := "idempotencyKey_example" // string | Required on transactional routes. Reuse the same key for every retry of one logical operation; the API replays the stored outcome. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.WorkspaceAPI.DeleteApiWorkspacesByWorkspaceIdMembersByUserId(context.Background(), workspaceId, userId).IdempotencyKey(idempotencyKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorkspaceAPI.DeleteApiWorkspacesByWorkspaceIdMembersByUserId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workspaceId** | **string** | Workspace id or its external id. | 
**userId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteApiWorkspacesByWorkspaceIdMembersByUserIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **idempotencyKey** | **string** | Required on transactional routes. Reuse the same key for every retry of one logical operation; the API replays the stored outcome. | 

### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetApiWorkspace

> GetApiWorkspace(ctx).Execute()

Get workspace profile for the authenticated tenant



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/seatschart/seatschart-go"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.WorkspaceAPI.GetApiWorkspace(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorkspaceAPI.GetApiWorkspace``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetApiWorkspaceRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetApiWorkspaceBuyerExtraData

> GetApiWorkspaceBuyerExtraData200Response GetApiWorkspaceBuyerExtraData(ctx).Execute()

Get the buyer extraData publication (opt-in, allow-listed keys)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/seatschart/seatschart-go"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WorkspaceAPI.GetApiWorkspaceBuyerExtraData(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorkspaceAPI.GetApiWorkspaceBuyerExtraData``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetApiWorkspaceBuyerExtraData`: GetApiWorkspaceBuyerExtraData200Response
	fmt.Fprintf(os.Stdout, "Response from `WorkspaceAPI.GetApiWorkspaceBuyerExtraData`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetApiWorkspaceBuyerExtraDataRequest struct via the builder pattern


### Return type

[**GetApiWorkspaceBuyerExtraData200Response**](GetApiWorkspaceBuyerExtraData200Response.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, multipart/form-data, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetApiWorkspaceChartValidation

> GetApiWorkspaceChartValidation(ctx).Execute()

Get chart validation severities



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/seatschart/seatschart-go"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.WorkspaceAPI.GetApiWorkspaceChartValidation(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorkspaceAPI.GetApiWorkspaceChartValidation``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetApiWorkspaceChartValidationRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetApiWorkspaceDefaults

> GetApiWorkspaceDefaults200Response GetApiWorkspaceDefaults(ctx).Execute()

Get the workspace defaults (renderer settings, draft chart drawings)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/seatschart/seatschart-go"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WorkspaceAPI.GetApiWorkspaceDefaults(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorkspaceAPI.GetApiWorkspaceDefaults``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetApiWorkspaceDefaults`: GetApiWorkspaceDefaults200Response
	fmt.Fprintf(os.Stdout, "Response from `WorkspaceAPI.GetApiWorkspaceDefaults`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetApiWorkspaceDefaultsRequest struct via the builder pattern


### Return type

[**GetApiWorkspaceDefaults200Response**](GetApiWorkspaceDefaults200Response.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, multipart/form-data, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetApiWorkspaceIntegration

> GetApiWorkspaceIntegration200Response GetApiWorkspaceIntegration(ctx).Execute()

Get the workspace integration mode (pricing source)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/seatschart/seatschart-go"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WorkspaceAPI.GetApiWorkspaceIntegration(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorkspaceAPI.GetApiWorkspaceIntegration``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetApiWorkspaceIntegration`: GetApiWorkspaceIntegration200Response
	fmt.Fprintf(os.Stdout, "Response from `WorkspaceAPI.GetApiWorkspaceIntegration`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetApiWorkspaceIntegrationRequest struct via the builder pattern


### Return type

[**GetApiWorkspaceIntegration200Response**](GetApiWorkspaceIntegration200Response.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, multipart/form-data, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetApiWorkspaceSettings

> GetApiWorkspaceSettings200Response GetApiWorkspaceSettings(ctx).Execute()

Get workspace settings (effective value and source)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/seatschart/seatschart-go"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WorkspaceAPI.GetApiWorkspaceSettings(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorkspaceAPI.GetApiWorkspaceSettings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetApiWorkspaceSettings`: GetApiWorkspaceSettings200Response
	fmt.Fprintf(os.Stdout, "Response from `WorkspaceAPI.GetApiWorkspaceSettings`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetApiWorkspaceSettingsRequest struct via the builder pattern


### Return type

[**GetApiWorkspaceSettings200Response**](GetApiWorkspaceSettings200Response.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, multipart/form-data, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetApiWorkspaces

> GetApiWorkspaces200Response GetApiWorkspaces(ctx).Filter(filter).Status(status).ExternalId(externalId).CreatedByApi(createdByApi).Execute()

List my workspaces



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/seatschart/seatschart-go"
)

func main() {
	filter := "filter_example" // string | Case-insensitive substring of the workspace name or id (seats.io `filter`). (optional)
	status := "status_example" // string | `active` (seats.io listActive), `inactive` (deactivated; seats.io listInactive) or `all` (default). (optional)
	externalId := "externalId_example" // string | Only the workspace with this external id (your organizer id; exact match). (optional)
	createdByApi := *openapiclient.NewGetApiWorkspacesCreatedByApiParameter() // GetApiWorkspacesCreatedByApiParameter | `true`: only workspaces provisioned by a workspace admin key; `false`: only the others. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WorkspaceAPI.GetApiWorkspaces(context.Background()).Filter(filter).Status(status).ExternalId(externalId).CreatedByApi(createdByApi).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorkspaceAPI.GetApiWorkspaces``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetApiWorkspaces`: GetApiWorkspaces200Response
	fmt.Fprintf(os.Stdout, "Response from `WorkspaceAPI.GetApiWorkspaces`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetApiWorkspacesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filter** | **string** | Case-insensitive substring of the workspace name or id (seats.io &#x60;filter&#x60;). | 
 **status** | **string** | &#x60;active&#x60; (seats.io listActive), &#x60;inactive&#x60; (deactivated; seats.io listInactive) or &#x60;all&#x60; (default). | 
 **externalId** | **string** | Only the workspace with this external id (your organizer id; exact match). | 
 **createdByApi** | [**GetApiWorkspacesCreatedByApiParameter**](GetApiWorkspacesCreatedByApiParameter.md) | &#x60;true&#x60;: only workspaces provisioned by a workspace admin key; &#x60;false&#x60;: only the others. | 

### Return type

[**GetApiWorkspaces200Response**](GetApiWorkspaces200Response.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, multipart/form-data, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetApiWorkspacesByExternalIdByExternalId

> GetApiWorkspaces200ResponseItemsInner GetApiWorkspacesByExternalIdByExternalId(ctx, externalId).Execute()

Get one of my workspaces by its external id



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/seatschart/seatschart-go"
)

func main() {
	externalId := "externalId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WorkspaceAPI.GetApiWorkspacesByExternalIdByExternalId(context.Background(), externalId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorkspaceAPI.GetApiWorkspacesByExternalIdByExternalId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetApiWorkspacesByExternalIdByExternalId`: GetApiWorkspaces200ResponseItemsInner
	fmt.Fprintf(os.Stdout, "Response from `WorkspaceAPI.GetApiWorkspacesByExternalIdByExternalId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**externalId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetApiWorkspacesByExternalIdByExternalIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetApiWorkspaces200ResponseItemsInner**](GetApiWorkspaces200ResponseItemsInner.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, multipart/form-data, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetApiWorkspacesByWorkspaceId

> GetApiWorkspaces200ResponseItemsInner GetApiWorkspacesByWorkspaceId(ctx, workspaceId).Execute()

Get one of my workspaces



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/seatschart/seatschart-go"
)

func main() {
	workspaceId := "workspaceId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WorkspaceAPI.GetApiWorkspacesByWorkspaceId(context.Background(), workspaceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorkspaceAPI.GetApiWorkspacesByWorkspaceId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetApiWorkspacesByWorkspaceId`: GetApiWorkspaces200ResponseItemsInner
	fmt.Fprintf(os.Stdout, "Response from `WorkspaceAPI.GetApiWorkspacesByWorkspaceId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workspaceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetApiWorkspacesByWorkspaceIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetApiWorkspaces200ResponseItemsInner**](GetApiWorkspaces200ResponseItemsInner.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, multipart/form-data, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetApiWorkspacesByWorkspaceIdInvitations

> GetApiWorkspacesByWorkspaceIdInvitations(ctx, workspaceId).Execute()

List the pending invitations of a workspace you administer

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/seatschart/seatschart-go"
)

func main() {
	workspaceId := "workspaceId_example" // string | Workspace id or its external id.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.WorkspaceAPI.GetApiWorkspacesByWorkspaceIdInvitations(context.Background(), workspaceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorkspaceAPI.GetApiWorkspacesByWorkspaceIdInvitations``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workspaceId** | **string** | Workspace id or its external id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetApiWorkspacesByWorkspaceIdInvitationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetApiWorkspacesByWorkspaceIdKeys

> GetApiWorkspacesByWorkspaceIdKeys200Response GetApiWorkspacesByWorkspaceIdKeys(ctx, workspaceId).Execute()

List the API keys of a workspace you administer



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/seatschart/seatschart-go"
)

func main() {
	workspaceId := "workspaceId_example" // string | Workspace id or its external id.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WorkspaceAPI.GetApiWorkspacesByWorkspaceIdKeys(context.Background(), workspaceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorkspaceAPI.GetApiWorkspacesByWorkspaceIdKeys``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetApiWorkspacesByWorkspaceIdKeys`: GetApiWorkspacesByWorkspaceIdKeys200Response
	fmt.Fprintf(os.Stdout, "Response from `WorkspaceAPI.GetApiWorkspacesByWorkspaceIdKeys`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workspaceId** | **string** | Workspace id or its external id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetApiWorkspacesByWorkspaceIdKeysRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetApiWorkspacesByWorkspaceIdKeys200Response**](GetApiWorkspacesByWorkspaceIdKeys200Response.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, multipart/form-data, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetApiWorkspacesByWorkspaceIdMembers

> GetApiWorkspacesByWorkspaceIdMembers(ctx, workspaceId).Execute()

List the members of a workspace you administer



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/seatschart/seatschart-go"
)

func main() {
	workspaceId := "workspaceId_example" // string | Workspace id or its external id.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.WorkspaceAPI.GetApiWorkspacesByWorkspaceIdMembers(context.Background(), workspaceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorkspaceAPI.GetApiWorkspacesByWorkspaceIdMembers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workspaceId** | **string** | Workspace id or its external id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetApiWorkspacesByWorkspaceIdMembersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetApiWorkspacesOverview

> GetApiWorkspacesOverview(ctx).ExternalId(externalId).Status(status).CreatedByApi(createdByApi).CreatedFrom(createdFrom).CreatedTo(createdTo).Month(month).Limit(limit).Cursor(cursor).Execute()

Consolidated overview of the organizer workspaces you administer



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/seatschart/seatschart-go"
)

func main() {
	externalId := "externalId_example" // string | Only this organizer. (optional)
	status := "status_example" // string | Default `all`. (optional)
	createdByApi := *openapiclient.NewGetApiWorkspacesOverviewCreatedByApiParameter() // GetApiWorkspacesOverviewCreatedByApiParameter | `true`: only workspaces provisioned by a workspace admin key. (optional)
	createdFrom := "createdFrom_example" // string | Workspaces created at or after this date / date-time (ISO 8601). (optional)
	createdTo := "createdTo_example" // string | Workspaces created at or before this date (whole UTC day) / date-time. (optional)
	month := "month_example" // string | Usage month `YYYY-MM` (default: the current UTC month). (optional)
	limit := float32(8.14) // float32 | Page size (1..200, default 50). (optional)
	cursor := "cursor_example" // string | Opaque `nextCursor` of the previous page. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.WorkspaceAPI.GetApiWorkspacesOverview(context.Background()).ExternalId(externalId).Status(status).CreatedByApi(createdByApi).CreatedFrom(createdFrom).CreatedTo(createdTo).Month(month).Limit(limit).Cursor(cursor).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorkspaceAPI.GetApiWorkspacesOverview``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetApiWorkspacesOverviewRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **externalId** | **string** | Only this organizer. | 
 **status** | **string** | Default &#x60;all&#x60;. | 
 **createdByApi** | [**GetApiWorkspacesOverviewCreatedByApiParameter**](GetApiWorkspacesOverviewCreatedByApiParameter.md) | &#x60;true&#x60;: only workspaces provisioned by a workspace admin key. | 
 **createdFrom** | **string** | Workspaces created at or after this date / date-time (ISO 8601). | 
 **createdTo** | **string** | Workspaces created at or before this date (whole UTC day) / date-time. | 
 **month** | **string** | Usage month &#x60;YYYY-MM&#x60; (default: the current UTC month). | 
 **limit** | **float32** | Page size (1..200, default 50). | 
 **cursor** | **string** | Opaque &#x60;nextCursor&#x60; of the previous page. | 

### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchApiWorkspace

> PatchApiWorkspace(ctx).PatchApiWorkspaceRequest(patchApiWorkspaceRequest).IdempotencyKey(idempotencyKey).Execute()

Update workspace display name

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/seatschart/seatschart-go"
)

func main() {
	patchApiWorkspaceRequest := *openapiclient.NewPatchApiWorkspaceRequest("DisplayName_example") // PatchApiWorkspaceRequest | 
	idempotencyKey := "idempotencyKey_example" // string | Required on transactional routes. Reuse the same key for every retry of one logical operation; the API replays the stored outcome. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.WorkspaceAPI.PatchApiWorkspace(context.Background()).PatchApiWorkspaceRequest(patchApiWorkspaceRequest).IdempotencyKey(idempotencyKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorkspaceAPI.PatchApiWorkspace``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPatchApiWorkspaceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **patchApiWorkspaceRequest** | [**PatchApiWorkspaceRequest**](PatchApiWorkspaceRequest.md) |  | 
 **idempotencyKey** | **string** | Required on transactional routes. Reuse the same key for every retry of one logical operation; the API replays the stored outcome. | 

### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json, multipart/form-data, text/plain
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchApiWorkspaceChartValidation

> PatchApiWorkspaceChartValidation(ctx).PatchApiWorkspaceChartValidationRequest(patchApiWorkspaceChartValidationRequest).IdempotencyKey(idempotencyKey).Execute()

Update chart validation severities



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/seatschart/seatschart-go"
)

func main() {
	patchApiWorkspaceChartValidationRequest := *openapiclient.NewPatchApiWorkspaceChartValidationRequest() // PatchApiWorkspaceChartValidationRequest | 
	idempotencyKey := "idempotencyKey_example" // string | Required on transactional routes. Reuse the same key for every retry of one logical operation; the API replays the stored outcome. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.WorkspaceAPI.PatchApiWorkspaceChartValidation(context.Background()).PatchApiWorkspaceChartValidationRequest(patchApiWorkspaceChartValidationRequest).IdempotencyKey(idempotencyKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorkspaceAPI.PatchApiWorkspaceChartValidation``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPatchApiWorkspaceChartValidationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **patchApiWorkspaceChartValidationRequest** | [**PatchApiWorkspaceChartValidationRequest**](PatchApiWorkspaceChartValidationRequest.md) |  | 
 **idempotencyKey** | **string** | Required on transactional routes. Reuse the same key for every retry of one logical operation; the API replays the stored outcome. | 

### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json, multipart/form-data, text/plain
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchApiWorkspaceDefaults

> PatchApiWorkspaceDefaults(ctx).PatchApiWorkspaceDefaultsRequest(patchApiWorkspaceDefaultsRequest).IdempotencyKey(idempotencyKey).Execute()

Change the workspace defaults (renderer settings, draft chart drawings)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/seatschart/seatschart-go"
)

func main() {
	patchApiWorkspaceDefaultsRequest := *openapiclient.NewPatchApiWorkspaceDefaultsRequest() // PatchApiWorkspaceDefaultsRequest | 
	idempotencyKey := "idempotencyKey_example" // string | Required on transactional routes. Reuse the same key for every retry of one logical operation; the API replays the stored outcome. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.WorkspaceAPI.PatchApiWorkspaceDefaults(context.Background()).PatchApiWorkspaceDefaultsRequest(patchApiWorkspaceDefaultsRequest).IdempotencyKey(idempotencyKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorkspaceAPI.PatchApiWorkspaceDefaults``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPatchApiWorkspaceDefaultsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **patchApiWorkspaceDefaultsRequest** | [**PatchApiWorkspaceDefaultsRequest**](PatchApiWorkspaceDefaultsRequest.md) |  | 
 **idempotencyKey** | **string** | Required on transactional routes. Reuse the same key for every retry of one logical operation; the API replays the stored outcome. | 

### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json, multipart/form-data, text/plain
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchApiWorkspaceIntegration

> PatchApiWorkspaceIntegration(ctx).PatchApiWorkspaceIntegrationRequest(patchApiWorkspaceIntegrationRequest).IdempotencyKey(idempotencyKey).Execute()

Change the workspace integration mode



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/seatschart/seatschart-go"
)

func main() {
	patchApiWorkspaceIntegrationRequest := *openapiclient.NewPatchApiWorkspaceIntegrationRequest() // PatchApiWorkspaceIntegrationRequest | 
	idempotencyKey := "idempotencyKey_example" // string | Required on transactional routes. Reuse the same key for every retry of one logical operation; the API replays the stored outcome. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.WorkspaceAPI.PatchApiWorkspaceIntegration(context.Background()).PatchApiWorkspaceIntegrationRequest(patchApiWorkspaceIntegrationRequest).IdempotencyKey(idempotencyKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorkspaceAPI.PatchApiWorkspaceIntegration``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPatchApiWorkspaceIntegrationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **patchApiWorkspaceIntegrationRequest** | [**PatchApiWorkspaceIntegrationRequest**](PatchApiWorkspaceIntegrationRequest.md) |  | 
 **idempotencyKey** | **string** | Required on transactional routes. Reuse the same key for every retry of one logical operation; the API replays the stored outcome. | 

### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json, multipart/form-data, text/plain
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchApiWorkspaceSettings

> PatchApiWorkspaceSettings(ctx).PatchApiWorkspaceSettingsRequest(patchApiWorkspaceSettingsRequest).IdempotencyKey(idempotencyKey).Execute()

Override workspace settings or return them to the default



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/seatschart/seatschart-go"
)

func main() {
	patchApiWorkspaceSettingsRequest := *openapiclient.NewPatchApiWorkspaceSettingsRequest() // PatchApiWorkspaceSettingsRequest | 
	idempotencyKey := "idempotencyKey_example" // string | Required on transactional routes. Reuse the same key for every retry of one logical operation; the API replays the stored outcome. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.WorkspaceAPI.PatchApiWorkspaceSettings(context.Background()).PatchApiWorkspaceSettingsRequest(patchApiWorkspaceSettingsRequest).IdempotencyKey(idempotencyKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorkspaceAPI.PatchApiWorkspaceSettings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPatchApiWorkspaceSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **patchApiWorkspaceSettingsRequest** | [**PatchApiWorkspaceSettingsRequest**](PatchApiWorkspaceSettingsRequest.md) |  | 
 **idempotencyKey** | **string** | Required on transactional routes. Reuse the same key for every retry of one logical operation; the API replays the stored outcome. | 

### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json, multipart/form-data, text/plain
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchApiWorkspacesByWorkspaceId

> GetApiWorkspaces200ResponseItemsInner PatchApiWorkspacesByWorkspaceId(ctx, workspaceId).PatchApiWorkspacesByWorkspaceIdRequest(patchApiWorkspacesByWorkspaceIdRequest).IdempotencyKey(idempotencyKey).Execute()

Rename one of my workspaces or change its external id



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/seatschart/seatschart-go"
)

func main() {
	workspaceId := "workspaceId_example" // string | 
	patchApiWorkspacesByWorkspaceIdRequest := *openapiclient.NewPatchApiWorkspacesByWorkspaceIdRequest() // PatchApiWorkspacesByWorkspaceIdRequest | 
	idempotencyKey := "idempotencyKey_example" // string | Required on transactional routes. Reuse the same key for every retry of one logical operation; the API replays the stored outcome. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WorkspaceAPI.PatchApiWorkspacesByWorkspaceId(context.Background(), workspaceId).PatchApiWorkspacesByWorkspaceIdRequest(patchApiWorkspacesByWorkspaceIdRequest).IdempotencyKey(idempotencyKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorkspaceAPI.PatchApiWorkspacesByWorkspaceId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchApiWorkspacesByWorkspaceId`: GetApiWorkspaces200ResponseItemsInner
	fmt.Fprintf(os.Stdout, "Response from `WorkspaceAPI.PatchApiWorkspacesByWorkspaceId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workspaceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchApiWorkspacesByWorkspaceIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **patchApiWorkspacesByWorkspaceIdRequest** | [**PatchApiWorkspacesByWorkspaceIdRequest**](PatchApiWorkspacesByWorkspaceIdRequest.md) |  | 
 **idempotencyKey** | **string** | Required on transactional routes. Reuse the same key for every retry of one logical operation; the API replays the stored outcome. | 

### Return type

[**GetApiWorkspaces200ResponseItemsInner**](GetApiWorkspaces200ResponseItemsInner.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json, multipart/form-data, text/plain
- **Accept**: application/json, multipart/form-data, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchApiWorkspacesByWorkspaceIdKeysByKeyId

> PatchApiWorkspacesByWorkspaceIdKeysByKeyId200Response PatchApiWorkspacesByWorkspaceIdKeysByKeyId(ctx, workspaceId, keyId).PatchApiWorkspacesByWorkspaceIdKeysByKeyIdRequest(patchApiWorkspacesByWorkspaceIdKeysByKeyIdRequest).IdempotencyKey(idempotencyKey).Execute()

Rename a key of a workspace you administer or change its scopes



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/seatschart/seatschart-go"
)

func main() {
	workspaceId := "workspaceId_example" // string | Workspace id or its external id.
	keyId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	patchApiWorkspacesByWorkspaceIdKeysByKeyIdRequest := *openapiclient.NewPatchApiWorkspacesByWorkspaceIdKeysByKeyIdRequest() // PatchApiWorkspacesByWorkspaceIdKeysByKeyIdRequest | 
	idempotencyKey := "idempotencyKey_example" // string | Required on transactional routes. Reuse the same key for every retry of one logical operation; the API replays the stored outcome. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WorkspaceAPI.PatchApiWorkspacesByWorkspaceIdKeysByKeyId(context.Background(), workspaceId, keyId).PatchApiWorkspacesByWorkspaceIdKeysByKeyIdRequest(patchApiWorkspacesByWorkspaceIdKeysByKeyIdRequest).IdempotencyKey(idempotencyKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorkspaceAPI.PatchApiWorkspacesByWorkspaceIdKeysByKeyId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchApiWorkspacesByWorkspaceIdKeysByKeyId`: PatchApiWorkspacesByWorkspaceIdKeysByKeyId200Response
	fmt.Fprintf(os.Stdout, "Response from `WorkspaceAPI.PatchApiWorkspacesByWorkspaceIdKeysByKeyId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workspaceId** | **string** | Workspace id or its external id. | 
**keyId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchApiWorkspacesByWorkspaceIdKeysByKeyIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **patchApiWorkspacesByWorkspaceIdKeysByKeyIdRequest** | [**PatchApiWorkspacesByWorkspaceIdKeysByKeyIdRequest**](PatchApiWorkspacesByWorkspaceIdKeysByKeyIdRequest.md) |  | 
 **idempotencyKey** | **string** | Required on transactional routes. Reuse the same key for every retry of one logical operation; the API replays the stored outcome. | 

### Return type

[**PatchApiWorkspacesByWorkspaceIdKeysByKeyId200Response**](PatchApiWorkspacesByWorkspaceIdKeysByKeyId200Response.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json, multipart/form-data, text/plain
- **Accept**: application/json, multipart/form-data, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostApiWorkspacePublicKeyRotate

> PostApiWorkspacePublicKeyRotate(ctx).IdempotencyKey(idempotencyKey).Execute()

Rotate the public workspace key



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/seatschart/seatschart-go"
)

func main() {
	idempotencyKey := "idempotencyKey_example" // string | Required on transactional routes. Reuse the same key for every retry of one logical operation; the API replays the stored outcome. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.WorkspaceAPI.PostApiWorkspacePublicKeyRotate(context.Background()).IdempotencyKey(idempotencyKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorkspaceAPI.PostApiWorkspacePublicKeyRotate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostApiWorkspacePublicKeyRotateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **idempotencyKey** | **string** | Required on transactional routes. Reuse the same key for every retry of one logical operation; the API replays the stored outcome. | 

### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostApiWorkspaces

> PostApiWorkspaces200Response PostApiWorkspaces(ctx).PostApiWorkspacesRequest(postApiWorkspacesRequest).IdempotencyKey(idempotencyKey).Execute()

Create a workspace (provision an organizer)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/seatschart/seatschart-go"
)

func main() {
	postApiWorkspacesRequest := *openapiclient.NewPostApiWorkspacesRequest("DisplayName_example") // PostApiWorkspacesRequest | 
	idempotencyKey := "idempotencyKey_example" // string | Required on transactional routes. Reuse the same key for every retry of one logical operation; the API replays the stored outcome. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WorkspaceAPI.PostApiWorkspaces(context.Background()).PostApiWorkspacesRequest(postApiWorkspacesRequest).IdempotencyKey(idempotencyKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorkspaceAPI.PostApiWorkspaces``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostApiWorkspaces`: PostApiWorkspaces200Response
	fmt.Fprintf(os.Stdout, "Response from `WorkspaceAPI.PostApiWorkspaces`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostApiWorkspacesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **postApiWorkspacesRequest** | [**PostApiWorkspacesRequest**](PostApiWorkspacesRequest.md) |  | 
 **idempotencyKey** | **string** | Required on transactional routes. Reuse the same key for every retry of one logical operation; the API replays the stored outcome. | 

### Return type

[**PostApiWorkspaces200Response**](PostApiWorkspaces200Response.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json, multipart/form-data, text/plain
- **Accept**: application/json, multipart/form-data, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostApiWorkspacesByWorkspaceIdActionsActivate

> DeleteApiWorkspacesByWorkspaceId200Response PostApiWorkspacesByWorkspaceIdActionsActivate(ctx, workspaceId).IdempotencyKey(idempotencyKey).Execute()

Reactivate a workspace



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/seatschart/seatschart-go"
)

func main() {
	workspaceId := "workspaceId_example" // string | 
	idempotencyKey := "idempotencyKey_example" // string | Required on transactional routes. Reuse the same key for every retry of one logical operation; the API replays the stored outcome. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WorkspaceAPI.PostApiWorkspacesByWorkspaceIdActionsActivate(context.Background(), workspaceId).IdempotencyKey(idempotencyKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorkspaceAPI.PostApiWorkspacesByWorkspaceIdActionsActivate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostApiWorkspacesByWorkspaceIdActionsActivate`: DeleteApiWorkspacesByWorkspaceId200Response
	fmt.Fprintf(os.Stdout, "Response from `WorkspaceAPI.PostApiWorkspacesByWorkspaceIdActionsActivate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workspaceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostApiWorkspacesByWorkspaceIdActionsActivateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **idempotencyKey** | **string** | Required on transactional routes. Reuse the same key for every retry of one logical operation; the API replays the stored outcome. | 

### Return type

[**DeleteApiWorkspacesByWorkspaceId200Response**](DeleteApiWorkspacesByWorkspaceId200Response.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, multipart/form-data, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostApiWorkspacesByWorkspaceIdActionsDeactivate

> DeleteApiWorkspacesByWorkspaceId200Response PostApiWorkspacesByWorkspaceIdActionsDeactivate(ctx, workspaceId).IdempotencyKey(idempotencyKey).Execute()

Deactivate a workspace



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/seatschart/seatschart-go"
)

func main() {
	workspaceId := "workspaceId_example" // string | 
	idempotencyKey := "idempotencyKey_example" // string | Required on transactional routes. Reuse the same key for every retry of one logical operation; the API replays the stored outcome. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WorkspaceAPI.PostApiWorkspacesByWorkspaceIdActionsDeactivate(context.Background(), workspaceId).IdempotencyKey(idempotencyKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorkspaceAPI.PostApiWorkspacesByWorkspaceIdActionsDeactivate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostApiWorkspacesByWorkspaceIdActionsDeactivate`: DeleteApiWorkspacesByWorkspaceId200Response
	fmt.Fprintf(os.Stdout, "Response from `WorkspaceAPI.PostApiWorkspacesByWorkspaceIdActionsDeactivate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workspaceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostApiWorkspacesByWorkspaceIdActionsDeactivateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **idempotencyKey** | **string** | Required on transactional routes. Reuse the same key for every retry of one logical operation; the API replays the stored outcome. | 

### Return type

[**DeleteApiWorkspacesByWorkspaceId200Response**](DeleteApiWorkspacesByWorkspaceId200Response.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, multipart/form-data, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostApiWorkspacesByWorkspaceIdActionsDelete

> PostApiWorkspacesByWorkspaceIdActionsDelete200Response PostApiWorkspacesByWorkspaceIdActionsDelete(ctx, workspaceId).PostApiWorkspacesByWorkspaceIdActionsDeleteRequest(postApiWorkspacesByWorkspaceIdActionsDeleteRequest).IdempotencyKey(idempotencyKey).Execute()

Delete a workspace



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/seatschart/seatschart-go"
)

func main() {
	workspaceId := "workspaceId_example" // string | 
	postApiWorkspacesByWorkspaceIdActionsDeleteRequest := *openapiclient.NewPostApiWorkspacesByWorkspaceIdActionsDeleteRequest() // PostApiWorkspacesByWorkspaceIdActionsDeleteRequest | 
	idempotencyKey := "idempotencyKey_example" // string | Required on transactional routes. Reuse the same key for every retry of one logical operation; the API replays the stored outcome. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WorkspaceAPI.PostApiWorkspacesByWorkspaceIdActionsDelete(context.Background(), workspaceId).PostApiWorkspacesByWorkspaceIdActionsDeleteRequest(postApiWorkspacesByWorkspaceIdActionsDeleteRequest).IdempotencyKey(idempotencyKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorkspaceAPI.PostApiWorkspacesByWorkspaceIdActionsDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostApiWorkspacesByWorkspaceIdActionsDelete`: PostApiWorkspacesByWorkspaceIdActionsDelete200Response
	fmt.Fprintf(os.Stdout, "Response from `WorkspaceAPI.PostApiWorkspacesByWorkspaceIdActionsDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workspaceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostApiWorkspacesByWorkspaceIdActionsDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **postApiWorkspacesByWorkspaceIdActionsDeleteRequest** | [**PostApiWorkspacesByWorkspaceIdActionsDeleteRequest**](PostApiWorkspacesByWorkspaceIdActionsDeleteRequest.md) |  | 
 **idempotencyKey** | **string** | Required on transactional routes. Reuse the same key for every retry of one logical operation; the API replays the stored outcome. | 

### Return type

[**PostApiWorkspacesByWorkspaceIdActionsDelete200Response**](PostApiWorkspacesByWorkspaceIdActionsDelete200Response.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json, multipart/form-data, text/plain
- **Accept**: application/json, multipart/form-data, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostApiWorkspacesByWorkspaceIdActionsSetDefault

> GetApiWorkspaces200ResponseItemsInner PostApiWorkspacesByWorkspaceIdActionsSetDefault(ctx, workspaceId).IdempotencyKey(idempotencyKey).Execute()

Set my default workspace



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/seatschart/seatschart-go"
)

func main() {
	workspaceId := "workspaceId_example" // string | 
	idempotencyKey := "idempotencyKey_example" // string | Required on transactional routes. Reuse the same key for every retry of one logical operation; the API replays the stored outcome. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WorkspaceAPI.PostApiWorkspacesByWorkspaceIdActionsSetDefault(context.Background(), workspaceId).IdempotencyKey(idempotencyKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorkspaceAPI.PostApiWorkspacesByWorkspaceIdActionsSetDefault``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostApiWorkspacesByWorkspaceIdActionsSetDefault`: GetApiWorkspaces200ResponseItemsInner
	fmt.Fprintf(os.Stdout, "Response from `WorkspaceAPI.PostApiWorkspacesByWorkspaceIdActionsSetDefault`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workspaceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostApiWorkspacesByWorkspaceIdActionsSetDefaultRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **idempotencyKey** | **string** | Required on transactional routes. Reuse the same key for every retry of one logical operation; the API replays the stored outcome. | 

### Return type

[**GetApiWorkspaces200ResponseItemsInner**](GetApiWorkspaces200ResponseItemsInner.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, multipart/form-data, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostApiWorkspacesByWorkspaceIdInvitations

> PostApiWorkspacesByWorkspaceIdInvitations(ctx, workspaceId).PostApiWorkspacesByWorkspaceIdInvitationsRequest(postApiWorkspacesByWorkspaceIdInvitationsRequest).IdempotencyKey(idempotencyKey).Execute()

Invite a person to a workspace you administer



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/seatschart/seatschart-go"
)

func main() {
	workspaceId := "workspaceId_example" // string | Workspace id or its external id.
	postApiWorkspacesByWorkspaceIdInvitationsRequest := *openapiclient.NewPostApiWorkspacesByWorkspaceIdInvitationsRequest("Email_example", "Role_example") // PostApiWorkspacesByWorkspaceIdInvitationsRequest | 
	idempotencyKey := "idempotencyKey_example" // string | Required on transactional routes. Reuse the same key for every retry of one logical operation; the API replays the stored outcome. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.WorkspaceAPI.PostApiWorkspacesByWorkspaceIdInvitations(context.Background(), workspaceId).PostApiWorkspacesByWorkspaceIdInvitationsRequest(postApiWorkspacesByWorkspaceIdInvitationsRequest).IdempotencyKey(idempotencyKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorkspaceAPI.PostApiWorkspacesByWorkspaceIdInvitations``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workspaceId** | **string** | Workspace id or its external id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostApiWorkspacesByWorkspaceIdInvitationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **postApiWorkspacesByWorkspaceIdInvitationsRequest** | [**PostApiWorkspacesByWorkspaceIdInvitationsRequest**](PostApiWorkspacesByWorkspaceIdInvitationsRequest.md) |  | 
 **idempotencyKey** | **string** | Required on transactional routes. Reuse the same key for every retry of one logical operation; the API replays the stored outcome. | 

### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json, multipart/form-data, text/plain
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostApiWorkspacesByWorkspaceIdKeys

> PostApiWorkspacesByWorkspaceIdKeys201Response PostApiWorkspacesByWorkspaceIdKeys(ctx, workspaceId).PostApiWorkspacesByWorkspaceIdKeysRequest(postApiWorkspacesByWorkspaceIdKeysRequest).IdempotencyKey(idempotencyKey).Execute()

Mint an API key in a workspace you administer



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/seatschart/seatschart-go"
)

func main() {
	workspaceId := "workspaceId_example" // string | Workspace id or its external id.
	postApiWorkspacesByWorkspaceIdKeysRequest := *openapiclient.NewPostApiWorkspacesByWorkspaceIdKeysRequest("Name_example") // PostApiWorkspacesByWorkspaceIdKeysRequest | 
	idempotencyKey := "idempotencyKey_example" // string | Required on transactional routes. Reuse the same key for every retry of one logical operation; the API replays the stored outcome. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WorkspaceAPI.PostApiWorkspacesByWorkspaceIdKeys(context.Background(), workspaceId).PostApiWorkspacesByWorkspaceIdKeysRequest(postApiWorkspacesByWorkspaceIdKeysRequest).IdempotencyKey(idempotencyKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorkspaceAPI.PostApiWorkspacesByWorkspaceIdKeys``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostApiWorkspacesByWorkspaceIdKeys`: PostApiWorkspacesByWorkspaceIdKeys201Response
	fmt.Fprintf(os.Stdout, "Response from `WorkspaceAPI.PostApiWorkspacesByWorkspaceIdKeys`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workspaceId** | **string** | Workspace id or its external id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostApiWorkspacesByWorkspaceIdKeysRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **postApiWorkspacesByWorkspaceIdKeysRequest** | [**PostApiWorkspacesByWorkspaceIdKeysRequest**](PostApiWorkspacesByWorkspaceIdKeysRequest.md) |  | 
 **idempotencyKey** | **string** | Required on transactional routes. Reuse the same key for every retry of one logical operation; the API replays the stored outcome. | 

### Return type

[**PostApiWorkspacesByWorkspaceIdKeys201Response**](PostApiWorkspacesByWorkspaceIdKeys201Response.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json, multipart/form-data, text/plain
- **Accept**: application/json, multipart/form-data, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostApiWorkspacesByWorkspaceIdKeysByKeyIdActionsRotate

> PostApiWorkspacesByWorkspaceIdKeys201Response PostApiWorkspacesByWorkspaceIdKeysByKeyIdActionsRotate(ctx, workspaceId, keyId).IdempotencyKey(idempotencyKey).Execute()

Rotate the secret of a key of a workspace you administer



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/seatschart/seatschart-go"
)

func main() {
	workspaceId := "workspaceId_example" // string | Workspace id or its external id.
	keyId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	idempotencyKey := "idempotencyKey_example" // string | Required on transactional routes. Reuse the same key for every retry of one logical operation; the API replays the stored outcome. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WorkspaceAPI.PostApiWorkspacesByWorkspaceIdKeysByKeyIdActionsRotate(context.Background(), workspaceId, keyId).IdempotencyKey(idempotencyKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorkspaceAPI.PostApiWorkspacesByWorkspaceIdKeysByKeyIdActionsRotate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostApiWorkspacesByWorkspaceIdKeysByKeyIdActionsRotate`: PostApiWorkspacesByWorkspaceIdKeys201Response
	fmt.Fprintf(os.Stdout, "Response from `WorkspaceAPI.PostApiWorkspacesByWorkspaceIdKeysByKeyIdActionsRotate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workspaceId** | **string** | Workspace id or its external id. | 
**keyId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostApiWorkspacesByWorkspaceIdKeysByKeyIdActionsRotateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **idempotencyKey** | **string** | Required on transactional routes. Reuse the same key for every retry of one logical operation; the API replays the stored outcome. | 

### Return type

[**PostApiWorkspacesByWorkspaceIdKeys201Response**](PostApiWorkspacesByWorkspaceIdKeys201Response.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, multipart/form-data, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutApiWorkspaceBuyerExtraData

> PutApiWorkspaceBuyerExtraData(ctx).PutApiWorkspaceBuyerExtraDataRequest(putApiWorkspaceBuyerExtraDataRequest).IdempotencyKey(idempotencyKey).Execute()

Publish selected extraData keys to buyers (or stop)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/seatschart/seatschart-go"
)

func main() {
	putApiWorkspaceBuyerExtraDataRequest := *openapiclient.NewPutApiWorkspaceBuyerExtraDataRequest(false, []string{"Keys_example"}) // PutApiWorkspaceBuyerExtraDataRequest | 
	idempotencyKey := "idempotencyKey_example" // string | Required on transactional routes. Reuse the same key for every retry of one logical operation; the API replays the stored outcome. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.WorkspaceAPI.PutApiWorkspaceBuyerExtraData(context.Background()).PutApiWorkspaceBuyerExtraDataRequest(putApiWorkspaceBuyerExtraDataRequest).IdempotencyKey(idempotencyKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorkspaceAPI.PutApiWorkspaceBuyerExtraData``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPutApiWorkspaceBuyerExtraDataRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **putApiWorkspaceBuyerExtraDataRequest** | [**PutApiWorkspaceBuyerExtraDataRequest**](PutApiWorkspaceBuyerExtraDataRequest.md) |  | 
 **idempotencyKey** | **string** | Required on transactional routes. Reuse the same key for every retry of one logical operation; the API replays the stored outcome. | 

### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json, multipart/form-data, text/plain
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

