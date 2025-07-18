/*
Infoblox DISCOVERY API

OpenAPI specification for Infoblox NIOS WAPI DISCOVERY objects

API version: 2.13.6
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package discovery

import (
	"encoding/json"
)

// checks if the ListDiscoveryStatusResponseObject type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListDiscoveryStatusResponseObject{}

// ListDiscoveryStatusResponseObject The response format to retrieve __DiscoveryStatus__ objects.
type ListDiscoveryStatusResponseObject struct {
	Result               []DiscoveryStatus `json:"result,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ListDiscoveryStatusResponseObject ListDiscoveryStatusResponseObject

// NewListDiscoveryStatusResponseObject instantiates a new ListDiscoveryStatusResponseObject object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListDiscoveryStatusResponseObject() *ListDiscoveryStatusResponseObject {
	this := ListDiscoveryStatusResponseObject{}
	return &this
}

// NewListDiscoveryStatusResponseObjectWithDefaults instantiates a new ListDiscoveryStatusResponseObject object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListDiscoveryStatusResponseObjectWithDefaults() *ListDiscoveryStatusResponseObject {
	this := ListDiscoveryStatusResponseObject{}
	return &this
}

// GetResult returns the Result field value if set, zero value otherwise.
func (o *ListDiscoveryStatusResponseObject) GetResult() []DiscoveryStatus {
	if o == nil || IsNil(o.Result) {
		var ret []DiscoveryStatus
		return ret
	}
	return o.Result
}

// GetResultOk returns a tuple with the Result field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListDiscoveryStatusResponseObject) GetResultOk() ([]DiscoveryStatus, bool) {
	if o == nil || IsNil(o.Result) {
		return nil, false
	}
	return o.Result, true
}

// HasResult returns a boolean if a field has been set.
func (o *ListDiscoveryStatusResponseObject) HasResult() bool {
	if o != nil && !IsNil(o.Result) {
		return true
	}

	return false
}

// SetResult gets a reference to the given []DiscoveryStatus and assigns it to the Result field.
func (o *ListDiscoveryStatusResponseObject) SetResult(v []DiscoveryStatus) {
	o.Result = v
}

func (o ListDiscoveryStatusResponseObject) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListDiscoveryStatusResponseObject) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Result) {
		toSerialize["result"] = o.Result
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ListDiscoveryStatusResponseObject) UnmarshalJSON(data []byte) (err error) {
	varListDiscoveryStatusResponseObject := _ListDiscoveryStatusResponseObject{}

	err = json.Unmarshal(data, &varListDiscoveryStatusResponseObject)

	if err != nil {
		return err
	}

	*o = ListDiscoveryStatusResponseObject(varListDiscoveryStatusResponseObject)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "result")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableListDiscoveryStatusResponseObject struct {
	value *ListDiscoveryStatusResponseObject
	isSet bool
}

func (v NullableListDiscoveryStatusResponseObject) Get() *ListDiscoveryStatusResponseObject {
	return v.value
}

func (v *NullableListDiscoveryStatusResponseObject) Set(val *ListDiscoveryStatusResponseObject) {
	v.value = val
	v.isSet = true
}

func (v NullableListDiscoveryStatusResponseObject) IsSet() bool {
	return v.isSet
}

func (v *NullableListDiscoveryStatusResponseObject) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListDiscoveryStatusResponseObject(val *ListDiscoveryStatusResponseObject) *NullableListDiscoveryStatusResponseObject {
	return &NullableListDiscoveryStatusResponseObject{value: val, isSet: true}
}

func (v NullableListDiscoveryStatusResponseObject) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListDiscoveryStatusResponseObject) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
