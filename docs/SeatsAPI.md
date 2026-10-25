# \SeatsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetApiChannels**](SeatsAPI.md#GetApiChannels) | **Get** /api/channels | List sales channels
[**GetApiHoldTokensByToken**](SeatsAPI.md#GetApiHoldTokensByToken) | **Get** /api/hold-tokens/{token} | Validate a hold-token session
[**GetApiSeatsGroup**](SeatsAPI.md#GetApiSeatsGroup) | **Get** /api/seats/group | Event-group availability intersection by label
[**PostApiGaConfirm**](SeatsAPI.md#PostApiGaConfirm) | **Post** /api/ga/confirm | Confirm GA purchase (Postgres sold_count)
[**PostApiGaRelease**](SeatsAPI.md#PostApiGaRelease) | **Post** /api/ga/release | Release a GA capacity hold
[**PostApiGaReserve**](SeatsAPI.md#PostApiGaReserve) | **Post** /api/ga/reserve | Hold GA area capacity (Redis)
[**PostApiHoldTokens**](SeatsAPI.md#PostApiHoldTokens) | **Post** /api/hold-tokens | Create an empty hold-token session
[**PostApiHoldTokensByToken**](SeatsAPI.md#PostApiHoldTokensByToken) | **Post** /api/hold-tokens/{token} | Change hold-token expiration (max 2h from creation)
[**PostApiSeatsBestAvailable**](SeatsAPI.md#PostApiSeatsBestAvailable) | **Post** /api/seats/best-available | Suggest or atomically hold best available seats
[**PostApiSeatsConfirm**](SeatsAPI.md#PostApiSeatsConfirm) | **Post** /api/seats/confirm | Accept a durable confirmation intent before asynchronous seat materialization
[**PostApiSeatsHoldsRelease**](SeatsAPI.md#PostApiSeatsHoldsRelease) | **Post** /api/seats/holds/release | Release all seats held by a hold token
[**PostApiSeatsHoldsRenew**](SeatsAPI.md#PostApiSeatsHoldsRenew) | **Post** /api/seats/holds/renew | Renew Redis TTL for all seats in a hold session
[**PostApiSeatsRelease**](SeatsAPI.md#PostApiSeatsRelease) | **Post** /api/seats/release | Clear Redis TTL holds before expiry (checkout abandoned / declined)
[**PostApiSeatsReserve**](SeatsAPI.md#PostApiSeatsReserve) | **Post** /api/seats/reserve | Place a temporary Redis hold (Lua TTL) before payment



## GetApiChannels

> GetApiChannels(ctx).Execute()

List sales channels



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
	r, err := apiClient.SeatsAPI.GetApiChannels(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SeatsAPI.GetApiChannels``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetApiChannelsRequest struct via the builder pattern


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


## GetApiHoldTokensByToken

> GetApiHoldTokensByToken(ctx, token).Execute()

Validate a hold-token session



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
	token := "token_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SeatsAPI.GetApiHoldTokensByToken(context.Background(), token).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SeatsAPI.GetApiHoldTokensByToken``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**token** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetApiHoldTokensByTokenRequest struct via the builder pattern


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


## GetApiSeatsGroup

> GetApiSeatsGroup(ctx).EventIds(eventIds).Execute()

Event-group availability intersection by label



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
	eventIds := "eventIds_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SeatsAPI.GetApiSeatsGroup(context.Background()).EventIds(eventIds).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SeatsAPI.GetApiSeatsGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetApiSeatsGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **eventIds** | **string** |  | 

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


## PostApiGaConfirm

> PostApiGaConfirm(ctx).IdempotencyKey(idempotencyKey).PostApiGaConfirmRequest(postApiGaConfirmRequest).Execute()

Confirm GA purchase (Postgres sold_count)



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
	idempotencyKey := "idempotencyKey_example" // string | Required. Reuse with the same request on retry; a different payload conflicts. Seat confirmation binds this key durably in PostgreSQL; hold mutations use bounded Redis receipts.
	postApiGaConfirmRequest := *openapiclient.NewPostApiGaConfirmRequest("EventId_example", "AreaId_example", int32(123), "UserId_example") // PostApiGaConfirmRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SeatsAPI.PostApiGaConfirm(context.Background()).IdempotencyKey(idempotencyKey).PostApiGaConfirmRequest(postApiGaConfirmRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SeatsAPI.PostApiGaConfirm``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostApiGaConfirmRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **idempotencyKey** | **string** | Required. Reuse with the same request on retry; a different payload conflicts. Seat confirmation binds this key durably in PostgreSQL; hold mutations use bounded Redis receipts. | 
 **postApiGaConfirmRequest** | [**PostApiGaConfirmRequest**](PostApiGaConfirmRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json, multipart/form-data, text/plain
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostApiGaRelease

> PostApiGaRelease(ctx).IdempotencyKey(idempotencyKey).PostApiGaReleaseRequest(postApiGaReleaseRequest).Execute()

Release a GA capacity hold



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
	idempotencyKey := "idempotencyKey_example" // string | Required. Reuse with the same request on retry; a different payload conflicts. Seat confirmation binds this key durably in PostgreSQL; hold mutations use bounded Redis receipts.
	postApiGaReleaseRequest := *openapiclient.NewPostApiGaReleaseRequest("EventId_example", "AreaId_example", "UserId_example") // PostApiGaReleaseRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SeatsAPI.PostApiGaRelease(context.Background()).IdempotencyKey(idempotencyKey).PostApiGaReleaseRequest(postApiGaReleaseRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SeatsAPI.PostApiGaRelease``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostApiGaReleaseRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **idempotencyKey** | **string** | Required. Reuse with the same request on retry; a different payload conflicts. Seat confirmation binds this key durably in PostgreSQL; hold mutations use bounded Redis receipts. | 
 **postApiGaReleaseRequest** | [**PostApiGaReleaseRequest**](PostApiGaReleaseRequest.md) |  | 

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


## PostApiGaReserve

> PostApiGaReserve(ctx).IdempotencyKey(idempotencyKey).PostApiGaReserveRequest(postApiGaReserveRequest).Execute()

Hold GA area capacity (Redis)



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
	idempotencyKey := "idempotencyKey_example" // string | Required. Reuse with the same request on retry; a different payload conflicts. Seat confirmation binds this key durably in PostgreSQL; hold mutations use bounded Redis receipts.
	postApiGaReserveRequest := *openapiclient.NewPostApiGaReserveRequest("EventId_example", "AreaId_example", int32(123), "UserId_example") // PostApiGaReserveRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SeatsAPI.PostApiGaReserve(context.Background()).IdempotencyKey(idempotencyKey).PostApiGaReserveRequest(postApiGaReserveRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SeatsAPI.PostApiGaReserve``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostApiGaReserveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **idempotencyKey** | **string** | Required. Reuse with the same request on retry; a different payload conflicts. Seat confirmation binds this key durably in PostgreSQL; hold mutations use bounded Redis receipts. | 
 **postApiGaReserveRequest** | [**PostApiGaReserveRequest**](PostApiGaReserveRequest.md) |  | 

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


## PostApiHoldTokens

> PostApiHoldTokens(ctx).PostApiHoldTokensRequest(postApiHoldTokensRequest).IdempotencyKey(idempotencyKey).Execute()

Create an empty hold-token session



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
	postApiHoldTokensRequest := *openapiclient.NewPostApiHoldTokensRequest("EventId_example", "UserId_example") // PostApiHoldTokensRequest | 
	idempotencyKey := "idempotencyKey_example" // string | Required on transactional routes. Reuse the same key for every retry of one logical operation; the API replays the stored outcome. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SeatsAPI.PostApiHoldTokens(context.Background()).PostApiHoldTokensRequest(postApiHoldTokensRequest).IdempotencyKey(idempotencyKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SeatsAPI.PostApiHoldTokens``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostApiHoldTokensRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **postApiHoldTokensRequest** | [**PostApiHoldTokensRequest**](PostApiHoldTokensRequest.md) |  | 
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


## PostApiHoldTokensByToken

> PostApiHoldTokensByToken(ctx, token).PostApiHoldTokensByTokenRequest(postApiHoldTokensByTokenRequest).IdempotencyKey(idempotencyKey).Execute()

Change hold-token expiration (max 2h from creation)

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
	token := "token_example" // string | 
	postApiHoldTokensByTokenRequest := *openapiclient.NewPostApiHoldTokensByTokenRequest(int32(123)) // PostApiHoldTokensByTokenRequest | 
	idempotencyKey := "idempotencyKey_example" // string | Required on transactional routes. Reuse the same key for every retry of one logical operation; the API replays the stored outcome. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SeatsAPI.PostApiHoldTokensByToken(context.Background(), token).PostApiHoldTokensByTokenRequest(postApiHoldTokensByTokenRequest).IdempotencyKey(idempotencyKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SeatsAPI.PostApiHoldTokensByToken``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**token** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostApiHoldTokensByTokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **postApiHoldTokensByTokenRequest** | [**PostApiHoldTokensByTokenRequest**](PostApiHoldTokensByTokenRequest.md) |  | 
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


## PostApiSeatsBestAvailable

> PostApiSeatsBestAvailable(ctx).PostApiSeatsBestAvailableRequest(postApiSeatsBestAvailableRequest).IdempotencyKey(idempotencyKey).Execute()

Suggest or atomically hold best available seats



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
	postApiSeatsBestAvailableRequest := *openapiclient.NewPostApiSeatsBestAvailableRequest("EventId_example") // PostApiSeatsBestAvailableRequest | 
	idempotencyKey := "idempotencyKey_example" // string | Required on transactional routes. Reuse the same key for every retry of one logical operation; the API replays the stored outcome. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SeatsAPI.PostApiSeatsBestAvailable(context.Background()).PostApiSeatsBestAvailableRequest(postApiSeatsBestAvailableRequest).IdempotencyKey(idempotencyKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SeatsAPI.PostApiSeatsBestAvailable``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostApiSeatsBestAvailableRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **postApiSeatsBestAvailableRequest** | [**PostApiSeatsBestAvailableRequest**](PostApiSeatsBestAvailableRequest.md) |  | 
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


## PostApiSeatsConfirm

> PostApiSeatsConfirm(ctx).IdempotencyKey(idempotencyKey).PostApiSeatsConfirmRequest(postApiSeatsConfirmRequest).XTenantId(xTenantId).Execute()

Accept a durable confirmation intent before asynchronous seat materialization



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
	idempotencyKey := "idempotencyKey_example" // string | Required key (max 255 chars after trim); durable fallback identity when externalOrderId is omitted.
	postApiSeatsConfirmRequest := *openapiclient.NewPostApiSeatsConfirmRequest("EventId_example", []string{"SeatIds_example"}, "UserId_example") // PostApiSeatsConfirmRequest | 
	xTenantId := "xTenantId_example" // string | Ignored for bootstrap env keys (bound to DEFAULT_TENANT_ID). Dashboard `sk_live_…` keys are tenant-scoped by hash. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SeatsAPI.PostApiSeatsConfirm(context.Background()).IdempotencyKey(idempotencyKey).PostApiSeatsConfirmRequest(postApiSeatsConfirmRequest).XTenantId(xTenantId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SeatsAPI.PostApiSeatsConfirm``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostApiSeatsConfirmRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **idempotencyKey** | **string** | Required key (max 255 chars after trim); durable fallback identity when externalOrderId is omitted. | 
 **postApiSeatsConfirmRequest** | [**PostApiSeatsConfirmRequest**](PostApiSeatsConfirmRequest.md) |  | 
 **xTenantId** | **string** | Ignored for bootstrap env keys (bound to DEFAULT_TENANT_ID). Dashboard &#x60;sk_live_…&#x60; keys are tenant-scoped by hash. | 

### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json, multipart/form-data, text/plain
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostApiSeatsHoldsRelease

> PostApiSeatsHoldsRelease(ctx).PostApiSeatsHoldsRenewRequest(postApiSeatsHoldsRenewRequest).IdempotencyKey(idempotencyKey).Execute()

Release all seats held by a hold token



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
	postApiSeatsHoldsRenewRequest := *openapiclient.NewPostApiSeatsHoldsRenewRequest("HoldToken_example") // PostApiSeatsHoldsRenewRequest | 
	idempotencyKey := "idempotencyKey_example" // string | Required on transactional routes. Reuse the same key for every retry of one logical operation; the API replays the stored outcome. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SeatsAPI.PostApiSeatsHoldsRelease(context.Background()).PostApiSeatsHoldsRenewRequest(postApiSeatsHoldsRenewRequest).IdempotencyKey(idempotencyKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SeatsAPI.PostApiSeatsHoldsRelease``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostApiSeatsHoldsReleaseRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **postApiSeatsHoldsRenewRequest** | [**PostApiSeatsHoldsRenewRequest**](PostApiSeatsHoldsRenewRequest.md) |  | 
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


## PostApiSeatsHoldsRenew

> PostApiSeatsHoldsRenew(ctx).PostApiSeatsHoldsRenewRequest(postApiSeatsHoldsRenewRequest).IdempotencyKey(idempotencyKey).Execute()

Renew Redis TTL for all seats in a hold session



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
	postApiSeatsHoldsRenewRequest := *openapiclient.NewPostApiSeatsHoldsRenewRequest("HoldToken_example") // PostApiSeatsHoldsRenewRequest | 
	idempotencyKey := "idempotencyKey_example" // string | Required on transactional routes. Reuse the same key for every retry of one logical operation; the API replays the stored outcome. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SeatsAPI.PostApiSeatsHoldsRenew(context.Background()).PostApiSeatsHoldsRenewRequest(postApiSeatsHoldsRenewRequest).IdempotencyKey(idempotencyKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SeatsAPI.PostApiSeatsHoldsRenew``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostApiSeatsHoldsRenewRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **postApiSeatsHoldsRenewRequest** | [**PostApiSeatsHoldsRenewRequest**](PostApiSeatsHoldsRenewRequest.md) |  | 
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


## PostApiSeatsRelease

> PostApiSeatsRelease(ctx).IdempotencyKey(idempotencyKey).PostApiSeatsReleaseRequest(postApiSeatsReleaseRequest).XTenantId(xTenantId).Execute()

Clear Redis TTL holds before expiry (checkout abandoned / declined)



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
	idempotencyKey := "idempotencyKey_example" // string | Required deduplication key (max 255 chars after trim); 24h cache replay in Redis per tenant + route.
	postApiSeatsReleaseRequest := *openapiclient.NewPostApiSeatsReleaseRequest("EventId_example", []string{"SeatIds_example"}, "UserId_example") // PostApiSeatsReleaseRequest | 
	xTenantId := "xTenantId_example" // string | Tenant id for partner integrations. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SeatsAPI.PostApiSeatsRelease(context.Background()).IdempotencyKey(idempotencyKey).PostApiSeatsReleaseRequest(postApiSeatsReleaseRequest).XTenantId(xTenantId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SeatsAPI.PostApiSeatsRelease``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostApiSeatsReleaseRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **idempotencyKey** | **string** | Required deduplication key (max 255 chars after trim); 24h cache replay in Redis per tenant + route. | 
 **postApiSeatsReleaseRequest** | [**PostApiSeatsReleaseRequest**](PostApiSeatsReleaseRequest.md) |  | 
 **xTenantId** | **string** | Tenant id for partner integrations. | 

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


## PostApiSeatsReserve

> PostApiSeatsReserve(ctx).IdempotencyKey(idempotencyKey).PostApiSeatsReserveRequest(postApiSeatsReserveRequest).XTenantId(xTenantId).Execute()

Place a temporary Redis hold (Lua TTL) before payment



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
	idempotencyKey := "idempotencyKey_example" // string | Required deduplication key (max 255 chars after trim; UUID v4 recommended). 24h idempotent cache replay in Redis per tenant.
	postApiSeatsReserveRequest := *openapiclient.NewPostApiSeatsReserveRequest("EventId_example", []string{"SeatIds_example"}, "UserId_example") // PostApiSeatsReserveRequest | 
	xTenantId := "xTenantId_example" // string | Tenant id for partner integrations. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SeatsAPI.PostApiSeatsReserve(context.Background()).IdempotencyKey(idempotencyKey).PostApiSeatsReserveRequest(postApiSeatsReserveRequest).XTenantId(xTenantId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SeatsAPI.PostApiSeatsReserve``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostApiSeatsReserveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **idempotencyKey** | **string** | Required deduplication key (max 255 chars after trim; UUID v4 recommended). 24h idempotent cache replay in Redis per tenant. | 
 **postApiSeatsReserveRequest** | [**PostApiSeatsReserveRequest**](PostApiSeatsReserveRequest.md) |  | 
 **xTenantId** | **string** | Tenant id for partner integrations. | 

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

