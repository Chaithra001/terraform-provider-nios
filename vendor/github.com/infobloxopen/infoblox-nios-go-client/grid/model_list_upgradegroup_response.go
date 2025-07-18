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

// ListUpgradegroupResponse - struct for ListUpgradegroupResponse
type ListUpgradegroupResponse struct {
	ListUpgradegroupResponseObject *ListUpgradegroupResponseObject
	ArrayOfUpgradegroup            *[]Upgradegroup
}

// ListUpgradegroupResponseObjectAsListUpgradegroupResponse is a convenience function that returns ListUpgradegroupResponseObject wrapped in ListUpgradegroupResponse
func ListUpgradegroupResponseObjectAsListUpgradegroupResponse(v *ListUpgradegroupResponseObject) ListUpgradegroupResponse {
	return ListUpgradegroupResponse{
		ListUpgradegroupResponseObject: v,
	}
}

// []UpgradegroupAsListUpgradegroupResponse is a convenience function that returns []Upgradegroup wrapped in ListUpgradegroupResponse
func ArrayOfUpgradegroupAsListUpgradegroupResponse(v *[]Upgradegroup) ListUpgradegroupResponse {
	return ListUpgradegroupResponse{
		ArrayOfUpgradegroup: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *ListUpgradegroupResponse) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into ListUpgradegroupResponseObject
	err = newStrictDecoder(data).Decode(&dst.ListUpgradegroupResponseObject)
	if err == nil {
		jsonListUpgradegroupResponseObject, _ := json.Marshal(dst.ListUpgradegroupResponseObject)
		if string(jsonListUpgradegroupResponseObject) == "{}" { // empty struct
			dst.ListUpgradegroupResponseObject = nil
		} else {
			match++
		}
	} else {
		dst.ListUpgradegroupResponseObject = nil
	}

	// try to unmarshal data into ArrayOfUpgradegroup
	err = newStrictDecoder(data).Decode(&dst.ArrayOfUpgradegroup)
	if err == nil {
		jsonArrayOfUpgradegroup, _ := json.Marshal(dst.ArrayOfUpgradegroup)
		if string(jsonArrayOfUpgradegroup) == "{}" { // empty struct
			dst.ArrayOfUpgradegroup = nil
		} else {
			match++
		}
	} else {
		dst.ArrayOfUpgradegroup = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.ListUpgradegroupResponseObject = nil
		dst.ArrayOfUpgradegroup = nil

		return fmt.Errorf("data matches more than one schema in oneOf(ListUpgradegroupResponse)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(ListUpgradegroupResponse)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src ListUpgradegroupResponse) MarshalJSON() ([]byte, error) {
	if src.ListUpgradegroupResponseObject != nil {
		return json.Marshal(&src.ListUpgradegroupResponseObject)
	}

	if src.ArrayOfUpgradegroup != nil {
		return json.Marshal(&src.ArrayOfUpgradegroup)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *ListUpgradegroupResponse) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.ListUpgradegroupResponseObject != nil {
		return obj.ListUpgradegroupResponseObject
	}

	if obj.ArrayOfUpgradegroup != nil {
		return obj.ArrayOfUpgradegroup
	}

	// all schemas are nil
	return nil
}

type NullableListUpgradegroupResponse struct {
	value *ListUpgradegroupResponse
	isSet bool
}

func (v NullableListUpgradegroupResponse) Get() *ListUpgradegroupResponse {
	return v.value
}

func (v *NullableListUpgradegroupResponse) Set(val *ListUpgradegroupResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableListUpgradegroupResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableListUpgradegroupResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListUpgradegroupResponse(val *ListUpgradegroupResponse) *NullableListUpgradegroupResponse {
	return &NullableListUpgradegroupResponse{value: val, isSet: true}
}

func (v NullableListUpgradegroupResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListUpgradegroupResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
