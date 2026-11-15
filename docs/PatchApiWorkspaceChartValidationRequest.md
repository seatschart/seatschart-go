# PatchApiWorkspaceChartValidationRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Severities** | Pointer to [**PatchApiWorkspaceChartValidationRequestSeverities**](PatchApiWorkspaceChartValidationRequestSeverities.md) |  | [optional] 
**Reset** | Pointer to **bool** | Restore every rule to its default before applying &#x60;severities&#x60;. | [optional] 

## Methods

### NewPatchApiWorkspaceChartValidationRequest

`func NewPatchApiWorkspaceChartValidationRequest() *PatchApiWorkspaceChartValidationRequest`

NewPatchApiWorkspaceChartValidationRequest instantiates a new PatchApiWorkspaceChartValidationRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchApiWorkspaceChartValidationRequestWithDefaults

`func NewPatchApiWorkspaceChartValidationRequestWithDefaults() *PatchApiWorkspaceChartValidationRequest`

NewPatchApiWorkspaceChartValidationRequestWithDefaults instantiates a new PatchApiWorkspaceChartValidationRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSeverities

`func (o *PatchApiWorkspaceChartValidationRequest) GetSeverities() PatchApiWorkspaceChartValidationRequestSeverities`

GetSeverities returns the Severities field if non-nil, zero value otherwise.

### GetSeveritiesOk

`func (o *PatchApiWorkspaceChartValidationRequest) GetSeveritiesOk() (*PatchApiWorkspaceChartValidationRequestSeverities, bool)`

GetSeveritiesOk returns a tuple with the Severities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverities

`func (o *PatchApiWorkspaceChartValidationRequest) SetSeverities(v PatchApiWorkspaceChartValidationRequestSeverities)`

SetSeverities sets Severities field to given value.

### HasSeverities

`func (o *PatchApiWorkspaceChartValidationRequest) HasSeverities() bool`

HasSeverities returns a boolean if a field has been set.

### GetReset

`func (o *PatchApiWorkspaceChartValidationRequest) GetReset() bool`

GetReset returns the Reset field if non-nil, zero value otherwise.

### GetResetOk

`func (o *PatchApiWorkspaceChartValidationRequest) GetResetOk() (*bool, bool)`

GetResetOk returns a tuple with the Reset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReset

`func (o *PatchApiWorkspaceChartValidationRequest) SetReset(v bool)`

SetReset sets Reset field to given value.

### HasReset

`func (o *PatchApiWorkspaceChartValidationRequest) HasReset() bool`

HasReset returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


