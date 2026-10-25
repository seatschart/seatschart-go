# PutApiEventsByEventIdSalesRulesRequestLotPlansInnerLotsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Name** | **string** |  | 
**Quantity** | Pointer to **int32** | Tickets sold (and held) in this lot before the next opens. | [optional] 
**UntilDate** | Pointer to **time.Time** | The lot closes at this instant. | [optional] 
**Prices** | Pointer to **map[string]int32** | Catalog modes: price per ticket type id in cents while this lot is active. | [optional] 

## Methods

### NewPutApiEventsByEventIdSalesRulesRequestLotPlansInnerLotsInner

`func NewPutApiEventsByEventIdSalesRulesRequestLotPlansInnerLotsInner(id string, name string, ) *PutApiEventsByEventIdSalesRulesRequestLotPlansInnerLotsInner`

NewPutApiEventsByEventIdSalesRulesRequestLotPlansInnerLotsInner instantiates a new PutApiEventsByEventIdSalesRulesRequestLotPlansInnerLotsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPutApiEventsByEventIdSalesRulesRequestLotPlansInnerLotsInnerWithDefaults

`func NewPutApiEventsByEventIdSalesRulesRequestLotPlansInnerLotsInnerWithDefaults() *PutApiEventsByEventIdSalesRulesRequestLotPlansInnerLotsInner`

NewPutApiEventsByEventIdSalesRulesRequestLotPlansInnerLotsInnerWithDefaults instantiates a new PutApiEventsByEventIdSalesRulesRequestLotPlansInnerLotsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PutApiEventsByEventIdSalesRulesRequestLotPlansInnerLotsInner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PutApiEventsByEventIdSalesRulesRequestLotPlansInnerLotsInner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PutApiEventsByEventIdSalesRulesRequestLotPlansInnerLotsInner) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *PutApiEventsByEventIdSalesRulesRequestLotPlansInnerLotsInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PutApiEventsByEventIdSalesRulesRequestLotPlansInnerLotsInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PutApiEventsByEventIdSalesRulesRequestLotPlansInnerLotsInner) SetName(v string)`

SetName sets Name field to given value.


### GetQuantity

`func (o *PutApiEventsByEventIdSalesRulesRequestLotPlansInnerLotsInner) GetQuantity() int32`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *PutApiEventsByEventIdSalesRulesRequestLotPlansInnerLotsInner) GetQuantityOk() (*int32, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *PutApiEventsByEventIdSalesRulesRequestLotPlansInnerLotsInner) SetQuantity(v int32)`

SetQuantity sets Quantity field to given value.

### HasQuantity

`func (o *PutApiEventsByEventIdSalesRulesRequestLotPlansInnerLotsInner) HasQuantity() bool`

HasQuantity returns a boolean if a field has been set.

### GetUntilDate

`func (o *PutApiEventsByEventIdSalesRulesRequestLotPlansInnerLotsInner) GetUntilDate() time.Time`

GetUntilDate returns the UntilDate field if non-nil, zero value otherwise.

### GetUntilDateOk

`func (o *PutApiEventsByEventIdSalesRulesRequestLotPlansInnerLotsInner) GetUntilDateOk() (*time.Time, bool)`

GetUntilDateOk returns a tuple with the UntilDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUntilDate

`func (o *PutApiEventsByEventIdSalesRulesRequestLotPlansInnerLotsInner) SetUntilDate(v time.Time)`

SetUntilDate sets UntilDate field to given value.

### HasUntilDate

`func (o *PutApiEventsByEventIdSalesRulesRequestLotPlansInnerLotsInner) HasUntilDate() bool`

HasUntilDate returns a boolean if a field has been set.

### GetPrices

`func (o *PutApiEventsByEventIdSalesRulesRequestLotPlansInnerLotsInner) GetPrices() map[string]int32`

GetPrices returns the Prices field if non-nil, zero value otherwise.

### GetPricesOk

`func (o *PutApiEventsByEventIdSalesRulesRequestLotPlansInnerLotsInner) GetPricesOk() (*map[string]int32, bool)`

GetPricesOk returns a tuple with the Prices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrices

`func (o *PutApiEventsByEventIdSalesRulesRequestLotPlansInnerLotsInner) SetPrices(v map[string]int32)`

SetPrices sets Prices field to given value.

### HasPrices

`func (o *PutApiEventsByEventIdSalesRulesRequestLotPlansInnerLotsInner) HasPrices() bool`

HasPrices returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


