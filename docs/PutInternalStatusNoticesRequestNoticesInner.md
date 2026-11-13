# PutInternalStatusNoticesRequestNoticesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Kind** | **string** |  | 
**Title** | **string** |  | 
**Message** | Pointer to **string** |  | [optional] 
**Impact** | Pointer to **string** |  | [optional] 
**Components** | Pointer to **[]string** |  | [optional] 
**StartsAt** | **time.Time** |  | 
**EndsAt** | Pointer to **NullableTime** |  | [optional] 

## Methods

### NewPutInternalStatusNoticesRequestNoticesInner

`func NewPutInternalStatusNoticesRequestNoticesInner(id string, kind string, title string, startsAt time.Time, ) *PutInternalStatusNoticesRequestNoticesInner`

NewPutInternalStatusNoticesRequestNoticesInner instantiates a new PutInternalStatusNoticesRequestNoticesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPutInternalStatusNoticesRequestNoticesInnerWithDefaults

`func NewPutInternalStatusNoticesRequestNoticesInnerWithDefaults() *PutInternalStatusNoticesRequestNoticesInner`

NewPutInternalStatusNoticesRequestNoticesInnerWithDefaults instantiates a new PutInternalStatusNoticesRequestNoticesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PutInternalStatusNoticesRequestNoticesInner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PutInternalStatusNoticesRequestNoticesInner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PutInternalStatusNoticesRequestNoticesInner) SetId(v string)`

SetId sets Id field to given value.


### GetKind

`func (o *PutInternalStatusNoticesRequestNoticesInner) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *PutInternalStatusNoticesRequestNoticesInner) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *PutInternalStatusNoticesRequestNoticesInner) SetKind(v string)`

SetKind sets Kind field to given value.


### GetTitle

`func (o *PutInternalStatusNoticesRequestNoticesInner) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *PutInternalStatusNoticesRequestNoticesInner) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *PutInternalStatusNoticesRequestNoticesInner) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetMessage

`func (o *PutInternalStatusNoticesRequestNoticesInner) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *PutInternalStatusNoticesRequestNoticesInner) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *PutInternalStatusNoticesRequestNoticesInner) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *PutInternalStatusNoticesRequestNoticesInner) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetImpact

`func (o *PutInternalStatusNoticesRequestNoticesInner) GetImpact() string`

GetImpact returns the Impact field if non-nil, zero value otherwise.

### GetImpactOk

`func (o *PutInternalStatusNoticesRequestNoticesInner) GetImpactOk() (*string, bool)`

GetImpactOk returns a tuple with the Impact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImpact

`func (o *PutInternalStatusNoticesRequestNoticesInner) SetImpact(v string)`

SetImpact sets Impact field to given value.

### HasImpact

`func (o *PutInternalStatusNoticesRequestNoticesInner) HasImpact() bool`

HasImpact returns a boolean if a field has been set.

### GetComponents

`func (o *PutInternalStatusNoticesRequestNoticesInner) GetComponents() []string`

GetComponents returns the Components field if non-nil, zero value otherwise.

### GetComponentsOk

`func (o *PutInternalStatusNoticesRequestNoticesInner) GetComponentsOk() (*[]string, bool)`

GetComponentsOk returns a tuple with the Components field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponents

`func (o *PutInternalStatusNoticesRequestNoticesInner) SetComponents(v []string)`

SetComponents sets Components field to given value.

### HasComponents

`func (o *PutInternalStatusNoticesRequestNoticesInner) HasComponents() bool`

HasComponents returns a boolean if a field has been set.

### GetStartsAt

`func (o *PutInternalStatusNoticesRequestNoticesInner) GetStartsAt() time.Time`

GetStartsAt returns the StartsAt field if non-nil, zero value otherwise.

### GetStartsAtOk

`func (o *PutInternalStatusNoticesRequestNoticesInner) GetStartsAtOk() (*time.Time, bool)`

GetStartsAtOk returns a tuple with the StartsAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartsAt

`func (o *PutInternalStatusNoticesRequestNoticesInner) SetStartsAt(v time.Time)`

SetStartsAt sets StartsAt field to given value.


### GetEndsAt

`func (o *PutInternalStatusNoticesRequestNoticesInner) GetEndsAt() time.Time`

GetEndsAt returns the EndsAt field if non-nil, zero value otherwise.

### GetEndsAtOk

`func (o *PutInternalStatusNoticesRequestNoticesInner) GetEndsAtOk() (*time.Time, bool)`

GetEndsAtOk returns a tuple with the EndsAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndsAt

`func (o *PutInternalStatusNoticesRequestNoticesInner) SetEndsAt(v time.Time)`

SetEndsAt sets EndsAt field to given value.

### HasEndsAt

`func (o *PutInternalStatusNoticesRequestNoticesInner) HasEndsAt() bool`

HasEndsAt returns a boolean if a field has been set.

### SetEndsAtNil

`func (o *PutInternalStatusNoticesRequestNoticesInner) SetEndsAtNil(b bool)`

 SetEndsAtNil sets the value for EndsAt to be an explicit nil

### UnsetEndsAt
`func (o *PutInternalStatusNoticesRequestNoticesInner) UnsetEndsAt()`

UnsetEndsAt ensures that no value is present for EndsAt, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


