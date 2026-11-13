# PostApiEventsRequestCategoriesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Key** | **string** | Category key, unique in the event; referenced by objectCategories. | 
**Label** | **string** | Display name; also the seat category value (overrides, price tables, sales rules). | 
**Color** | **string** | #RRGGBB. | 
**Accessible** | Pointer to **NullableBool** |  | [optional] 

## Methods

### NewPostApiEventsRequestCategoriesInner

`func NewPostApiEventsRequestCategoriesInner(key string, label string, color string, ) *PostApiEventsRequestCategoriesInner`

NewPostApiEventsRequestCategoriesInner instantiates a new PostApiEventsRequestCategoriesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostApiEventsRequestCategoriesInnerWithDefaults

`func NewPostApiEventsRequestCategoriesInnerWithDefaults() *PostApiEventsRequestCategoriesInner`

NewPostApiEventsRequestCategoriesInnerWithDefaults instantiates a new PostApiEventsRequestCategoriesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKey

`func (o *PostApiEventsRequestCategoriesInner) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *PostApiEventsRequestCategoriesInner) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *PostApiEventsRequestCategoriesInner) SetKey(v string)`

SetKey sets Key field to given value.


### GetLabel

`func (o *PostApiEventsRequestCategoriesInner) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *PostApiEventsRequestCategoriesInner) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *PostApiEventsRequestCategoriesInner) SetLabel(v string)`

SetLabel sets Label field to given value.


### GetColor

`func (o *PostApiEventsRequestCategoriesInner) GetColor() string`

GetColor returns the Color field if non-nil, zero value otherwise.

### GetColorOk

`func (o *PostApiEventsRequestCategoriesInner) GetColorOk() (*string, bool)`

GetColorOk returns a tuple with the Color field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColor

`func (o *PostApiEventsRequestCategoriesInner) SetColor(v string)`

SetColor sets Color field to given value.


### GetAccessible

`func (o *PostApiEventsRequestCategoriesInner) GetAccessible() bool`

GetAccessible returns the Accessible field if non-nil, zero value otherwise.

### GetAccessibleOk

`func (o *PostApiEventsRequestCategoriesInner) GetAccessibleOk() (*bool, bool)`

GetAccessibleOk returns a tuple with the Accessible field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessible

`func (o *PostApiEventsRequestCategoriesInner) SetAccessible(v bool)`

SetAccessible sets Accessible field to given value.

### HasAccessible

`func (o *PostApiEventsRequestCategoriesInner) HasAccessible() bool`

HasAccessible returns a boolean if a field has been set.

### SetAccessibleNil

`func (o *PostApiEventsRequestCategoriesInner) SetAccessibleNil(b bool)`

 SetAccessibleNil sets the value for Accessible to be an explicit nil

### UnsetAccessible
`func (o *PostApiEventsRequestCategoriesInner) UnsetAccessible()`

UnsetAccessible ensures that no value is present for Accessible, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


