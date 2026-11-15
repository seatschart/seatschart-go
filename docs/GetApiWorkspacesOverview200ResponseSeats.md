# GetApiWorkspacesOverview200ResponseSeats

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Month** | **string** |  | 
**PlanId** | **string** |  | 
**Mode** | **string** |  | 
**UsedSeats** | **int32** |  | 
**IncludedSeats** | **NullableInt32** |  | 
**OveragePriceCentsPerSeat** | **NullableFloat32** | US cents per seat beyond the included ones (may be fractional). | 
**Overage** | [**GetApiWorkspacesOverview200ResponseSeatsOverage**](GetApiWorkspacesOverview200ResponseSeatsOverage.md) |  | 
**AllowanceExceeded** | **bool** |  | 
**AllowanceNear** | **bool** |  | 
**Projected** | Pointer to [**NullableGetApiWorkspacesOverview200ResponseSeatsProjected**](GetApiWorkspacesOverview200ResponseSeatsProjected.md) |  | [optional] 
**BillingOwnerUserId** | **NullableString** |  | 
**RolledUp** | **bool** |  | 
**WorkspaceCount** | **int32** |  | 

## Methods

### NewGetApiWorkspacesOverview200ResponseSeats

`func NewGetApiWorkspacesOverview200ResponseSeats(month string, planId string, mode string, usedSeats int32, includedSeats NullableInt32, overagePriceCentsPerSeat NullableFloat32, overage GetApiWorkspacesOverview200ResponseSeatsOverage, allowanceExceeded bool, allowanceNear bool, billingOwnerUserId NullableString, rolledUp bool, workspaceCount int32, ) *GetApiWorkspacesOverview200ResponseSeats`

NewGetApiWorkspacesOverview200ResponseSeats instantiates a new GetApiWorkspacesOverview200ResponseSeats object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetApiWorkspacesOverview200ResponseSeatsWithDefaults

`func NewGetApiWorkspacesOverview200ResponseSeatsWithDefaults() *GetApiWorkspacesOverview200ResponseSeats`

NewGetApiWorkspacesOverview200ResponseSeatsWithDefaults instantiates a new GetApiWorkspacesOverview200ResponseSeats object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMonth

`func (o *GetApiWorkspacesOverview200ResponseSeats) GetMonth() string`

GetMonth returns the Month field if non-nil, zero value otherwise.

### GetMonthOk

`func (o *GetApiWorkspacesOverview200ResponseSeats) GetMonthOk() (*string, bool)`

GetMonthOk returns a tuple with the Month field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonth

`func (o *GetApiWorkspacesOverview200ResponseSeats) SetMonth(v string)`

SetMonth sets Month field to given value.


### GetPlanId

`func (o *GetApiWorkspacesOverview200ResponseSeats) GetPlanId() string`

GetPlanId returns the PlanId field if non-nil, zero value otherwise.

### GetPlanIdOk

`func (o *GetApiWorkspacesOverview200ResponseSeats) GetPlanIdOk() (*string, bool)`

GetPlanIdOk returns a tuple with the PlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanId

`func (o *GetApiWorkspacesOverview200ResponseSeats) SetPlanId(v string)`

SetPlanId sets PlanId field to given value.


### GetMode

`func (o *GetApiWorkspacesOverview200ResponseSeats) GetMode() string`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *GetApiWorkspacesOverview200ResponseSeats) GetModeOk() (*string, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *GetApiWorkspacesOverview200ResponseSeats) SetMode(v string)`

SetMode sets Mode field to given value.


### GetUsedSeats

`func (o *GetApiWorkspacesOverview200ResponseSeats) GetUsedSeats() int32`

GetUsedSeats returns the UsedSeats field if non-nil, zero value otherwise.

### GetUsedSeatsOk

`func (o *GetApiWorkspacesOverview200ResponseSeats) GetUsedSeatsOk() (*int32, bool)`

GetUsedSeatsOk returns a tuple with the UsedSeats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsedSeats

`func (o *GetApiWorkspacesOverview200ResponseSeats) SetUsedSeats(v int32)`

SetUsedSeats sets UsedSeats field to given value.


### GetIncludedSeats

`func (o *GetApiWorkspacesOverview200ResponseSeats) GetIncludedSeats() int32`

GetIncludedSeats returns the IncludedSeats field if non-nil, zero value otherwise.

### GetIncludedSeatsOk

`func (o *GetApiWorkspacesOverview200ResponseSeats) GetIncludedSeatsOk() (*int32, bool)`

GetIncludedSeatsOk returns a tuple with the IncludedSeats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludedSeats

`func (o *GetApiWorkspacesOverview200ResponseSeats) SetIncludedSeats(v int32)`

SetIncludedSeats sets IncludedSeats field to given value.


### SetIncludedSeatsNil

`func (o *GetApiWorkspacesOverview200ResponseSeats) SetIncludedSeatsNil(b bool)`

 SetIncludedSeatsNil sets the value for IncludedSeats to be an explicit nil

### UnsetIncludedSeats
`func (o *GetApiWorkspacesOverview200ResponseSeats) UnsetIncludedSeats()`

UnsetIncludedSeats ensures that no value is present for IncludedSeats, not even an explicit nil
### GetOveragePriceCentsPerSeat

`func (o *GetApiWorkspacesOverview200ResponseSeats) GetOveragePriceCentsPerSeat() float32`

GetOveragePriceCentsPerSeat returns the OveragePriceCentsPerSeat field if non-nil, zero value otherwise.

### GetOveragePriceCentsPerSeatOk

`func (o *GetApiWorkspacesOverview200ResponseSeats) GetOveragePriceCentsPerSeatOk() (*float32, bool)`

GetOveragePriceCentsPerSeatOk returns a tuple with the OveragePriceCentsPerSeat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOveragePriceCentsPerSeat

`func (o *GetApiWorkspacesOverview200ResponseSeats) SetOveragePriceCentsPerSeat(v float32)`

SetOveragePriceCentsPerSeat sets OveragePriceCentsPerSeat field to given value.


### SetOveragePriceCentsPerSeatNil

`func (o *GetApiWorkspacesOverview200ResponseSeats) SetOveragePriceCentsPerSeatNil(b bool)`

 SetOveragePriceCentsPerSeatNil sets the value for OveragePriceCentsPerSeat to be an explicit nil

### UnsetOveragePriceCentsPerSeat
`func (o *GetApiWorkspacesOverview200ResponseSeats) UnsetOveragePriceCentsPerSeat()`

UnsetOveragePriceCentsPerSeat ensures that no value is present for OveragePriceCentsPerSeat, not even an explicit nil
### GetOverage

`func (o *GetApiWorkspacesOverview200ResponseSeats) GetOverage() GetApiWorkspacesOverview200ResponseSeatsOverage`

GetOverage returns the Overage field if non-nil, zero value otherwise.

### GetOverageOk

`func (o *GetApiWorkspacesOverview200ResponseSeats) GetOverageOk() (*GetApiWorkspacesOverview200ResponseSeatsOverage, bool)`

GetOverageOk returns a tuple with the Overage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverage

`func (o *GetApiWorkspacesOverview200ResponseSeats) SetOverage(v GetApiWorkspacesOverview200ResponseSeatsOverage)`

SetOverage sets Overage field to given value.


### GetAllowanceExceeded

`func (o *GetApiWorkspacesOverview200ResponseSeats) GetAllowanceExceeded() bool`

GetAllowanceExceeded returns the AllowanceExceeded field if non-nil, zero value otherwise.

### GetAllowanceExceededOk

`func (o *GetApiWorkspacesOverview200ResponseSeats) GetAllowanceExceededOk() (*bool, bool)`

GetAllowanceExceededOk returns a tuple with the AllowanceExceeded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowanceExceeded

`func (o *GetApiWorkspacesOverview200ResponseSeats) SetAllowanceExceeded(v bool)`

SetAllowanceExceeded sets AllowanceExceeded field to given value.


### GetAllowanceNear

`func (o *GetApiWorkspacesOverview200ResponseSeats) GetAllowanceNear() bool`

GetAllowanceNear returns the AllowanceNear field if non-nil, zero value otherwise.

### GetAllowanceNearOk

`func (o *GetApiWorkspacesOverview200ResponseSeats) GetAllowanceNearOk() (*bool, bool)`

GetAllowanceNearOk returns a tuple with the AllowanceNear field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowanceNear

`func (o *GetApiWorkspacesOverview200ResponseSeats) SetAllowanceNear(v bool)`

SetAllowanceNear sets AllowanceNear field to given value.


### GetProjected

`func (o *GetApiWorkspacesOverview200ResponseSeats) GetProjected() GetApiWorkspacesOverview200ResponseSeatsProjected`

GetProjected returns the Projected field if non-nil, zero value otherwise.

### GetProjectedOk

`func (o *GetApiWorkspacesOverview200ResponseSeats) GetProjectedOk() (*GetApiWorkspacesOverview200ResponseSeatsProjected, bool)`

GetProjectedOk returns a tuple with the Projected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjected

`func (o *GetApiWorkspacesOverview200ResponseSeats) SetProjected(v GetApiWorkspacesOverview200ResponseSeatsProjected)`

SetProjected sets Projected field to given value.

### HasProjected

`func (o *GetApiWorkspacesOverview200ResponseSeats) HasProjected() bool`

HasProjected returns a boolean if a field has been set.

### SetProjectedNil

`func (o *GetApiWorkspacesOverview200ResponseSeats) SetProjectedNil(b bool)`

 SetProjectedNil sets the value for Projected to be an explicit nil

### UnsetProjected
`func (o *GetApiWorkspacesOverview200ResponseSeats) UnsetProjected()`

UnsetProjected ensures that no value is present for Projected, not even an explicit nil
### GetBillingOwnerUserId

`func (o *GetApiWorkspacesOverview200ResponseSeats) GetBillingOwnerUserId() string`

GetBillingOwnerUserId returns the BillingOwnerUserId field if non-nil, zero value otherwise.

### GetBillingOwnerUserIdOk

`func (o *GetApiWorkspacesOverview200ResponseSeats) GetBillingOwnerUserIdOk() (*string, bool)`

GetBillingOwnerUserIdOk returns a tuple with the BillingOwnerUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingOwnerUserId

`func (o *GetApiWorkspacesOverview200ResponseSeats) SetBillingOwnerUserId(v string)`

SetBillingOwnerUserId sets BillingOwnerUserId field to given value.


### SetBillingOwnerUserIdNil

`func (o *GetApiWorkspacesOverview200ResponseSeats) SetBillingOwnerUserIdNil(b bool)`

 SetBillingOwnerUserIdNil sets the value for BillingOwnerUserId to be an explicit nil

### UnsetBillingOwnerUserId
`func (o *GetApiWorkspacesOverview200ResponseSeats) UnsetBillingOwnerUserId()`

UnsetBillingOwnerUserId ensures that no value is present for BillingOwnerUserId, not even an explicit nil
### GetRolledUp

`func (o *GetApiWorkspacesOverview200ResponseSeats) GetRolledUp() bool`

GetRolledUp returns the RolledUp field if non-nil, zero value otherwise.

### GetRolledUpOk

`func (o *GetApiWorkspacesOverview200ResponseSeats) GetRolledUpOk() (*bool, bool)`

GetRolledUpOk returns a tuple with the RolledUp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRolledUp

`func (o *GetApiWorkspacesOverview200ResponseSeats) SetRolledUp(v bool)`

SetRolledUp sets RolledUp field to given value.


### GetWorkspaceCount

`func (o *GetApiWorkspacesOverview200ResponseSeats) GetWorkspaceCount() int32`

GetWorkspaceCount returns the WorkspaceCount field if non-nil, zero value otherwise.

### GetWorkspaceCountOk

`func (o *GetApiWorkspacesOverview200ResponseSeats) GetWorkspaceCountOk() (*int32, bool)`

GetWorkspaceCountOk returns a tuple with the WorkspaceCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkspaceCount

`func (o *GetApiWorkspacesOverview200ResponseSeats) SetWorkspaceCount(v int32)`

SetWorkspaceCount sets WorkspaceCount field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


