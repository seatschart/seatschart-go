# GetApiEventsByEventIdMapVersions200ResponseItemsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Version** | **int32** |  | 
**Kind** | **string** | &#x60;baseline&#x60; &#x3D; the map that was live before the first recorded publish (cloned from a chart, or published before history). | 
**GeometryHash** | **string** | SHA-256 of the canonical snapshot. Equal hashes mean identical published content. | 
**PublishedAt** | **time.Time** |  | 
**PublishedBy** | [**GetApiChartsByChartIdVersions200ResponseItemsInnerPublishedBy**](GetApiChartsByChartIdVersions200ResponseItemsInnerPublishedBy.md) |  | 
**SeatCount** | **int32** |  | 
**GaAreaCount** | **int32** |  | 
**ThumbnailUrl** | **string** | Relative API path (authenticated) of the placeholder preview of this version. | 
**IsCurrent** | **bool** | True for the newest version (the one currently published). | 

## Methods

### NewGetApiEventsByEventIdMapVersions200ResponseItemsInner

`func NewGetApiEventsByEventIdMapVersions200ResponseItemsInner(version int32, kind string, geometryHash string, publishedAt time.Time, publishedBy GetApiChartsByChartIdVersions200ResponseItemsInnerPublishedBy, seatCount int32, gaAreaCount int32, thumbnailUrl string, isCurrent bool, ) *GetApiEventsByEventIdMapVersions200ResponseItemsInner`

NewGetApiEventsByEventIdMapVersions200ResponseItemsInner instantiates a new GetApiEventsByEventIdMapVersions200ResponseItemsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetApiEventsByEventIdMapVersions200ResponseItemsInnerWithDefaults

`func NewGetApiEventsByEventIdMapVersions200ResponseItemsInnerWithDefaults() *GetApiEventsByEventIdMapVersions200ResponseItemsInner`

NewGetApiEventsByEventIdMapVersions200ResponseItemsInnerWithDefaults instantiates a new GetApiEventsByEventIdMapVersions200ResponseItemsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVersion

`func (o *GetApiEventsByEventIdMapVersions200ResponseItemsInner) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *GetApiEventsByEventIdMapVersions200ResponseItemsInner) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *GetApiEventsByEventIdMapVersions200ResponseItemsInner) SetVersion(v int32)`

SetVersion sets Version field to given value.


### GetKind

`func (o *GetApiEventsByEventIdMapVersions200ResponseItemsInner) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *GetApiEventsByEventIdMapVersions200ResponseItemsInner) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *GetApiEventsByEventIdMapVersions200ResponseItemsInner) SetKind(v string)`

SetKind sets Kind field to given value.


### GetGeometryHash

`func (o *GetApiEventsByEventIdMapVersions200ResponseItemsInner) GetGeometryHash() string`

GetGeometryHash returns the GeometryHash field if non-nil, zero value otherwise.

### GetGeometryHashOk

`func (o *GetApiEventsByEventIdMapVersions200ResponseItemsInner) GetGeometryHashOk() (*string, bool)`

GetGeometryHashOk returns a tuple with the GeometryHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeometryHash

`func (o *GetApiEventsByEventIdMapVersions200ResponseItemsInner) SetGeometryHash(v string)`

SetGeometryHash sets GeometryHash field to given value.


### GetPublishedAt

`func (o *GetApiEventsByEventIdMapVersions200ResponseItemsInner) GetPublishedAt() time.Time`

GetPublishedAt returns the PublishedAt field if non-nil, zero value otherwise.

### GetPublishedAtOk

`func (o *GetApiEventsByEventIdMapVersions200ResponseItemsInner) GetPublishedAtOk() (*time.Time, bool)`

GetPublishedAtOk returns a tuple with the PublishedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublishedAt

`func (o *GetApiEventsByEventIdMapVersions200ResponseItemsInner) SetPublishedAt(v time.Time)`

SetPublishedAt sets PublishedAt field to given value.


### GetPublishedBy

`func (o *GetApiEventsByEventIdMapVersions200ResponseItemsInner) GetPublishedBy() GetApiChartsByChartIdVersions200ResponseItemsInnerPublishedBy`

GetPublishedBy returns the PublishedBy field if non-nil, zero value otherwise.

### GetPublishedByOk

`func (o *GetApiEventsByEventIdMapVersions200ResponseItemsInner) GetPublishedByOk() (*GetApiChartsByChartIdVersions200ResponseItemsInnerPublishedBy, bool)`

GetPublishedByOk returns a tuple with the PublishedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublishedBy

`func (o *GetApiEventsByEventIdMapVersions200ResponseItemsInner) SetPublishedBy(v GetApiChartsByChartIdVersions200ResponseItemsInnerPublishedBy)`

SetPublishedBy sets PublishedBy field to given value.


### GetSeatCount

`func (o *GetApiEventsByEventIdMapVersions200ResponseItemsInner) GetSeatCount() int32`

GetSeatCount returns the SeatCount field if non-nil, zero value otherwise.

### GetSeatCountOk

`func (o *GetApiEventsByEventIdMapVersions200ResponseItemsInner) GetSeatCountOk() (*int32, bool)`

GetSeatCountOk returns a tuple with the SeatCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeatCount

`func (o *GetApiEventsByEventIdMapVersions200ResponseItemsInner) SetSeatCount(v int32)`

SetSeatCount sets SeatCount field to given value.


### GetGaAreaCount

`func (o *GetApiEventsByEventIdMapVersions200ResponseItemsInner) GetGaAreaCount() int32`

GetGaAreaCount returns the GaAreaCount field if non-nil, zero value otherwise.

### GetGaAreaCountOk

`func (o *GetApiEventsByEventIdMapVersions200ResponseItemsInner) GetGaAreaCountOk() (*int32, bool)`

GetGaAreaCountOk returns a tuple with the GaAreaCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGaAreaCount

`func (o *GetApiEventsByEventIdMapVersions200ResponseItemsInner) SetGaAreaCount(v int32)`

SetGaAreaCount sets GaAreaCount field to given value.


### GetThumbnailUrl

`func (o *GetApiEventsByEventIdMapVersions200ResponseItemsInner) GetThumbnailUrl() string`

GetThumbnailUrl returns the ThumbnailUrl field if non-nil, zero value otherwise.

### GetThumbnailUrlOk

`func (o *GetApiEventsByEventIdMapVersions200ResponseItemsInner) GetThumbnailUrlOk() (*string, bool)`

GetThumbnailUrlOk returns a tuple with the ThumbnailUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThumbnailUrl

`func (o *GetApiEventsByEventIdMapVersions200ResponseItemsInner) SetThumbnailUrl(v string)`

SetThumbnailUrl sets ThumbnailUrl field to given value.


### GetIsCurrent

`func (o *GetApiEventsByEventIdMapVersions200ResponseItemsInner) GetIsCurrent() bool`

GetIsCurrent returns the IsCurrent field if non-nil, zero value otherwise.

### GetIsCurrentOk

`func (o *GetApiEventsByEventIdMapVersions200ResponseItemsInner) GetIsCurrentOk() (*bool, bool)`

GetIsCurrentOk returns a tuple with the IsCurrent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsCurrent

`func (o *GetApiEventsByEventIdMapVersions200ResponseItemsInner) SetIsCurrent(v bool)`

SetIsCurrent sets IsCurrent field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


