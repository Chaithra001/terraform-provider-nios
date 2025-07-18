/*
Infoblox DHCP API

OpenAPI specification for Infoblox NIOS WAPI DHCP objects

API version: 2.13.6
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package dhcp

import (
	"encoding/json"
	"fmt"
)

// ListDhcpoptionspaceResponse - struct for ListDhcpoptionspaceResponse
type ListDhcpoptionspaceResponse struct {
	ListDhcpoptionspaceResponseObject *ListDhcpoptionspaceResponseObject
	ArrayOfDhcpoptionspace            *[]Dhcpoptionspace
}

// ListDhcpoptionspaceResponseObjectAsListDhcpoptionspaceResponse is a convenience function that returns ListDhcpoptionspaceResponseObject wrapped in ListDhcpoptionspaceResponse
func ListDhcpoptionspaceResponseObjectAsListDhcpoptionspaceResponse(v *ListDhcpoptionspaceResponseObject) ListDhcpoptionspaceResponse {
	return ListDhcpoptionspaceResponse{
		ListDhcpoptionspaceResponseObject: v,
	}
}

// []DhcpoptionspaceAsListDhcpoptionspaceResponse is a convenience function that returns []Dhcpoptionspace wrapped in ListDhcpoptionspaceResponse
func ArrayOfDhcpoptionspaceAsListDhcpoptionspaceResponse(v *[]Dhcpoptionspace) ListDhcpoptionspaceResponse {
	return ListDhcpoptionspaceResponse{
		ArrayOfDhcpoptionspace: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *ListDhcpoptionspaceResponse) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into ListDhcpoptionspaceResponseObject
	err = newStrictDecoder(data).Decode(&dst.ListDhcpoptionspaceResponseObject)
	if err == nil {
		jsonListDhcpoptionspaceResponseObject, _ := json.Marshal(dst.ListDhcpoptionspaceResponseObject)
		if string(jsonListDhcpoptionspaceResponseObject) == "{}" { // empty struct
			dst.ListDhcpoptionspaceResponseObject = nil
		} else {
			match++
		}
	} else {
		dst.ListDhcpoptionspaceResponseObject = nil
	}

	// try to unmarshal data into ArrayOfDhcpoptionspace
	err = newStrictDecoder(data).Decode(&dst.ArrayOfDhcpoptionspace)
	if err == nil {
		jsonArrayOfDhcpoptionspace, _ := json.Marshal(dst.ArrayOfDhcpoptionspace)
		if string(jsonArrayOfDhcpoptionspace) == "{}" { // empty struct
			dst.ArrayOfDhcpoptionspace = nil
		} else {
			match++
		}
	} else {
		dst.ArrayOfDhcpoptionspace = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.ListDhcpoptionspaceResponseObject = nil
		dst.ArrayOfDhcpoptionspace = nil

		return fmt.Errorf("data matches more than one schema in oneOf(ListDhcpoptionspaceResponse)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(ListDhcpoptionspaceResponse)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src ListDhcpoptionspaceResponse) MarshalJSON() ([]byte, error) {
	if src.ListDhcpoptionspaceResponseObject != nil {
		return json.Marshal(&src.ListDhcpoptionspaceResponseObject)
	}

	if src.ArrayOfDhcpoptionspace != nil {
		return json.Marshal(&src.ArrayOfDhcpoptionspace)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *ListDhcpoptionspaceResponse) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.ListDhcpoptionspaceResponseObject != nil {
		return obj.ListDhcpoptionspaceResponseObject
	}

	if obj.ArrayOfDhcpoptionspace != nil {
		return obj.ArrayOfDhcpoptionspace
	}

	// all schemas are nil
	return nil
}

type NullableListDhcpoptionspaceResponse struct {
	value *ListDhcpoptionspaceResponse
	isSet bool
}

func (v NullableListDhcpoptionspaceResponse) Get() *ListDhcpoptionspaceResponse {
	return v.value
}

func (v *NullableListDhcpoptionspaceResponse) Set(val *ListDhcpoptionspaceResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableListDhcpoptionspaceResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableListDhcpoptionspaceResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListDhcpoptionspaceResponse(val *ListDhcpoptionspaceResponse) *NullableListDhcpoptionspaceResponse {
	return &NullableListDhcpoptionspaceResponse{value: val, isSet: true}
}

func (v NullableListDhcpoptionspaceResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListDhcpoptionspaceResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
