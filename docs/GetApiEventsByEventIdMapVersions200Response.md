# GetApiEventsByEventIdMapVersions200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Items** | [**[]GetApiEventsByEventIdMapVersions200ResponseItemsInner**](GetApiEventsByEventIdMapVersions200ResponseItemsInner.md) |  | 
**NextCursor** | **NullableString** |  | 
**CurrentVersion** | **NullableInt32** |  | 

## Methods

### NewGetApiEventsByEventIdMapVersions200Response

`func NewGetApiEventsByEventIdMapVersions200Response(items []GetApiEventsByEventIdMapVersions200ResponseItemsInner, nextCursor NullableString, currentVersion NullableInt32, ) *GetApiEventsByEventIdMapVersions200Response`

NewGetApiEventsByEventIdMapVersions200Response instantiates a new GetApiEventsByEventIdMapVersions200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetApiEventsByEventIdMapVersions200ResponseWithDefaults

`func NewGetApiEventsByEventIdMapVersions200ResponseWithDefaults() *GetApiEventsByEventIdMapVersions200Response`

NewGetApiEventsByEventIdMapVersions200ResponseWithDefaults instantiates a new GetApiEventsByEventIdMapVersions200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItems

`func (o *GetApiEventsByEventIdMapVersions200Response) GetItems() []GetApiEventsByEventIdMapVersions200ResponseItemsInner`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *GetApiEventsByEventIdMapVersions200Response) GetItemsOk() (*[]GetApiEventsByEventIdMapVersions200ResponseItemsInner, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *GetApiEventsByEventIdMapVersions200Response) SetItems(v []GetApiEventsByEventIdMapVersions200ResponseItemsInner)`

SetItems sets Items field to given value.


### GetNextCursor

`func (o *GetApiEventsByEventIdMapVersions200Response) GetNextCursor() string`

GetNextCursor returns the NextCursor field if non-nil, zero value otherwise.

### GetNextCursorOk

`func (o *GetApiEventsByEventIdMapVersions200Response) GetNextCursorOk() (*string, bool)`

GetNextCursorOk returns a tuple with the NextCursor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextCursor

`func (o *GetApiEventsByEventIdMapVersions200Response) SetNextCursor(v string)`

SetNextCursor sets NextCursor field to given value.


### SetNextCursorNil

`func (o *GetApiEventsByEventIdMapVersions200Response) SetNextCursorNil(b bool)`

 SetNextCursorNil sets the value for NextCursor to be an explicit nil

### UnsetNextCursor
`func (o *GetApiEventsByEventIdMapVersions200Response) UnsetNextCursor()`

UnsetNextCursor ensures that no value is present for NextCursor, not even an explicit nil
### GetCurrentVersion

`func (o *GetApiEventsByEventIdMapVersions200Response) GetCurrentVersion() int32`

GetCurrentVersion returns the CurrentVersion field if non-nil, zero value otherwise.

### GetCurrentVersionOk

`func (o *GetApiEventsByEventIdMapVersions200Response) GetCurrentVersionOk() (*int32, bool)`

GetCurrentVersionOk returns a tuple with the CurrentVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentVersion

`func (o *GetApiEventsByEventIdMapVersions200Response) SetCurrentVersion(v int32)`

SetCurrentVersion sets CurrentVersion field to given value.


### SetCurrentVersionNil

`func (o *GetApiEventsByEventIdMapVersions200Response) SetCurrentVersionNil(b bool)`

 SetCurrentVersionNil sets the value for CurrentVersion to be an explicit nil

### UnsetCurrentVersion
`func (o *GetApiEventsByEventIdMapVersions200Response) UnsetCurrentVersion()`

UnsetCurrentVersion ensures that no value is present for CurrentVersion, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


