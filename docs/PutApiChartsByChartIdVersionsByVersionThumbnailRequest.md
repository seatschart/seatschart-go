# PutApiChartsByChartIdVersionsByVersionThumbnailRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Image** | **string** | Data URL (&#x60;data:image/webp;base64,...&#x60; preferred, PNG accepted) of at most 512 KB decoded, ~640px wide. | 

## Methods

### NewPutApiChartsByChartIdVersionsByVersionThumbnailRequest

`func NewPutApiChartsByChartIdVersionsByVersionThumbnailRequest(image string, ) *PutApiChartsByChartIdVersionsByVersionThumbnailRequest`

NewPutApiChartsByChartIdVersionsByVersionThumbnailRequest instantiates a new PutApiChartsByChartIdVersionsByVersionThumbnailRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPutApiChartsByChartIdVersionsByVersionThumbnailRequestWithDefaults

`func NewPutApiChartsByChartIdVersionsByVersionThumbnailRequestWithDefaults() *PutApiChartsByChartIdVersionsByVersionThumbnailRequest`

NewPutApiChartsByChartIdVersionsByVersionThumbnailRequestWithDefaults instantiates a new PutApiChartsByChartIdVersionsByVersionThumbnailRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetImage

`func (o *PutApiChartsByChartIdVersionsByVersionThumbnailRequest) GetImage() string`

GetImage returns the Image field if non-nil, zero value otherwise.

### GetImageOk

`func (o *PutApiChartsByChartIdVersionsByVersionThumbnailRequest) GetImageOk() (*string, bool)`

GetImageOk returns a tuple with the Image field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImage

`func (o *PutApiChartsByChartIdVersionsByVersionThumbnailRequest) SetImage(v string)`

SetImage sets Image field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


