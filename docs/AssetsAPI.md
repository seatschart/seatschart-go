# \AssetsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAssetsMapByTenantIdByEventIdByFilename**](AssetsAPI.md#GetAssetsMapByTenantIdByEventIdByFilename) | **Get** /assets/map/{tenantId}/{eventId}/{filename} | Serve persisted map reference/background images



## GetAssetsMapByTenantIdByEventIdByFilename

> GetAssetsMapByTenantIdByEventIdByFilename(ctx, tenantId, eventId, filename).Execute()

Serve persisted map reference/background images



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
	tenantId := "tenantId_example" // string | 
	eventId := "eventId_example" // string | 
	filename := "filename_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.AssetsAPI.GetAssetsMapByTenantIdByEventIdByFilename(context.Background(), tenantId, eventId, filename).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AssetsAPI.GetAssetsMapByTenantIdByEventIdByFilename``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tenantId** | **string** |  | 
**eventId** | **string** |  | 
**filename** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAssetsMapByTenantIdByEventIdByFilenameRequest struct via the builder pattern


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

