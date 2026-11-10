# GetApiWorkspaceBuyerExtraData200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | **bool** | Publish the listed keys to buyers (default false: nothing is published). | 
**Keys** | **[]string** | extraData keys buyers may read. | 

## Methods

### NewGetApiWorkspaceBuyerExtraData200Response

`func NewGetApiWorkspaceBuyerExtraData200Response(enabled bool, keys []string, ) *GetApiWorkspaceBuyerExtraData200Response`

NewGetApiWorkspaceBuyerExtraData200Response instantiates a new GetApiWorkspaceBuyerExtraData200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetApiWorkspaceBuyerExtraData200ResponseWithDefaults

`func NewGetApiWorkspaceBuyerExtraData200ResponseWithDefaults() *GetApiWorkspaceBuyerExtraData200Response`

NewGetApiWorkspaceBuyerExtraData200ResponseWithDefaults instantiates a new GetApiWorkspaceBuyerExtraData200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *GetApiWorkspaceBuyerExtraData200Response) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *GetApiWorkspaceBuyerExtraData200Response) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *GetApiWorkspaceBuyerExtraData200Response) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetKeys

`func (o *GetApiWorkspaceBuyerExtraData200Response) GetKeys() []string`

GetKeys returns the Keys field if non-nil, zero value otherwise.

### GetKeysOk

`func (o *GetApiWorkspaceBuyerExtraData200Response) GetKeysOk() (*[]string, bool)`

GetKeysOk returns a tuple with the Keys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeys

`func (o *GetApiWorkspaceBuyerExtraData200Response) SetKeys(v []string)`

SetKeys sets Keys field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


