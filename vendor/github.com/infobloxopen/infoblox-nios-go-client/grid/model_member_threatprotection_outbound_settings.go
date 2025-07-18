/*
Infoblox GRID API

OpenAPI specification for Infoblox NIOS WAPI GRID objects

API version: 2.13.6
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package grid

import (
	"encoding/json"
)

// checks if the MemberThreatprotectionOutboundSettings type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MemberThreatprotectionOutboundSettings{}

// MemberThreatprotectionOutboundSettings struct for MemberThreatprotectionOutboundSettings
type MemberThreatprotectionOutboundSettings struct {
	// Flag to enable using DNS query FQDN for Outbound.
	EnableQueryFqdn *bool `json:"enable_query_fqdn,omitempty"`
	// Max domain level for DNS Query FQDN
	QueryFqdnLimit       *int64 `json:"query_fqdn_limit,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _MemberThreatprotectionOutboundSettings MemberThreatprotectionOutboundSettings

// NewMemberThreatprotectionOutboundSettings instantiates a new MemberThreatprotectionOutboundSettings object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMemberThreatprotectionOutboundSettings() *MemberThreatprotectionOutboundSettings {
	this := MemberThreatprotectionOutboundSettings{}
	return &this
}

// NewMemberThreatprotectionOutboundSettingsWithDefaults instantiates a new MemberThreatprotectionOutboundSettings object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMemberThreatprotectionOutboundSettingsWithDefaults() *MemberThreatprotectionOutboundSettings {
	this := MemberThreatprotectionOutboundSettings{}
	return &this
}

// GetEnableQueryFqdn returns the EnableQueryFqdn field value if set, zero value otherwise.
func (o *MemberThreatprotectionOutboundSettings) GetEnableQueryFqdn() bool {
	if o == nil || IsNil(o.EnableQueryFqdn) {
		var ret bool
		return ret
	}
	return *o.EnableQueryFqdn
}

// GetEnableQueryFqdnOk returns a tuple with the EnableQueryFqdn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MemberThreatprotectionOutboundSettings) GetEnableQueryFqdnOk() (*bool, bool) {
	if o == nil || IsNil(o.EnableQueryFqdn) {
		return nil, false
	}
	return o.EnableQueryFqdn, true
}

// HasEnableQueryFqdn returns a boolean if a field has been set.
func (o *MemberThreatprotectionOutboundSettings) HasEnableQueryFqdn() bool {
	if o != nil && !IsNil(o.EnableQueryFqdn) {
		return true
	}

	return false
}

// SetEnableQueryFqdn gets a reference to the given bool and assigns it to the EnableQueryFqdn field.
func (o *MemberThreatprotectionOutboundSettings) SetEnableQueryFqdn(v bool) {
	o.EnableQueryFqdn = &v
}

// GetQueryFqdnLimit returns the QueryFqdnLimit field value if set, zero value otherwise.
func (o *MemberThreatprotectionOutboundSettings) GetQueryFqdnLimit() int64 {
	if o == nil || IsNil(o.QueryFqdnLimit) {
		var ret int64
		return ret
	}
	return *o.QueryFqdnLimit
}

// GetQueryFqdnLimitOk returns a tuple with the QueryFqdnLimit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MemberThreatprotectionOutboundSettings) GetQueryFqdnLimitOk() (*int64, bool) {
	if o == nil || IsNil(o.QueryFqdnLimit) {
		return nil, false
	}
	return o.QueryFqdnLimit, true
}

// HasQueryFqdnLimit returns a boolean if a field has been set.
func (o *MemberThreatprotectionOutboundSettings) HasQueryFqdnLimit() bool {
	if o != nil && !IsNil(o.QueryFqdnLimit) {
		return true
	}

	return false
}

// SetQueryFqdnLimit gets a reference to the given int64 and assigns it to the QueryFqdnLimit field.
func (o *MemberThreatprotectionOutboundSettings) SetQueryFqdnLimit(v int64) {
	o.QueryFqdnLimit = &v
}

func (o MemberThreatprotectionOutboundSettings) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MemberThreatprotectionOutboundSettings) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.EnableQueryFqdn) {
		toSerialize["enable_query_fqdn"] = o.EnableQueryFqdn
	}
	if !IsNil(o.QueryFqdnLimit) {
		toSerialize["query_fqdn_limit"] = o.QueryFqdnLimit
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *MemberThreatprotectionOutboundSettings) UnmarshalJSON(data []byte) (err error) {
	varMemberThreatprotectionOutboundSettings := _MemberThreatprotectionOutboundSettings{}

	err = json.Unmarshal(data, &varMemberThreatprotectionOutboundSettings)

	if err != nil {
		return err
	}

	*o = MemberThreatprotectionOutboundSettings(varMemberThreatprotectionOutboundSettings)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "enable_query_fqdn")
		delete(additionalProperties, "query_fqdn_limit")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableMemberThreatprotectionOutboundSettings struct {
	value *MemberThreatprotectionOutboundSettings
	isSet bool
}

func (v NullableMemberThreatprotectionOutboundSettings) Get() *MemberThreatprotectionOutboundSettings {
	return v.value
}

func (v *NullableMemberThreatprotectionOutboundSettings) Set(val *MemberThreatprotectionOutboundSettings) {
	v.value = val
	v.isSet = true
}

func (v NullableMemberThreatprotectionOutboundSettings) IsSet() bool {
	return v.isSet
}

func (v *NullableMemberThreatprotectionOutboundSettings) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMemberThreatprotectionOutboundSettings(val *MemberThreatprotectionOutboundSettings) *NullableMemberThreatprotectionOutboundSettings {
	return &NullableMemberThreatprotectionOutboundSettings{value: val, isSet: true}
}

func (v NullableMemberThreatprotectionOutboundSettings) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMemberThreatprotectionOutboundSettings) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
