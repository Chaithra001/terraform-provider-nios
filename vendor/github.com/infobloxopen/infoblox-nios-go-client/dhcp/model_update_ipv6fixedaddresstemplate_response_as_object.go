/*
Infoblox DHCP API

OpenAPI specification for Infoblox NIOS WAPI DHCP objects

API version: 2.13.6
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package dhcp

import (
	"encoding/json"
)

// checks if the UpdateIpv6fixedaddresstemplateResponseAsObject type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateIpv6fixedaddresstemplateResponseAsObject{}

// UpdateIpv6fixedaddresstemplateResponseAsObject The response format to update __Ipv6fixedaddresstemplate__ in object format.
type UpdateIpv6fixedaddresstemplateResponseAsObject struct {
	Result               *Ipv6fixedaddresstemplate `json:"result,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _UpdateIpv6fixedaddresstemplateResponseAsObject UpdateIpv6fixedaddresstemplateResponseAsObject

// NewUpdateIpv6fixedaddresstemplateResponseAsObject instantiates a new UpdateIpv6fixedaddresstemplateResponseAsObject object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateIpv6fixedaddresstemplateResponseAsObject() *UpdateIpv6fixedaddresstemplateResponseAsObject {
	this := UpdateIpv6fixedaddresstemplateResponseAsObject{}
	return &this
}

// NewUpdateIpv6fixedaddresstemplateResponseAsObjectWithDefaults instantiates a new UpdateIpv6fixedaddresstemplateResponseAsObject object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateIpv6fixedaddresstemplateResponseAsObjectWithDefaults() *UpdateIpv6fixedaddresstemplateResponseAsObject {
	this := UpdateIpv6fixedaddresstemplateResponseAsObject{}
	return &this
}

// GetResult returns the Result field value if set, zero value otherwise.
func (o *UpdateIpv6fixedaddresstemplateResponseAsObject) GetResult() Ipv6fixedaddresstemplate {
	if o == nil || IsNil(o.Result) {
		var ret Ipv6fixedaddresstemplate
		return ret
	}
	return *o.Result
}

// GetResultOk returns a tuple with the Result field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateIpv6fixedaddresstemplateResponseAsObject) GetResultOk() (*Ipv6fixedaddresstemplate, bool) {
	if o == nil || IsNil(o.Result) {
		return nil, false
	}
	return o.Result, true
}

// HasResult returns a boolean if a field has been set.
func (o *UpdateIpv6fixedaddresstemplateResponseAsObject) HasResult() bool {
	if o != nil && !IsNil(o.Result) {
		return true
	}

	return false
}

// SetResult gets a reference to the given Ipv6fixedaddresstemplate and assigns it to the Result field.
func (o *UpdateIpv6fixedaddresstemplateResponseAsObject) SetResult(v Ipv6fixedaddresstemplate) {
	o.Result = &v
}

func (o UpdateIpv6fixedaddresstemplateResponseAsObject) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateIpv6fixedaddresstemplateResponseAsObject) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Result) {
		toSerialize["result"] = o.Result
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *UpdateIpv6fixedaddresstemplateResponseAsObject) UnmarshalJSON(data []byte) (err error) {
	varUpdateIpv6fixedaddresstemplateResponseAsObject := _UpdateIpv6fixedaddresstemplateResponseAsObject{}

	err = json.Unmarshal(data, &varUpdateIpv6fixedaddresstemplateResponseAsObject)

	if err != nil {
		return err
	}

	*o = UpdateIpv6fixedaddresstemplateResponseAsObject(varUpdateIpv6fixedaddresstemplateResponseAsObject)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "result")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableUpdateIpv6fixedaddresstemplateResponseAsObject struct {
	value *UpdateIpv6fixedaddresstemplateResponseAsObject
	isSet bool
}

func (v NullableUpdateIpv6fixedaddresstemplateResponseAsObject) Get() *UpdateIpv6fixedaddresstemplateResponseAsObject {
	return v.value
}

func (v *NullableUpdateIpv6fixedaddresstemplateResponseAsObject) Set(val *UpdateIpv6fixedaddresstemplateResponseAsObject) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateIpv6fixedaddresstemplateResponseAsObject) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateIpv6fixedaddresstemplateResponseAsObject) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateIpv6fixedaddresstemplateResponseAsObject(val *UpdateIpv6fixedaddresstemplateResponseAsObject) *NullableUpdateIpv6fixedaddresstemplateResponseAsObject {
	return &NullableUpdateIpv6fixedaddresstemplateResponseAsObject{value: val, isSet: true}
}

func (v NullableUpdateIpv6fixedaddresstemplateResponseAsObject) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateIpv6fixedaddresstemplateResponseAsObject) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
