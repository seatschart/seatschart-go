# PostApiSandboxes201ResponseCopiedSeasonsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SourceSeasonId** | **string** |  | 
**SeasonId** | **string** |  | 
**Name** | **string** |  | 
**Key** | **NullableString** | The season key, or null when the sandbox already uses &#x60;sourceKey&#x60; (copied without a key). | 
**SourceKey** | **NullableString** |  | 
**Events** | **int32** | Member events copied into the season. | 
**PartialSeasons** | **int32** | Partial seasons copied with their copied member events. | 

## Methods

### NewPostApiSandboxes201ResponseCopiedSeasonsInner

`func NewPostApiSandboxes201ResponseCopiedSeasonsInner(sourceSeasonId string, seasonId string, name string, key NullableString, sourceKey NullableString, events int32, partialSeasons int32, ) *PostApiSandboxes201ResponseCopiedSeasonsInner`

NewPostApiSandboxes201ResponseCopiedSeasonsInner instantiates a new PostApiSandboxes201ResponseCopiedSeasonsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostApiSandboxes201ResponseCopiedSeasonsInnerWithDefaults

`func NewPostApiSandboxes201ResponseCopiedSeasonsInnerWithDefaults() *PostApiSandboxes201ResponseCopiedSeasonsInner`

NewPostApiSandboxes201ResponseCopiedSeasonsInnerWithDefaults instantiates a new PostApiSandboxes201ResponseCopiedSeasonsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSourceSeasonId

`func (o *PostApiSandboxes201ResponseCopiedSeasonsInner) GetSourceSeasonId() string`

GetSourceSeasonId returns the SourceSeasonId field if non-nil, zero value otherwise.

### GetSourceSeasonIdOk

`func (o *PostApiSandboxes201ResponseCopiedSeasonsInner) GetSourceSeasonIdOk() (*string, bool)`

GetSourceSeasonIdOk returns a tuple with the SourceSeasonId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceSeasonId

`func (o *PostApiSandboxes201ResponseCopiedSeasonsInner) SetSourceSeasonId(v string)`

SetSourceSeasonId sets SourceSeasonId field to given value.


### GetSeasonId

`func (o *PostApiSandboxes201ResponseCopiedSeasonsInner) GetSeasonId() string`

GetSeasonId returns the SeasonId field if non-nil, zero value otherwise.

### GetSeasonIdOk

`func (o *PostApiSandboxes201ResponseCopiedSeasonsInner) GetSeasonIdOk() (*string, bool)`

GetSeasonIdOk returns a tuple with the SeasonId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeasonId

`func (o *PostApiSandboxes201ResponseCopiedSeasonsInner) SetSeasonId(v string)`

SetSeasonId sets SeasonId field to given value.


### GetName

`func (o *PostApiSandboxes201ResponseCopiedSeasonsInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PostApiSandboxes201ResponseCopiedSeasonsInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PostApiSandboxes201ResponseCopiedSeasonsInner) SetName(v string)`

SetName sets Name field to given value.


### GetKey

`func (o *PostApiSandboxes201ResponseCopiedSeasonsInner) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *PostApiSandboxes201ResponseCopiedSeasonsInner) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *PostApiSandboxes201ResponseCopiedSeasonsInner) SetKey(v string)`

SetKey sets Key field to given value.


### SetKeyNil

`func (o *PostApiSandboxes201ResponseCopiedSeasonsInner) SetKeyNil(b bool)`

 SetKeyNil sets the value for Key to be an explicit nil

### UnsetKey
`func (o *PostApiSandboxes201ResponseCopiedSeasonsInner) UnsetKey()`

UnsetKey ensures that no value is present for Key, not even an explicit nil
### GetSourceKey

`func (o *PostApiSandboxes201ResponseCopiedSeasonsInner) GetSourceKey() string`

GetSourceKey returns the SourceKey field if non-nil, zero value otherwise.

### GetSourceKeyOk

`func (o *PostApiSandboxes201ResponseCopiedSeasonsInner) GetSourceKeyOk() (*string, bool)`

GetSourceKeyOk returns a tuple with the SourceKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceKey

`func (o *PostApiSandboxes201ResponseCopiedSeasonsInner) SetSourceKey(v string)`

SetSourceKey sets SourceKey field to given value.


### SetSourceKeyNil

`func (o *PostApiSandboxes201ResponseCopiedSeasonsInner) SetSourceKeyNil(b bool)`

 SetSourceKeyNil sets the value for SourceKey to be an explicit nil

### UnsetSourceKey
`func (o *PostApiSandboxes201ResponseCopiedSeasonsInner) UnsetSourceKey()`

UnsetSourceKey ensures that no value is present for SourceKey, not even an explicit nil
### GetEvents

`func (o *PostApiSandboxes201ResponseCopiedSeasonsInner) GetEvents() int32`

GetEvents returns the Events field if non-nil, zero value otherwise.

### GetEventsOk

`func (o *PostApiSandboxes201ResponseCopiedSeasonsInner) GetEventsOk() (*int32, bool)`

GetEventsOk returns a tuple with the Events field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvents

`func (o *PostApiSandboxes201ResponseCopiedSeasonsInner) SetEvents(v int32)`

SetEvents sets Events field to given value.


### GetPartialSeasons

`func (o *PostApiSandboxes201ResponseCopiedSeasonsInner) GetPartialSeasons() int32`

GetPartialSeasons returns the PartialSeasons field if non-nil, zero value otherwise.

### GetPartialSeasonsOk

`func (o *PostApiSandboxes201ResponseCopiedSeasonsInner) GetPartialSeasonsOk() (*int32, bool)`

GetPartialSeasonsOk returns a tuple with the PartialSeasons field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartialSeasons

`func (o *PostApiSandboxes201ResponseCopiedSeasonsInner) SetPartialSeasons(v int32)`

SetPartialSeasons sets PartialSeasons field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


