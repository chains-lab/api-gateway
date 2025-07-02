# CabinetData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | User ID of the elector. | 
**Type** | **string** |  | 
**Attributes** | [**CabinetAttributes**](CabinetAttributes.md) |  | 

## Methods

### NewCabinetData

`func NewCabinetData(id string, type_ string, attributes CabinetAttributes, ) *CabinetData`

NewCabinetData instantiates a new CabinetData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCabinetDataWithDefaults

`func NewCabinetDataWithDefaults() *CabinetData`

NewCabinetDataWithDefaults instantiates a new CabinetData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CabinetData) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CabinetData) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CabinetData) SetId(v string)`

SetId sets Id field to given value.


### GetType

`func (o *CabinetData) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CabinetData) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CabinetData) SetType(v string)`

SetType sets Type field to given value.


### GetAttributes

`func (o *CabinetData) GetAttributes() CabinetAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *CabinetData) GetAttributesOk() (*CabinetAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *CabinetData) SetAttributes(v CabinetAttributes)`

SetAttributes sets Attributes field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


