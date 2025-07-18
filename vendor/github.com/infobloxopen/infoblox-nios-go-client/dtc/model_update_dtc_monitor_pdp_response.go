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

// UpdateDtcMonitorPdpResponse - struct for UpdateDtcMonitorPdpResponse
type UpdateDtcMonitorPdpResponse struct {
	UpdateDtcMonitorPdpResponseAsObject *UpdateDtcMonitorPdpResponseAsObject
	String                              *string
}

// UpdateDtcMonitorPdpResponseAsObjectAsUpdateDtcMonitorPdpResponse is a convenience function that returns UpdateDtcMonitorPdpResponseAsObject wrapped in UpdateDtcMonitorPdpResponse
func UpdateDtcMonitorPdpResponseAsObjectAsUpdateDtcMonitorPdpResponse(v *UpdateDtcMonitorPdpResponseAsObject) UpdateDtcMonitorPdpResponse {
	return UpdateDtcMonitorPdpResponse{
		UpdateDtcMonitorPdpResponseAsObject: v,
	}
}

// stringAsUpdateDtcMonitorPdpResponse is a convenience function that returns string wrapped in UpdateDtcMonitorPdpResponse
func StringAsUpdateDtcMonitorPdpResponse(v *string) UpdateDtcMonitorPdpResponse {
	return UpdateDtcMonitorPdpResponse{
		String: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *UpdateDtcMonitorPdpResponse) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into UpdateDtcMonitorPdpResponseAsObject
	err = newStrictDecoder(data).Decode(&dst.UpdateDtcMonitorPdpResponseAsObject)
	if err == nil {
		jsonUpdateDtcMonitorPdpResponseAsObject, _ := json.Marshal(dst.UpdateDtcMonitorPdpResponseAsObject)
		if string(jsonUpdateDtcMonitorPdpResponseAsObject) == "{}" { // empty struct
			dst.UpdateDtcMonitorPdpResponseAsObject = nil
		} else {
			match++
		}
	} else {
		dst.UpdateDtcMonitorPdpResponseAsObject = nil
	}

	// try to unmarshal data into String
	err = newStrictDecoder(data).Decode(&dst.String)
	if err == nil {
		jsonString, _ := json.Marshal(dst.String)
		if string(jsonString) == "{}" { // empty struct
			dst.String = nil
		} else {
			match++
		}
	} else {
		dst.String = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.UpdateDtcMonitorPdpResponseAsObject = nil
		dst.String = nil

		return fmt.Errorf("data matches more than one schema in oneOf(UpdateDtcMonitorPdpResponse)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(UpdateDtcMonitorPdpResponse)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src UpdateDtcMonitorPdpResponse) MarshalJSON() ([]byte, error) {
	if src.UpdateDtcMonitorPdpResponseAsObject != nil {
		return json.Marshal(&src.UpdateDtcMonitorPdpResponseAsObject)
	}

	if src.String != nil {
		return json.Marshal(&src.String)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *UpdateDtcMonitorPdpResponse) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.UpdateDtcMonitorPdpResponseAsObject != nil {
		return obj.UpdateDtcMonitorPdpResponseAsObject
	}

	if obj.String != nil {
		return obj.String
	}

	// all schemas are nil
	return nil
}

type NullableUpdateDtcMonitorPdpResponse struct {
	value *UpdateDtcMonitorPdpResponse
	isSet bool
}

func (v NullableUpdateDtcMonitorPdpResponse) Get() *UpdateDtcMonitorPdpResponse {
	return v.value
}

func (v *NullableUpdateDtcMonitorPdpResponse) Set(val *UpdateDtcMonitorPdpResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateDtcMonitorPdpResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateDtcMonitorPdpResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateDtcMonitorPdpResponse(val *UpdateDtcMonitorPdpResponse) *NullableUpdateDtcMonitorPdpResponse {
	return &NullableUpdateDtcMonitorPdpResponse{value: val, isSet: true}
}

func (v NullableUpdateDtcMonitorPdpResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateDtcMonitorPdpResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
