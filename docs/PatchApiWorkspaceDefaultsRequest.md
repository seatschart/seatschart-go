# PatchApiWorkspaceDefaultsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Renderer** | Pointer to [**PatchApiWorkspaceDefaultsRequestRenderer**](PatchApiWorkspaceDefaultsRequestRenderer.md) |  | [optional] 
**DraftChartDrawingsEnabled** | Pointer to **bool** | &#x60;false&#x60;: publish designer changes directly (no draft step). | [optional] 

## Methods

### NewPatchApiWorkspaceDefaultsRequest

`func NewPatchApiWorkspaceDefaultsRequest() *PatchApiWorkspaceDefaultsRequest`

NewPatchApiWorkspaceDefaultsRequest instantiates a new PatchApiWorkspaceDefaultsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchApiWorkspaceDefaultsRequestWithDefaults

`func NewPatchApiWorkspaceDefaultsRequestWithDefaults() *PatchApiWorkspaceDefaultsRequest`

NewPatchApiWorkspaceDefaultsRequestWithDefaults instantiates a new PatchApiWorkspaceDefaultsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRenderer

`func (o *PatchApiWorkspaceDefaultsRequest) GetRenderer() PatchApiWorkspaceDefaultsRequestRenderer`

GetRenderer returns the Renderer field if non-nil, zero value otherwise.

### GetRendererOk

`func (o *PatchApiWorkspaceDefaultsRequest) GetRendererOk() (*PatchApiWorkspaceDefaultsRequestRenderer, bool)`

GetRendererOk returns a tuple with the Renderer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRenderer

`func (o *PatchApiWorkspaceDefaultsRequest) SetRenderer(v PatchApiWorkspaceDefaultsRequestRenderer)`

SetRenderer sets Renderer field to given value.

### HasRenderer

`func (o *PatchApiWorkspaceDefaultsRequest) HasRenderer() bool`

HasRenderer returns a boolean if a field has been set.

### GetDraftChartDrawingsEnabled

`func (o *PatchApiWorkspaceDefaultsRequest) GetDraftChartDrawingsEnabled() bool`

GetDraftChartDrawingsEnabled returns the DraftChartDrawingsEnabled field if non-nil, zero value otherwise.

### GetDraftChartDrawingsEnabledOk

`func (o *PatchApiWorkspaceDefaultsRequest) GetDraftChartDrawingsEnabledOk() (*bool, bool)`

GetDraftChartDrawingsEnabledOk returns a tuple with the DraftChartDrawingsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDraftChartDrawingsEnabled

`func (o *PatchApiWorkspaceDefaultsRequest) SetDraftChartDrawingsEnabled(v bool)`

SetDraftChartDrawingsEnabled sets DraftChartDrawingsEnabled field to given value.

### HasDraftChartDrawingsEnabled

`func (o *PatchApiWorkspaceDefaultsRequest) HasDraftChartDrawingsEnabled() bool`

HasDraftChartDrawingsEnabled returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


