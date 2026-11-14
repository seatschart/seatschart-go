# GetApiWorkspacesByWorkspaceIdKeys200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**WorkspaceId** | **string** |  | 
**ApiKeys** | [**[]GetApiWorkspacesByWorkspaceIdKeys200ResponseApiKeysInner**](GetApiWorkspacesByWorkspaceIdKeys200ResponseApiKeysInner.md) |  | 

## Methods

### NewGetApiWorkspacesByWorkspaceIdKeys200Response

`func NewGetApiWorkspacesByWorkspaceIdKeys200Response(workspaceId string, apiKeys []GetApiWorkspacesByWorkspaceIdKeys200ResponseApiKeysInner, ) *GetApiWorkspacesByWorkspaceIdKeys200Response`

NewGetApiWorkspacesByWorkspaceIdKeys200Response instantiates a new GetApiWorkspacesByWorkspaceIdKeys200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetApiWorkspacesByWorkspaceIdKeys200ResponseWithDefaults

`func NewGetApiWorkspacesByWorkspaceIdKeys200ResponseWithDefaults() *GetApiWorkspacesByWorkspaceIdKeys200Response`

NewGetApiWorkspacesByWorkspaceIdKeys200ResponseWithDefaults instantiates a new GetApiWorkspacesByWorkspaceIdKeys200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWorkspaceId

`func (o *GetApiWorkspacesByWorkspaceIdKeys200Response) GetWorkspaceId() string`

GetWorkspaceId returns the WorkspaceId field if non-nil, zero value otherwise.

### GetWorkspaceIdOk

`func (o *GetApiWorkspacesByWorkspaceIdKeys200Response) GetWorkspaceIdOk() (*string, bool)`

GetWorkspaceIdOk returns a tuple with the WorkspaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkspaceId

`func (o *GetApiWorkspacesByWorkspaceIdKeys200Response) SetWorkspaceId(v string)`

SetWorkspaceId sets WorkspaceId field to given value.


### GetApiKeys

`func (o *GetApiWorkspacesByWorkspaceIdKeys200Response) GetApiKeys() []GetApiWorkspacesByWorkspaceIdKeys200ResponseApiKeysInner`

GetApiKeys returns the ApiKeys field if non-nil, zero value otherwise.

### GetApiKeysOk

`func (o *GetApiWorkspacesByWorkspaceIdKeys200Response) GetApiKeysOk() (*[]GetApiWorkspacesByWorkspaceIdKeys200ResponseApiKeysInner, bool)`

GetApiKeysOk returns a tuple with the ApiKeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiKeys

`func (o *GetApiWorkspacesByWorkspaceIdKeys200Response) SetApiKeys(v []GetApiWorkspacesByWorkspaceIdKeys200ResponseApiKeysInner)`

SetApiKeys sets ApiKeys field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


