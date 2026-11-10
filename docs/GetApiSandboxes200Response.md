# GetApiSandboxes200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Items** | [**[]GetApiSandboxes200ResponseItemsInner**](GetApiSandboxes200ResponseItemsInner.md) |  | 
**Limit** | **int32** |  | 

## Methods

### NewGetApiSandboxes200Response

`func NewGetApiSandboxes200Response(items []GetApiSandboxes200ResponseItemsInner, limit int32, ) *GetApiSandboxes200Response`

NewGetApiSandboxes200Response instantiates a new GetApiSandboxes200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetApiSandboxes200ResponseWithDefaults

`func NewGetApiSandboxes200ResponseWithDefaults() *GetApiSandboxes200Response`

NewGetApiSandboxes200ResponseWithDefaults instantiates a new GetApiSandboxes200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItems

`func (o *GetApiSandboxes200Response) GetItems() []GetApiSandboxes200ResponseItemsInner`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *GetApiSandboxes200Response) GetItemsOk() (*[]GetApiSandboxes200ResponseItemsInner, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *GetApiSandboxes200Response) SetItems(v []GetApiSandboxes200ResponseItemsInner)`

SetItems sets Items field to given value.


### GetLimit

`func (o *GetApiSandboxes200Response) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *GetApiSandboxes200Response) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *GetApiSandboxes200Response) SetLimit(v int32)`

SetLimit sets Limit field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


