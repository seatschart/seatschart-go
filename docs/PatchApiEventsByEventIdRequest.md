# PatchApiEventsByEventIdRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsInThePast** | Pointer to **bool** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Currency** | Pointer to **string** |  | [optional] 
**Date** | Pointer to **string** |  | [optional] 
**TableBookingConfig** | Pointer to [**NullablePatchApiEventsByEventIdRequestTableBookingConfig**](PatchApiEventsByEventIdRequestTableBookingConfig.md) |  | [optional] 
**CategoryOverrides** | Pointer to [**map[string]PostApiEventsRequestCategoryOverridesValue**](PostApiEventsRequestCategoryOverridesValue.md) |  | [optional] 
**ForSalePerCategory** | Pointer to **map[string]bool** |  | [optional] 
**Key** | Pointer to **NullableString** | Set, rename (&#x60;\&quot;new-key\&quot;&#x60;) or clear (&#x60;null&#x60;) the event key. 409 &#x60;event_key_taken&#x60; when another event of the workspace uses it. | [optional] 
**ChartId** | Pointer to **string** | Switch the event to another published chart of the workspace (see description). | [optional] 
**Categories** | Pointer to [**[]PatchApiEventsByEventIdRequestCategoriesInner**](PatchApiEventsByEventIdRequestCategoriesInner.md) | Event-only categories (replaces the list; null clears). | [optional] 
**ObjectCategories** | Pointer to **map[string]string** | Object label (seat label or GA area name) → category key (event category key, chart category id or name). Replaces the map; null clears. | [optional] 

## Methods

### NewPatchApiEventsByEventIdRequest

`func NewPatchApiEventsByEventIdRequest() *PatchApiEventsByEventIdRequest`

NewPatchApiEventsByEventIdRequest instantiates a new PatchApiEventsByEventIdRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchApiEventsByEventIdRequestWithDefaults

`func NewPatchApiEventsByEventIdRequestWithDefaults() *PatchApiEventsByEventIdRequest`

NewPatchApiEventsByEventIdRequestWithDefaults instantiates a new PatchApiEventsByEventIdRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsInThePast

`func (o *PatchApiEventsByEventIdRequest) GetIsInThePast() bool`

GetIsInThePast returns the IsInThePast field if non-nil, zero value otherwise.

### GetIsInThePastOk

`func (o *PatchApiEventsByEventIdRequest) GetIsInThePastOk() (*bool, bool)`

GetIsInThePastOk returns a tuple with the IsInThePast field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsInThePast

`func (o *PatchApiEventsByEventIdRequest) SetIsInThePast(v bool)`

SetIsInThePast sets IsInThePast field to given value.

### HasIsInThePast

`func (o *PatchApiEventsByEventIdRequest) HasIsInThePast() bool`

HasIsInThePast returns a boolean if a field has been set.

### GetName

`func (o *PatchApiEventsByEventIdRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PatchApiEventsByEventIdRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PatchApiEventsByEventIdRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PatchApiEventsByEventIdRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetCurrency

`func (o *PatchApiEventsByEventIdRequest) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *PatchApiEventsByEventIdRequest) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *PatchApiEventsByEventIdRequest) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *PatchApiEventsByEventIdRequest) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetDate

`func (o *PatchApiEventsByEventIdRequest) GetDate() string`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *PatchApiEventsByEventIdRequest) GetDateOk() (*string, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *PatchApiEventsByEventIdRequest) SetDate(v string)`

SetDate sets Date field to given value.

### HasDate

`func (o *PatchApiEventsByEventIdRequest) HasDate() bool`

HasDate returns a boolean if a field has been set.

### GetTableBookingConfig

`func (o *PatchApiEventsByEventIdRequest) GetTableBookingConfig() PatchApiEventsByEventIdRequestTableBookingConfig`

GetTableBookingConfig returns the TableBookingConfig field if non-nil, zero value otherwise.

### GetTableBookingConfigOk

`func (o *PatchApiEventsByEventIdRequest) GetTableBookingConfigOk() (*PatchApiEventsByEventIdRequestTableBookingConfig, bool)`

GetTableBookingConfigOk returns a tuple with the TableBookingConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTableBookingConfig

`func (o *PatchApiEventsByEventIdRequest) SetTableBookingConfig(v PatchApiEventsByEventIdRequestTableBookingConfig)`

SetTableBookingConfig sets TableBookingConfig field to given value.

### HasTableBookingConfig

`func (o *PatchApiEventsByEventIdRequest) HasTableBookingConfig() bool`

HasTableBookingConfig returns a boolean if a field has been set.

### SetTableBookingConfigNil

`func (o *PatchApiEventsByEventIdRequest) SetTableBookingConfigNil(b bool)`

 SetTableBookingConfigNil sets the value for TableBookingConfig to be an explicit nil

### UnsetTableBookingConfig
`func (o *PatchApiEventsByEventIdRequest) UnsetTableBookingConfig()`

UnsetTableBookingConfig ensures that no value is present for TableBookingConfig, not even an explicit nil
### GetCategoryOverrides

`func (o *PatchApiEventsByEventIdRequest) GetCategoryOverrides() map[string]PostApiEventsRequestCategoryOverridesValue`

GetCategoryOverrides returns the CategoryOverrides field if non-nil, zero value otherwise.

### GetCategoryOverridesOk

`func (o *PatchApiEventsByEventIdRequest) GetCategoryOverridesOk() (*map[string]PostApiEventsRequestCategoryOverridesValue, bool)`

GetCategoryOverridesOk returns a tuple with the CategoryOverrides field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategoryOverrides

`func (o *PatchApiEventsByEventIdRequest) SetCategoryOverrides(v map[string]PostApiEventsRequestCategoryOverridesValue)`

SetCategoryOverrides sets CategoryOverrides field to given value.

### HasCategoryOverrides

`func (o *PatchApiEventsByEventIdRequest) HasCategoryOverrides() bool`

HasCategoryOverrides returns a boolean if a field has been set.

### SetCategoryOverridesNil

`func (o *PatchApiEventsByEventIdRequest) SetCategoryOverridesNil(b bool)`

 SetCategoryOverridesNil sets the value for CategoryOverrides to be an explicit nil

### UnsetCategoryOverrides
`func (o *PatchApiEventsByEventIdRequest) UnsetCategoryOverrides()`

UnsetCategoryOverrides ensures that no value is present for CategoryOverrides, not even an explicit nil
### GetForSalePerCategory

`func (o *PatchApiEventsByEventIdRequest) GetForSalePerCategory() map[string]bool`

GetForSalePerCategory returns the ForSalePerCategory field if non-nil, zero value otherwise.

### GetForSalePerCategoryOk

`func (o *PatchApiEventsByEventIdRequest) GetForSalePerCategoryOk() (*map[string]bool, bool)`

GetForSalePerCategoryOk returns a tuple with the ForSalePerCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForSalePerCategory

`func (o *PatchApiEventsByEventIdRequest) SetForSalePerCategory(v map[string]bool)`

SetForSalePerCategory sets ForSalePerCategory field to given value.

### HasForSalePerCategory

`func (o *PatchApiEventsByEventIdRequest) HasForSalePerCategory() bool`

HasForSalePerCategory returns a boolean if a field has been set.

### GetKey

`func (o *PatchApiEventsByEventIdRequest) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *PatchApiEventsByEventIdRequest) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *PatchApiEventsByEventIdRequest) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *PatchApiEventsByEventIdRequest) HasKey() bool`

HasKey returns a boolean if a field has been set.

### SetKeyNil

`func (o *PatchApiEventsByEventIdRequest) SetKeyNil(b bool)`

 SetKeyNil sets the value for Key to be an explicit nil

### UnsetKey
`func (o *PatchApiEventsByEventIdRequest) UnsetKey()`

UnsetKey ensures that no value is present for Key, not even an explicit nil
### GetChartId

`func (o *PatchApiEventsByEventIdRequest) GetChartId() string`

GetChartId returns the ChartId field if non-nil, zero value otherwise.

### GetChartIdOk

`func (o *PatchApiEventsByEventIdRequest) GetChartIdOk() (*string, bool)`

GetChartIdOk returns a tuple with the ChartId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChartId

`func (o *PatchApiEventsByEventIdRequest) SetChartId(v string)`

SetChartId sets ChartId field to given value.

### HasChartId

`func (o *PatchApiEventsByEventIdRequest) HasChartId() bool`

HasChartId returns a boolean if a field has been set.

### GetCategories

`func (o *PatchApiEventsByEventIdRequest) GetCategories() []PatchApiEventsByEventIdRequestCategoriesInner`

GetCategories returns the Categories field if non-nil, zero value otherwise.

### GetCategoriesOk

`func (o *PatchApiEventsByEventIdRequest) GetCategoriesOk() (*[]PatchApiEventsByEventIdRequestCategoriesInner, bool)`

GetCategoriesOk returns a tuple with the Categories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategories

`func (o *PatchApiEventsByEventIdRequest) SetCategories(v []PatchApiEventsByEventIdRequestCategoriesInner)`

SetCategories sets Categories field to given value.

### HasCategories

`func (o *PatchApiEventsByEventIdRequest) HasCategories() bool`

HasCategories returns a boolean if a field has been set.

### SetCategoriesNil

`func (o *PatchApiEventsByEventIdRequest) SetCategoriesNil(b bool)`

 SetCategoriesNil sets the value for Categories to be an explicit nil

### UnsetCategories
`func (o *PatchApiEventsByEventIdRequest) UnsetCategories()`

UnsetCategories ensures that no value is present for Categories, not even an explicit nil
### GetObjectCategories

`func (o *PatchApiEventsByEventIdRequest) GetObjectCategories() map[string]string`

GetObjectCategories returns the ObjectCategories field if non-nil, zero value otherwise.

### GetObjectCategoriesOk

`func (o *PatchApiEventsByEventIdRequest) GetObjectCategoriesOk() (*map[string]string, bool)`

GetObjectCategoriesOk returns a tuple with the ObjectCategories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectCategories

`func (o *PatchApiEventsByEventIdRequest) SetObjectCategories(v map[string]string)`

SetObjectCategories sets ObjectCategories field to given value.

### HasObjectCategories

`func (o *PatchApiEventsByEventIdRequest) HasObjectCategories() bool`

HasObjectCategories returns a boolean if a field has been set.

### SetObjectCategoriesNil

`func (o *PatchApiEventsByEventIdRequest) SetObjectCategoriesNil(b bool)`

 SetObjectCategoriesNil sets the value for ObjectCategories to be an explicit nil

### UnsetObjectCategories
`func (o *PatchApiEventsByEventIdRequest) UnsetObjectCategories()`

UnsetObjectCategories ensures that no value is present for ObjectCategories, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


