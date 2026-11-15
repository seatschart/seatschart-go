# PatchApiChartsByChartIdRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**Key** | Pointer to **NullableString** | Set, rename (&#x60;\&quot;new-key\&quot;&#x60;) or clear (&#x60;null&#x60;) the chart key. 409 &#x60;chart_key_taken&#x60; when another chart of the workspace uses it. | [optional] 

## Methods

### NewPatchApiChartsByChartIdRequest

`func NewPatchApiChartsByChartIdRequest() *PatchApiChartsByChartIdRequest`

NewPatchApiChartsByChartIdRequest instantiates a new PatchApiChartsByChartIdRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchApiChartsByChartIdRequestWithDefaults

`func NewPatchApiChartsByChartIdRequestWithDefaults() *PatchApiChartsByChartIdRequest`

NewPatchApiChartsByChartIdRequestWithDefaults instantiates a new PatchApiChartsByChartIdRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *PatchApiChartsByChartIdRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PatchApiChartsByChartIdRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PatchApiChartsByChartIdRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PatchApiChartsByChartIdRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetKey

`func (o *PatchApiChartsByChartIdRequest) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *PatchApiChartsByChartIdRequest) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *PatchApiChartsByChartIdRequest) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *PatchApiChartsByChartIdRequest) HasKey() bool`

HasKey returns a boolean if a field has been set.

### SetKeyNil

`func (o *PatchApiChartsByChartIdRequest) SetKeyNil(b bool)`

 SetKeyNil sets the value for Key to be an explicit nil

### UnsetKey
`func (o *PatchApiChartsByChartIdRequest) UnsetKey()`

UnsetKey ensures that no value is present for Key, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


