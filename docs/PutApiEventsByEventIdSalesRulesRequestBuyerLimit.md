# PutApiEventsByEventIdSalesRulesRequestBuyerLimit

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MaxTickets** | **int32** |  | 
**RequireBuyerId** | Pointer to **bool** | Refuse holds without a ticketBuyerId. | [optional] 

## Methods

### NewPutApiEventsByEventIdSalesRulesRequestBuyerLimit

`func NewPutApiEventsByEventIdSalesRulesRequestBuyerLimit(maxTickets int32, ) *PutApiEventsByEventIdSalesRulesRequestBuyerLimit`

NewPutApiEventsByEventIdSalesRulesRequestBuyerLimit instantiates a new PutApiEventsByEventIdSalesRulesRequestBuyerLimit object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPutApiEventsByEventIdSalesRulesRequestBuyerLimitWithDefaults

`func NewPutApiEventsByEventIdSalesRulesRequestBuyerLimitWithDefaults() *PutApiEventsByEventIdSalesRulesRequestBuyerLimit`

NewPutApiEventsByEventIdSalesRulesRequestBuyerLimitWithDefaults instantiates a new PutApiEventsByEventIdSalesRulesRequestBuyerLimit object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMaxTickets

`func (o *PutApiEventsByEventIdSalesRulesRequestBuyerLimit) GetMaxTickets() int32`

GetMaxTickets returns the MaxTickets field if non-nil, zero value otherwise.

### GetMaxTicketsOk

`func (o *PutApiEventsByEventIdSalesRulesRequestBuyerLimit) GetMaxTicketsOk() (*int32, bool)`

GetMaxTicketsOk returns a tuple with the MaxTickets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxTickets

`func (o *PutApiEventsByEventIdSalesRulesRequestBuyerLimit) SetMaxTickets(v int32)`

SetMaxTickets sets MaxTickets field to given value.


### GetRequireBuyerId

`func (o *PutApiEventsByEventIdSalesRulesRequestBuyerLimit) GetRequireBuyerId() bool`

GetRequireBuyerId returns the RequireBuyerId field if non-nil, zero value otherwise.

### GetRequireBuyerIdOk

`func (o *PutApiEventsByEventIdSalesRulesRequestBuyerLimit) GetRequireBuyerIdOk() (*bool, bool)`

GetRequireBuyerIdOk returns a tuple with the RequireBuyerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequireBuyerId

`func (o *PutApiEventsByEventIdSalesRulesRequestBuyerLimit) SetRequireBuyerId(v bool)`

SetRequireBuyerId sets RequireBuyerId field to given value.

### HasRequireBuyerId

`func (o *PutApiEventsByEventIdSalesRulesRequestBuyerLimit) HasRequireBuyerId() bool`

HasRequireBuyerId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


