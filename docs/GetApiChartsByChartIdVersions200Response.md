# GetApiChartsByChartIdVersions200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Items** | [**[]GetApiChartsByChartIdVersions200ResponseItemsInner**](GetApiChartsByChartIdVersions200ResponseItemsInner.md) |  | 
**NextCursor** | **NullableString** |  | 
**CurrentVersion** | **NullableInt32** |  | 

## Methods

### NewGetApiChartsByChartIdVersions200Response

`func NewGetApiChartsByChartIdVersions200Response(items []GetApiChartsByChartIdVersions200ResponseItemsInner, nextCursor NullableString, currentVersion NullableInt32, ) *GetApiChartsByChartIdVersions200Response`

NewGetApiChartsByChartIdVersions200Response instantiates a new GetApiChartsByChartIdVersions200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetApiChartsByChartIdVersions200ResponseWithDefaults

`func NewGetApiChartsByChartIdVersions200ResponseWithDefaults() *GetApiChartsByChartIdVersions200Response`

NewGetApiChartsByChartIdVersions200ResponseWithDefaults instantiates a new GetApiChartsByChartIdVersions200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItems

`func (o *GetApiChartsByChartIdVersions200Response) GetItems() []GetApiChartsByChartIdVersions200ResponseItemsInner`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *GetApiChartsByChartIdVersions200Response) GetItemsOk() (*[]GetApiChartsByChartIdVersions200ResponseItemsInner, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *GetApiChartsByChartIdVersions200Response) SetItems(v []GetApiChartsByChartIdVersions200ResponseItemsInner)`

SetItems sets Items field to given value.


### GetNextCursor

`func (o *GetApiChartsByChartIdVersions200Response) GetNextCursor() string`

GetNextCursor returns the NextCursor field if non-nil, zero value otherwise.

### GetNextCursorOk

`func (o *GetApiChartsByChartIdVersions200Response) GetNextCursorOk() (*string, bool)`

GetNextCursorOk returns a tuple with the NextCursor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextCursor

`func (o *GetApiChartsByChartIdVersions200Response) SetNextCursor(v string)`

SetNextCursor sets NextCursor field to given value.


### SetNextCursorNil

`func (o *GetApiChartsByChartIdVersions200Response) SetNextCursorNil(b bool)`

 SetNextCursorNil sets the value for NextCursor to be an explicit nil

### UnsetNextCursor
`func (o *GetApiChartsByChartIdVersions200Response) UnsetNextCursor()`

UnsetNextCursor ensures that no value is present for NextCursor, not even an explicit nil
### GetCurrentVersion

`func (o *GetApiChartsByChartIdVersions200Response) GetCurrentVersion() int32`

GetCurrentVersion returns the CurrentVersion field if non-nil, zero value otherwise.

### GetCurrentVersionOk

`func (o *GetApiChartsByChartIdVersions200Response) GetCurrentVersionOk() (*int32, bool)`

GetCurrentVersionOk returns a tuple with the CurrentVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentVersion

`func (o *GetApiChartsByChartIdVersions200Response) SetCurrentVersion(v int32)`

SetCurrentVersion sets CurrentVersion field to given value.


### SetCurrentVersionNil

`func (o *GetApiChartsByChartIdVersions200Response) SetCurrentVersionNil(b bool)`

 SetCurrentVersionNil sets the value for CurrentVersion to be an explicit nil

### UnsetCurrentVersion
`func (o *GetApiChartsByChartIdVersions200Response) UnsetCurrentVersion()`

UnsetCurrentVersion ensures that no value is present for CurrentVersion, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


