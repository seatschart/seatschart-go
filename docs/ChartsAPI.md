# \ChartsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteApiChartsByChartId**](ChartsAPI.md#DeleteApiChartsByChartId) | **Delete** /api/charts/{chartId} | Delete chart permanently
[**DeleteApiChartsByChartIdPreviewLink**](ChartsAPI.md#DeleteApiChartsByChartIdPreviewLink) | **Delete** /api/charts/{chartId}/preview-link | Revoke every public preview link of a chart
[**GetApiCharts**](ChartsAPI.md#GetApiCharts) | **Get** /api/charts | List charts
[**GetApiChartsByChartId**](ChartsAPI.md#GetApiChartsByChartId) | **Get** /api/charts/{chartId} | Retrieve one chart
[**GetApiChartsByChartIdCategories**](ChartsAPI.md#GetApiChartsByChartIdCategories) | **Get** /api/charts/{chartId}/categories | List the categories of a chart
[**GetApiChartsByChartIdMap**](ChartsAPI.md#GetApiChartsByChartIdMap) | **Get** /api/charts/{chartId}/map | Read one version of a chart map
[**GetApiChartsByChartIdMapDraft**](ChartsAPI.md#GetApiChartsByChartIdMapDraft) | **Get** /api/charts/{chartId}/map/draft | Load chart map draft for designer
[**GetApiChartsByChartIdReportsByCategoryCsv**](ChartsAPI.md#GetApiChartsByChartIdReportsByCategoryCsv) | **Get** /api/charts/{chartId}/reports/by-category.csv | Chart report by category (CSV)
[**GetApiChartsByChartIdReportsByCategoryJson**](ChartsAPI.md#GetApiChartsByChartIdReportsByCategoryJson) | **Get** /api/charts/{chartId}/reports/by-category.json | Chart report by category
[**GetApiChartsByChartIdReportsByCategoryKeyCsv**](ChartsAPI.md#GetApiChartsByChartIdReportsByCategoryKeyCsv) | **Get** /api/charts/{chartId}/reports/by-category-key.csv | Chart report by category key (CSV)
[**GetApiChartsByChartIdReportsByCategoryKeyJson**](ChartsAPI.md#GetApiChartsByChartIdReportsByCategoryKeyJson) | **Get** /api/charts/{chartId}/reports/by-category-key.json | Chart report by category key
[**GetApiChartsByChartIdReportsByCategoryLabelCsv**](ChartsAPI.md#GetApiChartsByChartIdReportsByCategoryLabelCsv) | **Get** /api/charts/{chartId}/reports/by-category-label.csv | Chart report by category label (CSV)
[**GetApiChartsByChartIdReportsByCategoryLabelJson**](ChartsAPI.md#GetApiChartsByChartIdReportsByCategoryLabelJson) | **Get** /api/charts/{chartId}/reports/by-category-label.json | Chart report by category label
[**GetApiChartsByChartIdReportsByLabelCsv**](ChartsAPI.md#GetApiChartsByChartIdReportsByLabelCsv) | **Get** /api/charts/{chartId}/reports/by-label.csv | Chart report by label (CSV)
[**GetApiChartsByChartIdReportsByLabelJson**](ChartsAPI.md#GetApiChartsByChartIdReportsByLabelJson) | **Get** /api/charts/{chartId}/reports/by-label.json | Chart report by label
[**GetApiChartsByChartIdReportsByObjectTypeCsv**](ChartsAPI.md#GetApiChartsByChartIdReportsByObjectTypeCsv) | **Get** /api/charts/{chartId}/reports/by-object-type.csv | Chart report by object type (CSV)
[**GetApiChartsByChartIdReportsByObjectTypeJson**](ChartsAPI.md#GetApiChartsByChartIdReportsByObjectTypeJson) | **Get** /api/charts/{chartId}/reports/by-object-type.json | Chart report by object type
[**GetApiChartsByChartIdReportsBySectionCsv**](ChartsAPI.md#GetApiChartsByChartIdReportsBySectionCsv) | **Get** /api/charts/{chartId}/reports/by-section.csv | Chart report by section (CSV)
[**GetApiChartsByChartIdReportsBySectionJson**](ChartsAPI.md#GetApiChartsByChartIdReportsBySectionJson) | **Get** /api/charts/{chartId}/reports/by-section.json | Chart report by section
[**GetApiChartsByChartIdReportsByZoneCsv**](ChartsAPI.md#GetApiChartsByChartIdReportsByZoneCsv) | **Get** /api/charts/{chartId}/reports/by-zone.csv | Chart report by zone (CSV)
[**GetApiChartsByChartIdReportsByZoneJson**](ChartsAPI.md#GetApiChartsByChartIdReportsByZoneJson) | **Get** /api/charts/{chartId}/reports/by-zone.json | Chart report by zone
[**GetApiChartsByChartIdReportsSummaryByCategoryCsv**](ChartsAPI.md#GetApiChartsByChartIdReportsSummaryByCategoryCsv) | **Get** /api/charts/{chartId}/reports/summary/by-category.csv | Chart summary by category (CSV)
[**GetApiChartsByChartIdReportsSummaryByCategoryJson**](ChartsAPI.md#GetApiChartsByChartIdReportsSummaryByCategoryJson) | **Get** /api/charts/{chartId}/reports/summary/by-category.json | Chart summary by category
[**GetApiChartsByChartIdReportsSummaryByCategoryKeyCsv**](ChartsAPI.md#GetApiChartsByChartIdReportsSummaryByCategoryKeyCsv) | **Get** /api/charts/{chartId}/reports/summary/by-category-key.csv | Chart summary by category key (CSV)
[**GetApiChartsByChartIdReportsSummaryByCategoryKeyJson**](ChartsAPI.md#GetApiChartsByChartIdReportsSummaryByCategoryKeyJson) | **Get** /api/charts/{chartId}/reports/summary/by-category-key.json | Chart summary by category key
[**GetApiChartsByChartIdReportsSummaryByCategoryLabelCsv**](ChartsAPI.md#GetApiChartsByChartIdReportsSummaryByCategoryLabelCsv) | **Get** /api/charts/{chartId}/reports/summary/by-category-label.csv | Chart summary by category label (CSV)
[**GetApiChartsByChartIdReportsSummaryByCategoryLabelJson**](ChartsAPI.md#GetApiChartsByChartIdReportsSummaryByCategoryLabelJson) | **Get** /api/charts/{chartId}/reports/summary/by-category-label.json | Chart summary by category label
[**GetApiChartsByChartIdReportsSummaryByObjectTypeCsv**](ChartsAPI.md#GetApiChartsByChartIdReportsSummaryByObjectTypeCsv) | **Get** /api/charts/{chartId}/reports/summary/by-object-type.csv | Chart summary by object type (CSV)
[**GetApiChartsByChartIdReportsSummaryByObjectTypeJson**](ChartsAPI.md#GetApiChartsByChartIdReportsSummaryByObjectTypeJson) | **Get** /api/charts/{chartId}/reports/summary/by-object-type.json | Chart summary by object type
[**GetApiChartsByChartIdReportsSummaryBySectionCsv**](ChartsAPI.md#GetApiChartsByChartIdReportsSummaryBySectionCsv) | **Get** /api/charts/{chartId}/reports/summary/by-section.csv | Chart summary by section (CSV)
[**GetApiChartsByChartIdReportsSummaryBySectionJson**](ChartsAPI.md#GetApiChartsByChartIdReportsSummaryBySectionJson) | **Get** /api/charts/{chartId}/reports/summary/by-section.json | Chart summary by section
[**GetApiChartsByChartIdReportsSummaryByZoneCsv**](ChartsAPI.md#GetApiChartsByChartIdReportsSummaryByZoneCsv) | **Get** /api/charts/{chartId}/reports/summary/by-zone.csv | Chart summary by zone (CSV)
[**GetApiChartsByChartIdReportsSummaryByZoneJson**](ChartsAPI.md#GetApiChartsByChartIdReportsSummaryByZoneJson) | **Get** /api/charts/{chartId}/reports/summary/by-zone.json | Chart summary by zone
[**GetApiChartsByChartIdThumbnail**](ChartsAPI.md#GetApiChartsByChartIdThumbnail) | **Get** /api/charts/{chartId}/thumbnail | Chart thumbnail
[**GetApiChartsByChartIdVersions**](ChartsAPI.md#GetApiChartsByChartIdVersions) | **Get** /api/charts/{chartId}/versions | List published versions
[**GetApiChartsByChartIdZones**](ChartsAPI.md#GetApiChartsByChartIdZones) | **Get** /api/charts/{chartId}/zones | List the zones of a chart
[**GetApiChartsTags**](ChartsAPI.md#GetApiChartsTags) | **Get** /api/charts/tags | List chart tags
[**PatchApiChartsByChartId**](ChartsAPI.md#PatchApiChartsByChartId) | **Patch** /api/charts/{chartId} | Update chart name or key
[**PostApiCharts**](ChartsAPI.md#PostApiCharts) | **Post** /api/charts | Create chart
[**PostApiChartsByChartIdActionsArchive**](ChartsAPI.md#PostApiChartsByChartIdActionsArchive) | **Post** /api/charts/{chartId}/actions/archive | Archive chart
[**PostApiChartsByChartIdActionsCopyToWorkspace**](ChartsAPI.md#PostApiChartsByChartIdActionsCopyToWorkspace) | **Post** /api/charts/{chartId}/actions/copy-to-workspace | Copy a chart to another workspace
[**PostApiChartsByChartIdActionsUnarchive**](ChartsAPI.md#PostApiChartsByChartIdActionsUnarchive) | **Post** /api/charts/{chartId}/actions/unarchive | Unarchive chart
[**PostApiChartsByChartIdActionsValidate**](ChartsAPI.md#PostApiChartsByChartIdActionsValidate) | **Post** /api/charts/{chartId}/actions/validate | Validate a chart map (dry run)
[**PostApiChartsByChartIdDuplicate**](ChartsAPI.md#PostApiChartsByChartIdDuplicate) | **Post** /api/charts/{chartId}/duplicate | Duplicate chart
[**PostApiChartsByChartIdMapDraftActionsDiscard**](ChartsAPI.md#PostApiChartsByChartIdMapDraftActionsDiscard) | **Post** /api/charts/{chartId}/map/draft/actions/discard | Discard chart draft
[**PostApiChartsByChartIdMapPublish**](ChartsAPI.md#PostApiChartsByChartIdMapPublish) | **Post** /api/charts/{chartId}/map/publish | Publish chart map template
[**PostApiChartsByChartIdPreviewLink**](ChartsAPI.md#PostApiChartsByChartIdPreviewLink) | **Post** /api/charts/{chartId}/preview-link | Create a public preview link
[**PostApiChartsByChartIdVersionsByVersionActionsRestoreToDraft**](ChartsAPI.md#PostApiChartsByChartIdVersionsByVersionActionsRestoreToDraft) | **Post** /api/charts/{chartId}/versions/{version}/actions/restore-to-draft | Restore a version into the draft
[**PutApiChartsByChartIdMapDraft**](ChartsAPI.md#PutApiChartsByChartIdMapDraft) | **Put** /api/charts/{chartId}/map/draft | Save chart map draft
[**PutApiChartsByChartIdTags**](ChartsAPI.md#PutApiChartsByChartIdTags) | **Put** /api/charts/{chartId}/tags | Replace chart tags
[**PutApiChartsByChartIdVersionsByVersionThumbnail**](ChartsAPI.md#PutApiChartsByChartIdVersionsByVersionThumbnail) | **Put** /api/charts/{chartId}/versions/{version}/thumbnail | Attach a rendered thumbnail to a version



## DeleteApiChartsByChartId

> DeleteApiChartsByChartId(ctx, chartId).IdempotencyKey(idempotencyKey).Execute()

Delete chart permanently



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
	r, err := apiClient.ChartsAPI.DeleteApiChartsByChartId(context.Background(), chartId).IdempotencyKey(idempotencyKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChartsAPI.DeleteApiChartsByChartId``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDeleteApiChartsByChartIdRequest struct via the builder pattern


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


## DeleteApiChartsByChartIdPreviewLink

> DeleteApiChartsByChartIdPreviewLink(ctx, chartId).IdempotencyKey(idempotencyKey).Execute()

Revoke every public preview link of a chart



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
	r, err := apiClient.ChartsAPI.DeleteApiChartsByChartIdPreviewLink(context.Background(), chartId).IdempotencyKey(idempotencyKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChartsAPI.DeleteApiChartsByChartIdPreviewLink``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDeleteApiChartsByChartIdPreviewLinkRequest struct via the builder pattern


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


## GetApiCharts

> GetApiCharts(ctx).Limit(limit).Cursor(cursor).Q(q).Archived(archived).Tag(tag).Sort(sort).Expand(expand).Execute()

List charts



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
	q := "q_example" // string | Case-insensitive substring match on the chart name, or the exact chart key. (optional)
	archived := "archived_example" // string | `false` (default) hides archived charts, `true` returns only archived charts, `all` returns both. (optional)
	tag := "tag_example" // string | Comma-separated tags (case-insensitive, at most 5); a chart must carry all of them. (optional)
	sort := "sort_example" // string | `name` (default): name A→Z. `recent`: last draft save / publish first, then creation time. (optional)
	expand := "expand_example" // string | Comma-separated. `events` adds `events[]` (`{ id, key, name, date, isInThePast }`, soonest first, at most 100 per chart; `eventCount` has the total). `validation` is accepted for seats.io compatibility (`validation` is always part of paginated items). (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ChartsAPI.GetApiCharts(context.Background()).Limit(limit).Cursor(cursor).Q(q).Archived(archived).Tag(tag).Sort(sort).Expand(expand).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChartsAPI.GetApiCharts``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetApiChartsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **float32** | Page size (1..100, default 20). | 
 **cursor** | **string** | Opaque &#x60;nextCursor&#x60; value from the previous page. | 
 **q** | **string** | Case-insensitive substring match on the chart name, or the exact chart key. | 
 **archived** | **string** | &#x60;false&#x60; (default) hides archived charts, &#x60;true&#x60; returns only archived charts, &#x60;all&#x60; returns both. | 
 **tag** | **string** | Comma-separated tags (case-insensitive, at most 5); a chart must carry all of them. | 
 **sort** | **string** | &#x60;name&#x60; (default): name A→Z. &#x60;recent&#x60;: last draft save / publish first, then creation time. | 
 **expand** | **string** | Comma-separated. &#x60;events&#x60; adds &#x60;events[]&#x60; (&#x60;{ id, key, name, date, isInThePast }&#x60;, soonest first, at most 100 per chart; &#x60;eventCount&#x60; has the total). &#x60;validation&#x60; is accepted for seats.io compatibility (&#x60;validation&#x60; is always part of paginated items). | 

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


## GetApiChartsByChartId

> GetApiChartsByChartId(ctx, chartId).Expand(expand).Version(version).Execute()

Retrieve one chart



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
	chartId := "chartId_example" // string | Chart id (UUID) or chart key.
	expand := "expand_example" // string | `zones` adds `zones[]` (same items as GET /charts/{chartId}/zones) of the version selected by `version`. (optional)
	version := "version_example" // string | `published` or `draft` (the unpublished designer draft). (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ChartsAPI.GetApiChartsByChartId(context.Background(), chartId).Expand(expand).Version(version).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChartsAPI.GetApiChartsByChartId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**chartId** | **string** | Chart id (UUID) or chart key. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetApiChartsByChartIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **expand** | **string** | &#x60;zones&#x60; adds &#x60;zones[]&#x60; (same items as GET /charts/{chartId}/zones) of the version selected by &#x60;version&#x60;. | 
 **version** | **string** | &#x60;published&#x60; or &#x60;draft&#x60; (the unpublished designer draft). | 

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


## GetApiChartsByChartIdCategories

> GetApiChartsByChartIdCategories(ctx, chartId).Version(version).Execute()

List the categories of a chart



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
	version := "version_example" // string | `published` or `draft`. Omitted: the draft when there is one, else the published map. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ChartsAPI.GetApiChartsByChartIdCategories(context.Background(), chartId).Version(version).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChartsAPI.GetApiChartsByChartIdCategories``: %v\n", err)
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

Other parameters are passed through a pointer to a apiGetApiChartsByChartIdCategoriesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **version** | **string** | &#x60;published&#x60; or &#x60;draft&#x60;. Omitted: the draft when there is one, else the published map. | 

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


## GetApiChartsByChartIdMap

> GetApiChartsByChartIdMap(ctx, chartId).Version(version).Execute()

Read one version of a chart map



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
	chartId := "chartId_example" // string | Chart id (UUID) or chart key.
	version := "version_example" // string | `published` or `draft` (the unpublished designer draft). (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ChartsAPI.GetApiChartsByChartIdMap(context.Background(), chartId).Version(version).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChartsAPI.GetApiChartsByChartIdMap``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**chartId** | **string** | Chart id (UUID) or chart key. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetApiChartsByChartIdMapRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **version** | **string** | &#x60;published&#x60; or &#x60;draft&#x60; (the unpublished designer draft). | 

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


## GetApiChartsByChartIdMapDraft

> GetApiChartsByChartIdMapDraft(ctx, chartId).Execute()

Load chart map draft for designer

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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ChartsAPI.GetApiChartsByChartIdMapDraft(context.Background(), chartId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChartsAPI.GetApiChartsByChartIdMapDraft``: %v\n", err)
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

Other parameters are passed through a pointer to a apiGetApiChartsByChartIdMapDraftRequest struct via the builder pattern


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


## GetApiChartsByChartIdReportsByCategoryCsv

> GetApiChartsByChartIdReportsByCategoryCsv(ctx, chartId).Category(category).Version(version).Execute()

Chart report by category (CSV)



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
	category := "category_example" // string | Only this category bucket. (optional)
	version := "version_example" // string | `published` (default) or `draft` (the unpublished designer draft). (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ChartsAPI.GetApiChartsByChartIdReportsByCategoryCsv(context.Background(), chartId).Category(category).Version(version).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChartsAPI.GetApiChartsByChartIdReportsByCategoryCsv``: %v\n", err)
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

Other parameters are passed through a pointer to a apiGetApiChartsByChartIdReportsByCategoryCsvRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **category** | **string** | Only this category bucket. | 
 **version** | **string** | &#x60;published&#x60; (default) or &#x60;draft&#x60; (the unpublished designer draft). | 

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


## GetApiChartsByChartIdReportsByCategoryJson

> GetApiChartsByChartIdReportsByCategoryJson(ctx, chartId).Category(category).Version(version).Execute()

Chart report by category



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
	category := "category_example" // string | Only this category bucket. (optional)
	version := "version_example" // string | `published` (default) or `draft` (the unpublished designer draft). (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ChartsAPI.GetApiChartsByChartIdReportsByCategoryJson(context.Background(), chartId).Category(category).Version(version).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChartsAPI.GetApiChartsByChartIdReportsByCategoryJson``: %v\n", err)
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

Other parameters are passed through a pointer to a apiGetApiChartsByChartIdReportsByCategoryJsonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **category** | **string** | Only this category bucket. | 
 **version** | **string** | &#x60;published&#x60; (default) or &#x60;draft&#x60; (the unpublished designer draft). | 

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


## GetApiChartsByChartIdReportsByCategoryKeyCsv

> GetApiChartsByChartIdReportsByCategoryKeyCsv(ctx, chartId).CategoryKey(categoryKey).Version(version).Execute()

Chart report by category key (CSV)



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
	categoryKey := "categoryKey_example" // string | Only this categoryKey bucket. (optional)
	version := "version_example" // string | `published` (default) or `draft` (the unpublished designer draft). (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ChartsAPI.GetApiChartsByChartIdReportsByCategoryKeyCsv(context.Background(), chartId).CategoryKey(categoryKey).Version(version).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChartsAPI.GetApiChartsByChartIdReportsByCategoryKeyCsv``: %v\n", err)
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

Other parameters are passed through a pointer to a apiGetApiChartsByChartIdReportsByCategoryKeyCsvRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **categoryKey** | **string** | Only this categoryKey bucket. | 
 **version** | **string** | &#x60;published&#x60; (default) or &#x60;draft&#x60; (the unpublished designer draft). | 

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


## GetApiChartsByChartIdReportsByCategoryKeyJson

> GetApiChartsByChartIdReportsByCategoryKeyJson(ctx, chartId).CategoryKey(categoryKey).Version(version).Execute()

Chart report by category key



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
	categoryKey := "categoryKey_example" // string | Only this categoryKey bucket. (optional)
	version := "version_example" // string | `published` (default) or `draft` (the unpublished designer draft). (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ChartsAPI.GetApiChartsByChartIdReportsByCategoryKeyJson(context.Background(), chartId).CategoryKey(categoryKey).Version(version).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChartsAPI.GetApiChartsByChartIdReportsByCategoryKeyJson``: %v\n", err)
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

Other parameters are passed through a pointer to a apiGetApiChartsByChartIdReportsByCategoryKeyJsonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **categoryKey** | **string** | Only this categoryKey bucket. | 
 **version** | **string** | &#x60;published&#x60; (default) or &#x60;draft&#x60; (the unpublished designer draft). | 

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


## GetApiChartsByChartIdReportsByCategoryLabelCsv

> GetApiChartsByChartIdReportsByCategoryLabelCsv(ctx, chartId).CategoryLabel(categoryLabel).Version(version).Execute()

Chart report by category label (CSV)



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
	categoryLabel := "categoryLabel_example" // string | Only this categoryLabel bucket. (optional)
	version := "version_example" // string | `published` (default) or `draft` (the unpublished designer draft). (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ChartsAPI.GetApiChartsByChartIdReportsByCategoryLabelCsv(context.Background(), chartId).CategoryLabel(categoryLabel).Version(version).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChartsAPI.GetApiChartsByChartIdReportsByCategoryLabelCsv``: %v\n", err)
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

Other parameters are passed through a pointer to a apiGetApiChartsByChartIdReportsByCategoryLabelCsvRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **categoryLabel** | **string** | Only this categoryLabel bucket. | 
 **version** | **string** | &#x60;published&#x60; (default) or &#x60;draft&#x60; (the unpublished designer draft). | 

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


## GetApiChartsByChartIdReportsByCategoryLabelJson

> GetApiChartsByChartIdReportsByCategoryLabelJson(ctx, chartId).CategoryLabel(categoryLabel).Version(version).Execute()

Chart report by category label



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
	categoryLabel := "categoryLabel_example" // string | Only this categoryLabel bucket. (optional)
	version := "version_example" // string | `published` (default) or `draft` (the unpublished designer draft). (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ChartsAPI.GetApiChartsByChartIdReportsByCategoryLabelJson(context.Background(), chartId).CategoryLabel(categoryLabel).Version(version).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChartsAPI.GetApiChartsByChartIdReportsByCategoryLabelJson``: %v\n", err)
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

Other parameters are passed through a pointer to a apiGetApiChartsByChartIdReportsByCategoryLabelJsonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **categoryLabel** | **string** | Only this categoryLabel bucket. | 
 **version** | **string** | &#x60;published&#x60; (default) or &#x60;draft&#x60; (the unpublished designer draft). | 

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


## GetApiChartsByChartIdReportsByLabelCsv

> GetApiChartsByChartIdReportsByLabelCsv(ctx, chartId).Label(label).Version(version).Execute()

Chart report by label (CSV)



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
	label := "label_example" // string | Only this label bucket. (optional)
	version := "version_example" // string | `published` (default) or `draft` (the unpublished designer draft). (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ChartsAPI.GetApiChartsByChartIdReportsByLabelCsv(context.Background(), chartId).Label(label).Version(version).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChartsAPI.GetApiChartsByChartIdReportsByLabelCsv``: %v\n", err)
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

Other parameters are passed through a pointer to a apiGetApiChartsByChartIdReportsByLabelCsvRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **label** | **string** | Only this label bucket. | 
 **version** | **string** | &#x60;published&#x60; (default) or &#x60;draft&#x60; (the unpublished designer draft). | 

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


## GetApiChartsByChartIdReportsByLabelJson

> GetApiChartsByChartIdReportsByLabelJson(ctx, chartId).Label(label).Version(version).Execute()

Chart report by label



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
	label := "label_example" // string | Only this label bucket. (optional)
	version := "version_example" // string | `published` (default) or `draft` (the unpublished designer draft). (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ChartsAPI.GetApiChartsByChartIdReportsByLabelJson(context.Background(), chartId).Label(label).Version(version).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChartsAPI.GetApiChartsByChartIdReportsByLabelJson``: %v\n", err)
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

Other parameters are passed through a pointer to a apiGetApiChartsByChartIdReportsByLabelJsonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **label** | **string** | Only this label bucket. | 
 **version** | **string** | &#x60;published&#x60; (default) or &#x60;draft&#x60; (the unpublished designer draft). | 

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


## GetApiChartsByChartIdReportsByObjectTypeCsv

> GetApiChartsByChartIdReportsByObjectTypeCsv(ctx, chartId).ObjectType(objectType).Version(version).Execute()

Chart report by object type (CSV)



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
	objectType := "objectType_example" // string | Only this objectType bucket. (optional)
	version := "version_example" // string | `published` (default) or `draft` (the unpublished designer draft). (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ChartsAPI.GetApiChartsByChartIdReportsByObjectTypeCsv(context.Background(), chartId).ObjectType(objectType).Version(version).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChartsAPI.GetApiChartsByChartIdReportsByObjectTypeCsv``: %v\n", err)
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

Other parameters are passed through a pointer to a apiGetApiChartsByChartIdReportsByObjectTypeCsvRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **objectType** | **string** | Only this objectType bucket. | 
 **version** | **string** | &#x60;published&#x60; (default) or &#x60;draft&#x60; (the unpublished designer draft). | 

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


## GetApiChartsByChartIdReportsByObjectTypeJson

> GetApiChartsByChartIdReportsByObjectTypeJson(ctx, chartId).ObjectType(objectType).Version(version).Execute()

Chart report by object type



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
	objectType := "objectType_example" // string | Only this objectType bucket. (optional)
	version := "version_example" // string | `published` (default) or `draft` (the unpublished designer draft). (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ChartsAPI.GetApiChartsByChartIdReportsByObjectTypeJson(context.Background(), chartId).ObjectType(objectType).Version(version).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChartsAPI.GetApiChartsByChartIdReportsByObjectTypeJson``: %v\n", err)
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

Other parameters are passed through a pointer to a apiGetApiChartsByChartIdReportsByObjectTypeJsonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **objectType** | **string** | Only this objectType bucket. | 
 **version** | **string** | &#x60;published&#x60; (default) or &#x60;draft&#x60; (the unpublished designer draft). | 

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


## GetApiChartsByChartIdReportsBySectionCsv

> GetApiChartsByChartIdReportsBySectionCsv(ctx, chartId).Section(section).Version(version).Execute()

Chart report by section (CSV)



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
	section := "section_example" // string | Only this section bucket. (optional)
	version := "version_example" // string | `published` (default) or `draft` (the unpublished designer draft). (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ChartsAPI.GetApiChartsByChartIdReportsBySectionCsv(context.Background(), chartId).Section(section).Version(version).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChartsAPI.GetApiChartsByChartIdReportsBySectionCsv``: %v\n", err)
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

Other parameters are passed through a pointer to a apiGetApiChartsByChartIdReportsBySectionCsvRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **section** | **string** | Only this section bucket. | 
 **version** | **string** | &#x60;published&#x60; (default) or &#x60;draft&#x60; (the unpublished designer draft). | 

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


## GetApiChartsByChartIdReportsBySectionJson

> GetApiChartsByChartIdReportsBySectionJson(ctx, chartId).Section(section).Version(version).Execute()

Chart report by section



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
	section := "section_example" // string | Only this section bucket. (optional)
	version := "version_example" // string | `published` (default) or `draft` (the unpublished designer draft). (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ChartsAPI.GetApiChartsByChartIdReportsBySectionJson(context.Background(), chartId).Section(section).Version(version).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChartsAPI.GetApiChartsByChartIdReportsBySectionJson``: %v\n", err)
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

Other parameters are passed through a pointer to a apiGetApiChartsByChartIdReportsBySectionJsonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **section** | **string** | Only this section bucket. | 
 **version** | **string** | &#x60;published&#x60; (default) or &#x60;draft&#x60; (the unpublished designer draft). | 

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


## GetApiChartsByChartIdReportsByZoneCsv

> GetApiChartsByChartIdReportsByZoneCsv(ctx, chartId).Zone(zone).Version(version).Execute()

Chart report by zone (CSV)



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
	zone := "zone_example" // string | Only this zone bucket. (optional)
	version := "version_example" // string | `published` (default) or `draft` (the unpublished designer draft). (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ChartsAPI.GetApiChartsByChartIdReportsByZoneCsv(context.Background(), chartId).Zone(zone).Version(version).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChartsAPI.GetApiChartsByChartIdReportsByZoneCsv``: %v\n", err)
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

Other parameters are passed through a pointer to a apiGetApiChartsByChartIdReportsByZoneCsvRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **zone** | **string** | Only this zone bucket. | 
 **version** | **string** | &#x60;published&#x60; (default) or &#x60;draft&#x60; (the unpublished designer draft). | 

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


## GetApiChartsByChartIdReportsByZoneJson

> GetApiChartsByChartIdReportsByZoneJson(ctx, chartId).Zone(zone).Version(version).Execute()

Chart report by zone



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
	zone := "zone_example" // string | Only this zone bucket. (optional)
	version := "version_example" // string | `published` (default) or `draft` (the unpublished designer draft). (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ChartsAPI.GetApiChartsByChartIdReportsByZoneJson(context.Background(), chartId).Zone(zone).Version(version).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChartsAPI.GetApiChartsByChartIdReportsByZoneJson``: %v\n", err)
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

Other parameters are passed through a pointer to a apiGetApiChartsByChartIdReportsByZoneJsonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **zone** | **string** | Only this zone bucket. | 
 **version** | **string** | &#x60;published&#x60; (default) or &#x60;draft&#x60; (the unpublished designer draft). | 

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


## GetApiChartsByChartIdReportsSummaryByCategoryCsv

> GetApiChartsByChartIdReportsSummaryByCategoryCsv(ctx, chartId).Version(version).Execute()

Chart summary by category (CSV)



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
	version := "version_example" // string | `published` (default) or `draft` (the unpublished designer draft). (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ChartsAPI.GetApiChartsByChartIdReportsSummaryByCategoryCsv(context.Background(), chartId).Version(version).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChartsAPI.GetApiChartsByChartIdReportsSummaryByCategoryCsv``: %v\n", err)
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

Other parameters are passed through a pointer to a apiGetApiChartsByChartIdReportsSummaryByCategoryCsvRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **version** | **string** | &#x60;published&#x60; (default) or &#x60;draft&#x60; (the unpublished designer draft). | 

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


## GetApiChartsByChartIdReportsSummaryByCategoryJson

> GetApiChartsByChartIdReportsSummaryByCategoryJson(ctx, chartId).Version(version).Execute()

Chart summary by category



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
	version := "version_example" // string | `published` (default) or `draft` (the unpublished designer draft). (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ChartsAPI.GetApiChartsByChartIdReportsSummaryByCategoryJson(context.Background(), chartId).Version(version).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChartsAPI.GetApiChartsByChartIdReportsSummaryByCategoryJson``: %v\n", err)
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

Other parameters are passed through a pointer to a apiGetApiChartsByChartIdReportsSummaryByCategoryJsonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **version** | **string** | &#x60;published&#x60; (default) or &#x60;draft&#x60; (the unpublished designer draft). | 

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


## GetApiChartsByChartIdReportsSummaryByCategoryKeyCsv

> GetApiChartsByChartIdReportsSummaryByCategoryKeyCsv(ctx, chartId).Version(version).Execute()

Chart summary by category key (CSV)



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
	version := "version_example" // string | `published` (default) or `draft` (the unpublished designer draft). (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ChartsAPI.GetApiChartsByChartIdReportsSummaryByCategoryKeyCsv(context.Background(), chartId).Version(version).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChartsAPI.GetApiChartsByChartIdReportsSummaryByCategoryKeyCsv``: %v\n", err)
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

Other parameters are passed through a pointer to a apiGetApiChartsByChartIdReportsSummaryByCategoryKeyCsvRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **version** | **string** | &#x60;published&#x60; (default) or &#x60;draft&#x60; (the unpublished designer draft). | 

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


## GetApiChartsByChartIdReportsSummaryByCategoryKeyJson

> GetApiChartsByChartIdReportsSummaryByCategoryKeyJson(ctx, chartId).Version(version).Execute()

Chart summary by category key



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
	version := "version_example" // string | `published` (default) or `draft` (the unpublished designer draft). (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ChartsAPI.GetApiChartsByChartIdReportsSummaryByCategoryKeyJson(context.Background(), chartId).Version(version).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChartsAPI.GetApiChartsByChartIdReportsSummaryByCategoryKeyJson``: %v\n", err)
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

Other parameters are passed through a pointer to a apiGetApiChartsByChartIdReportsSummaryByCategoryKeyJsonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **version** | **string** | &#x60;published&#x60; (default) or &#x60;draft&#x60; (the unpublished designer draft). | 

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


## GetApiChartsByChartIdReportsSummaryByCategoryLabelCsv

> GetApiChartsByChartIdReportsSummaryByCategoryLabelCsv(ctx, chartId).Version(version).Execute()

Chart summary by category label (CSV)



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
	version := "version_example" // string | `published` (default) or `draft` (the unpublished designer draft). (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ChartsAPI.GetApiChartsByChartIdReportsSummaryByCategoryLabelCsv(context.Background(), chartId).Version(version).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChartsAPI.GetApiChartsByChartIdReportsSummaryByCategoryLabelCsv``: %v\n", err)
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

Other parameters are passed through a pointer to a apiGetApiChartsByChartIdReportsSummaryByCategoryLabelCsvRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **version** | **string** | &#x60;published&#x60; (default) or &#x60;draft&#x60; (the unpublished designer draft). | 

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


## GetApiChartsByChartIdReportsSummaryByCategoryLabelJson

> GetApiChartsByChartIdReportsSummaryByCategoryLabelJson(ctx, chartId).Version(version).Execute()

Chart summary by category label



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
	version := "version_example" // string | `published` (default) or `draft` (the unpublished designer draft). (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ChartsAPI.GetApiChartsByChartIdReportsSummaryByCategoryLabelJson(context.Background(), chartId).Version(version).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChartsAPI.GetApiChartsByChartIdReportsSummaryByCategoryLabelJson``: %v\n", err)
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

Other parameters are passed through a pointer to a apiGetApiChartsByChartIdReportsSummaryByCategoryLabelJsonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **version** | **string** | &#x60;published&#x60; (default) or &#x60;draft&#x60; (the unpublished designer draft). | 

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


## GetApiChartsByChartIdReportsSummaryByObjectTypeCsv

> GetApiChartsByChartIdReportsSummaryByObjectTypeCsv(ctx, chartId).Version(version).Execute()

Chart summary by object type (CSV)



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
	version := "version_example" // string | `published` (default) or `draft` (the unpublished designer draft). (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ChartsAPI.GetApiChartsByChartIdReportsSummaryByObjectTypeCsv(context.Background(), chartId).Version(version).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChartsAPI.GetApiChartsByChartIdReportsSummaryByObjectTypeCsv``: %v\n", err)
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

Other parameters are passed through a pointer to a apiGetApiChartsByChartIdReportsSummaryByObjectTypeCsvRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **version** | **string** | &#x60;published&#x60; (default) or &#x60;draft&#x60; (the unpublished designer draft). | 

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


## GetApiChartsByChartIdReportsSummaryByObjectTypeJson

> GetApiChartsByChartIdReportsSummaryByObjectTypeJson(ctx, chartId).Version(version).Execute()

Chart summary by object type



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
	version := "version_example" // string | `published` (default) or `draft` (the unpublished designer draft). (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ChartsAPI.GetApiChartsByChartIdReportsSummaryByObjectTypeJson(context.Background(), chartId).Version(version).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChartsAPI.GetApiChartsByChartIdReportsSummaryByObjectTypeJson``: %v\n", err)
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

Other parameters are passed through a pointer to a apiGetApiChartsByChartIdReportsSummaryByObjectTypeJsonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **version** | **string** | &#x60;published&#x60; (default) or &#x60;draft&#x60; (the unpublished designer draft). | 

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


## GetApiChartsByChartIdReportsSummaryBySectionCsv

> GetApiChartsByChartIdReportsSummaryBySectionCsv(ctx, chartId).Version(version).Execute()

Chart summary by section (CSV)



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
	version := "version_example" // string | `published` (default) or `draft` (the unpublished designer draft). (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ChartsAPI.GetApiChartsByChartIdReportsSummaryBySectionCsv(context.Background(), chartId).Version(version).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChartsAPI.GetApiChartsByChartIdReportsSummaryBySectionCsv``: %v\n", err)
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

Other parameters are passed through a pointer to a apiGetApiChartsByChartIdReportsSummaryBySectionCsvRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **version** | **string** | &#x60;published&#x60; (default) or &#x60;draft&#x60; (the unpublished designer draft). | 

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


## GetApiChartsByChartIdReportsSummaryBySectionJson

> GetApiChartsByChartIdReportsSummaryBySectionJson(ctx, chartId).Version(version).Execute()

Chart summary by section



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
	version := "version_example" // string | `published` (default) or `draft` (the unpublished designer draft). (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ChartsAPI.GetApiChartsByChartIdReportsSummaryBySectionJson(context.Background(), chartId).Version(version).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChartsAPI.GetApiChartsByChartIdReportsSummaryBySectionJson``: %v\n", err)
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

Other parameters are passed through a pointer to a apiGetApiChartsByChartIdReportsSummaryBySectionJsonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **version** | **string** | &#x60;published&#x60; (default) or &#x60;draft&#x60; (the unpublished designer draft). | 

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


## GetApiChartsByChartIdReportsSummaryByZoneCsv

> GetApiChartsByChartIdReportsSummaryByZoneCsv(ctx, chartId).Version(version).Execute()

Chart summary by zone (CSV)



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
	version := "version_example" // string | `published` (default) or `draft` (the unpublished designer draft). (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ChartsAPI.GetApiChartsByChartIdReportsSummaryByZoneCsv(context.Background(), chartId).Version(version).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChartsAPI.GetApiChartsByChartIdReportsSummaryByZoneCsv``: %v\n", err)
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

Other parameters are passed through a pointer to a apiGetApiChartsByChartIdReportsSummaryByZoneCsvRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **version** | **string** | &#x60;published&#x60; (default) or &#x60;draft&#x60; (the unpublished designer draft). | 

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


## GetApiChartsByChartIdReportsSummaryByZoneJson

> GetApiChartsByChartIdReportsSummaryByZoneJson(ctx, chartId).Version(version).Execute()

Chart summary by zone



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
	version := "version_example" // string | `published` (default) or `draft` (the unpublished designer draft). (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ChartsAPI.GetApiChartsByChartIdReportsSummaryByZoneJson(context.Background(), chartId).Version(version).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChartsAPI.GetApiChartsByChartIdReportsSummaryByZoneJson``: %v\n", err)
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

Other parameters are passed through a pointer to a apiGetApiChartsByChartIdReportsSummaryByZoneJsonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **version** | **string** | &#x60;published&#x60; (default) or &#x60;draft&#x60; (the unpublished designer draft). | 

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


## GetApiChartsByChartIdThumbnail

> GetApiChartsByChartIdThumbnail(ctx, chartId).Version(version).Execute()

Chart thumbnail



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
	version := float32(8.14) // float32 | Version number (1-based). (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ChartsAPI.GetApiChartsByChartIdThumbnail(context.Background(), chartId).Version(version).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChartsAPI.GetApiChartsByChartIdThumbnail``: %v\n", err)
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

Other parameters are passed through a pointer to a apiGetApiChartsByChartIdThumbnailRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **version** | **float32** | Version number (1-based). | 

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


## GetApiChartsByChartIdVersions

> GetApiChartsByChartIdVersions200Response GetApiChartsByChartIdVersions(ctx, chartId).Limit(limit).Cursor(cursor).Execute()

List published versions



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
	limit := float32(8.14) // float32 | Page size (1..100, default 20). (optional)
	cursor := "cursor_example" // string | Opaque `nextCursor` value from the previous page. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ChartsAPI.GetApiChartsByChartIdVersions(context.Background(), chartId).Limit(limit).Cursor(cursor).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChartsAPI.GetApiChartsByChartIdVersions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetApiChartsByChartIdVersions`: GetApiChartsByChartIdVersions200Response
	fmt.Fprintf(os.Stdout, "Response from `ChartsAPI.GetApiChartsByChartIdVersions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**chartId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetApiChartsByChartIdVersionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **float32** | Page size (1..100, default 20). | 
 **cursor** | **string** | Opaque &#x60;nextCursor&#x60; value from the previous page. | 

### Return type

[**GetApiChartsByChartIdVersions200Response**](GetApiChartsByChartIdVersions200Response.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, multipart/form-data, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetApiChartsByChartIdZones

> GetApiChartsByChartIdZones(ctx, chartId).Version(version).Execute()

List the zones of a chart



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
	chartId := "chartId_example" // string | Chart id (UUID) or chart key.
	version := "version_example" // string | `published` or `draft` (the unpublished designer draft). (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ChartsAPI.GetApiChartsByChartIdZones(context.Background(), chartId).Version(version).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChartsAPI.GetApiChartsByChartIdZones``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**chartId** | **string** | Chart id (UUID) or chart key. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetApiChartsByChartIdZonesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **version** | **string** | &#x60;published&#x60; or &#x60;draft&#x60; (the unpublished designer draft). | 

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


## GetApiChartsTags

> GetApiChartsTags(ctx).Archived(archived).Execute()

List chart tags



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
	archived := "archived_example" // string | `false` (default) hides archived charts, `true` returns only archived charts, `all` returns both. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ChartsAPI.GetApiChartsTags(context.Background()).Archived(archived).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChartsAPI.GetApiChartsTags``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetApiChartsTagsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **archived** | **string** | &#x60;false&#x60; (default) hides archived charts, &#x60;true&#x60; returns only archived charts, &#x60;all&#x60; returns both. | 

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


## PatchApiChartsByChartId

> PatchApiChartsByChartId(ctx, chartId).PatchApiChartsByChartIdRequest(patchApiChartsByChartIdRequest).IdempotencyKey(idempotencyKey).Execute()

Update chart name or key



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
	chartId := "chartId_example" // string | Chart id (UUID) or chart key.
	patchApiChartsByChartIdRequest := *openapiclient.NewPatchApiChartsByChartIdRequest() // PatchApiChartsByChartIdRequest | 
	idempotencyKey := "idempotencyKey_example" // string | Required on transactional routes. Reuse the same key for every retry of one logical operation; the API replays the stored outcome. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ChartsAPI.PatchApiChartsByChartId(context.Background(), chartId).PatchApiChartsByChartIdRequest(patchApiChartsByChartIdRequest).IdempotencyKey(idempotencyKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChartsAPI.PatchApiChartsByChartId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**chartId** | **string** | Chart id (UUID) or chart key. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchApiChartsByChartIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **patchApiChartsByChartIdRequest** | [**PatchApiChartsByChartIdRequest**](PatchApiChartsByChartIdRequest.md) |  | 
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


## PostApiCharts

> PostApiCharts(ctx).PostApiChartsRequest(postApiChartsRequest).IdempotencyKey(idempotencyKey).Execute()

Create chart



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
	postApiChartsRequest := *openapiclient.NewPostApiChartsRequest("Name_example") // PostApiChartsRequest | 
	idempotencyKey := "idempotencyKey_example" // string | Required on transactional routes. Reuse the same key for every retry of one logical operation; the API replays the stored outcome. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ChartsAPI.PostApiCharts(context.Background()).PostApiChartsRequest(postApiChartsRequest).IdempotencyKey(idempotencyKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChartsAPI.PostApiCharts``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostApiChartsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **postApiChartsRequest** | [**PostApiChartsRequest**](PostApiChartsRequest.md) |  | 
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


## PostApiChartsByChartIdActionsArchive

> PostApiChartsByChartIdActionsArchive200Response PostApiChartsByChartIdActionsArchive(ctx, chartId).IdempotencyKey(idempotencyKey).Execute()

Archive chart



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
	resp, r, err := apiClient.ChartsAPI.PostApiChartsByChartIdActionsArchive(context.Background(), chartId).IdempotencyKey(idempotencyKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChartsAPI.PostApiChartsByChartIdActionsArchive``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostApiChartsByChartIdActionsArchive`: PostApiChartsByChartIdActionsArchive200Response
	fmt.Fprintf(os.Stdout, "Response from `ChartsAPI.PostApiChartsByChartIdActionsArchive`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**chartId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostApiChartsByChartIdActionsArchiveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **idempotencyKey** | **string** | Required on transactional routes. Reuse the same key for every retry of one logical operation; the API replays the stored outcome. | 

### Return type

[**PostApiChartsByChartIdActionsArchive200Response**](PostApiChartsByChartIdActionsArchive200Response.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, multipart/form-data, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostApiChartsByChartIdActionsCopyToWorkspace

> PostApiChartsByChartIdActionsCopyToWorkspace(ctx, chartId).PostApiChartsByChartIdActionsCopyToWorkspaceRequest(postApiChartsByChartIdActionsCopyToWorkspaceRequest).IdempotencyKey(idempotencyKey).Execute()

Copy a chart to another workspace



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
	postApiChartsByChartIdActionsCopyToWorkspaceRequest := *openapiclient.NewPostApiChartsByChartIdActionsCopyToWorkspaceRequest("WorkspaceId_example") // PostApiChartsByChartIdActionsCopyToWorkspaceRequest | 
	idempotencyKey := "idempotencyKey_example" // string | Required on transactional routes. Reuse the same key for every retry of one logical operation; the API replays the stored outcome. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ChartsAPI.PostApiChartsByChartIdActionsCopyToWorkspace(context.Background(), chartId).PostApiChartsByChartIdActionsCopyToWorkspaceRequest(postApiChartsByChartIdActionsCopyToWorkspaceRequest).IdempotencyKey(idempotencyKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChartsAPI.PostApiChartsByChartIdActionsCopyToWorkspace``: %v\n", err)
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

Other parameters are passed through a pointer to a apiPostApiChartsByChartIdActionsCopyToWorkspaceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **postApiChartsByChartIdActionsCopyToWorkspaceRequest** | [**PostApiChartsByChartIdActionsCopyToWorkspaceRequest**](PostApiChartsByChartIdActionsCopyToWorkspaceRequest.md) |  | 
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


## PostApiChartsByChartIdActionsUnarchive

> PostApiChartsByChartIdActionsArchive200Response PostApiChartsByChartIdActionsUnarchive(ctx, chartId).IdempotencyKey(idempotencyKey).Execute()

Unarchive chart



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
	resp, r, err := apiClient.ChartsAPI.PostApiChartsByChartIdActionsUnarchive(context.Background(), chartId).IdempotencyKey(idempotencyKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChartsAPI.PostApiChartsByChartIdActionsUnarchive``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostApiChartsByChartIdActionsUnarchive`: PostApiChartsByChartIdActionsArchive200Response
	fmt.Fprintf(os.Stdout, "Response from `ChartsAPI.PostApiChartsByChartIdActionsUnarchive`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**chartId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostApiChartsByChartIdActionsUnarchiveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **idempotencyKey** | **string** | Required on transactional routes. Reuse the same key for every retry of one logical operation; the API replays the stored outcome. | 

### Return type

[**PostApiChartsByChartIdActionsArchive200Response**](PostApiChartsByChartIdActionsArchive200Response.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, multipart/form-data, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostApiChartsByChartIdActionsValidate

> PostApiChartsByChartIdActionsValidate(ctx, chartId).PostApiChartsByChartIdActionsValidateRequest(postApiChartsByChartIdActionsValidateRequest).Version(version).IdempotencyKey(idempotencyKey).Execute()

Validate a chart map (dry run)



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
	postApiChartsByChartIdActionsValidateRequest := *openapiclient.NewPostApiChartsByChartIdActionsValidateRequest([]openapiclient.PostApiChartsByChartIdActionsValidateRequestSeatsInner{*openapiclient.NewPostApiChartsByChartIdActionsValidateRequestSeatsInner("Label_example", float32(123), float32(123))}) // PostApiChartsByChartIdActionsValidateRequest | 
	version := "version_example" // string | `published` or `draft`. Omitted: the draft when there is one, else the published map. (optional)
	idempotencyKey := "idempotencyKey_example" // string | Required on transactional routes. Reuse the same key for every retry of one logical operation; the API replays the stored outcome. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ChartsAPI.PostApiChartsByChartIdActionsValidate(context.Background(), chartId).PostApiChartsByChartIdActionsValidateRequest(postApiChartsByChartIdActionsValidateRequest).Version(version).IdempotencyKey(idempotencyKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChartsAPI.PostApiChartsByChartIdActionsValidate``: %v\n", err)
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

Other parameters are passed through a pointer to a apiPostApiChartsByChartIdActionsValidateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **postApiChartsByChartIdActionsValidateRequest** | [**PostApiChartsByChartIdActionsValidateRequest**](PostApiChartsByChartIdActionsValidateRequest.md) |  | 
 **version** | **string** | &#x60;published&#x60; or &#x60;draft&#x60;. Omitted: the draft when there is one, else the published map. | 
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


## PostApiChartsByChartIdDuplicate

> PostApiChartsByChartIdDuplicate(ctx, chartId).PostApiChartsByChartIdDuplicateRequest(postApiChartsByChartIdDuplicateRequest).IdempotencyKey(idempotencyKey).Execute()

Duplicate chart



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
	postApiChartsByChartIdDuplicateRequest := *openapiclient.NewPostApiChartsByChartIdDuplicateRequest() // PostApiChartsByChartIdDuplicateRequest | 
	idempotencyKey := "idempotencyKey_example" // string | Required on transactional routes. Reuse the same key for every retry of one logical operation; the API replays the stored outcome. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ChartsAPI.PostApiChartsByChartIdDuplicate(context.Background(), chartId).PostApiChartsByChartIdDuplicateRequest(postApiChartsByChartIdDuplicateRequest).IdempotencyKey(idempotencyKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChartsAPI.PostApiChartsByChartIdDuplicate``: %v\n", err)
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

Other parameters are passed through a pointer to a apiPostApiChartsByChartIdDuplicateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **postApiChartsByChartIdDuplicateRequest** | [**PostApiChartsByChartIdDuplicateRequest**](PostApiChartsByChartIdDuplicateRequest.md) |  | 
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


## PostApiChartsByChartIdMapDraftActionsDiscard

> PostApiChartsByChartIdMapDraftActionsDiscard(ctx, chartId).IdempotencyKey(idempotencyKey).Execute()

Discard chart draft



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
	r, err := apiClient.ChartsAPI.PostApiChartsByChartIdMapDraftActionsDiscard(context.Background(), chartId).IdempotencyKey(idempotencyKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChartsAPI.PostApiChartsByChartIdMapDraftActionsDiscard``: %v\n", err)
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

Other parameters are passed through a pointer to a apiPostApiChartsByChartIdMapDraftActionsDiscardRequest struct via the builder pattern


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


## PostApiChartsByChartIdMapPublish

> PostApiChartsByChartIdMapPublish(ctx, chartId).PostApiChartsByChartIdMapPublishRequest(postApiChartsByChartIdMapPublishRequest).IdempotencyKey(idempotencyKey).Execute()

Publish chart map template



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
	postApiChartsByChartIdMapPublishRequest := *openapiclient.NewPostApiChartsByChartIdMapPublishRequest() // PostApiChartsByChartIdMapPublishRequest | 
	idempotencyKey := "idempotencyKey_example" // string | Required on transactional routes. Reuse the same key for every retry of one logical operation; the API replays the stored outcome. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ChartsAPI.PostApiChartsByChartIdMapPublish(context.Background(), chartId).PostApiChartsByChartIdMapPublishRequest(postApiChartsByChartIdMapPublishRequest).IdempotencyKey(idempotencyKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChartsAPI.PostApiChartsByChartIdMapPublish``: %v\n", err)
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

Other parameters are passed through a pointer to a apiPostApiChartsByChartIdMapPublishRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **postApiChartsByChartIdMapPublishRequest** | [**PostApiChartsByChartIdMapPublishRequest**](PostApiChartsByChartIdMapPublishRequest.md) |  | 
 **idempotencyKey** | **string** | Required on transactional routes. Reuse the same key for every retry of one logical operation; the API replays the stored outcome. | 

### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json, multipart/form-data, text/plain
- **Accept**: application/json, multipart/form-data, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostApiChartsByChartIdPreviewLink

> PostApiChartsByChartIdPreviewLink(ctx, chartId).PostApiChartsByChartIdPreviewLinkRequest(postApiChartsByChartIdPreviewLinkRequest).IdempotencyKey(idempotencyKey).Execute()

Create a public preview link



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
	postApiChartsByChartIdPreviewLinkRequest := *openapiclient.NewPostApiChartsByChartIdPreviewLinkRequest() // PostApiChartsByChartIdPreviewLinkRequest | 
	idempotencyKey := "idempotencyKey_example" // string | Required on transactional routes. Reuse the same key for every retry of one logical operation; the API replays the stored outcome. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ChartsAPI.PostApiChartsByChartIdPreviewLink(context.Background(), chartId).PostApiChartsByChartIdPreviewLinkRequest(postApiChartsByChartIdPreviewLinkRequest).IdempotencyKey(idempotencyKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChartsAPI.PostApiChartsByChartIdPreviewLink``: %v\n", err)
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

Other parameters are passed through a pointer to a apiPostApiChartsByChartIdPreviewLinkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **postApiChartsByChartIdPreviewLinkRequest** | [**PostApiChartsByChartIdPreviewLinkRequest**](PostApiChartsByChartIdPreviewLinkRequest.md) |  | 
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


## PostApiChartsByChartIdVersionsByVersionActionsRestoreToDraft

> PostApiChartsByChartIdVersionsByVersionActionsRestoreToDraft(ctx, chartId, version).IdempotencyKey(idempotencyKey).Execute()

Restore a version into the draft



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
	version := float32(8.14) // float32 | Version number (1-based).
	idempotencyKey := "idempotencyKey_example" // string | Required on transactional routes. Reuse the same key for every retry of one logical operation; the API replays the stored outcome. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ChartsAPI.PostApiChartsByChartIdVersionsByVersionActionsRestoreToDraft(context.Background(), chartId, version).IdempotencyKey(idempotencyKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChartsAPI.PostApiChartsByChartIdVersionsByVersionActionsRestoreToDraft``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**chartId** | **string** |  | 
**version** | **float32** | Version number (1-based). | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostApiChartsByChartIdVersionsByVersionActionsRestoreToDraftRequest struct via the builder pattern


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


## PutApiChartsByChartIdMapDraft

> PutApiChartsByChartIdMapDraft(ctx, chartId).PostApiChartsByChartIdActionsValidateRequest(postApiChartsByChartIdActionsValidateRequest).IdempotencyKey(idempotencyKey).Execute()

Save chart map draft



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
	postApiChartsByChartIdActionsValidateRequest := *openapiclient.NewPostApiChartsByChartIdActionsValidateRequest([]openapiclient.PostApiChartsByChartIdActionsValidateRequestSeatsInner{*openapiclient.NewPostApiChartsByChartIdActionsValidateRequestSeatsInner("Label_example", float32(123), float32(123))}) // PostApiChartsByChartIdActionsValidateRequest | 
	idempotencyKey := "idempotencyKey_example" // string | Required on transactional routes. Reuse the same key for every retry of one logical operation; the API replays the stored outcome. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ChartsAPI.PutApiChartsByChartIdMapDraft(context.Background(), chartId).PostApiChartsByChartIdActionsValidateRequest(postApiChartsByChartIdActionsValidateRequest).IdempotencyKey(idempotencyKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChartsAPI.PutApiChartsByChartIdMapDraft``: %v\n", err)
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

Other parameters are passed through a pointer to a apiPutApiChartsByChartIdMapDraftRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **postApiChartsByChartIdActionsValidateRequest** | [**PostApiChartsByChartIdActionsValidateRequest**](PostApiChartsByChartIdActionsValidateRequest.md) |  | 
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


## PutApiChartsByChartIdTags

> PutApiChartsByChartIdTags(ctx, chartId).PutApiChartsByChartIdTagsRequest(putApiChartsByChartIdTagsRequest).IdempotencyKey(idempotencyKey).Execute()

Replace chart tags



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
	putApiChartsByChartIdTagsRequest := *openapiclient.NewPutApiChartsByChartIdTagsRequest([]string{"Tags_example"}) // PutApiChartsByChartIdTagsRequest | 
	idempotencyKey := "idempotencyKey_example" // string | Required on transactional routes. Reuse the same key for every retry of one logical operation; the API replays the stored outcome. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ChartsAPI.PutApiChartsByChartIdTags(context.Background(), chartId).PutApiChartsByChartIdTagsRequest(putApiChartsByChartIdTagsRequest).IdempotencyKey(idempotencyKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChartsAPI.PutApiChartsByChartIdTags``: %v\n", err)
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

Other parameters are passed through a pointer to a apiPutApiChartsByChartIdTagsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **putApiChartsByChartIdTagsRequest** | [**PutApiChartsByChartIdTagsRequest**](PutApiChartsByChartIdTagsRequest.md) |  | 
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


## PutApiChartsByChartIdVersionsByVersionThumbnail

> PutApiChartsByChartIdVersionsByVersionThumbnail(ctx, chartId, version).PutApiChartsByChartIdVersionsByVersionThumbnailRequest(putApiChartsByChartIdVersionsByVersionThumbnailRequest).IdempotencyKey(idempotencyKey).Execute()

Attach a rendered thumbnail to a version



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
	version := float32(8.14) // float32 | Version number (1-based).
	putApiChartsByChartIdVersionsByVersionThumbnailRequest := *openapiclient.NewPutApiChartsByChartIdVersionsByVersionThumbnailRequest("Image_example") // PutApiChartsByChartIdVersionsByVersionThumbnailRequest | 
	idempotencyKey := "idempotencyKey_example" // string | Required on transactional routes. Reuse the same key for every retry of one logical operation; the API replays the stored outcome. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ChartsAPI.PutApiChartsByChartIdVersionsByVersionThumbnail(context.Background(), chartId, version).PutApiChartsByChartIdVersionsByVersionThumbnailRequest(putApiChartsByChartIdVersionsByVersionThumbnailRequest).IdempotencyKey(idempotencyKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChartsAPI.PutApiChartsByChartIdVersionsByVersionThumbnail``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**chartId** | **string** |  | 
**version** | **float32** | Version number (1-based). | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutApiChartsByChartIdVersionsByVersionThumbnailRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **putApiChartsByChartIdVersionsByVersionThumbnailRequest** | [**PutApiChartsByChartIdVersionsByVersionThumbnailRequest**](PutApiChartsByChartIdVersionsByVersionThumbnailRequest.md) |  | 
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

