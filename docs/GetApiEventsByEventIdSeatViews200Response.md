# GetApiEventsByEventIdSeatViews200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EventId** | **string** |  | 
**Views** | [**map[string]GetApiEventsByEventIdSeatViews200ResponseViewsValue**](GetApiEventsByEventIdSeatViews200ResponseViewsValue.md) |  | 

## Methods

### NewGetApiEventsByEventIdSeatViews200Response

`func NewGetApiEventsByEventIdSeatViews200Response(eventId string, views map[string]GetApiEventsByEventIdSeatViews200ResponseViewsValue, ) *GetApiEventsByEventIdSeatViews200Response`

NewGetApiEventsByEventIdSeatViews200Response instantiates a new GetApiEventsByEventIdSeatViews200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetApiEventsByEventIdSeatViews200ResponseWithDefaults

`func NewGetApiEventsByEventIdSeatViews200ResponseWithDefaults() *GetApiEventsByEventIdSeatViews200Response`

NewGetApiEventsByEventIdSeatViews200ResponseWithDefaults instantiates a new GetApiEventsByEventIdSeatViews200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEventId

`func (o *GetApiEventsByEventIdSeatViews200Response) GetEventId() string`

GetEventId returns the EventId field if non-nil, zero value otherwise.

### GetEventIdOk

`func (o *GetApiEventsByEventIdSeatViews200Response) GetEventIdOk() (*string, bool)`

GetEventIdOk returns a tuple with the EventId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventId

`func (o *GetApiEventsByEventIdSeatViews200Response) SetEventId(v string)`

SetEventId sets EventId field to given value.


### GetViews

`func (o *GetApiEventsByEventIdSeatViews200Response) GetViews() map[string]GetApiEventsByEventIdSeatViews200ResponseViewsValue`

GetViews returns the Views field if non-nil, zero value otherwise.

### GetViewsOk

`func (o *GetApiEventsByEventIdSeatViews200Response) GetViewsOk() (*map[string]GetApiEventsByEventIdSeatViews200ResponseViewsValue, bool)`

GetViewsOk returns a tuple with the Views field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViews

`func (o *GetApiEventsByEventIdSeatViews200Response) SetViews(v map[string]GetApiEventsByEventIdSeatViews200ResponseViewsValue)`

SetViews sets Views field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


