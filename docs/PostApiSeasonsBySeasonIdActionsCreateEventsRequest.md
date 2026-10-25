# PostApiSeasonsBySeasonIdActionsCreateEventsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EventKeys** | Pointer to **[]string** | One new event per key (named \&quot;{season} · {key}\&quot;). | [optional] 
**NumberOfEvents** | Pointer to **int32** | This many new events, numbered after the current members (\&quot;{season} · N\&quot;). | [optional] 
**Events** | Pointer to [**[]PostApiSeasonsBySeasonIdActionsCreateEventsRequestEventsInner**](PostApiSeasonsBySeasonIdActionsCreateEventsRequestEventsInner.md) |  | [optional] 

## Methods

### NewPostApiSeasonsBySeasonIdActionsCreateEventsRequest

`func NewPostApiSeasonsBySeasonIdActionsCreateEventsRequest() *PostApiSeasonsBySeasonIdActionsCreateEventsRequest`

NewPostApiSeasonsBySeasonIdActionsCreateEventsRequest instantiates a new PostApiSeasonsBySeasonIdActionsCreateEventsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostApiSeasonsBySeasonIdActionsCreateEventsRequestWithDefaults

`func NewPostApiSeasonsBySeasonIdActionsCreateEventsRequestWithDefaults() *PostApiSeasonsBySeasonIdActionsCreateEventsRequest`

NewPostApiSeasonsBySeasonIdActionsCreateEventsRequestWithDefaults instantiates a new PostApiSeasonsBySeasonIdActionsCreateEventsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEventKeys

`func (o *PostApiSeasonsBySeasonIdActionsCreateEventsRequest) GetEventKeys() []string`

GetEventKeys returns the EventKeys field if non-nil, zero value otherwise.

### GetEventKeysOk

`func (o *PostApiSeasonsBySeasonIdActionsCreateEventsRequest) GetEventKeysOk() (*[]string, bool)`

GetEventKeysOk returns a tuple with the EventKeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventKeys

`func (o *PostApiSeasonsBySeasonIdActionsCreateEventsRequest) SetEventKeys(v []string)`

SetEventKeys sets EventKeys field to given value.

### HasEventKeys

`func (o *PostApiSeasonsBySeasonIdActionsCreateEventsRequest) HasEventKeys() bool`

HasEventKeys returns a boolean if a field has been set.

### GetNumberOfEvents

`func (o *PostApiSeasonsBySeasonIdActionsCreateEventsRequest) GetNumberOfEvents() int32`

GetNumberOfEvents returns the NumberOfEvents field if non-nil, zero value otherwise.

### GetNumberOfEventsOk

`func (o *PostApiSeasonsBySeasonIdActionsCreateEventsRequest) GetNumberOfEventsOk() (*int32, bool)`

GetNumberOfEventsOk returns a tuple with the NumberOfEvents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberOfEvents

`func (o *PostApiSeasonsBySeasonIdActionsCreateEventsRequest) SetNumberOfEvents(v int32)`

SetNumberOfEvents sets NumberOfEvents field to given value.

### HasNumberOfEvents

`func (o *PostApiSeasonsBySeasonIdActionsCreateEventsRequest) HasNumberOfEvents() bool`

HasNumberOfEvents returns a boolean if a field has been set.

### GetEvents

`func (o *PostApiSeasonsBySeasonIdActionsCreateEventsRequest) GetEvents() []PostApiSeasonsBySeasonIdActionsCreateEventsRequestEventsInner`

GetEvents returns the Events field if non-nil, zero value otherwise.

### GetEventsOk

`func (o *PostApiSeasonsBySeasonIdActionsCreateEventsRequest) GetEventsOk() (*[]PostApiSeasonsBySeasonIdActionsCreateEventsRequestEventsInner, bool)`

GetEventsOk returns a tuple with the Events field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvents

`func (o *PostApiSeasonsBySeasonIdActionsCreateEventsRequest) SetEvents(v []PostApiSeasonsBySeasonIdActionsCreateEventsRequestEventsInner)`

SetEvents sets Events field to given value.

### HasEvents

`func (o *PostApiSeasonsBySeasonIdActionsCreateEventsRequest) HasEvents() bool`

HasEvents returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


