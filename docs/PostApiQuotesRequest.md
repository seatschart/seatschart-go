# PostApiQuotesRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EventId** | **string** |  | 
**UserId** | **string** |  | 
**HoldToken** | **string** |  | 
**SeatTicketTypes** | Pointer to **map[string]string** |  | [optional] 
**TicketTypeId** | Pointer to **string** |  | [optional] 

## Methods

### NewPostApiQuotesRequest

`func NewPostApiQuotesRequest(eventId string, userId string, holdToken string, ) *PostApiQuotesRequest`

NewPostApiQuotesRequest instantiates a new PostApiQuotesRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostApiQuotesRequestWithDefaults

`func NewPostApiQuotesRequestWithDefaults() *PostApiQuotesRequest`

NewPostApiQuotesRequestWithDefaults instantiates a new PostApiQuotesRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEventId

`func (o *PostApiQuotesRequest) GetEventId() string`

GetEventId returns the EventId field if non-nil, zero value otherwise.

### GetEventIdOk

`func (o *PostApiQuotesRequest) GetEventIdOk() (*string, bool)`

GetEventIdOk returns a tuple with the EventId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventId

`func (o *PostApiQuotesRequest) SetEventId(v string)`

SetEventId sets EventId field to given value.


### GetUserId

`func (o *PostApiQuotesRequest) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *PostApiQuotesRequest) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *PostApiQuotesRequest) SetUserId(v string)`

SetUserId sets UserId field to given value.


### GetHoldToken

`func (o *PostApiQuotesRequest) GetHoldToken() string`

GetHoldToken returns the HoldToken field if non-nil, zero value otherwise.

### GetHoldTokenOk

`func (o *PostApiQuotesRequest) GetHoldTokenOk() (*string, bool)`

GetHoldTokenOk returns a tuple with the HoldToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHoldToken

`func (o *PostApiQuotesRequest) SetHoldToken(v string)`

SetHoldToken sets HoldToken field to given value.


### GetSeatTicketTypes

`func (o *PostApiQuotesRequest) GetSeatTicketTypes() map[string]string`

GetSeatTicketTypes returns the SeatTicketTypes field if non-nil, zero value otherwise.

### GetSeatTicketTypesOk

`func (o *PostApiQuotesRequest) GetSeatTicketTypesOk() (*map[string]string, bool)`

GetSeatTicketTypesOk returns a tuple with the SeatTicketTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeatTicketTypes

`func (o *PostApiQuotesRequest) SetSeatTicketTypes(v map[string]string)`

SetSeatTicketTypes sets SeatTicketTypes field to given value.

### HasSeatTicketTypes

`func (o *PostApiQuotesRequest) HasSeatTicketTypes() bool`

HasSeatTicketTypes returns a boolean if a field has been set.

### GetTicketTypeId

`func (o *PostApiQuotesRequest) GetTicketTypeId() string`

GetTicketTypeId returns the TicketTypeId field if non-nil, zero value otherwise.

### GetTicketTypeIdOk

`func (o *PostApiQuotesRequest) GetTicketTypeIdOk() (*string, bool)`

GetTicketTypeIdOk returns a tuple with the TicketTypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTicketTypeId

`func (o *PostApiQuotesRequest) SetTicketTypeId(v string)`

SetTicketTypeId sets TicketTypeId field to given value.

### HasTicketTypeId

`func (o *PostApiQuotesRequest) HasTicketTypeId() bool`

HasTicketTypeId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


