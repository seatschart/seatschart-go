# PostApiChartsByChartIdActionsCopyToWorkspaceRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**WorkspaceId** | **string** | Target workspace id (see GET /workspaces). | 
**FromWorkspaceId** | Pointer to **string** | Source workspace id (default: the calling workspace). seats.io copyFromWorkspaceTo: copy a chart of another workspace the same admin administers; pass the chart id (UUID), keys resolve in the calling workspace only. | [optional] 
**Name** | Pointer to **string** |  | [optional] 

## Methods

### NewPostApiChartsByChartIdActionsCopyToWorkspaceRequest

`func NewPostApiChartsByChartIdActionsCopyToWorkspaceRequest(workspaceId string, ) *PostApiChartsByChartIdActionsCopyToWorkspaceRequest`

NewPostApiChartsByChartIdActionsCopyToWorkspaceRequest instantiates a new PostApiChartsByChartIdActionsCopyToWorkspaceRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostApiChartsByChartIdActionsCopyToWorkspaceRequestWithDefaults

`func NewPostApiChartsByChartIdActionsCopyToWorkspaceRequestWithDefaults() *PostApiChartsByChartIdActionsCopyToWorkspaceRequest`

NewPostApiChartsByChartIdActionsCopyToWorkspaceRequestWithDefaults instantiates a new PostApiChartsByChartIdActionsCopyToWorkspaceRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWorkspaceId

`func (o *PostApiChartsByChartIdActionsCopyToWorkspaceRequest) GetWorkspaceId() string`

GetWorkspaceId returns the WorkspaceId field if non-nil, zero value otherwise.

### GetWorkspaceIdOk

`func (o *PostApiChartsByChartIdActionsCopyToWorkspaceRequest) GetWorkspaceIdOk() (*string, bool)`

GetWorkspaceIdOk returns a tuple with the WorkspaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkspaceId

`func (o *PostApiChartsByChartIdActionsCopyToWorkspaceRequest) SetWorkspaceId(v string)`

SetWorkspaceId sets WorkspaceId field to given value.


### GetFromWorkspaceId

`func (o *PostApiChartsByChartIdActionsCopyToWorkspaceRequest) GetFromWorkspaceId() string`

GetFromWorkspaceId returns the FromWorkspaceId field if non-nil, zero value otherwise.

### GetFromWorkspaceIdOk

`func (o *PostApiChartsByChartIdActionsCopyToWorkspaceRequest) GetFromWorkspaceIdOk() (*string, bool)`

GetFromWorkspaceIdOk returns a tuple with the FromWorkspaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromWorkspaceId

`func (o *PostApiChartsByChartIdActionsCopyToWorkspaceRequest) SetFromWorkspaceId(v string)`

SetFromWorkspaceId sets FromWorkspaceId field to given value.

### HasFromWorkspaceId

`func (o *PostApiChartsByChartIdActionsCopyToWorkspaceRequest) HasFromWorkspaceId() bool`

HasFromWorkspaceId returns a boolean if a field has been set.

### GetName

`func (o *PostApiChartsByChartIdActionsCopyToWorkspaceRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PostApiChartsByChartIdActionsCopyToWorkspaceRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PostApiChartsByChartIdActionsCopyToWorkspaceRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PostApiChartsByChartIdActionsCopyToWorkspaceRequest) HasName() bool`

HasName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


