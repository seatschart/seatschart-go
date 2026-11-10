# PostApiSeatsBestAvailableRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EventId** | **string** |  | 
**Quantity** | Pointer to **int32** |  | [optional] 
**Category** | Pointer to **NullableString** |  | [optional] 
**Categories** | Pointer to **[]string** |  | [optional] 
**TicketTypes** | Pointer to **map[string]int32** |  | [optional] 
**SectionId** | Pointer to **NullableString** |  | [optional] 
**Sections** | Pointer to **[]string** | Layout section ids or names; the candidate pool is the union of their seats (XOR with sectionId). | [optional] 
**ChannelKeys** | Pointer to **[]string** | API keys only: channel secret keys or channel ids whose seats may be picked (public seats always can). | [optional] 
**IgnoreChannels** | Pointer to **bool** | API keys only: pick from every channel. | [optional] 
**OrderId** | Pointer to **string** | Requires hold: partner order reference stored on the held seats (extraData.orderId) and echoed back. | [optional] 
**Zone** | Pointer to **NullableString** | Layout zone id or name. | [optional] 
**Channel** | Pointer to **string** |  | [optional] 
**PreferAdjacent** | Pointer to **bool** |  | [optional] 
**TryToPreventOrphanSeats** | Pointer to **bool** |  | [optional] 
**AccessibleSeats** | Pointer to **int32** |  | [optional] 
**Hold** | Pointer to **bool** |  | [optional] 
**HoldToken** | Pointer to **string** |  | [optional] 
**ExtraData** | Pointer to **map[string]interface{}** |  | [optional] 
**UserId** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** | &#x60;booked&#x60;: book the best window directly (API keys with the bookings scope / dashboard); &#x60;reservedByToken&#x60;: same as hold: true. | [optional] 
**KeepExtraData** | Pointer to **bool** | Keep the seats&#39; stored extraData; extraData / orderId are merged on top. | [optional] 

## Methods

### NewPostApiSeatsBestAvailableRequest

`func NewPostApiSeatsBestAvailableRequest(eventId string, ) *PostApiSeatsBestAvailableRequest`

NewPostApiSeatsBestAvailableRequest instantiates a new PostApiSeatsBestAvailableRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostApiSeatsBestAvailableRequestWithDefaults

`func NewPostApiSeatsBestAvailableRequestWithDefaults() *PostApiSeatsBestAvailableRequest`

NewPostApiSeatsBestAvailableRequestWithDefaults instantiates a new PostApiSeatsBestAvailableRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEventId

`func (o *PostApiSeatsBestAvailableRequest) GetEventId() string`

GetEventId returns the EventId field if non-nil, zero value otherwise.

### GetEventIdOk

`func (o *PostApiSeatsBestAvailableRequest) GetEventIdOk() (*string, bool)`

GetEventIdOk returns a tuple with the EventId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventId

`func (o *PostApiSeatsBestAvailableRequest) SetEventId(v string)`

SetEventId sets EventId field to given value.


### GetQuantity

`func (o *PostApiSeatsBestAvailableRequest) GetQuantity() int32`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *PostApiSeatsBestAvailableRequest) GetQuantityOk() (*int32, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *PostApiSeatsBestAvailableRequest) SetQuantity(v int32)`

SetQuantity sets Quantity field to given value.

### HasQuantity

`func (o *PostApiSeatsBestAvailableRequest) HasQuantity() bool`

HasQuantity returns a boolean if a field has been set.

### GetCategory

`func (o *PostApiSeatsBestAvailableRequest) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *PostApiSeatsBestAvailableRequest) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *PostApiSeatsBestAvailableRequest) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *PostApiSeatsBestAvailableRequest) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### SetCategoryNil

`func (o *PostApiSeatsBestAvailableRequest) SetCategoryNil(b bool)`

 SetCategoryNil sets the value for Category to be an explicit nil

### UnsetCategory
`func (o *PostApiSeatsBestAvailableRequest) UnsetCategory()`

UnsetCategory ensures that no value is present for Category, not even an explicit nil
### GetCategories

`func (o *PostApiSeatsBestAvailableRequest) GetCategories() []string`

GetCategories returns the Categories field if non-nil, zero value otherwise.

### GetCategoriesOk

`func (o *PostApiSeatsBestAvailableRequest) GetCategoriesOk() (*[]string, bool)`

GetCategoriesOk returns a tuple with the Categories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategories

`func (o *PostApiSeatsBestAvailableRequest) SetCategories(v []string)`

SetCategories sets Categories field to given value.

### HasCategories

`func (o *PostApiSeatsBestAvailableRequest) HasCategories() bool`

HasCategories returns a boolean if a field has been set.

### GetTicketTypes

`func (o *PostApiSeatsBestAvailableRequest) GetTicketTypes() map[string]int32`

GetTicketTypes returns the TicketTypes field if non-nil, zero value otherwise.

### GetTicketTypesOk

`func (o *PostApiSeatsBestAvailableRequest) GetTicketTypesOk() (*map[string]int32, bool)`

GetTicketTypesOk returns a tuple with the TicketTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTicketTypes

`func (o *PostApiSeatsBestAvailableRequest) SetTicketTypes(v map[string]int32)`

SetTicketTypes sets TicketTypes field to given value.

### HasTicketTypes

`func (o *PostApiSeatsBestAvailableRequest) HasTicketTypes() bool`

HasTicketTypes returns a boolean if a field has been set.

### GetSectionId

`func (o *PostApiSeatsBestAvailableRequest) GetSectionId() string`

GetSectionId returns the SectionId field if non-nil, zero value otherwise.

### GetSectionIdOk

`func (o *PostApiSeatsBestAvailableRequest) GetSectionIdOk() (*string, bool)`

GetSectionIdOk returns a tuple with the SectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSectionId

`func (o *PostApiSeatsBestAvailableRequest) SetSectionId(v string)`

SetSectionId sets SectionId field to given value.

### HasSectionId

`func (o *PostApiSeatsBestAvailableRequest) HasSectionId() bool`

HasSectionId returns a boolean if a field has been set.

### SetSectionIdNil

`func (o *PostApiSeatsBestAvailableRequest) SetSectionIdNil(b bool)`

 SetSectionIdNil sets the value for SectionId to be an explicit nil

### UnsetSectionId
`func (o *PostApiSeatsBestAvailableRequest) UnsetSectionId()`

UnsetSectionId ensures that no value is present for SectionId, not even an explicit nil
### GetSections

`func (o *PostApiSeatsBestAvailableRequest) GetSections() []string`

GetSections returns the Sections field if non-nil, zero value otherwise.

### GetSectionsOk

`func (o *PostApiSeatsBestAvailableRequest) GetSectionsOk() (*[]string, bool)`

GetSectionsOk returns a tuple with the Sections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSections

`func (o *PostApiSeatsBestAvailableRequest) SetSections(v []string)`

SetSections sets Sections field to given value.

### HasSections

`func (o *PostApiSeatsBestAvailableRequest) HasSections() bool`

HasSections returns a boolean if a field has been set.

### GetChannelKeys

`func (o *PostApiSeatsBestAvailableRequest) GetChannelKeys() []string`

GetChannelKeys returns the ChannelKeys field if non-nil, zero value otherwise.

### GetChannelKeysOk

`func (o *PostApiSeatsBestAvailableRequest) GetChannelKeysOk() (*[]string, bool)`

GetChannelKeysOk returns a tuple with the ChannelKeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelKeys

`func (o *PostApiSeatsBestAvailableRequest) SetChannelKeys(v []string)`

SetChannelKeys sets ChannelKeys field to given value.

### HasChannelKeys

`func (o *PostApiSeatsBestAvailableRequest) HasChannelKeys() bool`

HasChannelKeys returns a boolean if a field has been set.

### GetIgnoreChannels

`func (o *PostApiSeatsBestAvailableRequest) GetIgnoreChannels() bool`

GetIgnoreChannels returns the IgnoreChannels field if non-nil, zero value otherwise.

### GetIgnoreChannelsOk

`func (o *PostApiSeatsBestAvailableRequest) GetIgnoreChannelsOk() (*bool, bool)`

GetIgnoreChannelsOk returns a tuple with the IgnoreChannels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIgnoreChannels

`func (o *PostApiSeatsBestAvailableRequest) SetIgnoreChannels(v bool)`

SetIgnoreChannels sets IgnoreChannels field to given value.

### HasIgnoreChannels

`func (o *PostApiSeatsBestAvailableRequest) HasIgnoreChannels() bool`

HasIgnoreChannels returns a boolean if a field has been set.

### GetOrderId

`func (o *PostApiSeatsBestAvailableRequest) GetOrderId() string`

GetOrderId returns the OrderId field if non-nil, zero value otherwise.

### GetOrderIdOk

`func (o *PostApiSeatsBestAvailableRequest) GetOrderIdOk() (*string, bool)`

GetOrderIdOk returns a tuple with the OrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderId

`func (o *PostApiSeatsBestAvailableRequest) SetOrderId(v string)`

SetOrderId sets OrderId field to given value.

### HasOrderId

`func (o *PostApiSeatsBestAvailableRequest) HasOrderId() bool`

HasOrderId returns a boolean if a field has been set.

### GetZone

`func (o *PostApiSeatsBestAvailableRequest) GetZone() string`

GetZone returns the Zone field if non-nil, zero value otherwise.

### GetZoneOk

`func (o *PostApiSeatsBestAvailableRequest) GetZoneOk() (*string, bool)`

GetZoneOk returns a tuple with the Zone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZone

`func (o *PostApiSeatsBestAvailableRequest) SetZone(v string)`

SetZone sets Zone field to given value.

### HasZone

`func (o *PostApiSeatsBestAvailableRequest) HasZone() bool`

HasZone returns a boolean if a field has been set.

### SetZoneNil

`func (o *PostApiSeatsBestAvailableRequest) SetZoneNil(b bool)`

 SetZoneNil sets the value for Zone to be an explicit nil

### UnsetZone
`func (o *PostApiSeatsBestAvailableRequest) UnsetZone()`

UnsetZone ensures that no value is present for Zone, not even an explicit nil
### GetChannel

`func (o *PostApiSeatsBestAvailableRequest) GetChannel() string`

GetChannel returns the Channel field if non-nil, zero value otherwise.

### GetChannelOk

`func (o *PostApiSeatsBestAvailableRequest) GetChannelOk() (*string, bool)`

GetChannelOk returns a tuple with the Channel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannel

`func (o *PostApiSeatsBestAvailableRequest) SetChannel(v string)`

SetChannel sets Channel field to given value.

### HasChannel

`func (o *PostApiSeatsBestAvailableRequest) HasChannel() bool`

HasChannel returns a boolean if a field has been set.

### GetPreferAdjacent

`func (o *PostApiSeatsBestAvailableRequest) GetPreferAdjacent() bool`

GetPreferAdjacent returns the PreferAdjacent field if non-nil, zero value otherwise.

### GetPreferAdjacentOk

`func (o *PostApiSeatsBestAvailableRequest) GetPreferAdjacentOk() (*bool, bool)`

GetPreferAdjacentOk returns a tuple with the PreferAdjacent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreferAdjacent

`func (o *PostApiSeatsBestAvailableRequest) SetPreferAdjacent(v bool)`

SetPreferAdjacent sets PreferAdjacent field to given value.

### HasPreferAdjacent

`func (o *PostApiSeatsBestAvailableRequest) HasPreferAdjacent() bool`

HasPreferAdjacent returns a boolean if a field has been set.

### GetTryToPreventOrphanSeats

`func (o *PostApiSeatsBestAvailableRequest) GetTryToPreventOrphanSeats() bool`

GetTryToPreventOrphanSeats returns the TryToPreventOrphanSeats field if non-nil, zero value otherwise.

### GetTryToPreventOrphanSeatsOk

`func (o *PostApiSeatsBestAvailableRequest) GetTryToPreventOrphanSeatsOk() (*bool, bool)`

GetTryToPreventOrphanSeatsOk returns a tuple with the TryToPreventOrphanSeats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTryToPreventOrphanSeats

`func (o *PostApiSeatsBestAvailableRequest) SetTryToPreventOrphanSeats(v bool)`

SetTryToPreventOrphanSeats sets TryToPreventOrphanSeats field to given value.

### HasTryToPreventOrphanSeats

`func (o *PostApiSeatsBestAvailableRequest) HasTryToPreventOrphanSeats() bool`

HasTryToPreventOrphanSeats returns a boolean if a field has been set.

### GetAccessibleSeats

`func (o *PostApiSeatsBestAvailableRequest) GetAccessibleSeats() int32`

GetAccessibleSeats returns the AccessibleSeats field if non-nil, zero value otherwise.

### GetAccessibleSeatsOk

`func (o *PostApiSeatsBestAvailableRequest) GetAccessibleSeatsOk() (*int32, bool)`

GetAccessibleSeatsOk returns a tuple with the AccessibleSeats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessibleSeats

`func (o *PostApiSeatsBestAvailableRequest) SetAccessibleSeats(v int32)`

SetAccessibleSeats sets AccessibleSeats field to given value.

### HasAccessibleSeats

`func (o *PostApiSeatsBestAvailableRequest) HasAccessibleSeats() bool`

HasAccessibleSeats returns a boolean if a field has been set.

### GetHold

`func (o *PostApiSeatsBestAvailableRequest) GetHold() bool`

GetHold returns the Hold field if non-nil, zero value otherwise.

### GetHoldOk

`func (o *PostApiSeatsBestAvailableRequest) GetHoldOk() (*bool, bool)`

GetHoldOk returns a tuple with the Hold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHold

`func (o *PostApiSeatsBestAvailableRequest) SetHold(v bool)`

SetHold sets Hold field to given value.

### HasHold

`func (o *PostApiSeatsBestAvailableRequest) HasHold() bool`

HasHold returns a boolean if a field has been set.

### GetHoldToken

`func (o *PostApiSeatsBestAvailableRequest) GetHoldToken() string`

GetHoldToken returns the HoldToken field if non-nil, zero value otherwise.

### GetHoldTokenOk

`func (o *PostApiSeatsBestAvailableRequest) GetHoldTokenOk() (*string, bool)`

GetHoldTokenOk returns a tuple with the HoldToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHoldToken

`func (o *PostApiSeatsBestAvailableRequest) SetHoldToken(v string)`

SetHoldToken sets HoldToken field to given value.

### HasHoldToken

`func (o *PostApiSeatsBestAvailableRequest) HasHoldToken() bool`

HasHoldToken returns a boolean if a field has been set.

### GetExtraData

`func (o *PostApiSeatsBestAvailableRequest) GetExtraData() map[string]interface{}`

GetExtraData returns the ExtraData field if non-nil, zero value otherwise.

### GetExtraDataOk

`func (o *PostApiSeatsBestAvailableRequest) GetExtraDataOk() (*map[string]interface{}, bool)`

GetExtraDataOk returns a tuple with the ExtraData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtraData

`func (o *PostApiSeatsBestAvailableRequest) SetExtraData(v map[string]interface{})`

SetExtraData sets ExtraData field to given value.

### HasExtraData

`func (o *PostApiSeatsBestAvailableRequest) HasExtraData() bool`

HasExtraData returns a boolean if a field has been set.

### GetUserId

`func (o *PostApiSeatsBestAvailableRequest) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *PostApiSeatsBestAvailableRequest) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *PostApiSeatsBestAvailableRequest) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *PostApiSeatsBestAvailableRequest) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### GetStatus

`func (o *PostApiSeatsBestAvailableRequest) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *PostApiSeatsBestAvailableRequest) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *PostApiSeatsBestAvailableRequest) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *PostApiSeatsBestAvailableRequest) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetKeepExtraData

`func (o *PostApiSeatsBestAvailableRequest) GetKeepExtraData() bool`

GetKeepExtraData returns the KeepExtraData field if non-nil, zero value otherwise.

### GetKeepExtraDataOk

`func (o *PostApiSeatsBestAvailableRequest) GetKeepExtraDataOk() (*bool, bool)`

GetKeepExtraDataOk returns a tuple with the KeepExtraData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeepExtraData

`func (o *PostApiSeatsBestAvailableRequest) SetKeepExtraData(v bool)`

SetKeepExtraData sets KeepExtraData field to given value.

### HasKeepExtraData

`func (o *PostApiSeatsBestAvailableRequest) HasKeepExtraData() bool`

HasKeepExtraData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


