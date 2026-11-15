# GetApiWorkspacesOverview200ResponseItemsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**WorkspaceId** | **string** |  | 
**DisplayName** | **string** |  | 
**ExternalId** | **NullableString** |  | 
**Status** | **string** |  | 
**CreatedByApi** | **bool** |  | 
**CreatedAt** | **time.Time** |  | 
**PricingSource** | **string** |  | 
**Charts** | **int32** |  | 
**Events** | **int32** |  | 
**Seats** | [**GetApiWorkspacesOverview200ResponseTotalsSeats**](GetApiWorkspacesOverview200ResponseTotalsSeats.md) |  | 
**Revenue** | [**GetApiWorkspacesOverview200ResponseItemsInnerRevenue**](GetApiWorkspacesOverview200ResponseItemsInnerRevenue.md) |  | 
**Usage** | [**GetApiWorkspacesOverview200ResponseItemsInnerUsage**](GetApiWorkspacesOverview200ResponseItemsInnerUsage.md) |  | 

## Methods

### NewGetApiWorkspacesOverview200ResponseItemsInner

`func NewGetApiWorkspacesOverview200ResponseItemsInner(workspaceId string, displayName string, externalId NullableString, status string, createdByApi bool, createdAt time.Time, pricingSource string, charts int32, events int32, seats GetApiWorkspacesOverview200ResponseTotalsSeats, revenue GetApiWorkspacesOverview200ResponseItemsInnerRevenue, usage GetApiWorkspacesOverview200ResponseItemsInnerUsage, ) *GetApiWorkspacesOverview200ResponseItemsInner`

NewGetApiWorkspacesOverview200ResponseItemsInner instantiates a new GetApiWorkspacesOverview200ResponseItemsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetApiWorkspacesOverview200ResponseItemsInnerWithDefaults

`func NewGetApiWorkspacesOverview200ResponseItemsInnerWithDefaults() *GetApiWorkspacesOverview200ResponseItemsInner`

NewGetApiWorkspacesOverview200ResponseItemsInnerWithDefaults instantiates a new GetApiWorkspacesOverview200ResponseItemsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWorkspaceId

`func (o *GetApiWorkspacesOverview200ResponseItemsInner) GetWorkspaceId() string`

GetWorkspaceId returns the WorkspaceId field if non-nil, zero value otherwise.

### GetWorkspaceIdOk

`func (o *GetApiWorkspacesOverview200ResponseItemsInner) GetWorkspaceIdOk() (*string, bool)`

GetWorkspaceIdOk returns a tuple with the WorkspaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkspaceId

`func (o *GetApiWorkspacesOverview200ResponseItemsInner) SetWorkspaceId(v string)`

SetWorkspaceId sets WorkspaceId field to given value.


### GetDisplayName

`func (o *GetApiWorkspacesOverview200ResponseItemsInner) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *GetApiWorkspacesOverview200ResponseItemsInner) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *GetApiWorkspacesOverview200ResponseItemsInner) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.


### GetExternalId

`func (o *GetApiWorkspacesOverview200ResponseItemsInner) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *GetApiWorkspacesOverview200ResponseItemsInner) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *GetApiWorkspacesOverview200ResponseItemsInner) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.


### SetExternalIdNil

`func (o *GetApiWorkspacesOverview200ResponseItemsInner) SetExternalIdNil(b bool)`

 SetExternalIdNil sets the value for ExternalId to be an explicit nil

### UnsetExternalId
`func (o *GetApiWorkspacesOverview200ResponseItemsInner) UnsetExternalId()`

UnsetExternalId ensures that no value is present for ExternalId, not even an explicit nil
### GetStatus

`func (o *GetApiWorkspacesOverview200ResponseItemsInner) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetApiWorkspacesOverview200ResponseItemsInner) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetApiWorkspacesOverview200ResponseItemsInner) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetCreatedByApi

`func (o *GetApiWorkspacesOverview200ResponseItemsInner) GetCreatedByApi() bool`

GetCreatedByApi returns the CreatedByApi field if non-nil, zero value otherwise.

### GetCreatedByApiOk

`func (o *GetApiWorkspacesOverview200ResponseItemsInner) GetCreatedByApiOk() (*bool, bool)`

GetCreatedByApiOk returns a tuple with the CreatedByApi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedByApi

`func (o *GetApiWorkspacesOverview200ResponseItemsInner) SetCreatedByApi(v bool)`

SetCreatedByApi sets CreatedByApi field to given value.


### GetCreatedAt

`func (o *GetApiWorkspacesOverview200ResponseItemsInner) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *GetApiWorkspacesOverview200ResponseItemsInner) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *GetApiWorkspacesOverview200ResponseItemsInner) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetPricingSource

`func (o *GetApiWorkspacesOverview200ResponseItemsInner) GetPricingSource() string`

GetPricingSource returns the PricingSource field if non-nil, zero value otherwise.

### GetPricingSourceOk

`func (o *GetApiWorkspacesOverview200ResponseItemsInner) GetPricingSourceOk() (*string, bool)`

GetPricingSourceOk returns a tuple with the PricingSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPricingSource

`func (o *GetApiWorkspacesOverview200ResponseItemsInner) SetPricingSource(v string)`

SetPricingSource sets PricingSource field to given value.


### GetCharts

`func (o *GetApiWorkspacesOverview200ResponseItemsInner) GetCharts() int32`

GetCharts returns the Charts field if non-nil, zero value otherwise.

### GetChartsOk

`func (o *GetApiWorkspacesOverview200ResponseItemsInner) GetChartsOk() (*int32, bool)`

GetChartsOk returns a tuple with the Charts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCharts

`func (o *GetApiWorkspacesOverview200ResponseItemsInner) SetCharts(v int32)`

SetCharts sets Charts field to given value.


### GetEvents

`func (o *GetApiWorkspacesOverview200ResponseItemsInner) GetEvents() int32`

GetEvents returns the Events field if non-nil, zero value otherwise.

### GetEventsOk

`func (o *GetApiWorkspacesOverview200ResponseItemsInner) GetEventsOk() (*int32, bool)`

GetEventsOk returns a tuple with the Events field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvents

`func (o *GetApiWorkspacesOverview200ResponseItemsInner) SetEvents(v int32)`

SetEvents sets Events field to given value.


### GetSeats

`func (o *GetApiWorkspacesOverview200ResponseItemsInner) GetSeats() GetApiWorkspacesOverview200ResponseTotalsSeats`

GetSeats returns the Seats field if non-nil, zero value otherwise.

### GetSeatsOk

`func (o *GetApiWorkspacesOverview200ResponseItemsInner) GetSeatsOk() (*GetApiWorkspacesOverview200ResponseTotalsSeats, bool)`

GetSeatsOk returns a tuple with the Seats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeats

`func (o *GetApiWorkspacesOverview200ResponseItemsInner) SetSeats(v GetApiWorkspacesOverview200ResponseTotalsSeats)`

SetSeats sets Seats field to given value.


### GetRevenue

`func (o *GetApiWorkspacesOverview200ResponseItemsInner) GetRevenue() GetApiWorkspacesOverview200ResponseItemsInnerRevenue`

GetRevenue returns the Revenue field if non-nil, zero value otherwise.

### GetRevenueOk

`func (o *GetApiWorkspacesOverview200ResponseItemsInner) GetRevenueOk() (*GetApiWorkspacesOverview200ResponseItemsInnerRevenue, bool)`

GetRevenueOk returns a tuple with the Revenue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevenue

`func (o *GetApiWorkspacesOverview200ResponseItemsInner) SetRevenue(v GetApiWorkspacesOverview200ResponseItemsInnerRevenue)`

SetRevenue sets Revenue field to given value.


### GetUsage

`func (o *GetApiWorkspacesOverview200ResponseItemsInner) GetUsage() GetApiWorkspacesOverview200ResponseItemsInnerUsage`

GetUsage returns the Usage field if non-nil, zero value otherwise.

### GetUsageOk

`func (o *GetApiWorkspacesOverview200ResponseItemsInner) GetUsageOk() (*GetApiWorkspacesOverview200ResponseItemsInnerUsage, bool)`

GetUsageOk returns a tuple with the Usage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsage

`func (o *GetApiWorkspacesOverview200ResponseItemsInner) SetUsage(v GetApiWorkspacesOverview200ResponseItemsInnerUsage)`

SetUsage sets Usage field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


