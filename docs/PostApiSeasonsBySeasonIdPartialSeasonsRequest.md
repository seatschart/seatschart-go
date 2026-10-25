# PostApiSeasonsBySeasonIdPartialSeasonsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Key** | Pointer to **NullableString** | Partner &#x60;key&#x60;: unique within the season, 1–100 URL-safe characters, never UUID-shaped; usable in place of the partial season id in every partial season path (&#x60;/seasons/{seasonIdOrKey}/partial-seasons/{partialSeasonIdOrKey}&#x60;), and as body &#x60;seasonId&#x60; of booking and status changes when no season and no other partial season of the workspace uses it. | [optional] 
**EventIds** | Pointer to **[]string** |  | [optional] 

## Methods

### NewPostApiSeasonsBySeasonIdPartialSeasonsRequest

`func NewPostApiSeasonsBySeasonIdPartialSeasonsRequest(name string, ) *PostApiSeasonsBySeasonIdPartialSeasonsRequest`

NewPostApiSeasonsBySeasonIdPartialSeasonsRequest instantiates a new PostApiSeasonsBySeasonIdPartialSeasonsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostApiSeasonsBySeasonIdPartialSeasonsRequestWithDefaults

`func NewPostApiSeasonsBySeasonIdPartialSeasonsRequestWithDefaults() *PostApiSeasonsBySeasonIdPartialSeasonsRequest`

NewPostApiSeasonsBySeasonIdPartialSeasonsRequestWithDefaults instantiates a new PostApiSeasonsBySeasonIdPartialSeasonsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *PostApiSeasonsBySeasonIdPartialSeasonsRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PostApiSeasonsBySeasonIdPartialSeasonsRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PostApiSeasonsBySeasonIdPartialSeasonsRequest) SetName(v string)`

SetName sets Name field to given value.


### GetKey

`func (o *PostApiSeasonsBySeasonIdPartialSeasonsRequest) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *PostApiSeasonsBySeasonIdPartialSeasonsRequest) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *PostApiSeasonsBySeasonIdPartialSeasonsRequest) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *PostApiSeasonsBySeasonIdPartialSeasonsRequest) HasKey() bool`

HasKey returns a boolean if a field has been set.

### SetKeyNil

`func (o *PostApiSeasonsBySeasonIdPartialSeasonsRequest) SetKeyNil(b bool)`

 SetKeyNil sets the value for Key to be an explicit nil

### UnsetKey
`func (o *PostApiSeasonsBySeasonIdPartialSeasonsRequest) UnsetKey()`

UnsetKey ensures that no value is present for Key, not even an explicit nil
### GetEventIds

`func (o *PostApiSeasonsBySeasonIdPartialSeasonsRequest) GetEventIds() []string`

GetEventIds returns the EventIds field if non-nil, zero value otherwise.

### GetEventIdsOk

`func (o *PostApiSeasonsBySeasonIdPartialSeasonsRequest) GetEventIdsOk() (*[]string, bool)`

GetEventIdsOk returns a tuple with the EventIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventIds

`func (o *PostApiSeasonsBySeasonIdPartialSeasonsRequest) SetEventIds(v []string)`

SetEventIds sets EventIds field to given value.

### HasEventIds

`func (o *PostApiSeasonsBySeasonIdPartialSeasonsRequest) HasEventIds() bool`

HasEventIds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


