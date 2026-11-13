# PostApiAuth2faEnrollStartRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PendingToken** | **string** | Pending sign-in token returned by &#x60;POST /auth/login&#x60;. | 

## Methods

### NewPostApiAuth2faEnrollStartRequest

`func NewPostApiAuth2faEnrollStartRequest(pendingToken string, ) *PostApiAuth2faEnrollStartRequest`

NewPostApiAuth2faEnrollStartRequest instantiates a new PostApiAuth2faEnrollStartRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostApiAuth2faEnrollStartRequestWithDefaults

`func NewPostApiAuth2faEnrollStartRequestWithDefaults() *PostApiAuth2faEnrollStartRequest`

NewPostApiAuth2faEnrollStartRequestWithDefaults instantiates a new PostApiAuth2faEnrollStartRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPendingToken

`func (o *PostApiAuth2faEnrollStartRequest) GetPendingToken() string`

GetPendingToken returns the PendingToken field if non-nil, zero value otherwise.

### GetPendingTokenOk

`func (o *PostApiAuth2faEnrollStartRequest) GetPendingTokenOk() (*string, bool)`

GetPendingTokenOk returns a tuple with the PendingToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPendingToken

`func (o *PostApiAuth2faEnrollStartRequest) SetPendingToken(v string)`

SetPendingToken sets PendingToken field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


