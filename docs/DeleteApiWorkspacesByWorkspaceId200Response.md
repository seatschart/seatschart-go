# DeleteApiWorkspacesByWorkspaceId200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**WorkspaceId** | **string** |  | 
**Status** | **string** |  | 
**DeactivatedAt** | **NullableTime** |  | 

## Methods

### NewDeleteApiWorkspacesByWorkspaceId200Response

`func NewDeleteApiWorkspacesByWorkspaceId200Response(workspaceId string, status string, deactivatedAt NullableTime, ) *DeleteApiWorkspacesByWorkspaceId200Response`

NewDeleteApiWorkspacesByWorkspaceId200Response instantiates a new DeleteApiWorkspacesByWorkspaceId200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeleteApiWorkspacesByWorkspaceId200ResponseWithDefaults

`func NewDeleteApiWorkspacesByWorkspaceId200ResponseWithDefaults() *DeleteApiWorkspacesByWorkspaceId200Response`

NewDeleteApiWorkspacesByWorkspaceId200ResponseWithDefaults instantiates a new DeleteApiWorkspacesByWorkspaceId200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWorkspaceId

`func (o *DeleteApiWorkspacesByWorkspaceId200Response) GetWorkspaceId() string`

GetWorkspaceId returns the WorkspaceId field if non-nil, zero value otherwise.

### GetWorkspaceIdOk

`func (o *DeleteApiWorkspacesByWorkspaceId200Response) GetWorkspaceIdOk() (*string, bool)`

GetWorkspaceIdOk returns a tuple with the WorkspaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkspaceId

`func (o *DeleteApiWorkspacesByWorkspaceId200Response) SetWorkspaceId(v string)`

SetWorkspaceId sets WorkspaceId field to given value.


### GetStatus

`func (o *DeleteApiWorkspacesByWorkspaceId200Response) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *DeleteApiWorkspacesByWorkspaceId200Response) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *DeleteApiWorkspacesByWorkspaceId200Response) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetDeactivatedAt

`func (o *DeleteApiWorkspacesByWorkspaceId200Response) GetDeactivatedAt() time.Time`

GetDeactivatedAt returns the DeactivatedAt field if non-nil, zero value otherwise.

### GetDeactivatedAtOk

`func (o *DeleteApiWorkspacesByWorkspaceId200Response) GetDeactivatedAtOk() (*time.Time, bool)`

GetDeactivatedAtOk returns a tuple with the DeactivatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeactivatedAt

`func (o *DeleteApiWorkspacesByWorkspaceId200Response) SetDeactivatedAt(v time.Time)`

SetDeactivatedAt sets DeactivatedAt field to given value.


### SetDeactivatedAtNil

`func (o *DeleteApiWorkspacesByWorkspaceId200Response) SetDeactivatedAtNil(b bool)`

 SetDeactivatedAtNil sets the value for DeactivatedAt to be an explicit nil

### UnsetDeactivatedAt
`func (o *DeleteApiWorkspacesByWorkspaceId200Response) UnsetDeactivatedAt()`

UnsetDeactivatedAt ensures that no value is present for DeactivatedAt, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


