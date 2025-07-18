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

// ListSearchResponse - struct for ListSearchResponse
type ListSearchResponse struct {
	ListSearchResponseObject *ListSearchResponseObject
	ArrayOfSearch            *[]Search
}

// ListSearchResponseObjectAsListSearchResponse is a convenience function that returns ListSearchResponseObject wrapped in ListSearchResponse
func ListSearchResponseObjectAsListSearchResponse(v *ListSearchResponseObject) ListSearchResponse {
	return ListSearchResponse{
		ListSearchResponseObject: v,
	}
}

// []SearchAsListSearchResponse is a convenience function that returns []Search wrapped in ListSearchResponse
func ArrayOfSearchAsListSearchResponse(v *[]Search) ListSearchResponse {
	return ListSearchResponse{
		ArrayOfSearch: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *ListSearchResponse) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into ListSearchResponseObject
	err = newStrictDecoder(data).Decode(&dst.ListSearchResponseObject)
	if err == nil {
		jsonListSearchResponseObject, _ := json.Marshal(dst.ListSearchResponseObject)
		if string(jsonListSearchResponseObject) == "{}" { // empty struct
			dst.ListSearchResponseObject = nil
		} else {
			match++
		}
	} else {
		dst.ListSearchResponseObject = nil
	}

	// try to unmarshal data into ArrayOfSearch
	err = newStrictDecoder(data).Decode(&dst.ArrayOfSearch)
	if err == nil {
		jsonArrayOfSearch, _ := json.Marshal(dst.ArrayOfSearch)
		if string(jsonArrayOfSearch) == "{}" { // empty struct
			dst.ArrayOfSearch = nil
		} else {
			match++
		}
	} else {
		dst.ArrayOfSearch = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.ListSearchResponseObject = nil
		dst.ArrayOfSearch = nil

		return fmt.Errorf("data matches more than one schema in oneOf(ListSearchResponse)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(ListSearchResponse)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src ListSearchResponse) MarshalJSON() ([]byte, error) {
	if src.ListSearchResponseObject != nil {
		return json.Marshal(&src.ListSearchResponseObject)
	}

	if src.ArrayOfSearch != nil {
		return json.Marshal(&src.ArrayOfSearch)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *ListSearchResponse) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.ListSearchResponseObject != nil {
		return obj.ListSearchResponseObject
	}

	if obj.ArrayOfSearch != nil {
		return obj.ArrayOfSearch
	}

	// all schemas are nil
	return nil
}

type NullableListSearchResponse struct {
	value *ListSearchResponse
	isSet bool
}

func (v NullableListSearchResponse) Get() *ListSearchResponse {
	return v.value
}

func (v *NullableListSearchResponse) Set(val *ListSearchResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableListSearchResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableListSearchResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListSearchResponse(val *ListSearchResponse) *NullableListSearchResponse {
	return &NullableListSearchResponse{value: val, isSet: true}
}

func (v NullableListSearchResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListSearchResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
