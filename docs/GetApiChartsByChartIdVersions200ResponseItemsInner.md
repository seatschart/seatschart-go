# GetApiChartsByChartIdVersions200ResponseItemsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Version** | **int32** |  | 
**Kind** | **string** | &#x60;baseline&#x60; &#x3D; the map that was live before version history existed (recorded on the first publish after upgrade). | 
**GeometryHash** | **string** | SHA-256 of the canonical snapshot. Equal hashes mean identical published content. | 
**PublishedAt** | **time.Time** |  | 
**PublishedBy** | [**GetApiChartsByChartIdVersions200ResponseItemsInnerPublishedBy**](GetApiChartsByChartIdVersions200ResponseItemsInnerPublishedBy.md) |  | 
**SeatCount** | **int32** |  | 
**GaAreaCount** | **int32** |  | 
**Note** | **NullableString** |  | 
**HasThumbnail** | **bool** | True when the designer uploaded a rendered thumbnail; otherwise the thumbnail URL serves a placeholder. | 
**ThumbnailUrl** | **string** | Relative API path (authenticated) of the immutable thumbnail for this version. | 
**IsCurrent** | **bool** | True for the newest version (the one currently published). | 

## Methods

### NewGetApiChartsByChartIdVersions200ResponseItemsInner

`func NewGetApiChartsByChartIdVersions200ResponseItemsInner(version int32, kind string, geometryHash string, publishedAt time.Time, publishedBy GetApiChartsByChartIdVersions200ResponseItemsInnerPublishedBy, seatCount int32, gaAreaCount int32, note NullableString, hasThumbnail bool, thumbnailUrl string, isCurrent bool, ) *GetApiChartsByChartIdVersions200ResponseItemsInner`

NewGetApiChartsByChartIdVersions200ResponseItemsInner instantiates a new GetApiChartsByChartIdVersions200ResponseItemsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetApiChartsByChartIdVersions200ResponseItemsInnerWithDefaults

`func NewGetApiChartsByChartIdVersions200ResponseItemsInnerWithDefaults() *GetApiChartsByChartIdVersions200ResponseItemsInner`

NewGetApiChartsByChartIdVersions200ResponseItemsInnerWithDefaults instantiates a new GetApiChartsByChartIdVersions200ResponseItemsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVersion

`func (o *GetApiChartsByChartIdVersions200ResponseItemsInner) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *GetApiChartsByChartIdVersions200ResponseItemsInner) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *GetApiChartsByChartIdVersions200ResponseItemsInner) SetVersion(v int32)`

SetVersion sets Version field to given value.


### GetKind

`func (o *GetApiChartsByChartIdVersions200ResponseItemsInner) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *GetApiChartsByChartIdVersions200ResponseItemsInner) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *GetApiChartsByChartIdVersions200ResponseItemsInner) SetKind(v string)`

SetKind sets Kind field to given value.


### GetGeometryHash

`func (o *GetApiChartsByChartIdVersions200ResponseItemsInner) GetGeometryHash() string`

GetGeometryHash returns the GeometryHash field if non-nil, zero value otherwise.

### GetGeometryHashOk

`func (o *GetApiChartsByChartIdVersions200ResponseItemsInner) GetGeometryHashOk() (*string, bool)`

GetGeometryHashOk returns a tuple with the GeometryHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeometryHash

`func (o *GetApiChartsByChartIdVersions200ResponseItemsInner) SetGeometryHash(v string)`

SetGeometryHash sets GeometryHash field to given value.


### GetPublishedAt

`func (o *GetApiChartsByChartIdVersions200ResponseItemsInner) GetPublishedAt() time.Time`

GetPublishedAt returns the PublishedAt field if non-nil, zero value otherwise.

### GetPublishedAtOk

`func (o *GetApiChartsByChartIdVersions200ResponseItemsInner) GetPublishedAtOk() (*time.Time, bool)`

GetPublishedAtOk returns a tuple with the PublishedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublishedAt

`func (o *GetApiChartsByChartIdVersions200ResponseItemsInner) SetPublishedAt(v time.Time)`

SetPublishedAt sets PublishedAt field to given value.


### GetPublishedBy

`func (o *GetApiChartsByChartIdVersions200ResponseItemsInner) GetPublishedBy() GetApiChartsByChartIdVersions200ResponseItemsInnerPublishedBy`

GetPublishedBy returns the PublishedBy field if non-nil, zero value otherwise.

### GetPublishedByOk

`func (o *GetApiChartsByChartIdVersions200ResponseItemsInner) GetPublishedByOk() (*GetApiChartsByChartIdVersions200ResponseItemsInnerPublishedBy, bool)`

GetPublishedByOk returns a tuple with the PublishedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublishedBy

`func (o *GetApiChartsByChartIdVersions200ResponseItemsInner) SetPublishedBy(v GetApiChartsByChartIdVersions200ResponseItemsInnerPublishedBy)`

SetPublishedBy sets PublishedBy field to given value.


### GetSeatCount

`func (o *GetApiChartsByChartIdVersions200ResponseItemsInner) GetSeatCount() int32`

GetSeatCount returns the SeatCount field if non-nil, zero value otherwise.

### GetSeatCountOk

`func (o *GetApiChartsByChartIdVersions200ResponseItemsInner) GetSeatCountOk() (*int32, bool)`

GetSeatCountOk returns a tuple with the SeatCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeatCount

`func (o *GetApiChartsByChartIdVersions200ResponseItemsInner) SetSeatCount(v int32)`

SetSeatCount sets SeatCount field to given value.


### GetGaAreaCount

`func (o *GetApiChartsByChartIdVersions200ResponseItemsInner) GetGaAreaCount() int32`

GetGaAreaCount returns the GaAreaCount field if non-nil, zero value otherwise.

### GetGaAreaCountOk

`func (o *GetApiChartsByChartIdVersions200ResponseItemsInner) GetGaAreaCountOk() (*int32, bool)`

GetGaAreaCountOk returns a tuple with the GaAreaCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGaAreaCount

`func (o *GetApiChartsByChartIdVersions200ResponseItemsInner) SetGaAreaCount(v int32)`

SetGaAreaCount sets GaAreaCount field to given value.


### GetNote

`func (o *GetApiChartsByChartIdVersions200ResponseItemsInner) GetNote() string`

GetNote returns the Note field if non-nil, zero value otherwise.

### GetNoteOk

`func (o *GetApiChartsByChartIdVersions200ResponseItemsInner) GetNoteOk() (*string, bool)`

GetNoteOk returns a tuple with the Note field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNote

`func (o *GetApiChartsByChartIdVersions200ResponseItemsInner) SetNote(v string)`

SetNote sets Note field to given value.


### SetNoteNil

`func (o *GetApiChartsByChartIdVersions200ResponseItemsInner) SetNoteNil(b bool)`

 SetNoteNil sets the value for Note to be an explicit nil

### UnsetNote
`func (o *GetApiChartsByChartIdVersions200ResponseItemsInner) UnsetNote()`

UnsetNote ensures that no value is present for Note, not even an explicit nil
### GetHasThumbnail

`func (o *GetApiChartsByChartIdVersions200ResponseItemsInner) GetHasThumbnail() bool`

GetHasThumbnail returns the HasThumbnail field if non-nil, zero value otherwise.

### GetHasThumbnailOk

`func (o *GetApiChartsByChartIdVersions200ResponseItemsInner) GetHasThumbnailOk() (*bool, bool)`

GetHasThumbnailOk returns a tuple with the HasThumbnail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasThumbnail

`func (o *GetApiChartsByChartIdVersions200ResponseItemsInner) SetHasThumbnail(v bool)`

SetHasThumbnail sets HasThumbnail field to given value.


### GetThumbnailUrl

`func (o *GetApiChartsByChartIdVersions200ResponseItemsInner) GetThumbnailUrl() string`

GetThumbnailUrl returns the ThumbnailUrl field if non-nil, zero value otherwise.

### GetThumbnailUrlOk

`func (o *GetApiChartsByChartIdVersions200ResponseItemsInner) GetThumbnailUrlOk() (*string, bool)`

GetThumbnailUrlOk returns a tuple with the ThumbnailUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThumbnailUrl

`func (o *GetApiChartsByChartIdVersions200ResponseItemsInner) SetThumbnailUrl(v string)`

SetThumbnailUrl sets ThumbnailUrl field to given value.


### GetIsCurrent

`func (o *GetApiChartsByChartIdVersions200ResponseItemsInner) GetIsCurrent() bool`

GetIsCurrent returns the IsCurrent field if non-nil, zero value otherwise.

### GetIsCurrentOk

`func (o *GetApiChartsByChartIdVersions200ResponseItemsInner) GetIsCurrentOk() (*bool, bool)`

GetIsCurrentOk returns a tuple with the IsCurrent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsCurrent

`func (o *GetApiChartsByChartIdVersions200ResponseItemsInner) SetIsCurrent(v bool)`

SetIsCurrent sets IsCurrent field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


