# UpdateOwnProfileData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | User ID of the elector. | 
**Type** | **string** |  | 
**Attributes** | [**UpdateOwnProfileDataAttributes**](UpdateOwnProfileDataAttributes.md) |  | 

## Methods

### NewUpdateOwnProfileData

`func NewUpdateOwnProfileData(id string, type_ string, attributes UpdateOwnProfileDataAttributes, ) *UpdateOwnProfileData`

NewUpdateOwnProfileData instantiates a new UpdateOwnProfileData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateOwnProfileDataWithDefaults

`func NewUpdateOwnProfileDataWithDefaults() *UpdateOwnProfileData`

NewUpdateOwnProfileDataWithDefaults instantiates a new UpdateOwnProfileData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *UpdateOwnProfileData) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UpdateOwnProfileData) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UpdateOwnProfileData) SetId(v string)`

SetId sets Id field to given value.


### GetType

`func (o *UpdateOwnProfileData) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *UpdateOwnProfileData) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *UpdateOwnProfileData) SetType(v string)`

SetType sets Type field to given value.


### GetAttributes

`func (o *UpdateOwnProfileData) GetAttributes() UpdateOwnProfileDataAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *UpdateOwnProfileData) GetAttributesOk() (*UpdateOwnProfileDataAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *UpdateOwnProfileData) SetAttributes(v UpdateOwnProfileDataAttributes)`

SetAttributes sets Attributes field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


