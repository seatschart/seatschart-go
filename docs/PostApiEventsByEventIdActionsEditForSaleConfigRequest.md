# PostApiEventsByEventIdActionsEditForSaleConfigRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ForSale** | Pointer to [**[]PostApiEventsByEventIdActionsEditForSaleConfigRequestForSaleInner**](PostApiEventsByEventIdActionsEditForSaleConfigRequestForSaleInner.md) | Objects, GA places and categories to put on sale. | [optional] 
**NotForSale** | Pointer to [**[]PostApiEventsByEventIdActionsEditForSaleConfigRequestForSaleInner**](PostApiEventsByEventIdActionsEditForSaleConfigRequestForSaleInner.md) | Objects, GA places and categories to take off sale. | [optional] 
**CascadeSeason** | Pointer to **bool** | Sync live season siblings when this event is the season source (default false). | [optional] 

## Methods

### NewPostApiEventsByEventIdActionsEditForSaleConfigRequest

`func NewPostApiEventsByEventIdActionsEditForSaleConfigRequest() *PostApiEventsByEventIdActionsEditForSaleConfigRequest`

NewPostApiEventsByEventIdActionsEditForSaleConfigRequest instantiates a new PostApiEventsByEventIdActionsEditForSaleConfigRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostApiEventsByEventIdActionsEditForSaleConfigRequestWithDefaults

`func NewPostApiEventsByEventIdActionsEditForSaleConfigRequestWithDefaults() *PostApiEventsByEventIdActionsEditForSaleConfigRequest`

NewPostApiEventsByEventIdActionsEditForSaleConfigRequestWithDefaults instantiates a new PostApiEventsByEventIdActionsEditForSaleConfigRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetForSale

`func (o *PostApiEventsByEventIdActionsEditForSaleConfigRequest) GetForSale() []PostApiEventsByEventIdActionsEditForSaleConfigRequestForSaleInner`

GetForSale returns the ForSale field if non-nil, zero value otherwise.

### GetForSaleOk

`func (o *PostApiEventsByEventIdActionsEditForSaleConfigRequest) GetForSaleOk() (*[]PostApiEventsByEventIdActionsEditForSaleConfigRequestForSaleInner, bool)`

GetForSaleOk returns a tuple with the ForSale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForSale

`func (o *PostApiEventsByEventIdActionsEditForSaleConfigRequest) SetForSale(v []PostApiEventsByEventIdActionsEditForSaleConfigRequestForSaleInner)`

SetForSale sets ForSale field to given value.

### HasForSale

`func (o *PostApiEventsByEventIdActionsEditForSaleConfigRequest) HasForSale() bool`

HasForSale returns a boolean if a field has been set.

### GetNotForSale

`func (o *PostApiEventsByEventIdActionsEditForSaleConfigRequest) GetNotForSale() []PostApiEventsByEventIdActionsEditForSaleConfigRequestForSaleInner`

GetNotForSale returns the NotForSale field if non-nil, zero value otherwise.

### GetNotForSaleOk

`func (o *PostApiEventsByEventIdActionsEditForSaleConfigRequest) GetNotForSaleOk() (*[]PostApiEventsByEventIdActionsEditForSaleConfigRequestForSaleInner, bool)`

GetNotForSaleOk returns a tuple with the NotForSale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotForSale

`func (o *PostApiEventsByEventIdActionsEditForSaleConfigRequest) SetNotForSale(v []PostApiEventsByEventIdActionsEditForSaleConfigRequestForSaleInner)`

SetNotForSale sets NotForSale field to given value.

### HasNotForSale

`func (o *PostApiEventsByEventIdActionsEditForSaleConfigRequest) HasNotForSale() bool`

HasNotForSale returns a boolean if a field has been set.

### GetCascadeSeason

`func (o *PostApiEventsByEventIdActionsEditForSaleConfigRequest) GetCascadeSeason() bool`

GetCascadeSeason returns the CascadeSeason field if non-nil, zero value otherwise.

### GetCascadeSeasonOk

`func (o *PostApiEventsByEventIdActionsEditForSaleConfigRequest) GetCascadeSeasonOk() (*bool, bool)`

GetCascadeSeasonOk returns a tuple with the CascadeSeason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCascadeSeason

`func (o *PostApiEventsByEventIdActionsEditForSaleConfigRequest) SetCascadeSeason(v bool)`

SetCascadeSeason sets CascadeSeason field to given value.

### HasCascadeSeason

`func (o *PostApiEventsByEventIdActionsEditForSaleConfigRequest) HasCascadeSeason() bool`

HasCascadeSeason returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


