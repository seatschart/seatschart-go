# PostApiWorkspaces200ResponseTemplate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EmbedOrigins** | **[]string** |  | 
**SkippedEmbedOrigins** | **[]string** | Template embed domains beyond the new workspace plan limit (not copied). | 
**Webhooks** | [**[]PostApiWorkspaces200ResponseTemplateWebhooksInner**](PostApiWorkspaces200ResponseTemplateWebhooksInner.md) |  | 

## Methods

### NewPostApiWorkspaces200ResponseTemplate

`func NewPostApiWorkspaces200ResponseTemplate(embedOrigins []string, skippedEmbedOrigins []string, webhooks []PostApiWorkspaces200ResponseTemplateWebhooksInner, ) *PostApiWorkspaces200ResponseTemplate`

NewPostApiWorkspaces200ResponseTemplate instantiates a new PostApiWorkspaces200ResponseTemplate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostApiWorkspaces200ResponseTemplateWithDefaults

`func NewPostApiWorkspaces200ResponseTemplateWithDefaults() *PostApiWorkspaces200ResponseTemplate`

NewPostApiWorkspaces200ResponseTemplateWithDefaults instantiates a new PostApiWorkspaces200ResponseTemplate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmbedOrigins

`func (o *PostApiWorkspaces200ResponseTemplate) GetEmbedOrigins() []string`

GetEmbedOrigins returns the EmbedOrigins field if non-nil, zero value otherwise.

### GetEmbedOriginsOk

`func (o *PostApiWorkspaces200ResponseTemplate) GetEmbedOriginsOk() (*[]string, bool)`

GetEmbedOriginsOk returns a tuple with the EmbedOrigins field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmbedOrigins

`func (o *PostApiWorkspaces200ResponseTemplate) SetEmbedOrigins(v []string)`

SetEmbedOrigins sets EmbedOrigins field to given value.


### GetSkippedEmbedOrigins

`func (o *PostApiWorkspaces200ResponseTemplate) GetSkippedEmbedOrigins() []string`

GetSkippedEmbedOrigins returns the SkippedEmbedOrigins field if non-nil, zero value otherwise.

### GetSkippedEmbedOriginsOk

`func (o *PostApiWorkspaces200ResponseTemplate) GetSkippedEmbedOriginsOk() (*[]string, bool)`

GetSkippedEmbedOriginsOk returns a tuple with the SkippedEmbedOrigins field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkippedEmbedOrigins

`func (o *PostApiWorkspaces200ResponseTemplate) SetSkippedEmbedOrigins(v []string)`

SetSkippedEmbedOrigins sets SkippedEmbedOrigins field to given value.


### GetWebhooks

`func (o *PostApiWorkspaces200ResponseTemplate) GetWebhooks() []PostApiWorkspaces200ResponseTemplateWebhooksInner`

GetWebhooks returns the Webhooks field if non-nil, zero value otherwise.

### GetWebhooksOk

`func (o *PostApiWorkspaces200ResponseTemplate) GetWebhooksOk() (*[]PostApiWorkspaces200ResponseTemplateWebhooksInner, bool)`

GetWebhooksOk returns a tuple with the Webhooks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhooks

`func (o *PostApiWorkspaces200ResponseTemplate) SetWebhooks(v []PostApiWorkspaces200ResponseTemplateWebhooksInner)`

SetWebhooks sets Webhooks field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


