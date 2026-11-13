# GetApiSandboxes200ResponseItemsInnerUsage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Month** | **string** | UTC month &#x60;YYYY-MM&#x60;. | 
**UsedObjects** | **int32** | Objects booked in the sandbox this month: shown, never billed or counted against the plan. | 
**Charts** | **int32** |  | 
**Events** | **int32** |  | 
**ActiveApiKeys** | **int32** |  | 

## Methods

### NewGetApiSandboxes200ResponseItemsInnerUsage

`func NewGetApiSandboxes200ResponseItemsInnerUsage(month string, usedObjects int32, charts int32, events int32, activeApiKeys int32, ) *GetApiSandboxes200ResponseItemsInnerUsage`

NewGetApiSandboxes200ResponseItemsInnerUsage instantiates a new GetApiSandboxes200ResponseItemsInnerUsage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetApiSandboxes200ResponseItemsInnerUsageWithDefaults

`func NewGetApiSandboxes200ResponseItemsInnerUsageWithDefaults() *GetApiSandboxes200ResponseItemsInnerUsage`

NewGetApiSandboxes200ResponseItemsInnerUsageWithDefaults instantiates a new GetApiSandboxes200ResponseItemsInnerUsage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMonth

`func (o *GetApiSandboxes200ResponseItemsInnerUsage) GetMonth() string`

GetMonth returns the Month field if non-nil, zero value otherwise.

### GetMonthOk

`func (o *GetApiSandboxes200ResponseItemsInnerUsage) GetMonthOk() (*string, bool)`

GetMonthOk returns a tuple with the Month field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonth

`func (o *GetApiSandboxes200ResponseItemsInnerUsage) SetMonth(v string)`

SetMonth sets Month field to given value.


### GetUsedObjects

`func (o *GetApiSandboxes200ResponseItemsInnerUsage) GetUsedObjects() int32`

GetUsedObjects returns the UsedObjects field if non-nil, zero value otherwise.

### GetUsedObjectsOk

`func (o *GetApiSandboxes200ResponseItemsInnerUsage) GetUsedObjectsOk() (*int32, bool)`

GetUsedObjectsOk returns a tuple with the UsedObjects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsedObjects

`func (o *GetApiSandboxes200ResponseItemsInnerUsage) SetUsedObjects(v int32)`

SetUsedObjects sets UsedObjects field to given value.


### GetCharts

`func (o *GetApiSandboxes200ResponseItemsInnerUsage) GetCharts() int32`

GetCharts returns the Charts field if non-nil, zero value otherwise.

### GetChartsOk

`func (o *GetApiSandboxes200ResponseItemsInnerUsage) GetChartsOk() (*int32, bool)`

GetChartsOk returns a tuple with the Charts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCharts

`func (o *GetApiSandboxes200ResponseItemsInnerUsage) SetCharts(v int32)`

SetCharts sets Charts field to given value.


### GetEvents

`func (o *GetApiSandboxes200ResponseItemsInnerUsage) GetEvents() int32`

GetEvents returns the Events field if non-nil, zero value otherwise.

### GetEventsOk

`func (o *GetApiSandboxes200ResponseItemsInnerUsage) GetEventsOk() (*int32, bool)`

GetEventsOk returns a tuple with the Events field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvents

`func (o *GetApiSandboxes200ResponseItemsInnerUsage) SetEvents(v int32)`

SetEvents sets Events field to given value.


### GetActiveApiKeys

`func (o *GetApiSandboxes200ResponseItemsInnerUsage) GetActiveApiKeys() int32`

GetActiveApiKeys returns the ActiveApiKeys field if non-nil, zero value otherwise.

### GetActiveApiKeysOk

`func (o *GetApiSandboxes200ResponseItemsInnerUsage) GetActiveApiKeysOk() (*int32, bool)`

GetActiveApiKeysOk returns a tuple with the ActiveApiKeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveApiKeys

`func (o *GetApiSandboxes200ResponseItemsInnerUsage) SetActiveApiKeys(v int32)`

SetActiveApiKeys sets ActiveApiKeys field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


