# PostApiEventsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Date** | Pointer to **string** |  | [optional] 
**ChartId** | Pointer to **string** |  | [optional] 
**SeasonId** | Pointer to **string** |  | [optional] 
**Currency** | Pointer to **string** |  | [optional] 
**Key** | Pointer to **string** | Optional partner key, unique per workspace, accepted wherever an event id is (URL-safe, not UUID-shaped). | [optional] 
**TableBookingConfig** | Pointer to [**NullablePostApiEventsRequestTableBookingConfig**](PostApiEventsRequestTableBookingConfig.md) |  | [optional] 
**CategoryOverrides** | Pointer to [**map[string]PostApiEventsRequestCategoryOverridesValue**](PostApiEventsRequestCategoryOverridesValue.md) | Per-event category overrides (same semantics as PATCH /events/{eventId}). | [optional] 
**Channels** | Pointer to [**[]PostApiEventsRequestChannelsInner**](PostApiEventsRequestChannelsInner.md) | Inventory channels created with the event; &#x60;channelKey&#x60; secrets are returned once. | [optional] 
**ForSaleConfig** | Pointer to [**PostApiEventsRequestForSaleConfig**](PostApiEventsRequestForSaleConfig.md) |  | [optional] 

## Methods

### NewPostApiEventsRequest

`func NewPostApiEventsRequest(name string, ) *PostApiEventsRequest`

NewPostApiEventsRequest instantiates a new PostApiEventsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostApiEventsRequestWithDefaults

`func NewPostApiEventsRequestWithDefaults() *PostApiEventsRequest`

NewPostApiEventsRequestWithDefaults instantiates a new PostApiEventsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *PostApiEventsRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PostApiEventsRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PostApiEventsRequest) SetName(v string)`

SetName sets Name field to given value.


### GetDate

`func (o *PostApiEventsRequest) GetDate() string`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *PostApiEventsRequest) GetDateOk() (*string, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *PostApiEventsRequest) SetDate(v string)`

SetDate sets Date field to given value.

### HasDate

`func (o *PostApiEventsRequest) HasDate() bool`

HasDate returns a boolean if a field has been set.

### GetChartId

`func (o *PostApiEventsRequest) GetChartId() string`

GetChartId returns the ChartId field if non-nil, zero value otherwise.

### GetChartIdOk

`func (o *PostApiEventsRequest) GetChartIdOk() (*string, bool)`

GetChartIdOk returns a tuple with the ChartId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChartId

`func (o *PostApiEventsRequest) SetChartId(v string)`

SetChartId sets ChartId field to given value.

### HasChartId

`func (o *PostApiEventsRequest) HasChartId() bool`

HasChartId returns a boolean if a field has been set.

### GetSeasonId

`func (o *PostApiEventsRequest) GetSeasonId() string`

GetSeasonId returns the SeasonId field if non-nil, zero value otherwise.

### GetSeasonIdOk

`func (o *PostApiEventsRequest) GetSeasonIdOk() (*string, bool)`

GetSeasonIdOk returns a tuple with the SeasonId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeasonId

`func (o *PostApiEventsRequest) SetSeasonId(v string)`

SetSeasonId sets SeasonId field to given value.

### HasSeasonId

`func (o *PostApiEventsRequest) HasSeasonId() bool`

HasSeasonId returns a boolean if a field has been set.

### GetCurrency

`func (o *PostApiEventsRequest) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *PostApiEventsRequest) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *PostApiEventsRequest) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *PostApiEventsRequest) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetKey

`func (o *PostApiEventsRequest) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *PostApiEventsRequest) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *PostApiEventsRequest) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *PostApiEventsRequest) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetTableBookingConfig

`func (o *PostApiEventsRequest) GetTableBookingConfig() PostApiEventsRequestTableBookingConfig`

GetTableBookingConfig returns the TableBookingConfig field if non-nil, zero value otherwise.

### GetTableBookingConfigOk

`func (o *PostApiEventsRequest) GetTableBookingConfigOk() (*PostApiEventsRequestTableBookingConfig, bool)`

GetTableBookingConfigOk returns a tuple with the TableBookingConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTableBookingConfig

`func (o *PostApiEventsRequest) SetTableBookingConfig(v PostApiEventsRequestTableBookingConfig)`

SetTableBookingConfig sets TableBookingConfig field to given value.

### HasTableBookingConfig

`func (o *PostApiEventsRequest) HasTableBookingConfig() bool`

HasTableBookingConfig returns a boolean if a field has been set.

### SetTableBookingConfigNil

`func (o *PostApiEventsRequest) SetTableBookingConfigNil(b bool)`

 SetTableBookingConfigNil sets the value for TableBookingConfig to be an explicit nil

### UnsetTableBookingConfig
`func (o *PostApiEventsRequest) UnsetTableBookingConfig()`

UnsetTableBookingConfig ensures that no value is present for TableBookingConfig, not even an explicit nil
### GetCategoryOverrides

`func (o *PostApiEventsRequest) GetCategoryOverrides() map[string]PostApiEventsRequestCategoryOverridesValue`

GetCategoryOverrides returns the CategoryOverrides field if non-nil, zero value otherwise.

### GetCategoryOverridesOk

`func (o *PostApiEventsRequest) GetCategoryOverridesOk() (*map[string]PostApiEventsRequestCategoryOverridesValue, bool)`

GetCategoryOverridesOk returns a tuple with the CategoryOverrides field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategoryOverrides

`func (o *PostApiEventsRequest) SetCategoryOverrides(v map[string]PostApiEventsRequestCategoryOverridesValue)`

SetCategoryOverrides sets CategoryOverrides field to given value.

### HasCategoryOverrides

`func (o *PostApiEventsRequest) HasCategoryOverrides() bool`

HasCategoryOverrides returns a boolean if a field has been set.

### SetCategoryOverridesNil

`func (o *PostApiEventsRequest) SetCategoryOverridesNil(b bool)`

 SetCategoryOverridesNil sets the value for CategoryOverrides to be an explicit nil

### UnsetCategoryOverrides
`func (o *PostApiEventsRequest) UnsetCategoryOverrides()`

UnsetCategoryOverrides ensures that no value is present for CategoryOverrides, not even an explicit nil
### GetChannels

`func (o *PostApiEventsRequest) GetChannels() []PostApiEventsRequestChannelsInner`

GetChannels returns the Channels field if non-nil, zero value otherwise.

### GetChannelsOk

`func (o *PostApiEventsRequest) GetChannelsOk() (*[]PostApiEventsRequestChannelsInner, bool)`

GetChannelsOk returns a tuple with the Channels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannels

`func (o *PostApiEventsRequest) SetChannels(v []PostApiEventsRequestChannelsInner)`

SetChannels sets Channels field to given value.

### HasChannels

`func (o *PostApiEventsRequest) HasChannels() bool`

HasChannels returns a boolean if a field has been set.

### GetForSaleConfig

`func (o *PostApiEventsRequest) GetForSaleConfig() PostApiEventsRequestForSaleConfig`

GetForSaleConfig returns the ForSaleConfig field if non-nil, zero value otherwise.

### GetForSaleConfigOk

`func (o *PostApiEventsRequest) GetForSaleConfigOk() (*PostApiEventsRequestForSaleConfig, bool)`

GetForSaleConfigOk returns a tuple with the ForSaleConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForSaleConfig

`func (o *PostApiEventsRequest) SetForSaleConfig(v PostApiEventsRequestForSaleConfig)`

SetForSaleConfig sets ForSaleConfig field to given value.

### HasForSaleConfig

`func (o *PostApiEventsRequest) HasForSaleConfig() bool`

HasForSaleConfig returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


