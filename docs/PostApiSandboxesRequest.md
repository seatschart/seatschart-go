# PostApiSandboxesRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DisplayName** | **string** |  | 
**CopyCharts** | Pointer to **bool** | Copy the production charts (every non-archived chart, up to 50). Default false. | [optional] 
**ChartIds** | Pointer to **[]string** | Charts of the production workspace to copy (default: every non-archived chart, up to 50, by name). | [optional] 
**IncludeEvents** | Pointer to **bool** | Also copy the events of each copied published chart (up to 200 per call, newest first) with their for-sale configuration: for-sale flags, channels (new channel keys), GA place quotas, category overrides, event categories, table booking, sales rules, synced prices and the event-level layout. Seasons and partial seasons of the copied events are copied too (fresh ids, remapped members). Nothing sold or held is copied. | [optional] 
**CreateAdminKey** | Pointer to **bool** | Mint a full-access &#x60;sk_test_&#x60; key for the sandbox, returned once in &#x60;adminKey&#x60;. Default true. | [optional] 

## Methods

### NewPostApiSandboxesRequest

`func NewPostApiSandboxesRequest(displayName string, ) *PostApiSandboxesRequest`

NewPostApiSandboxesRequest instantiates a new PostApiSandboxesRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostApiSandboxesRequestWithDefaults

`func NewPostApiSandboxesRequestWithDefaults() *PostApiSandboxesRequest`

NewPostApiSandboxesRequestWithDefaults instantiates a new PostApiSandboxesRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDisplayName

`func (o *PostApiSandboxesRequest) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *PostApiSandboxesRequest) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *PostApiSandboxesRequest) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.


### GetCopyCharts

`func (o *PostApiSandboxesRequest) GetCopyCharts() bool`

GetCopyCharts returns the CopyCharts field if non-nil, zero value otherwise.

### GetCopyChartsOk

`func (o *PostApiSandboxesRequest) GetCopyChartsOk() (*bool, bool)`

GetCopyChartsOk returns a tuple with the CopyCharts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCopyCharts

`func (o *PostApiSandboxesRequest) SetCopyCharts(v bool)`

SetCopyCharts sets CopyCharts field to given value.

### HasCopyCharts

`func (o *PostApiSandboxesRequest) HasCopyCharts() bool`

HasCopyCharts returns a boolean if a field has been set.

### GetChartIds

`func (o *PostApiSandboxesRequest) GetChartIds() []string`

GetChartIds returns the ChartIds field if non-nil, zero value otherwise.

### GetChartIdsOk

`func (o *PostApiSandboxesRequest) GetChartIdsOk() (*[]string, bool)`

GetChartIdsOk returns a tuple with the ChartIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChartIds

`func (o *PostApiSandboxesRequest) SetChartIds(v []string)`

SetChartIds sets ChartIds field to given value.

### HasChartIds

`func (o *PostApiSandboxesRequest) HasChartIds() bool`

HasChartIds returns a boolean if a field has been set.

### GetIncludeEvents

`func (o *PostApiSandboxesRequest) GetIncludeEvents() bool`

GetIncludeEvents returns the IncludeEvents field if non-nil, zero value otherwise.

### GetIncludeEventsOk

`func (o *PostApiSandboxesRequest) GetIncludeEventsOk() (*bool, bool)`

GetIncludeEventsOk returns a tuple with the IncludeEvents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeEvents

`func (o *PostApiSandboxesRequest) SetIncludeEvents(v bool)`

SetIncludeEvents sets IncludeEvents field to given value.

### HasIncludeEvents

`func (o *PostApiSandboxesRequest) HasIncludeEvents() bool`

HasIncludeEvents returns a boolean if a field has been set.

### GetCreateAdminKey

`func (o *PostApiSandboxesRequest) GetCreateAdminKey() bool`

GetCreateAdminKey returns the CreateAdminKey field if non-nil, zero value otherwise.

### GetCreateAdminKeyOk

`func (o *PostApiSandboxesRequest) GetCreateAdminKeyOk() (*bool, bool)`

GetCreateAdminKeyOk returns a tuple with the CreateAdminKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateAdminKey

`func (o *PostApiSandboxesRequest) SetCreateAdminKey(v bool)`

SetCreateAdminKey sets CreateAdminKey field to given value.

### HasCreateAdminKey

`func (o *PostApiSandboxesRequest) HasCreateAdminKey() bool`

HasCreateAdminKey returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


