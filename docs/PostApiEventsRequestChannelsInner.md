# PostApiEventsRequestChannelsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Color** | Pointer to **string** |  | [optional] 
**ChannelKey** | Pointer to **string** | Secret channel key; generated when omitted. | [optional] 
**Objects** | Pointer to **[]string** | Seat labels or UUIDs and GA area names or UUIDs of the new event. | [optional] 
**AreaPlaces** | Pointer to **map[string]int32** | GA area name or UUID → number of places. | [optional] 

## Methods

### NewPostApiEventsRequestChannelsInner

`func NewPostApiEventsRequestChannelsInner(name string, ) *PostApiEventsRequestChannelsInner`

NewPostApiEventsRequestChannelsInner instantiates a new PostApiEventsRequestChannelsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostApiEventsRequestChannelsInnerWithDefaults

`func NewPostApiEventsRequestChannelsInnerWithDefaults() *PostApiEventsRequestChannelsInner`

NewPostApiEventsRequestChannelsInnerWithDefaults instantiates a new PostApiEventsRequestChannelsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *PostApiEventsRequestChannelsInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PostApiEventsRequestChannelsInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PostApiEventsRequestChannelsInner) SetName(v string)`

SetName sets Name field to given value.


### GetColor

`func (o *PostApiEventsRequestChannelsInner) GetColor() string`

GetColor returns the Color field if non-nil, zero value otherwise.

### GetColorOk

`func (o *PostApiEventsRequestChannelsInner) GetColorOk() (*string, bool)`

GetColorOk returns a tuple with the Color field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColor

`func (o *PostApiEventsRequestChannelsInner) SetColor(v string)`

SetColor sets Color field to given value.

### HasColor

`func (o *PostApiEventsRequestChannelsInner) HasColor() bool`

HasColor returns a boolean if a field has been set.

### GetChannelKey

`func (o *PostApiEventsRequestChannelsInner) GetChannelKey() string`

GetChannelKey returns the ChannelKey field if non-nil, zero value otherwise.

### GetChannelKeyOk

`func (o *PostApiEventsRequestChannelsInner) GetChannelKeyOk() (*string, bool)`

GetChannelKeyOk returns a tuple with the ChannelKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelKey

`func (o *PostApiEventsRequestChannelsInner) SetChannelKey(v string)`

SetChannelKey sets ChannelKey field to given value.

### HasChannelKey

`func (o *PostApiEventsRequestChannelsInner) HasChannelKey() bool`

HasChannelKey returns a boolean if a field has been set.

### GetObjects

`func (o *PostApiEventsRequestChannelsInner) GetObjects() []string`

GetObjects returns the Objects field if non-nil, zero value otherwise.

### GetObjectsOk

`func (o *PostApiEventsRequestChannelsInner) GetObjectsOk() (*[]string, bool)`

GetObjectsOk returns a tuple with the Objects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjects

`func (o *PostApiEventsRequestChannelsInner) SetObjects(v []string)`

SetObjects sets Objects field to given value.

### HasObjects

`func (o *PostApiEventsRequestChannelsInner) HasObjects() bool`

HasObjects returns a boolean if a field has been set.

### GetAreaPlaces

`func (o *PostApiEventsRequestChannelsInner) GetAreaPlaces() map[string]int32`

GetAreaPlaces returns the AreaPlaces field if non-nil, zero value otherwise.

### GetAreaPlacesOk

`func (o *PostApiEventsRequestChannelsInner) GetAreaPlacesOk() (*map[string]int32, bool)`

GetAreaPlacesOk returns a tuple with the AreaPlaces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAreaPlaces

`func (o *PostApiEventsRequestChannelsInner) SetAreaPlaces(v map[string]int32)`

SetAreaPlaces sets AreaPlaces field to given value.

### HasAreaPlaces

`func (o *PostApiEventsRequestChannelsInner) HasAreaPlaces() bool`

HasAreaPlaces returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


