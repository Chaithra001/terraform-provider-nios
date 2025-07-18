/*
Infoblox IPAM API

OpenAPI specification for Infoblox NIOS WAPI IPAM objects

API version: 2.13.6
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package ipam

import (
	"encoding/json"
)

// checks if the GetNetworkDiscoveryResponseObjectAsResult type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetNetworkDiscoveryResponseObjectAsResult{}

// GetNetworkDiscoveryResponseObjectAsResult The response format to retrieve __NetworkDiscovery__ objects.
type GetNetworkDiscoveryResponseObjectAsResult struct {
	Result *NetworkDiscovery `json:"result,omitempty"`
}

// NewGetNetworkDiscoveryResponseObjectAsResult instantiates a new GetNetworkDiscoveryResponseObjectAsResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetNetworkDiscoveryResponseObjectAsResult() *GetNetworkDiscoveryResponseObjectAsResult {
	this := GetNetworkDiscoveryResponseObjectAsResult{}
	return &this
}

// NewGetNetworkDiscoveryResponseObjectAsResultWithDefaults instantiates a new GetNetworkDiscoveryResponseObjectAsResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetNetworkDiscoveryResponseObjectAsResultWithDefaults() *GetNetworkDiscoveryResponseObjectAsResult {
	this := GetNetworkDiscoveryResponseObjectAsResult{}
	return &this
}

// GetResult returns the Result field value if set, zero value otherwise.
func (o *GetNetworkDiscoveryResponseObjectAsResult) GetResult() NetworkDiscovery {
	if o == nil || IsNil(o.Result) {
		var ret NetworkDiscovery
		return ret
	}
	return *o.Result
}

// GetResultOk returns a tuple with the Result field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkDiscoveryResponseObjectAsResult) GetResultOk() (*NetworkDiscovery, bool) {
	if o == nil || IsNil(o.Result) {
		return nil, false
	}
	return o.Result, true
}

// HasResult returns a boolean if a field has been set.
func (o *GetNetworkDiscoveryResponseObjectAsResult) HasResult() bool {
	if o != nil && !IsNil(o.Result) {
		return true
	}

	return false
}

// SetResult gets a reference to the given NetworkDiscovery and assigns it to the Result field.
func (o *GetNetworkDiscoveryResponseObjectAsResult) SetResult(v NetworkDiscovery) {
	o.Result = &v
}

func (o GetNetworkDiscoveryResponseObjectAsResult) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetNetworkDiscoveryResponseObjectAsResult) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Result) {
		toSerialize["result"] = o.Result
	}
	return toSerialize, nil
}

type NullableGetNetworkDiscoveryResponseObjectAsResult struct {
	value *GetNetworkDiscoveryResponseObjectAsResult
	isSet bool
}

func (v NullableGetNetworkDiscoveryResponseObjectAsResult) Get() *GetNetworkDiscoveryResponseObjectAsResult {
	return v.value
}

func (v *NullableGetNetworkDiscoveryResponseObjectAsResult) Set(val *GetNetworkDiscoveryResponseObjectAsResult) {
	v.value = val
	v.isSet = true
}

func (v NullableGetNetworkDiscoveryResponseObjectAsResult) IsSet() bool {
	return v.isSet
}

func (v *NullableGetNetworkDiscoveryResponseObjectAsResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetNetworkDiscoveryResponseObjectAsResult(val *GetNetworkDiscoveryResponseObjectAsResult) *NullableGetNetworkDiscoveryResponseObjectAsResult {
	return &NullableGetNetworkDiscoveryResponseObjectAsResult{value: val, isSet: true}
}

func (v NullableGetNetworkDiscoveryResponseObjectAsResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetNetworkDiscoveryResponseObjectAsResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
