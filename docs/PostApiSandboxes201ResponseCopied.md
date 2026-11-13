# PostApiSandboxes201ResponseCopied

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Charts** | [**[]PostApiSandboxes201ResponseCopiedChartsInner**](PostApiSandboxes201ResponseCopiedChartsInner.md) |  | 
**Events** | [**[]PostApiSandboxes201ResponseCopiedEventsInner**](PostApiSandboxes201ResponseCopiedEventsInner.md) |  | 
**KeyConflicts** | [**[]PostApiSandboxes201ResponseCopiedKeyConflictsInner**](PostApiSandboxes201ResponseCopiedKeyConflictsInner.md) |  | 
**TruncatedEvents** | **int32** |  | 
**EventsWithOwnMap** | [**[]PostApiSandboxes201ResponseCopiedEventsWithOwnMapInner**](PostApiSandboxes201ResponseCopiedEventsWithOwnMapInner.md) | Events with their own map (no chart) copied with their event layout: only when every chart is copied (no &#x60;chartIds&#x60;) with &#x60;includeEvents&#x60;. | 
**Seasons** | [**[]PostApiSandboxes201ResponseCopiedSeasonsInner**](PostApiSandboxes201ResponseCopiedSeasonsInner.md) | Seasons (with partial seasons) whose member events were copied (&#x60;includeEvents&#x60;), with fresh ids and remapped members. | 

## Methods

### NewPostApiSandboxes201ResponseCopied

`func NewPostApiSandboxes201ResponseCopied(charts []PostApiSandboxes201ResponseCopiedChartsInner, events []PostApiSandboxes201ResponseCopiedEventsInner, keyConflicts []PostApiSandboxes201ResponseCopiedKeyConflictsInner, truncatedEvents int32, eventsWithOwnMap []PostApiSandboxes201ResponseCopiedEventsWithOwnMapInner, seasons []PostApiSandboxes201ResponseCopiedSeasonsInner, ) *PostApiSandboxes201ResponseCopied`

NewPostApiSandboxes201ResponseCopied instantiates a new PostApiSandboxes201ResponseCopied object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostApiSandboxes201ResponseCopiedWithDefaults

`func NewPostApiSandboxes201ResponseCopiedWithDefaults() *PostApiSandboxes201ResponseCopied`

NewPostApiSandboxes201ResponseCopiedWithDefaults instantiates a new PostApiSandboxes201ResponseCopied object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCharts

`func (o *PostApiSandboxes201ResponseCopied) GetCharts() []PostApiSandboxes201ResponseCopiedChartsInner`

GetCharts returns the Charts field if non-nil, zero value otherwise.

### GetChartsOk

`func (o *PostApiSandboxes201ResponseCopied) GetChartsOk() (*[]PostApiSandboxes201ResponseCopiedChartsInner, bool)`

GetChartsOk returns a tuple with the Charts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCharts

`func (o *PostApiSandboxes201ResponseCopied) SetCharts(v []PostApiSandboxes201ResponseCopiedChartsInner)`

SetCharts sets Charts field to given value.


### GetEvents

`func (o *PostApiSandboxes201ResponseCopied) GetEvents() []PostApiSandboxes201ResponseCopiedEventsInner`

GetEvents returns the Events field if non-nil, zero value otherwise.

### GetEventsOk

`func (o *PostApiSandboxes201ResponseCopied) GetEventsOk() (*[]PostApiSandboxes201ResponseCopiedEventsInner, bool)`

GetEventsOk returns a tuple with the Events field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvents

`func (o *PostApiSandboxes201ResponseCopied) SetEvents(v []PostApiSandboxes201ResponseCopiedEventsInner)`

SetEvents sets Events field to given value.


### GetKeyConflicts

`func (o *PostApiSandboxes201ResponseCopied) GetKeyConflicts() []PostApiSandboxes201ResponseCopiedKeyConflictsInner`

GetKeyConflicts returns the KeyConflicts field if non-nil, zero value otherwise.

### GetKeyConflictsOk

`func (o *PostApiSandboxes201ResponseCopied) GetKeyConflictsOk() (*[]PostApiSandboxes201ResponseCopiedKeyConflictsInner, bool)`

GetKeyConflictsOk returns a tuple with the KeyConflicts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyConflicts

`func (o *PostApiSandboxes201ResponseCopied) SetKeyConflicts(v []PostApiSandboxes201ResponseCopiedKeyConflictsInner)`

SetKeyConflicts sets KeyConflicts field to given value.


### GetTruncatedEvents

`func (o *PostApiSandboxes201ResponseCopied) GetTruncatedEvents() int32`

GetTruncatedEvents returns the TruncatedEvents field if non-nil, zero value otherwise.

### GetTruncatedEventsOk

`func (o *PostApiSandboxes201ResponseCopied) GetTruncatedEventsOk() (*int32, bool)`

GetTruncatedEventsOk returns a tuple with the TruncatedEvents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTruncatedEvents

`func (o *PostApiSandboxes201ResponseCopied) SetTruncatedEvents(v int32)`

SetTruncatedEvents sets TruncatedEvents field to given value.


### GetEventsWithOwnMap

`func (o *PostApiSandboxes201ResponseCopied) GetEventsWithOwnMap() []PostApiSandboxes201ResponseCopiedEventsWithOwnMapInner`

GetEventsWithOwnMap returns the EventsWithOwnMap field if non-nil, zero value otherwise.

### GetEventsWithOwnMapOk

`func (o *PostApiSandboxes201ResponseCopied) GetEventsWithOwnMapOk() (*[]PostApiSandboxes201ResponseCopiedEventsWithOwnMapInner, bool)`

GetEventsWithOwnMapOk returns a tuple with the EventsWithOwnMap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventsWithOwnMap

`func (o *PostApiSandboxes201ResponseCopied) SetEventsWithOwnMap(v []PostApiSandboxes201ResponseCopiedEventsWithOwnMapInner)`

SetEventsWithOwnMap sets EventsWithOwnMap field to given value.


### GetSeasons

`func (o *PostApiSandboxes201ResponseCopied) GetSeasons() []PostApiSandboxes201ResponseCopiedSeasonsInner`

GetSeasons returns the Seasons field if non-nil, zero value otherwise.

### GetSeasonsOk

`func (o *PostApiSandboxes201ResponseCopied) GetSeasonsOk() (*[]PostApiSandboxes201ResponseCopiedSeasonsInner, bool)`

GetSeasonsOk returns a tuple with the Seasons field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeasons

`func (o *PostApiSandboxes201ResponseCopied) SetSeasons(v []PostApiSandboxes201ResponseCopiedSeasonsInner)`

SetSeasons sets Seasons field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


