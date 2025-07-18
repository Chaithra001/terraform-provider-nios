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

// checks if the ParentalcontrolAvp type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ParentalcontrolAvp{}

// ParentalcontrolAvp struct for ParentalcontrolAvp
type ParentalcontrolAvp struct {
	// The reference to the object.
	Ref *string `json:"_ref,omitempty"`
	// The comment for the AVP.
	Comment *string `json:"comment,omitempty"`
	// The list of domains applicable to AVP.
	DomainTypes []string `json:"domain_types,omitempty"`
	// Determines if AVP is restricted to domains.
	IsRestricted *bool `json:"is_restricted,omitempty"`
	// The name of AVP.
	Name *string `json:"name,omitempty"`
	// The type of AVP as per RFC 2865/2866.
	Type *int64 `json:"type,omitempty"`
	// Determines if AVP was defined by user.
	UserDefined *bool `json:"user_defined,omitempty"`
	// The type of value.
	ValueType *string `json:"value_type,omitempty"`
	// The vendor ID as per RFC 2865/2866.
	VendorId *int64 `json:"vendor_id,omitempty"`
	// The vendor type as per RFC 2865/2866.
	VendorType *int64 `json:"vendor_type,omitempty"`
}

// NewParentalcontrolAvp instantiates a new ParentalcontrolAvp object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewParentalcontrolAvp() *ParentalcontrolAvp {
	this := ParentalcontrolAvp{}
	return &this
}

// NewParentalcontrolAvpWithDefaults instantiates a new ParentalcontrolAvp object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewParentalcontrolAvpWithDefaults() *ParentalcontrolAvp {
	this := ParentalcontrolAvp{}
	return &this
}

// GetRef returns the Ref field value if set, zero value otherwise.
func (o *ParentalcontrolAvp) GetRef() string {
	if o == nil || IsNil(o.Ref) {
		var ret string
		return ret
	}
	return *o.Ref
}

// GetRefOk returns a tuple with the Ref field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ParentalcontrolAvp) GetRefOk() (*string, bool) {
	if o == nil || IsNil(o.Ref) {
		return nil, false
	}
	return o.Ref, true
}

// HasRef returns a boolean if a field has been set.
func (o *ParentalcontrolAvp) HasRef() bool {
	if o != nil && !IsNil(o.Ref) {
		return true
	}

	return false
}

// SetRef gets a reference to the given string and assigns it to the Ref field.
func (o *ParentalcontrolAvp) SetRef(v string) {
	o.Ref = &v
}

// GetComment returns the Comment field value if set, zero value otherwise.
func (o *ParentalcontrolAvp) GetComment() string {
	if o == nil || IsNil(o.Comment) {
		var ret string
		return ret
	}
	return *o.Comment
}

// GetCommentOk returns a tuple with the Comment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ParentalcontrolAvp) GetCommentOk() (*string, bool) {
	if o == nil || IsNil(o.Comment) {
		return nil, false
	}
	return o.Comment, true
}

// HasComment returns a boolean if a field has been set.
func (o *ParentalcontrolAvp) HasComment() bool {
	if o != nil && !IsNil(o.Comment) {
		return true
	}

	return false
}

// SetComment gets a reference to the given string and assigns it to the Comment field.
func (o *ParentalcontrolAvp) SetComment(v string) {
	o.Comment = &v
}

// GetDomainTypes returns the DomainTypes field value if set, zero value otherwise.
func (o *ParentalcontrolAvp) GetDomainTypes() []string {
	if o == nil || IsNil(o.DomainTypes) {
		var ret []string
		return ret
	}
	return o.DomainTypes
}

// GetDomainTypesOk returns a tuple with the DomainTypes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ParentalcontrolAvp) GetDomainTypesOk() ([]string, bool) {
	if o == nil || IsNil(o.DomainTypes) {
		return nil, false
	}
	return o.DomainTypes, true
}

// HasDomainTypes returns a boolean if a field has been set.
func (o *ParentalcontrolAvp) HasDomainTypes() bool {
	if o != nil && !IsNil(o.DomainTypes) {
		return true
	}

	return false
}

// SetDomainTypes gets a reference to the given []string and assigns it to the DomainTypes field.
func (o *ParentalcontrolAvp) SetDomainTypes(v []string) {
	o.DomainTypes = v
}

// GetIsRestricted returns the IsRestricted field value if set, zero value otherwise.
func (o *ParentalcontrolAvp) GetIsRestricted() bool {
	if o == nil || IsNil(o.IsRestricted) {
		var ret bool
		return ret
	}
	return *o.IsRestricted
}

// GetIsRestrictedOk returns a tuple with the IsRestricted field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ParentalcontrolAvp) GetIsRestrictedOk() (*bool, bool) {
	if o == nil || IsNil(o.IsRestricted) {
		return nil, false
	}
	return o.IsRestricted, true
}

// HasIsRestricted returns a boolean if a field has been set.
func (o *ParentalcontrolAvp) HasIsRestricted() bool {
	if o != nil && !IsNil(o.IsRestricted) {
		return true
	}

	return false
}

// SetIsRestricted gets a reference to the given bool and assigns it to the IsRestricted field.
func (o *ParentalcontrolAvp) SetIsRestricted(v bool) {
	o.IsRestricted = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ParentalcontrolAvp) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ParentalcontrolAvp) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ParentalcontrolAvp) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ParentalcontrolAvp) SetName(v string) {
	o.Name = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *ParentalcontrolAvp) GetType() int64 {
	if o == nil || IsNil(o.Type) {
		var ret int64
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ParentalcontrolAvp) GetTypeOk() (*int64, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *ParentalcontrolAvp) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given int64 and assigns it to the Type field.
func (o *ParentalcontrolAvp) SetType(v int64) {
	o.Type = &v
}

// GetUserDefined returns the UserDefined field value if set, zero value otherwise.
func (o *ParentalcontrolAvp) GetUserDefined() bool {
	if o == nil || IsNil(o.UserDefined) {
		var ret bool
		return ret
	}
	return *o.UserDefined
}

// GetUserDefinedOk returns a tuple with the UserDefined field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ParentalcontrolAvp) GetUserDefinedOk() (*bool, bool) {
	if o == nil || IsNil(o.UserDefined) {
		return nil, false
	}
	return o.UserDefined, true
}

// HasUserDefined returns a boolean if a field has been set.
func (o *ParentalcontrolAvp) HasUserDefined() bool {
	if o != nil && !IsNil(o.UserDefined) {
		return true
	}

	return false
}

// SetUserDefined gets a reference to the given bool and assigns it to the UserDefined field.
func (o *ParentalcontrolAvp) SetUserDefined(v bool) {
	o.UserDefined = &v
}

// GetValueType returns the ValueType field value if set, zero value otherwise.
func (o *ParentalcontrolAvp) GetValueType() string {
	if o == nil || IsNil(o.ValueType) {
		var ret string
		return ret
	}
	return *o.ValueType
}

// GetValueTypeOk returns a tuple with the ValueType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ParentalcontrolAvp) GetValueTypeOk() (*string, bool) {
	if o == nil || IsNil(o.ValueType) {
		return nil, false
	}
	return o.ValueType, true
}

// HasValueType returns a boolean if a field has been set.
func (o *ParentalcontrolAvp) HasValueType() bool {
	if o != nil && !IsNil(o.ValueType) {
		return true
	}

	return false
}

// SetValueType gets a reference to the given string and assigns it to the ValueType field.
func (o *ParentalcontrolAvp) SetValueType(v string) {
	o.ValueType = &v
}

// GetVendorId returns the VendorId field value if set, zero value otherwise.
func (o *ParentalcontrolAvp) GetVendorId() int64 {
	if o == nil || IsNil(o.VendorId) {
		var ret int64
		return ret
	}
	return *o.VendorId
}

// GetVendorIdOk returns a tuple with the VendorId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ParentalcontrolAvp) GetVendorIdOk() (*int64, bool) {
	if o == nil || IsNil(o.VendorId) {
		return nil, false
	}
	return o.VendorId, true
}

// HasVendorId returns a boolean if a field has been set.
func (o *ParentalcontrolAvp) HasVendorId() bool {
	if o != nil && !IsNil(o.VendorId) {
		return true
	}

	return false
}

// SetVendorId gets a reference to the given int64 and assigns it to the VendorId field.
func (o *ParentalcontrolAvp) SetVendorId(v int64) {
	o.VendorId = &v
}

// GetVendorType returns the VendorType field value if set, zero value otherwise.
func (o *ParentalcontrolAvp) GetVendorType() int64 {
	if o == nil || IsNil(o.VendorType) {
		var ret int64
		return ret
	}
	return *o.VendorType
}

// GetVendorTypeOk returns a tuple with the VendorType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ParentalcontrolAvp) GetVendorTypeOk() (*int64, bool) {
	if o == nil || IsNil(o.VendorType) {
		return nil, false
	}
	return o.VendorType, true
}

// HasVendorType returns a boolean if a field has been set.
func (o *ParentalcontrolAvp) HasVendorType() bool {
	if o != nil && !IsNil(o.VendorType) {
		return true
	}

	return false
}

// SetVendorType gets a reference to the given int64 and assigns it to the VendorType field.
func (o *ParentalcontrolAvp) SetVendorType(v int64) {
	o.VendorType = &v
}

func (o ParentalcontrolAvp) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ParentalcontrolAvp) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Ref) {
		toSerialize["_ref"] = o.Ref
	}
	if !IsNil(o.Comment) {
		toSerialize["comment"] = o.Comment
	}
	if !IsNil(o.DomainTypes) {
		toSerialize["domain_types"] = o.DomainTypes
	}
	if !IsNil(o.IsRestricted) {
		toSerialize["is_restricted"] = o.IsRestricted
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.UserDefined) {
		toSerialize["user_defined"] = o.UserDefined
	}
	if !IsNil(o.ValueType) {
		toSerialize["value_type"] = o.ValueType
	}
	if !IsNil(o.VendorId) {
		toSerialize["vendor_id"] = o.VendorId
	}
	if !IsNil(o.VendorType) {
		toSerialize["vendor_type"] = o.VendorType
	}
	return toSerialize, nil
}

type NullableParentalcontrolAvp struct {
	value *ParentalcontrolAvp
	isSet bool
}

func (v NullableParentalcontrolAvp) Get() *ParentalcontrolAvp {
	return v.value
}

func (v *NullableParentalcontrolAvp) Set(val *ParentalcontrolAvp) {
	v.value = val
	v.isSet = true
}

func (v NullableParentalcontrolAvp) IsSet() bool {
	return v.isSet
}

func (v *NullableParentalcontrolAvp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableParentalcontrolAvp(val *ParentalcontrolAvp) *NullableParentalcontrolAvp {
	return &NullableParentalcontrolAvp{value: val, isSet: true}
}

func (v NullableParentalcontrolAvp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableParentalcontrolAvp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
