# PostApiChartsByChartIdMapPublishRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Note** | Pointer to **string** | Optional release note stored on the version record (shown in version history). | [optional] 
**AutoPublish** | Pointer to **bool** | Set by the designer auto-publish option. The publish is refused with 409 &#x60;ERR_AUTO_PUBLISH_CHART_HAS_EVENTS&#x60; (params.eventCount) when any event uses the chart. | [optional] 

## Methods

### NewPostApiChartsByChartIdMapPublishRequest

`func NewPostApiChartsByChartIdMapPublishRequest() *PostApiChartsByChartIdMapPublishRequest`

NewPostApiChartsByChartIdMapPublishRequest instantiates a new PostApiChartsByChartIdMapPublishRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostApiChartsByChartIdMapPublishRequestWithDefaults

`func NewPostApiChartsByChartIdMapPublishRequestWithDefaults() *PostApiChartsByChartIdMapPublishRequest`

NewPostApiChartsByChartIdMapPublishRequestWithDefaults instantiates a new PostApiChartsByChartIdMapPublishRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNote

`func (o *PostApiChartsByChartIdMapPublishRequest) GetNote() string`

GetNote returns the Note field if non-nil, zero value otherwise.

### GetNoteOk

`func (o *PostApiChartsByChartIdMapPublishRequest) GetNoteOk() (*string, bool)`

GetNoteOk returns a tuple with the Note field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNote

`func (o *PostApiChartsByChartIdMapPublishRequest) SetNote(v string)`

SetNote sets Note field to given value.

### HasNote

`func (o *PostApiChartsByChartIdMapPublishRequest) HasNote() bool`

HasNote returns a boolean if a field has been set.

### GetAutoPublish

`func (o *PostApiChartsByChartIdMapPublishRequest) GetAutoPublish() bool`

GetAutoPublish returns the AutoPublish field if non-nil, zero value otherwise.

### GetAutoPublishOk

`func (o *PostApiChartsByChartIdMapPublishRequest) GetAutoPublishOk() (*bool, bool)`

GetAutoPublishOk returns a tuple with the AutoPublish field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoPublish

`func (o *PostApiChartsByChartIdMapPublishRequest) SetAutoPublish(v bool)`

SetAutoPublish sets AutoPublish field to given value.

### HasAutoPublish

`func (o *PostApiChartsByChartIdMapPublishRequest) HasAutoPublish() bool`

HasAutoPublish returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


