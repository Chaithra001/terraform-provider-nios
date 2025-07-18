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

// GetDtcAllrecordsResponse - struct for GetDtcAllrecordsResponse
type GetDtcAllrecordsResponse struct {
	DtcAllrecords                          *DtcAllrecords
	GetDtcAllrecordsResponseObjectAsResult *GetDtcAllrecordsResponseObjectAsResult
}

// DtcAllrecordsAsGetDtcAllrecordsResponse is a convenience function that returns DtcAllrecords wrapped in GetDtcAllrecordsResponse
func DtcAllrecordsAsGetDtcAllrecordsResponse(v *DtcAllrecords) GetDtcAllrecordsResponse {
	return GetDtcAllrecordsResponse{
		DtcAllrecords: v,
	}
}

// GetDtcAllrecordsResponseObjectAsResultAsGetDtcAllrecordsResponse is a convenience function that returns GetDtcAllrecordsResponseObjectAsResult wrapped in GetDtcAllrecordsResponse
func GetDtcAllrecordsResponseObjectAsResultAsGetDtcAllrecordsResponse(v *GetDtcAllrecordsResponseObjectAsResult) GetDtcAllrecordsResponse {
	return GetDtcAllrecordsResponse{
		GetDtcAllrecordsResponseObjectAsResult: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *GetDtcAllrecordsResponse) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into DtcAllrecords
	err = newStrictDecoder(data).Decode(&dst.DtcAllrecords)
	if err == nil {
		jsonDtcAllrecords, _ := json.Marshal(dst.DtcAllrecords)
		if string(jsonDtcAllrecords) == "{}" { // empty struct
			dst.DtcAllrecords = nil
		} else {
			match++
		}
	} else {
		dst.DtcAllrecords = nil
	}

	// try to unmarshal data into GetDtcAllrecordsResponseObjectAsResult
	err = newStrictDecoder(data).Decode(&dst.GetDtcAllrecordsResponseObjectAsResult)
	if err == nil {
		jsonGetDtcAllrecordsResponseObjectAsResult, _ := json.Marshal(dst.GetDtcAllrecordsResponseObjectAsResult)
		if string(jsonGetDtcAllrecordsResponseObjectAsResult) == "{}" { // empty struct
			dst.GetDtcAllrecordsResponseObjectAsResult = nil
		} else {
			match++
		}
	} else {
		dst.GetDtcAllrecordsResponseObjectAsResult = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.DtcAllrecords = nil
		dst.GetDtcAllrecordsResponseObjectAsResult = nil

		return fmt.Errorf("data matches more than one schema in oneOf(GetDtcAllrecordsResponse)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(GetDtcAllrecordsResponse)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src GetDtcAllrecordsResponse) MarshalJSON() ([]byte, error) {
	if src.DtcAllrecords != nil {
		return json.Marshal(&src.DtcAllrecords)
	}

	if src.GetDtcAllrecordsResponseObjectAsResult != nil {
		return json.Marshal(&src.GetDtcAllrecordsResponseObjectAsResult)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *GetDtcAllrecordsResponse) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.DtcAllrecords != nil {
		return obj.DtcAllrecords
	}

	if obj.GetDtcAllrecordsResponseObjectAsResult != nil {
		return obj.GetDtcAllrecordsResponseObjectAsResult
	}

	// all schemas are nil
	return nil
}

type NullableGetDtcAllrecordsResponse struct {
	value *GetDtcAllrecordsResponse
	isSet bool
}

func (v NullableGetDtcAllrecordsResponse) Get() *GetDtcAllrecordsResponse {
	return v.value
}

func (v *NullableGetDtcAllrecordsResponse) Set(val *GetDtcAllrecordsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetDtcAllrecordsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetDtcAllrecordsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetDtcAllrecordsResponse(val *GetDtcAllrecordsResponse) *NullableGetDtcAllrecordsResponse {
	return &NullableGetDtcAllrecordsResponse{value: val, isSet: true}
}

func (v NullableGetDtcAllrecordsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetDtcAllrecordsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
