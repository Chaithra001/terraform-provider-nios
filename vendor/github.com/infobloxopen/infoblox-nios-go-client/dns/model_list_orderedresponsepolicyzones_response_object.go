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

// checks if the ListOrderedresponsepolicyzonesResponseObject type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListOrderedresponsepolicyzonesResponseObject{}

// ListOrderedresponsepolicyzonesResponseObject The response format to retrieve __Orderedresponsepolicyzones__ objects.
type ListOrderedresponsepolicyzonesResponseObject struct {
	Result               []Orderedresponsepolicyzones `json:"result,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ListOrderedresponsepolicyzonesResponseObject ListOrderedresponsepolicyzonesResponseObject

// NewListOrderedresponsepolicyzonesResponseObject instantiates a new ListOrderedresponsepolicyzonesResponseObject object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListOrderedresponsepolicyzonesResponseObject() *ListOrderedresponsepolicyzonesResponseObject {
	this := ListOrderedresponsepolicyzonesResponseObject{}
	return &this
}

// NewListOrderedresponsepolicyzonesResponseObjectWithDefaults instantiates a new ListOrderedresponsepolicyzonesResponseObject object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListOrderedresponsepolicyzonesResponseObjectWithDefaults() *ListOrderedresponsepolicyzonesResponseObject {
	this := ListOrderedresponsepolicyzonesResponseObject{}
	return &this
}

// GetResult returns the Result field value if set, zero value otherwise.
func (o *ListOrderedresponsepolicyzonesResponseObject) GetResult() []Orderedresponsepolicyzones {
	if o == nil || IsNil(o.Result) {
		var ret []Orderedresponsepolicyzones
		return ret
	}
	return o.Result
}

// GetResultOk returns a tuple with the Result field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListOrderedresponsepolicyzonesResponseObject) GetResultOk() ([]Orderedresponsepolicyzones, bool) {
	if o == nil || IsNil(o.Result) {
		return nil, false
	}
	return o.Result, true
}

// HasResult returns a boolean if a field has been set.
func (o *ListOrderedresponsepolicyzonesResponseObject) HasResult() bool {
	if o != nil && !IsNil(o.Result) {
		return true
	}

	return false
}

// SetResult gets a reference to the given []Orderedresponsepolicyzones and assigns it to the Result field.
func (o *ListOrderedresponsepolicyzonesResponseObject) SetResult(v []Orderedresponsepolicyzones) {
	o.Result = v
}

func (o ListOrderedresponsepolicyzonesResponseObject) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListOrderedresponsepolicyzonesResponseObject) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Result) {
		toSerialize["result"] = o.Result
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ListOrderedresponsepolicyzonesResponseObject) UnmarshalJSON(data []byte) (err error) {
	varListOrderedresponsepolicyzonesResponseObject := _ListOrderedresponsepolicyzonesResponseObject{}

	err = json.Unmarshal(data, &varListOrderedresponsepolicyzonesResponseObject)

	if err != nil {
		return err
	}

	*o = ListOrderedresponsepolicyzonesResponseObject(varListOrderedresponsepolicyzonesResponseObject)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "result")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableListOrderedresponsepolicyzonesResponseObject struct {
	value *ListOrderedresponsepolicyzonesResponseObject
	isSet bool
}

func (v NullableListOrderedresponsepolicyzonesResponseObject) Get() *ListOrderedresponsepolicyzonesResponseObject {
	return v.value
}

func (v *NullableListOrderedresponsepolicyzonesResponseObject) Set(val *ListOrderedresponsepolicyzonesResponseObject) {
	v.value = val
	v.isSet = true
}

func (v NullableListOrderedresponsepolicyzonesResponseObject) IsSet() bool {
	return v.isSet
}

func (v *NullableListOrderedresponsepolicyzonesResponseObject) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListOrderedresponsepolicyzonesResponseObject(val *ListOrderedresponsepolicyzonesResponseObject) *NullableListOrderedresponsepolicyzonesResponseObject {
	return &NullableListOrderedresponsepolicyzonesResponseObject{value: val, isSet: true}
}

func (v NullableListOrderedresponsepolicyzonesResponseObject) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListOrderedresponsepolicyzonesResponseObject) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
