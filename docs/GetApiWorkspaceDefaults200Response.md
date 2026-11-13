# GetApiWorkspaceDefaults200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Renderer** | [**GetApiWorkspaceDefaults200ResponseRenderer**](GetApiWorkspaceDefaults200ResponseRenderer.md) |  | 
**RendererBuiltIn** | [**GetApiWorkspaceDefaults200ResponseRendererBuiltIn**](GetApiWorkspaceDefaults200ResponseRendererBuiltIn.md) |  | 
**DraftChartDrawingsEnabled** | **bool** | Designer changes go to a draft first (default true). false: every designer save is published directly. | 

## Methods

### NewGetApiWorkspaceDefaults200Response

`func NewGetApiWorkspaceDefaults200Response(renderer GetApiWorkspaceDefaults200ResponseRenderer, rendererBuiltIn GetApiWorkspaceDefaults200ResponseRendererBuiltIn, draftChartDrawingsEnabled bool, ) *GetApiWorkspaceDefaults200Response`

NewGetApiWorkspaceDefaults200Response instantiates a new GetApiWorkspaceDefaults200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetApiWorkspaceDefaults200ResponseWithDefaults

`func NewGetApiWorkspaceDefaults200ResponseWithDefaults() *GetApiWorkspaceDefaults200Response`

NewGetApiWorkspaceDefaults200ResponseWithDefaults instantiates a new GetApiWorkspaceDefaults200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRenderer

`func (o *GetApiWorkspaceDefaults200Response) GetRenderer() GetApiWorkspaceDefaults200ResponseRenderer`

GetRenderer returns the Renderer field if non-nil, zero value otherwise.

### GetRendererOk

`func (o *GetApiWorkspaceDefaults200Response) GetRendererOk() (*GetApiWorkspaceDefaults200ResponseRenderer, bool)`

GetRendererOk returns a tuple with the Renderer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRenderer

`func (o *GetApiWorkspaceDefaults200Response) SetRenderer(v GetApiWorkspaceDefaults200ResponseRenderer)`

SetRenderer sets Renderer field to given value.


### GetRendererBuiltIn

`func (o *GetApiWorkspaceDefaults200Response) GetRendererBuiltIn() GetApiWorkspaceDefaults200ResponseRendererBuiltIn`

GetRendererBuiltIn returns the RendererBuiltIn field if non-nil, zero value otherwise.

### GetRendererBuiltInOk

`func (o *GetApiWorkspaceDefaults200Response) GetRendererBuiltInOk() (*GetApiWorkspaceDefaults200ResponseRendererBuiltIn, bool)`

GetRendererBuiltInOk returns a tuple with the RendererBuiltIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRendererBuiltIn

`func (o *GetApiWorkspaceDefaults200Response) SetRendererBuiltIn(v GetApiWorkspaceDefaults200ResponseRendererBuiltIn)`

SetRendererBuiltIn sets RendererBuiltIn field to given value.


### GetDraftChartDrawingsEnabled

`func (o *GetApiWorkspaceDefaults200Response) GetDraftChartDrawingsEnabled() bool`

GetDraftChartDrawingsEnabled returns the DraftChartDrawingsEnabled field if non-nil, zero value otherwise.

### GetDraftChartDrawingsEnabledOk

`func (o *GetApiWorkspaceDefaults200Response) GetDraftChartDrawingsEnabledOk() (*bool, bool)`

GetDraftChartDrawingsEnabledOk returns a tuple with the DraftChartDrawingsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDraftChartDrawingsEnabled

`func (o *GetApiWorkspaceDefaults200Response) SetDraftChartDrawingsEnabled(v bool)`

SetDraftChartDrawingsEnabled sets DraftChartDrawingsEnabled field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


