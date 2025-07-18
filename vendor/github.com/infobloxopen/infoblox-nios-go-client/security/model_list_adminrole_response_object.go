/*
Infoblox SECURITY API

OpenAPI specification for Infoblox NIOS WAPI SECURITY objects

API version: 2.13.6
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package security

import (
	"encoding/json"
)

// checks if the ListAdminroleResponseObject type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListAdminroleResponseObject{}

// ListAdminroleResponseObject The response format to retrieve __Adminrole__ objects.
type ListAdminroleResponseObject struct {
	Result               []Adminrole `json:"result,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ListAdminroleResponseObject ListAdminroleResponseObject

// NewListAdminroleResponseObject instantiates a new ListAdminroleResponseObject object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListAdminroleResponseObject() *ListAdminroleResponseObject {
	this := ListAdminroleResponseObject{}
	return &this
}

// NewListAdminroleResponseObjectWithDefaults instantiates a new ListAdminroleResponseObject object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListAdminroleResponseObjectWithDefaults() *ListAdminroleResponseObject {
	this := ListAdminroleResponseObject{}
	return &this
}

// GetResult returns the Result field value if set, zero value otherwise.
func (o *ListAdminroleResponseObject) GetResult() []Adminrole {
	if o == nil || IsNil(o.Result) {
		var ret []Adminrole
		return ret
	}
	return o.Result
}

// GetResultOk returns a tuple with the Result field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListAdminroleResponseObject) GetResultOk() ([]Adminrole, bool) {
	if o == nil || IsNil(o.Result) {
		return nil, false
	}
	return o.Result, true
}

// HasResult returns a boolean if a field has been set.
func (o *ListAdminroleResponseObject) HasResult() bool {
	if o != nil && !IsNil(o.Result) {
		return true
	}

	return false
}

// SetResult gets a reference to the given []Adminrole and assigns it to the Result field.
func (o *ListAdminroleResponseObject) SetResult(v []Adminrole) {
	o.Result = v
}

func (o ListAdminroleResponseObject) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListAdminroleResponseObject) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Result) {
		toSerialize["result"] = o.Result
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ListAdminroleResponseObject) UnmarshalJSON(data []byte) (err error) {
	varListAdminroleResponseObject := _ListAdminroleResponseObject{}

	err = json.Unmarshal(data, &varListAdminroleResponseObject)

	if err != nil {
		return err
	}

	*o = ListAdminroleResponseObject(varListAdminroleResponseObject)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "result")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableListAdminroleResponseObject struct {
	value *ListAdminroleResponseObject
	isSet bool
}

func (v NullableListAdminroleResponseObject) Get() *ListAdminroleResponseObject {
	return v.value
}

func (v *NullableListAdminroleResponseObject) Set(val *ListAdminroleResponseObject) {
	v.value = val
	v.isSet = true
}

func (v NullableListAdminroleResponseObject) IsSet() bool {
	return v.isSet
}

func (v *NullableListAdminroleResponseObject) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListAdminroleResponseObject(val *ListAdminroleResponseObject) *NullableListAdminroleResponseObject {
	return &NullableListAdminroleResponseObject{value: val, isSet: true}
}

func (v NullableListAdminroleResponseObject) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListAdminroleResponseObject) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
