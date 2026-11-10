# PostApiChartsByChartIdActionsValidateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Seats** | [**[]PostApiChartsByChartIdActionsValidateRequestSeatsInner**](PostApiChartsByChartIdActionsValidateRequestSeatsInner.md) |  | 
**Shapes** | Pointer to **[]interface{}** | Layout shapes (sections, zones, GA areas, tables, icons…) and one &#x60;layoutMeta&#x60;. &#x60;layoutMeta.settings3d&#x60; holds the authored 3D (elevation / rake per section or zone, floors, light preset, spectator target) and the 3D billboards: &#x60;billboards&#x60; (at most 64 boards, each &#x60;{ id, points: [{x, y}] (2-64, map px), height?: m (0.2-20, default 1), elevation?: m (-20-100), scroll?: \&quot;reverse\&quot; | \&quot;off\&quot; }&#x60;) and &#x60;billboardImages&#x60; (at most 12 sponsor images shared by every board, in scroll order: base64 WebP / PNG / JPEG data URLs, stored to the map-asset CDN when configured, or map-asset URLs). Invalid values are refused with 400 &#x60;validation_error&#x60;, never clamped. &#x60;settings3d.seatPicking&#x60; (boolean) is the chart-level \&quot;Allow 3D seat picking\&quot; switch (absent: on for charts that author 3D). Optional fields published only when set: tables &#x60;oneSide&#x60; (rectangular, every guest on one long side) and &#x60;openSpaces&#x60; (1-20 empty chair positions); text &#x60;bold&#x60; / &#x60;italic&#x60;; decorative shapes &#x60;layer: \&quot;foreground\&quot;&#x60; (drawn above the seats for buyers, never interactive); &#x60;layoutMeta.floors[].abbreviation&#x60; (at most 6 characters) and &#x60;layoutMeta.displayTypes&#x60; (object id → displayed type, at most 64 characters, seats.io &#x60;displayObjectType&#x60;). Designer-only metadata (&#x60;bytes&#x60; on images, &#x60;layoutMeta.backgroundBytes&#x60;) feeds the &#x60;images_size&#x60; check and is never published. | [optional] 
**Background** | Pointer to **NullableString** |  | [optional] 
**Reference** | Pointer to **NullableString** |  | [optional] 
**BaseDraftUpdatedAt** | Pointer to **NullableString** | Draft version this save is based on (&#x60;draftUpdatedAt&#x60; from the last load or save; null for a never-saved draft). When given and the stored draft changed meanwhile, the save is refused with 409 &#x60;draft_conflict&#x60; (designer autosave). Omit it for last-writer-wins. | [optional] 

## Methods

### NewPostApiChartsByChartIdActionsValidateRequest

`func NewPostApiChartsByChartIdActionsValidateRequest(seats []PostApiChartsByChartIdActionsValidateRequestSeatsInner, ) *PostApiChartsByChartIdActionsValidateRequest`

NewPostApiChartsByChartIdActionsValidateRequest instantiates a new PostApiChartsByChartIdActionsValidateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostApiChartsByChartIdActionsValidateRequestWithDefaults

`func NewPostApiChartsByChartIdActionsValidateRequestWithDefaults() *PostApiChartsByChartIdActionsValidateRequest`

NewPostApiChartsByChartIdActionsValidateRequestWithDefaults instantiates a new PostApiChartsByChartIdActionsValidateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSeats

`func (o *PostApiChartsByChartIdActionsValidateRequest) GetSeats() []PostApiChartsByChartIdActionsValidateRequestSeatsInner`

GetSeats returns the Seats field if non-nil, zero value otherwise.

### GetSeatsOk

`func (o *PostApiChartsByChartIdActionsValidateRequest) GetSeatsOk() (*[]PostApiChartsByChartIdActionsValidateRequestSeatsInner, bool)`

GetSeatsOk returns a tuple with the Seats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeats

`func (o *PostApiChartsByChartIdActionsValidateRequest) SetSeats(v []PostApiChartsByChartIdActionsValidateRequestSeatsInner)`

SetSeats sets Seats field to given value.


### GetShapes

`func (o *PostApiChartsByChartIdActionsValidateRequest) GetShapes() []interface{}`

GetShapes returns the Shapes field if non-nil, zero value otherwise.

### GetShapesOk

`func (o *PostApiChartsByChartIdActionsValidateRequest) GetShapesOk() (*[]interface{}, bool)`

GetShapesOk returns a tuple with the Shapes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShapes

`func (o *PostApiChartsByChartIdActionsValidateRequest) SetShapes(v []interface{})`

SetShapes sets Shapes field to given value.

### HasShapes

`func (o *PostApiChartsByChartIdActionsValidateRequest) HasShapes() bool`

HasShapes returns a boolean if a field has been set.

### GetBackground

`func (o *PostApiChartsByChartIdActionsValidateRequest) GetBackground() string`

GetBackground returns the Background field if non-nil, zero value otherwise.

### GetBackgroundOk

`func (o *PostApiChartsByChartIdActionsValidateRequest) GetBackgroundOk() (*string, bool)`

GetBackgroundOk returns a tuple with the Background field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackground

`func (o *PostApiChartsByChartIdActionsValidateRequest) SetBackground(v string)`

SetBackground sets Background field to given value.

### HasBackground

`func (o *PostApiChartsByChartIdActionsValidateRequest) HasBackground() bool`

HasBackground returns a boolean if a field has been set.

### SetBackgroundNil

`func (o *PostApiChartsByChartIdActionsValidateRequest) SetBackgroundNil(b bool)`

 SetBackgroundNil sets the value for Background to be an explicit nil

### UnsetBackground
`func (o *PostApiChartsByChartIdActionsValidateRequest) UnsetBackground()`

UnsetBackground ensures that no value is present for Background, not even an explicit nil
### GetReference

`func (o *PostApiChartsByChartIdActionsValidateRequest) GetReference() string`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *PostApiChartsByChartIdActionsValidateRequest) GetReferenceOk() (*string, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *PostApiChartsByChartIdActionsValidateRequest) SetReference(v string)`

SetReference sets Reference field to given value.

### HasReference

`func (o *PostApiChartsByChartIdActionsValidateRequest) HasReference() bool`

HasReference returns a boolean if a field has been set.

### SetReferenceNil

`func (o *PostApiChartsByChartIdActionsValidateRequest) SetReferenceNil(b bool)`

 SetReferenceNil sets the value for Reference to be an explicit nil

### UnsetReference
`func (o *PostApiChartsByChartIdActionsValidateRequest) UnsetReference()`

UnsetReference ensures that no value is present for Reference, not even an explicit nil
### GetBaseDraftUpdatedAt

`func (o *PostApiChartsByChartIdActionsValidateRequest) GetBaseDraftUpdatedAt() string`

GetBaseDraftUpdatedAt returns the BaseDraftUpdatedAt field if non-nil, zero value otherwise.

### GetBaseDraftUpdatedAtOk

`func (o *PostApiChartsByChartIdActionsValidateRequest) GetBaseDraftUpdatedAtOk() (*string, bool)`

GetBaseDraftUpdatedAtOk returns a tuple with the BaseDraftUpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseDraftUpdatedAt

`func (o *PostApiChartsByChartIdActionsValidateRequest) SetBaseDraftUpdatedAt(v string)`

SetBaseDraftUpdatedAt sets BaseDraftUpdatedAt field to given value.

### HasBaseDraftUpdatedAt

`func (o *PostApiChartsByChartIdActionsValidateRequest) HasBaseDraftUpdatedAt() bool`

HasBaseDraftUpdatedAt returns a boolean if a field has been set.

### SetBaseDraftUpdatedAtNil

`func (o *PostApiChartsByChartIdActionsValidateRequest) SetBaseDraftUpdatedAtNil(b bool)`

 SetBaseDraftUpdatedAtNil sets the value for BaseDraftUpdatedAt to be an explicit nil

### UnsetBaseDraftUpdatedAt
`func (o *PostApiChartsByChartIdActionsValidateRequest) UnsetBaseDraftUpdatedAt()`

UnsetBaseDraftUpdatedAt ensures that no value is present for BaseDraftUpdatedAt, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


