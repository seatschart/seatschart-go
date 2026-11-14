# PostApiWorkspacesByWorkspaceIdKeys201ResponseApiKey

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Name** | **string** |  | 
**Prefix** | **string** | Non-secret fingerprint (&#x60;sk_live_1a2b3c4d…&#x60;). | 
**CreatedAt** | **time.Time** |  | 
**Scopes** | **[]string** |  | 
**FullAccess** | **bool** |  | 
**WorkspaceAdmin** | **bool** | A workspace admin key (read only here). | 
**Secret** | **string** | Shown once. | 

## Methods

### NewPostApiWorkspacesByWorkspaceIdKeys201ResponseApiKey

`func NewPostApiWorkspacesByWorkspaceIdKeys201ResponseApiKey(id string, name string, prefix string, createdAt time.Time, scopes []string, fullAccess bool, workspaceAdmin bool, secret string, ) *PostApiWorkspacesByWorkspaceIdKeys201ResponseApiKey`

NewPostApiWorkspacesByWorkspaceIdKeys201ResponseApiKey instantiates a new PostApiWorkspacesByWorkspaceIdKeys201ResponseApiKey object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostApiWorkspacesByWorkspaceIdKeys201ResponseApiKeyWithDefaults

`func NewPostApiWorkspacesByWorkspaceIdKeys201ResponseApiKeyWithDefaults() *PostApiWorkspacesByWorkspaceIdKeys201ResponseApiKey`

NewPostApiWorkspacesByWorkspaceIdKeys201ResponseApiKeyWithDefaults instantiates a new PostApiWorkspacesByWorkspaceIdKeys201ResponseApiKey object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PostApiWorkspacesByWorkspaceIdKeys201ResponseApiKey) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PostApiWorkspacesByWorkspaceIdKeys201ResponseApiKey) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PostApiWorkspacesByWorkspaceIdKeys201ResponseApiKey) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *PostApiWorkspacesByWorkspaceIdKeys201ResponseApiKey) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PostApiWorkspacesByWorkspaceIdKeys201ResponseApiKey) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PostApiWorkspacesByWorkspaceIdKeys201ResponseApiKey) SetName(v string)`

SetName sets Name field to given value.


### GetPrefix

`func (o *PostApiWorkspacesByWorkspaceIdKeys201ResponseApiKey) GetPrefix() string`

GetPrefix returns the Prefix field if non-nil, zero value otherwise.

### GetPrefixOk

`func (o *PostApiWorkspacesByWorkspaceIdKeys201ResponseApiKey) GetPrefixOk() (*string, bool)`

GetPrefixOk returns a tuple with the Prefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrefix

`func (o *PostApiWorkspacesByWorkspaceIdKeys201ResponseApiKey) SetPrefix(v string)`

SetPrefix sets Prefix field to given value.


### GetCreatedAt

`func (o *PostApiWorkspacesByWorkspaceIdKeys201ResponseApiKey) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *PostApiWorkspacesByWorkspaceIdKeys201ResponseApiKey) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *PostApiWorkspacesByWorkspaceIdKeys201ResponseApiKey) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetScopes

`func (o *PostApiWorkspacesByWorkspaceIdKeys201ResponseApiKey) GetScopes() []string`

GetScopes returns the Scopes field if non-nil, zero value otherwise.

### GetScopesOk

`func (o *PostApiWorkspacesByWorkspaceIdKeys201ResponseApiKey) GetScopesOk() (*[]string, bool)`

GetScopesOk returns a tuple with the Scopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopes

`func (o *PostApiWorkspacesByWorkspaceIdKeys201ResponseApiKey) SetScopes(v []string)`

SetScopes sets Scopes field to given value.


### GetFullAccess

`func (o *PostApiWorkspacesByWorkspaceIdKeys201ResponseApiKey) GetFullAccess() bool`

GetFullAccess returns the FullAccess field if non-nil, zero value otherwise.

### GetFullAccessOk

`func (o *PostApiWorkspacesByWorkspaceIdKeys201ResponseApiKey) GetFullAccessOk() (*bool, bool)`

GetFullAccessOk returns a tuple with the FullAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFullAccess

`func (o *PostApiWorkspacesByWorkspaceIdKeys201ResponseApiKey) SetFullAccess(v bool)`

SetFullAccess sets FullAccess field to given value.


### GetWorkspaceAdmin

`func (o *PostApiWorkspacesByWorkspaceIdKeys201ResponseApiKey) GetWorkspaceAdmin() bool`

GetWorkspaceAdmin returns the WorkspaceAdmin field if non-nil, zero value otherwise.

### GetWorkspaceAdminOk

`func (o *PostApiWorkspacesByWorkspaceIdKeys201ResponseApiKey) GetWorkspaceAdminOk() (*bool, bool)`

GetWorkspaceAdminOk returns a tuple with the WorkspaceAdmin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkspaceAdmin

`func (o *PostApiWorkspacesByWorkspaceIdKeys201ResponseApiKey) SetWorkspaceAdmin(v bool)`

SetWorkspaceAdmin sets WorkspaceAdmin field to given value.


### GetSecret

`func (o *PostApiWorkspacesByWorkspaceIdKeys201ResponseApiKey) GetSecret() string`

GetSecret returns the Secret field if non-nil, zero value otherwise.

### GetSecretOk

`func (o *PostApiWorkspacesByWorkspaceIdKeys201ResponseApiKey) GetSecretOk() (*string, bool)`

GetSecretOk returns a tuple with the Secret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecret

`func (o *PostApiWorkspacesByWorkspaceIdKeys201ResponseApiKey) SetSecret(v string)`

SetSecret sets Secret field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


