# CreateUserByAdminDataAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Email** | **string** | The email address of the admin to be created. | 
**Role** | **string** | The role of the admin to be created. | 

## Methods

### NewCreateUserByAdminDataAttributes

`func NewCreateUserByAdminDataAttributes(email string, role string, ) *CreateUserByAdminDataAttributes`

NewCreateUserByAdminDataAttributes instantiates a new CreateUserByAdminDataAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateUserByAdminDataAttributesWithDefaults

`func NewCreateUserByAdminDataAttributesWithDefaults() *CreateUserByAdminDataAttributes`

NewCreateUserByAdminDataAttributesWithDefaults instantiates a new CreateUserByAdminDataAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmail

`func (o *CreateUserByAdminDataAttributes) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *CreateUserByAdminDataAttributes) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *CreateUserByAdminDataAttributes) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetRole

`func (o *CreateUserByAdminDataAttributes) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *CreateUserByAdminDataAttributes) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *CreateUserByAdminDataAttributes) SetRole(v string)`

SetRole sets Role field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


