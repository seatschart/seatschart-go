# PostApiEventsByEventIdListingsRequestObjectsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ObjectId** | **string** | Seat id or label, or general-admission area id or name (with quantity). | 
**Quantity** | Pointer to **int32** | General-admission places (required for an area). | [optional] 
**PriceCents** | Pointer to **int32** |  | [optional] 

## Methods

### NewPostApiEventsByEventIdListingsRequestObjectsInner

`func NewPostApiEventsByEventIdListingsRequestObjectsInner(objectId string, ) *PostApiEventsByEventIdListingsRequestObjectsInner`

NewPostApiEventsByEventIdListingsRequestObjectsInner instantiates a new PostApiEventsByEventIdListingsRequestObjectsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostApiEventsByEventIdListingsRequestObjectsInnerWithDefaults

`func NewPostApiEventsByEventIdListingsRequestObjectsInnerWithDefaults() *PostApiEventsByEventIdListingsRequestObjectsInner`

NewPostApiEventsByEventIdListingsRequestObjectsInnerWithDefaults instantiates a new PostApiEventsByEventIdListingsRequestObjectsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetObjectId

`func (o *PostApiEventsByEventIdListingsRequestObjectsInner) GetObjectId() string`

GetObjectId returns the ObjectId field if non-nil, zero value otherwise.

### GetObjectIdOk

`func (o *PostApiEventsByEventIdListingsRequestObjectsInner) GetObjectIdOk() (*string, bool)`

GetObjectIdOk returns a tuple with the ObjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectId

`func (o *PostApiEventsByEventIdListingsRequestObjectsInner) SetObjectId(v string)`

SetObjectId sets ObjectId field to given value.


### GetQuantity

`func (o *PostApiEventsByEventIdListingsRequestObjectsInner) GetQuantity() int32`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *PostApiEventsByEventIdListingsRequestObjectsInner) GetQuantityOk() (*int32, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *PostApiEventsByEventIdListingsRequestObjectsInner) SetQuantity(v int32)`

SetQuantity sets Quantity field to given value.

### HasQuantity

`func (o *PostApiEventsByEventIdListingsRequestObjectsInner) HasQuantity() bool`

HasQuantity returns a boolean if a field has been set.

### GetPriceCents

`func (o *PostApiEventsByEventIdListingsRequestObjectsInner) GetPriceCents() int32`

GetPriceCents returns the PriceCents field if non-nil, zero value otherwise.

### GetPriceCentsOk

`func (o *PostApiEventsByEventIdListingsRequestObjectsInner) GetPriceCentsOk() (*int32, bool)`

GetPriceCentsOk returns a tuple with the PriceCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceCents

`func (o *PostApiEventsByEventIdListingsRequestObjectsInner) SetPriceCents(v int32)`

SetPriceCents sets PriceCents field to given value.

### HasPriceCents

`func (o *PostApiEventsByEventIdListingsRequestObjectsInner) HasPriceCents() bool`

HasPriceCents returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


