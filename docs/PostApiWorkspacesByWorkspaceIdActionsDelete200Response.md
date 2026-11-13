# PostApiWorkspacesByWorkspaceIdActionsDelete200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**WorkspaceId** | **string** |  | 
**Status** | **string** |  | 
**DeletedAt** | **time.Time** |  | 
**SwitchTo** | **NullableString** | Set when the deleted workspace was the one of the calling session: the workspace to switch to (POST /auth/switch-workspace). It is also your default workspace when the deleted one was. | 
**DeletedSandboxIds** | **[]string** | Sandboxes of the workspace, soft-deleted with it. | 

## Methods

### NewPostApiWorkspacesByWorkspaceIdActionsDelete200Response

`func NewPostApiWorkspacesByWorkspaceIdActionsDelete200Response(workspaceId string, status string, deletedAt time.Time, switchTo NullableString, deletedSandboxIds []string, ) *PostApiWorkspacesByWorkspaceIdActionsDelete200Response`

NewPostApiWorkspacesByWorkspaceIdActionsDelete200Response instantiates a new PostApiWorkspacesByWorkspaceIdActionsDelete200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostApiWorkspacesByWorkspaceIdActionsDelete200ResponseWithDefaults

`func NewPostApiWorkspacesByWorkspaceIdActionsDelete200ResponseWithDefaults() *PostApiWorkspacesByWorkspaceIdActionsDelete200Response`

NewPostApiWorkspacesByWorkspaceIdActionsDelete200ResponseWithDefaults instantiates a new PostApiWorkspacesByWorkspaceIdActionsDelete200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWorkspaceId

`func (o *PostApiWorkspacesByWorkspaceIdActionsDelete200Response) GetWorkspaceId() string`

GetWorkspaceId returns the WorkspaceId field if non-nil, zero value otherwise.

### GetWorkspaceIdOk

`func (o *PostApiWorkspacesByWorkspaceIdActionsDelete200Response) GetWorkspaceIdOk() (*string, bool)`

GetWorkspaceIdOk returns a tuple with the WorkspaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkspaceId

`func (o *PostApiWorkspacesByWorkspaceIdActionsDelete200Response) SetWorkspaceId(v string)`

SetWorkspaceId sets WorkspaceId field to given value.


### GetStatus

`func (o *PostApiWorkspacesByWorkspaceIdActionsDelete200Response) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *PostApiWorkspacesByWorkspaceIdActionsDelete200Response) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *PostApiWorkspacesByWorkspaceIdActionsDelete200Response) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetDeletedAt

`func (o *PostApiWorkspacesByWorkspaceIdActionsDelete200Response) GetDeletedAt() time.Time`

GetDeletedAt returns the DeletedAt field if non-nil, zero value otherwise.

### GetDeletedAtOk

`func (o *PostApiWorkspacesByWorkspaceIdActionsDelete200Response) GetDeletedAtOk() (*time.Time, bool)`

GetDeletedAtOk returns a tuple with the DeletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedAt

`func (o *PostApiWorkspacesByWorkspaceIdActionsDelete200Response) SetDeletedAt(v time.Time)`

SetDeletedAt sets DeletedAt field to given value.


### GetSwitchTo

`func (o *PostApiWorkspacesByWorkspaceIdActionsDelete200Response) GetSwitchTo() string`

GetSwitchTo returns the SwitchTo field if non-nil, zero value otherwise.

### GetSwitchToOk

`func (o *PostApiWorkspacesByWorkspaceIdActionsDelete200Response) GetSwitchToOk() (*string, bool)`

GetSwitchToOk returns a tuple with the SwitchTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSwitchTo

`func (o *PostApiWorkspacesByWorkspaceIdActionsDelete200Response) SetSwitchTo(v string)`

SetSwitchTo sets SwitchTo field to given value.


### SetSwitchToNil

`func (o *PostApiWorkspacesByWorkspaceIdActionsDelete200Response) SetSwitchToNil(b bool)`

 SetSwitchToNil sets the value for SwitchTo to be an explicit nil

### UnsetSwitchTo
`func (o *PostApiWorkspacesByWorkspaceIdActionsDelete200Response) UnsetSwitchTo()`

UnsetSwitchTo ensures that no value is present for SwitchTo, not even an explicit nil
### GetDeletedSandboxIds

`func (o *PostApiWorkspacesByWorkspaceIdActionsDelete200Response) GetDeletedSandboxIds() []string`

GetDeletedSandboxIds returns the DeletedSandboxIds field if non-nil, zero value otherwise.

### GetDeletedSandboxIdsOk

`func (o *PostApiWorkspacesByWorkspaceIdActionsDelete200Response) GetDeletedSandboxIdsOk() (*[]string, bool)`

GetDeletedSandboxIdsOk returns a tuple with the DeletedSandboxIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedSandboxIds

`func (o *PostApiWorkspacesByWorkspaceIdActionsDelete200Response) SetDeletedSandboxIds(v []string)`

SetDeletedSandboxIds sets DeletedSandboxIds field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


