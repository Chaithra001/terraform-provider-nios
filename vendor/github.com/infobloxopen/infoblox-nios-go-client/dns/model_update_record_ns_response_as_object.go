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

// checks if the UpdateRecordNsResponseAsObject type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateRecordNsResponseAsObject{}

// UpdateRecordNsResponseAsObject The response format to update __RecordNs__ in object format.
type UpdateRecordNsResponseAsObject struct {
	Result               *RecordNs `json:"result,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _UpdateRecordNsResponseAsObject UpdateRecordNsResponseAsObject

// NewUpdateRecordNsResponseAsObject instantiates a new UpdateRecordNsResponseAsObject object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateRecordNsResponseAsObject() *UpdateRecordNsResponseAsObject {
	this := UpdateRecordNsResponseAsObject{}
	return &this
}

// NewUpdateRecordNsResponseAsObjectWithDefaults instantiates a new UpdateRecordNsResponseAsObject object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateRecordNsResponseAsObjectWithDefaults() *UpdateRecordNsResponseAsObject {
	this := UpdateRecordNsResponseAsObject{}
	return &this
}

// GetResult returns the Result field value if set, zero value otherwise.
func (o *UpdateRecordNsResponseAsObject) GetResult() RecordNs {
	if o == nil || IsNil(o.Result) {
		var ret RecordNs
		return ret
	}
	return *o.Result
}

// GetResultOk returns a tuple with the Result field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateRecordNsResponseAsObject) GetResultOk() (*RecordNs, bool) {
	if o == nil || IsNil(o.Result) {
		return nil, false
	}
	return o.Result, true
}

// HasResult returns a boolean if a field has been set.
func (o *UpdateRecordNsResponseAsObject) HasResult() bool {
	if o != nil && !IsNil(o.Result) {
		return true
	}

	return false
}

// SetResult gets a reference to the given RecordNs and assigns it to the Result field.
func (o *UpdateRecordNsResponseAsObject) SetResult(v RecordNs) {
	o.Result = &v
}

func (o UpdateRecordNsResponseAsObject) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateRecordNsResponseAsObject) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Result) {
		toSerialize["result"] = o.Result
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *UpdateRecordNsResponseAsObject) UnmarshalJSON(data []byte) (err error) {
	varUpdateRecordNsResponseAsObject := _UpdateRecordNsResponseAsObject{}

	err = json.Unmarshal(data, &varUpdateRecordNsResponseAsObject)

	if err != nil {
		return err
	}

	*o = UpdateRecordNsResponseAsObject(varUpdateRecordNsResponseAsObject)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "result")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableUpdateRecordNsResponseAsObject struct {
	value *UpdateRecordNsResponseAsObject
	isSet bool
}

func (v NullableUpdateRecordNsResponseAsObject) Get() *UpdateRecordNsResponseAsObject {
	return v.value
}

func (v *NullableUpdateRecordNsResponseAsObject) Set(val *UpdateRecordNsResponseAsObject) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateRecordNsResponseAsObject) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateRecordNsResponseAsObject) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateRecordNsResponseAsObject(val *UpdateRecordNsResponseAsObject) *NullableUpdateRecordNsResponseAsObject {
	return &NullableUpdateRecordNsResponseAsObject{value: val, isSet: true}
}

func (v NullableUpdateRecordNsResponseAsObject) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateRecordNsResponseAsObject) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
