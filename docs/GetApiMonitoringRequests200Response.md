# GetApiMonitoringRequests200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Range** | **string** |  | 
**BucketSeconds** | **int32** | 60 for 10m/1h, 900 for 24h, 3600 for 7d. | 
**From** | **time.Time** |  | 
**To** | **time.Time** |  | 
**Plan** | **string** | Plan whose limits apply (a test workspace uses its production workspace plan). | 
**IsTest** | **bool** |  | 
**WindowSeconds** | **int32** | Rate limit window (RATE_LIMIT_TENANT_WINDOW_SEC). | 
**Lanes** | [**GetApiMonitoringRequests200ResponseLanes**](GetApiMonitoringRequests200ResponseLanes.md) |  | 
**Points** | [**[]GetApiMonitoringRequests200ResponsePointsInner**](GetApiMonitoringRequests200ResponsePointsInner.md) |  | 

## Methods

### NewGetApiMonitoringRequests200Response

`func NewGetApiMonitoringRequests200Response(range_ string, bucketSeconds int32, from time.Time, to time.Time, plan string, isTest bool, windowSeconds int32, lanes GetApiMonitoringRequests200ResponseLanes, points []GetApiMonitoringRequests200ResponsePointsInner, ) *GetApiMonitoringRequests200Response`

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


