/*
Infoblox DNS API

OpenAPI specification for Infoblox NIOS WAPI DNS objects

API version: 2.13.6
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package dns

import (
	"encoding/json"
	"fmt"
)

// GetAllnsgroupResponse - struct for GetAllnsgroupResponse
type GetAllnsgroupResponse struct {
	Allnsgroup                          *Allnsgroup
	GetAllnsgroupResponseObjectAsResult *GetAllnsgroupResponseObjectAsResult
}

// AllnsgroupAsGetAllnsgroupResponse is a convenience function that returns Allnsgroup wrapped in GetAllnsgroupResponse
func AllnsgroupAsGetAllnsgroupResponse(v *Allnsgroup) GetAllnsgroupResponse {
	return GetAllnsgroupResponse{
		Allnsgroup: v,
	}
}

// GetAllnsgroupResponseObjectAsResultAsGetAllnsgroupResponse is a convenience function that returns GetAllnsgroupResponseObjectAsResult wrapped in GetAllnsgroupResponse
func GetAllnsgroupResponseObjectAsResultAsGetAllnsgroupResponse(v *GetAllnsgroupResponseObjectAsResult) GetAllnsgroupResponse {
	return GetAllnsgroupResponse{
		GetAllnsgroupResponseObjectAsResult: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *GetAllnsgroupResponse) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into Allnsgroup
	err = newStrictDecoder(data).Decode(&dst.Allnsgroup)
	if err == nil {
		jsonAllnsgroup, _ := json.Marshal(dst.Allnsgroup)
		if string(jsonAllnsgroup) == "{}" { // empty struct
			dst.Allnsgroup = nil
		} else {
			match++
		}
	} else {
		dst.Allnsgroup = nil
	}

	// try to unmarshal data into GetAllnsgroupResponseObjectAsResult
	err = newStrictDecoder(data).Decode(&dst.GetAllnsgroupResponseObjectAsResult)
	if err == nil {
		jsonGetAllnsgroupResponseObjectAsResult, _ := json.Marshal(dst.GetAllnsgroupResponseObjectAsResult)
		if string(jsonGetAllnsgroupResponseObjectAsResult) == "{}" { // empty struct
			dst.GetAllnsgroupResponseObjectAsResult = nil
		} else {
			match++
		}
	} else {
		dst.GetAllnsgroupResponseObjectAsResult = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.Allnsgroup = nil
		dst.GetAllnsgroupResponseObjectAsResult = nil

		return fmt.Errorf("data matches more than one schema in oneOf(GetAllnsgroupResponse)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(GetAllnsgroupResponse)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src GetAllnsgroupResponse) MarshalJSON() ([]byte, error) {
	if src.Allnsgroup != nil {
		return json.Marshal(&src.Allnsgroup)
	}

	if src.GetAllnsgroupResponseObjectAsResult != nil {
		return json.Marshal(&src.GetAllnsgroupResponseObjectAsResult)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *GetAllnsgroupResponse) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.Allnsgroup != nil {
		return obj.Allnsgroup
	}

	if obj.GetAllnsgroupResponseObjectAsResult != nil {
		return obj.GetAllnsgroupResponseObjectAsResult
	}

	// all schemas are nil
	return nil
}

type NullableGetAllnsgroupResponse struct {
	value *GetAllnsgroupResponse
	isSet bool
}

func (v NullableGetAllnsgroupResponse) Get() *GetAllnsgroupResponse {
	return v.value
}

func (v *NullableGetAllnsgroupResponse) Set(val *GetAllnsgroupResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetAllnsgroupResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetAllnsgroupResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetAllnsgroupResponse(val *GetAllnsgroupResponse) *NullableGetAllnsgroupResponse {
	return &NullableGetAllnsgroupResponse{value: val, isSet: true}
}

func (v NullableGetAllnsgroupResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetAllnsgroupResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
