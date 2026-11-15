# PutApiEventsByEventIdInventoryChannelsDefinitionsRequestChannelsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Existing channel to keep (else matched by channelKey, else created). | [optional] 
**ChannelKey** | Pointer to **string** | Channel secret; a new channel without one gets a generated secret, returned once. | [optional] 
**Name** | **string** |  | 
**Color** | Pointer to **string** |  | [optional] 
**SeatIds** | Pointer to **[]string** |  | [optional] 
**GaAreaIds** | Pointer to **[]string** |  | [optional] 
**AreaPlaces** | Pointer to **map[string]int32** |  | [optional] 

## Methods

### NewPutApiEventsByEventIdInventoryChannelsDefinitionsRequestChannelsInner

`func NewPutApiEventsByEventIdInventoryChannelsDefinitionsRequestChannelsInner(name string, ) *PutApiEventsByEventIdInventoryChannelsDefinitionsRequestChannelsInner`

NewPutApiEventsByEventIdInventoryChannelsDefinitionsRequestChannelsInner instantiates a new PutApiEventsByEventIdInventoryChannelsDefinitionsRequestChannelsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPutApiEventsByEventIdInventoryChannelsDefinitionsRequestChannelsInnerWithDefaults

`func NewPutApiEventsByEventIdInventoryChannelsDefinitionsRequestChannelsInnerWithDefaults() *PutApiEventsByEventIdInventoryChannelsDefinitionsRequestChannelsInner`

NewPutApiEventsByEventIdInventoryChannelsDefinitionsRequestChannelsInnerWithDefaults instantiates a new PutApiEventsByEventIdInventoryChannelsDefinitionsRequestChannelsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PutApiEventsByEventIdInventoryChannelsDefinitionsRequestChannelsInner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PutApiEventsByEventIdInventoryChannelsDefinitionsRequestChannelsInner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PutApiEventsByEventIdInventoryChannelsDefinitionsRequestChannelsInner) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *PutApiEventsByEventIdInventoryChannelsDefinitionsRequestChannelsInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetChannelKey

`func (o *PutApiEventsByEventIdInventoryChannelsDefinitionsRequestChannelsInner) GetChannelKey() string`

GetChannelKey returns the ChannelKey field if non-nil, zero value otherwise.

### GetChannelKeyOk

`func (o *PutApiEventsByEventIdInventoryChannelsDefinitionsRequestChannelsInner) GetChannelKeyOk() (*string, bool)`

GetChannelKeyOk returns a tuple with the ChannelKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelKey

`func (o *PutApiEventsByEventIdInventoryChannelsDefinitionsRequestChannelsInner) SetChannelKey(v string)`

SetChannelKey sets ChannelKey field to given value.

### HasChannelKey

`func (o *PutApiEventsByEventIdInventoryChannelsDefinitionsRequestChannelsInner) HasChannelKey() bool`

HasChannelKey returns a boolean if a field has been set.

### GetName

`func (o *PutApiEventsByEventIdInventoryChannelsDefinitionsRequestChannelsInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PutApiEventsByEventIdInventoryChannelsDefinitionsRequestChannelsInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PutApiEventsByEventIdInventoryChannelsDefinitionsRequestChannelsInner) SetName(v string)`

SetName sets Name field to given value.


### GetColor

`func (o *PutApiEventsByEventIdInventoryChannelsDefinitionsRequestChannelsInner) GetColor() string`

GetColor returns the Color field if non-nil, zero value otherwise.

### GetColorOk

`func (o *PutApiEventsByEventIdInventoryChannelsDefinitionsRequestChannelsInner) GetColorOk() (*string, bool)`

GetColorOk returns a tuple with the Color field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColor

`func (o *PutApiEventsByEventIdInventoryChannelsDefinitionsRequestChannelsInner) SetColor(v string)`

SetColor sets Color field to given value.

### HasColor

`func (o *PutApiEventsByEventIdInventoryChannelsDefinitionsRequestChannelsInner) HasColor() bool`

HasColor returns a boolean if a field has been set.

### GetSeatIds

`func (o *PutApiEventsByEventIdInventoryChannelsDefinitionsRequestChannelsInner) GetSeatIds() []string`

GetSeatIds returns the SeatIds field if non-nil, zero value otherwise.

### GetSeatIdsOk

`func (o *PutApiEventsByEventIdInventoryChannelsDefinitionsRequestChannelsInner) GetSeatIdsOk() (*[]string, bool)`

GetSeatIdsOk returns a tuple with the SeatIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeatIds

`func (o *PutApiEventsByEventIdInventoryChannelsDefinitionsRequestChannelsInner) SetSeatIds(v []string)`

SetSeatIds sets SeatIds field to given value.

### HasSeatIds

`func (o *PutApiEventsByEventIdInventoryChannelsDefinitionsRequestChannelsInner) HasSeatIds() bool`

HasSeatIds returns a boolean if a field has been set.

### GetGaAreaIds

`func (o *PutApiEventsByEventIdInventoryChannelsDefinitionsRequestChannelsInner) GetGaAreaIds() []string`

GetGaAreaIds returns the GaAreaIds field if non-nil, zero value otherwise.

### GetGaAreaIdsOk

`func (o *PutApiEventsByEventIdInventoryChannelsDefinitionsRequestChannelsInner) GetGaAreaIdsOk() (*[]string, bool)`

GetGaAreaIdsOk returns a tuple with the GaAreaIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGaAreaIds

`func (o *PutApiEventsByEventIdInventoryChannelsDefinitionsRequestChannelsInner) SetGaAreaIds(v []string)`

SetGaAreaIds sets GaAreaIds field to given value.

### HasGaAreaIds

`func (o *PutApiEventsByEventIdInventoryChannelsDefinitionsRequestChannelsInner) HasGaAreaIds() bool`

HasGaAreaIds returns a boolean if a field has been set.

### GetAreaPlaces

`func (o *PutApiEventsByEventIdInventoryChannelsDefinitionsRequestChannelsInner) GetAreaPlaces() map[string]int32`

GetAreaPlaces returns the AreaPlaces field if non-nil, zero value otherwise.

### GetAreaPlacesOk

`func (o *PutApiEventsByEventIdInventoryChannelsDefinitionsRequestChannelsInner) GetAreaPlacesOk() (*map[string]int32, bool)`

GetAreaPlacesOk returns a tuple with the AreaPlaces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAreaPlaces

`func (o *PutApiEventsByEventIdInventoryChannelsDefinitionsRequestChannelsInner) SetAreaPlaces(v map[string]int32)`

SetAreaPlaces sets AreaPlaces field to given value.

### HasAreaPlaces

`func (o *PutApiEventsByEventIdInventoryChannelsDefinitionsRequestChannelsInner) HasAreaPlaces() bool`

HasAreaPlaces returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


