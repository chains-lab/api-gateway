# ChainsAuthV1OwnGet400ResponseErrorsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | **int32** | Status is the HTTP status code applicable to this problem | 
**Title** | **string** | Title is a short, human-readable summary of the problem | 
**Code** | **string** | Code is an application-specific error code, expressed as a string | 
**Detail** | **string** | Detail is a human-readable explanation specific to this occurrence of the problem | 
**Meta** | [**ChainsAuthV1OwnGet400ResponseErrorsInnerMeta**](ChainsAuthV1OwnGet400ResponseErrorsInnerMeta.md) |  | 

## Methods

### NewChainsAuthV1OwnGet400ResponseErrorsInner

`func NewChainsAuthV1OwnGet400ResponseErrorsInner(status int32, title string, code string, detail string, meta ChainsAuthV1OwnGet400ResponseErrorsInnerMeta, ) *ChainsAuthV1OwnGet400ResponseErrorsInner`

NewChainsAuthV1OwnGet400ResponseErrorsInner instantiates a new ChainsAuthV1OwnGet400ResponseErrorsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChainsAuthV1OwnGet400ResponseErrorsInnerWithDefaults

`func NewChainsAuthV1OwnGet400ResponseErrorsInnerWithDefaults() *ChainsAuthV1OwnGet400ResponseErrorsInner`

NewChainsAuthV1OwnGet400ResponseErrorsInnerWithDefaults instantiates a new ChainsAuthV1OwnGet400ResponseErrorsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *ChainsAuthV1OwnGet400ResponseErrorsInner) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ChainsAuthV1OwnGet400ResponseErrorsInner) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ChainsAuthV1OwnGet400ResponseErrorsInner) SetStatus(v int32)`

SetStatus sets Status field to given value.


### GetTitle

`func (o *ChainsAuthV1OwnGet400ResponseErrorsInner) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *ChainsAuthV1OwnGet400ResponseErrorsInner) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *ChainsAuthV1OwnGet400ResponseErrorsInner) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetCode

`func (o *ChainsAuthV1OwnGet400ResponseErrorsInner) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *ChainsAuthV1OwnGet400ResponseErrorsInner) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *ChainsAuthV1OwnGet400ResponseErrorsInner) SetCode(v string)`

SetCode sets Code field to given value.


### GetDetail

`func (o *ChainsAuthV1OwnGet400ResponseErrorsInner) GetDetail() string`

GetDetail returns the Detail field if non-nil, zero value otherwise.

### GetDetailOk

`func (o *ChainsAuthV1OwnGet400ResponseErrorsInner) GetDetailOk() (*string, bool)`

GetDetailOk returns a tuple with the Detail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetail

`func (o *ChainsAuthV1OwnGet400ResponseErrorsInner) SetDetail(v string)`

SetDetail sets Detail field to given value.


### GetMeta

`func (o *ChainsAuthV1OwnGet400ResponseErrorsInner) GetMeta() ChainsAuthV1OwnGet400ResponseErrorsInnerMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *ChainsAuthV1OwnGet400ResponseErrorsInner) GetMetaOk() (*ChainsAuthV1OwnGet400ResponseErrorsInnerMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *ChainsAuthV1OwnGet400ResponseErrorsInner) SetMeta(v ChainsAuthV1OwnGet400ResponseErrorsInnerMeta)`

SetMeta sets Meta field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


