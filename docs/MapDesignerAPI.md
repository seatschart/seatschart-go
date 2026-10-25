# \MapDesignerAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PostApiExtractSeats**](MapDesignerAPI.md#PostApiExtractSeats) | **Post** /api/extract-seats | Extract seat candidates with OpenCV



## PostApiExtractSeats

> PostApiExtractSeats(ctx).PostApiExtractSeatsRequest(postApiExtractSeatsRequest).SeatDiameter(seatDiameter).ClosingKernel(closingKernel).CannyLow(cannyLow).CannyHigh(cannyHigh).IdempotencyKey(idempotencyKey).Execute()

Extract seat candidates with OpenCV



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
	postApiExtractSeatsRequest := *openapiclient.NewPostApiExtractSeatsRequest("TODO") // PostApiExtractSeatsRequest | 
	seatDiameter := float32(8.14) // float32 |  (optional)
	closingKernel := "closingKernel_example" // string |  (optional)
	cannyLow := float32(8.14) // float32 |  (optional)
	cannyHigh := float32(8.14) // float32 |  (optional)
	idempotencyKey := "idempotencyKey_example" // string | Required on transactional routes. Reuse the same key for every retry of one logical operation; the API replays the stored outcome. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.MapDesignerAPI.PostApiExtractSeats(context.Background()).PostApiExtractSeatsRequest(postApiExtractSeatsRequest).SeatDiameter(seatDiameter).ClosingKernel(closingKernel).CannyLow(cannyLow).CannyHigh(cannyHigh).IdempotencyKey(idempotencyKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MapDesignerAPI.PostApiExtractSeats``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostApiExtractSeatsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **postApiExtractSeatsRequest** | [**PostApiExtractSeatsRequest**](PostApiExtractSeatsRequest.md) |  | 
 **seatDiameter** | **float32** |  | 
 **closingKernel** | **string** |  | 
 **cannyLow** | **float32** |  | 
 **cannyHigh** | **float32** |  | 
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

