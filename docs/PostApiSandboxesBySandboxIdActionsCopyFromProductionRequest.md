# PostApiSandboxesBySandboxIdActionsCopyFromProductionRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ChartIds** | Pointer to **[]string** | Charts of the production workspace to copy (default: every non-archived chart, up to 50, by name). | [optional] 
**IncludeEvents** | Pointer to **bool** | Also copy the events of each copied published chart (up to 200 per call, newest first) with their for-sale configuration: for-sale flags, channels (new channel keys), GA place quotas, category overrides, event categories, table booking, sales rules, synced prices and the event-level layout. Seasons and partial seasons of the copied events are copied too (fresh ids, remapped members). Nothing sold or held is copied. | [optional] 

## Methods

### NewPostApiSandboxesBySandboxIdActionsCopyFromProductionRequest

`func NewPostApiSandboxesBySandboxIdActionsCopyFromProductionRequest() *PostApiSandboxesBySandboxIdActionsCopyFromProductionRequest`

NewPostApiSandboxesBySandboxIdActionsCopyFromProductionRequest instantiates a new PostApiSandboxesBySandboxIdActionsCopyFromProductionRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostApiSandboxesBySandboxIdActionsCopyFromProductionRequestWithDefaults

`func NewPostApiSandboxesBySandboxIdActionsCopyFromProductionRequestWithDefaults() *PostApiSandboxesBySandboxIdActionsCopyFromProductionRequest`

NewPostApiSandboxesBySandboxIdActionsCopyFromProductionRequestWithDefaults instantiates a new PostApiSandboxesBySandboxIdActionsCopyFromProductionRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChartIds

`func (o *PostApiSandboxesBySandboxIdActionsCopyFromProductionRequest) GetChartIds() []string`

GetChartIds returns the ChartIds field if non-nil, zero value otherwise.

### GetChartIdsOk

`func (o *PostApiSandboxesBySandboxIdActionsCopyFromProductionRequest) GetChartIdsOk() (*[]string, bool)`

GetChartIdsOk returns a tuple with the ChartIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChartIds

`func (o *PostApiSandboxesBySandboxIdActionsCopyFromProductionRequest) SetChartIds(v []string)`

SetChartIds sets ChartIds field to given value.

### HasChartIds

`func (o *PostApiSandboxesBySandboxIdActionsCopyFromProductionRequest) HasChartIds() bool`

HasChartIds returns a boolean if a field has been set.

### GetIncludeEvents

`func (o *PostApiSandboxesBySandboxIdActionsCopyFromProductionRequest) GetIncludeEvents() bool`

GetIncludeEvents returns the IncludeEvents field if non-nil, zero value otherwise.

### GetIncludeEventsOk

`func (o *PostApiSandboxesBySandboxIdActionsCopyFromProductionRequest) GetIncludeEventsOk() (*bool, bool)`

GetIncludeEventsOk returns a tuple with the IncludeEvents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeEvents

`func (o *PostApiSandboxesBySandboxIdActionsCopyFromProductionRequest) SetIncludeEvents(v bool)`

SetIncludeEvents sets IncludeEvents field to given value.

### HasIncludeEvents

`func (o *PostApiSandboxesBySandboxIdActionsCopyFromProductionRequest) HasIncludeEvents() bool`

HasIncludeEvents returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


