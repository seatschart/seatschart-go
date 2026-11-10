# PostApiEventsByEventIdActionsUpdateExtraDataRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Objects** | [**[]PostApiEventsByEventIdActionsUpdateExtraDataRequestObjectsInner**](PostApiEventsByEventIdActionsUpdateExtraDataRequestObjectsInner.md) |  | 
**Merge** | Pointer to **bool** |  | [optional] 

## Methods

### NewPostApiEventsByEventIdActionsUpdateExtraDataRequest

`func NewPostApiEventsByEventIdActionsUpdateExtraDataRequest(objects []PostApiEventsByEventIdActionsUpdateExtraDataRequestObjectsInner, ) *PostApiEventsByEventIdActionsUpdateExtraDataRequest`

NewPostApiEventsByEventIdActionsUpdateExtraDataRequest instantiates a new PostApiEventsByEventIdActionsUpdateExtraDataRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostApiEventsByEventIdActionsUpdateExtraDataRequestWithDefaults

`func NewPostApiEventsByEventIdActionsUpdateExtraDataRequestWithDefaults() *PostApiEventsByEventIdActionsUpdateExtraDataRequest`

NewPostApiEventsByEventIdActionsUpdateExtraDataRequestWithDefaults instantiates a new PostApiEventsByEventIdActionsUpdateExtraDataRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetObjects

`func (o *PostApiEventsByEventIdActionsUpdateExtraDataRequest) GetObjects() []PostApiEventsByEventIdActionsUpdateExtraDataRequestObjectsInner`

GetObjects returns the Objects field if non-nil, zero value otherwise.

### GetObjectsOk

`func (o *PostApiEventsByEventIdActionsUpdateExtraDataRequest) GetObjectsOk() (*[]PostApiEventsByEventIdActionsUpdateExtraDataRequestObjectsInner, bool)`

GetObjectsOk returns a tuple with the Objects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjects

`func (o *PostApiEventsByEventIdActionsUpdateExtraDataRequest) SetObjects(v []PostApiEventsByEventIdActionsUpdateExtraDataRequestObjectsInner)`

SetObjects sets Objects field to given value.


### GetMerge

`func (o *PostApiEventsByEventIdActionsUpdateExtraDataRequest) GetMerge() bool`

GetMerge returns the Merge field if non-nil, zero value otherwise.

### GetMergeOk

`func (o *PostApiEventsByEventIdActionsUpdateExtraDataRequest) GetMergeOk() (*bool, bool)`

GetMergeOk returns a tuple with the Merge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerge

`func (o *PostApiEventsByEventIdActionsUpdateExtraDataRequest) SetMerge(v bool)`

SetMerge sets Merge field to given value.

### HasMerge

`func (o *PostApiEventsByEventIdActionsUpdateExtraDataRequest) HasMerge() bool`

HasMerge returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


