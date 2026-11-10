# PostApiEventsByEventIdActionsEditForSaleConfigRequestForSaleInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Object** | Pointer to **string** | Seat UUID or label, or GA area UUID or name (exactly one of object / category). | [optional] 
**Quantity** | Pointer to **int32** | GA places: &#x60;notForSale&#x60; withholds this many more places, &#x60;forSale&#x60; returns this many withheld places. | [optional] 
**Category** | Pointer to **string** | Every seat and GA area of this category. | [optional] 

## Methods

### NewPostApiEventsByEventIdActionsEditForSaleConfigRequestForSaleInner

`func NewPostApiEventsByEventIdActionsEditForSaleConfigRequestForSaleInner() *PostApiEventsByEventIdActionsEditForSaleConfigRequestForSaleInner`

NewPostApiEventsByEventIdActionsEditForSaleConfigRequestForSaleInner instantiates a new PostApiEventsByEventIdActionsEditForSaleConfigRequestForSaleInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostApiEventsByEventIdActionsEditForSaleConfigRequestForSaleInnerWithDefaults

`func NewPostApiEventsByEventIdActionsEditForSaleConfigRequestForSaleInnerWithDefaults() *PostApiEventsByEventIdActionsEditForSaleConfigRequestForSaleInner`

NewPostApiEventsByEventIdActionsEditForSaleConfigRequestForSaleInnerWithDefaults instantiates a new PostApiEventsByEventIdActionsEditForSaleConfigRequestForSaleInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetObject

`func (o *PostApiEventsByEventIdActionsEditForSaleConfigRequestForSaleInner) GetObject() string`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *PostApiEventsByEventIdActionsEditForSaleConfigRequestForSaleInner) GetObjectOk() (*string, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *PostApiEventsByEventIdActionsEditForSaleConfigRequestForSaleInner) SetObject(v string)`

SetObject sets Object field to given value.

### HasObject

`func (o *PostApiEventsByEventIdActionsEditForSaleConfigRequestForSaleInner) HasObject() bool`

HasObject returns a boolean if a field has been set.

### GetQuantity

`func (o *PostApiEventsByEventIdActionsEditForSaleConfigRequestForSaleInner) GetQuantity() int32`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *PostApiEventsByEventIdActionsEditForSaleConfigRequestForSaleInner) GetQuantityOk() (*int32, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *PostApiEventsByEventIdActionsEditForSaleConfigRequestForSaleInner) SetQuantity(v int32)`

SetQuantity sets Quantity field to given value.

### HasQuantity

`func (o *PostApiEventsByEventIdActionsEditForSaleConfigRequestForSaleInner) HasQuantity() bool`

HasQuantity returns a boolean if a field has been set.

### GetCategory

`func (o *PostApiEventsByEventIdActionsEditForSaleConfigRequestForSaleInner) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *PostApiEventsByEventIdActionsEditForSaleConfigRequestForSaleInner) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *PostApiEventsByEventIdActionsEditForSaleConfigRequestForSaleInner) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *PostApiEventsByEventIdActionsEditForSaleConfigRequestForSaleInner) HasCategory() bool`

HasCategory returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


