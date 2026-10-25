# PostApiAuth2faEnrollConfirmRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PendingToken** | **string** | Pending sign-in token returned by &#x60;POST /auth/login&#x60;. | 
**Code** | **string** | 6-digit code from the authenticator app. | 

## Methods

### NewPostApiAuth2faEnrollConfirmRequest

`func NewPostApiAuth2faEnrollConfirmRequest(pendingToken string, code string, ) *PostApiAuth2faEnrollConfirmRequest`

NewPostApiAuth2faEnrollConfirmRequest instantiates a new PostApiAuth2faEnrollConfirmRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostApiAuth2faEnrollConfirmRequestWithDefaults

`func NewPostApiAuth2faEnrollConfirmRequestWithDefaults() *PostApiAuth2faEnrollConfirmRequest`

NewPostApiAuth2faEnrollConfirmRequestWithDefaults instantiates a new PostApiAuth2faEnrollConfirmRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPendingToken

`func (o *PostApiAuth2faEnrollConfirmRequest) GetPendingToken() string`

GetPendingToken returns the PendingToken field if non-nil, zero value otherwise.

### GetPendingTokenOk

`func (o *PostApiAuth2faEnrollConfirmRequest) GetPendingTokenOk() (*string, bool)`

GetPendingTokenOk returns a tuple with the PendingToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPendingToken

`func (o *PostApiAuth2faEnrollConfirmRequest) SetPendingToken(v string)`

SetPendingToken sets PendingToken field to given value.


### GetCode

`func (o *PostApiAuth2faEnrollConfirmRequest) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *PostApiAuth2faEnrollConfirmRequest) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *PostApiAuth2faEnrollConfirmRequest) SetCode(v string)`

SetCode sets Code field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


