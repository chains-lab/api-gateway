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

// checks if the RefreshToken type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RefreshToken{}

// RefreshToken struct for RefreshToken
type RefreshToken struct {
	Data RefreshTokenData `json:"data"`
}

type _RefreshToken RefreshToken

// NewRefreshToken instantiates a new RefreshToken object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRefreshToken(data RefreshTokenData) *RefreshToken {
	this := RefreshToken{}
	this.Data = data
	return &this
}

// NewRefreshTokenWithDefaults instantiates a new RefreshToken object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRefreshTokenWithDefaults() *RefreshToken {
	this := RefreshToken{}
	return &this
}

// GetData returns the Data field value
func (o *RefreshToken) GetData() RefreshTokenData {
	if o == nil {
		var ret RefreshTokenData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *RefreshToken) GetDataOk() (*RefreshTokenData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *RefreshToken) SetData(v RefreshTokenData) {
	o.Data = v
}

func (o RefreshToken) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RefreshToken) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

func (o *RefreshToken) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"data",
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

	varRefreshToken := _RefreshToken{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varRefreshToken)

	if err != nil {
		return err
	}

	*o = RefreshToken(varRefreshToken)

	return err
}

type NullableRefreshToken struct {
	value *RefreshToken
	isSet bool
}

func (v NullableRefreshToken) Get() *RefreshToken {
	return v.value
}

func (v *NullableRefreshToken) Set(val *RefreshToken) {
	v.value = val
	v.isSet = true
}

func (v NullableRefreshToken) IsSet() bool {
	return v.isSet
}

func (v *NullableRefreshToken) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRefreshToken(val *RefreshToken) *NullableRefreshToken {
	return &NullableRefreshToken{value: val, isSet: true}
}

func (v NullableRefreshToken) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRefreshToken) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


