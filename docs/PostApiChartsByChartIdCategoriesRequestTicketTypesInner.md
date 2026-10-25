# PostApiChartsByChartIdCategoriesRequestTicketTypesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Name** | **string** |  | 
**PriceCents** | **int32** |  | 
**Description** | Pointer to **string** |  | [optional] 
**Primary** | Pointer to **bool** |  | [optional] 
**Unavailable** | Pointer to **bool** |  | [optional] 
**UnavailableReason** | Pointer to **string** |  | [optional] 
**OriginalPriceCents** | Pointer to **int32** |  | [optional] 

## Methods

### NewPostApiChartsByChartIdCategoriesRequestTicketTypesInner

`func NewPostApiChartsByChartIdCategoriesRequestTicketTypesInner(id string, name string, priceCents int32, ) *PostApiChartsByChartIdCategoriesRequestTicketTypesInner`

NewPostApiChartsByChartIdCategoriesRequestTicketTypesInner instantiates a new PostApiChartsByChartIdCategoriesRequestTicketTypesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostApiChartsByChartIdCategoriesRequestTicketTypesInnerWithDefaults

`func NewPostApiChartsByChartIdCategoriesRequestTicketTypesInnerWithDefaults() *PostApiChartsByChartIdCategoriesRequestTicketTypesInner`

NewPostApiChartsByChartIdCategoriesRequestTicketTypesInnerWithDefaults instantiates a new PostApiChartsByChartIdCategoriesRequestTicketTypesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PostApiChartsByChartIdCategoriesRequestTicketTypesInner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PostApiChartsByChartIdCategoriesRequestTicketTypesInner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PostApiChartsByChartIdCategoriesRequestTicketTypesInner) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *PostApiChartsByChartIdCategoriesRequestTicketTypesInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PostApiChartsByChartIdCategoriesRequestTicketTypesInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PostApiChartsByChartIdCategoriesRequestTicketTypesInner) SetName(v string)`

SetName sets Name field to given value.


### GetPriceCents

`func (o *PostApiChartsByChartIdCategoriesRequestTicketTypesInner) GetPriceCents() int32`

GetPriceCents returns the PriceCents field if non-nil, zero value otherwise.

### GetPriceCentsOk

`func (o *PostApiChartsByChartIdCategoriesRequestTicketTypesInner) GetPriceCentsOk() (*int32, bool)`

GetPriceCentsOk returns a tuple with the PriceCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceCents

`func (o *PostApiChartsByChartIdCategoriesRequestTicketTypesInner) SetPriceCents(v int32)`

SetPriceCents sets PriceCents field to given value.


### GetDescription

`func (o *PostApiChartsByChartIdCategoriesRequestTicketTypesInner) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PostApiChartsByChartIdCategoriesRequestTicketTypesInner) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PostApiChartsByChartIdCategoriesRequestTicketTypesInner) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PostApiChartsByChartIdCategoriesRequestTicketTypesInner) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetPrimary

`func (o *PostApiChartsByChartIdCategoriesRequestTicketTypesInner) GetPrimary() bool`

GetPrimary returns the Primary field if non-nil, zero value otherwise.

### GetPrimaryOk

`func (o *PostApiChartsByChartIdCategoriesRequestTicketTypesInner) GetPrimaryOk() (*bool, bool)`

GetPrimaryOk returns a tuple with the Primary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimary

`func (o *PostApiChartsByChartIdCategoriesRequestTicketTypesInner) SetPrimary(v bool)`

SetPrimary sets Primary field to given value.

### HasPrimary

`func (o *PostApiChartsByChartIdCategoriesRequestTicketTypesInner) HasPrimary() bool`

HasPrimary returns a boolean if a field has been set.

### GetUnavailable

`func (o *PostApiChartsByChartIdCategoriesRequestTicketTypesInner) GetUnavailable() bool`

GetUnavailable returns the Unavailable field if non-nil, zero value otherwise.

### GetUnavailableOk

`func (o *PostApiChartsByChartIdCategoriesRequestTicketTypesInner) GetUnavailableOk() (*bool, bool)`

GetUnavailableOk returns a tuple with the Unavailable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnavailable

`func (o *PostApiChartsByChartIdCategoriesRequestTicketTypesInner) SetUnavailable(v bool)`

SetUnavailable sets Unavailable field to given value.

### HasUnavailable

`func (o *PostApiChartsByChartIdCategoriesRequestTicketTypesInner) HasUnavailable() bool`

HasUnavailable returns a boolean if a field has been set.

### GetUnavailableReason

`func (o *PostApiChartsByChartIdCategoriesRequestTicketTypesInner) GetUnavailableReason() string`

GetUnavailableReason returns the UnavailableReason field if non-nil, zero value otherwise.

### GetUnavailableReasonOk

`func (o *PostApiChartsByChartIdCategoriesRequestTicketTypesInner) GetUnavailableReasonOk() (*string, bool)`

GetUnavailableReasonOk returns a tuple with the UnavailableReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnavailableReason

`func (o *PostApiChartsByChartIdCategoriesRequestTicketTypesInner) SetUnavailableReason(v string)`

SetUnavailableReason sets UnavailableReason field to given value.

### HasUnavailableReason

`func (o *PostApiChartsByChartIdCategoriesRequestTicketTypesInner) HasUnavailableReason() bool`

HasUnavailableReason returns a boolean if a field has been set.

### GetOriginalPriceCents

`func (o *PostApiChartsByChartIdCategoriesRequestTicketTypesInner) GetOriginalPriceCents() int32`

GetOriginalPriceCents returns the OriginalPriceCents field if non-nil, zero value otherwise.

### GetOriginalPriceCentsOk

`func (o *PostApiChartsByChartIdCategoriesRequestTicketTypesInner) GetOriginalPriceCentsOk() (*int32, bool)`

GetOriginalPriceCentsOk returns a tuple with the OriginalPriceCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginalPriceCents

`func (o *PostApiChartsByChartIdCategoriesRequestTicketTypesInner) SetOriginalPriceCents(v int32)`

SetOriginalPriceCents sets OriginalPriceCents field to given value.

### HasOriginalPriceCents

`func (o *PostApiChartsByChartIdCategoriesRequestTicketTypesInner) HasOriginalPriceCents() bool`

HasOriginalPriceCents returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


