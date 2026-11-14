# \SandboxesAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteApiSandboxesBySandboxId**](SandboxesAPI.md#DeleteApiSandboxesBySandboxId) | **Delete** /api/sandboxes/{sandboxId} | Delete a sandbox
[**GetApiSandboxes**](SandboxesAPI.md#GetApiSandboxes) | **Get** /api/sandboxes | List the sandboxes of this workspace
[**GetApiSandboxesBySandboxId**](SandboxesAPI.md#GetApiSandboxesBySandboxId) | **Get** /api/sandboxes/{sandboxId} | Get a sandbox
[**PostApiSandboxes**](SandboxesAPI.md#PostApiSandboxes) | **Post** /api/sandboxes | Create a sandbox
[**PostApiSandboxesBySandboxIdActionsCopyFromProduction**](SandboxesAPI.md#PostApiSandboxesBySandboxIdActionsCopyFromProduction) | **Post** /api/sandboxes/{sandboxId}/actions/copy-from-production | Copy charts (and events) from production into a sandbox
[**PostApiSandboxesBySandboxIdActionsJoin**](SandboxesAPI.md#PostApiSandboxesBySandboxIdActionsJoin) | **Post** /api/sandboxes/{sandboxId}/actions/join | Join a sandbox (dashboard)
[**PostApiSandboxesBySandboxIdActionsReset**](SandboxesAPI.md#PostApiSandboxesBySandboxIdActionsReset) | **Post** /api/sandboxes/{sandboxId}/actions/reset | Reset a sandbox
[**PostApiSandboxesBySandboxIdAdminKey**](SandboxesAPI.md#PostApiSandboxesBySandboxIdAdminKey) | **Post** /api/sandboxes/{sandboxId}/admin-key | Create or rotate the sandbox admin key



## DeleteApiSandboxesBySandboxId

> DeleteApiSandboxesBySandboxId200Response DeleteApiSandboxesBySandboxId(ctx, sandboxId).IdempotencyKey(idempotencyKey).Execute()

Delete a sandbox



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
	sandboxId := "sandboxId_example" // string | 
	idempotencyKey := "idempotencyKey_example" // string | Required on transactional routes. Reuse the same key for every retry of one logical operation; the API replays the stored outcome. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SandboxesAPI.DeleteApiSandboxesBySandboxId(context.Background(), sandboxId).IdempotencyKey(idempotencyKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SandboxesAPI.DeleteApiSandboxesBySandboxId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteApiSandboxesBySandboxId`: DeleteApiSandboxesBySandboxId200Response
	fmt.Fprintf(os.Stdout, "Response from `SandboxesAPI.DeleteApiSandboxesBySandboxId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sandboxId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteApiSandboxesBySandboxIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **idempotencyKey** | **string** | Required on transactional routes. Reuse the same key for every retry of one logical operation; the API replays the stored outcome. | 

### Return type

[**DeleteApiSandboxesBySandboxId200Response**](DeleteApiSandboxesBySandboxId200Response.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, multipart/form-data, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetApiSandboxes

> GetApiSandboxes200Response GetApiSandboxes(ctx).Execute()

List the sandboxes of this workspace



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
	resp, r, err := apiClient.SandboxesAPI.GetApiSandboxes(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SandboxesAPI.GetApiSandboxes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetApiSandboxes`: GetApiSandboxes200Response
	fmt.Fprintf(os.Stdout, "Response from `SandboxesAPI.GetApiSandboxes`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetApiSandboxesRequest struct via the builder pattern


### Return type

[**GetApiSandboxes200Response**](GetApiSandboxes200Response.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, multipart/form-data, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetApiSandboxesBySandboxId

> GetApiSandboxes200ResponseItemsInner GetApiSandboxesBySandboxId(ctx, sandboxId).Execute()

Get a sandbox



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
	sandboxId := "sandboxId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SandboxesAPI.GetApiSandboxesBySandboxId(context.Background(), sandboxId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SandboxesAPI.GetApiSandboxesBySandboxId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetApiSandboxesBySandboxId`: GetApiSandboxes200ResponseItemsInner
	fmt.Fprintf(os.Stdout, "Response from `SandboxesAPI.GetApiSandboxesBySandboxId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sandboxId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetApiSandboxesBySandboxIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetApiSandboxes200ResponseItemsInner**](GetApiSandboxes200ResponseItemsInner.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, multipart/form-data, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostApiSandboxes

> PostApiSandboxes201Response PostApiSandboxes(ctx).PostApiSandboxesRequest(postApiSandboxesRequest).IdempotencyKey(idempotencyKey).Execute()

Create a sandbox



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
	postApiSandboxesRequest := *openapiclient.NewPostApiSandboxesRequest("DisplayName_example") // PostApiSandboxesRequest | 
	idempotencyKey := "idempotencyKey_example" // string | Required on transactional routes. Reuse the same key for every retry of one logical operation; the API replays the stored outcome. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SandboxesAPI.PostApiSandboxes(context.Background()).PostApiSandboxesRequest(postApiSandboxesRequest).IdempotencyKey(idempotencyKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SandboxesAPI.PostApiSandboxes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostApiSandboxes`: PostApiSandboxes201Response
	fmt.Fprintf(os.Stdout, "Response from `SandboxesAPI.PostApiSandboxes`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostApiSandboxesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **postApiSandboxesRequest** | [**PostApiSandboxesRequest**](PostApiSandboxesRequest.md) |  | 
 **idempotencyKey** | **string** | Required on transactional routes. Reuse the same key for every retry of one logical operation; the API replays the stored outcome. | 

### Return type

[**PostApiSandboxes201Response**](PostApiSandboxes201Response.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json, multipart/form-data, text/plain
- **Accept**: application/json, multipart/form-data, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostApiSandboxesBySandboxIdActionsCopyFromProduction

> PostApiSandboxesBySandboxIdActionsCopyFromProduction200Response PostApiSandboxesBySandboxIdActionsCopyFromProduction(ctx, sandboxId).PostApiSandboxesBySandboxIdActionsCopyFromProductionRequest(postApiSandboxesBySandboxIdActionsCopyFromProductionRequest).IdempotencyKey(idempotencyKey).Execute()

Copy charts (and events) from production into a sandbox



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
	sandboxId := "sandboxId_example" // string | 
	postApiSandboxesBySandboxIdActionsCopyFromProductionRequest := *openapiclient.NewPostApiSandboxesBySandboxIdActionsCopyFromProductionRequest() // PostApiSandboxesBySandboxIdActionsCopyFromProductionRequest | 
	idempotencyKey := "idempotencyKey_example" // string | Required on transactional routes. Reuse the same key for every retry of one logical operation; the API replays the stored outcome. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SandboxesAPI.PostApiSandboxesBySandboxIdActionsCopyFromProduction(context.Background(), sandboxId).PostApiSandboxesBySandboxIdActionsCopyFromProductionRequest(postApiSandboxesBySandboxIdActionsCopyFromProductionRequest).IdempotencyKey(idempotencyKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SandboxesAPI.PostApiSandboxesBySandboxIdActionsCopyFromProduction``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostApiSandboxesBySandboxIdActionsCopyFromProduction`: PostApiSandboxesBySandboxIdActionsCopyFromProduction200Response
	fmt.Fprintf(os.Stdout, "Response from `SandboxesAPI.PostApiSandboxesBySandboxIdActionsCopyFromProduction`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sandboxId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostApiSandboxesBySandboxIdActionsCopyFromProductionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **postApiSandboxesBySandboxIdActionsCopyFromProductionRequest** | [**PostApiSandboxesBySandboxIdActionsCopyFromProductionRequest**](PostApiSandboxesBySandboxIdActionsCopyFromProductionRequest.md) |  | 
 **idempotencyKey** | **string** | Required on transactional routes. Reuse the same key for every retry of one logical operation; the API replays the stored outcome. | 

### Return type

[**PostApiSandboxesBySandboxIdActionsCopyFromProduction200Response**](PostApiSandboxesBySandboxIdActionsCopyFromProduction200Response.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json, multipart/form-data, text/plain
- **Accept**: application/json, multipart/form-data, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostApiSandboxesBySandboxIdActionsJoin

> PostApiSandboxesBySandboxIdActionsJoin200Response PostApiSandboxesBySandboxIdActionsJoin(ctx, sandboxId).IdempotencyKey(idempotencyKey).Execute()

Join a sandbox (dashboard)



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
	sandboxId := "sandboxId_example" // string | 
	idempotencyKey := "idempotencyKey_example" // string | Required on transactional routes. Reuse the same key for every retry of one logical operation; the API replays the stored outcome. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SandboxesAPI.PostApiSandboxesBySandboxIdActionsJoin(context.Background(), sandboxId).IdempotencyKey(idempotencyKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SandboxesAPI.PostApiSandboxesBySandboxIdActionsJoin``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostApiSandboxesBySandboxIdActionsJoin`: PostApiSandboxesBySandboxIdActionsJoin200Response
	fmt.Fprintf(os.Stdout, "Response from `SandboxesAPI.PostApiSandboxesBySandboxIdActionsJoin`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sandboxId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostApiSandboxesBySandboxIdActionsJoinRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **idempotencyKey** | **string** | Required on transactional routes. Reuse the same key for every retry of one logical operation; the API replays the stored outcome. | 

### Return type

[**PostApiSandboxesBySandboxIdActionsJoin200Response**](PostApiSandboxesBySandboxIdActionsJoin200Response.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, multipart/form-data, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostApiSandboxesBySandboxIdActionsReset

> PostApiSandboxesBySandboxIdActionsReset200Response PostApiSandboxesBySandboxIdActionsReset(ctx, sandboxId).PostApiSandboxesBySandboxIdActionsResetRequest(postApiSandboxesBySandboxIdActionsResetRequest).IdempotencyKey(idempotencyKey).Execute()

Reset a sandbox



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
	sandboxId := "sandboxId_example" // string | 
	postApiSandboxesBySandboxIdActionsResetRequest := *openapiclient.NewPostApiSandboxesBySandboxIdActionsResetRequest() // PostApiSandboxesBySandboxIdActionsResetRequest | 
	idempotencyKey := "idempotencyKey_example" // string | Required on transactional routes. Reuse the same key for every retry of one logical operation; the API replays the stored outcome. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SandboxesAPI.PostApiSandboxesBySandboxIdActionsReset(context.Background(), sandboxId).PostApiSandboxesBySandboxIdActionsResetRequest(postApiSandboxesBySandboxIdActionsResetRequest).IdempotencyKey(idempotencyKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SandboxesAPI.PostApiSandboxesBySandboxIdActionsReset``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostApiSandboxesBySandboxIdActionsReset`: PostApiSandboxesBySandboxIdActionsReset200Response
	fmt.Fprintf(os.Stdout, "Response from `SandboxesAPI.PostApiSandboxesBySandboxIdActionsReset`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sandboxId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostApiSandboxesBySandboxIdActionsResetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **postApiSandboxesBySandboxIdActionsResetRequest** | [**PostApiSandboxesBySandboxIdActionsResetRequest**](PostApiSandboxesBySandboxIdActionsResetRequest.md) |  | 
 **idempotencyKey** | **string** | Required on transactional routes. Reuse the same key for every retry of one logical operation; the API replays the stored outcome. | 

### Return type

[**PostApiSandboxesBySandboxIdActionsReset200Response**](PostApiSandboxesBySandboxIdActionsReset200Response.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json, multipart/form-data, text/plain
- **Accept**: application/json, multipart/form-data, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostApiSandboxesBySandboxIdAdminKey

> PostApiSandboxesBySandboxIdAdminKey201Response PostApiSandboxesBySandboxIdAdminKey(ctx, sandboxId).PostApiSandboxesBySandboxIdAdminKeyRequest(postApiSandboxesBySandboxIdAdminKeyRequest).IdempotencyKey(idempotencyKey).Execute()

Create or rotate the sandbox admin key



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
	sandboxId := "sandboxId_example" // string | 
	postApiSandboxesBySandboxIdAdminKeyRequest := *openapiclient.NewPostApiSandboxesBySandboxIdAdminKeyRequest() // PostApiSandboxesBySandboxIdAdminKeyRequest | 
	idempotencyKey := "idempotencyKey_example" // string | Required on transactional routes. Reuse the same key for every retry of one logical operation; the API replays the stored outcome. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SandboxesAPI.PostApiSandboxesBySandboxIdAdminKey(context.Background(), sandboxId).PostApiSandboxesBySandboxIdAdminKeyRequest(postApiSandboxesBySandboxIdAdminKeyRequest).IdempotencyKey(idempotencyKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SandboxesAPI.PostApiSandboxesBySandboxIdAdminKey``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostApiSandboxesBySandboxIdAdminKey`: PostApiSandboxesBySandboxIdAdminKey201Response
	fmt.Fprintf(os.Stdout, "Response from `SandboxesAPI.PostApiSandboxesBySandboxIdAdminKey`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sandboxId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostApiSandboxesBySandboxIdAdminKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **postApiSandboxesBySandboxIdAdminKeyRequest** | [**PostApiSandboxesBySandboxIdAdminKeyRequest**](PostApiSandboxesBySandboxIdAdminKeyRequest.md) |  | 
 **idempotencyKey** | **string** | Required on transactional routes. Reuse the same key for every retry of one logical operation; the API replays the stored outcome. | 

### Return type

[**PostApiSandboxesBySandboxIdAdminKey201Response**](PostApiSandboxesBySandboxIdAdminKey201Response.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json, multipart/form-data, text/plain
- **Accept**: application/json, multipart/form-data, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

