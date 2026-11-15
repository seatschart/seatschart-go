# GetApiStatus200ResponseComponentsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Component: &#x60;api&#x60; (REST API readiness), &#x60;realtime&#x60; (Pub/Sub fan-out behind the SSE streams), &#x60;queue&#x60; (reservation worker queue), &#x60;database&#x60; (PostgreSQL), &#x60;cache&#x60; (KeyDB / Redis), &#x60;assets&#x60; (map asset CDN), &#x60;dashboard&#x60; (web app and dashboard). | 
**Status** | **string** | &#x60;not_monitored&#x60; when the component has no probe configured on this deployment. | 
**LatencyMs** | **NullableFloat32** | Latest probe round trip in ms (for &#x60;queue&#x60;: age of the oldest unprocessed job). | 
**Uptime** | **NullableFloat32** | Uptime over the sampled minutes of the 90 days; null without samples. | 
**Days** | [**[]GetApiStatus200ResponseComponentsInnerDaysInner**](GetApiStatus200ResponseComponentsInnerDaysInner.md) | Oldest first; the last entry is today (UTC). | 

## Methods

### NewGetApiStatus200ResponseComponentsInner

`func NewGetApiStatus200ResponseComponentsInner(id string, status string, latencyMs NullableFloat32, uptime NullableFloat32, days []GetApiStatus200ResponseComponentsInnerDaysInner, ) *GetApiStatus200ResponseComponentsInner`

NewGetApiStatus200ResponseComponentsInner instantiates a new GetApiStatus200ResponseComponentsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetApiStatus200ResponseComponentsInnerWithDefaults

`func NewGetApiStatus200ResponseComponentsInnerWithDefaults() *GetApiStatus200ResponseComponentsInner`

NewGetApiStatus200ResponseComponentsInnerWithDefaults instantiates a new GetApiStatus200ResponseComponentsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GetApiStatus200ResponseComponentsInner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetApiStatus200ResponseComponentsInner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetApiStatus200ResponseComponentsInner) SetId(v string)`

SetId sets Id field to given value.


### GetStatus

`func (o *GetApiStatus200ResponseComponentsInner) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetApiStatus200ResponseComponentsInner) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetApiStatus200ResponseComponentsInner) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetLatencyMs

`func (o *GetApiStatus200ResponseComponentsInner) GetLatencyMs() float32`

GetLatencyMs returns the LatencyMs field if non-nil, zero value otherwise.

### GetLatencyMsOk

`func (o *GetApiStatus200ResponseComponentsInner) GetLatencyMsOk() (*float32, bool)`

GetLatencyMsOk returns a tuple with the LatencyMs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatencyMs

`func (o *GetApiStatus200ResponseComponentsInner) SetLatencyMs(v float32)`

SetLatencyMs sets LatencyMs field to given value.


### SetLatencyMsNil

`func (o *GetApiStatus200ResponseComponentsInner) SetLatencyMsNil(b bool)`

 SetLatencyMsNil sets the value for LatencyMs to be an explicit nil

### UnsetLatencyMs
`func (o *GetApiStatus200ResponseComponentsInner) UnsetLatencyMs()`

UnsetLatencyMs ensures that no value is present for LatencyMs, not even an explicit nil
### GetUptime

`func (o *GetApiStatus200ResponseComponentsInner) GetUptime() float32`

GetUptime returns the Uptime field if non-nil, zero value otherwise.

### GetUptimeOk

`func (o *GetApiStatus200ResponseComponentsInner) GetUptimeOk() (*float32, bool)`

GetUptimeOk returns a tuple with the Uptime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUptime

`func (o *GetApiStatus200ResponseComponentsInner) SetUptime(v float32)`

SetUptime sets Uptime field to given value.


### SetUptimeNil

`func (o *GetApiStatus200ResponseComponentsInner) SetUptimeNil(b bool)`

 SetUptimeNil sets the value for Uptime to be an explicit nil

### UnsetUptime
`func (o *GetApiStatus200ResponseComponentsInner) UnsetUptime()`

UnsetUptime ensures that no value is present for Uptime, not even an explicit nil
### GetDays

`func (o *GetApiStatus200ResponseComponentsInner) GetDays() []GetApiStatus200ResponseComponentsInnerDaysInner`

GetDays returns the Days field if non-nil, zero value otherwise.

### GetDaysOk

`func (o *GetApiStatus200ResponseComponentsInner) GetDaysOk() (*[]GetApiStatus200ResponseComponentsInnerDaysInner, bool)`

GetDaysOk returns a tuple with the Days field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDays

`func (o *GetApiStatus200ResponseComponentsInner) SetDays(v []GetApiStatus200ResponseComponentsInnerDaysInner)`

SetDays sets Days field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


