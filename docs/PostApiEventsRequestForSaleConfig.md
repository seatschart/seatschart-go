# PostApiEventsRequestForSaleConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ForSale** | **bool** | &#x60;false&#x60;: the listed objects/categories are not for sale and &#x60;areaPlaces&#x60; are withheld places. &#x60;true&#x60;: only the listed objects/categories are for sale and &#x60;areaPlaces&#x60; are the places left for sale. | 
**Objects** | Pointer to **[]string** | Seat labels or UUIDs and GA area names or UUIDs of the new event. | [optional] 
**Categories** | Pointer to **[]string** |  | [optional] 
**AreaPlaces** | Pointer to **map[string]int32** | GA area name or UUID → number of places. | [optional] 

## Methods

### NewPostApiEventsRequestForSaleConfig

`func NewPostApiEventsRequestForSaleConfig(forSale bool, ) *PostApiEventsRequestForSaleConfig`

NewPostApiEventsRequestForSaleConfig instantiates a new PostApiEventsRequestForSaleConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostApiEventsRequestForSaleConfigWithDefaults

`func NewPostApiEventsRequestForSaleConfigWithDefaults() *PostApiEventsRequestForSaleConfig`

NewPostApiEventsRequestForSaleConfigWithDefaults instantiates a new PostApiEventsRequestForSaleConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetForSale

`func (o *PostApiEventsRequestForSaleConfig) GetForSale() bool`

GetForSale returns the ForSale field if non-nil, zero value otherwise.

### GetForSaleOk

`func (o *PostApiEventsRequestForSaleConfig) GetForSaleOk() (*bool, bool)`

GetForSaleOk returns a tuple with the ForSale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForSale

`func (o *PostApiEventsRequestForSaleConfig) SetForSale(v bool)`

SetForSale sets ForSale field to given value.


### GetObjects

`func (o *PostApiEventsRequestForSaleConfig) GetObjects() []string`

GetObjects returns the Objects field if non-nil, zero value otherwise.

### GetObjectsOk

`func (o *PostApiEventsRequestForSaleConfig) GetObjectsOk() (*[]string, bool)`

GetObjectsOk returns a tuple with the Objects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjects

`func (o *PostApiEventsRequestForSaleConfig) SetObjects(v []string)`

SetObjects sets Objects field to given value.

### HasObjects

`func (o *PostApiEventsRequestForSaleConfig) HasObjects() bool`

HasObjects returns a boolean if a field has been set.

### GetCategories

`func (o *PostApiEventsRequestForSaleConfig) GetCategories() []string`

GetCategories returns the Categories field if non-nil, zero value otherwise.

### GetCategoriesOk

`func (o *PostApiEventsRequestForSaleConfig) GetCategoriesOk() (*[]string, bool)`

GetCategoriesOk returns a tuple with the Categories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategories

`func (o *PostApiEventsRequestForSaleConfig) SetCategories(v []string)`

SetCategories sets Categories field to given value.

### HasCategories

`func (o *PostApiEventsRequestForSaleConfig) HasCategories() bool`

HasCategories returns a boolean if a field has been set.

### GetAreaPlaces

`func (o *PostApiEventsRequestForSaleConfig) GetAreaPlaces() map[string]int32`

GetAreaPlaces returns the AreaPlaces field if non-nil, zero value otherwise.

### GetAreaPlacesOk

`func (o *PostApiEventsRequestForSaleConfig) GetAreaPlacesOk() (*map[string]int32, bool)`

GetAreaPlacesOk returns a tuple with the AreaPlaces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAreaPlaces

`func (o *PostApiEventsRequestForSaleConfig) SetAreaPlaces(v map[string]int32)`

SetAreaPlaces sets AreaPlaces field to given value.

### HasAreaPlaces

`func (o *PostApiEventsRequestForSaleConfig) HasAreaPlaces() bool`

HasAreaPlaces returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


