# PatchApiWorkspacesByWorkspaceIdRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DisplayName** | Pointer to **string** |  | [optional] 
**ExternalId** | Pointer to **NullableString** | Your organizer id (letters, digits and &#x60;_ . ~ -&#x60;); &#x60;null&#x60; removes it. Unique among the workspaces you administer: 409 &#x60;external_id_taken&#x60; names the workspace that has it. | [optional] 

## Methods

### NewPatchApiWorkspacesByWorkspaceIdRequest

`func NewPatchApiWorkspacesByWorkspaceIdRequest() *PatchApiWorkspacesByWorkspaceIdRequest`

NewPatchApiWorkspacesByWorkspaceIdRequest instantiates a new PatchApiWorkspacesByWorkspaceIdRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchApiWorkspacesByWorkspaceIdRequestWithDefaults

`func NewPatchApiWorkspacesByWorkspaceIdRequestWithDefaults() *PatchApiWorkspacesByWorkspaceIdRequest`

NewPatchApiWorkspacesByWorkspaceIdRequestWithDefaults instantiates a new PatchApiWorkspacesByWorkspaceIdRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDisplayName

`func (o *PatchApiWorkspacesByWorkspaceIdRequest) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *PatchApiWorkspacesByWorkspaceIdRequest) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *PatchApiWorkspacesByWorkspaceIdRequest) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *PatchApiWorkspacesByWorkspaceIdRequest) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetExternalId

`func (o *PatchApiWorkspacesByWorkspaceIdRequest) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *PatchApiWorkspacesByWorkspaceIdRequest) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *PatchApiWorkspacesByWorkspaceIdRequest) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *PatchApiWorkspacesByWorkspaceIdRequest) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### SetExternalIdNil

`func (o *PatchApiWorkspacesByWorkspaceIdRequest) SetExternalIdNil(b bool)`

 SetExternalIdNil sets the value for ExternalId to be an explicit nil

### UnsetExternalId
`func (o *PatchApiWorkspacesByWorkspaceIdRequest) UnsetExternalId()`

UnsetExternalId ensures that no value is present for ExternalId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


