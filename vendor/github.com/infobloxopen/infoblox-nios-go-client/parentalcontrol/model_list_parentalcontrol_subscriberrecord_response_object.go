/*
Infoblox PARENTALCONTROL API

OpenAPI specification for Infoblox NIOS WAPI PARENTALCONTROL objects

API version: 2.13.6
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package parentalcontrol

import (
	"encoding/json"
)

// checks if the ListParentalcontrolSubscriberrecordResponseObject type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListParentalcontrolSubscriberrecordResponseObject{}

// ListParentalcontrolSubscriberrecordResponseObject The response format to retrieve __ParentalcontrolSubscriberrecord__ objects.
type ListParentalcontrolSubscriberrecordResponseObject struct {
	Result               []ParentalcontrolSubscriberrecord `json:"result,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ListParentalcontrolSubscriberrecordResponseObject ListParentalcontrolSubscriberrecordResponseObject

// NewListParentalcontrolSubscriberrecordResponseObject instantiates a new ListParentalcontrolSubscriberrecordResponseObject object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListParentalcontrolSubscriberrecordResponseObject() *ListParentalcontrolSubscriberrecordResponseObject {
	this := ListParentalcontrolSubscriberrecordResponseObject{}
	return &this
}

// NewListParentalcontrolSubscriberrecordResponseObjectWithDefaults instantiates a new ListParentalcontrolSubscriberrecordResponseObject object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListParentalcontrolSubscriberrecordResponseObjectWithDefaults() *ListParentalcontrolSubscriberrecordResponseObject {
	this := ListParentalcontrolSubscriberrecordResponseObject{}
	return &this
}

// GetResult returns the Result field value if set, zero value otherwise.
func (o *ListParentalcontrolSubscriberrecordResponseObject) GetResult() []ParentalcontrolSubscriberrecord {
	if o == nil || IsNil(o.Result) {
		var ret []ParentalcontrolSubscriberrecord
		return ret
	}
	return o.Result
}

// GetResultOk returns a tuple with the Result field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListParentalcontrolSubscriberrecordResponseObject) GetResultOk() ([]ParentalcontrolSubscriberrecord, bool) {
	if o == nil || IsNil(o.Result) {
		return nil, false
	}
	return o.Result, true
}

// HasResult returns a boolean if a field has been set.
func (o *ListParentalcontrolSubscriberrecordResponseObject) HasResult() bool {
	if o != nil && !IsNil(o.Result) {
		return true
	}

	return false
}

// SetResult gets a reference to the given []ParentalcontrolSubscriberrecord and assigns it to the Result field.
func (o *ListParentalcontrolSubscriberrecordResponseObject) SetResult(v []ParentalcontrolSubscriberrecord) {
	o.Result = v
}

func (o ListParentalcontrolSubscriberrecordResponseObject) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListParentalcontrolSubscriberrecordResponseObject) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Result) {
		toSerialize["result"] = o.Result
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ListParentalcontrolSubscriberrecordResponseObject) UnmarshalJSON(data []byte) (err error) {
	varListParentalcontrolSubscriberrecordResponseObject := _ListParentalcontrolSubscriberrecordResponseObject{}

	err = json.Unmarshal(data, &varListParentalcontrolSubscriberrecordResponseObject)

	if err != nil {
		return err
	}

	*o = ListParentalcontrolSubscriberrecordResponseObject(varListParentalcontrolSubscriberrecordResponseObject)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "result")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableListParentalcontrolSubscriberrecordResponseObject struct {
	value *ListParentalcontrolSubscriberrecordResponseObject
	isSet bool
}

func (v NullableListParentalcontrolSubscriberrecordResponseObject) Get() *ListParentalcontrolSubscriberrecordResponseObject {
	return v.value
}

func (v *NullableListParentalcontrolSubscriberrecordResponseObject) Set(val *ListParentalcontrolSubscriberrecordResponseObject) {
	v.value = val
	v.isSet = true
}

func (v NullableListParentalcontrolSubscriberrecordResponseObject) IsSet() bool {
	return v.isSet
}

func (v *NullableListParentalcontrolSubscriberrecordResponseObject) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListParentalcontrolSubscriberrecordResponseObject(val *ListParentalcontrolSubscriberrecordResponseObject) *NullableListParentalcontrolSubscriberrecordResponseObject {
	return &NullableListParentalcontrolSubscriberrecordResponseObject{value: val, isSet: true}
}

func (v NullableListParentalcontrolSubscriberrecordResponseObject) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListParentalcontrolSubscriberrecordResponseObject) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
