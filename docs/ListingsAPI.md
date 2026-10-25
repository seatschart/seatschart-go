# \ListingsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteApiListingTypesByKey**](ListingsAPI.md#DeleteApiListingTypesByKey) | **Delete** /api/listing-types/{key} | Delete a listing type
[**GetApiEventsByEventIdListings**](ListingsAPI.md#GetApiEventsByEventIdListings) | **Get** /api/events/{eventId}/listings | List the listings of an event
[**GetApiEventsByEventIdListingsByListingId**](ListingsAPI.md#GetApiEventsByEventIdListingsByListingId) | **Get** /api/events/{eventId}/listings/{listingId} | Read a listing
[**GetApiListingTypes**](ListingsAPI.md#GetApiListingTypes) | **Get** /api/listing-types | List the listing types of the workspace
[**PatchApiEventsByEventIdListingsByListingId**](ListingsAPI.md#PatchApiEventsByEventIdListingsByListingId) | **Patch** /api/events/{eventId}/listings/{listingId} | Update an active listing
[**PatchApiListingTypesByKey**](ListingsAPI.md#PatchApiListingTypesByKey) | **Patch** /api/listing-types/{key} | Update a listing type
[**PostApiEventsByEventIdListings**](ListingsAPI.md#PostApiEventsByEventIdListings) | **Post** /api/events/{eventId}/listings | Create a listing
[**PostApiEventsByEventIdListingsByListingIdActionsMarkSold**](ListingsAPI.md#PostApiEventsByEventIdListingsByListingIdActionsMarkSold) | **Post** /api/events/{eventId}/listings/{listingId}/actions/mark-sold | Mark a listing as sold
[**PostApiEventsByEventIdListingsByListingIdActionsWithdraw**](ListingsAPI.md#PostApiEventsByEventIdListingsByListingIdActionsWithdraw) | **Post** /api/events/{eventId}/listings/{listingId}/actions/withdraw | Withdraw a listing
[**PostApiListingTypes**](ListingsAPI.md#PostApiListingTypes) | **Post** /api/listing-types | Create a listing type



## DeleteApiListingTypesByKey

> DeleteApiListingTypesByKey(ctx, key).IdempotencyKey(idempotencyKey).Execute()

Delete a listing type



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
	key := "key_example" // string | 
	idempotencyKey := "idempotencyKey_example" // string | Required on transactional routes. Reuse the same key for every retry of one logical operation; the API replays the stored outcome. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ListingsAPI.DeleteApiListingTypesByKey(context.Background(), key).IdempotencyKey(idempotencyKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ListingsAPI.DeleteApiListingTypesByKey``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**key** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteApiListingTypesByKeyRequest struct via the builder pattern


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


## GetApiEventsByEventIdListings

> GetApiEventsByEventIdListings(ctx, eventId).Status(status).Type_(type_).Limit(limit).Cursor(cursor).Execute()

List the listings of an event



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
	eventId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	status := "status_example" // string |  (optional)
	type_ := "type__example" // string |  (optional)
	limit := int32(56) // int32 | Page size (1..100, default 20). (optional)
	cursor := "cursor_example" // string | Opaque `nextCursor` value from the previous page. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ListingsAPI.GetApiEventsByEventIdListings(context.Background(), eventId).Status(status).Type_(type_).Limit(limit).Cursor(cursor).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ListingsAPI.GetApiEventsByEventIdListings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**eventId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetApiEventsByEventIdListingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **status** | **string** |  | 
 **type_** | **string** |  | 
 **limit** | **int32** | Page size (1..100, default 20). | 
 **cursor** | **string** | Opaque &#x60;nextCursor&#x60; value from the previous page. | 

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


## GetApiEventsByEventIdListingsByListingId

> GetApiEventsByEventIdListingsByListingId(ctx, eventId, listingId).Execute()

Read a listing



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
	eventId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	listingId := "listingId_example" // string | Listing id or key.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ListingsAPI.GetApiEventsByEventIdListingsByListingId(context.Background(), eventId, listingId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ListingsAPI.GetApiEventsByEventIdListingsByListingId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**eventId** | **string** |  | 
**listingId** | **string** | Listing id or key. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetApiEventsByEventIdListingsByListingIdRequest struct via the builder pattern


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


## GetApiListingTypes

> GetApiListingTypes(ctx).Execute()

List the listing types of the workspace



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
	r, err := apiClient.ListingsAPI.GetApiListingTypes(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ListingsAPI.GetApiListingTypes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetApiListingTypesRequest struct via the builder pattern


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


## PatchApiEventsByEventIdListingsByListingId

> PatchApiEventsByEventIdListingsByListingId(ctx, eventId, listingId).PatchApiEventsByEventIdListingsByListingIdRequest(patchApiEventsByEventIdListingsByListingIdRequest).IdempotencyKey(idempotencyKey).Execute()

Update an active listing



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
	eventId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	listingId := "listingId_example" // string | Listing id or key.
	patchApiEventsByEventIdListingsByListingIdRequest := *openapiclient.NewPatchApiEventsByEventIdListingsByListingIdRequest() // PatchApiEventsByEventIdListingsByListingIdRequest | 
	idempotencyKey := "idempotencyKey_example" // string | Required on transactional routes. Reuse the same key for every retry of one logical operation; the API replays the stored outcome. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ListingsAPI.PatchApiEventsByEventIdListingsByListingId(context.Background(), eventId, listingId).PatchApiEventsByEventIdListingsByListingIdRequest(patchApiEventsByEventIdListingsByListingIdRequest).IdempotencyKey(idempotencyKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ListingsAPI.PatchApiEventsByEventIdListingsByListingId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**eventId** | **string** |  | 
**listingId** | **string** | Listing id or key. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchApiEventsByEventIdListingsByListingIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **patchApiEventsByEventIdListingsByListingIdRequest** | [**PatchApiEventsByEventIdListingsByListingIdRequest**](PatchApiEventsByEventIdListingsByListingIdRequest.md) |  | 
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


## PatchApiListingTypesByKey

> PatchApiListingTypesByKey(ctx, key).PatchApiListingTypesByKeyRequest(patchApiListingTypesByKeyRequest).IdempotencyKey(idempotencyKey).Execute()

Update a listing type



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
	key := "key_example" // string | 
	patchApiListingTypesByKeyRequest := *openapiclient.NewPatchApiListingTypesByKeyRequest() // PatchApiListingTypesByKeyRequest | 
	idempotencyKey := "idempotencyKey_example" // string | Required on transactional routes. Reuse the same key for every retry of one logical operation; the API replays the stored outcome. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ListingsAPI.PatchApiListingTypesByKey(context.Background(), key).PatchApiListingTypesByKeyRequest(patchApiListingTypesByKeyRequest).IdempotencyKey(idempotencyKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ListingsAPI.PatchApiListingTypesByKey``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**key** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchApiListingTypesByKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **patchApiListingTypesByKeyRequest** | [**PatchApiListingTypesByKeyRequest**](PatchApiListingTypesByKeyRequest.md) |  | 
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


## PostApiEventsByEventIdListings

> PostApiEventsByEventIdListings(ctx, eventId).PostApiEventsByEventIdListingsRequest(postApiEventsByEventIdListingsRequest).IdempotencyKey(idempotencyKey).Execute()

Create a listing



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
	eventId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	postApiEventsByEventIdListingsRequest := *openapiclient.NewPostApiEventsByEventIdListingsRequest([]openapiclient.PostApiEventsByEventIdListingsRequestObjectsInner{*openapiclient.NewPostApiEventsByEventIdListingsRequestObjectsInner("ObjectId_example")}) // PostApiEventsByEventIdListingsRequest | 
	idempotencyKey := "idempotencyKey_example" // string | Required on transactional routes. Reuse the same key for every retry of one logical operation; the API replays the stored outcome. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ListingsAPI.PostApiEventsByEventIdListings(context.Background(), eventId).PostApiEventsByEventIdListingsRequest(postApiEventsByEventIdListingsRequest).IdempotencyKey(idempotencyKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ListingsAPI.PostApiEventsByEventIdListings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**eventId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostApiEventsByEventIdListingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **postApiEventsByEventIdListingsRequest** | [**PostApiEventsByEventIdListingsRequest**](PostApiEventsByEventIdListingsRequest.md) |  | 
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


## PostApiEventsByEventIdListingsByListingIdActionsMarkSold

> PostApiEventsByEventIdListingsByListingIdActionsMarkSold(ctx, eventId, listingId).PostApiEventsByEventIdListingsByListingIdActionsMarkSoldRequest(postApiEventsByEventIdListingsByListingIdActionsMarkSoldRequest).IdempotencyKey(idempotencyKey).Execute()

Mark a listing as sold



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
	eventId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	listingId := "listingId_example" // string | Listing id or key.
	postApiEventsByEventIdListingsByListingIdActionsMarkSoldRequest := *openapiclient.NewPostApiEventsByEventIdListingsByListingIdActionsMarkSoldRequest() // PostApiEventsByEventIdListingsByListingIdActionsMarkSoldRequest | 
	idempotencyKey := "idempotencyKey_example" // string | Required on transactional routes. Reuse the same key for every retry of one logical operation; the API replays the stored outcome. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ListingsAPI.PostApiEventsByEventIdListingsByListingIdActionsMarkSold(context.Background(), eventId, listingId).PostApiEventsByEventIdListingsByListingIdActionsMarkSoldRequest(postApiEventsByEventIdListingsByListingIdActionsMarkSoldRequest).IdempotencyKey(idempotencyKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ListingsAPI.PostApiEventsByEventIdListingsByListingIdActionsMarkSold``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**eventId** | **string** |  | 
**listingId** | **string** | Listing id or key. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostApiEventsByEventIdListingsByListingIdActionsMarkSoldRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **postApiEventsByEventIdListingsByListingIdActionsMarkSoldRequest** | [**PostApiEventsByEventIdListingsByListingIdActionsMarkSoldRequest**](PostApiEventsByEventIdListingsByListingIdActionsMarkSoldRequest.md) |  | 
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


## PostApiEventsByEventIdListingsByListingIdActionsWithdraw

> PostApiEventsByEventIdListingsByListingIdActionsWithdraw(ctx, eventId, listingId).IdempotencyKey(idempotencyKey).Execute()

Withdraw a listing



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
	eventId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	listingId := "listingId_example" // string | Listing id or key.
	idempotencyKey := "idempotencyKey_example" // string | Required on transactional routes. Reuse the same key for every retry of one logical operation; the API replays the stored outcome. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ListingsAPI.PostApiEventsByEventIdListingsByListingIdActionsWithdraw(context.Background(), eventId, listingId).IdempotencyKey(idempotencyKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ListingsAPI.PostApiEventsByEventIdListingsByListingIdActionsWithdraw``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**eventId** | **string** |  | 
**listingId** | **string** | Listing id or key. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostApiEventsByEventIdListingsByListingIdActionsWithdrawRequest struct via the builder pattern


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


## PostApiListingTypes

> PostApiListingTypes(ctx).PostApiListingTypesRequest(postApiListingTypesRequest).IdempotencyKey(idempotencyKey).Execute()

Create a listing type



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
	postApiListingTypesRequest := *openapiclient.NewPostApiListingTypesRequest("Key_example", "Name_example") // PostApiListingTypesRequest | 
	idempotencyKey := "idempotencyKey_example" // string | Required on transactional routes. Reuse the same key for every retry of one logical operation; the API replays the stored outcome. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ListingsAPI.PostApiListingTypes(context.Background()).PostApiListingTypesRequest(postApiListingTypesRequest).IdempotencyKey(idempotencyKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ListingsAPI.PostApiListingTypes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostApiListingTypesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **postApiListingTypesRequest** | [**PostApiListingTypesRequest**](PostApiListingTypesRequest.md) |  | 
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

