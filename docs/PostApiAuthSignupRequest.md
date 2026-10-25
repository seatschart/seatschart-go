# PostApiAuthSignupRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Email** | **string** |  | 
**Password** | **string** |  | 
**WorkspaceName** | Pointer to **string** |  | [optional] 

## Methods

### NewPostApiAuthSignupRequest

`func NewPostApiAuthSignupRequest(email string, password string, ) *PostApiAuthSignupRequest`

NewPostApiAuthSignupRequest instantiates a new PostApiAuthSignupRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostApiAuthSignupRequestWithDefaults

`func NewPostApiAuthSignupRequestWithDefaults() *PostApiAuthSignupRequest`

NewPostApiAuthSignupRequestWithDefaults instantiates a new PostApiAuthSignupRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmail

`func (o *PostApiAuthSignupRequest) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *PostApiAuthSignupRequest) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *PostApiAuthSignupRequest) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetPassword

`func (o *PostApiAuthSignupRequest) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *PostApiAuthSignupRequest) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *PostApiAuthSignupRequest) SetPassword(v string)`

SetPassword sets Password field to given value.


### GetWorkspaceName

`func (o *PostApiAuthSignupRequest) GetWorkspaceName() string`

GetWorkspaceName returns the WorkspaceName field if non-nil, zero value otherwise.

### GetWorkspaceNameOk

`func (o *PostApiAuthSignupRequest) GetWorkspaceNameOk() (*string, bool)`

GetWorkspaceNameOk returns a tuple with the WorkspaceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkspaceName

`func (o *PostApiAuthSignupRequest) SetWorkspaceName(v string)`

SetWorkspaceName sets WorkspaceName field to given value.

### HasWorkspaceName

`func (o *PostApiAuthSignupRequest) HasWorkspaceName() bool`

HasWorkspaceName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


