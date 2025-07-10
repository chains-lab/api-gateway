# BiographyAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Birthday** | Pointer to **time.Time** | The birth date of the elector. | [optional] 
**Sex** | Pointer to **string** | sex of the elector | [optional] 
**City** | Pointer to **string** | The city where the elector resides. | [optional] 
**Region** | Pointer to **string** | The region where the elector resides. | [optional] 
**Country** | Pointer to **string** | The country where the elector resides. | [optional] 
**SexUpdatedAt** | Pointer to **time.Time** |  | [optional] 
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


