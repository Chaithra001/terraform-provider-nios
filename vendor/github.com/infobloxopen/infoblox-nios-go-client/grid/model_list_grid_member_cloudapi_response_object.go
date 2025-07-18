/*
Infoblox GRID API

OpenAPI specification for Infoblox NIOS WAPI GRID objects

API version: 2.13.6
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package grid

import (
	"encoding/json"
)

// checks if the ListGridMemberCloudapiResponseObject type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListGridMemberCloudapiResponseObject{}

// ListGridMemberCloudapiResponseObject The response format to retrieve __GridMemberCloudapi__ objects.
type ListGridMemberCloudapiResponseObject struct {
	Result               []GridMemberCloudapi `json:"result,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ListGridMemberCloudapiResponseObject ListGridMemberCloudapiResponseObject

// NewListGridMemberCloudapiResponseObject instantiates a new ListGridMemberCloudapiResponseObject object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListGridMemberCloudapiResponseObject() *ListGridMemberCloudapiResponseObject {
	this := ListGridMemberCloudapiResponseObject{}
	return &this
}

// NewListGridMemberCloudapiResponseObjectWithDefaults instantiates a new ListGridMemberCloudapiResponseObject object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListGridMemberCloudapiResponseObjectWithDefaults() *ListGridMemberCloudapiResponseObject {
	this := ListGridMemberCloudapiResponseObject{}
	return &this
}

// GetResult returns the Result field value if set, zero value otherwise.
func (o *ListGridMemberCloudapiResponseObject) GetResult() []GridMemberCloudapi {
	if o == nil || IsNil(o.Result) {
		var ret []GridMemberCloudapi
		return ret
	}
	return o.Result
}

// GetResultOk returns a tuple with the Result field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListGridMemberCloudapiResponseObject) GetResultOk() ([]GridMemberCloudapi, bool) {
	if o == nil || IsNil(o.Result) {
		return nil, false
	}
	return o.Result, true
}

// HasResult returns a boolean if a field has been set.
func (o *ListGridMemberCloudapiResponseObject) HasResult() bool {
	if o != nil && !IsNil(o.Result) {
		return true
	}

	return false
}

// SetResult gets a reference to the given []GridMemberCloudapi and assigns it to the Result field.
func (o *ListGridMemberCloudapiResponseObject) SetResult(v []GridMemberCloudapi) {
	o.Result = v
}

func (o ListGridMemberCloudapiResponseObject) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListGridMemberCloudapiResponseObject) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Result) {
		toSerialize["result"] = o.Result
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ListGridMemberCloudapiResponseObject) UnmarshalJSON(data []byte) (err error) {
	varListGridMemberCloudapiResponseObject := _ListGridMemberCloudapiResponseObject{}

	err = json.Unmarshal(data, &varListGridMemberCloudapiResponseObject)

	if err != nil {
		return err
	}

	*o = ListGridMemberCloudapiResponseObject(varListGridMemberCloudapiResponseObject)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "result")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableListGridMemberCloudapiResponseObject struct {
	value *ListGridMemberCloudapiResponseObject
	isSet bool
}

func (v NullableListGridMemberCloudapiResponseObject) Get() *ListGridMemberCloudapiResponseObject {
	return v.value
}

func (v *NullableListGridMemberCloudapiResponseObject) Set(val *ListGridMemberCloudapiResponseObject) {
	v.value = val
	v.isSet = true
}

func (v NullableListGridMemberCloudapiResponseObject) IsSet() bool {
	return v.isSet
}

func (v *NullableListGridMemberCloudapiResponseObject) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListGridMemberCloudapiResponseObject(val *ListGridMemberCloudapiResponseObject) *NullableListGridMemberCloudapiResponseObject {
	return &NullableListGridMemberCloudapiResponseObject{value: val, isSet: true}
}

func (v NullableListGridMemberCloudapiResponseObject) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListGridMemberCloudapiResponseObject) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
