# GetApiWorkspacesByWorkspaceIdKeys200ResponseApiKeysInner

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

## Methods

### NewGetApiWorkspacesByWorkspaceIdKeys200ResponseApiKeysInner

`func NewGetApiWorkspacesByWorkspaceIdKeys200ResponseApiKeysInner(id string, name string, prefix string, createdAt time.Time, scopes []string, fullAccess bool, workspaceAdmin bool, ) *GetApiWorkspacesByWorkspaceIdKeys200ResponseApiKeysInner`

NewGetApiWorkspacesByWorkspaceIdKeys200ResponseApiKeysInner instantiates a new GetApiWorkspacesByWorkspaceIdKeys200ResponseApiKeysInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetApiWorkspacesByWorkspaceIdKeys200ResponseApiKeysInnerWithDefaults

`func NewGetApiWorkspacesByWorkspaceIdKeys200ResponseApiKeysInnerWithDefaults() *GetApiWorkspacesByWorkspaceIdKeys200ResponseApiKeysInner`

NewGetApiWorkspacesByWorkspaceIdKeys200ResponseApiKeysInnerWithDefaults instantiates a new GetApiWorkspacesByWorkspaceIdKeys200ResponseApiKeysInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GetApiWorkspacesByWorkspaceIdKeys200ResponseApiKeysInner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetApiWorkspacesByWorkspaceIdKeys200ResponseApiKeysInner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetApiWorkspacesByWorkspaceIdKeys200ResponseApiKeysInner) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *GetApiWorkspacesByWorkspaceIdKeys200ResponseApiKeysInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetApiWorkspacesByWorkspaceIdKeys200ResponseApiKeysInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetApiWorkspacesByWorkspaceIdKeys200ResponseApiKeysInner) SetName(v string)`

SetName sets Name field to given value.


### GetPrefix

`func (o *GetApiWorkspacesByWorkspaceIdKeys200ResponseApiKeysInner) GetPrefix() string`

GetPrefix returns the Prefix field if non-nil, zero value otherwise.

### GetPrefixOk

`func (o *GetApiWorkspacesByWorkspaceIdKeys200ResponseApiKeysInner) GetPrefixOk() (*string, bool)`

GetPrefixOk returns a tuple with the Prefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrefix

`func (o *GetApiWorkspacesByWorkspaceIdKeys200ResponseApiKeysInner) SetPrefix(v string)`

SetPrefix sets Prefix field to given value.


### GetCreatedAt

`func (o *GetApiWorkspacesByWorkspaceIdKeys200ResponseApiKeysInner) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *GetApiWorkspacesByWorkspaceIdKeys200ResponseApiKeysInner) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *GetApiWorkspacesByWorkspaceIdKeys200ResponseApiKeysInner) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetScopes

`func (o *GetApiWorkspacesByWorkspaceIdKeys200ResponseApiKeysInner) GetScopes() []string`

GetScopes returns the Scopes field if non-nil, zero value otherwise.

### GetScopesOk

`func (o *GetApiWorkspacesByWorkspaceIdKeys200ResponseApiKeysInner) GetScopesOk() (*[]string, bool)`

GetScopesOk returns a tuple with the Scopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopes

`func (o *GetApiWorkspacesByWorkspaceIdKeys200ResponseApiKeysInner) SetScopes(v []string)`

SetScopes sets Scopes field to given value.


### GetFullAccess

`func (o *GetApiWorkspacesByWorkspaceIdKeys200ResponseApiKeysInner) GetFullAccess() bool`

GetFullAccess returns the FullAccess field if non-nil, zero value otherwise.

### GetFullAccessOk

`func (o *GetApiWorkspacesByWorkspaceIdKeys200ResponseApiKeysInner) GetFullAccessOk() (*bool, bool)`

GetFullAccessOk returns a tuple with the FullAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFullAccess

`func (o *GetApiWorkspacesByWorkspaceIdKeys200ResponseApiKeysInner) SetFullAccess(v bool)`

SetFullAccess sets FullAccess field to given value.


### GetWorkspaceAdmin

`func (o *GetApiWorkspacesByWorkspaceIdKeys200ResponseApiKeysInner) GetWorkspaceAdmin() bool`

GetWorkspaceAdmin returns the WorkspaceAdmin field if non-nil, zero value otherwise.

### GetWorkspaceAdminOk

`func (o *GetApiWorkspacesByWorkspaceIdKeys200ResponseApiKeysInner) GetWorkspaceAdminOk() (*bool, bool)`

GetWorkspaceAdminOk returns a tuple with the WorkspaceAdmin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkspaceAdmin

`func (o *GetApiWorkspacesByWorkspaceIdKeys200ResponseApiKeysInner) SetWorkspaceAdmin(v bool)`

SetWorkspaceAdmin sets WorkspaceAdmin field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


