# GetApiStatus200ResponseMaintenanceInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Title** | **string** |  | 
**Message** | **string** |  | 
**Components** | **[]string** |  | 
**StartsAt** | **time.Time** |  | 
**EndsAt** | **NullableTime** | Null when open-ended. | 
**State** | **string** |  | 

## Methods

### NewGetApiStatus200ResponseMaintenanceInner

`func NewGetApiStatus200ResponseMaintenanceInner(id string, title string, message string, components []string, startsAt time.Time, endsAt NullableTime, state string, ) *GetApiStatus200ResponseMaintenanceInner`

NewGetApiStatus200ResponseMaintenanceInner instantiates a new GetApiStatus200ResponseMaintenanceInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetApiStatus200ResponseMaintenanceInnerWithDefaults

`func NewGetApiStatus200ResponseMaintenanceInnerWithDefaults() *GetApiStatus200ResponseMaintenanceInner`

NewGetApiStatus200ResponseMaintenanceInnerWithDefaults instantiates a new GetApiStatus200ResponseMaintenanceInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GetApiStatus200ResponseMaintenanceInner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetApiStatus200ResponseMaintenanceInner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetApiStatus200ResponseMaintenanceInner) SetId(v string)`

SetId sets Id field to given value.


### GetTitle

`func (o *GetApiStatus200ResponseMaintenanceInner) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *GetApiStatus200ResponseMaintenanceInner) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *GetApiStatus200ResponseMaintenanceInner) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetMessage

`func (o *GetApiStatus200ResponseMaintenanceInner) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *GetApiStatus200ResponseMaintenanceInner) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *GetApiStatus200ResponseMaintenanceInner) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetComponents

`func (o *GetApiStatus200ResponseMaintenanceInner) GetComponents() []string`

GetComponents returns the Components field if non-nil, zero value otherwise.

### GetComponentsOk

`func (o *GetApiStatus200ResponseMaintenanceInner) GetComponentsOk() (*[]string, bool)`

GetComponentsOk returns a tuple with the Components field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponents

`func (o *GetApiStatus200ResponseMaintenanceInner) SetComponents(v []string)`

SetComponents sets Components field to given value.


### GetStartsAt

`func (o *GetApiStatus200ResponseMaintenanceInner) GetStartsAt() time.Time`

GetStartsAt returns the StartsAt field if non-nil, zero value otherwise.

### GetStartsAtOk

`func (o *GetApiStatus200ResponseMaintenanceInner) GetStartsAtOk() (*time.Time, bool)`

GetStartsAtOk returns a tuple with the StartsAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartsAt

`func (o *GetApiStatus200ResponseMaintenanceInner) SetStartsAt(v time.Time)`

SetStartsAt sets StartsAt field to given value.


### GetEndsAt

`func (o *GetApiStatus200ResponseMaintenanceInner) GetEndsAt() time.Time`

GetEndsAt returns the EndsAt field if non-nil, zero value otherwise.

### GetEndsAtOk

`func (o *GetApiStatus200ResponseMaintenanceInner) GetEndsAtOk() (*time.Time, bool)`

GetEndsAtOk returns a tuple with the EndsAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndsAt

`func (o *GetApiStatus200ResponseMaintenanceInner) SetEndsAt(v time.Time)`

SetEndsAt sets EndsAt field to given value.


### SetEndsAtNil

`func (o *GetApiStatus200ResponseMaintenanceInner) SetEndsAtNil(b bool)`

 SetEndsAtNil sets the value for EndsAt to be an explicit nil

### UnsetEndsAt
`func (o *GetApiStatus200ResponseMaintenanceInner) UnsetEndsAt()`

UnsetEndsAt ensures that no value is present for EndsAt, not even an explicit nil
### GetState

`func (o *GetApiStatus200ResponseMaintenanceInner) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *GetApiStatus200ResponseMaintenanceInner) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *GetApiStatus200ResponseMaintenanceInner) SetState(v string)`

SetState sets State field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


