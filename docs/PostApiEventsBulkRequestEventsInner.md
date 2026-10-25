# PostApiEventsBulkRequestEventsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Date** | Pointer to **string** |  | [optional] 
**Key** | Pointer to **string** | Optional partner key, unique per workspace, accepted wherever an event id is (URL-safe, not UUID-shaped). | [optional] 
**TableBookingConfig** | Pointer to [**NullablePostApiEventsRequestTableBookingConfig**](PostApiEventsRequestTableBookingConfig.md) |  | [optional] 
**CategoryOverrides** | Pointer to [**map[string]PostApiEventsRequestCategoryOverridesValue**](PostApiEventsRequestCategoryOverridesValue.md) | Per-event category overrides (same semantics as PATCH /events/{eventId}). | [optional] 
**Channels** | Pointer to [**[]PostApiEventsRequestChannelsInner**](PostApiEventsRequestChannelsInner.md) | Inventory channels created with the event; &#x60;channelKey&#x60; secrets are returned once. | [optional] 
**ForSaleConfig** | Pointer to [**PostApiEventsRequestForSaleConfig**](PostApiEventsRequestForSaleConfig.md) |  | [optional] 

## Methods

### NewPostApiEventsBulkRequestEventsInner

`func NewPostApiEventsBulkRequestEventsInner(name string, ) *PostApiEventsBulkRequestEventsInner`

NewPostApiEventsBulkRequestEventsInner instantiates a new PostApiEventsBulkRequestEventsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostApiEventsBulkRequestEventsInnerWithDefaults

`func NewPostApiEventsBulkRequestEventsInnerWithDefaults() *PostApiEventsBulkRequestEventsInner`

NewPostApiEventsBulkRequestEventsInnerWithDefaults instantiates a new PostApiEventsBulkRequestEventsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *PostApiEventsBulkRequestEventsInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PostApiEventsBulkRequestEventsInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PostApiEventsBulkRequestEventsInner) SetName(v string)`

SetName sets Name field to given value.


### GetDate

`func (o *PostApiEventsBulkRequestEventsInner) GetDate() string`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *PostApiEventsBulkRequestEventsInner) GetDateOk() (*string, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *PostApiEventsBulkRequestEventsInner) SetDate(v string)`

SetDate sets Date field to given value.

### HasDate

`func (o *PostApiEventsBulkRequestEventsInner) HasDate() bool`

HasDate returns a boolean if a field has been set.

### GetKey

`func (o *PostApiEventsBulkRequestEventsInner) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *PostApiEventsBulkRequestEventsInner) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *PostApiEventsBulkRequestEventsInner) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *PostApiEventsBulkRequestEventsInner) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetTableBookingConfig

`func (o *PostApiEventsBulkRequestEventsInner) GetTableBookingConfig() PostApiEventsRequestTableBookingConfig`

GetTableBookingConfig returns the TableBookingConfig field if non-nil, zero value otherwise.

### GetTableBookingConfigOk

`func (o *PostApiEventsBulkRequestEventsInner) GetTableBookingConfigOk() (*PostApiEventsRequestTableBookingConfig, bool)`

GetTableBookingConfigOk returns a tuple with the TableBookingConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTableBookingConfig

`func (o *PostApiEventsBulkRequestEventsInner) SetTableBookingConfig(v PostApiEventsRequestTableBookingConfig)`

SetTableBookingConfig sets TableBookingConfig field to given value.

### HasTableBookingConfig

`func (o *PostApiEventsBulkRequestEventsInner) HasTableBookingConfig() bool`

HasTableBookingConfig returns a boolean if a field has been set.

### SetTableBookingConfigNil

`func (o *PostApiEventsBulkRequestEventsInner) SetTableBookingConfigNil(b bool)`

 SetTableBookingConfigNil sets the value for TableBookingConfig to be an explicit nil

### UnsetTableBookingConfig
`func (o *PostApiEventsBulkRequestEventsInner) UnsetTableBookingConfig()`

UnsetTableBookingConfig ensures that no value is present for TableBookingConfig, not even an explicit nil
### GetCategoryOverrides

`func (o *PostApiEventsBulkRequestEventsInner) GetCategoryOverrides() map[string]PostApiEventsRequestCategoryOverridesValue`

GetCategoryOverrides returns the CategoryOverrides field if non-nil, zero value otherwise.

### GetCategoryOverridesOk

`func (o *PostApiEventsBulkRequestEventsInner) GetCategoryOverridesOk() (*map[string]PostApiEventsRequestCategoryOverridesValue, bool)`

GetCategoryOverridesOk returns a tuple with the CategoryOverrides field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategoryOverrides

`func (o *PostApiEventsBulkRequestEventsInner) SetCategoryOverrides(v map[string]PostApiEventsRequestCategoryOverridesValue)`

SetCategoryOverrides sets CategoryOverrides field to given value.

### HasCategoryOverrides

`func (o *PostApiEventsBulkRequestEventsInner) HasCategoryOverrides() bool`

HasCategoryOverrides returns a boolean if a field has been set.

### SetCategoryOverridesNil

`func (o *PostApiEventsBulkRequestEventsInner) SetCategoryOverridesNil(b bool)`

 SetCategoryOverridesNil sets the value for CategoryOverrides to be an explicit nil

### UnsetCategoryOverrides
`func (o *PostApiEventsBulkRequestEventsInner) UnsetCategoryOverrides()`

UnsetCategoryOverrides ensures that no value is present for CategoryOverrides, not even an explicit nil
### GetChannels

`func (o *PostApiEventsBulkRequestEventsInner) GetChannels() []PostApiEventsRequestChannelsInner`

GetChannels returns the Channels field if non-nil, zero value otherwise.

### GetChannelsOk

`func (o *PostApiEventsBulkRequestEventsInner) GetChannelsOk() (*[]PostApiEventsRequestChannelsInner, bool)`

GetChannelsOk returns a tuple with the Channels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannels

`func (o *PostApiEventsBulkRequestEventsInner) SetChannels(v []PostApiEventsRequestChannelsInner)`

SetChannels sets Channels field to given value.

### HasChannels

`func (o *PostApiEventsBulkRequestEventsInner) HasChannels() bool`

HasChannels returns a boolean if a field has been set.

### GetForSaleConfig

`func (o *PostApiEventsBulkRequestEventsInner) GetForSaleConfig() PostApiEventsRequestForSaleConfig`

GetForSaleConfig returns the ForSaleConfig field if non-nil, zero value otherwise.

### GetForSaleConfigOk

`func (o *PostApiEventsBulkRequestEventsInner) GetForSaleConfigOk() (*PostApiEventsRequestForSaleConfig, bool)`

GetForSaleConfigOk returns a tuple with the ForSaleConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForSaleConfig

`func (o *PostApiEventsBulkRequestEventsInner) SetForSaleConfig(v PostApiEventsRequestForSaleConfig)`

SetForSaleConfig sets ForSaleConfig field to given value.

### HasForSaleConfig

`func (o *PostApiEventsBulkRequestEventsInner) HasForSaleConfig() bool`

HasForSaleConfig returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


