# PostApiSeasonsBySeasonIdActionsAddEventsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EventIds** | **[]string** |  | 
**Sync** | Pointer to **bool** | Sync inventory from the source into the new members (default true). | [optional] 

## Methods

### NewPostApiSeasonsBySeasonIdActionsAddEventsRequest

`func NewPostApiSeasonsBySeasonIdActionsAddEventsRequest(eventIds []string, ) *PostApiSeasonsBySeasonIdActionsAddEventsRequest`

NewPostApiSeasonsBySeasonIdActionsAddEventsRequest instantiates a new PostApiSeasonsBySeasonIdActionsAddEventsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostApiSeasonsBySeasonIdActionsAddEventsRequestWithDefaults

`func NewPostApiSeasonsBySeasonIdActionsAddEventsRequestWithDefaults() *PostApiSeasonsBySeasonIdActionsAddEventsRequest`

NewPostApiSeasonsBySeasonIdActionsAddEventsRequestWithDefaults instantiates a new PostApiSeasonsBySeasonIdActionsAddEventsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEventIds

`func (o *PostApiSeasonsBySeasonIdActionsAddEventsRequest) GetEventIds() []string`

GetEventIds returns the EventIds field if non-nil, zero value otherwise.

### GetEventIdsOk

`func (o *PostApiSeasonsBySeasonIdActionsAddEventsRequest) GetEventIdsOk() (*[]string, bool)`

GetEventIdsOk returns a tuple with the EventIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventIds

`func (o *PostApiSeasonsBySeasonIdActionsAddEventsRequest) SetEventIds(v []string)`

SetEventIds sets EventIds field to given value.


### GetSync

`func (o *PostApiSeasonsBySeasonIdActionsAddEventsRequest) GetSync() bool`

GetSync returns the Sync field if non-nil, zero value otherwise.

### GetSyncOk

`func (o *PostApiSeasonsBySeasonIdActionsAddEventsRequest) GetSyncOk() (*bool, bool)`

GetSyncOk returns a tuple with the Sync field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSync

`func (o *PostApiSeasonsBySeasonIdActionsAddEventsRequest) SetSync(v bool)`

SetSync sets Sync field to given value.

### HasSync

`func (o *PostApiSeasonsBySeasonIdActionsAddEventsRequest) HasSync() bool`

HasSync returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


