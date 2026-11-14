# GetApiWorkspacesOverview200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Month** | **string** | &#x60;YYYY-MM&#x60;. | 
**Totals** | [**GetApiWorkspacesOverview200ResponseTotals**](GetApiWorkspacesOverview200ResponseTotals.md) |  | 
**Items** | [**[]GetApiWorkspacesOverview200ResponseItemsInner**](GetApiWorkspacesOverview200ResponseItemsInner.md) |  | 
**NextCursor** | **NullableString** |  | 
**Seats** | Pointer to [**NullableGetApiWorkspacesOverview200ResponseSeats**](GetApiWorkspacesOverview200ResponseSeats.md) |  | [optional] 
**Account** | Pointer to [**NullableGetApiWorkspacesOverview200ResponseAccount**](GetApiWorkspacesOverview200ResponseAccount.md) |  | [optional] 

## Methods

### NewGetApiWorkspacesOverview200Response

`func NewGetApiWorkspacesOverview200Response(month string, totals GetApiWorkspacesOverview200ResponseTotals, items []GetApiWorkspacesOverview200ResponseItemsInner, nextCursor NullableString, ) *GetApiWorkspacesOverview200Response`

NewGetApiWorkspacesOverview200Response instantiates a new GetApiWorkspacesOverview200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetApiWorkspacesOverview200ResponseWithDefaults

`func NewGetApiWorkspacesOverview200ResponseWithDefaults() *GetApiWorkspacesOverview200Response`

NewGetApiWorkspacesOverview200ResponseWithDefaults instantiates a new GetApiWorkspacesOverview200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMonth

`func (o *GetApiWorkspacesOverview200Response) GetMonth() string`

GetMonth returns the Month field if non-nil, zero value otherwise.

### GetMonthOk

`func (o *GetApiWorkspacesOverview200Response) GetMonthOk() (*string, bool)`

GetMonthOk returns a tuple with the Month field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonth

`func (o *GetApiWorkspacesOverview200Response) SetMonth(v string)`

SetMonth sets Month field to given value.


### GetTotals

`func (o *GetApiWorkspacesOverview200Response) GetTotals() GetApiWorkspacesOverview200ResponseTotals`

GetTotals returns the Totals field if non-nil, zero value otherwise.

### GetTotalsOk

`func (o *GetApiWorkspacesOverview200Response) GetTotalsOk() (*GetApiWorkspacesOverview200ResponseTotals, bool)`

GetTotalsOk returns a tuple with the Totals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotals

`func (o *GetApiWorkspacesOverview200Response) SetTotals(v GetApiWorkspacesOverview200ResponseTotals)`

SetTotals sets Totals field to given value.


### GetItems

`func (o *GetApiWorkspacesOverview200Response) GetItems() []GetApiWorkspacesOverview200ResponseItemsInner`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *GetApiWorkspacesOverview200Response) GetItemsOk() (*[]GetApiWorkspacesOverview200ResponseItemsInner, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *GetApiWorkspacesOverview200Response) SetItems(v []GetApiWorkspacesOverview200ResponseItemsInner)`

SetItems sets Items field to given value.


### GetNextCursor

`func (o *GetApiWorkspacesOverview200Response) GetNextCursor() string`

GetNextCursor returns the NextCursor field if non-nil, zero value otherwise.

### GetNextCursorOk

`func (o *GetApiWorkspacesOverview200Response) GetNextCursorOk() (*string, bool)`

GetNextCursorOk returns a tuple with the NextCursor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextCursor

`func (o *GetApiWorkspacesOverview200Response) SetNextCursor(v string)`

SetNextCursor sets NextCursor field to given value.


### SetNextCursorNil

`func (o *GetApiWorkspacesOverview200Response) SetNextCursorNil(b bool)`

 SetNextCursorNil sets the value for NextCursor to be an explicit nil

### UnsetNextCursor
`func (o *GetApiWorkspacesOverview200Response) UnsetNextCursor()`

UnsetNextCursor ensures that no value is present for NextCursor, not even an explicit nil
### GetSeats

`func (o *GetApiWorkspacesOverview200Response) GetSeats() GetApiWorkspacesOverview200ResponseSeats`

GetSeats returns the Seats field if non-nil, zero value otherwise.

### GetSeatsOk

`func (o *GetApiWorkspacesOverview200Response) GetSeatsOk() (*GetApiWorkspacesOverview200ResponseSeats, bool)`

GetSeatsOk returns a tuple with the Seats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeats

`func (o *GetApiWorkspacesOverview200Response) SetSeats(v GetApiWorkspacesOverview200ResponseSeats)`

SetSeats sets Seats field to given value.

### HasSeats

`func (o *GetApiWorkspacesOverview200Response) HasSeats() bool`

HasSeats returns a boolean if a field has been set.

### SetSeatsNil

`func (o *GetApiWorkspacesOverview200Response) SetSeatsNil(b bool)`

 SetSeatsNil sets the value for Seats to be an explicit nil

### UnsetSeats
`func (o *GetApiWorkspacesOverview200Response) UnsetSeats()`

UnsetSeats ensures that no value is present for Seats, not even an explicit nil
### GetAccount

`func (o *GetApiWorkspacesOverview200Response) GetAccount() GetApiWorkspacesOverview200ResponseAccount`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *GetApiWorkspacesOverview200Response) GetAccountOk() (*GetApiWorkspacesOverview200ResponseAccount, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *GetApiWorkspacesOverview200Response) SetAccount(v GetApiWorkspacesOverview200ResponseAccount)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *GetApiWorkspacesOverview200Response) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### SetAccountNil

`func (o *GetApiWorkspacesOverview200Response) SetAccountNil(b bool)`

 SetAccountNil sets the value for Account to be an explicit nil

### UnsetAccount
`func (o *GetApiWorkspacesOverview200Response) UnsetAccount()`

UnsetAccount ensures that no value is present for Account, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


