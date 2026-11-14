# \WebhooksAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteApiWebhooksById**](WebhooksAPI.md#DeleteApiWebhooksById) | **Delete** /api/webhooks/{id} | Delete a webhook
[**GetApiWebhooks**](WebhooksAPI.md#GetApiWebhooks) | **Get** /api/webhooks | List webhooks for the authenticated tenant
[**GetApiWebhooksByIdDeliveries**](WebhooksAPI.md#GetApiWebhooksByIdDeliveries) | **Get** /api/webhooks/{id}/deliveries | List deliveries of a webhook endpoint
[**GetApiWebhooksByIdDeliveriesByDeliveryId**](WebhooksAPI.md#GetApiWebhooksByIdDeliveriesByDeliveryId) | **Get** /api/webhooks/{id}/deliveries/{deliveryId} | Get one webhook delivery with its request payload
[**GetApiWebhooksEventTypes**](WebhooksAPI.md#GetApiWebhooksEventTypes) | **Get** /api/webhooks/event-types | List subscribable webhook event types
[**PatchApiWebhooksById**](WebhooksAPI.md#PatchApiWebhooksById) | **Patch** /api/webhooks/{id} | Update a webhook
[**PostApiWebhooks**](WebhooksAPI.md#PostApiWebhooks) | **Post** /api/webhooks | Create a webhook endpoint
[**PostApiWebhooksByIdDeliveriesByDeliveryIdReplay**](WebhooksAPI.md#PostApiWebhooksByIdDeliveriesByDeliveryIdReplay) | **Post** /api/webhooks/{id}/deliveries/{deliveryId}/replay | Replay a webhook delivery
[**PostApiWebhooksByIdRotateSecret**](WebhooksAPI.md#PostApiWebhooksByIdRotateSecret) | **Post** /api/webhooks/{id}/rotate-secret | Rotate a webhook signing secret
[**PostApiWebhooksByIdTest**](WebhooksAPI.md#PostApiWebhooksByIdTest) | **Post** /api/webhooks/{id}/test | Send a signed test event



## DeleteApiWebhooksById

> DeleteApiWebhooksById(ctx, id).IdempotencyKey(idempotencyKey).Execute()

Delete a webhook



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
	id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Webhook UUID (stored lowercase)
	idempotencyKey := "idempotencyKey_example" // string | Required on transactional routes. Reuse the same key for every retry of one logical operation; the API replays the stored outcome. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.WebhooksAPI.DeleteApiWebhooksById(context.Background(), id).IdempotencyKey(idempotencyKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebhooksAPI.DeleteApiWebhooksById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Webhook UUID (stored lowercase) | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteApiWebhooksByIdRequest struct via the builder pattern


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


## GetApiWebhooks

> GetApiWebhooks(ctx).Execute()

List webhooks for the authenticated tenant



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
	r, err := apiClient.WebhooksAPI.GetApiWebhooks(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebhooksAPI.GetApiWebhooks``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetApiWebhooksRequest struct via the builder pattern


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


## GetApiWebhooksByIdDeliveries

> GetApiWebhooksByIdDeliveries(ctx, id).Limit(limit).Cursor(cursor).Status(status).EventType(eventType).Source(source).Execute()

List deliveries of a webhook endpoint



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
	id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Webhook UUID (stored lowercase)
	limit := float32(8.14) // float32 | Page size (1..100, default 20). (optional)
	cursor := "cursor_example" // string | Opaque `nextCursor` value from the previous page. (optional)
	status := "status_example" // string | `pending` also matches deliveries currently being attempted (`processing`). (optional)
	eventType := "eventType_example" // string |  (optional)
	source := "source_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.WebhooksAPI.GetApiWebhooksByIdDeliveries(context.Background(), id).Limit(limit).Cursor(cursor).Status(status).EventType(eventType).Source(source).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebhooksAPI.GetApiWebhooksByIdDeliveries``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Webhook UUID (stored lowercase) | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetApiWebhooksByIdDeliveriesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **float32** | Page size (1..100, default 20). | 
 **cursor** | **string** | Opaque &#x60;nextCursor&#x60; value from the previous page. | 
 **status** | **string** | &#x60;pending&#x60; also matches deliveries currently being attempted (&#x60;processing&#x60;). | 
 **eventType** | **string** |  | 
 **source** | **string** |  | 

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


## GetApiWebhooksByIdDeliveriesByDeliveryId

> GetApiWebhooksByIdDeliveriesByDeliveryId(ctx, id, deliveryId).Execute()

Get one webhook delivery with its request payload



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
	id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Webhook UUID (stored lowercase)
	deliveryId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.WebhooksAPI.GetApiWebhooksByIdDeliveriesByDeliveryId(context.Background(), id, deliveryId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebhooksAPI.GetApiWebhooksByIdDeliveriesByDeliveryId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Webhook UUID (stored lowercase) | 
**deliveryId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetApiWebhooksByIdDeliveriesByDeliveryIdRequest struct via the builder pattern


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


## GetApiWebhooksEventTypes

> GetApiWebhooksEventTypes(ctx).Execute()

List subscribable webhook event types



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
	r, err := apiClient.WebhooksAPI.GetApiWebhooksEventTypes(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebhooksAPI.GetApiWebhooksEventTypes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetApiWebhooksEventTypesRequest struct via the builder pattern


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


## PatchApiWebhooksById

> PatchApiWebhooksById(ctx, id).PatchApiWebhooksByIdRequest(patchApiWebhooksByIdRequest).IdempotencyKey(idempotencyKey).Execute()

Update a webhook



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
	id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Webhook UUID (stored lowercase)
	patchApiWebhooksByIdRequest := *openapiclient.NewPatchApiWebhooksByIdRequest() // PatchApiWebhooksByIdRequest | 
	idempotencyKey := "idempotencyKey_example" // string | Required on transactional routes. Reuse the same key for every retry of one logical operation; the API replays the stored outcome. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.WebhooksAPI.PatchApiWebhooksById(context.Background(), id).PatchApiWebhooksByIdRequest(patchApiWebhooksByIdRequest).IdempotencyKey(idempotencyKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebhooksAPI.PatchApiWebhooksById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Webhook UUID (stored lowercase) | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchApiWebhooksByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **patchApiWebhooksByIdRequest** | [**PatchApiWebhooksByIdRequest**](PatchApiWebhooksByIdRequest.md) |  | 
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


## PostApiWebhooks

> PostApiWebhooks(ctx).PostApiWebhooksRequest(postApiWebhooksRequest).IdempotencyKey(idempotencyKey).Execute()

Create a webhook endpoint



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
	postApiWebhooksRequest := *openapiclient.NewPostApiWebhooksRequest("Url_example") // PostApiWebhooksRequest | 
	idempotencyKey := "idempotencyKey_example" // string | Required on transactional routes. Reuse the same key for every retry of one logical operation; the API replays the stored outcome. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.WebhooksAPI.PostApiWebhooks(context.Background()).PostApiWebhooksRequest(postApiWebhooksRequest).IdempotencyKey(idempotencyKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebhooksAPI.PostApiWebhooks``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostApiWebhooksRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **postApiWebhooksRequest** | [**PostApiWebhooksRequest**](PostApiWebhooksRequest.md) |  | 
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


## PostApiWebhooksByIdDeliveriesByDeliveryIdReplay

> PostApiWebhooksByIdDeliveriesByDeliveryIdReplay(ctx, id, deliveryId).IdempotencyKey(idempotencyKey).Execute()

Replay a webhook delivery



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
	id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Webhook UUID (stored lowercase)
	deliveryId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	idempotencyKey := "idempotencyKey_example" // string | Optional. Repeating a replay with the same key returns the first result instead of sending again. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.WebhooksAPI.PostApiWebhooksByIdDeliveriesByDeliveryIdReplay(context.Background(), id, deliveryId).IdempotencyKey(idempotencyKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebhooksAPI.PostApiWebhooksByIdDeliveriesByDeliveryIdReplay``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Webhook UUID (stored lowercase) | 
**deliveryId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostApiWebhooksByIdDeliveriesByDeliveryIdReplayRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **idempotencyKey** | **string** | Optional. Repeating a replay with the same key returns the first result instead of sending again. | 

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


## PostApiWebhooksByIdRotateSecret

> PostApiWebhooksByIdRotateSecret(ctx, id).IdempotencyKey(idempotencyKey).Execute()

Rotate a webhook signing secret



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
	id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Webhook UUID (stored lowercase)
	idempotencyKey := "idempotencyKey_example" // string | Required on transactional routes. Reuse the same key for every retry of one logical operation; the API replays the stored outcome. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.WebhooksAPI.PostApiWebhooksByIdRotateSecret(context.Background(), id).IdempotencyKey(idempotencyKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebhooksAPI.PostApiWebhooksByIdRotateSecret``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Webhook UUID (stored lowercase) | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostApiWebhooksByIdRotateSecretRequest struct via the builder pattern


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


## PostApiWebhooksByIdTest

> PostApiWebhooksByIdTest(ctx, id).IdempotencyKey(idempotencyKey).Execute()

Send a signed test event



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
	id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Webhook UUID (stored lowercase)
	idempotencyKey := "idempotencyKey_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.WebhooksAPI.PostApiWebhooksByIdTest(context.Background(), id).IdempotencyKey(idempotencyKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebhooksAPI.PostApiWebhooksByIdTest``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Webhook UUID (stored lowercase) | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostApiWebhooksByIdTestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **idempotencyKey** | **string** |  | 

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

