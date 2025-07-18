/*
Infoblox MISC API

OpenAPI specification for Infoblox NIOS WAPI MISC objects

API version: 2.13.6
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package misc

import (
	"encoding/json"
)

// checks if the CreateTftpfiledirResponseAsObject type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateTftpfiledirResponseAsObject{}

// CreateTftpfiledirResponseAsObject The response format to create __Tftpfiledir__ in object format.
type CreateTftpfiledirResponseAsObject struct {
	Result               *Tftpfiledir `json:"result,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CreateTftpfiledirResponseAsObject CreateTftpfiledirResponseAsObject

// NewCreateTftpfiledirResponseAsObject instantiates a new CreateTftpfiledirResponseAsObject object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateTftpfiledirResponseAsObject() *CreateTftpfiledirResponseAsObject {
	this := CreateTftpfiledirResponseAsObject{}
	return &this
}

// NewCreateTftpfiledirResponseAsObjectWithDefaults instantiates a new CreateTftpfiledirResponseAsObject object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateTftpfiledirResponseAsObjectWithDefaults() *CreateTftpfiledirResponseAsObject {
	this := CreateTftpfiledirResponseAsObject{}
	return &this
}

// GetResult returns the Result field value if set, zero value otherwise.
func (o *CreateTftpfiledirResponseAsObject) GetResult() Tftpfiledir {
	if o == nil || IsNil(o.Result) {
		var ret Tftpfiledir
		return ret
	}
	return *o.Result
}

// GetResultOk returns a tuple with the Result field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateTftpfiledirResponseAsObject) GetResultOk() (*Tftpfiledir, bool) {
	if o == nil || IsNil(o.Result) {
		return nil, false
	}
	return o.Result, true
}

// HasResult returns a boolean if a field has been set.
func (o *CreateTftpfiledirResponseAsObject) HasResult() bool {
	if o != nil && !IsNil(o.Result) {
		return true
	}

	return false
}

// SetResult gets a reference to the given Tftpfiledir and assigns it to the Result field.
func (o *CreateTftpfiledirResponseAsObject) SetResult(v Tftpfiledir) {
	o.Result = &v
}

func (o CreateTftpfiledirResponseAsObject) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateTftpfiledirResponseAsObject) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Result) {
		toSerialize["result"] = o.Result
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CreateTftpfiledirResponseAsObject) UnmarshalJSON(data []byte) (err error) {
	varCreateTftpfiledirResponseAsObject := _CreateTftpfiledirResponseAsObject{}

	err = json.Unmarshal(data, &varCreateTftpfiledirResponseAsObject)

	if err != nil {
		return err
	}

	*o = CreateTftpfiledirResponseAsObject(varCreateTftpfiledirResponseAsObject)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "result")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCreateTftpfiledirResponseAsObject struct {
	value *CreateTftpfiledirResponseAsObject
	isSet bool
}

func (v NullableCreateTftpfiledirResponseAsObject) Get() *CreateTftpfiledirResponseAsObject {
	return v.value
}

func (v *NullableCreateTftpfiledirResponseAsObject) Set(val *CreateTftpfiledirResponseAsObject) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateTftpfiledirResponseAsObject) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateTftpfiledirResponseAsObject) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateTftpfiledirResponseAsObject(val *CreateTftpfiledirResponseAsObject) *NullableCreateTftpfiledirResponseAsObject {
	return &NullableCreateTftpfiledirResponseAsObject{value: val, isSet: true}
}

func (v NullableCreateTftpfiledirResponseAsObject) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateTftpfiledirResponseAsObject) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
