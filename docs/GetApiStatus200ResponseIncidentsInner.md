# GetApiStatus200ResponseIncidentsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Source** | **string** | &#x60;automatic&#x60;: derived from three consecutive failed samples (resolved after three good ones); &#x60;manual&#x60;: posted by an operator. | 
**Components** | **[]string** |  | 
**Severity** | **string** | &#x60;major&#x60; when the component was down, &#x60;minor&#x60; when degraded. | 
**Title** | **NullableString** | Manual incidents only (English). | 
**Message** | **NullableString** | Manual incidents only (English). | 
**StartedAt** | **time.Time** |  | 
**ResolvedAt** | **NullableTime** | Null while the incident is open. | 

## Methods

### NewGetApiStatus200ResponseIncidentsInner

`func NewGetApiStatus200ResponseIncidentsInner(id string, source string, components []string, severity string, title NullableString, message NullableString, startedAt time.Time, resolvedAt NullableTime, ) *GetApiStatus200ResponseIncidentsInner`

NewGetApiStatus200ResponseIncidentsInner instantiates a new GetApiStatus200ResponseIncidentsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetApiStatus200ResponseIncidentsInnerWithDefaults

`func NewGetApiStatus200ResponseIncidentsInnerWithDefaults() *GetApiStatus200ResponseIncidentsInner`

NewGetApiStatus200ResponseIncidentsInnerWithDefaults instantiates a new GetApiStatus200ResponseIncidentsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GetApiStatus200ResponseIncidentsInner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetApiStatus200ResponseIncidentsInner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetApiStatus200ResponseIncidentsInner) SetId(v string)`

SetId sets Id field to given value.


### GetSource

`func (o *GetApiStatus200ResponseIncidentsInner) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *GetApiStatus200ResponseIncidentsInner) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *GetApiStatus200ResponseIncidentsInner) SetSource(v string)`

SetSource sets Source field to given value.


### GetComponents

`func (o *GetApiStatus200ResponseIncidentsInner) GetComponents() []string`

GetComponents returns the Components field if non-nil, zero value otherwise.

### GetComponentsOk

`func (o *GetApiStatus200ResponseIncidentsInner) GetComponentsOk() (*[]string, bool)`

GetComponentsOk returns a tuple with the Components field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponents

`func (o *GetApiStatus200ResponseIncidentsInner) SetComponents(v []string)`

SetComponents sets Components field to given value.


### GetSeverity

`func (o *GetApiStatus200ResponseIncidentsInner) GetSeverity() string`

GetSeverity returns the Severity field if non-nil, zero value otherwise.

### GetSeverityOk

`func (o *GetApiStatus200ResponseIncidentsInner) GetSeverityOk() (*string, bool)`

GetSeverityOk returns a tuple with the Severity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverity

`func (o *GetApiStatus200ResponseIncidentsInner) SetSeverity(v string)`

SetSeverity sets Severity field to given value.


### GetTitle

`func (o *GetApiStatus200ResponseIncidentsInner) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *GetApiStatus200ResponseIncidentsInner) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *GetApiStatus200ResponseIncidentsInner) SetTitle(v string)`

SetTitle sets Title field to given value.


### SetTitleNil

`func (o *GetApiStatus200ResponseIncidentsInner) SetTitleNil(b bool)`

 SetTitleNil sets the value for Title to be an explicit nil

### UnsetTitle
`func (o *GetApiStatus200ResponseIncidentsInner) UnsetTitle()`

UnsetTitle ensures that no value is present for Title, not even an explicit nil
### GetMessage

`func (o *GetApiStatus200ResponseIncidentsInner) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *GetApiStatus200ResponseIncidentsInner) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *GetApiStatus200ResponseIncidentsInner) SetMessage(v string)`

SetMessage sets Message field to given value.


### SetMessageNil

`func (o *GetApiStatus200ResponseIncidentsInner) SetMessageNil(b bool)`

 SetMessageNil sets the value for Message to be an explicit nil

### UnsetMessage
`func (o *GetApiStatus200ResponseIncidentsInner) UnsetMessage()`

UnsetMessage ensures that no value is present for Message, not even an explicit nil
### GetStartedAt

`func (o *GetApiStatus200ResponseIncidentsInner) GetStartedAt() time.Time`

GetStartedAt returns the StartedAt field if non-nil, zero value otherwise.

### GetStartedAtOk

`func (o *GetApiStatus200ResponseIncidentsInner) GetStartedAtOk() (*time.Time, bool)`

GetStartedAtOk returns a tuple with the StartedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartedAt

`func (o *GetApiStatus200ResponseIncidentsInner) SetStartedAt(v time.Time)`

SetStartedAt sets StartedAt field to given value.


### GetResolvedAt

`func (o *GetApiStatus200ResponseIncidentsInner) GetResolvedAt() time.Time`

GetResolvedAt returns the ResolvedAt field if non-nil, zero value otherwise.

### GetResolvedAtOk

`func (o *GetApiStatus200ResponseIncidentsInner) GetResolvedAtOk() (*time.Time, bool)`

GetResolvedAtOk returns a tuple with the ResolvedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResolvedAt

`func (o *GetApiStatus200ResponseIncidentsInner) SetResolvedAt(v time.Time)`

SetResolvedAt sets ResolvedAt field to given value.


### SetResolvedAtNil

`func (o *GetApiStatus200ResponseIncidentsInner) SetResolvedAtNil(b bool)`

 SetResolvedAtNil sets the value for ResolvedAt to be an explicit nil

### UnsetResolvedAt
`func (o *GetApiStatus200ResponseIncidentsInner) UnsetResolvedAt()`

UnsetResolvedAt ensures that no value is present for ResolvedAt, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


