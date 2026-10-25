# PostApiChartsByChartIdMapPublish422ResponseIssuesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | **string** | Chart validation rule id. | 
**Severity** | **string** |  | 
**ObjectIds** | **[]string** | Seat, GA area, section, zone or floor ids, depending on the code. | 
**Message** | **string** |  | 
**Count** | **int32** |  | 
**Label** | Pointer to **string** |  | [optional] 
**SectionId** | Pointer to **string** |  | [optional] 
**TableIds** | Pointer to **[]string** |  | [optional] 
**Scope** | Pointer to **string** | Section or floor name the offenders belong to (empty &#x3D; whole map). | [optional] 
**FloorId** | Pointer to **string** |  | [optional] 

## Methods

### NewPostApiChartsByChartIdMapPublish422ResponseIssuesInner

`func NewPostApiChartsByChartIdMapPublish422ResponseIssuesInner(code string, severity string, objectIds []string, message string, count int32, ) *PostApiChartsByChartIdMapPublish422ResponseIssuesInner`

NewPostApiChartsByChartIdMapPublish422ResponseIssuesInner instantiates a new PostApiChartsByChartIdMapPublish422ResponseIssuesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostApiChartsByChartIdMapPublish422ResponseIssuesInnerWithDefaults

`func NewPostApiChartsByChartIdMapPublish422ResponseIssuesInnerWithDefaults() *PostApiChartsByChartIdMapPublish422ResponseIssuesInner`

NewPostApiChartsByChartIdMapPublish422ResponseIssuesInnerWithDefaults instantiates a new PostApiChartsByChartIdMapPublish422ResponseIssuesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *PostApiChartsByChartIdMapPublish422ResponseIssuesInner) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *PostApiChartsByChartIdMapPublish422ResponseIssuesInner) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *PostApiChartsByChartIdMapPublish422ResponseIssuesInner) SetCode(v string)`

SetCode sets Code field to given value.


### GetSeverity

`func (o *PostApiChartsByChartIdMapPublish422ResponseIssuesInner) GetSeverity() string`

GetSeverity returns the Severity field if non-nil, zero value otherwise.

### GetSeverityOk

`func (o *PostApiChartsByChartIdMapPublish422ResponseIssuesInner) GetSeverityOk() (*string, bool)`

GetSeverityOk returns a tuple with the Severity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverity

`func (o *PostApiChartsByChartIdMapPublish422ResponseIssuesInner) SetSeverity(v string)`

SetSeverity sets Severity field to given value.


### GetObjectIds

`func (o *PostApiChartsByChartIdMapPublish422ResponseIssuesInner) GetObjectIds() []string`

GetObjectIds returns the ObjectIds field if non-nil, zero value otherwise.

### GetObjectIdsOk

`func (o *PostApiChartsByChartIdMapPublish422ResponseIssuesInner) GetObjectIdsOk() (*[]string, bool)`

GetObjectIdsOk returns a tuple with the ObjectIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectIds

`func (o *PostApiChartsByChartIdMapPublish422ResponseIssuesInner) SetObjectIds(v []string)`

SetObjectIds sets ObjectIds field to given value.


### GetMessage

`func (o *PostApiChartsByChartIdMapPublish422ResponseIssuesInner) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *PostApiChartsByChartIdMapPublish422ResponseIssuesInner) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *PostApiChartsByChartIdMapPublish422ResponseIssuesInner) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetCount

`func (o *PostApiChartsByChartIdMapPublish422ResponseIssuesInner) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *PostApiChartsByChartIdMapPublish422ResponseIssuesInner) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *PostApiChartsByChartIdMapPublish422ResponseIssuesInner) SetCount(v int32)`

SetCount sets Count field to given value.


### GetLabel

`func (o *PostApiChartsByChartIdMapPublish422ResponseIssuesInner) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *PostApiChartsByChartIdMapPublish422ResponseIssuesInner) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *PostApiChartsByChartIdMapPublish422ResponseIssuesInner) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *PostApiChartsByChartIdMapPublish422ResponseIssuesInner) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetSectionId

`func (o *PostApiChartsByChartIdMapPublish422ResponseIssuesInner) GetSectionId() string`

GetSectionId returns the SectionId field if non-nil, zero value otherwise.

### GetSectionIdOk

`func (o *PostApiChartsByChartIdMapPublish422ResponseIssuesInner) GetSectionIdOk() (*string, bool)`

GetSectionIdOk returns a tuple with the SectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSectionId

`func (o *PostApiChartsByChartIdMapPublish422ResponseIssuesInner) SetSectionId(v string)`

SetSectionId sets SectionId field to given value.

### HasSectionId

`func (o *PostApiChartsByChartIdMapPublish422ResponseIssuesInner) HasSectionId() bool`

HasSectionId returns a boolean if a field has been set.

### GetTableIds

`func (o *PostApiChartsByChartIdMapPublish422ResponseIssuesInner) GetTableIds() []string`

GetTableIds returns the TableIds field if non-nil, zero value otherwise.

### GetTableIdsOk

`func (o *PostApiChartsByChartIdMapPublish422ResponseIssuesInner) GetTableIdsOk() (*[]string, bool)`

GetTableIdsOk returns a tuple with the TableIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTableIds

`func (o *PostApiChartsByChartIdMapPublish422ResponseIssuesInner) SetTableIds(v []string)`

SetTableIds sets TableIds field to given value.

### HasTableIds

`func (o *PostApiChartsByChartIdMapPublish422ResponseIssuesInner) HasTableIds() bool`

HasTableIds returns a boolean if a field has been set.

### GetScope

`func (o *PostApiChartsByChartIdMapPublish422ResponseIssuesInner) GetScope() string`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *PostApiChartsByChartIdMapPublish422ResponseIssuesInner) GetScopeOk() (*string, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *PostApiChartsByChartIdMapPublish422ResponseIssuesInner) SetScope(v string)`

SetScope sets Scope field to given value.

### HasScope

`func (o *PostApiChartsByChartIdMapPublish422ResponseIssuesInner) HasScope() bool`

HasScope returns a boolean if a field has been set.

### GetFloorId

`func (o *PostApiChartsByChartIdMapPublish422ResponseIssuesInner) GetFloorId() string`

GetFloorId returns the FloorId field if non-nil, zero value otherwise.

### GetFloorIdOk

`func (o *PostApiChartsByChartIdMapPublish422ResponseIssuesInner) GetFloorIdOk() (*string, bool)`

GetFloorIdOk returns a tuple with the FloorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFloorId

`func (o *PostApiChartsByChartIdMapPublish422ResponseIssuesInner) SetFloorId(v string)`

SetFloorId sets FloorId field to given value.

### HasFloorId

`func (o *PostApiChartsByChartIdMapPublish422ResponseIssuesInner) HasFloorId() bool`

HasFloorId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


