# \WorkspaceAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteApiWorkspacesByWorkspaceId**](WorkspaceAPI.md#DeleteApiWorkspacesByWorkspaceId) | **Delete** /api/workspaces/{workspaceId} | Delete (soft) a workspace
[**GetApiWorkspace**](WorkspaceAPI.md#GetApiWorkspace) | **Get** /api/workspace | Get workspace profile for the authenticated tenant
[**GetApiWorkspaceChartValidation**](WorkspaceAPI.md#GetApiWorkspaceChartValidation) | **Get** /api/workspace/chart-validation | Get chart validation severities
[**GetApiWorkspaceIntegration**](WorkspaceAPI.md#GetApiWorkspaceIntegration) | **Get** /api/workspace/integration | Get the workspace integration mode (pricing source)
[**GetApiWorkspaceSettings**](WorkspaceAPI.md#GetApiWorkspaceSettings) | **Get** /api/workspace/settings | Get workspace settings (effective value and source)
[**GetApiWorkspaces**](WorkspaceAPI.md#GetApiWorkspaces) | **Get** /api/workspaces | List my workspaces
[**GetApiWorkspacesByWorkspaceId**](WorkspaceAPI.md#GetApiWorkspacesByWorkspaceId) | **Get** /api/workspaces/{workspaceId} | Get one of my workspaces
[**PatchApiWorkspace**](WorkspaceAPI.md#PatchApiWorkspace) | **Patch** /api/workspace | Update workspace display name
[**PatchApiWorkspaceChartValidation**](WorkspaceAPI.md#PatchApiWorkspaceChartValidation) | **Patch** /api/workspace/chart-validation | Update chart validation severities
[**PatchApiWorkspaceIntegration**](WorkspaceAPI.md#PatchApiWorkspaceIntegration) | **Patch** /api/workspace/integration | Change the workspace integration mode
[**PatchApiWorkspaceSettings**](WorkspaceAPI.md#PatchApiWorkspaceSettings) | **Patch** /api/workspace/settings | Override workspace settings or return them to the default
[**PatchApiWorkspacesByWorkspaceId**](WorkspaceAPI.md#PatchApiWorkspacesByWorkspaceId) | **Patch** /api/workspaces/{workspaceId} | Rename one of my workspaces
[**PostApiWorkspacePublicKeyRotate**](WorkspaceAPI.md#PostApiWorkspacePublicKeyRotate) | **Post** /api/workspace/public-key/rotate | Rotate the public workspace key
[**PostApiWorkspaces**](WorkspaceAPI.md#PostApiWorkspaces) | **Post** /api/workspaces | Create an additional workspace
[**PostApiWorkspacesByWorkspaceIdActionsActivate**](WorkspaceAPI.md#PostApiWorkspacesByWorkspaceIdActionsActivate) | **Post** /api/workspaces/{workspaceId}/actions/activate | Reactivate a workspace
[**PostApiWorkspacesByWorkspaceIdActionsDeactivate**](WorkspaceAPI.md#PostApiWorkspacesByWorkspaceIdActionsDeactivate) | **Post** /api/workspaces/{workspaceId}/actions/deactivate | Deactivate a workspace
[**PostApiWorkspacesByWorkspaceIdActionsSetDefault**](WorkspaceAPI.md#PostApiWorkspacesByWorkspaceIdActionsSetDefault) | **Post** /api/workspaces/{workspaceId}/actions/set-default | Set my default workspace



## DeleteApiWorkspacesByWorkspaceId

> DeleteApiWorkspacesByWorkspaceId200Response DeleteApiWorkspacesByWorkspaceId(ctx, workspaceId).IdempotencyKey(idempotencyKey).Execute()

Delete (soft) a workspace



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

> GetApiWorkspaces200Response GetApiWorkspaces(ctx).Execute()

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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WorkspaceAPI.GetApiWorkspaces(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorkspaceAPI.GetApiWorkspaces``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetApiWorkspaces`: GetApiWorkspaces200Response
	fmt.Fprintf(os.Stdout, "Response from `WorkspaceAPI.GetApiWorkspaces`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetApiWorkspacesRequest struct via the builder pattern


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

> GetApiWorkspaces200ResponseItemsInner PatchApiWorkspacesByWorkspaceId(ctx, workspaceId).PatchApiWorkspaceRequest(patchApiWorkspaceRequest).IdempotencyKey(idempotencyKey).Execute()

Rename one of my workspaces



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
	patchApiWorkspaceRequest := *openapiclient.NewPatchApiWorkspaceRequest("DisplayName_example") // PatchApiWorkspaceRequest | 
	idempotencyKey := "idempotencyKey_example" // string | Required on transactional routes. Reuse the same key for every retry of one logical operation; the API replays the stored outcome. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WorkspaceAPI.PatchApiWorkspacesByWorkspaceId(context.Background(), workspaceId).PatchApiWorkspaceRequest(patchApiWorkspaceRequest).IdempotencyKey(idempotencyKey).Execute()
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

 **patchApiWorkspaceRequest** | [**PatchApiWorkspaceRequest**](PatchApiWorkspaceRequest.md) |  | 
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

> PostApiWorkspaces(ctx).PostApiWorkspacesRequest(postApiWorkspacesRequest).IdempotencyKey(idempotencyKey).Execute()

Create an additional workspace



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
	r, err := apiClient.WorkspaceAPI.PostApiWorkspaces(context.Background()).PostApiWorkspacesRequest(postApiWorkspacesRequest).IdempotencyKey(idempotencyKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorkspaceAPI.PostApiWorkspaces``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
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

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json, multipart/form-data, text/plain
- **Accept**: Not defined

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

