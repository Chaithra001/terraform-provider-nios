/*
Infoblox MICROSOFTSERVER API

OpenAPI specification for Infoblox NIOS WAPI MICROSOFTSERVER objects

API version: 2.13.6
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package microsoftserver

import (
	"encoding/json"
)

// checks if the GetMsserverDhcpResponseObjectAsResult type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetMsserverDhcpResponseObjectAsResult{}

// GetMsserverDhcpResponseObjectAsResult The response format to retrieve __MsserverDhcp__ objects.
type GetMsserverDhcpResponseObjectAsResult struct {
	Result *MsserverDhcp `json:"result,omitempty"`
}

// NewGetMsserverDhcpResponseObjectAsResult instantiates a new GetMsserverDhcpResponseObjectAsResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetMsserverDhcpResponseObjectAsResult() *GetMsserverDhcpResponseObjectAsResult {
	this := GetMsserverDhcpResponseObjectAsResult{}
	return &this
}

// NewGetMsserverDhcpResponseObjectAsResultWithDefaults instantiates a new GetMsserverDhcpResponseObjectAsResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetMsserverDhcpResponseObjectAsResultWithDefaults() *GetMsserverDhcpResponseObjectAsResult {
	this := GetMsserverDhcpResponseObjectAsResult{}
	return &this
}

// GetResult returns the Result field value if set, zero value otherwise.
func (o *GetMsserverDhcpResponseObjectAsResult) GetResult() MsserverDhcp {
	if o == nil || IsNil(o.Result) {
		var ret MsserverDhcp
		return ret
	}
	return *o.Result
}

// GetResultOk returns a tuple with the Result field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetMsserverDhcpResponseObjectAsResult) GetResultOk() (*MsserverDhcp, bool) {
	if o == nil || IsNil(o.Result) {
		return nil, false
	}
	return o.Result, true
}

// HasResult returns a boolean if a field has been set.
func (o *GetMsserverDhcpResponseObjectAsResult) HasResult() bool {
	if o != nil && !IsNil(o.Result) {
		return true
	}

	return false
}

// SetResult gets a reference to the given MsserverDhcp and assigns it to the Result field.
func (o *GetMsserverDhcpResponseObjectAsResult) SetResult(v MsserverDhcp) {
	o.Result = &v
}

func (o GetMsserverDhcpResponseObjectAsResult) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetMsserverDhcpResponseObjectAsResult) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Result) {
		toSerialize["result"] = o.Result
	}
	return toSerialize, nil
}

type NullableGetMsserverDhcpResponseObjectAsResult struct {
	value *GetMsserverDhcpResponseObjectAsResult
	isSet bool
}

func (v NullableGetMsserverDhcpResponseObjectAsResult) Get() *GetMsserverDhcpResponseObjectAsResult {
	return v.value
}

func (v *NullableGetMsserverDhcpResponseObjectAsResult) Set(val *GetMsserverDhcpResponseObjectAsResult) {
	v.value = val
	v.isSet = true
}

func (v NullableGetMsserverDhcpResponseObjectAsResult) IsSet() bool {
	return v.isSet
}

func (v *NullableGetMsserverDhcpResponseObjectAsResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetMsserverDhcpResponseObjectAsResult(val *GetMsserverDhcpResponseObjectAsResult) *NullableGetMsserverDhcpResponseObjectAsResult {
	return &NullableGetMsserverDhcpResponseObjectAsResult{value: val, isSet: true}
}

func (v NullableGetMsserverDhcpResponseObjectAsResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetMsserverDhcpResponseObjectAsResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
