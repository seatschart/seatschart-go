# PostApiChartsByChartIdCategoriesRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Stable id (generated when omitted). | [optional] 
**Name** | **string** | Seats and GA areas reference categories by name. | 
**Color** | Pointer to **string** | Category colour &#x60;#RRGGBB&#x60;. | [optional] 
**TicketTypes** | Pointer to [**[]PostApiChartsByChartIdCategoriesRequestTicketTypesInner**](PostApiChartsByChartIdCategoriesRequestTicketTypesInner.md) | Ticket types (tiers) sold in this category. | [optional] 

## Methods

### NewPostApiChartsByChartIdCategoriesRequest

`func NewPostApiChartsByChartIdCategoriesRequest(name string, ) *PostApiChartsByChartIdCategoriesRequest`

NewPostApiChartsByChartIdCategoriesRequest instantiates a new PostApiChartsByChartIdCategoriesRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostApiChartsByChartIdCategoriesRequestWithDefaults

`func NewPostApiChartsByChartIdCategoriesRequestWithDefaults() *PostApiChartsByChartIdCategoriesRequest`

NewPostApiChartsByChartIdCategoriesRequestWithDefaults instantiates a new PostApiChartsByChartIdCategoriesRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PostApiChartsByChartIdCategoriesRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PostApiChartsByChartIdCategoriesRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PostApiChartsByChartIdCategoriesRequest) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *PostApiChartsByChartIdCategoriesRequest) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *PostApiChartsByChartIdCategoriesRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PostApiChartsByChartIdCategoriesRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PostApiChartsByChartIdCategoriesRequest) SetName(v string)`

SetName sets Name field to given value.


### GetColor

`func (o *PostApiChartsByChartIdCategoriesRequest) GetColor() string`

GetColor returns the Color field if non-nil, zero value otherwise.

### GetColorOk

`func (o *PostApiChartsByChartIdCategoriesRequest) GetColorOk() (*string, bool)`

GetColorOk returns a tuple with the Color field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColor

`func (o *PostApiChartsByChartIdCategoriesRequest) SetColor(v string)`

SetColor sets Color field to given value.

### HasColor

`func (o *PostApiChartsByChartIdCategoriesRequest) HasColor() bool`

HasColor returns a boolean if a field has been set.

### GetTicketTypes

`func (o *PostApiChartsByChartIdCategoriesRequest) GetTicketTypes() []PostApiChartsByChartIdCategoriesRequestTicketTypesInner`

GetTicketTypes returns the TicketTypes field if non-nil, zero value otherwise.

### GetTicketTypesOk

`func (o *PostApiChartsByChartIdCategoriesRequest) GetTicketTypesOk() (*[]PostApiChartsByChartIdCategoriesRequestTicketTypesInner, bool)`

GetTicketTypesOk returns a tuple with the TicketTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTicketTypes

`func (o *PostApiChartsByChartIdCategoriesRequest) SetTicketTypes(v []PostApiChartsByChartIdCategoriesRequestTicketTypesInner)`

SetTicketTypes sets TicketTypes field to given value.

### HasTicketTypes

`func (o *PostApiChartsByChartIdCategoriesRequest) HasTicketTypes() bool`

HasTicketTypes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


