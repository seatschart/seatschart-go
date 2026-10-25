# PostApiHoldTokensRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EventId** | **string** |  | 
**UserId** | **string** |  | 
**TicketBuyerId** | Pointer to **string** | Optional opaque ticket buyer id from your system (≤ 256 printable characters). Persisted on holds, orders, reservations, GA receipts, object statuses and status history; never interpreted as an internal user id. Query with GET /events/{eventId}/objects?ticketBuyerId&#x3D;…. | [optional] 
**ExpiresInMinutes** | Pointer to **int32** | Session lifetime in minutes (1–120; default: the workspace hold period). API keys and dashboard sessions only. | [optional] 

## Methods

### NewPostApiHoldTokensRequest

`func NewPostApiHoldTokensRequest(eventId string, userId string, ) *PostApiHoldTokensRequest`

NewPostApiHoldTokensRequest instantiates a new PostApiHoldTokensRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostApiHoldTokensRequestWithDefaults

`func NewPostApiHoldTokensRequestWithDefaults() *PostApiHoldTokensRequest`

NewPostApiHoldTokensRequestWithDefaults instantiates a new PostApiHoldTokensRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEventId

`func (o *PostApiHoldTokensRequest) GetEventId() string`

GetEventId returns the EventId field if non-nil, zero value otherwise.

### GetEventIdOk

`func (o *PostApiHoldTokensRequest) GetEventIdOk() (*string, bool)`

GetEventIdOk returns a tuple with the EventId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventId

`func (o *PostApiHoldTokensRequest) SetEventId(v string)`

SetEventId sets EventId field to given value.


### GetUserId

`func (o *PostApiHoldTokensRequest) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *PostApiHoldTokensRequest) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *PostApiHoldTokensRequest) SetUserId(v string)`

SetUserId sets UserId field to given value.


### GetTicketBuyerId

`func (o *PostApiHoldTokensRequest) GetTicketBuyerId() string`

GetTicketBuyerId returns the TicketBuyerId field if non-nil, zero value otherwise.

### GetTicketBuyerIdOk

`func (o *PostApiHoldTokensRequest) GetTicketBuyerIdOk() (*string, bool)`

GetTicketBuyerIdOk returns a tuple with the TicketBuyerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTicketBuyerId

`func (o *PostApiHoldTokensRequest) SetTicketBuyerId(v string)`

SetTicketBuyerId sets TicketBuyerId field to given value.

### HasTicketBuyerId

`func (o *PostApiHoldTokensRequest) HasTicketBuyerId() bool`

HasTicketBuyerId returns a boolean if a field has been set.

### GetExpiresInMinutes

`func (o *PostApiHoldTokensRequest) GetExpiresInMinutes() int32`

GetExpiresInMinutes returns the ExpiresInMinutes field if non-nil, zero value otherwise.

### GetExpiresInMinutesOk

`func (o *PostApiHoldTokensRequest) GetExpiresInMinutesOk() (*int32, bool)`

GetExpiresInMinutesOk returns a tuple with the ExpiresInMinutes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresInMinutes

`func (o *PostApiHoldTokensRequest) SetExpiresInMinutes(v int32)`

SetExpiresInMinutes sets ExpiresInMinutes field to given value.

### HasExpiresInMinutes

`func (o *PostApiHoldTokensRequest) HasExpiresInMinutes() bool`

HasExpiresInMinutes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


