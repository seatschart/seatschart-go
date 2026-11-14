# PostApiGaReserveRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EventId** | **string** |  | 
**AreaId** | **string** |  | 
**Quantity** | **int32** |  | 
**UserId** | **string** |  | 
**Channel** | Pointer to **string** |  | [optional] 
**TicketBuyerId** | Pointer to **string** | Optional opaque ticket buyer id from your system (≤ 256 printable characters). Persisted on holds, orders, reservations, GA receipts, object statuses and status history; never interpreted as an internal user id. Query with GET /events/{eventId}/objects?ticketBuyerId&#x3D;…. | [optional] 
**TicketTypeId** | Pointer to **string** | Optional ticket type of every place (e.g. meia), counted against the event quotas at hold time. | [optional] 

## Methods

### NewPostApiGaReserveRequest

`func NewPostApiGaReserveRequest(eventId string, areaId string, quantity int32, userId string, ) *PostApiGaReserveRequest`

NewPostApiGaReserveRequest instantiates a new PostApiGaReserveRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostApiGaReserveRequestWithDefaults

`func NewPostApiGaReserveRequestWithDefaults() *PostApiGaReserveRequest`

NewPostApiGaReserveRequestWithDefaults instantiates a new PostApiGaReserveRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEventId

`func (o *PostApiGaReserveRequest) GetEventId() string`

GetEventId returns the EventId field if non-nil, zero value otherwise.

### GetEventIdOk

`func (o *PostApiGaReserveRequest) GetEventIdOk() (*string, bool)`

GetEventIdOk returns a tuple with the EventId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventId

`func (o *PostApiGaReserveRequest) SetEventId(v string)`

SetEventId sets EventId field to given value.


### GetAreaId

`func (o *PostApiGaReserveRequest) GetAreaId() string`

GetAreaId returns the AreaId field if non-nil, zero value otherwise.

### GetAreaIdOk

`func (o *PostApiGaReserveRequest) GetAreaIdOk() (*string, bool)`

GetAreaIdOk returns a tuple with the AreaId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAreaId

`func (o *PostApiGaReserveRequest) SetAreaId(v string)`

SetAreaId sets AreaId field to given value.


### GetQuantity

`func (o *PostApiGaReserveRequest) GetQuantity() int32`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *PostApiGaReserveRequest) GetQuantityOk() (*int32, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *PostApiGaReserveRequest) SetQuantity(v int32)`

SetQuantity sets Quantity field to given value.


### GetUserId

`func (o *PostApiGaReserveRequest) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *PostApiGaReserveRequest) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *PostApiGaReserveRequest) SetUserId(v string)`

SetUserId sets UserId field to given value.


### GetChannel

`func (o *PostApiGaReserveRequest) GetChannel() string`

GetChannel returns the Channel field if non-nil, zero value otherwise.

### GetChannelOk

`func (o *PostApiGaReserveRequest) GetChannelOk() (*string, bool)`

GetChannelOk returns a tuple with the Channel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannel

`func (o *PostApiGaReserveRequest) SetChannel(v string)`

SetChannel sets Channel field to given value.

### HasChannel

`func (o *PostApiGaReserveRequest) HasChannel() bool`

HasChannel returns a boolean if a field has been set.

### GetTicketBuyerId

`func (o *PostApiGaReserveRequest) GetTicketBuyerId() string`

GetTicketBuyerId returns the TicketBuyerId field if non-nil, zero value otherwise.

### GetTicketBuyerIdOk

`func (o *PostApiGaReserveRequest) GetTicketBuyerIdOk() (*string, bool)`

GetTicketBuyerIdOk returns a tuple with the TicketBuyerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTicketBuyerId

`func (o *PostApiGaReserveRequest) SetTicketBuyerId(v string)`

SetTicketBuyerId sets TicketBuyerId field to given value.

### HasTicketBuyerId

`func (o *PostApiGaReserveRequest) HasTicketBuyerId() bool`

HasTicketBuyerId returns a boolean if a field has been set.

### GetTicketTypeId

`func (o *PostApiGaReserveRequest) GetTicketTypeId() string`

GetTicketTypeId returns the TicketTypeId field if non-nil, zero value otherwise.

### GetTicketTypeIdOk

`func (o *PostApiGaReserveRequest) GetTicketTypeIdOk() (*string, bool)`

GetTicketTypeIdOk returns a tuple with the TicketTypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTicketTypeId

`func (o *PostApiGaReserveRequest) SetTicketTypeId(v string)`

SetTicketTypeId sets TicketTypeId field to given value.

### HasTicketTypeId

`func (o *PostApiGaReserveRequest) HasTicketTypeId() bool`

HasTicketTypeId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


