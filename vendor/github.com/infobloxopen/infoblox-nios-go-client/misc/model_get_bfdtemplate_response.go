/*
Infoblox MISC API

OpenAPI specification for Infoblox NIOS WAPI MISC objects

API version: 2.13.6
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package misc

import (
	"encoding/json"
	"fmt"
)

// GetBfdtemplateResponse - struct for GetBfdtemplateResponse
type GetBfdtemplateResponse struct {
	Bfdtemplate                          *Bfdtemplate
	GetBfdtemplateResponseObjectAsResult *GetBfdtemplateResponseObjectAsResult
}

// BfdtemplateAsGetBfdtemplateResponse is a convenience function that returns Bfdtemplate wrapped in GetBfdtemplateResponse
func BfdtemplateAsGetBfdtemplateResponse(v *Bfdtemplate) GetBfdtemplateResponse {
	return GetBfdtemplateResponse{
		Bfdtemplate: v,
	}
}

// GetBfdtemplateResponseObjectAsResultAsGetBfdtemplateResponse is a convenience function that returns GetBfdtemplateResponseObjectAsResult wrapped in GetBfdtemplateResponse
func GetBfdtemplateResponseObjectAsResultAsGetBfdtemplateResponse(v *GetBfdtemplateResponseObjectAsResult) GetBfdtemplateResponse {
	return GetBfdtemplateResponse{
		GetBfdtemplateResponseObjectAsResult: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *GetBfdtemplateResponse) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into Bfdtemplate
	err = newStrictDecoder(data).Decode(&dst.Bfdtemplate)
	if err == nil {
		jsonBfdtemplate, _ := json.Marshal(dst.Bfdtemplate)
		if string(jsonBfdtemplate) == "{}" { // empty struct
			dst.Bfdtemplate = nil
		} else {
			match++
		}
	} else {
		dst.Bfdtemplate = nil
	}

	// try to unmarshal data into GetBfdtemplateResponseObjectAsResult
	err = newStrictDecoder(data).Decode(&dst.GetBfdtemplateResponseObjectAsResult)
	if err == nil {
		jsonGetBfdtemplateResponseObjectAsResult, _ := json.Marshal(dst.GetBfdtemplateResponseObjectAsResult)
		if string(jsonGetBfdtemplateResponseObjectAsResult) == "{}" { // empty struct
			dst.GetBfdtemplateResponseObjectAsResult = nil
		} else {
			match++
		}
	} else {
		dst.GetBfdtemplateResponseObjectAsResult = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.Bfdtemplate = nil
		dst.GetBfdtemplateResponseObjectAsResult = nil

		return fmt.Errorf("data matches more than one schema in oneOf(GetBfdtemplateResponse)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(GetBfdtemplateResponse)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src GetBfdtemplateResponse) MarshalJSON() ([]byte, error) {
	if src.Bfdtemplate != nil {
		return json.Marshal(&src.Bfdtemplate)
	}

	if src.GetBfdtemplateResponseObjectAsResult != nil {
		return json.Marshal(&src.GetBfdtemplateResponseObjectAsResult)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *GetBfdtemplateResponse) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.Bfdtemplate != nil {
		return obj.Bfdtemplate
	}

	if obj.GetBfdtemplateResponseObjectAsResult != nil {
		return obj.GetBfdtemplateResponseObjectAsResult
	}

	// all schemas are nil
	return nil
}

type NullableGetBfdtemplateResponse struct {
	value *GetBfdtemplateResponse
	isSet bool
}

func (v NullableGetBfdtemplateResponse) Get() *GetBfdtemplateResponse {
	return v.value
}

func (v *NullableGetBfdtemplateResponse) Set(val *GetBfdtemplateResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetBfdtemplateResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetBfdtemplateResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetBfdtemplateResponse(val *GetBfdtemplateResponse) *NullableGetBfdtemplateResponse {
	return &NullableGetBfdtemplateResponse{value: val, isSet: true}
}

func (v NullableGetBfdtemplateResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetBfdtemplateResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
