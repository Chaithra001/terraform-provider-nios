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

// checks if the PxgridEndpoint type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PxgridEndpoint{}

// PxgridEndpoint struct for PxgridEndpoint
type PxgridEndpoint struct {
	// The reference to the object.
	Ref *string `json:"_ref,omitempty"`
	// The pxgrid endpoint IPv4 Address or IPv6 Address or Fully-Qualified Domain Name (FQDN)
	Address *string `json:"address,omitempty"`
	// The Cisco ISE client certificate subject.
	ClientCertificateSubject *string `json:"client_certificate_subject,omitempty"`
	// The token returned by the uploadinit function call in object fileop for Cisco ISE client certificate.
	ClientCertificateToken *string `json:"client_certificate_token,omitempty"`
	// The pxgrid endpoint client certificate valid from.
	ClientCertificateValidFrom *int64 `json:"client_certificate_valid_from,omitempty"`
	// The pxgrid endpoint client certificate valid to.
	ClientCertificateValidTo *int64 `json:"client_certificate_valid_to,omitempty"`
	// The Cisco ISE endpoint descriptive comment.
	Comment *string `json:"comment,omitempty"`
	// Determines whether a Cisco ISE endpoint is disabled or not. When this is set to False, the Cisco ISE endpoint is enabled.
	Disable *bool `json:"disable,omitempty"`
	// Extensible attributes associated with the object. For valid values for extensible attributes, see {extattrs:values}.
	ExtAttrs *map[string]ExtAttrs `json:"extattrs,omitempty"`
	// The log level for a notification pxgrid endpoint.
	LogLevel *string `json:"log_level,omitempty"`
	// The name of the pxgrid endpoint.
	Name *string `json:"name,omitempty"`
	// The pxgrid network view name.
	NetworkView *string `json:"network_view,omitempty"`
	// The outbound member that will generate events.
	OutboundMemberType *string `json:"outbound_member_type,omitempty"`
	// The list of members for outbound events.
	OutboundMembers   []string                         `json:"outbound_members,omitempty"`
	PublishSettings   *PxgridEndpointPublishSettings   `json:"publish_settings,omitempty"`
	SubscribeSettings *PxgridEndpointSubscribeSettings `json:"subscribe_settings,omitempty"`
	TemplateInstance  *PxgridEndpointTemplateInstance  `json:"template_instance,omitempty"`
	// The timeout of session management (in seconds).
	Timeout *int64 `json:"timeout,omitempty"`
	// The vendor identifier.
	VendorIdentifier *string `json:"vendor_identifier,omitempty"`
	// The user name for WAPI integration.
	WapiUserName *string `json:"wapi_user_name,omitempty"`
	// The user password for WAPI integration.
	WapiUserPassword *string `json:"wapi_user_password,omitempty"`
}

// NewPxgridEndpoint instantiates a new PxgridEndpoint object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPxgridEndpoint() *PxgridEndpoint {
	this := PxgridEndpoint{}
	return &this
}

// NewPxgridEndpointWithDefaults instantiates a new PxgridEndpoint object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPxgridEndpointWithDefaults() *PxgridEndpoint {
	this := PxgridEndpoint{}
	return &this
}

// GetRef returns the Ref field value if set, zero value otherwise.
func (o *PxgridEndpoint) GetRef() string {
	if o == nil || IsNil(o.Ref) {
		var ret string
		return ret
	}
	return *o.Ref
}

// GetRefOk returns a tuple with the Ref field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PxgridEndpoint) GetRefOk() (*string, bool) {
	if o == nil || IsNil(o.Ref) {
		return nil, false
	}
	return o.Ref, true
}

// HasRef returns a boolean if a field has been set.
func (o *PxgridEndpoint) HasRef() bool {
	if o != nil && !IsNil(o.Ref) {
		return true
	}

	return false
}

// SetRef gets a reference to the given string and assigns it to the Ref field.
func (o *PxgridEndpoint) SetRef(v string) {
	o.Ref = &v
}

// GetAddress returns the Address field value if set, zero value otherwise.
func (o *PxgridEndpoint) GetAddress() string {
	if o == nil || IsNil(o.Address) {
		var ret string
		return ret
	}
	return *o.Address
}

// GetAddressOk returns a tuple with the Address field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PxgridEndpoint) GetAddressOk() (*string, bool) {
	if o == nil || IsNil(o.Address) {
		return nil, false
	}
	return o.Address, true
}

// HasAddress returns a boolean if a field has been set.
func (o *PxgridEndpoint) HasAddress() bool {
	if o != nil && !IsNil(o.Address) {
		return true
	}

	return false
}

// SetAddress gets a reference to the given string and assigns it to the Address field.
func (o *PxgridEndpoint) SetAddress(v string) {
	o.Address = &v
}

// GetClientCertificateSubject returns the ClientCertificateSubject field value if set, zero value otherwise.
func (o *PxgridEndpoint) GetClientCertificateSubject() string {
	if o == nil || IsNil(o.ClientCertificateSubject) {
		var ret string
		return ret
	}
	return *o.ClientCertificateSubject
}

// GetClientCertificateSubjectOk returns a tuple with the ClientCertificateSubject field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PxgridEndpoint) GetClientCertificateSubjectOk() (*string, bool) {
	if o == nil || IsNil(o.ClientCertificateSubject) {
		return nil, false
	}
	return o.ClientCertificateSubject, true
}

// HasClientCertificateSubject returns a boolean if a field has been set.
func (o *PxgridEndpoint) HasClientCertificateSubject() bool {
	if o != nil && !IsNil(o.ClientCertificateSubject) {
		return true
	}

	return false
}

// SetClientCertificateSubject gets a reference to the given string and assigns it to the ClientCertificateSubject field.
func (o *PxgridEndpoint) SetClientCertificateSubject(v string) {
	o.ClientCertificateSubject = &v
}

// GetClientCertificateToken returns the ClientCertificateToken field value if set, zero value otherwise.
func (o *PxgridEndpoint) GetClientCertificateToken() string {
	if o == nil || IsNil(o.ClientCertificateToken) {
		var ret string
		return ret
	}
	return *o.ClientCertificateToken
}

// GetClientCertificateTokenOk returns a tuple with the ClientCertificateToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PxgridEndpoint) GetClientCertificateTokenOk() (*string, bool) {
	if o == nil || IsNil(o.ClientCertificateToken) {
		return nil, false
	}
	return o.ClientCertificateToken, true
}

// HasClientCertificateToken returns a boolean if a field has been set.
func (o *PxgridEndpoint) HasClientCertificateToken() bool {
	if o != nil && !IsNil(o.ClientCertificateToken) {
		return true
	}

	return false
}

// SetClientCertificateToken gets a reference to the given string and assigns it to the ClientCertificateToken field.
func (o *PxgridEndpoint) SetClientCertificateToken(v string) {
	o.ClientCertificateToken = &v
}

// GetClientCertificateValidFrom returns the ClientCertificateValidFrom field value if set, zero value otherwise.
func (o *PxgridEndpoint) GetClientCertificateValidFrom() int64 {
	if o == nil || IsNil(o.ClientCertificateValidFrom) {
		var ret int64
		return ret
	}
	return *o.ClientCertificateValidFrom
}

// GetClientCertificateValidFromOk returns a tuple with the ClientCertificateValidFrom field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PxgridEndpoint) GetClientCertificateValidFromOk() (*int64, bool) {
	if o == nil || IsNil(o.ClientCertificateValidFrom) {
		return nil, false
	}
	return o.ClientCertificateValidFrom, true
}

// HasClientCertificateValidFrom returns a boolean if a field has been set.
func (o *PxgridEndpoint) HasClientCertificateValidFrom() bool {
	if o != nil && !IsNil(o.ClientCertificateValidFrom) {
		return true
	}

	return false
}

// SetClientCertificateValidFrom gets a reference to the given int64 and assigns it to the ClientCertificateValidFrom field.
func (o *PxgridEndpoint) SetClientCertificateValidFrom(v int64) {
	o.ClientCertificateValidFrom = &v
}

// GetClientCertificateValidTo returns the ClientCertificateValidTo field value if set, zero value otherwise.
func (o *PxgridEndpoint) GetClientCertificateValidTo() int64 {
	if o == nil || IsNil(o.ClientCertificateValidTo) {
		var ret int64
		return ret
	}
	return *o.ClientCertificateValidTo
}

// GetClientCertificateValidToOk returns a tuple with the ClientCertificateValidTo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PxgridEndpoint) GetClientCertificateValidToOk() (*int64, bool) {
	if o == nil || IsNil(o.ClientCertificateValidTo) {
		return nil, false
	}
	return o.ClientCertificateValidTo, true
}

// HasClientCertificateValidTo returns a boolean if a field has been set.
func (o *PxgridEndpoint) HasClientCertificateValidTo() bool {
	if o != nil && !IsNil(o.ClientCertificateValidTo) {
		return true
	}

	return false
}

// SetClientCertificateValidTo gets a reference to the given int64 and assigns it to the ClientCertificateValidTo field.
func (o *PxgridEndpoint) SetClientCertificateValidTo(v int64) {
	o.ClientCertificateValidTo = &v
}

// GetComment returns the Comment field value if set, zero value otherwise.
func (o *PxgridEndpoint) GetComment() string {
	if o == nil || IsNil(o.Comment) {
		var ret string
		return ret
	}
	return *o.Comment
}

// GetCommentOk returns a tuple with the Comment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PxgridEndpoint) GetCommentOk() (*string, bool) {
	if o == nil || IsNil(o.Comment) {
		return nil, false
	}
	return o.Comment, true
}

// HasComment returns a boolean if a field has been set.
func (o *PxgridEndpoint) HasComment() bool {
	if o != nil && !IsNil(o.Comment) {
		return true
	}

	return false
}

// SetComment gets a reference to the given string and assigns it to the Comment field.
func (o *PxgridEndpoint) SetComment(v string) {
	o.Comment = &v
}

// GetDisable returns the Disable field value if set, zero value otherwise.
func (o *PxgridEndpoint) GetDisable() bool {
	if o == nil || IsNil(o.Disable) {
		var ret bool
		return ret
	}
	return *o.Disable
}

// GetDisableOk returns a tuple with the Disable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PxgridEndpoint) GetDisableOk() (*bool, bool) {
	if o == nil || IsNil(o.Disable) {
		return nil, false
	}
	return o.Disable, true
}

// HasDisable returns a boolean if a field has been set.
func (o *PxgridEndpoint) HasDisable() bool {
	if o != nil && !IsNil(o.Disable) {
		return true
	}

	return false
}

// SetDisable gets a reference to the given bool and assigns it to the Disable field.
func (o *PxgridEndpoint) SetDisable(v bool) {
	o.Disable = &v
}

// GetExtAttrs returns the ExtAttrs field value if set, zero value otherwise.
func (o *PxgridEndpoint) GetExtAttrs() map[string]ExtAttrs {
	if o == nil || IsNil(o.ExtAttrs) {
		var ret map[string]ExtAttrs
		return ret
	}
	return *o.ExtAttrs
}

// GetExtAttrsOk returns a tuple with the ExtAttrs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PxgridEndpoint) GetExtAttrsOk() (*map[string]ExtAttrs, bool) {
	if o == nil || IsNil(o.ExtAttrs) {
		return nil, false
	}
	return o.ExtAttrs, true
}

// HasExtAttrs returns a boolean if a field has been set.
func (o *PxgridEndpoint) HasExtAttrs() bool {
	if o != nil && !IsNil(o.ExtAttrs) {
		return true
	}

	return false
}

// SetExtAttrs gets a reference to the given map[string]ExtAttrs and assigns it to the ExtAttrs field.
func (o *PxgridEndpoint) SetExtAttrs(v map[string]ExtAttrs) {
	o.ExtAttrs = &v
}

// GetLogLevel returns the LogLevel field value if set, zero value otherwise.
func (o *PxgridEndpoint) GetLogLevel() string {
	if o == nil || IsNil(o.LogLevel) {
		var ret string
		return ret
	}
	return *o.LogLevel
}

// GetLogLevelOk returns a tuple with the LogLevel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PxgridEndpoint) GetLogLevelOk() (*string, bool) {
	if o == nil || IsNil(o.LogLevel) {
		return nil, false
	}
	return o.LogLevel, true
}

// HasLogLevel returns a boolean if a field has been set.
func (o *PxgridEndpoint) HasLogLevel() bool {
	if o != nil && !IsNil(o.LogLevel) {
		return true
	}

	return false
}

// SetLogLevel gets a reference to the given string and assigns it to the LogLevel field.
func (o *PxgridEndpoint) SetLogLevel(v string) {
	o.LogLevel = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *PxgridEndpoint) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PxgridEndpoint) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *PxgridEndpoint) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *PxgridEndpoint) SetName(v string) {
	o.Name = &v
}

// GetNetworkView returns the NetworkView field value if set, zero value otherwise.
func (o *PxgridEndpoint) GetNetworkView() string {
	if o == nil || IsNil(o.NetworkView) {
		var ret string
		return ret
	}
	return *o.NetworkView
}

// GetNetworkViewOk returns a tuple with the NetworkView field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PxgridEndpoint) GetNetworkViewOk() (*string, bool) {
	if o == nil || IsNil(o.NetworkView) {
		return nil, false
	}
	return o.NetworkView, true
}

// HasNetworkView returns a boolean if a field has been set.
func (o *PxgridEndpoint) HasNetworkView() bool {
	if o != nil && !IsNil(o.NetworkView) {
		return true
	}

	return false
}

// SetNetworkView gets a reference to the given string and assigns it to the NetworkView field.
func (o *PxgridEndpoint) SetNetworkView(v string) {
	o.NetworkView = &v
}

// GetOutboundMemberType returns the OutboundMemberType field value if set, zero value otherwise.
func (o *PxgridEndpoint) GetOutboundMemberType() string {
	if o == nil || IsNil(o.OutboundMemberType) {
		var ret string
		return ret
	}
	return *o.OutboundMemberType
}

// GetOutboundMemberTypeOk returns a tuple with the OutboundMemberType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PxgridEndpoint) GetOutboundMemberTypeOk() (*string, bool) {
	if o == nil || IsNil(o.OutboundMemberType) {
		return nil, false
	}
	return o.OutboundMemberType, true
}

// HasOutboundMemberType returns a boolean if a field has been set.
func (o *PxgridEndpoint) HasOutboundMemberType() bool {
	if o != nil && !IsNil(o.OutboundMemberType) {
		return true
	}

	return false
}

// SetOutboundMemberType gets a reference to the given string and assigns it to the OutboundMemberType field.
func (o *PxgridEndpoint) SetOutboundMemberType(v string) {
	o.OutboundMemberType = &v
}

// GetOutboundMembers returns the OutboundMembers field value if set, zero value otherwise.
func (o *PxgridEndpoint) GetOutboundMembers() []string {
	if o == nil || IsNil(o.OutboundMembers) {
		var ret []string
		return ret
	}
	return o.OutboundMembers
}

// GetOutboundMembersOk returns a tuple with the OutboundMembers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PxgridEndpoint) GetOutboundMembersOk() ([]string, bool) {
	if o == nil || IsNil(o.OutboundMembers) {
		return nil, false
	}
	return o.OutboundMembers, true
}

// HasOutboundMembers returns a boolean if a field has been set.
func (o *PxgridEndpoint) HasOutboundMembers() bool {
	if o != nil && !IsNil(o.OutboundMembers) {
		return true
	}

	return false
}

// SetOutboundMembers gets a reference to the given []string and assigns it to the OutboundMembers field.
func (o *PxgridEndpoint) SetOutboundMembers(v []string) {
	o.OutboundMembers = v
}

// GetPublishSettings returns the PublishSettings field value if set, zero value otherwise.
func (o *PxgridEndpoint) GetPublishSettings() PxgridEndpointPublishSettings {
	if o == nil || IsNil(o.PublishSettings) {
		var ret PxgridEndpointPublishSettings
		return ret
	}
	return *o.PublishSettings
}

// GetPublishSettingsOk returns a tuple with the PublishSettings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PxgridEndpoint) GetPublishSettingsOk() (*PxgridEndpointPublishSettings, bool) {
	if o == nil || IsNil(o.PublishSettings) {
		return nil, false
	}
	return o.PublishSettings, true
}

// HasPublishSettings returns a boolean if a field has been set.
func (o *PxgridEndpoint) HasPublishSettings() bool {
	if o != nil && !IsNil(o.PublishSettings) {
		return true
	}

	return false
}

// SetPublishSettings gets a reference to the given PxgridEndpointPublishSettings and assigns it to the PublishSettings field.
func (o *PxgridEndpoint) SetPublishSettings(v PxgridEndpointPublishSettings) {
	o.PublishSettings = &v
}

// GetSubscribeSettings returns the SubscribeSettings field value if set, zero value otherwise.
func (o *PxgridEndpoint) GetSubscribeSettings() PxgridEndpointSubscribeSettings {
	if o == nil || IsNil(o.SubscribeSettings) {
		var ret PxgridEndpointSubscribeSettings
		return ret
	}
	return *o.SubscribeSettings
}

// GetSubscribeSettingsOk returns a tuple with the SubscribeSettings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PxgridEndpoint) GetSubscribeSettingsOk() (*PxgridEndpointSubscribeSettings, bool) {
	if o == nil || IsNil(o.SubscribeSettings) {
		return nil, false
	}
	return o.SubscribeSettings, true
}

// HasSubscribeSettings returns a boolean if a field has been set.
func (o *PxgridEndpoint) HasSubscribeSettings() bool {
	if o != nil && !IsNil(o.SubscribeSettings) {
		return true
	}

	return false
}

// SetSubscribeSettings gets a reference to the given PxgridEndpointSubscribeSettings and assigns it to the SubscribeSettings field.
func (o *PxgridEndpoint) SetSubscribeSettings(v PxgridEndpointSubscribeSettings) {
	o.SubscribeSettings = &v
}

// GetTemplateInstance returns the TemplateInstance field value if set, zero value otherwise.
func (o *PxgridEndpoint) GetTemplateInstance() PxgridEndpointTemplateInstance {
	if o == nil || IsNil(o.TemplateInstance) {
		var ret PxgridEndpointTemplateInstance
		return ret
	}
	return *o.TemplateInstance
}

// GetTemplateInstanceOk returns a tuple with the TemplateInstance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PxgridEndpoint) GetTemplateInstanceOk() (*PxgridEndpointTemplateInstance, bool) {
	if o == nil || IsNil(o.TemplateInstance) {
		return nil, false
	}
	return o.TemplateInstance, true
}

// HasTemplateInstance returns a boolean if a field has been set.
func (o *PxgridEndpoint) HasTemplateInstance() bool {
	if o != nil && !IsNil(o.TemplateInstance) {
		return true
	}

	return false
}

// SetTemplateInstance gets a reference to the given PxgridEndpointTemplateInstance and assigns it to the TemplateInstance field.
func (o *PxgridEndpoint) SetTemplateInstance(v PxgridEndpointTemplateInstance) {
	o.TemplateInstance = &v
}

// GetTimeout returns the Timeout field value if set, zero value otherwise.
func (o *PxgridEndpoint) GetTimeout() int64 {
	if o == nil || IsNil(o.Timeout) {
		var ret int64
		return ret
	}
	return *o.Timeout
}

// GetTimeoutOk returns a tuple with the Timeout field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PxgridEndpoint) GetTimeoutOk() (*int64, bool) {
	if o == nil || IsNil(o.Timeout) {
		return nil, false
	}
	return o.Timeout, true
}

// HasTimeout returns a boolean if a field has been set.
func (o *PxgridEndpoint) HasTimeout() bool {
	if o != nil && !IsNil(o.Timeout) {
		return true
	}

	return false
}

// SetTimeout gets a reference to the given int64 and assigns it to the Timeout field.
func (o *PxgridEndpoint) SetTimeout(v int64) {
	o.Timeout = &v
}

// GetVendorIdentifier returns the VendorIdentifier field value if set, zero value otherwise.
func (o *PxgridEndpoint) GetVendorIdentifier() string {
	if o == nil || IsNil(o.VendorIdentifier) {
		var ret string
		return ret
	}
	return *o.VendorIdentifier
}

// GetVendorIdentifierOk returns a tuple with the VendorIdentifier field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PxgridEndpoint) GetVendorIdentifierOk() (*string, bool) {
	if o == nil || IsNil(o.VendorIdentifier) {
		return nil, false
	}
	return o.VendorIdentifier, true
}

// HasVendorIdentifier returns a boolean if a field has been set.
func (o *PxgridEndpoint) HasVendorIdentifier() bool {
	if o != nil && !IsNil(o.VendorIdentifier) {
		return true
	}

	return false
}

// SetVendorIdentifier gets a reference to the given string and assigns it to the VendorIdentifier field.
func (o *PxgridEndpoint) SetVendorIdentifier(v string) {
	o.VendorIdentifier = &v
}

// GetWapiUserName returns the WapiUserName field value if set, zero value otherwise.
func (o *PxgridEndpoint) GetWapiUserName() string {
	if o == nil || IsNil(o.WapiUserName) {
		var ret string
		return ret
	}
	return *o.WapiUserName
}

// GetWapiUserNameOk returns a tuple with the WapiUserName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PxgridEndpoint) GetWapiUserNameOk() (*string, bool) {
	if o == nil || IsNil(o.WapiUserName) {
		return nil, false
	}
	return o.WapiUserName, true
}

// HasWapiUserName returns a boolean if a field has been set.
func (o *PxgridEndpoint) HasWapiUserName() bool {
	if o != nil && !IsNil(o.WapiUserName) {
		return true
	}

	return false
}

// SetWapiUserName gets a reference to the given string and assigns it to the WapiUserName field.
func (o *PxgridEndpoint) SetWapiUserName(v string) {
	o.WapiUserName = &v
}

// GetWapiUserPassword returns the WapiUserPassword field value if set, zero value otherwise.
func (o *PxgridEndpoint) GetWapiUserPassword() string {
	if o == nil || IsNil(o.WapiUserPassword) {
		var ret string
		return ret
	}
	return *o.WapiUserPassword
}

// GetWapiUserPasswordOk returns a tuple with the WapiUserPassword field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PxgridEndpoint) GetWapiUserPasswordOk() (*string, bool) {
	if o == nil || IsNil(o.WapiUserPassword) {
		return nil, false
	}
	return o.WapiUserPassword, true
}

// HasWapiUserPassword returns a boolean if a field has been set.
func (o *PxgridEndpoint) HasWapiUserPassword() bool {
	if o != nil && !IsNil(o.WapiUserPassword) {
		return true
	}

	return false
}

// SetWapiUserPassword gets a reference to the given string and assigns it to the WapiUserPassword field.
func (o *PxgridEndpoint) SetWapiUserPassword(v string) {
	o.WapiUserPassword = &v
}

func (o PxgridEndpoint) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PxgridEndpoint) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Ref) {
		toSerialize["_ref"] = o.Ref
	}
	if !IsNil(o.Address) {
		toSerialize["address"] = o.Address
	}
	if !IsNil(o.ClientCertificateSubject) {
		toSerialize["client_certificate_subject"] = o.ClientCertificateSubject
	}
	if !IsNil(o.ClientCertificateToken) {
		toSerialize["client_certificate_token"] = o.ClientCertificateToken
	}
	if !IsNil(o.ClientCertificateValidFrom) {
		toSerialize["client_certificate_valid_from"] = o.ClientCertificateValidFrom
	}
	if !IsNil(o.ClientCertificateValidTo) {
		toSerialize["client_certificate_valid_to"] = o.ClientCertificateValidTo
	}
	if !IsNil(o.Comment) {
		toSerialize["comment"] = o.Comment
	}
	if !IsNil(o.Disable) {
		toSerialize["disable"] = o.Disable
	}
	if !IsNil(o.ExtAttrs) {
		toSerialize["extattrs"] = o.ExtAttrs
	}
	if !IsNil(o.LogLevel) {
		toSerialize["log_level"] = o.LogLevel
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.NetworkView) {
		toSerialize["network_view"] = o.NetworkView
	}
	if !IsNil(o.OutboundMemberType) {
		toSerialize["outbound_member_type"] = o.OutboundMemberType
	}
	if !IsNil(o.OutboundMembers) {
		toSerialize["outbound_members"] = o.OutboundMembers
	}
	if !IsNil(o.PublishSettings) {
		toSerialize["publish_settings"] = o.PublishSettings
	}
	if !IsNil(o.SubscribeSettings) {
		toSerialize["subscribe_settings"] = o.SubscribeSettings
	}
	if !IsNil(o.TemplateInstance) {
		toSerialize["template_instance"] = o.TemplateInstance
	}
	if !IsNil(o.Timeout) {
		toSerialize["timeout"] = o.Timeout
	}
	if !IsNil(o.VendorIdentifier) {
		toSerialize["vendor_identifier"] = o.VendorIdentifier
	}
	if !IsNil(o.WapiUserName) {
		toSerialize["wapi_user_name"] = o.WapiUserName
	}
	if !IsNil(o.WapiUserPassword) {
		toSerialize["wapi_user_password"] = o.WapiUserPassword
	}
	return toSerialize, nil
}

type NullablePxgridEndpoint struct {
	value *PxgridEndpoint
	isSet bool
}

func (v NullablePxgridEndpoint) Get() *PxgridEndpoint {
	return v.value
}

func (v *NullablePxgridEndpoint) Set(val *PxgridEndpoint) {
	v.value = val
	v.isSet = true
}

func (v NullablePxgridEndpoint) IsSet() bool {
	return v.isSet
}

func (v *NullablePxgridEndpoint) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePxgridEndpoint(val *PxgridEndpoint) *NullablePxgridEndpoint {
	return &NullablePxgridEndpoint{value: val, isSet: true}
}

func (v NullablePxgridEndpoint) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePxgridEndpoint) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
