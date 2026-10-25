# PutApiChartsByChartIdMapDraftRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Seats** | [**[]PostApiChartsByChartIdActionsValidateRequestSeatsInner**](PostApiChartsByChartIdActionsValidateRequestSeatsInner.md) |  | 
**Shapes** | Pointer to **[]interface{}** |  | [optional] 
**Background** | Pointer to **NullableString** |  | [optional] 
**Reference** | Pointer to **NullableString** |  | [optional] 
**BaseDraftUpdatedAt** | Pointer to **NullableString** | Draft version this save is based on (&#x60;draftUpdatedAt&#x60; from the last load or save; null for a never-saved draft). When given and the stored draft changed meanwhile, the save is refused with 409 &#x60;draft_conflict&#x60; (designer autosave). Omit it for last-writer-wins. | [optional] 

## Methods

### NewPutApiChartsByChartIdMapDraftRequest

`func NewPutApiChartsByChartIdMapDraftRequest(seats []PostApiChartsByChartIdActionsValidateRequestSeatsInner, ) *PutApiChartsByChartIdMapDraftRequest`

NewPutApiChartsByChartIdMapDraftRequest instantiates a new PutApiChartsByChartIdMapDraftRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPutApiChartsByChartIdMapDraftRequestWithDefaults

`func NewPutApiChartsByChartIdMapDraftRequestWithDefaults() *PutApiChartsByChartIdMapDraftRequest`

NewPutApiChartsByChartIdMapDraftRequestWithDefaults instantiates a new PutApiChartsByChartIdMapDraftRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSeats

`func (o *PutApiChartsByChartIdMapDraftRequest) GetSeats() []PostApiChartsByChartIdActionsValidateRequestSeatsInner`

GetSeats returns the Seats field if non-nil, zero value otherwise.

### GetSeatsOk

`func (o *PutApiChartsByChartIdMapDraftRequest) GetSeatsOk() (*[]PostApiChartsByChartIdActionsValidateRequestSeatsInner, bool)`

GetSeatsOk returns a tuple with the Seats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeats

`func (o *PutApiChartsByChartIdMapDraftRequest) SetSeats(v []PostApiChartsByChartIdActionsValidateRequestSeatsInner)`

SetSeats sets Seats field to given value.


### GetShapes

`func (o *PutApiChartsByChartIdMapDraftRequest) GetShapes() []interface{}`

GetShapes returns the Shapes field if non-nil, zero value otherwise.

### GetShapesOk

`func (o *PutApiChartsByChartIdMapDraftRequest) GetShapesOk() (*[]interface{}, bool)`

GetShapesOk returns a tuple with the Shapes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShapes

`func (o *PutApiChartsByChartIdMapDraftRequest) SetShapes(v []interface{})`

SetShapes sets Shapes field to given value.

### HasShapes

`func (o *PutApiChartsByChartIdMapDraftRequest) HasShapes() bool`

HasShapes returns a boolean if a field has been set.

### GetBackground

`func (o *PutApiChartsByChartIdMapDraftRequest) GetBackground() string`

GetBackground returns the Background field if non-nil, zero value otherwise.

### GetBackgroundOk

`func (o *PutApiChartsByChartIdMapDraftRequest) GetBackgroundOk() (*string, bool)`

GetBackgroundOk returns a tuple with the Background field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackground

`func (o *PutApiChartsByChartIdMapDraftRequest) SetBackground(v string)`

SetBackground sets Background field to given value.

### HasBackground

`func (o *PutApiChartsByChartIdMapDraftRequest) HasBackground() bool`

HasBackground returns a boolean if a field has been set.

### SetBackgroundNil

`func (o *PutApiChartsByChartIdMapDraftRequest) SetBackgroundNil(b bool)`

 SetBackgroundNil sets the value for Background to be an explicit nil

### UnsetBackground
`func (o *PutApiChartsByChartIdMapDraftRequest) UnsetBackground()`

UnsetBackground ensures that no value is present for Background, not even an explicit nil
### GetReference

`func (o *PutApiChartsByChartIdMapDraftRequest) GetReference() string`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *PutApiChartsByChartIdMapDraftRequest) GetReferenceOk() (*string, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *PutApiChartsByChartIdMapDraftRequest) SetReference(v string)`

SetReference sets Reference field to given value.

### HasReference

`func (o *PutApiChartsByChartIdMapDraftRequest) HasReference() bool`

HasReference returns a boolean if a field has been set.

### SetReferenceNil

`func (o *PutApiChartsByChartIdMapDraftRequest) SetReferenceNil(b bool)`

 SetReferenceNil sets the value for Reference to be an explicit nil

### UnsetReference
`func (o *PutApiChartsByChartIdMapDraftRequest) UnsetReference()`

UnsetReference ensures that no value is present for Reference, not even an explicit nil
### GetBaseDraftUpdatedAt

`func (o *PutApiChartsByChartIdMapDraftRequest) GetBaseDraftUpdatedAt() string`

GetBaseDraftUpdatedAt returns the BaseDraftUpdatedAt field if non-nil, zero value otherwise.

### GetBaseDraftUpdatedAtOk

`func (o *PutApiChartsByChartIdMapDraftRequest) GetBaseDraftUpdatedAtOk() (*string, bool)`

GetBaseDraftUpdatedAtOk returns a tuple with the BaseDraftUpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseDraftUpdatedAt

`func (o *PutApiChartsByChartIdMapDraftRequest) SetBaseDraftUpdatedAt(v string)`

SetBaseDraftUpdatedAt sets BaseDraftUpdatedAt field to given value.

### HasBaseDraftUpdatedAt

`func (o *PutApiChartsByChartIdMapDraftRequest) HasBaseDraftUpdatedAt() bool`

HasBaseDraftUpdatedAt returns a boolean if a field has been set.

### SetBaseDraftUpdatedAtNil

`func (o *PutApiChartsByChartIdMapDraftRequest) SetBaseDraftUpdatedAtNil(b bool)`

 SetBaseDraftUpdatedAtNil sets the value for BaseDraftUpdatedAt to be an explicit nil

### UnsetBaseDraftUpdatedAt
`func (o *PutApiChartsByChartIdMapDraftRequest) UnsetBaseDraftUpdatedAt()`

UnsetBaseDraftUpdatedAt ensures that no value is present for BaseDraftUpdatedAt, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


