# PostApiEventsByEventIdInventoryForSaleApplyRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ForSale** | Pointer to **bool** |  | [optional] 
**SeatIds** | Pointer to **[]string** |  | [optional] 
**GaAreaIds** | Pointer to **[]string** |  | [optional] 
**Categories** | Pointer to **[]string** |  | [optional] 
**AreaPlaces** | Pointer to **map[string]int32** |  | [optional] 
**Everything** | Pointer to **bool** |  | [optional] 
**CascadeSeason** | Pointer to **bool** |  | [optional] 

## Methods

### NewPostApiEventsByEventIdInventoryForSaleApplyRequest

`func NewPostApiEventsByEventIdInventoryForSaleApplyRequest() *PostApiEventsByEventIdInventoryForSaleApplyRequest`

NewPostApiEventsByEventIdInventoryForSaleApplyRequest instantiates a new PostApiEventsByEventIdInventoryForSaleApplyRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostApiEventsByEventIdInventoryForSaleApplyRequestWithDefaults

`func NewPostApiEventsByEventIdInventoryForSaleApplyRequestWithDefaults() *PostApiEventsByEventIdInventoryForSaleApplyRequest`

NewPostApiEventsByEventIdInventoryForSaleApplyRequestWithDefaults instantiates a new PostApiEventsByEventIdInventoryForSaleApplyRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetForSale

`func (o *PostApiEventsByEventIdInventoryForSaleApplyRequest) GetForSale() bool`

GetForSale returns the ForSale field if non-nil, zero value otherwise.

### GetForSaleOk

`func (o *PostApiEventsByEventIdInventoryForSaleApplyRequest) GetForSaleOk() (*bool, bool)`

GetForSaleOk returns a tuple with the ForSale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForSale

`func (o *PostApiEventsByEventIdInventoryForSaleApplyRequest) SetForSale(v bool)`

SetForSale sets ForSale field to given value.

### HasForSale

`func (o *PostApiEventsByEventIdInventoryForSaleApplyRequest) HasForSale() bool`

HasForSale returns a boolean if a field has been set.

### GetSeatIds

`func (o *PostApiEventsByEventIdInventoryForSaleApplyRequest) GetSeatIds() []string`

GetSeatIds returns the SeatIds field if non-nil, zero value otherwise.

### GetSeatIdsOk

`func (o *PostApiEventsByEventIdInventoryForSaleApplyRequest) GetSeatIdsOk() (*[]string, bool)`

GetSeatIdsOk returns a tuple with the SeatIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeatIds

`func (o *PostApiEventsByEventIdInventoryForSaleApplyRequest) SetSeatIds(v []string)`

SetSeatIds sets SeatIds field to given value.

### HasSeatIds

`func (o *PostApiEventsByEventIdInventoryForSaleApplyRequest) HasSeatIds() bool`

HasSeatIds returns a boolean if a field has been set.

### GetGaAreaIds

`func (o *PostApiEventsByEventIdInventoryForSaleApplyRequest) GetGaAreaIds() []string`

GetGaAreaIds returns the GaAreaIds field if non-nil, zero value otherwise.

### GetGaAreaIdsOk

`func (o *PostApiEventsByEventIdInventoryForSaleApplyRequest) GetGaAreaIdsOk() (*[]string, bool)`

GetGaAreaIdsOk returns a tuple with the GaAreaIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGaAreaIds

`func (o *PostApiEventsByEventIdInventoryForSaleApplyRequest) SetGaAreaIds(v []string)`

SetGaAreaIds sets GaAreaIds field to given value.

### HasGaAreaIds

`func (o *PostApiEventsByEventIdInventoryForSaleApplyRequest) HasGaAreaIds() bool`

HasGaAreaIds returns a boolean if a field has been set.

### GetCategories

`func (o *PostApiEventsByEventIdInventoryForSaleApplyRequest) GetCategories() []string`

GetCategories returns the Categories field if non-nil, zero value otherwise.

### GetCategoriesOk

`func (o *PostApiEventsByEventIdInventoryForSaleApplyRequest) GetCategoriesOk() (*[]string, bool)`

GetCategoriesOk returns a tuple with the Categories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategories

`func (o *PostApiEventsByEventIdInventoryForSaleApplyRequest) SetCategories(v []string)`

SetCategories sets Categories field to given value.

### HasCategories

`func (o *PostApiEventsByEventIdInventoryForSaleApplyRequest) HasCategories() bool`

HasCategories returns a boolean if a field has been set.

### GetAreaPlaces

`func (o *PostApiEventsByEventIdInventoryForSaleApplyRequest) GetAreaPlaces() map[string]int32`

GetAreaPlaces returns the AreaPlaces field if non-nil, zero value otherwise.

### GetAreaPlacesOk

`func (o *PostApiEventsByEventIdInventoryForSaleApplyRequest) GetAreaPlacesOk() (*map[string]int32, bool)`

GetAreaPlacesOk returns a tuple with the AreaPlaces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAreaPlaces

`func (o *PostApiEventsByEventIdInventoryForSaleApplyRequest) SetAreaPlaces(v map[string]int32)`

SetAreaPlaces sets AreaPlaces field to given value.

### HasAreaPlaces

`func (o *PostApiEventsByEventIdInventoryForSaleApplyRequest) HasAreaPlaces() bool`

HasAreaPlaces returns a boolean if a field has been set.

### GetEverything

`func (o *PostApiEventsByEventIdInventoryForSaleApplyRequest) GetEverything() bool`

GetEverything returns the Everything field if non-nil, zero value otherwise.

### GetEverythingOk

`func (o *PostApiEventsByEventIdInventoryForSaleApplyRequest) GetEverythingOk() (*bool, bool)`

GetEverythingOk returns a tuple with the Everything field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEverything

`func (o *PostApiEventsByEventIdInventoryForSaleApplyRequest) SetEverything(v bool)`

SetEverything sets Everything field to given value.

### HasEverything

`func (o *PostApiEventsByEventIdInventoryForSaleApplyRequest) HasEverything() bool`

HasEverything returns a boolean if a field has been set.

### GetCascadeSeason

`func (o *PostApiEventsByEventIdInventoryForSaleApplyRequest) GetCascadeSeason() bool`

GetCascadeSeason returns the CascadeSeason field if non-nil, zero value otherwise.

### GetCascadeSeasonOk

`func (o *PostApiEventsByEventIdInventoryForSaleApplyRequest) GetCascadeSeasonOk() (*bool, bool)`

GetCascadeSeasonOk returns a tuple with the CascadeSeason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCascadeSeason

`func (o *PostApiEventsByEventIdInventoryForSaleApplyRequest) SetCascadeSeason(v bool)`

SetCascadeSeason sets CascadeSeason field to given value.

### HasCascadeSeason

`func (o *PostApiEventsByEventIdInventoryForSaleApplyRequest) HasCascadeSeason() bool`

HasCascadeSeason returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


