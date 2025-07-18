/*
Infoblox DHCP API

OpenAPI specification for Infoblox NIOS WAPI DHCP objects

API version: 2.13.6
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package dhcp

import (
	"encoding/json"
	"fmt"
)

// GetDhcpfailoverResponse - struct for GetDhcpfailoverResponse
type GetDhcpfailoverResponse struct {
	Dhcpfailover                          *Dhcpfailover
	GetDhcpfailoverResponseObjectAsResult *GetDhcpfailoverResponseObjectAsResult
}

// DhcpfailoverAsGetDhcpfailoverResponse is a convenience function that returns Dhcpfailover wrapped in GetDhcpfailoverResponse
func DhcpfailoverAsGetDhcpfailoverResponse(v *Dhcpfailover) GetDhcpfailoverResponse {
	return GetDhcpfailoverResponse{
		Dhcpfailover: v,
	}
}

// GetDhcpfailoverResponseObjectAsResultAsGetDhcpfailoverResponse is a convenience function that returns GetDhcpfailoverResponseObjectAsResult wrapped in GetDhcpfailoverResponse
func GetDhcpfailoverResponseObjectAsResultAsGetDhcpfailoverResponse(v *GetDhcpfailoverResponseObjectAsResult) GetDhcpfailoverResponse {
	return GetDhcpfailoverResponse{
		GetDhcpfailoverResponseObjectAsResult: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *GetDhcpfailoverResponse) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into Dhcpfailover
	err = newStrictDecoder(data).Decode(&dst.Dhcpfailover)
	if err == nil {
		jsonDhcpfailover, _ := json.Marshal(dst.Dhcpfailover)
		if string(jsonDhcpfailover) == "{}" { // empty struct
			dst.Dhcpfailover = nil
		} else {
			match++
		}
	} else {
		dst.Dhcpfailover = nil
	}

	// try to unmarshal data into GetDhcpfailoverResponseObjectAsResult
	err = newStrictDecoder(data).Decode(&dst.GetDhcpfailoverResponseObjectAsResult)
	if err == nil {
		jsonGetDhcpfailoverResponseObjectAsResult, _ := json.Marshal(dst.GetDhcpfailoverResponseObjectAsResult)
		if string(jsonGetDhcpfailoverResponseObjectAsResult) == "{}" { // empty struct
			dst.GetDhcpfailoverResponseObjectAsResult = nil
		} else {
			match++
		}
	} else {
		dst.GetDhcpfailoverResponseObjectAsResult = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.Dhcpfailover = nil
		dst.GetDhcpfailoverResponseObjectAsResult = nil

		return fmt.Errorf("data matches more than one schema in oneOf(GetDhcpfailoverResponse)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(GetDhcpfailoverResponse)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src GetDhcpfailoverResponse) MarshalJSON() ([]byte, error) {
	if src.Dhcpfailover != nil {
		return json.Marshal(&src.Dhcpfailover)
	}

	if src.GetDhcpfailoverResponseObjectAsResult != nil {
		return json.Marshal(&src.GetDhcpfailoverResponseObjectAsResult)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *GetDhcpfailoverResponse) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.Dhcpfailover != nil {
		return obj.Dhcpfailover
	}

	if obj.GetDhcpfailoverResponseObjectAsResult != nil {
		return obj.GetDhcpfailoverResponseObjectAsResult
	}

	// all schemas are nil
	return nil
}

type NullableGetDhcpfailoverResponse struct {
	value *GetDhcpfailoverResponse
	isSet bool
}

func (v NullableGetDhcpfailoverResponse) Get() *GetDhcpfailoverResponse {
	return v.value
}

func (v *NullableGetDhcpfailoverResponse) Set(val *GetDhcpfailoverResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetDhcpfailoverResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetDhcpfailoverResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetDhcpfailoverResponse(val *GetDhcpfailoverResponse) *NullableGetDhcpfailoverResponse {
	return &NullableGetDhcpfailoverResponse{value: val, isSet: true}
}

func (v NullableGetDhcpfailoverResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetDhcpfailoverResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
