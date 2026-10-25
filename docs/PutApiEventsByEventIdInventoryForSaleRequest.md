# PutApiEventsByEventIdInventoryForSaleRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ForSale** | **bool** | &#x60;false&#x60;: the listed objects are not for sale, everything else is. &#x60;true&#x60;: only the listed objects are for sale. | 
**SeatIds** | Pointer to **[]string** |  | [optional] 
**GaAreaIds** | Pointer to **[]string** |  | [optional] 
**Categories** | Pointer to **[]string** |  | [optional] 
**AreaPlaces** | Pointer to **map[string]int32** | GA area id → places: withheld (&#x60;forSale: false&#x60;) or the only places on sale (&#x60;forSale: true&#x60;). | [optional] 
**CascadeSeason** | Pointer to **bool** |  | [optional] 

## Methods

### NewPutApiEventsByEventIdInventoryForSaleRequest

`func NewPutApiEventsByEventIdInventoryForSaleRequest(forSale bool, ) *PutApiEventsByEventIdInventoryForSaleRequest`

NewPutApiEventsByEventIdInventoryForSaleRequest instantiates a new PutApiEventsByEventIdInventoryForSaleRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPutApiEventsByEventIdInventoryForSaleRequestWithDefaults

`func NewPutApiEventsByEventIdInventoryForSaleRequestWithDefaults() *PutApiEventsByEventIdInventoryForSaleRequest`

NewPutApiEventsByEventIdInventoryForSaleRequestWithDefaults instantiates a new PutApiEventsByEventIdInventoryForSaleRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetForSale

`func (o *PutApiEventsByEventIdInventoryForSaleRequest) GetForSale() bool`

GetForSale returns the ForSale field if non-nil, zero value otherwise.

### GetForSaleOk

`func (o *PutApiEventsByEventIdInventoryForSaleRequest) GetForSaleOk() (*bool, bool)`

GetForSaleOk returns a tuple with the ForSale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForSale

`func (o *PutApiEventsByEventIdInventoryForSaleRequest) SetForSale(v bool)`

SetForSale sets ForSale field to given value.


### GetSeatIds

`func (o *PutApiEventsByEventIdInventoryForSaleRequest) GetSeatIds() []string`

GetSeatIds returns the SeatIds field if non-nil, zero value otherwise.

### GetSeatIdsOk

`func (o *PutApiEventsByEventIdInventoryForSaleRequest) GetSeatIdsOk() (*[]string, bool)`

GetSeatIdsOk returns a tuple with the SeatIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeatIds

`func (o *PutApiEventsByEventIdInventoryForSaleRequest) SetSeatIds(v []string)`

SetSeatIds sets SeatIds field to given value.

### HasSeatIds

`func (o *PutApiEventsByEventIdInventoryForSaleRequest) HasSeatIds() bool`

HasSeatIds returns a boolean if a field has been set.

### GetGaAreaIds

`func (o *PutApiEventsByEventIdInventoryForSaleRequest) GetGaAreaIds() []string`

GetGaAreaIds returns the GaAreaIds field if non-nil, zero value otherwise.

### GetGaAreaIdsOk

`func (o *PutApiEventsByEventIdInventoryForSaleRequest) GetGaAreaIdsOk() (*[]string, bool)`

GetGaAreaIdsOk returns a tuple with the GaAreaIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGaAreaIds

`func (o *PutApiEventsByEventIdInventoryForSaleRequest) SetGaAreaIds(v []string)`

SetGaAreaIds sets GaAreaIds field to given value.

### HasGaAreaIds

`func (o *PutApiEventsByEventIdInventoryForSaleRequest) HasGaAreaIds() bool`

HasGaAreaIds returns a boolean if a field has been set.

### GetCategories

`func (o *PutApiEventsByEventIdInventoryForSaleRequest) GetCategories() []string`

GetCategories returns the Categories field if non-nil, zero value otherwise.

### GetCategoriesOk

`func (o *PutApiEventsByEventIdInventoryForSaleRequest) GetCategoriesOk() (*[]string, bool)`

GetCategoriesOk returns a tuple with the Categories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategories

`func (o *PutApiEventsByEventIdInventoryForSaleRequest) SetCategories(v []string)`

SetCategories sets Categories field to given value.

### HasCategories

`func (o *PutApiEventsByEventIdInventoryForSaleRequest) HasCategories() bool`

HasCategories returns a boolean if a field has been set.

### GetAreaPlaces

`func (o *PutApiEventsByEventIdInventoryForSaleRequest) GetAreaPlaces() map[string]int32`

GetAreaPlaces returns the AreaPlaces field if non-nil, zero value otherwise.

### GetAreaPlacesOk

`func (o *PutApiEventsByEventIdInventoryForSaleRequest) GetAreaPlacesOk() (*map[string]int32, bool)`

GetAreaPlacesOk returns a tuple with the AreaPlaces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAreaPlaces

`func (o *PutApiEventsByEventIdInventoryForSaleRequest) SetAreaPlaces(v map[string]int32)`

SetAreaPlaces sets AreaPlaces field to given value.

### HasAreaPlaces

`func (o *PutApiEventsByEventIdInventoryForSaleRequest) HasAreaPlaces() bool`

HasAreaPlaces returns a boolean if a field has been set.

### GetCascadeSeason

`func (o *PutApiEventsByEventIdInventoryForSaleRequest) GetCascadeSeason() bool`

GetCascadeSeason returns the CascadeSeason field if non-nil, zero value otherwise.

### GetCascadeSeasonOk

`func (o *PutApiEventsByEventIdInventoryForSaleRequest) GetCascadeSeasonOk() (*bool, bool)`

GetCascadeSeasonOk returns a tuple with the CascadeSeason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCascadeSeason

`func (o *PutApiEventsByEventIdInventoryForSaleRequest) SetCascadeSeason(v bool)`

SetCascadeSeason sets CascadeSeason field to given value.

### HasCascadeSeason

`func (o *PutApiEventsByEventIdInventoryForSaleRequest) HasCascadeSeason() bool`

HasCascadeSeason returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


