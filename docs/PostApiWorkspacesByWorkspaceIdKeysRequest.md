# PostApiWorkspacesByWorkspaceIdKeysRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**FullAccess** | Pointer to **bool** | Every scope. Default when &#x60;scopes&#x60; is absent. | [optional] 
**Scopes** | Pointer to **[]string** | Granted scopes: &#x60;read&#x60;, &#x60;holds&#x60;, &#x60;bookings&#x60;, &#x60;inventory_admin&#x60;, &#x60;charts_write&#x60;, &#x60;events_write&#x60;, &#x60;reports&#x60;, &#x60;webhooks&#x60;. | [optional] 

## Methods

### NewPostApiWorkspacesByWorkspaceIdKeysRequest

`func NewPostApiWorkspacesByWorkspaceIdKeysRequest(name string, ) *PostApiWorkspacesByWorkspaceIdKeysRequest`

NewPostApiWorkspacesByWorkspaceIdKeysRequest instantiates a new PostApiWorkspacesByWorkspaceIdKeysRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostApiWorkspacesByWorkspaceIdKeysRequestWithDefaults

`func NewPostApiWorkspacesByWorkspaceIdKeysRequestWithDefaults() *PostApiWorkspacesByWorkspaceIdKeysRequest`

NewPostApiWorkspacesByWorkspaceIdKeysRequestWithDefaults instantiates a new PostApiWorkspacesByWorkspaceIdKeysRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *PostApiWorkspacesByWorkspaceIdKeysRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PostApiWorkspacesByWorkspaceIdKeysRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PostApiWorkspacesByWorkspaceIdKeysRequest) SetName(v string)`

SetName sets Name field to given value.


### GetFullAccess

`func (o *PostApiWorkspacesByWorkspaceIdKeysRequest) GetFullAccess() bool`

GetFullAccess returns the FullAccess field if non-nil, zero value otherwise.

### GetFullAccessOk

`func (o *PostApiWorkspacesByWorkspaceIdKeysRequest) GetFullAccessOk() (*bool, bool)`

GetFullAccessOk returns a tuple with the FullAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFullAccess

`func (o *PostApiWorkspacesByWorkspaceIdKeysRequest) SetFullAccess(v bool)`

SetFullAccess sets FullAccess field to given value.

### HasFullAccess

`func (o *PostApiWorkspacesByWorkspaceIdKeysRequest) HasFullAccess() bool`

HasFullAccess returns a boolean if a field has been set.

### GetScopes

`func (o *PostApiWorkspacesByWorkspaceIdKeysRequest) GetScopes() []string`

GetScopes returns the Scopes field if non-nil, zero value otherwise.

### GetScopesOk

`func (o *PostApiWorkspacesByWorkspaceIdKeysRequest) GetScopesOk() (*[]string, bool)`

GetScopesOk returns a tuple with the Scopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopes

`func (o *PostApiWorkspacesByWorkspaceIdKeysRequest) SetScopes(v []string)`

SetScopes sets Scopes field to given value.

### HasScopes

`func (o *PostApiWorkspacesByWorkspaceIdKeysRequest) HasScopes() bool`

HasScopes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


