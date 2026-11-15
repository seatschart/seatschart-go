# PostApiGaReleaseRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EventId** | **string** |  | 
**AreaId** | **string** |  | 
**UserId** | **string** |  | 
**HoldToken** | Pointer to **string** |  | [optional] 

## Methods

### NewPostApiGaReleaseRequest

`func NewPostApiGaReleaseRequest(eventId string, areaId string, userId string, ) *PostApiGaReleaseRequest`

NewPostApiGaReleaseRequest instantiates a new PostApiGaReleaseRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostApiGaReleaseRequestWithDefaults

`func NewPostApiGaReleaseRequestWithDefaults() *PostApiGaReleaseRequest`

NewPostApiGaReleaseRequestWithDefaults instantiates a new PostApiGaReleaseRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEventId

`func (o *PostApiGaReleaseRequest) GetEventId() string`

GetEventId returns the EventId field if non-nil, zero value otherwise.

### GetEventIdOk

`func (o *PostApiGaReleaseRequest) GetEventIdOk() (*string, bool)`

GetEventIdOk returns a tuple with the EventId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventId

`func (o *PostApiGaReleaseRequest) SetEventId(v string)`

SetEventId sets EventId field to given value.


### GetAreaId

`func (o *PostApiGaReleaseRequest) GetAreaId() string`

GetAreaId returns the AreaId field if non-nil, zero value otherwise.

### GetAreaIdOk

`func (o *PostApiGaReleaseRequest) GetAreaIdOk() (*string, bool)`

GetAreaIdOk returns a tuple with the AreaId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAreaId

`func (o *PostApiGaReleaseRequest) SetAreaId(v string)`

SetAreaId sets AreaId field to given value.


### GetUserId

`func (o *PostApiGaReleaseRequest) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *PostApiGaReleaseRequest) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *PostApiGaReleaseRequest) SetUserId(v string)`

SetUserId sets UserId field to given value.


### GetHoldToken

`func (o *PostApiGaReleaseRequest) GetHoldToken() string`

GetHoldToken returns the HoldToken field if non-nil, zero value otherwise.

### GetHoldTokenOk

`func (o *PostApiGaReleaseRequest) GetHoldTokenOk() (*string, bool)`

GetHoldTokenOk returns a tuple with the HoldToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHoldToken

`func (o *PostApiGaReleaseRequest) SetHoldToken(v string)`

SetHoldToken sets HoldToken field to given value.

### HasHoldToken

`func (o *PostApiGaReleaseRequest) HasHoldToken() bool`

HasHoldToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


