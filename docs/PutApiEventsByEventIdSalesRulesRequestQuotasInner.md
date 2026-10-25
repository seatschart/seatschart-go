# PutApiEventsByEventIdSalesRulesRequestQuotasInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Label** | **string** |  | 
**TicketTypes** | **[]string** | Ticket type ids counted by this quota (e.g. [\&quot;meia\&quot;]). | 
**MaxShare** | Pointer to **float32** | Share of the event capacity (0.4 &#x3D; 40%, Lei 12.933). | [optional] 
**MaxCount** | Pointer to **int32** | Absolute ticket count; with maxShare the lower limit applies. | [optional] 

## Methods

### NewPutApiEventsByEventIdSalesRulesRequestQuotasInner

`func NewPutApiEventsByEventIdSalesRulesRequestQuotasInner(id string, label string, ticketTypes []string, ) *PutApiEventsByEventIdSalesRulesRequestQuotasInner`

NewPutApiEventsByEventIdSalesRulesRequestQuotasInner instantiates a new PutApiEventsByEventIdSalesRulesRequestQuotasInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPutApiEventsByEventIdSalesRulesRequestQuotasInnerWithDefaults

`func NewPutApiEventsByEventIdSalesRulesRequestQuotasInnerWithDefaults() *PutApiEventsByEventIdSalesRulesRequestQuotasInner`

NewPutApiEventsByEventIdSalesRulesRequestQuotasInnerWithDefaults instantiates a new PutApiEventsByEventIdSalesRulesRequestQuotasInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PutApiEventsByEventIdSalesRulesRequestQuotasInner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PutApiEventsByEventIdSalesRulesRequestQuotasInner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PutApiEventsByEventIdSalesRulesRequestQuotasInner) SetId(v string)`

SetId sets Id field to given value.


### GetLabel

`func (o *PutApiEventsByEventIdSalesRulesRequestQuotasInner) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *PutApiEventsByEventIdSalesRulesRequestQuotasInner) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *PutApiEventsByEventIdSalesRulesRequestQuotasInner) SetLabel(v string)`

SetLabel sets Label field to given value.


### GetTicketTypes

`func (o *PutApiEventsByEventIdSalesRulesRequestQuotasInner) GetTicketTypes() []string`

GetTicketTypes returns the TicketTypes field if non-nil, zero value otherwise.

### GetTicketTypesOk

`func (o *PutApiEventsByEventIdSalesRulesRequestQuotasInner) GetTicketTypesOk() (*[]string, bool)`

GetTicketTypesOk returns a tuple with the TicketTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTicketTypes

`func (o *PutApiEventsByEventIdSalesRulesRequestQuotasInner) SetTicketTypes(v []string)`

SetTicketTypes sets TicketTypes field to given value.


### GetMaxShare

`func (o *PutApiEventsByEventIdSalesRulesRequestQuotasInner) GetMaxShare() float32`

GetMaxShare returns the MaxShare field if non-nil, zero value otherwise.

### GetMaxShareOk

`func (o *PutApiEventsByEventIdSalesRulesRequestQuotasInner) GetMaxShareOk() (*float32, bool)`

GetMaxShareOk returns a tuple with the MaxShare field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxShare

`func (o *PutApiEventsByEventIdSalesRulesRequestQuotasInner) SetMaxShare(v float32)`

SetMaxShare sets MaxShare field to given value.

### HasMaxShare

`func (o *PutApiEventsByEventIdSalesRulesRequestQuotasInner) HasMaxShare() bool`

HasMaxShare returns a boolean if a field has been set.

### GetMaxCount

`func (o *PutApiEventsByEventIdSalesRulesRequestQuotasInner) GetMaxCount() int32`

GetMaxCount returns the MaxCount field if non-nil, zero value otherwise.

### GetMaxCountOk

`func (o *PutApiEventsByEventIdSalesRulesRequestQuotasInner) GetMaxCountOk() (*int32, bool)`

GetMaxCountOk returns a tuple with the MaxCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxCount

`func (o *PutApiEventsByEventIdSalesRulesRequestQuotasInner) SetMaxCount(v int32)`

SetMaxCount sets MaxCount field to given value.

### HasMaxCount

`func (o *PutApiEventsByEventIdSalesRulesRequestQuotasInner) HasMaxCount() bool`

HasMaxCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


