# \UsageAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetApiMonitoringRequests**](UsageAPI.md#GetApiMonitoringRequests) | **Get** /api/monitoring/requests | Request monitoring per rate-limit lane, with in-flight (concurrent) requests
[**GetApiUsage**](UsageAPI.md#GetApiUsage) | **Get** /api/usage | Monthly usage: used objects, charts and events
[**GetApiUsageCharts**](UsageAPI.md#GetApiUsageCharts) | **Get** /api/usage/charts | Monthly used objects per chart and event
[**GetApiUsageEvents**](UsageAPI.md#GetApiUsageEvents) | **Get** /api/usage/events | Monthly used objects per event (paginated)
[**GetApiUsageEventsByEventId**](UsageAPI.md#GetApiUsageEventsByEventId) | **Get** /api/usage/events/{eventId} | Monthly used objects of one event (paginated)
[**GetApiUsageHistory**](UsageAPI.md#GetApiUsageHistory) | **Get** /api/usage/history | Used objects per month over a range
[**GetApiUsageWorkspaces**](UsageAPI.md#GetApiUsageWorkspaces) | **Get** /api/usage/workspaces | Monthly used objects across my workspaces (admin view)



## GetApiMonitoringRequests

> GetApiMonitoringRequests200Response GetApiMonitoringRequests(ctx).Range_(range_).Execute()

Request monitoring per rate-limit lane, with in-flight (concurrent) requests



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
	range_ := "range__example" // string | Time range (default `1h`). (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UsageAPI.GetApiMonitoringRequests(context.Background()).Range_(range_).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsageAPI.GetApiMonitoringRequests``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetApiMonitoringRequests`: GetApiMonitoringRequests200Response
	fmt.Fprintf(os.Stdout, "Response from `UsageAPI.GetApiMonitoringRequests`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetApiMonitoringRequestsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **range_** | **string** | Time range (default &#x60;1h&#x60;). | 

### Return type

[**GetApiMonitoringRequests200Response**](GetApiMonitoringRequests200Response.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, multipart/form-data, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetApiUsage

> GetApiUsage(ctx).Month(month).Execute()

Monthly usage: used objects, charts and events



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
	month := "month_example" // string | UTC calendar month `YYYY-MM` (default: current month). (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.UsageAPI.GetApiUsage(context.Background()).Month(month).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsageAPI.GetApiUsage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetApiUsageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **month** | **string** | UTC calendar month &#x60;YYYY-MM&#x60; (default: current month). | 

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


## GetApiUsageCharts

> GetApiUsageCharts(ctx).Month(month).Execute()

Monthly used objects per chart and event



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
	month := "month_example" // string | UTC calendar month `YYYY-MM` (default: current month). (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.UsageAPI.GetApiUsageCharts(context.Background()).Month(month).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsageAPI.GetApiUsageCharts``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetApiUsageChartsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **month** | **string** | UTC calendar month &#x60;YYYY-MM&#x60; (default: current month). | 

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


## GetApiUsageEvents

> GetApiUsageEvents(ctx).Month(month).Limit(limit).Cursor(cursor).Execute()

Monthly used objects per event (paginated)



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
	month := "month_example" // string | UTC calendar month `YYYY-MM` (default: current month). (optional)
	limit := float32(8.14) // float32 | Page size (1..100, default 50). (optional)
	cursor := "cursor_example" // string | Opaque `nextCursor` of the previous page. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.UsageAPI.GetApiUsageEvents(context.Background()).Month(month).Limit(limit).Cursor(cursor).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsageAPI.GetApiUsageEvents``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetApiUsageEventsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **month** | **string** | UTC calendar month &#x60;YYYY-MM&#x60; (default: current month). | 
 **limit** | **float32** | Page size (1..100, default 50). | 
 **cursor** | **string** | Opaque &#x60;nextCursor&#x60; of the previous page. | 

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


## GetApiUsageEventsByEventId

> GetApiUsageEventsByEventId(ctx, eventId).Month(month).Limit(limit).Cursor(cursor).Execute()

Monthly used objects of one event (paginated)



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
	month := "month_example" // string | UTC calendar month `YYYY-MM` (default: current month). (optional)
	limit := float32(8.14) // float32 | Page size (1..1000, default 1000). (optional)
	cursor := "cursor_example" // string | Opaque `nextCursor` of the previous page. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.UsageAPI.GetApiUsageEventsByEventId(context.Background(), eventId).Month(month).Limit(limit).Cursor(cursor).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsageAPI.GetApiUsageEventsByEventId``: %v\n", err)
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

Other parameters are passed through a pointer to a apiGetApiUsageEventsByEventIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **month** | **string** | UTC calendar month &#x60;YYYY-MM&#x60; (default: current month). | 
 **limit** | **float32** | Page size (1..1000, default 1000). | 
 **cursor** | **string** | Opaque &#x60;nextCursor&#x60; of the previous page. | 

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


## GetApiUsageHistory

> GetApiUsageHistory(ctx).From(from).To(to).Execute()

Used objects per month over a range



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
	from := "from_example" // string | First month `YYYY-MM` (default: 11 months before `to`). (optional)
	to := "to_example" // string | Last month `YYYY-MM` (default: current month; not in the future). (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.UsageAPI.GetApiUsageHistory(context.Background()).From(from).To(to).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsageAPI.GetApiUsageHistory``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetApiUsageHistoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **from** | **string** | First month &#x60;YYYY-MM&#x60; (default: 11 months before &#x60;to&#x60;). | 
 **to** | **string** | Last month &#x60;YYYY-MM&#x60; (default: current month; not in the future). | 

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


## GetApiUsageWorkspaces

> GetApiUsageWorkspaces(ctx).Month(month).Execute()

Monthly used objects across my workspaces (admin view)



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
	month := "month_example" // string | UTC calendar month `YYYY-MM` (default: current month). (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.UsageAPI.GetApiUsageWorkspaces(context.Background()).Month(month).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsageAPI.GetApiUsageWorkspaces``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetApiUsageWorkspacesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **month** | **string** | UTC calendar month &#x60;YYYY-MM&#x60; (default: current month). | 

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

