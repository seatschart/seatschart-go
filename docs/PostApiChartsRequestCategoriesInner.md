# PostApiChartsRequestCategoriesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Key** | Pointer to **string** | Stable category id (generated when omitted). | [optional] 
**Label** | **string** | Category name; seats and GA areas reference it. | 
**Color** | **string** | Category colour &#x60;#RRGGBB&#x60;. | 
**Accessible** | Pointer to **bool** |  | [optional] 

## Methods

### NewPostApiChartsRequestCategoriesInner

`func NewPostApiChartsRequestCategoriesInner(label string, color string, ) *PostApiChartsRequestCategoriesInner`

NewPostApiChartsRequestCategoriesInner instantiates a new PostApiChartsRequestCategoriesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostApiChartsRequestCategoriesInnerWithDefaults

`func NewPostApiChartsRequestCategoriesInnerWithDefaults() *PostApiChartsRequestCategoriesInner`

NewPostApiChartsRequestCategoriesInnerWithDefaults instantiates a new PostApiChartsRequestCategoriesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKey

`func (o *PostApiChartsRequestCategoriesInner) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *PostApiChartsRequestCategoriesInner) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *PostApiChartsRequestCategoriesInner) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *PostApiChartsRequestCategoriesInner) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetLabel

`func (o *PostApiChartsRequestCategoriesInner) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *PostApiChartsRequestCategoriesInner) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *PostApiChartsRequestCategoriesInner) SetLabel(v string)`

SetLabel sets Label field to given value.


### GetColor

`func (o *PostApiChartsRequestCategoriesInner) GetColor() string`

GetColor returns the Color field if non-nil, zero value otherwise.

### GetColorOk

`func (o *PostApiChartsRequestCategoriesInner) GetColorOk() (*string, bool)`

GetColorOk returns a tuple with the Color field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColor

`func (o *PostApiChartsRequestCategoriesInner) SetColor(v string)`

SetColor sets Color field to given value.


### GetAccessible

`func (o *PostApiChartsRequestCategoriesInner) GetAccessible() bool`

GetAccessible returns the Accessible field if non-nil, zero value otherwise.

### GetAccessibleOk

`func (o *PostApiChartsRequestCategoriesInner) GetAccessibleOk() (*bool, bool)`

GetAccessibleOk returns a tuple with the Accessible field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessible

`func (o *PostApiChartsRequestCategoriesInner) SetAccessible(v bool)`

SetAccessible sets Accessible field to given value.

### HasAccessible

`func (o *PostApiChartsRequestCategoriesInner) HasAccessible() bool`

HasAccessible returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


