# BiographyDataAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Birthday** | Pointer to **time.Time** | The birth date of the elector. | [optional] 
**Sex** | Pointer to **string** | sex of the elector | [optional] 
**City** | Pointer to **string** | The city where the elector resides. | [optional] 
**Region** | Pointer to **string** | The region where the elector resides. | [optional] 
**Country** | Pointer to **string** | The country where the elector resides. | [optional] 
**Nationality** | Pointer to **string** | nationality of the elector | [optional] 
**PrimaryLanguage** | Pointer to **string** | The primary language spoken by the elector. | [optional] 
**SexUpdatedAt** | Pointer to **time.Time** |  | [optional] 
**NationalityUpdatedAt** | Pointer to **time.Time** |  | [optional] 
**PrimaryLanguageUpdatedAt** | Pointer to **time.Time** |  | [optional] 
**ResidenceUpdatedAt** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewBiographyDataAttributes

`func NewBiographyDataAttributes() *BiographyDataAttributes`

NewBiographyDataAttributes instantiates a new BiographyDataAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBiographyDataAttributesWithDefaults

`func NewBiographyDataAttributesWithDefaults() *BiographyDataAttributes`

NewBiographyDataAttributesWithDefaults instantiates a new BiographyDataAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBirthday

`func (o *BiographyDataAttributes) GetBirthday() time.Time`

GetBirthday returns the Birthday field if non-nil, zero value otherwise.

### GetBirthdayOk

`func (o *BiographyDataAttributes) GetBirthdayOk() (*time.Time, bool)`

GetBirthdayOk returns a tuple with the Birthday field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBirthday

`func (o *BiographyDataAttributes) SetBirthday(v time.Time)`

SetBirthday sets Birthday field to given value.

### HasBirthday

`func (o *BiographyDataAttributes) HasBirthday() bool`

HasBirthday returns a boolean if a field has been set.

### GetSex

`func (o *BiographyDataAttributes) GetSex() string`

GetSex returns the Sex field if non-nil, zero value otherwise.

### GetSexOk

`func (o *BiographyDataAttributes) GetSexOk() (*string, bool)`

GetSexOk returns a tuple with the Sex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSex

`func (o *BiographyDataAttributes) SetSex(v string)`

SetSex sets Sex field to given value.

### HasSex

`func (o *BiographyDataAttributes) HasSex() bool`

HasSex returns a boolean if a field has been set.

### GetCity

`func (o *BiographyDataAttributes) GetCity() string`

GetCity returns the City field if non-nil, zero value otherwise.

### GetCityOk

`func (o *BiographyDataAttributes) GetCityOk() (*string, bool)`

GetCityOk returns a tuple with the City field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCity

`func (o *BiographyDataAttributes) SetCity(v string)`

SetCity sets City field to given value.

### HasCity

`func (o *BiographyDataAttributes) HasCity() bool`

HasCity returns a boolean if a field has been set.

### GetRegion

`func (o *BiographyDataAttributes) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *BiographyDataAttributes) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *BiographyDataAttributes) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *BiographyDataAttributes) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### GetCountry

`func (o *BiographyDataAttributes) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *BiographyDataAttributes) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *BiographyDataAttributes) SetCountry(v string)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *BiographyDataAttributes) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### GetNationality

`func (o *BiographyDataAttributes) GetNationality() string`

GetNationality returns the Nationality field if non-nil, zero value otherwise.

### GetNationalityOk

`func (o *BiographyDataAttributes) GetNationalityOk() (*string, bool)`

GetNationalityOk returns a tuple with the Nationality field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNationality

`func (o *BiographyDataAttributes) SetNationality(v string)`

SetNationality sets Nationality field to given value.

### HasNationality

`func (o *BiographyDataAttributes) HasNationality() bool`

HasNationality returns a boolean if a field has been set.

### GetPrimaryLanguage

`func (o *BiographyDataAttributes) GetPrimaryLanguage() string`

GetPrimaryLanguage returns the PrimaryLanguage field if non-nil, zero value otherwise.

### GetPrimaryLanguageOk

`func (o *BiographyDataAttributes) GetPrimaryLanguageOk() (*string, bool)`

GetPrimaryLanguageOk returns a tuple with the PrimaryLanguage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryLanguage

`func (o *BiographyDataAttributes) SetPrimaryLanguage(v string)`

SetPrimaryLanguage sets PrimaryLanguage field to given value.

### HasPrimaryLanguage

`func (o *BiographyDataAttributes) HasPrimaryLanguage() bool`

HasPrimaryLanguage returns a boolean if a field has been set.

### GetSexUpdatedAt

`func (o *BiographyDataAttributes) GetSexUpdatedAt() time.Time`

GetSexUpdatedAt returns the SexUpdatedAt field if non-nil, zero value otherwise.

### GetSexUpdatedAtOk

`func (o *BiographyDataAttributes) GetSexUpdatedAtOk() (*time.Time, bool)`

GetSexUpdatedAtOk returns a tuple with the SexUpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSexUpdatedAt

`func (o *BiographyDataAttributes) SetSexUpdatedAt(v time.Time)`

SetSexUpdatedAt sets SexUpdatedAt field to given value.

### HasSexUpdatedAt

`func (o *BiographyDataAttributes) HasSexUpdatedAt() bool`

HasSexUpdatedAt returns a boolean if a field has been set.

### GetNationalityUpdatedAt

`func (o *BiographyDataAttributes) GetNationalityUpdatedAt() time.Time`

GetNationalityUpdatedAt returns the NationalityUpdatedAt field if non-nil, zero value otherwise.

### GetNationalityUpdatedAtOk

`func (o *BiographyDataAttributes) GetNationalityUpdatedAtOk() (*time.Time, bool)`

GetNationalityUpdatedAtOk returns a tuple with the NationalityUpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNationalityUpdatedAt

`func (o *BiographyDataAttributes) SetNationalityUpdatedAt(v time.Time)`

SetNationalityUpdatedAt sets NationalityUpdatedAt field to given value.

### HasNationalityUpdatedAt

`func (o *BiographyDataAttributes) HasNationalityUpdatedAt() bool`

HasNationalityUpdatedAt returns a boolean if a field has been set.

### GetPrimaryLanguageUpdatedAt

`func (o *BiographyDataAttributes) GetPrimaryLanguageUpdatedAt() time.Time`

GetPrimaryLanguageUpdatedAt returns the PrimaryLanguageUpdatedAt field if non-nil, zero value otherwise.

### GetPrimaryLanguageUpdatedAtOk

`func (o *BiographyDataAttributes) GetPrimaryLanguageUpdatedAtOk() (*time.Time, bool)`

GetPrimaryLanguageUpdatedAtOk returns a tuple with the PrimaryLanguageUpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryLanguageUpdatedAt

`func (o *BiographyDataAttributes) SetPrimaryLanguageUpdatedAt(v time.Time)`

SetPrimaryLanguageUpdatedAt sets PrimaryLanguageUpdatedAt field to given value.

### HasPrimaryLanguageUpdatedAt

`func (o *BiographyDataAttributes) HasPrimaryLanguageUpdatedAt() bool`

HasPrimaryLanguageUpdatedAt returns a boolean if a field has been set.

### GetResidenceUpdatedAt

`func (o *BiographyDataAttributes) GetResidenceUpdatedAt() time.Time`

GetResidenceUpdatedAt returns the ResidenceUpdatedAt field if non-nil, zero value otherwise.

### GetResidenceUpdatedAtOk

`func (o *BiographyDataAttributes) GetResidenceUpdatedAtOk() (*time.Time, bool)`

GetResidenceUpdatedAtOk returns a tuple with the ResidenceUpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResidenceUpdatedAt

`func (o *BiographyDataAttributes) SetResidenceUpdatedAt(v time.Time)`

SetResidenceUpdatedAt sets ResidenceUpdatedAt field to given value.

### HasResidenceUpdatedAt

`func (o *BiographyDataAttributes) HasResidenceUpdatedAt() bool`

HasResidenceUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


