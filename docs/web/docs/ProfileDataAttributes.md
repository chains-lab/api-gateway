# ProfileDataAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Username** | **string** | The username of the elector. | 
**Pseudonym** | Pointer to **string** | The pseudonym of the elector, if applicable. | [optional] 
**Description** | Pointer to **string** | A brief description of the elector. | [optional] 
**Avatar** | Pointer to **string** | The URL of the elector&#39;s avatar image. | [optional] 
**Official** | **bool** | Indicates if the elector is an official candidate. | 
**UpdatedAt** | **time.Time** | The date and time when the profile was last updated. | 
**CreatedAt** | **time.Time** | The date and time when the profile was created. | 

## Methods

### NewProfileDataAttributes

`func NewProfileDataAttributes(username string, official bool, updatedAt time.Time, createdAt time.Time, ) *ProfileDataAttributes`

NewProfileDataAttributes instantiates a new ProfileDataAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProfileDataAttributesWithDefaults

`func NewProfileDataAttributesWithDefaults() *ProfileDataAttributes`

NewProfileDataAttributesWithDefaults instantiates a new ProfileDataAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUsername

`func (o *ProfileDataAttributes) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *ProfileDataAttributes) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *ProfileDataAttributes) SetUsername(v string)`

SetUsername sets Username field to given value.


### GetPseudonym

`func (o *ProfileDataAttributes) GetPseudonym() string`

GetPseudonym returns the Pseudonym field if non-nil, zero value otherwise.

### GetPseudonymOk

`func (o *ProfileDataAttributes) GetPseudonymOk() (*string, bool)`

GetPseudonymOk returns a tuple with the Pseudonym field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPseudonym

`func (o *ProfileDataAttributes) SetPseudonym(v string)`

SetPseudonym sets Pseudonym field to given value.

### HasPseudonym

`func (o *ProfileDataAttributes) HasPseudonym() bool`

HasPseudonym returns a boolean if a field has been set.

### GetDescription

`func (o *ProfileDataAttributes) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ProfileDataAttributes) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ProfileDataAttributes) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ProfileDataAttributes) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetAvatar

`func (o *ProfileDataAttributes) GetAvatar() string`

GetAvatar returns the Avatar field if non-nil, zero value otherwise.

### GetAvatarOk

`func (o *ProfileDataAttributes) GetAvatarOk() (*string, bool)`

GetAvatarOk returns a tuple with the Avatar field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvatar

`func (o *ProfileDataAttributes) SetAvatar(v string)`

SetAvatar sets Avatar field to given value.

### HasAvatar

`func (o *ProfileDataAttributes) HasAvatar() bool`

HasAvatar returns a boolean if a field has been set.

### GetOfficial

`func (o *ProfileDataAttributes) GetOfficial() bool`

GetOfficial returns the Official field if non-nil, zero value otherwise.

### GetOfficialOk

`func (o *ProfileDataAttributes) GetOfficialOk() (*bool, bool)`

GetOfficialOk returns a tuple with the Official field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOfficial

`func (o *ProfileDataAttributes) SetOfficial(v bool)`

SetOfficial sets Official field to given value.


### GetUpdatedAt

`func (o *ProfileDataAttributes) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *ProfileDataAttributes) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *ProfileDataAttributes) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetCreatedAt

`func (o *ProfileDataAttributes) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ProfileDataAttributes) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ProfileDataAttributes) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


