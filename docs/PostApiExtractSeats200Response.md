# PostApiExtractSeats200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Width** | **int32** |  | 
**Height** | **int32** |  | 
**Seats** | [**[]PostApiExtractSeats200ResponseSeatsInner**](PostApiExtractSeats200ResponseSeatsInner.md) |  | 
**Mode** | Pointer to **string** |  | [optional] 
**Legibility** | Pointer to **string** | &#x60;ok&#x60;; &#x60;low_resolution&#x60;: the seats are drawn too small to be read one by one; &#x60;no_individual_seats&#x60;: the plan shows sections or blocks, not seats. When not &#x60;ok&#x60;, every seat is &#x60;doubtful&#x60; and the count must not be trusted. | [optional] 
**Review** | Pointer to **int32** | How many seats are flagged &#x60;doubtful&#x60;. | [optional] 
**Vector** | Pointer to **string** | Only when a &#x60;vector&#x60; file was sent: &#x60;used&#x60; when its drawn geometry checked the seats, &#x60;unusable&#x60; when it held none (a scanned page, no seat-sized shapes, over the work budget) and the image alone was scanned. | [optional] 
**Suggested** | Pointer to [**[]PostApiExtractSeats200ResponseSuggestedInner**](PostApiExtractSeats200ResponseSuggestedInner.md) | Auto and examples modes: places where a seat was probably missed (a gap in a row, the end of a row, a mark that looks like the clicked seats but scored below the threshold), best first. They are not part of &#x60;seats&#x60;; the designer offers them in the review, where one key adds one. Omitted when there are none. | [optional] 

## Methods

### NewPostApiExtractSeats200Response

`func NewPostApiExtractSeats200Response(width int32, height int32, seats []PostApiExtractSeats200ResponseSeatsInner, ) *PostApiExtractSeats200Response`

NewPostApiExtractSeats200Response instantiates a new PostApiExtractSeats200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostApiExtractSeats200ResponseWithDefaults

`func NewPostApiExtractSeats200ResponseWithDefaults() *PostApiExtractSeats200Response`

NewPostApiExtractSeats200ResponseWithDefaults instantiates a new PostApiExtractSeats200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWidth

`func (o *PostApiExtractSeats200Response) GetWidth() int32`

GetWidth returns the Width field if non-nil, zero value otherwise.

### GetWidthOk

`func (o *PostApiExtractSeats200Response) GetWidthOk() (*int32, bool)`

GetWidthOk returns a tuple with the Width field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWidth

`func (o *PostApiExtractSeats200Response) SetWidth(v int32)`

SetWidth sets Width field to given value.


### GetHeight

`func (o *PostApiExtractSeats200Response) GetHeight() int32`

GetHeight returns the Height field if non-nil, zero value otherwise.

### GetHeightOk

`func (o *PostApiExtractSeats200Response) GetHeightOk() (*int32, bool)`

GetHeightOk returns a tuple with the Height field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeight

`func (o *PostApiExtractSeats200Response) SetHeight(v int32)`

SetHeight sets Height field to given value.


### GetSeats

`func (o *PostApiExtractSeats200Response) GetSeats() []PostApiExtractSeats200ResponseSeatsInner`

GetSeats returns the Seats field if non-nil, zero value otherwise.

### GetSeatsOk

`func (o *PostApiExtractSeats200Response) GetSeatsOk() (*[]PostApiExtractSeats200ResponseSeatsInner, bool)`

GetSeatsOk returns a tuple with the Seats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeats

`func (o *PostApiExtractSeats200Response) SetSeats(v []PostApiExtractSeats200ResponseSeatsInner)`

SetSeats sets Seats field to given value.


### GetMode

`func (o *PostApiExtractSeats200Response) GetMode() string`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *PostApiExtractSeats200Response) GetModeOk() (*string, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *PostApiExtractSeats200Response) SetMode(v string)`

SetMode sets Mode field to given value.

### HasMode

`func (o *PostApiExtractSeats200Response) HasMode() bool`

HasMode returns a boolean if a field has been set.

### GetLegibility

`func (o *PostApiExtractSeats200Response) GetLegibility() string`

GetLegibility returns the Legibility field if non-nil, zero value otherwise.

### GetLegibilityOk

`func (o *PostApiExtractSeats200Response) GetLegibilityOk() (*string, bool)`

GetLegibilityOk returns a tuple with the Legibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLegibility

`func (o *PostApiExtractSeats200Response) SetLegibility(v string)`

SetLegibility sets Legibility field to given value.

### HasLegibility

`func (o *PostApiExtractSeats200Response) HasLegibility() bool`

HasLegibility returns a boolean if a field has been set.

### GetReview

`func (o *PostApiExtractSeats200Response) GetReview() int32`

GetReview returns the Review field if non-nil, zero value otherwise.

### GetReviewOk

`func (o *PostApiExtractSeats200Response) GetReviewOk() (*int32, bool)`

GetReviewOk returns a tuple with the Review field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReview

`func (o *PostApiExtractSeats200Response) SetReview(v int32)`

SetReview sets Review field to given value.

### HasReview

`func (o *PostApiExtractSeats200Response) HasReview() bool`

HasReview returns a boolean if a field has been set.

### GetVector

`func (o *PostApiExtractSeats200Response) GetVector() string`

GetVector returns the Vector field if non-nil, zero value otherwise.

### GetVectorOk

`func (o *PostApiExtractSeats200Response) GetVectorOk() (*string, bool)`

GetVectorOk returns a tuple with the Vector field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVector

`func (o *PostApiExtractSeats200Response) SetVector(v string)`

SetVector sets Vector field to given value.

### HasVector

`func (o *PostApiExtractSeats200Response) HasVector() bool`

HasVector returns a boolean if a field has been set.

### GetSuggested

`func (o *PostApiExtractSeats200Response) GetSuggested() []PostApiExtractSeats200ResponseSuggestedInner`

GetSuggested returns the Suggested field if non-nil, zero value otherwise.

### GetSuggestedOk

`func (o *PostApiExtractSeats200Response) GetSuggestedOk() (*[]PostApiExtractSeats200ResponseSuggestedInner, bool)`

GetSuggestedOk returns a tuple with the Suggested field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuggested

`func (o *PostApiExtractSeats200Response) SetSuggested(v []PostApiExtractSeats200ResponseSuggestedInner)`

SetSuggested sets Suggested field to given value.

### HasSuggested

`func (o *PostApiExtractSeats200Response) HasSuggested() bool`

HasSuggested returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


