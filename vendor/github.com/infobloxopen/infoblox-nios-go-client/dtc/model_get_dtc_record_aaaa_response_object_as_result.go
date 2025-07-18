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

// checks if the GetDtcRecordAaaaResponseObjectAsResult type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetDtcRecordAaaaResponseObjectAsResult{}

// GetDtcRecordAaaaResponseObjectAsResult The response format to retrieve __DtcRecordAaaa__ objects.
type GetDtcRecordAaaaResponseObjectAsResult struct {
	Result *DtcRecordAaaa `json:"result,omitempty"`
}

// NewGetDtcRecordAaaaResponseObjectAsResult instantiates a new GetDtcRecordAaaaResponseObjectAsResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetDtcRecordAaaaResponseObjectAsResult() *GetDtcRecordAaaaResponseObjectAsResult {
	this := GetDtcRecordAaaaResponseObjectAsResult{}
	return &this
}

// NewGetDtcRecordAaaaResponseObjectAsResultWithDefaults instantiates a new GetDtcRecordAaaaResponseObjectAsResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetDtcRecordAaaaResponseObjectAsResultWithDefaults() *GetDtcRecordAaaaResponseObjectAsResult {
	this := GetDtcRecordAaaaResponseObjectAsResult{}
	return &this
}

// GetResult returns the Result field value if set, zero value otherwise.
func (o *GetDtcRecordAaaaResponseObjectAsResult) GetResult() DtcRecordAaaa {
	if o == nil || IsNil(o.Result) {
		var ret DtcRecordAaaa
		return ret
	}
	return *o.Result
}

// GetResultOk returns a tuple with the Result field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetDtcRecordAaaaResponseObjectAsResult) GetResultOk() (*DtcRecordAaaa, bool) {
	if o == nil || IsNil(o.Result) {
		return nil, false
	}
	return o.Result, true
}

// HasResult returns a boolean if a field has been set.
func (o *GetDtcRecordAaaaResponseObjectAsResult) HasResult() bool {
	if o != nil && !IsNil(o.Result) {
		return true
	}

	return false
}

// SetResult gets a reference to the given DtcRecordAaaa and assigns it to the Result field.
func (o *GetDtcRecordAaaaResponseObjectAsResult) SetResult(v DtcRecordAaaa) {
	o.Result = &v
}

func (o GetDtcRecordAaaaResponseObjectAsResult) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetDtcRecordAaaaResponseObjectAsResult) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Result) {
		toSerialize["result"] = o.Result
	}
	return toSerialize, nil
}

type NullableGetDtcRecordAaaaResponseObjectAsResult struct {
	value *GetDtcRecordAaaaResponseObjectAsResult
	isSet bool
}

func (v NullableGetDtcRecordAaaaResponseObjectAsResult) Get() *GetDtcRecordAaaaResponseObjectAsResult {
	return v.value
}

func (v *NullableGetDtcRecordAaaaResponseObjectAsResult) Set(val *GetDtcRecordAaaaResponseObjectAsResult) {
	v.value = val
	v.isSet = true
}

func (v NullableGetDtcRecordAaaaResponseObjectAsResult) IsSet() bool {
	return v.isSet
}

func (v *NullableGetDtcRecordAaaaResponseObjectAsResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetDtcRecordAaaaResponseObjectAsResult(val *GetDtcRecordAaaaResponseObjectAsResult) *NullableGetDtcRecordAaaaResponseObjectAsResult {
	return &NullableGetDtcRecordAaaaResponseObjectAsResult{value: val, isSet: true}
}

func (v NullableGetDtcRecordAaaaResponseObjectAsResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetDtcRecordAaaaResponseObjectAsResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
