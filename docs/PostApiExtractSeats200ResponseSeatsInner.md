# PostApiExtractSeats200ResponseSeatsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**X** | **float32** |  | 
**Y** | **float32** |  | 
**Confidence** | Pointer to **float32** | How sure the scanner is that this is a seat (auto and examples modes). | [optional] 
**Doubtful** | Pointer to **bool** | Present (true) on seats to review before keeping them. | [optional] 

## Methods

### NewPostApiExtractSeats200ResponseSeatsInner

`func NewPostApiExtractSeats200ResponseSeatsInner(x float32, y float32, ) *PostApiExtractSeats200ResponseSeatsInner`

NewPostApiExtractSeats200ResponseSeatsInner instantiates a new PostApiExtractSeats200ResponseSeatsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostApiExtractSeats200ResponseSeatsInnerWithDefaults

`func NewPostApiExtractSeats200ResponseSeatsInnerWithDefaults() *PostApiExtractSeats200ResponseSeatsInner`

NewPostApiExtractSeats200ResponseSeatsInnerWithDefaults instantiates a new PostApiExtractSeats200ResponseSeatsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetX

`func (o *PostApiExtractSeats200ResponseSeatsInner) GetX() float32`

GetX returns the X field if non-nil, zero value otherwise.

### GetXOk

`func (o *PostApiExtractSeats200ResponseSeatsInner) GetXOk() (*float32, bool)`

GetXOk returns a tuple with the X field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetX

`func (o *PostApiExtractSeats200ResponseSeatsInner) SetX(v float32)`

SetX sets X field to given value.


### GetY

`func (o *PostApiExtractSeats200ResponseSeatsInner) GetY() float32`

GetY returns the Y field if non-nil, zero value otherwise.

### GetYOk

`func (o *PostApiExtractSeats200ResponseSeatsInner) GetYOk() (*float32, bool)`

GetYOk returns a tuple with the Y field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetY

`func (o *PostApiExtractSeats200ResponseSeatsInner) SetY(v float32)`

SetY sets Y field to given value.


### GetConfidence

`func (o *PostApiExtractSeats200ResponseSeatsInner) GetConfidence() float32`

GetConfidence returns the Confidence field if non-nil, zero value otherwise.

### GetConfidenceOk

`func (o *PostApiExtractSeats200ResponseSeatsInner) GetConfidenceOk() (*float32, bool)`

GetConfidenceOk returns a tuple with the Confidence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfidence

`func (o *PostApiExtractSeats200ResponseSeatsInner) SetConfidence(v float32)`

SetConfidence sets Confidence field to given value.

### HasConfidence

`func (o *PostApiExtractSeats200ResponseSeatsInner) HasConfidence() bool`

HasConfidence returns a boolean if a field has been set.

### GetDoubtful

`func (o *PostApiExtractSeats200ResponseSeatsInner) GetDoubtful() bool`

GetDoubtful returns the Doubtful field if non-nil, zero value otherwise.

### GetDoubtfulOk

`func (o *PostApiExtractSeats200ResponseSeatsInner) GetDoubtfulOk() (*bool, bool)`

GetDoubtfulOk returns a tuple with the Doubtful field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDoubtful

`func (o *PostApiExtractSeats200ResponseSeatsInner) SetDoubtful(v bool)`

SetDoubtful sets Doubtful field to given value.

### HasDoubtful

`func (o *PostApiExtractSeats200ResponseSeatsInner) HasDoubtful() bool`

HasDoubtful returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


