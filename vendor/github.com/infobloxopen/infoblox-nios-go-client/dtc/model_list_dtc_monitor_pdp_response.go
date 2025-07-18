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

// ListDtcMonitorPdpResponse - struct for ListDtcMonitorPdpResponse
type ListDtcMonitorPdpResponse struct {
	ListDtcMonitorPdpResponseObject *ListDtcMonitorPdpResponseObject
	ArrayOfDtcMonitorPdp            *[]DtcMonitorPdp
}

// ListDtcMonitorPdpResponseObjectAsListDtcMonitorPdpResponse is a convenience function that returns ListDtcMonitorPdpResponseObject wrapped in ListDtcMonitorPdpResponse
func ListDtcMonitorPdpResponseObjectAsListDtcMonitorPdpResponse(v *ListDtcMonitorPdpResponseObject) ListDtcMonitorPdpResponse {
	return ListDtcMonitorPdpResponse{
		ListDtcMonitorPdpResponseObject: v,
	}
}

// []DtcMonitorPdpAsListDtcMonitorPdpResponse is a convenience function that returns []DtcMonitorPdp wrapped in ListDtcMonitorPdpResponse
func ArrayOfDtcMonitorPdpAsListDtcMonitorPdpResponse(v *[]DtcMonitorPdp) ListDtcMonitorPdpResponse {
	return ListDtcMonitorPdpResponse{
		ArrayOfDtcMonitorPdp: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *ListDtcMonitorPdpResponse) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into ListDtcMonitorPdpResponseObject
	err = newStrictDecoder(data).Decode(&dst.ListDtcMonitorPdpResponseObject)
	if err == nil {
		jsonListDtcMonitorPdpResponseObject, _ := json.Marshal(dst.ListDtcMonitorPdpResponseObject)
		if string(jsonListDtcMonitorPdpResponseObject) == "{}" { // empty struct
			dst.ListDtcMonitorPdpResponseObject = nil
		} else {
			match++
		}
	} else {
		dst.ListDtcMonitorPdpResponseObject = nil
	}

	// try to unmarshal data into ArrayOfDtcMonitorPdp
	err = newStrictDecoder(data).Decode(&dst.ArrayOfDtcMonitorPdp)
	if err == nil {
		jsonArrayOfDtcMonitorPdp, _ := json.Marshal(dst.ArrayOfDtcMonitorPdp)
		if string(jsonArrayOfDtcMonitorPdp) == "{}" { // empty struct
			dst.ArrayOfDtcMonitorPdp = nil
		} else {
			match++
		}
	} else {
		dst.ArrayOfDtcMonitorPdp = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.ListDtcMonitorPdpResponseObject = nil
		dst.ArrayOfDtcMonitorPdp = nil

		return fmt.Errorf("data matches more than one schema in oneOf(ListDtcMonitorPdpResponse)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(ListDtcMonitorPdpResponse)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src ListDtcMonitorPdpResponse) MarshalJSON() ([]byte, error) {
	if src.ListDtcMonitorPdpResponseObject != nil {
		return json.Marshal(&src.ListDtcMonitorPdpResponseObject)
	}

	if src.ArrayOfDtcMonitorPdp != nil {
		return json.Marshal(&src.ArrayOfDtcMonitorPdp)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *ListDtcMonitorPdpResponse) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.ListDtcMonitorPdpResponseObject != nil {
		return obj.ListDtcMonitorPdpResponseObject
	}

	if obj.ArrayOfDtcMonitorPdp != nil {
		return obj.ArrayOfDtcMonitorPdp
	}

	// all schemas are nil
	return nil
}

type NullableListDtcMonitorPdpResponse struct {
	value *ListDtcMonitorPdpResponse
	isSet bool
}

func (v NullableListDtcMonitorPdpResponse) Get() *ListDtcMonitorPdpResponse {
	return v.value
}

func (v *NullableListDtcMonitorPdpResponse) Set(val *ListDtcMonitorPdpResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableListDtcMonitorPdpResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableListDtcMonitorPdpResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListDtcMonitorPdpResponse(val *ListDtcMonitorPdpResponse) *NullableListDtcMonitorPdpResponse {
	return &NullableListDtcMonitorPdpResponse{value: val, isSet: true}
}

func (v NullableListDtcMonitorPdpResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListDtcMonitorPdpResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
