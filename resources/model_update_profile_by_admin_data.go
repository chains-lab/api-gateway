/*
Chains API

Chains API

API version: 0.0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package resources

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the UpdateProfileByAdminData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateProfileByAdminData{}

// UpdateProfileByAdminData struct for UpdateProfileByAdminData
type UpdateProfileByAdminData struct {
	// User ID of the elector.
	Id string `json:"id"`
	Type string `json:"type"`
	Attributes UpdateProfileByAdminDataAttributes `json:"attributes"`
}

type _UpdateProfileByAdminData UpdateProfileByAdminData

// NewUpdateProfileByAdminData instantiates a new UpdateProfileByAdminData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateProfileByAdminData(id string, type_ string, attributes UpdateProfileByAdminDataAttributes) *UpdateProfileByAdminData {
	this := UpdateProfileByAdminData{}
	this.Id = id
	this.Type = type_
	this.Attributes = attributes
	return &this
}

// NewUpdateProfileByAdminDataWithDefaults instantiates a new UpdateProfileByAdminData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateProfileByAdminDataWithDefaults() *UpdateProfileByAdminData {
	this := UpdateProfileByAdminData{}
	return &this
}

// GetId returns the Id field value
func (o *UpdateProfileByAdminData) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *UpdateProfileByAdminData) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *UpdateProfileByAdminData) SetId(v string) {
	o.Id = v
}

// GetType returns the Type field value
func (o *UpdateProfileByAdminData) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *UpdateProfileByAdminData) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *UpdateProfileByAdminData) SetType(v string) {
	o.Type = v
}

// GetAttributes returns the Attributes field value
func (o *UpdateProfileByAdminData) GetAttributes() UpdateProfileByAdminDataAttributes {
	if o == nil {
		var ret UpdateProfileByAdminDataAttributes
		return ret
	}

	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *UpdateProfileByAdminData) GetAttributesOk() (*UpdateProfileByAdminDataAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value
func (o *UpdateProfileByAdminData) SetAttributes(v UpdateProfileByAdminDataAttributes) {
	o.Attributes = v
}

func (o UpdateProfileByAdminData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateProfileByAdminData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["type"] = o.Type
	toSerialize["attributes"] = o.Attributes
	return toSerialize, nil
}

func (o *UpdateProfileByAdminData) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"type",
		"attributes",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varUpdateProfileByAdminData := _UpdateProfileByAdminData{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varUpdateProfileByAdminData)

	if err != nil {
		return err
	}

	*o = UpdateProfileByAdminData(varUpdateProfileByAdminData)

	return err
}

type NullableUpdateProfileByAdminData struct {
	value *UpdateProfileByAdminData
	isSet bool
}

func (v NullableUpdateProfileByAdminData) Get() *UpdateProfileByAdminData {
	return v.value
}

func (v *NullableUpdateProfileByAdminData) Set(val *UpdateProfileByAdminData) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateProfileByAdminData) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateProfileByAdminData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateProfileByAdminData(val *UpdateProfileByAdminData) *NullableUpdateProfileByAdminData {
	return &NullableUpdateProfileByAdminData{value: val, isSet: true}
}

func (v NullableUpdateProfileByAdminData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateProfileByAdminData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


