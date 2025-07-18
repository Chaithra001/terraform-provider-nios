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

// checks if the UpdateIpv4addressResponseAsObject type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateIpv4addressResponseAsObject{}

// UpdateIpv4addressResponseAsObject The response format to update __Ipv4address__ in object format.
type UpdateIpv4addressResponseAsObject struct {
	Result               *Ipv4address `json:"result,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _UpdateIpv4addressResponseAsObject UpdateIpv4addressResponseAsObject

// NewUpdateIpv4addressResponseAsObject instantiates a new UpdateIpv4addressResponseAsObject object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateIpv4addressResponseAsObject() *UpdateIpv4addressResponseAsObject {
	this := UpdateIpv4addressResponseAsObject{}
	return &this
}

// NewUpdateIpv4addressResponseAsObjectWithDefaults instantiates a new UpdateIpv4addressResponseAsObject object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateIpv4addressResponseAsObjectWithDefaults() *UpdateIpv4addressResponseAsObject {
	this := UpdateIpv4addressResponseAsObject{}
	return &this
}

// GetResult returns the Result field value if set, zero value otherwise.
func (o *UpdateIpv4addressResponseAsObject) GetResult() Ipv4address {
	if o == nil || IsNil(o.Result) {
		var ret Ipv4address
		return ret
	}
	return *o.Result
}

// GetResultOk returns a tuple with the Result field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateIpv4addressResponseAsObject) GetResultOk() (*Ipv4address, bool) {
	if o == nil || IsNil(o.Result) {
		return nil, false
	}
	return o.Result, true
}

// HasResult returns a boolean if a field has been set.
func (o *UpdateIpv4addressResponseAsObject) HasResult() bool {
	if o != nil && !IsNil(o.Result) {
		return true
	}

	return false
}

// SetResult gets a reference to the given Ipv4address and assigns it to the Result field.
func (o *UpdateIpv4addressResponseAsObject) SetResult(v Ipv4address) {
	o.Result = &v
}

func (o UpdateIpv4addressResponseAsObject) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateIpv4addressResponseAsObject) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Result) {
		toSerialize["result"] = o.Result
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *UpdateIpv4addressResponseAsObject) UnmarshalJSON(data []byte) (err error) {
	varUpdateIpv4addressResponseAsObject := _UpdateIpv4addressResponseAsObject{}

	err = json.Unmarshal(data, &varUpdateIpv4addressResponseAsObject)

	if err != nil {
		return err
	}

	*o = UpdateIpv4addressResponseAsObject(varUpdateIpv4addressResponseAsObject)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "result")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableUpdateIpv4addressResponseAsObject struct {
	value *UpdateIpv4addressResponseAsObject
	isSet bool
}

func (v NullableUpdateIpv4addressResponseAsObject) Get() *UpdateIpv4addressResponseAsObject {
	return v.value
}

func (v *NullableUpdateIpv4addressResponseAsObject) Set(val *UpdateIpv4addressResponseAsObject) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateIpv4addressResponseAsObject) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateIpv4addressResponseAsObject) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateIpv4addressResponseAsObject(val *UpdateIpv4addressResponseAsObject) *NullableUpdateIpv4addressResponseAsObject {
	return &NullableUpdateIpv4addressResponseAsObject{value: val, isSet: true}
}

func (v NullableUpdateIpv4addressResponseAsObject) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateIpv4addressResponseAsObject) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
