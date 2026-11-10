# GetApiMonitoringRequests200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Range** | **string** |  | 
**BucketSeconds** | **int32** | 60 for 10m/1h, 900 for 24h, 3600 for 7d, 86400 for 30d. | 
**From** | **time.Time** | Window start (first bucket). | 
**To** | **time.Time** | Window end: now for the live window, else the (exclusive) &#x60;before&#x60;. | 
**Before** | **NullableTime** | Aligned cursor of this window; null for the live window. | 
**PreviousBefore** | **NullableTime** | &#x60;before&#x60; of the previous (older) window; null when nothing older is kept. | 
**NextBefore** | **NullableTime** | &#x60;before&#x60; of the next (newer) window; null when the next window is the live one (omit &#x60;before&#x60;) or this is the live window. | 
**RetentionStart** | **time.Time** | Oldest instant kept at this window size (24 hours of minutes, 31 days of hours and days). | 
**Plan** | **string** | Plan whose limits apply (a test workspace uses its production workspace plan). | 
**IsTest** | **bool** |  | 
**WindowSeconds** | **int32** | Rate limit window (RATE_LIMIT_TENANT_WINDOW_SEC). | 
**Lanes** | [**GetApiMonitoringRequests200ResponseLanes**](GetApiMonitoringRequests200ResponseLanes.md) |  | 
**Concurrency** | [**GetApiMonitoringRequests200ResponseConcurrency**](GetApiMonitoringRequests200ResponseConcurrency.md) |  | 
**Points** | [**[]GetApiMonitoringRequests200ResponsePointsInner**](GetApiMonitoringRequests200ResponsePointsInner.md) |  | 

## Methods

### NewGetApiMonitoringRequests200Response

`func NewGetApiMonitoringRequests200Response(range_ string, bucketSeconds int32, from time.Time, to time.Time, before NullableTime, previousBefore NullableTime, nextBefore NullableTime, retentionStart time.Time, plan string, isTest bool, windowSeconds int32, lanes GetApiMonitoringRequests200ResponseLanes, concurrency GetApiMonitoringRequests200ResponseConcurrency, points []GetApiMonitoringRequests200ResponsePointsInner, ) *GetApiMonitoringRequests200Response`

NewGetApiMonitoringRequests200Response instantiates a new GetApiMonitoringRequests200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetApiMonitoringRequests200ResponseWithDefaults

`func NewGetApiMonitoringRequests200ResponseWithDefaults() *GetApiMonitoringRequests200Response`

NewGetApiMonitoringRequests200ResponseWithDefaults instantiates a new GetApiMonitoringRequests200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRange

`func (o *GetApiMonitoringRequests200Response) GetRange() string`

GetRange returns the Range field if non-nil, zero value otherwise.

### GetRangeOk

`func (o *GetApiMonitoringRequests200Response) GetRangeOk() (*string, bool)`

GetRangeOk returns a tuple with the Range field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRange

`func (o *GetApiMonitoringRequests200Response) SetRange(v string)`

SetRange sets Range field to given value.


### GetBucketSeconds

`func (o *GetApiMonitoringRequests200Response) GetBucketSeconds() int32`

GetBucketSeconds returns the BucketSeconds field if non-nil, zero value otherwise.

### GetBucketSecondsOk

`func (o *GetApiMonitoringRequests200Response) GetBucketSecondsOk() (*int32, bool)`

GetBucketSecondsOk returns a tuple with the BucketSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBucketSeconds

`func (o *GetApiMonitoringRequests200Response) SetBucketSeconds(v int32)`

SetBucketSeconds sets BucketSeconds field to given value.


### GetFrom

`func (o *GetApiMonitoringRequests200Response) GetFrom() time.Time`

GetFrom returns the From field if non-nil, zero value otherwise.

### GetFromOk

`func (o *GetApiMonitoringRequests200Response) GetFromOk() (*time.Time, bool)`

GetFromOk returns a tuple with the From field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrom

`func (o *GetApiMonitoringRequests200Response) SetFrom(v time.Time)`

SetFrom sets From field to given value.


### GetTo

`func (o *GetApiMonitoringRequests200Response) GetTo() time.Time`

GetTo returns the To field if non-nil, zero value otherwise.

### GetToOk

`func (o *GetApiMonitoringRequests200Response) GetToOk() (*time.Time, bool)`

GetToOk returns a tuple with the To field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTo

`func (o *GetApiMonitoringRequests200Response) SetTo(v time.Time)`

SetTo sets To field to given value.


### GetBefore

`func (o *GetApiMonitoringRequests200Response) GetBefore() time.Time`

GetBefore returns the Before field if non-nil, zero value otherwise.

### GetBeforeOk

`func (o *GetApiMonitoringRequests200Response) GetBeforeOk() (*time.Time, bool)`

GetBeforeOk returns a tuple with the Before field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBefore

`func (o *GetApiMonitoringRequests200Response) SetBefore(v time.Time)`

SetBefore sets Before field to given value.


### SetBeforeNil

`func (o *GetApiMonitoringRequests200Response) SetBeforeNil(b bool)`

 SetBeforeNil sets the value for Before to be an explicit nil

### UnsetBefore
`func (o *GetApiMonitoringRequests200Response) UnsetBefore()`

UnsetBefore ensures that no value is present for Before, not even an explicit nil
### GetPreviousBefore

`func (o *GetApiMonitoringRequests200Response) GetPreviousBefore() time.Time`

GetPreviousBefore returns the PreviousBefore field if non-nil, zero value otherwise.

### GetPreviousBeforeOk

`func (o *GetApiMonitoringRequests200Response) GetPreviousBeforeOk() (*time.Time, bool)`

GetPreviousBeforeOk returns a tuple with the PreviousBefore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreviousBefore

`func (o *GetApiMonitoringRequests200Response) SetPreviousBefore(v time.Time)`

SetPreviousBefore sets PreviousBefore field to given value.


### SetPreviousBeforeNil

`func (o *GetApiMonitoringRequests200Response) SetPreviousBeforeNil(b bool)`

 SetPreviousBeforeNil sets the value for PreviousBefore to be an explicit nil

### UnsetPreviousBefore
`func (o *GetApiMonitoringRequests200Response) UnsetPreviousBefore()`

UnsetPreviousBefore ensures that no value is present for PreviousBefore, not even an explicit nil
### GetNextBefore

`func (o *GetApiMonitoringRequests200Response) GetNextBefore() time.Time`

GetNextBefore returns the NextBefore field if non-nil, zero value otherwise.

### GetNextBeforeOk

`func (o *GetApiMonitoringRequests200Response) GetNextBeforeOk() (*time.Time, bool)`

GetNextBeforeOk returns a tuple with the NextBefore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextBefore

`func (o *GetApiMonitoringRequests200Response) SetNextBefore(v time.Time)`

SetNextBefore sets NextBefore field to given value.


### SetNextBeforeNil

`func (o *GetApiMonitoringRequests200Response) SetNextBeforeNil(b bool)`

 SetNextBeforeNil sets the value for NextBefore to be an explicit nil

### UnsetNextBefore
`func (o *GetApiMonitoringRequests200Response) UnsetNextBefore()`

UnsetNextBefore ensures that no value is present for NextBefore, not even an explicit nil
### GetRetentionStart

`func (o *GetApiMonitoringRequests200Response) GetRetentionStart() time.Time`

GetRetentionStart returns the RetentionStart field if non-nil, zero value otherwise.

### GetRetentionStartOk

`func (o *GetApiMonitoringRequests200Response) GetRetentionStartOk() (*time.Time, bool)`

GetRetentionStartOk returns a tuple with the RetentionStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetentionStart

`func (o *GetApiMonitoringRequests200Response) SetRetentionStart(v time.Time)`

SetRetentionStart sets RetentionStart field to given value.


### GetPlan

`func (o *GetApiMonitoringRequests200Response) GetPlan() string`

GetPlan returns the Plan field if non-nil, zero value otherwise.

### GetPlanOk

`func (o *GetApiMonitoringRequests200Response) GetPlanOk() (*string, bool)`

GetPlanOk returns a tuple with the Plan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlan

`func (o *GetApiMonitoringRequests200Response) SetPlan(v string)`

SetPlan sets Plan field to given value.


### GetIsTest

`func (o *GetApiMonitoringRequests200Response) GetIsTest() bool`

GetIsTest returns the IsTest field if non-nil, zero value otherwise.

### GetIsTestOk

`func (o *GetApiMonitoringRequests200Response) GetIsTestOk() (*bool, bool)`

GetIsTestOk returns a tuple with the IsTest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsTest

`func (o *GetApiMonitoringRequests200Response) SetIsTest(v bool)`

SetIsTest sets IsTest field to given value.


### GetWindowSeconds

`func (o *GetApiMonitoringRequests200Response) GetWindowSeconds() int32`

GetWindowSeconds returns the WindowSeconds field if non-nil, zero value otherwise.

### GetWindowSecondsOk

`func (o *GetApiMonitoringRequests200Response) GetWindowSecondsOk() (*int32, bool)`

GetWindowSecondsOk returns a tuple with the WindowSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWindowSeconds

`func (o *GetApiMonitoringRequests200Response) SetWindowSeconds(v int32)`

SetWindowSeconds sets WindowSeconds field to given value.


### GetLanes

`func (o *GetApiMonitoringRequests200Response) GetLanes() GetApiMonitoringRequests200ResponseLanes`

GetLanes returns the Lanes field if non-nil, zero value otherwise.

### GetLanesOk

`func (o *GetApiMonitoringRequests200Response) GetLanesOk() (*GetApiMonitoringRequests200ResponseLanes, bool)`

GetLanesOk returns a tuple with the Lanes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanes

`func (o *GetApiMonitoringRequests200Response) SetLanes(v GetApiMonitoringRequests200ResponseLanes)`

SetLanes sets Lanes field to given value.


### GetConcurrency

`func (o *GetApiMonitoringRequests200Response) GetConcurrency() GetApiMonitoringRequests200ResponseConcurrency`

GetConcurrency returns the Concurrency field if non-nil, zero value otherwise.

### GetConcurrencyOk

`func (o *GetApiMonitoringRequests200Response) GetConcurrencyOk() (*GetApiMonitoringRequests200ResponseConcurrency, bool)`

GetConcurrencyOk returns a tuple with the Concurrency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConcurrency

`func (o *GetApiMonitoringRequests200Response) SetConcurrency(v GetApiMonitoringRequests200ResponseConcurrency)`

SetConcurrency sets Concurrency field to given value.


### GetPoints

`func (o *GetApiMonitoringRequests200Response) GetPoints() []GetApiMonitoringRequests200ResponsePointsInner`

GetPoints returns the Points field if non-nil, zero value otherwise.

### GetPointsOk

`func (o *GetApiMonitoringRequests200Response) GetPointsOk() (*[]GetApiMonitoringRequests200ResponsePointsInner, bool)`

GetPointsOk returns a tuple with the Points field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoints

`func (o *GetApiMonitoringRequests200Response) SetPoints(v []GetApiMonitoringRequests200ResponsePointsInner)`

SetPoints sets Points field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


