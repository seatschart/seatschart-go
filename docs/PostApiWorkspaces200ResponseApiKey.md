# PostApiWorkspaces200ResponseApiKey

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Name** | **string** |  | 
**Prefix** | **string** |  | 
**CreatedAt** | **string** |  | 
**Scopes** | **[]string** |  | 
**FullAccess** | **bool** |  | 
**WorkspaceAdmin** | **bool** |  | 
**Secret** | **NullableString** | Shown once; null on an idempotent replay. | 

## Methods

### NewPostApiWorkspaces200ResponseApiKey

`func NewPostApiWorkspaces200ResponseApiKey(id string, name string, prefix string, createdAt string, scopes []string, fullAccess bool, workspaceAdmin bool, secret NullableString, ) *PostApiWorkspaces200ResponseApiKey`

NewPostApiWorkspaces200ResponseApiKey instantiates a new PostApiWorkspaces200ResponseApiKey object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostApiWorkspaces200ResponseApiKeyWithDefaults

`func NewPostApiWorkspaces200ResponseApiKeyWithDefaults() *PostApiWorkspaces200ResponseApiKey`

NewPostApiWorkspaces200ResponseApiKeyWithDefaults instantiates a new PostApiWorkspaces200ResponseApiKey object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PostApiWorkspaces200ResponseApiKey) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PostApiWorkspaces200ResponseApiKey) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PostApiWorkspaces200ResponseApiKey) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *PostApiWorkspaces200ResponseApiKey) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PostApiWorkspaces200ResponseApiKey) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PostApiWorkspaces200ResponseApiKey) SetName(v string)`

SetName sets Name field to given value.


### GetPrefix

`func (o *PostApiWorkspaces200ResponseApiKey) GetPrefix() string`

GetPrefix returns the Prefix field if non-nil, zero value otherwise.

### GetPrefixOk

`func (o *PostApiWorkspaces200ResponseApiKey) GetPrefixOk() (*string, bool)`

GetPrefixOk returns a tuple with the Prefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrefix

`func (o *PostApiWorkspaces200ResponseApiKey) SetPrefix(v string)`

SetPrefix sets Prefix field to given value.


### GetCreatedAt

`func (o *PostApiWorkspaces200ResponseApiKey) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *PostApiWorkspaces200ResponseApiKey) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *PostApiWorkspaces200ResponseApiKey) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.


### GetScopes

`func (o *PostApiWorkspaces200ResponseApiKey) GetScopes() []string`

GetScopes returns the Scopes field if non-nil, zero value otherwise.

### GetScopesOk

`func (o *PostApiWorkspaces200ResponseApiKey) GetScopesOk() (*[]string, bool)`

GetScopesOk returns a tuple with the Scopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopes

`func (o *PostApiWorkspaces200ResponseApiKey) SetScopes(v []string)`

SetScopes sets Scopes field to given value.


### GetFullAccess

`func (o *PostApiWorkspaces200ResponseApiKey) GetFullAccess() bool`

GetFullAccess returns the FullAccess field if non-nil, zero value otherwise.

### GetFullAccessOk

`func (o *PostApiWorkspaces200ResponseApiKey) GetFullAccessOk() (*bool, bool)`

GetFullAccessOk returns a tuple with the FullAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFullAccess

`func (o *PostApiWorkspaces200ResponseApiKey) SetFullAccess(v bool)`

SetFullAccess sets FullAccess field to given value.


### GetWorkspaceAdmin

`func (o *PostApiWorkspaces200ResponseApiKey) GetWorkspaceAdmin() bool`

GetWorkspaceAdmin returns the WorkspaceAdmin field if non-nil, zero value otherwise.

### GetWorkspaceAdminOk

`func (o *PostApiWorkspaces200ResponseApiKey) GetWorkspaceAdminOk() (*bool, bool)`

GetWorkspaceAdminOk returns a tuple with the WorkspaceAdmin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkspaceAdmin

`func (o *PostApiWorkspaces200ResponseApiKey) SetWorkspaceAdmin(v bool)`

SetWorkspaceAdmin sets WorkspaceAdmin field to given value.


### GetSecret

`func (o *PostApiWorkspaces200ResponseApiKey) GetSecret() string`

GetSecret returns the Secret field if non-nil, zero value otherwise.

### GetSecretOk

`func (o *PostApiWorkspaces200ResponseApiKey) GetSecretOk() (*string, bool)`

GetSecretOk returns a tuple with the Secret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecret

`func (o *PostApiWorkspaces200ResponseApiKey) SetSecret(v string)`

SetSecret sets Secret field to given value.


### SetSecretNil

`func (o *PostApiWorkspaces200ResponseApiKey) SetSecretNil(b bool)`

 SetSecretNil sets the value for Secret to be an explicit nil

### UnsetSecret
`func (o *PostApiWorkspaces200ResponseApiKey) UnsetSecret()`

UnsetSecret ensures that no value is present for Secret, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


