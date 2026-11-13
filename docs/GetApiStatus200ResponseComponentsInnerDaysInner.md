# GetApiStatus200ResponseComponentsInnerDaysInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Date** | **string** | UTC day (&#x60;YYYY-MM-DD&#x60;). | 
**Uptime** | **NullableFloat32** | Percent of the sampled minutes that were up (degraded counts as up), rounded down to 2 decimals; null when the day has no sample (\&quot;no data\&quot;, never \&quot;up\&quot;). | 
**Coverage** | **float32** | Share of the day (so far, for today) with a sample: 0–1. Below 1 the sampler was not running part of the day. | 
**DegradedMinutes** | **int32** |  | 
**DownMinutes** | **int32** |  | 
**Incidents** | **int32** | Incidents (automatic and manual) that overlap the day. | 

## Methods

### NewGetApiStatus200ResponseComponentsInnerDaysInner

`func NewGetApiStatus200ResponseComponentsInnerDaysInner(date string, uptime NullableFloat32, coverage float32, degradedMinutes int32, downMinutes int32, incidents int32, ) *GetApiStatus200ResponseComponentsInnerDaysInner`

NewGetApiStatus200ResponseComponentsInnerDaysInner instantiates a new GetApiStatus200ResponseComponentsInnerDaysInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetApiStatus200ResponseComponentsInnerDaysInnerWithDefaults

`func NewGetApiStatus200ResponseComponentsInnerDaysInnerWithDefaults() *GetApiStatus200ResponseComponentsInnerDaysInner`

NewGetApiStatus200ResponseComponentsInnerDaysInnerWithDefaults instantiates a new GetApiStatus200ResponseComponentsInnerDaysInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDate

`func (o *GetApiStatus200ResponseComponentsInnerDaysInner) GetDate() string`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *GetApiStatus200ResponseComponentsInnerDaysInner) GetDateOk() (*string, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *GetApiStatus200ResponseComponentsInnerDaysInner) SetDate(v string)`

SetDate sets Date field to given value.


### GetUptime

`func (o *GetApiStatus200ResponseComponentsInnerDaysInner) GetUptime() float32`

GetUptime returns the Uptime field if non-nil, zero value otherwise.

### GetUptimeOk

`func (o *GetApiStatus200ResponseComponentsInnerDaysInner) GetUptimeOk() (*float32, bool)`

GetUptimeOk returns a tuple with the Uptime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUptime

`func (o *GetApiStatus200ResponseComponentsInnerDaysInner) SetUptime(v float32)`

SetUptime sets Uptime field to given value.


### SetUptimeNil

`func (o *GetApiStatus200ResponseComponentsInnerDaysInner) SetUptimeNil(b bool)`

 SetUptimeNil sets the value for Uptime to be an explicit nil

### UnsetUptime
`func (o *GetApiStatus200ResponseComponentsInnerDaysInner) UnsetUptime()`

UnsetUptime ensures that no value is present for Uptime, not even an explicit nil
### GetCoverage

`func (o *GetApiStatus200ResponseComponentsInnerDaysInner) GetCoverage() float32`

GetCoverage returns the Coverage field if non-nil, zero value otherwise.

### GetCoverageOk

`func (o *GetApiStatus200ResponseComponentsInnerDaysInner) GetCoverageOk() (*float32, bool)`

GetCoverageOk returns a tuple with the Coverage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoverage

`func (o *GetApiStatus200ResponseComponentsInnerDaysInner) SetCoverage(v float32)`

SetCoverage sets Coverage field to given value.


### GetDegradedMinutes

`func (o *GetApiStatus200ResponseComponentsInnerDaysInner) GetDegradedMinutes() int32`

GetDegradedMinutes returns the DegradedMinutes field if non-nil, zero value otherwise.

### GetDegradedMinutesOk

`func (o *GetApiStatus200ResponseComponentsInnerDaysInner) GetDegradedMinutesOk() (*int32, bool)`

GetDegradedMinutesOk returns a tuple with the DegradedMinutes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDegradedMinutes

`func (o *GetApiStatus200ResponseComponentsInnerDaysInner) SetDegradedMinutes(v int32)`

SetDegradedMinutes sets DegradedMinutes field to given value.


### GetDownMinutes

`func (o *GetApiStatus200ResponseComponentsInnerDaysInner) GetDownMinutes() int32`

GetDownMinutes returns the DownMinutes field if non-nil, zero value otherwise.

### GetDownMinutesOk

`func (o *GetApiStatus200ResponseComponentsInnerDaysInner) GetDownMinutesOk() (*int32, bool)`

GetDownMinutesOk returns a tuple with the DownMinutes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownMinutes

`func (o *GetApiStatus200ResponseComponentsInnerDaysInner) SetDownMinutes(v int32)`

SetDownMinutes sets DownMinutes field to given value.


### GetIncidents

`func (o *GetApiStatus200ResponseComponentsInnerDaysInner) GetIncidents() int32`

GetIncidents returns the Incidents field if non-nil, zero value otherwise.

### GetIncidentsOk

`func (o *GetApiStatus200ResponseComponentsInnerDaysInner) GetIncidentsOk() (*int32, bool)`

GetIncidentsOk returns a tuple with the Incidents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncidents

`func (o *GetApiStatus200ResponseComponentsInnerDaysInner) SetIncidents(v int32)`

SetIncidents sets Incidents field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


