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

// ListGridMemberCloudapiResponse - struct for ListGridMemberCloudapiResponse
type ListGridMemberCloudapiResponse struct {
	ListGridMemberCloudapiResponseObject *ListGridMemberCloudapiResponseObject
	ArrayOfGridMemberCloudapi            *[]GridMemberCloudapi
}

// ListGridMemberCloudapiResponseObjectAsListGridMemberCloudapiResponse is a convenience function that returns ListGridMemberCloudapiResponseObject wrapped in ListGridMemberCloudapiResponse
func ListGridMemberCloudapiResponseObjectAsListGridMemberCloudapiResponse(v *ListGridMemberCloudapiResponseObject) ListGridMemberCloudapiResponse {
	return ListGridMemberCloudapiResponse{
		ListGridMemberCloudapiResponseObject: v,
	}
}

// []GridMemberCloudapiAsListGridMemberCloudapiResponse is a convenience function that returns []GridMemberCloudapi wrapped in ListGridMemberCloudapiResponse
func ArrayOfGridMemberCloudapiAsListGridMemberCloudapiResponse(v *[]GridMemberCloudapi) ListGridMemberCloudapiResponse {
	return ListGridMemberCloudapiResponse{
		ArrayOfGridMemberCloudapi: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *ListGridMemberCloudapiResponse) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into ListGridMemberCloudapiResponseObject
	err = newStrictDecoder(data).Decode(&dst.ListGridMemberCloudapiResponseObject)
	if err == nil {
		jsonListGridMemberCloudapiResponseObject, _ := json.Marshal(dst.ListGridMemberCloudapiResponseObject)
		if string(jsonListGridMemberCloudapiResponseObject) == "{}" { // empty struct
			dst.ListGridMemberCloudapiResponseObject = nil
		} else {
			match++
		}
	} else {
		dst.ListGridMemberCloudapiResponseObject = nil
	}

	// try to unmarshal data into ArrayOfGridMemberCloudapi
	err = newStrictDecoder(data).Decode(&dst.ArrayOfGridMemberCloudapi)
	if err == nil {
		jsonArrayOfGridMemberCloudapi, _ := json.Marshal(dst.ArrayOfGridMemberCloudapi)
		if string(jsonArrayOfGridMemberCloudapi) == "{}" { // empty struct
			dst.ArrayOfGridMemberCloudapi = nil
		} else {
			match++
		}
	} else {
		dst.ArrayOfGridMemberCloudapi = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.ListGridMemberCloudapiResponseObject = nil
		dst.ArrayOfGridMemberCloudapi = nil

		return fmt.Errorf("data matches more than one schema in oneOf(ListGridMemberCloudapiResponse)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(ListGridMemberCloudapiResponse)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src ListGridMemberCloudapiResponse) MarshalJSON() ([]byte, error) {
	if src.ListGridMemberCloudapiResponseObject != nil {
		return json.Marshal(&src.ListGridMemberCloudapiResponseObject)
	}

	if src.ArrayOfGridMemberCloudapi != nil {
		return json.Marshal(&src.ArrayOfGridMemberCloudapi)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *ListGridMemberCloudapiResponse) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.ListGridMemberCloudapiResponseObject != nil {
		return obj.ListGridMemberCloudapiResponseObject
	}

	if obj.ArrayOfGridMemberCloudapi != nil {
		return obj.ArrayOfGridMemberCloudapi
	}

	// all schemas are nil
	return nil
}

type NullableListGridMemberCloudapiResponse struct {
	value *ListGridMemberCloudapiResponse
	isSet bool
}

func (v NullableListGridMemberCloudapiResponse) Get() *ListGridMemberCloudapiResponse {
	return v.value
}

func (v *NullableListGridMemberCloudapiResponse) Set(val *ListGridMemberCloudapiResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableListGridMemberCloudapiResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableListGridMemberCloudapiResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListGridMemberCloudapiResponse(val *ListGridMemberCloudapiResponse) *NullableListGridMemberCloudapiResponse {
	return &NullableListGridMemberCloudapiResponse{value: val, isSet: true}
}

func (v NullableListGridMemberCloudapiResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListGridMemberCloudapiResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
