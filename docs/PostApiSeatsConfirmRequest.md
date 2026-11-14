# PostApiSeatsConfirmRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Scope** | Pointer to **string** |  | [optional] 
**ExternalOrderId** | Pointer to **string** |  | [optional] 
**ExternalPaymentId** | Pointer to **string** |  | [optional] 
**QuoteId** | Pointer to **string** |  | [optional] 
**PaidAmountCents** | Pointer to **int32** |  | [optional] 
**PaidCurrency** | Pointer to **string** |  | [optional] 
**EventId** | **string** |  | 
**SeatIds** | **[]string** |  | 
**UserId** | **string** |  | 
**SeatTicketTypes** | Pointer to **map[string]string** |  | [optional] 
**HoldToken** | Pointer to **string** |  | [optional] 
**Channel** | Pointer to **string** |  | [optional] 
**TicketBuyerId** | Pointer to **string** | Optional opaque ticket buyer id from your system (≤ 256 printable characters). Persisted on holds, orders, reservations, GA receipts, object statuses and status history; never interpreted as an internal user id. Query with GET /events/{eventId}/objects?ticketBuyerId&#x3D;…. | [optional] 
**SeatPrices** | Pointer to **map[string]int32** | Partner and synced-catalog pricing: amount charged per seat in cents (stored on the order items for reports; validated against the synced price table). | [optional] 

## Methods

### NewPostApiSeatsConfirmRequest

`func NewPostApiSeatsConfirmRequest(eventId string, seatIds []string, userId string, ) *PostApiSeatsConfirmRequest`

NewPostApiSeatsConfirmRequest instantiates a new PostApiSeatsConfirmRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostApiSeatsConfirmRequestWithDefaults

`func NewPostApiSeatsConfirmRequestWithDefaults() *PostApiSeatsConfirmRequest`

NewPostApiSeatsConfirmRequestWithDefaults instantiates a new PostApiSeatsConfirmRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetScope

`func (o *PostApiSeatsConfirmRequest) GetScope() string`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *PostApiSeatsConfirmRequest) GetScopeOk() (*string, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *PostApiSeatsConfirmRequest) SetScope(v string)`

SetScope sets Scope field to given value.

### HasScope

`func (o *PostApiSeatsConfirmRequest) HasScope() bool`

HasScope returns a boolean if a field has been set.

### GetExternalOrderId

`func (o *PostApiSeatsConfirmRequest) GetExternalOrderId() string`

GetExternalOrderId returns the ExternalOrderId field if non-nil, zero value otherwise.

### GetExternalOrderIdOk

`func (o *PostApiSeatsConfirmRequest) GetExternalOrderIdOk() (*string, bool)`

GetExternalOrderIdOk returns a tuple with the ExternalOrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalOrderId

`func (o *PostApiSeatsConfirmRequest) SetExternalOrderId(v string)`

SetExternalOrderId sets ExternalOrderId field to given value.

### HasExternalOrderId

`func (o *PostApiSeatsConfirmRequest) HasExternalOrderId() bool`

HasExternalOrderId returns a boolean if a field has been set.

### GetExternalPaymentId

`func (o *PostApiSeatsConfirmRequest) GetExternalPaymentId() string`

GetExternalPaymentId returns the ExternalPaymentId field if non-nil, zero value otherwise.

### GetExternalPaymentIdOk

`func (o *PostApiSeatsConfirmRequest) GetExternalPaymentIdOk() (*string, bool)`

GetExternalPaymentIdOk returns a tuple with the ExternalPaymentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalPaymentId

`func (o *PostApiSeatsConfirmRequest) SetExternalPaymentId(v string)`

SetExternalPaymentId sets ExternalPaymentId field to given value.

### HasExternalPaymentId

`func (o *PostApiSeatsConfirmRequest) HasExternalPaymentId() bool`

HasExternalPaymentId returns a boolean if a field has been set.

### GetQuoteId

`func (o *PostApiSeatsConfirmRequest) GetQuoteId() string`

GetQuoteId returns the QuoteId field if non-nil, zero value otherwise.

### GetQuoteIdOk

`func (o *PostApiSeatsConfirmRequest) GetQuoteIdOk() (*string, bool)`

GetQuoteIdOk returns a tuple with the QuoteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuoteId

`func (o *PostApiSeatsConfirmRequest) SetQuoteId(v string)`

SetQuoteId sets QuoteId field to given value.

### HasQuoteId

`func (o *PostApiSeatsConfirmRequest) HasQuoteId() bool`

HasQuoteId returns a boolean if a field has been set.

### GetPaidAmountCents

`func (o *PostApiSeatsConfirmRequest) GetPaidAmountCents() int32`

GetPaidAmountCents returns the PaidAmountCents field if non-nil, zero value otherwise.

### GetPaidAmountCentsOk

`func (o *PostApiSeatsConfirmRequest) GetPaidAmountCentsOk() (*int32, bool)`

GetPaidAmountCentsOk returns a tuple with the PaidAmountCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaidAmountCents

`func (o *PostApiSeatsConfirmRequest) SetPaidAmountCents(v int32)`

SetPaidAmountCents sets PaidAmountCents field to given value.

### HasPaidAmountCents

`func (o *PostApiSeatsConfirmRequest) HasPaidAmountCents() bool`

HasPaidAmountCents returns a boolean if a field has been set.

### GetPaidCurrency

`func (o *PostApiSeatsConfirmRequest) GetPaidCurrency() string`

GetPaidCurrency returns the PaidCurrency field if non-nil, zero value otherwise.

### GetPaidCurrencyOk

`func (o *PostApiSeatsConfirmRequest) GetPaidCurrencyOk() (*string, bool)`

GetPaidCurrencyOk returns a tuple with the PaidCurrency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaidCurrency

`func (o *PostApiSeatsConfirmRequest) SetPaidCurrency(v string)`

SetPaidCurrency sets PaidCurrency field to given value.

### HasPaidCurrency

`func (o *PostApiSeatsConfirmRequest) HasPaidCurrency() bool`

HasPaidCurrency returns a boolean if a field has been set.

### GetEventId

`func (o *PostApiSeatsConfirmRequest) GetEventId() string`

GetEventId returns the EventId field if non-nil, zero value otherwise.

### GetEventIdOk

`func (o *PostApiSeatsConfirmRequest) GetEventIdOk() (*string, bool)`

GetEventIdOk returns a tuple with the EventId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventId

`func (o *PostApiSeatsConfirmRequest) SetEventId(v string)`

SetEventId sets EventId field to given value.


### GetSeatIds

`func (o *PostApiSeatsConfirmRequest) GetSeatIds() []string`

GetSeatIds returns the SeatIds field if non-nil, zero value otherwise.

### GetSeatIdsOk

`func (o *PostApiSeatsConfirmRequest) GetSeatIdsOk() (*[]string, bool)`

GetSeatIdsOk returns a tuple with the SeatIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeatIds

`func (o *PostApiSeatsConfirmRequest) SetSeatIds(v []string)`

SetSeatIds sets SeatIds field to given value.


### GetUserId

`func (o *PostApiSeatsConfirmRequest) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *PostApiSeatsConfirmRequest) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *PostApiSeatsConfirmRequest) SetUserId(v string)`

SetUserId sets UserId field to given value.


### GetSeatTicketTypes

`func (o *PostApiSeatsConfirmRequest) GetSeatTicketTypes() map[string]string`

GetSeatTicketTypes returns the SeatTicketTypes field if non-nil, zero value otherwise.

### GetSeatTicketTypesOk

`func (o *PostApiSeatsConfirmRequest) GetSeatTicketTypesOk() (*map[string]string, bool)`

GetSeatTicketTypesOk returns a tuple with the SeatTicketTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeatTicketTypes

`func (o *PostApiSeatsConfirmRequest) SetSeatTicketTypes(v map[string]string)`

SetSeatTicketTypes sets SeatTicketTypes field to given value.

### HasSeatTicketTypes

`func (o *PostApiSeatsConfirmRequest) HasSeatTicketTypes() bool`

HasSeatTicketTypes returns a boolean if a field has been set.

### GetHoldToken

`func (o *PostApiSeatsConfirmRequest) GetHoldToken() string`

GetHoldToken returns the HoldToken field if non-nil, zero value otherwise.

### GetHoldTokenOk

`func (o *PostApiSeatsConfirmRequest) GetHoldTokenOk() (*string, bool)`

GetHoldTokenOk returns a tuple with the HoldToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHoldToken

`func (o *PostApiSeatsConfirmRequest) SetHoldToken(v string)`

SetHoldToken sets HoldToken field to given value.

### HasHoldToken

`func (o *PostApiSeatsConfirmRequest) HasHoldToken() bool`

HasHoldToken returns a boolean if a field has been set.

### GetChannel

`func (o *PostApiSeatsConfirmRequest) GetChannel() string`

GetChannel returns the Channel field if non-nil, zero value otherwise.

### GetChannelOk

`func (o *PostApiSeatsConfirmRequest) GetChannelOk() (*string, bool)`

GetChannelOk returns a tuple with the Channel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannel

`func (o *PostApiSeatsConfirmRequest) SetChannel(v string)`

SetChannel sets Channel field to given value.

### HasChannel

`func (o *PostApiSeatsConfirmRequest) HasChannel() bool`

HasChannel returns a boolean if a field has been set.

### GetTicketBuyerId

`func (o *PostApiSeatsConfirmRequest) GetTicketBuyerId() string`

GetTicketBuyerId returns the TicketBuyerId field if non-nil, zero value otherwise.

### GetTicketBuyerIdOk

`func (o *PostApiSeatsConfirmRequest) GetTicketBuyerIdOk() (*string, bool)`

GetTicketBuyerIdOk returns a tuple with the TicketBuyerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTicketBuyerId

`func (o *PostApiSeatsConfirmRequest) SetTicketBuyerId(v string)`

SetTicketBuyerId sets TicketBuyerId field to given value.

### HasTicketBuyerId

`func (o *PostApiSeatsConfirmRequest) HasTicketBuyerId() bool`

HasTicketBuyerId returns a boolean if a field has been set.

### GetSeatPrices

`func (o *PostApiSeatsConfirmRequest) GetSeatPrices() map[string]int32`

GetSeatPrices returns the SeatPrices field if non-nil, zero value otherwise.

### GetSeatPricesOk

`func (o *PostApiSeatsConfirmRequest) GetSeatPricesOk() (*map[string]int32, bool)`

GetSeatPricesOk returns a tuple with the SeatPrices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeatPrices

`func (o *PostApiSeatsConfirmRequest) SetSeatPrices(v map[string]int32)`

SetSeatPrices sets SeatPrices field to given value.

### HasSeatPrices

`func (o *PostApiSeatsConfirmRequest) HasSeatPrices() bool`

HasSeatPrices returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


