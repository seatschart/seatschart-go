# PostApiWorkspacesRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DisplayName** | **string** |  | 
**IsTest** | Pointer to **bool** | Create a test (sandbox) workspace: never billed, not metered, no chart/event caps, &#x60;sk_test_&#x60; API keys, a TEST badge for buyers. Permanent. | [optional] 
**CopyCharts** | Pointer to **bool** | Test workspaces only: copy the current workspace&#39;s charts as unpublished drafts (up to 50, archived skipped). Default false. | [optional] 

## Methods

### NewPostApiWorkspacesRequest

`func NewPostApiWorkspacesRequest(displayName string, ) *PostApiWorkspacesRequest`

NewPostApiWorkspacesRequest instantiates a new PostApiWorkspacesRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostApiWorkspacesRequestWithDefaults

`func NewPostApiWorkspacesRequestWithDefaults() *PostApiWorkspacesRequest`

NewPostApiWorkspacesRequestWithDefaults instantiates a new PostApiWorkspacesRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDisplayName

`func (o *PostApiWorkspacesRequest) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *PostApiWorkspacesRequest) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *PostApiWorkspacesRequest) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.


### GetIsTest

`func (o *PostApiWorkspacesRequest) GetIsTest() bool`

GetIsTest returns the IsTest field if non-nil, zero value otherwise.

### GetIsTestOk

`func (o *PostApiWorkspacesRequest) GetIsTestOk() (*bool, bool)`

GetIsTestOk returns a tuple with the IsTest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsTest

`func (o *PostApiWorkspacesRequest) SetIsTest(v bool)`

SetIsTest sets IsTest field to given value.

### HasIsTest

`func (o *PostApiWorkspacesRequest) HasIsTest() bool`

HasIsTest returns a boolean if a field has been set.

### GetCopyCharts

`func (o *PostApiWorkspacesRequest) GetCopyCharts() bool`

GetCopyCharts returns the CopyCharts field if non-nil, zero value otherwise.

### GetCopyChartsOk

`func (o *PostApiWorkspacesRequest) GetCopyChartsOk() (*bool, bool)`

GetCopyChartsOk returns a tuple with the CopyCharts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCopyCharts

`func (o *PostApiWorkspacesRequest) SetCopyCharts(v bool)`

SetCopyCharts sets CopyCharts field to given value.

### HasCopyCharts

`func (o *PostApiWorkspacesRequest) HasCopyCharts() bool`

HasCopyCharts returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


