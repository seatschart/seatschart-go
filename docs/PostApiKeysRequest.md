# PostApiKeysRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Scopes** | Pointer to **[]string** | Granted scopes: &#x60;read&#x60;, &#x60;holds&#x60;, &#x60;bookings&#x60;, &#x60;inventory_admin&#x60;, &#x60;charts_write&#x60;, &#x60;events_write&#x60;, &#x60;reports&#x60;, &#x60;webhooks&#x60;. Granting every scope is full access. | [optional] 
**WorkspaceAdmin** | Pointer to **bool** | Workspace admin key: it may also list, activate and deactivate every workspace **you** administer (&#x60;/api/workspaces&#x60;). Bound to the calling admin and re-checked on each call; requires the &#x60;admin&#x60; role in this workspace. &#x60;false&#x60; removes the capability. | [optional] 

## Methods

### NewPostApiKeysRequest

`func NewPostApiKeysRequest(name string, ) *PostApiKeysRequest`

NewPostApiKeysRequest instantiates a new PostApiKeysRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostApiKeysRequestWithDefaults

`func NewPostApiKeysRequestWithDefaults() *PostApiKeysRequest`

NewPostApiKeysRequestWithDefaults instantiates a new PostApiKeysRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *PostApiKeysRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PostApiKeysRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PostApiKeysRequest) SetName(v string)`

SetName sets Name field to given value.


### GetScopes

`func (o *PostApiKeysRequest) GetScopes() []string`

GetScopes returns the Scopes field if non-nil, zero value otherwise.

### GetScopesOk

`func (o *PostApiKeysRequest) GetScopesOk() (*[]string, bool)`

GetScopesOk returns a tuple with the Scopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopes

`func (o *PostApiKeysRequest) SetScopes(v []string)`

SetScopes sets Scopes field to given value.

### HasScopes

`func (o *PostApiKeysRequest) HasScopes() bool`

HasScopes returns a boolean if a field has been set.

### GetWorkspaceAdmin

`func (o *PostApiKeysRequest) GetWorkspaceAdmin() bool`

GetWorkspaceAdmin returns the WorkspaceAdmin field if non-nil, zero value otherwise.

### GetWorkspaceAdminOk

`func (o *PostApiKeysRequest) GetWorkspaceAdminOk() (*bool, bool)`

GetWorkspaceAdminOk returns a tuple with the WorkspaceAdmin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkspaceAdmin

`func (o *PostApiKeysRequest) SetWorkspaceAdmin(v bool)`

SetWorkspaceAdmin sets WorkspaceAdmin field to given value.

### HasWorkspaceAdmin

`func (o *PostApiKeysRequest) HasWorkspaceAdmin() bool`

HasWorkspaceAdmin returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


