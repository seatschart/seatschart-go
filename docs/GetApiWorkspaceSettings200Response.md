# GetApiWorkspaceSettings200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Effective** | [**GetApiWorkspaceSettings200ResponseEffective**](GetApiWorkspaceSettings200ResponseEffective.md) |  | 
**Defaults** | [**GetApiWorkspaceSettings200ResponseEffective**](GetApiWorkspaceSettings200ResponseEffective.md) |  | 
**Sources** | [**GetApiWorkspaceSettings200ResponseSources**](GetApiWorkspaceSettings200ResponseSources.md) |  | 
**Overrides** | [**GetApiWorkspaceSettings200ResponseOverrides**](GetApiWorkspaceSettings200ResponseOverrides.md) |  | 
**Bounds** | [**GetApiWorkspaceSettings200ResponseBounds**](GetApiWorkspaceSettings200ResponseBounds.md) |  | 
**RequireTwoFactor** | **bool** |  | 
**IsTest** | **bool** |  | 
**ParentWorkspaceId** | **NullableString** |  | 

## Methods

### NewGetApiWorkspaceSettings200Response

`func NewGetApiWorkspaceSettings200Response(effective GetApiWorkspaceSettings200ResponseEffective, defaults GetApiWorkspaceSettings200ResponseEffective, sources GetApiWorkspaceSettings200ResponseSources, overrides GetApiWorkspaceSettings200ResponseOverrides, bounds GetApiWorkspaceSettings200ResponseBounds, requireTwoFactor bool, isTest bool, parentWorkspaceId NullableString, ) *GetApiWorkspaceSettings200Response`

NewGetApiWorkspaceSettings200Response instantiates a new GetApiWorkspaceSettings200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetApiWorkspaceSettings200ResponseWithDefaults

`func NewGetApiWorkspaceSettings200ResponseWithDefaults() *GetApiWorkspaceSettings200Response`

NewGetApiWorkspaceSettings200ResponseWithDefaults instantiates a new GetApiWorkspaceSettings200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEffective

`func (o *GetApiWorkspaceSettings200Response) GetEffective() GetApiWorkspaceSettings200ResponseEffective`

GetEffective returns the Effective field if non-nil, zero value otherwise.

### GetEffectiveOk

`func (o *GetApiWorkspaceSettings200Response) GetEffectiveOk() (*GetApiWorkspaceSettings200ResponseEffective, bool)`

GetEffectiveOk returns a tuple with the Effective field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEffective

`func (o *GetApiWorkspaceSettings200Response) SetEffective(v GetApiWorkspaceSettings200ResponseEffective)`

SetEffective sets Effective field to given value.


### GetDefaults

`func (o *GetApiWorkspaceSettings200Response) GetDefaults() GetApiWorkspaceSettings200ResponseEffective`

GetDefaults returns the Defaults field if non-nil, zero value otherwise.

### GetDefaultsOk

`func (o *GetApiWorkspaceSettings200Response) GetDefaultsOk() (*GetApiWorkspaceSettings200ResponseEffective, bool)`

GetDefaultsOk returns a tuple with the Defaults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaults

`func (o *GetApiWorkspaceSettings200Response) SetDefaults(v GetApiWorkspaceSettings200ResponseEffective)`

SetDefaults sets Defaults field to given value.


### GetSources

`func (o *GetApiWorkspaceSettings200Response) GetSources() GetApiWorkspaceSettings200ResponseSources`

GetSources returns the Sources field if non-nil, zero value otherwise.

### GetSourcesOk

`func (o *GetApiWorkspaceSettings200Response) GetSourcesOk() (*GetApiWorkspaceSettings200ResponseSources, bool)`

GetSourcesOk returns a tuple with the Sources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSources

`func (o *GetApiWorkspaceSettings200Response) SetSources(v GetApiWorkspaceSettings200ResponseSources)`

SetSources sets Sources field to given value.


### GetOverrides

`func (o *GetApiWorkspaceSettings200Response) GetOverrides() GetApiWorkspaceSettings200ResponseOverrides`

GetOverrides returns the Overrides field if non-nil, zero value otherwise.

### GetOverridesOk

`func (o *GetApiWorkspaceSettings200Response) GetOverridesOk() (*GetApiWorkspaceSettings200ResponseOverrides, bool)`

GetOverridesOk returns a tuple with the Overrides field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverrides

`func (o *GetApiWorkspaceSettings200Response) SetOverrides(v GetApiWorkspaceSettings200ResponseOverrides)`

SetOverrides sets Overrides field to given value.


### GetBounds

`func (o *GetApiWorkspaceSettings200Response) GetBounds() GetApiWorkspaceSettings200ResponseBounds`

GetBounds returns the Bounds field if non-nil, zero value otherwise.

### GetBoundsOk

`func (o *GetApiWorkspaceSettings200Response) GetBoundsOk() (*GetApiWorkspaceSettings200ResponseBounds, bool)`

GetBoundsOk returns a tuple with the Bounds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBounds

`func (o *GetApiWorkspaceSettings200Response) SetBounds(v GetApiWorkspaceSettings200ResponseBounds)`

SetBounds sets Bounds field to given value.


### GetRequireTwoFactor

`func (o *GetApiWorkspaceSettings200Response) GetRequireTwoFactor() bool`

GetRequireTwoFactor returns the RequireTwoFactor field if non-nil, zero value otherwise.

### GetRequireTwoFactorOk

`func (o *GetApiWorkspaceSettings200Response) GetRequireTwoFactorOk() (*bool, bool)`

GetRequireTwoFactorOk returns a tuple with the RequireTwoFactor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequireTwoFactor

`func (o *GetApiWorkspaceSettings200Response) SetRequireTwoFactor(v bool)`

SetRequireTwoFactor sets RequireTwoFactor field to given value.


### GetIsTest

`func (o *GetApiWorkspaceSettings200Response) GetIsTest() bool`

GetIsTest returns the IsTest field if non-nil, zero value otherwise.

### GetIsTestOk

`func (o *GetApiWorkspaceSettings200Response) GetIsTestOk() (*bool, bool)`

GetIsTestOk returns a tuple with the IsTest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsTest

`func (o *GetApiWorkspaceSettings200Response) SetIsTest(v bool)`

SetIsTest sets IsTest field to given value.


### GetParentWorkspaceId

`func (o *GetApiWorkspaceSettings200Response) GetParentWorkspaceId() string`

GetParentWorkspaceId returns the ParentWorkspaceId field if non-nil, zero value otherwise.

### GetParentWorkspaceIdOk

`func (o *GetApiWorkspaceSettings200Response) GetParentWorkspaceIdOk() (*string, bool)`

GetParentWorkspaceIdOk returns a tuple with the ParentWorkspaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentWorkspaceId

`func (o *GetApiWorkspaceSettings200Response) SetParentWorkspaceId(v string)`

SetParentWorkspaceId sets ParentWorkspaceId field to given value.


### SetParentWorkspaceIdNil

`func (o *GetApiWorkspaceSettings200Response) SetParentWorkspaceIdNil(b bool)`

 SetParentWorkspaceIdNil sets the value for ParentWorkspaceId to be an explicit nil

### UnsetParentWorkspaceId
`func (o *GetApiWorkspaceSettings200Response) UnsetParentWorkspaceId()`

UnsetParentWorkspaceId ensures that no value is present for ParentWorkspaceId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


