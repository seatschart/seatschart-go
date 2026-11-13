# PostApiWorkspacesByWorkspaceIdInvitationsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Email** | **string** |  | 
**Role** | **string** | &#x60;admin&#x60; (everything in this workspace), &#x60;map_editor&#x60; (charts, events, designer) or &#x60;viewer&#x60; (read and reports). | 

## Methods

### NewPostApiWorkspacesByWorkspaceIdInvitationsRequest

`func NewPostApiWorkspacesByWorkspaceIdInvitationsRequest(email string, role string, ) *PostApiWorkspacesByWorkspaceIdInvitationsRequest`

NewPostApiWorkspacesByWorkspaceIdInvitationsRequest instantiates a new PostApiWorkspacesByWorkspaceIdInvitationsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostApiWorkspacesByWorkspaceIdInvitationsRequestWithDefaults

`func NewPostApiWorkspacesByWorkspaceIdInvitationsRequestWithDefaults() *PostApiWorkspacesByWorkspaceIdInvitationsRequest`

NewPostApiWorkspacesByWorkspaceIdInvitationsRequestWithDefaults instantiates a new PostApiWorkspacesByWorkspaceIdInvitationsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmail

`func (o *PostApiWorkspacesByWorkspaceIdInvitationsRequest) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *PostApiWorkspacesByWorkspaceIdInvitationsRequest) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *PostApiWorkspacesByWorkspaceIdInvitationsRequest) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetRole

`func (o *PostApiWorkspacesByWorkspaceIdInvitationsRequest) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *PostApiWorkspacesByWorkspaceIdInvitationsRequest) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *PostApiWorkspacesByWorkspaceIdInvitationsRequest) SetRole(v string)`

SetRole sets Role field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


