# PostApiEventsActionsEditForSaleConfigRequestEventsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Event** | **string** | Event UUID or event key. | 
**ForSale** | Pointer to [**[]PostApiEventsByEventIdActionsEditForSaleConfigRequestForSaleInner**](PostApiEventsByEventIdActionsEditForSaleConfigRequestForSaleInner.md) | Objects, GA places and categories to put on sale. | [optional] 
**NotForSale** | Pointer to [**[]PostApiEventsByEventIdActionsEditForSaleConfigRequestForSaleInner**](PostApiEventsByEventIdActionsEditForSaleConfigRequestForSaleInner.md) | Objects, GA places and categories to take off sale. | [optional] 
**CascadeSeason** | Pointer to **bool** | Sync live season siblings when this event is the season source (default false). | [optional] 

## Methods

### NewPostApiEventsActionsEditForSaleConfigRequestEventsInner

`func NewPostApiEventsActionsEditForSaleConfigRequestEventsInner(event string, ) *PostApiEventsActionsEditForSaleConfigRequestEventsInner`

NewPostApiEventsActionsEditForSaleConfigRequestEventsInner instantiates a new PostApiEventsActionsEditForSaleConfigRequestEventsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostApiEventsActionsEditForSaleConfigRequestEventsInnerWithDefaults

`func NewPostApiEventsActionsEditForSaleConfigRequestEventsInnerWithDefaults() *PostApiEventsActionsEditForSaleConfigRequestEventsInner`

NewPostApiEventsActionsEditForSaleConfigRequestEventsInnerWithDefaults instantiates a new PostApiEventsActionsEditForSaleConfigRequestEventsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEvent

`func (o *PostApiEventsActionsEditForSaleConfigRequestEventsInner) GetEvent() string`

GetEvent returns the Event field if non-nil, zero value otherwise.

### GetEventOk

`func (o *PostApiEventsActionsEditForSaleConfigRequestEventsInner) GetEventOk() (*string, bool)`

GetEventOk returns a tuple with the Event field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvent

`func (o *PostApiEventsActionsEditForSaleConfigRequestEventsInner) SetEvent(v string)`

SetEvent sets Event field to given value.


### GetForSale

`func (o *PostApiEventsActionsEditForSaleConfigRequestEventsInner) GetForSale() []PostApiEventsByEventIdActionsEditForSaleConfigRequestForSaleInner`

GetForSale returns the ForSale field if non-nil, zero value otherwise.

### GetForSaleOk

`func (o *PostApiEventsActionsEditForSaleConfigRequestEventsInner) GetForSaleOk() (*[]PostApiEventsByEventIdActionsEditForSaleConfigRequestForSaleInner, bool)`

GetForSaleOk returns a tuple with the ForSale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForSale

`func (o *PostApiEventsActionsEditForSaleConfigRequestEventsInner) SetForSale(v []PostApiEventsByEventIdActionsEditForSaleConfigRequestForSaleInner)`

SetForSale sets ForSale field to given value.

### HasForSale

`func (o *PostApiEventsActionsEditForSaleConfigRequestEventsInner) HasForSale() bool`

HasForSale returns a boolean if a field has been set.

### GetNotForSale

`func (o *PostApiEventsActionsEditForSaleConfigRequestEventsInner) GetNotForSale() []PostApiEventsByEventIdActionsEditForSaleConfigRequestForSaleInner`

GetNotForSale returns the NotForSale field if non-nil, zero value otherwise.

### GetNotForSaleOk

`func (o *PostApiEventsActionsEditForSaleConfigRequestEventsInner) GetNotForSaleOk() (*[]PostApiEventsByEventIdActionsEditForSaleConfigRequestForSaleInner, bool)`

GetNotForSaleOk returns a tuple with the NotForSale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotForSale

`func (o *PostApiEventsActionsEditForSaleConfigRequestEventsInner) SetNotForSale(v []PostApiEventsByEventIdActionsEditForSaleConfigRequestForSaleInner)`

SetNotForSale sets NotForSale field to given value.

### HasNotForSale

`func (o *PostApiEventsActionsEditForSaleConfigRequestEventsInner) HasNotForSale() bool`

HasNotForSale returns a boolean if a field has been set.

### GetCascadeSeason

`func (o *PostApiEventsActionsEditForSaleConfigRequestEventsInner) GetCascadeSeason() bool`

GetCascadeSeason returns the CascadeSeason field if non-nil, zero value otherwise.

### GetCascadeSeasonOk

`func (o *PostApiEventsActionsEditForSaleConfigRequestEventsInner) GetCascadeSeasonOk() (*bool, bool)`

GetCascadeSeasonOk returns a tuple with the CascadeSeason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCascadeSeason

`func (o *PostApiEventsActionsEditForSaleConfigRequestEventsInner) SetCascadeSeason(v bool)`

SetCascadeSeason sets CascadeSeason field to given value.

### HasCascadeSeason

`func (o *PostApiEventsActionsEditForSaleConfigRequestEventsInner) HasCascadeSeason() bool`

HasCascadeSeason returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


