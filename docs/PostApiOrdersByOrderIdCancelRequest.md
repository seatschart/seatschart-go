# PostApiOrdersByOrderIdCancelRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Reason** | **string** |  | 
**SeatIds** | Pointer to **[]string** | Cancel only these seats of the order (1 to 64 seat ids of its items). Omitted: the whole order. | [optional] 

## Methods

### NewPostApiOrdersByOrderIdCancelRequest

`func NewPostApiOrdersByOrderIdCancelRequest(reason string, ) *PostApiOrdersByOrderIdCancelRequest`

NewPostApiOrdersByOrderIdCancelRequest instantiates a new PostApiOrdersByOrderIdCancelRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostApiOrdersByOrderIdCancelRequestWithDefaults

`func NewPostApiOrdersByOrderIdCancelRequestWithDefaults() *PostApiOrdersByOrderIdCancelRequest`

NewPostApiOrdersByOrderIdCancelRequestWithDefaults instantiates a new PostApiOrdersByOrderIdCancelRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReason

`func (o *PostApiOrdersByOrderIdCancelRequest) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *PostApiOrdersByOrderIdCancelRequest) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *PostApiOrdersByOrderIdCancelRequest) SetReason(v string)`

SetReason sets Reason field to given value.


### GetSeatIds

`func (o *PostApiOrdersByOrderIdCancelRequest) GetSeatIds() []string`

GetSeatIds returns the SeatIds field if non-nil, zero value otherwise.

### GetSeatIdsOk

`func (o *PostApiOrdersByOrderIdCancelRequest) GetSeatIdsOk() (*[]string, bool)`

GetSeatIdsOk returns a tuple with the SeatIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeatIds

`func (o *PostApiOrdersByOrderIdCancelRequest) SetSeatIds(v []string)`

SetSeatIds sets SeatIds field to given value.

### HasSeatIds

`func (o *PostApiOrdersByOrderIdCancelRequest) HasSeatIds() bool`

HasSeatIds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


