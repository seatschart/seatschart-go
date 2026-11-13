# PostApiEventsByEventIdListingsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Key** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **NullableString** | Listing type key; omitted &#x3D; untyped (resale). | [optional] 
**Objects** | [**[]PostApiEventsByEventIdListingsRequestObjectsInner**](PostApiEventsByEventIdListingsRequestObjectsInner.md) |  | 
**SellAsWhole** | Pointer to **bool** |  | [optional] 
**PriceCents** | Pointer to **NullableInt32** |  | [optional] 
**TotalPriceCents** | Pointer to **NullableInt32** |  | [optional] 
**PartnerMetadata** | Pointer to **map[string]interface{}** | Partner data (≤ 4 KiB JSON), returned as is. | [optional] 

## Methods

### NewPostApiEventsByEventIdListingsRequest

`func NewPostApiEventsByEventIdListingsRequest(objects []PostApiEventsByEventIdListingsRequestObjectsInner, ) *PostApiEventsByEventIdListingsRequest`

NewPostApiEventsByEventIdListingsRequest instantiates a new PostApiEventsByEventIdListingsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostApiEventsByEventIdListingsRequestWithDefaults

`func NewPostApiEventsByEventIdListingsRequestWithDefaults() *PostApiEventsByEventIdListingsRequest`

NewPostApiEventsByEventIdListingsRequestWithDefaults instantiates a new PostApiEventsByEventIdListingsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKey

`func (o *PostApiEventsByEventIdListingsRequest) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *PostApiEventsByEventIdListingsRequest) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *PostApiEventsByEventIdListingsRequest) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *PostApiEventsByEventIdListingsRequest) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetType

`func (o *PostApiEventsByEventIdListingsRequest) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PostApiEventsByEventIdListingsRequest) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PostApiEventsByEventIdListingsRequest) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *PostApiEventsByEventIdListingsRequest) HasType() bool`

HasType returns a boolean if a field has been set.

### SetTypeNil

`func (o *PostApiEventsByEventIdListingsRequest) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *PostApiEventsByEventIdListingsRequest) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetObjects

`func (o *PostApiEventsByEventIdListingsRequest) GetObjects() []PostApiEventsByEventIdListingsRequestObjectsInner`

GetObjects returns the Objects field if non-nil, zero value otherwise.

### GetObjectsOk

`func (o *PostApiEventsByEventIdListingsRequest) GetObjectsOk() (*[]PostApiEventsByEventIdListingsRequestObjectsInner, bool)`

GetObjectsOk returns a tuple with the Objects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjects

`func (o *PostApiEventsByEventIdListingsRequest) SetObjects(v []PostApiEventsByEventIdListingsRequestObjectsInner)`

SetObjects sets Objects field to given value.


### GetSellAsWhole

`func (o *PostApiEventsByEventIdListingsRequest) GetSellAsWhole() bool`

GetSellAsWhole returns the SellAsWhole field if non-nil, zero value otherwise.

### GetSellAsWholeOk

`func (o *PostApiEventsByEventIdListingsRequest) GetSellAsWholeOk() (*bool, bool)`

GetSellAsWholeOk returns a tuple with the SellAsWhole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSellAsWhole

`func (o *PostApiEventsByEventIdListingsRequest) SetSellAsWhole(v bool)`

SetSellAsWhole sets SellAsWhole field to given value.

### HasSellAsWhole

`func (o *PostApiEventsByEventIdListingsRequest) HasSellAsWhole() bool`

HasSellAsWhole returns a boolean if a field has been set.

### GetPriceCents

`func (o *PostApiEventsByEventIdListingsRequest) GetPriceCents() int32`

GetPriceCents returns the PriceCents field if non-nil, zero value otherwise.

### GetPriceCentsOk

`func (o *PostApiEventsByEventIdListingsRequest) GetPriceCentsOk() (*int32, bool)`

GetPriceCentsOk returns a tuple with the PriceCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceCents

`func (o *PostApiEventsByEventIdListingsRequest) SetPriceCents(v int32)`

SetPriceCents sets PriceCents field to given value.

### HasPriceCents

`func (o *PostApiEventsByEventIdListingsRequest) HasPriceCents() bool`

HasPriceCents returns a boolean if a field has been set.

### SetPriceCentsNil

`func (o *PostApiEventsByEventIdListingsRequest) SetPriceCentsNil(b bool)`

 SetPriceCentsNil sets the value for PriceCents to be an explicit nil

### UnsetPriceCents
`func (o *PostApiEventsByEventIdListingsRequest) UnsetPriceCents()`

UnsetPriceCents ensures that no value is present for PriceCents, not even an explicit nil
### GetTotalPriceCents

`func (o *PostApiEventsByEventIdListingsRequest) GetTotalPriceCents() int32`

GetTotalPriceCents returns the TotalPriceCents field if non-nil, zero value otherwise.

### GetTotalPriceCentsOk

`func (o *PostApiEventsByEventIdListingsRequest) GetTotalPriceCentsOk() (*int32, bool)`

GetTotalPriceCentsOk returns a tuple with the TotalPriceCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalPriceCents

`func (o *PostApiEventsByEventIdListingsRequest) SetTotalPriceCents(v int32)`

SetTotalPriceCents sets TotalPriceCents field to given value.

### HasTotalPriceCents

`func (o *PostApiEventsByEventIdListingsRequest) HasTotalPriceCents() bool`

HasTotalPriceCents returns a boolean if a field has been set.

### SetTotalPriceCentsNil

`func (o *PostApiEventsByEventIdListingsRequest) SetTotalPriceCentsNil(b bool)`

 SetTotalPriceCentsNil sets the value for TotalPriceCents to be an explicit nil

### UnsetTotalPriceCents
`func (o *PostApiEventsByEventIdListingsRequest) UnsetTotalPriceCents()`

UnsetTotalPriceCents ensures that no value is present for TotalPriceCents, not even an explicit nil
### GetPartnerMetadata

`func (o *PostApiEventsByEventIdListingsRequest) GetPartnerMetadata() map[string]interface{}`

GetPartnerMetadata returns the PartnerMetadata field if non-nil, zero value otherwise.

### GetPartnerMetadataOk

`func (o *PostApiEventsByEventIdListingsRequest) GetPartnerMetadataOk() (*map[string]interface{}, bool)`

GetPartnerMetadataOk returns a tuple with the PartnerMetadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartnerMetadata

`func (o *PostApiEventsByEventIdListingsRequest) SetPartnerMetadata(v map[string]interface{})`

SetPartnerMetadata sets PartnerMetadata field to given value.

### HasPartnerMetadata

`func (o *PostApiEventsByEventIdListingsRequest) HasPartnerMetadata() bool`

HasPartnerMetadata returns a boolean if a field has been set.

### SetPartnerMetadataNil

`func (o *PostApiEventsByEventIdListingsRequest) SetPartnerMetadataNil(b bool)`

 SetPartnerMetadataNil sets the value for PartnerMetadata to be an explicit nil

### UnsetPartnerMetadata
`func (o *PostApiEventsByEventIdListingsRequest) UnsetPartnerMetadata()`

UnsetPartnerMetadata ensures that no value is present for PartnerMetadata, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


