# GetApiStatus200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GeneratedAt** | **time.Time** |  | 
**HistoryDays** | **int32** | Days in every &#x60;days&#x60; array (90). | 
**SampleIntervalSeconds** | **NullableInt32** | Sampler interval; null before the first sample. | 
**LastSampleAt** | **NullableTime** | Time of the latest sample; null when nothing was sampled yet. | 
**Stale** | **bool** | True when the latest sample is older than three intervals (at least 3 minutes): the sampler is not running and current states are &#x60;unknown&#x60;. | 
**Overall** | **string** |  | 
**Components** | [**[]GetApiStatus200ResponseComponentsInner**](GetApiStatus200ResponseComponentsInner.md) |  | 
**Incidents** | [**[]GetApiStatus200ResponseIncidentsInner**](GetApiStatus200ResponseIncidentsInner.md) | Newest first, last 90 days, at most 50. | 
**Maintenance** | [**[]GetApiStatus200ResponseMaintenanceInner**](GetApiStatus200ResponseMaintenanceInner.md) | Scheduled and in-progress windows, and those completed in the last 7 days. | 

## Methods

### NewGetApiStatus200Response

`func NewGetApiStatus200Response(generatedAt time.Time, historyDays int32, sampleIntervalSeconds NullableInt32, lastSampleAt NullableTime, stale bool, overall string, components []GetApiStatus200ResponseComponentsInner, incidents []GetApiStatus200ResponseIncidentsInner, maintenance []GetApiStatus200ResponseMaintenanceInner, ) *GetApiStatus200Response`

NewGetApiStatus200Response instantiates a new GetApiStatus200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetApiStatus200ResponseWithDefaults

`func NewGetApiStatus200ResponseWithDefaults() *GetApiStatus200Response`

NewGetApiStatus200ResponseWithDefaults instantiates a new GetApiStatus200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGeneratedAt

`func (o *GetApiStatus200Response) GetGeneratedAt() time.Time`

GetGeneratedAt returns the GeneratedAt field if non-nil, zero value otherwise.

### GetGeneratedAtOk

`func (o *GetApiStatus200Response) GetGeneratedAtOk() (*time.Time, bool)`

GetGeneratedAtOk returns a tuple with the GeneratedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeneratedAt

`func (o *GetApiStatus200Response) SetGeneratedAt(v time.Time)`

SetGeneratedAt sets GeneratedAt field to given value.


### GetHistoryDays

`func (o *GetApiStatus200Response) GetHistoryDays() int32`

GetHistoryDays returns the HistoryDays field if non-nil, zero value otherwise.

### GetHistoryDaysOk

`func (o *GetApiStatus200Response) GetHistoryDaysOk() (*int32, bool)`

GetHistoryDaysOk returns a tuple with the HistoryDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHistoryDays

`func (o *GetApiStatus200Response) SetHistoryDays(v int32)`

SetHistoryDays sets HistoryDays field to given value.


### GetSampleIntervalSeconds

`func (o *GetApiStatus200Response) GetSampleIntervalSeconds() int32`

GetSampleIntervalSeconds returns the SampleIntervalSeconds field if non-nil, zero value otherwise.

### GetSampleIntervalSecondsOk

`func (o *GetApiStatus200Response) GetSampleIntervalSecondsOk() (*int32, bool)`

GetSampleIntervalSecondsOk returns a tuple with the SampleIntervalSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSampleIntervalSeconds

`func (o *GetApiStatus200Response) SetSampleIntervalSeconds(v int32)`

SetSampleIntervalSeconds sets SampleIntervalSeconds field to given value.


### SetSampleIntervalSecondsNil

`func (o *GetApiStatus200Response) SetSampleIntervalSecondsNil(b bool)`

 SetSampleIntervalSecondsNil sets the value for SampleIntervalSeconds to be an explicit nil

### UnsetSampleIntervalSeconds
`func (o *GetApiStatus200Response) UnsetSampleIntervalSeconds()`

UnsetSampleIntervalSeconds ensures that no value is present for SampleIntervalSeconds, not even an explicit nil
### GetLastSampleAt

`func (o *GetApiStatus200Response) GetLastSampleAt() time.Time`

GetLastSampleAt returns the LastSampleAt field if non-nil, zero value otherwise.

### GetLastSampleAtOk

`func (o *GetApiStatus200Response) GetLastSampleAtOk() (*time.Time, bool)`

GetLastSampleAtOk returns a tuple with the LastSampleAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSampleAt

`func (o *GetApiStatus200Response) SetLastSampleAt(v time.Time)`

SetLastSampleAt sets LastSampleAt field to given value.


### SetLastSampleAtNil

`func (o *GetApiStatus200Response) SetLastSampleAtNil(b bool)`

 SetLastSampleAtNil sets the value for LastSampleAt to be an explicit nil

### UnsetLastSampleAt
`func (o *GetApiStatus200Response) UnsetLastSampleAt()`

UnsetLastSampleAt ensures that no value is present for LastSampleAt, not even an explicit nil
### GetStale

`func (o *GetApiStatus200Response) GetStale() bool`

GetStale returns the Stale field if non-nil, zero value otherwise.

### GetStaleOk

`func (o *GetApiStatus200Response) GetStaleOk() (*bool, bool)`

GetStaleOk returns a tuple with the Stale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStale

`func (o *GetApiStatus200Response) SetStale(v bool)`

SetStale sets Stale field to given value.


### GetOverall

`func (o *GetApiStatus200Response) GetOverall() string`

GetOverall returns the Overall field if non-nil, zero value otherwise.

### GetOverallOk

`func (o *GetApiStatus200Response) GetOverallOk() (*string, bool)`

GetOverallOk returns a tuple with the Overall field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverall

`func (o *GetApiStatus200Response) SetOverall(v string)`

SetOverall sets Overall field to given value.


### GetComponents

`func (o *GetApiStatus200Response) GetComponents() []GetApiStatus200ResponseComponentsInner`

GetComponents returns the Components field if non-nil, zero value otherwise.

### GetComponentsOk

`func (o *GetApiStatus200Response) GetComponentsOk() (*[]GetApiStatus200ResponseComponentsInner, bool)`

GetComponentsOk returns a tuple with the Components field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponents

`func (o *GetApiStatus200Response) SetComponents(v []GetApiStatus200ResponseComponentsInner)`

SetComponents sets Components field to given value.


### GetIncidents

`func (o *GetApiStatus200Response) GetIncidents() []GetApiStatus200ResponseIncidentsInner`

GetIncidents returns the Incidents field if non-nil, zero value otherwise.

### GetIncidentsOk

`func (o *GetApiStatus200Response) GetIncidentsOk() (*[]GetApiStatus200ResponseIncidentsInner, bool)`

GetIncidentsOk returns a tuple with the Incidents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncidents

`func (o *GetApiStatus200Response) SetIncidents(v []GetApiStatus200ResponseIncidentsInner)`

SetIncidents sets Incidents field to given value.


### GetMaintenance

`func (o *GetApiStatus200Response) GetMaintenance() []GetApiStatus200ResponseMaintenanceInner`

GetMaintenance returns the Maintenance field if non-nil, zero value otherwise.

### GetMaintenanceOk

`func (o *GetApiStatus200Response) GetMaintenanceOk() (*[]GetApiStatus200ResponseMaintenanceInner, bool)`

GetMaintenanceOk returns a tuple with the Maintenance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaintenance

`func (o *GetApiStatus200Response) SetMaintenance(v []GetApiStatus200ResponseMaintenanceInner)`

SetMaintenance sets Maintenance field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


