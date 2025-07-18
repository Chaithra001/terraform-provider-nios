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

// ListMemberThreatinsightResponse - struct for ListMemberThreatinsightResponse
type ListMemberThreatinsightResponse struct {
	ListMemberThreatinsightResponseObject *ListMemberThreatinsightResponseObject
	ArrayOfMemberThreatinsight            *[]MemberThreatinsight
}

// ListMemberThreatinsightResponseObjectAsListMemberThreatinsightResponse is a convenience function that returns ListMemberThreatinsightResponseObject wrapped in ListMemberThreatinsightResponse
func ListMemberThreatinsightResponseObjectAsListMemberThreatinsightResponse(v *ListMemberThreatinsightResponseObject) ListMemberThreatinsightResponse {
	return ListMemberThreatinsightResponse{
		ListMemberThreatinsightResponseObject: v,
	}
}

// []MemberThreatinsightAsListMemberThreatinsightResponse is a convenience function that returns []MemberThreatinsight wrapped in ListMemberThreatinsightResponse
func ArrayOfMemberThreatinsightAsListMemberThreatinsightResponse(v *[]MemberThreatinsight) ListMemberThreatinsightResponse {
	return ListMemberThreatinsightResponse{
		ArrayOfMemberThreatinsight: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *ListMemberThreatinsightResponse) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into ListMemberThreatinsightResponseObject
	err = newStrictDecoder(data).Decode(&dst.ListMemberThreatinsightResponseObject)
	if err == nil {
		jsonListMemberThreatinsightResponseObject, _ := json.Marshal(dst.ListMemberThreatinsightResponseObject)
		if string(jsonListMemberThreatinsightResponseObject) == "{}" { // empty struct
			dst.ListMemberThreatinsightResponseObject = nil
		} else {
			match++
		}
	} else {
		dst.ListMemberThreatinsightResponseObject = nil
	}

	// try to unmarshal data into ArrayOfMemberThreatinsight
	err = newStrictDecoder(data).Decode(&dst.ArrayOfMemberThreatinsight)
	if err == nil {
		jsonArrayOfMemberThreatinsight, _ := json.Marshal(dst.ArrayOfMemberThreatinsight)
		if string(jsonArrayOfMemberThreatinsight) == "{}" { // empty struct
			dst.ArrayOfMemberThreatinsight = nil
		} else {
			match++
		}
	} else {
		dst.ArrayOfMemberThreatinsight = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.ListMemberThreatinsightResponseObject = nil
		dst.ArrayOfMemberThreatinsight = nil

		return fmt.Errorf("data matches more than one schema in oneOf(ListMemberThreatinsightResponse)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(ListMemberThreatinsightResponse)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src ListMemberThreatinsightResponse) MarshalJSON() ([]byte, error) {
	if src.ListMemberThreatinsightResponseObject != nil {
		return json.Marshal(&src.ListMemberThreatinsightResponseObject)
	}

	if src.ArrayOfMemberThreatinsight != nil {
		return json.Marshal(&src.ArrayOfMemberThreatinsight)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *ListMemberThreatinsightResponse) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.ListMemberThreatinsightResponseObject != nil {
		return obj.ListMemberThreatinsightResponseObject
	}

	if obj.ArrayOfMemberThreatinsight != nil {
		return obj.ArrayOfMemberThreatinsight
	}

	// all schemas are nil
	return nil
}

type NullableListMemberThreatinsightResponse struct {
	value *ListMemberThreatinsightResponse
	isSet bool
}

func (v NullableListMemberThreatinsightResponse) Get() *ListMemberThreatinsightResponse {
	return v.value
}

func (v *NullableListMemberThreatinsightResponse) Set(val *ListMemberThreatinsightResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableListMemberThreatinsightResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableListMemberThreatinsightResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListMemberThreatinsightResponse(val *ListMemberThreatinsightResponse) *NullableListMemberThreatinsightResponse {
	return &NullableListMemberThreatinsightResponse{value: val, isSet: true}
}

func (v NullableListMemberThreatinsightResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListMemberThreatinsightResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
