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

// ListViewResponse - struct for ListViewResponse
type ListViewResponse struct {
	ListViewResponseObject *ListViewResponseObject
	ArrayOfView            *[]View
}

// ListViewResponseObjectAsListViewResponse is a convenience function that returns ListViewResponseObject wrapped in ListViewResponse
func ListViewResponseObjectAsListViewResponse(v *ListViewResponseObject) ListViewResponse {
	return ListViewResponse{
		ListViewResponseObject: v,
	}
}

// []ViewAsListViewResponse is a convenience function that returns []View wrapped in ListViewResponse
func ArrayOfViewAsListViewResponse(v *[]View) ListViewResponse {
	return ListViewResponse{
		ArrayOfView: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *ListViewResponse) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into ListViewResponseObject
	err = newStrictDecoder(data).Decode(&dst.ListViewResponseObject)
	if err == nil {
		jsonListViewResponseObject, _ := json.Marshal(dst.ListViewResponseObject)
		if string(jsonListViewResponseObject) == "{}" { // empty struct
			dst.ListViewResponseObject = nil
		} else {
			match++
		}
	} else {
		dst.ListViewResponseObject = nil
	}

	// try to unmarshal data into ArrayOfView
	err = newStrictDecoder(data).Decode(&dst.ArrayOfView)
	if err == nil {
		jsonArrayOfView, _ := json.Marshal(dst.ArrayOfView)
		if string(jsonArrayOfView) == "{}" { // empty struct
			dst.ArrayOfView = nil
		} else {
			match++
		}
	} else {
		dst.ArrayOfView = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.ListViewResponseObject = nil
		dst.ArrayOfView = nil

		return fmt.Errorf("data matches more than one schema in oneOf(ListViewResponse)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(ListViewResponse)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src ListViewResponse) MarshalJSON() ([]byte, error) {
	if src.ListViewResponseObject != nil {
		return json.Marshal(&src.ListViewResponseObject)
	}

	if src.ArrayOfView != nil {
		return json.Marshal(&src.ArrayOfView)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *ListViewResponse) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.ListViewResponseObject != nil {
		return obj.ListViewResponseObject
	}

	if obj.ArrayOfView != nil {
		return obj.ArrayOfView
	}

	// all schemas are nil
	return nil
}

type NullableListViewResponse struct {
	value *ListViewResponse
	isSet bool
}

func (v NullableListViewResponse) Get() *ListViewResponse {
	return v.value
}

func (v *NullableListViewResponse) Set(val *ListViewResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableListViewResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableListViewResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListViewResponse(val *ListViewResponse) *NullableListViewResponse {
	return &NullableListViewResponse{value: val, isSet: true}
}

func (v NullableListViewResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListViewResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
