/*
Infoblox THREATPROTECTION API

OpenAPI specification for Infoblox NIOS WAPI THREATPROTECTION objects

API version: 2.13.6
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package threatprotection

import (
	"encoding/json"
	"fmt"
)

// ListThreatprotectionRulesetResponse - struct for ListThreatprotectionRulesetResponse
type ListThreatprotectionRulesetResponse struct {
	ListThreatprotectionRulesetResponseObject *ListThreatprotectionRulesetResponseObject
	ArrayOfThreatprotectionRuleset            *[]ThreatprotectionRuleset
}

// ListThreatprotectionRulesetResponseObjectAsListThreatprotectionRulesetResponse is a convenience function that returns ListThreatprotectionRulesetResponseObject wrapped in ListThreatprotectionRulesetResponse
func ListThreatprotectionRulesetResponseObjectAsListThreatprotectionRulesetResponse(v *ListThreatprotectionRulesetResponseObject) ListThreatprotectionRulesetResponse {
	return ListThreatprotectionRulesetResponse{
		ListThreatprotectionRulesetResponseObject: v,
	}
}

// []ThreatprotectionRulesetAsListThreatprotectionRulesetResponse is a convenience function that returns []ThreatprotectionRuleset wrapped in ListThreatprotectionRulesetResponse
func ArrayOfThreatprotectionRulesetAsListThreatprotectionRulesetResponse(v *[]ThreatprotectionRuleset) ListThreatprotectionRulesetResponse {
	return ListThreatprotectionRulesetResponse{
		ArrayOfThreatprotectionRuleset: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *ListThreatprotectionRulesetResponse) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into ListThreatprotectionRulesetResponseObject
	err = newStrictDecoder(data).Decode(&dst.ListThreatprotectionRulesetResponseObject)
	if err == nil {
		jsonListThreatprotectionRulesetResponseObject, _ := json.Marshal(dst.ListThreatprotectionRulesetResponseObject)
		if string(jsonListThreatprotectionRulesetResponseObject) == "{}" { // empty struct
			dst.ListThreatprotectionRulesetResponseObject = nil
		} else {
			match++
		}
	} else {
		dst.ListThreatprotectionRulesetResponseObject = nil
	}

	// try to unmarshal data into ArrayOfThreatprotectionRuleset
	err = newStrictDecoder(data).Decode(&dst.ArrayOfThreatprotectionRuleset)
	if err == nil {
		jsonArrayOfThreatprotectionRuleset, _ := json.Marshal(dst.ArrayOfThreatprotectionRuleset)
		if string(jsonArrayOfThreatprotectionRuleset) == "{}" { // empty struct
			dst.ArrayOfThreatprotectionRuleset = nil
		} else {
			match++
		}
	} else {
		dst.ArrayOfThreatprotectionRuleset = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.ListThreatprotectionRulesetResponseObject = nil
		dst.ArrayOfThreatprotectionRuleset = nil

		return fmt.Errorf("data matches more than one schema in oneOf(ListThreatprotectionRulesetResponse)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(ListThreatprotectionRulesetResponse)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src ListThreatprotectionRulesetResponse) MarshalJSON() ([]byte, error) {
	if src.ListThreatprotectionRulesetResponseObject != nil {
		return json.Marshal(&src.ListThreatprotectionRulesetResponseObject)
	}

	if src.ArrayOfThreatprotectionRuleset != nil {
		return json.Marshal(&src.ArrayOfThreatprotectionRuleset)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *ListThreatprotectionRulesetResponse) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.ListThreatprotectionRulesetResponseObject != nil {
		return obj.ListThreatprotectionRulesetResponseObject
	}

	if obj.ArrayOfThreatprotectionRuleset != nil {
		return obj.ArrayOfThreatprotectionRuleset
	}

	// all schemas are nil
	return nil
}

type NullableListThreatprotectionRulesetResponse struct {
	value *ListThreatprotectionRulesetResponse
	isSet bool
}

func (v NullableListThreatprotectionRulesetResponse) Get() *ListThreatprotectionRulesetResponse {
	return v.value
}

func (v *NullableListThreatprotectionRulesetResponse) Set(val *ListThreatprotectionRulesetResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableListThreatprotectionRulesetResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableListThreatprotectionRulesetResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListThreatprotectionRulesetResponse(val *ListThreatprotectionRulesetResponse) *NullableListThreatprotectionRulesetResponse {
	return &NullableListThreatprotectionRulesetResponse{value: val, isSet: true}
}

func (v NullableListThreatprotectionRulesetResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListThreatprotectionRulesetResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
