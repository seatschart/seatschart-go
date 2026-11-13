# PostApiSandboxesBySandboxIdActionsCopyFromProduction200Response

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

### NewPostApiSandboxesBySandboxIdActionsCopyFromProduction200Response

`func NewPostApiSandboxesBySandboxIdActionsCopyFromProduction200Response(charts []PostApiSandboxes201ResponseCopiedChartsInner, events []PostApiSandboxes201ResponseCopiedEventsInner, keyConflicts []PostApiSandboxes201ResponseCopiedKeyConflictsInner, truncatedEvents int32, eventsWithOwnMap []PostApiSandboxes201ResponseCopiedEventsWithOwnMapInner, seasons []PostApiSandboxes201ResponseCopiedSeasonsInner, ) *PostApiSandboxesBySandboxIdActionsCopyFromProduction200Response`

NewPostApiSandboxesBySandboxIdActionsCopyFromProduction200Response instantiates a new PostApiSandboxesBySandboxIdActionsCopyFromProduction200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostApiSandboxesBySandboxIdActionsCopyFromProduction200ResponseWithDefaults

`func NewPostApiSandboxesBySandboxIdActionsCopyFromProduction200ResponseWithDefaults() *PostApiSandboxesBySandboxIdActionsCopyFromProduction200Response`

NewPostApiSandboxesBySandboxIdActionsCopyFromProduction200ResponseWithDefaults instantiates a new PostApiSandboxesBySandboxIdActionsCopyFromProduction200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCharts

`func (o *PostApiSandboxesBySandboxIdActionsCopyFromProduction200Response) GetCharts() []PostApiSandboxes201ResponseCopiedChartsInner`

GetCharts returns the Charts field if non-nil, zero value otherwise.

### GetChartsOk

`func (o *PostApiSandboxesBySandboxIdActionsCopyFromProduction200Response) GetChartsOk() (*[]PostApiSandboxes201ResponseCopiedChartsInner, bool)`

GetChartsOk returns a tuple with the Charts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCharts

`func (o *PostApiSandboxesBySandboxIdActionsCopyFromProduction200Response) SetCharts(v []PostApiSandboxes201ResponseCopiedChartsInner)`

SetCharts sets Charts field to given value.


### GetEvents

`func (o *PostApiSandboxesBySandboxIdActionsCopyFromProduction200Response) GetEvents() []PostApiSandboxes201ResponseCopiedEventsInner`

GetEvents returns the Events field if non-nil, zero value otherwise.

### GetEventsOk

`func (o *PostApiSandboxesBySandboxIdActionsCopyFromProduction200Response) GetEventsOk() (*[]PostApiSandboxes201ResponseCopiedEventsInner, bool)`

GetEventsOk returns a tuple with the Events field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvents

`func (o *PostApiSandboxesBySandboxIdActionsCopyFromProduction200Response) SetEvents(v []PostApiSandboxes201ResponseCopiedEventsInner)`

SetEvents sets Events field to given value.


### GetKeyConflicts

`func (o *PostApiSandboxesBySandboxIdActionsCopyFromProduction200Response) GetKeyConflicts() []PostApiSandboxes201ResponseCopiedKeyConflictsInner`

GetKeyConflicts returns the KeyConflicts field if non-nil, zero value otherwise.

### GetKeyConflictsOk

`func (o *PostApiSandboxesBySandboxIdActionsCopyFromProduction200Response) GetKeyConflictsOk() (*[]PostApiSandboxes201ResponseCopiedKeyConflictsInner, bool)`

GetKeyConflictsOk returns a tuple with the KeyConflicts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyConflicts

`func (o *PostApiSandboxesBySandboxIdActionsCopyFromProduction200Response) SetKeyConflicts(v []PostApiSandboxes201ResponseCopiedKeyConflictsInner)`

SetKeyConflicts sets KeyConflicts field to given value.


### GetTruncatedEvents

`func (o *PostApiSandboxesBySandboxIdActionsCopyFromProduction200Response) GetTruncatedEvents() int32`

GetTruncatedEvents returns the TruncatedEvents field if non-nil, zero value otherwise.

### GetTruncatedEventsOk

`func (o *PostApiSandboxesBySandboxIdActionsCopyFromProduction200Response) GetTruncatedEventsOk() (*int32, bool)`

GetTruncatedEventsOk returns a tuple with the TruncatedEvents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTruncatedEvents

`func (o *PostApiSandboxesBySandboxIdActionsCopyFromProduction200Response) SetTruncatedEvents(v int32)`

SetTruncatedEvents sets TruncatedEvents field to given value.


### GetEventsWithOwnMap

`func (o *PostApiSandboxesBySandboxIdActionsCopyFromProduction200Response) GetEventsWithOwnMap() []PostApiSandboxes201ResponseCopiedEventsWithOwnMapInner`

GetEventsWithOwnMap returns the EventsWithOwnMap field if non-nil, zero value otherwise.

### GetEventsWithOwnMapOk

`func (o *PostApiSandboxesBySandboxIdActionsCopyFromProduction200Response) GetEventsWithOwnMapOk() (*[]PostApiSandboxes201ResponseCopiedEventsWithOwnMapInner, bool)`

GetEventsWithOwnMapOk returns a tuple with the EventsWithOwnMap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventsWithOwnMap

`func (o *PostApiSandboxesBySandboxIdActionsCopyFromProduction200Response) SetEventsWithOwnMap(v []PostApiSandboxes201ResponseCopiedEventsWithOwnMapInner)`

SetEventsWithOwnMap sets EventsWithOwnMap field to given value.


### GetSeasons

`func (o *PostApiSandboxesBySandboxIdActionsCopyFromProduction200Response) GetSeasons() []PostApiSandboxes201ResponseCopiedSeasonsInner`

GetSeasons returns the Seasons field if non-nil, zero value otherwise.

### GetSeasonsOk

`func (o *PostApiSandboxesBySandboxIdActionsCopyFromProduction200Response) GetSeasonsOk() (*[]PostApiSandboxes201ResponseCopiedSeasonsInner, bool)`

GetSeasonsOk returns a tuple with the Seasons field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeasons

`func (o *PostApiSandboxesBySandboxIdActionsCopyFromProduction200Response) SetSeasons(v []PostApiSandboxes201ResponseCopiedSeasonsInner)`

SetSeasons sets Seasons field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


