# PatchApiKeysByIdRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**Scopes** | Pointer to **[]string** | Granted scopes: &#x60;read&#x60;, &#x60;holds&#x60;, &#x60;bookings&#x60;, &#x60;inventory_admin&#x60;, &#x60;charts_write&#x60;, &#x60;events_write&#x60;, &#x60;reports&#x60;, &#x60;webhooks&#x60;. Granting every scope is full access. | [optional] 
**WorkspaceAdmin** | Pointer to **bool** | Workspace admin key: it may also list, activate and deactivate every workspace **you** administer (&#x60;/api/workspaces&#x60;). Bound to the calling admin and re-checked on each call; requires the &#x60;admin&#x60; role in this workspace. &#x60;false&#x60; removes the capability. | [optional] 

## Methods

### NewPatchApiKeysByIdRequest

`func NewPatchApiKeysByIdRequest() *PatchApiKeysByIdRequest`

NewPatchApiKeysByIdRequest instantiates a new PatchApiKeysByIdRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchApiKeysByIdRequestWithDefaults

`func NewPatchApiKeysByIdRequestWithDefaults() *PatchApiKeysByIdRequest`

NewPatchApiKeysByIdRequestWithDefaults instantiates a new PatchApiKeysByIdRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *PatchApiKeysByIdRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PatchApiKeysByIdRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PatchApiKeysByIdRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PatchApiKeysByIdRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetScopes

`func (o *PatchApiKeysByIdRequest) GetScopes() []string`

GetScopes returns the Scopes field if non-nil, zero value otherwise.

### GetScopesOk

`func (o *PatchApiKeysByIdRequest) GetScopesOk() (*[]string, bool)`

GetScopesOk returns a tuple with the Scopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopes

`func (o *PatchApiKeysByIdRequest) SetScopes(v []string)`

SetScopes sets Scopes field to given value.

### HasScopes

`func (o *PatchApiKeysByIdRequest) HasScopes() bool`

HasScopes returns a boolean if a field has been set.

### GetWorkspaceAdmin

`func (o *PatchApiKeysByIdRequest) GetWorkspaceAdmin() bool`

GetWorkspaceAdmin returns the WorkspaceAdmin field if non-nil, zero value otherwise.

### GetWorkspaceAdminOk

`func (o *PatchApiKeysByIdRequest) GetWorkspaceAdminOk() (*bool, bool)`

GetWorkspaceAdminOk returns a tuple with the WorkspaceAdmin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkspaceAdmin

`func (o *PatchApiKeysByIdRequest) SetWorkspaceAdmin(v bool)`

SetWorkspaceAdmin sets WorkspaceAdmin field to given value.

### HasWorkspaceAdmin

`func (o *PatchApiKeysByIdRequest) HasWorkspaceAdmin() bool`

HasWorkspaceAdmin returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


