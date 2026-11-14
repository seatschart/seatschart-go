# PostApiWorkspacesByWorkspaceIdInvitations201Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Invitation** | [**PostApiWorkspacesByWorkspaceIdInvitations201ResponseInvitation**](PostApiWorkspacesByWorkspaceIdInvitations201ResponseInvitation.md) |  | 
**InviteToken** | **string** | One-time token of the accept link (&#x60;/accept-invite?token&#x3D;…&#x60;), returned once. | 

## Methods

### NewPostApiWorkspacesByWorkspaceIdInvitations201Response

`func NewPostApiWorkspacesByWorkspaceIdInvitations201Response(invitation PostApiWorkspacesByWorkspaceIdInvitations201ResponseInvitation, inviteToken string, ) *PostApiWorkspacesByWorkspaceIdInvitations201Response`

NewPostApiWorkspacesByWorkspaceIdInvitations201Response instantiates a new PostApiWorkspacesByWorkspaceIdInvitations201Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostApiWorkspacesByWorkspaceIdInvitations201ResponseWithDefaults

`func NewPostApiWorkspacesByWorkspaceIdInvitations201ResponseWithDefaults() *PostApiWorkspacesByWorkspaceIdInvitations201Response`

NewPostApiWorkspacesByWorkspaceIdInvitations201ResponseWithDefaults instantiates a new PostApiWorkspacesByWorkspaceIdInvitations201Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInvitation

`func (o *PostApiWorkspacesByWorkspaceIdInvitations201Response) GetInvitation() PostApiWorkspacesByWorkspaceIdInvitations201ResponseInvitation`

GetInvitation returns the Invitation field if non-nil, zero value otherwise.

### GetInvitationOk

`func (o *PostApiWorkspacesByWorkspaceIdInvitations201Response) GetInvitationOk() (*PostApiWorkspacesByWorkspaceIdInvitations201ResponseInvitation, bool)`

GetInvitationOk returns a tuple with the Invitation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvitation

`func (o *PostApiWorkspacesByWorkspaceIdInvitations201Response) SetInvitation(v PostApiWorkspacesByWorkspaceIdInvitations201ResponseInvitation)`

SetInvitation sets Invitation field to given value.


### GetInviteToken

`func (o *PostApiWorkspacesByWorkspaceIdInvitations201Response) GetInviteToken() string`

GetInviteToken returns the InviteToken field if non-nil, zero value otherwise.

### GetInviteTokenOk

`func (o *PostApiWorkspacesByWorkspaceIdInvitations201Response) GetInviteTokenOk() (*string, bool)`

GetInviteTokenOk returns a tuple with the InviteToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInviteToken

`func (o *PostApiWorkspacesByWorkspaceIdInvitations201Response) SetInviteToken(v string)`

SetInviteToken sets InviteToken field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


