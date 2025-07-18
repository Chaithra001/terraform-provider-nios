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

// CreateSharedrecordAaaaResponse - struct for CreateSharedrecordAaaaResponse
type CreateSharedrecordAaaaResponse struct {
	CreateSharedrecordAaaaResponseAsObject *CreateSharedrecordAaaaResponseAsObject
	String                                 *string
}

// CreateSharedrecordAaaaResponseAsObjectAsCreateSharedrecordAaaaResponse is a convenience function that returns CreateSharedrecordAaaaResponseAsObject wrapped in CreateSharedrecordAaaaResponse
func CreateSharedrecordAaaaResponseAsObjectAsCreateSharedrecordAaaaResponse(v *CreateSharedrecordAaaaResponseAsObject) CreateSharedrecordAaaaResponse {
	return CreateSharedrecordAaaaResponse{
		CreateSharedrecordAaaaResponseAsObject: v,
	}
}

// stringAsCreateSharedrecordAaaaResponse is a convenience function that returns string wrapped in CreateSharedrecordAaaaResponse
func StringAsCreateSharedrecordAaaaResponse(v *string) CreateSharedrecordAaaaResponse {
	return CreateSharedrecordAaaaResponse{
		String: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *CreateSharedrecordAaaaResponse) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into CreateSharedrecordAaaaResponseAsObject
	err = newStrictDecoder(data).Decode(&dst.CreateSharedrecordAaaaResponseAsObject)
	if err == nil {
		jsonCreateSharedrecordAaaaResponseAsObject, _ := json.Marshal(dst.CreateSharedrecordAaaaResponseAsObject)
		if string(jsonCreateSharedrecordAaaaResponseAsObject) == "{}" { // empty struct
			dst.CreateSharedrecordAaaaResponseAsObject = nil
		} else {
			match++
		}
	} else {
		dst.CreateSharedrecordAaaaResponseAsObject = nil
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
		dst.CreateSharedrecordAaaaResponseAsObject = nil
		dst.String = nil

		return fmt.Errorf("data matches more than one schema in oneOf(CreateSharedrecordAaaaResponse)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(CreateSharedrecordAaaaResponse)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src CreateSharedrecordAaaaResponse) MarshalJSON() ([]byte, error) {
	if src.CreateSharedrecordAaaaResponseAsObject != nil {
		return json.Marshal(&src.CreateSharedrecordAaaaResponseAsObject)
	}

	if src.String != nil {
		return json.Marshal(&src.String)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *CreateSharedrecordAaaaResponse) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.CreateSharedrecordAaaaResponseAsObject != nil {
		return obj.CreateSharedrecordAaaaResponseAsObject
	}

	if obj.String != nil {
		return obj.String
	}

	// all schemas are nil
	return nil
}

type NullableCreateSharedrecordAaaaResponse struct {
	value *CreateSharedrecordAaaaResponse
	isSet bool
}

func (v NullableCreateSharedrecordAaaaResponse) Get() *CreateSharedrecordAaaaResponse {
	return v.value
}

func (v *NullableCreateSharedrecordAaaaResponse) Set(val *CreateSharedrecordAaaaResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateSharedrecordAaaaResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateSharedrecordAaaaResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateSharedrecordAaaaResponse(val *CreateSharedrecordAaaaResponse) *NullableCreateSharedrecordAaaaResponse {
	return &NullableCreateSharedrecordAaaaResponse{value: val, isSet: true}
}

func (v NullableCreateSharedrecordAaaaResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateSharedrecordAaaaResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
