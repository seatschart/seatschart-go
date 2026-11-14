# GetApiWorkspacesOverview200ResponseItemsInnerRevenue

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Known** | **bool** |  | 
**Cents** | **NullableInt32** | Catalog price of the sold seats (US cents); &#x60;null&#x60; when not known. | 

## Methods

### NewGetApiWorkspacesOverview200ResponseItemsInnerRevenue

`func NewGetApiWorkspacesOverview200ResponseItemsInnerRevenue(known bool, cents NullableInt32, ) *GetApiWorkspacesOverview200ResponseItemsInnerRevenue`

NewGetApiWorkspacesOverview200ResponseItemsInnerRevenue instantiates a new GetApiWorkspacesOverview200ResponseItemsInnerRevenue object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetApiWorkspacesOverview200ResponseItemsInnerRevenueWithDefaults

`func NewGetApiWorkspacesOverview200ResponseItemsInnerRevenueWithDefaults() *GetApiWorkspacesOverview200ResponseItemsInnerRevenue`

NewGetApiWorkspacesOverview200ResponseItemsInnerRevenueWithDefaults instantiates a new GetApiWorkspacesOverview200ResponseItemsInnerRevenue object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKnown

`func (o *GetApiWorkspacesOverview200ResponseItemsInnerRevenue) GetKnown() bool`

GetKnown returns the Known field if non-nil, zero value otherwise.

### GetKnownOk

`func (o *GetApiWorkspacesOverview200ResponseItemsInnerRevenue) GetKnownOk() (*bool, bool)`

GetKnownOk returns a tuple with the Known field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKnown

`func (o *GetApiWorkspacesOverview200ResponseItemsInnerRevenue) SetKnown(v bool)`

SetKnown sets Known field to given value.


### GetCents

`func (o *GetApiWorkspacesOverview200ResponseItemsInnerRevenue) GetCents() int32`

GetCents returns the Cents field if non-nil, zero value otherwise.

### GetCentsOk

`func (o *GetApiWorkspacesOverview200ResponseItemsInnerRevenue) GetCentsOk() (*int32, bool)`

GetCentsOk returns a tuple with the Cents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCents

`func (o *GetApiWorkspacesOverview200ResponseItemsInnerRevenue) SetCents(v int32)`

SetCents sets Cents field to given value.


### SetCentsNil

`func (o *GetApiWorkspacesOverview200ResponseItemsInnerRevenue) SetCentsNil(b bool)`

 SetCentsNil sets the value for Cents to be an explicit nil

### UnsetCents
`func (o *GetApiWorkspacesOverview200ResponseItemsInnerRevenue) UnsetCents()`

UnsetCents ensures that no value is present for Cents, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


