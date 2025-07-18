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

// checks if the ListNsgroupForwardstubserverResponseObject type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListNsgroupForwardstubserverResponseObject{}

// ListNsgroupForwardstubserverResponseObject The response format to retrieve __NsgroupForwardstubserver__ objects.
type ListNsgroupForwardstubserverResponseObject struct {
	Result               []NsgroupForwardstubserver `json:"result,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ListNsgroupForwardstubserverResponseObject ListNsgroupForwardstubserverResponseObject

// NewListNsgroupForwardstubserverResponseObject instantiates a new ListNsgroupForwardstubserverResponseObject object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListNsgroupForwardstubserverResponseObject() *ListNsgroupForwardstubserverResponseObject {
	this := ListNsgroupForwardstubserverResponseObject{}
	return &this
}

// NewListNsgroupForwardstubserverResponseObjectWithDefaults instantiates a new ListNsgroupForwardstubserverResponseObject object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListNsgroupForwardstubserverResponseObjectWithDefaults() *ListNsgroupForwardstubserverResponseObject {
	this := ListNsgroupForwardstubserverResponseObject{}
	return &this
}

// GetResult returns the Result field value if set, zero value otherwise.
func (o *ListNsgroupForwardstubserverResponseObject) GetResult() []NsgroupForwardstubserver {
	if o == nil || IsNil(o.Result) {
		var ret []NsgroupForwardstubserver
		return ret
	}
	return o.Result
}

// GetResultOk returns a tuple with the Result field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListNsgroupForwardstubserverResponseObject) GetResultOk() ([]NsgroupForwardstubserver, bool) {
	if o == nil || IsNil(o.Result) {
		return nil, false
	}
	return o.Result, true
}

// HasResult returns a boolean if a field has been set.
func (o *ListNsgroupForwardstubserverResponseObject) HasResult() bool {
	if o != nil && !IsNil(o.Result) {
		return true
	}

	return false
}

// SetResult gets a reference to the given []NsgroupForwardstubserver and assigns it to the Result field.
func (o *ListNsgroupForwardstubserverResponseObject) SetResult(v []NsgroupForwardstubserver) {
	o.Result = v
}

func (o ListNsgroupForwardstubserverResponseObject) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListNsgroupForwardstubserverResponseObject) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Result) {
		toSerialize["result"] = o.Result
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ListNsgroupForwardstubserverResponseObject) UnmarshalJSON(data []byte) (err error) {
	varListNsgroupForwardstubserverResponseObject := _ListNsgroupForwardstubserverResponseObject{}

	err = json.Unmarshal(data, &varListNsgroupForwardstubserverResponseObject)

	if err != nil {
		return err
	}

	*o = ListNsgroupForwardstubserverResponseObject(varListNsgroupForwardstubserverResponseObject)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "result")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableListNsgroupForwardstubserverResponseObject struct {
	value *ListNsgroupForwardstubserverResponseObject
	isSet bool
}

func (v NullableListNsgroupForwardstubserverResponseObject) Get() *ListNsgroupForwardstubserverResponseObject {
	return v.value
}

func (v *NullableListNsgroupForwardstubserverResponseObject) Set(val *ListNsgroupForwardstubserverResponseObject) {
	v.value = val
	v.isSet = true
}

func (v NullableListNsgroupForwardstubserverResponseObject) IsSet() bool {
	return v.isSet
}

func (v *NullableListNsgroupForwardstubserverResponseObject) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListNsgroupForwardstubserverResponseObject(val *ListNsgroupForwardstubserverResponseObject) *NullableListNsgroupForwardstubserverResponseObject {
	return &NullableListNsgroupForwardstubserverResponseObject{value: val, isSet: true}
}

func (v NullableListNsgroupForwardstubserverResponseObject) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListNsgroupForwardstubserverResponseObject) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
