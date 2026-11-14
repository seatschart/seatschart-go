# PostApiSandboxes201Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Sandbox** | [**GetApiSandboxes200ResponseItemsInner**](GetApiSandboxes200ResponseItemsInner.md) |  | 
**AdminKey** | Pointer to [**NullablePostApiSandboxes201ResponseAdminKey**](PostApiSandboxes201ResponseAdminKey.md) |  | [optional] 
**Copied** | Pointer to [**NullablePostApiSandboxes201ResponseCopied**](PostApiSandboxes201ResponseCopied.md) |  | [optional] 

## Methods

### NewPostApiSandboxes201Response

`func NewPostApiSandboxes201Response(sandbox GetApiSandboxes200ResponseItemsInner, ) *PostApiSandboxes201Response`

NewPostApiSandboxes201Response instantiates a new PostApiSandboxes201Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostApiSandboxes201ResponseWithDefaults

`func NewPostApiSandboxes201ResponseWithDefaults() *PostApiSandboxes201Response`

NewPostApiSandboxes201ResponseWithDefaults instantiates a new PostApiSandboxes201Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSandbox

`func (o *PostApiSandboxes201Response) GetSandbox() GetApiSandboxes200ResponseItemsInner`

GetSandbox returns the Sandbox field if non-nil, zero value otherwise.

### GetSandboxOk

`func (o *PostApiSandboxes201Response) GetSandboxOk() (*GetApiSandboxes200ResponseItemsInner, bool)`

GetSandboxOk returns a tuple with the Sandbox field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSandbox

`func (o *PostApiSandboxes201Response) SetSandbox(v GetApiSandboxes200ResponseItemsInner)`

SetSandbox sets Sandbox field to given value.


### GetAdminKey

`func (o *PostApiSandboxes201Response) GetAdminKey() PostApiSandboxes201ResponseAdminKey`

GetAdminKey returns the AdminKey field if non-nil, zero value otherwise.

### GetAdminKeyOk

`func (o *PostApiSandboxes201Response) GetAdminKeyOk() (*PostApiSandboxes201ResponseAdminKey, bool)`

GetAdminKeyOk returns a tuple with the AdminKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminKey

`func (o *PostApiSandboxes201Response) SetAdminKey(v PostApiSandboxes201ResponseAdminKey)`

SetAdminKey sets AdminKey field to given value.

### HasAdminKey

`func (o *PostApiSandboxes201Response) HasAdminKey() bool`

HasAdminKey returns a boolean if a field has been set.

### SetAdminKeyNil

`func (o *PostApiSandboxes201Response) SetAdminKeyNil(b bool)`

 SetAdminKeyNil sets the value for AdminKey to be an explicit nil

### UnsetAdminKey
`func (o *PostApiSandboxes201Response) UnsetAdminKey()`

UnsetAdminKey ensures that no value is present for AdminKey, not even an explicit nil
### GetCopied

`func (o *PostApiSandboxes201Response) GetCopied() PostApiSandboxes201ResponseCopied`

GetCopied returns the Copied field if non-nil, zero value otherwise.

### GetCopiedOk

`func (o *PostApiSandboxes201Response) GetCopiedOk() (*PostApiSandboxes201ResponseCopied, bool)`

GetCopiedOk returns a tuple with the Copied field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCopied

`func (o *PostApiSandboxes201Response) SetCopied(v PostApiSandboxes201ResponseCopied)`

SetCopied sets Copied field to given value.

### HasCopied

`func (o *PostApiSandboxes201Response) HasCopied() bool`

HasCopied returns a boolean if a field has been set.

### SetCopiedNil

`func (o *PostApiSandboxes201Response) SetCopiedNil(b bool)`

 SetCopiedNil sets the value for Copied to be an explicit nil

### UnsetCopied
`func (o *PostApiSandboxes201Response) UnsetCopied()`

UnsetCopied ensures that no value is present for Copied, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


