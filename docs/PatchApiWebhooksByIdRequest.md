# PatchApiWebhooksByIdRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Url** | Pointer to **string** |  | [optional] 
**IsActive** | Pointer to **bool** |  | [optional] 
**EventTypes** | Pointer to **[]string** | Subscribed event types: &#x60;[\&quot;*\&quot;]&#x60; (default) for every type, or a list of &#x60;reservations.confirmed&#x60;, &#x60;holds.created&#x60;, &#x60;holds.released&#x60;, &#x60;ga.confirmed&#x60;, &#x60;ga.released&#x60;, &#x60;objects.updated&#x60;, &#x60;season.objects.updated&#x60;, &#x60;hold.expired&#x60;, &#x60;event.created&#x60;, &#x60;event.updated&#x60;, &#x60;event.deleted&#x60;, &#x60;chart.published&#x60;, &#x60;lots.switched&#x60;, &#x60;chart.created&#x60;, &#x60;event.for_sale_changed&#x60;, &#x60;event.channels_updated&#x60;, &#x60;event.categories_updated&#x60;, &#x60;event.table_booking_changed&#x60;, &#x60;event.listings_updated&#x60;. Seats.io-named aliases (&#x60;event.for.sale.config.updated&#x60;, &#x60;event.channels.updated&#x60;, &#x60;event.categories.updated&#x60;, &#x60;event.table.booking.config.updated&#x60;) are opt-in: listed explicitly, the endpoint also receives the aliased event under that name (same payload, &#x60;type&#x60; &#x3D; the alias); &#x60;[\&quot;*\&quot;]&#x60; never includes them. | [optional] 

## Methods

### NewPatchApiWebhooksByIdRequest

`func NewPatchApiWebhooksByIdRequest() *PatchApiWebhooksByIdRequest`

NewPatchApiWebhooksByIdRequest instantiates a new PatchApiWebhooksByIdRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchApiWebhooksByIdRequestWithDefaults

`func NewPatchApiWebhooksByIdRequestWithDefaults() *PatchApiWebhooksByIdRequest`

NewPatchApiWebhooksByIdRequestWithDefaults instantiates a new PatchApiWebhooksByIdRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUrl

`func (o *PatchApiWebhooksByIdRequest) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *PatchApiWebhooksByIdRequest) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *PatchApiWebhooksByIdRequest) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *PatchApiWebhooksByIdRequest) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetIsActive

`func (o *PatchApiWebhooksByIdRequest) GetIsActive() bool`

GetIsActive returns the IsActive field if non-nil, zero value otherwise.

### GetIsActiveOk

`func (o *PatchApiWebhooksByIdRequest) GetIsActiveOk() (*bool, bool)`

GetIsActiveOk returns a tuple with the IsActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsActive

`func (o *PatchApiWebhooksByIdRequest) SetIsActive(v bool)`

SetIsActive sets IsActive field to given value.

### HasIsActive

`func (o *PatchApiWebhooksByIdRequest) HasIsActive() bool`

HasIsActive returns a boolean if a field has been set.

### GetEventTypes

`func (o *PatchApiWebhooksByIdRequest) GetEventTypes() []string`

GetEventTypes returns the EventTypes field if non-nil, zero value otherwise.

### GetEventTypesOk

`func (o *PatchApiWebhooksByIdRequest) GetEventTypesOk() (*[]string, bool)`

GetEventTypesOk returns a tuple with the EventTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventTypes

`func (o *PatchApiWebhooksByIdRequest) SetEventTypes(v []string)`

SetEventTypes sets EventTypes field to given value.

### HasEventTypes

`func (o *PatchApiWebhooksByIdRequest) HasEventTypes() bool`

HasEventTypes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


