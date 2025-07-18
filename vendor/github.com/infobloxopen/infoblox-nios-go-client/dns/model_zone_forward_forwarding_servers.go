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

// checks if the ZoneForwardForwardingServers type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ZoneForwardForwardingServers{}

// ZoneForwardForwardingServers struct for ZoneForwardForwardingServers
type ZoneForwardForwardingServers struct {
	// The name of this Grid member in FQDN format.
	Name *string `json:"name,omitempty"`
	// Determines if the appliance sends queries to forwarders only, and not to other internal or Internet root servers.
	ForwardersOnly *bool `json:"forwarders_only,omitempty"`
	// The information for the remote name server to which you want the Infoblox appliance to forward queries for a specified domain name.
	ForwardTo []ZoneforwardforwardingserversForwardTo `json:"forward_to,omitempty"`
	// Use flag for: forward_to
	UseOverrideForwarders *bool `json:"use_override_forwarders,omitempty"`
	AdditionalProperties  map[string]interface{}
}

type _ZoneForwardForwardingServers ZoneForwardForwardingServers

// NewZoneForwardForwardingServers instantiates a new ZoneForwardForwardingServers object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewZoneForwardForwardingServers() *ZoneForwardForwardingServers {
	this := ZoneForwardForwardingServers{}
	return &this
}

// NewZoneForwardForwardingServersWithDefaults instantiates a new ZoneForwardForwardingServers object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewZoneForwardForwardingServersWithDefaults() *ZoneForwardForwardingServers {
	this := ZoneForwardForwardingServers{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ZoneForwardForwardingServers) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ZoneForwardForwardingServers) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ZoneForwardForwardingServers) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ZoneForwardForwardingServers) SetName(v string) {
	o.Name = &v
}

// GetForwardersOnly returns the ForwardersOnly field value if set, zero value otherwise.
func (o *ZoneForwardForwardingServers) GetForwardersOnly() bool {
	if o == nil || IsNil(o.ForwardersOnly) {
		var ret bool
		return ret
	}
	return *o.ForwardersOnly
}

// GetForwardersOnlyOk returns a tuple with the ForwardersOnly field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ZoneForwardForwardingServers) GetForwardersOnlyOk() (*bool, bool) {
	if o == nil || IsNil(o.ForwardersOnly) {
		return nil, false
	}
	return o.ForwardersOnly, true
}

// HasForwardersOnly returns a boolean if a field has been set.
func (o *ZoneForwardForwardingServers) HasForwardersOnly() bool {
	if o != nil && !IsNil(o.ForwardersOnly) {
		return true
	}

	return false
}

// SetForwardersOnly gets a reference to the given bool and assigns it to the ForwardersOnly field.
func (o *ZoneForwardForwardingServers) SetForwardersOnly(v bool) {
	o.ForwardersOnly = &v
}

// GetForwardTo returns the ForwardTo field value if set, zero value otherwise.
func (o *ZoneForwardForwardingServers) GetForwardTo() []ZoneforwardforwardingserversForwardTo {
	if o == nil || IsNil(o.ForwardTo) {
		var ret []ZoneforwardforwardingserversForwardTo
		return ret
	}
	return o.ForwardTo
}

// GetForwardToOk returns a tuple with the ForwardTo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ZoneForwardForwardingServers) GetForwardToOk() ([]ZoneforwardforwardingserversForwardTo, bool) {
	if o == nil || IsNil(o.ForwardTo) {
		return nil, false
	}
	return o.ForwardTo, true
}

// HasForwardTo returns a boolean if a field has been set.
func (o *ZoneForwardForwardingServers) HasForwardTo() bool {
	if o != nil && !IsNil(o.ForwardTo) {
		return true
	}

	return false
}

// SetForwardTo gets a reference to the given []ZoneforwardforwardingserversForwardTo and assigns it to the ForwardTo field.
func (o *ZoneForwardForwardingServers) SetForwardTo(v []ZoneforwardforwardingserversForwardTo) {
	o.ForwardTo = v
}

// GetUseOverrideForwarders returns the UseOverrideForwarders field value if set, zero value otherwise.
func (o *ZoneForwardForwardingServers) GetUseOverrideForwarders() bool {
	if o == nil || IsNil(o.UseOverrideForwarders) {
		var ret bool
		return ret
	}
	return *o.UseOverrideForwarders
}

// GetUseOverrideForwardersOk returns a tuple with the UseOverrideForwarders field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ZoneForwardForwardingServers) GetUseOverrideForwardersOk() (*bool, bool) {
	if o == nil || IsNil(o.UseOverrideForwarders) {
		return nil, false
	}
	return o.UseOverrideForwarders, true
}

// HasUseOverrideForwarders returns a boolean if a field has been set.
func (o *ZoneForwardForwardingServers) HasUseOverrideForwarders() bool {
	if o != nil && !IsNil(o.UseOverrideForwarders) {
		return true
	}

	return false
}

// SetUseOverrideForwarders gets a reference to the given bool and assigns it to the UseOverrideForwarders field.
func (o *ZoneForwardForwardingServers) SetUseOverrideForwarders(v bool) {
	o.UseOverrideForwarders = &v
}

func (o ZoneForwardForwardingServers) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ZoneForwardForwardingServers) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.ForwardersOnly) {
		toSerialize["forwarders_only"] = o.ForwardersOnly
	}
	if !IsNil(o.ForwardTo) {
		toSerialize["forward_to"] = o.ForwardTo
	}
	if !IsNil(o.UseOverrideForwarders) {
		toSerialize["use_override_forwarders"] = o.UseOverrideForwarders
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ZoneForwardForwardingServers) UnmarshalJSON(data []byte) (err error) {
	varZoneForwardForwardingServers := _ZoneForwardForwardingServers{}

	err = json.Unmarshal(data, &varZoneForwardForwardingServers)

	if err != nil {
		return err
	}

	*o = ZoneForwardForwardingServers(varZoneForwardForwardingServers)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "name")
		delete(additionalProperties, "forwarders_only")
		delete(additionalProperties, "forward_to")
		delete(additionalProperties, "use_override_forwarders")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableZoneForwardForwardingServers struct {
	value *ZoneForwardForwardingServers
	isSet bool
}

func (v NullableZoneForwardForwardingServers) Get() *ZoneForwardForwardingServers {
	return v.value
}

func (v *NullableZoneForwardForwardingServers) Set(val *ZoneForwardForwardingServers) {
	v.value = val
	v.isSet = true
}

func (v NullableZoneForwardForwardingServers) IsSet() bool {
	return v.isSet
}

func (v *NullableZoneForwardForwardingServers) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableZoneForwardForwardingServers(val *ZoneForwardForwardingServers) *NullableZoneForwardForwardingServers {
	return &NullableZoneForwardForwardingServers{value: val, isSet: true}
}

func (v NullableZoneForwardForwardingServers) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableZoneForwardForwardingServers) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
