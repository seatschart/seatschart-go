# PatchApiChartsByChartIdCategoriesByCategoryIdRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**Color** | Pointer to **string** | Category colour &#x60;#RRGGBB&#x60;. | [optional] 
**TicketTypes** | Pointer to [**[]PostApiChartsByChartIdCategoriesRequestTicketTypesInner**](PostApiChartsByChartIdCategoriesRequestTicketTypesInner.md) | Ticket types (tiers) sold in this category. | [optional] 

## Methods

### NewPatchApiChartsByChartIdCategoriesByCategoryIdRequest

`func NewPatchApiChartsByChartIdCategoriesByCategoryIdRequest() *PatchApiChartsByChartIdCategoriesByCategoryIdRequest`

NewPatchApiChartsByChartIdCategoriesByCategoryIdRequest instantiates a new PatchApiChartsByChartIdCategoriesByCategoryIdRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchApiChartsByChartIdCategoriesByCategoryIdRequestWithDefaults

`func NewPatchApiChartsByChartIdCategoriesByCategoryIdRequestWithDefaults() *PatchApiChartsByChartIdCategoriesByCategoryIdRequest`

NewPatchApiChartsByChartIdCategoriesByCategoryIdRequestWithDefaults instantiates a new PatchApiChartsByChartIdCategoriesByCategoryIdRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *PatchApiChartsByChartIdCategoriesByCategoryIdRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PatchApiChartsByChartIdCategoriesByCategoryIdRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PatchApiChartsByChartIdCategoriesByCategoryIdRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PatchApiChartsByChartIdCategoriesByCategoryIdRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetColor

`func (o *PatchApiChartsByChartIdCategoriesByCategoryIdRequest) GetColor() string`

GetColor returns the Color field if non-nil, zero value otherwise.

### GetColorOk

`func (o *PatchApiChartsByChartIdCategoriesByCategoryIdRequest) GetColorOk() (*string, bool)`

GetColorOk returns a tuple with the Color field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColor

`func (o *PatchApiChartsByChartIdCategoriesByCategoryIdRequest) SetColor(v string)`

SetColor sets Color field to given value.

### HasColor

`func (o *PatchApiChartsByChartIdCategoriesByCategoryIdRequest) HasColor() bool`

HasColor returns a boolean if a field has been set.

### GetTicketTypes

`func (o *PatchApiChartsByChartIdCategoriesByCategoryIdRequest) GetTicketTypes() []PostApiChartsByChartIdCategoriesRequestTicketTypesInner`

GetTicketTypes returns the TicketTypes field if non-nil, zero value otherwise.

### GetTicketTypesOk

`func (o *PatchApiChartsByChartIdCategoriesByCategoryIdRequest) GetTicketTypesOk() (*[]PostApiChartsByChartIdCategoriesRequestTicketTypesInner, bool)`

GetTicketTypesOk returns a tuple with the TicketTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTicketTypes

`func (o *PatchApiChartsByChartIdCategoriesByCategoryIdRequest) SetTicketTypes(v []PostApiChartsByChartIdCategoriesRequestTicketTypesInner)`

SetTicketTypes sets TicketTypes field to given value.

### HasTicketTypes

`func (o *PatchApiChartsByChartIdCategoriesByCategoryIdRequest) HasTicketTypes() bool`

HasTicketTypes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


