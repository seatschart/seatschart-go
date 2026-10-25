# PatchApiSeasonsBySeasonIdRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**Key** | Pointer to **NullableString** | Set or rename the key; &#x60;null&#x60; or &#x60;\&quot;\&quot;&#x60; clears it. 409 &#x60;season_key_taken&#x60;. Partner key: unique per workspace, 1–100 URL-safe characters (letters, digits, &#x60;_ . ~ -&#x60;, starting with a letter or digit), never UUID-shaped. Every season route accepts the key in place of the id (path, body and query &#x60;seasonId&#x60;), like event and chart keys. | [optional] 
**SourceEventId** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewPatchApiSeasonsBySeasonIdRequest

`func NewPatchApiSeasonsBySeasonIdRequest() *PatchApiSeasonsBySeasonIdRequest`

NewPatchApiSeasonsBySeasonIdRequest instantiates a new PatchApiSeasonsBySeasonIdRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchApiSeasonsBySeasonIdRequestWithDefaults

`func NewPatchApiSeasonsBySeasonIdRequestWithDefaults() *PatchApiSeasonsBySeasonIdRequest`

NewPatchApiSeasonsBySeasonIdRequestWithDefaults instantiates a new PatchApiSeasonsBySeasonIdRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *PatchApiSeasonsBySeasonIdRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PatchApiSeasonsBySeasonIdRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PatchApiSeasonsBySeasonIdRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PatchApiSeasonsBySeasonIdRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetKey

`func (o *PatchApiSeasonsBySeasonIdRequest) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *PatchApiSeasonsBySeasonIdRequest) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *PatchApiSeasonsBySeasonIdRequest) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *PatchApiSeasonsBySeasonIdRequest) HasKey() bool`

HasKey returns a boolean if a field has been set.

### SetKeyNil

`func (o *PatchApiSeasonsBySeasonIdRequest) SetKeyNil(b bool)`

 SetKeyNil sets the value for Key to be an explicit nil

### UnsetKey
`func (o *PatchApiSeasonsBySeasonIdRequest) UnsetKey()`

UnsetKey ensures that no value is present for Key, not even an explicit nil
### GetSourceEventId

`func (o *PatchApiSeasonsBySeasonIdRequest) GetSourceEventId() string`

GetSourceEventId returns the SourceEventId field if non-nil, zero value otherwise.

### GetSourceEventIdOk

`func (o *PatchApiSeasonsBySeasonIdRequest) GetSourceEventIdOk() (*string, bool)`

GetSourceEventIdOk returns a tuple with the SourceEventId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceEventId

`func (o *PatchApiSeasonsBySeasonIdRequest) SetSourceEventId(v string)`

SetSourceEventId sets SourceEventId field to given value.

### HasSourceEventId

`func (o *PatchApiSeasonsBySeasonIdRequest) HasSourceEventId() bool`

HasSourceEventId returns a boolean if a field has been set.

### SetSourceEventIdNil

`func (o *PatchApiSeasonsBySeasonIdRequest) SetSourceEventIdNil(b bool)`

 SetSourceEventIdNil sets the value for SourceEventId to be an explicit nil

### UnsetSourceEventId
`func (o *PatchApiSeasonsBySeasonIdRequest) UnsetSourceEventId()`

UnsetSourceEventId ensures that no value is present for SourceEventId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


