# PostApiSandboxesBySandboxIdAdminKeyRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**RevokeExisting** | Pointer to **bool** | Rotate: revoke every other key of the sandbox in the same transaction. Default false. | [optional] 

## Methods

### NewPostApiSandboxesBySandboxIdAdminKeyRequest

`func NewPostApiSandboxesBySandboxIdAdminKeyRequest() *PostApiSandboxesBySandboxIdAdminKeyRequest`

NewPostApiSandboxesBySandboxIdAdminKeyRequest instantiates a new PostApiSandboxesBySandboxIdAdminKeyRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostApiSandboxesBySandboxIdAdminKeyRequestWithDefaults

`func NewPostApiSandboxesBySandboxIdAdminKeyRequestWithDefaults() *PostApiSandboxesBySandboxIdAdminKeyRequest`

NewPostApiSandboxesBySandboxIdAdminKeyRequestWithDefaults instantiates a new PostApiSandboxesBySandboxIdAdminKeyRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *PostApiSandboxesBySandboxIdAdminKeyRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PostApiSandboxesBySandboxIdAdminKeyRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PostApiSandboxesBySandboxIdAdminKeyRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PostApiSandboxesBySandboxIdAdminKeyRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetRevokeExisting

`func (o *PostApiSandboxesBySandboxIdAdminKeyRequest) GetRevokeExisting() bool`

GetRevokeExisting returns the RevokeExisting field if non-nil, zero value otherwise.

### GetRevokeExistingOk

`func (o *PostApiSandboxesBySandboxIdAdminKeyRequest) GetRevokeExistingOk() (*bool, bool)`

GetRevokeExistingOk returns a tuple with the RevokeExisting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevokeExisting

`func (o *PostApiSandboxesBySandboxIdAdminKeyRequest) SetRevokeExisting(v bool)`

SetRevokeExisting sets RevokeExisting field to given value.

### HasRevokeExisting

`func (o *PostApiSandboxesBySandboxIdAdminKeyRequest) HasRevokeExisting() bool`

HasRevokeExisting returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


