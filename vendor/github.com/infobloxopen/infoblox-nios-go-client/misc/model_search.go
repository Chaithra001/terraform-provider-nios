/*
Infoblox MISC API

OpenAPI specification for Infoblox NIOS WAPI MISC objects

API version: 2.13.6
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package misc

import (
	"encoding/json"
)

// checks if the Search type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Search{}

// Search struct for Search
type Search struct {
	// The reference to the object.
	Ref *string `json:"_ref,omitempty"`
}

// NewSearch instantiates a new Search object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSearch() *Search {
	this := Search{}
	return &this
}

// NewSearchWithDefaults instantiates a new Search object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSearchWithDefaults() *Search {
	this := Search{}
	return &this
}

// GetRef returns the Ref field value if set, zero value otherwise.
func (o *Search) GetRef() string {
	if o == nil || IsNil(o.Ref) {
		var ret string
		return ret
	}
	return *o.Ref
}

// GetRefOk returns a tuple with the Ref field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Search) GetRefOk() (*string, bool) {
	if o == nil || IsNil(o.Ref) {
		return nil, false
	}
	return o.Ref, true
}

// HasRef returns a boolean if a field has been set.
func (o *Search) HasRef() bool {
	if o != nil && !IsNil(o.Ref) {
		return true
	}

	return false
}

// SetRef gets a reference to the given string and assigns it to the Ref field.
func (o *Search) SetRef(v string) {
	o.Ref = &v
}

func (o Search) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Search) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Ref) {
		toSerialize["_ref"] = o.Ref
	}
	return toSerialize, nil
}

type NullableSearch struct {
	value *Search
	isSet bool
}

func (v NullableSearch) Get() *Search {
	return v.value
}

func (v *NullableSearch) Set(val *Search) {
	v.value = val
	v.isSet = true
}

func (v NullableSearch) IsSet() bool {
	return v.isSet
}

func (v *NullableSearch) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSearch(val *Search) *NullableSearch {
	return &NullableSearch{value: val, isSet: true}
}

func (v NullableSearch) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSearch) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
