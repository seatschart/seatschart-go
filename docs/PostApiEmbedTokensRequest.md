# PostApiEmbedTokensRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Purpose** | **string** |  | 
**ChartId** | Pointer to **string** |  | [optional] 
**EventId** | Pointer to **string** |  | [optional] 
**Mode** | Pointer to **string** | Designer tokens only, enforced by the API (default &#x60;normal&#x60;, full edit). &#x60;safe&#x60;: saves, publishes and version restores that would delete or relabel published objects, change the booking type of a published table or area, or change the capacity of a published area are refused with 403 &#x60;forbidden&#x60; / &#x60;ERR_DESIGNER_SAFE_MODE&#x60; (&#x60;reason&#x60;, &#x60;objectIds&#x60;). &#x60;readOnly&#x60;: every write is refused with 403 &#x60;forbidden&#x60; / &#x60;ERR_DESIGNER_READ_ONLY&#x60;. | [optional] 
**SafeModeOptions** | Pointer to [**PostApiEmbedTokensRequestSafeModeOptions**](PostApiEmbedTokensRequestSafeModeOptions.md) |  | [optional] 
**TtlSeconds** | Pointer to **int32** |  | [optional] 

## Methods

### NewPostApiEmbedTokensRequest

`func NewPostApiEmbedTokensRequest(purpose string, ) *PostApiEmbedTokensRequest`

NewPostApiEmbedTokensRequest instantiates a new PostApiEmbedTokensRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostApiEmbedTokensRequestWithDefaults

`func NewPostApiEmbedTokensRequestWithDefaults() *PostApiEmbedTokensRequest`

NewPostApiEmbedTokensRequestWithDefaults instantiates a new PostApiEmbedTokensRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPurpose

`func (o *PostApiEmbedTokensRequest) GetPurpose() string`

GetPurpose returns the Purpose field if non-nil, zero value otherwise.

### GetPurposeOk

`func (o *PostApiEmbedTokensRequest) GetPurposeOk() (*string, bool)`

GetPurposeOk returns a tuple with the Purpose field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurpose

`func (o *PostApiEmbedTokensRequest) SetPurpose(v string)`

SetPurpose sets Purpose field to given value.


### GetChartId

`func (o *PostApiEmbedTokensRequest) GetChartId() string`

GetChartId returns the ChartId field if non-nil, zero value otherwise.

### GetChartIdOk

`func (o *PostApiEmbedTokensRequest) GetChartIdOk() (*string, bool)`

GetChartIdOk returns a tuple with the ChartId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChartId

`func (o *PostApiEmbedTokensRequest) SetChartId(v string)`

SetChartId sets ChartId field to given value.

### HasChartId

`func (o *PostApiEmbedTokensRequest) HasChartId() bool`

HasChartId returns a boolean if a field has been set.

### GetEventId

`func (o *PostApiEmbedTokensRequest) GetEventId() string`

GetEventId returns the EventId field if non-nil, zero value otherwise.

### GetEventIdOk

`func (o *PostApiEmbedTokensRequest) GetEventIdOk() (*string, bool)`

GetEventIdOk returns a tuple with the EventId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventId

`func (o *PostApiEmbedTokensRequest) SetEventId(v string)`

SetEventId sets EventId field to given value.

### HasEventId

`func (o *PostApiEmbedTokensRequest) HasEventId() bool`

HasEventId returns a boolean if a field has been set.

### GetMode

`func (o *PostApiEmbedTokensRequest) GetMode() string`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *PostApiEmbedTokensRequest) GetModeOk() (*string, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *PostApiEmbedTokensRequest) SetMode(v string)`

SetMode sets Mode field to given value.

### HasMode

`func (o *PostApiEmbedTokensRequest) HasMode() bool`

HasMode returns a boolean if a field has been set.

### GetSafeModeOptions

`func (o *PostApiEmbedTokensRequest) GetSafeModeOptions() PostApiEmbedTokensRequestSafeModeOptions`

GetSafeModeOptions returns the SafeModeOptions field if non-nil, zero value otherwise.

### GetSafeModeOptionsOk

`func (o *PostApiEmbedTokensRequest) GetSafeModeOptionsOk() (*PostApiEmbedTokensRequestSafeModeOptions, bool)`

GetSafeModeOptionsOk returns a tuple with the SafeModeOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSafeModeOptions

`func (o *PostApiEmbedTokensRequest) SetSafeModeOptions(v PostApiEmbedTokensRequestSafeModeOptions)`

SetSafeModeOptions sets SafeModeOptions field to given value.

### HasSafeModeOptions

`func (o *PostApiEmbedTokensRequest) HasSafeModeOptions() bool`

HasSafeModeOptions returns a boolean if a field has been set.

### GetTtlSeconds

`func (o *PostApiEmbedTokensRequest) GetTtlSeconds() int32`

GetTtlSeconds returns the TtlSeconds field if non-nil, zero value otherwise.

### GetTtlSecondsOk

`func (o *PostApiEmbedTokensRequest) GetTtlSecondsOk() (*int32, bool)`

GetTtlSecondsOk returns a tuple with the TtlSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTtlSeconds

`func (o *PostApiEmbedTokensRequest) SetTtlSeconds(v int32)`

SetTtlSeconds sets TtlSeconds field to given value.

### HasTtlSeconds

`func (o *PostApiEmbedTokensRequest) HasTtlSeconds() bool`

HasTtlSeconds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


