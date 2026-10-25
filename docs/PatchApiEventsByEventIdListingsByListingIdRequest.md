# PatchApiEventsByEventIdListingsByListingIdRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **NullableString** |  | [optional] 
**Objects** | Pointer to [**[]PostApiEventsByEventIdListingsRequestObjectsInner**](PostApiEventsByEventIdListingsRequestObjectsInner.md) |  | [optional] 
**SellAsWhole** | Pointer to **bool** |  | [optional] 
**PriceCents** | Pointer to **NullableInt32** |  | [optional] 
**TotalPriceCents** | Pointer to **NullableInt32** |  | [optional] 
**PartnerMetadata** | Pointer to **map[string]interface{}** | Partner data (≤ 4 KiB JSON), returned as is. | [optional] 

## Methods

### NewPatchApiEventsByEventIdListingsByListingIdRequest

`func NewPatchApiEventsByEventIdListingsByListingIdRequest() *PatchApiEventsByEventIdListingsByListingIdRequest`

NewPatchApiEventsByEventIdListingsByListingIdRequest instantiates a new PatchApiEventsByEventIdListingsByListingIdRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchApiEventsByEventIdListingsByListingIdRequestWithDefaults

`func NewPatchApiEventsByEventIdListingsByListingIdRequestWithDefaults() *PatchApiEventsByEventIdListingsByListingIdRequest`

NewPatchApiEventsByEventIdListingsByListingIdRequestWithDefaults instantiates a new PatchApiEventsByEventIdListingsByListingIdRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *PatchApiEventsByEventIdListingsByListingIdRequest) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PatchApiEventsByEventIdListingsByListingIdRequest) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PatchApiEventsByEventIdListingsByListingIdRequest) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *PatchApiEventsByEventIdListingsByListingIdRequest) HasType() bool`

HasType returns a boolean if a field has been set.

### SetTypeNil

`func (o *PatchApiEventsByEventIdListingsByListingIdRequest) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *PatchApiEventsByEventIdListingsByListingIdRequest) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetObjects

`func (o *PatchApiEventsByEventIdListingsByListingIdRequest) GetObjects() []PostApiEventsByEventIdListingsRequestObjectsInner`

GetObjects returns the Objects field if non-nil, zero value otherwise.

### GetObjectsOk

`func (o *PatchApiEventsByEventIdListingsByListingIdRequest) GetObjectsOk() (*[]PostApiEventsByEventIdListingsRequestObjectsInner, bool)`

GetObjectsOk returns a tuple with the Objects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjects

`func (o *PatchApiEventsByEventIdListingsByListingIdRequest) SetObjects(v []PostApiEventsByEventIdListingsRequestObjectsInner)`

SetObjects sets Objects field to given value.

### HasObjects

`func (o *PatchApiEventsByEventIdListingsByListingIdRequest) HasObjects() bool`

HasObjects returns a boolean if a field has been set.

### GetSellAsWhole

`func (o *PatchApiEventsByEventIdListingsByListingIdRequest) GetSellAsWhole() bool`

GetSellAsWhole returns the SellAsWhole field if non-nil, zero value otherwise.

### GetSellAsWholeOk

`func (o *PatchApiEventsByEventIdListingsByListingIdRequest) GetSellAsWholeOk() (*bool, bool)`

GetSellAsWholeOk returns a tuple with the SellAsWhole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSellAsWhole

`func (o *PatchApiEventsByEventIdListingsByListingIdRequest) SetSellAsWhole(v bool)`

SetSellAsWhole sets SellAsWhole field to given value.

### HasSellAsWhole

`func (o *PatchApiEventsByEventIdListingsByListingIdRequest) HasSellAsWhole() bool`

HasSellAsWhole returns a boolean if a field has been set.

### GetPriceCents

`func (o *PatchApiEventsByEventIdListingsByListingIdRequest) GetPriceCents() int32`

GetPriceCents returns the PriceCents field if non-nil, zero value otherwise.

### GetPriceCentsOk

`func (o *PatchApiEventsByEventIdListingsByListingIdRequest) GetPriceCentsOk() (*int32, bool)`

GetPriceCentsOk returns a tuple with the PriceCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceCents

`func (o *PatchApiEventsByEventIdListingsByListingIdRequest) SetPriceCents(v int32)`

SetPriceCents sets PriceCents field to given value.

### HasPriceCents

`func (o *PatchApiEventsByEventIdListingsByListingIdRequest) HasPriceCents() bool`

HasPriceCents returns a boolean if a field has been set.

### SetPriceCentsNil

`func (o *PatchApiEventsByEventIdListingsByListingIdRequest) SetPriceCentsNil(b bool)`

 SetPriceCentsNil sets the value for PriceCents to be an explicit nil

### UnsetPriceCents
`func (o *PatchApiEventsByEventIdListingsByListingIdRequest) UnsetPriceCents()`

UnsetPriceCents ensures that no value is present for PriceCents, not even an explicit nil
### GetTotalPriceCents

`func (o *PatchApiEventsByEventIdListingsByListingIdRequest) GetTotalPriceCents() int32`

GetTotalPriceCents returns the TotalPriceCents field if non-nil, zero value otherwise.

### GetTotalPriceCentsOk

`func (o *PatchApiEventsByEventIdListingsByListingIdRequest) GetTotalPriceCentsOk() (*int32, bool)`

GetTotalPriceCentsOk returns a tuple with the TotalPriceCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalPriceCents

`func (o *PatchApiEventsByEventIdListingsByListingIdRequest) SetTotalPriceCents(v int32)`

SetTotalPriceCents sets TotalPriceCents field to given value.

### HasTotalPriceCents

`func (o *PatchApiEventsByEventIdListingsByListingIdRequest) HasTotalPriceCents() bool`

HasTotalPriceCents returns a boolean if a field has been set.

### SetTotalPriceCentsNil

`func (o *PatchApiEventsByEventIdListingsByListingIdRequest) SetTotalPriceCentsNil(b bool)`

 SetTotalPriceCentsNil sets the value for TotalPriceCents to be an explicit nil

### UnsetTotalPriceCents
`func (o *PatchApiEventsByEventIdListingsByListingIdRequest) UnsetTotalPriceCents()`

UnsetTotalPriceCents ensures that no value is present for TotalPriceCents, not even an explicit nil
### GetPartnerMetadata

`func (o *PatchApiEventsByEventIdListingsByListingIdRequest) GetPartnerMetadata() map[string]interface{}`

GetPartnerMetadata returns the PartnerMetadata field if non-nil, zero value otherwise.

### GetPartnerMetadataOk

`func (o *PatchApiEventsByEventIdListingsByListingIdRequest) GetPartnerMetadataOk() (*map[string]interface{}, bool)`

GetPartnerMetadataOk returns a tuple with the PartnerMetadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartnerMetadata

`func (o *PatchApiEventsByEventIdListingsByListingIdRequest) SetPartnerMetadata(v map[string]interface{})`

SetPartnerMetadata sets PartnerMetadata field to given value.

### HasPartnerMetadata

`func (o *PatchApiEventsByEventIdListingsByListingIdRequest) HasPartnerMetadata() bool`

HasPartnerMetadata returns a boolean if a field has been set.

### SetPartnerMetadataNil

`func (o *PatchApiEventsByEventIdListingsByListingIdRequest) SetPartnerMetadataNil(b bool)`

 SetPartnerMetadataNil sets the value for PartnerMetadata to be an explicit nil

### UnsetPartnerMetadata
`func (o *PatchApiEventsByEventIdListingsByListingIdRequest) UnsetPartnerMetadata()`

UnsetPartnerMetadata ensures that no value is present for PartnerMetadata, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


