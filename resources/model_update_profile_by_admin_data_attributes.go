/*
Chains API

Chains API

API version: 0.0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package resources

import (
	"encoding/json"
)

// checks if the UpdateProfileByAdminDataAttributes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateProfileByAdminDataAttributes{}

// UpdateProfileByAdminDataAttributes struct for UpdateProfileByAdminDataAttributes
type UpdateProfileByAdminDataAttributes struct {
	Pseudonym *bool `json:"pseudonym,omitempty"`
	Description *bool `json:"description,omitempty"`
	Avatar *bool `json:"avatar,omitempty"`
}

// NewUpdateProfileByAdminDataAttributes instantiates a new UpdateProfileByAdminDataAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateProfileByAdminDataAttributes() *UpdateProfileByAdminDataAttributes {
	this := UpdateProfileByAdminDataAttributes{}
	return &this
}

// NewUpdateProfileByAdminDataAttributesWithDefaults instantiates a new UpdateProfileByAdminDataAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateProfileByAdminDataAttributesWithDefaults() *UpdateProfileByAdminDataAttributes {
	this := UpdateProfileByAdminDataAttributes{}
	return &this
}

// GetPseudonym returns the Pseudonym field value if set, zero value otherwise.
func (o *UpdateProfileByAdminDataAttributes) GetPseudonym() bool {
	if o == nil || IsNil(o.Pseudonym) {
		var ret bool
		return ret
	}
	return *o.Pseudonym
}

// GetPseudonymOk returns a tuple with the Pseudonym field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateProfileByAdminDataAttributes) GetPseudonymOk() (*bool, bool) {
	if o == nil || IsNil(o.Pseudonym) {
		return nil, false
	}
	return o.Pseudonym, true
}

// HasPseudonym returns a boolean if a field has been set.
func (o *UpdateProfileByAdminDataAttributes) HasPseudonym() bool {
	if o != nil && !IsNil(o.Pseudonym) {
		return true
	}

	return false
}

// SetPseudonym gets a reference to the given bool and assigns it to the Pseudonym field.
func (o *UpdateProfileByAdminDataAttributes) SetPseudonym(v bool) {
	o.Pseudonym = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *UpdateProfileByAdminDataAttributes) GetDescription() bool {
	if o == nil || IsNil(o.Description) {
		var ret bool
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateProfileByAdminDataAttributes) GetDescriptionOk() (*bool, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *UpdateProfileByAdminDataAttributes) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given bool and assigns it to the Description field.
func (o *UpdateProfileByAdminDataAttributes) SetDescription(v bool) {
	o.Description = &v
}

// GetAvatar returns the Avatar field value if set, zero value otherwise.
func (o *UpdateProfileByAdminDataAttributes) GetAvatar() bool {
	if o == nil || IsNil(o.Avatar) {
		var ret bool
		return ret
	}
	return *o.Avatar
}

// GetAvatarOk returns a tuple with the Avatar field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateProfileByAdminDataAttributes) GetAvatarOk() (*bool, bool) {
	if o == nil || IsNil(o.Avatar) {
		return nil, false
	}
	return o.Avatar, true
}

// HasAvatar returns a boolean if a field has been set.
func (o *UpdateProfileByAdminDataAttributes) HasAvatar() bool {
	if o != nil && !IsNil(o.Avatar) {
		return true
	}

	return false
}

// SetAvatar gets a reference to the given bool and assigns it to the Avatar field.
func (o *UpdateProfileByAdminDataAttributes) SetAvatar(v bool) {
	o.Avatar = &v
}

func (o UpdateProfileByAdminDataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateProfileByAdminDataAttributes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Pseudonym) {
		toSerialize["pseudonym"] = o.Pseudonym
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Avatar) {
		toSerialize["avatar"] = o.Avatar
	}
	return toSerialize, nil
}

type NullableUpdateProfileByAdminDataAttributes struct {
	value *UpdateProfileByAdminDataAttributes
	isSet bool
}

func (v NullableUpdateProfileByAdminDataAttributes) Get() *UpdateProfileByAdminDataAttributes {
	return v.value
}

func (v *NullableUpdateProfileByAdminDataAttributes) Set(val *UpdateProfileByAdminDataAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateProfileByAdminDataAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateProfileByAdminDataAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateProfileByAdminDataAttributes(val *UpdateProfileByAdminDataAttributes) *NullableUpdateProfileByAdminDataAttributes {
	return &NullableUpdateProfileByAdminDataAttributes{value: val, isSet: true}
}

func (v NullableUpdateProfileByAdminDataAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateProfileByAdminDataAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


