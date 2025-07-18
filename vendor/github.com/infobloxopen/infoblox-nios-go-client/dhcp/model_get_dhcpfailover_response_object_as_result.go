/*
Infoblox DHCP API

OpenAPI specification for Infoblox NIOS WAPI DHCP objects

API version: 2.13.6
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package dhcp

import (
	"encoding/json"
)

// checks if the GetDhcpfailoverResponseObjectAsResult type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetDhcpfailoverResponseObjectAsResult{}

// GetDhcpfailoverResponseObjectAsResult The response format to retrieve __Dhcpfailover__ objects.
type GetDhcpfailoverResponseObjectAsResult struct {
	Result *Dhcpfailover `json:"result,omitempty"`
}

// NewGetDhcpfailoverResponseObjectAsResult instantiates a new GetDhcpfailoverResponseObjectAsResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetDhcpfailoverResponseObjectAsResult() *GetDhcpfailoverResponseObjectAsResult {
	this := GetDhcpfailoverResponseObjectAsResult{}
	return &this
}

// NewGetDhcpfailoverResponseObjectAsResultWithDefaults instantiates a new GetDhcpfailoverResponseObjectAsResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetDhcpfailoverResponseObjectAsResultWithDefaults() *GetDhcpfailoverResponseObjectAsResult {
	this := GetDhcpfailoverResponseObjectAsResult{}
	return &this
}

// GetResult returns the Result field value if set, zero value otherwise.
func (o *GetDhcpfailoverResponseObjectAsResult) GetResult() Dhcpfailover {
	if o == nil || IsNil(o.Result) {
		var ret Dhcpfailover
		return ret
	}
	return *o.Result
}

// GetResultOk returns a tuple with the Result field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetDhcpfailoverResponseObjectAsResult) GetResultOk() (*Dhcpfailover, bool) {
	if o == nil || IsNil(o.Result) {
		return nil, false
	}
	return o.Result, true
}

// HasResult returns a boolean if a field has been set.
func (o *GetDhcpfailoverResponseObjectAsResult) HasResult() bool {
	if o != nil && !IsNil(o.Result) {
		return true
	}

	return false
}

// SetResult gets a reference to the given Dhcpfailover and assigns it to the Result field.
func (o *GetDhcpfailoverResponseObjectAsResult) SetResult(v Dhcpfailover) {
	o.Result = &v
}

func (o GetDhcpfailoverResponseObjectAsResult) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetDhcpfailoverResponseObjectAsResult) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Result) {
		toSerialize["result"] = o.Result
	}
	return toSerialize, nil
}

type NullableGetDhcpfailoverResponseObjectAsResult struct {
	value *GetDhcpfailoverResponseObjectAsResult
	isSet bool
}

func (v NullableGetDhcpfailoverResponseObjectAsResult) Get() *GetDhcpfailoverResponseObjectAsResult {
	return v.value
}

func (v *NullableGetDhcpfailoverResponseObjectAsResult) Set(val *GetDhcpfailoverResponseObjectAsResult) {
	v.value = val
	v.isSet = true
}

func (v NullableGetDhcpfailoverResponseObjectAsResult) IsSet() bool {
	return v.isSet
}

func (v *NullableGetDhcpfailoverResponseObjectAsResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetDhcpfailoverResponseObjectAsResult(val *GetDhcpfailoverResponseObjectAsResult) *NullableGetDhcpfailoverResponseObjectAsResult {
	return &NullableGetDhcpfailoverResponseObjectAsResult{value: val, isSet: true}
}

func (v NullableGetDhcpfailoverResponseObjectAsResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetDhcpfailoverResponseObjectAsResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
