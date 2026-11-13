# PatchApiWorkspacesByWorkspaceIdKeysByKeyIdRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**FullAccess** | Pointer to **bool** | Every scope. Default when &#x60;scopes&#x60; is absent. | [optional] 
**Scopes** | Pointer to **[]string** | Granted scopes: &#x60;read&#x60;, &#x60;holds&#x60;, &#x60;bookings&#x60;, &#x60;inventory_admin&#x60;, &#x60;charts_write&#x60;, &#x60;events_write&#x60;, &#x60;reports&#x60;, &#x60;webhooks&#x60;. | [optional] 

## Methods

### NewPatchApiWorkspacesByWorkspaceIdKeysByKeyIdRequest

`func NewPatchApiWorkspacesByWorkspaceIdKeysByKeyIdRequest() *PatchApiWorkspacesByWorkspaceIdKeysByKeyIdRequest`

NewPatchApiWorkspacesByWorkspaceIdKeysByKeyIdRequest instantiates a new PatchApiWorkspacesByWorkspaceIdKeysByKeyIdRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchApiWorkspacesByWorkspaceIdKeysByKeyIdRequestWithDefaults

`func NewPatchApiWorkspacesByWorkspaceIdKeysByKeyIdRequestWithDefaults() *PatchApiWorkspacesByWorkspaceIdKeysByKeyIdRequest`

NewPatchApiWorkspacesByWorkspaceIdKeysByKeyIdRequestWithDefaults instantiates a new PatchApiWorkspacesByWorkspaceIdKeysByKeyIdRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *PatchApiWorkspacesByWorkspaceIdKeysByKeyIdRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PatchApiWorkspacesByWorkspaceIdKeysByKeyIdRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PatchApiWorkspacesByWorkspaceIdKeysByKeyIdRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PatchApiWorkspacesByWorkspaceIdKeysByKeyIdRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetFullAccess

`func (o *PatchApiWorkspacesByWorkspaceIdKeysByKeyIdRequest) GetFullAccess() bool`

GetFullAccess returns the FullAccess field if non-nil, zero value otherwise.

### GetFullAccessOk

`func (o *PatchApiWorkspacesByWorkspaceIdKeysByKeyIdRequest) GetFullAccessOk() (*bool, bool)`

GetFullAccessOk returns a tuple with the FullAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFullAccess

`func (o *PatchApiWorkspacesByWorkspaceIdKeysByKeyIdRequest) SetFullAccess(v bool)`

SetFullAccess sets FullAccess field to given value.

### HasFullAccess

`func (o *PatchApiWorkspacesByWorkspaceIdKeysByKeyIdRequest) HasFullAccess() bool`

HasFullAccess returns a boolean if a field has been set.

### GetScopes

`func (o *PatchApiWorkspacesByWorkspaceIdKeysByKeyIdRequest) GetScopes() []string`

GetScopes returns the Scopes field if non-nil, zero value otherwise.

### GetScopesOk

`func (o *PatchApiWorkspacesByWorkspaceIdKeysByKeyIdRequest) GetScopesOk() (*[]string, bool)`

GetScopesOk returns a tuple with the Scopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopes

`func (o *PatchApiWorkspacesByWorkspaceIdKeysByKeyIdRequest) SetScopes(v []string)`

SetScopes sets Scopes field to given value.

### HasScopes

`func (o *PatchApiWorkspacesByWorkspaceIdKeysByKeyIdRequest) HasScopes() bool`

HasScopes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


