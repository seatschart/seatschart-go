# PatchApiWorkspaceSettingsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HoldPeriodMinutes** | Pointer to **NullableInt32** | Minutes (1–120); &#x60;null&#x60; &#x3D; use the default. | [optional] 
**MaxSeatsPerHoldToken** | Pointer to **NullableInt32** | Seats per hold token (1–64); &#x60;null&#x60; &#x3D; use the default. | [optional] 
**TicketBuyerValidation** | Pointer to **NullableBool** | &#x60;null&#x60; &#x3D; use the default (off). | [optional] 
**GaPlacesToChannels** | Pointer to **NullableBool** | &#x60;null&#x60; &#x3D; use the default (on). | [optional] 
**NewEventsNotForSale** | Pointer to **NullableBool** | New events start not for sale; &#x60;null&#x60; &#x3D; use the default (off). | [optional] 
**RequireTwoFactor** | Pointer to **bool** | Members must sign in with two-factor authentication (enforced at login with a guided enrolment). Dashboard session only; the caller must have 2FA enabled. | [optional] 

## Methods

### NewPatchApiWorkspaceSettingsRequest

`func NewPatchApiWorkspaceSettingsRequest() *PatchApiWorkspaceSettingsRequest`

NewPatchApiWorkspaceSettingsRequest instantiates a new PatchApiWorkspaceSettingsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchApiWorkspaceSettingsRequestWithDefaults

`func NewPatchApiWorkspaceSettingsRequestWithDefaults() *PatchApiWorkspaceSettingsRequest`

NewPatchApiWorkspaceSettingsRequestWithDefaults instantiates a new PatchApiWorkspaceSettingsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHoldPeriodMinutes

`func (o *PatchApiWorkspaceSettingsRequest) GetHoldPeriodMinutes() int32`

GetHoldPeriodMinutes returns the HoldPeriodMinutes field if non-nil, zero value otherwise.

### GetHoldPeriodMinutesOk

`func (o *PatchApiWorkspaceSettingsRequest) GetHoldPeriodMinutesOk() (*int32, bool)`

GetHoldPeriodMinutesOk returns a tuple with the HoldPeriodMinutes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHoldPeriodMinutes

`func (o *PatchApiWorkspaceSettingsRequest) SetHoldPeriodMinutes(v int32)`

SetHoldPeriodMinutes sets HoldPeriodMinutes field to given value.

### HasHoldPeriodMinutes

`func (o *PatchApiWorkspaceSettingsRequest) HasHoldPeriodMinutes() bool`

HasHoldPeriodMinutes returns a boolean if a field has been set.

### SetHoldPeriodMinutesNil

`func (o *PatchApiWorkspaceSettingsRequest) SetHoldPeriodMinutesNil(b bool)`

 SetHoldPeriodMinutesNil sets the value for HoldPeriodMinutes to be an explicit nil

### UnsetHoldPeriodMinutes
`func (o *PatchApiWorkspaceSettingsRequest) UnsetHoldPeriodMinutes()`

UnsetHoldPeriodMinutes ensures that no value is present for HoldPeriodMinutes, not even an explicit nil
### GetMaxSeatsPerHoldToken

`func (o *PatchApiWorkspaceSettingsRequest) GetMaxSeatsPerHoldToken() int32`

GetMaxSeatsPerHoldToken returns the MaxSeatsPerHoldToken field if non-nil, zero value otherwise.

### GetMaxSeatsPerHoldTokenOk

`func (o *PatchApiWorkspaceSettingsRequest) GetMaxSeatsPerHoldTokenOk() (*int32, bool)`

GetMaxSeatsPerHoldTokenOk returns a tuple with the MaxSeatsPerHoldToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxSeatsPerHoldToken

`func (o *PatchApiWorkspaceSettingsRequest) SetMaxSeatsPerHoldToken(v int32)`

SetMaxSeatsPerHoldToken sets MaxSeatsPerHoldToken field to given value.

### HasMaxSeatsPerHoldToken

`func (o *PatchApiWorkspaceSettingsRequest) HasMaxSeatsPerHoldToken() bool`

HasMaxSeatsPerHoldToken returns a boolean if a field has been set.

### SetMaxSeatsPerHoldTokenNil

`func (o *PatchApiWorkspaceSettingsRequest) SetMaxSeatsPerHoldTokenNil(b bool)`

 SetMaxSeatsPerHoldTokenNil sets the value for MaxSeatsPerHoldToken to be an explicit nil

### UnsetMaxSeatsPerHoldToken
`func (o *PatchApiWorkspaceSettingsRequest) UnsetMaxSeatsPerHoldToken()`

UnsetMaxSeatsPerHoldToken ensures that no value is present for MaxSeatsPerHoldToken, not even an explicit nil
### GetTicketBuyerValidation

`func (o *PatchApiWorkspaceSettingsRequest) GetTicketBuyerValidation() bool`

GetTicketBuyerValidation returns the TicketBuyerValidation field if non-nil, zero value otherwise.

### GetTicketBuyerValidationOk

`func (o *PatchApiWorkspaceSettingsRequest) GetTicketBuyerValidationOk() (*bool, bool)`

GetTicketBuyerValidationOk returns a tuple with the TicketBuyerValidation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTicketBuyerValidation

`func (o *PatchApiWorkspaceSettingsRequest) SetTicketBuyerValidation(v bool)`

SetTicketBuyerValidation sets TicketBuyerValidation field to given value.

### HasTicketBuyerValidation

`func (o *PatchApiWorkspaceSettingsRequest) HasTicketBuyerValidation() bool`

HasTicketBuyerValidation returns a boolean if a field has been set.

### SetTicketBuyerValidationNil

`func (o *PatchApiWorkspaceSettingsRequest) SetTicketBuyerValidationNil(b bool)`

 SetTicketBuyerValidationNil sets the value for TicketBuyerValidation to be an explicit nil

### UnsetTicketBuyerValidation
`func (o *PatchApiWorkspaceSettingsRequest) UnsetTicketBuyerValidation()`

UnsetTicketBuyerValidation ensures that no value is present for TicketBuyerValidation, not even an explicit nil
### GetGaPlacesToChannels

`func (o *PatchApiWorkspaceSettingsRequest) GetGaPlacesToChannels() bool`

GetGaPlacesToChannels returns the GaPlacesToChannels field if non-nil, zero value otherwise.

### GetGaPlacesToChannelsOk

`func (o *PatchApiWorkspaceSettingsRequest) GetGaPlacesToChannelsOk() (*bool, bool)`

GetGaPlacesToChannelsOk returns a tuple with the GaPlacesToChannels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGaPlacesToChannels

`func (o *PatchApiWorkspaceSettingsRequest) SetGaPlacesToChannels(v bool)`

SetGaPlacesToChannels sets GaPlacesToChannels field to given value.

### HasGaPlacesToChannels

`func (o *PatchApiWorkspaceSettingsRequest) HasGaPlacesToChannels() bool`

HasGaPlacesToChannels returns a boolean if a field has been set.

### SetGaPlacesToChannelsNil

`func (o *PatchApiWorkspaceSettingsRequest) SetGaPlacesToChannelsNil(b bool)`

 SetGaPlacesToChannelsNil sets the value for GaPlacesToChannels to be an explicit nil

### UnsetGaPlacesToChannels
`func (o *PatchApiWorkspaceSettingsRequest) UnsetGaPlacesToChannels()`

UnsetGaPlacesToChannels ensures that no value is present for GaPlacesToChannels, not even an explicit nil
### GetNewEventsNotForSale

`func (o *PatchApiWorkspaceSettingsRequest) GetNewEventsNotForSale() bool`

GetNewEventsNotForSale returns the NewEventsNotForSale field if non-nil, zero value otherwise.

### GetNewEventsNotForSaleOk

`func (o *PatchApiWorkspaceSettingsRequest) GetNewEventsNotForSaleOk() (*bool, bool)`

GetNewEventsNotForSaleOk returns a tuple with the NewEventsNotForSale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewEventsNotForSale

`func (o *PatchApiWorkspaceSettingsRequest) SetNewEventsNotForSale(v bool)`

SetNewEventsNotForSale sets NewEventsNotForSale field to given value.

### HasNewEventsNotForSale

`func (o *PatchApiWorkspaceSettingsRequest) HasNewEventsNotForSale() bool`

HasNewEventsNotForSale returns a boolean if a field has been set.

### SetNewEventsNotForSaleNil

`func (o *PatchApiWorkspaceSettingsRequest) SetNewEventsNotForSaleNil(b bool)`

 SetNewEventsNotForSaleNil sets the value for NewEventsNotForSale to be an explicit nil

### UnsetNewEventsNotForSale
`func (o *PatchApiWorkspaceSettingsRequest) UnsetNewEventsNotForSale()`

UnsetNewEventsNotForSale ensures that no value is present for NewEventsNotForSale, not even an explicit nil
### GetRequireTwoFactor

`func (o *PatchApiWorkspaceSettingsRequest) GetRequireTwoFactor() bool`

GetRequireTwoFactor returns the RequireTwoFactor field if non-nil, zero value otherwise.

### GetRequireTwoFactorOk

`func (o *PatchApiWorkspaceSettingsRequest) GetRequireTwoFactorOk() (*bool, bool)`

GetRequireTwoFactorOk returns a tuple with the RequireTwoFactor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequireTwoFactor

`func (o *PatchApiWorkspaceSettingsRequest) SetRequireTwoFactor(v bool)`

SetRequireTwoFactor sets RequireTwoFactor field to given value.

### HasRequireTwoFactor

`func (o *PatchApiWorkspaceSettingsRequest) HasRequireTwoFactor() bool`

HasRequireTwoFactor returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


