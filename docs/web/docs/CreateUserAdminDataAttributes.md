# CreateUserAdminDataAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Email** | **string** | The email address of the admin to be created. | 
**Role** | **string** | The role of the admin to be created. | 

## Methods

### NewCreateUserAdminDataAttributes

`func NewCreateUserAdminDataAttributes(email string, role string, ) *CreateUserAdminDataAttributes`

NewCreateUserAdminDataAttributes instantiates a new CreateUserAdminDataAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateUserAdminDataAttributesWithDefaults

`func NewCreateUserAdminDataAttributesWithDefaults() *CreateUserAdminDataAttributes`

NewCreateUserAdminDataAttributesWithDefaults instantiates a new CreateUserAdminDataAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmail

`func (o *CreateUserAdminDataAttributes) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *CreateUserAdminDataAttributes) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *CreateUserAdminDataAttributes) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetRole

`func (o *CreateUserAdminDataAttributes) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *CreateUserAdminDataAttributes) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *CreateUserAdminDataAttributes) SetRole(v string)`

SetRole sets Role field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


