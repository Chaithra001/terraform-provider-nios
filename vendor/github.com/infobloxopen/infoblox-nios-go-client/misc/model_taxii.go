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

// checks if the Taxii type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Taxii{}

// Taxii struct for Taxii
type Taxii struct {
	// The reference to the object.
	Ref *string `json:"_ref,omitempty"`
	// Indicates whether the Taxii service is running on the given member or not.
	EnableService *bool `json:"enable_service,omitempty"`
	// The IPv4 Address of the Grid member.
	Ipv4addr *string `json:"ipv4addr,omitempty"`
	// The IPv6 Address of the Grid member.
	Ipv6addr *string `json:"ipv6addr,omitempty"`
	// The name of the Taxii Member.
	Name *string `json:"name,omitempty"`
	// Taxii service RPZ configuration list.
	TaxiiRpzConfig []TaxiiTaxiiRpzConfig `json:"taxii_rpz_config,omitempty"`
}

// NewTaxii instantiates a new Taxii object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTaxii() *Taxii {
	this := Taxii{}
	return &this
}

// NewTaxiiWithDefaults instantiates a new Taxii object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTaxiiWithDefaults() *Taxii {
	this := Taxii{}
	return &this
}

// GetRef returns the Ref field value if set, zero value otherwise.
func (o *Taxii) GetRef() string {
	if o == nil || IsNil(o.Ref) {
		var ret string
		return ret
	}
	return *o.Ref
}

// GetRefOk returns a tuple with the Ref field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Taxii) GetRefOk() (*string, bool) {
	if o == nil || IsNil(o.Ref) {
		return nil, false
	}
	return o.Ref, true
}

// HasRef returns a boolean if a field has been set.
func (o *Taxii) HasRef() bool {
	if o != nil && !IsNil(o.Ref) {
		return true
	}

	return false
}

// SetRef gets a reference to the given string and assigns it to the Ref field.
func (o *Taxii) SetRef(v string) {
	o.Ref = &v
}

// GetEnableService returns the EnableService field value if set, zero value otherwise.
func (o *Taxii) GetEnableService() bool {
	if o == nil || IsNil(o.EnableService) {
		var ret bool
		return ret
	}
	return *o.EnableService
}

// GetEnableServiceOk returns a tuple with the EnableService field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Taxii) GetEnableServiceOk() (*bool, bool) {
	if o == nil || IsNil(o.EnableService) {
		return nil, false
	}
	return o.EnableService, true
}

// HasEnableService returns a boolean if a field has been set.
func (o *Taxii) HasEnableService() bool {
	if o != nil && !IsNil(o.EnableService) {
		return true
	}

	return false
}

// SetEnableService gets a reference to the given bool and assigns it to the EnableService field.
func (o *Taxii) SetEnableService(v bool) {
	o.EnableService = &v
}

// GetIpv4addr returns the Ipv4addr field value if set, zero value otherwise.
func (o *Taxii) GetIpv4addr() string {
	if o == nil || IsNil(o.Ipv4addr) {
		var ret string
		return ret
	}
	return *o.Ipv4addr
}

// GetIpv4addrOk returns a tuple with the Ipv4addr field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Taxii) GetIpv4addrOk() (*string, bool) {
	if o == nil || IsNil(o.Ipv4addr) {
		return nil, false
	}
	return o.Ipv4addr, true
}

// HasIpv4addr returns a boolean if a field has been set.
func (o *Taxii) HasIpv4addr() bool {
	if o != nil && !IsNil(o.Ipv4addr) {
		return true
	}

	return false
}

// SetIpv4addr gets a reference to the given string and assigns it to the Ipv4addr field.
func (o *Taxii) SetIpv4addr(v string) {
	o.Ipv4addr = &v
}

// GetIpv6addr returns the Ipv6addr field value if set, zero value otherwise.
func (o *Taxii) GetIpv6addr() string {
	if o == nil || IsNil(o.Ipv6addr) {
		var ret string
		return ret
	}
	return *o.Ipv6addr
}

// GetIpv6addrOk returns a tuple with the Ipv6addr field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Taxii) GetIpv6addrOk() (*string, bool) {
	if o == nil || IsNil(o.Ipv6addr) {
		return nil, false
	}
	return o.Ipv6addr, true
}

// HasIpv6addr returns a boolean if a field has been set.
func (o *Taxii) HasIpv6addr() bool {
	if o != nil && !IsNil(o.Ipv6addr) {
		return true
	}

	return false
}

// SetIpv6addr gets a reference to the given string and assigns it to the Ipv6addr field.
func (o *Taxii) SetIpv6addr(v string) {
	o.Ipv6addr = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *Taxii) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Taxii) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *Taxii) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *Taxii) SetName(v string) {
	o.Name = &v
}

// GetTaxiiRpzConfig returns the TaxiiRpzConfig field value if set, zero value otherwise.
func (o *Taxii) GetTaxiiRpzConfig() []TaxiiTaxiiRpzConfig {
	if o == nil || IsNil(o.TaxiiRpzConfig) {
		var ret []TaxiiTaxiiRpzConfig
		return ret
	}
	return o.TaxiiRpzConfig
}

// GetTaxiiRpzConfigOk returns a tuple with the TaxiiRpzConfig field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Taxii) GetTaxiiRpzConfigOk() ([]TaxiiTaxiiRpzConfig, bool) {
	if o == nil || IsNil(o.TaxiiRpzConfig) {
		return nil, false
	}
	return o.TaxiiRpzConfig, true
}

// HasTaxiiRpzConfig returns a boolean if a field has been set.
func (o *Taxii) HasTaxiiRpzConfig() bool {
	if o != nil && !IsNil(o.TaxiiRpzConfig) {
		return true
	}

	return false
}

// SetTaxiiRpzConfig gets a reference to the given []TaxiiTaxiiRpzConfig and assigns it to the TaxiiRpzConfig field.
func (o *Taxii) SetTaxiiRpzConfig(v []TaxiiTaxiiRpzConfig) {
	o.TaxiiRpzConfig = v
}

func (o Taxii) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Taxii) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Ref) {
		toSerialize["_ref"] = o.Ref
	}
	if !IsNil(o.EnableService) {
		toSerialize["enable_service"] = o.EnableService
	}
	if !IsNil(o.Ipv4addr) {
		toSerialize["ipv4addr"] = o.Ipv4addr
	}
	if !IsNil(o.Ipv6addr) {
		toSerialize["ipv6addr"] = o.Ipv6addr
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.TaxiiRpzConfig) {
		toSerialize["taxii_rpz_config"] = o.TaxiiRpzConfig
	}
	return toSerialize, nil
}

type NullableTaxii struct {
	value *Taxii
	isSet bool
}

func (v NullableTaxii) Get() *Taxii {
	return v.value
}

func (v *NullableTaxii) Set(val *Taxii) {
	v.value = val
	v.isSet = true
}

func (v NullableTaxii) IsSet() bool {
	return v.isSet
}

func (v *NullableTaxii) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTaxii(val *Taxii) *NullableTaxii {
	return &NullableTaxii{value: val, isSet: true}
}

func (v NullableTaxii) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTaxii) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
