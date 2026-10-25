# GetApiMonitoringRequests200ResponseConcurrencyApiKeysInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | API key id. | 
**Name** | **NullableString** | Key name (null when the key was deleted). | 
**Prefix** | **NullableString** | Non-secret key prefix. | 
**Current** | **int32** | Requests of this key in flight now on the API instance that answered. | 
**PeakLastHour** | **int32** |  | 
**PeakLastDay** | **int32** |  | 

## Methods

### NewGetApiMonitoringRequests200ResponseConcurrencyApiKeysInner

`func NewGetApiMonitoringRequests200ResponseConcurrencyApiKeysInner(id string, name NullableString, prefix NullableString, current int32, peakLastHour int32, peakLastDay int32, ) *GetApiMonitoringRequests200ResponseConcurrencyApiKeysInner`

NewGetApiMonitoringRequests200ResponseConcurrencyApiKeysInner instantiates a new GetApiMonitoringRequests200ResponseConcurrencyApiKeysInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetApiMonitoringRequests200ResponseConcurrencyApiKeysInnerWithDefaults

`func NewGetApiMonitoringRequests200ResponseConcurrencyApiKeysInnerWithDefaults() *GetApiMonitoringRequests200ResponseConcurrencyApiKeysInner`

NewGetApiMonitoringRequests200ResponseConcurrencyApiKeysInnerWithDefaults instantiates a new GetApiMonitoringRequests200ResponseConcurrencyApiKeysInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GetApiMonitoringRequests200ResponseConcurrencyApiKeysInner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetApiMonitoringRequests200ResponseConcurrencyApiKeysInner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetApiMonitoringRequests200ResponseConcurrencyApiKeysInner) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *GetApiMonitoringRequests200ResponseConcurrencyApiKeysInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetApiMonitoringRequests200ResponseConcurrencyApiKeysInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetApiMonitoringRequests200ResponseConcurrencyApiKeysInner) SetName(v string)`

SetName sets Name field to given value.


### SetNameNil

`func (o *GetApiMonitoringRequests200ResponseConcurrencyApiKeysInner) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *GetApiMonitoringRequests200ResponseConcurrencyApiKeysInner) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetPrefix

`func (o *GetApiMonitoringRequests200ResponseConcurrencyApiKeysInner) GetPrefix() string`

GetPrefix returns the Prefix field if non-nil, zero value otherwise.

### GetPrefixOk

`func (o *GetApiMonitoringRequests200ResponseConcurrencyApiKeysInner) GetPrefixOk() (*string, bool)`

GetPrefixOk returns a tuple with the Prefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrefix

`func (o *GetApiMonitoringRequests200ResponseConcurrencyApiKeysInner) SetPrefix(v string)`

SetPrefix sets Prefix field to given value.


### SetPrefixNil

`func (o *GetApiMonitoringRequests200ResponseConcurrencyApiKeysInner) SetPrefixNil(b bool)`

 SetPrefixNil sets the value for Prefix to be an explicit nil

### UnsetPrefix
`func (o *GetApiMonitoringRequests200ResponseConcurrencyApiKeysInner) UnsetPrefix()`

UnsetPrefix ensures that no value is present for Prefix, not even an explicit nil
### GetCurrent

`func (o *GetApiMonitoringRequests200ResponseConcurrencyApiKeysInner) GetCurrent() int32`

GetCurrent returns the Current field if non-nil, zero value otherwise.

### GetCurrentOk

`func (o *GetApiMonitoringRequests200ResponseConcurrencyApiKeysInner) GetCurrentOk() (*int32, bool)`

GetCurrentOk returns a tuple with the Current field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrent

`func (o *GetApiMonitoringRequests200ResponseConcurrencyApiKeysInner) SetCurrent(v int32)`

SetCurrent sets Current field to given value.


### GetPeakLastHour

`func (o *GetApiMonitoringRequests200ResponseConcurrencyApiKeysInner) GetPeakLastHour() int32`

GetPeakLastHour returns the PeakLastHour field if non-nil, zero value otherwise.

### GetPeakLastHourOk

`func (o *GetApiMonitoringRequests200ResponseConcurrencyApiKeysInner) GetPeakLastHourOk() (*int32, bool)`

GetPeakLastHourOk returns a tuple with the PeakLastHour field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeakLastHour

`func (o *GetApiMonitoringRequests200ResponseConcurrencyApiKeysInner) SetPeakLastHour(v int32)`

SetPeakLastHour sets PeakLastHour field to given value.


### GetPeakLastDay

`func (o *GetApiMonitoringRequests200ResponseConcurrencyApiKeysInner) GetPeakLastDay() int32`

GetPeakLastDay returns the PeakLastDay field if non-nil, zero value otherwise.

### GetPeakLastDayOk

`func (o *GetApiMonitoringRequests200ResponseConcurrencyApiKeysInner) GetPeakLastDayOk() (*int32, bool)`

GetPeakLastDayOk returns a tuple with the PeakLastDay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeakLastDay

`func (o *GetApiMonitoringRequests200ResponseConcurrencyApiKeysInner) SetPeakLastDay(v int32)`

SetPeakLastDay sets PeakLastDay field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


