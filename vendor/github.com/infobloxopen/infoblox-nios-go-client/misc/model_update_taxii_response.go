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

// UpdateTaxiiResponse - struct for UpdateTaxiiResponse
type UpdateTaxiiResponse struct {
	UpdateTaxiiResponseAsObject *UpdateTaxiiResponseAsObject
	String                      *string
}

// UpdateTaxiiResponseAsObjectAsUpdateTaxiiResponse is a convenience function that returns UpdateTaxiiResponseAsObject wrapped in UpdateTaxiiResponse
func UpdateTaxiiResponseAsObjectAsUpdateTaxiiResponse(v *UpdateTaxiiResponseAsObject) UpdateTaxiiResponse {
	return UpdateTaxiiResponse{
		UpdateTaxiiResponseAsObject: v,
	}
}

// stringAsUpdateTaxiiResponse is a convenience function that returns string wrapped in UpdateTaxiiResponse
func StringAsUpdateTaxiiResponse(v *string) UpdateTaxiiResponse {
	return UpdateTaxiiResponse{
		String: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *UpdateTaxiiResponse) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into UpdateTaxiiResponseAsObject
	err = newStrictDecoder(data).Decode(&dst.UpdateTaxiiResponseAsObject)
	if err == nil {
		jsonUpdateTaxiiResponseAsObject, _ := json.Marshal(dst.UpdateTaxiiResponseAsObject)
		if string(jsonUpdateTaxiiResponseAsObject) == "{}" { // empty struct
			dst.UpdateTaxiiResponseAsObject = nil
		} else {
			match++
		}
	} else {
		dst.UpdateTaxiiResponseAsObject = nil
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
		dst.UpdateTaxiiResponseAsObject = nil
		dst.String = nil

		return fmt.Errorf("data matches more than one schema in oneOf(UpdateTaxiiResponse)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(UpdateTaxiiResponse)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src UpdateTaxiiResponse) MarshalJSON() ([]byte, error) {
	if src.UpdateTaxiiResponseAsObject != nil {
		return json.Marshal(&src.UpdateTaxiiResponseAsObject)
	}

	if src.String != nil {
		return json.Marshal(&src.String)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *UpdateTaxiiResponse) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.UpdateTaxiiResponseAsObject != nil {
		return obj.UpdateTaxiiResponseAsObject
	}

	if obj.String != nil {
		return obj.String
	}

	// all schemas are nil
	return nil
}

type NullableUpdateTaxiiResponse struct {
	value *UpdateTaxiiResponse
	isSet bool
}

func (v NullableUpdateTaxiiResponse) Get() *UpdateTaxiiResponse {
	return v.value
}

func (v *NullableUpdateTaxiiResponse) Set(val *UpdateTaxiiResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateTaxiiResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateTaxiiResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateTaxiiResponse(val *UpdateTaxiiResponse) *NullableUpdateTaxiiResponse {
	return &NullableUpdateTaxiiResponse{value: val, isSet: true}
}

func (v NullableUpdateTaxiiResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateTaxiiResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
