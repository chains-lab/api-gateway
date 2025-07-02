# UpdateProfileByAdminData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | User ID of the elector. | 
**Type** | **string** |  | 
**Attributes** | [**UpdateProfileByAdminDataAttributes**](UpdateProfileByAdminDataAttributes.md) |  | 

## Methods

### NewUpdateProfileByAdminData

`func NewUpdateProfileByAdminData(id string, type_ string, attributes UpdateProfileByAdminDataAttributes, ) *UpdateProfileByAdminData`

NewUpdateProfileByAdminData instantiates a new UpdateProfileByAdminData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateProfileByAdminDataWithDefaults

`func NewUpdateProfileByAdminDataWithDefaults() *UpdateProfileByAdminData`

NewUpdateProfileByAdminDataWithDefaults instantiates a new UpdateProfileByAdminData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *UpdateProfileByAdminData) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UpdateProfileByAdminData) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UpdateProfileByAdminData) SetId(v string)`

SetId sets Id field to given value.


### GetType

`func (o *UpdateProfileByAdminData) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *UpdateProfileByAdminData) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *UpdateProfileByAdminData) SetType(v string)`

SetType sets Type field to given value.


### GetAttributes

`func (o *UpdateProfileByAdminData) GetAttributes() UpdateProfileByAdminDataAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *UpdateProfileByAdminData) GetAttributesOk() (*UpdateProfileByAdminDataAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *UpdateProfileByAdminData) SetAttributes(v UpdateProfileByAdminDataAttributes)`

SetAttributes sets Attributes field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


