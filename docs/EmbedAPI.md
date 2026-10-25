# \EmbedAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetApiTicketBuyers**](EmbedAPI.md#GetApiTicketBuyers) | **Get** /api/ticket-buyers | List registered ticket buyer ids
[**PostApiEmbedTokens**](EmbedAPI.md#PostApiEmbedTokens) | **Post** /api/embed-tokens | Mint a short-lived designer/manager embed token
[**PostApiTicketBuyers**](EmbedAPI.md#PostApiTicketBuyers) | **Post** /api/ticket-buyers | Register ticket buyer ids
[**PostApiTicketBuyersActionsRemove**](EmbedAPI.md#PostApiTicketBuyersActionsRemove) | **Post** /api/ticket-buyers/actions/remove | Remove ticket buyer ids
[**PostApiViewerTokens**](EmbedAPI.md#PostApiViewerTokens) | **Post** /api/viewer-tokens | Mint a short-lived buyer (viewer) token with the partner API key



## GetApiTicketBuyers

> GetApiTicketBuyers(ctx).Limit(limit).Cursor(cursor).Execute()

List registered ticket buyer ids



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
	limit := int32(56) // int32 |  (optional)
	cursor := "cursor_example" // string | `nextCursor` of the previous page. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.EmbedAPI.GetApiTicketBuyers(context.Background()).Limit(limit).Cursor(cursor).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EmbedAPI.GetApiTicketBuyers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetApiTicketBuyersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** |  | 
 **cursor** | **string** | &#x60;nextCursor&#x60; of the previous page. | 

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


## PostApiEmbedTokens

> PostApiEmbedTokens(ctx).PostApiEmbedTokensRequest(postApiEmbedTokensRequest).IdempotencyKey(idempotencyKey).Execute()

Mint a short-lived designer/manager embed token



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
	postApiEmbedTokensRequest := *openapiclient.NewPostApiEmbedTokensRequest("Purpose_example") // PostApiEmbedTokensRequest | 
	idempotencyKey := "idempotencyKey_example" // string | Required on transactional routes. Reuse the same key for every retry of one logical operation; the API replays the stored outcome. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.EmbedAPI.PostApiEmbedTokens(context.Background()).PostApiEmbedTokensRequest(postApiEmbedTokensRequest).IdempotencyKey(idempotencyKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EmbedAPI.PostApiEmbedTokens``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostApiEmbedTokensRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **postApiEmbedTokensRequest** | [**PostApiEmbedTokensRequest**](PostApiEmbedTokensRequest.md) |  | 
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


## PostApiTicketBuyers

> PostApiTicketBuyers(ctx).PostApiTicketBuyersRequest(postApiTicketBuyersRequest).IdempotencyKey(idempotencyKey).Execute()

Register ticket buyer ids



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
	postApiTicketBuyersRequest := *openapiclient.NewPostApiTicketBuyersRequest([]string{"Ids_example"}) // PostApiTicketBuyersRequest | 
	idempotencyKey := "idempotencyKey_example" // string | Required on transactional routes. Reuse the same key for every retry of one logical operation; the API replays the stored outcome. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.EmbedAPI.PostApiTicketBuyers(context.Background()).PostApiTicketBuyersRequest(postApiTicketBuyersRequest).IdempotencyKey(idempotencyKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EmbedAPI.PostApiTicketBuyers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostApiTicketBuyersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **postApiTicketBuyersRequest** | [**PostApiTicketBuyersRequest**](PostApiTicketBuyersRequest.md) |  | 
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


## PostApiTicketBuyersActionsRemove

> PostApiTicketBuyersActionsRemove(ctx).PostApiTicketBuyersRequest(postApiTicketBuyersRequest).IdempotencyKey(idempotencyKey).Execute()

Remove ticket buyer ids



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
	postApiTicketBuyersRequest := *openapiclient.NewPostApiTicketBuyersRequest([]string{"Ids_example"}) // PostApiTicketBuyersRequest | 
	idempotencyKey := "idempotencyKey_example" // string | Required on transactional routes. Reuse the same key for every retry of one logical operation; the API replays the stored outcome. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.EmbedAPI.PostApiTicketBuyersActionsRemove(context.Background()).PostApiTicketBuyersRequest(postApiTicketBuyersRequest).IdempotencyKey(idempotencyKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EmbedAPI.PostApiTicketBuyersActionsRemove``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostApiTicketBuyersActionsRemoveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **postApiTicketBuyersRequest** | [**PostApiTicketBuyersRequest**](PostApiTicketBuyersRequest.md) |  | 
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


## PostApiViewerTokens

> PostApiViewerTokens(ctx).PostApiViewerTokensRequest(postApiViewerTokensRequest).IdempotencyKey(idempotencyKey).Execute()

Mint a short-lived buyer (viewer) token with the partner API key



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
	postApiViewerTokensRequest := *openapiclient.NewPostApiViewerTokensRequest() // PostApiViewerTokensRequest | 
	idempotencyKey := "idempotencyKey_example" // string | Required on transactional routes. Reuse the same key for every retry of one logical operation; the API replays the stored outcome. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.EmbedAPI.PostApiViewerTokens(context.Background()).PostApiViewerTokensRequest(postApiViewerTokensRequest).IdempotencyKey(idempotencyKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EmbedAPI.PostApiViewerTokens``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostApiViewerTokensRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **postApiViewerTokensRequest** | [**PostApiViewerTokensRequest**](PostApiViewerTokensRequest.md) |  | 
 **idempotencyKey** | **string** | Required on transactional routes. Reuse the same key for every retry of one logical operation; the API replays the stored outcome. | 

### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json, multipart/form-data, text/plain
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

