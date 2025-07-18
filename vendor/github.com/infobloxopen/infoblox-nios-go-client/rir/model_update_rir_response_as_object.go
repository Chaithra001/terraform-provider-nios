/*
Infoblox RIR API

OpenAPI specification for Infoblox NIOS WAPI RIR objects

API version: 2.13.6
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package rir

import (
	"encoding/json"
)

// checks if the UpdateRirResponseAsObject type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateRirResponseAsObject{}

// UpdateRirResponseAsObject The response format to update __Rir__ in object format.
type UpdateRirResponseAsObject struct {
	Result               *Rir `json:"result,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _UpdateRirResponseAsObject UpdateRirResponseAsObject

// NewUpdateRirResponseAsObject instantiates a new UpdateRirResponseAsObject object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateRirResponseAsObject() *UpdateRirResponseAsObject {
	this := UpdateRirResponseAsObject{}
	return &this
}

// NewUpdateRirResponseAsObjectWithDefaults instantiates a new UpdateRirResponseAsObject object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateRirResponseAsObjectWithDefaults() *UpdateRirResponseAsObject {
	this := UpdateRirResponseAsObject{}
	return &this
}

// GetResult returns the Result field value if set, zero value otherwise.
func (o *UpdateRirResponseAsObject) GetResult() Rir {
	if o == nil || IsNil(o.Result) {
		var ret Rir
		return ret
	}
	return *o.Result
}

// GetResultOk returns a tuple with the Result field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateRirResponseAsObject) GetResultOk() (*Rir, bool) {
	if o == nil || IsNil(o.Result) {
		return nil, false
	}
	return o.Result, true
}

// HasResult returns a boolean if a field has been set.
func (o *UpdateRirResponseAsObject) HasResult() bool {
	if o != nil && !IsNil(o.Result) {
		return true
	}

	return false
}

// SetResult gets a reference to the given Rir and assigns it to the Result field.
func (o *UpdateRirResponseAsObject) SetResult(v Rir) {
	o.Result = &v
}

func (o UpdateRirResponseAsObject) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateRirResponseAsObject) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Result) {
		toSerialize["result"] = o.Result
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *UpdateRirResponseAsObject) UnmarshalJSON(data []byte) (err error) {
	varUpdateRirResponseAsObject := _UpdateRirResponseAsObject{}

	err = json.Unmarshal(data, &varUpdateRirResponseAsObject)

	if err != nil {
		return err
	}

	*o = UpdateRirResponseAsObject(varUpdateRirResponseAsObject)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "result")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableUpdateRirResponseAsObject struct {
	value *UpdateRirResponseAsObject
	isSet bool
}

func (v NullableUpdateRirResponseAsObject) Get() *UpdateRirResponseAsObject {
	return v.value
}

func (v *NullableUpdateRirResponseAsObject) Set(val *UpdateRirResponseAsObject) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateRirResponseAsObject) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateRirResponseAsObject) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateRirResponseAsObject(val *UpdateRirResponseAsObject) *NullableUpdateRirResponseAsObject {
	return &NullableUpdateRirResponseAsObject{value: val, isSet: true}
}

func (v NullableUpdateRirResponseAsObject) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateRirResponseAsObject) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
