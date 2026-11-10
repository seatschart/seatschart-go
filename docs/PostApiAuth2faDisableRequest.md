# PostApiAuth2faDisableRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | Pointer to **string** | 6-digit code from the authenticator app. | [optional] 
**RecoveryCode** | Pointer to **string** | One-time recovery code (&#x60;xxxxx-xxxxx&#x60;). | [optional] 

## Methods

### NewPostApiAuth2faDisableRequest

`func NewPostApiAuth2faDisableRequest() *PostApiAuth2faDisableRequest`

NewPostApiAuth2faDisableRequest instantiates a new PostApiAuth2faDisableRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostApiAuth2faDisableRequestWithDefaults

`func NewPostApiAuth2faDisableRequestWithDefaults() *PostApiAuth2faDisableRequest`

NewPostApiAuth2faDisableRequestWithDefaults instantiates a new PostApiAuth2faDisableRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *PostApiAuth2faDisableRequest) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *PostApiAuth2faDisableRequest) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *PostApiAuth2faDisableRequest) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *PostApiAuth2faDisableRequest) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetRecoveryCode

`func (o *PostApiAuth2faDisableRequest) GetRecoveryCode() string`

GetRecoveryCode returns the RecoveryCode field if non-nil, zero value otherwise.

### GetRecoveryCodeOk

`func (o *PostApiAuth2faDisableRequest) GetRecoveryCodeOk() (*string, bool)`

GetRecoveryCodeOk returns a tuple with the RecoveryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecoveryCode

`func (o *PostApiAuth2faDisableRequest) SetRecoveryCode(v string)`

SetRecoveryCode sets RecoveryCode field to given value.

### HasRecoveryCode

`func (o *PostApiAuth2faDisableRequest) HasRecoveryCode() bool`

HasRecoveryCode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


