# PostApiEventsByEventIdMapPublishRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**WithdrawAffectedListings** | Pointer to **bool** | When the draft removes seats of active listings, withdraw those listings in the publish transaction (refused with 409 &#x60;listing_has_holds&#x60; while any object of them is held). Without it such a publish is refused with 409 &#x60;listing_objects_removed&#x60; (&#x60;listings&#x60;: id, key and removed objects of each affected listing). | [optional] 

## Methods

### NewPostApiEventsByEventIdMapPublishRequest

`func NewPostApiEventsByEventIdMapPublishRequest() *PostApiEventsByEventIdMapPublishRequest`

NewPostApiEventsByEventIdMapPublishRequest instantiates a new PostApiEventsByEventIdMapPublishRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostApiEventsByEventIdMapPublishRequestWithDefaults

`func NewPostApiEventsByEventIdMapPublishRequestWithDefaults() *PostApiEventsByEventIdMapPublishRequest`

NewPostApiEventsByEventIdMapPublishRequestWithDefaults instantiates a new PostApiEventsByEventIdMapPublishRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWithdrawAffectedListings

`func (o *PostApiEventsByEventIdMapPublishRequest) GetWithdrawAffectedListings() bool`

GetWithdrawAffectedListings returns the WithdrawAffectedListings field if non-nil, zero value otherwise.

### GetWithdrawAffectedListingsOk

`func (o *PostApiEventsByEventIdMapPublishRequest) GetWithdrawAffectedListingsOk() (*bool, bool)`

GetWithdrawAffectedListingsOk returns a tuple with the WithdrawAffectedListings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWithdrawAffectedListings

`func (o *PostApiEventsByEventIdMapPublishRequest) SetWithdrawAffectedListings(v bool)`

SetWithdrawAffectedListings sets WithdrawAffectedListings field to given value.

### HasWithdrawAffectedListings

`func (o *PostApiEventsByEventIdMapPublishRequest) HasWithdrawAffectedListings() bool`

HasWithdrawAffectedListings returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


