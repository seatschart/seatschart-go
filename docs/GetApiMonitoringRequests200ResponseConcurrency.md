# GetApiMonitoringRequests200ResponseConcurrency

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Current** | **int32** | Workspace requests in flight now on the API instance that answered (this request excluded). | 
**PeakLastHour** | **int32** | Highest concurrent requests in any minute of the last hour. | 
**PeakLastHourAt** | **NullableTime** | Start of the minute (hour for the day) of the peak. | 
**PeakLastDay** | **int32** | Highest concurrent requests in any hour of the last 24 hours. | 
**PeakLastDayAt** | **NullableTime** | Start of the minute (hour for the day) of the peak. | 
**Instances** | **int32** | API instances that reported in-flight peaks in the last 24 hours (peaks are summed across them). | 
**ApiKeys** | [**[]GetApiMonitoringRequests200ResponseConcurrencyApiKeysInner**](GetApiMonitoringRequests200ResponseConcurrencyApiKeysInner.md) | Per API key, busiest first (keys without traffic in the last 24 hours are omitted). | 

## Methods

### NewGetApiMonitoringRequests200ResponseConcurrency

`func NewGetApiMonitoringRequests200ResponseConcurrency(current int32, peakLastHour int32, peakLastHourAt NullableTime, peakLastDay int32, peakLastDayAt NullableTime, instances int32, apiKeys []GetApiMonitoringRequests200ResponseConcurrencyApiKeysInner, ) *GetApiMonitoringRequests200ResponseConcurrency`

NewGetApiMonitoringRequests200ResponseConcurrency instantiates a new GetApiMonitoringRequests200ResponseConcurrency object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetApiMonitoringRequests200ResponseConcurrencyWithDefaults

`func NewGetApiMonitoringRequests200ResponseConcurrencyWithDefaults() *GetApiMonitoringRequests200ResponseConcurrency`

NewGetApiMonitoringRequests200ResponseConcurrencyWithDefaults instantiates a new GetApiMonitoringRequests200ResponseConcurrency object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCurrent

`func (o *GetApiMonitoringRequests200ResponseConcurrency) GetCurrent() int32`

GetCurrent returns the Current field if non-nil, zero value otherwise.

### GetCurrentOk

`func (o *GetApiMonitoringRequests200ResponseConcurrency) GetCurrentOk() (*int32, bool)`

GetCurrentOk returns a tuple with the Current field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrent

`func (o *GetApiMonitoringRequests200ResponseConcurrency) SetCurrent(v int32)`

SetCurrent sets Current field to given value.


### GetPeakLastHour

`func (o *GetApiMonitoringRequests200ResponseConcurrency) GetPeakLastHour() int32`

GetPeakLastHour returns the PeakLastHour field if non-nil, zero value otherwise.

### GetPeakLastHourOk

`func (o *GetApiMonitoringRequests200ResponseConcurrency) GetPeakLastHourOk() (*int32, bool)`

GetPeakLastHourOk returns a tuple with the PeakLastHour field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeakLastHour

`func (o *GetApiMonitoringRequests200ResponseConcurrency) SetPeakLastHour(v int32)`

SetPeakLastHour sets PeakLastHour field to given value.


### GetPeakLastHourAt

`func (o *GetApiMonitoringRequests200ResponseConcurrency) GetPeakLastHourAt() time.Time`

GetPeakLastHourAt returns the PeakLastHourAt field if non-nil, zero value otherwise.

### GetPeakLastHourAtOk

`func (o *GetApiMonitoringRequests200ResponseConcurrency) GetPeakLastHourAtOk() (*time.Time, bool)`

GetPeakLastHourAtOk returns a tuple with the PeakLastHourAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeakLastHourAt

`func (o *GetApiMonitoringRequests200ResponseConcurrency) SetPeakLastHourAt(v time.Time)`

SetPeakLastHourAt sets PeakLastHourAt field to given value.


### SetPeakLastHourAtNil

`func (o *GetApiMonitoringRequests200ResponseConcurrency) SetPeakLastHourAtNil(b bool)`

 SetPeakLastHourAtNil sets the value for PeakLastHourAt to be an explicit nil

### UnsetPeakLastHourAt
`func (o *GetApiMonitoringRequests200ResponseConcurrency) UnsetPeakLastHourAt()`

UnsetPeakLastHourAt ensures that no value is present for PeakLastHourAt, not even an explicit nil
### GetPeakLastDay

`func (o *GetApiMonitoringRequests200ResponseConcurrency) GetPeakLastDay() int32`

GetPeakLastDay returns the PeakLastDay field if non-nil, zero value otherwise.

### GetPeakLastDayOk

`func (o *GetApiMonitoringRequests200ResponseConcurrency) GetPeakLastDayOk() (*int32, bool)`

GetPeakLastDayOk returns a tuple with the PeakLastDay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeakLastDay

`func (o *GetApiMonitoringRequests200ResponseConcurrency) SetPeakLastDay(v int32)`

SetPeakLastDay sets PeakLastDay field to given value.


### GetPeakLastDayAt

`func (o *GetApiMonitoringRequests200ResponseConcurrency) GetPeakLastDayAt() time.Time`

GetPeakLastDayAt returns the PeakLastDayAt field if non-nil, zero value otherwise.

### GetPeakLastDayAtOk

`func (o *GetApiMonitoringRequests200ResponseConcurrency) GetPeakLastDayAtOk() (*time.Time, bool)`

GetPeakLastDayAtOk returns a tuple with the PeakLastDayAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeakLastDayAt

`func (o *GetApiMonitoringRequests200ResponseConcurrency) SetPeakLastDayAt(v time.Time)`

SetPeakLastDayAt sets PeakLastDayAt field to given value.


### SetPeakLastDayAtNil

`func (o *GetApiMonitoringRequests200ResponseConcurrency) SetPeakLastDayAtNil(b bool)`

 SetPeakLastDayAtNil sets the value for PeakLastDayAt to be an explicit nil

### UnsetPeakLastDayAt
`func (o *GetApiMonitoringRequests200ResponseConcurrency) UnsetPeakLastDayAt()`

UnsetPeakLastDayAt ensures that no value is present for PeakLastDayAt, not even an explicit nil
### GetInstances

`func (o *GetApiMonitoringRequests200ResponseConcurrency) GetInstances() int32`

GetInstances returns the Instances field if non-nil, zero value otherwise.

### GetInstancesOk

`func (o *GetApiMonitoringRequests200ResponseConcurrency) GetInstancesOk() (*int32, bool)`

GetInstancesOk returns a tuple with the Instances field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstances

`func (o *GetApiMonitoringRequests200ResponseConcurrency) SetInstances(v int32)`

SetInstances sets Instances field to given value.


### GetApiKeys

`func (o *GetApiMonitoringRequests200ResponseConcurrency) GetApiKeys() []GetApiMonitoringRequests200ResponseConcurrencyApiKeysInner`

GetApiKeys returns the ApiKeys field if non-nil, zero value otherwise.

### GetApiKeysOk

`func (o *GetApiMonitoringRequests200ResponseConcurrency) GetApiKeysOk() (*[]GetApiMonitoringRequests200ResponseConcurrencyApiKeysInner, bool)`

GetApiKeysOk returns a tuple with the ApiKeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiKeys

`func (o *GetApiMonitoringRequests200ResponseConcurrency) SetApiKeys(v []GetApiMonitoringRequests200ResponseConcurrencyApiKeysInner)`

SetApiKeys sets ApiKeys field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


