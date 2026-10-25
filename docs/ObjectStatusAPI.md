# \ObjectStatusAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetApiEventsByEventIdObjectStatuses**](ObjectStatusAPI.md#GetApiEventsByEventIdObjectStatuses) | **Get** /api/events/{eventId}/object-statuses | List object status overlays for an event
[**PostApiEventsActionsChangeObjectStatus**](ObjectStatusAPI.md#PostApiEventsActionsChangeObjectStatus) | **Post** /api/events/actions/change-object-status | Atomically change object statuses across events (all-or-nothing batch)
[**PostApiEventsByEventIdActionsBook**](ObjectStatusAPI.md#PostApiEventsByEventIdActionsBook) | **Post** /api/events/{eventId}/actions/book | Create a durable administrative booking on unheld inventory (or on a hold token)
[**PostApiEventsByEventIdActionsChangeObjectStatus**](ObjectStatusAPI.md#PostApiEventsByEventIdActionsChangeObjectStatus) | **Post** /api/events/{eventId}/actions/change-object-status | Change object status (booked / free / resale / retired / custom)
[**PostApiEventsByEventIdActionsHold**](ObjectStatusAPI.md#PostApiEventsByEventIdActionsHold) | **Post** /api/events/{eventId}/actions/hold | Hold general admission places administratively (issues a GA hold token)
[**PostApiEventsByEventIdActionsPutUpForResale**](ObjectStatusAPI.md#PostApiEventsByEventIdActionsPutUpForResale) | **Post** /api/events/{eventId}/actions/put-up-for-resale | Put booked objects up for resale
[**PostApiEventsByEventIdActionsRelease**](ObjectStatusAPI.md#PostApiEventsByEventIdActionsRelease) | **Post** /api/events/{eventId}/actions/release | Release unsold objects; sales require explicit order cancellation
[**PostApiEventsByEventIdActionsRetire**](ObjectStatusAPI.md#PostApiEventsByEventIdActionsRetire) | **Post** /api/events/{eventId}/actions/retire | Retire objects (not for sale, history kept)
[**PostApiEventsByEventIdActionsUpdateExtraData**](ObjectStatusAPI.md#PostApiEventsByEventIdActionsUpdateExtraData) | **Post** /api/events/{eventId}/actions/update-extra-data | Update extraData of objects without changing their status



## GetApiEventsByEventIdObjectStatuses

> GetApiEventsByEventIdObjectStatuses(ctx, eventId).Execute()

List object status overlays for an event



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
	eventId := "eventId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ObjectStatusAPI.GetApiEventsByEventIdObjectStatuses(context.Background(), eventId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectStatusAPI.GetApiEventsByEventIdObjectStatuses``: %v\n", err)
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

Other parameters are passed through a pointer to a apiGetApiEventsByEventIdObjectStatusesRequest struct via the builder pattern


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


## PostApiEventsActionsChangeObjectStatus

> PostApiEventsActionsChangeObjectStatus(ctx).PostApiEventsActionsChangeObjectStatusRequest(postApiEventsActionsChangeObjectStatusRequest).IdempotencyKey(idempotencyKey).Execute()

Atomically change object statuses across events (all-or-nothing batch)



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
	postApiEventsActionsChangeObjectStatusRequest := *openapiclient.NewPostApiEventsActionsChangeObjectStatusRequest([]openapiclient.PostApiEventsActionsChangeObjectStatusRequestStatusChangesInner{*openapiclient.NewPostApiEventsActionsChangeObjectStatusRequestStatusChangesInner([]openapiclient.PostApiEventsActionsChangeObjectStatusRequestStatusChangesInnerObjectsInner{*openapiclient.NewPostApiEventsActionsChangeObjectStatusRequestStatusChangesInnerObjectsInner("ObjectId_example")})}) // PostApiEventsActionsChangeObjectStatusRequest | 
	idempotencyKey := "idempotencyKey_example" // string | Required on transactional routes. Reuse the same key for every retry of one logical operation; the API replays the stored outcome. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ObjectStatusAPI.PostApiEventsActionsChangeObjectStatus(context.Background()).PostApiEventsActionsChangeObjectStatusRequest(postApiEventsActionsChangeObjectStatusRequest).IdempotencyKey(idempotencyKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectStatusAPI.PostApiEventsActionsChangeObjectStatus``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostApiEventsActionsChangeObjectStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **postApiEventsActionsChangeObjectStatusRequest** | [**PostApiEventsActionsChangeObjectStatusRequest**](PostApiEventsActionsChangeObjectStatusRequest.md) |  | 
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


## PostApiEventsByEventIdActionsBook

> PostApiEventsByEventIdActionsBook(ctx, eventId).PostApiEventsByEventIdActionsChangeObjectStatusRequest(postApiEventsByEventIdActionsChangeObjectStatusRequest).IdempotencyKey(idempotencyKey).Execute()

Create a durable administrative booking on unheld inventory (or on a hold token)



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
	eventId := "eventId_example" // string | 
	postApiEventsByEventIdActionsChangeObjectStatusRequest := *openapiclient.NewPostApiEventsByEventIdActionsChangeObjectStatusRequest([]openapiclient.PostApiEventsActionsChangeObjectStatusRequestStatusChangesInnerObjectsInner{*openapiclient.NewPostApiEventsActionsChangeObjectStatusRequestStatusChangesInnerObjectsInner("ObjectId_example")}) // PostApiEventsByEventIdActionsChangeObjectStatusRequest | 
	idempotencyKey := "idempotencyKey_example" // string | Required on transactional routes. Reuse the same key for every retry of one logical operation; the API replays the stored outcome. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ObjectStatusAPI.PostApiEventsByEventIdActionsBook(context.Background(), eventId).PostApiEventsByEventIdActionsChangeObjectStatusRequest(postApiEventsByEventIdActionsChangeObjectStatusRequest).IdempotencyKey(idempotencyKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectStatusAPI.PostApiEventsByEventIdActionsBook``: %v\n", err)
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

Other parameters are passed through a pointer to a apiPostApiEventsByEventIdActionsBookRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **postApiEventsByEventIdActionsChangeObjectStatusRequest** | [**PostApiEventsByEventIdActionsChangeObjectStatusRequest**](PostApiEventsByEventIdActionsChangeObjectStatusRequest.md) |  | 
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


## PostApiEventsByEventIdActionsChangeObjectStatus

> PostApiEventsByEventIdActionsChangeObjectStatus(ctx, eventId).PostApiEventsByEventIdActionsChangeObjectStatusRequest(postApiEventsByEventIdActionsChangeObjectStatusRequest).IdempotencyKey(idempotencyKey).Execute()

Change object status (booked / free / resale / retired / custom)



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
	eventId := "eventId_example" // string | 
	postApiEventsByEventIdActionsChangeObjectStatusRequest := *openapiclient.NewPostApiEventsByEventIdActionsChangeObjectStatusRequest([]openapiclient.PostApiEventsActionsChangeObjectStatusRequestStatusChangesInnerObjectsInner{*openapiclient.NewPostApiEventsActionsChangeObjectStatusRequestStatusChangesInnerObjectsInner("ObjectId_example")}) // PostApiEventsByEventIdActionsChangeObjectStatusRequest | 
	idempotencyKey := "idempotencyKey_example" // string | Required on transactional routes. Reuse the same key for every retry of one logical operation; the API replays the stored outcome. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ObjectStatusAPI.PostApiEventsByEventIdActionsChangeObjectStatus(context.Background(), eventId).PostApiEventsByEventIdActionsChangeObjectStatusRequest(postApiEventsByEventIdActionsChangeObjectStatusRequest).IdempotencyKey(idempotencyKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectStatusAPI.PostApiEventsByEventIdActionsChangeObjectStatus``: %v\n", err)
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

Other parameters are passed through a pointer to a apiPostApiEventsByEventIdActionsChangeObjectStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **postApiEventsByEventIdActionsChangeObjectStatusRequest** | [**PostApiEventsByEventIdActionsChangeObjectStatusRequest**](PostApiEventsByEventIdActionsChangeObjectStatusRequest.md) |  | 
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


## PostApiEventsByEventIdActionsHold

> PostApiEventsByEventIdActionsHold(ctx, eventId).PostApiEventsByEventIdActionsChangeObjectStatusRequest(postApiEventsByEventIdActionsChangeObjectStatusRequest).IdempotencyKey(idempotencyKey).Execute()

Hold general admission places administratively (issues a GA hold token)



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
	eventId := "eventId_example" // string | 
	postApiEventsByEventIdActionsChangeObjectStatusRequest := *openapiclient.NewPostApiEventsByEventIdActionsChangeObjectStatusRequest([]openapiclient.PostApiEventsActionsChangeObjectStatusRequestStatusChangesInnerObjectsInner{*openapiclient.NewPostApiEventsActionsChangeObjectStatusRequestStatusChangesInnerObjectsInner("ObjectId_example")}) // PostApiEventsByEventIdActionsChangeObjectStatusRequest | 
	idempotencyKey := "idempotencyKey_example" // string | Required on transactional routes. Reuse the same key for every retry of one logical operation; the API replays the stored outcome. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ObjectStatusAPI.PostApiEventsByEventIdActionsHold(context.Background(), eventId).PostApiEventsByEventIdActionsChangeObjectStatusRequest(postApiEventsByEventIdActionsChangeObjectStatusRequest).IdempotencyKey(idempotencyKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectStatusAPI.PostApiEventsByEventIdActionsHold``: %v\n", err)
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

Other parameters are passed through a pointer to a apiPostApiEventsByEventIdActionsHoldRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **postApiEventsByEventIdActionsChangeObjectStatusRequest** | [**PostApiEventsByEventIdActionsChangeObjectStatusRequest**](PostApiEventsByEventIdActionsChangeObjectStatusRequest.md) |  | 
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


## PostApiEventsByEventIdActionsPutUpForResale

> PostApiEventsByEventIdActionsPutUpForResale(ctx, eventId).PostApiEventsByEventIdActionsChangeObjectStatusRequest(postApiEventsByEventIdActionsChangeObjectStatusRequest).IdempotencyKey(idempotencyKey).Execute()

Put booked objects up for resale



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
	eventId := "eventId_example" // string | 
	postApiEventsByEventIdActionsChangeObjectStatusRequest := *openapiclient.NewPostApiEventsByEventIdActionsChangeObjectStatusRequest([]openapiclient.PostApiEventsActionsChangeObjectStatusRequestStatusChangesInnerObjectsInner{*openapiclient.NewPostApiEventsActionsChangeObjectStatusRequestStatusChangesInnerObjectsInner("ObjectId_example")}) // PostApiEventsByEventIdActionsChangeObjectStatusRequest | 
	idempotencyKey := "idempotencyKey_example" // string | Required on transactional routes. Reuse the same key for every retry of one logical operation; the API replays the stored outcome. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ObjectStatusAPI.PostApiEventsByEventIdActionsPutUpForResale(context.Background(), eventId).PostApiEventsByEventIdActionsChangeObjectStatusRequest(postApiEventsByEventIdActionsChangeObjectStatusRequest).IdempotencyKey(idempotencyKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectStatusAPI.PostApiEventsByEventIdActionsPutUpForResale``: %v\n", err)
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

Other parameters are passed through a pointer to a apiPostApiEventsByEventIdActionsPutUpForResaleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **postApiEventsByEventIdActionsChangeObjectStatusRequest** | [**PostApiEventsByEventIdActionsChangeObjectStatusRequest**](PostApiEventsByEventIdActionsChangeObjectStatusRequest.md) |  | 
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


## PostApiEventsByEventIdActionsRelease

> PostApiEventsByEventIdActionsRelease(ctx, eventId).PostApiEventsByEventIdActionsChangeObjectStatusRequest(postApiEventsByEventIdActionsChangeObjectStatusRequest).IdempotencyKey(idempotencyKey).Execute()

Release unsold objects; sales require explicit order cancellation



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
	eventId := "eventId_example" // string | 
	postApiEventsByEventIdActionsChangeObjectStatusRequest := *openapiclient.NewPostApiEventsByEventIdActionsChangeObjectStatusRequest([]openapiclient.PostApiEventsActionsChangeObjectStatusRequestStatusChangesInnerObjectsInner{*openapiclient.NewPostApiEventsActionsChangeObjectStatusRequestStatusChangesInnerObjectsInner("ObjectId_example")}) // PostApiEventsByEventIdActionsChangeObjectStatusRequest | 
	idempotencyKey := "idempotencyKey_example" // string | Required on transactional routes. Reuse the same key for every retry of one logical operation; the API replays the stored outcome. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ObjectStatusAPI.PostApiEventsByEventIdActionsRelease(context.Background(), eventId).PostApiEventsByEventIdActionsChangeObjectStatusRequest(postApiEventsByEventIdActionsChangeObjectStatusRequest).IdempotencyKey(idempotencyKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectStatusAPI.PostApiEventsByEventIdActionsRelease``: %v\n", err)
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

Other parameters are passed through a pointer to a apiPostApiEventsByEventIdActionsReleaseRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **postApiEventsByEventIdActionsChangeObjectStatusRequest** | [**PostApiEventsByEventIdActionsChangeObjectStatusRequest**](PostApiEventsByEventIdActionsChangeObjectStatusRequest.md) |  | 
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


## PostApiEventsByEventIdActionsRetire

> PostApiEventsByEventIdActionsRetire(ctx, eventId).PostApiEventsByEventIdActionsChangeObjectStatusRequest(postApiEventsByEventIdActionsChangeObjectStatusRequest).IdempotencyKey(idempotencyKey).Execute()

Retire objects (not for sale, history kept)



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
	eventId := "eventId_example" // string | 
	postApiEventsByEventIdActionsChangeObjectStatusRequest := *openapiclient.NewPostApiEventsByEventIdActionsChangeObjectStatusRequest([]openapiclient.PostApiEventsActionsChangeObjectStatusRequestStatusChangesInnerObjectsInner{*openapiclient.NewPostApiEventsActionsChangeObjectStatusRequestStatusChangesInnerObjectsInner("ObjectId_example")}) // PostApiEventsByEventIdActionsChangeObjectStatusRequest | 
	idempotencyKey := "idempotencyKey_example" // string | Required on transactional routes. Reuse the same key for every retry of one logical operation; the API replays the stored outcome. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ObjectStatusAPI.PostApiEventsByEventIdActionsRetire(context.Background(), eventId).PostApiEventsByEventIdActionsChangeObjectStatusRequest(postApiEventsByEventIdActionsChangeObjectStatusRequest).IdempotencyKey(idempotencyKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectStatusAPI.PostApiEventsByEventIdActionsRetire``: %v\n", err)
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

Other parameters are passed through a pointer to a apiPostApiEventsByEventIdActionsRetireRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **postApiEventsByEventIdActionsChangeObjectStatusRequest** | [**PostApiEventsByEventIdActionsChangeObjectStatusRequest**](PostApiEventsByEventIdActionsChangeObjectStatusRequest.md) |  | 
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


## PostApiEventsByEventIdActionsUpdateExtraData

> PostApiEventsByEventIdActionsUpdateExtraData(ctx, eventId).PostApiEventsByEventIdActionsUpdateExtraDataRequest(postApiEventsByEventIdActionsUpdateExtraDataRequest).IdempotencyKey(idempotencyKey).Execute()

Update extraData of objects without changing their status



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
	eventId := "eventId_example" // string | 
	postApiEventsByEventIdActionsUpdateExtraDataRequest := *openapiclient.NewPostApiEventsByEventIdActionsUpdateExtraDataRequest([]openapiclient.PostApiEventsByEventIdActionsUpdateExtraDataRequestObjectsInner{*openapiclient.NewPostApiEventsByEventIdActionsUpdateExtraDataRequestObjectsInner("ObjectId_example", map[string]interface{}{"key": interface{}(123)})}) // PostApiEventsByEventIdActionsUpdateExtraDataRequest | 
	idempotencyKey := "idempotencyKey_example" // string | Required on transactional routes. Reuse the same key for every retry of one logical operation; the API replays the stored outcome. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ObjectStatusAPI.PostApiEventsByEventIdActionsUpdateExtraData(context.Background(), eventId).PostApiEventsByEventIdActionsUpdateExtraDataRequest(postApiEventsByEventIdActionsUpdateExtraDataRequest).IdempotencyKey(idempotencyKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectStatusAPI.PostApiEventsByEventIdActionsUpdateExtraData``: %v\n", err)
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

Other parameters are passed through a pointer to a apiPostApiEventsByEventIdActionsUpdateExtraDataRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **postApiEventsByEventIdActionsUpdateExtraDataRequest** | [**PostApiEventsByEventIdActionsUpdateExtraDataRequest**](PostApiEventsByEventIdActionsUpdateExtraDataRequest.md) |  | 
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

