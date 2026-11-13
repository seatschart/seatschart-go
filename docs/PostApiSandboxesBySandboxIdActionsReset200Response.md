# PostApiSandboxesBySandboxIdActionsReset200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Sandbox** | [**GetApiSandboxes200ResponseItemsInner**](GetApiSandboxes200ResponseItemsInner.md) |  | 
**Deleted** | **map[string]int32** | Rows deleted per table. | 
**RevokedKeys** | **int32** |  | 
**AdminKey** | [**NullablePostApiSandboxes201ResponseAdminKey**](PostApiSandboxes201ResponseAdminKey.md) |  | 

## Methods

### NewPostApiSandboxesBySandboxIdActionsReset200Response

`func NewPostApiSandboxesBySandboxIdActionsReset200Response(sandbox GetApiSandboxes200ResponseItemsInner, deleted map[string]int32, revokedKeys int32, adminKey NullablePostApiSandboxes201ResponseAdminKey, ) *PostApiSandboxesBySandboxIdActionsReset200Response`

NewPostApiSandboxesBySandboxIdActionsReset200Response instantiates a new PostApiSandboxesBySandboxIdActionsReset200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostApiSandboxesBySandboxIdActionsReset200ResponseWithDefaults

`func NewPostApiSandboxesBySandboxIdActionsReset200ResponseWithDefaults() *PostApiSandboxesBySandboxIdActionsReset200Response`

NewPostApiSandboxesBySandboxIdActionsReset200ResponseWithDefaults instantiates a new PostApiSandboxesBySandboxIdActionsReset200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSandbox

`func (o *PostApiSandboxesBySandboxIdActionsReset200Response) GetSandbox() GetApiSandboxes200ResponseItemsInner`

GetSandbox returns the Sandbox field if non-nil, zero value otherwise.

### GetSandboxOk

`func (o *PostApiSandboxesBySandboxIdActionsReset200Response) GetSandboxOk() (*GetApiSandboxes200ResponseItemsInner, bool)`

GetSandboxOk returns a tuple with the Sandbox field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSandbox

`func (o *PostApiSandboxesBySandboxIdActionsReset200Response) SetSandbox(v GetApiSandboxes200ResponseItemsInner)`

SetSandbox sets Sandbox field to given value.


### GetDeleted

`func (o *PostApiSandboxesBySandboxIdActionsReset200Response) GetDeleted() map[string]int32`

GetDeleted returns the Deleted field if non-nil, zero value otherwise.

### GetDeletedOk

`func (o *PostApiSandboxesBySandboxIdActionsReset200Response) GetDeletedOk() (*map[string]int32, bool)`

GetDeletedOk returns a tuple with the Deleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleted

`func (o *PostApiSandboxesBySandboxIdActionsReset200Response) SetDeleted(v map[string]int32)`

SetDeleted sets Deleted field to given value.


### GetRevokedKeys

`func (o *PostApiSandboxesBySandboxIdActionsReset200Response) GetRevokedKeys() int32`

GetRevokedKeys returns the RevokedKeys field if non-nil, zero value otherwise.

### GetRevokedKeysOk

`func (o *PostApiSandboxesBySandboxIdActionsReset200Response) GetRevokedKeysOk() (*int32, bool)`

GetRevokedKeysOk returns a tuple with the RevokedKeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevokedKeys

`func (o *PostApiSandboxesBySandboxIdActionsReset200Response) SetRevokedKeys(v int32)`

SetRevokedKeys sets RevokedKeys field to given value.


### GetAdminKey

`func (o *PostApiSandboxesBySandboxIdActionsReset200Response) GetAdminKey() PostApiSandboxes201ResponseAdminKey`

GetAdminKey returns the AdminKey field if non-nil, zero value otherwise.

### GetAdminKeyOk

`func (o *PostApiSandboxesBySandboxIdActionsReset200Response) GetAdminKeyOk() (*PostApiSandboxes201ResponseAdminKey, bool)`

GetAdminKeyOk returns a tuple with the AdminKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminKey

`func (o *PostApiSandboxesBySandboxIdActionsReset200Response) SetAdminKey(v PostApiSandboxes201ResponseAdminKey)`

SetAdminKey sets AdminKey field to given value.


### SetAdminKeyNil

`func (o *PostApiSandboxesBySandboxIdActionsReset200Response) SetAdminKeyNil(b bool)`

 SetAdminKeyNil sets the value for AdminKey to be an explicit nil

### UnsetAdminKey
`func (o *PostApiSandboxesBySandboxIdActionsReset200Response) UnsetAdminKey()`

UnsetAdminKey ensures that no value is present for AdminKey, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


