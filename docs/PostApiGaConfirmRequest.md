# PostApiGaConfirmRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EventId** | **string** |  | 
**AreaId** | **string** |  | 
**Quantity** | **int32** |  | 
**UserId** | **string** |  | 
**HoldToken** | Pointer to **string** |  | [optional] 
**TicketTypeId** | Pointer to **string** |  | [optional] 
**QuoteId** | Pointer to **string** |  | [optional] 
**ExternalPaymentId** | Pointer to **string** |  | [optional] 
**PaidAmountCents** | Pointer to **int32** |  | [optional] 
**PaidCurrency** | Pointer to **string** |  | [optional] 
**Channel** | Pointer to **string** |  | [optional] 
**TicketBuyerId** | Pointer to **string** | Optional opaque ticket buyer id from your system (≤ 256 printable characters). Persisted on holds, orders, reservations, GA receipts, object statuses and status history; never interpreted as an internal user id. Query with GET /events/{eventId}/objects?ticketBuyerId&#x3D;…. | [optional] 

## Methods

### NewPostApiGaConfirmRequest

`func NewPostApiGaConfirmRequest(eventId string, areaId string, quantity int32, userId string, ) *PostApiGaConfirmRequest`

NewPostApiGaConfirmRequest instantiates a new PostApiGaConfirmRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostApiGaConfirmRequestWithDefaults

`func NewPostApiGaConfirmRequestWithDefaults() *PostApiGaConfirmRequest`

NewPostApiGaConfirmRequestWithDefaults instantiates a new PostApiGaConfirmRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEventId

`func (o *PostApiGaConfirmRequest) GetEventId() string`

GetEventId returns the EventId field if non-nil, zero value otherwise.

### GetEventIdOk

`func (o *PostApiGaConfirmRequest) GetEventIdOk() (*string, bool)`

GetEventIdOk returns a tuple with the EventId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventId

`func (o *PostApiGaConfirmRequest) SetEventId(v string)`

SetEventId sets EventId field to given value.


### GetAreaId

`func (o *PostApiGaConfirmRequest) GetAreaId() string`

GetAreaId returns the AreaId field if non-nil, zero value otherwise.

### GetAreaIdOk

`func (o *PostApiGaConfirmRequest) GetAreaIdOk() (*string, bool)`

GetAreaIdOk returns a tuple with the AreaId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAreaId

`func (o *PostApiGaConfirmRequest) SetAreaId(v string)`

SetAreaId sets AreaId field to given value.


### GetQuantity

`func (o *PostApiGaConfirmRequest) GetQuantity() int32`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *PostApiGaConfirmRequest) GetQuantityOk() (*int32, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *PostApiGaConfirmRequest) SetQuantity(v int32)`

SetQuantity sets Quantity field to given value.


### GetUserId

`func (o *PostApiGaConfirmRequest) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *PostApiGaConfirmRequest) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *PostApiGaConfirmRequest) SetUserId(v string)`

SetUserId sets UserId field to given value.


### GetHoldToken

`func (o *PostApiGaConfirmRequest) GetHoldToken() string`

GetHoldToken returns the HoldToken field if non-nil, zero value otherwise.

### GetHoldTokenOk

`func (o *PostApiGaConfirmRequest) GetHoldTokenOk() (*string, bool)`

GetHoldTokenOk returns a tuple with the HoldToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHoldToken

`func (o *PostApiGaConfirmRequest) SetHoldToken(v string)`

SetHoldToken sets HoldToken field to given value.

### HasHoldToken

`func (o *PostApiGaConfirmRequest) HasHoldToken() bool`

HasHoldToken returns a boolean if a field has been set.

### GetTicketTypeId

`func (o *PostApiGaConfirmRequest) GetTicketTypeId() string`

GetTicketTypeId returns the TicketTypeId field if non-nil, zero value otherwise.

### GetTicketTypeIdOk

`func (o *PostApiGaConfirmRequest) GetTicketTypeIdOk() (*string, bool)`

GetTicketTypeIdOk returns a tuple with the TicketTypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTicketTypeId

`func (o *PostApiGaConfirmRequest) SetTicketTypeId(v string)`

SetTicketTypeId sets TicketTypeId field to given value.

### HasTicketTypeId

`func (o *PostApiGaConfirmRequest) HasTicketTypeId() bool`

HasTicketTypeId returns a boolean if a field has been set.

### GetQuoteId

`func (o *PostApiGaConfirmRequest) GetQuoteId() string`

GetQuoteId returns the QuoteId field if non-nil, zero value otherwise.

### GetQuoteIdOk

`func (o *PostApiGaConfirmRequest) GetQuoteIdOk() (*string, bool)`

GetQuoteIdOk returns a tuple with the QuoteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuoteId

`func (o *PostApiGaConfirmRequest) SetQuoteId(v string)`

SetQuoteId sets QuoteId field to given value.

### HasQuoteId

`func (o *PostApiGaConfirmRequest) HasQuoteId() bool`

HasQuoteId returns a boolean if a field has been set.

### GetExternalPaymentId

`func (o *PostApiGaConfirmRequest) GetExternalPaymentId() string`

GetExternalPaymentId returns the ExternalPaymentId field if non-nil, zero value otherwise.

### GetExternalPaymentIdOk

`func (o *PostApiGaConfirmRequest) GetExternalPaymentIdOk() (*string, bool)`

GetExternalPaymentIdOk returns a tuple with the ExternalPaymentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalPaymentId

`func (o *PostApiGaConfirmRequest) SetExternalPaymentId(v string)`

SetExternalPaymentId sets ExternalPaymentId field to given value.

### HasExternalPaymentId

`func (o *PostApiGaConfirmRequest) HasExternalPaymentId() bool`

HasExternalPaymentId returns a boolean if a field has been set.

### GetPaidAmountCents

`func (o *PostApiGaConfirmRequest) GetPaidAmountCents() int32`

GetPaidAmountCents returns the PaidAmountCents field if non-nil, zero value otherwise.

### GetPaidAmountCentsOk

`func (o *PostApiGaConfirmRequest) GetPaidAmountCentsOk() (*int32, bool)`

GetPaidAmountCentsOk returns a tuple with the PaidAmountCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaidAmountCents

`func (o *PostApiGaConfirmRequest) SetPaidAmountCents(v int32)`

SetPaidAmountCents sets PaidAmountCents field to given value.

### HasPaidAmountCents

`func (o *PostApiGaConfirmRequest) HasPaidAmountCents() bool`

HasPaidAmountCents returns a boolean if a field has been set.

### GetPaidCurrency

`func (o *PostApiGaConfirmRequest) GetPaidCurrency() string`

GetPaidCurrency returns the PaidCurrency field if non-nil, zero value otherwise.

### GetPaidCurrencyOk

`func (o *PostApiGaConfirmRequest) GetPaidCurrencyOk() (*string, bool)`

GetPaidCurrencyOk returns a tuple with the PaidCurrency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaidCurrency

`func (o *PostApiGaConfirmRequest) SetPaidCurrency(v string)`

SetPaidCurrency sets PaidCurrency field to given value.

### HasPaidCurrency

`func (o *PostApiGaConfirmRequest) HasPaidCurrency() bool`

HasPaidCurrency returns a boolean if a field has been set.

### GetChannel

`func (o *PostApiGaConfirmRequest) GetChannel() string`

GetChannel returns the Channel field if non-nil, zero value otherwise.

### GetChannelOk

`func (o *PostApiGaConfirmRequest) GetChannelOk() (*string, bool)`

GetChannelOk returns a tuple with the Channel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannel

`func (o *PostApiGaConfirmRequest) SetChannel(v string)`

SetChannel sets Channel field to given value.

### HasChannel

`func (o *PostApiGaConfirmRequest) HasChannel() bool`

HasChannel returns a boolean if a field has been set.

### GetTicketBuyerId

`func (o *PostApiGaConfirmRequest) GetTicketBuyerId() string`

GetTicketBuyerId returns the TicketBuyerId field if non-nil, zero value otherwise.

### GetTicketBuyerIdOk

`func (o *PostApiGaConfirmRequest) GetTicketBuyerIdOk() (*string, bool)`

GetTicketBuyerIdOk returns a tuple with the TicketBuyerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTicketBuyerId

`func (o *PostApiGaConfirmRequest) SetTicketBuyerId(v string)`

SetTicketBuyerId sets TicketBuyerId field to given value.

### HasTicketBuyerId

`func (o *PostApiGaConfirmRequest) HasTicketBuyerId() bool`

HasTicketBuyerId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


