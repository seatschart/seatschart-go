# \EventsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteApiEventsByEventId**](EventsAPI.md#DeleteApiEventsByEventId) | **Delete** /api/events/{eventId} | Delete event
[**DeleteApiEventsByEventIdInventoryChannelsByChannelKeyId**](EventsAPI.md#DeleteApiEventsByEventIdInventoryChannelsByChannelKeyId) | **Delete** /api/events/{eventId}/inventory/channels/{channelKeyId} | Delete inventory channel
[**GetApiEvents**](EventsAPI.md#GetApiEvents) | **Get** /api/events | List events for the authenticated tenant
[**GetApiEventsByEventId**](EventsAPI.md#GetApiEventsByEventId) | **Get** /api/events/{eventId} | Retrieve one event
[**GetApiEventsByEventIdChannelMembership**](EventsAPI.md#GetApiEventsByEventIdChannelMembership) | **Get** /api/events/{eventId}/channel-membership | Do objects belong to my inventory channel? (renderer isObjectInChannel)
[**GetApiEventsByEventIdInventoryChannels**](EventsAPI.md#GetApiEventsByEventIdInventoryChannels) | **Get** /api/events/{eventId}/inventory/channels | List inventory channels (ops)
[**GetApiEventsByEventIdMapDraft**](EventsAPI.md#GetApiEventsByEventIdMapDraft) | **Get** /api/events/{eventId}/map/draft | Load map draft for designer
[**GetApiEventsByEventIdMapVersions**](EventsAPI.md#GetApiEventsByEventIdMapVersions) | **Get** /api/events/{eventId}/map/versions | List event map versions
[**GetApiEventsByEventIdMapVersionsByVersionThumbnail**](EventsAPI.md#GetApiEventsByEventIdMapVersionsByVersionThumbnail) | **Get** /api/events/{eventId}/map/versions/{version}/thumbnail | Event map version preview
[**GetApiEventsByEventIdObjects**](EventsAPI.md#GetApiEventsByEventIdObjects) | **Get** /api/events/{eventId}/objects | List the objects of a ticket buyer
[**GetApiEventsByEventIdObjectsByObjectId**](EventsAPI.md#GetApiEventsByEventIdObjectsByObjectId) | **Get** /api/events/{eventId}/objects/{objectId} | Retrieve one object (seat or GA area)
[**GetApiEventsByEventIdObjectsByObjectIdStatusChanges**](EventsAPI.md#GetApiEventsByEventIdObjectsByObjectIdStatusChanges) | **Get** /api/events/{eventId}/objects/{objectId}/status-changes | List status changes of one object
[**GetApiEventsByEventIdPricing**](EventsAPI.md#GetApiEventsByEventIdPricing) | **Get** /api/events/{eventId}/pricing | Read the synced price table of an event
[**GetApiEventsByEventIdReportsByAvailabilityCsv**](EventsAPI.md#GetApiEventsByEventIdReportsByAvailabilityCsv) | **Get** /api/events/{eventId}/reports/by-availability.csv | Detailed report grouped by availability (available / not_available) (CSV)
[**GetApiEventsByEventIdReportsByAvailabilityJson**](EventsAPI.md#GetApiEventsByEventIdReportsByAvailabilityJson) | **Get** /api/events/{eventId}/reports/by-availability.json | Detailed report grouped by availability (available / not_available)
[**GetApiEventsByEventIdReportsByAvailabilityReasonCsv**](EventsAPI.md#GetApiEventsByEventIdReportsByAvailabilityReasonCsv) | **Get** /api/events/{eventId}/reports/by-availability-reason.csv | Detailed report grouped by availability reason (CSV)
[**GetApiEventsByEventIdReportsByAvailabilityReasonJson**](EventsAPI.md#GetApiEventsByEventIdReportsByAvailabilityReasonJson) | **Get** /api/events/{eventId}/reports/by-availability-reason.json | Detailed report grouped by availability reason
[**GetApiEventsByEventIdReportsByCategoryJson**](EventsAPI.md#GetApiEventsByEventIdReportsByCategoryJson) | **Get** /api/events/{eventId}/reports/by-category.json | Seat and GA counts by category
[**GetApiEventsByEventIdReportsByCategoryKeyCsv**](EventsAPI.md#GetApiEventsByEventIdReportsByCategoryKeyCsv) | **Get** /api/events/{eventId}/reports/by-category-key.csv | Detailed report grouped by category key (CSV)
[**GetApiEventsByEventIdReportsByCategoryKeyJson**](EventsAPI.md#GetApiEventsByEventIdReportsByCategoryKeyJson) | **Get** /api/events/{eventId}/reports/by-category-key.json | Detailed report grouped by category key
[**GetApiEventsByEventIdReportsByCategoryLabelCsv**](EventsAPI.md#GetApiEventsByEventIdReportsByCategoryLabelCsv) | **Get** /api/events/{eventId}/reports/by-category-label.csv | Detailed report grouped by category label (CSV)
[**GetApiEventsByEventIdReportsByCategoryLabelJson**](EventsAPI.md#GetApiEventsByEventIdReportsByCategoryLabelJson) | **Get** /api/events/{eventId}/reports/by-category-label.json | Detailed report grouped by category label
[**GetApiEventsByEventIdReportsByChannelJson**](EventsAPI.md#GetApiEventsByEventIdReportsByChannelJson) | **Get** /api/events/{eventId}/reports/by-channel.json | Detailed report grouped by inventory channel
[**GetApiEventsByEventIdReportsByLabelJson**](EventsAPI.md#GetApiEventsByEventIdReportsByLabelJson) | **Get** /api/events/{eventId}/reports/by-label.json | Detailed report grouped by object label (Postgres)
[**GetApiEventsByEventIdReportsByObjectTypeJson**](EventsAPI.md#GetApiEventsByEventIdReportsByObjectTypeJson) | **Get** /api/events/{eventId}/reports/by-object-type.json | Detailed report grouped by object type (seat|ga)
[**GetApiEventsByEventIdReportsByOrderIdJson**](EventsAPI.md#GetApiEventsByEventIdReportsByOrderIdJson) | **Get** /api/events/{eventId}/reports/by-order-id.json | Detailed report grouped by orderId (Postgres)
[**GetApiEventsByEventIdReportsBySectionJson**](EventsAPI.md#GetApiEventsByEventIdReportsBySectionJson) | **Get** /api/events/{eventId}/reports/by-section.json | Detailed report grouped by section
[**GetApiEventsByEventIdReportsByStatusJson**](EventsAPI.md#GetApiEventsByEventIdReportsByStatusJson) | **Get** /api/events/{eventId}/reports/by-status.json | Seat counts by Postgres status
[**GetApiEventsByEventIdReportsByZoneCsv**](EventsAPI.md#GetApiEventsByEventIdReportsByZoneCsv) | **Get** /api/events/{eventId}/reports/by-zone.csv | Detailed report grouped by zone (CSV)
[**GetApiEventsByEventIdReportsByZoneJson**](EventsAPI.md#GetApiEventsByEventIdReportsByZoneJson) | **Get** /api/events/{eventId}/reports/by-zone.json | Detailed report grouped by zone
[**GetApiEventsByEventIdReportsDeepSummaryJson**](EventsAPI.md#GetApiEventsByEventIdReportsDeepSummaryJson) | **Get** /api/events/{eventId}/reports/deep-summary.json | Deep summary: status counts nested by category
[**GetApiEventsByEventIdReportsOccupancyCsv**](EventsAPI.md#GetApiEventsByEventIdReportsOccupancyCsv) | **Get** /api/events/{eventId}/reports/occupancy.csv | Export seat and GA occupancy as CSV
[**GetApiEventsByEventIdReportsSalesCsv**](EventsAPI.md#GetApiEventsByEventIdReportsSalesCsv) | **Get** /api/events/{eventId}/reports/sales.csv | Export confirmed seat and GA sales as CSV
[**GetApiEventsByEventIdReportsSummaryByAvailabilityCsv**](EventsAPI.md#GetApiEventsByEventIdReportsSummaryByAvailabilityCsv) | **Get** /api/events/{eventId}/reports/summary/by-availability.csv | Summary report by availability (CSV)
[**GetApiEventsByEventIdReportsSummaryByAvailabilityJson**](EventsAPI.md#GetApiEventsByEventIdReportsSummaryByAvailabilityJson) | **Get** /api/events/{eventId}/reports/summary/by-availability.json | Summary report by availability
[**GetApiEventsByEventIdReportsSummaryByAvailabilityReasonCsv**](EventsAPI.md#GetApiEventsByEventIdReportsSummaryByAvailabilityReasonCsv) | **Get** /api/events/{eventId}/reports/summary/by-availability-reason.csv | Summary report by availability reason (CSV)
[**GetApiEventsByEventIdReportsSummaryByAvailabilityReasonJson**](EventsAPI.md#GetApiEventsByEventIdReportsSummaryByAvailabilityReasonJson) | **Get** /api/events/{eventId}/reports/summary/by-availability-reason.json | Summary report by availability reason
[**GetApiEventsByEventIdReportsSummaryByCategoryCsv**](EventsAPI.md#GetApiEventsByEventIdReportsSummaryByCategoryCsv) | **Get** /api/events/{eventId}/reports/summary/by-category.csv | Summary report by category (CSV)
[**GetApiEventsByEventIdReportsSummaryByCategoryJson**](EventsAPI.md#GetApiEventsByEventIdReportsSummaryByCategoryJson) | **Get** /api/events/{eventId}/reports/summary/by-category.json | Summary report by category
[**GetApiEventsByEventIdReportsSummaryByCategoryKeyCsv**](EventsAPI.md#GetApiEventsByEventIdReportsSummaryByCategoryKeyCsv) | **Get** /api/events/{eventId}/reports/summary/by-category-key.csv | Summary report by category key (CSV)
[**GetApiEventsByEventIdReportsSummaryByCategoryKeyJson**](EventsAPI.md#GetApiEventsByEventIdReportsSummaryByCategoryKeyJson) | **Get** /api/events/{eventId}/reports/summary/by-category-key.json | Summary report by category key
[**GetApiEventsByEventIdReportsSummaryByCategoryLabelCsv**](EventsAPI.md#GetApiEventsByEventIdReportsSummaryByCategoryLabelCsv) | **Get** /api/events/{eventId}/reports/summary/by-category-label.csv | Summary report by category label (CSV)
[**GetApiEventsByEventIdReportsSummaryByCategoryLabelJson**](EventsAPI.md#GetApiEventsByEventIdReportsSummaryByCategoryLabelJson) | **Get** /api/events/{eventId}/reports/summary/by-category-label.json | Summary report by category label
[**GetApiEventsByEventIdReportsSummaryByChannelCsv**](EventsAPI.md#GetApiEventsByEventIdReportsSummaryByChannelCsv) | **Get** /api/events/{eventId}/reports/summary/by-channel.csv | Summary report by channel (CSV)
[**GetApiEventsByEventIdReportsSummaryByChannelJson**](EventsAPI.md#GetApiEventsByEventIdReportsSummaryByChannelJson) | **Get** /api/events/{eventId}/reports/summary/by-channel.json | Summary report by channel
[**GetApiEventsByEventIdReportsSummaryByObjectTypeCsv**](EventsAPI.md#GetApiEventsByEventIdReportsSummaryByObjectTypeCsv) | **Get** /api/events/{eventId}/reports/summary/by-object-type.csv | Summary report by object type (CSV)
[**GetApiEventsByEventIdReportsSummaryByObjectTypeJson**](EventsAPI.md#GetApiEventsByEventIdReportsSummaryByObjectTypeJson) | **Get** /api/events/{eventId}/reports/summary/by-object-type.json | Summary report by object type
[**GetApiEventsByEventIdReportsSummaryBySectionCsv**](EventsAPI.md#GetApiEventsByEventIdReportsSummaryBySectionCsv) | **Get** /api/events/{eventId}/reports/summary/by-section.csv | Summary report by section (CSV)
[**GetApiEventsByEventIdReportsSummaryBySectionJson**](EventsAPI.md#GetApiEventsByEventIdReportsSummaryBySectionJson) | **Get** /api/events/{eventId}/reports/summary/by-section.json | Summary report by section
[**GetApiEventsByEventIdReportsSummaryByStatusCsv**](EventsAPI.md#GetApiEventsByEventIdReportsSummaryByStatusCsv) | **Get** /api/events/{eventId}/reports/summary/by-status.csv | Summary report by status (CSV)
[**GetApiEventsByEventIdReportsSummaryByStatusJson**](EventsAPI.md#GetApiEventsByEventIdReportsSummaryByStatusJson) | **Get** /api/events/{eventId}/reports/summary/by-status.json | Summary report by status
[**GetApiEventsByEventIdReportsSummaryByZoneCsv**](EventsAPI.md#GetApiEventsByEventIdReportsSummaryByZoneCsv) | **Get** /api/events/{eventId}/reports/summary/by-zone.csv | Summary report by zone (CSV)
[**GetApiEventsByEventIdReportsSummaryByZoneJson**](EventsAPI.md#GetApiEventsByEventIdReportsSummaryByZoneJson) | **Get** /api/events/{eventId}/reports/summary/by-zone.json | Summary report by zone
[**GetApiEventsByEventIdReportsSummaryJson**](EventsAPI.md#GetApiEventsByEventIdReportsSummaryJson) | **Get** /api/events/{eventId}/reports/summary.json | Occupancy summary JSON (Postgres only)
[**GetApiEventsByEventIdSalesCounters**](EventsAPI.md#GetApiEventsByEventIdSalesCounters) | **Get** /api/events/{eventId}/sales-counters | Live sales counters (quotas, lots) of an event
[**GetApiEventsByEventIdSalesRules**](EventsAPI.md#GetApiEventsByEventIdSalesRules) | **Get** /api/events/{eventId}/sales-rules | Read the sales rules of an event
[**GetApiEventsByEventIdSeatViews**](EventsAPI.md#GetApiEventsByEventIdSeatViews) | **Get** /api/events/{eventId}/seat-views | List view-from-seat image and optional 360 URLs
[**GetApiEventsByEventIdSeats**](EventsAPI.md#GetApiEventsByEventIdSeats) | **Get** /api/events/{eventId}/seats | List seats for an event
[**GetApiEventsByEventIdState**](EventsAPI.md#GetApiEventsByEventIdState) | **Get** /api/events/{eventId}/state | Compact visual state (protocol 5)
[**GetApiEventsByEventIdStatusChanges**](EventsAPI.md#GetApiEventsByEventIdStatusChanges) | **Get** /api/events/{eventId}/status-changes | List status changes of an event
[**GetApiEventsByEventIdZones**](EventsAPI.md#GetApiEventsByEventIdZones) | **Get** /api/events/{eventId}/zones | List the zones of an event
[**PatchApiEventsByEventId**](EventsAPI.md#PatchApiEventsByEventId) | **Patch** /api/events/{eventId} | Patch event: name, date, currency, overrides (isInThePast is irreversible)
[**PatchApiEventsByEventIdInventoryChannelsByChannelKeyId**](EventsAPI.md#PatchApiEventsByEventIdInventoryChannelsByChannelKeyId) | **Patch** /api/events/{eventId}/inventory/channels/{channelKeyId} | Rename or recolor an inventory channel
[**PostApiEvents**](EventsAPI.md#PostApiEvents) | **Post** /api/events | Create event
[**PostApiEventsBulk**](EventsAPI.md#PostApiEventsBulk) | **Post** /api/events/bulk | Create up to 50 events for one chart
[**PostApiEventsByEventIdActionsMarkEverythingAsForSale**](EventsAPI.md#PostApiEventsByEventIdActionsMarkEverythingAsForSale) | **Post** /api/events/{eventId}/actions/mark-everything-as-for-sale | Put every object of the event back on sale
[**PostApiEventsByEventIdActionsMarkEverythingAsNotForSale**](EventsAPI.md#PostApiEventsByEventIdActionsMarkEverythingAsNotForSale) | **Post** /api/events/{eventId}/actions/mark-everything-as-not-for-sale | Mark every object of the event not for sale
[**PostApiEventsByEventIdActionsMoveToNewChartCopy**](EventsAPI.md#PostApiEventsByEventIdActionsMoveToNewChartCopy) | **Post** /api/events/{eventId}/actions/move-to-new-chart-copy | Move the event to a new copy of its chart
[**PostApiEventsByEventIdDuplicate**](EventsAPI.md#PostApiEventsByEventIdDuplicate) | **Post** /api/events/{eventId}/duplicate | Duplicate event
[**PostApiEventsByEventIdInventoryChannels**](EventsAPI.md#PostApiEventsByEventIdInventoryChannels) | **Post** /api/events/{eventId}/inventory/channels | Create inventory channel
[**PostApiEventsByEventIdInventoryChannelsAssign**](EventsAPI.md#PostApiEventsByEventIdInventoryChannelsAssign) | **Post** /api/events/{eventId}/inventory/channels/assign | Assign seats/GA to inventory channel
[**PostApiEventsByEventIdInventoryForSaleApply**](EventsAPI.md#PostApiEventsByEventIdInventoryForSaleApply) | **Post** /api/events/{eventId}/inventory/for-sale/apply | Apply for-sale / not-for-sale to a selection (incremental)
[**PostApiEventsByEventIdMapDraftActionsDiscard**](EventsAPI.md#PostApiEventsByEventIdMapDraftActionsDiscard) | **Post** /api/events/{eventId}/map/draft/actions/discard | Discard event map draft
[**PostApiEventsByEventIdMapPublish**](EventsAPI.md#PostApiEventsByEventIdMapPublish) | **Post** /api/events/{eventId}/map/publish | Publish draft map to live seats
[**PostApiEventsByEventIdMapVersionsByVersionActionsRestoreToDraft**](EventsAPI.md#PostApiEventsByEventIdMapVersionsByVersionActionsRestoreToDraft) | **Post** /api/events/{eventId}/map/versions/{version}/actions/restore-to-draft | Restore an event map version into the draft
[**PostApiEventsByEventIdObjectsActionsRetrieve**](EventsAPI.md#PostApiEventsByEventIdObjectsActionsRetrieve) | **Post** /api/events/{eventId}/objects/actions/retrieve | Retrieve several objects at once (seats.io retrieveObjectInfos)
[**PostApiEventsByEventIdSalesCountersActionsRebuild**](EventsAPI.md#PostApiEventsByEventIdSalesCountersActionsRebuild) | **Post** /api/events/{eventId}/sales-counters/actions/rebuild | Rebuild the sold counters of an event from its orders
[**PutApiEventsByEventIdInventoryChannels**](EventsAPI.md#PutApiEventsByEventIdInventoryChannels) | **Put** /api/events/{eventId}/inventory/channels | Replace all inventory channel assignments atomically
[**PutApiEventsByEventIdInventoryChannelsDefinitions**](EventsAPI.md#PutApiEventsByEventIdInventoryChannelsDefinitions) | **Put** /api/events/{eventId}/inventory/channels/definitions | Replace all inventory channels (definitions and assignments) atomically
[**PutApiEventsByEventIdInventoryForSale**](EventsAPI.md#PutApiEventsByEventIdInventoryForSale) | **Put** /api/events/{eventId}/inventory/for-sale | Replace the whole for-sale configuration atomically
[**PutApiEventsByEventIdMap**](EventsAPI.md#PutApiEventsByEventIdMap) | **Put** /api/events/{eventId}/map | Save map draft (legacy path)
[**PutApiEventsByEventIdMapDraft**](EventsAPI.md#PutApiEventsByEventIdMapDraft) | **Put** /api/events/{eventId}/map/draft | Save map draft (does not go live)
[**PutApiEventsByEventIdPricing**](EventsAPI.md#PutApiEventsByEventIdPricing) | **Put** /api/events/{eventId}/pricing | Push the synced price table of an event (synced catalog mode)
[**PutApiEventsByEventIdSalesRules**](EventsAPI.md#PutApiEventsByEventIdSalesRules) | **Put** /api/events/{eventId}/sales-rules | Replace the sales rules of an event (quotas, lots, buyer limit, pricing source)



## DeleteApiEventsByEventId

> DeleteApiEventsByEventId(ctx, eventId).IdempotencyKey(idempotencyKey).Execute()

Delete event



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
	r, err := apiClient.EventsAPI.DeleteApiEventsByEventId(context.Background(), eventId).IdempotencyKey(idempotencyKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EventsAPI.DeleteApiEventsByEventId``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDeleteApiEventsByEventIdRequest struct via the builder pattern


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


## DeleteApiEventsByEventIdInventoryChannelsByChannelKeyId

> DeleteApiEventsByEventIdInventoryChannelsByChannelKeyId(ctx, eventId, channelKeyId).IdempotencyKey(idempotencyKey).Execute()

Delete inventory channel



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
	channelKeyId := "channelKeyId_example" // string | 
	idempotencyKey := "idempotencyKey_example" // string | Required on transactional routes. Reuse the same key for every retry of one logical operation; the API replays the stored outcome. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.EventsAPI.DeleteApiEventsByEventIdInventoryChannelsByChannelKeyId(context.Background(), eventId, channelKeyId).IdempotencyKey(idempotencyKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EventsAPI.DeleteApiEventsByEventIdInventoryChannelsByChannelKeyId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**eventId** | **string** |  | 
**channelKeyId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteApiEventsByEventIdInventoryChannelsByChannelKeyIdRequest struct via the builder pattern


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


## GetApiEvents

> GetApiEvents(ctx).Limit(limit).Cursor(cursor).ChartId(chartId).From(from).To(to).Q(q).Key(key).Ended(ended).Sort(sort).Execute()

List events for the authenticated tenant



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/seatschart/seatschart-go"
)

func main() {
	limit := float32(8.14) // float32 | Page size (1..100, default 20). (optional)
	cursor := "cursor_example" // string | Opaque `nextCursor` value from the previous page. (optional)
	chartId := "chartId_example" // string | Only events created from this chart. (optional)
	from := time.Now() // time.Time | Inclusive lower bound on the event date (ISO 8601 date-time). (optional)
	to := time.Now() // time.Time | Inclusive upper bound on the event date (ISO 8601 date-time). (optional)
	q := "q_example" // string | Case-insensitive substring match on the event name. (optional)
	key := "key_example" // string | Exact event key. (optional)
	ended := "ended_example" // string | Events marked as ended (`isInThePast`): `include` (default), `exclude` or `only`. (optional)
	sort := "sort_example" // string | `date` (default): event date ascending. `created`: creation time, newest first. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.EventsAPI.GetApiEvents(context.Background()).Limit(limit).Cursor(cursor).ChartId(chartId).From(from).To(to).Q(q).Key(key).Ended(ended).Sort(sort).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EventsAPI.GetApiEvents``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetApiEventsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **float32** | Page size (1..100, default 20). | 
 **cursor** | **string** | Opaque &#x60;nextCursor&#x60; value from the previous page. | 
 **chartId** | **string** | Only events created from this chart. | 
 **from** | **time.Time** | Inclusive lower bound on the event date (ISO 8601 date-time). | 
 **to** | **time.Time** | Inclusive upper bound on the event date (ISO 8601 date-time). | 
 **q** | **string** | Case-insensitive substring match on the event name. | 
 **key** | **string** | Exact event key. | 
 **ended** | **string** | Events marked as ended (&#x60;isInThePast&#x60;): &#x60;include&#x60; (default), &#x60;exclude&#x60; or &#x60;only&#x60;. | 
 **sort** | **string** | &#x60;date&#x60; (default): event date ascending. &#x60;created&#x60;: creation time, newest first. | 

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


## GetApiEventsByEventId

> GetApiEventsByEventId(ctx, eventId).Expand(expand).Execute()

Retrieve one event



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
	eventId := "eventId_example" // string | Event id (UUID) or event key.
	expand := "expand_example" // string | Comma-separated. `chartTableBooking` adds the table booking of the chart (`{ mode, tables? }` or null) — what `tableBookingConfig` overrides. `zones` adds `zones[]` (same items as GET /events/{eventId}/zones). (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.EventsAPI.GetApiEventsByEventId(context.Background(), eventId).Expand(expand).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EventsAPI.GetApiEventsByEventId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**eventId** | **string** | Event id (UUID) or event key. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetApiEventsByEventIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **expand** | **string** | Comma-separated. &#x60;chartTableBooking&#x60; adds the table booking of the chart (&#x60;{ mode, tables? }&#x60; or null) — what &#x60;tableBookingConfig&#x60; overrides. &#x60;zones&#x60; adds &#x60;zones[]&#x60; (same items as GET /events/{eventId}/zones). | 

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


## GetApiEventsByEventIdChannelMembership

> GetApiEventsByEventIdChannelMembership(ctx, eventId).Objects(objects).Execute()

Do objects belong to my inventory channel? (renderer isObjectInChannel)



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
	objects := "objects_example" // string | Comma-separated seat ids / labels or GA area ids / names (max 100).

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.EventsAPI.GetApiEventsByEventIdChannelMembership(context.Background(), eventId).Objects(objects).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EventsAPI.GetApiEventsByEventIdChannelMembership``: %v\n", err)
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

Other parameters are passed through a pointer to a apiGetApiEventsByEventIdChannelMembershipRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **objects** | **string** | Comma-separated seat ids / labels or GA area ids / names (max 100). | 

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


## GetApiEventsByEventIdInventoryChannels

> GetApiEventsByEventIdInventoryChannels(ctx, eventId).Execute()

List inventory channels (ops)



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
	r, err := apiClient.EventsAPI.GetApiEventsByEventIdInventoryChannels(context.Background(), eventId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EventsAPI.GetApiEventsByEventIdInventoryChannels``: %v\n", err)
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

Other parameters are passed through a pointer to a apiGetApiEventsByEventIdInventoryChannelsRequest struct via the builder pattern


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


## GetApiEventsByEventIdMapDraft

> GetApiEventsByEventIdMapDraft(ctx, eventId).Execute()

Load map draft for designer



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
	r, err := apiClient.EventsAPI.GetApiEventsByEventIdMapDraft(context.Background(), eventId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EventsAPI.GetApiEventsByEventIdMapDraft``: %v\n", err)
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

Other parameters are passed through a pointer to a apiGetApiEventsByEventIdMapDraftRequest struct via the builder pattern


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


## GetApiEventsByEventIdMapVersions

> GetApiEventsByEventIdMapVersions200Response GetApiEventsByEventIdMapVersions(ctx, eventId).Limit(limit).Cursor(cursor).Execute()

List event map versions



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
	limit := float32(8.14) // float32 | Page size (1..100, default 20). (optional)
	cursor := "cursor_example" // string | Opaque `nextCursor` value from the previous page. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EventsAPI.GetApiEventsByEventIdMapVersions(context.Background(), eventId).Limit(limit).Cursor(cursor).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EventsAPI.GetApiEventsByEventIdMapVersions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetApiEventsByEventIdMapVersions`: GetApiEventsByEventIdMapVersions200Response
	fmt.Fprintf(os.Stdout, "Response from `EventsAPI.GetApiEventsByEventIdMapVersions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**eventId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetApiEventsByEventIdMapVersionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **float32** | Page size (1..100, default 20). | 
 **cursor** | **string** | Opaque &#x60;nextCursor&#x60; value from the previous page. | 

### Return type

[**GetApiEventsByEventIdMapVersions200Response**](GetApiEventsByEventIdMapVersions200Response.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, multipart/form-data, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetApiEventsByEventIdMapVersionsByVersionThumbnail

> GetApiEventsByEventIdMapVersionsByVersionThumbnail(ctx, eventId, version).Execute()

Event map version preview



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
	version := float32(8.14) // float32 | Version number (1-based).

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.EventsAPI.GetApiEventsByEventIdMapVersionsByVersionThumbnail(context.Background(), eventId, version).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EventsAPI.GetApiEventsByEventIdMapVersionsByVersionThumbnail``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**eventId** | **string** |  | 
**version** | **float32** | Version number (1-based). | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetApiEventsByEventIdMapVersionsByVersionThumbnailRequest struct via the builder pattern


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


## GetApiEventsByEventIdObjects

> GetApiEventsByEventIdObjects(ctx, eventId).TicketBuyerId(ticketBuyerId).Execute()

List the objects of a ticket buyer



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
	eventId := "eventId_example" // string | Event id (UUID) or event key.
	ticketBuyerId := "ticketBuyerId_example" // string | Opaque ticket buyer id (required). (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.EventsAPI.GetApiEventsByEventIdObjects(context.Background(), eventId).TicketBuyerId(ticketBuyerId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EventsAPI.GetApiEventsByEventIdObjects``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**eventId** | **string** | Event id (UUID) or event key. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetApiEventsByEventIdObjectsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **ticketBuyerId** | **string** | Opaque ticket buyer id (required). | 

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


## GetApiEventsByEventIdObjectsByObjectId

> GetApiEventsByEventIdObjectsByObjectId(ctx, eventId, objectId).Execute()

Retrieve one object (seat or GA area)



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
	objectId := "objectId_example" // string | Seat or GA area id (UUID), or its exact label / GA name.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.EventsAPI.GetApiEventsByEventIdObjectsByObjectId(context.Background(), eventId, objectId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EventsAPI.GetApiEventsByEventIdObjectsByObjectId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**eventId** | **string** |  | 
**objectId** | **string** | Seat or GA area id (UUID), or its exact label / GA name. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetApiEventsByEventIdObjectsByObjectIdRequest struct via the builder pattern


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


## GetApiEventsByEventIdObjectsByObjectIdStatusChanges

> GetApiEventsByEventIdObjectsByObjectIdStatusChanges(ctx, eventId, objectId).Limit(limit).Cursor(cursor).Action(action).OrderId(orderId).Label(label).LabelMatch(labelMatch).Origin(origin).Sort(sort).Execute()

List status changes of one object



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
	objectId := "objectId_example" // string | Seat id (UUID) or exact seat label.
	limit := float32(8.14) // float32 | Page size (1..100, default 20). (optional)
	cursor := "cursor_example" // string | Opaque `nextCursor` value from the previous page. (optional)
	action := "action_example" // string | Exact action: durable `sell`, `cancel`, `booked`, `free`, `resale`, `retired`, a custom status, or Redis-only `hold`, `release`, `hold_expired`. (optional)
	orderId := "orderId_example" // string | Only transitions of this order. (optional)
	label := "label_example" // string | Seat label or GA area name to search (case-insensitive), see `labelMatch`. (optional)
	labelMatch := "labelMatch_example" // string | `exact` (default), `contains`, `prefix` (begins with) or `suffix` (ends with). Exact/prefix/suffix use an index. (optional)
	origin := "origin_example" // string | Only changes of this origin: `api_key`, `dashboard_user`, `buyer_session`, `system_expiry`, or `unknown` (recorded before origins existed). (optional)
	sort := "sort_example" // string | `newest` (default) or `oldest` first. Keep the same sort while following `nextCursor`. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.EventsAPI.GetApiEventsByEventIdObjectsByObjectIdStatusChanges(context.Background(), eventId, objectId).Limit(limit).Cursor(cursor).Action(action).OrderId(orderId).Label(label).LabelMatch(labelMatch).Origin(origin).Sort(sort).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EventsAPI.GetApiEventsByEventIdObjectsByObjectIdStatusChanges``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**eventId** | **string** |  | 
**objectId** | **string** | Seat id (UUID) or exact seat label. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetApiEventsByEventIdObjectsByObjectIdStatusChangesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **limit** | **float32** | Page size (1..100, default 20). | 
 **cursor** | **string** | Opaque &#x60;nextCursor&#x60; value from the previous page. | 
 **action** | **string** | Exact action: durable &#x60;sell&#x60;, &#x60;cancel&#x60;, &#x60;booked&#x60;, &#x60;free&#x60;, &#x60;resale&#x60;, &#x60;retired&#x60;, a custom status, or Redis-only &#x60;hold&#x60;, &#x60;release&#x60;, &#x60;hold_expired&#x60;. | 
 **orderId** | **string** | Only transitions of this order. | 
 **label** | **string** | Seat label or GA area name to search (case-insensitive), see &#x60;labelMatch&#x60;. | 
 **labelMatch** | **string** | &#x60;exact&#x60; (default), &#x60;contains&#x60;, &#x60;prefix&#x60; (begins with) or &#x60;suffix&#x60; (ends with). Exact/prefix/suffix use an index. | 
 **origin** | **string** | Only changes of this origin: &#x60;api_key&#x60;, &#x60;dashboard_user&#x60;, &#x60;buyer_session&#x60;, &#x60;system_expiry&#x60;, or &#x60;unknown&#x60; (recorded before origins existed). | 
 **sort** | **string** | &#x60;newest&#x60; (default) or &#x60;oldest&#x60; first. Keep the same sort while following &#x60;nextCursor&#x60;. | 

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


## GetApiEventsByEventIdPricing

> GetApiEventsByEventIdPricing(ctx, eventId).Execute()

Read the synced price table of an event

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
	r, err := apiClient.EventsAPI.GetApiEventsByEventIdPricing(context.Background(), eventId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EventsAPI.GetApiEventsByEventIdPricing``: %v\n", err)
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

Other parameters are passed through a pointer to a apiGetApiEventsByEventIdPricingRequest struct via the builder pattern


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


## GetApiEventsByEventIdReportsByAvailabilityCsv

> GetApiEventsByEventIdReportsByAvailabilityCsv(ctx, eventId).Availability(availability).LiveHolds(liveHolds).Execute()

Detailed report grouped by availability (available / not_available) (CSV)



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
	availability := "availability_example" // string | Only this availability bucket. (optional)
	liveHolds := "liveHolds_example" // string | Merge active buyer holds from the per-event hold index (default `true`). `false` = Postgres only. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.EventsAPI.GetApiEventsByEventIdReportsByAvailabilityCsv(context.Background(), eventId).Availability(availability).LiveHolds(liveHolds).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EventsAPI.GetApiEventsByEventIdReportsByAvailabilityCsv``: %v\n", err)
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

Other parameters are passed through a pointer to a apiGetApiEventsByEventIdReportsByAvailabilityCsvRequest struct via the builder pattern


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


## GetApiEventsByEventIdReportsByAvailabilityJson

> GetApiEventsByEventIdReportsByAvailabilityJson(ctx, eventId).Availability(availability).LiveHolds(liveHolds).Execute()

Detailed report grouped by availability (available / not_available)



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
	availability := "availability_example" // string | Only this availability bucket. (optional)
	liveHolds := "liveHolds_example" // string | Merge active buyer holds from the per-event hold index (default `true`). `false` = Postgres only. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.EventsAPI.GetApiEventsByEventIdReportsByAvailabilityJson(context.Background(), eventId).Availability(availability).LiveHolds(liveHolds).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EventsAPI.GetApiEventsByEventIdReportsByAvailabilityJson``: %v\n", err)
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

Other parameters are passed through a pointer to a apiGetApiEventsByEventIdReportsByAvailabilityJsonRequest struct via the builder pattern


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


## GetApiEventsByEventIdReportsByAvailabilityReasonCsv

> GetApiEventsByEventIdReportsByAvailabilityReasonCsv(ctx, eventId).Reason(reason).LiveHolds(liveHolds).Execute()

Detailed report grouped by availability reason (CSV)



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
	reason := "reason_example" // string | Only this reason bucket. (optional)
	liveHolds := "liveHolds_example" // string | Merge active buyer holds from the per-event hold index (default `true`). `false` = Postgres only. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.EventsAPI.GetApiEventsByEventIdReportsByAvailabilityReasonCsv(context.Background(), eventId).Reason(reason).LiveHolds(liveHolds).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EventsAPI.GetApiEventsByEventIdReportsByAvailabilityReasonCsv``: %v\n", err)
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

Other parameters are passed through a pointer to a apiGetApiEventsByEventIdReportsByAvailabilityReasonCsvRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **reason** | **string** | Only this reason bucket. | 
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


## GetApiEventsByEventIdReportsByAvailabilityReasonJson

> GetApiEventsByEventIdReportsByAvailabilityReasonJson(ctx, eventId).Reason(reason).LiveHolds(liveHolds).Execute()

Detailed report grouped by availability reason



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
	reason := "reason_example" // string | Only this reason bucket. (optional)
	liveHolds := "liveHolds_example" // string | Merge active buyer holds from the per-event hold index (default `true`). `false` = Postgres only. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.EventsAPI.GetApiEventsByEventIdReportsByAvailabilityReasonJson(context.Background(), eventId).Reason(reason).LiveHolds(liveHolds).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EventsAPI.GetApiEventsByEventIdReportsByAvailabilityReasonJson``: %v\n", err)
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

Other parameters are passed through a pointer to a apiGetApiEventsByEventIdReportsByAvailabilityReasonJsonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **reason** | **string** | Only this reason bucket. | 
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


## GetApiEventsByEventIdReportsByCategoryJson

> GetApiEventsByEventIdReportsByCategoryJson(ctx, eventId).Execute()

Seat and GA counts by category

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
	r, err := apiClient.EventsAPI.GetApiEventsByEventIdReportsByCategoryJson(context.Background(), eventId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EventsAPI.GetApiEventsByEventIdReportsByCategoryJson``: %v\n", err)
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

Other parameters are passed through a pointer to a apiGetApiEventsByEventIdReportsByCategoryJsonRequest struct via the builder pattern


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


## GetApiEventsByEventIdReportsByCategoryKeyCsv

> GetApiEventsByEventIdReportsByCategoryKeyCsv(ctx, eventId).CategoryKey(categoryKey).LiveHolds(liveHolds).Execute()

Detailed report grouped by category key (CSV)



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
	categoryKey := "categoryKey_example" // string | Only this categoryKey bucket. (optional)
	liveHolds := "liveHolds_example" // string | Merge active buyer holds from the per-event hold index (default `true`). `false` = Postgres only. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.EventsAPI.GetApiEventsByEventIdReportsByCategoryKeyCsv(context.Background(), eventId).CategoryKey(categoryKey).LiveHolds(liveHolds).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EventsAPI.GetApiEventsByEventIdReportsByCategoryKeyCsv``: %v\n", err)
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

Other parameters are passed through a pointer to a apiGetApiEventsByEventIdReportsByCategoryKeyCsvRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **categoryKey** | **string** | Only this categoryKey bucket. | 
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


## GetApiEventsByEventIdReportsByCategoryKeyJson

> GetApiEventsByEventIdReportsByCategoryKeyJson(ctx, eventId).CategoryKey(categoryKey).LiveHolds(liveHolds).Execute()

Detailed report grouped by category key



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
	categoryKey := "categoryKey_example" // string | Only this categoryKey bucket. (optional)
	liveHolds := "liveHolds_example" // string | Merge active buyer holds from the per-event hold index (default `true`). `false` = Postgres only. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.EventsAPI.GetApiEventsByEventIdReportsByCategoryKeyJson(context.Background(), eventId).CategoryKey(categoryKey).LiveHolds(liveHolds).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EventsAPI.GetApiEventsByEventIdReportsByCategoryKeyJson``: %v\n", err)
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

Other parameters are passed through a pointer to a apiGetApiEventsByEventIdReportsByCategoryKeyJsonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **categoryKey** | **string** | Only this categoryKey bucket. | 
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


## GetApiEventsByEventIdReportsByCategoryLabelCsv

> GetApiEventsByEventIdReportsByCategoryLabelCsv(ctx, eventId).CategoryLabel(categoryLabel).LiveHolds(liveHolds).Execute()

Detailed report grouped by category label (CSV)



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
	categoryLabel := "categoryLabel_example" // string | Only this categoryLabel bucket. (optional)
	liveHolds := "liveHolds_example" // string | Merge active buyer holds from the per-event hold index (default `true`). `false` = Postgres only. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.EventsAPI.GetApiEventsByEventIdReportsByCategoryLabelCsv(context.Background(), eventId).CategoryLabel(categoryLabel).LiveHolds(liveHolds).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EventsAPI.GetApiEventsByEventIdReportsByCategoryLabelCsv``: %v\n", err)
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

Other parameters are passed through a pointer to a apiGetApiEventsByEventIdReportsByCategoryLabelCsvRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **categoryLabel** | **string** | Only this categoryLabel bucket. | 
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


## GetApiEventsByEventIdReportsByCategoryLabelJson

> GetApiEventsByEventIdReportsByCategoryLabelJson(ctx, eventId).CategoryLabel(categoryLabel).LiveHolds(liveHolds).Execute()

Detailed report grouped by category label



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
	categoryLabel := "categoryLabel_example" // string | Only this categoryLabel bucket. (optional)
	liveHolds := "liveHolds_example" // string | Merge active buyer holds from the per-event hold index (default `true`). `false` = Postgres only. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.EventsAPI.GetApiEventsByEventIdReportsByCategoryLabelJson(context.Background(), eventId).CategoryLabel(categoryLabel).LiveHolds(liveHolds).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EventsAPI.GetApiEventsByEventIdReportsByCategoryLabelJson``: %v\n", err)
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

Other parameters are passed through a pointer to a apiGetApiEventsByEventIdReportsByCategoryLabelJsonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **categoryLabel** | **string** | Only this categoryLabel bucket. | 
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


## GetApiEventsByEventIdReportsByChannelJson

> GetApiEventsByEventIdReportsByChannelJson(ctx, eventId).ChannelKeyId(channelKeyId).Execute()

Detailed report grouped by inventory channel



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
	channelKeyId := "channelKeyId_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.EventsAPI.GetApiEventsByEventIdReportsByChannelJson(context.Background(), eventId).ChannelKeyId(channelKeyId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EventsAPI.GetApiEventsByEventIdReportsByChannelJson``: %v\n", err)
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

Other parameters are passed through a pointer to a apiGetApiEventsByEventIdReportsByChannelJsonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **channelKeyId** | **string** |  | 

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


## GetApiEventsByEventIdReportsByLabelJson

> GetApiEventsByEventIdReportsByLabelJson(ctx, eventId).Label(label).Execute()

Detailed report grouped by object label (Postgres)



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
	label := "label_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.EventsAPI.GetApiEventsByEventIdReportsByLabelJson(context.Background(), eventId).Label(label).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EventsAPI.GetApiEventsByEventIdReportsByLabelJson``: %v\n", err)
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

Other parameters are passed through a pointer to a apiGetApiEventsByEventIdReportsByLabelJsonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **label** | **string** |  | 

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


## GetApiEventsByEventIdReportsByObjectTypeJson

> GetApiEventsByEventIdReportsByObjectTypeJson(ctx, eventId).Execute()

Detailed report grouped by object type (seat|ga)

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
	r, err := apiClient.EventsAPI.GetApiEventsByEventIdReportsByObjectTypeJson(context.Background(), eventId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EventsAPI.GetApiEventsByEventIdReportsByObjectTypeJson``: %v\n", err)
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

Other parameters are passed through a pointer to a apiGetApiEventsByEventIdReportsByObjectTypeJsonRequest struct via the builder pattern


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


## GetApiEventsByEventIdReportsByOrderIdJson

> GetApiEventsByEventIdReportsByOrderIdJson(ctx, eventId).OrderId(orderId).Execute()

Detailed report grouped by orderId (Postgres)



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
	orderId := "orderId_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.EventsAPI.GetApiEventsByEventIdReportsByOrderIdJson(context.Background(), eventId).OrderId(orderId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EventsAPI.GetApiEventsByEventIdReportsByOrderIdJson``: %v\n", err)
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

Other parameters are passed through a pointer to a apiGetApiEventsByEventIdReportsByOrderIdJsonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **orderId** | **string** |  | 

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


## GetApiEventsByEventIdReportsBySectionJson

> GetApiEventsByEventIdReportsBySectionJson(ctx, eventId).Section(section).Execute()

Detailed report grouped by section



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
	section := "section_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.EventsAPI.GetApiEventsByEventIdReportsBySectionJson(context.Background(), eventId).Section(section).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EventsAPI.GetApiEventsByEventIdReportsBySectionJson``: %v\n", err)
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

Other parameters are passed through a pointer to a apiGetApiEventsByEventIdReportsBySectionJsonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **section** | **string** |  | 

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


## GetApiEventsByEventIdReportsByStatusJson

> GetApiEventsByEventIdReportsByStatusJson(ctx, eventId).Execute()

Seat counts by Postgres status

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
	r, err := apiClient.EventsAPI.GetApiEventsByEventIdReportsByStatusJson(context.Background(), eventId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EventsAPI.GetApiEventsByEventIdReportsByStatusJson``: %v\n", err)
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

Other parameters are passed through a pointer to a apiGetApiEventsByEventIdReportsByStatusJsonRequest struct via the builder pattern


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


## GetApiEventsByEventIdReportsByZoneCsv

> GetApiEventsByEventIdReportsByZoneCsv(ctx, eventId).Zone(zone).LiveHolds(liveHolds).Execute()

Detailed report grouped by zone (CSV)



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
	zone := "zone_example" // string | Only this zone bucket. (optional)
	liveHolds := "liveHolds_example" // string | Merge active buyer holds from the per-event hold index (default `true`). `false` = Postgres only. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.EventsAPI.GetApiEventsByEventIdReportsByZoneCsv(context.Background(), eventId).Zone(zone).LiveHolds(liveHolds).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EventsAPI.GetApiEventsByEventIdReportsByZoneCsv``: %v\n", err)
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

Other parameters are passed through a pointer to a apiGetApiEventsByEventIdReportsByZoneCsvRequest struct via the builder pattern


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


## GetApiEventsByEventIdReportsByZoneJson

> GetApiEventsByEventIdReportsByZoneJson(ctx, eventId).Zone(zone).LiveHolds(liveHolds).Execute()

Detailed report grouped by zone



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
	zone := "zone_example" // string | Only this zone bucket. (optional)
	liveHolds := "liveHolds_example" // string | Merge active buyer holds from the per-event hold index (default `true`). `false` = Postgres only. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.EventsAPI.GetApiEventsByEventIdReportsByZoneJson(context.Background(), eventId).Zone(zone).LiveHolds(liveHolds).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EventsAPI.GetApiEventsByEventIdReportsByZoneJson``: %v\n", err)
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

Other parameters are passed through a pointer to a apiGetApiEventsByEventIdReportsByZoneJsonRequest struct via the builder pattern


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


## GetApiEventsByEventIdReportsDeepSummaryJson

> GetApiEventsByEventIdReportsDeepSummaryJson(ctx, eventId).Execute()

Deep summary: status counts nested by category

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
	r, err := apiClient.EventsAPI.GetApiEventsByEventIdReportsDeepSummaryJson(context.Background(), eventId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EventsAPI.GetApiEventsByEventIdReportsDeepSummaryJson``: %v\n", err)
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

Other parameters are passed through a pointer to a apiGetApiEventsByEventIdReportsDeepSummaryJsonRequest struct via the builder pattern


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


## GetApiEventsByEventIdReportsOccupancyCsv

> GetApiEventsByEventIdReportsOccupancyCsv(ctx, eventId).Execute()

Export seat and GA occupancy as CSV



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
	r, err := apiClient.EventsAPI.GetApiEventsByEventIdReportsOccupancyCsv(context.Background(), eventId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EventsAPI.GetApiEventsByEventIdReportsOccupancyCsv``: %v\n", err)
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

Other parameters are passed through a pointer to a apiGetApiEventsByEventIdReportsOccupancyCsvRequest struct via the builder pattern


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


## GetApiEventsByEventIdReportsSalesCsv

> GetApiEventsByEventIdReportsSalesCsv(ctx, eventId).Execute()

Export confirmed seat and GA sales as CSV



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
	r, err := apiClient.EventsAPI.GetApiEventsByEventIdReportsSalesCsv(context.Background(), eventId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EventsAPI.GetApiEventsByEventIdReportsSalesCsv``: %v\n", err)
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

Other parameters are passed through a pointer to a apiGetApiEventsByEventIdReportsSalesCsvRequest struct via the builder pattern


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


## GetApiEventsByEventIdReportsSummaryByAvailabilityCsv

> GetApiEventsByEventIdReportsSummaryByAvailabilityCsv(ctx, eventId).LiveHolds(liveHolds).Execute()

Summary report by availability (CSV)



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
	liveHolds := "liveHolds_example" // string | Merge active buyer holds from the per-event hold index (default `true`). `false` = Postgres only. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.EventsAPI.GetApiEventsByEventIdReportsSummaryByAvailabilityCsv(context.Background(), eventId).LiveHolds(liveHolds).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EventsAPI.GetApiEventsByEventIdReportsSummaryByAvailabilityCsv``: %v\n", err)
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

Other parameters are passed through a pointer to a apiGetApiEventsByEventIdReportsSummaryByAvailabilityCsvRequest struct via the builder pattern


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


## GetApiEventsByEventIdReportsSummaryByAvailabilityJson

> GetApiEventsByEventIdReportsSummaryByAvailabilityJson(ctx, eventId).LiveHolds(liveHolds).Execute()

Summary report by availability



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
	liveHolds := "liveHolds_example" // string | Merge active buyer holds from the per-event hold index (default `true`). `false` = Postgres only. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.EventsAPI.GetApiEventsByEventIdReportsSummaryByAvailabilityJson(context.Background(), eventId).LiveHolds(liveHolds).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EventsAPI.GetApiEventsByEventIdReportsSummaryByAvailabilityJson``: %v\n", err)
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

Other parameters are passed through a pointer to a apiGetApiEventsByEventIdReportsSummaryByAvailabilityJsonRequest struct via the builder pattern


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


## GetApiEventsByEventIdReportsSummaryByAvailabilityReasonCsv

> GetApiEventsByEventIdReportsSummaryByAvailabilityReasonCsv(ctx, eventId).LiveHolds(liveHolds).Execute()

Summary report by availability reason (CSV)



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
	liveHolds := "liveHolds_example" // string | Merge active buyer holds from the per-event hold index (default `true`). `false` = Postgres only. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.EventsAPI.GetApiEventsByEventIdReportsSummaryByAvailabilityReasonCsv(context.Background(), eventId).LiveHolds(liveHolds).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EventsAPI.GetApiEventsByEventIdReportsSummaryByAvailabilityReasonCsv``: %v\n", err)
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

Other parameters are passed through a pointer to a apiGetApiEventsByEventIdReportsSummaryByAvailabilityReasonCsvRequest struct via the builder pattern


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


## GetApiEventsByEventIdReportsSummaryByAvailabilityReasonJson

> GetApiEventsByEventIdReportsSummaryByAvailabilityReasonJson(ctx, eventId).LiveHolds(liveHolds).Execute()

Summary report by availability reason



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
	liveHolds := "liveHolds_example" // string | Merge active buyer holds from the per-event hold index (default `true`). `false` = Postgres only. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.EventsAPI.GetApiEventsByEventIdReportsSummaryByAvailabilityReasonJson(context.Background(), eventId).LiveHolds(liveHolds).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EventsAPI.GetApiEventsByEventIdReportsSummaryByAvailabilityReasonJson``: %v\n", err)
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

Other parameters are passed through a pointer to a apiGetApiEventsByEventIdReportsSummaryByAvailabilityReasonJsonRequest struct via the builder pattern


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


## GetApiEventsByEventIdReportsSummaryByCategoryCsv

> GetApiEventsByEventIdReportsSummaryByCategoryCsv(ctx, eventId).LiveHolds(liveHolds).Execute()

Summary report by category (CSV)



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
	liveHolds := "liveHolds_example" // string | Merge active buyer holds from the per-event hold index (default `true`). `false` = Postgres only. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.EventsAPI.GetApiEventsByEventIdReportsSummaryByCategoryCsv(context.Background(), eventId).LiveHolds(liveHolds).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EventsAPI.GetApiEventsByEventIdReportsSummaryByCategoryCsv``: %v\n", err)
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

Other parameters are passed through a pointer to a apiGetApiEventsByEventIdReportsSummaryByCategoryCsvRequest struct via the builder pattern


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


## GetApiEventsByEventIdReportsSummaryByCategoryJson

> GetApiEventsByEventIdReportsSummaryByCategoryJson(ctx, eventId).LiveHolds(liveHolds).Execute()

Summary report by category



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
	liveHolds := "liveHolds_example" // string | Merge active buyer holds from the per-event hold index (default `true`). `false` = Postgres only. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.EventsAPI.GetApiEventsByEventIdReportsSummaryByCategoryJson(context.Background(), eventId).LiveHolds(liveHolds).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EventsAPI.GetApiEventsByEventIdReportsSummaryByCategoryJson``: %v\n", err)
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

Other parameters are passed through a pointer to a apiGetApiEventsByEventIdReportsSummaryByCategoryJsonRequest struct via the builder pattern


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


## GetApiEventsByEventIdReportsSummaryByCategoryKeyCsv

> GetApiEventsByEventIdReportsSummaryByCategoryKeyCsv(ctx, eventId).LiveHolds(liveHolds).Execute()

Summary report by category key (CSV)



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
	liveHolds := "liveHolds_example" // string | Merge active buyer holds from the per-event hold index (default `true`). `false` = Postgres only. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.EventsAPI.GetApiEventsByEventIdReportsSummaryByCategoryKeyCsv(context.Background(), eventId).LiveHolds(liveHolds).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EventsAPI.GetApiEventsByEventIdReportsSummaryByCategoryKeyCsv``: %v\n", err)
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

Other parameters are passed through a pointer to a apiGetApiEventsByEventIdReportsSummaryByCategoryKeyCsvRequest struct via the builder pattern


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


## GetApiEventsByEventIdReportsSummaryByCategoryKeyJson

> GetApiEventsByEventIdReportsSummaryByCategoryKeyJson(ctx, eventId).LiveHolds(liveHolds).Execute()

Summary report by category key



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
	liveHolds := "liveHolds_example" // string | Merge active buyer holds from the per-event hold index (default `true`). `false` = Postgres only. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.EventsAPI.GetApiEventsByEventIdReportsSummaryByCategoryKeyJson(context.Background(), eventId).LiveHolds(liveHolds).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EventsAPI.GetApiEventsByEventIdReportsSummaryByCategoryKeyJson``: %v\n", err)
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

Other parameters are passed through a pointer to a apiGetApiEventsByEventIdReportsSummaryByCategoryKeyJsonRequest struct via the builder pattern


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


## GetApiEventsByEventIdReportsSummaryByCategoryLabelCsv

> GetApiEventsByEventIdReportsSummaryByCategoryLabelCsv(ctx, eventId).LiveHolds(liveHolds).Execute()

Summary report by category label (CSV)



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
	liveHolds := "liveHolds_example" // string | Merge active buyer holds from the per-event hold index (default `true`). `false` = Postgres only. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.EventsAPI.GetApiEventsByEventIdReportsSummaryByCategoryLabelCsv(context.Background(), eventId).LiveHolds(liveHolds).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EventsAPI.GetApiEventsByEventIdReportsSummaryByCategoryLabelCsv``: %v\n", err)
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

Other parameters are passed through a pointer to a apiGetApiEventsByEventIdReportsSummaryByCategoryLabelCsvRequest struct via the builder pattern


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


## GetApiEventsByEventIdReportsSummaryByCategoryLabelJson

> GetApiEventsByEventIdReportsSummaryByCategoryLabelJson(ctx, eventId).LiveHolds(liveHolds).Execute()

Summary report by category label



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
	liveHolds := "liveHolds_example" // string | Merge active buyer holds from the per-event hold index (default `true`). `false` = Postgres only. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.EventsAPI.GetApiEventsByEventIdReportsSummaryByCategoryLabelJson(context.Background(), eventId).LiveHolds(liveHolds).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EventsAPI.GetApiEventsByEventIdReportsSummaryByCategoryLabelJson``: %v\n", err)
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

Other parameters are passed through a pointer to a apiGetApiEventsByEventIdReportsSummaryByCategoryLabelJsonRequest struct via the builder pattern


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


## GetApiEventsByEventIdReportsSummaryByChannelCsv

> GetApiEventsByEventIdReportsSummaryByChannelCsv(ctx, eventId).LiveHolds(liveHolds).Execute()

Summary report by channel (CSV)



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
	liveHolds := "liveHolds_example" // string | Merge active buyer holds from the per-event hold index (default `true`). `false` = Postgres only. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.EventsAPI.GetApiEventsByEventIdReportsSummaryByChannelCsv(context.Background(), eventId).LiveHolds(liveHolds).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EventsAPI.GetApiEventsByEventIdReportsSummaryByChannelCsv``: %v\n", err)
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

Other parameters are passed through a pointer to a apiGetApiEventsByEventIdReportsSummaryByChannelCsvRequest struct via the builder pattern


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


## GetApiEventsByEventIdReportsSummaryByChannelJson

> GetApiEventsByEventIdReportsSummaryByChannelJson(ctx, eventId).LiveHolds(liveHolds).Execute()

Summary report by channel



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
	liveHolds := "liveHolds_example" // string | Merge active buyer holds from the per-event hold index (default `true`). `false` = Postgres only. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.EventsAPI.GetApiEventsByEventIdReportsSummaryByChannelJson(context.Background(), eventId).LiveHolds(liveHolds).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EventsAPI.GetApiEventsByEventIdReportsSummaryByChannelJson``: %v\n", err)
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

Other parameters are passed through a pointer to a apiGetApiEventsByEventIdReportsSummaryByChannelJsonRequest struct via the builder pattern


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


## GetApiEventsByEventIdReportsSummaryByObjectTypeCsv

> GetApiEventsByEventIdReportsSummaryByObjectTypeCsv(ctx, eventId).LiveHolds(liveHolds).Execute()

Summary report by object type (CSV)



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
	liveHolds := "liveHolds_example" // string | Merge active buyer holds from the per-event hold index (default `true`). `false` = Postgres only. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.EventsAPI.GetApiEventsByEventIdReportsSummaryByObjectTypeCsv(context.Background(), eventId).LiveHolds(liveHolds).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EventsAPI.GetApiEventsByEventIdReportsSummaryByObjectTypeCsv``: %v\n", err)
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

Other parameters are passed through a pointer to a apiGetApiEventsByEventIdReportsSummaryByObjectTypeCsvRequest struct via the builder pattern


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


## GetApiEventsByEventIdReportsSummaryByObjectTypeJson

> GetApiEventsByEventIdReportsSummaryByObjectTypeJson(ctx, eventId).LiveHolds(liveHolds).Execute()

Summary report by object type



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
	liveHolds := "liveHolds_example" // string | Merge active buyer holds from the per-event hold index (default `true`). `false` = Postgres only. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.EventsAPI.GetApiEventsByEventIdReportsSummaryByObjectTypeJson(context.Background(), eventId).LiveHolds(liveHolds).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EventsAPI.GetApiEventsByEventIdReportsSummaryByObjectTypeJson``: %v\n", err)
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

Other parameters are passed through a pointer to a apiGetApiEventsByEventIdReportsSummaryByObjectTypeJsonRequest struct via the builder pattern


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


## GetApiEventsByEventIdReportsSummaryBySectionCsv

> GetApiEventsByEventIdReportsSummaryBySectionCsv(ctx, eventId).LiveHolds(liveHolds).Execute()

Summary report by section (CSV)



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
	liveHolds := "liveHolds_example" // string | Merge active buyer holds from the per-event hold index (default `true`). `false` = Postgres only. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.EventsAPI.GetApiEventsByEventIdReportsSummaryBySectionCsv(context.Background(), eventId).LiveHolds(liveHolds).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EventsAPI.GetApiEventsByEventIdReportsSummaryBySectionCsv``: %v\n", err)
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

Other parameters are passed through a pointer to a apiGetApiEventsByEventIdReportsSummaryBySectionCsvRequest struct via the builder pattern


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


## GetApiEventsByEventIdReportsSummaryBySectionJson

> GetApiEventsByEventIdReportsSummaryBySectionJson(ctx, eventId).LiveHolds(liveHolds).Execute()

Summary report by section



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
	liveHolds := "liveHolds_example" // string | Merge active buyer holds from the per-event hold index (default `true`). `false` = Postgres only. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.EventsAPI.GetApiEventsByEventIdReportsSummaryBySectionJson(context.Background(), eventId).LiveHolds(liveHolds).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EventsAPI.GetApiEventsByEventIdReportsSummaryBySectionJson``: %v\n", err)
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

Other parameters are passed through a pointer to a apiGetApiEventsByEventIdReportsSummaryBySectionJsonRequest struct via the builder pattern


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


## GetApiEventsByEventIdReportsSummaryByStatusCsv

> GetApiEventsByEventIdReportsSummaryByStatusCsv(ctx, eventId).LiveHolds(liveHolds).Execute()

Summary report by status (CSV)



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
	liveHolds := "liveHolds_example" // string | Merge active buyer holds from the per-event hold index (default `true`). `false` = Postgres only. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.EventsAPI.GetApiEventsByEventIdReportsSummaryByStatusCsv(context.Background(), eventId).LiveHolds(liveHolds).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EventsAPI.GetApiEventsByEventIdReportsSummaryByStatusCsv``: %v\n", err)
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

Other parameters are passed through a pointer to a apiGetApiEventsByEventIdReportsSummaryByStatusCsvRequest struct via the builder pattern


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


## GetApiEventsByEventIdReportsSummaryByStatusJson

> GetApiEventsByEventIdReportsSummaryByStatusJson(ctx, eventId).LiveHolds(liveHolds).Execute()

Summary report by status



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
	liveHolds := "liveHolds_example" // string | Merge active buyer holds from the per-event hold index (default `true`). `false` = Postgres only. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.EventsAPI.GetApiEventsByEventIdReportsSummaryByStatusJson(context.Background(), eventId).LiveHolds(liveHolds).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EventsAPI.GetApiEventsByEventIdReportsSummaryByStatusJson``: %v\n", err)
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

Other parameters are passed through a pointer to a apiGetApiEventsByEventIdReportsSummaryByStatusJsonRequest struct via the builder pattern


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


## GetApiEventsByEventIdReportsSummaryByZoneCsv

> GetApiEventsByEventIdReportsSummaryByZoneCsv(ctx, eventId).LiveHolds(liveHolds).Execute()

Summary report by zone (CSV)



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
	liveHolds := "liveHolds_example" // string | Merge active buyer holds from the per-event hold index (default `true`). `false` = Postgres only. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.EventsAPI.GetApiEventsByEventIdReportsSummaryByZoneCsv(context.Background(), eventId).LiveHolds(liveHolds).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EventsAPI.GetApiEventsByEventIdReportsSummaryByZoneCsv``: %v\n", err)
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

Other parameters are passed through a pointer to a apiGetApiEventsByEventIdReportsSummaryByZoneCsvRequest struct via the builder pattern


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


## GetApiEventsByEventIdReportsSummaryByZoneJson

> GetApiEventsByEventIdReportsSummaryByZoneJson(ctx, eventId).LiveHolds(liveHolds).Execute()

Summary report by zone



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
	liveHolds := "liveHolds_example" // string | Merge active buyer holds from the per-event hold index (default `true`). `false` = Postgres only. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.EventsAPI.GetApiEventsByEventIdReportsSummaryByZoneJson(context.Background(), eventId).LiveHolds(liveHolds).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EventsAPI.GetApiEventsByEventIdReportsSummaryByZoneJson``: %v\n", err)
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

Other parameters are passed through a pointer to a apiGetApiEventsByEventIdReportsSummaryByZoneJsonRequest struct via the builder pattern


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


## GetApiEventsByEventIdReportsSummaryJson

> GetApiEventsByEventIdReportsSummaryJson(ctx, eventId).Execute()

Occupancy summary JSON (Postgres only)



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
	r, err := apiClient.EventsAPI.GetApiEventsByEventIdReportsSummaryJson(context.Background(), eventId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EventsAPI.GetApiEventsByEventIdReportsSummaryJson``: %v\n", err)
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

Other parameters are passed through a pointer to a apiGetApiEventsByEventIdReportsSummaryJsonRequest struct via the builder pattern


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


## GetApiEventsByEventIdSalesCounters

> GetApiEventsByEventIdSalesCounters(ctx, eventId).Execute()

Live sales counters (quotas, lots) of an event



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
	r, err := apiClient.EventsAPI.GetApiEventsByEventIdSalesCounters(context.Background(), eventId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EventsAPI.GetApiEventsByEventIdSalesCounters``: %v\n", err)
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

Other parameters are passed through a pointer to a apiGetApiEventsByEventIdSalesCountersRequest struct via the builder pattern


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


## GetApiEventsByEventIdSalesRules

> GetApiEventsByEventIdSalesRules(ctx, eventId).Execute()

Read the sales rules of an event



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
	r, err := apiClient.EventsAPI.GetApiEventsByEventIdSalesRules(context.Background(), eventId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EventsAPI.GetApiEventsByEventIdSalesRules``: %v\n", err)
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

Other parameters are passed through a pointer to a apiGetApiEventsByEventIdSalesRulesRequest struct via the builder pattern


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


## GetApiEventsByEventIdSeatViews

> GetApiEventsByEventIdSeatViews200Response GetApiEventsByEventIdSeatViews(ctx, eventId).Execute()

List view-from-seat image and optional 360 URLs



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
	resp, r, err := apiClient.EventsAPI.GetApiEventsByEventIdSeatViews(context.Background(), eventId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EventsAPI.GetApiEventsByEventIdSeatViews``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetApiEventsByEventIdSeatViews`: GetApiEventsByEventIdSeatViews200Response
	fmt.Fprintf(os.Stdout, "Response from `EventsAPI.GetApiEventsByEventIdSeatViews`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**eventId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetApiEventsByEventIdSeatViewsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetApiEventsByEventIdSeatViews200Response**](GetApiEventsByEventIdSeatViews200Response.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, multipart/form-data, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetApiEventsByEventIdSeats

> GetApiEventsByEventIdSeats(ctx, eventId).Execute()

List seats for an event



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
	r, err := apiClient.EventsAPI.GetApiEventsByEventIdSeats(context.Background(), eventId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EventsAPI.GetApiEventsByEventIdSeats``: %v\n", err)
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

Other parameters are passed through a pointer to a apiGetApiEventsByEventIdSeatsRequest struct via the builder pattern


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


## GetApiEventsByEventIdState

> GetApiEventsByEventIdState(ctx, eventId).Execute()

Compact visual state (protocol 5)



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
	r, err := apiClient.EventsAPI.GetApiEventsByEventIdState(context.Background(), eventId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EventsAPI.GetApiEventsByEventIdState``: %v\n", err)
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

Other parameters are passed through a pointer to a apiGetApiEventsByEventIdStateRequest struct via the builder pattern


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


## GetApiEventsByEventIdStatusChanges

> GetApiEventsByEventIdStatusChanges(ctx, eventId).Limit(limit).Cursor(cursor).Action(action).OrderId(orderId).Label(label).LabelMatch(labelMatch).Origin(origin).Sort(sort).ObjectId(objectId).Execute()

List status changes of an event



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
	limit := float32(8.14) // float32 | Page size (1..100, default 20). (optional)
	cursor := "cursor_example" // string | Opaque `nextCursor` value from the previous page. (optional)
	action := "action_example" // string | Exact action: durable `sell`, `cancel`, `booked`, `free`, `resale`, `retired`, a custom status, or Redis-only `hold`, `release`, `hold_expired`. (optional)
	orderId := "orderId_example" // string | Only transitions of this order. (optional)
	label := "label_example" // string | Seat label or GA area name to search (case-insensitive), see `labelMatch`. (optional)
	labelMatch := "labelMatch_example" // string | `exact` (default), `contains`, `prefix` (begins with) or `suffix` (ends with). Exact/prefix/suffix use an index. (optional)
	origin := "origin_example" // string | Only changes of this origin: `api_key`, `dashboard_user`, `buyer_session`, `system_expiry`, or `unknown` (recorded before origins existed). (optional)
	sort := "sort_example" // string | `newest` (default) or `oldest` first. Keep the same sort while following `nextCursor`. (optional)
	objectId := "objectId_example" // string | Seat id (UUID) or exact seat label. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.EventsAPI.GetApiEventsByEventIdStatusChanges(context.Background(), eventId).Limit(limit).Cursor(cursor).Action(action).OrderId(orderId).Label(label).LabelMatch(labelMatch).Origin(origin).Sort(sort).ObjectId(objectId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EventsAPI.GetApiEventsByEventIdStatusChanges``: %v\n", err)
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

Other parameters are passed through a pointer to a apiGetApiEventsByEventIdStatusChangesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **float32** | Page size (1..100, default 20). | 
 **cursor** | **string** | Opaque &#x60;nextCursor&#x60; value from the previous page. | 
 **action** | **string** | Exact action: durable &#x60;sell&#x60;, &#x60;cancel&#x60;, &#x60;booked&#x60;, &#x60;free&#x60;, &#x60;resale&#x60;, &#x60;retired&#x60;, a custom status, or Redis-only &#x60;hold&#x60;, &#x60;release&#x60;, &#x60;hold_expired&#x60;. | 
 **orderId** | **string** | Only transitions of this order. | 
 **label** | **string** | Seat label or GA area name to search (case-insensitive), see &#x60;labelMatch&#x60;. | 
 **labelMatch** | **string** | &#x60;exact&#x60; (default), &#x60;contains&#x60;, &#x60;prefix&#x60; (begins with) or &#x60;suffix&#x60; (ends with). Exact/prefix/suffix use an index. | 
 **origin** | **string** | Only changes of this origin: &#x60;api_key&#x60;, &#x60;dashboard_user&#x60;, &#x60;buyer_session&#x60;, &#x60;system_expiry&#x60;, or &#x60;unknown&#x60; (recorded before origins existed). | 
 **sort** | **string** | &#x60;newest&#x60; (default) or &#x60;oldest&#x60; first. Keep the same sort while following &#x60;nextCursor&#x60;. | 
 **objectId** | **string** | Seat id (UUID) or exact seat label. | 

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


## GetApiEventsByEventIdZones

> GetApiEventsByEventIdZones(ctx, eventId).Execute()

List the zones of an event



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
	eventId := "eventId_example" // string | Event id (UUID) or event key.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.EventsAPI.GetApiEventsByEventIdZones(context.Background(), eventId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EventsAPI.GetApiEventsByEventIdZones``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**eventId** | **string** | Event id (UUID) or event key. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetApiEventsByEventIdZonesRequest struct via the builder pattern


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


## PatchApiEventsByEventId

> PatchApiEventsByEventId(ctx, eventId).PatchApiEventsByEventIdRequest(patchApiEventsByEventIdRequest).IdempotencyKey(idempotencyKey).Execute()

Patch event: name, date, currency, overrides (isInThePast is irreversible)



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
	patchApiEventsByEventIdRequest := *openapiclient.NewPatchApiEventsByEventIdRequest() // PatchApiEventsByEventIdRequest | 
	idempotencyKey := "idempotencyKey_example" // string | Required on transactional routes. Reuse the same key for every retry of one logical operation; the API replays the stored outcome. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.EventsAPI.PatchApiEventsByEventId(context.Background(), eventId).PatchApiEventsByEventIdRequest(patchApiEventsByEventIdRequest).IdempotencyKey(idempotencyKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EventsAPI.PatchApiEventsByEventId``: %v\n", err)
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

Other parameters are passed through a pointer to a apiPatchApiEventsByEventIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **patchApiEventsByEventIdRequest** | [**PatchApiEventsByEventIdRequest**](PatchApiEventsByEventIdRequest.md) |  | 
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


## PatchApiEventsByEventIdInventoryChannelsByChannelKeyId

> PatchApiEventsByEventIdInventoryChannelsByChannelKeyId(ctx, eventId, channelKeyId).PatchApiEventsByEventIdInventoryChannelsByChannelKeyIdRequest(patchApiEventsByEventIdInventoryChannelsByChannelKeyIdRequest).IdempotencyKey(idempotencyKey).Execute()

Rename or recolor an inventory channel



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
	channelKeyId := "channelKeyId_example" // string | 
	patchApiEventsByEventIdInventoryChannelsByChannelKeyIdRequest := *openapiclient.NewPatchApiEventsByEventIdInventoryChannelsByChannelKeyIdRequest() // PatchApiEventsByEventIdInventoryChannelsByChannelKeyIdRequest | 
	idempotencyKey := "idempotencyKey_example" // string | Required on transactional routes. Reuse the same key for every retry of one logical operation; the API replays the stored outcome. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.EventsAPI.PatchApiEventsByEventIdInventoryChannelsByChannelKeyId(context.Background(), eventId, channelKeyId).PatchApiEventsByEventIdInventoryChannelsByChannelKeyIdRequest(patchApiEventsByEventIdInventoryChannelsByChannelKeyIdRequest).IdempotencyKey(idempotencyKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EventsAPI.PatchApiEventsByEventIdInventoryChannelsByChannelKeyId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**eventId** | **string** |  | 
**channelKeyId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchApiEventsByEventIdInventoryChannelsByChannelKeyIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **patchApiEventsByEventIdInventoryChannelsByChannelKeyIdRequest** | [**PatchApiEventsByEventIdInventoryChannelsByChannelKeyIdRequest**](PatchApiEventsByEventIdInventoryChannelsByChannelKeyIdRequest.md) |  | 
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


## PostApiEvents

> PostApiEvents(ctx).PostApiEventsRequest(postApiEventsRequest).IdempotencyKey(idempotencyKey).Execute()

Create event



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
	postApiEventsRequest := *openapiclient.NewPostApiEventsRequest("Name_example") // PostApiEventsRequest | 
	idempotencyKey := "idempotencyKey_example" // string | Required on transactional routes. Reuse the same key for every retry of one logical operation; the API replays the stored outcome. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.EventsAPI.PostApiEvents(context.Background()).PostApiEventsRequest(postApiEventsRequest).IdempotencyKey(idempotencyKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EventsAPI.PostApiEvents``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostApiEventsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **postApiEventsRequest** | [**PostApiEventsRequest**](PostApiEventsRequest.md) |  | 
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


## PostApiEventsBulk

> PostApiEventsBulk(ctx).PostApiEventsBulkRequest(postApiEventsBulkRequest).IdempotencyKey(idempotencyKey).Execute()

Create up to 50 events for one chart



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
	postApiEventsBulkRequest := *openapiclient.NewPostApiEventsBulkRequest("ChartId_example", []openapiclient.PostApiEventsBulkRequestEventsInner{*openapiclient.NewPostApiEventsBulkRequestEventsInner("Name_example")}) // PostApiEventsBulkRequest | 
	idempotencyKey := "idempotencyKey_example" // string | Required on transactional routes. Reuse the same key for every retry of one logical operation; the API replays the stored outcome. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.EventsAPI.PostApiEventsBulk(context.Background()).PostApiEventsBulkRequest(postApiEventsBulkRequest).IdempotencyKey(idempotencyKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EventsAPI.PostApiEventsBulk``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostApiEventsBulkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **postApiEventsBulkRequest** | [**PostApiEventsBulkRequest**](PostApiEventsBulkRequest.md) |  | 
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


## PostApiEventsByEventIdActionsMarkEverythingAsForSale

> PostApiEventsByEventIdActionsMarkEverythingAsForSale(ctx, eventId).PostApiEventsByEventIdActionsMarkEverythingAsNotForSaleRequest(postApiEventsByEventIdActionsMarkEverythingAsNotForSaleRequest).IdempotencyKey(idempotencyKey).Execute()

Put every object of the event back on sale



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
	postApiEventsByEventIdActionsMarkEverythingAsNotForSaleRequest := *openapiclient.NewPostApiEventsByEventIdActionsMarkEverythingAsNotForSaleRequest() // PostApiEventsByEventIdActionsMarkEverythingAsNotForSaleRequest | 
	idempotencyKey := "idempotencyKey_example" // string | Required on transactional routes. Reuse the same key for every retry of one logical operation; the API replays the stored outcome. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.EventsAPI.PostApiEventsByEventIdActionsMarkEverythingAsForSale(context.Background(), eventId).PostApiEventsByEventIdActionsMarkEverythingAsNotForSaleRequest(postApiEventsByEventIdActionsMarkEverythingAsNotForSaleRequest).IdempotencyKey(idempotencyKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EventsAPI.PostApiEventsByEventIdActionsMarkEverythingAsForSale``: %v\n", err)
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

Other parameters are passed through a pointer to a apiPostApiEventsByEventIdActionsMarkEverythingAsForSaleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **postApiEventsByEventIdActionsMarkEverythingAsNotForSaleRequest** | [**PostApiEventsByEventIdActionsMarkEverythingAsNotForSaleRequest**](PostApiEventsByEventIdActionsMarkEverythingAsNotForSaleRequest.md) |  | 
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


## PostApiEventsByEventIdActionsMarkEverythingAsNotForSale

> PostApiEventsByEventIdActionsMarkEverythingAsNotForSale(ctx, eventId).PostApiEventsByEventIdActionsMarkEverythingAsNotForSaleRequest(postApiEventsByEventIdActionsMarkEverythingAsNotForSaleRequest).IdempotencyKey(idempotencyKey).Execute()

Mark every object of the event not for sale



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
	postApiEventsByEventIdActionsMarkEverythingAsNotForSaleRequest := *openapiclient.NewPostApiEventsByEventIdActionsMarkEverythingAsNotForSaleRequest() // PostApiEventsByEventIdActionsMarkEverythingAsNotForSaleRequest | 
	idempotencyKey := "idempotencyKey_example" // string | Required on transactional routes. Reuse the same key for every retry of one logical operation; the API replays the stored outcome. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.EventsAPI.PostApiEventsByEventIdActionsMarkEverythingAsNotForSale(context.Background(), eventId).PostApiEventsByEventIdActionsMarkEverythingAsNotForSaleRequest(postApiEventsByEventIdActionsMarkEverythingAsNotForSaleRequest).IdempotencyKey(idempotencyKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EventsAPI.PostApiEventsByEventIdActionsMarkEverythingAsNotForSale``: %v\n", err)
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

Other parameters are passed through a pointer to a apiPostApiEventsByEventIdActionsMarkEverythingAsNotForSaleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **postApiEventsByEventIdActionsMarkEverythingAsNotForSaleRequest** | [**PostApiEventsByEventIdActionsMarkEverythingAsNotForSaleRequest**](PostApiEventsByEventIdActionsMarkEverythingAsNotForSaleRequest.md) |  | 
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


## PostApiEventsByEventIdActionsMoveToNewChartCopy

> PostApiEventsByEventIdActionsMoveToNewChartCopy(ctx, eventId).PostApiChartsByChartIdVersionsByVersionActionsCopyRequest(postApiChartsByChartIdVersionsByVersionActionsCopyRequest).IdempotencyKey(idempotencyKey).Execute()

Move the event to a new copy of its chart



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
	postApiChartsByChartIdVersionsByVersionActionsCopyRequest := *openapiclient.NewPostApiChartsByChartIdVersionsByVersionActionsCopyRequest() // PostApiChartsByChartIdVersionsByVersionActionsCopyRequest | 
	idempotencyKey := "idempotencyKey_example" // string | Required on transactional routes. Reuse the same key for every retry of one logical operation; the API replays the stored outcome. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.EventsAPI.PostApiEventsByEventIdActionsMoveToNewChartCopy(context.Background(), eventId).PostApiChartsByChartIdVersionsByVersionActionsCopyRequest(postApiChartsByChartIdVersionsByVersionActionsCopyRequest).IdempotencyKey(idempotencyKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EventsAPI.PostApiEventsByEventIdActionsMoveToNewChartCopy``: %v\n", err)
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

Other parameters are passed through a pointer to a apiPostApiEventsByEventIdActionsMoveToNewChartCopyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **postApiChartsByChartIdVersionsByVersionActionsCopyRequest** | [**PostApiChartsByChartIdVersionsByVersionActionsCopyRequest**](PostApiChartsByChartIdVersionsByVersionActionsCopyRequest.md) |  | 
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


## PostApiEventsByEventIdDuplicate

> PostApiEventsByEventIdDuplicate(ctx, eventId).PostApiEventsByEventIdDuplicateRequest(postApiEventsByEventIdDuplicateRequest).IdempotencyKey(idempotencyKey).Execute()

Duplicate event



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
	postApiEventsByEventIdDuplicateRequest := *openapiclient.NewPostApiEventsByEventIdDuplicateRequest() // PostApiEventsByEventIdDuplicateRequest | 
	idempotencyKey := "idempotencyKey_example" // string | Required on transactional routes. Reuse the same key for every retry of one logical operation; the API replays the stored outcome. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.EventsAPI.PostApiEventsByEventIdDuplicate(context.Background(), eventId).PostApiEventsByEventIdDuplicateRequest(postApiEventsByEventIdDuplicateRequest).IdempotencyKey(idempotencyKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EventsAPI.PostApiEventsByEventIdDuplicate``: %v\n", err)
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

Other parameters are passed through a pointer to a apiPostApiEventsByEventIdDuplicateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **postApiEventsByEventIdDuplicateRequest** | [**PostApiEventsByEventIdDuplicateRequest**](PostApiEventsByEventIdDuplicateRequest.md) |  | 
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


## PostApiEventsByEventIdInventoryChannels

> PostApiEventsByEventIdInventoryChannels(ctx, eventId).PostApiEventsByEventIdInventoryChannelsRequest(postApiEventsByEventIdInventoryChannelsRequest).IdempotencyKey(idempotencyKey).Execute()

Create inventory channel



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
	postApiEventsByEventIdInventoryChannelsRequest := *openapiclient.NewPostApiEventsByEventIdInventoryChannelsRequest("Name_example") // PostApiEventsByEventIdInventoryChannelsRequest | 
	idempotencyKey := "idempotencyKey_example" // string | Required on transactional routes. Reuse the same key for every retry of one logical operation; the API replays the stored outcome. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.EventsAPI.PostApiEventsByEventIdInventoryChannels(context.Background(), eventId).PostApiEventsByEventIdInventoryChannelsRequest(postApiEventsByEventIdInventoryChannelsRequest).IdempotencyKey(idempotencyKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EventsAPI.PostApiEventsByEventIdInventoryChannels``: %v\n", err)
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

Other parameters are passed through a pointer to a apiPostApiEventsByEventIdInventoryChannelsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **postApiEventsByEventIdInventoryChannelsRequest** | [**PostApiEventsByEventIdInventoryChannelsRequest**](PostApiEventsByEventIdInventoryChannelsRequest.md) |  | 
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


## PostApiEventsByEventIdInventoryChannelsAssign

> PostApiEventsByEventIdInventoryChannelsAssign(ctx, eventId).PostApiEventsByEventIdInventoryChannelsAssignRequest(postApiEventsByEventIdInventoryChannelsAssignRequest).IdempotencyKey(idempotencyKey).Execute()

Assign seats/GA to inventory channel



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
	postApiEventsByEventIdInventoryChannelsAssignRequest := *openapiclient.NewPostApiEventsByEventIdInventoryChannelsAssignRequest("ChannelKeyId_example") // PostApiEventsByEventIdInventoryChannelsAssignRequest | 
	idempotencyKey := "idempotencyKey_example" // string | Required on transactional routes. Reuse the same key for every retry of one logical operation; the API replays the stored outcome. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.EventsAPI.PostApiEventsByEventIdInventoryChannelsAssign(context.Background(), eventId).PostApiEventsByEventIdInventoryChannelsAssignRequest(postApiEventsByEventIdInventoryChannelsAssignRequest).IdempotencyKey(idempotencyKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EventsAPI.PostApiEventsByEventIdInventoryChannelsAssign``: %v\n", err)
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

Other parameters are passed through a pointer to a apiPostApiEventsByEventIdInventoryChannelsAssignRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **postApiEventsByEventIdInventoryChannelsAssignRequest** | [**PostApiEventsByEventIdInventoryChannelsAssignRequest**](PostApiEventsByEventIdInventoryChannelsAssignRequest.md) |  | 
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


## PostApiEventsByEventIdInventoryForSaleApply

> PostApiEventsByEventIdInventoryForSaleApply(ctx, eventId).PostApiEventsByEventIdInventoryForSaleApplyRequest(postApiEventsByEventIdInventoryForSaleApplyRequest).IdempotencyKey(idempotencyKey).Execute()

Apply for-sale / not-for-sale to a selection (incremental)



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
	postApiEventsByEventIdInventoryForSaleApplyRequest := *openapiclient.NewPostApiEventsByEventIdInventoryForSaleApplyRequest() // PostApiEventsByEventIdInventoryForSaleApplyRequest | 
	idempotencyKey := "idempotencyKey_example" // string | Required on transactional routes. Reuse the same key for every retry of one logical operation; the API replays the stored outcome. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.EventsAPI.PostApiEventsByEventIdInventoryForSaleApply(context.Background(), eventId).PostApiEventsByEventIdInventoryForSaleApplyRequest(postApiEventsByEventIdInventoryForSaleApplyRequest).IdempotencyKey(idempotencyKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EventsAPI.PostApiEventsByEventIdInventoryForSaleApply``: %v\n", err)
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

Other parameters are passed through a pointer to a apiPostApiEventsByEventIdInventoryForSaleApplyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **postApiEventsByEventIdInventoryForSaleApplyRequest** | [**PostApiEventsByEventIdInventoryForSaleApplyRequest**](PostApiEventsByEventIdInventoryForSaleApplyRequest.md) |  | 
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


## PostApiEventsByEventIdMapDraftActionsDiscard

> PostApiEventsByEventIdMapDraftActionsDiscard(ctx, eventId).IdempotencyKey(idempotencyKey).Execute()

Discard event map draft



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
	r, err := apiClient.EventsAPI.PostApiEventsByEventIdMapDraftActionsDiscard(context.Background(), eventId).IdempotencyKey(idempotencyKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EventsAPI.PostApiEventsByEventIdMapDraftActionsDiscard``: %v\n", err)
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

Other parameters are passed through a pointer to a apiPostApiEventsByEventIdMapDraftActionsDiscardRequest struct via the builder pattern


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


## PostApiEventsByEventIdMapPublish

> PostApiEventsByEventIdMapPublish(ctx, eventId).IdempotencyKey(idempotencyKey).Execute()

Publish draft map to live seats



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
	r, err := apiClient.EventsAPI.PostApiEventsByEventIdMapPublish(context.Background(), eventId).IdempotencyKey(idempotencyKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EventsAPI.PostApiEventsByEventIdMapPublish``: %v\n", err)
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

Other parameters are passed through a pointer to a apiPostApiEventsByEventIdMapPublishRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **idempotencyKey** | **string** | Required on transactional routes. Reuse the same key for every retry of one logical operation; the API replays the stored outcome. | 

### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, multipart/form-data, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostApiEventsByEventIdMapVersionsByVersionActionsRestoreToDraft

> PostApiEventsByEventIdMapVersionsByVersionActionsRestoreToDraft(ctx, eventId, version).IdempotencyKey(idempotencyKey).Execute()

Restore an event map version into the draft



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
	version := float32(8.14) // float32 | Version number (1-based).
	idempotencyKey := "idempotencyKey_example" // string | Required on transactional routes. Reuse the same key for every retry of one logical operation; the API replays the stored outcome. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.EventsAPI.PostApiEventsByEventIdMapVersionsByVersionActionsRestoreToDraft(context.Background(), eventId, version).IdempotencyKey(idempotencyKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EventsAPI.PostApiEventsByEventIdMapVersionsByVersionActionsRestoreToDraft``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**eventId** | **string** |  | 
**version** | **float32** | Version number (1-based). | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostApiEventsByEventIdMapVersionsByVersionActionsRestoreToDraftRequest struct via the builder pattern


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


## PostApiEventsByEventIdObjectsActionsRetrieve

> PostApiEventsByEventIdObjectsActionsRetrieve(ctx, eventId).PostApiEventsByEventIdObjectsActionsRetrieveRequest(postApiEventsByEventIdObjectsActionsRetrieveRequest).IdempotencyKey(idempotencyKey).Execute()

Retrieve several objects at once (seats.io retrieveObjectInfos)



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
	postApiEventsByEventIdObjectsActionsRetrieveRequest := *openapiclient.NewPostApiEventsByEventIdObjectsActionsRetrieveRequest([]string{"Objects_example"}) // PostApiEventsByEventIdObjectsActionsRetrieveRequest | 
	idempotencyKey := "idempotencyKey_example" // string | Required on transactional routes. Reuse the same key for every retry of one logical operation; the API replays the stored outcome. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.EventsAPI.PostApiEventsByEventIdObjectsActionsRetrieve(context.Background(), eventId).PostApiEventsByEventIdObjectsActionsRetrieveRequest(postApiEventsByEventIdObjectsActionsRetrieveRequest).IdempotencyKey(idempotencyKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EventsAPI.PostApiEventsByEventIdObjectsActionsRetrieve``: %v\n", err)
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

Other parameters are passed through a pointer to a apiPostApiEventsByEventIdObjectsActionsRetrieveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **postApiEventsByEventIdObjectsActionsRetrieveRequest** | [**PostApiEventsByEventIdObjectsActionsRetrieveRequest**](PostApiEventsByEventIdObjectsActionsRetrieveRequest.md) |  | 
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


## PostApiEventsByEventIdSalesCountersActionsRebuild

> PostApiEventsByEventIdSalesCountersActionsRebuild(ctx, eventId).IdempotencyKey(idempotencyKey).Execute()

Rebuild the sold counters of an event from its orders



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
	idempotencyKey := "idempotencyKey_example" // string | Required on transactional routes. Reuse the same key for every retry of one logical operation; the API replays the stored outcome. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.EventsAPI.PostApiEventsByEventIdSalesCountersActionsRebuild(context.Background(), eventId).IdempotencyKey(idempotencyKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EventsAPI.PostApiEventsByEventIdSalesCountersActionsRebuild``: %v\n", err)
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

Other parameters are passed through a pointer to a apiPostApiEventsByEventIdSalesCountersActionsRebuildRequest struct via the builder pattern


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


## PutApiEventsByEventIdInventoryChannels

> PutApiEventsByEventIdInventoryChannels(ctx, eventId).PutApiEventsByEventIdInventoryChannelsRequest(putApiEventsByEventIdInventoryChannelsRequest).IdempotencyKey(idempotencyKey).Execute()

Replace all inventory channel assignments atomically



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
	putApiEventsByEventIdInventoryChannelsRequest := *openapiclient.NewPutApiEventsByEventIdInventoryChannelsRequest([]openapiclient.PutApiEventsByEventIdInventoryChannelsRequestAssignmentsInner{*openapiclient.NewPutApiEventsByEventIdInventoryChannelsRequestAssignmentsInner("ChannelKeyId_example")}) // PutApiEventsByEventIdInventoryChannelsRequest | 
	idempotencyKey := "idempotencyKey_example" // string | Required on transactional routes. Reuse the same key for every retry of one logical operation; the API replays the stored outcome. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.EventsAPI.PutApiEventsByEventIdInventoryChannels(context.Background(), eventId).PutApiEventsByEventIdInventoryChannelsRequest(putApiEventsByEventIdInventoryChannelsRequest).IdempotencyKey(idempotencyKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EventsAPI.PutApiEventsByEventIdInventoryChannels``: %v\n", err)
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

Other parameters are passed through a pointer to a apiPutApiEventsByEventIdInventoryChannelsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **putApiEventsByEventIdInventoryChannelsRequest** | [**PutApiEventsByEventIdInventoryChannelsRequest**](PutApiEventsByEventIdInventoryChannelsRequest.md) |  | 
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


## PutApiEventsByEventIdInventoryChannelsDefinitions

> PutApiEventsByEventIdInventoryChannelsDefinitions(ctx, eventId).PutApiEventsByEventIdInventoryChannelsDefinitionsRequest(putApiEventsByEventIdInventoryChannelsDefinitionsRequest).IdempotencyKey(idempotencyKey).Execute()

Replace all inventory channels (definitions and assignments) atomically



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
	putApiEventsByEventIdInventoryChannelsDefinitionsRequest := *openapiclient.NewPutApiEventsByEventIdInventoryChannelsDefinitionsRequest([]openapiclient.PutApiEventsByEventIdInventoryChannelsDefinitionsRequestChannelsInner{*openapiclient.NewPutApiEventsByEventIdInventoryChannelsDefinitionsRequestChannelsInner("Name_example")}) // PutApiEventsByEventIdInventoryChannelsDefinitionsRequest | 
	idempotencyKey := "idempotencyKey_example" // string | Required on transactional routes. Reuse the same key for every retry of one logical operation; the API replays the stored outcome. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.EventsAPI.PutApiEventsByEventIdInventoryChannelsDefinitions(context.Background(), eventId).PutApiEventsByEventIdInventoryChannelsDefinitionsRequest(putApiEventsByEventIdInventoryChannelsDefinitionsRequest).IdempotencyKey(idempotencyKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EventsAPI.PutApiEventsByEventIdInventoryChannelsDefinitions``: %v\n", err)
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

Other parameters are passed through a pointer to a apiPutApiEventsByEventIdInventoryChannelsDefinitionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **putApiEventsByEventIdInventoryChannelsDefinitionsRequest** | [**PutApiEventsByEventIdInventoryChannelsDefinitionsRequest**](PutApiEventsByEventIdInventoryChannelsDefinitionsRequest.md) |  | 
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


## PutApiEventsByEventIdInventoryForSale

> PutApiEventsByEventIdInventoryForSale(ctx, eventId).PutApiEventsByEventIdInventoryForSaleRequest(putApiEventsByEventIdInventoryForSaleRequest).IdempotencyKey(idempotencyKey).Execute()

Replace the whole for-sale configuration atomically



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
	putApiEventsByEventIdInventoryForSaleRequest := *openapiclient.NewPutApiEventsByEventIdInventoryForSaleRequest(false) // PutApiEventsByEventIdInventoryForSaleRequest | 
	idempotencyKey := "idempotencyKey_example" // string | Required on transactional routes. Reuse the same key for every retry of one logical operation; the API replays the stored outcome. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.EventsAPI.PutApiEventsByEventIdInventoryForSale(context.Background(), eventId).PutApiEventsByEventIdInventoryForSaleRequest(putApiEventsByEventIdInventoryForSaleRequest).IdempotencyKey(idempotencyKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EventsAPI.PutApiEventsByEventIdInventoryForSale``: %v\n", err)
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

Other parameters are passed through a pointer to a apiPutApiEventsByEventIdInventoryForSaleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **putApiEventsByEventIdInventoryForSaleRequest** | [**PutApiEventsByEventIdInventoryForSaleRequest**](PutApiEventsByEventIdInventoryForSaleRequest.md) |  | 
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


## PutApiEventsByEventIdMap

> PutApiEventsByEventIdMap(ctx, eventId).PutApiChartsByChartIdMapDraftRequest(putApiChartsByChartIdMapDraftRequest).IdempotencyKey(idempotencyKey).Execute()

Save map draft (legacy path)



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
	putApiChartsByChartIdMapDraftRequest := *openapiclient.NewPutApiChartsByChartIdMapDraftRequest([]openapiclient.PostApiChartsByChartIdActionsValidateRequestSeatsInner{*openapiclient.NewPostApiChartsByChartIdActionsValidateRequestSeatsInner("Label_example", float32(123), float32(123))}) // PutApiChartsByChartIdMapDraftRequest | 
	idempotencyKey := "idempotencyKey_example" // string | Required on transactional routes. Reuse the same key for every retry of one logical operation; the API replays the stored outcome. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.EventsAPI.PutApiEventsByEventIdMap(context.Background(), eventId).PutApiChartsByChartIdMapDraftRequest(putApiChartsByChartIdMapDraftRequest).IdempotencyKey(idempotencyKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EventsAPI.PutApiEventsByEventIdMap``: %v\n", err)
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

Other parameters are passed through a pointer to a apiPutApiEventsByEventIdMapRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **putApiChartsByChartIdMapDraftRequest** | [**PutApiChartsByChartIdMapDraftRequest**](PutApiChartsByChartIdMapDraftRequest.md) |  | 
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


## PutApiEventsByEventIdMapDraft

> PutApiEventsByEventIdMapDraft(ctx, eventId).PutApiChartsByChartIdMapDraftRequest(putApiChartsByChartIdMapDraftRequest).IdempotencyKey(idempotencyKey).Execute()

Save map draft (does not go live)



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
	putApiChartsByChartIdMapDraftRequest := *openapiclient.NewPutApiChartsByChartIdMapDraftRequest([]openapiclient.PostApiChartsByChartIdActionsValidateRequestSeatsInner{*openapiclient.NewPostApiChartsByChartIdActionsValidateRequestSeatsInner("Label_example", float32(123), float32(123))}) // PutApiChartsByChartIdMapDraftRequest | 
	idempotencyKey := "idempotencyKey_example" // string | Required on transactional routes. Reuse the same key for every retry of one logical operation; the API replays the stored outcome. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.EventsAPI.PutApiEventsByEventIdMapDraft(context.Background(), eventId).PutApiChartsByChartIdMapDraftRequest(putApiChartsByChartIdMapDraftRequest).IdempotencyKey(idempotencyKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EventsAPI.PutApiEventsByEventIdMapDraft``: %v\n", err)
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

Other parameters are passed through a pointer to a apiPutApiEventsByEventIdMapDraftRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **putApiChartsByChartIdMapDraftRequest** | [**PutApiChartsByChartIdMapDraftRequest**](PutApiChartsByChartIdMapDraftRequest.md) |  | 
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


## PutApiEventsByEventIdPricing

> PutApiEventsByEventIdPricing(ctx, eventId).PutApiEventsByEventIdPricingRequest(putApiEventsByEventIdPricingRequest).IdempotencyKey(idempotencyKey).Execute()

Push the synced price table of an event (synced catalog mode)



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
	putApiEventsByEventIdPricingRequest := *openapiclient.NewPutApiEventsByEventIdPricingRequest("Currency_example", []openapiclient.PutApiEventsByEventIdPricingRequestCategoriesInner{*openapiclient.NewPutApiEventsByEventIdPricingRequestCategoriesInner("Category_example", []openapiclient.PutApiEventsByEventIdPricingRequestCategoriesInnerTicketTypesInner{*openapiclient.NewPutApiEventsByEventIdPricingRequestCategoriesInnerTicketTypesInner("Id_example", "Name_example", int32(123))})}) // PutApiEventsByEventIdPricingRequest | 
	idempotencyKey := "idempotencyKey_example" // string | Required on transactional routes. Reuse the same key for every retry of one logical operation; the API replays the stored outcome. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.EventsAPI.PutApiEventsByEventIdPricing(context.Background(), eventId).PutApiEventsByEventIdPricingRequest(putApiEventsByEventIdPricingRequest).IdempotencyKey(idempotencyKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EventsAPI.PutApiEventsByEventIdPricing``: %v\n", err)
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

Other parameters are passed through a pointer to a apiPutApiEventsByEventIdPricingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **putApiEventsByEventIdPricingRequest** | [**PutApiEventsByEventIdPricingRequest**](PutApiEventsByEventIdPricingRequest.md) |  | 
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


## PutApiEventsByEventIdSalesRules

> PutApiEventsByEventIdSalesRules(ctx, eventId).PutApiEventsByEventIdSalesRulesRequest(putApiEventsByEventIdSalesRulesRequest).IdempotencyKey(idempotencyKey).Execute()

Replace the sales rules of an event (quotas, lots, buyer limit, pricing source)



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
	putApiEventsByEventIdSalesRulesRequest := *openapiclient.NewPutApiEventsByEventIdSalesRulesRequest() // PutApiEventsByEventIdSalesRulesRequest | 
	idempotencyKey := "idempotencyKey_example" // string | Required on transactional routes. Reuse the same key for every retry of one logical operation; the API replays the stored outcome. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.EventsAPI.PutApiEventsByEventIdSalesRules(context.Background(), eventId).PutApiEventsByEventIdSalesRulesRequest(putApiEventsByEventIdSalesRulesRequest).IdempotencyKey(idempotencyKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EventsAPI.PutApiEventsByEventIdSalesRules``: %v\n", err)
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

Other parameters are passed through a pointer to a apiPutApiEventsByEventIdSalesRulesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **putApiEventsByEventIdSalesRulesRequest** | [**PutApiEventsByEventIdSalesRulesRequest**](PutApiEventsByEventIdSalesRulesRequest.md) |  | 
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

