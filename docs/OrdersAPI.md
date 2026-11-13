# \OrdersAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetApiOrdersByOrderId**](OrdersAPI.md#GetApiOrdersByOrderId) | **Get** /api/orders/{orderId} | Read a durable confirmation and its requested items
[**PostApiOrdersByOrderIdCancel**](OrdersAPI.md#PostApiOrdersByOrderIdCancel) | **Post** /api/orders/{orderId}/cancel | Cancel the entire order and preserve its history; does not refund payment



## GetApiOrdersByOrderId

> GetApiOrdersByOrderId(ctx, orderId).Execute()

Read a durable confirmation and its requested items



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
	orderId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.OrdersAPI.GetApiOrdersByOrderId(context.Background(), orderId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrdersAPI.GetApiOrdersByOrderId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orderId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetApiOrdersByOrderIdRequest struct via the builder pattern


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


## PostApiOrdersByOrderIdCancel

> PostApiOrdersByOrderIdCancel(ctx, orderId).PostApiOrdersByOrderIdCancelRequest(postApiOrdersByOrderIdCancelRequest).IdempotencyKey(idempotencyKey).Execute()

Cancel the entire order and preserve its history; does not refund payment



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
	orderId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	postApiOrdersByOrderIdCancelRequest := *openapiclient.NewPostApiOrdersByOrderIdCancelRequest("Reason_example") // PostApiOrdersByOrderIdCancelRequest | 
	idempotencyKey := "idempotencyKey_example" // string | Required on transactional routes. Reuse the same key for every retry of one logical operation; the API replays the stored outcome. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.OrdersAPI.PostApiOrdersByOrderIdCancel(context.Background(), orderId).PostApiOrdersByOrderIdCancelRequest(postApiOrdersByOrderIdCancelRequest).IdempotencyKey(idempotencyKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrdersAPI.PostApiOrdersByOrderIdCancel``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orderId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostApiOrdersByOrderIdCancelRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **postApiOrdersByOrderIdCancelRequest** | [**PostApiOrdersByOrderIdCancelRequest**](PostApiOrdersByOrderIdCancelRequest.md) |  | 
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

