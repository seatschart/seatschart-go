# PostApiWorkspaces200ResponseTemplateWebhooksInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Url** | **string** |  | 
**EventTypes** | **[]string** |  | 
**IsActive** | **bool** |  | 
**SigningSecret** | **NullableString** | NEW secret of this endpoint, shown once; null on a replay. | 

## Methods

### NewPostApiWorkspaces200ResponseTemplateWebhooksInner

`func NewPostApiWorkspaces200ResponseTemplateWebhooksInner(id string, url string, eventTypes []string, isActive bool, signingSecret NullableString, ) *PostApiWorkspaces200ResponseTemplateWebhooksInner`

NewPostApiWorkspaces200ResponseTemplateWebhooksInner instantiates a new PostApiWorkspaces200ResponseTemplateWebhooksInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostApiWorkspaces200ResponseTemplateWebhooksInnerWithDefaults

`func NewPostApiWorkspaces200ResponseTemplateWebhooksInnerWithDefaults() *PostApiWorkspaces200ResponseTemplateWebhooksInner`

NewPostApiWorkspaces200ResponseTemplateWebhooksInnerWithDefaults instantiates a new PostApiWorkspaces200ResponseTemplateWebhooksInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PostApiWorkspaces200ResponseTemplateWebhooksInner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PostApiWorkspaces200ResponseTemplateWebhooksInner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PostApiWorkspaces200ResponseTemplateWebhooksInner) SetId(v string)`

SetId sets Id field to given value.


### GetUrl

`func (o *PostApiWorkspaces200ResponseTemplateWebhooksInner) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *PostApiWorkspaces200ResponseTemplateWebhooksInner) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *PostApiWorkspaces200ResponseTemplateWebhooksInner) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetEventTypes

`func (o *PostApiWorkspaces200ResponseTemplateWebhooksInner) GetEventTypes() []string`

GetEventTypes returns the EventTypes field if non-nil, zero value otherwise.

### GetEventTypesOk

`func (o *PostApiWorkspaces200ResponseTemplateWebhooksInner) GetEventTypesOk() (*[]string, bool)`

GetEventTypesOk returns a tuple with the EventTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventTypes

`func (o *PostApiWorkspaces200ResponseTemplateWebhooksInner) SetEventTypes(v []string)`

SetEventTypes sets EventTypes field to given value.


### GetIsActive

`func (o *PostApiWorkspaces200ResponseTemplateWebhooksInner) GetIsActive() bool`

GetIsActive returns the IsActive field if non-nil, zero value otherwise.

### GetIsActiveOk

`func (o *PostApiWorkspaces200ResponseTemplateWebhooksInner) GetIsActiveOk() (*bool, bool)`

GetIsActiveOk returns a tuple with the IsActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsActive

`func (o *PostApiWorkspaces200ResponseTemplateWebhooksInner) SetIsActive(v bool)`

SetIsActive sets IsActive field to given value.


### GetSigningSecret

`func (o *PostApiWorkspaces200ResponseTemplateWebhooksInner) GetSigningSecret() string`

GetSigningSecret returns the SigningSecret field if non-nil, zero value otherwise.

### GetSigningSecretOk

`func (o *PostApiWorkspaces200ResponseTemplateWebhooksInner) GetSigningSecretOk() (*string, bool)`

GetSigningSecretOk returns a tuple with the SigningSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSigningSecret

`func (o *PostApiWorkspaces200ResponseTemplateWebhooksInner) SetSigningSecret(v string)`

SetSigningSecret sets SigningSecret field to given value.


### SetSigningSecretNil

`func (o *PostApiWorkspaces200ResponseTemplateWebhooksInner) SetSigningSecretNil(b bool)`

 SetSigningSecretNil sets the value for SigningSecret to be an explicit nil

### UnsetSigningSecret
`func (o *PostApiWorkspaces200ResponseTemplateWebhooksInner) UnsetSigningSecret()`

UnsetSigningSecret ensures that no value is present for SigningSecret, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


