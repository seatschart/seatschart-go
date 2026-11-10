# GetApiSandboxes200ResponseItemsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SandboxId** | **string** | Workspace id of the sandbox (&#x60;ws_…&#x60;): switch to it, or call it with its own &#x60;sk_test_&#x60; key. | 
**DisplayName** | **string** |  | 
**Kind** | **string** |  | 
**ParentWorkspaceId** | **string** | The production workspace the sandbox belongs to. | 
**PublicKey** | **NullableString** | Public renderer key (&#x60;pub_…&#x60;) of the sandbox; not a secret. | 
**Status** | **string** |  | 
**CreatedAt** | **time.Time** |  | 
**LastResetAt** | **NullableTime** |  | 
**Usage** | [**GetApiSandboxes200ResponseItemsInnerUsage**](GetApiSandboxes200ResponseItemsInnerUsage.md) |  | 
**RateLimits** | [**GetApiSandboxes200ResponseItemsInnerRateLimits**](GetApiSandboxes200ResponseItemsInnerRateLimits.md) |  | 
**Role** | **NullableString** | Role of the calling dashboard user in the sandbox (null for API keys or non-members). | 

## Methods

### NewGetApiSandboxes200ResponseItemsInner

`func NewGetApiSandboxes200ResponseItemsInner(sandboxId string, displayName string, kind string, parentWorkspaceId string, publicKey NullableString, status string, createdAt time.Time, lastResetAt NullableTime, usage GetApiSandboxes200ResponseItemsInnerUsage, rateLimits GetApiSandboxes200ResponseItemsInnerRateLimits, role NullableString, ) *GetApiSandboxes200ResponseItemsInner`

NewGetApiSandboxes200ResponseItemsInner instantiates a new GetApiSandboxes200ResponseItemsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetApiSandboxes200ResponseItemsInnerWithDefaults

`func NewGetApiSandboxes200ResponseItemsInnerWithDefaults() *GetApiSandboxes200ResponseItemsInner`

NewGetApiSandboxes200ResponseItemsInnerWithDefaults instantiates a new GetApiSandboxes200ResponseItemsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSandboxId

`func (o *GetApiSandboxes200ResponseItemsInner) GetSandboxId() string`

GetSandboxId returns the SandboxId field if non-nil, zero value otherwise.

### GetSandboxIdOk

`func (o *GetApiSandboxes200ResponseItemsInner) GetSandboxIdOk() (*string, bool)`

GetSandboxIdOk returns a tuple with the SandboxId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSandboxId

`func (o *GetApiSandboxes200ResponseItemsInner) SetSandboxId(v string)`

SetSandboxId sets SandboxId field to given value.


### GetDisplayName

`func (o *GetApiSandboxes200ResponseItemsInner) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *GetApiSandboxes200ResponseItemsInner) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *GetApiSandboxes200ResponseItemsInner) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.


### GetKind

`func (o *GetApiSandboxes200ResponseItemsInner) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *GetApiSandboxes200ResponseItemsInner) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *GetApiSandboxes200ResponseItemsInner) SetKind(v string)`

SetKind sets Kind field to given value.


### GetParentWorkspaceId

`func (o *GetApiSandboxes200ResponseItemsInner) GetParentWorkspaceId() string`

GetParentWorkspaceId returns the ParentWorkspaceId field if non-nil, zero value otherwise.

### GetParentWorkspaceIdOk

`func (o *GetApiSandboxes200ResponseItemsInner) GetParentWorkspaceIdOk() (*string, bool)`

GetParentWorkspaceIdOk returns a tuple with the ParentWorkspaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentWorkspaceId

`func (o *GetApiSandboxes200ResponseItemsInner) SetParentWorkspaceId(v string)`

SetParentWorkspaceId sets ParentWorkspaceId field to given value.


### GetPublicKey

`func (o *GetApiSandboxes200ResponseItemsInner) GetPublicKey() string`

GetPublicKey returns the PublicKey field if non-nil, zero value otherwise.

### GetPublicKeyOk

`func (o *GetApiSandboxes200ResponseItemsInner) GetPublicKeyOk() (*string, bool)`

GetPublicKeyOk returns a tuple with the PublicKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicKey

`func (o *GetApiSandboxes200ResponseItemsInner) SetPublicKey(v string)`

SetPublicKey sets PublicKey field to given value.


### SetPublicKeyNil

`func (o *GetApiSandboxes200ResponseItemsInner) SetPublicKeyNil(b bool)`

 SetPublicKeyNil sets the value for PublicKey to be an explicit nil

### UnsetPublicKey
`func (o *GetApiSandboxes200ResponseItemsInner) UnsetPublicKey()`

UnsetPublicKey ensures that no value is present for PublicKey, not even an explicit nil
### GetStatus

`func (o *GetApiSandboxes200ResponseItemsInner) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetApiSandboxes200ResponseItemsInner) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetApiSandboxes200ResponseItemsInner) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetCreatedAt

`func (o *GetApiSandboxes200ResponseItemsInner) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *GetApiSandboxes200ResponseItemsInner) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *GetApiSandboxes200ResponseItemsInner) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetLastResetAt

`func (o *GetApiSandboxes200ResponseItemsInner) GetLastResetAt() time.Time`

GetLastResetAt returns the LastResetAt field if non-nil, zero value otherwise.

### GetLastResetAtOk

`func (o *GetApiSandboxes200ResponseItemsInner) GetLastResetAtOk() (*time.Time, bool)`

GetLastResetAtOk returns a tuple with the LastResetAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastResetAt

`func (o *GetApiSandboxes200ResponseItemsInner) SetLastResetAt(v time.Time)`

SetLastResetAt sets LastResetAt field to given value.


### SetLastResetAtNil

`func (o *GetApiSandboxes200ResponseItemsInner) SetLastResetAtNil(b bool)`

 SetLastResetAtNil sets the value for LastResetAt to be an explicit nil

### UnsetLastResetAt
`func (o *GetApiSandboxes200ResponseItemsInner) UnsetLastResetAt()`

UnsetLastResetAt ensures that no value is present for LastResetAt, not even an explicit nil
### GetUsage

`func (o *GetApiSandboxes200ResponseItemsInner) GetUsage() GetApiSandboxes200ResponseItemsInnerUsage`

GetUsage returns the Usage field if non-nil, zero value otherwise.

### GetUsageOk

`func (o *GetApiSandboxes200ResponseItemsInner) GetUsageOk() (*GetApiSandboxes200ResponseItemsInnerUsage, bool)`

GetUsageOk returns a tuple with the Usage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsage

`func (o *GetApiSandboxes200ResponseItemsInner) SetUsage(v GetApiSandboxes200ResponseItemsInnerUsage)`

SetUsage sets Usage field to given value.


### GetRateLimits

`func (o *GetApiSandboxes200ResponseItemsInner) GetRateLimits() GetApiSandboxes200ResponseItemsInnerRateLimits`

GetRateLimits returns the RateLimits field if non-nil, zero value otherwise.

### GetRateLimitsOk

`func (o *GetApiSandboxes200ResponseItemsInner) GetRateLimitsOk() (*GetApiSandboxes200ResponseItemsInnerRateLimits, bool)`

GetRateLimitsOk returns a tuple with the RateLimits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRateLimits

`func (o *GetApiSandboxes200ResponseItemsInner) SetRateLimits(v GetApiSandboxes200ResponseItemsInnerRateLimits)`

SetRateLimits sets RateLimits field to given value.


### GetRole

`func (o *GetApiSandboxes200ResponseItemsInner) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *GetApiSandboxes200ResponseItemsInner) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *GetApiSandboxes200ResponseItemsInner) SetRole(v string)`

SetRole sets Role field to given value.


### SetRoleNil

`func (o *GetApiSandboxes200ResponseItemsInner) SetRoleNil(b bool)`

 SetRoleNil sets the value for Role to be an explicit nil

### UnsetRole
`func (o *GetApiSandboxes200ResponseItemsInner) UnsetRole()`

UnsetRole ensures that no value is present for Role, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


