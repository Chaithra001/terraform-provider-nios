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

// checks if the GetGmcscheduleResponseObjectAsResult type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetGmcscheduleResponseObjectAsResult{}

// GetGmcscheduleResponseObjectAsResult The response format to retrieve __Gmcschedule__ objects.
type GetGmcscheduleResponseObjectAsResult struct {
	Result *Gmcschedule `json:"result,omitempty"`
}

// NewGetGmcscheduleResponseObjectAsResult instantiates a new GetGmcscheduleResponseObjectAsResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetGmcscheduleResponseObjectAsResult() *GetGmcscheduleResponseObjectAsResult {
	this := GetGmcscheduleResponseObjectAsResult{}
	return &this
}

// NewGetGmcscheduleResponseObjectAsResultWithDefaults instantiates a new GetGmcscheduleResponseObjectAsResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetGmcscheduleResponseObjectAsResultWithDefaults() *GetGmcscheduleResponseObjectAsResult {
	this := GetGmcscheduleResponseObjectAsResult{}
	return &this
}

// GetResult returns the Result field value if set, zero value otherwise.
func (o *GetGmcscheduleResponseObjectAsResult) GetResult() Gmcschedule {
	if o == nil || IsNil(o.Result) {
		var ret Gmcschedule
		return ret
	}
	return *o.Result
}

// GetResultOk returns a tuple with the Result field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetGmcscheduleResponseObjectAsResult) GetResultOk() (*Gmcschedule, bool) {
	if o == nil || IsNil(o.Result) {
		return nil, false
	}
	return o.Result, true
}

// HasResult returns a boolean if a field has been set.
func (o *GetGmcscheduleResponseObjectAsResult) HasResult() bool {
	if o != nil && !IsNil(o.Result) {
		return true
	}

	return false
}

// SetResult gets a reference to the given Gmcschedule and assigns it to the Result field.
func (o *GetGmcscheduleResponseObjectAsResult) SetResult(v Gmcschedule) {
	o.Result = &v
}

func (o GetGmcscheduleResponseObjectAsResult) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetGmcscheduleResponseObjectAsResult) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Result) {
		toSerialize["result"] = o.Result
	}
	return toSerialize, nil
}

type NullableGetGmcscheduleResponseObjectAsResult struct {
	value *GetGmcscheduleResponseObjectAsResult
	isSet bool
}

func (v NullableGetGmcscheduleResponseObjectAsResult) Get() *GetGmcscheduleResponseObjectAsResult {
	return v.value
}

func (v *NullableGetGmcscheduleResponseObjectAsResult) Set(val *GetGmcscheduleResponseObjectAsResult) {
	v.value = val
	v.isSet = true
}

func (v NullableGetGmcscheduleResponseObjectAsResult) IsSet() bool {
	return v.isSet
}

func (v *NullableGetGmcscheduleResponseObjectAsResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetGmcscheduleResponseObjectAsResult(val *GetGmcscheduleResponseObjectAsResult) *NullableGetGmcscheduleResponseObjectAsResult {
	return &NullableGetGmcscheduleResponseObjectAsResult{value: val, isSet: true}
}

func (v NullableGetGmcscheduleResponseObjectAsResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetGmcscheduleResponseObjectAsResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
