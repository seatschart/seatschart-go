# PatchApiListingTypesByKeyRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**Kind** | Pointer to **string** | &#x60;resale&#x60;: objects resold by their owner (listingType filter &#x60;resale&#x60;); &#x60;primary&#x60;: first-hand inventory sold as a listing. | [optional] 
**BackgroundColor** | Pointer to **NullableString** | CSS hex color &#x60;#rrggbb&#x60;. | [optional] 
**IconColor** | Pointer to **NullableString** | CSS hex color &#x60;#rrggbb&#x60;. | [optional] 
**Icon** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewPatchApiListingTypesByKeyRequest

`func NewPatchApiListingTypesByKeyRequest() *PatchApiListingTypesByKeyRequest`

NewPatchApiListingTypesByKeyRequest instantiates a new PatchApiListingTypesByKeyRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchApiListingTypesByKeyRequestWithDefaults

`func NewPatchApiListingTypesByKeyRequestWithDefaults() *PatchApiListingTypesByKeyRequest`

NewPatchApiListingTypesByKeyRequestWithDefaults instantiates a new PatchApiListingTypesByKeyRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *PatchApiListingTypesByKeyRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PatchApiListingTypesByKeyRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PatchApiListingTypesByKeyRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PatchApiListingTypesByKeyRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetKind

`func (o *PatchApiListingTypesByKeyRequest) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *PatchApiListingTypesByKeyRequest) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *PatchApiListingTypesByKeyRequest) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *PatchApiListingTypesByKeyRequest) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetBackgroundColor

`func (o *PatchApiListingTypesByKeyRequest) GetBackgroundColor() string`

GetBackgroundColor returns the BackgroundColor field if non-nil, zero value otherwise.

### GetBackgroundColorOk

`func (o *PatchApiListingTypesByKeyRequest) GetBackgroundColorOk() (*string, bool)`

GetBackgroundColorOk returns a tuple with the BackgroundColor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackgroundColor

`func (o *PatchApiListingTypesByKeyRequest) SetBackgroundColor(v string)`

SetBackgroundColor sets BackgroundColor field to given value.

### HasBackgroundColor

`func (o *PatchApiListingTypesByKeyRequest) HasBackgroundColor() bool`

HasBackgroundColor returns a boolean if a field has been set.

### SetBackgroundColorNil

`func (o *PatchApiListingTypesByKeyRequest) SetBackgroundColorNil(b bool)`

 SetBackgroundColorNil sets the value for BackgroundColor to be an explicit nil

### UnsetBackgroundColor
`func (o *PatchApiListingTypesByKeyRequest) UnsetBackgroundColor()`

UnsetBackgroundColor ensures that no value is present for BackgroundColor, not even an explicit nil
### GetIconColor

`func (o *PatchApiListingTypesByKeyRequest) GetIconColor() string`

GetIconColor returns the IconColor field if non-nil, zero value otherwise.

### GetIconColorOk

`func (o *PatchApiListingTypesByKeyRequest) GetIconColorOk() (*string, bool)`

GetIconColorOk returns a tuple with the IconColor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIconColor

`func (o *PatchApiListingTypesByKeyRequest) SetIconColor(v string)`

SetIconColor sets IconColor field to given value.

### HasIconColor

`func (o *PatchApiListingTypesByKeyRequest) HasIconColor() bool`

HasIconColor returns a boolean if a field has been set.

### SetIconColorNil

`func (o *PatchApiListingTypesByKeyRequest) SetIconColorNil(b bool)`

 SetIconColorNil sets the value for IconColor to be an explicit nil

### UnsetIconColor
`func (o *PatchApiListingTypesByKeyRequest) UnsetIconColor()`

UnsetIconColor ensures that no value is present for IconColor, not even an explicit nil
### GetIcon

`func (o *PatchApiListingTypesByKeyRequest) GetIcon() string`

GetIcon returns the Icon field if non-nil, zero value otherwise.

### GetIconOk

`func (o *PatchApiListingTypesByKeyRequest) GetIconOk() (*string, bool)`

GetIconOk returns a tuple with the Icon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIcon

`func (o *PatchApiListingTypesByKeyRequest) SetIcon(v string)`

SetIcon sets Icon field to given value.

### HasIcon

`func (o *PatchApiListingTypesByKeyRequest) HasIcon() bool`

HasIcon returns a boolean if a field has been set.

### SetIconNil

`func (o *PatchApiListingTypesByKeyRequest) SetIconNil(b bool)`

 SetIconNil sets the value for Icon to be an explicit nil

### UnsetIcon
`func (o *PatchApiListingTypesByKeyRequest) UnsetIcon()`

UnsetIcon ensures that no value is present for Icon, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


