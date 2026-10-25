# \DefaultAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteApiChartsByChartIdCategoriesByCategoryId**](DefaultAPI.md#DeleteApiChartsByChartIdCategoriesByCategoryId) | **Delete** /api/charts/{chartId}/categories/{categoryId} | 
[**GetApiEventsByEventIdGeometryByHash**](DefaultAPI.md#GetApiEventsByEventIdGeometryByHash) | **Get** /api/events/{eventId}/geometry/{hash} | 
[**GetApiEventsByEventIdManifest**](DefaultAPI.md#GetApiEventsByEventIdManifest) | **Get** /api/events/{eventId}/manifest | 
[**GetHealthLive**](DefaultAPI.md#GetHealthLive) | **Get** /health/live | 
[**GetInternalMetrics**](DefaultAPI.md#GetInternalMetrics) | **Get** /internal/metrics | 
[**PatchApiChartsByChartIdCategoriesByCategoryId**](DefaultAPI.md#PatchApiChartsByChartIdCategoriesByCategoryId) | **Patch** /api/charts/{chartId}/categories/{categoryId} | 
[**PostApiChartsByChartIdActionsScanChart**](DefaultAPI.md#PostApiChartsByChartIdActionsScanChart) | **Post** /api/charts/{chartId}/actions/scan-chart | 
[**PostApiChartsByChartIdCategories**](DefaultAPI.md#PostApiChartsByChartIdCategories) | **Post** /api/charts/{chartId}/categories | 
[**PostApiChartsByChartIdVersionsByVersionActionsCopy**](DefaultAPI.md#PostApiChartsByChartIdVersionsByVersionActionsCopy) | **Post** /api/charts/{chartId}/versions/{version}/actions/copy | 
[**PostApiEventsByEventIdActionsScanChart**](DefaultAPI.md#PostApiEventsByEventIdActionsScanChart) | **Post** /api/events/{eventId}/actions/scan-chart | 



## DeleteApiChartsByChartIdCategoriesByCategoryId

> DeleteApiChartsByChartIdCategoriesByCategoryId(ctx, chartId, categoryId).ReassignTo(reassignTo).IdempotencyKey(idempotencyKey).Execute()



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
	categoryId := "categoryId_example" // string | 
	reassignTo := "reassignTo_example" // string | Category id that receives the seats/GA areas of the deleted category. (optional)
	idempotencyKey := "idempotencyKey_example" // string | Required on transactional routes. Reuse the same key for every retry of one logical operation; the API replays the stored outcome. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DefaultAPI.DeleteApiChartsByChartIdCategoriesByCategoryId(context.Background(), chartId, categoryId).ReassignTo(reassignTo).IdempotencyKey(idempotencyKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.DeleteApiChartsByChartIdCategoriesByCategoryId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**chartId** | **string** |  | 
**categoryId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteApiChartsByChartIdCategoriesByCategoryIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **reassignTo** | **string** | Category id that receives the seats/GA areas of the deleted category. | 
 **idempotencyKey** | **string** | Required on transactional routes. Reuse the same key for every retry of one logical operation; the API replays the stored outcome. | 

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


## GetApiEventsByEventIdGeometryByHash

> GetApiEventsByEventIdGeometryByHash(ctx, eventId, hash).Execute()



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
	hash := "hash_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DefaultAPI.GetApiEventsByEventIdGeometryByHash(context.Background(), eventId, hash).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.GetApiEventsByEventIdGeometryByHash``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**eventId** | **string** |  | 
**hash** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetApiEventsByEventIdGeometryByHashRequest struct via the builder pattern


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


## GetApiEventsByEventIdManifest

> GetApiEventsByEventIdManifest(ctx, eventId).Execute()



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DefaultAPI.GetApiEventsByEventIdManifest(context.Background(), eventId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.GetApiEventsByEventIdManifest``: %v\n", err)
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

Other parameters are passed through a pointer to a apiGetApiEventsByEventIdManifestRequest struct via the builder pattern


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


## GetHealthLive

> GetHealthLive(ctx).Execute()



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
	r, err := apiClient.DefaultAPI.GetHealthLive(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.GetHealthLive``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetHealthLiveRequest struct via the builder pattern


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


## GetInternalMetrics

> GetInternalMetrics(ctx).Execute()



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
	r, err := apiClient.DefaultAPI.GetInternalMetrics(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.GetInternalMetrics``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetInternalMetricsRequest struct via the builder pattern


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


## PatchApiChartsByChartIdCategoriesByCategoryId

> PatchApiChartsByChartIdCategoriesByCategoryId(ctx, chartId, categoryId).PatchApiChartsByChartIdCategoriesByCategoryIdRequest(patchApiChartsByChartIdCategoriesByCategoryIdRequest).IdempotencyKey(idempotencyKey).Execute()



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
	categoryId := "categoryId_example" // string | 
	patchApiChartsByChartIdCategoriesByCategoryIdRequest := *openapiclient.NewPatchApiChartsByChartIdCategoriesByCategoryIdRequest() // PatchApiChartsByChartIdCategoriesByCategoryIdRequest | 
	idempotencyKey := "idempotencyKey_example" // string | Required on transactional routes. Reuse the same key for every retry of one logical operation; the API replays the stored outcome. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DefaultAPI.PatchApiChartsByChartIdCategoriesByCategoryId(context.Background(), chartId, categoryId).PatchApiChartsByChartIdCategoriesByCategoryIdRequest(patchApiChartsByChartIdCategoriesByCategoryIdRequest).IdempotencyKey(idempotencyKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.PatchApiChartsByChartIdCategoriesByCategoryId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**chartId** | **string** |  | 
**categoryId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchApiChartsByChartIdCategoriesByCategoryIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **patchApiChartsByChartIdCategoriesByCategoryIdRequest** | [**PatchApiChartsByChartIdCategoriesByCategoryIdRequest**](PatchApiChartsByChartIdCategoriesByCategoryIdRequest.md) |  | 
 **idempotencyKey** | **string** | Required on transactional routes. Reuse the same key for every retry of one logical operation; the API replays the stored outcome. | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, multipart/form-data, text/plain
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostApiChartsByChartIdActionsScanChart

> PostApiChartsByChartIdActionsScanChart(ctx, chartId).IdempotencyKey(idempotencyKey).Execute()



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
	idempotencyKey := "idempotencyKey_example" // string | Required on transactional routes. Reuse the same key for every retry of one logical operation; the API replays the stored outcome. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DefaultAPI.PostApiChartsByChartIdActionsScanChart(context.Background(), chartId).IdempotencyKey(idempotencyKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.PostApiChartsByChartIdActionsScanChart``: %v\n", err)
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

Other parameters are passed through a pointer to a apiPostApiChartsByChartIdActionsScanChartRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **idempotencyKey** | **string** | Required on transactional routes. Reuse the same key for every retry of one logical operation; the API replays the stored outcome. | 

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


## PostApiChartsByChartIdCategories

> PostApiChartsByChartIdCategories(ctx, chartId).PostApiChartsByChartIdCategoriesRequest(postApiChartsByChartIdCategoriesRequest).IdempotencyKey(idempotencyKey).Execute()



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
	postApiChartsByChartIdCategoriesRequest := *openapiclient.NewPostApiChartsByChartIdCategoriesRequest("Name_example") // PostApiChartsByChartIdCategoriesRequest | 
	idempotencyKey := "idempotencyKey_example" // string | Required on transactional routes. Reuse the same key for every retry of one logical operation; the API replays the stored outcome. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DefaultAPI.PostApiChartsByChartIdCategories(context.Background(), chartId).PostApiChartsByChartIdCategoriesRequest(postApiChartsByChartIdCategoriesRequest).IdempotencyKey(idempotencyKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.PostApiChartsByChartIdCategories``: %v\n", err)
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

Other parameters are passed through a pointer to a apiPostApiChartsByChartIdCategoriesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **postApiChartsByChartIdCategoriesRequest** | [**PostApiChartsByChartIdCategoriesRequest**](PostApiChartsByChartIdCategoriesRequest.md) |  | 
 **idempotencyKey** | **string** | Required on transactional routes. Reuse the same key for every retry of one logical operation; the API replays the stored outcome. | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, multipart/form-data, text/plain
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostApiChartsByChartIdVersionsByVersionActionsCopy

> PostApiChartsByChartIdVersionsByVersionActionsCopy(ctx, chartId, version).PostApiChartsByChartIdVersionsByVersionActionsCopyRequest(postApiChartsByChartIdVersionsByVersionActionsCopyRequest).IdempotencyKey(idempotencyKey).Execute()



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
	version := float32(8.14) // float32 | 
	postApiChartsByChartIdVersionsByVersionActionsCopyRequest := *openapiclient.NewPostApiChartsByChartIdVersionsByVersionActionsCopyRequest() // PostApiChartsByChartIdVersionsByVersionActionsCopyRequest | 
	idempotencyKey := "idempotencyKey_example" // string | Required on transactional routes. Reuse the same key for every retry of one logical operation; the API replays the stored outcome. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DefaultAPI.PostApiChartsByChartIdVersionsByVersionActionsCopy(context.Background(), chartId, version).PostApiChartsByChartIdVersionsByVersionActionsCopyRequest(postApiChartsByChartIdVersionsByVersionActionsCopyRequest).IdempotencyKey(idempotencyKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.PostApiChartsByChartIdVersionsByVersionActionsCopy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**chartId** | **string** |  | 
**version** | **float32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostApiChartsByChartIdVersionsByVersionActionsCopyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **postApiChartsByChartIdVersionsByVersionActionsCopyRequest** | [**PostApiChartsByChartIdVersionsByVersionActionsCopyRequest**](PostApiChartsByChartIdVersionsByVersionActionsCopyRequest.md) |  | 
 **idempotencyKey** | **string** | Required on transactional routes. Reuse the same key for every retry of one logical operation; the API replays the stored outcome. | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, multipart/form-data, text/plain
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostApiEventsByEventIdActionsScanChart

> PostApiEventsByEventIdActionsScanChart(ctx, eventId).IdempotencyKey(idempotencyKey).Execute()



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
	idempotencyKey := "idempotencyKey_example" // string | Required on transactional routes. Reuse the same key for every retry of one logical operation; the API replays the stored outcome. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DefaultAPI.PostApiEventsByEventIdActionsScanChart(context.Background(), eventId).IdempotencyKey(idempotencyKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.PostApiEventsByEventIdActionsScanChart``: %v\n", err)
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

Other parameters are passed through a pointer to a apiPostApiEventsByEventIdActionsScanChartRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **idempotencyKey** | **string** | Required on transactional routes. Reuse the same key for every retry of one logical operation; the API replays the stored outcome. | 

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

