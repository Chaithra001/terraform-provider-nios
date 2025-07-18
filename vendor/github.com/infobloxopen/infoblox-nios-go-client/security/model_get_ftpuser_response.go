/*
Infoblox SECURITY API

OpenAPI specification for Infoblox NIOS WAPI SECURITY objects

API version: 2.13.6
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package security

import (
	"encoding/json"
	"fmt"
)

// GetFtpuserResponse - struct for GetFtpuserResponse
type GetFtpuserResponse struct {
	Ftpuser                          *Ftpuser
	GetFtpuserResponseObjectAsResult *GetFtpuserResponseObjectAsResult
}

// FtpuserAsGetFtpuserResponse is a convenience function that returns Ftpuser wrapped in GetFtpuserResponse
func FtpuserAsGetFtpuserResponse(v *Ftpuser) GetFtpuserResponse {
	return GetFtpuserResponse{
		Ftpuser: v,
	}
}

// GetFtpuserResponseObjectAsResultAsGetFtpuserResponse is a convenience function that returns GetFtpuserResponseObjectAsResult wrapped in GetFtpuserResponse
func GetFtpuserResponseObjectAsResultAsGetFtpuserResponse(v *GetFtpuserResponseObjectAsResult) GetFtpuserResponse {
	return GetFtpuserResponse{
		GetFtpuserResponseObjectAsResult: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *GetFtpuserResponse) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into Ftpuser
	err = newStrictDecoder(data).Decode(&dst.Ftpuser)
	if err == nil {
		jsonFtpuser, _ := json.Marshal(dst.Ftpuser)
		if string(jsonFtpuser) == "{}" { // empty struct
			dst.Ftpuser = nil
		} else {
			match++
		}
	} else {
		dst.Ftpuser = nil
	}

	// try to unmarshal data into GetFtpuserResponseObjectAsResult
	err = newStrictDecoder(data).Decode(&dst.GetFtpuserResponseObjectAsResult)
	if err == nil {
		jsonGetFtpuserResponseObjectAsResult, _ := json.Marshal(dst.GetFtpuserResponseObjectAsResult)
		if string(jsonGetFtpuserResponseObjectAsResult) == "{}" { // empty struct
			dst.GetFtpuserResponseObjectAsResult = nil
		} else {
			match++
		}
	} else {
		dst.GetFtpuserResponseObjectAsResult = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.Ftpuser = nil
		dst.GetFtpuserResponseObjectAsResult = nil

		return fmt.Errorf("data matches more than one schema in oneOf(GetFtpuserResponse)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(GetFtpuserResponse)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src GetFtpuserResponse) MarshalJSON() ([]byte, error) {
	if src.Ftpuser != nil {
		return json.Marshal(&src.Ftpuser)
	}

	if src.GetFtpuserResponseObjectAsResult != nil {
		return json.Marshal(&src.GetFtpuserResponseObjectAsResult)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *GetFtpuserResponse) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.Ftpuser != nil {
		return obj.Ftpuser
	}

	if obj.GetFtpuserResponseObjectAsResult != nil {
		return obj.GetFtpuserResponseObjectAsResult
	}

	// all schemas are nil
	return nil
}

type NullableGetFtpuserResponse struct {
	value *GetFtpuserResponse
	isSet bool
}

func (v NullableGetFtpuserResponse) Get() *GetFtpuserResponse {
	return v.value
}

func (v *NullableGetFtpuserResponse) Set(val *GetFtpuserResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetFtpuserResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetFtpuserResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetFtpuserResponse(val *GetFtpuserResponse) *NullableGetFtpuserResponse {
	return &NullableGetFtpuserResponse{value: val, isSet: true}
}

func (v NullableGetFtpuserResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetFtpuserResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
