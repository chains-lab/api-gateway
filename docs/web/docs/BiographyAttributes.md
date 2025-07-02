# BiographyAttributes

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

### NewBiographyAttributes

`func NewBiographyAttributes() *BiographyAttributes`

NewBiographyAttributes instantiates a new BiographyAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBiographyAttributesWithDefaults

`func NewBiographyAttributesWithDefaults() *BiographyAttributes`

NewBiographyAttributesWithDefaults instantiates a new BiographyAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBirthday

`func (o *BiographyAttributes) GetBirthday() time.Time`

GetBirthday returns the Birthday field if non-nil, zero value otherwise.

### GetBirthdayOk

`func (o *BiographyAttributes) GetBirthdayOk() (*time.Time, bool)`

GetBirthdayOk returns a tuple with the Birthday field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBirthday

`func (o *BiographyAttributes) SetBirthday(v time.Time)`

SetBirthday sets Birthday field to given value.

### HasBirthday

`func (o *BiographyAttributes) HasBirthday() bool`

HasBirthday returns a boolean if a field has been set.

### GetSex

`func (o *BiographyAttributes) GetSex() string`

GetSex returns the Sex field if non-nil, zero value otherwise.

### GetSexOk

`func (o *BiographyAttributes) GetSexOk() (*string, bool)`

GetSexOk returns a tuple with the Sex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSex

`func (o *BiographyAttributes) SetSex(v string)`

SetSex sets Sex field to given value.

### HasSex

`func (o *BiographyAttributes) HasSex() bool`

HasSex returns a boolean if a field has been set.

### GetCity

`func (o *BiographyAttributes) GetCity() string`

GetCity returns the City field if non-nil, zero value otherwise.

### GetCityOk

`func (o *BiographyAttributes) GetCityOk() (*string, bool)`

GetCityOk returns a tuple with the City field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCity

`func (o *BiographyAttributes) SetCity(v string)`

SetCity sets City field to given value.

### HasCity

`func (o *BiographyAttributes) HasCity() bool`

HasCity returns a boolean if a field has been set.

### GetRegion

`func (o *BiographyAttributes) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *BiographyAttributes) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *BiographyAttributes) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *BiographyAttributes) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### GetCountry

`func (o *BiographyAttributes) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *BiographyAttributes) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *BiographyAttributes) SetCountry(v string)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *BiographyAttributes) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### GetNationality

`func (o *BiographyAttributes) GetNationality() string`

GetNationality returns the Nationality field if non-nil, zero value otherwise.

### GetNationalityOk

`func (o *BiographyAttributes) GetNationalityOk() (*string, bool)`

GetNationalityOk returns a tuple with the Nationality field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNationality

`func (o *BiographyAttributes) SetNationality(v string)`

SetNationality sets Nationality field to given value.

### HasNationality

`func (o *BiographyAttributes) HasNationality() bool`

HasNationality returns a boolean if a field has been set.

### GetPrimaryLanguage

`func (o *BiographyAttributes) GetPrimaryLanguage() string`

GetPrimaryLanguage returns the PrimaryLanguage field if non-nil, zero value otherwise.

### GetPrimaryLanguageOk

`func (o *BiographyAttributes) GetPrimaryLanguageOk() (*string, bool)`

GetPrimaryLanguageOk returns a tuple with the PrimaryLanguage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryLanguage

`func (o *BiographyAttributes) SetPrimaryLanguage(v string)`

SetPrimaryLanguage sets PrimaryLanguage field to given value.

### HasPrimaryLanguage

`func (o *BiographyAttributes) HasPrimaryLanguage() bool`

HasPrimaryLanguage returns a boolean if a field has been set.

### GetSexUpdatedAt

`func (o *BiographyAttributes) GetSexUpdatedAt() time.Time`

GetSexUpdatedAt returns the SexUpdatedAt field if non-nil, zero value otherwise.

### GetSexUpdatedAtOk

`func (o *BiographyAttributes) GetSexUpdatedAtOk() (*time.Time, bool)`

GetSexUpdatedAtOk returns a tuple with the SexUpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSexUpdatedAt

`func (o *BiographyAttributes) SetSexUpdatedAt(v time.Time)`

SetSexUpdatedAt sets SexUpdatedAt field to given value.

### HasSexUpdatedAt

`func (o *BiographyAttributes) HasSexUpdatedAt() bool`

HasSexUpdatedAt returns a boolean if a field has been set.

### GetNationalityUpdatedAt

`func (o *BiographyAttributes) GetNationalityUpdatedAt() time.Time`

GetNationalityUpdatedAt returns the NationalityUpdatedAt field if non-nil, zero value otherwise.

### GetNationalityUpdatedAtOk

`func (o *BiographyAttributes) GetNationalityUpdatedAtOk() (*time.Time, bool)`

GetNationalityUpdatedAtOk returns a tuple with the NationalityUpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNationalityUpdatedAt

`func (o *BiographyAttributes) SetNationalityUpdatedAt(v time.Time)`

SetNationalityUpdatedAt sets NationalityUpdatedAt field to given value.

### HasNationalityUpdatedAt

`func (o *BiographyAttributes) HasNationalityUpdatedAt() bool`

HasNationalityUpdatedAt returns a boolean if a field has been set.

### GetPrimaryLanguageUpdatedAt

`func (o *BiographyAttributes) GetPrimaryLanguageUpdatedAt() time.Time`

GetPrimaryLanguageUpdatedAt returns the PrimaryLanguageUpdatedAt field if non-nil, zero value otherwise.

### GetPrimaryLanguageUpdatedAtOk

`func (o *BiographyAttributes) GetPrimaryLanguageUpdatedAtOk() (*time.Time, bool)`

GetPrimaryLanguageUpdatedAtOk returns a tuple with the PrimaryLanguageUpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryLanguageUpdatedAt

`func (o *BiographyAttributes) SetPrimaryLanguageUpdatedAt(v time.Time)`

SetPrimaryLanguageUpdatedAt sets PrimaryLanguageUpdatedAt field to given value.

### HasPrimaryLanguageUpdatedAt

`func (o *BiographyAttributes) HasPrimaryLanguageUpdatedAt() bool`

HasPrimaryLanguageUpdatedAt returns a boolean if a field has been set.

### GetResidenceUpdatedAt

`func (o *BiographyAttributes) GetResidenceUpdatedAt() time.Time`

GetResidenceUpdatedAt returns the ResidenceUpdatedAt field if non-nil, zero value otherwise.

### GetResidenceUpdatedAtOk

`func (o *BiographyAttributes) GetResidenceUpdatedAtOk() (*time.Time, bool)`

GetResidenceUpdatedAtOk returns a tuple with the ResidenceUpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResidenceUpdatedAt

`func (o *BiographyAttributes) SetResidenceUpdatedAt(v time.Time)`

SetResidenceUpdatedAt sets ResidenceUpdatedAt field to given value.

### HasResidenceUpdatedAt

`func (o *BiographyAttributes) HasResidenceUpdatedAt() bool`

HasResidenceUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


