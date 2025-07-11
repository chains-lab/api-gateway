# ErrorsErrorsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | **int32** | Status is the HTTP status code applicable to this problem | 
**Title** | **string** | Title is a short, human-readable summary of the problem | 
**Code** | Pointer to **string** | Code is an application-specific error code, expressed as a string | [optional] 
**Detail** | Pointer to **string** | Detail is a human-readable explanation specific to this occurrence of the problem | [optional] 
**Meta** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewErrorsErrorsInner

`func NewErrorsErrorsInner(status int32, title string, ) *ErrorsErrorsInner`

NewErrorsErrorsInner instantiates a new ErrorsErrorsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewErrorsErrorsInnerWithDefaults

`func NewErrorsErrorsInnerWithDefaults() *ErrorsErrorsInner`

NewErrorsErrorsInnerWithDefaults instantiates a new ErrorsErrorsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *ErrorsErrorsInner) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ErrorsErrorsInner) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ErrorsErrorsInner) SetStatus(v int32)`

SetStatus sets Status field to given value.


### GetTitle

`func (o *ErrorsErrorsInner) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *ErrorsErrorsInner) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *ErrorsErrorsInner) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetCode

`func (o *ErrorsErrorsInner) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *ErrorsErrorsInner) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *ErrorsErrorsInner) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *ErrorsErrorsInner) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetDetail

`func (o *ErrorsErrorsInner) GetDetail() string`

GetDetail returns the Detail field if non-nil, zero value otherwise.

### GetDetailOk

`func (o *ErrorsErrorsInner) GetDetailOk() (*string, bool)`

GetDetailOk returns a tuple with the Detail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetail

`func (o *ErrorsErrorsInner) SetDetail(v string)`

SetDetail sets Detail field to given value.

### HasDetail

`func (o *ErrorsErrorsInner) HasDetail() bool`

HasDetail returns a boolean if a field has been set.

### GetMeta

`func (o *ErrorsErrorsInner) GetMeta() map[string]interface{}`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *ErrorsErrorsInner) GetMetaOk() (*map[string]interface{}, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *ErrorsErrorsInner) SetMeta(v map[string]interface{})`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *ErrorsErrorsInner) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


