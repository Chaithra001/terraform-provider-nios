/*
Infoblox IPAM API

OpenAPI specification for Infoblox NIOS WAPI IPAM objects

API version: 2.13.6
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package ipam

import (
	"encoding/json"
	"fmt"
)

// CreateSuperhostResponse - struct for CreateSuperhostResponse
type CreateSuperhostResponse struct {
	CreateSuperhostResponseAsObject *CreateSuperhostResponseAsObject
	String                          *string
}

// CreateSuperhostResponseAsObjectAsCreateSuperhostResponse is a convenience function that returns CreateSuperhostResponseAsObject wrapped in CreateSuperhostResponse
func CreateSuperhostResponseAsObjectAsCreateSuperhostResponse(v *CreateSuperhostResponseAsObject) CreateSuperhostResponse {
	return CreateSuperhostResponse{
		CreateSuperhostResponseAsObject: v,
	}
}

// stringAsCreateSuperhostResponse is a convenience function that returns string wrapped in CreateSuperhostResponse
func StringAsCreateSuperhostResponse(v *string) CreateSuperhostResponse {
	return CreateSuperhostResponse{
		String: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *CreateSuperhostResponse) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into CreateSuperhostResponseAsObject
	err = newStrictDecoder(data).Decode(&dst.CreateSuperhostResponseAsObject)
	if err == nil {
		jsonCreateSuperhostResponseAsObject, _ := json.Marshal(dst.CreateSuperhostResponseAsObject)
		if string(jsonCreateSuperhostResponseAsObject) == "{}" { // empty struct
			dst.CreateSuperhostResponseAsObject = nil
		} else {
			match++
		}
	} else {
		dst.CreateSuperhostResponseAsObject = nil
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
		dst.CreateSuperhostResponseAsObject = nil
		dst.String = nil

		return fmt.Errorf("data matches more than one schema in oneOf(CreateSuperhostResponse)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(CreateSuperhostResponse)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src CreateSuperhostResponse) MarshalJSON() ([]byte, error) {
	if src.CreateSuperhostResponseAsObject != nil {
		return json.Marshal(&src.CreateSuperhostResponseAsObject)
	}

	if src.String != nil {
		return json.Marshal(&src.String)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *CreateSuperhostResponse) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.CreateSuperhostResponseAsObject != nil {
		return obj.CreateSuperhostResponseAsObject
	}

	if obj.String != nil {
		return obj.String
	}

	// all schemas are nil
	return nil
}

type NullableCreateSuperhostResponse struct {
	value *CreateSuperhostResponse
	isSet bool
}

func (v NullableCreateSuperhostResponse) Get() *CreateSuperhostResponse {
	return v.value
}

func (v *NullableCreateSuperhostResponse) Set(val *CreateSuperhostResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateSuperhostResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateSuperhostResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateSuperhostResponse(val *CreateSuperhostResponse) *NullableCreateSuperhostResponse {
	return &NullableCreateSuperhostResponse{value: val, isSet: true}
}

func (v NullableCreateSuperhostResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateSuperhostResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
