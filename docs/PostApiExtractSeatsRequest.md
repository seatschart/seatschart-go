# PostApiExtractSeatsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**File** | ***os.File** |  | 
**Calibration** | Pointer to **string** | Calibration JSON (sample patches, rotations, grouping) for the legacy matcher. | [optional] 
**Profile** | Pointer to **string** |  | [optional] 
**Mode** | Pointer to **string** | &#x60;auto&#x60; (default): discover the seat glyphs and match them, no clicks needed. &#x60;examples&#x60;: use the seats clicked in &#x60;examples&#x60;. &#x60;edges&#x60;: legacy contour mode tuned by &#x60;seat_diameter&#x60; and &#x60;closing_kernel&#x60;. | [optional] 
**Examples** | Pointer to [**PostApiExtractSeatsRequestExamples**](PostApiExtractSeatsRequestExamples.md) |  | [optional] 
**Negatives** | Pointer to [**PostApiExtractSeatsRequestNegatives**](PostApiExtractSeatsRequestNegatives.md) |  | [optional] 
**Region** | Pointer to [**PostApiExtractSeatsRequestRegion**](PostApiExtractSeatsRequestRegion.md) |  | [optional] 
**Vector** | Pointer to ***os.File** | Optional vector source of &#x60;file&#x60; (a PDF or an SVG, at most 8 MiB): &#x60;file&#x60; must be that PDF page or SVG rendered whole (any scale; the page box, or the SVG viewBox, maps onto the full image). The scanner reads the drawn geometry (closed paths, rectangles, circles, repeated symbols, text runs) and uses it to check, add and drop seats; a scanned page or one without seat-sized shapes falls back to the image alone (&#x60;vector: \&quot;unusable\&quot;&#x60;). Works with &#x60;auto&#x60; and &#x60;examples&#x60;. | [optional] 
**VectorPage** | Pointer to **float32** | PDF page (1-based) that &#x60;file&#x60; shows; default 1. | [optional] 

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

### GetMode

`func (o *PostApiExtractSeatsRequest) GetMode() string`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *PostApiExtractSeatsRequest) GetModeOk() (*string, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *PostApiExtractSeatsRequest) SetMode(v string)`

SetMode sets Mode field to given value.

### HasMode

`func (o *PostApiExtractSeatsRequest) HasMode() bool`

HasMode returns a boolean if a field has been set.

### GetExamples

`func (o *PostApiExtractSeatsRequest) GetExamples() PostApiExtractSeatsRequestExamples`

GetExamples returns the Examples field if non-nil, zero value otherwise.

### GetExamplesOk

`func (o *PostApiExtractSeatsRequest) GetExamplesOk() (*PostApiExtractSeatsRequestExamples, bool)`

GetExamplesOk returns a tuple with the Examples field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExamples

`func (o *PostApiExtractSeatsRequest) SetExamples(v PostApiExtractSeatsRequestExamples)`

SetExamples sets Examples field to given value.

### HasExamples

`func (o *PostApiExtractSeatsRequest) HasExamples() bool`

HasExamples returns a boolean if a field has been set.

### GetNegatives

`func (o *PostApiExtractSeatsRequest) GetNegatives() PostApiExtractSeatsRequestNegatives`

GetNegatives returns the Negatives field if non-nil, zero value otherwise.

### GetNegativesOk

`func (o *PostApiExtractSeatsRequest) GetNegativesOk() (*PostApiExtractSeatsRequestNegatives, bool)`

GetNegativesOk returns a tuple with the Negatives field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNegatives

`func (o *PostApiExtractSeatsRequest) SetNegatives(v PostApiExtractSeatsRequestNegatives)`

SetNegatives sets Negatives field to given value.

### HasNegatives

`func (o *PostApiExtractSeatsRequest) HasNegatives() bool`

HasNegatives returns a boolean if a field has been set.

### GetRegion

`func (o *PostApiExtractSeatsRequest) GetRegion() PostApiExtractSeatsRequestRegion`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *PostApiExtractSeatsRequest) GetRegionOk() (*PostApiExtractSeatsRequestRegion, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *PostApiExtractSeatsRequest) SetRegion(v PostApiExtractSeatsRequestRegion)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *PostApiExtractSeatsRequest) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### GetVector

`func (o *PostApiExtractSeatsRequest) GetVector() *os.File`

GetVector returns the Vector field if non-nil, zero value otherwise.

### GetVectorOk

`func (o *PostApiExtractSeatsRequest) GetVectorOk() (**os.File, bool)`

GetVectorOk returns a tuple with the Vector field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVector

`func (o *PostApiExtractSeatsRequest) SetVector(v *os.File)`

SetVector sets Vector field to given value.

### HasVector

`func (o *PostApiExtractSeatsRequest) HasVector() bool`

HasVector returns a boolean if a field has been set.

### GetVectorPage

`func (o *PostApiExtractSeatsRequest) GetVectorPage() float32`

GetVectorPage returns the VectorPage field if non-nil, zero value otherwise.

### GetVectorPageOk

`func (o *PostApiExtractSeatsRequest) GetVectorPageOk() (*float32, bool)`

GetVectorPageOk returns a tuple with the VectorPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVectorPage

`func (o *PostApiExtractSeatsRequest) SetVectorPage(v float32)`

SetVectorPage sets VectorPage field to given value.

### HasVectorPage

`func (o *PostApiExtractSeatsRequest) HasVectorPage() bool`

HasVectorPage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


