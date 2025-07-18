/*
Infoblox IPAM API

OpenAPI specification for Infoblox NIOS WAPI IPAM objects

API version: 2.13.6
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package ipam

import (
	"encoding/json"
)

// checks if the Ipv6networkMsAdUserData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Ipv6networkMsAdUserData{}

// Ipv6networkMsAdUserData struct for Ipv6networkMsAdUserData
type Ipv6networkMsAdUserData struct {
	// The number of active users.
	ActiveUsersCount     *int64 `json:"active_users_count,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _Ipv6networkMsAdUserData Ipv6networkMsAdUserData

// NewIpv6networkMsAdUserData instantiates a new Ipv6networkMsAdUserData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIpv6networkMsAdUserData() *Ipv6networkMsAdUserData {
	this := Ipv6networkMsAdUserData{}
	return &this
}

// NewIpv6networkMsAdUserDataWithDefaults instantiates a new Ipv6networkMsAdUserData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIpv6networkMsAdUserDataWithDefaults() *Ipv6networkMsAdUserData {
	this := Ipv6networkMsAdUserData{}
	return &this
}

// GetActiveUsersCount returns the ActiveUsersCount field value if set, zero value otherwise.
func (o *Ipv6networkMsAdUserData) GetActiveUsersCount() int64 {
	if o == nil || IsNil(o.ActiveUsersCount) {
		var ret int64
		return ret
	}
	return *o.ActiveUsersCount
}

// GetActiveUsersCountOk returns a tuple with the ActiveUsersCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Ipv6networkMsAdUserData) GetActiveUsersCountOk() (*int64, bool) {
	if o == nil || IsNil(o.ActiveUsersCount) {
		return nil, false
	}
	return o.ActiveUsersCount, true
}

// HasActiveUsersCount returns a boolean if a field has been set.
func (o *Ipv6networkMsAdUserData) HasActiveUsersCount() bool {
	if o != nil && !IsNil(o.ActiveUsersCount) {
		return true
	}

	return false
}

// SetActiveUsersCount gets a reference to the given int64 and assigns it to the ActiveUsersCount field.
func (o *Ipv6networkMsAdUserData) SetActiveUsersCount(v int64) {
	o.ActiveUsersCount = &v
}

func (o Ipv6networkMsAdUserData) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Ipv6networkMsAdUserData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ActiveUsersCount) {
		toSerialize["active_users_count"] = o.ActiveUsersCount
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *Ipv6networkMsAdUserData) UnmarshalJSON(data []byte) (err error) {
	varIpv6networkMsAdUserData := _Ipv6networkMsAdUserData{}

	err = json.Unmarshal(data, &varIpv6networkMsAdUserData)

	if err != nil {
		return err
	}

	*o = Ipv6networkMsAdUserData(varIpv6networkMsAdUserData)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "active_users_count")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableIpv6networkMsAdUserData struct {
	value *Ipv6networkMsAdUserData
	isSet bool
}

func (v NullableIpv6networkMsAdUserData) Get() *Ipv6networkMsAdUserData {
	return v.value
}

func (v *NullableIpv6networkMsAdUserData) Set(val *Ipv6networkMsAdUserData) {
	v.value = val
	v.isSet = true
}

func (v NullableIpv6networkMsAdUserData) IsSet() bool {
	return v.isSet
}

func (v *NullableIpv6networkMsAdUserData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIpv6networkMsAdUserData(val *Ipv6networkMsAdUserData) *NullableIpv6networkMsAdUserData {
	return &NullableIpv6networkMsAdUserData{value: val, isSet: true}
}

func (v NullableIpv6networkMsAdUserData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIpv6networkMsAdUserData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
