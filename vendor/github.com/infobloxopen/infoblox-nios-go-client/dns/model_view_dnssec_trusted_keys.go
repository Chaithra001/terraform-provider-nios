/*
Infoblox DNS API

OpenAPI specification for Infoblox NIOS WAPI DNS objects

API version: 2.13.6
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package dns

import (
	"encoding/json"
)

// checks if the ViewDnssecTrustedKeys type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ViewDnssecTrustedKeys{}

// ViewDnssecTrustedKeys struct for ViewDnssecTrustedKeys
type ViewDnssecTrustedKeys struct {
	// The FQDN of the domain for which the member validates responses to recursive queries.
	Fqdn *string `json:"fqdn,omitempty"`
	// The DNSSEC algorithm used to generate the key.
	Algorithm *string `json:"algorithm,omitempty"`
	// The DNSSEC key.
	Key *string `json:"key,omitempty"`
	// The secure entry point flag, if set it means this is a KSK configuration.
	SecureEntryPoint *bool `json:"secure_entry_point,omitempty"`
	// Responses must be DNSSEC secure for this hierarchy/domain.
	DnssecMustBeSecure   *bool `json:"dnssec_must_be_secure,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ViewDnssecTrustedKeys ViewDnssecTrustedKeys

// NewViewDnssecTrustedKeys instantiates a new ViewDnssecTrustedKeys object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewViewDnssecTrustedKeys() *ViewDnssecTrustedKeys {
	this := ViewDnssecTrustedKeys{}
	return &this
}

// NewViewDnssecTrustedKeysWithDefaults instantiates a new ViewDnssecTrustedKeys object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewViewDnssecTrustedKeysWithDefaults() *ViewDnssecTrustedKeys {
	this := ViewDnssecTrustedKeys{}
	return &this
}

// GetFqdn returns the Fqdn field value if set, zero value otherwise.
func (o *ViewDnssecTrustedKeys) GetFqdn() string {
	if o == nil || IsNil(o.Fqdn) {
		var ret string
		return ret
	}
	return *o.Fqdn
}

// GetFqdnOk returns a tuple with the Fqdn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ViewDnssecTrustedKeys) GetFqdnOk() (*string, bool) {
	if o == nil || IsNil(o.Fqdn) {
		return nil, false
	}
	return o.Fqdn, true
}

// HasFqdn returns a boolean if a field has been set.
func (o *ViewDnssecTrustedKeys) HasFqdn() bool {
	if o != nil && !IsNil(o.Fqdn) {
		return true
	}

	return false
}

// SetFqdn gets a reference to the given string and assigns it to the Fqdn field.
func (o *ViewDnssecTrustedKeys) SetFqdn(v string) {
	o.Fqdn = &v
}

// GetAlgorithm returns the Algorithm field value if set, zero value otherwise.
func (o *ViewDnssecTrustedKeys) GetAlgorithm() string {
	if o == nil || IsNil(o.Algorithm) {
		var ret string
		return ret
	}
	return *o.Algorithm
}

// GetAlgorithmOk returns a tuple with the Algorithm field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ViewDnssecTrustedKeys) GetAlgorithmOk() (*string, bool) {
	if o == nil || IsNil(o.Algorithm) {
		return nil, false
	}
	return o.Algorithm, true
}

// HasAlgorithm returns a boolean if a field has been set.
func (o *ViewDnssecTrustedKeys) HasAlgorithm() bool {
	if o != nil && !IsNil(o.Algorithm) {
		return true
	}

	return false
}

// SetAlgorithm gets a reference to the given string and assigns it to the Algorithm field.
func (o *ViewDnssecTrustedKeys) SetAlgorithm(v string) {
	o.Algorithm = &v
}

// GetKey returns the Key field value if set, zero value otherwise.
func (o *ViewDnssecTrustedKeys) GetKey() string {
	if o == nil || IsNil(o.Key) {
		var ret string
		return ret
	}
	return *o.Key
}

// GetKeyOk returns a tuple with the Key field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ViewDnssecTrustedKeys) GetKeyOk() (*string, bool) {
	if o == nil || IsNil(o.Key) {
		return nil, false
	}
	return o.Key, true
}

// HasKey returns a boolean if a field has been set.
func (o *ViewDnssecTrustedKeys) HasKey() bool {
	if o != nil && !IsNil(o.Key) {
		return true
	}

	return false
}

// SetKey gets a reference to the given string and assigns it to the Key field.
func (o *ViewDnssecTrustedKeys) SetKey(v string) {
	o.Key = &v
}

// GetSecureEntryPoint returns the SecureEntryPoint field value if set, zero value otherwise.
func (o *ViewDnssecTrustedKeys) GetSecureEntryPoint() bool {
	if o == nil || IsNil(o.SecureEntryPoint) {
		var ret bool
		return ret
	}
	return *o.SecureEntryPoint
}

// GetSecureEntryPointOk returns a tuple with the SecureEntryPoint field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ViewDnssecTrustedKeys) GetSecureEntryPointOk() (*bool, bool) {
	if o == nil || IsNil(o.SecureEntryPoint) {
		return nil, false
	}
	return o.SecureEntryPoint, true
}

// HasSecureEntryPoint returns a boolean if a field has been set.
func (o *ViewDnssecTrustedKeys) HasSecureEntryPoint() bool {
	if o != nil && !IsNil(o.SecureEntryPoint) {
		return true
	}

	return false
}

// SetSecureEntryPoint gets a reference to the given bool and assigns it to the SecureEntryPoint field.
func (o *ViewDnssecTrustedKeys) SetSecureEntryPoint(v bool) {
	o.SecureEntryPoint = &v
}

// GetDnssecMustBeSecure returns the DnssecMustBeSecure field value if set, zero value otherwise.
func (o *ViewDnssecTrustedKeys) GetDnssecMustBeSecure() bool {
	if o == nil || IsNil(o.DnssecMustBeSecure) {
		var ret bool
		return ret
	}
	return *o.DnssecMustBeSecure
}

// GetDnssecMustBeSecureOk returns a tuple with the DnssecMustBeSecure field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ViewDnssecTrustedKeys) GetDnssecMustBeSecureOk() (*bool, bool) {
	if o == nil || IsNil(o.DnssecMustBeSecure) {
		return nil, false
	}
	return o.DnssecMustBeSecure, true
}

// HasDnssecMustBeSecure returns a boolean if a field has been set.
func (o *ViewDnssecTrustedKeys) HasDnssecMustBeSecure() bool {
	if o != nil && !IsNil(o.DnssecMustBeSecure) {
		return true
	}

	return false
}

// SetDnssecMustBeSecure gets a reference to the given bool and assigns it to the DnssecMustBeSecure field.
func (o *ViewDnssecTrustedKeys) SetDnssecMustBeSecure(v bool) {
	o.DnssecMustBeSecure = &v
}

func (o ViewDnssecTrustedKeys) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ViewDnssecTrustedKeys) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Fqdn) {
		toSerialize["fqdn"] = o.Fqdn
	}
	if !IsNil(o.Algorithm) {
		toSerialize["algorithm"] = o.Algorithm
	}
	if !IsNil(o.Key) {
		toSerialize["key"] = o.Key
	}
	if !IsNil(o.SecureEntryPoint) {
		toSerialize["secure_entry_point"] = o.SecureEntryPoint
	}
	if !IsNil(o.DnssecMustBeSecure) {
		toSerialize["dnssec_must_be_secure"] = o.DnssecMustBeSecure
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ViewDnssecTrustedKeys) UnmarshalJSON(data []byte) (err error) {
	varViewDnssecTrustedKeys := _ViewDnssecTrustedKeys{}

	err = json.Unmarshal(data, &varViewDnssecTrustedKeys)

	if err != nil {
		return err
	}

	*o = ViewDnssecTrustedKeys(varViewDnssecTrustedKeys)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "fqdn")
		delete(additionalProperties, "algorithm")
		delete(additionalProperties, "key")
		delete(additionalProperties, "secure_entry_point")
		delete(additionalProperties, "dnssec_must_be_secure")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableViewDnssecTrustedKeys struct {
	value *ViewDnssecTrustedKeys
	isSet bool
}

func (v NullableViewDnssecTrustedKeys) Get() *ViewDnssecTrustedKeys {
	return v.value
}

func (v *NullableViewDnssecTrustedKeys) Set(val *ViewDnssecTrustedKeys) {
	v.value = val
	v.isSet = true
}

func (v NullableViewDnssecTrustedKeys) IsSet() bool {
	return v.isSet
}

func (v *NullableViewDnssecTrustedKeys) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableViewDnssecTrustedKeys(val *ViewDnssecTrustedKeys) *NullableViewDnssecTrustedKeys {
	return &NullableViewDnssecTrustedKeys{value: val, isSet: true}
}

func (v NullableViewDnssecTrustedKeys) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableViewDnssecTrustedKeys) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
