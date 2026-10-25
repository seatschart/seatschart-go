# \EventLogAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetApiEventLog**](EventLogAPI.md#GetApiEventLog) | **Get** /api/event-log | Event log: admin changes of charts and events (cursor-paginated)



## GetApiEventLog

> GetApiEventLog(ctx).Limit(limit).Cursor(cursor).Types(types).From(from).To(to).ResourceId(resourceId).Execute()

Event log: admin changes of charts and events (cursor-paginated)



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
	limit := float32(8.14) // float32 | Page size (1..100, default 20). (optional)
	cursor := "cursor_example" // string | Opaque `nextCursor` value from the previous page. (optional)
	types := "types_example" // string | Comma-separated event log types (default: all). (optional)
	from := "from_example" // string | Inclusive lower bound (ISO-8601 instant). (optional)
	to := "to_example" // string | Exclusive upper bound (ISO-8601 instant). (optional)
	resourceId := "resourceId_example" // string | Only entries of this chart or event id. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.EventLogAPI.GetApiEventLog(context.Background()).Limit(limit).Cursor(cursor).Types(types).From(from).To(to).ResourceId(resourceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EventLogAPI.GetApiEventLog``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetApiEventLogRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **float32** | Page size (1..100, default 20). | 
 **cursor** | **string** | Opaque &#x60;nextCursor&#x60; value from the previous page. | 
 **types** | **string** | Comma-separated event log types (default: all). | 
 **from** | **string** | Inclusive lower bound (ISO-8601 instant). | 
 **to** | **string** | Exclusive upper bound (ISO-8601 instant). | 
 **resourceId** | **string** | Only entries of this chart or event id. | 

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

