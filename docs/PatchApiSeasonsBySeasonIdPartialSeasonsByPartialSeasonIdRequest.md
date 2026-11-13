# PatchApiSeasonsBySeasonIdPartialSeasonsByPartialSeasonIdRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**Key** | Pointer to **NullableString** | Partner &#x60;key&#x60;: unique within the season, 1–100 URL-safe characters, never UUID-shaped; usable in place of the partial season id in every partial season path (&#x60;/seasons/{seasonIdOrKey}/partial-seasons/{partialSeasonIdOrKey}&#x60;), and as body &#x60;seasonId&#x60; of booking and status changes when no season and no other partial season of the workspace uses it. | [optional] 

## Methods

### NewPatchApiSeasonsBySeasonIdPartialSeasonsByPartialSeasonIdRequest

`func NewPatchApiSeasonsBySeasonIdPartialSeasonsByPartialSeasonIdRequest() *PatchApiSeasonsBySeasonIdPartialSeasonsByPartialSeasonIdRequest`

NewPatchApiSeasonsBySeasonIdPartialSeasonsByPartialSeasonIdRequest instantiates a new PatchApiSeasonsBySeasonIdPartialSeasonsByPartialSeasonIdRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchApiSeasonsBySeasonIdPartialSeasonsByPartialSeasonIdRequestWithDefaults

`func NewPatchApiSeasonsBySeasonIdPartialSeasonsByPartialSeasonIdRequestWithDefaults() *PatchApiSeasonsBySeasonIdPartialSeasonsByPartialSeasonIdRequest`

NewPatchApiSeasonsBySeasonIdPartialSeasonsByPartialSeasonIdRequestWithDefaults instantiates a new PatchApiSeasonsBySeasonIdPartialSeasonsByPartialSeasonIdRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *PatchApiSeasonsBySeasonIdPartialSeasonsByPartialSeasonIdRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PatchApiSeasonsBySeasonIdPartialSeasonsByPartialSeasonIdRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PatchApiSeasonsBySeasonIdPartialSeasonsByPartialSeasonIdRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PatchApiSeasonsBySeasonIdPartialSeasonsByPartialSeasonIdRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetKey

`func (o *PatchApiSeasonsBySeasonIdPartialSeasonsByPartialSeasonIdRequest) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *PatchApiSeasonsBySeasonIdPartialSeasonsByPartialSeasonIdRequest) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *PatchApiSeasonsBySeasonIdPartialSeasonsByPartialSeasonIdRequest) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *PatchApiSeasonsBySeasonIdPartialSeasonsByPartialSeasonIdRequest) HasKey() bool`

HasKey returns a boolean if a field has been set.

### SetKeyNil

`func (o *PatchApiSeasonsBySeasonIdPartialSeasonsByPartialSeasonIdRequest) SetKeyNil(b bool)`

 SetKeyNil sets the value for Key to be an explicit nil

### UnsetKey
`func (o *PatchApiSeasonsBySeasonIdPartialSeasonsByPartialSeasonIdRequest) UnsetKey()`

UnsetKey ensures that no value is present for Key, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


