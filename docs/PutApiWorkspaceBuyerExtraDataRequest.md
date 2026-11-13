# PutApiWorkspaceBuyerExtraDataRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | **bool** | Publish the listed keys to buyers. | 
**Keys** | **[]string** | extraData keys buyers may read (replaces the list). | 

## Methods

### NewPutApiWorkspaceBuyerExtraDataRequest

`func NewPutApiWorkspaceBuyerExtraDataRequest(enabled bool, keys []string, ) *PutApiWorkspaceBuyerExtraDataRequest`

NewPutApiWorkspaceBuyerExtraDataRequest instantiates a new PutApiWorkspaceBuyerExtraDataRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPutApiWorkspaceBuyerExtraDataRequestWithDefaults

`func NewPutApiWorkspaceBuyerExtraDataRequestWithDefaults() *PutApiWorkspaceBuyerExtraDataRequest`

NewPutApiWorkspaceBuyerExtraDataRequestWithDefaults instantiates a new PutApiWorkspaceBuyerExtraDataRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *PutApiWorkspaceBuyerExtraDataRequest) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *PutApiWorkspaceBuyerExtraDataRequest) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *PutApiWorkspaceBuyerExtraDataRequest) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetKeys

`func (o *PutApiWorkspaceBuyerExtraDataRequest) GetKeys() []string`

GetKeys returns the Keys field if non-nil, zero value otherwise.

### GetKeysOk

`func (o *PutApiWorkspaceBuyerExtraDataRequest) GetKeysOk() (*[]string, bool)`

GetKeysOk returns a tuple with the Keys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeys

`func (o *PutApiWorkspaceBuyerExtraDataRequest) SetKeys(v []string)`

SetKeys sets Keys field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


