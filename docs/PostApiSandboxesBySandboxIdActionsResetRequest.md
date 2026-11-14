# PostApiSandboxesBySandboxIdActionsResetRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RotateKeys** | Pointer to **bool** | Revoke every API key of the sandbox and mint a new full-access key (returned once). Default false: keys are kept. | [optional] 
**KeepWebhooks** | Pointer to **bool** | Keep the webhook endpoints (their delivery log is always cleared). Default true. | [optional] 

## Methods

### NewPostApiSandboxesBySandboxIdActionsResetRequest

`func NewPostApiSandboxesBySandboxIdActionsResetRequest() *PostApiSandboxesBySandboxIdActionsResetRequest`

NewPostApiSandboxesBySandboxIdActionsResetRequest instantiates a new PostApiSandboxesBySandboxIdActionsResetRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostApiSandboxesBySandboxIdActionsResetRequestWithDefaults

`func NewPostApiSandboxesBySandboxIdActionsResetRequestWithDefaults() *PostApiSandboxesBySandboxIdActionsResetRequest`

NewPostApiSandboxesBySandboxIdActionsResetRequestWithDefaults instantiates a new PostApiSandboxesBySandboxIdActionsResetRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRotateKeys

`func (o *PostApiSandboxesBySandboxIdActionsResetRequest) GetRotateKeys() bool`

GetRotateKeys returns the RotateKeys field if non-nil, zero value otherwise.

### GetRotateKeysOk

`func (o *PostApiSandboxesBySandboxIdActionsResetRequest) GetRotateKeysOk() (*bool, bool)`

GetRotateKeysOk returns a tuple with the RotateKeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRotateKeys

`func (o *PostApiSandboxesBySandboxIdActionsResetRequest) SetRotateKeys(v bool)`

SetRotateKeys sets RotateKeys field to given value.

### HasRotateKeys

`func (o *PostApiSandboxesBySandboxIdActionsResetRequest) HasRotateKeys() bool`

HasRotateKeys returns a boolean if a field has been set.

### GetKeepWebhooks

`func (o *PostApiSandboxesBySandboxIdActionsResetRequest) GetKeepWebhooks() bool`

GetKeepWebhooks returns the KeepWebhooks field if non-nil, zero value otherwise.

### GetKeepWebhooksOk

`func (o *PostApiSandboxesBySandboxIdActionsResetRequest) GetKeepWebhooksOk() (*bool, bool)`

GetKeepWebhooksOk returns a tuple with the KeepWebhooks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeepWebhooks

`func (o *PostApiSandboxesBySandboxIdActionsResetRequest) SetKeepWebhooks(v bool)`

SetKeepWebhooks sets KeepWebhooks field to given value.

### HasKeepWebhooks

`func (o *PostApiSandboxesBySandboxIdActionsResetRequest) HasKeepWebhooks() bool`

HasKeepWebhooks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


