# \PublicAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetApiPublicByPublicKeyEventsByEventIdGeometryByHash**](PublicAPI.md#GetApiPublicByPublicKeyEventsByEventIdGeometryByHash) | **Get** /api/public/{publicKey}/events/{eventId}/geometry/{hash} | Public immutable geometry (CDN-cacheable)
[**GetApiPublicByPublicKeyEventsByEventIdManifest**](PublicAPI.md#GetApiPublicByPublicKeyEventsByEventIdManifest) | **Get** /api/public/{publicKey}/events/{eventId}/manifest | Public geometry manifest (CDN-cacheable)
[**GetApiPublicByPublicKeyEventsByEventIdState**](PublicAPI.md#GetApiPublicByPublicKeyEventsByEventIdState) | **Get** /api/public/{publicKey}/events/{eventId}/state | Public compact visual state (CDN-cacheable)
[**GetApiPublicChartPreviewsByChartIdGeometryByHash**](PublicAPI.md#GetApiPublicChartPreviewsByChartIdGeometryByHash) | **Get** /api/public/chart-previews/{chartId}/geometry/{hash} | Chart preview geometry (signed link)
[**GetApiPublicChartPreviewsByChartIdManifest**](PublicAPI.md#GetApiPublicChartPreviewsByChartIdManifest) | **Get** /api/public/chart-previews/{chartId}/manifest | Chart preview manifest (signed link)



## GetApiPublicByPublicKeyEventsByEventIdGeometryByHash

> GetApiPublicByPublicKeyEventsByEventIdGeometryByHash(ctx, publicKey, eventId, hash).Execute()

Public immutable geometry (CDN-cacheable)



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
	publicKey := "publicKey_example" // string | Workspace public key (`pub_…`).
	eventId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	hash := "hash_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.PublicAPI.GetApiPublicByPublicKeyEventsByEventIdGeometryByHash(context.Background(), publicKey, eventId, hash).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PublicAPI.GetApiPublicByPublicKeyEventsByEventIdGeometryByHash``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**publicKey** | **string** | Workspace public key (&#x60;pub_…&#x60;). | 
**eventId** | **string** |  | 
**hash** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetApiPublicByPublicKeyEventsByEventIdGeometryByHashRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetApiPublicByPublicKeyEventsByEventIdManifest

> GetApiPublicByPublicKeyEventsByEventIdManifest(ctx, publicKey, eventId).Execute()

Public geometry manifest (CDN-cacheable)



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
	publicKey := "publicKey_example" // string | Workspace public key (`pub_…`).
	eventId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.PublicAPI.GetApiPublicByPublicKeyEventsByEventIdManifest(context.Background(), publicKey, eventId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PublicAPI.GetApiPublicByPublicKeyEventsByEventIdManifest``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**publicKey** | **string** | Workspace public key (&#x60;pub_…&#x60;). | 
**eventId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetApiPublicByPublicKeyEventsByEventIdManifestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetApiPublicByPublicKeyEventsByEventIdState

> GetApiPublicByPublicKeyEventsByEventIdState(ctx, publicKey, eventId).Execute()

Public compact visual state (CDN-cacheable)



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
	publicKey := "publicKey_example" // string | Workspace public key (`pub_…`).
	eventId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.PublicAPI.GetApiPublicByPublicKeyEventsByEventIdState(context.Background(), publicKey, eventId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PublicAPI.GetApiPublicByPublicKeyEventsByEventIdState``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**publicKey** | **string** | Workspace public key (&#x60;pub_…&#x60;). | 
**eventId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetApiPublicByPublicKeyEventsByEventIdStateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetApiPublicChartPreviewsByChartIdGeometryByHash

> GetApiPublicChartPreviewsByChartIdGeometryByHash(ctx, chartId, hash).Exp(exp).Sig(sig).Execute()

Chart preview geometry (signed link)



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
	chartId := "chartId_example" // string | 
	hash := "hash_example" // string | 
	exp := "exp_example" // string | Expiry (unix seconds) from the link.
	sig := "sig_example" // string | Signature from the link.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.PublicAPI.GetApiPublicChartPreviewsByChartIdGeometryByHash(context.Background(), chartId, hash).Exp(exp).Sig(sig).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PublicAPI.GetApiPublicChartPreviewsByChartIdGeometryByHash``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**chartId** | **string** |  | 
**hash** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetApiPublicChartPreviewsByChartIdGeometryByHashRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **exp** | **string** | Expiry (unix seconds) from the link. | 
 **sig** | **string** | Signature from the link. | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetApiPublicChartPreviewsByChartIdManifest

> GetApiPublicChartPreviewsByChartIdManifest(ctx, chartId).Exp(exp).Sig(sig).Execute()

Chart preview manifest (signed link)



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
	chartId := "chartId_example" // string | 
	exp := "exp_example" // string | Expiry (unix seconds) from the link.
	sig := "sig_example" // string | Signature from the link.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.PublicAPI.GetApiPublicChartPreviewsByChartIdManifest(context.Background(), chartId).Exp(exp).Sig(sig).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PublicAPI.GetApiPublicChartPreviewsByChartIdManifest``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**chartId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetApiPublicChartPreviewsByChartIdManifestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **exp** | **string** | Expiry (unix seconds) from the link. | 
 **sig** | **string** | Signature from the link. | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

