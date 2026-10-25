# PostApiSeasonsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**ChartId** | Pointer to **NullableString** |  | [optional] 
**SourceEventId** | Pointer to **NullableString** |  | [optional] 
**EventKeys** | Pointer to **[]string** | Create one season event per key (requires chartId). | [optional] 
**NumberOfEvents** | Pointer to **int32** | Create this many season events (requires chartId). | [optional] 

## Methods

### NewPostApiSeasonsRequest

`func NewPostApiSeasonsRequest(name string, ) *PostApiSeasonsRequest`

NewPostApiSeasonsRequest instantiates a new PostApiSeasonsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostApiSeasonsRequestWithDefaults

`func NewPostApiSeasonsRequestWithDefaults() *PostApiSeasonsRequest`

NewPostApiSeasonsRequestWithDefaults instantiates a new PostApiSeasonsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *PostApiSeasonsRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PostApiSeasonsRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PostApiSeasonsRequest) SetName(v string)`

SetName sets Name field to given value.


### GetChartId

`func (o *PostApiSeasonsRequest) GetChartId() string`

GetChartId returns the ChartId field if non-nil, zero value otherwise.

### GetChartIdOk

`func (o *PostApiSeasonsRequest) GetChartIdOk() (*string, bool)`

GetChartIdOk returns a tuple with the ChartId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChartId

`func (o *PostApiSeasonsRequest) SetChartId(v string)`

SetChartId sets ChartId field to given value.

### HasChartId

`func (o *PostApiSeasonsRequest) HasChartId() bool`

HasChartId returns a boolean if a field has been set.

### SetChartIdNil

`func (o *PostApiSeasonsRequest) SetChartIdNil(b bool)`

 SetChartIdNil sets the value for ChartId to be an explicit nil

### UnsetChartId
`func (o *PostApiSeasonsRequest) UnsetChartId()`

UnsetChartId ensures that no value is present for ChartId, not even an explicit nil
### GetSourceEventId

`func (o *PostApiSeasonsRequest) GetSourceEventId() string`

GetSourceEventId returns the SourceEventId field if non-nil, zero value otherwise.

### GetSourceEventIdOk

`func (o *PostApiSeasonsRequest) GetSourceEventIdOk() (*string, bool)`

GetSourceEventIdOk returns a tuple with the SourceEventId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceEventId

`func (o *PostApiSeasonsRequest) SetSourceEventId(v string)`

SetSourceEventId sets SourceEventId field to given value.

### HasSourceEventId

`func (o *PostApiSeasonsRequest) HasSourceEventId() bool`

HasSourceEventId returns a boolean if a field has been set.

### SetSourceEventIdNil

`func (o *PostApiSeasonsRequest) SetSourceEventIdNil(b bool)`

 SetSourceEventIdNil sets the value for SourceEventId to be an explicit nil

### UnsetSourceEventId
`func (o *PostApiSeasonsRequest) UnsetSourceEventId()`

UnsetSourceEventId ensures that no value is present for SourceEventId, not even an explicit nil
### GetEventKeys

`func (o *PostApiSeasonsRequest) GetEventKeys() []string`

GetEventKeys returns the EventKeys field if non-nil, zero value otherwise.

### GetEventKeysOk

`func (o *PostApiSeasonsRequest) GetEventKeysOk() (*[]string, bool)`

GetEventKeysOk returns a tuple with the EventKeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventKeys

`func (o *PostApiSeasonsRequest) SetEventKeys(v []string)`

SetEventKeys sets EventKeys field to given value.

### HasEventKeys

`func (o *PostApiSeasonsRequest) HasEventKeys() bool`

HasEventKeys returns a boolean if a field has been set.

### GetNumberOfEvents

`func (o *PostApiSeasonsRequest) GetNumberOfEvents() int32`

GetNumberOfEvents returns the NumberOfEvents field if non-nil, zero value otherwise.

### GetNumberOfEventsOk

`func (o *PostApiSeasonsRequest) GetNumberOfEventsOk() (*int32, bool)`

GetNumberOfEventsOk returns a tuple with the NumberOfEvents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberOfEvents

`func (o *PostApiSeasonsRequest) SetNumberOfEvents(v int32)`

SetNumberOfEvents sets NumberOfEvents field to given value.

### HasNumberOfEvents

`func (o *PostApiSeasonsRequest) HasNumberOfEvents() bool`

HasNumberOfEvents returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


