# PostApiSeasonsBySeasonIdEventsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EventId** | **string** |  | 
**SetAsSource** | Pointer to **bool** |  | [optional] 
**Sync** | Pointer to **bool** |  | [optional] 

## Methods

### NewPostApiSeasonsBySeasonIdEventsRequest

`func NewPostApiSeasonsBySeasonIdEventsRequest(eventId string, ) *PostApiSeasonsBySeasonIdEventsRequest`

NewPostApiSeasonsBySeasonIdEventsRequest instantiates a new PostApiSeasonsBySeasonIdEventsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostApiSeasonsBySeasonIdEventsRequestWithDefaults

`func NewPostApiSeasonsBySeasonIdEventsRequestWithDefaults() *PostApiSeasonsBySeasonIdEventsRequest`

NewPostApiSeasonsBySeasonIdEventsRequestWithDefaults instantiates a new PostApiSeasonsBySeasonIdEventsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEventId

`func (o *PostApiSeasonsBySeasonIdEventsRequest) GetEventId() string`

GetEventId returns the EventId field if non-nil, zero value otherwise.

### GetEventIdOk

`func (o *PostApiSeasonsBySeasonIdEventsRequest) GetEventIdOk() (*string, bool)`

GetEventIdOk returns a tuple with the EventId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventId

`func (o *PostApiSeasonsBySeasonIdEventsRequest) SetEventId(v string)`

SetEventId sets EventId field to given value.


### GetSetAsSource

`func (o *PostApiSeasonsBySeasonIdEventsRequest) GetSetAsSource() bool`

GetSetAsSource returns the SetAsSource field if non-nil, zero value otherwise.

### GetSetAsSourceOk

`func (o *PostApiSeasonsBySeasonIdEventsRequest) GetSetAsSourceOk() (*bool, bool)`

GetSetAsSourceOk returns a tuple with the SetAsSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSetAsSource

`func (o *PostApiSeasonsBySeasonIdEventsRequest) SetSetAsSource(v bool)`

SetSetAsSource sets SetAsSource field to given value.

### HasSetAsSource

`func (o *PostApiSeasonsBySeasonIdEventsRequest) HasSetAsSource() bool`

HasSetAsSource returns a boolean if a field has been set.

### GetSync

`func (o *PostApiSeasonsBySeasonIdEventsRequest) GetSync() bool`

GetSync returns the Sync field if non-nil, zero value otherwise.

### GetSyncOk

`func (o *PostApiSeasonsBySeasonIdEventsRequest) GetSyncOk() (*bool, bool)`

GetSyncOk returns a tuple with the Sync field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSync

`func (o *PostApiSeasonsBySeasonIdEventsRequest) SetSync(v bool)`

SetSync sets Sync field to given value.

### HasSync

`func (o *PostApiSeasonsBySeasonIdEventsRequest) HasSync() bool`

HasSync returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


