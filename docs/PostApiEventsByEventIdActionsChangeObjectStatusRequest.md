# PostApiEventsByEventIdActionsChangeObjectStatusRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Objects** | [**[]PostApiEventsActionsChangeObjectStatusRequestStatusChangesInnerObjectsInner**](PostApiEventsActionsChangeObjectStatusRequestStatusChangesInnerObjectsInner.md) |  | 
**Status** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**OrderId** | Pointer to **string** |  | [optional] 
**KeepExtraData** | Pointer to **bool** |  | [optional] 
**SeasonId** | Pointer to **string** |  | [optional] 
**ResalePriceCents** | Pointer to **int32** |  | [optional] 
**ResaleListingId** | Pointer to **string** | Resale listing: stored on extraData.resaleListingId of every object; the renderer can select a listing as one group (resaleListings.selectAsGroup). | [optional] 
**HoldToken** | Pointer to **string** |  | [optional] 
**AllowedPreviousStatuses** | Pointer to **[]string** |  | [optional] 
**RejectedPreviousStatuses** | Pointer to **[]string** |  | [optional] 
**ChannelKeys** | Pointer to **[]string** |  | [optional] 
**IgnoreChannels** | Pointer to **bool** |  | [optional] 
**Events** | Pointer to **[]string** |  | [optional] 
**Atomic** | Pointer to **bool** | With &#x60;events&#x60;: all-or-nothing across the events (one SQL transaction). Default false (each event on its own). | [optional] 
**UserId** | Pointer to **string** | GA holds only: customer identity of the hold (POST /ga/confirm must send the same userId). | [optional] 
**HoldExpiresInMinutes** | Pointer to **int32** | GA holds only: hold lifetime (default: seat hold TTL). | [optional] 
**TicketBuyerId** | Pointer to **string** | Optional opaque ticket buyer id from your system (≤ 256 printable characters). Persisted on holds, orders, reservations, GA receipts, object statuses and status history; never interpreted as an internal user id. Query with GET /events/{eventId}/objects?ticketBuyerId&#x3D;…. | [optional] 

## Methods

### NewPostApiEventsByEventIdActionsChangeObjectStatusRequest

`func NewPostApiEventsByEventIdActionsChangeObjectStatusRequest(objects []PostApiEventsActionsChangeObjectStatusRequestStatusChangesInnerObjectsInner, ) *PostApiEventsByEventIdActionsChangeObjectStatusRequest`

NewPostApiEventsByEventIdActionsChangeObjectStatusRequest instantiates a new PostApiEventsByEventIdActionsChangeObjectStatusRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostApiEventsByEventIdActionsChangeObjectStatusRequestWithDefaults

`func NewPostApiEventsByEventIdActionsChangeObjectStatusRequestWithDefaults() *PostApiEventsByEventIdActionsChangeObjectStatusRequest`

NewPostApiEventsByEventIdActionsChangeObjectStatusRequestWithDefaults instantiates a new PostApiEventsByEventIdActionsChangeObjectStatusRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetObjects

`func (o *PostApiEventsByEventIdActionsChangeObjectStatusRequest) GetObjects() []PostApiEventsActionsChangeObjectStatusRequestStatusChangesInnerObjectsInner`

GetObjects returns the Objects field if non-nil, zero value otherwise.

### GetObjectsOk

`func (o *PostApiEventsByEventIdActionsChangeObjectStatusRequest) GetObjectsOk() (*[]PostApiEventsActionsChangeObjectStatusRequestStatusChangesInnerObjectsInner, bool)`

GetObjectsOk returns a tuple with the Objects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjects

`func (o *PostApiEventsByEventIdActionsChangeObjectStatusRequest) SetObjects(v []PostApiEventsActionsChangeObjectStatusRequestStatusChangesInnerObjectsInner)`

SetObjects sets Objects field to given value.


### GetStatus

`func (o *PostApiEventsByEventIdActionsChangeObjectStatusRequest) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *PostApiEventsByEventIdActionsChangeObjectStatusRequest) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *PostApiEventsByEventIdActionsChangeObjectStatusRequest) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *PostApiEventsByEventIdActionsChangeObjectStatusRequest) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetType

`func (o *PostApiEventsByEventIdActionsChangeObjectStatusRequest) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PostApiEventsByEventIdActionsChangeObjectStatusRequest) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PostApiEventsByEventIdActionsChangeObjectStatusRequest) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *PostApiEventsByEventIdActionsChangeObjectStatusRequest) HasType() bool`

HasType returns a boolean if a field has been set.

### GetOrderId

`func (o *PostApiEventsByEventIdActionsChangeObjectStatusRequest) GetOrderId() string`

GetOrderId returns the OrderId field if non-nil, zero value otherwise.

### GetOrderIdOk

`func (o *PostApiEventsByEventIdActionsChangeObjectStatusRequest) GetOrderIdOk() (*string, bool)`

GetOrderIdOk returns a tuple with the OrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderId

`func (o *PostApiEventsByEventIdActionsChangeObjectStatusRequest) SetOrderId(v string)`

SetOrderId sets OrderId field to given value.

### HasOrderId

`func (o *PostApiEventsByEventIdActionsChangeObjectStatusRequest) HasOrderId() bool`

HasOrderId returns a boolean if a field has been set.

### GetKeepExtraData

`func (o *PostApiEventsByEventIdActionsChangeObjectStatusRequest) GetKeepExtraData() bool`

GetKeepExtraData returns the KeepExtraData field if non-nil, zero value otherwise.

### GetKeepExtraDataOk

`func (o *PostApiEventsByEventIdActionsChangeObjectStatusRequest) GetKeepExtraDataOk() (*bool, bool)`

GetKeepExtraDataOk returns a tuple with the KeepExtraData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeepExtraData

`func (o *PostApiEventsByEventIdActionsChangeObjectStatusRequest) SetKeepExtraData(v bool)`

SetKeepExtraData sets KeepExtraData field to given value.

### HasKeepExtraData

`func (o *PostApiEventsByEventIdActionsChangeObjectStatusRequest) HasKeepExtraData() bool`

HasKeepExtraData returns a boolean if a field has been set.

### GetSeasonId

`func (o *PostApiEventsByEventIdActionsChangeObjectStatusRequest) GetSeasonId() string`

GetSeasonId returns the SeasonId field if non-nil, zero value otherwise.

### GetSeasonIdOk

`func (o *PostApiEventsByEventIdActionsChangeObjectStatusRequest) GetSeasonIdOk() (*string, bool)`

GetSeasonIdOk returns a tuple with the SeasonId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeasonId

`func (o *PostApiEventsByEventIdActionsChangeObjectStatusRequest) SetSeasonId(v string)`

SetSeasonId sets SeasonId field to given value.

### HasSeasonId

`func (o *PostApiEventsByEventIdActionsChangeObjectStatusRequest) HasSeasonId() bool`

HasSeasonId returns a boolean if a field has been set.

### GetResalePriceCents

`func (o *PostApiEventsByEventIdActionsChangeObjectStatusRequest) GetResalePriceCents() int32`

GetResalePriceCents returns the ResalePriceCents field if non-nil, zero value otherwise.

### GetResalePriceCentsOk

`func (o *PostApiEventsByEventIdActionsChangeObjectStatusRequest) GetResalePriceCentsOk() (*int32, bool)`

GetResalePriceCentsOk returns a tuple with the ResalePriceCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResalePriceCents

`func (o *PostApiEventsByEventIdActionsChangeObjectStatusRequest) SetResalePriceCents(v int32)`

SetResalePriceCents sets ResalePriceCents field to given value.

### HasResalePriceCents

`func (o *PostApiEventsByEventIdActionsChangeObjectStatusRequest) HasResalePriceCents() bool`

HasResalePriceCents returns a boolean if a field has been set.

### GetResaleListingId

`func (o *PostApiEventsByEventIdActionsChangeObjectStatusRequest) GetResaleListingId() string`

GetResaleListingId returns the ResaleListingId field if non-nil, zero value otherwise.

### GetResaleListingIdOk

`func (o *PostApiEventsByEventIdActionsChangeObjectStatusRequest) GetResaleListingIdOk() (*string, bool)`

GetResaleListingIdOk returns a tuple with the ResaleListingId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResaleListingId

`func (o *PostApiEventsByEventIdActionsChangeObjectStatusRequest) SetResaleListingId(v string)`

SetResaleListingId sets ResaleListingId field to given value.

### HasResaleListingId

`func (o *PostApiEventsByEventIdActionsChangeObjectStatusRequest) HasResaleListingId() bool`

HasResaleListingId returns a boolean if a field has been set.

### GetHoldToken

`func (o *PostApiEventsByEventIdActionsChangeObjectStatusRequest) GetHoldToken() string`

GetHoldToken returns the HoldToken field if non-nil, zero value otherwise.

### GetHoldTokenOk

`func (o *PostApiEventsByEventIdActionsChangeObjectStatusRequest) GetHoldTokenOk() (*string, bool)`

GetHoldTokenOk returns a tuple with the HoldToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHoldToken

`func (o *PostApiEventsByEventIdActionsChangeObjectStatusRequest) SetHoldToken(v string)`

SetHoldToken sets HoldToken field to given value.

### HasHoldToken

`func (o *PostApiEventsByEventIdActionsChangeObjectStatusRequest) HasHoldToken() bool`

HasHoldToken returns a boolean if a field has been set.

### GetAllowedPreviousStatuses

`func (o *PostApiEventsByEventIdActionsChangeObjectStatusRequest) GetAllowedPreviousStatuses() []string`

GetAllowedPreviousStatuses returns the AllowedPreviousStatuses field if non-nil, zero value otherwise.

### GetAllowedPreviousStatusesOk

`func (o *PostApiEventsByEventIdActionsChangeObjectStatusRequest) GetAllowedPreviousStatusesOk() (*[]string, bool)`

GetAllowedPreviousStatusesOk returns a tuple with the AllowedPreviousStatuses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedPreviousStatuses

`func (o *PostApiEventsByEventIdActionsChangeObjectStatusRequest) SetAllowedPreviousStatuses(v []string)`

SetAllowedPreviousStatuses sets AllowedPreviousStatuses field to given value.

### HasAllowedPreviousStatuses

`func (o *PostApiEventsByEventIdActionsChangeObjectStatusRequest) HasAllowedPreviousStatuses() bool`

HasAllowedPreviousStatuses returns a boolean if a field has been set.

### GetRejectedPreviousStatuses

`func (o *PostApiEventsByEventIdActionsChangeObjectStatusRequest) GetRejectedPreviousStatuses() []string`

GetRejectedPreviousStatuses returns the RejectedPreviousStatuses field if non-nil, zero value otherwise.

### GetRejectedPreviousStatusesOk

`func (o *PostApiEventsByEventIdActionsChangeObjectStatusRequest) GetRejectedPreviousStatusesOk() (*[]string, bool)`

GetRejectedPreviousStatusesOk returns a tuple with the RejectedPreviousStatuses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRejectedPreviousStatuses

`func (o *PostApiEventsByEventIdActionsChangeObjectStatusRequest) SetRejectedPreviousStatuses(v []string)`

SetRejectedPreviousStatuses sets RejectedPreviousStatuses field to given value.

### HasRejectedPreviousStatuses

`func (o *PostApiEventsByEventIdActionsChangeObjectStatusRequest) HasRejectedPreviousStatuses() bool`

HasRejectedPreviousStatuses returns a boolean if a field has been set.

### GetChannelKeys

`func (o *PostApiEventsByEventIdActionsChangeObjectStatusRequest) GetChannelKeys() []string`

GetChannelKeys returns the ChannelKeys field if non-nil, zero value otherwise.

### GetChannelKeysOk

`func (o *PostApiEventsByEventIdActionsChangeObjectStatusRequest) GetChannelKeysOk() (*[]string, bool)`

GetChannelKeysOk returns a tuple with the ChannelKeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelKeys

`func (o *PostApiEventsByEventIdActionsChangeObjectStatusRequest) SetChannelKeys(v []string)`

SetChannelKeys sets ChannelKeys field to given value.

### HasChannelKeys

`func (o *PostApiEventsByEventIdActionsChangeObjectStatusRequest) HasChannelKeys() bool`

HasChannelKeys returns a boolean if a field has been set.

### GetIgnoreChannels

`func (o *PostApiEventsByEventIdActionsChangeObjectStatusRequest) GetIgnoreChannels() bool`

GetIgnoreChannels returns the IgnoreChannels field if non-nil, zero value otherwise.

### GetIgnoreChannelsOk

`func (o *PostApiEventsByEventIdActionsChangeObjectStatusRequest) GetIgnoreChannelsOk() (*bool, bool)`

GetIgnoreChannelsOk returns a tuple with the IgnoreChannels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIgnoreChannels

`func (o *PostApiEventsByEventIdActionsChangeObjectStatusRequest) SetIgnoreChannels(v bool)`

SetIgnoreChannels sets IgnoreChannels field to given value.

### HasIgnoreChannels

`func (o *PostApiEventsByEventIdActionsChangeObjectStatusRequest) HasIgnoreChannels() bool`

HasIgnoreChannels returns a boolean if a field has been set.

### GetEvents

`func (o *PostApiEventsByEventIdActionsChangeObjectStatusRequest) GetEvents() []string`

GetEvents returns the Events field if non-nil, zero value otherwise.

### GetEventsOk

`func (o *PostApiEventsByEventIdActionsChangeObjectStatusRequest) GetEventsOk() (*[]string, bool)`

GetEventsOk returns a tuple with the Events field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvents

`func (o *PostApiEventsByEventIdActionsChangeObjectStatusRequest) SetEvents(v []string)`

SetEvents sets Events field to given value.

### HasEvents

`func (o *PostApiEventsByEventIdActionsChangeObjectStatusRequest) HasEvents() bool`

HasEvents returns a boolean if a field has been set.

### GetAtomic

`func (o *PostApiEventsByEventIdActionsChangeObjectStatusRequest) GetAtomic() bool`

GetAtomic returns the Atomic field if non-nil, zero value otherwise.

### GetAtomicOk

`func (o *PostApiEventsByEventIdActionsChangeObjectStatusRequest) GetAtomicOk() (*bool, bool)`

GetAtomicOk returns a tuple with the Atomic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAtomic

`func (o *PostApiEventsByEventIdActionsChangeObjectStatusRequest) SetAtomic(v bool)`

SetAtomic sets Atomic field to given value.

### HasAtomic

`func (o *PostApiEventsByEventIdActionsChangeObjectStatusRequest) HasAtomic() bool`

HasAtomic returns a boolean if a field has been set.

### GetUserId

`func (o *PostApiEventsByEventIdActionsChangeObjectStatusRequest) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *PostApiEventsByEventIdActionsChangeObjectStatusRequest) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *PostApiEventsByEventIdActionsChangeObjectStatusRequest) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *PostApiEventsByEventIdActionsChangeObjectStatusRequest) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### GetHoldExpiresInMinutes

`func (o *PostApiEventsByEventIdActionsChangeObjectStatusRequest) GetHoldExpiresInMinutes() int32`

GetHoldExpiresInMinutes returns the HoldExpiresInMinutes field if non-nil, zero value otherwise.

### GetHoldExpiresInMinutesOk

`func (o *PostApiEventsByEventIdActionsChangeObjectStatusRequest) GetHoldExpiresInMinutesOk() (*int32, bool)`

GetHoldExpiresInMinutesOk returns a tuple with the HoldExpiresInMinutes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHoldExpiresInMinutes

`func (o *PostApiEventsByEventIdActionsChangeObjectStatusRequest) SetHoldExpiresInMinutes(v int32)`

SetHoldExpiresInMinutes sets HoldExpiresInMinutes field to given value.

### HasHoldExpiresInMinutes

`func (o *PostApiEventsByEventIdActionsChangeObjectStatusRequest) HasHoldExpiresInMinutes() bool`

HasHoldExpiresInMinutes returns a boolean if a field has been set.

### GetTicketBuyerId

`func (o *PostApiEventsByEventIdActionsChangeObjectStatusRequest) GetTicketBuyerId() string`

GetTicketBuyerId returns the TicketBuyerId field if non-nil, zero value otherwise.

### GetTicketBuyerIdOk

`func (o *PostApiEventsByEventIdActionsChangeObjectStatusRequest) GetTicketBuyerIdOk() (*string, bool)`

GetTicketBuyerIdOk returns a tuple with the TicketBuyerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTicketBuyerId

`func (o *PostApiEventsByEventIdActionsChangeObjectStatusRequest) SetTicketBuyerId(v string)`

SetTicketBuyerId sets TicketBuyerId field to given value.

### HasTicketBuyerId

`func (o *PostApiEventsByEventIdActionsChangeObjectStatusRequest) HasTicketBuyerId() bool`

HasTicketBuyerId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


