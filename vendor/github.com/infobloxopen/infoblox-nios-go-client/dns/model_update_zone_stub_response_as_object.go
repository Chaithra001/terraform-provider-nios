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

// checks if the UpdateZoneStubResponseAsObject type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateZoneStubResponseAsObject{}

// UpdateZoneStubResponseAsObject The response format to update __ZoneStub__ in object format.
type UpdateZoneStubResponseAsObject struct {
	Result               *ZoneStub `json:"result,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _UpdateZoneStubResponseAsObject UpdateZoneStubResponseAsObject

// NewUpdateZoneStubResponseAsObject instantiates a new UpdateZoneStubResponseAsObject object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateZoneStubResponseAsObject() *UpdateZoneStubResponseAsObject {
	this := UpdateZoneStubResponseAsObject{}
	return &this
}

// NewUpdateZoneStubResponseAsObjectWithDefaults instantiates a new UpdateZoneStubResponseAsObject object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateZoneStubResponseAsObjectWithDefaults() *UpdateZoneStubResponseAsObject {
	this := UpdateZoneStubResponseAsObject{}
	return &this
}

// GetResult returns the Result field value if set, zero value otherwise.
func (o *UpdateZoneStubResponseAsObject) GetResult() ZoneStub {
	if o == nil || IsNil(o.Result) {
		var ret ZoneStub
		return ret
	}
	return *o.Result
}

// GetResultOk returns a tuple with the Result field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateZoneStubResponseAsObject) GetResultOk() (*ZoneStub, bool) {
	if o == nil || IsNil(o.Result) {
		return nil, false
	}
	return o.Result, true
}

// HasResult returns a boolean if a field has been set.
func (o *UpdateZoneStubResponseAsObject) HasResult() bool {
	if o != nil && !IsNil(o.Result) {
		return true
	}

	return false
}

// SetResult gets a reference to the given ZoneStub and assigns it to the Result field.
func (o *UpdateZoneStubResponseAsObject) SetResult(v ZoneStub) {
	o.Result = &v
}

func (o UpdateZoneStubResponseAsObject) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateZoneStubResponseAsObject) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Result) {
		toSerialize["result"] = o.Result
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *UpdateZoneStubResponseAsObject) UnmarshalJSON(data []byte) (err error) {
	varUpdateZoneStubResponseAsObject := _UpdateZoneStubResponseAsObject{}

	err = json.Unmarshal(data, &varUpdateZoneStubResponseAsObject)

	if err != nil {
		return err
	}

	*o = UpdateZoneStubResponseAsObject(varUpdateZoneStubResponseAsObject)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "result")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableUpdateZoneStubResponseAsObject struct {
	value *UpdateZoneStubResponseAsObject
	isSet bool
}

func (v NullableUpdateZoneStubResponseAsObject) Get() *UpdateZoneStubResponseAsObject {
	return v.value
}

func (v *NullableUpdateZoneStubResponseAsObject) Set(val *UpdateZoneStubResponseAsObject) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateZoneStubResponseAsObject) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateZoneStubResponseAsObject) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateZoneStubResponseAsObject(val *UpdateZoneStubResponseAsObject) *NullableUpdateZoneStubResponseAsObject {
	return &NullableUpdateZoneStubResponseAsObject{value: val, isSet: true}
}

func (v NullableUpdateZoneStubResponseAsObject) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateZoneStubResponseAsObject) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
