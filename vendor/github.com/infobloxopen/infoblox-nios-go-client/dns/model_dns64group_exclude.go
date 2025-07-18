/*
Infoblox DNS API

OpenAPI specification for Infoblox NIOS WAPI DNS objects

API version: 2.13.6
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package dns

import (
	"encoding/json"
)

// checks if the Dns64groupExclude type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Dns64groupExclude{}

// Dns64groupExclude struct for Dns64groupExclude
type Dns64groupExclude struct {
	// The address this rule applies to or \"Any\".
	Address *string `json:"address,omitempty"`
	// The permission to use for this address.
	Permission           *string `json:"permission,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _Dns64groupExclude Dns64groupExclude

// NewDns64groupExclude instantiates a new Dns64groupExclude object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDns64groupExclude() *Dns64groupExclude {
	this := Dns64groupExclude{}
	return &this
}

// NewDns64groupExcludeWithDefaults instantiates a new Dns64groupExclude object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDns64groupExcludeWithDefaults() *Dns64groupExclude {
	this := Dns64groupExclude{}
	return &this
}

// GetAddress returns the Address field value if set, zero value otherwise.
func (o *Dns64groupExclude) GetAddress() string {
	if o == nil || IsNil(o.Address) {
		var ret string
		return ret
	}
	return *o.Address
}

// GetAddressOk returns a tuple with the Address field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Dns64groupExclude) GetAddressOk() (*string, bool) {
	if o == nil || IsNil(o.Address) {
		return nil, false
	}
	return o.Address, true
}

// HasAddress returns a boolean if a field has been set.
func (o *Dns64groupExclude) HasAddress() bool {
	if o != nil && !IsNil(o.Address) {
		return true
	}

	return false
}

// SetAddress gets a reference to the given string and assigns it to the Address field.
func (o *Dns64groupExclude) SetAddress(v string) {
	o.Address = &v
}

// GetPermission returns the Permission field value if set, zero value otherwise.
func (o *Dns64groupExclude) GetPermission() string {
	if o == nil || IsNil(o.Permission) {
		var ret string
		return ret
	}
	return *o.Permission
}

// GetPermissionOk returns a tuple with the Permission field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Dns64groupExclude) GetPermissionOk() (*string, bool) {
	if o == nil || IsNil(o.Permission) {
		return nil, false
	}
	return o.Permission, true
}

// HasPermission returns a boolean if a field has been set.
func (o *Dns64groupExclude) HasPermission() bool {
	if o != nil && !IsNil(o.Permission) {
		return true
	}

	return false
}

// SetPermission gets a reference to the given string and assigns it to the Permission field.
func (o *Dns64groupExclude) SetPermission(v string) {
	o.Permission = &v
}

func (o Dns64groupExclude) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Dns64groupExclude) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Address) {
		toSerialize["address"] = o.Address
	}
	if !IsNil(o.Permission) {
		toSerialize["permission"] = o.Permission
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *Dns64groupExclude) UnmarshalJSON(data []byte) (err error) {
	varDns64groupExclude := _Dns64groupExclude{}

	err = json.Unmarshal(data, &varDns64groupExclude)

	if err != nil {
		return err
	}

	*o = Dns64groupExclude(varDns64groupExclude)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "address")
		delete(additionalProperties, "permission")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableDns64groupExclude struct {
	value *Dns64groupExclude
	isSet bool
}

func (v NullableDns64groupExclude) Get() *Dns64groupExclude {
	return v.value
}

func (v *NullableDns64groupExclude) Set(val *Dns64groupExclude) {
	v.value = val
	v.isSet = true
}

func (v NullableDns64groupExclude) IsSet() bool {
	return v.isSet
}

func (v *NullableDns64groupExclude) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDns64groupExclude(val *Dns64groupExclude) *NullableDns64groupExclude {
	return &NullableDns64groupExclude{value: val, isSet: true}
}

func (v NullableDns64groupExclude) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDns64groupExclude) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
