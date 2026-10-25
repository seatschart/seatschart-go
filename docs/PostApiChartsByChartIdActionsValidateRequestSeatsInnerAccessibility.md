# PostApiChartsByChartIdActionsValidateRequestSeatsInnerAccessibility

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Wheelchair** | Pointer to **bool** | Wheelchair space. | [optional] 
**Companion** | Pointer to **bool** | Companion seat next to a wheelchair space. | [optional] 
**RestrictedView** | Pointer to **bool** | Partially restricted view. | [optional] 
**ObstructedView** | Pointer to **bool** | View obstructed (e.g. by a pillar). | [optional] 
**ViewNote** | Pointer to **string** | Short buyer-facing note about the view. | [optional] 
**Amenities** | Pointer to **int32** | Compact bitmask of the other accessibility and amenity flags: 16 wheelchair space without a seat in place (requires wheelchair), 32 semi-ambulatory, 64 plus-size, 128 hearing impaired (CART services), 256 hearing impaired (sign language interpretation), 512 not suitable for vertigo sufferers, 1024 limited legroom, 2048 lift-up armrests. Bits 1/2/4/8 mirror wheelchair / companion / restrictedView / obstructedView, which are always stored and returned as booleans. | [optional] 
**SeatType** | Pointer to **string** | Seat display type. \&quot;chair\&quot; is the default and is not stored. | [optional] 

## Methods

### NewPostApiChartsByChartIdActionsValidateRequestSeatsInnerAccessibility

`func NewPostApiChartsByChartIdActionsValidateRequestSeatsInnerAccessibility() *PostApiChartsByChartIdActionsValidateRequestSeatsInnerAccessibility`

NewPostApiChartsByChartIdActionsValidateRequestSeatsInnerAccessibility instantiates a new PostApiChartsByChartIdActionsValidateRequestSeatsInnerAccessibility object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostApiChartsByChartIdActionsValidateRequestSeatsInnerAccessibilityWithDefaults

`func NewPostApiChartsByChartIdActionsValidateRequestSeatsInnerAccessibilityWithDefaults() *PostApiChartsByChartIdActionsValidateRequestSeatsInnerAccessibility`

NewPostApiChartsByChartIdActionsValidateRequestSeatsInnerAccessibilityWithDefaults instantiates a new PostApiChartsByChartIdActionsValidateRequestSeatsInnerAccessibility object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWheelchair

`func (o *PostApiChartsByChartIdActionsValidateRequestSeatsInnerAccessibility) GetWheelchair() bool`

GetWheelchair returns the Wheelchair field if non-nil, zero value otherwise.

### GetWheelchairOk

`func (o *PostApiChartsByChartIdActionsValidateRequestSeatsInnerAccessibility) GetWheelchairOk() (*bool, bool)`

GetWheelchairOk returns a tuple with the Wheelchair field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWheelchair

`func (o *PostApiChartsByChartIdActionsValidateRequestSeatsInnerAccessibility) SetWheelchair(v bool)`

SetWheelchair sets Wheelchair field to given value.

### HasWheelchair

`func (o *PostApiChartsByChartIdActionsValidateRequestSeatsInnerAccessibility) HasWheelchair() bool`

HasWheelchair returns a boolean if a field has been set.

### GetCompanion

`func (o *PostApiChartsByChartIdActionsValidateRequestSeatsInnerAccessibility) GetCompanion() bool`

GetCompanion returns the Companion field if non-nil, zero value otherwise.

### GetCompanionOk

`func (o *PostApiChartsByChartIdActionsValidateRequestSeatsInnerAccessibility) GetCompanionOk() (*bool, bool)`

GetCompanionOk returns a tuple with the Companion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanion

`func (o *PostApiChartsByChartIdActionsValidateRequestSeatsInnerAccessibility) SetCompanion(v bool)`

SetCompanion sets Companion field to given value.

### HasCompanion

`func (o *PostApiChartsByChartIdActionsValidateRequestSeatsInnerAccessibility) HasCompanion() bool`

HasCompanion returns a boolean if a field has been set.

### GetRestrictedView

`func (o *PostApiChartsByChartIdActionsValidateRequestSeatsInnerAccessibility) GetRestrictedView() bool`

GetRestrictedView returns the RestrictedView field if non-nil, zero value otherwise.

### GetRestrictedViewOk

`func (o *PostApiChartsByChartIdActionsValidateRequestSeatsInnerAccessibility) GetRestrictedViewOk() (*bool, bool)`

GetRestrictedViewOk returns a tuple with the RestrictedView field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestrictedView

`func (o *PostApiChartsByChartIdActionsValidateRequestSeatsInnerAccessibility) SetRestrictedView(v bool)`

SetRestrictedView sets RestrictedView field to given value.

### HasRestrictedView

`func (o *PostApiChartsByChartIdActionsValidateRequestSeatsInnerAccessibility) HasRestrictedView() bool`

HasRestrictedView returns a boolean if a field has been set.

### GetObstructedView

`func (o *PostApiChartsByChartIdActionsValidateRequestSeatsInnerAccessibility) GetObstructedView() bool`

GetObstructedView returns the ObstructedView field if non-nil, zero value otherwise.

### GetObstructedViewOk

`func (o *PostApiChartsByChartIdActionsValidateRequestSeatsInnerAccessibility) GetObstructedViewOk() (*bool, bool)`

GetObstructedViewOk returns a tuple with the ObstructedView field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObstructedView

`func (o *PostApiChartsByChartIdActionsValidateRequestSeatsInnerAccessibility) SetObstructedView(v bool)`

SetObstructedView sets ObstructedView field to given value.

### HasObstructedView

`func (o *PostApiChartsByChartIdActionsValidateRequestSeatsInnerAccessibility) HasObstructedView() bool`

HasObstructedView returns a boolean if a field has been set.

### GetViewNote

`func (o *PostApiChartsByChartIdActionsValidateRequestSeatsInnerAccessibility) GetViewNote() string`

GetViewNote returns the ViewNote field if non-nil, zero value otherwise.

### GetViewNoteOk

`func (o *PostApiChartsByChartIdActionsValidateRequestSeatsInnerAccessibility) GetViewNoteOk() (*string, bool)`

GetViewNoteOk returns a tuple with the ViewNote field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewNote

`func (o *PostApiChartsByChartIdActionsValidateRequestSeatsInnerAccessibility) SetViewNote(v string)`

SetViewNote sets ViewNote field to given value.

### HasViewNote

`func (o *PostApiChartsByChartIdActionsValidateRequestSeatsInnerAccessibility) HasViewNote() bool`

HasViewNote returns a boolean if a field has been set.

### GetAmenities

`func (o *PostApiChartsByChartIdActionsValidateRequestSeatsInnerAccessibility) GetAmenities() int32`

GetAmenities returns the Amenities field if non-nil, zero value otherwise.

### GetAmenitiesOk

`func (o *PostApiChartsByChartIdActionsValidateRequestSeatsInnerAccessibility) GetAmenitiesOk() (*int32, bool)`

GetAmenitiesOk returns a tuple with the Amenities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmenities

`func (o *PostApiChartsByChartIdActionsValidateRequestSeatsInnerAccessibility) SetAmenities(v int32)`

SetAmenities sets Amenities field to given value.

### HasAmenities

`func (o *PostApiChartsByChartIdActionsValidateRequestSeatsInnerAccessibility) HasAmenities() bool`

HasAmenities returns a boolean if a field has been set.

### GetSeatType

`func (o *PostApiChartsByChartIdActionsValidateRequestSeatsInnerAccessibility) GetSeatType() string`

GetSeatType returns the SeatType field if non-nil, zero value otherwise.

### GetSeatTypeOk

`func (o *PostApiChartsByChartIdActionsValidateRequestSeatsInnerAccessibility) GetSeatTypeOk() (*string, bool)`

GetSeatTypeOk returns a tuple with the SeatType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeatType

`func (o *PostApiChartsByChartIdActionsValidateRequestSeatsInnerAccessibility) SetSeatType(v string)`

SetSeatType sets SeatType field to given value.

### HasSeatType

`func (o *PostApiChartsByChartIdActionsValidateRequestSeatsInnerAccessibility) HasSeatType() bool`

HasSeatType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


