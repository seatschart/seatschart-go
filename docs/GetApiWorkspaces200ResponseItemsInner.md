# GetApiWorkspaces200ResponseItemsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**WorkspaceId** | **string** |  | 
**DisplayName** | **string** |  | 
**Role** | **string** |  | 
**PlanId** | **string** |  | 
**SubscriptionStatus** | **string** |  | 
**Status** | **string** |  | 
**DeactivatedAt** | **NullableTime** |  | 
**IsCurrent** | **bool** | True for the workspace of the calling session (or of the calling workspace admin key). | 
**IsDefault** | **bool** | The default workspace of the user (for a workspace admin key: of the admin it is bound to): dashboard sign-in opens it. See POST /workspaces/{workspaceId}/actions/set-default. | 
**IsTest** | **bool** | Test (sandbox) workspace: not billed or metered, &#x60;sk_test_&#x60; keys, TEST badge for buyers. | 
**ParentWorkspaceId** | **NullableString** | Production workspace a test workspace was created from. | 

## Methods

### NewGetApiWorkspaces200ResponseItemsInner

`func NewGetApiWorkspaces200ResponseItemsInner(workspaceId string, displayName string, role string, planId string, subscriptionStatus string, status string, deactivatedAt NullableTime, isCurrent bool, isDefault bool, isTest bool, parentWorkspaceId NullableString, ) *GetApiWorkspaces200ResponseItemsInner`

NewGetApiWorkspaces200ResponseItemsInner instantiates a new GetApiWorkspaces200ResponseItemsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetApiWorkspaces200ResponseItemsInnerWithDefaults

`func NewGetApiWorkspaces200ResponseItemsInnerWithDefaults() *GetApiWorkspaces200ResponseItemsInner`

NewGetApiWorkspaces200ResponseItemsInnerWithDefaults instantiates a new GetApiWorkspaces200ResponseItemsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWorkspaceId

`func (o *GetApiWorkspaces200ResponseItemsInner) GetWorkspaceId() string`

GetWorkspaceId returns the WorkspaceId field if non-nil, zero value otherwise.

### GetWorkspaceIdOk

`func (o *GetApiWorkspaces200ResponseItemsInner) GetWorkspaceIdOk() (*string, bool)`

GetWorkspaceIdOk returns a tuple with the WorkspaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkspaceId

`func (o *GetApiWorkspaces200ResponseItemsInner) SetWorkspaceId(v string)`

SetWorkspaceId sets WorkspaceId field to given value.


### GetDisplayName

`func (o *GetApiWorkspaces200ResponseItemsInner) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *GetApiWorkspaces200ResponseItemsInner) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *GetApiWorkspaces200ResponseItemsInner) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.


### GetRole

`func (o *GetApiWorkspaces200ResponseItemsInner) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *GetApiWorkspaces200ResponseItemsInner) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *GetApiWorkspaces200ResponseItemsInner) SetRole(v string)`

SetRole sets Role field to given value.


### GetPlanId

`func (o *GetApiWorkspaces200ResponseItemsInner) GetPlanId() string`

GetPlanId returns the PlanId field if non-nil, zero value otherwise.

### GetPlanIdOk

`func (o *GetApiWorkspaces200ResponseItemsInner) GetPlanIdOk() (*string, bool)`

GetPlanIdOk returns a tuple with the PlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanId

`func (o *GetApiWorkspaces200ResponseItemsInner) SetPlanId(v string)`

SetPlanId sets PlanId field to given value.


### GetSubscriptionStatus

`func (o *GetApiWorkspaces200ResponseItemsInner) GetSubscriptionStatus() string`

GetSubscriptionStatus returns the SubscriptionStatus field if non-nil, zero value otherwise.

### GetSubscriptionStatusOk

`func (o *GetApiWorkspaces200ResponseItemsInner) GetSubscriptionStatusOk() (*string, bool)`

GetSubscriptionStatusOk returns a tuple with the SubscriptionStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionStatus

`func (o *GetApiWorkspaces200ResponseItemsInner) SetSubscriptionStatus(v string)`

SetSubscriptionStatus sets SubscriptionStatus field to given value.


### GetStatus

`func (o *GetApiWorkspaces200ResponseItemsInner) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetApiWorkspaces200ResponseItemsInner) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetApiWorkspaces200ResponseItemsInner) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetDeactivatedAt

`func (o *GetApiWorkspaces200ResponseItemsInner) GetDeactivatedAt() time.Time`

GetDeactivatedAt returns the DeactivatedAt field if non-nil, zero value otherwise.

### GetDeactivatedAtOk

`func (o *GetApiWorkspaces200ResponseItemsInner) GetDeactivatedAtOk() (*time.Time, bool)`

GetDeactivatedAtOk returns a tuple with the DeactivatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeactivatedAt

`func (o *GetApiWorkspaces200ResponseItemsInner) SetDeactivatedAt(v time.Time)`

SetDeactivatedAt sets DeactivatedAt field to given value.


### SetDeactivatedAtNil

`func (o *GetApiWorkspaces200ResponseItemsInner) SetDeactivatedAtNil(b bool)`

 SetDeactivatedAtNil sets the value for DeactivatedAt to be an explicit nil

### UnsetDeactivatedAt
`func (o *GetApiWorkspaces200ResponseItemsInner) UnsetDeactivatedAt()`

UnsetDeactivatedAt ensures that no value is present for DeactivatedAt, not even an explicit nil
### GetIsCurrent

`func (o *GetApiWorkspaces200ResponseItemsInner) GetIsCurrent() bool`

GetIsCurrent returns the IsCurrent field if non-nil, zero value otherwise.

### GetIsCurrentOk

`func (o *GetApiWorkspaces200ResponseItemsInner) GetIsCurrentOk() (*bool, bool)`

GetIsCurrentOk returns a tuple with the IsCurrent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsCurrent

`func (o *GetApiWorkspaces200ResponseItemsInner) SetIsCurrent(v bool)`

SetIsCurrent sets IsCurrent field to given value.


### GetIsDefault

`func (o *GetApiWorkspaces200ResponseItemsInner) GetIsDefault() bool`

GetIsDefault returns the IsDefault field if non-nil, zero value otherwise.

### GetIsDefaultOk

`func (o *GetApiWorkspaces200ResponseItemsInner) GetIsDefaultOk() (*bool, bool)`

GetIsDefaultOk returns a tuple with the IsDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDefault

`func (o *GetApiWorkspaces200ResponseItemsInner) SetIsDefault(v bool)`

SetIsDefault sets IsDefault field to given value.


### GetIsTest

`func (o *GetApiWorkspaces200ResponseItemsInner) GetIsTest() bool`

GetIsTest returns the IsTest field if non-nil, zero value otherwise.

### GetIsTestOk

`func (o *GetApiWorkspaces200ResponseItemsInner) GetIsTestOk() (*bool, bool)`

GetIsTestOk returns a tuple with the IsTest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsTest

`func (o *GetApiWorkspaces200ResponseItemsInner) SetIsTest(v bool)`

SetIsTest sets IsTest field to given value.


### GetParentWorkspaceId

`func (o *GetApiWorkspaces200ResponseItemsInner) GetParentWorkspaceId() string`

GetParentWorkspaceId returns the ParentWorkspaceId field if non-nil, zero value otherwise.

### GetParentWorkspaceIdOk

`func (o *GetApiWorkspaces200ResponseItemsInner) GetParentWorkspaceIdOk() (*string, bool)`

GetParentWorkspaceIdOk returns a tuple with the ParentWorkspaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentWorkspaceId

`func (o *GetApiWorkspaces200ResponseItemsInner) SetParentWorkspaceId(v string)`

SetParentWorkspaceId sets ParentWorkspaceId field to given value.


### SetParentWorkspaceIdNil

`func (o *GetApiWorkspaces200ResponseItemsInner) SetParentWorkspaceIdNil(b bool)`

 SetParentWorkspaceIdNil sets the value for ParentWorkspaceId to be an explicit nil

### UnsetParentWorkspaceId
`func (o *GetApiWorkspaces200ResponseItemsInner) UnsetParentWorkspaceId()`

UnsetParentWorkspaceId ensures that no value is present for ParentWorkspaceId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


