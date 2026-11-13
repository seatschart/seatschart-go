# DeleteApiSandboxesBySandboxId200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SandboxId** | **string** |  | 
**Deleted** | **bool** |  | 
**Removed** | **map[string]int32** |  | 

## Methods

### NewDeleteApiSandboxesBySandboxId200Response

`func NewDeleteApiSandboxesBySandboxId200Response(sandboxId string, deleted bool, removed map[string]int32, ) *DeleteApiSandboxesBySandboxId200Response`

NewDeleteApiSandboxesBySandboxId200Response instantiates a new DeleteApiSandboxesBySandboxId200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeleteApiSandboxesBySandboxId200ResponseWithDefaults

`func NewDeleteApiSandboxesBySandboxId200ResponseWithDefaults() *DeleteApiSandboxesBySandboxId200Response`

NewDeleteApiSandboxesBySandboxId200ResponseWithDefaults instantiates a new DeleteApiSandboxesBySandboxId200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSandboxId

`func (o *DeleteApiSandboxesBySandboxId200Response) GetSandboxId() string`

GetSandboxId returns the SandboxId field if non-nil, zero value otherwise.

### GetSandboxIdOk

`func (o *DeleteApiSandboxesBySandboxId200Response) GetSandboxIdOk() (*string, bool)`

GetSandboxIdOk returns a tuple with the SandboxId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSandboxId

`func (o *DeleteApiSandboxesBySandboxId200Response) SetSandboxId(v string)`

SetSandboxId sets SandboxId field to given value.


### GetDeleted

`func (o *DeleteApiSandboxesBySandboxId200Response) GetDeleted() bool`

GetDeleted returns the Deleted field if non-nil, zero value otherwise.

### GetDeletedOk

`func (o *DeleteApiSandboxesBySandboxId200Response) GetDeletedOk() (*bool, bool)`

GetDeletedOk returns a tuple with the Deleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleted

`func (o *DeleteApiSandboxesBySandboxId200Response) SetDeleted(v bool)`

SetDeleted sets Deleted field to given value.


### GetRemoved

`func (o *DeleteApiSandboxesBySandboxId200Response) GetRemoved() map[string]int32`

GetRemoved returns the Removed field if non-nil, zero value otherwise.

### GetRemovedOk

`func (o *DeleteApiSandboxesBySandboxId200Response) GetRemovedOk() (*map[string]int32, bool)`

GetRemovedOk returns a tuple with the Removed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoved

`func (o *DeleteApiSandboxesBySandboxId200Response) SetRemoved(v map[string]int32)`

SetRemoved sets Removed field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


