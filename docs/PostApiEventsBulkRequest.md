# PostApiEventsBulkRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ChartId** | **string** |  | 
**Events** | [**[]PostApiEventsBulkRequestEventsInner**](PostApiEventsBulkRequestEventsInner.md) |  | 
**SeasonId** | Pointer to **string** |  | [optional] 

## Methods

### NewPostApiEventsBulkRequest

`func NewPostApiEventsBulkRequest(chartId string, events []PostApiEventsBulkRequestEventsInner, ) *PostApiEventsBulkRequest`

NewPostApiEventsBulkRequest instantiates a new PostApiEventsBulkRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostApiEventsBulkRequestWithDefaults

`func NewPostApiEventsBulkRequestWithDefaults() *PostApiEventsBulkRequest`

NewPostApiEventsBulkRequestWithDefaults instantiates a new PostApiEventsBulkRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChartId

`func (o *PostApiEventsBulkRequest) GetChartId() string`

GetChartId returns the ChartId field if non-nil, zero value otherwise.

### GetChartIdOk

`func (o *PostApiEventsBulkRequest) GetChartIdOk() (*string, bool)`

GetChartIdOk returns a tuple with the ChartId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChartId

`func (o *PostApiEventsBulkRequest) SetChartId(v string)`

SetChartId sets ChartId field to given value.


### GetEvents

`func (o *PostApiEventsBulkRequest) GetEvents() []PostApiEventsBulkRequestEventsInner`

GetEvents returns the Events field if non-nil, zero value otherwise.

### GetEventsOk

`func (o *PostApiEventsBulkRequest) GetEventsOk() (*[]PostApiEventsBulkRequestEventsInner, bool)`

GetEventsOk returns a tuple with the Events field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvents

`func (o *PostApiEventsBulkRequest) SetEvents(v []PostApiEventsBulkRequestEventsInner)`

SetEvents sets Events field to given value.


### GetSeasonId

`func (o *PostApiEventsBulkRequest) GetSeasonId() string`

GetSeasonId returns the SeasonId field if non-nil, zero value otherwise.

### GetSeasonIdOk

`func (o *PostApiEventsBulkRequest) GetSeasonIdOk() (*string, bool)`

GetSeasonIdOk returns a tuple with the SeasonId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeasonId

`func (o *PostApiEventsBulkRequest) SetSeasonId(v string)`

SetSeasonId sets SeasonId field to given value.

### HasSeasonId

`func (o *PostApiEventsBulkRequest) HasSeasonId() bool`

HasSeasonId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


