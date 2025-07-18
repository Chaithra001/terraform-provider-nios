/*
Infoblox THREATINSIGHT API

OpenAPI specification for Infoblox NIOS WAPI THREATINSIGHT objects

API version: 2.13.6
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package threatinsight

import (
	"encoding/json"
	"fmt"
)

// GetThreatinsightInsightAllowlistResponse - struct for GetThreatinsightInsightAllowlistResponse
type GetThreatinsightInsightAllowlistResponse struct {
	GetThreatinsightInsightAllowlistResponseObjectAsResult *GetThreatinsightInsightAllowlistResponseObjectAsResult
	ThreatinsightInsightAllowlist                          *ThreatinsightInsightAllowlist
}

// GetThreatinsightInsightAllowlistResponseObjectAsResultAsGetThreatinsightInsightAllowlistResponse is a convenience function that returns GetThreatinsightInsightAllowlistResponseObjectAsResult wrapped in GetThreatinsightInsightAllowlistResponse
func GetThreatinsightInsightAllowlistResponseObjectAsResultAsGetThreatinsightInsightAllowlistResponse(v *GetThreatinsightInsightAllowlistResponseObjectAsResult) GetThreatinsightInsightAllowlistResponse {
	return GetThreatinsightInsightAllowlistResponse{
		GetThreatinsightInsightAllowlistResponseObjectAsResult: v,
	}
}

// ThreatinsightInsightAllowlistAsGetThreatinsightInsightAllowlistResponse is a convenience function that returns ThreatinsightInsightAllowlist wrapped in GetThreatinsightInsightAllowlistResponse
func ThreatinsightInsightAllowlistAsGetThreatinsightInsightAllowlistResponse(v *ThreatinsightInsightAllowlist) GetThreatinsightInsightAllowlistResponse {
	return GetThreatinsightInsightAllowlistResponse{
		ThreatinsightInsightAllowlist: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *GetThreatinsightInsightAllowlistResponse) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into GetThreatinsightInsightAllowlistResponseObjectAsResult
	err = newStrictDecoder(data).Decode(&dst.GetThreatinsightInsightAllowlistResponseObjectAsResult)
	if err == nil {
		jsonGetThreatinsightInsightAllowlistResponseObjectAsResult, _ := json.Marshal(dst.GetThreatinsightInsightAllowlistResponseObjectAsResult)
		if string(jsonGetThreatinsightInsightAllowlistResponseObjectAsResult) == "{}" { // empty struct
			dst.GetThreatinsightInsightAllowlistResponseObjectAsResult = nil
		} else {
			match++
		}
	} else {
		dst.GetThreatinsightInsightAllowlistResponseObjectAsResult = nil
	}

	// try to unmarshal data into ThreatinsightInsightAllowlist
	err = newStrictDecoder(data).Decode(&dst.ThreatinsightInsightAllowlist)
	if err == nil {
		jsonThreatinsightInsightAllowlist, _ := json.Marshal(dst.ThreatinsightInsightAllowlist)
		if string(jsonThreatinsightInsightAllowlist) == "{}" { // empty struct
			dst.ThreatinsightInsightAllowlist = nil
		} else {
			match++
		}
	} else {
		dst.ThreatinsightInsightAllowlist = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.GetThreatinsightInsightAllowlistResponseObjectAsResult = nil
		dst.ThreatinsightInsightAllowlist = nil

		return fmt.Errorf("data matches more than one schema in oneOf(GetThreatinsightInsightAllowlistResponse)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(GetThreatinsightInsightAllowlistResponse)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src GetThreatinsightInsightAllowlistResponse) MarshalJSON() ([]byte, error) {
	if src.GetThreatinsightInsightAllowlistResponseObjectAsResult != nil {
		return json.Marshal(&src.GetThreatinsightInsightAllowlistResponseObjectAsResult)
	}

	if src.ThreatinsightInsightAllowlist != nil {
		return json.Marshal(&src.ThreatinsightInsightAllowlist)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *GetThreatinsightInsightAllowlistResponse) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.GetThreatinsightInsightAllowlistResponseObjectAsResult != nil {
		return obj.GetThreatinsightInsightAllowlistResponseObjectAsResult
	}

	if obj.ThreatinsightInsightAllowlist != nil {
		return obj.ThreatinsightInsightAllowlist
	}

	// all schemas are nil
	return nil
}

type NullableGetThreatinsightInsightAllowlistResponse struct {
	value *GetThreatinsightInsightAllowlistResponse
	isSet bool
}

func (v NullableGetThreatinsightInsightAllowlistResponse) Get() *GetThreatinsightInsightAllowlistResponse {
	return v.value
}

func (v *NullableGetThreatinsightInsightAllowlistResponse) Set(val *GetThreatinsightInsightAllowlistResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetThreatinsightInsightAllowlistResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetThreatinsightInsightAllowlistResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetThreatinsightInsightAllowlistResponse(val *GetThreatinsightInsightAllowlistResponse) *NullableGetThreatinsightInsightAllowlistResponse {
	return &NullableGetThreatinsightInsightAllowlistResponse{value: val, isSet: true}
}

func (v NullableGetThreatinsightInsightAllowlistResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetThreatinsightInsightAllowlistResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
