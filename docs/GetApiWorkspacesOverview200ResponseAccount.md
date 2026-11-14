# GetApiWorkspacesOverview200ResponseAccount

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**WorkspaceCount** | **int32** |  | 
**Plan** | [**GetApiWorkspacesOverview200ResponseAccountPlan**](GetApiWorkspacesOverview200ResponseAccountPlan.md) |  | 
**Charts** | [**GetApiWorkspacesOverview200ResponseAccountCharts**](GetApiWorkspacesOverview200ResponseAccountCharts.md) |  | 
**Events** | [**GetApiWorkspacesOverview200ResponseAccountCharts**](GetApiWorkspacesOverview200ResponseAccountCharts.md) |  | 

## Methods

### NewGetApiWorkspacesOverview200ResponseAccount

`func NewGetApiWorkspacesOverview200ResponseAccount(workspaceCount int32, plan GetApiWorkspacesOverview200ResponseAccountPlan, charts GetApiWorkspacesOverview200ResponseAccountCharts, events GetApiWorkspacesOverview200ResponseAccountCharts, ) *GetApiWorkspacesOverview200ResponseAccount`

NewGetApiWorkspacesOverview200ResponseAccount instantiates a new GetApiWorkspacesOverview200ResponseAccount object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetApiWorkspacesOverview200ResponseAccountWithDefaults

`func NewGetApiWorkspacesOverview200ResponseAccountWithDefaults() *GetApiWorkspacesOverview200ResponseAccount`

NewGetApiWorkspacesOverview200ResponseAccountWithDefaults instantiates a new GetApiWorkspacesOverview200ResponseAccount object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWorkspaceCount

`func (o *GetApiWorkspacesOverview200ResponseAccount) GetWorkspaceCount() int32`

GetWorkspaceCount returns the WorkspaceCount field if non-nil, zero value otherwise.

### GetWorkspaceCountOk

`func (o *GetApiWorkspacesOverview200ResponseAccount) GetWorkspaceCountOk() (*int32, bool)`

GetWorkspaceCountOk returns a tuple with the WorkspaceCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkspaceCount

`func (o *GetApiWorkspacesOverview200ResponseAccount) SetWorkspaceCount(v int32)`

SetWorkspaceCount sets WorkspaceCount field to given value.


### GetPlan

`func (o *GetApiWorkspacesOverview200ResponseAccount) GetPlan() GetApiWorkspacesOverview200ResponseAccountPlan`

GetPlan returns the Plan field if non-nil, zero value otherwise.

### GetPlanOk

`func (o *GetApiWorkspacesOverview200ResponseAccount) GetPlanOk() (*GetApiWorkspacesOverview200ResponseAccountPlan, bool)`

GetPlanOk returns a tuple with the Plan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlan

`func (o *GetApiWorkspacesOverview200ResponseAccount) SetPlan(v GetApiWorkspacesOverview200ResponseAccountPlan)`

SetPlan sets Plan field to given value.


### GetCharts

`func (o *GetApiWorkspacesOverview200ResponseAccount) GetCharts() GetApiWorkspacesOverview200ResponseAccountCharts`

GetCharts returns the Charts field if non-nil, zero value otherwise.

### GetChartsOk

`func (o *GetApiWorkspacesOverview200ResponseAccount) GetChartsOk() (*GetApiWorkspacesOverview200ResponseAccountCharts, bool)`

GetChartsOk returns a tuple with the Charts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCharts

`func (o *GetApiWorkspacesOverview200ResponseAccount) SetCharts(v GetApiWorkspacesOverview200ResponseAccountCharts)`

SetCharts sets Charts field to given value.


### GetEvents

`func (o *GetApiWorkspacesOverview200ResponseAccount) GetEvents() GetApiWorkspacesOverview200ResponseAccountCharts`

GetEvents returns the Events field if non-nil, zero value otherwise.

### GetEventsOk

`func (o *GetApiWorkspacesOverview200ResponseAccount) GetEventsOk() (*GetApiWorkspacesOverview200ResponseAccountCharts, bool)`

GetEventsOk returns a tuple with the Events field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvents

`func (o *GetApiWorkspacesOverview200ResponseAccount) SetEvents(v GetApiWorkspacesOverview200ResponseAccountCharts)`

SetEvents sets Events field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


