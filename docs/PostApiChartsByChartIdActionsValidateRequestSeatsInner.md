# PostApiChartsByChartIdActionsValidateRequestSeatsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Label** | **string** |  | 
**PosX** | **float32** |  | 
**PosY** | **float32** |  | 
**Category** | Pointer to **string** |  | [optional] 
**Price** | Pointer to **float32** |  | [optional] 
**RowLabel** | Pointer to **string** | Row label (structural; e.g. \&quot;A\&quot;). | [optional] 
**SectionId** | Pointer to **string** | Section scope (structural). Seat labels must be unique within a section. | [optional] 
**DisplayedLabel** | Pointer to **string** | Short label printed on the seat (e.g. \&quot;12\&quot;). &#x60;label&#x60; stays the API identity. | [optional] 
**Accessibility** | Pointer to [**PostApiChartsByChartIdActionsValidateRequestSeatsInnerAccessibility**](PostApiChartsByChartIdActionsValidateRequestSeatsInnerAccessibility.md) |  | [optional] 
**Entrance** | Pointer to **string** | Entrance shown to buyers (e.g. \&quot;Gate 3\&quot;). Rows (layoutMeta.rowSettings), tables, sections and GA areas (an \&quot;entrance\&quot; field on their layout shape) can set one too; the most specific wins: seat, row, table, section. | [optional] 

## Methods

### NewPostApiChartsByChartIdActionsValidateRequestSeatsInner

`func NewPostApiChartsByChartIdActionsValidateRequestSeatsInner(label string, posX float32, posY float32, ) *PostApiChartsByChartIdActionsValidateRequestSeatsInner`

NewPostApiChartsByChartIdActionsValidateRequestSeatsInner instantiates a new PostApiChartsByChartIdActionsValidateRequestSeatsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostApiChartsByChartIdActionsValidateRequestSeatsInnerWithDefaults

`func NewPostApiChartsByChartIdActionsValidateRequestSeatsInnerWithDefaults() *PostApiChartsByChartIdActionsValidateRequestSeatsInner`

NewPostApiChartsByChartIdActionsValidateRequestSeatsInnerWithDefaults instantiates a new PostApiChartsByChartIdActionsValidateRequestSeatsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PostApiChartsByChartIdActionsValidateRequestSeatsInner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PostApiChartsByChartIdActionsValidateRequestSeatsInner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PostApiChartsByChartIdActionsValidateRequestSeatsInner) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *PostApiChartsByChartIdActionsValidateRequestSeatsInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLabel

`func (o *PostApiChartsByChartIdActionsValidateRequestSeatsInner) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *PostApiChartsByChartIdActionsValidateRequestSeatsInner) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *PostApiChartsByChartIdActionsValidateRequestSeatsInner) SetLabel(v string)`

SetLabel sets Label field to given value.


### GetPosX

`func (o *PostApiChartsByChartIdActionsValidateRequestSeatsInner) GetPosX() float32`

GetPosX returns the PosX field if non-nil, zero value otherwise.

### GetPosXOk

`func (o *PostApiChartsByChartIdActionsValidateRequestSeatsInner) GetPosXOk() (*float32, bool)`

GetPosXOk returns a tuple with the PosX field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPosX

`func (o *PostApiChartsByChartIdActionsValidateRequestSeatsInner) SetPosX(v float32)`

SetPosX sets PosX field to given value.


### GetPosY

`func (o *PostApiChartsByChartIdActionsValidateRequestSeatsInner) GetPosY() float32`

GetPosY returns the PosY field if non-nil, zero value otherwise.

### GetPosYOk

`func (o *PostApiChartsByChartIdActionsValidateRequestSeatsInner) GetPosYOk() (*float32, bool)`

GetPosYOk returns a tuple with the PosY field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPosY

`func (o *PostApiChartsByChartIdActionsValidateRequestSeatsInner) SetPosY(v float32)`

SetPosY sets PosY field to given value.


### GetCategory

`func (o *PostApiChartsByChartIdActionsValidateRequestSeatsInner) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *PostApiChartsByChartIdActionsValidateRequestSeatsInner) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *PostApiChartsByChartIdActionsValidateRequestSeatsInner) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *PostApiChartsByChartIdActionsValidateRequestSeatsInner) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### GetPrice

`func (o *PostApiChartsByChartIdActionsValidateRequestSeatsInner) GetPrice() float32`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *PostApiChartsByChartIdActionsValidateRequestSeatsInner) GetPriceOk() (*float32, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *PostApiChartsByChartIdActionsValidateRequestSeatsInner) SetPrice(v float32)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *PostApiChartsByChartIdActionsValidateRequestSeatsInner) HasPrice() bool`

HasPrice returns a boolean if a field has been set.

### GetRowLabel

`func (o *PostApiChartsByChartIdActionsValidateRequestSeatsInner) GetRowLabel() string`

GetRowLabel returns the RowLabel field if non-nil, zero value otherwise.

### GetRowLabelOk

`func (o *PostApiChartsByChartIdActionsValidateRequestSeatsInner) GetRowLabelOk() (*string, bool)`

GetRowLabelOk returns a tuple with the RowLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRowLabel

`func (o *PostApiChartsByChartIdActionsValidateRequestSeatsInner) SetRowLabel(v string)`

SetRowLabel sets RowLabel field to given value.

### HasRowLabel

`func (o *PostApiChartsByChartIdActionsValidateRequestSeatsInner) HasRowLabel() bool`

HasRowLabel returns a boolean if a field has been set.

### GetSectionId

`func (o *PostApiChartsByChartIdActionsValidateRequestSeatsInner) GetSectionId() string`

GetSectionId returns the SectionId field if non-nil, zero value otherwise.

### GetSectionIdOk

`func (o *PostApiChartsByChartIdActionsValidateRequestSeatsInner) GetSectionIdOk() (*string, bool)`

GetSectionIdOk returns a tuple with the SectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSectionId

`func (o *PostApiChartsByChartIdActionsValidateRequestSeatsInner) SetSectionId(v string)`

SetSectionId sets SectionId field to given value.

### HasSectionId

`func (o *PostApiChartsByChartIdActionsValidateRequestSeatsInner) HasSectionId() bool`

HasSectionId returns a boolean if a field has been set.

### GetDisplayedLabel

`func (o *PostApiChartsByChartIdActionsValidateRequestSeatsInner) GetDisplayedLabel() string`

GetDisplayedLabel returns the DisplayedLabel field if non-nil, zero value otherwise.

### GetDisplayedLabelOk

`func (o *PostApiChartsByChartIdActionsValidateRequestSeatsInner) GetDisplayedLabelOk() (*string, bool)`

GetDisplayedLabelOk returns a tuple with the DisplayedLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayedLabel

`func (o *PostApiChartsByChartIdActionsValidateRequestSeatsInner) SetDisplayedLabel(v string)`

SetDisplayedLabel sets DisplayedLabel field to given value.

### HasDisplayedLabel

`func (o *PostApiChartsByChartIdActionsValidateRequestSeatsInner) HasDisplayedLabel() bool`

HasDisplayedLabel returns a boolean if a field has been set.

### GetAccessibility

`func (o *PostApiChartsByChartIdActionsValidateRequestSeatsInner) GetAccessibility() PostApiChartsByChartIdActionsValidateRequestSeatsInnerAccessibility`

GetAccessibility returns the Accessibility field if non-nil, zero value otherwise.

### GetAccessibilityOk

`func (o *PostApiChartsByChartIdActionsValidateRequestSeatsInner) GetAccessibilityOk() (*PostApiChartsByChartIdActionsValidateRequestSeatsInnerAccessibility, bool)`

GetAccessibilityOk returns a tuple with the Accessibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessibility

`func (o *PostApiChartsByChartIdActionsValidateRequestSeatsInner) SetAccessibility(v PostApiChartsByChartIdActionsValidateRequestSeatsInnerAccessibility)`

SetAccessibility sets Accessibility field to given value.

### HasAccessibility

`func (o *PostApiChartsByChartIdActionsValidateRequestSeatsInner) HasAccessibility() bool`

HasAccessibility returns a boolean if a field has been set.

### GetEntrance

`func (o *PostApiChartsByChartIdActionsValidateRequestSeatsInner) GetEntrance() string`

GetEntrance returns the Entrance field if non-nil, zero value otherwise.

### GetEntranceOk

`func (o *PostApiChartsByChartIdActionsValidateRequestSeatsInner) GetEntranceOk() (*string, bool)`

GetEntranceOk returns a tuple with the Entrance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntrance

`func (o *PostApiChartsByChartIdActionsValidateRequestSeatsInner) SetEntrance(v string)`

SetEntrance sets Entrance field to given value.

### HasEntrance

`func (o *PostApiChartsByChartIdActionsValidateRequestSeatsInner) HasEntrance() bool`

HasEntrance returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


