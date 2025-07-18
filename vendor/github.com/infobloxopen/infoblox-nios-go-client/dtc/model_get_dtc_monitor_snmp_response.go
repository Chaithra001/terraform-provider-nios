/*
Infoblox DTC API

OpenAPI specification for Infoblox NIOS WAPI DTC objects

API version: 2.13.6
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package dtc

import (
	"encoding/json"
	"fmt"
)

// GetDtcMonitorSnmpResponse - struct for GetDtcMonitorSnmpResponse
type GetDtcMonitorSnmpResponse struct {
	DtcMonitorSnmp                          *DtcMonitorSnmp
	GetDtcMonitorSnmpResponseObjectAsResult *GetDtcMonitorSnmpResponseObjectAsResult
}

// DtcMonitorSnmpAsGetDtcMonitorSnmpResponse is a convenience function that returns DtcMonitorSnmp wrapped in GetDtcMonitorSnmpResponse
func DtcMonitorSnmpAsGetDtcMonitorSnmpResponse(v *DtcMonitorSnmp) GetDtcMonitorSnmpResponse {
	return GetDtcMonitorSnmpResponse{
		DtcMonitorSnmp: v,
	}
}

// GetDtcMonitorSnmpResponseObjectAsResultAsGetDtcMonitorSnmpResponse is a convenience function that returns GetDtcMonitorSnmpResponseObjectAsResult wrapped in GetDtcMonitorSnmpResponse
func GetDtcMonitorSnmpResponseObjectAsResultAsGetDtcMonitorSnmpResponse(v *GetDtcMonitorSnmpResponseObjectAsResult) GetDtcMonitorSnmpResponse {
	return GetDtcMonitorSnmpResponse{
		GetDtcMonitorSnmpResponseObjectAsResult: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *GetDtcMonitorSnmpResponse) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into DtcMonitorSnmp
	err = newStrictDecoder(data).Decode(&dst.DtcMonitorSnmp)
	if err == nil {
		jsonDtcMonitorSnmp, _ := json.Marshal(dst.DtcMonitorSnmp)
		if string(jsonDtcMonitorSnmp) == "{}" { // empty struct
			dst.DtcMonitorSnmp = nil
		} else {
			match++
		}
	} else {
		dst.DtcMonitorSnmp = nil
	}

	// try to unmarshal data into GetDtcMonitorSnmpResponseObjectAsResult
	err = newStrictDecoder(data).Decode(&dst.GetDtcMonitorSnmpResponseObjectAsResult)
	if err == nil {
		jsonGetDtcMonitorSnmpResponseObjectAsResult, _ := json.Marshal(dst.GetDtcMonitorSnmpResponseObjectAsResult)
		if string(jsonGetDtcMonitorSnmpResponseObjectAsResult) == "{}" { // empty struct
			dst.GetDtcMonitorSnmpResponseObjectAsResult = nil
		} else {
			match++
		}
	} else {
		dst.GetDtcMonitorSnmpResponseObjectAsResult = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.DtcMonitorSnmp = nil
		dst.GetDtcMonitorSnmpResponseObjectAsResult = nil

		return fmt.Errorf("data matches more than one schema in oneOf(GetDtcMonitorSnmpResponse)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(GetDtcMonitorSnmpResponse)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src GetDtcMonitorSnmpResponse) MarshalJSON() ([]byte, error) {
	if src.DtcMonitorSnmp != nil {
		return json.Marshal(&src.DtcMonitorSnmp)
	}

	if src.GetDtcMonitorSnmpResponseObjectAsResult != nil {
		return json.Marshal(&src.GetDtcMonitorSnmpResponseObjectAsResult)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *GetDtcMonitorSnmpResponse) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.DtcMonitorSnmp != nil {
		return obj.DtcMonitorSnmp
	}

	if obj.GetDtcMonitorSnmpResponseObjectAsResult != nil {
		return obj.GetDtcMonitorSnmpResponseObjectAsResult
	}

	// all schemas are nil
	return nil
}

type NullableGetDtcMonitorSnmpResponse struct {
	value *GetDtcMonitorSnmpResponse
	isSet bool
}

func (v NullableGetDtcMonitorSnmpResponse) Get() *GetDtcMonitorSnmpResponse {
	return v.value
}

func (v *NullableGetDtcMonitorSnmpResponse) Set(val *GetDtcMonitorSnmpResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetDtcMonitorSnmpResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetDtcMonitorSnmpResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetDtcMonitorSnmpResponse(val *GetDtcMonitorSnmpResponse) *NullableGetDtcMonitorSnmpResponse {
	return &NullableGetDtcMonitorSnmpResponse{value: val, isSet: true}
}

func (v NullableGetDtcMonitorSnmpResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetDtcMonitorSnmpResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
