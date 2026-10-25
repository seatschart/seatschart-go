# \QuotesAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PostApiQuotes**](QuotesAPI.md#PostApiQuotes) | **Post** /api/quotes | Freeze a server-calculated catalog quote before external payment



## PostApiQuotes

> PostApiQuotes(ctx).IdempotencyKey(idempotencyKey).PostApiQuotesRequest(postApiQuotesRequest).Execute()

Freeze a server-calculated catalog quote before external payment



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
	idempotencyKey := "idempotencyKey_example" // string | 
	postApiQuotesRequest := *openapiclient.NewPostApiQuotesRequest("EventId_example", "UserId_example", "HoldToken_example") // PostApiQuotesRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.QuotesAPI.PostApiQuotes(context.Background()).IdempotencyKey(idempotencyKey).PostApiQuotesRequest(postApiQuotesRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `QuotesAPI.PostApiQuotes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostApiQuotesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **idempotencyKey** | **string** |  | 
 **postApiQuotesRequest** | [**PostApiQuotesRequest**](PostApiQuotesRequest.md) |  | 

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

