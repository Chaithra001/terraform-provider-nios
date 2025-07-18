/*
Infoblox MICROSOFTSERVER API

OpenAPI specification for Infoblox NIOS WAPI MICROSOFTSERVER objects

API version: 2.13.6
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package microsoftserver

import (
	"encoding/json"
)

// checks if the Mssuperscope type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Mssuperscope{}

// Mssuperscope struct for Mssuperscope
type Mssuperscope struct {
	// The reference to the object.
	Ref *string `json:"_ref,omitempty"`
	// The superscope descriptive comment.
	Comment *string `json:"comment,omitempty"`
	// The percentage of the total DHCP usage of the ranges in the superscope.
	DhcpUtilization *int64 `json:"dhcp_utilization,omitempty"`
	// Utilization level of the DHCP range objects that belong to the superscope.
	DhcpUtilizationStatus *string `json:"dhcp_utilization_status,omitempty"`
	// Determines whether the superscope is disabled.
	Disable *bool `json:"disable,omitempty"`
	// The total number of DHCP leases issued for the DHCP range objects that belong to the superscope.
	DynamicHosts *int64 `json:"dynamic_hosts,omitempty"`
	// Extensible attributes associated with the object. For valid values for extensible attributes, see {extattrs:values}.
	ExtAttrs *map[string]ExtAttrs `json:"extattrs,omitempty"`
	// The percentage value for DHCP range usage after which an alarm will be active.
	HighWaterMark *int64 `json:"high_water_mark,omitempty"`
	// The percentage value for DHCP range usage after which an alarm will be reset.
	HighWaterMarkReset *int64 `json:"high_water_mark_reset,omitempty"`
	// The percentage value for DHCP range usage below which an alarm will be active.
	LowWaterMark *int64 `json:"low_water_mark,omitempty"`
	// The percentage value for DHCP range usage below which an alarm will be reset.
	LowWaterMarkReset *int64 `json:"low_water_mark_reset,omitempty"`
	// The name of the Microsoft DHCP superscope.
	Name *string `json:"name,omitempty"`
	// The name of the network view in which the superscope resides.
	NetworkView *string `json:"network_view,omitempty"`
	// The list of DHCP ranges that are associated with the superscope.
	Ranges []string `json:"ranges,omitempty"`
	// The number of static DHCP addresses configured in DHCP range objects that belong to the superscope.
	StaticHosts *int64 `json:"static_hosts,omitempty"`
	// The total number of DHCP addresses configured in DHCP range objects that belong to the superscope.
	TotalHosts *int64 `json:"total_hosts,omitempty"`
}

// NewMssuperscope instantiates a new Mssuperscope object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMssuperscope() *Mssuperscope {
	this := Mssuperscope{}
	return &this
}

// NewMssuperscopeWithDefaults instantiates a new Mssuperscope object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMssuperscopeWithDefaults() *Mssuperscope {
	this := Mssuperscope{}
	return &this
}

// GetRef returns the Ref field value if set, zero value otherwise.
func (o *Mssuperscope) GetRef() string {
	if o == nil || IsNil(o.Ref) {
		var ret string
		return ret
	}
	return *o.Ref
}

// GetRefOk returns a tuple with the Ref field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Mssuperscope) GetRefOk() (*string, bool) {
	if o == nil || IsNil(o.Ref) {
		return nil, false
	}
	return o.Ref, true
}

// HasRef returns a boolean if a field has been set.
func (o *Mssuperscope) HasRef() bool {
	if o != nil && !IsNil(o.Ref) {
		return true
	}

	return false
}

// SetRef gets a reference to the given string and assigns it to the Ref field.
func (o *Mssuperscope) SetRef(v string) {
	o.Ref = &v
}

// GetComment returns the Comment field value if set, zero value otherwise.
func (o *Mssuperscope) GetComment() string {
	if o == nil || IsNil(o.Comment) {
		var ret string
		return ret
	}
	return *o.Comment
}

// GetCommentOk returns a tuple with the Comment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Mssuperscope) GetCommentOk() (*string, bool) {
	if o == nil || IsNil(o.Comment) {
		return nil, false
	}
	return o.Comment, true
}

// HasComment returns a boolean if a field has been set.
func (o *Mssuperscope) HasComment() bool {
	if o != nil && !IsNil(o.Comment) {
		return true
	}

	return false
}

// SetComment gets a reference to the given string and assigns it to the Comment field.
func (o *Mssuperscope) SetComment(v string) {
	o.Comment = &v
}

// GetDhcpUtilization returns the DhcpUtilization field value if set, zero value otherwise.
func (o *Mssuperscope) GetDhcpUtilization() int64 {
	if o == nil || IsNil(o.DhcpUtilization) {
		var ret int64
		return ret
	}
	return *o.DhcpUtilization
}

// GetDhcpUtilizationOk returns a tuple with the DhcpUtilization field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Mssuperscope) GetDhcpUtilizationOk() (*int64, bool) {
	if o == nil || IsNil(o.DhcpUtilization) {
		return nil, false
	}
	return o.DhcpUtilization, true
}

// HasDhcpUtilization returns a boolean if a field has been set.
func (o *Mssuperscope) HasDhcpUtilization() bool {
	if o != nil && !IsNil(o.DhcpUtilization) {
		return true
	}

	return false
}

// SetDhcpUtilization gets a reference to the given int64 and assigns it to the DhcpUtilization field.
func (o *Mssuperscope) SetDhcpUtilization(v int64) {
	o.DhcpUtilization = &v
}

// GetDhcpUtilizationStatus returns the DhcpUtilizationStatus field value if set, zero value otherwise.
func (o *Mssuperscope) GetDhcpUtilizationStatus() string {
	if o == nil || IsNil(o.DhcpUtilizationStatus) {
		var ret string
		return ret
	}
	return *o.DhcpUtilizationStatus
}

// GetDhcpUtilizationStatusOk returns a tuple with the DhcpUtilizationStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Mssuperscope) GetDhcpUtilizationStatusOk() (*string, bool) {
	if o == nil || IsNil(o.DhcpUtilizationStatus) {
		return nil, false
	}
	return o.DhcpUtilizationStatus, true
}

// HasDhcpUtilizationStatus returns a boolean if a field has been set.
func (o *Mssuperscope) HasDhcpUtilizationStatus() bool {
	if o != nil && !IsNil(o.DhcpUtilizationStatus) {
		return true
	}

	return false
}

// SetDhcpUtilizationStatus gets a reference to the given string and assigns it to the DhcpUtilizationStatus field.
func (o *Mssuperscope) SetDhcpUtilizationStatus(v string) {
	o.DhcpUtilizationStatus = &v
}

// GetDisable returns the Disable field value if set, zero value otherwise.
func (o *Mssuperscope) GetDisable() bool {
	if o == nil || IsNil(o.Disable) {
		var ret bool
		return ret
	}
	return *o.Disable
}

// GetDisableOk returns a tuple with the Disable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Mssuperscope) GetDisableOk() (*bool, bool) {
	if o == nil || IsNil(o.Disable) {
		return nil, false
	}
	return o.Disable, true
}

// HasDisable returns a boolean if a field has been set.
func (o *Mssuperscope) HasDisable() bool {
	if o != nil && !IsNil(o.Disable) {
		return true
	}

	return false
}

// SetDisable gets a reference to the given bool and assigns it to the Disable field.
func (o *Mssuperscope) SetDisable(v bool) {
	o.Disable = &v
}

// GetDynamicHosts returns the DynamicHosts field value if set, zero value otherwise.
func (o *Mssuperscope) GetDynamicHosts() int64 {
	if o == nil || IsNil(o.DynamicHosts) {
		var ret int64
		return ret
	}
	return *o.DynamicHosts
}

// GetDynamicHostsOk returns a tuple with the DynamicHosts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Mssuperscope) GetDynamicHostsOk() (*int64, bool) {
	if o == nil || IsNil(o.DynamicHosts) {
		return nil, false
	}
	return o.DynamicHosts, true
}

// HasDynamicHosts returns a boolean if a field has been set.
func (o *Mssuperscope) HasDynamicHosts() bool {
	if o != nil && !IsNil(o.DynamicHosts) {
		return true
	}

	return false
}

// SetDynamicHosts gets a reference to the given int64 and assigns it to the DynamicHosts field.
func (o *Mssuperscope) SetDynamicHosts(v int64) {
	o.DynamicHosts = &v
}

// GetExtAttrs returns the ExtAttrs field value if set, zero value otherwise.
func (o *Mssuperscope) GetExtAttrs() map[string]ExtAttrs {
	if o == nil || IsNil(o.ExtAttrs) {
		var ret map[string]ExtAttrs
		return ret
	}
	return *o.ExtAttrs
}

// GetExtAttrsOk returns a tuple with the ExtAttrs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Mssuperscope) GetExtAttrsOk() (*map[string]ExtAttrs, bool) {
	if o == nil || IsNil(o.ExtAttrs) {
		return nil, false
	}
	return o.ExtAttrs, true
}

// HasExtAttrs returns a boolean if a field has been set.
func (o *Mssuperscope) HasExtAttrs() bool {
	if o != nil && !IsNil(o.ExtAttrs) {
		return true
	}

	return false
}

// SetExtAttrs gets a reference to the given map[string]ExtAttrs and assigns it to the ExtAttrs field.
func (o *Mssuperscope) SetExtAttrs(v map[string]ExtAttrs) {
	o.ExtAttrs = &v
}

// GetHighWaterMark returns the HighWaterMark field value if set, zero value otherwise.
func (o *Mssuperscope) GetHighWaterMark() int64 {
	if o == nil || IsNil(o.HighWaterMark) {
		var ret int64
		return ret
	}
	return *o.HighWaterMark
}

// GetHighWaterMarkOk returns a tuple with the HighWaterMark field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Mssuperscope) GetHighWaterMarkOk() (*int64, bool) {
	if o == nil || IsNil(o.HighWaterMark) {
		return nil, false
	}
	return o.HighWaterMark, true
}

// HasHighWaterMark returns a boolean if a field has been set.
func (o *Mssuperscope) HasHighWaterMark() bool {
	if o != nil && !IsNil(o.HighWaterMark) {
		return true
	}

	return false
}

// SetHighWaterMark gets a reference to the given int64 and assigns it to the HighWaterMark field.
func (o *Mssuperscope) SetHighWaterMark(v int64) {
	o.HighWaterMark = &v
}

// GetHighWaterMarkReset returns the HighWaterMarkReset field value if set, zero value otherwise.
func (o *Mssuperscope) GetHighWaterMarkReset() int64 {
	if o == nil || IsNil(o.HighWaterMarkReset) {
		var ret int64
		return ret
	}
	return *o.HighWaterMarkReset
}

// GetHighWaterMarkResetOk returns a tuple with the HighWaterMarkReset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Mssuperscope) GetHighWaterMarkResetOk() (*int64, bool) {
	if o == nil || IsNil(o.HighWaterMarkReset) {
		return nil, false
	}
	return o.HighWaterMarkReset, true
}

// HasHighWaterMarkReset returns a boolean if a field has been set.
func (o *Mssuperscope) HasHighWaterMarkReset() bool {
	if o != nil && !IsNil(o.HighWaterMarkReset) {
		return true
	}

	return false
}

// SetHighWaterMarkReset gets a reference to the given int64 and assigns it to the HighWaterMarkReset field.
func (o *Mssuperscope) SetHighWaterMarkReset(v int64) {
	o.HighWaterMarkReset = &v
}

// GetLowWaterMark returns the LowWaterMark field value if set, zero value otherwise.
func (o *Mssuperscope) GetLowWaterMark() int64 {
	if o == nil || IsNil(o.LowWaterMark) {
		var ret int64
		return ret
	}
	return *o.LowWaterMark
}

// GetLowWaterMarkOk returns a tuple with the LowWaterMark field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Mssuperscope) GetLowWaterMarkOk() (*int64, bool) {
	if o == nil || IsNil(o.LowWaterMark) {
		return nil, false
	}
	return o.LowWaterMark, true
}

// HasLowWaterMark returns a boolean if a field has been set.
func (o *Mssuperscope) HasLowWaterMark() bool {
	if o != nil && !IsNil(o.LowWaterMark) {
		return true
	}

	return false
}

// SetLowWaterMark gets a reference to the given int64 and assigns it to the LowWaterMark field.
func (o *Mssuperscope) SetLowWaterMark(v int64) {
	o.LowWaterMark = &v
}

// GetLowWaterMarkReset returns the LowWaterMarkReset field value if set, zero value otherwise.
func (o *Mssuperscope) GetLowWaterMarkReset() int64 {
	if o == nil || IsNil(o.LowWaterMarkReset) {
		var ret int64
		return ret
	}
	return *o.LowWaterMarkReset
}

// GetLowWaterMarkResetOk returns a tuple with the LowWaterMarkReset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Mssuperscope) GetLowWaterMarkResetOk() (*int64, bool) {
	if o == nil || IsNil(o.LowWaterMarkReset) {
		return nil, false
	}
	return o.LowWaterMarkReset, true
}

// HasLowWaterMarkReset returns a boolean if a field has been set.
func (o *Mssuperscope) HasLowWaterMarkReset() bool {
	if o != nil && !IsNil(o.LowWaterMarkReset) {
		return true
	}

	return false
}

// SetLowWaterMarkReset gets a reference to the given int64 and assigns it to the LowWaterMarkReset field.
func (o *Mssuperscope) SetLowWaterMarkReset(v int64) {
	o.LowWaterMarkReset = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *Mssuperscope) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Mssuperscope) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *Mssuperscope) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *Mssuperscope) SetName(v string) {
	o.Name = &v
}

// GetNetworkView returns the NetworkView field value if set, zero value otherwise.
func (o *Mssuperscope) GetNetworkView() string {
	if o == nil || IsNil(o.NetworkView) {
		var ret string
		return ret
	}
	return *o.NetworkView
}

// GetNetworkViewOk returns a tuple with the NetworkView field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Mssuperscope) GetNetworkViewOk() (*string, bool) {
	if o == nil || IsNil(o.NetworkView) {
		return nil, false
	}
	return o.NetworkView, true
}

// HasNetworkView returns a boolean if a field has been set.
func (o *Mssuperscope) HasNetworkView() bool {
	if o != nil && !IsNil(o.NetworkView) {
		return true
	}

	return false
}

// SetNetworkView gets a reference to the given string and assigns it to the NetworkView field.
func (o *Mssuperscope) SetNetworkView(v string) {
	o.NetworkView = &v
}

// GetRanges returns the Ranges field value if set, zero value otherwise.
func (o *Mssuperscope) GetRanges() []string {
	if o == nil || IsNil(o.Ranges) {
		var ret []string
		return ret
	}
	return o.Ranges
}

// GetRangesOk returns a tuple with the Ranges field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Mssuperscope) GetRangesOk() ([]string, bool) {
	if o == nil || IsNil(o.Ranges) {
		return nil, false
	}
	return o.Ranges, true
}

// HasRanges returns a boolean if a field has been set.
func (o *Mssuperscope) HasRanges() bool {
	if o != nil && !IsNil(o.Ranges) {
		return true
	}

	return false
}

// SetRanges gets a reference to the given []string and assigns it to the Ranges field.
func (o *Mssuperscope) SetRanges(v []string) {
	o.Ranges = v
}

// GetStaticHosts returns the StaticHosts field value if set, zero value otherwise.
func (o *Mssuperscope) GetStaticHosts() int64 {
	if o == nil || IsNil(o.StaticHosts) {
		var ret int64
		return ret
	}
	return *o.StaticHosts
}

// GetStaticHostsOk returns a tuple with the StaticHosts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Mssuperscope) GetStaticHostsOk() (*int64, bool) {
	if o == nil || IsNil(o.StaticHosts) {
		return nil, false
	}
	return o.StaticHosts, true
}

// HasStaticHosts returns a boolean if a field has been set.
func (o *Mssuperscope) HasStaticHosts() bool {
	if o != nil && !IsNil(o.StaticHosts) {
		return true
	}

	return false
}

// SetStaticHosts gets a reference to the given int64 and assigns it to the StaticHosts field.
func (o *Mssuperscope) SetStaticHosts(v int64) {
	o.StaticHosts = &v
}

// GetTotalHosts returns the TotalHosts field value if set, zero value otherwise.
func (o *Mssuperscope) GetTotalHosts() int64 {
	if o == nil || IsNil(o.TotalHosts) {
		var ret int64
		return ret
	}
	return *o.TotalHosts
}

// GetTotalHostsOk returns a tuple with the TotalHosts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Mssuperscope) GetTotalHostsOk() (*int64, bool) {
	if o == nil || IsNil(o.TotalHosts) {
		return nil, false
	}
	return o.TotalHosts, true
}

// HasTotalHosts returns a boolean if a field has been set.
func (o *Mssuperscope) HasTotalHosts() bool {
	if o != nil && !IsNil(o.TotalHosts) {
		return true
	}

	return false
}

// SetTotalHosts gets a reference to the given int64 and assigns it to the TotalHosts field.
func (o *Mssuperscope) SetTotalHosts(v int64) {
	o.TotalHosts = &v
}

func (o Mssuperscope) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Mssuperscope) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Ref) {
		toSerialize["_ref"] = o.Ref
	}
	if !IsNil(o.Comment) {
		toSerialize["comment"] = o.Comment
	}
	if !IsNil(o.DhcpUtilization) {
		toSerialize["dhcp_utilization"] = o.DhcpUtilization
	}
	if !IsNil(o.DhcpUtilizationStatus) {
		toSerialize["dhcp_utilization_status"] = o.DhcpUtilizationStatus
	}
	if !IsNil(o.Disable) {
		toSerialize["disable"] = o.Disable
	}
	if !IsNil(o.DynamicHosts) {
		toSerialize["dynamic_hosts"] = o.DynamicHosts
	}
	if !IsNil(o.ExtAttrs) {
		toSerialize["extattrs"] = o.ExtAttrs
	}
	if !IsNil(o.HighWaterMark) {
		toSerialize["high_water_mark"] = o.HighWaterMark
	}
	if !IsNil(o.HighWaterMarkReset) {
		toSerialize["high_water_mark_reset"] = o.HighWaterMarkReset
	}
	if !IsNil(o.LowWaterMark) {
		toSerialize["low_water_mark"] = o.LowWaterMark
	}
	if !IsNil(o.LowWaterMarkReset) {
		toSerialize["low_water_mark_reset"] = o.LowWaterMarkReset
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.NetworkView) {
		toSerialize["network_view"] = o.NetworkView
	}
	if !IsNil(o.Ranges) {
		toSerialize["ranges"] = o.Ranges
	}
	if !IsNil(o.StaticHosts) {
		toSerialize["static_hosts"] = o.StaticHosts
	}
	if !IsNil(o.TotalHosts) {
		toSerialize["total_hosts"] = o.TotalHosts
	}
	return toSerialize, nil
}

type NullableMssuperscope struct {
	value *Mssuperscope
	isSet bool
}

func (v NullableMssuperscope) Get() *Mssuperscope {
	return v.value
}

func (v *NullableMssuperscope) Set(val *Mssuperscope) {
	v.value = val
	v.isSet = true
}

func (v NullableMssuperscope) IsSet() bool {
	return v.isSet
}

func (v *NullableMssuperscope) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMssuperscope(val *Mssuperscope) *NullableMssuperscope {
	return &NullableMssuperscope{value: val, isSet: true}
}

func (v NullableMssuperscope) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMssuperscope) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
