# PostApiWorkspaces200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TenantId** | **string** | Id of the new workspace (also its &#x60;workspaceId&#x60;). | 
**DisplayName** | **string** |  | 
**PublicKey** | **string** | Public renderer key (&#x60;pub_…&#x60;): not a secret, not a credential. | 
**Role** | **string** |  | 
**IsTest** | **bool** |  | 
**Kind** | **string** |  | 
**ParentWorkspaceId** | **NullableString** |  | 
**CopiedCharts** | **int32** |  | 
**ExternalId** | **NullableString** |  | 
**CreatedByApi** | **bool** | Provisioned by a workspace admin key: billed to the platform account, no upgrade prompts for the organizer. | 
**TemplateWorkspaceId** | **NullableString** |  | 
**Template** | [**NullablePostApiWorkspaces200ResponseTemplate**](PostApiWorkspaces200ResponseTemplate.md) |  | 
**ApiKey** | [**NullablePostApiWorkspaces200ResponseApiKey**](PostApiWorkspaces200ResponseApiKey.md) |  | 
**IdempotentReplay** | **bool** | True when this is the stored result of an earlier call with the same Idempotency-Key (secrets are null). | 
**Warning** | Pointer to **string** |  | [optional] 

## Methods

### NewPostApiWorkspaces200Response

`func NewPostApiWorkspaces200Response(tenantId string, displayName string, publicKey string, role string, isTest bool, kind string, parentWorkspaceId NullableString, copiedCharts int32, externalId NullableString, createdByApi bool, templateWorkspaceId NullableString, template NullablePostApiWorkspaces200ResponseTemplate, apiKey NullablePostApiWorkspaces200ResponseApiKey, idempotentReplay bool, ) *PostApiWorkspaces200Response`

NewPostApiWorkspaces200Response instantiates a new PostApiWorkspaces200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostApiWorkspaces200ResponseWithDefaults

`func NewPostApiWorkspaces200ResponseWithDefaults() *PostApiWorkspaces200Response`

NewPostApiWorkspaces200ResponseWithDefaults instantiates a new PostApiWorkspaces200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTenantId

`func (o *PostApiWorkspaces200Response) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *PostApiWorkspaces200Response) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *PostApiWorkspaces200Response) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.


### GetDisplayName

`func (o *PostApiWorkspaces200Response) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *PostApiWorkspaces200Response) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *PostApiWorkspaces200Response) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.


### GetPublicKey

`func (o *PostApiWorkspaces200Response) GetPublicKey() string`

GetPublicKey returns the PublicKey field if non-nil, zero value otherwise.

### GetPublicKeyOk

`func (o *PostApiWorkspaces200Response) GetPublicKeyOk() (*string, bool)`

GetPublicKeyOk returns a tuple with the PublicKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicKey

`func (o *PostApiWorkspaces200Response) SetPublicKey(v string)`

SetPublicKey sets PublicKey field to given value.


### GetRole

`func (o *PostApiWorkspaces200Response) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *PostApiWorkspaces200Response) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *PostApiWorkspaces200Response) SetRole(v string)`

SetRole sets Role field to given value.


### GetIsTest

`func (o *PostApiWorkspaces200Response) GetIsTest() bool`

GetIsTest returns the IsTest field if non-nil, zero value otherwise.

### GetIsTestOk

`func (o *PostApiWorkspaces200Response) GetIsTestOk() (*bool, bool)`

GetIsTestOk returns a tuple with the IsTest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsTest

`func (o *PostApiWorkspaces200Response) SetIsTest(v bool)`

SetIsTest sets IsTest field to given value.


### GetKind

`func (o *PostApiWorkspaces200Response) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *PostApiWorkspaces200Response) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *PostApiWorkspaces200Response) SetKind(v string)`

SetKind sets Kind field to given value.


### GetParentWorkspaceId

`func (o *PostApiWorkspaces200Response) GetParentWorkspaceId() string`

GetParentWorkspaceId returns the ParentWorkspaceId field if non-nil, zero value otherwise.

### GetParentWorkspaceIdOk

`func (o *PostApiWorkspaces200Response) GetParentWorkspaceIdOk() (*string, bool)`

GetParentWorkspaceIdOk returns a tuple with the ParentWorkspaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentWorkspaceId

`func (o *PostApiWorkspaces200Response) SetParentWorkspaceId(v string)`

SetParentWorkspaceId sets ParentWorkspaceId field to given value.


### SetParentWorkspaceIdNil

`func (o *PostApiWorkspaces200Response) SetParentWorkspaceIdNil(b bool)`

 SetParentWorkspaceIdNil sets the value for ParentWorkspaceId to be an explicit nil

### UnsetParentWorkspaceId
`func (o *PostApiWorkspaces200Response) UnsetParentWorkspaceId()`

UnsetParentWorkspaceId ensures that no value is present for ParentWorkspaceId, not even an explicit nil
### GetCopiedCharts

`func (o *PostApiWorkspaces200Response) GetCopiedCharts() int32`

GetCopiedCharts returns the CopiedCharts field if non-nil, zero value otherwise.

### GetCopiedChartsOk

`func (o *PostApiWorkspaces200Response) GetCopiedChartsOk() (*int32, bool)`

GetCopiedChartsOk returns a tuple with the CopiedCharts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCopiedCharts

`func (o *PostApiWorkspaces200Response) SetCopiedCharts(v int32)`

SetCopiedCharts sets CopiedCharts field to given value.


### GetExternalId

`func (o *PostApiWorkspaces200Response) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *PostApiWorkspaces200Response) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *PostApiWorkspaces200Response) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.


### SetExternalIdNil

`func (o *PostApiWorkspaces200Response) SetExternalIdNil(b bool)`

 SetExternalIdNil sets the value for ExternalId to be an explicit nil

### UnsetExternalId
`func (o *PostApiWorkspaces200Response) UnsetExternalId()`

UnsetExternalId ensures that no value is present for ExternalId, not even an explicit nil
### GetCreatedByApi

`func (o *PostApiWorkspaces200Response) GetCreatedByApi() bool`

GetCreatedByApi returns the CreatedByApi field if non-nil, zero value otherwise.

### GetCreatedByApiOk

`func (o *PostApiWorkspaces200Response) GetCreatedByApiOk() (*bool, bool)`

GetCreatedByApiOk returns a tuple with the CreatedByApi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedByApi

`func (o *PostApiWorkspaces200Response) SetCreatedByApi(v bool)`

SetCreatedByApi sets CreatedByApi field to given value.


### GetTemplateWorkspaceId

`func (o *PostApiWorkspaces200Response) GetTemplateWorkspaceId() string`

GetTemplateWorkspaceId returns the TemplateWorkspaceId field if non-nil, zero value otherwise.

### GetTemplateWorkspaceIdOk

`func (o *PostApiWorkspaces200Response) GetTemplateWorkspaceIdOk() (*string, bool)`

GetTemplateWorkspaceIdOk returns a tuple with the TemplateWorkspaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateWorkspaceId

`func (o *PostApiWorkspaces200Response) SetTemplateWorkspaceId(v string)`

SetTemplateWorkspaceId sets TemplateWorkspaceId field to given value.


### SetTemplateWorkspaceIdNil

`func (o *PostApiWorkspaces200Response) SetTemplateWorkspaceIdNil(b bool)`

 SetTemplateWorkspaceIdNil sets the value for TemplateWorkspaceId to be an explicit nil

### UnsetTemplateWorkspaceId
`func (o *PostApiWorkspaces200Response) UnsetTemplateWorkspaceId()`

UnsetTemplateWorkspaceId ensures that no value is present for TemplateWorkspaceId, not even an explicit nil
### GetTemplate

`func (o *PostApiWorkspaces200Response) GetTemplate() PostApiWorkspaces200ResponseTemplate`

GetTemplate returns the Template field if non-nil, zero value otherwise.

### GetTemplateOk

`func (o *PostApiWorkspaces200Response) GetTemplateOk() (*PostApiWorkspaces200ResponseTemplate, bool)`

GetTemplateOk returns a tuple with the Template field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplate

`func (o *PostApiWorkspaces200Response) SetTemplate(v PostApiWorkspaces200ResponseTemplate)`

SetTemplate sets Template field to given value.


### SetTemplateNil

`func (o *PostApiWorkspaces200Response) SetTemplateNil(b bool)`

 SetTemplateNil sets the value for Template to be an explicit nil

### UnsetTemplate
`func (o *PostApiWorkspaces200Response) UnsetTemplate()`

UnsetTemplate ensures that no value is present for Template, not even an explicit nil
### GetApiKey

`func (o *PostApiWorkspaces200Response) GetApiKey() PostApiWorkspaces200ResponseApiKey`

GetApiKey returns the ApiKey field if non-nil, zero value otherwise.

### GetApiKeyOk

`func (o *PostApiWorkspaces200Response) GetApiKeyOk() (*PostApiWorkspaces200ResponseApiKey, bool)`

GetApiKeyOk returns a tuple with the ApiKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiKey

`func (o *PostApiWorkspaces200Response) SetApiKey(v PostApiWorkspaces200ResponseApiKey)`

SetApiKey sets ApiKey field to given value.


### SetApiKeyNil

`func (o *PostApiWorkspaces200Response) SetApiKeyNil(b bool)`

 SetApiKeyNil sets the value for ApiKey to be an explicit nil

### UnsetApiKey
`func (o *PostApiWorkspaces200Response) UnsetApiKey()`

UnsetApiKey ensures that no value is present for ApiKey, not even an explicit nil
### GetIdempotentReplay

`func (o *PostApiWorkspaces200Response) GetIdempotentReplay() bool`

GetIdempotentReplay returns the IdempotentReplay field if non-nil, zero value otherwise.

### GetIdempotentReplayOk

`func (o *PostApiWorkspaces200Response) GetIdempotentReplayOk() (*bool, bool)`

GetIdempotentReplayOk returns a tuple with the IdempotentReplay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdempotentReplay

`func (o *PostApiWorkspaces200Response) SetIdempotentReplay(v bool)`

SetIdempotentReplay sets IdempotentReplay field to given value.


### GetWarning

`func (o *PostApiWorkspaces200Response) GetWarning() string`

GetWarning returns the Warning field if non-nil, zero value otherwise.

### GetWarningOk

`func (o *PostApiWorkspaces200Response) GetWarningOk() (*string, bool)`

GetWarningOk returns a tuple with the Warning field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarning

`func (o *PostApiWorkspaces200Response) SetWarning(v string)`

SetWarning sets Warning field to given value.

### HasWarning

`func (o *PostApiWorkspaces200Response) HasWarning() bool`

HasWarning returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


