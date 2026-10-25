# PostApiChartsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Key** | Pointer to **string** | Partner key, unique in the workspace (409 &#x60;chart_key_taken&#x60;), accepted wherever a chart id is. Never UUID-shaped. | [optional] 
**VenueType** | Pointer to **string** | Designer template: &#x60;simple&#x60; (rows), &#x60;sections&#x60; (sections and floors) or &#x60;zones&#x60;. Default &#x60;simple&#x60;. | [optional] 
**Categories** | Pointer to [**[]PostApiChartsRequestCategoriesInner**](PostApiChartsRequestCategoriesInner.md) |  | [optional] 

## Methods

### NewPostApiChartsRequest

`func NewPostApiChartsRequest(name string, ) *PostApiChartsRequest`

NewPostApiChartsRequest instantiates a new PostApiChartsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostApiChartsRequestWithDefaults

`func NewPostApiChartsRequestWithDefaults() *PostApiChartsRequest`

NewPostApiChartsRequestWithDefaults instantiates a new PostApiChartsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *PostApiChartsRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PostApiChartsRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PostApiChartsRequest) SetName(v string)`

SetName sets Name field to given value.


### GetKey

`func (o *PostApiChartsRequest) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *PostApiChartsRequest) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *PostApiChartsRequest) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *PostApiChartsRequest) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetVenueType

`func (o *PostApiChartsRequest) GetVenueType() string`

GetVenueType returns the VenueType field if non-nil, zero value otherwise.

### GetVenueTypeOk

`func (o *PostApiChartsRequest) GetVenueTypeOk() (*string, bool)`

GetVenueTypeOk returns a tuple with the VenueType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVenueType

`func (o *PostApiChartsRequest) SetVenueType(v string)`

SetVenueType sets VenueType field to given value.

### HasVenueType

`func (o *PostApiChartsRequest) HasVenueType() bool`

HasVenueType returns a boolean if a field has been set.

### GetCategories

`func (o *PostApiChartsRequest) GetCategories() []PostApiChartsRequestCategoriesInner`

GetCategories returns the Categories field if non-nil, zero value otherwise.

### GetCategoriesOk

`func (o *PostApiChartsRequest) GetCategoriesOk() (*[]PostApiChartsRequestCategoriesInner, bool)`

GetCategoriesOk returns a tuple with the Categories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategories

`func (o *PostApiChartsRequest) SetCategories(v []PostApiChartsRequestCategoriesInner)`

SetCategories sets Categories field to given value.

### HasCategories

`func (o *PostApiChartsRequest) HasCategories() bool`

HasCategories returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


