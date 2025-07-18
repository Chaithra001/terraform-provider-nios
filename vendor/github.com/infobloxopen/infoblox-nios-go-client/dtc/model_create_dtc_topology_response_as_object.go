/*
Infoblox DTC API

OpenAPI specification for Infoblox NIOS WAPI DTC objects

API version: 2.13.6
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package dtc

import (
	"encoding/json"
)

// checks if the CreateDtcTopologyResponseAsObject type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateDtcTopologyResponseAsObject{}

// CreateDtcTopologyResponseAsObject The response format to create __DtcTopology__ in object format.
type CreateDtcTopologyResponseAsObject struct {
	Result               *DtcTopology `json:"result,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CreateDtcTopologyResponseAsObject CreateDtcTopologyResponseAsObject

// NewCreateDtcTopologyResponseAsObject instantiates a new CreateDtcTopologyResponseAsObject object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateDtcTopologyResponseAsObject() *CreateDtcTopologyResponseAsObject {
	this := CreateDtcTopologyResponseAsObject{}
	return &this
}

// NewCreateDtcTopologyResponseAsObjectWithDefaults instantiates a new CreateDtcTopologyResponseAsObject object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateDtcTopologyResponseAsObjectWithDefaults() *CreateDtcTopologyResponseAsObject {
	this := CreateDtcTopologyResponseAsObject{}
	return &this
}

// GetResult returns the Result field value if set, zero value otherwise.
func (o *CreateDtcTopologyResponseAsObject) GetResult() DtcTopology {
	if o == nil || IsNil(o.Result) {
		var ret DtcTopology
		return ret
	}
	return *o.Result
}

// GetResultOk returns a tuple with the Result field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateDtcTopologyResponseAsObject) GetResultOk() (*DtcTopology, bool) {
	if o == nil || IsNil(o.Result) {
		return nil, false
	}
	return o.Result, true
}

// HasResult returns a boolean if a field has been set.
func (o *CreateDtcTopologyResponseAsObject) HasResult() bool {
	if o != nil && !IsNil(o.Result) {
		return true
	}

	return false
}

// SetResult gets a reference to the given DtcTopology and assigns it to the Result field.
func (o *CreateDtcTopologyResponseAsObject) SetResult(v DtcTopology) {
	o.Result = &v
}

func (o CreateDtcTopologyResponseAsObject) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateDtcTopologyResponseAsObject) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Result) {
		toSerialize["result"] = o.Result
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CreateDtcTopologyResponseAsObject) UnmarshalJSON(data []byte) (err error) {
	varCreateDtcTopologyResponseAsObject := _CreateDtcTopologyResponseAsObject{}

	err = json.Unmarshal(data, &varCreateDtcTopologyResponseAsObject)

	if err != nil {
		return err
	}

	*o = CreateDtcTopologyResponseAsObject(varCreateDtcTopologyResponseAsObject)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "result")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCreateDtcTopologyResponseAsObject struct {
	value *CreateDtcTopologyResponseAsObject
	isSet bool
}

func (v NullableCreateDtcTopologyResponseAsObject) Get() *CreateDtcTopologyResponseAsObject {
	return v.value
}

func (v *NullableCreateDtcTopologyResponseAsObject) Set(val *CreateDtcTopologyResponseAsObject) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateDtcTopologyResponseAsObject) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateDtcTopologyResponseAsObject) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateDtcTopologyResponseAsObject(val *CreateDtcTopologyResponseAsObject) *NullableCreateDtcTopologyResponseAsObject {
	return &NullableCreateDtcTopologyResponseAsObject{value: val, isSet: true}
}

func (v NullableCreateDtcTopologyResponseAsObject) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateDtcTopologyResponseAsObject) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
