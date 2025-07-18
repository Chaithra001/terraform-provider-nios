/*
Infoblox CLOUD API

OpenAPI specification for Infoblox NIOS WAPI CLOUD objects

API version: 2.13.6
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package cloud

import (
	"encoding/json"
)

// checks if the UpdateAzurednstaskgroupResponseAsObject type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateAzurednstaskgroupResponseAsObject{}

// UpdateAzurednstaskgroupResponseAsObject The response format to update __Azurednstaskgroup__ in object format.
type UpdateAzurednstaskgroupResponseAsObject struct {
	Result               *Azurednstaskgroup `json:"result,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _UpdateAzurednstaskgroupResponseAsObject UpdateAzurednstaskgroupResponseAsObject

// NewUpdateAzurednstaskgroupResponseAsObject instantiates a new UpdateAzurednstaskgroupResponseAsObject object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateAzurednstaskgroupResponseAsObject() *UpdateAzurednstaskgroupResponseAsObject {
	this := UpdateAzurednstaskgroupResponseAsObject{}
	return &this
}

// NewUpdateAzurednstaskgroupResponseAsObjectWithDefaults instantiates a new UpdateAzurednstaskgroupResponseAsObject object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateAzurednstaskgroupResponseAsObjectWithDefaults() *UpdateAzurednstaskgroupResponseAsObject {
	this := UpdateAzurednstaskgroupResponseAsObject{}
	return &this
}

// GetResult returns the Result field value if set, zero value otherwise.
func (o *UpdateAzurednstaskgroupResponseAsObject) GetResult() Azurednstaskgroup {
	if o == nil || IsNil(o.Result) {
		var ret Azurednstaskgroup
		return ret
	}
	return *o.Result
}

// GetResultOk returns a tuple with the Result field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateAzurednstaskgroupResponseAsObject) GetResultOk() (*Azurednstaskgroup, bool) {
	if o == nil || IsNil(o.Result) {
		return nil, false
	}
	return o.Result, true
}

// HasResult returns a boolean if a field has been set.
func (o *UpdateAzurednstaskgroupResponseAsObject) HasResult() bool {
	if o != nil && !IsNil(o.Result) {
		return true
	}

	return false
}

// SetResult gets a reference to the given Azurednstaskgroup and assigns it to the Result field.
func (o *UpdateAzurednstaskgroupResponseAsObject) SetResult(v Azurednstaskgroup) {
	o.Result = &v
}

func (o UpdateAzurednstaskgroupResponseAsObject) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateAzurednstaskgroupResponseAsObject) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Result) {
		toSerialize["result"] = o.Result
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *UpdateAzurednstaskgroupResponseAsObject) UnmarshalJSON(data []byte) (err error) {
	varUpdateAzurednstaskgroupResponseAsObject := _UpdateAzurednstaskgroupResponseAsObject{}

	err = json.Unmarshal(data, &varUpdateAzurednstaskgroupResponseAsObject)

	if err != nil {
		return err
	}

	*o = UpdateAzurednstaskgroupResponseAsObject(varUpdateAzurednstaskgroupResponseAsObject)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "result")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableUpdateAzurednstaskgroupResponseAsObject struct {
	value *UpdateAzurednstaskgroupResponseAsObject
	isSet bool
}

func (v NullableUpdateAzurednstaskgroupResponseAsObject) Get() *UpdateAzurednstaskgroupResponseAsObject {
	return v.value
}

func (v *NullableUpdateAzurednstaskgroupResponseAsObject) Set(val *UpdateAzurednstaskgroupResponseAsObject) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateAzurednstaskgroupResponseAsObject) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateAzurednstaskgroupResponseAsObject) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateAzurednstaskgroupResponseAsObject(val *UpdateAzurednstaskgroupResponseAsObject) *NullableUpdateAzurednstaskgroupResponseAsObject {
	return &NullableUpdateAzurednstaskgroupResponseAsObject{value: val, isSet: true}
}

func (v NullableUpdateAzurednstaskgroupResponseAsObject) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateAzurednstaskgroupResponseAsObject) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
