# PostApiWorkspacesRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DisplayName** | **string** |  | 
**IsTest** | Pointer to **bool** | Create a sandbox of the current workspace (prefer &#x60;POST /sandboxes&#x60;): never billed (usage shown, not counted), no chart/event caps, &#x60;sk_test_&#x60; API keys, a Sandbox ribbon for buyers. Permanent. | [optional] 
**CopyCharts** | Pointer to **bool** | Test workspaces only: copy the current workspace&#39;s charts as unpublished drafts (up to 50, archived skipped). Default false. | [optional] 
**ExternalId** | Pointer to **string** | Your own organizer id (letters, digits and &#x60;_ . ~ -&#x60;), unique among the workspaces you administer: 409 &#x60;external_id_taken&#x60; names the workspace that has it. Usable in &#x60;X-Workspace-Key&#x60; and &#x60;GET /workspaces/by-external-id/{externalId}&#x60;. | [optional] 
**TemplateWorkspaceId** | Pointer to **string** | A production workspace you administer (id or external id) whose integration settings the new workspace copies: every setting override (pricing source, own catalog, hold period, seats per hold token, ticket buyer validation, GA places to channels, new events not for sale, renderer and draft defaults, buyer extraData), chart validation levels, embed domains (up to the new plan limit) and webhook endpoints (same URL, event types and active flag, each with a NEW signing secret returned once). Not copied: members, API keys, charts, events and the two-factor requirement. 404 &#x60;template_not_found&#x60; otherwise. | [optional] 
**IssueKey** | Pointer to [**PostApiWorkspacesRequestIssueKey**](PostApiWorkspacesRequestIssueKey.md) |  | [optional] 

## Methods

### NewPostApiWorkspacesRequest

`func NewPostApiWorkspacesRequest(displayName string, ) *PostApiWorkspacesRequest`

NewPostApiWorkspacesRequest instantiates a new PostApiWorkspacesRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostApiWorkspacesRequestWithDefaults

`func NewPostApiWorkspacesRequestWithDefaults() *PostApiWorkspacesRequest`

NewPostApiWorkspacesRequestWithDefaults instantiates a new PostApiWorkspacesRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDisplayName

`func (o *PostApiWorkspacesRequest) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *PostApiWorkspacesRequest) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *PostApiWorkspacesRequest) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.


### GetIsTest

`func (o *PostApiWorkspacesRequest) GetIsTest() bool`

GetIsTest returns the IsTest field if non-nil, zero value otherwise.

### GetIsTestOk

`func (o *PostApiWorkspacesRequest) GetIsTestOk() (*bool, bool)`

GetIsTestOk returns a tuple with the IsTest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsTest

`func (o *PostApiWorkspacesRequest) SetIsTest(v bool)`

SetIsTest sets IsTest field to given value.

### HasIsTest

`func (o *PostApiWorkspacesRequest) HasIsTest() bool`

HasIsTest returns a boolean if a field has been set.

### GetCopyCharts

`func (o *PostApiWorkspacesRequest) GetCopyCharts() bool`

GetCopyCharts returns the CopyCharts field if non-nil, zero value otherwise.

### GetCopyChartsOk

`func (o *PostApiWorkspacesRequest) GetCopyChartsOk() (*bool, bool)`

GetCopyChartsOk returns a tuple with the CopyCharts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCopyCharts

`func (o *PostApiWorkspacesRequest) SetCopyCharts(v bool)`

SetCopyCharts sets CopyCharts field to given value.

### HasCopyCharts

`func (o *PostApiWorkspacesRequest) HasCopyCharts() bool`

HasCopyCharts returns a boolean if a field has been set.

### GetExternalId

`func (o *PostApiWorkspacesRequest) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *PostApiWorkspacesRequest) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *PostApiWorkspacesRequest) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *PostApiWorkspacesRequest) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### GetTemplateWorkspaceId

`func (o *PostApiWorkspacesRequest) GetTemplateWorkspaceId() string`

GetTemplateWorkspaceId returns the TemplateWorkspaceId field if non-nil, zero value otherwise.

### GetTemplateWorkspaceIdOk

`func (o *PostApiWorkspacesRequest) GetTemplateWorkspaceIdOk() (*string, bool)`

GetTemplateWorkspaceIdOk returns a tuple with the TemplateWorkspaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateWorkspaceId

`func (o *PostApiWorkspacesRequest) SetTemplateWorkspaceId(v string)`

SetTemplateWorkspaceId sets TemplateWorkspaceId field to given value.

### HasTemplateWorkspaceId

`func (o *PostApiWorkspacesRequest) HasTemplateWorkspaceId() bool`

HasTemplateWorkspaceId returns a boolean if a field has been set.

### GetIssueKey

`func (o *PostApiWorkspacesRequest) GetIssueKey() PostApiWorkspacesRequestIssueKey`

GetIssueKey returns the IssueKey field if non-nil, zero value otherwise.

### GetIssueKeyOk

`func (o *PostApiWorkspacesRequest) GetIssueKeyOk() (*PostApiWorkspacesRequestIssueKey, bool)`

GetIssueKeyOk returns a tuple with the IssueKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssueKey

`func (o *PostApiWorkspacesRequest) SetIssueKey(v PostApiWorkspacesRequestIssueKey)`

SetIssueKey sets IssueKey field to given value.

### HasIssueKey

`func (o *PostApiWorkspacesRequest) HasIssueKey() bool`

HasIssueKey returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


