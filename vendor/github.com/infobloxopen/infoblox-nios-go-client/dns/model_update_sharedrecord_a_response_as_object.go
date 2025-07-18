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

// checks if the UpdateSharedrecordAResponseAsObject type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateSharedrecordAResponseAsObject{}

// UpdateSharedrecordAResponseAsObject The response format to update __SharedrecordA__ in object format.
type UpdateSharedrecordAResponseAsObject struct {
	Result               *SharedrecordA `json:"result,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _UpdateSharedrecordAResponseAsObject UpdateSharedrecordAResponseAsObject

// NewUpdateSharedrecordAResponseAsObject instantiates a new UpdateSharedrecordAResponseAsObject object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateSharedrecordAResponseAsObject() *UpdateSharedrecordAResponseAsObject {
	this := UpdateSharedrecordAResponseAsObject{}
	return &this
}

// NewUpdateSharedrecordAResponseAsObjectWithDefaults instantiates a new UpdateSharedrecordAResponseAsObject object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateSharedrecordAResponseAsObjectWithDefaults() *UpdateSharedrecordAResponseAsObject {
	this := UpdateSharedrecordAResponseAsObject{}
	return &this
}

// GetResult returns the Result field value if set, zero value otherwise.
func (o *UpdateSharedrecordAResponseAsObject) GetResult() SharedrecordA {
	if o == nil || IsNil(o.Result) {
		var ret SharedrecordA
		return ret
	}
	return *o.Result
}

// GetResultOk returns a tuple with the Result field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateSharedrecordAResponseAsObject) GetResultOk() (*SharedrecordA, bool) {
	if o == nil || IsNil(o.Result) {
		return nil, false
	}
	return o.Result, true
}

// HasResult returns a boolean if a field has been set.
func (o *UpdateSharedrecordAResponseAsObject) HasResult() bool {
	if o != nil && !IsNil(o.Result) {
		return true
	}

	return false
}

// SetResult gets a reference to the given SharedrecordA and assigns it to the Result field.
func (o *UpdateSharedrecordAResponseAsObject) SetResult(v SharedrecordA) {
	o.Result = &v
}

func (o UpdateSharedrecordAResponseAsObject) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateSharedrecordAResponseAsObject) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Result) {
		toSerialize["result"] = o.Result
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *UpdateSharedrecordAResponseAsObject) UnmarshalJSON(data []byte) (err error) {
	varUpdateSharedrecordAResponseAsObject := _UpdateSharedrecordAResponseAsObject{}

	err = json.Unmarshal(data, &varUpdateSharedrecordAResponseAsObject)

	if err != nil {
		return err
	}

	*o = UpdateSharedrecordAResponseAsObject(varUpdateSharedrecordAResponseAsObject)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "result")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableUpdateSharedrecordAResponseAsObject struct {
	value *UpdateSharedrecordAResponseAsObject
	isSet bool
}

func (v NullableUpdateSharedrecordAResponseAsObject) Get() *UpdateSharedrecordAResponseAsObject {
	return v.value
}

func (v *NullableUpdateSharedrecordAResponseAsObject) Set(val *UpdateSharedrecordAResponseAsObject) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateSharedrecordAResponseAsObject) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateSharedrecordAResponseAsObject) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateSharedrecordAResponseAsObject(val *UpdateSharedrecordAResponseAsObject) *NullableUpdateSharedrecordAResponseAsObject {
	return &NullableUpdateSharedrecordAResponseAsObject{value: val, isSet: true}
}

func (v NullableUpdateSharedrecordAResponseAsObject) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateSharedrecordAResponseAsObject) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
