/*
Infoblox IPAM API

OpenAPI specification for Infoblox NIOS WAPI IPAM objects

API version: 2.13.6
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package ipam

import (
	"encoding/json"
	"fmt"
)

// GetNetworkviewResponse - struct for GetNetworkviewResponse
type GetNetworkviewResponse struct {
	GetNetworkviewResponseObjectAsResult *GetNetworkviewResponseObjectAsResult
	Networkview                          *Networkview
}

// GetNetworkviewResponseObjectAsResultAsGetNetworkviewResponse is a convenience function that returns GetNetworkviewResponseObjectAsResult wrapped in GetNetworkviewResponse
func GetNetworkviewResponseObjectAsResultAsGetNetworkviewResponse(v *GetNetworkviewResponseObjectAsResult) GetNetworkviewResponse {
	return GetNetworkviewResponse{
		GetNetworkviewResponseObjectAsResult: v,
	}
}

// NetworkviewAsGetNetworkviewResponse is a convenience function that returns Networkview wrapped in GetNetworkviewResponse
func NetworkviewAsGetNetworkviewResponse(v *Networkview) GetNetworkviewResponse {
	return GetNetworkviewResponse{
		Networkview: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *GetNetworkviewResponse) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into GetNetworkviewResponseObjectAsResult
	err = newStrictDecoder(data).Decode(&dst.GetNetworkviewResponseObjectAsResult)
	if err == nil {
		jsonGetNetworkviewResponseObjectAsResult, _ := json.Marshal(dst.GetNetworkviewResponseObjectAsResult)
		if string(jsonGetNetworkviewResponseObjectAsResult) == "{}" { // empty struct
			dst.GetNetworkviewResponseObjectAsResult = nil
		} else {
			match++
		}
	} else {
		dst.GetNetworkviewResponseObjectAsResult = nil
	}

	// try to unmarshal data into Networkview
	err = newStrictDecoder(data).Decode(&dst.Networkview)
	if err == nil {
		jsonNetworkview, _ := json.Marshal(dst.Networkview)
		if string(jsonNetworkview) == "{}" { // empty struct
			dst.Networkview = nil
		} else {
			match++
		}
	} else {
		dst.Networkview = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.GetNetworkviewResponseObjectAsResult = nil
		dst.Networkview = nil

		return fmt.Errorf("data matches more than one schema in oneOf(GetNetworkviewResponse)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(GetNetworkviewResponse)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src GetNetworkviewResponse) MarshalJSON() ([]byte, error) {
	if src.GetNetworkviewResponseObjectAsResult != nil {
		return json.Marshal(&src.GetNetworkviewResponseObjectAsResult)
	}

	if src.Networkview != nil {
		return json.Marshal(&src.Networkview)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *GetNetworkviewResponse) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.GetNetworkviewResponseObjectAsResult != nil {
		return obj.GetNetworkviewResponseObjectAsResult
	}

	if obj.Networkview != nil {
		return obj.Networkview
	}

	// all schemas are nil
	return nil
}

type NullableGetNetworkviewResponse struct {
	value *GetNetworkviewResponse
	isSet bool
}

func (v NullableGetNetworkviewResponse) Get() *GetNetworkviewResponse {
	return v.value
}

func (v *NullableGetNetworkviewResponse) Set(val *GetNetworkviewResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetNetworkviewResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetNetworkviewResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetNetworkviewResponse(val *GetNetworkviewResponse) *NullableGetNetworkviewResponse {
	return &NullableGetNetworkviewResponse{value: val, isSet: true}
}

func (v NullableGetNetworkviewResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetNetworkviewResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
