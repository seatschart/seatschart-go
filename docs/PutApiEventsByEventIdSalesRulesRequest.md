# PutApiEventsByEventIdSalesRulesRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PricingSource** | Pointer to **NullableString** | Event override of the workspace pricing source; null &#x3D; workspace setting. | [optional] 
**Quotas** | Pointer to [**[]PutApiEventsByEventIdSalesRulesRequestQuotasInner**](PutApiEventsByEventIdSalesRulesRequestQuotasInner.md) |  | [optional] 
**LotPlans** | Pointer to [**[]PutApiEventsByEventIdSalesRulesRequestLotPlansInner**](PutApiEventsByEventIdSalesRulesRequestLotPlansInner.md) |  | [optional] 
**BuyerLimit** | Pointer to [**NullablePutApiEventsByEventIdSalesRulesRequestBuyerLimit**](PutApiEventsByEventIdSalesRulesRequestBuyerLimit.md) |  | [optional] 

## Methods

### NewPutApiEventsByEventIdSalesRulesRequest

`func NewPutApiEventsByEventIdSalesRulesRequest() *PutApiEventsByEventIdSalesRulesRequest`

NewPutApiEventsByEventIdSalesRulesRequest instantiates a new PutApiEventsByEventIdSalesRulesRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPutApiEventsByEventIdSalesRulesRequestWithDefaults

`func NewPutApiEventsByEventIdSalesRulesRequestWithDefaults() *PutApiEventsByEventIdSalesRulesRequest`

NewPutApiEventsByEventIdSalesRulesRequestWithDefaults instantiates a new PutApiEventsByEventIdSalesRulesRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPricingSource

`func (o *PutApiEventsByEventIdSalesRulesRequest) GetPricingSource() string`

GetPricingSource returns the PricingSource field if non-nil, zero value otherwise.

### GetPricingSourceOk

`func (o *PutApiEventsByEventIdSalesRulesRequest) GetPricingSourceOk() (*string, bool)`

GetPricingSourceOk returns a tuple with the PricingSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPricingSource

`func (o *PutApiEventsByEventIdSalesRulesRequest) SetPricingSource(v string)`

SetPricingSource sets PricingSource field to given value.

### HasPricingSource

`func (o *PutApiEventsByEventIdSalesRulesRequest) HasPricingSource() bool`

HasPricingSource returns a boolean if a field has been set.

### SetPricingSourceNil

`func (o *PutApiEventsByEventIdSalesRulesRequest) SetPricingSourceNil(b bool)`

 SetPricingSourceNil sets the value for PricingSource to be an explicit nil

### UnsetPricingSource
`func (o *PutApiEventsByEventIdSalesRulesRequest) UnsetPricingSource()`

UnsetPricingSource ensures that no value is present for PricingSource, not even an explicit nil
### GetQuotas

`func (o *PutApiEventsByEventIdSalesRulesRequest) GetQuotas() []PutApiEventsByEventIdSalesRulesRequestQuotasInner`

GetQuotas returns the Quotas field if non-nil, zero value otherwise.

### GetQuotasOk

`func (o *PutApiEventsByEventIdSalesRulesRequest) GetQuotasOk() (*[]PutApiEventsByEventIdSalesRulesRequestQuotasInner, bool)`

GetQuotasOk returns a tuple with the Quotas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuotas

`func (o *PutApiEventsByEventIdSalesRulesRequest) SetQuotas(v []PutApiEventsByEventIdSalesRulesRequestQuotasInner)`

SetQuotas sets Quotas field to given value.

### HasQuotas

`func (o *PutApiEventsByEventIdSalesRulesRequest) HasQuotas() bool`

HasQuotas returns a boolean if a field has been set.

### GetLotPlans

`func (o *PutApiEventsByEventIdSalesRulesRequest) GetLotPlans() []PutApiEventsByEventIdSalesRulesRequestLotPlansInner`

GetLotPlans returns the LotPlans field if non-nil, zero value otherwise.

### GetLotPlansOk

`func (o *PutApiEventsByEventIdSalesRulesRequest) GetLotPlansOk() (*[]PutApiEventsByEventIdSalesRulesRequestLotPlansInner, bool)`

GetLotPlansOk returns a tuple with the LotPlans field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLotPlans

`func (o *PutApiEventsByEventIdSalesRulesRequest) SetLotPlans(v []PutApiEventsByEventIdSalesRulesRequestLotPlansInner)`

SetLotPlans sets LotPlans field to given value.

### HasLotPlans

`func (o *PutApiEventsByEventIdSalesRulesRequest) HasLotPlans() bool`

HasLotPlans returns a boolean if a field has been set.

### GetBuyerLimit

`func (o *PutApiEventsByEventIdSalesRulesRequest) GetBuyerLimit() PutApiEventsByEventIdSalesRulesRequestBuyerLimit`

GetBuyerLimit returns the BuyerLimit field if non-nil, zero value otherwise.

### GetBuyerLimitOk

`func (o *PutApiEventsByEventIdSalesRulesRequest) GetBuyerLimitOk() (*PutApiEventsByEventIdSalesRulesRequestBuyerLimit, bool)`

GetBuyerLimitOk returns a tuple with the BuyerLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuyerLimit

`func (o *PutApiEventsByEventIdSalesRulesRequest) SetBuyerLimit(v PutApiEventsByEventIdSalesRulesRequestBuyerLimit)`

SetBuyerLimit sets BuyerLimit field to given value.

### HasBuyerLimit

`func (o *PutApiEventsByEventIdSalesRulesRequest) HasBuyerLimit() bool`

HasBuyerLimit returns a boolean if a field has been set.

### SetBuyerLimitNil

`func (o *PutApiEventsByEventIdSalesRulesRequest) SetBuyerLimitNil(b bool)`

 SetBuyerLimitNil sets the value for BuyerLimit to be an explicit nil

### UnsetBuyerLimit
`func (o *PutApiEventsByEventIdSalesRulesRequest) UnsetBuyerLimit()`

UnsetBuyerLimit ensures that no value is present for BuyerLimit, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


