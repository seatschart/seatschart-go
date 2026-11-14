# PostApiEventsByEventIdActionsChangeObjectStatusRequestListing

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Key** | Pointer to **string** | Listing key (default: &#x60;resaleListingId&#x60;, else a generated UUID). | [optional] 
**Type** | Pointer to **NullableString** | Listing type key; omitted &#x3D; untyped (resale). | [optional] 
**SellAsWhole** | Pointer to **bool** | Default true: the objects are held and sold together. | [optional] 
**PriceCents** | Pointer to **NullableInt32** | Price per object (default: &#x60;resalePriceCents&#x60;). | [optional] 
**TotalPriceCents** | Pointer to **NullableInt32** |  | [optional] 
**PartnerMetadata** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewPostApiEventsByEventIdActionsChangeObjectStatusRequestListing

`func NewPostApiEventsByEventIdActionsChangeObjectStatusRequestListing() *PostApiEventsByEventIdActionsChangeObjectStatusRequestListing`

NewPostApiEventsByEventIdActionsChangeObjectStatusRequestListing instantiates a new PostApiEventsByEventIdActionsChangeObjectStatusRequestListing object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostApiEventsByEventIdActionsChangeObjectStatusRequestListingWithDefaults

`func NewPostApiEventsByEventIdActionsChangeObjectStatusRequestListingWithDefaults() *PostApiEventsByEventIdActionsChangeObjectStatusRequestListing`

NewPostApiEventsByEventIdActionsChangeObjectStatusRequestListingWithDefaults instantiates a new PostApiEventsByEventIdActionsChangeObjectStatusRequestListing object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKey

`func (o *PostApiEventsByEventIdActionsChangeObjectStatusRequestListing) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *PostApiEventsByEventIdActionsChangeObjectStatusRequestListing) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *PostApiEventsByEventIdActionsChangeObjectStatusRequestListing) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *PostApiEventsByEventIdActionsChangeObjectStatusRequestListing) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetType

`func (o *PostApiEventsByEventIdActionsChangeObjectStatusRequestListing) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PostApiEventsByEventIdActionsChangeObjectStatusRequestListing) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PostApiEventsByEventIdActionsChangeObjectStatusRequestListing) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *PostApiEventsByEventIdActionsChangeObjectStatusRequestListing) HasType() bool`

HasType returns a boolean if a field has been set.

### SetTypeNil

`func (o *PostApiEventsByEventIdActionsChangeObjectStatusRequestListing) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *PostApiEventsByEventIdActionsChangeObjectStatusRequestListing) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetSellAsWhole

`func (o *PostApiEventsByEventIdActionsChangeObjectStatusRequestListing) GetSellAsWhole() bool`

GetSellAsWhole returns the SellAsWhole field if non-nil, zero value otherwise.

### GetSellAsWholeOk

`func (o *PostApiEventsByEventIdActionsChangeObjectStatusRequestListing) GetSellAsWholeOk() (*bool, bool)`

GetSellAsWholeOk returns a tuple with the SellAsWhole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSellAsWhole

`func (o *PostApiEventsByEventIdActionsChangeObjectStatusRequestListing) SetSellAsWhole(v bool)`

SetSellAsWhole sets SellAsWhole field to given value.

### HasSellAsWhole

`func (o *PostApiEventsByEventIdActionsChangeObjectStatusRequestListing) HasSellAsWhole() bool`

HasSellAsWhole returns a boolean if a field has been set.

### GetPriceCents

`func (o *PostApiEventsByEventIdActionsChangeObjectStatusRequestListing) GetPriceCents() int32`

GetPriceCents returns the PriceCents field if non-nil, zero value otherwise.

### GetPriceCentsOk

`func (o *PostApiEventsByEventIdActionsChangeObjectStatusRequestListing) GetPriceCentsOk() (*int32, bool)`

GetPriceCentsOk returns a tuple with the PriceCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceCents

`func (o *PostApiEventsByEventIdActionsChangeObjectStatusRequestListing) SetPriceCents(v int32)`

SetPriceCents sets PriceCents field to given value.

### HasPriceCents

`func (o *PostApiEventsByEventIdActionsChangeObjectStatusRequestListing) HasPriceCents() bool`

HasPriceCents returns a boolean if a field has been set.

### SetPriceCentsNil

`func (o *PostApiEventsByEventIdActionsChangeObjectStatusRequestListing) SetPriceCentsNil(b bool)`

 SetPriceCentsNil sets the value for PriceCents to be an explicit nil

### UnsetPriceCents
`func (o *PostApiEventsByEventIdActionsChangeObjectStatusRequestListing) UnsetPriceCents()`

UnsetPriceCents ensures that no value is present for PriceCents, not even an explicit nil
### GetTotalPriceCents

`func (o *PostApiEventsByEventIdActionsChangeObjectStatusRequestListing) GetTotalPriceCents() int32`

GetTotalPriceCents returns the TotalPriceCents field if non-nil, zero value otherwise.

### GetTotalPriceCentsOk

`func (o *PostApiEventsByEventIdActionsChangeObjectStatusRequestListing) GetTotalPriceCentsOk() (*int32, bool)`

GetTotalPriceCentsOk returns a tuple with the TotalPriceCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalPriceCents

`func (o *PostApiEventsByEventIdActionsChangeObjectStatusRequestListing) SetTotalPriceCents(v int32)`

SetTotalPriceCents sets TotalPriceCents field to given value.

### HasTotalPriceCents

`func (o *PostApiEventsByEventIdActionsChangeObjectStatusRequestListing) HasTotalPriceCents() bool`

HasTotalPriceCents returns a boolean if a field has been set.

### SetTotalPriceCentsNil

`func (o *PostApiEventsByEventIdActionsChangeObjectStatusRequestListing) SetTotalPriceCentsNil(b bool)`

 SetTotalPriceCentsNil sets the value for TotalPriceCents to be an explicit nil

### UnsetTotalPriceCents
`func (o *PostApiEventsByEventIdActionsChangeObjectStatusRequestListing) UnsetTotalPriceCents()`

UnsetTotalPriceCents ensures that no value is present for TotalPriceCents, not even an explicit nil
### GetPartnerMetadata

`func (o *PostApiEventsByEventIdActionsChangeObjectStatusRequestListing) GetPartnerMetadata() map[string]interface{}`

GetPartnerMetadata returns the PartnerMetadata field if non-nil, zero value otherwise.

### GetPartnerMetadataOk

`func (o *PostApiEventsByEventIdActionsChangeObjectStatusRequestListing) GetPartnerMetadataOk() (*map[string]interface{}, bool)`

GetPartnerMetadataOk returns a tuple with the PartnerMetadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartnerMetadata

`func (o *PostApiEventsByEventIdActionsChangeObjectStatusRequestListing) SetPartnerMetadata(v map[string]interface{})`

SetPartnerMetadata sets PartnerMetadata field to given value.

### HasPartnerMetadata

`func (o *PostApiEventsByEventIdActionsChangeObjectStatusRequestListing) HasPartnerMetadata() bool`

HasPartnerMetadata returns a boolean if a field has been set.

### SetPartnerMetadataNil

`func (o *PostApiEventsByEventIdActionsChangeObjectStatusRequestListing) SetPartnerMetadataNil(b bool)`

 SetPartnerMetadataNil sets the value for PartnerMetadata to be an explicit nil

### UnsetPartnerMetadata
`func (o *PostApiEventsByEventIdActionsChangeObjectStatusRequestListing) UnsetPartnerMetadata()`

UnsetPartnerMetadata ensures that no value is present for PartnerMetadata, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


