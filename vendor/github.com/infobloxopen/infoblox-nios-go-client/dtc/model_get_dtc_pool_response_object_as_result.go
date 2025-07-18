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

// checks if the GetDtcPoolResponseObjectAsResult type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetDtcPoolResponseObjectAsResult{}

// GetDtcPoolResponseObjectAsResult The response format to retrieve __DtcPool__ objects.
type GetDtcPoolResponseObjectAsResult struct {
	Result *DtcPool `json:"result,omitempty"`
}

// NewGetDtcPoolResponseObjectAsResult instantiates a new GetDtcPoolResponseObjectAsResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetDtcPoolResponseObjectAsResult() *GetDtcPoolResponseObjectAsResult {
	this := GetDtcPoolResponseObjectAsResult{}
	return &this
}

// NewGetDtcPoolResponseObjectAsResultWithDefaults instantiates a new GetDtcPoolResponseObjectAsResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetDtcPoolResponseObjectAsResultWithDefaults() *GetDtcPoolResponseObjectAsResult {
	this := GetDtcPoolResponseObjectAsResult{}
	return &this
}

// GetResult returns the Result field value if set, zero value otherwise.
func (o *GetDtcPoolResponseObjectAsResult) GetResult() DtcPool {
	if o == nil || IsNil(o.Result) {
		var ret DtcPool
		return ret
	}
	return *o.Result
}

// GetResultOk returns a tuple with the Result field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetDtcPoolResponseObjectAsResult) GetResultOk() (*DtcPool, bool) {
	if o == nil || IsNil(o.Result) {
		return nil, false
	}
	return o.Result, true
}

// HasResult returns a boolean if a field has been set.
func (o *GetDtcPoolResponseObjectAsResult) HasResult() bool {
	if o != nil && !IsNil(o.Result) {
		return true
	}

	return false
}

// SetResult gets a reference to the given DtcPool and assigns it to the Result field.
func (o *GetDtcPoolResponseObjectAsResult) SetResult(v DtcPool) {
	o.Result = &v
}

func (o GetDtcPoolResponseObjectAsResult) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetDtcPoolResponseObjectAsResult) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Result) {
		toSerialize["result"] = o.Result
	}
	return toSerialize, nil
}

type NullableGetDtcPoolResponseObjectAsResult struct {
	value *GetDtcPoolResponseObjectAsResult
	isSet bool
}

func (v NullableGetDtcPoolResponseObjectAsResult) Get() *GetDtcPoolResponseObjectAsResult {
	return v.value
}

func (v *NullableGetDtcPoolResponseObjectAsResult) Set(val *GetDtcPoolResponseObjectAsResult) {
	v.value = val
	v.isSet = true
}

func (v NullableGetDtcPoolResponseObjectAsResult) IsSet() bool {
	return v.isSet
}

func (v *NullableGetDtcPoolResponseObjectAsResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetDtcPoolResponseObjectAsResult(val *GetDtcPoolResponseObjectAsResult) *NullableGetDtcPoolResponseObjectAsResult {
	return &NullableGetDtcPoolResponseObjectAsResult{value: val, isSet: true}
}

func (v NullableGetDtcPoolResponseObjectAsResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetDtcPoolResponseObjectAsResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
