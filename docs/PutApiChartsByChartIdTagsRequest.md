# PutApiChartsByChartIdTagsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Tags** | **[]string** | Full replacement set. Each tag is trimmed, inner whitespace collapsed and NFC-normalized; 1-40 characters, no commas or control characters. Duplicates are merged case-insensitively (first spelling wins). At most 20 distinct tags. Send &#x60;[]&#x60; to clear. | 

## Methods

### NewPutApiChartsByChartIdTagsRequest

`func NewPutApiChartsByChartIdTagsRequest(tags []string, ) *PutApiChartsByChartIdTagsRequest`

NewPutApiChartsByChartIdTagsRequest instantiates a new PutApiChartsByChartIdTagsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPutApiChartsByChartIdTagsRequestWithDefaults

`func NewPutApiChartsByChartIdTagsRequestWithDefaults() *PutApiChartsByChartIdTagsRequest`

NewPutApiChartsByChartIdTagsRequestWithDefaults instantiates a new PutApiChartsByChartIdTagsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTags

`func (o *PutApiChartsByChartIdTagsRequest) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *PutApiChartsByChartIdTagsRequest) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *PutApiChartsByChartIdTagsRequest) SetTags(v []string)`

SetTags sets Tags field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


