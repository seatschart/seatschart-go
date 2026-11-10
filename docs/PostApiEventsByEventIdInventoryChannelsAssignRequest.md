# PostApiEventsByEventIdInventoryChannelsAssignRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ChannelKeyId** | **NullableString** |  | 
**SeatIds** | Pointer to **[]string** |  | [optional] 
**GaAreaIds** | Pointer to **[]string** |  | [optional] 
**AreaPlaces** | Pointer to **map[string]int32** |  | [optional] 
**CascadeSeason** | Pointer to **bool** |  | [optional] 

## Methods

### NewPostApiEventsByEventIdInventoryChannelsAssignRequest

`func NewPostApiEventsByEventIdInventoryChannelsAssignRequest(channelKeyId NullableString, ) *PostApiEventsByEventIdInventoryChannelsAssignRequest`

NewPostApiEventsByEventIdInventoryChannelsAssignRequest instantiates a new PostApiEventsByEventIdInventoryChannelsAssignRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostApiEventsByEventIdInventoryChannelsAssignRequestWithDefaults

`func NewPostApiEventsByEventIdInventoryChannelsAssignRequestWithDefaults() *PostApiEventsByEventIdInventoryChannelsAssignRequest`

NewPostApiEventsByEventIdInventoryChannelsAssignRequestWithDefaults instantiates a new PostApiEventsByEventIdInventoryChannelsAssignRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChannelKeyId

`func (o *PostApiEventsByEventIdInventoryChannelsAssignRequest) GetChannelKeyId() string`

GetChannelKeyId returns the ChannelKeyId field if non-nil, zero value otherwise.

### GetChannelKeyIdOk

`func (o *PostApiEventsByEventIdInventoryChannelsAssignRequest) GetChannelKeyIdOk() (*string, bool)`

GetChannelKeyIdOk returns a tuple with the ChannelKeyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelKeyId

`func (o *PostApiEventsByEventIdInventoryChannelsAssignRequest) SetChannelKeyId(v string)`

SetChannelKeyId sets ChannelKeyId field to given value.


### SetChannelKeyIdNil

`func (o *PostApiEventsByEventIdInventoryChannelsAssignRequest) SetChannelKeyIdNil(b bool)`

 SetChannelKeyIdNil sets the value for ChannelKeyId to be an explicit nil

### UnsetChannelKeyId
`func (o *PostApiEventsByEventIdInventoryChannelsAssignRequest) UnsetChannelKeyId()`

UnsetChannelKeyId ensures that no value is present for ChannelKeyId, not even an explicit nil
### GetSeatIds

`func (o *PostApiEventsByEventIdInventoryChannelsAssignRequest) GetSeatIds() []string`

GetSeatIds returns the SeatIds field if non-nil, zero value otherwise.

### GetSeatIdsOk

`func (o *PostApiEventsByEventIdInventoryChannelsAssignRequest) GetSeatIdsOk() (*[]string, bool)`

GetSeatIdsOk returns a tuple with the SeatIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeatIds

`func (o *PostApiEventsByEventIdInventoryChannelsAssignRequest) SetSeatIds(v []string)`

SetSeatIds sets SeatIds field to given value.

### HasSeatIds

`func (o *PostApiEventsByEventIdInventoryChannelsAssignRequest) HasSeatIds() bool`

HasSeatIds returns a boolean if a field has been set.

### GetGaAreaIds

`func (o *PostApiEventsByEventIdInventoryChannelsAssignRequest) GetGaAreaIds() []string`

GetGaAreaIds returns the GaAreaIds field if non-nil, zero value otherwise.

### GetGaAreaIdsOk

`func (o *PostApiEventsByEventIdInventoryChannelsAssignRequest) GetGaAreaIdsOk() (*[]string, bool)`

GetGaAreaIdsOk returns a tuple with the GaAreaIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGaAreaIds

`func (o *PostApiEventsByEventIdInventoryChannelsAssignRequest) SetGaAreaIds(v []string)`

SetGaAreaIds sets GaAreaIds field to given value.

### HasGaAreaIds

`func (o *PostApiEventsByEventIdInventoryChannelsAssignRequest) HasGaAreaIds() bool`

HasGaAreaIds returns a boolean if a field has been set.

### GetAreaPlaces

`func (o *PostApiEventsByEventIdInventoryChannelsAssignRequest) GetAreaPlaces() map[string]int32`

GetAreaPlaces returns the AreaPlaces field if non-nil, zero value otherwise.

### GetAreaPlacesOk

`func (o *PostApiEventsByEventIdInventoryChannelsAssignRequest) GetAreaPlacesOk() (*map[string]int32, bool)`

GetAreaPlacesOk returns a tuple with the AreaPlaces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAreaPlaces

`func (o *PostApiEventsByEventIdInventoryChannelsAssignRequest) SetAreaPlaces(v map[string]int32)`

SetAreaPlaces sets AreaPlaces field to given value.

### HasAreaPlaces

`func (o *PostApiEventsByEventIdInventoryChannelsAssignRequest) HasAreaPlaces() bool`

HasAreaPlaces returns a boolean if a field has been set.

### GetCascadeSeason

`func (o *PostApiEventsByEventIdInventoryChannelsAssignRequest) GetCascadeSeason() bool`

GetCascadeSeason returns the CascadeSeason field if non-nil, zero value otherwise.

### GetCascadeSeasonOk

`func (o *PostApiEventsByEventIdInventoryChannelsAssignRequest) GetCascadeSeasonOk() (*bool, bool)`

GetCascadeSeasonOk returns a tuple with the CascadeSeason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCascadeSeason

`func (o *PostApiEventsByEventIdInventoryChannelsAssignRequest) SetCascadeSeason(v bool)`

SetCascadeSeason sets CascadeSeason field to given value.

### HasCascadeSeason

`func (o *PostApiEventsByEventIdInventoryChannelsAssignRequest) HasCascadeSeason() bool`

HasCascadeSeason returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


