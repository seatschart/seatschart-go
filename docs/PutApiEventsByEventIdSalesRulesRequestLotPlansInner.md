# PutApiEventsByEventIdSalesRulesRequestLotPlansInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Name** | Pointer to **string** |  | [optional] 
**Categories** | Pointer to **[]string** | Seat categories sold through this plan; omit for the plan that sells every other category. | [optional] 
**Lots** | [**[]PutApiEventsByEventIdSalesRulesRequestLotPlansInnerLotsInner**](PutApiEventsByEventIdSalesRulesRequestLotPlansInnerLotsInner.md) |  | 

## Methods

### NewPutApiEventsByEventIdSalesRulesRequestLotPlansInner

`func NewPutApiEventsByEventIdSalesRulesRequestLotPlansInner(id string, lots []PutApiEventsByEventIdSalesRulesRequestLotPlansInnerLotsInner, ) *PutApiEventsByEventIdSalesRulesRequestLotPlansInner`

NewPutApiEventsByEventIdSalesRulesRequestLotPlansInner instantiates a new PutApiEventsByEventIdSalesRulesRequestLotPlansInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPutApiEventsByEventIdSalesRulesRequestLotPlansInnerWithDefaults

`func NewPutApiEventsByEventIdSalesRulesRequestLotPlansInnerWithDefaults() *PutApiEventsByEventIdSalesRulesRequestLotPlansInner`

NewPutApiEventsByEventIdSalesRulesRequestLotPlansInnerWithDefaults instantiates a new PutApiEventsByEventIdSalesRulesRequestLotPlansInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PutApiEventsByEventIdSalesRulesRequestLotPlansInner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PutApiEventsByEventIdSalesRulesRequestLotPlansInner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PutApiEventsByEventIdSalesRulesRequestLotPlansInner) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *PutApiEventsByEventIdSalesRulesRequestLotPlansInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PutApiEventsByEventIdSalesRulesRequestLotPlansInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PutApiEventsByEventIdSalesRulesRequestLotPlansInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PutApiEventsByEventIdSalesRulesRequestLotPlansInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetCategories

`func (o *PutApiEventsByEventIdSalesRulesRequestLotPlansInner) GetCategories() []string`

GetCategories returns the Categories field if non-nil, zero value otherwise.

### GetCategoriesOk

`func (o *PutApiEventsByEventIdSalesRulesRequestLotPlansInner) GetCategoriesOk() (*[]string, bool)`

GetCategoriesOk returns a tuple with the Categories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategories

`func (o *PutApiEventsByEventIdSalesRulesRequestLotPlansInner) SetCategories(v []string)`

SetCategories sets Categories field to given value.

### HasCategories

`func (o *PutApiEventsByEventIdSalesRulesRequestLotPlansInner) HasCategories() bool`

HasCategories returns a boolean if a field has been set.

### GetLots

`func (o *PutApiEventsByEventIdSalesRulesRequestLotPlansInner) GetLots() []PutApiEventsByEventIdSalesRulesRequestLotPlansInnerLotsInner`

GetLots returns the Lots field if non-nil, zero value otherwise.

### GetLotsOk

`func (o *PutApiEventsByEventIdSalesRulesRequestLotPlansInner) GetLotsOk() (*[]PutApiEventsByEventIdSalesRulesRequestLotPlansInnerLotsInner, bool)`

GetLotsOk returns a tuple with the Lots field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLots

`func (o *PutApiEventsByEventIdSalesRulesRequestLotPlansInner) SetLots(v []PutApiEventsByEventIdSalesRulesRequestLotPlansInnerLotsInner)`

SetLots sets Lots field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


