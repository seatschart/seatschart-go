# PostApiSeatsReserveRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EventId** | **string** |  | 
**SeatIds** | **[]string** |  | 
**UserId** | **string** |  | 
**Channel** | Pointer to **string** |  | [optional] 
**SeasonId** | Pointer to **string** |  | [optional] 
**EventIds** | Pointer to **[]string** |  | [optional] 
**HoldToken** | Pointer to **string** |  | [optional] 
**TicketBuyerId** | Pointer to **string** | Optional opaque ticket buyer id from your system (≤ 256 printable characters). Persisted on holds, orders, reservations, GA receipts, object statuses and status history; never interpreted as an internal user id. Query with GET /events/{eventId}/objects?ticketBuyerId&#x3D;…. | [optional] 
**SeatTicketTypes** | Pointer to **map[string]string** | Optional ticket type per seat (e.g. meia / inteira). Counted atomically against the event quotas when the hold is placed; quote and confirm may still change it within the quota. | [optional] 

## Methods

### NewPostApiSeatsReserveRequest

`func NewPostApiSeatsReserveRequest(eventId string, seatIds []string, userId string, ) *PostApiSeatsReserveRequest`

NewPostApiSeatsReserveRequest instantiates a new PostApiSeatsReserveRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostApiSeatsReserveRequestWithDefaults

`func NewPostApiSeatsReserveRequestWithDefaults() *PostApiSeatsReserveRequest`

NewPostApiSeatsReserveRequestWithDefaults instantiates a new PostApiSeatsReserveRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEventId

`func (o *PostApiSeatsReserveRequest) GetEventId() string`

GetEventId returns the EventId field if non-nil, zero value otherwise.

### GetEventIdOk

`func (o *PostApiSeatsReserveRequest) GetEventIdOk() (*string, bool)`

GetEventIdOk returns a tuple with the EventId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventId

`func (o *PostApiSeatsReserveRequest) SetEventId(v string)`

SetEventId sets EventId field to given value.


### GetSeatIds

`func (o *PostApiSeatsReserveRequest) GetSeatIds() []string`

GetSeatIds returns the SeatIds field if non-nil, zero value otherwise.

### GetSeatIdsOk

`func (o *PostApiSeatsReserveRequest) GetSeatIdsOk() (*[]string, bool)`

GetSeatIdsOk returns a tuple with the SeatIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeatIds

`func (o *PostApiSeatsReserveRequest) SetSeatIds(v []string)`

SetSeatIds sets SeatIds field to given value.


### GetUserId

`func (o *PostApiSeatsReserveRequest) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *PostApiSeatsReserveRequest) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *PostApiSeatsReserveRequest) SetUserId(v string)`

SetUserId sets UserId field to given value.


### GetChannel

`func (o *PostApiSeatsReserveRequest) GetChannel() string`

GetChannel returns the Channel field if non-nil, zero value otherwise.

### GetChannelOk

`func (o *PostApiSeatsReserveRequest) GetChannelOk() (*string, bool)`

GetChannelOk returns a tuple with the Channel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannel

`func (o *PostApiSeatsReserveRequest) SetChannel(v string)`

SetChannel sets Channel field to given value.

### HasChannel

`func (o *PostApiSeatsReserveRequest) HasChannel() bool`

HasChannel returns a boolean if a field has been set.

### GetSeasonId

`func (o *PostApiSeatsReserveRequest) GetSeasonId() string`

GetSeasonId returns the SeasonId field if non-nil, zero value otherwise.

### GetSeasonIdOk

`func (o *PostApiSeatsReserveRequest) GetSeasonIdOk() (*string, bool)`

GetSeasonIdOk returns a tuple with the SeasonId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeasonId

`func (o *PostApiSeatsReserveRequest) SetSeasonId(v string)`

SetSeasonId sets SeasonId field to given value.

### HasSeasonId

`func (o *PostApiSeatsReserveRequest) HasSeasonId() bool`

HasSeasonId returns a boolean if a field has been set.

### GetEventIds

`func (o *PostApiSeatsReserveRequest) GetEventIds() []string`

GetEventIds returns the EventIds field if non-nil, zero value otherwise.

### GetEventIdsOk

`func (o *PostApiSeatsReserveRequest) GetEventIdsOk() (*[]string, bool)`

GetEventIdsOk returns a tuple with the EventIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventIds

`func (o *PostApiSeatsReserveRequest) SetEventIds(v []string)`

SetEventIds sets EventIds field to given value.

### HasEventIds

`func (o *PostApiSeatsReserveRequest) HasEventIds() bool`

HasEventIds returns a boolean if a field has been set.

### GetHoldToken

`func (o *PostApiSeatsReserveRequest) GetHoldToken() string`

GetHoldToken returns the HoldToken field if non-nil, zero value otherwise.

### GetHoldTokenOk

`func (o *PostApiSeatsReserveRequest) GetHoldTokenOk() (*string, bool)`

GetHoldTokenOk returns a tuple with the HoldToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHoldToken

`func (o *PostApiSeatsReserveRequest) SetHoldToken(v string)`

SetHoldToken sets HoldToken field to given value.

### HasHoldToken

`func (o *PostApiSeatsReserveRequest) HasHoldToken() bool`

HasHoldToken returns a boolean if a field has been set.

### GetTicketBuyerId

`func (o *PostApiSeatsReserveRequest) GetTicketBuyerId() string`

GetTicketBuyerId returns the TicketBuyerId field if non-nil, zero value otherwise.

### GetTicketBuyerIdOk

`func (o *PostApiSeatsReserveRequest) GetTicketBuyerIdOk() (*string, bool)`

GetTicketBuyerIdOk returns a tuple with the TicketBuyerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTicketBuyerId

`func (o *PostApiSeatsReserveRequest) SetTicketBuyerId(v string)`

SetTicketBuyerId sets TicketBuyerId field to given value.

### HasTicketBuyerId

`func (o *PostApiSeatsReserveRequest) HasTicketBuyerId() bool`

HasTicketBuyerId returns a boolean if a field has been set.

### GetSeatTicketTypes

`func (o *PostApiSeatsReserveRequest) GetSeatTicketTypes() map[string]string`

GetSeatTicketTypes returns the SeatTicketTypes field if non-nil, zero value otherwise.

### GetSeatTicketTypesOk

`func (o *PostApiSeatsReserveRequest) GetSeatTicketTypesOk() (*map[string]string, bool)`

GetSeatTicketTypesOk returns a tuple with the SeatTicketTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeatTicketTypes

`func (o *PostApiSeatsReserveRequest) SetSeatTicketTypes(v map[string]string)`

SetSeatTicketTypes sets SeatTicketTypes field to given value.

### HasSeatTicketTypes

`func (o *PostApiSeatsReserveRequest) HasSeatTicketTypes() bool`

HasSeatTicketTypes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


