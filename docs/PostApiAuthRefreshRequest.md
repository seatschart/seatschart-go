# PostApiAuthRefreshRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RefreshToken** | **string** |  | 
**Audience** | Pointer to **string** |  | [optional] 

## Methods

### NewPostApiAuthRefreshRequest

`func NewPostApiAuthRefreshRequest(refreshToken string, ) *PostApiAuthRefreshRequest`

NewPostApiAuthRefreshRequest instantiates a new PostApiAuthRefreshRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostApiAuthRefreshRequestWithDefaults

`func NewPostApiAuthRefreshRequestWithDefaults() *PostApiAuthRefreshRequest`

NewPostApiAuthRefreshRequestWithDefaults instantiates a new PostApiAuthRefreshRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRefreshToken

`func (o *PostApiAuthRefreshRequest) GetRefreshToken() string`

GetRefreshToken returns the RefreshToken field if non-nil, zero value otherwise.

### GetRefreshTokenOk

`func (o *PostApiAuthRefreshRequest) GetRefreshTokenOk() (*string, bool)`

GetRefreshTokenOk returns a tuple with the RefreshToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefreshToken

`func (o *PostApiAuthRefreshRequest) SetRefreshToken(v string)`

SetRefreshToken sets RefreshToken field to given value.


### GetAudience

`func (o *PostApiAuthRefreshRequest) GetAudience() string`

GetAudience returns the Audience field if non-nil, zero value otherwise.

### GetAudienceOk

`func (o *PostApiAuthRefreshRequest) GetAudienceOk() (*string, bool)`

GetAudienceOk returns a tuple with the Audience field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAudience

`func (o *PostApiAuthRefreshRequest) SetAudience(v string)`

SetAudience sets Audience field to given value.

### HasAudience

`func (o *PostApiAuthRefreshRequest) HasAudience() bool`

HasAudience returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


