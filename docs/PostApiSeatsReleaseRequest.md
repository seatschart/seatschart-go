# PostApiSeatsReleaseRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EventId** | **string** |  | 
**SeatIds** | **[]string** |  | 
**UserId** | **string** |  | 
**HoldToken** | Pointer to **string** |  | [optional] 

## Methods

### NewPostApiSeatsReleaseRequest

`func NewPostApiSeatsReleaseRequest(eventId string, seatIds []string, userId string, ) *PostApiSeatsReleaseRequest`

NewPostApiSeatsReleaseRequest instantiates a new PostApiSeatsReleaseRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostApiSeatsReleaseRequestWithDefaults

`func NewPostApiSeatsReleaseRequestWithDefaults() *PostApiSeatsReleaseRequest`

NewPostApiSeatsReleaseRequestWithDefaults instantiates a new PostApiSeatsReleaseRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEventId

`func (o *PostApiSeatsReleaseRequest) GetEventId() string`

GetEventId returns the EventId field if non-nil, zero value otherwise.

### GetEventIdOk

`func (o *PostApiSeatsReleaseRequest) GetEventIdOk() (*string, bool)`

GetEventIdOk returns a tuple with the EventId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventId

`func (o *PostApiSeatsReleaseRequest) SetEventId(v string)`

SetEventId sets EventId field to given value.


### GetSeatIds

`func (o *PostApiSeatsReleaseRequest) GetSeatIds() []string`

GetSeatIds returns the SeatIds field if non-nil, zero value otherwise.

### GetSeatIdsOk

`func (o *PostApiSeatsReleaseRequest) GetSeatIdsOk() (*[]string, bool)`

GetSeatIdsOk returns a tuple with the SeatIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeatIds

`func (o *PostApiSeatsReleaseRequest) SetSeatIds(v []string)`

SetSeatIds sets SeatIds field to given value.


### GetUserId

`func (o *PostApiSeatsReleaseRequest) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *PostApiSeatsReleaseRequest) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *PostApiSeatsReleaseRequest) SetUserId(v string)`

SetUserId sets UserId field to given value.


### GetHoldToken

`func (o *PostApiSeatsReleaseRequest) GetHoldToken() string`

GetHoldToken returns the HoldToken field if non-nil, zero value otherwise.

### GetHoldTokenOk

`func (o *PostApiSeatsReleaseRequest) GetHoldTokenOk() (*string, bool)`

GetHoldTokenOk returns a tuple with the HoldToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHoldToken

`func (o *PostApiSeatsReleaseRequest) SetHoldToken(v string)`

SetHoldToken sets HoldToken field to given value.

### HasHoldToken

`func (o *PostApiSeatsReleaseRequest) HasHoldToken() bool`

HasHoldToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


