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

// checks if the GetRecordPtrResponseObjectAsResult type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetRecordPtrResponseObjectAsResult{}

// GetRecordPtrResponseObjectAsResult The response format to retrieve __RecordPtr__ objects.
type GetRecordPtrResponseObjectAsResult struct {
	Result *RecordPtr `json:"result,omitempty"`
}

// NewGetRecordPtrResponseObjectAsResult instantiates a new GetRecordPtrResponseObjectAsResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetRecordPtrResponseObjectAsResult() *GetRecordPtrResponseObjectAsResult {
	this := GetRecordPtrResponseObjectAsResult{}
	return &this
}

// NewGetRecordPtrResponseObjectAsResultWithDefaults instantiates a new GetRecordPtrResponseObjectAsResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetRecordPtrResponseObjectAsResultWithDefaults() *GetRecordPtrResponseObjectAsResult {
	this := GetRecordPtrResponseObjectAsResult{}
	return &this
}

// GetResult returns the Result field value if set, zero value otherwise.
func (o *GetRecordPtrResponseObjectAsResult) GetResult() RecordPtr {
	if o == nil || IsNil(o.Result) {
		var ret RecordPtr
		return ret
	}
	return *o.Result
}

// GetResultOk returns a tuple with the Result field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetRecordPtrResponseObjectAsResult) GetResultOk() (*RecordPtr, bool) {
	if o == nil || IsNil(o.Result) {
		return nil, false
	}
	return o.Result, true
}

// HasResult returns a boolean if a field has been set.
func (o *GetRecordPtrResponseObjectAsResult) HasResult() bool {
	if o != nil && !IsNil(o.Result) {
		return true
	}

	return false
}

// SetResult gets a reference to the given RecordPtr and assigns it to the Result field.
func (o *GetRecordPtrResponseObjectAsResult) SetResult(v RecordPtr) {
	o.Result = &v
}

func (o GetRecordPtrResponseObjectAsResult) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetRecordPtrResponseObjectAsResult) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Result) {
		toSerialize["result"] = o.Result
	}
	return toSerialize, nil
}

type NullableGetRecordPtrResponseObjectAsResult struct {
	value *GetRecordPtrResponseObjectAsResult
	isSet bool
}

func (v NullableGetRecordPtrResponseObjectAsResult) Get() *GetRecordPtrResponseObjectAsResult {
	return v.value
}

func (v *NullableGetRecordPtrResponseObjectAsResult) Set(val *GetRecordPtrResponseObjectAsResult) {
	v.value = val
	v.isSet = true
}

func (v NullableGetRecordPtrResponseObjectAsResult) IsSet() bool {
	return v.isSet
}

func (v *NullableGetRecordPtrResponseObjectAsResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetRecordPtrResponseObjectAsResult(val *GetRecordPtrResponseObjectAsResult) *NullableGetRecordPtrResponseObjectAsResult {
	return &NullableGetRecordPtrResponseObjectAsResult{value: val, isSet: true}
}

func (v NullableGetRecordPtrResponseObjectAsResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetRecordPtrResponseObjectAsResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
