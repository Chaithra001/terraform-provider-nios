/*
Infoblox ACL API

OpenAPI specification for Infoblox NIOS WAPI ACL objects

API version: 2.13.6
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package acl

import (
	"encoding/json"
	"fmt"
)

// CreateNamedaclResponse - struct for CreateNamedaclResponse
type CreateNamedaclResponse struct {
	CreateNamedaclResponseAsObject *CreateNamedaclResponseAsObject
	String                         *string
}

// CreateNamedaclResponseAsObjectAsCreateNamedaclResponse is a convenience function that returns CreateNamedaclResponseAsObject wrapped in CreateNamedaclResponse
func CreateNamedaclResponseAsObjectAsCreateNamedaclResponse(v *CreateNamedaclResponseAsObject) CreateNamedaclResponse {
	return CreateNamedaclResponse{
		CreateNamedaclResponseAsObject: v,
	}
}

// stringAsCreateNamedaclResponse is a convenience function that returns string wrapped in CreateNamedaclResponse
func StringAsCreateNamedaclResponse(v *string) CreateNamedaclResponse {
	return CreateNamedaclResponse{
		String: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *CreateNamedaclResponse) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into CreateNamedaclResponseAsObject
	err = newStrictDecoder(data).Decode(&dst.CreateNamedaclResponseAsObject)
	if err == nil {
		jsonCreateNamedaclResponseAsObject, _ := json.Marshal(dst.CreateNamedaclResponseAsObject)
		if string(jsonCreateNamedaclResponseAsObject) == "{}" { // empty struct
			dst.CreateNamedaclResponseAsObject = nil
		} else {
			match++
		}
	} else {
		dst.CreateNamedaclResponseAsObject = nil
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
		dst.CreateNamedaclResponseAsObject = nil
		dst.String = nil

		return fmt.Errorf("data matches more than one schema in oneOf(CreateNamedaclResponse)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(CreateNamedaclResponse)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src CreateNamedaclResponse) MarshalJSON() ([]byte, error) {
	if src.CreateNamedaclResponseAsObject != nil {
		return json.Marshal(&src.CreateNamedaclResponseAsObject)
	}

	if src.String != nil {
		return json.Marshal(&src.String)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *CreateNamedaclResponse) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.CreateNamedaclResponseAsObject != nil {
		return obj.CreateNamedaclResponseAsObject
	}

	if obj.String != nil {
		return obj.String
	}

	// all schemas are nil
	return nil
}

type NullableCreateNamedaclResponse struct {
	value *CreateNamedaclResponse
	isSet bool
}

func (v NullableCreateNamedaclResponse) Get() *CreateNamedaclResponse {
	return v.value
}

func (v *NullableCreateNamedaclResponse) Set(val *CreateNamedaclResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateNamedaclResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateNamedaclResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateNamedaclResponse(val *CreateNamedaclResponse) *NullableCreateNamedaclResponse {
	return &NullableCreateNamedaclResponse{value: val, isSet: true}
}

func (v NullableCreateNamedaclResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateNamedaclResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
