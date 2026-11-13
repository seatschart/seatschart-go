# PostApiListingTypesRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Key** | **string** |  | 
**Name** | **string** |  | 
**Kind** | Pointer to **string** | &#x60;resale&#x60;: objects resold by their owner (listingType filter &#x60;resale&#x60;); &#x60;primary&#x60;: first-hand inventory sold as a listing. | [optional] 
**BackgroundColor** | Pointer to **NullableString** | CSS hex color &#x60;#rrggbb&#x60;. | [optional] 
**IconColor** | Pointer to **NullableString** | CSS hex color &#x60;#rrggbb&#x60;. | [optional] 
**Icon** | Pointer to **NullableString** | Icon name (FontAwesome 4.7 names, as seats.io). | [optional] 

## Methods

### NewPostApiListingTypesRequest

`func NewPostApiListingTypesRequest(key string, name string, ) *PostApiListingTypesRequest`

NewPostApiListingTypesRequest instantiates a new PostApiListingTypesRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostApiListingTypesRequestWithDefaults

`func NewPostApiListingTypesRequestWithDefaults() *PostApiListingTypesRequest`

NewPostApiListingTypesRequestWithDefaults instantiates a new PostApiListingTypesRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKey

`func (o *PostApiListingTypesRequest) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *PostApiListingTypesRequest) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *PostApiListingTypesRequest) SetKey(v string)`

SetKey sets Key field to given value.


### GetName

`func (o *PostApiListingTypesRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PostApiListingTypesRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PostApiListingTypesRequest) SetName(v string)`

SetName sets Name field to given value.


### GetKind

`func (o *PostApiListingTypesRequest) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *PostApiListingTypesRequest) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *PostApiListingTypesRequest) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *PostApiListingTypesRequest) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetBackgroundColor

`func (o *PostApiListingTypesRequest) GetBackgroundColor() string`

GetBackgroundColor returns the BackgroundColor field if non-nil, zero value otherwise.

### GetBackgroundColorOk

`func (o *PostApiListingTypesRequest) GetBackgroundColorOk() (*string, bool)`

GetBackgroundColorOk returns a tuple with the BackgroundColor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackgroundColor

`func (o *PostApiListingTypesRequest) SetBackgroundColor(v string)`

SetBackgroundColor sets BackgroundColor field to given value.

### HasBackgroundColor

`func (o *PostApiListingTypesRequest) HasBackgroundColor() bool`

HasBackgroundColor returns a boolean if a field has been set.

### SetBackgroundColorNil

`func (o *PostApiListingTypesRequest) SetBackgroundColorNil(b bool)`

 SetBackgroundColorNil sets the value for BackgroundColor to be an explicit nil

### UnsetBackgroundColor
`func (o *PostApiListingTypesRequest) UnsetBackgroundColor()`

UnsetBackgroundColor ensures that no value is present for BackgroundColor, not even an explicit nil
### GetIconColor

`func (o *PostApiListingTypesRequest) GetIconColor() string`

GetIconColor returns the IconColor field if non-nil, zero value otherwise.

### GetIconColorOk

`func (o *PostApiListingTypesRequest) GetIconColorOk() (*string, bool)`

GetIconColorOk returns a tuple with the IconColor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIconColor

`func (o *PostApiListingTypesRequest) SetIconColor(v string)`

SetIconColor sets IconColor field to given value.

### HasIconColor

`func (o *PostApiListingTypesRequest) HasIconColor() bool`

HasIconColor returns a boolean if a field has been set.

### SetIconColorNil

`func (o *PostApiListingTypesRequest) SetIconColorNil(b bool)`

 SetIconColorNil sets the value for IconColor to be an explicit nil

### UnsetIconColor
`func (o *PostApiListingTypesRequest) UnsetIconColor()`

UnsetIconColor ensures that no value is present for IconColor, not even an explicit nil
### GetIcon

`func (o *PostApiListingTypesRequest) GetIcon() string`

GetIcon returns the Icon field if non-nil, zero value otherwise.

### GetIconOk

`func (o *PostApiListingTypesRequest) GetIconOk() (*string, bool)`

GetIconOk returns a tuple with the Icon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIcon

`func (o *PostApiListingTypesRequest) SetIcon(v string)`

SetIcon sets Icon field to given value.

### HasIcon

`func (o *PostApiListingTypesRequest) HasIcon() bool`

HasIcon returns a boolean if a field has been set.

### SetIconNil

`func (o *PostApiListingTypesRequest) SetIconNil(b bool)`

 SetIconNil sets the value for Icon to be an explicit nil

### UnsetIcon
`func (o *PostApiListingTypesRequest) UnsetIcon()`

UnsetIcon ensures that no value is present for Icon, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


