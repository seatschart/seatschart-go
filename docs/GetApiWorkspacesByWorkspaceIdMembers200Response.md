# GetApiWorkspacesByWorkspaceIdMembers200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**WorkspaceId** | **string** |  | 
**Members** | [**[]GetApiWorkspacesByWorkspaceIdMembers200ResponseMembersInner**](GetApiWorkspacesByWorkspaceIdMembers200ResponseMembersInner.md) |  | 

## Methods

### NewGetApiWorkspacesByWorkspaceIdMembers200Response

`func NewGetApiWorkspacesByWorkspaceIdMembers200Response(workspaceId string, members []GetApiWorkspacesByWorkspaceIdMembers200ResponseMembersInner, ) *GetApiWorkspacesByWorkspaceIdMembers200Response`

NewGetApiWorkspacesByWorkspaceIdMembers200Response instantiates a new GetApiWorkspacesByWorkspaceIdMembers200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetApiWorkspacesByWorkspaceIdMembers200ResponseWithDefaults

`func NewGetApiWorkspacesByWorkspaceIdMembers200ResponseWithDefaults() *GetApiWorkspacesByWorkspaceIdMembers200Response`

NewGetApiWorkspacesByWorkspaceIdMembers200ResponseWithDefaults instantiates a new GetApiWorkspacesByWorkspaceIdMembers200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWorkspaceId

`func (o *GetApiWorkspacesByWorkspaceIdMembers200Response) GetWorkspaceId() string`

GetWorkspaceId returns the WorkspaceId field if non-nil, zero value otherwise.

### GetWorkspaceIdOk

`func (o *GetApiWorkspacesByWorkspaceIdMembers200Response) GetWorkspaceIdOk() (*string, bool)`

GetWorkspaceIdOk returns a tuple with the WorkspaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkspaceId

`func (o *GetApiWorkspacesByWorkspaceIdMembers200Response) SetWorkspaceId(v string)`

SetWorkspaceId sets WorkspaceId field to given value.


### GetMembers

`func (o *GetApiWorkspacesByWorkspaceIdMembers200Response) GetMembers() []GetApiWorkspacesByWorkspaceIdMembers200ResponseMembersInner`

GetMembers returns the Members field if non-nil, zero value otherwise.

### GetMembersOk

`func (o *GetApiWorkspacesByWorkspaceIdMembers200Response) GetMembersOk() (*[]GetApiWorkspacesByWorkspaceIdMembers200ResponseMembersInner, bool)`

GetMembersOk returns a tuple with the Members field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMembers

`func (o *GetApiWorkspacesByWorkspaceIdMembers200Response) SetMembers(v []GetApiWorkspacesByWorkspaceIdMembers200ResponseMembersInner)`

SetMembers sets Members field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


