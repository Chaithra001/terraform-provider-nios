/*
Infoblox CLOUD API

OpenAPI specification for Infoblox NIOS WAPI CLOUD objects

API version: 2.13.6
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package cloud

import (
	"encoding/json"
	"fmt"
)

// ListAwsrte53taskgroupResponse - struct for ListAwsrte53taskgroupResponse
type ListAwsrte53taskgroupResponse struct {
	ListAwsrte53taskgroupResponseObject *ListAwsrte53taskgroupResponseObject
	ArrayOfAwsrte53taskgroup            *[]Awsrte53taskgroup
}

// ListAwsrte53taskgroupResponseObjectAsListAwsrte53taskgroupResponse is a convenience function that returns ListAwsrte53taskgroupResponseObject wrapped in ListAwsrte53taskgroupResponse
func ListAwsrte53taskgroupResponseObjectAsListAwsrte53taskgroupResponse(v *ListAwsrte53taskgroupResponseObject) ListAwsrte53taskgroupResponse {
	return ListAwsrte53taskgroupResponse{
		ListAwsrte53taskgroupResponseObject: v,
	}
}

// []Awsrte53taskgroupAsListAwsrte53taskgroupResponse is a convenience function that returns []Awsrte53taskgroup wrapped in ListAwsrte53taskgroupResponse
func ArrayOfAwsrte53taskgroupAsListAwsrte53taskgroupResponse(v *[]Awsrte53taskgroup) ListAwsrte53taskgroupResponse {
	return ListAwsrte53taskgroupResponse{
		ArrayOfAwsrte53taskgroup: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *ListAwsrte53taskgroupResponse) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into ListAwsrte53taskgroupResponseObject
	err = newStrictDecoder(data).Decode(&dst.ListAwsrte53taskgroupResponseObject)
	if err == nil {
		jsonListAwsrte53taskgroupResponseObject, _ := json.Marshal(dst.ListAwsrte53taskgroupResponseObject)
		if string(jsonListAwsrte53taskgroupResponseObject) == "{}" { // empty struct
			dst.ListAwsrte53taskgroupResponseObject = nil
		} else {
			match++
		}
	} else {
		dst.ListAwsrte53taskgroupResponseObject = nil
	}

	// try to unmarshal data into ArrayOfAwsrte53taskgroup
	err = newStrictDecoder(data).Decode(&dst.ArrayOfAwsrte53taskgroup)
	if err == nil {
		jsonArrayOfAwsrte53taskgroup, _ := json.Marshal(dst.ArrayOfAwsrte53taskgroup)
		if string(jsonArrayOfAwsrte53taskgroup) == "{}" { // empty struct
			dst.ArrayOfAwsrte53taskgroup = nil
		} else {
			match++
		}
	} else {
		dst.ArrayOfAwsrte53taskgroup = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.ListAwsrte53taskgroupResponseObject = nil
		dst.ArrayOfAwsrte53taskgroup = nil

		return fmt.Errorf("data matches more than one schema in oneOf(ListAwsrte53taskgroupResponse)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(ListAwsrte53taskgroupResponse)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src ListAwsrte53taskgroupResponse) MarshalJSON() ([]byte, error) {
	if src.ListAwsrte53taskgroupResponseObject != nil {
		return json.Marshal(&src.ListAwsrte53taskgroupResponseObject)
	}

	if src.ArrayOfAwsrte53taskgroup != nil {
		return json.Marshal(&src.ArrayOfAwsrte53taskgroup)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *ListAwsrte53taskgroupResponse) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.ListAwsrte53taskgroupResponseObject != nil {
		return obj.ListAwsrte53taskgroupResponseObject
	}

	if obj.ArrayOfAwsrte53taskgroup != nil {
		return obj.ArrayOfAwsrte53taskgroup
	}

	// all schemas are nil
	return nil
}

type NullableListAwsrte53taskgroupResponse struct {
	value *ListAwsrte53taskgroupResponse
	isSet bool
}

func (v NullableListAwsrte53taskgroupResponse) Get() *ListAwsrte53taskgroupResponse {
	return v.value
}

func (v *NullableListAwsrte53taskgroupResponse) Set(val *ListAwsrte53taskgroupResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableListAwsrte53taskgroupResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableListAwsrte53taskgroupResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListAwsrte53taskgroupResponse(val *ListAwsrte53taskgroupResponse) *NullableListAwsrte53taskgroupResponse {
	return &NullableListAwsrte53taskgroupResponse{value: val, isSet: true}
}

func (v NullableListAwsrte53taskgroupResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListAwsrte53taskgroupResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
