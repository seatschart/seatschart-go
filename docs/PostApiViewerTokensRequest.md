# PostApiViewerTokensRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TtlSeconds** | Pointer to **int32** | Token lifetime in seconds (default 900). | [optional] 
**TicketBuyerId** | Pointer to **string** | Optional opaque ticket buyer id from your system (≤ 256 printable characters). Persisted on holds, orders, reservations, GA receipts, object statuses and status history; never interpreted as an internal user id. Query with GET /events/{eventId}/objects?ticketBuyerId&#x3D;…. | [optional] 

## Methods

### NewPostApiViewerTokensRequest

`func NewPostApiViewerTokensRequest() *PostApiViewerTokensRequest`

NewPostApiViewerTokensRequest instantiates a new PostApiViewerTokensRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostApiViewerTokensRequestWithDefaults

`func NewPostApiViewerTokensRequestWithDefaults() *PostApiViewerTokensRequest`

NewPostApiViewerTokensRequestWithDefaults instantiates a new PostApiViewerTokensRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTtlSeconds

`func (o *PostApiViewerTokensRequest) GetTtlSeconds() int32`

GetTtlSeconds returns the TtlSeconds field if non-nil, zero value otherwise.

### GetTtlSecondsOk

`func (o *PostApiViewerTokensRequest) GetTtlSecondsOk() (*int32, bool)`

GetTtlSecondsOk returns a tuple with the TtlSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTtlSeconds

`func (o *PostApiViewerTokensRequest) SetTtlSeconds(v int32)`

SetTtlSeconds sets TtlSeconds field to given value.

### HasTtlSeconds

`func (o *PostApiViewerTokensRequest) HasTtlSeconds() bool`

HasTtlSeconds returns a boolean if a field has been set.

### GetTicketBuyerId

`func (o *PostApiViewerTokensRequest) GetTicketBuyerId() string`

GetTicketBuyerId returns the TicketBuyerId field if non-nil, zero value otherwise.

### GetTicketBuyerIdOk

`func (o *PostApiViewerTokensRequest) GetTicketBuyerIdOk() (*string, bool)`

GetTicketBuyerIdOk returns a tuple with the TicketBuyerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTicketBuyerId

`func (o *PostApiViewerTokensRequest) SetTicketBuyerId(v string)`

SetTicketBuyerId sets TicketBuyerId field to given value.

### HasTicketBuyerId

`func (o *PostApiViewerTokensRequest) HasTicketBuyerId() bool`

HasTicketBuyerId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


