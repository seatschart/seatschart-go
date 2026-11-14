# PostApiBillingActionsChangePlanRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PlanId** | **string** | The plan to switch to. Only &#x60;free&#x60; changes here; paid plans (&#x60;starter&#x60;, &#x60;pro&#x60;, &#x60;business&#x60;, &#x60;scale&#x60;) start through checkout. | 

## Methods

### NewPostApiBillingActionsChangePlanRequest

`func NewPostApiBillingActionsChangePlanRequest(planId string, ) *PostApiBillingActionsChangePlanRequest`

NewPostApiBillingActionsChangePlanRequest instantiates a new PostApiBillingActionsChangePlanRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostApiBillingActionsChangePlanRequestWithDefaults

`func NewPostApiBillingActionsChangePlanRequestWithDefaults() *PostApiBillingActionsChangePlanRequest`

NewPostApiBillingActionsChangePlanRequestWithDefaults instantiates a new PostApiBillingActionsChangePlanRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPlanId

`func (o *PostApiBillingActionsChangePlanRequest) GetPlanId() string`

GetPlanId returns the PlanId field if non-nil, zero value otherwise.

### GetPlanIdOk

`func (o *PostApiBillingActionsChangePlanRequest) GetPlanIdOk() (*string, bool)`

GetPlanIdOk returns a tuple with the PlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanId

`func (o *PostApiBillingActionsChangePlanRequest) SetPlanId(v string)`

SetPlanId sets PlanId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


