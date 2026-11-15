# GetApiWorkspacesByWorkspaceIdMembers200ResponseMembersInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | User id. | 
**Email** | **string** |  | 
**Role** | **string** | Workspace role: &#x60;admin&#x60;, &#x60;map_editor&#x60; or &#x60;viewer&#x60;. | 
**CreatedAt** | **time.Time** | Joined at. | 
**TwoFactorEnabled** | **bool** |  | 

## Methods

### NewGetApiWorkspacesByWorkspaceIdMembers200ResponseMembersInner

`func NewGetApiWorkspacesByWorkspaceIdMembers200ResponseMembersInner(id string, email string, role string, createdAt time.Time, twoFactorEnabled bool, ) *GetApiWorkspacesByWorkspaceIdMembers200ResponseMembersInner`

NewGetApiWorkspacesByWorkspaceIdMembers200ResponseMembersInner instantiates a new GetApiWorkspacesByWorkspaceIdMembers200ResponseMembersInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetApiWorkspacesByWorkspaceIdMembers200ResponseMembersInnerWithDefaults

`func NewGetApiWorkspacesByWorkspaceIdMembers200ResponseMembersInnerWithDefaults() *GetApiWorkspacesByWorkspaceIdMembers200ResponseMembersInner`

NewGetApiWorkspacesByWorkspaceIdMembers200ResponseMembersInnerWithDefaults instantiates a new GetApiWorkspacesByWorkspaceIdMembers200ResponseMembersInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GetApiWorkspacesByWorkspaceIdMembers200ResponseMembersInner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetApiWorkspacesByWorkspaceIdMembers200ResponseMembersInner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetApiWorkspacesByWorkspaceIdMembers200ResponseMembersInner) SetId(v string)`

SetId sets Id field to given value.


### GetEmail

`func (o *GetApiWorkspacesByWorkspaceIdMembers200ResponseMembersInner) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *GetApiWorkspacesByWorkspaceIdMembers200ResponseMembersInner) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *GetApiWorkspacesByWorkspaceIdMembers200ResponseMembersInner) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetRole

`func (o *GetApiWorkspacesByWorkspaceIdMembers200ResponseMembersInner) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *GetApiWorkspacesByWorkspaceIdMembers200ResponseMembersInner) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *GetApiWorkspacesByWorkspaceIdMembers200ResponseMembersInner) SetRole(v string)`

SetRole sets Role field to given value.


### GetCreatedAt

`func (o *GetApiWorkspacesByWorkspaceIdMembers200ResponseMembersInner) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *GetApiWorkspacesByWorkspaceIdMembers200ResponseMembersInner) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *GetApiWorkspacesByWorkspaceIdMembers200ResponseMembersInner) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetTwoFactorEnabled

`func (o *GetApiWorkspacesByWorkspaceIdMembers200ResponseMembersInner) GetTwoFactorEnabled() bool`

GetTwoFactorEnabled returns the TwoFactorEnabled field if non-nil, zero value otherwise.

### GetTwoFactorEnabledOk

`func (o *GetApiWorkspacesByWorkspaceIdMembers200ResponseMembersInner) GetTwoFactorEnabledOk() (*bool, bool)`

GetTwoFactorEnabledOk returns a tuple with the TwoFactorEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTwoFactorEnabled

`func (o *GetApiWorkspacesByWorkspaceIdMembers200ResponseMembersInner) SetTwoFactorEnabled(v bool)`

SetTwoFactorEnabled sets TwoFactorEnabled field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


