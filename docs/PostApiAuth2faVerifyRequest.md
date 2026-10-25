# PostApiAuth2faVerifyRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PendingToken** | **string** | Pending sign-in token returned by &#x60;POST /auth/login&#x60;. | 
**Code** | Pointer to **string** | 6-digit code from the authenticator app. | [optional] 
**RecoveryCode** | Pointer to **string** | One-time recovery code (&#x60;xxxxx-xxxxx&#x60;). | [optional] 

## Methods

### NewPostApiAuth2faVerifyRequest

`func NewPostApiAuth2faVerifyRequest(pendingToken string, ) *PostApiAuth2faVerifyRequest`

NewPostApiAuth2faVerifyRequest instantiates a new PostApiAuth2faVerifyRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostApiAuth2faVerifyRequestWithDefaults

`func NewPostApiAuth2faVerifyRequestWithDefaults() *PostApiAuth2faVerifyRequest`

NewPostApiAuth2faVerifyRequestWithDefaults instantiates a new PostApiAuth2faVerifyRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPendingToken

`func (o *PostApiAuth2faVerifyRequest) GetPendingToken() string`

GetPendingToken returns the PendingToken field if non-nil, zero value otherwise.

### GetPendingTokenOk

`func (o *PostApiAuth2faVerifyRequest) GetPendingTokenOk() (*string, bool)`

GetPendingTokenOk returns a tuple with the PendingToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPendingToken

`func (o *PostApiAuth2faVerifyRequest) SetPendingToken(v string)`

SetPendingToken sets PendingToken field to given value.


### GetCode

`func (o *PostApiAuth2faVerifyRequest) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *PostApiAuth2faVerifyRequest) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *PostApiAuth2faVerifyRequest) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *PostApiAuth2faVerifyRequest) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetRecoveryCode

`func (o *PostApiAuth2faVerifyRequest) GetRecoveryCode() string`

GetRecoveryCode returns the RecoveryCode field if non-nil, zero value otherwise.

### GetRecoveryCodeOk

`func (o *PostApiAuth2faVerifyRequest) GetRecoveryCodeOk() (*string, bool)`

GetRecoveryCodeOk returns a tuple with the RecoveryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecoveryCode

`func (o *PostApiAuth2faVerifyRequest) SetRecoveryCode(v string)`

SetRecoveryCode sets RecoveryCode field to given value.

### HasRecoveryCode

`func (o *PostApiAuth2faVerifyRequest) HasRecoveryCode() bool`

HasRecoveryCode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


