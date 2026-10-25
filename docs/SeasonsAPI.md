# \SeasonsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteApiSeasonsBySeasonId**](SeasonsAPI.md#DeleteApiSeasonsBySeasonId) | **Delete** /api/seasons/{seasonId} | Delete season
[**DeleteApiSeasonsBySeasonIdEventsByEventId**](SeasonsAPI.md#DeleteApiSeasonsBySeasonIdEventsByEventId) | **Delete** /api/seasons/{seasonId}/events/{eventId} | Detach event from season
[**DeleteApiSeasonsBySeasonIdPartialSeasonsByPartialSeasonId**](SeasonsAPI.md#DeleteApiSeasonsBySeasonIdPartialSeasonsByPartialSeasonId) | **Delete** /api/seasons/{seasonId}/partial-seasons/{partialSeasonId} | Delete a partial season
[**GetApiEventsByEventIdSeasonStatus**](SeasonsAPI.md#GetApiEventsByEventIdSeasonStatus) | **Get** /api/events/{eventId}/season-status | Season membership and status overrides of an event
[**GetApiSeasons**](SeasonsAPI.md#GetApiSeasons) | **Get** /api/seasons | List seasons
[**GetApiSeasonsBySeasonId**](SeasonsAPI.md#GetApiSeasonsBySeasonId) | **Get** /api/seasons/{seasonId} | Get season detail
[**GetApiSeasonsBySeasonIdPartialSeasons**](SeasonsAPI.md#GetApiSeasonsBySeasonIdPartialSeasons) | **Get** /api/seasons/{seasonId}/partial-seasons | List partial seasons of a season
[**GetApiSeasonsBySeasonIdPartialSeasonsByPartialSeasonId**](SeasonsAPI.md#GetApiSeasonsBySeasonIdPartialSeasonsByPartialSeasonId) | **Get** /api/seasons/{seasonId}/partial-seasons/{partialSeasonId} | Get one partial season
[**GetApiSeasonsBySeasonIdPartialSeasonsByPartialSeasonIdReportsByAvailabilityCsv**](SeasonsAPI.md#GetApiSeasonsBySeasonIdPartialSeasonsByPartialSeasonIdReportsByAvailabilityCsv) | **Get** /api/seasons/{seasonId}/partial-seasons/{partialSeasonId}/reports/by-availability.csv | Partial season report by availability (CSV)
[**GetApiSeasonsBySeasonIdPartialSeasonsByPartialSeasonIdReportsByAvailabilityJson**](SeasonsAPI.md#GetApiSeasonsBySeasonIdPartialSeasonsByPartialSeasonIdReportsByAvailabilityJson) | **Get** /api/seasons/{seasonId}/partial-seasons/{partialSeasonId}/reports/by-availability.json | Partial season report by availability
[**GetApiSeasonsBySeasonIdPartialSeasonsByPartialSeasonIdReportsByCategoryCsv**](SeasonsAPI.md#GetApiSeasonsBySeasonIdPartialSeasonsByPartialSeasonIdReportsByCategoryCsv) | **Get** /api/seasons/{seasonId}/partial-seasons/{partialSeasonId}/reports/by-category.csv | Partial season report by category (CSV)
[**GetApiSeasonsBySeasonIdPartialSeasonsByPartialSeasonIdReportsByCategoryJson**](SeasonsAPI.md#GetApiSeasonsBySeasonIdPartialSeasonsByPartialSeasonIdReportsByCategoryJson) | **Get** /api/seasons/{seasonId}/partial-seasons/{partialSeasonId}/reports/by-category.json | Partial season report by category
[**GetApiSeasonsBySeasonIdPartialSeasonsByPartialSeasonIdReportsByStatusCsv**](SeasonsAPI.md#GetApiSeasonsBySeasonIdPartialSeasonsByPartialSeasonIdReportsByStatusCsv) | **Get** /api/seasons/{seasonId}/partial-seasons/{partialSeasonId}/reports/by-status.csv | Partial season report by status (CSV)
[**GetApiSeasonsBySeasonIdPartialSeasonsByPartialSeasonIdReportsByStatusJson**](SeasonsAPI.md#GetApiSeasonsBySeasonIdPartialSeasonsByPartialSeasonIdReportsByStatusJson) | **Get** /api/seasons/{seasonId}/partial-seasons/{partialSeasonId}/reports/by-status.json | Partial season report by status
[**GetApiSeasonsBySeasonIdPartialSeasonsByPartialSeasonIdReportsByZoneCsv**](SeasonsAPI.md#GetApiSeasonsBySeasonIdPartialSeasonsByPartialSeasonIdReportsByZoneCsv) | **Get** /api/seasons/{seasonId}/partial-seasons/{partialSeasonId}/reports/by-zone.csv | Partial season report by zone (CSV)
[**GetApiSeasonsBySeasonIdPartialSeasonsByPartialSeasonIdReportsByZoneJson**](SeasonsAPI.md#GetApiSeasonsBySeasonIdPartialSeasonsByPartialSeasonIdReportsByZoneJson) | **Get** /api/seasons/{seasonId}/partial-seasons/{partialSeasonId}/reports/by-zone.json | Partial season report by zone
[**GetApiSeasonsBySeasonIdPartialSeasonsByPartialSeasonIdReportsSummaryByAvailabilityCsv**](SeasonsAPI.md#GetApiSeasonsBySeasonIdPartialSeasonsByPartialSeasonIdReportsSummaryByAvailabilityCsv) | **Get** /api/seasons/{seasonId}/partial-seasons/{partialSeasonId}/reports/summary/by-availability.csv | Partial season summary by availability (CSV)
[**GetApiSeasonsBySeasonIdPartialSeasonsByPartialSeasonIdReportsSummaryByAvailabilityJson**](SeasonsAPI.md#GetApiSeasonsBySeasonIdPartialSeasonsByPartialSeasonIdReportsSummaryByAvailabilityJson) | **Get** /api/seasons/{seasonId}/partial-seasons/{partialSeasonId}/reports/summary/by-availability.json | Partial season summary by availability
[**GetApiSeasonsBySeasonIdPartialSeasonsByPartialSeasonIdReportsSummaryByAvailabilityReasonCsv**](SeasonsAPI.md#GetApiSeasonsBySeasonIdPartialSeasonsByPartialSeasonIdReportsSummaryByAvailabilityReasonCsv) | **Get** /api/seasons/{seasonId}/partial-seasons/{partialSeasonId}/reports/summary/by-availability-reason.csv | Partial season summary by availability reason (CSV)
[**GetApiSeasonsBySeasonIdPartialSeasonsByPartialSeasonIdReportsSummaryByAvailabilityReasonJson**](SeasonsAPI.md#GetApiSeasonsBySeasonIdPartialSeasonsByPartialSeasonIdReportsSummaryByAvailabilityReasonJson) | **Get** /api/seasons/{seasonId}/partial-seasons/{partialSeasonId}/reports/summary/by-availability-reason.json | Partial season summary by availability reason
[**GetApiSeasonsBySeasonIdPartialSeasonsByPartialSeasonIdReportsSummaryByCategoryCsv**](SeasonsAPI.md#GetApiSeasonsBySeasonIdPartialSeasonsByPartialSeasonIdReportsSummaryByCategoryCsv) | **Get** /api/seasons/{seasonId}/partial-seasons/{partialSeasonId}/reports/summary/by-category.csv | Partial season summary by category (CSV)
[**GetApiSeasonsBySeasonIdPartialSeasonsByPartialSeasonIdReportsSummaryByCategoryJson**](SeasonsAPI.md#GetApiSeasonsBySeasonIdPartialSeasonsByPartialSeasonIdReportsSummaryByCategoryJson) | **Get** /api/seasons/{seasonId}/partial-seasons/{partialSeasonId}/reports/summary/by-category.json | Partial season summary by category
[**GetApiSeasonsBySeasonIdPartialSeasonsByPartialSeasonIdReportsSummaryBySectionCsv**](SeasonsAPI.md#GetApiSeasonsBySeasonIdPartialSeasonsByPartialSeasonIdReportsSummaryBySectionCsv) | **Get** /api/seasons/{seasonId}/partial-seasons/{partialSeasonId}/reports/summary/by-section.csv | Partial season summary by section (CSV)
[**GetApiSeasonsBySeasonIdPartialSeasonsByPartialSeasonIdReportsSummaryBySectionJson**](SeasonsAPI.md#GetApiSeasonsBySeasonIdPartialSeasonsByPartialSeasonIdReportsSummaryBySectionJson) | **Get** /api/seasons/{seasonId}/partial-seasons/{partialSeasonId}/reports/summary/by-section.json | Partial season summary by section
[**GetApiSeasonsBySeasonIdPartialSeasonsByPartialSeasonIdReportsSummaryByStatusCsv**](SeasonsAPI.md#GetApiSeasonsBySeasonIdPartialSeasonsByPartialSeasonIdReportsSummaryByStatusCsv) | **Get** /api/seasons/{seasonId}/partial-seasons/{partialSeasonId}/reports/summary/by-status.csv | Partial season summary by status (CSV)
[**GetApiSeasonsBySeasonIdPartialSeasonsByPartialSeasonIdReportsSummaryByStatusJson**](SeasonsAPI.md#GetApiSeasonsBySeasonIdPartialSeasonsByPartialSeasonIdReportsSummaryByStatusJson) | **Get** /api/seasons/{seasonId}/partial-seasons/{partialSeasonId}/reports/summary/by-status.json | Partial season summary by status
[**GetApiSeasonsBySeasonIdPartialSeasonsByPartialSeasonIdReportsSummaryByZoneCsv**](SeasonsAPI.md#GetApiSeasonsBySeasonIdPartialSeasonsByPartialSeasonIdReportsSummaryByZoneCsv) | **Get** /api/seasons/{seasonId}/partial-seasons/{partialSeasonId}/reports/summary/by-zone.csv | Partial season summary by zone (CSV)
[**GetApiSeasonsBySeasonIdPartialSeasonsByPartialSeasonIdReportsSummaryByZoneJson**](SeasonsAPI.md#GetApiSeasonsBySeasonIdPartialSeasonsByPartialSeasonIdReportsSummaryByZoneJson) | **Get** /api/seasons/{seasonId}/partial-seasons/{partialSeasonId}/reports/summary/by-zone.json | Partial season summary by zone
[**GetApiSeasonsBySeasonIdPartialSeasonsByPartialSeasonIdReportsSummaryCsv**](SeasonsAPI.md#GetApiSeasonsBySeasonIdPartialSeasonsByPartialSeasonIdReportsSummaryCsv) | **Get** /api/seasons/{seasonId}/partial-seasons/{partialSeasonId}/reports/summary.csv | Partial season summary across member events (CSV)
[**GetApiSeasonsBySeasonIdPartialSeasonsByPartialSeasonIdReportsSummaryJson**](SeasonsAPI.md#GetApiSeasonsBySeasonIdPartialSeasonsByPartialSeasonIdReportsSummaryJson) | **Get** /api/seasons/{seasonId}/partial-seasons/{partialSeasonId}/reports/summary.json | Partial season summary across member events
[**GetApiSeasonsBySeasonIdReportsByAvailabilityCsv**](SeasonsAPI.md#GetApiSeasonsBySeasonIdReportsByAvailabilityCsv) | **Get** /api/seasons/{seasonId}/reports/by-availability.csv | Season report by availability (CSV)
[**GetApiSeasonsBySeasonIdReportsByAvailabilityJson**](SeasonsAPI.md#GetApiSeasonsBySeasonIdReportsByAvailabilityJson) | **Get** /api/seasons/{seasonId}/reports/by-availability.json | Season report by availability
[**GetApiSeasonsBySeasonIdReportsByCategoryCsv**](SeasonsAPI.md#GetApiSeasonsBySeasonIdReportsByCategoryCsv) | **Get** /api/seasons/{seasonId}/reports/by-category.csv | Season report by category (CSV)
[**GetApiSeasonsBySeasonIdReportsByCategoryJson**](SeasonsAPI.md#GetApiSeasonsBySeasonIdReportsByCategoryJson) | **Get** /api/seasons/{seasonId}/reports/by-category.json | Season report by category
[**GetApiSeasonsBySeasonIdReportsByStatusCsv**](SeasonsAPI.md#GetApiSeasonsBySeasonIdReportsByStatusCsv) | **Get** /api/seasons/{seasonId}/reports/by-status.csv | Season report by status (CSV)
[**GetApiSeasonsBySeasonIdReportsByStatusJson**](SeasonsAPI.md#GetApiSeasonsBySeasonIdReportsByStatusJson) | **Get** /api/seasons/{seasonId}/reports/by-status.json | Season report by status
[**GetApiSeasonsBySeasonIdReportsByZoneCsv**](SeasonsAPI.md#GetApiSeasonsBySeasonIdReportsByZoneCsv) | **Get** /api/seasons/{seasonId}/reports/by-zone.csv | Season report by zone (CSV)
[**GetApiSeasonsBySeasonIdReportsByZoneJson**](SeasonsAPI.md#GetApiSeasonsBySeasonIdReportsByZoneJson) | **Get** /api/seasons/{seasonId}/reports/by-zone.json | Season report by zone
[**GetApiSeasonsBySeasonIdReportsSummaryByAvailabilityCsv**](SeasonsAPI.md#GetApiSeasonsBySeasonIdReportsSummaryByAvailabilityCsv) | **Get** /api/seasons/{seasonId}/reports/summary/by-availability.csv | Season summary by availability (CSV)
[**GetApiSeasonsBySeasonIdReportsSummaryByAvailabilityJson**](SeasonsAPI.md#GetApiSeasonsBySeasonIdReportsSummaryByAvailabilityJson) | **Get** /api/seasons/{seasonId}/reports/summary/by-availability.json | Season summary by availability
[**GetApiSeasonsBySeasonIdReportsSummaryByAvailabilityReasonCsv**](SeasonsAPI.md#GetApiSeasonsBySeasonIdReportsSummaryByAvailabilityReasonCsv) | **Get** /api/seasons/{seasonId}/reports/summary/by-availability-reason.csv | Season summary by availability reason (CSV)
[**GetApiSeasonsBySeasonIdReportsSummaryByAvailabilityReasonJson**](SeasonsAPI.md#GetApiSeasonsBySeasonIdReportsSummaryByAvailabilityReasonJson) | **Get** /api/seasons/{seasonId}/reports/summary/by-availability-reason.json | Season summary by availability reason
[**GetApiSeasonsBySeasonIdReportsSummaryByCategoryCsv**](SeasonsAPI.md#GetApiSeasonsBySeasonIdReportsSummaryByCategoryCsv) | **Get** /api/seasons/{seasonId}/reports/summary/by-category.csv | Season summary by category (CSV)
[**GetApiSeasonsBySeasonIdReportsSummaryByCategoryJson**](SeasonsAPI.md#GetApiSeasonsBySeasonIdReportsSummaryByCategoryJson) | **Get** /api/seasons/{seasonId}/reports/summary/by-category.json | Season summary by category
[**GetApiSeasonsBySeasonIdReportsSummaryBySectionCsv**](SeasonsAPI.md#GetApiSeasonsBySeasonIdReportsSummaryBySectionCsv) | **Get** /api/seasons/{seasonId}/reports/summary/by-section.csv | Season summary by section (CSV)
[**GetApiSeasonsBySeasonIdReportsSummaryBySectionJson**](SeasonsAPI.md#GetApiSeasonsBySeasonIdReportsSummaryBySectionJson) | **Get** /api/seasons/{seasonId}/reports/summary/by-section.json | Season summary by section
[**GetApiSeasonsBySeasonIdReportsSummaryByStatusCsv**](SeasonsAPI.md#GetApiSeasonsBySeasonIdReportsSummaryByStatusCsv) | **Get** /api/seasons/{seasonId}/reports/summary/by-status.csv | Season summary by status (CSV)
[**GetApiSeasonsBySeasonIdReportsSummaryByStatusJson**](SeasonsAPI.md#GetApiSeasonsBySeasonIdReportsSummaryByStatusJson) | **Get** /api/seasons/{seasonId}/reports/summary/by-status.json | Season summary by status
[**GetApiSeasonsBySeasonIdReportsSummaryByZoneCsv**](SeasonsAPI.md#GetApiSeasonsBySeasonIdReportsSummaryByZoneCsv) | **Get** /api/seasons/{seasonId}/reports/summary/by-zone.csv | Season summary by zone (CSV)
[**GetApiSeasonsBySeasonIdReportsSummaryByZoneJson**](SeasonsAPI.md#GetApiSeasonsBySeasonIdReportsSummaryByZoneJson) | **Get** /api/seasons/{seasonId}/reports/summary/by-zone.json | Season summary by zone
[**GetApiSeasonsBySeasonIdReportsSummaryCsv**](SeasonsAPI.md#GetApiSeasonsBySeasonIdReportsSummaryCsv) | **Get** /api/seasons/{seasonId}/reports/summary.csv | Season summary across member events (CSV)
[**GetApiSeasonsBySeasonIdReportsSummaryJson**](SeasonsAPI.md#GetApiSeasonsBySeasonIdReportsSummaryJson) | **Get** /api/seasons/{seasonId}/reports/summary.json | Season summary across member events
[**GetApiSeasonsBySeasonIdSeats**](SeasonsAPI.md#GetApiSeasonsBySeasonIdSeats) | **Get** /api/seasons/{seasonId}/seats | Season booking view (source geometry + intersection status)
[**PatchApiSeasonsBySeasonId**](SeasonsAPI.md#PatchApiSeasonsBySeasonId) | **Patch** /api/seasons/{seasonId} | Update season
[**PatchApiSeasonsBySeasonIdEventsByEventId**](SeasonsAPI.md#PatchApiSeasonsBySeasonIdEventsByEventId) | **Patch** /api/seasons/{seasonId}/events/{eventId} | Season member settings (forSalePropagated)
[**PatchApiSeasonsBySeasonIdPartialSeasonsByPartialSeasonId**](SeasonsAPI.md#PatchApiSeasonsBySeasonIdPartialSeasonsByPartialSeasonId) | **Patch** /api/seasons/{seasonId}/partial-seasons/{partialSeasonId} | Rename a partial season or set its key
[**PostApiEventsByEventIdActionsOverrideSeasonStatus**](SeasonsAPI.md#PostApiEventsByEventIdActionsOverrideSeasonStatus) | **Post** /api/events/{eventId}/actions/override-season-status | Override the season status of objects in an event (OVERRIDE_SEASON_STATUS)
[**PostApiEventsByEventIdActionsUseSeasonStatus**](SeasonsAPI.md#PostApiEventsByEventIdActionsUseSeasonStatus) | **Post** /api/events/{eventId}/actions/use-season-status | Make objects follow the season status again (USE_SEASON_STATUS)
[**PostApiSeasons**](SeasonsAPI.md#PostApiSeasons) | **Post** /api/seasons | Create season
[**PostApiSeasonsBySeasonIdActionsAddEvents**](SeasonsAPI.md#PostApiSeasonsBySeasonIdActionsAddEvents) | **Post** /api/seasons/{seasonId}/actions/add-events | Add several events to a season (atomic)
[**PostApiSeasonsBySeasonIdActionsCreateEvents**](SeasonsAPI.md#PostApiSeasonsBySeasonIdActionsCreateEvents) | **Post** /api/seasons/{seasonId}/actions/create-events | Create events in a season
[**PostApiSeasonsBySeasonIdActionsRemoveEvents**](SeasonsAPI.md#PostApiSeasonsBySeasonIdActionsRemoveEvents) | **Post** /api/seasons/{seasonId}/actions/remove-events | Remove several events from a season (atomic)
[**PostApiSeasonsBySeasonIdEvents**](SeasonsAPI.md#PostApiSeasonsBySeasonIdEvents) | **Post** /api/seasons/{seasonId}/events | Attach event to season
[**PostApiSeasonsBySeasonIdInventorySync**](SeasonsAPI.md#PostApiSeasonsBySeasonIdInventorySync) | **Post** /api/seasons/{seasonId}/inventory/sync | Sync season inventory from source
[**PostApiSeasonsBySeasonIdPartialSeasons**](SeasonsAPI.md#PostApiSeasonsBySeasonIdPartialSeasons) | **Post** /api/seasons/{seasonId}/partial-seasons | Create a partial season
[**PostApiSeasonsBySeasonIdPartialSeasonsByPartialSeasonIdActionsAddEvents**](SeasonsAPI.md#PostApiSeasonsBySeasonIdPartialSeasonsByPartialSeasonIdActionsAddEvents) | **Post** /api/seasons/{seasonId}/partial-seasons/{partialSeasonId}/actions/add-events | Add events to a partial season (bulk, all-or-nothing)
[**PostApiSeasonsBySeasonIdPartialSeasonsByPartialSeasonIdActionsRemoveEvents**](SeasonsAPI.md#PostApiSeasonsBySeasonIdPartialSeasonsByPartialSeasonIdActionsRemoveEvents) | **Post** /api/seasons/{seasonId}/partial-seasons/{partialSeasonId}/actions/remove-events | Remove events from a partial season (bulk)



## DeleteApiSeasonsBySeasonId

> DeleteApiSeasonsBySeasonId(ctx, seasonId).IdempotencyKey(idempotencyKey).Execute()

Delete season



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
	seasonId := "seasonId_example" // string | 
	idempotencyKey := "idempotencyKey_example" // string | Required on transactional routes. Reuse the same key for every retry of one logical operation; the API replays the stored outcome. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SeasonsAPI.DeleteApiSeasonsBySeasonId(context.Background(), seasonId).IdempotencyKey(idempotencyKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SeasonsAPI.DeleteApiSeasonsBySeasonId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**seasonId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteApiSeasonsBySeasonIdRequest struct via the builder pattern


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


## DeleteApiSeasonsBySeasonIdEventsByEventId

> DeleteApiSeasonsBySeasonIdEventsByEventId(ctx, seasonId, eventId).IdempotencyKey(idempotencyKey).Execute()

Detach event from season



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
	seasonId := "seasonId_example" // string | 
	eventId := "eventId_example" // string | 
	idempotencyKey := "idempotencyKey_example" // string | Required on transactional routes. Reuse the same key for every retry of one logical operation; the API replays the stored outcome. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SeasonsAPI.DeleteApiSeasonsBySeasonIdEventsByEventId(context.Background(), seasonId, eventId).IdempotencyKey(idempotencyKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SeasonsAPI.DeleteApiSeasonsBySeasonIdEventsByEventId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**seasonId** | **string** |  | 
**eventId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteApiSeasonsBySeasonIdEventsByEventIdRequest struct via the builder pattern


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


## DeleteApiSeasonsBySeasonIdPartialSeasonsByPartialSeasonId

> DeleteApiSeasonsBySeasonIdPartialSeasonsByPartialSeasonId(ctx, seasonId, partialSeasonId).IdempotencyKey(idempotencyKey).Execute()

Delete a partial season



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
	seasonId := "seasonId_example" // string | 
	partialSeasonId := "partialSeasonId_example" // string | 
	idempotencyKey := "idempotencyKey_example" // string | Required on transactional routes. Reuse the same key for every retry of one logical operation; the API replays the stored outcome. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SeasonsAPI.DeleteApiSeasonsBySeasonIdPartialSeasonsByPartialSeasonId(context.Background(), seasonId, partialSeasonId).IdempotencyKey(idempotencyKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SeasonsAPI.DeleteApiSeasonsBySeasonIdPartialSeasonsByPartialSeasonId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**seasonId** | **string** |  | 
**partialSeasonId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteApiSeasonsBySeasonIdPartialSeasonsByPartialSeasonIdRequest struct via the builder pattern


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


## GetApiEventsByEventIdSeasonStatus

> GetApiEventsByEventIdSeasonStatus(ctx, eventId).Execute()

Season membership and status overrides of an event



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SeasonsAPI.GetApiEventsByEventIdSeasonStatus(context.Background(), eventId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SeasonsAPI.GetApiEventsByEventIdSeasonStatus``: %v\n", err)
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

Other parameters are passed through a pointer to a apiGetApiEventsByEventIdSeasonStatusRequest struct via the builder pattern


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


## GetApiSeasons

> GetApiSeasons(ctx).Execute()

List seasons



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
	r, err := apiClient.SeasonsAPI.GetApiSeasons(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SeasonsAPI.GetApiSeasons``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetApiSeasonsRequest struct via the builder pattern


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


## GetApiSeasonsBySeasonId

> GetApiSeasonsBySeasonId(ctx, seasonId).Execute()

Get season detail

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
	seasonId := "seasonId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SeasonsAPI.GetApiSeasonsBySeasonId(context.Background(), seasonId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SeasonsAPI.GetApiSeasonsBySeasonId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**seasonId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetApiSeasonsBySeasonIdRequest struct via the builder pattern


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


## GetApiSeasonsBySeasonIdPartialSeasons

> GetApiSeasonsBySeasonIdPartialSeasons(ctx, seasonId).Execute()

List partial seasons of a season



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
	seasonId := "seasonId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SeasonsAPI.GetApiSeasonsBySeasonIdPartialSeasons(context.Background(), seasonId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SeasonsAPI.GetApiSeasonsBySeasonIdPartialSeasons``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**seasonId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetApiSeasonsBySeasonIdPartialSeasonsRequest struct via the builder pattern


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


## GetApiSeasonsBySeasonIdPartialSeasonsByPartialSeasonId

> GetApiSeasonsBySeasonIdPartialSeasonsByPartialSeasonId(ctx, seasonId, partialSeasonId).Execute()

Get one partial season



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
	seasonId := "seasonId_example" // string | 
	partialSeasonId := "partialSeasonId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SeasonsAPI.GetApiSeasonsBySeasonIdPartialSeasonsByPartialSeasonId(context.Background(), seasonId, partialSeasonId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SeasonsAPI.GetApiSeasonsBySeasonIdPartialSeasonsByPartialSeasonId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**seasonId** | **string** |  | 
**partialSeasonId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetApiSeasonsBySeasonIdPartialSeasonsByPartialSeasonIdRequest struct via the builder pattern


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


## GetApiSeasonsBySeasonIdPartialSeasonsByPartialSeasonIdReportsByAvailabilityCsv

> GetApiSeasonsBySeasonIdPartialSeasonsByPartialSeasonIdReportsByAvailabilityCsv(ctx, seasonId, partialSeasonId).Availability(availability).LiveHolds(liveHolds).Execute()

Partial season report by availability (CSV)



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
	seasonId := "seasonId_example" // string | 
	partialSeasonId := "partialSeasonId_example" // string | 
	availability := "availability_example" // string | Only this availability bucket. (optional)
	liveHolds := "liveHolds_example" // string | Merge active buyer holds from the per-event hold index (default `true`). `false` = Postgres only. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SeasonsAPI.GetApiSeasonsBySeasonIdPartialSeasonsByPartialSeasonIdReportsByAvailabilityCsv(context.Background(), seasonId, partialSeasonId).Availability(availability).LiveHolds(liveHolds).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SeasonsAPI.GetApiSeasonsBySeasonIdPartialSeasonsByPartialSeasonIdReportsByAvailabilityCsv``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**seasonId** | **string** |  | 
**partialSeasonId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetApiSeasonsBySeasonIdPartialSeasonsByPartialSeasonIdReportsByAvailabilityCsvRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **availability** | **string** | Only this availability bucket. | 
 **liveHolds** | **string** | Merge active buyer holds from the per-event hold index (default &#x60;true&#x60;). &#x60;false&#x60; &#x3D; Postgres only. | 

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


## GetApiSeasonsBySeasonIdPartialSeasonsByPartialSeasonIdReportsByAvailabilityJson

> GetApiSeasonsBySeasonIdPartialSeasonsByPartialSeasonIdReportsByAvailabilityJson(ctx, seasonId, partialSeasonId).Availability(availability).LiveHolds(liveHolds).Execute()

Partial season report by availability



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
	seasonId := "seasonId_example" // string | 
	partialSeasonId := "partialSeasonId_example" // string | 
	availability := "availability_example" // string | Only this availability bucket. (optional)
	liveHolds := "liveHolds_example" // string | Merge active buyer holds from the per-event hold index (default `true`). `false` = Postgres only. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SeasonsAPI.GetApiSeasonsBySeasonIdPartialSeasonsByPartialSeasonIdReportsByAvailabilityJson(context.Background(), seasonId, partialSeasonId).Availability(availability).LiveHolds(liveHolds).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SeasonsAPI.GetApiSeasonsBySeasonIdPartialSeasonsByPartialSeasonIdReportsByAvailabilityJson``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**seasonId** | **string** |  | 
**partialSeasonId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetApiSeasonsBySeasonIdPartialSeasonsByPartialSeasonIdReportsByAvailabilityJsonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **availability** | **string** | Only this availability bucket. | 
 **liveHolds** | **string** | Merge active buyer holds from the per-event hold index (default &#x60;true&#x60;). &#x60;false&#x60; &#x3D; Postgres only. | 

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


## GetApiSeasonsBySeasonIdPartialSeasonsByPartialSeasonIdReportsByCategoryCsv

> GetApiSeasonsBySeasonIdPartialSeasonsByPartialSeasonIdReportsByCategoryCsv(ctx, seasonId, partialSeasonId).Category(category).LiveHolds(liveHolds).Execute()

Partial season report by category (CSV)



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
	seasonId := "seasonId_example" // string | 
	partialSeasonId := "partialSeasonId_example" // string | 
	category := "category_example" // string | Only this category bucket. (optional)
	liveHolds := "liveHolds_example" // string | Merge active buyer holds from the per-event hold index (default `true`). `false` = Postgres only. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SeasonsAPI.GetApiSeasonsBySeasonIdPartialSeasonsByPartialSeasonIdReportsByCategoryCsv(context.Background(), seasonId, partialSeasonId).Category(category).LiveHolds(liveHolds).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SeasonsAPI.GetApiSeasonsBySeasonIdPartialSeasonsByPartialSeasonIdReportsByCategoryCsv``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**seasonId** | **string** |  | 
**partialSeasonId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetApiSeasonsBySeasonIdPartialSeasonsByPartialSeasonIdReportsByCategoryCsvRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **category** | **string** | Only this category bucket. | 
 **liveHolds** | **string** | Merge active buyer holds from the per-event hold index (default &#x60;true&#x60;). &#x60;false&#x60; &#x3D; Postgres only. | 

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


## GetApiSeasonsBySeasonIdPartialSeasonsByPartialSeasonIdReportsByCategoryJson

> GetApiSeasonsBySeasonIdPartialSeasonsByPartialSeasonIdReportsByCategoryJson(ctx, seasonId, partialSeasonId).Category(category).LiveHolds(liveHolds).Execute()

Partial season report by category



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
	seasonId := "seasonId_example" // string | 
	partialSeasonId := "partialSeasonId_example" // string | 
	category := "category_example" // string | Only this category bucket. (optional)
	liveHolds := "liveHolds_example" // string | Merge active buyer holds from the per-event hold index (default `true`). `false` = Postgres only. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SeasonsAPI.GetApiSeasonsBySeasonIdPartialSeasonsByPartialSeasonIdReportsByCategoryJson(context.Background(), seasonId, partialSeasonId).Category(category).LiveHolds(liveHolds).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SeasonsAPI.GetApiSeasonsBySeasonIdPartialSeasonsByPartialSeasonIdReportsByCategoryJson``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**seasonId** | **string** |  | 
**partialSeasonId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetApiSeasonsBySeasonIdPartialSeasonsByPartialSeasonIdReportsByCategoryJsonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **category** | **string** | Only this category bucket. | 
 **liveHolds** | **string** | Merge active buyer holds from the per-event hold index (default &#x60;true&#x60;). &#x60;false&#x60; &#x3D; Postgres only. | 

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


## GetApiSeasonsBySeasonIdPartialSeasonsByPartialSeasonIdReportsByStatusCsv

> GetApiSeasonsBySeasonIdPartialSeasonsByPartialSeasonIdReportsByStatusCsv(ctx, seasonId, partialSeasonId).Status(status).LiveHolds(liveHolds).Execute()

Partial season report by status (CSV)



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
	seasonId := "seasonId_example" // string | 
	partialSeasonId := "partialSeasonId_example" // string | 
	status := "status_example" // string | Only this status bucket. (optional)
	liveHolds := "liveHolds_example" // string | Merge active buyer holds from the per-event hold index (default `true`). `false` = Postgres only. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SeasonsAPI.GetApiSeasonsBySeasonIdPartialSeasonsByPartialSeasonIdReportsByStatusCsv(context.Background(), seasonId, partialSeasonId).Status(status).LiveHolds(liveHolds).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SeasonsAPI.GetApiSeasonsBySeasonIdPartialSeasonsByPartialSeasonIdReportsByStatusCsv``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**seasonId** | **string** |  | 
**partialSeasonId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetApiSeasonsBySeasonIdPartialSeasonsByPartialSeasonIdReportsByStatusCsvRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **status** | **string** | Only this status bucket. | 
 **liveHolds** | **string** | Merge active buyer holds from the per-event hold index (default &#x60;true&#x60;). &#x60;false&#x60; &#x3D; Postgres only. | 

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


## GetApiSeasonsBySeasonIdPartialSeasonsByPartialSeasonIdReportsByStatusJson

> GetApiSeasonsBySeasonIdPartialSeasonsByPartialSeasonIdReportsByStatusJson(ctx, seasonId, partialSeasonId).Status(status).LiveHolds(liveHolds).Execute()

Partial season report by status



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
	seasonId := "seasonId_example" // string | 
	partialSeasonId := "partialSeasonId_example" // string | 
	status := "status_example" // string | Only this status bucket. (optional)
	liveHolds := "liveHolds_example" // string | Merge active buyer holds from the per-event hold index (default `true`). `false` = Postgres only. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SeasonsAPI.GetApiSeasonsBySeasonIdPartialSeasonsByPartialSeasonIdReportsByStatusJson(context.Background(), seasonId, partialSeasonId).Status(status).LiveHolds(liveHolds).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SeasonsAPI.GetApiSeasonsBySeasonIdPartialSeasonsByPartialSeasonIdReportsByStatusJson``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**seasonId** | **string** |  | 
**partialSeasonId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetApiSeasonsBySeasonIdPartialSeasonsByPartialSeasonIdReportsByStatusJsonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **status** | **string** | Only this status bucket. | 
 **liveHolds** | **string** | Merge active buyer holds from the per-event hold index (default &#x60;true&#x60;). &#x60;false&#x60; &#x3D; Postgres only. | 

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


## GetApiSeasonsBySeasonIdPartialSeasonsByPartialSeasonIdReportsByZoneCsv

> GetApiSeasonsBySeasonIdPartialSeasonsByPartialSeasonIdReportsByZoneCsv(ctx, seasonId, partialSeasonId).Zone(zone).LiveHolds(liveHolds).Execute()

Partial season report by zone (CSV)



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
	seasonId := "seasonId_example" // string | 
	partialSeasonId := "partialSeasonId_example" // string | 
	zone := "zone_example" // string | Only this zone bucket. (optional)
	liveHolds := "liveHolds_example" // string | Merge active buyer holds from the per-event hold index (default `true`). `false` = Postgres only. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SeasonsAPI.GetApiSeasonsBySeasonIdPartialSeasonsByPartialSeasonIdReportsByZoneCsv(context.Background(), seasonId, partialSeasonId).Zone(zone).LiveHolds(liveHolds).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SeasonsAPI.GetApiSeasonsBySeasonIdPartialSeasonsByPartialSeasonIdReportsByZoneCsv``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**seasonId** | **string** |  | 
**partialSeasonId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetApiSeasonsBySeasonIdPartialSeasonsByPartialSeasonIdReportsByZoneCsvRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **zone** | **string** | Only this zone bucket. | 
 **liveHolds** | **string** | Merge active buyer holds from the per-event hold index (default &#x60;true&#x60;). &#x60;false&#x60; &#x3D; Postgres only. | 

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


## GetApiSeasonsBySeasonIdPartialSeasonsByPartialSeasonIdReportsByZoneJson

> GetApiSeasonsBySeasonIdPartialSeasonsByPartialSeasonIdReportsByZoneJson(ctx, seasonId, partialSeasonId).Zone(zone).LiveHolds(liveHolds).Execute()

Partial season report by zone



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
	seasonId := "seasonId_example" // string | 
	partialSeasonId := "partialSeasonId_example" // string | 
	zone := "zone_example" // string | Only this zone bucket. (optional)
	liveHolds := "liveHolds_example" // string | Merge active buyer holds from the per-event hold index (default `true`). `false` = Postgres only. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SeasonsAPI.GetApiSeasonsBySeasonIdPartialSeasonsByPartialSeasonIdReportsByZoneJson(context.Background(), seasonId, partialSeasonId).Zone(zone).LiveHolds(liveHolds).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SeasonsAPI.GetApiSeasonsBySeasonIdPartialSeasonsByPartialSeasonIdReportsByZoneJson``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**seasonId** | **string** |  | 
**partialSeasonId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetApiSeasonsBySeasonIdPartialSeasonsByPartialSeasonIdReportsByZoneJsonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **zone** | **string** | Only this zone bucket. | 
 **liveHolds** | **string** | Merge active buyer holds from the per-event hold index (default &#x60;true&#x60;). &#x60;false&#x60; &#x3D; Postgres only. | 

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


## GetApiSeasonsBySeasonIdPartialSeasonsByPartialSeasonIdReportsSummaryByAvailabilityCsv

> GetApiSeasonsBySeasonIdPartialSeasonsByPartialSeasonIdReportsSummaryByAvailabilityCsv(ctx, seasonId, partialSeasonId).LiveHolds(liveHolds).Execute()

Partial season summary by availability (CSV)



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
	seasonId := "seasonId_example" // string | 
	partialSeasonId := "partialSeasonId_example" // string | 
	liveHolds := "liveHolds_example" // string | Merge active buyer holds from the per-event hold index (default `true`). `false` = Postgres only. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SeasonsAPI.GetApiSeasonsBySeasonIdPartialSeasonsByPartialSeasonIdReportsSummaryByAvailabilityCsv(context.Background(), seasonId, partialSeasonId).LiveHolds(liveHolds).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SeasonsAPI.GetApiSeasonsBySeasonIdPartialSeasonsByPartialSeasonIdReportsSummaryByAvailabilityCsv``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**seasonId** | **string** |  | 
**partialSeasonId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetApiSeasonsBySeasonIdPartialSeasonsByPartialSeasonIdReportsSummaryByAvailabilityCsvRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **liveHolds** | **string** | Merge active buyer holds from the per-event hold index (default &#x60;true&#x60;). &#x60;false&#x60; &#x3D; Postgres only. | 

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


## GetApiSeasonsBySeasonIdPartialSeasonsByPartialSeasonIdReportsSummaryByAvailabilityJson

> GetApiSeasonsBySeasonIdPartialSeasonsByPartialSeasonIdReportsSummaryByAvailabilityJson(ctx, seasonId, partialSeasonId).LiveHolds(liveHolds).Execute()

Partial season summary by availability



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
	seasonId := "seasonId_example" // string | 
	partialSeasonId := "partialSeasonId_example" // string | 
	liveHolds := "liveHolds_example" // string | Merge active buyer holds from the per-event hold index (default `true`). `false` = Postgres only. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SeasonsAPI.GetApiSeasonsBySeasonIdPartialSeasonsByPartialSeasonIdReportsSummaryByAvailabilityJson(context.Background(), seasonId, partialSeasonId).LiveHolds(liveHolds).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SeasonsAPI.GetApiSeasonsBySeasonIdPartialSeasonsByPartialSeasonIdReportsSummaryByAvailabilityJson``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**seasonId** | **string** |  | 
**partialSeasonId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetApiSeasonsBySeasonIdPartialSeasonsByPartialSeasonIdReportsSummaryByAvailabilityJsonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **liveHolds** | **string** | Merge active buyer holds from the per-event hold index (default &#x60;true&#x60;). &#x60;false&#x60; &#x3D; Postgres only. | 

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


## GetApiSeasonsBySeasonIdPartialSeasonsByPartialSeasonIdReportsSummaryByAvailabilityReasonCsv

> GetApiSeasonsBySeasonIdPartialSeasonsByPartialSeasonIdReportsSummaryByAvailabilityReasonCsv(ctx, seasonId, partialSeasonId).LiveHolds(liveHolds).Execute()

Partial season summary by availability reason (CSV)



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
	seasonId := "seasonId_example" // string | 
	partialSeasonId := "partialSeasonId_example" // string | 
	liveHolds := "liveHolds_example" // string | Merge active buyer holds from the per-event hold index (default `true`). `false` = Postgres only. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SeasonsAPI.GetApiSeasonsBySeasonIdPartialSeasonsByPartialSeasonIdReportsSummaryByAvailabilityReasonCsv(context.Background(), seasonId, partialSeasonId).LiveHolds(liveHolds).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SeasonsAPI.GetApiSeasonsBySeasonIdPartialSeasonsByPartialSeasonIdReportsSummaryByAvailabilityReasonCsv``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**seasonId** | **string** |  | 
**partialSeasonId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetApiSeasonsBySeasonIdPartialSeasonsByPartialSeasonIdReportsSummaryByAvailabilityReasonCsvRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **liveHolds** | **string** | Merge active buyer holds from the per-event hold index (default &#x60;true&#x60;). &#x60;false&#x60; &#x3D; Postgres only. | 

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


## GetApiSeasonsBySeasonIdPartialSeasonsByPartialSeasonIdReportsSummaryByAvailabilityReasonJson

> GetApiSeasonsBySeasonIdPartialSeasonsByPartialSeasonIdReportsSummaryByAvailabilityReasonJson(ctx, seasonId, partialSeasonId).LiveHolds(liveHolds).Execute()

Partial season summary by availability reason



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
	seasonId := "seasonId_example" // string | 
	partialSeasonId := "partialSeasonId_example" // string | 
	liveHolds := "liveHolds_example" // string | Merge active buyer holds from the per-event hold index (default `true`). `false` = Postgres only. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SeasonsAPI.GetApiSeasonsBySeasonIdPartialSeasonsByPartialSeasonIdReportsSummaryByAvailabilityReasonJson(context.Background(), seasonId, partialSeasonId).LiveHolds(liveHolds).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SeasonsAPI.GetApiSeasonsBySeasonIdPartialSeasonsByPartialSeasonIdReportsSummaryByAvailabilityReasonJson``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**seasonId** | **string** |  | 
**partialSeasonId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetApiSeasonsBySeasonIdPartialSeasonsByPartialSeasonIdReportsSummaryByAvailabilityReasonJsonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **liveHolds** | **string** | Merge active buyer holds from the per-event hold index (default &#x60;true&#x60;). &#x60;false&#x60; &#x3D; Postgres only. | 

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


## GetApiSeasonsBySeasonIdPartialSeasonsByPartialSeasonIdReportsSummaryByCategoryCsv

> GetApiSeasonsBySeasonIdPartialSeasonsByPartialSeasonIdReportsSummaryByCategoryCsv(ctx, seasonId, partialSeasonId).LiveHolds(liveHolds).Execute()

Partial season summary by category (CSV)



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
	seasonId := "seasonId_example" // string | 
	partialSeasonId := "partialSeasonId_example" // string | 
	liveHolds := "liveHolds_example" // string | Merge active buyer holds from the per-event hold index (default `true`). `false` = Postgres only. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SeasonsAPI.GetApiSeasonsBySeasonIdPartialSeasonsByPartialSeasonIdReportsSummaryByCategoryCsv(context.Background(), seasonId, partialSeasonId).LiveHolds(liveHolds).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SeasonsAPI.GetApiSeasonsBySeasonIdPartialSeasonsByPartialSeasonIdReportsSummaryByCategoryCsv``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**seasonId** | **string** |  | 
**partialSeasonId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetApiSeasonsBySeasonIdPartialSeasonsByPartialSeasonIdReportsSummaryByCategoryCsvRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **liveHolds** | **string** | Merge active buyer holds from the per-event hold index (default &#x60;true&#x60;). &#x60;false&#x60; &#x3D; Postgres only. | 

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


## GetApiSeasonsBySeasonIdPartialSeasonsByPartialSeasonIdReportsSummaryByCategoryJson

> GetApiSeasonsBySeasonIdPartialSeasonsByPartialSeasonIdReportsSummaryByCategoryJson(ctx, seasonId, partialSeasonId).LiveHolds(liveHolds).Execute()

Partial season summary by category



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
	seasonId := "seasonId_example" // string | 
	partialSeasonId := "partialSeasonId_example" // string | 
	liveHolds := "liveHolds_example" // string | Merge active buyer holds from the per-event hold index (default `true`). `false` = Postgres only. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SeasonsAPI.GetApiSeasonsBySeasonIdPartialSeasonsByPartialSeasonIdReportsSummaryByCategoryJson(context.Background(), seasonId, partialSeasonId).LiveHolds(liveHolds).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SeasonsAPI.GetApiSeasonsBySeasonIdPartialSeasonsByPartialSeasonIdReportsSummaryByCategoryJson``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**seasonId** | **string** |  | 
**partialSeasonId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetApiSeasonsBySeasonIdPartialSeasonsByPartialSeasonIdReportsSummaryByCategoryJsonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **liveHolds** | **string** | Merge active buyer holds from the per-event hold index (default &#x60;true&#x60;). &#x60;false&#x60; &#x3D; Postgres only. | 

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


## GetApiSeasonsBySeasonIdPartialSeasonsByPartialSeasonIdReportsSummaryBySectionCsv

> GetApiSeasonsBySeasonIdPartialSeasonsByPartialSeasonIdReportsSummaryBySectionCsv(ctx, seasonId, partialSeasonId).LiveHolds(liveHolds).Execute()

Partial season summary by section (CSV)



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
	seasonId := "seasonId_example" // string | 
	partialSeasonId := "partialSeasonId_example" // string | 
	liveHolds := "liveHolds_example" // string | Merge active buyer holds from the per-event hold index (default `true`). `false` = Postgres only. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SeasonsAPI.GetApiSeasonsBySeasonIdPartialSeasonsByPartialSeasonIdReportsSummaryBySectionCsv(context.Background(), seasonId, partialSeasonId).LiveHolds(liveHolds).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SeasonsAPI.GetApiSeasonsBySeasonIdPartialSeasonsByPartialSeasonIdReportsSummaryBySectionCsv``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**seasonId** | **string** |  | 
**partialSeasonId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetApiSeasonsBySeasonIdPartialSeasonsByPartialSeasonIdReportsSummaryBySectionCsvRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **liveHolds** | **string** | Merge active buyer holds from the per-event hold index (default &#x60;true&#x60;). &#x60;false&#x60; &#x3D; Postgres only. | 

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


## GetApiSeasonsBySeasonIdPartialSeasonsByPartialSeasonIdReportsSummaryBySectionJson

> GetApiSeasonsBySeasonIdPartialSeasonsByPartialSeasonIdReportsSummaryBySectionJson(ctx, seasonId, partialSeasonId).LiveHolds(liveHolds).Execute()

Partial season summary by section



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
	seasonId := "seasonId_example" // string | 
	partialSeasonId := "partialSeasonId_example" // string | 
	liveHolds := "liveHolds_example" // string | Merge active buyer holds from the per-event hold index (default `true`). `false` = Postgres only. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SeasonsAPI.GetApiSeasonsBySeasonIdPartialSeasonsByPartialSeasonIdReportsSummaryBySectionJson(context.Background(), seasonId, partialSeasonId).LiveHolds(liveHolds).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SeasonsAPI.GetApiSeasonsBySeasonIdPartialSeasonsByPartialSeasonIdReportsSummaryBySectionJson``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**seasonId** | **string** |  | 
**partialSeasonId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetApiSeasonsBySeasonIdPartialSeasonsByPartialSeasonIdReportsSummaryBySectionJsonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **liveHolds** | **string** | Merge active buyer holds from the per-event hold index (default &#x60;true&#x60;). &#x60;false&#x60; &#x3D; Postgres only. | 

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


## GetApiSeasonsBySeasonIdPartialSeasonsByPartialSeasonIdReportsSummaryByStatusCsv

> GetApiSeasonsBySeasonIdPartialSeasonsByPartialSeasonIdReportsSummaryByStatusCsv(ctx, seasonId, partialSeasonId).LiveHolds(liveHolds).Execute()

Partial season summary by status (CSV)



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
	seasonId := "seasonId_example" // string | 
	partialSeasonId := "partialSeasonId_example" // string | 
	liveHolds := "liveHolds_example" // string | Merge active buyer holds from the per-event hold index (default `true`). `false` = Postgres only. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SeasonsAPI.GetApiSeasonsBySeasonIdPartialSeasonsByPartialSeasonIdReportsSummaryByStatusCsv(context.Background(), seasonId, partialSeasonId).LiveHolds(liveHolds).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SeasonsAPI.GetApiSeasonsBySeasonIdPartialSeasonsByPartialSeasonIdReportsSummaryByStatusCsv``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**seasonId** | **string** |  | 
**partialSeasonId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetApiSeasonsBySeasonIdPartialSeasonsByPartialSeasonIdReportsSummaryByStatusCsvRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **liveHolds** | **string** | Merge active buyer holds from the per-event hold index (default &#x60;true&#x60;). &#x60;false&#x60; &#x3D; Postgres only. | 

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


## GetApiSeasonsBySeasonIdPartialSeasonsByPartialSeasonIdReportsSummaryByStatusJson

> GetApiSeasonsBySeasonIdPartialSeasonsByPartialSeasonIdReportsSummaryByStatusJson(ctx, seasonId, partialSeasonId).LiveHolds(liveHolds).Execute()

Partial season summary by status



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
	seasonId := "seasonId_example" // string | 
	partialSeasonId := "partialSeasonId_example" // string | 
	liveHolds := "liveHolds_example" // string | Merge active buyer holds from the per-event hold index (default `true`). `false` = Postgres only. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SeasonsAPI.GetApiSeasonsBySeasonIdPartialSeasonsByPartialSeasonIdReportsSummaryByStatusJson(context.Background(), seasonId, partialSeasonId).LiveHolds(liveHolds).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SeasonsAPI.GetApiSeasonsBySeasonIdPartialSeasonsByPartialSeasonIdReportsSummaryByStatusJson``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**seasonId** | **string** |  | 
**partialSeasonId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetApiSeasonsBySeasonIdPartialSeasonsByPartialSeasonIdReportsSummaryByStatusJsonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **liveHolds** | **string** | Merge active buyer holds from the per-event hold index (default &#x60;true&#x60;). &#x60;false&#x60; &#x3D; Postgres only. | 

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


## GetApiSeasonsBySeasonIdPartialSeasonsByPartialSeasonIdReportsSummaryByZoneCsv

> GetApiSeasonsBySeasonIdPartialSeasonsByPartialSeasonIdReportsSummaryByZoneCsv(ctx, seasonId, partialSeasonId).LiveHolds(liveHolds).Execute()

Partial season summary by zone (CSV)



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
	seasonId := "seasonId_example" // string | 
	partialSeasonId := "partialSeasonId_example" // string | 
	liveHolds := "liveHolds_example" // string | Merge active buyer holds from the per-event hold index (default `true`). `false` = Postgres only. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SeasonsAPI.GetApiSeasonsBySeasonIdPartialSeasonsByPartialSeasonIdReportsSummaryByZoneCsv(context.Background(), seasonId, partialSeasonId).LiveHolds(liveHolds).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SeasonsAPI.GetApiSeasonsBySeasonIdPartialSeasonsByPartialSeasonIdReportsSummaryByZoneCsv``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**seasonId** | **string** |  | 
**partialSeasonId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetApiSeasonsBySeasonIdPartialSeasonsByPartialSeasonIdReportsSummaryByZoneCsvRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **liveHolds** | **string** | Merge active buyer holds from the per-event hold index (default &#x60;true&#x60;). &#x60;false&#x60; &#x3D; Postgres only. | 

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


## GetApiSeasonsBySeasonIdPartialSeasonsByPartialSeasonIdReportsSummaryByZoneJson

> GetApiSeasonsBySeasonIdPartialSeasonsByPartialSeasonIdReportsSummaryByZoneJson(ctx, seasonId, partialSeasonId).LiveHolds(liveHolds).Execute()

Partial season summary by zone



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
	seasonId := "seasonId_example" // string | 
	partialSeasonId := "partialSeasonId_example" // string | 
	liveHolds := "liveHolds_example" // string | Merge active buyer holds from the per-event hold index (default `true`). `false` = Postgres only. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SeasonsAPI.GetApiSeasonsBySeasonIdPartialSeasonsByPartialSeasonIdReportsSummaryByZoneJson(context.Background(), seasonId, partialSeasonId).LiveHolds(liveHolds).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SeasonsAPI.GetApiSeasonsBySeasonIdPartialSeasonsByPartialSeasonIdReportsSummaryByZoneJson``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**seasonId** | **string** |  | 
**partialSeasonId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetApiSeasonsBySeasonIdPartialSeasonsByPartialSeasonIdReportsSummaryByZoneJsonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **liveHolds** | **string** | Merge active buyer holds from the per-event hold index (default &#x60;true&#x60;). &#x60;false&#x60; &#x3D; Postgres only. | 

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


## GetApiSeasonsBySeasonIdPartialSeasonsByPartialSeasonIdReportsSummaryCsv

> GetApiSeasonsBySeasonIdPartialSeasonsByPartialSeasonIdReportsSummaryCsv(ctx, seasonId, partialSeasonId).LiveHolds(liveHolds).Execute()

Partial season summary across member events (CSV)



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
	seasonId := "seasonId_example" // string | 
	partialSeasonId := "partialSeasonId_example" // string | 
	liveHolds := "liveHolds_example" // string | Merge active buyer holds from the per-event hold index (default `true`). `false` = Postgres only. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SeasonsAPI.GetApiSeasonsBySeasonIdPartialSeasonsByPartialSeasonIdReportsSummaryCsv(context.Background(), seasonId, partialSeasonId).LiveHolds(liveHolds).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SeasonsAPI.GetApiSeasonsBySeasonIdPartialSeasonsByPartialSeasonIdReportsSummaryCsv``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**seasonId** | **string** |  | 
**partialSeasonId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetApiSeasonsBySeasonIdPartialSeasonsByPartialSeasonIdReportsSummaryCsvRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **liveHolds** | **string** | Merge active buyer holds from the per-event hold index (default &#x60;true&#x60;). &#x60;false&#x60; &#x3D; Postgres only. | 

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


## GetApiSeasonsBySeasonIdPartialSeasonsByPartialSeasonIdReportsSummaryJson

> GetApiSeasonsBySeasonIdPartialSeasonsByPartialSeasonIdReportsSummaryJson(ctx, seasonId, partialSeasonId).LiveHolds(liveHolds).Execute()

Partial season summary across member events



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
	seasonId := "seasonId_example" // string | 
	partialSeasonId := "partialSeasonId_example" // string | 
	liveHolds := "liveHolds_example" // string | Merge active buyer holds from the per-event hold index (default `true`). `false` = Postgres only. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SeasonsAPI.GetApiSeasonsBySeasonIdPartialSeasonsByPartialSeasonIdReportsSummaryJson(context.Background(), seasonId, partialSeasonId).LiveHolds(liveHolds).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SeasonsAPI.GetApiSeasonsBySeasonIdPartialSeasonsByPartialSeasonIdReportsSummaryJson``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**seasonId** | **string** |  | 
**partialSeasonId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetApiSeasonsBySeasonIdPartialSeasonsByPartialSeasonIdReportsSummaryJsonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **liveHolds** | **string** | Merge active buyer holds from the per-event hold index (default &#x60;true&#x60;). &#x60;false&#x60; &#x3D; Postgres only. | 

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


## GetApiSeasonsBySeasonIdReportsByAvailabilityCsv

> GetApiSeasonsBySeasonIdReportsByAvailabilityCsv(ctx, seasonId).Availability(availability).LiveHolds(liveHolds).Execute()

Season report by availability (CSV)



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
	seasonId := "seasonId_example" // string | 
	availability := "availability_example" // string | Only this availability bucket. (optional)
	liveHolds := "liveHolds_example" // string | Merge active buyer holds from the per-event hold index (default `true`). `false` = Postgres only. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SeasonsAPI.GetApiSeasonsBySeasonIdReportsByAvailabilityCsv(context.Background(), seasonId).Availability(availability).LiveHolds(liveHolds).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SeasonsAPI.GetApiSeasonsBySeasonIdReportsByAvailabilityCsv``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**seasonId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetApiSeasonsBySeasonIdReportsByAvailabilityCsvRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **availability** | **string** | Only this availability bucket. | 
 **liveHolds** | **string** | Merge active buyer holds from the per-event hold index (default &#x60;true&#x60;). &#x60;false&#x60; &#x3D; Postgres only. | 

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


## GetApiSeasonsBySeasonIdReportsByAvailabilityJson

> GetApiSeasonsBySeasonIdReportsByAvailabilityJson(ctx, seasonId).Availability(availability).LiveHolds(liveHolds).Execute()

Season report by availability



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
	seasonId := "seasonId_example" // string | 
	availability := "availability_example" // string | Only this availability bucket. (optional)
	liveHolds := "liveHolds_example" // string | Merge active buyer holds from the per-event hold index (default `true`). `false` = Postgres only. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SeasonsAPI.GetApiSeasonsBySeasonIdReportsByAvailabilityJson(context.Background(), seasonId).Availability(availability).LiveHolds(liveHolds).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SeasonsAPI.GetApiSeasonsBySeasonIdReportsByAvailabilityJson``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**seasonId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetApiSeasonsBySeasonIdReportsByAvailabilityJsonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **availability** | **string** | Only this availability bucket. | 
 **liveHolds** | **string** | Merge active buyer holds from the per-event hold index (default &#x60;true&#x60;). &#x60;false&#x60; &#x3D; Postgres only. | 

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


## GetApiSeasonsBySeasonIdReportsByCategoryCsv

> GetApiSeasonsBySeasonIdReportsByCategoryCsv(ctx, seasonId).Category(category).LiveHolds(liveHolds).Execute()

Season report by category (CSV)



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
	seasonId := "seasonId_example" // string | 
	category := "category_example" // string | Only this category bucket. (optional)
	liveHolds := "liveHolds_example" // string | Merge active buyer holds from the per-event hold index (default `true`). `false` = Postgres only. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SeasonsAPI.GetApiSeasonsBySeasonIdReportsByCategoryCsv(context.Background(), seasonId).Category(category).LiveHolds(liveHolds).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SeasonsAPI.GetApiSeasonsBySeasonIdReportsByCategoryCsv``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**seasonId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetApiSeasonsBySeasonIdReportsByCategoryCsvRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **category** | **string** | Only this category bucket. | 
 **liveHolds** | **string** | Merge active buyer holds from the per-event hold index (default &#x60;true&#x60;). &#x60;false&#x60; &#x3D; Postgres only. | 

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


## GetApiSeasonsBySeasonIdReportsByCategoryJson

> GetApiSeasonsBySeasonIdReportsByCategoryJson(ctx, seasonId).Category(category).LiveHolds(liveHolds).Execute()

Season report by category



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
	seasonId := "seasonId_example" // string | 
	category := "category_example" // string | Only this category bucket. (optional)
	liveHolds := "liveHolds_example" // string | Merge active buyer holds from the per-event hold index (default `true`). `false` = Postgres only. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SeasonsAPI.GetApiSeasonsBySeasonIdReportsByCategoryJson(context.Background(), seasonId).Category(category).LiveHolds(liveHolds).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SeasonsAPI.GetApiSeasonsBySeasonIdReportsByCategoryJson``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**seasonId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetApiSeasonsBySeasonIdReportsByCategoryJsonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **category** | **string** | Only this category bucket. | 
 **liveHolds** | **string** | Merge active buyer holds from the per-event hold index (default &#x60;true&#x60;). &#x60;false&#x60; &#x3D; Postgres only. | 

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


## GetApiSeasonsBySeasonIdReportsByStatusCsv

> GetApiSeasonsBySeasonIdReportsByStatusCsv(ctx, seasonId).Status(status).LiveHolds(liveHolds).Execute()

Season report by status (CSV)



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
	seasonId := "seasonId_example" // string | 
	status := "status_example" // string | Only this status bucket. (optional)
	liveHolds := "liveHolds_example" // string | Merge active buyer holds from the per-event hold index (default `true`). `false` = Postgres only. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SeasonsAPI.GetApiSeasonsBySeasonIdReportsByStatusCsv(context.Background(), seasonId).Status(status).LiveHolds(liveHolds).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SeasonsAPI.GetApiSeasonsBySeasonIdReportsByStatusCsv``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**seasonId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetApiSeasonsBySeasonIdReportsByStatusCsvRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **status** | **string** | Only this status bucket. | 
 **liveHolds** | **string** | Merge active buyer holds from the per-event hold index (default &#x60;true&#x60;). &#x60;false&#x60; &#x3D; Postgres only. | 

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


## GetApiSeasonsBySeasonIdReportsByStatusJson

> GetApiSeasonsBySeasonIdReportsByStatusJson(ctx, seasonId).Status(status).LiveHolds(liveHolds).Execute()

Season report by status



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
	seasonId := "seasonId_example" // string | 
	status := "status_example" // string | Only this status bucket. (optional)
	liveHolds := "liveHolds_example" // string | Merge active buyer holds from the per-event hold index (default `true`). `false` = Postgres only. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SeasonsAPI.GetApiSeasonsBySeasonIdReportsByStatusJson(context.Background(), seasonId).Status(status).LiveHolds(liveHolds).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SeasonsAPI.GetApiSeasonsBySeasonIdReportsByStatusJson``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**seasonId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetApiSeasonsBySeasonIdReportsByStatusJsonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **status** | **string** | Only this status bucket. | 
 **liveHolds** | **string** | Merge active buyer holds from the per-event hold index (default &#x60;true&#x60;). &#x60;false&#x60; &#x3D; Postgres only. | 

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


## GetApiSeasonsBySeasonIdReportsByZoneCsv

> GetApiSeasonsBySeasonIdReportsByZoneCsv(ctx, seasonId).Zone(zone).LiveHolds(liveHolds).Execute()

Season report by zone (CSV)



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
	seasonId := "seasonId_example" // string | 
	zone := "zone_example" // string | Only this zone bucket. (optional)
	liveHolds := "liveHolds_example" // string | Merge active buyer holds from the per-event hold index (default `true`). `false` = Postgres only. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SeasonsAPI.GetApiSeasonsBySeasonIdReportsByZoneCsv(context.Background(), seasonId).Zone(zone).LiveHolds(liveHolds).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SeasonsAPI.GetApiSeasonsBySeasonIdReportsByZoneCsv``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**seasonId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetApiSeasonsBySeasonIdReportsByZoneCsvRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **zone** | **string** | Only this zone bucket. | 
 **liveHolds** | **string** | Merge active buyer holds from the per-event hold index (default &#x60;true&#x60;). &#x60;false&#x60; &#x3D; Postgres only. | 

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


## GetApiSeasonsBySeasonIdReportsByZoneJson

> GetApiSeasonsBySeasonIdReportsByZoneJson(ctx, seasonId).Zone(zone).LiveHolds(liveHolds).Execute()

Season report by zone



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
	seasonId := "seasonId_example" // string | 
	zone := "zone_example" // string | Only this zone bucket. (optional)
	liveHolds := "liveHolds_example" // string | Merge active buyer holds from the per-event hold index (default `true`). `false` = Postgres only. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SeasonsAPI.GetApiSeasonsBySeasonIdReportsByZoneJson(context.Background(), seasonId).Zone(zone).LiveHolds(liveHolds).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SeasonsAPI.GetApiSeasonsBySeasonIdReportsByZoneJson``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**seasonId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetApiSeasonsBySeasonIdReportsByZoneJsonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **zone** | **string** | Only this zone bucket. | 
 **liveHolds** | **string** | Merge active buyer holds from the per-event hold index (default &#x60;true&#x60;). &#x60;false&#x60; &#x3D; Postgres only. | 

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


## GetApiSeasonsBySeasonIdReportsSummaryByAvailabilityCsv

> GetApiSeasonsBySeasonIdReportsSummaryByAvailabilityCsv(ctx, seasonId).LiveHolds(liveHolds).Execute()

Season summary by availability (CSV)



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
	seasonId := "seasonId_example" // string | 
	liveHolds := "liveHolds_example" // string | Merge active buyer holds from the per-event hold index (default `true`). `false` = Postgres only. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SeasonsAPI.GetApiSeasonsBySeasonIdReportsSummaryByAvailabilityCsv(context.Background(), seasonId).LiveHolds(liveHolds).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SeasonsAPI.GetApiSeasonsBySeasonIdReportsSummaryByAvailabilityCsv``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**seasonId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetApiSeasonsBySeasonIdReportsSummaryByAvailabilityCsvRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **liveHolds** | **string** | Merge active buyer holds from the per-event hold index (default &#x60;true&#x60;). &#x60;false&#x60; &#x3D; Postgres only. | 

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


## GetApiSeasonsBySeasonIdReportsSummaryByAvailabilityJson

> GetApiSeasonsBySeasonIdReportsSummaryByAvailabilityJson(ctx, seasonId).LiveHolds(liveHolds).Execute()

Season summary by availability



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
	seasonId := "seasonId_example" // string | 
	liveHolds := "liveHolds_example" // string | Merge active buyer holds from the per-event hold index (default `true`). `false` = Postgres only. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SeasonsAPI.GetApiSeasonsBySeasonIdReportsSummaryByAvailabilityJson(context.Background(), seasonId).LiveHolds(liveHolds).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SeasonsAPI.GetApiSeasonsBySeasonIdReportsSummaryByAvailabilityJson``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**seasonId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetApiSeasonsBySeasonIdReportsSummaryByAvailabilityJsonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **liveHolds** | **string** | Merge active buyer holds from the per-event hold index (default &#x60;true&#x60;). &#x60;false&#x60; &#x3D; Postgres only. | 

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


## GetApiSeasonsBySeasonIdReportsSummaryByAvailabilityReasonCsv

> GetApiSeasonsBySeasonIdReportsSummaryByAvailabilityReasonCsv(ctx, seasonId).LiveHolds(liveHolds).Execute()

Season summary by availability reason (CSV)



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
	seasonId := "seasonId_example" // string | 
	liveHolds := "liveHolds_example" // string | Merge active buyer holds from the per-event hold index (default `true`). `false` = Postgres only. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SeasonsAPI.GetApiSeasonsBySeasonIdReportsSummaryByAvailabilityReasonCsv(context.Background(), seasonId).LiveHolds(liveHolds).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SeasonsAPI.GetApiSeasonsBySeasonIdReportsSummaryByAvailabilityReasonCsv``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**seasonId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetApiSeasonsBySeasonIdReportsSummaryByAvailabilityReasonCsvRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **liveHolds** | **string** | Merge active buyer holds from the per-event hold index (default &#x60;true&#x60;). &#x60;false&#x60; &#x3D; Postgres only. | 

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


## GetApiSeasonsBySeasonIdReportsSummaryByAvailabilityReasonJson

> GetApiSeasonsBySeasonIdReportsSummaryByAvailabilityReasonJson(ctx, seasonId).LiveHolds(liveHolds).Execute()

Season summary by availability reason



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
	seasonId := "seasonId_example" // string | 
	liveHolds := "liveHolds_example" // string | Merge active buyer holds from the per-event hold index (default `true`). `false` = Postgres only. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SeasonsAPI.GetApiSeasonsBySeasonIdReportsSummaryByAvailabilityReasonJson(context.Background(), seasonId).LiveHolds(liveHolds).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SeasonsAPI.GetApiSeasonsBySeasonIdReportsSummaryByAvailabilityReasonJson``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**seasonId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetApiSeasonsBySeasonIdReportsSummaryByAvailabilityReasonJsonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **liveHolds** | **string** | Merge active buyer holds from the per-event hold index (default &#x60;true&#x60;). &#x60;false&#x60; &#x3D; Postgres only. | 

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


## GetApiSeasonsBySeasonIdReportsSummaryByCategoryCsv

> GetApiSeasonsBySeasonIdReportsSummaryByCategoryCsv(ctx, seasonId).LiveHolds(liveHolds).Execute()

Season summary by category (CSV)



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
	seasonId := "seasonId_example" // string | 
	liveHolds := "liveHolds_example" // string | Merge active buyer holds from the per-event hold index (default `true`). `false` = Postgres only. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SeasonsAPI.GetApiSeasonsBySeasonIdReportsSummaryByCategoryCsv(context.Background(), seasonId).LiveHolds(liveHolds).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SeasonsAPI.GetApiSeasonsBySeasonIdReportsSummaryByCategoryCsv``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**seasonId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetApiSeasonsBySeasonIdReportsSummaryByCategoryCsvRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **liveHolds** | **string** | Merge active buyer holds from the per-event hold index (default &#x60;true&#x60;). &#x60;false&#x60; &#x3D; Postgres only. | 

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


## GetApiSeasonsBySeasonIdReportsSummaryByCategoryJson

> GetApiSeasonsBySeasonIdReportsSummaryByCategoryJson(ctx, seasonId).LiveHolds(liveHolds).Execute()

Season summary by category



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
	seasonId := "seasonId_example" // string | 
	liveHolds := "liveHolds_example" // string | Merge active buyer holds from the per-event hold index (default `true`). `false` = Postgres only. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SeasonsAPI.GetApiSeasonsBySeasonIdReportsSummaryByCategoryJson(context.Background(), seasonId).LiveHolds(liveHolds).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SeasonsAPI.GetApiSeasonsBySeasonIdReportsSummaryByCategoryJson``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**seasonId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetApiSeasonsBySeasonIdReportsSummaryByCategoryJsonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **liveHolds** | **string** | Merge active buyer holds from the per-event hold index (default &#x60;true&#x60;). &#x60;false&#x60; &#x3D; Postgres only. | 

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


## GetApiSeasonsBySeasonIdReportsSummaryBySectionCsv

> GetApiSeasonsBySeasonIdReportsSummaryBySectionCsv(ctx, seasonId).LiveHolds(liveHolds).Execute()

Season summary by section (CSV)



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
	seasonId := "seasonId_example" // string | 
	liveHolds := "liveHolds_example" // string | Merge active buyer holds from the per-event hold index (default `true`). `false` = Postgres only. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SeasonsAPI.GetApiSeasonsBySeasonIdReportsSummaryBySectionCsv(context.Background(), seasonId).LiveHolds(liveHolds).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SeasonsAPI.GetApiSeasonsBySeasonIdReportsSummaryBySectionCsv``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**seasonId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetApiSeasonsBySeasonIdReportsSummaryBySectionCsvRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **liveHolds** | **string** | Merge active buyer holds from the per-event hold index (default &#x60;true&#x60;). &#x60;false&#x60; &#x3D; Postgres only. | 

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


## GetApiSeasonsBySeasonIdReportsSummaryBySectionJson

> GetApiSeasonsBySeasonIdReportsSummaryBySectionJson(ctx, seasonId).LiveHolds(liveHolds).Execute()

Season summary by section



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
	seasonId := "seasonId_example" // string | 
	liveHolds := "liveHolds_example" // string | Merge active buyer holds from the per-event hold index (default `true`). `false` = Postgres only. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SeasonsAPI.GetApiSeasonsBySeasonIdReportsSummaryBySectionJson(context.Background(), seasonId).LiveHolds(liveHolds).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SeasonsAPI.GetApiSeasonsBySeasonIdReportsSummaryBySectionJson``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**seasonId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetApiSeasonsBySeasonIdReportsSummaryBySectionJsonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **liveHolds** | **string** | Merge active buyer holds from the per-event hold index (default &#x60;true&#x60;). &#x60;false&#x60; &#x3D; Postgres only. | 

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


## GetApiSeasonsBySeasonIdReportsSummaryByStatusCsv

> GetApiSeasonsBySeasonIdReportsSummaryByStatusCsv(ctx, seasonId).LiveHolds(liveHolds).Execute()

Season summary by status (CSV)



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
	seasonId := "seasonId_example" // string | 
	liveHolds := "liveHolds_example" // string | Merge active buyer holds from the per-event hold index (default `true`). `false` = Postgres only. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SeasonsAPI.GetApiSeasonsBySeasonIdReportsSummaryByStatusCsv(context.Background(), seasonId).LiveHolds(liveHolds).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SeasonsAPI.GetApiSeasonsBySeasonIdReportsSummaryByStatusCsv``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**seasonId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetApiSeasonsBySeasonIdReportsSummaryByStatusCsvRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **liveHolds** | **string** | Merge active buyer holds from the per-event hold index (default &#x60;true&#x60;). &#x60;false&#x60; &#x3D; Postgres only. | 

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


## GetApiSeasonsBySeasonIdReportsSummaryByStatusJson

> GetApiSeasonsBySeasonIdReportsSummaryByStatusJson(ctx, seasonId).LiveHolds(liveHolds).Execute()

Season summary by status



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
	seasonId := "seasonId_example" // string | 
	liveHolds := "liveHolds_example" // string | Merge active buyer holds from the per-event hold index (default `true`). `false` = Postgres only. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SeasonsAPI.GetApiSeasonsBySeasonIdReportsSummaryByStatusJson(context.Background(), seasonId).LiveHolds(liveHolds).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SeasonsAPI.GetApiSeasonsBySeasonIdReportsSummaryByStatusJson``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**seasonId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetApiSeasonsBySeasonIdReportsSummaryByStatusJsonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **liveHolds** | **string** | Merge active buyer holds from the per-event hold index (default &#x60;true&#x60;). &#x60;false&#x60; &#x3D; Postgres only. | 

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


## GetApiSeasonsBySeasonIdReportsSummaryByZoneCsv

> GetApiSeasonsBySeasonIdReportsSummaryByZoneCsv(ctx, seasonId).LiveHolds(liveHolds).Execute()

Season summary by zone (CSV)



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
	seasonId := "seasonId_example" // string | 
	liveHolds := "liveHolds_example" // string | Merge active buyer holds from the per-event hold index (default `true`). `false` = Postgres only. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SeasonsAPI.GetApiSeasonsBySeasonIdReportsSummaryByZoneCsv(context.Background(), seasonId).LiveHolds(liveHolds).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SeasonsAPI.GetApiSeasonsBySeasonIdReportsSummaryByZoneCsv``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**seasonId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetApiSeasonsBySeasonIdReportsSummaryByZoneCsvRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **liveHolds** | **string** | Merge active buyer holds from the per-event hold index (default &#x60;true&#x60;). &#x60;false&#x60; &#x3D; Postgres only. | 

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


## GetApiSeasonsBySeasonIdReportsSummaryByZoneJson

> GetApiSeasonsBySeasonIdReportsSummaryByZoneJson(ctx, seasonId).LiveHolds(liveHolds).Execute()

Season summary by zone



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
	seasonId := "seasonId_example" // string | 
	liveHolds := "liveHolds_example" // string | Merge active buyer holds from the per-event hold index (default `true`). `false` = Postgres only. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SeasonsAPI.GetApiSeasonsBySeasonIdReportsSummaryByZoneJson(context.Background(), seasonId).LiveHolds(liveHolds).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SeasonsAPI.GetApiSeasonsBySeasonIdReportsSummaryByZoneJson``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**seasonId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetApiSeasonsBySeasonIdReportsSummaryByZoneJsonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **liveHolds** | **string** | Merge active buyer holds from the per-event hold index (default &#x60;true&#x60;). &#x60;false&#x60; &#x3D; Postgres only. | 

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


## GetApiSeasonsBySeasonIdReportsSummaryCsv

> GetApiSeasonsBySeasonIdReportsSummaryCsv(ctx, seasonId).LiveHolds(liveHolds).Execute()

Season summary across member events (CSV)



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
	seasonId := "seasonId_example" // string | 
	liveHolds := "liveHolds_example" // string | Merge active buyer holds from the per-event hold index (default `true`). `false` = Postgres only. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SeasonsAPI.GetApiSeasonsBySeasonIdReportsSummaryCsv(context.Background(), seasonId).LiveHolds(liveHolds).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SeasonsAPI.GetApiSeasonsBySeasonIdReportsSummaryCsv``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**seasonId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetApiSeasonsBySeasonIdReportsSummaryCsvRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **liveHolds** | **string** | Merge active buyer holds from the per-event hold index (default &#x60;true&#x60;). &#x60;false&#x60; &#x3D; Postgres only. | 

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


## GetApiSeasonsBySeasonIdReportsSummaryJson

> GetApiSeasonsBySeasonIdReportsSummaryJson(ctx, seasonId).LiveHolds(liveHolds).Execute()

Season summary across member events



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
	seasonId := "seasonId_example" // string | 
	liveHolds := "liveHolds_example" // string | Merge active buyer holds from the per-event hold index (default `true`). `false` = Postgres only. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SeasonsAPI.GetApiSeasonsBySeasonIdReportsSummaryJson(context.Background(), seasonId).LiveHolds(liveHolds).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SeasonsAPI.GetApiSeasonsBySeasonIdReportsSummaryJson``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**seasonId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetApiSeasonsBySeasonIdReportsSummaryJsonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **liveHolds** | **string** | Merge active buyer holds from the per-event hold index (default &#x60;true&#x60;). &#x60;false&#x60; &#x3D; Postgres only. | 

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


## GetApiSeasonsBySeasonIdSeats

> GetApiSeasonsBySeasonIdSeats(ctx, seasonId).Execute()

Season booking view (source geometry + intersection status)



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
	seasonId := "seasonId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SeasonsAPI.GetApiSeasonsBySeasonIdSeats(context.Background(), seasonId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SeasonsAPI.GetApiSeasonsBySeasonIdSeats``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**seasonId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetApiSeasonsBySeasonIdSeatsRequest struct via the builder pattern


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


## PatchApiSeasonsBySeasonId

> PatchApiSeasonsBySeasonId(ctx, seasonId).PatchApiSeasonsBySeasonIdRequest(patchApiSeasonsBySeasonIdRequest).IdempotencyKey(idempotencyKey).Execute()

Update season



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
	seasonId := "seasonId_example" // string | 
	patchApiSeasonsBySeasonIdRequest := *openapiclient.NewPatchApiSeasonsBySeasonIdRequest() // PatchApiSeasonsBySeasonIdRequest | 
	idempotencyKey := "idempotencyKey_example" // string | Required on transactional routes. Reuse the same key for every retry of one logical operation; the API replays the stored outcome. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SeasonsAPI.PatchApiSeasonsBySeasonId(context.Background(), seasonId).PatchApiSeasonsBySeasonIdRequest(patchApiSeasonsBySeasonIdRequest).IdempotencyKey(idempotencyKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SeasonsAPI.PatchApiSeasonsBySeasonId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**seasonId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchApiSeasonsBySeasonIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **patchApiSeasonsBySeasonIdRequest** | [**PatchApiSeasonsBySeasonIdRequest**](PatchApiSeasonsBySeasonIdRequest.md) |  | 
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


## PatchApiSeasonsBySeasonIdEventsByEventId

> PatchApiSeasonsBySeasonIdEventsByEventId(ctx, seasonId, eventId).PatchApiSeasonsBySeasonIdEventsByEventIdRequest(patchApiSeasonsBySeasonIdEventsByEventIdRequest).IdempotencyKey(idempotencyKey).Execute()

Season member settings (forSalePropagated)



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
	seasonId := "seasonId_example" // string | 
	eventId := "eventId_example" // string | 
	patchApiSeasonsBySeasonIdEventsByEventIdRequest := *openapiclient.NewPatchApiSeasonsBySeasonIdEventsByEventIdRequest(false) // PatchApiSeasonsBySeasonIdEventsByEventIdRequest | 
	idempotencyKey := "idempotencyKey_example" // string | Required on transactional routes. Reuse the same key for every retry of one logical operation; the API replays the stored outcome. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SeasonsAPI.PatchApiSeasonsBySeasonIdEventsByEventId(context.Background(), seasonId, eventId).PatchApiSeasonsBySeasonIdEventsByEventIdRequest(patchApiSeasonsBySeasonIdEventsByEventIdRequest).IdempotencyKey(idempotencyKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SeasonsAPI.PatchApiSeasonsBySeasonIdEventsByEventId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**seasonId** | **string** |  | 
**eventId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchApiSeasonsBySeasonIdEventsByEventIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **patchApiSeasonsBySeasonIdEventsByEventIdRequest** | [**PatchApiSeasonsBySeasonIdEventsByEventIdRequest**](PatchApiSeasonsBySeasonIdEventsByEventIdRequest.md) |  | 
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


## PatchApiSeasonsBySeasonIdPartialSeasonsByPartialSeasonId

> PatchApiSeasonsBySeasonIdPartialSeasonsByPartialSeasonId(ctx, seasonId, partialSeasonId).PatchApiSeasonsBySeasonIdPartialSeasonsByPartialSeasonIdRequest(patchApiSeasonsBySeasonIdPartialSeasonsByPartialSeasonIdRequest).IdempotencyKey(idempotencyKey).Execute()

Rename a partial season or set its key



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
	seasonId := "seasonId_example" // string | 
	partialSeasonId := "partialSeasonId_example" // string | 
	patchApiSeasonsBySeasonIdPartialSeasonsByPartialSeasonIdRequest := *openapiclient.NewPatchApiSeasonsBySeasonIdPartialSeasonsByPartialSeasonIdRequest() // PatchApiSeasonsBySeasonIdPartialSeasonsByPartialSeasonIdRequest | 
	idempotencyKey := "idempotencyKey_example" // string | Required on transactional routes. Reuse the same key for every retry of one logical operation; the API replays the stored outcome. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SeasonsAPI.PatchApiSeasonsBySeasonIdPartialSeasonsByPartialSeasonId(context.Background(), seasonId, partialSeasonId).PatchApiSeasonsBySeasonIdPartialSeasonsByPartialSeasonIdRequest(patchApiSeasonsBySeasonIdPartialSeasonsByPartialSeasonIdRequest).IdempotencyKey(idempotencyKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SeasonsAPI.PatchApiSeasonsBySeasonIdPartialSeasonsByPartialSeasonId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**seasonId** | **string** |  | 
**partialSeasonId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchApiSeasonsBySeasonIdPartialSeasonsByPartialSeasonIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **patchApiSeasonsBySeasonIdPartialSeasonsByPartialSeasonIdRequest** | [**PatchApiSeasonsBySeasonIdPartialSeasonsByPartialSeasonIdRequest**](PatchApiSeasonsBySeasonIdPartialSeasonsByPartialSeasonIdRequest.md) |  | 
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


## PostApiEventsByEventIdActionsOverrideSeasonStatus

> PostApiEventsByEventIdActionsOverrideSeasonStatus(ctx, eventId).PostApiEventsByEventIdActionsOverrideSeasonStatusRequest(postApiEventsByEventIdActionsOverrideSeasonStatusRequest).IdempotencyKey(idempotencyKey).Execute()

Override the season status of objects in an event (OVERRIDE_SEASON_STATUS)



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
	postApiEventsByEventIdActionsOverrideSeasonStatusRequest := *openapiclient.NewPostApiEventsByEventIdActionsOverrideSeasonStatusRequest([]string{"Objects_example"}) // PostApiEventsByEventIdActionsOverrideSeasonStatusRequest | 
	idempotencyKey := "idempotencyKey_example" // string | Required on transactional routes. Reuse the same key for every retry of one logical operation; the API replays the stored outcome. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SeasonsAPI.PostApiEventsByEventIdActionsOverrideSeasonStatus(context.Background(), eventId).PostApiEventsByEventIdActionsOverrideSeasonStatusRequest(postApiEventsByEventIdActionsOverrideSeasonStatusRequest).IdempotencyKey(idempotencyKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SeasonsAPI.PostApiEventsByEventIdActionsOverrideSeasonStatus``: %v\n", err)
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

Other parameters are passed through a pointer to a apiPostApiEventsByEventIdActionsOverrideSeasonStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **postApiEventsByEventIdActionsOverrideSeasonStatusRequest** | [**PostApiEventsByEventIdActionsOverrideSeasonStatusRequest**](PostApiEventsByEventIdActionsOverrideSeasonStatusRequest.md) |  | 
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


## PostApiEventsByEventIdActionsUseSeasonStatus

> PostApiEventsByEventIdActionsUseSeasonStatus(ctx, eventId).PostApiEventsByEventIdActionsUseSeasonStatusRequest(postApiEventsByEventIdActionsUseSeasonStatusRequest).IdempotencyKey(idempotencyKey).Execute()

Make objects follow the season status again (USE_SEASON_STATUS)



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
	postApiEventsByEventIdActionsUseSeasonStatusRequest := *openapiclient.NewPostApiEventsByEventIdActionsUseSeasonStatusRequest([]string{"Objects_example"}) // PostApiEventsByEventIdActionsUseSeasonStatusRequest | 
	idempotencyKey := "idempotencyKey_example" // string | Required on transactional routes. Reuse the same key for every retry of one logical operation; the API replays the stored outcome. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SeasonsAPI.PostApiEventsByEventIdActionsUseSeasonStatus(context.Background(), eventId).PostApiEventsByEventIdActionsUseSeasonStatusRequest(postApiEventsByEventIdActionsUseSeasonStatusRequest).IdempotencyKey(idempotencyKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SeasonsAPI.PostApiEventsByEventIdActionsUseSeasonStatus``: %v\n", err)
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

Other parameters are passed through a pointer to a apiPostApiEventsByEventIdActionsUseSeasonStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **postApiEventsByEventIdActionsUseSeasonStatusRequest** | [**PostApiEventsByEventIdActionsUseSeasonStatusRequest**](PostApiEventsByEventIdActionsUseSeasonStatusRequest.md) |  | 
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


## PostApiSeasons

> PostApiSeasons(ctx).PostApiSeasonsRequest(postApiSeasonsRequest).IdempotencyKey(idempotencyKey).Execute()

Create season



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
	postApiSeasonsRequest := *openapiclient.NewPostApiSeasonsRequest("Name_example") // PostApiSeasonsRequest | 
	idempotencyKey := "idempotencyKey_example" // string | Required on transactional routes. Reuse the same key for every retry of one logical operation; the API replays the stored outcome. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SeasonsAPI.PostApiSeasons(context.Background()).PostApiSeasonsRequest(postApiSeasonsRequest).IdempotencyKey(idempotencyKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SeasonsAPI.PostApiSeasons``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostApiSeasonsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **postApiSeasonsRequest** | [**PostApiSeasonsRequest**](PostApiSeasonsRequest.md) |  | 
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


## PostApiSeasonsBySeasonIdActionsAddEvents

> PostApiSeasonsBySeasonIdActionsAddEvents(ctx, seasonId).PostApiSeasonsBySeasonIdActionsAddEventsRequest(postApiSeasonsBySeasonIdActionsAddEventsRequest).IdempotencyKey(idempotencyKey).Execute()

Add several events to a season (atomic)



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
	seasonId := "seasonId_example" // string | 
	postApiSeasonsBySeasonIdActionsAddEventsRequest := *openapiclient.NewPostApiSeasonsBySeasonIdActionsAddEventsRequest([]string{"EventIds_example"}) // PostApiSeasonsBySeasonIdActionsAddEventsRequest | 
	idempotencyKey := "idempotencyKey_example" // string | Required on transactional routes. Reuse the same key for every retry of one logical operation; the API replays the stored outcome. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SeasonsAPI.PostApiSeasonsBySeasonIdActionsAddEvents(context.Background(), seasonId).PostApiSeasonsBySeasonIdActionsAddEventsRequest(postApiSeasonsBySeasonIdActionsAddEventsRequest).IdempotencyKey(idempotencyKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SeasonsAPI.PostApiSeasonsBySeasonIdActionsAddEvents``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**seasonId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostApiSeasonsBySeasonIdActionsAddEventsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **postApiSeasonsBySeasonIdActionsAddEventsRequest** | [**PostApiSeasonsBySeasonIdActionsAddEventsRequest**](PostApiSeasonsBySeasonIdActionsAddEventsRequest.md) |  | 
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


## PostApiSeasonsBySeasonIdActionsCreateEvents

> PostApiSeasonsBySeasonIdActionsCreateEvents(ctx, seasonId).PostApiSeasonsBySeasonIdActionsCreateEventsRequest(postApiSeasonsBySeasonIdActionsCreateEventsRequest).IdempotencyKey(idempotencyKey).Execute()

Create events in a season



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
	seasonId := "seasonId_example" // string | 
	postApiSeasonsBySeasonIdActionsCreateEventsRequest := *openapiclient.NewPostApiSeasonsBySeasonIdActionsCreateEventsRequest() // PostApiSeasonsBySeasonIdActionsCreateEventsRequest | 
	idempotencyKey := "idempotencyKey_example" // string | Required on transactional routes. Reuse the same key for every retry of one logical operation; the API replays the stored outcome. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SeasonsAPI.PostApiSeasonsBySeasonIdActionsCreateEvents(context.Background(), seasonId).PostApiSeasonsBySeasonIdActionsCreateEventsRequest(postApiSeasonsBySeasonIdActionsCreateEventsRequest).IdempotencyKey(idempotencyKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SeasonsAPI.PostApiSeasonsBySeasonIdActionsCreateEvents``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**seasonId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostApiSeasonsBySeasonIdActionsCreateEventsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **postApiSeasonsBySeasonIdActionsCreateEventsRequest** | [**PostApiSeasonsBySeasonIdActionsCreateEventsRequest**](PostApiSeasonsBySeasonIdActionsCreateEventsRequest.md) |  | 
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


## PostApiSeasonsBySeasonIdActionsRemoveEvents

> PostApiSeasonsBySeasonIdActionsRemoveEvents(ctx, seasonId).PostApiSeasonsBySeasonIdActionsRemoveEventsRequest(postApiSeasonsBySeasonIdActionsRemoveEventsRequest).IdempotencyKey(idempotencyKey).Execute()

Remove several events from a season (atomic)



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
	seasonId := "seasonId_example" // string | 
	postApiSeasonsBySeasonIdActionsRemoveEventsRequest := *openapiclient.NewPostApiSeasonsBySeasonIdActionsRemoveEventsRequest([]string{"EventIds_example"}) // PostApiSeasonsBySeasonIdActionsRemoveEventsRequest | 
	idempotencyKey := "idempotencyKey_example" // string | Required on transactional routes. Reuse the same key for every retry of one logical operation; the API replays the stored outcome. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SeasonsAPI.PostApiSeasonsBySeasonIdActionsRemoveEvents(context.Background(), seasonId).PostApiSeasonsBySeasonIdActionsRemoveEventsRequest(postApiSeasonsBySeasonIdActionsRemoveEventsRequest).IdempotencyKey(idempotencyKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SeasonsAPI.PostApiSeasonsBySeasonIdActionsRemoveEvents``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**seasonId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostApiSeasonsBySeasonIdActionsRemoveEventsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **postApiSeasonsBySeasonIdActionsRemoveEventsRequest** | [**PostApiSeasonsBySeasonIdActionsRemoveEventsRequest**](PostApiSeasonsBySeasonIdActionsRemoveEventsRequest.md) |  | 
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


## PostApiSeasonsBySeasonIdEvents

> PostApiSeasonsBySeasonIdEvents(ctx, seasonId).PostApiSeasonsBySeasonIdEventsRequest(postApiSeasonsBySeasonIdEventsRequest).IdempotencyKey(idempotencyKey).Execute()

Attach event to season



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
	seasonId := "seasonId_example" // string | 
	postApiSeasonsBySeasonIdEventsRequest := *openapiclient.NewPostApiSeasonsBySeasonIdEventsRequest("EventId_example") // PostApiSeasonsBySeasonIdEventsRequest | 
	idempotencyKey := "idempotencyKey_example" // string | Required on transactional routes. Reuse the same key for every retry of one logical operation; the API replays the stored outcome. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SeasonsAPI.PostApiSeasonsBySeasonIdEvents(context.Background(), seasonId).PostApiSeasonsBySeasonIdEventsRequest(postApiSeasonsBySeasonIdEventsRequest).IdempotencyKey(idempotencyKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SeasonsAPI.PostApiSeasonsBySeasonIdEvents``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**seasonId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostApiSeasonsBySeasonIdEventsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **postApiSeasonsBySeasonIdEventsRequest** | [**PostApiSeasonsBySeasonIdEventsRequest**](PostApiSeasonsBySeasonIdEventsRequest.md) |  | 
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


## PostApiSeasonsBySeasonIdInventorySync

> PostApiSeasonsBySeasonIdInventorySync(ctx, seasonId).PostApiSeasonsBySeasonIdInventorySyncRequest(postApiSeasonsBySeasonIdInventorySyncRequest).IdempotencyKey(idempotencyKey).Execute()

Sync season inventory from source



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
	seasonId := "seasonId_example" // string | 
	postApiSeasonsBySeasonIdInventorySyncRequest := *openapiclient.NewPostApiSeasonsBySeasonIdInventorySyncRequest() // PostApiSeasonsBySeasonIdInventorySyncRequest | 
	idempotencyKey := "idempotencyKey_example" // string | Required on transactional routes. Reuse the same key for every retry of one logical operation; the API replays the stored outcome. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SeasonsAPI.PostApiSeasonsBySeasonIdInventorySync(context.Background(), seasonId).PostApiSeasonsBySeasonIdInventorySyncRequest(postApiSeasonsBySeasonIdInventorySyncRequest).IdempotencyKey(idempotencyKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SeasonsAPI.PostApiSeasonsBySeasonIdInventorySync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**seasonId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostApiSeasonsBySeasonIdInventorySyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **postApiSeasonsBySeasonIdInventorySyncRequest** | [**PostApiSeasonsBySeasonIdInventorySyncRequest**](PostApiSeasonsBySeasonIdInventorySyncRequest.md) |  | 
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


## PostApiSeasonsBySeasonIdPartialSeasons

> PostApiSeasonsBySeasonIdPartialSeasons(ctx, seasonId).PostApiSeasonsBySeasonIdPartialSeasonsRequest(postApiSeasonsBySeasonIdPartialSeasonsRequest).IdempotencyKey(idempotencyKey).Execute()

Create a partial season



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
	seasonId := "seasonId_example" // string | 
	postApiSeasonsBySeasonIdPartialSeasonsRequest := *openapiclient.NewPostApiSeasonsBySeasonIdPartialSeasonsRequest("Name_example") // PostApiSeasonsBySeasonIdPartialSeasonsRequest | 
	idempotencyKey := "idempotencyKey_example" // string | Required on transactional routes. Reuse the same key for every retry of one logical operation; the API replays the stored outcome. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SeasonsAPI.PostApiSeasonsBySeasonIdPartialSeasons(context.Background(), seasonId).PostApiSeasonsBySeasonIdPartialSeasonsRequest(postApiSeasonsBySeasonIdPartialSeasonsRequest).IdempotencyKey(idempotencyKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SeasonsAPI.PostApiSeasonsBySeasonIdPartialSeasons``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**seasonId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostApiSeasonsBySeasonIdPartialSeasonsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **postApiSeasonsBySeasonIdPartialSeasonsRequest** | [**PostApiSeasonsBySeasonIdPartialSeasonsRequest**](PostApiSeasonsBySeasonIdPartialSeasonsRequest.md) |  | 
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


## PostApiSeasonsBySeasonIdPartialSeasonsByPartialSeasonIdActionsAddEvents

> PostApiSeasonsBySeasonIdPartialSeasonsByPartialSeasonIdActionsAddEvents(ctx, seasonId, partialSeasonId).PostApiSeasonsBySeasonIdActionsRemoveEventsRequest(postApiSeasonsBySeasonIdActionsRemoveEventsRequest).IdempotencyKey(idempotencyKey).Execute()

Add events to a partial season (bulk, all-or-nothing)



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
	seasonId := "seasonId_example" // string | 
	partialSeasonId := "partialSeasonId_example" // string | 
	postApiSeasonsBySeasonIdActionsRemoveEventsRequest := *openapiclient.NewPostApiSeasonsBySeasonIdActionsRemoveEventsRequest([]string{"EventIds_example"}) // PostApiSeasonsBySeasonIdActionsRemoveEventsRequest | 
	idempotencyKey := "idempotencyKey_example" // string | Required on transactional routes. Reuse the same key for every retry of one logical operation; the API replays the stored outcome. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SeasonsAPI.PostApiSeasonsBySeasonIdPartialSeasonsByPartialSeasonIdActionsAddEvents(context.Background(), seasonId, partialSeasonId).PostApiSeasonsBySeasonIdActionsRemoveEventsRequest(postApiSeasonsBySeasonIdActionsRemoveEventsRequest).IdempotencyKey(idempotencyKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SeasonsAPI.PostApiSeasonsBySeasonIdPartialSeasonsByPartialSeasonIdActionsAddEvents``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**seasonId** | **string** |  | 
**partialSeasonId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostApiSeasonsBySeasonIdPartialSeasonsByPartialSeasonIdActionsAddEventsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **postApiSeasonsBySeasonIdActionsRemoveEventsRequest** | [**PostApiSeasonsBySeasonIdActionsRemoveEventsRequest**](PostApiSeasonsBySeasonIdActionsRemoveEventsRequest.md) |  | 
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


## PostApiSeasonsBySeasonIdPartialSeasonsByPartialSeasonIdActionsRemoveEvents

> PostApiSeasonsBySeasonIdPartialSeasonsByPartialSeasonIdActionsRemoveEvents(ctx, seasonId, partialSeasonId).PostApiSeasonsBySeasonIdActionsRemoveEventsRequest(postApiSeasonsBySeasonIdActionsRemoveEventsRequest).IdempotencyKey(idempotencyKey).Execute()

Remove events from a partial season (bulk)



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
	seasonId := "seasonId_example" // string | 
	partialSeasonId := "partialSeasonId_example" // string | 
	postApiSeasonsBySeasonIdActionsRemoveEventsRequest := *openapiclient.NewPostApiSeasonsBySeasonIdActionsRemoveEventsRequest([]string{"EventIds_example"}) // PostApiSeasonsBySeasonIdActionsRemoveEventsRequest | 
	idempotencyKey := "idempotencyKey_example" // string | Required on transactional routes. Reuse the same key for every retry of one logical operation; the API replays the stored outcome. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SeasonsAPI.PostApiSeasonsBySeasonIdPartialSeasonsByPartialSeasonIdActionsRemoveEvents(context.Background(), seasonId, partialSeasonId).PostApiSeasonsBySeasonIdActionsRemoveEventsRequest(postApiSeasonsBySeasonIdActionsRemoveEventsRequest).IdempotencyKey(idempotencyKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SeasonsAPI.PostApiSeasonsBySeasonIdPartialSeasonsByPartialSeasonIdActionsRemoveEvents``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**seasonId** | **string** |  | 
**partialSeasonId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostApiSeasonsBySeasonIdPartialSeasonsByPartialSeasonIdActionsRemoveEventsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **postApiSeasonsBySeasonIdActionsRemoveEventsRequest** | [**PostApiSeasonsBySeasonIdActionsRemoveEventsRequest**](PostApiSeasonsBySeasonIdActionsRemoveEventsRequest.md) |  | 
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

