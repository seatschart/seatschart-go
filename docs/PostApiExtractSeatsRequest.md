# PostApiExtractSeatsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**File** | ***os.File** |  | 
**Calibration** | Pointer to **string** |  | [optional] 
**Profile** | Pointer to **string** |  | [optional] 

## Methods

### NewPostApiExtractSeatsRequest

`func NewPostApiExtractSeatsRequest(file *os.File, ) *PostApiExtractSeatsRequest`

NewPostApiExtractSeatsRequest instantiates a new PostApiExtractSeatsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostApiExtractSeatsRequestWithDefaults

`func NewPostApiExtractSeatsRequestWithDefaults() *PostApiExtractSeatsRequest`

NewPostApiExtractSeatsRequestWithDefaults instantiates a new PostApiExtractSeatsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFile

`func (o *PostApiExtractSeatsRequest) GetFile() *os.File`

GetFile returns the File field if non-nil, zero value otherwise.

### GetFileOk

`func (o *PostApiExtractSeatsRequest) GetFileOk() (**os.File, bool)`

GetFileOk returns a tuple with the File field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFile

`func (o *PostApiExtractSeatsRequest) SetFile(v *os.File)`

SetFile sets File field to given value.


### GetCalibration

`func (o *PostApiExtractSeatsRequest) GetCalibration() string`

GetCalibration returns the Calibration field if non-nil, zero value otherwise.

### GetCalibrationOk

`func (o *PostApiExtractSeatsRequest) GetCalibrationOk() (*string, bool)`

GetCalibrationOk returns a tuple with the Calibration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCalibration

`func (o *PostApiExtractSeatsRequest) SetCalibration(v string)`

SetCalibration sets Calibration field to given value.

### HasCalibration

`func (o *PostApiExtractSeatsRequest) HasCalibration() bool`

HasCalibration returns a boolean if a field has been set.

### GetProfile

`func (o *PostApiExtractSeatsRequest) GetProfile() string`

GetProfile returns the Profile field if non-nil, zero value otherwise.

### GetProfileOk

`func (o *PostApiExtractSeatsRequest) GetProfileOk() (*string, bool)`

GetProfileOk returns a tuple with the Profile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfile

`func (o *PostApiExtractSeatsRequest) SetProfile(v string)`

SetProfile sets Profile field to given value.

### HasProfile

`func (o *PostApiExtractSeatsRequest) HasProfile() bool`

HasProfile returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


