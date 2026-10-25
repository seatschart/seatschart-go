# PostApiEventsActionsChangeObjectStatusRequestStatusChangesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EventId** | Pointer to **string** | Event UUID (exactly one of eventId / eventKey). | [optional] 
**EventKey** | Pointer to **string** | Event key (your event id; exactly one of eventId / eventKey). | [optional] 
**Objects** | [**[]PostApiEventsActionsChangeObjectStatusRequestStatusChangesInnerObjectsInner**](PostApiEventsActionsChangeObjectStatusRequestStatusChangesInnerObjectsInner.md) |  | 
**Status** | Pointer to **string** | &#x60;booked&#x60;, &#x60;free&#x60;, &#x60;resale&#x60;, &#x60;retired&#x60; or a custom status (required unless type is RELEASE). | [optional] 
**Type** | Pointer to **string** | &#x60;RELEASE&#x60; frees the objects (same as status &#x60;free&#x60;). | [optional] 
**OrderId** | Pointer to **string** | Existing booking order to act on (UUID). | [optional] 
**HoldToken** | Pointer to **string** | Books seats held by this buyer hold token (status &#x60;booked&#x60;). | [optional] 
**KeepExtraData** | Pointer to **bool** |  | [optional] 
**ResalePriceCents** | Pointer to **int32** |  | [optional] 
**ResaleListingId** | Pointer to **string** | Resale listing: stored on extraData.resaleListingId of every object; the renderer can select a listing as one group (resaleListings.selectAsGroup). | [optional] 
**AllowedPreviousStatuses** | Pointer to **[]string** |  | [optional] 
**RejectedPreviousStatuses** | Pointer to **[]string** |  | [optional] 
**ChannelKeys** | Pointer to **[]string** |  | [optional] 
**IgnoreChannels** | Pointer to **bool** |  | [optional] 
**TicketBuyerId** | Pointer to **string** | Optional opaque ticket buyer id from your system (≤ 256 printable characters). Persisted on holds, orders, reservations, GA receipts, object statuses and status history; never interpreted as an internal user id. Query with GET /events/{eventId}/objects?ticketBuyerId&#x3D;…. | [optional] 

## Methods

### NewPostApiEventsActionsChangeObjectStatusRequestStatusChangesInner

`func NewPostApiEventsActionsChangeObjectStatusRequestStatusChangesInner(objects []PostApiEventsActionsChangeObjectStatusRequestStatusChangesInnerObjectsInner, ) *PostApiEventsActionsChangeObjectStatusRequestStatusChangesInner`

NewPostApiEventsActionsChangeObjectStatusRequestStatusChangesInner instantiates a new PostApiEventsActionsChangeObjectStatusRequestStatusChangesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostApiEventsActionsChangeObjectStatusRequestStatusChangesInnerWithDefaults

`func NewPostApiEventsActionsChangeObjectStatusRequestStatusChangesInnerWithDefaults() *PostApiEventsActionsChangeObjectStatusRequestStatusChangesInner`

NewPostApiEventsActionsChangeObjectStatusRequestStatusChangesInnerWithDefaults instantiates a new PostApiEventsActionsChangeObjectStatusRequestStatusChangesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEventId

`func (o *PostApiEventsActionsChangeObjectStatusRequestStatusChangesInner) GetEventId() string`

GetEventId returns the EventId field if non-nil, zero value otherwise.

### GetEventIdOk

`func (o *PostApiEventsActionsChangeObjectStatusRequestStatusChangesInner) GetEventIdOk() (*string, bool)`

GetEventIdOk returns a tuple with the EventId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventId

`func (o *PostApiEventsActionsChangeObjectStatusRequestStatusChangesInner) SetEventId(v string)`

SetEventId sets EventId field to given value.

### HasEventId

`func (o *PostApiEventsActionsChangeObjectStatusRequestStatusChangesInner) HasEventId() bool`

HasEventId returns a boolean if a field has been set.

### GetEventKey

`func (o *PostApiEventsActionsChangeObjectStatusRequestStatusChangesInner) GetEventKey() string`

GetEventKey returns the EventKey field if non-nil, zero value otherwise.

### GetEventKeyOk

`func (o *PostApiEventsActionsChangeObjectStatusRequestStatusChangesInner) GetEventKeyOk() (*string, bool)`

GetEventKeyOk returns a tuple with the EventKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventKey

`func (o *PostApiEventsActionsChangeObjectStatusRequestStatusChangesInner) SetEventKey(v string)`

SetEventKey sets EventKey field to given value.

### HasEventKey

`func (o *PostApiEventsActionsChangeObjectStatusRequestStatusChangesInner) HasEventKey() bool`

HasEventKey returns a boolean if a field has been set.

### GetObjects

`func (o *PostApiEventsActionsChangeObjectStatusRequestStatusChangesInner) GetObjects() []PostApiEventsActionsChangeObjectStatusRequestStatusChangesInnerObjectsInner`

GetObjects returns the Objects field if non-nil, zero value otherwise.

### GetObjectsOk

`func (o *PostApiEventsActionsChangeObjectStatusRequestStatusChangesInner) GetObjectsOk() (*[]PostApiEventsActionsChangeObjectStatusRequestStatusChangesInnerObjectsInner, bool)`

GetObjectsOk returns a tuple with the Objects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjects

`func (o *PostApiEventsActionsChangeObjectStatusRequestStatusChangesInner) SetObjects(v []PostApiEventsActionsChangeObjectStatusRequestStatusChangesInnerObjectsInner)`

SetObjects sets Objects field to given value.


### GetStatus

`func (o *PostApiEventsActionsChangeObjectStatusRequestStatusChangesInner) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *PostApiEventsActionsChangeObjectStatusRequestStatusChangesInner) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *PostApiEventsActionsChangeObjectStatusRequestStatusChangesInner) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *PostApiEventsActionsChangeObjectStatusRequestStatusChangesInner) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetType

`func (o *PostApiEventsActionsChangeObjectStatusRequestStatusChangesInner) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PostApiEventsActionsChangeObjectStatusRequestStatusChangesInner) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PostApiEventsActionsChangeObjectStatusRequestStatusChangesInner) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *PostApiEventsActionsChangeObjectStatusRequestStatusChangesInner) HasType() bool`

HasType returns a boolean if a field has been set.

### GetOrderId

`func (o *PostApiEventsActionsChangeObjectStatusRequestStatusChangesInner) GetOrderId() string`

GetOrderId returns the OrderId field if non-nil, zero value otherwise.

### GetOrderIdOk

`func (o *PostApiEventsActionsChangeObjectStatusRequestStatusChangesInner) GetOrderIdOk() (*string, bool)`

GetOrderIdOk returns a tuple with the OrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderId

`func (o *PostApiEventsActionsChangeObjectStatusRequestStatusChangesInner) SetOrderId(v string)`

SetOrderId sets OrderId field to given value.

### HasOrderId

`func (o *PostApiEventsActionsChangeObjectStatusRequestStatusChangesInner) HasOrderId() bool`

HasOrderId returns a boolean if a field has been set.

### GetHoldToken

`func (o *PostApiEventsActionsChangeObjectStatusRequestStatusChangesInner) GetHoldToken() string`

GetHoldToken returns the HoldToken field if non-nil, zero value otherwise.

### GetHoldTokenOk

`func (o *PostApiEventsActionsChangeObjectStatusRequestStatusChangesInner) GetHoldTokenOk() (*string, bool)`

GetHoldTokenOk returns a tuple with the HoldToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHoldToken

`func (o *PostApiEventsActionsChangeObjectStatusRequestStatusChangesInner) SetHoldToken(v string)`

SetHoldToken sets HoldToken field to given value.

### HasHoldToken

`func (o *PostApiEventsActionsChangeObjectStatusRequestStatusChangesInner) HasHoldToken() bool`

HasHoldToken returns a boolean if a field has been set.

### GetKeepExtraData

`func (o *PostApiEventsActionsChangeObjectStatusRequestStatusChangesInner) GetKeepExtraData() bool`

GetKeepExtraData returns the KeepExtraData field if non-nil, zero value otherwise.

### GetKeepExtraDataOk

`func (o *PostApiEventsActionsChangeObjectStatusRequestStatusChangesInner) GetKeepExtraDataOk() (*bool, bool)`

GetKeepExtraDataOk returns a tuple with the KeepExtraData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeepExtraData

`func (o *PostApiEventsActionsChangeObjectStatusRequestStatusChangesInner) SetKeepExtraData(v bool)`

SetKeepExtraData sets KeepExtraData field to given value.

### HasKeepExtraData

`func (o *PostApiEventsActionsChangeObjectStatusRequestStatusChangesInner) HasKeepExtraData() bool`

HasKeepExtraData returns a boolean if a field has been set.

### GetResalePriceCents

`func (o *PostApiEventsActionsChangeObjectStatusRequestStatusChangesInner) GetResalePriceCents() int32`

GetResalePriceCents returns the ResalePriceCents field if non-nil, zero value otherwise.

### GetResalePriceCentsOk

`func (o *PostApiEventsActionsChangeObjectStatusRequestStatusChangesInner) GetResalePriceCentsOk() (*int32, bool)`

GetResalePriceCentsOk returns a tuple with the ResalePriceCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResalePriceCents

`func (o *PostApiEventsActionsChangeObjectStatusRequestStatusChangesInner) SetResalePriceCents(v int32)`

SetResalePriceCents sets ResalePriceCents field to given value.

### HasResalePriceCents

`func (o *PostApiEventsActionsChangeObjectStatusRequestStatusChangesInner) HasResalePriceCents() bool`

HasResalePriceCents returns a boolean if a field has been set.

### GetResaleListingId

`func (o *PostApiEventsActionsChangeObjectStatusRequestStatusChangesInner) GetResaleListingId() string`

GetResaleListingId returns the ResaleListingId field if non-nil, zero value otherwise.

### GetResaleListingIdOk

`func (o *PostApiEventsActionsChangeObjectStatusRequestStatusChangesInner) GetResaleListingIdOk() (*string, bool)`

GetResaleListingIdOk returns a tuple with the ResaleListingId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResaleListingId

`func (o *PostApiEventsActionsChangeObjectStatusRequestStatusChangesInner) SetResaleListingId(v string)`

SetResaleListingId sets ResaleListingId field to given value.

### HasResaleListingId

`func (o *PostApiEventsActionsChangeObjectStatusRequestStatusChangesInner) HasResaleListingId() bool`

HasResaleListingId returns a boolean if a field has been set.

### GetAllowedPreviousStatuses

`func (o *PostApiEventsActionsChangeObjectStatusRequestStatusChangesInner) GetAllowedPreviousStatuses() []string`

GetAllowedPreviousStatuses returns the AllowedPreviousStatuses field if non-nil, zero value otherwise.

### GetAllowedPreviousStatusesOk

`func (o *PostApiEventsActionsChangeObjectStatusRequestStatusChangesInner) GetAllowedPreviousStatusesOk() (*[]string, bool)`

GetAllowedPreviousStatusesOk returns a tuple with the AllowedPreviousStatuses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedPreviousStatuses

`func (o *PostApiEventsActionsChangeObjectStatusRequestStatusChangesInner) SetAllowedPreviousStatuses(v []string)`

SetAllowedPreviousStatuses sets AllowedPreviousStatuses field to given value.

### HasAllowedPreviousStatuses

`func (o *PostApiEventsActionsChangeObjectStatusRequestStatusChangesInner) HasAllowedPreviousStatuses() bool`

HasAllowedPreviousStatuses returns a boolean if a field has been set.

### GetRejectedPreviousStatuses

`func (o *PostApiEventsActionsChangeObjectStatusRequestStatusChangesInner) GetRejectedPreviousStatuses() []string`

GetRejectedPreviousStatuses returns the RejectedPreviousStatuses field if non-nil, zero value otherwise.

### GetRejectedPreviousStatusesOk

`func (o *PostApiEventsActionsChangeObjectStatusRequestStatusChangesInner) GetRejectedPreviousStatusesOk() (*[]string, bool)`

GetRejectedPreviousStatusesOk returns a tuple with the RejectedPreviousStatuses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRejectedPreviousStatuses

`func (o *PostApiEventsActionsChangeObjectStatusRequestStatusChangesInner) SetRejectedPreviousStatuses(v []string)`

SetRejectedPreviousStatuses sets RejectedPreviousStatuses field to given value.

### HasRejectedPreviousStatuses

`func (o *PostApiEventsActionsChangeObjectStatusRequestStatusChangesInner) HasRejectedPreviousStatuses() bool`

HasRejectedPreviousStatuses returns a boolean if a field has been set.

### GetChannelKeys

`func (o *PostApiEventsActionsChangeObjectStatusRequestStatusChangesInner) GetChannelKeys() []string`

GetChannelKeys returns the ChannelKeys field if non-nil, zero value otherwise.

### GetChannelKeysOk

`func (o *PostApiEventsActionsChangeObjectStatusRequestStatusChangesInner) GetChannelKeysOk() (*[]string, bool)`

GetChannelKeysOk returns a tuple with the ChannelKeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelKeys

`func (o *PostApiEventsActionsChangeObjectStatusRequestStatusChangesInner) SetChannelKeys(v []string)`

SetChannelKeys sets ChannelKeys field to given value.

### HasChannelKeys

`func (o *PostApiEventsActionsChangeObjectStatusRequestStatusChangesInner) HasChannelKeys() bool`

HasChannelKeys returns a boolean if a field has been set.

### GetIgnoreChannels

`func (o *PostApiEventsActionsChangeObjectStatusRequestStatusChangesInner) GetIgnoreChannels() bool`

GetIgnoreChannels returns the IgnoreChannels field if non-nil, zero value otherwise.

### GetIgnoreChannelsOk

`func (o *PostApiEventsActionsChangeObjectStatusRequestStatusChangesInner) GetIgnoreChannelsOk() (*bool, bool)`

GetIgnoreChannelsOk returns a tuple with the IgnoreChannels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIgnoreChannels

`func (o *PostApiEventsActionsChangeObjectStatusRequestStatusChangesInner) SetIgnoreChannels(v bool)`

SetIgnoreChannels sets IgnoreChannels field to given value.

### HasIgnoreChannels

`func (o *PostApiEventsActionsChangeObjectStatusRequestStatusChangesInner) HasIgnoreChannels() bool`

HasIgnoreChannels returns a boolean if a field has been set.

### GetTicketBuyerId

`func (o *PostApiEventsActionsChangeObjectStatusRequestStatusChangesInner) GetTicketBuyerId() string`

GetTicketBuyerId returns the TicketBuyerId field if non-nil, zero value otherwise.

### GetTicketBuyerIdOk

`func (o *PostApiEventsActionsChangeObjectStatusRequestStatusChangesInner) GetTicketBuyerIdOk() (*string, bool)`

GetTicketBuyerIdOk returns a tuple with the TicketBuyerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTicketBuyerId

`func (o *PostApiEventsActionsChangeObjectStatusRequestStatusChangesInner) SetTicketBuyerId(v string)`

SetTicketBuyerId sets TicketBuyerId field to given value.

### HasTicketBuyerId

`func (o *PostApiEventsActionsChangeObjectStatusRequestStatusChangesInner) HasTicketBuyerId() bool`

HasTicketBuyerId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


