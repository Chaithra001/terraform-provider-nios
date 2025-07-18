/*
Infoblox GRID API

OpenAPI specification for Infoblox NIOS WAPI GRID objects

API version: 2.13.6
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package grid

import (
	"encoding/json"
	"fmt"
)

// GetGridServicerestartRequestChangedobjectResponse - struct for GetGridServicerestartRequestChangedobjectResponse
type GetGridServicerestartRequestChangedobjectResponse struct {
	GetGridServicerestartRequestChangedobjectResponseObjectAsResult *GetGridServicerestartRequestChangedobjectResponseObjectAsResult
	GridServicerestartRequestChangedobject                          *GridServicerestartRequestChangedobject
}

// GetGridServicerestartRequestChangedobjectResponseObjectAsResultAsGetGridServicerestartRequestChangedobjectResponse is a convenience function that returns GetGridServicerestartRequestChangedobjectResponseObjectAsResult wrapped in GetGridServicerestartRequestChangedobjectResponse
func GetGridServicerestartRequestChangedobjectResponseObjectAsResultAsGetGridServicerestartRequestChangedobjectResponse(v *GetGridServicerestartRequestChangedobjectResponseObjectAsResult) GetGridServicerestartRequestChangedobjectResponse {
	return GetGridServicerestartRequestChangedobjectResponse{
		GetGridServicerestartRequestChangedobjectResponseObjectAsResult: v,
	}
}

// GridServicerestartRequestChangedobjectAsGetGridServicerestartRequestChangedobjectResponse is a convenience function that returns GridServicerestartRequestChangedobject wrapped in GetGridServicerestartRequestChangedobjectResponse
func GridServicerestartRequestChangedobjectAsGetGridServicerestartRequestChangedobjectResponse(v *GridServicerestartRequestChangedobject) GetGridServicerestartRequestChangedobjectResponse {
	return GetGridServicerestartRequestChangedobjectResponse{
		GridServicerestartRequestChangedobject: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *GetGridServicerestartRequestChangedobjectResponse) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into GetGridServicerestartRequestChangedobjectResponseObjectAsResult
	err = newStrictDecoder(data).Decode(&dst.GetGridServicerestartRequestChangedobjectResponseObjectAsResult)
	if err == nil {
		jsonGetGridServicerestartRequestChangedobjectResponseObjectAsResult, _ := json.Marshal(dst.GetGridServicerestartRequestChangedobjectResponseObjectAsResult)
		if string(jsonGetGridServicerestartRequestChangedobjectResponseObjectAsResult) == "{}" { // empty struct
			dst.GetGridServicerestartRequestChangedobjectResponseObjectAsResult = nil
		} else {
			match++
		}
	} else {
		dst.GetGridServicerestartRequestChangedobjectResponseObjectAsResult = nil
	}

	// try to unmarshal data into GridServicerestartRequestChangedobject
	err = newStrictDecoder(data).Decode(&dst.GridServicerestartRequestChangedobject)
	if err == nil {
		jsonGridServicerestartRequestChangedobject, _ := json.Marshal(dst.GridServicerestartRequestChangedobject)
		if string(jsonGridServicerestartRequestChangedobject) == "{}" { // empty struct
			dst.GridServicerestartRequestChangedobject = nil
		} else {
			match++
		}
	} else {
		dst.GridServicerestartRequestChangedobject = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.GetGridServicerestartRequestChangedobjectResponseObjectAsResult = nil
		dst.GridServicerestartRequestChangedobject = nil

		return fmt.Errorf("data matches more than one schema in oneOf(GetGridServicerestartRequestChangedobjectResponse)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(GetGridServicerestartRequestChangedobjectResponse)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src GetGridServicerestartRequestChangedobjectResponse) MarshalJSON() ([]byte, error) {
	if src.GetGridServicerestartRequestChangedobjectResponseObjectAsResult != nil {
		return json.Marshal(&src.GetGridServicerestartRequestChangedobjectResponseObjectAsResult)
	}

	if src.GridServicerestartRequestChangedobject != nil {
		return json.Marshal(&src.GridServicerestartRequestChangedobject)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *GetGridServicerestartRequestChangedobjectResponse) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.GetGridServicerestartRequestChangedobjectResponseObjectAsResult != nil {
		return obj.GetGridServicerestartRequestChangedobjectResponseObjectAsResult
	}

	if obj.GridServicerestartRequestChangedobject != nil {
		return obj.GridServicerestartRequestChangedobject
	}

	// all schemas are nil
	return nil
}

type NullableGetGridServicerestartRequestChangedobjectResponse struct {
	value *GetGridServicerestartRequestChangedobjectResponse
	isSet bool
}

func (v NullableGetGridServicerestartRequestChangedobjectResponse) Get() *GetGridServicerestartRequestChangedobjectResponse {
	return v.value
}

func (v *NullableGetGridServicerestartRequestChangedobjectResponse) Set(val *GetGridServicerestartRequestChangedobjectResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetGridServicerestartRequestChangedobjectResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetGridServicerestartRequestChangedobjectResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetGridServicerestartRequestChangedobjectResponse(val *GetGridServicerestartRequestChangedobjectResponse) *NullableGetGridServicerestartRequestChangedobjectResponse {
	return &NullableGetGridServicerestartRequestChangedobjectResponse{value: val, isSet: true}
}

func (v NullableGetGridServicerestartRequestChangedobjectResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetGridServicerestartRequestChangedobjectResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
