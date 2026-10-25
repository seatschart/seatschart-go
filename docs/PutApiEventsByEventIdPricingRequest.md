# PutApiEventsByEventIdPricingRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Currency** | **string** |  | 
**AllFeesIncluded** | Pointer to **bool** |  | [optional] 
**Categories** | [**[]PutApiEventsByEventIdPricingRequestCategoriesInner**](PutApiEventsByEventIdPricingRequestCategoriesInner.md) |  | 

## Methods

### NewPutApiEventsByEventIdPricingRequest

`func NewPutApiEventsByEventIdPricingRequest(currency string, categories []PutApiEventsByEventIdPricingRequestCategoriesInner, ) *PutApiEventsByEventIdPricingRequest`

NewPutApiEventsByEventIdPricingRequest instantiates a new PutApiEventsByEventIdPricingRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPutApiEventsByEventIdPricingRequestWithDefaults

`func NewPutApiEventsByEventIdPricingRequestWithDefaults() *PutApiEventsByEventIdPricingRequest`

NewPutApiEventsByEventIdPricingRequestWithDefaults instantiates a new PutApiEventsByEventIdPricingRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCurrency

`func (o *PutApiEventsByEventIdPricingRequest) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *PutApiEventsByEventIdPricingRequest) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *PutApiEventsByEventIdPricingRequest) SetCurrency(v string)`

SetCurrency sets Currency field to given value.


### GetAllFeesIncluded

`func (o *PutApiEventsByEventIdPricingRequest) GetAllFeesIncluded() bool`

GetAllFeesIncluded returns the AllFeesIncluded field if non-nil, zero value otherwise.

### GetAllFeesIncludedOk

`func (o *PutApiEventsByEventIdPricingRequest) GetAllFeesIncludedOk() (*bool, bool)`

GetAllFeesIncludedOk returns a tuple with the AllFeesIncluded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllFeesIncluded

`func (o *PutApiEventsByEventIdPricingRequest) SetAllFeesIncluded(v bool)`

SetAllFeesIncluded sets AllFeesIncluded field to given value.

### HasAllFeesIncluded

`func (o *PutApiEventsByEventIdPricingRequest) HasAllFeesIncluded() bool`

HasAllFeesIncluded returns a boolean if a field has been set.

### GetCategories

`func (o *PutApiEventsByEventIdPricingRequest) GetCategories() []PutApiEventsByEventIdPricingRequestCategoriesInner`

GetCategories returns the Categories field if non-nil, zero value otherwise.

### GetCategoriesOk

`func (o *PutApiEventsByEventIdPricingRequest) GetCategoriesOk() (*[]PutApiEventsByEventIdPricingRequestCategoriesInner, bool)`

GetCategoriesOk returns a tuple with the Categories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategories

`func (o *PutApiEventsByEventIdPricingRequest) SetCategories(v []PutApiEventsByEventIdPricingRequestCategoriesInner)`

SetCategories sets Categories field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


