# PostApiEventsByEventIdActionsOverrideSeasonStatusRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Objects** | **[]string** | Seat labels or seat UUIDs of the path event. | 
**Release** | Pointer to **bool** | Also free the objects in this event (max 64). | [optional] 
**AllowBookedRelease** | Pointer to **bool** | With release: allow freeing booked / resale objects (season tickets). | [optional] 

## Methods

### NewPostApiEventsByEventIdActionsOverrideSeasonStatusRequest

`func NewPostApiEventsByEventIdActionsOverrideSeasonStatusRequest(objects []string, ) *PostApiEventsByEventIdActionsOverrideSeasonStatusRequest`

NewPostApiEventsByEventIdActionsOverrideSeasonStatusRequest instantiates a new PostApiEventsByEventIdActionsOverrideSeasonStatusRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostApiEventsByEventIdActionsOverrideSeasonStatusRequestWithDefaults

`func NewPostApiEventsByEventIdActionsOverrideSeasonStatusRequestWithDefaults() *PostApiEventsByEventIdActionsOverrideSeasonStatusRequest`

NewPostApiEventsByEventIdActionsOverrideSeasonStatusRequestWithDefaults instantiates a new PostApiEventsByEventIdActionsOverrideSeasonStatusRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetObjects

`func (o *PostApiEventsByEventIdActionsOverrideSeasonStatusRequest) GetObjects() []string`

GetObjects returns the Objects field if non-nil, zero value otherwise.

### GetObjectsOk

`func (o *PostApiEventsByEventIdActionsOverrideSeasonStatusRequest) GetObjectsOk() (*[]string, bool)`

GetObjectsOk returns a tuple with the Objects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjects

`func (o *PostApiEventsByEventIdActionsOverrideSeasonStatusRequest) SetObjects(v []string)`

SetObjects sets Objects field to given value.


### GetRelease

`func (o *PostApiEventsByEventIdActionsOverrideSeasonStatusRequest) GetRelease() bool`

GetRelease returns the Release field if non-nil, zero value otherwise.

### GetReleaseOk

`func (o *PostApiEventsByEventIdActionsOverrideSeasonStatusRequest) GetReleaseOk() (*bool, bool)`

GetReleaseOk returns a tuple with the Release field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelease

`func (o *PostApiEventsByEventIdActionsOverrideSeasonStatusRequest) SetRelease(v bool)`

SetRelease sets Release field to given value.

### HasRelease

`func (o *PostApiEventsByEventIdActionsOverrideSeasonStatusRequest) HasRelease() bool`

HasRelease returns a boolean if a field has been set.

### GetAllowBookedRelease

`func (o *PostApiEventsByEventIdActionsOverrideSeasonStatusRequest) GetAllowBookedRelease() bool`

GetAllowBookedRelease returns the AllowBookedRelease field if non-nil, zero value otherwise.

### GetAllowBookedReleaseOk

`func (o *PostApiEventsByEventIdActionsOverrideSeasonStatusRequest) GetAllowBookedReleaseOk() (*bool, bool)`

GetAllowBookedReleaseOk returns a tuple with the AllowBookedRelease field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowBookedRelease

`func (o *PostApiEventsByEventIdActionsOverrideSeasonStatusRequest) SetAllowBookedRelease(v bool)`

SetAllowBookedRelease sets AllowBookedRelease field to given value.

### HasAllowBookedRelease

`func (o *PostApiEventsByEventIdActionsOverrideSeasonStatusRequest) HasAllowBookedRelease() bool`

HasAllowBookedRelease returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


