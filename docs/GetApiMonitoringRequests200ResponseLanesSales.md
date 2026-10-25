# GetApiMonitoringRequests200ResponseLanesSales

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LimitPerMinute** | **int32** | Requests per minute this lane admits on the workspace plan. | 
**Peak** | **int32** | Highest requests in one minute over the range. | 
**Rejected** | **int32** | Requests refused with 429 over the range. | 

## Methods

### NewGetApiMonitoringRequests200ResponseLanesSales

`func NewGetApiMonitoringRequests200ResponseLanesSales(limitPerMinute int32, peak int32, rejected int32, ) *GetApiMonitoringRequests200ResponseLanesSales`

NewGetApiMonitoringRequests200ResponseLanesSales instantiates a new GetApiMonitoringRequests200ResponseLanesSales object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetApiMonitoringRequests200ResponseLanesSalesWithDefaults

`func NewGetApiMonitoringRequests200ResponseLanesSalesWithDefaults() *GetApiMonitoringRequests200ResponseLanesSales`

NewGetApiMonitoringRequests200ResponseLanesSalesWithDefaults instantiates a new GetApiMonitoringRequests200ResponseLanesSales object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLimitPerMinute

`func (o *GetApiMonitoringRequests200ResponseLanesSales) GetLimitPerMinute() int32`

GetLimitPerMinute returns the LimitPerMinute field if non-nil, zero value otherwise.

### GetLimitPerMinuteOk

`func (o *GetApiMonitoringRequests200ResponseLanesSales) GetLimitPerMinuteOk() (*int32, bool)`

GetLimitPerMinuteOk returns a tuple with the LimitPerMinute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimitPerMinute

`func (o *GetApiMonitoringRequests200ResponseLanesSales) SetLimitPerMinute(v int32)`

SetLimitPerMinute sets LimitPerMinute field to given value.


### GetPeak

`func (o *GetApiMonitoringRequests200ResponseLanesSales) GetPeak() int32`

GetPeak returns the Peak field if non-nil, zero value otherwise.

### GetPeakOk

`func (o *GetApiMonitoringRequests200ResponseLanesSales) GetPeakOk() (*int32, bool)`

GetPeakOk returns a tuple with the Peak field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeak

`func (o *GetApiMonitoringRequests200ResponseLanesSales) SetPeak(v int32)`

SetPeak sets Peak field to given value.


### GetRejected

`func (o *GetApiMonitoringRequests200ResponseLanesSales) GetRejected() int32`

GetRejected returns the Rejected field if non-nil, zero value otherwise.

### GetRejectedOk

`func (o *GetApiMonitoringRequests200ResponseLanesSales) GetRejectedOk() (*int32, bool)`

GetRejectedOk returns a tuple with the Rejected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRejected

`func (o *GetApiMonitoringRequests200ResponseLanesSales) SetRejected(v int32)`

SetRejected sets Rejected field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


