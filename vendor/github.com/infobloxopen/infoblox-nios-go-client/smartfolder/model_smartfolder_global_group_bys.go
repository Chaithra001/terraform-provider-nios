/*
Infoblox SMARTFOLDER API

OpenAPI specification for Infoblox NIOS WAPI SMARTFOLDER objects

API version: 2.13.6
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package smartfolder

import (
	"encoding/json"
)

// checks if the SmartfolderGlobalGroupBys type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SmartfolderGlobalGroupBys{}

// SmartfolderGlobalGroupBys struct for SmartfolderGlobalGroupBys
type SmartfolderGlobalGroupBys struct {
	// The name of the Smart Folder grouping attribute.
	Value *string `json:"value,omitempty"`
	// The type of the Smart Folder grouping attribute value.
	ValueType *string `json:"value_type,omitempty"`
	// Determines whether the grouping is enabled.
	EnableGrouping       *bool `json:"enable_grouping,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _SmartfolderGlobalGroupBys SmartfolderGlobalGroupBys

// NewSmartfolderGlobalGroupBys instantiates a new SmartfolderGlobalGroupBys object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSmartfolderGlobalGroupBys() *SmartfolderGlobalGroupBys {
	this := SmartfolderGlobalGroupBys{}
	return &this
}

// NewSmartfolderGlobalGroupBysWithDefaults instantiates a new SmartfolderGlobalGroupBys object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSmartfolderGlobalGroupBysWithDefaults() *SmartfolderGlobalGroupBys {
	this := SmartfolderGlobalGroupBys{}
	return &this
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *SmartfolderGlobalGroupBys) GetValue() string {
	if o == nil || IsNil(o.Value) {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SmartfolderGlobalGroupBys) GetValueOk() (*string, bool) {
	if o == nil || IsNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *SmartfolderGlobalGroupBys) HasValue() bool {
	if o != nil && !IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *SmartfolderGlobalGroupBys) SetValue(v string) {
	o.Value = &v
}

// GetValueType returns the ValueType field value if set, zero value otherwise.
func (o *SmartfolderGlobalGroupBys) GetValueType() string {
	if o == nil || IsNil(o.ValueType) {
		var ret string
		return ret
	}
	return *o.ValueType
}

// GetValueTypeOk returns a tuple with the ValueType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SmartfolderGlobalGroupBys) GetValueTypeOk() (*string, bool) {
	if o == nil || IsNil(o.ValueType) {
		return nil, false
	}
	return o.ValueType, true
}

// HasValueType returns a boolean if a field has been set.
func (o *SmartfolderGlobalGroupBys) HasValueType() bool {
	if o != nil && !IsNil(o.ValueType) {
		return true
	}

	return false
}

// SetValueType gets a reference to the given string and assigns it to the ValueType field.
func (o *SmartfolderGlobalGroupBys) SetValueType(v string) {
	o.ValueType = &v
}

// GetEnableGrouping returns the EnableGrouping field value if set, zero value otherwise.
func (o *SmartfolderGlobalGroupBys) GetEnableGrouping() bool {
	if o == nil || IsNil(o.EnableGrouping) {
		var ret bool
		return ret
	}
	return *o.EnableGrouping
}

// GetEnableGroupingOk returns a tuple with the EnableGrouping field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SmartfolderGlobalGroupBys) GetEnableGroupingOk() (*bool, bool) {
	if o == nil || IsNil(o.EnableGrouping) {
		return nil, false
	}
	return o.EnableGrouping, true
}

// HasEnableGrouping returns a boolean if a field has been set.
func (o *SmartfolderGlobalGroupBys) HasEnableGrouping() bool {
	if o != nil && !IsNil(o.EnableGrouping) {
		return true
	}

	return false
}

// SetEnableGrouping gets a reference to the given bool and assigns it to the EnableGrouping field.
func (o *SmartfolderGlobalGroupBys) SetEnableGrouping(v bool) {
	o.EnableGrouping = &v
}

func (o SmartfolderGlobalGroupBys) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SmartfolderGlobalGroupBys) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Value) {
		toSerialize["value"] = o.Value
	}
	if !IsNil(o.ValueType) {
		toSerialize["value_type"] = o.ValueType
	}
	if !IsNil(o.EnableGrouping) {
		toSerialize["enable_grouping"] = o.EnableGrouping
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *SmartfolderGlobalGroupBys) UnmarshalJSON(data []byte) (err error) {
	varSmartfolderGlobalGroupBys := _SmartfolderGlobalGroupBys{}

	err = json.Unmarshal(data, &varSmartfolderGlobalGroupBys)

	if err != nil {
		return err
	}

	*o = SmartfolderGlobalGroupBys(varSmartfolderGlobalGroupBys)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "value")
		delete(additionalProperties, "value_type")
		delete(additionalProperties, "enable_grouping")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableSmartfolderGlobalGroupBys struct {
	value *SmartfolderGlobalGroupBys
	isSet bool
}

func (v NullableSmartfolderGlobalGroupBys) Get() *SmartfolderGlobalGroupBys {
	return v.value
}

func (v *NullableSmartfolderGlobalGroupBys) Set(val *SmartfolderGlobalGroupBys) {
	v.value = val
	v.isSet = true
}

func (v NullableSmartfolderGlobalGroupBys) IsSet() bool {
	return v.isSet
}

func (v *NullableSmartfolderGlobalGroupBys) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSmartfolderGlobalGroupBys(val *SmartfolderGlobalGroupBys) *NullableSmartfolderGlobalGroupBys {
	return &NullableSmartfolderGlobalGroupBys{value: val, isSet: true}
}

func (v NullableSmartfolderGlobalGroupBys) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSmartfolderGlobalGroupBys) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
