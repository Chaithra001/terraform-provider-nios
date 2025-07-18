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

// checks if the ZoneAuthDnssecKeyParams type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ZoneAuthDnssecKeyParams{}

// ZoneAuthDnssecKeyParams struct for ZoneAuthDnssecKeyParams
type ZoneAuthDnssecKeyParams struct {
	// If set to True, automatic rollovers for the signing key is enabled.
	EnableKskAutoRollover *bool `json:"enable_ksk_auto_rollover,omitempty"`
	// Key Signing Key algorithm. Deprecated.
	KskAlgorithm *string `json:"ksk_algorithm,omitempty"`
	// A list of Key Signing Key Algorithms.
	KskAlgorithms []ZoneauthdnsseckeyparamsKskAlgorithms `json:"ksk_algorithms,omitempty"`
	// Key Signing Key rollover interval, in seconds.
	KskRollover *int64 `json:"ksk_rollover,omitempty"`
	// Key Signing Key size, in bits. Deprecated.
	KskSize *int64 `json:"ksk_size,omitempty"`
	// NSEC (next secure) types.
	NextSecureType *string `json:"next_secure_type,omitempty"`
	// This field controls events for which users will be notified.
	KskRolloverNotificationConfig *string `json:"ksk_rollover_notification_config,omitempty"`
	// Enable SNMP notifications for KSK related events.
	KskSnmpNotificationEnabled *bool `json:"ksk_snmp_notification_enabled,omitempty"`
	// Enable email notifications for KSK related events.
	KskEmailNotificationEnabled *bool `json:"ksk_email_notification_enabled,omitempty"`
	// The minimum length for NSEC3 salts.
	Nsec3SaltMinLength *int64 `json:"nsec3_salt_min_length,omitempty"`
	// The maximum length for NSEC3 salts.
	Nsec3SaltMaxLength *int64 `json:"nsec3_salt_max_length,omitempty"`
	// The number of iterations used for hashing NSEC3.
	Nsec3Iterations *int64 `json:"nsec3_iterations,omitempty"`
	// Signature expiration time, in seconds.
	SignatureExpiration *int64 `json:"signature_expiration,omitempty"`
	// Zone Signing Key algorithm. Deprecated.
	ZskAlgorithm *string `json:"zsk_algorithm,omitempty"`
	// A list of Zone Signing Key Algorithms.
	ZskAlgorithms []ZoneauthdnsseckeyparamsZskAlgorithms `json:"zsk_algorithms,omitempty"`
	// Zone Signing Key rollover interval, in seconds.
	ZskRollover *int64 `json:"zsk_rollover,omitempty"`
	// Zone Signing Key rollover mechanism.
	ZskRolloverMechanism *string `json:"zsk_rollover_mechanism,omitempty"`
	// Zone Signing Key size, in bits. Deprecated.
	ZskSize              *int64 `json:"zsk_size,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ZoneAuthDnssecKeyParams ZoneAuthDnssecKeyParams

// NewZoneAuthDnssecKeyParams instantiates a new ZoneAuthDnssecKeyParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewZoneAuthDnssecKeyParams() *ZoneAuthDnssecKeyParams {
	this := ZoneAuthDnssecKeyParams{}
	return &this
}

// NewZoneAuthDnssecKeyParamsWithDefaults instantiates a new ZoneAuthDnssecKeyParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewZoneAuthDnssecKeyParamsWithDefaults() *ZoneAuthDnssecKeyParams {
	this := ZoneAuthDnssecKeyParams{}
	return &this
}

// GetEnableKskAutoRollover returns the EnableKskAutoRollover field value if set, zero value otherwise.
func (o *ZoneAuthDnssecKeyParams) GetEnableKskAutoRollover() bool {
	if o == nil || IsNil(o.EnableKskAutoRollover) {
		var ret bool
		return ret
	}
	return *o.EnableKskAutoRollover
}

// GetEnableKskAutoRolloverOk returns a tuple with the EnableKskAutoRollover field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ZoneAuthDnssecKeyParams) GetEnableKskAutoRolloverOk() (*bool, bool) {
	if o == nil || IsNil(o.EnableKskAutoRollover) {
		return nil, false
	}
	return o.EnableKskAutoRollover, true
}

// HasEnableKskAutoRollover returns a boolean if a field has been set.
func (o *ZoneAuthDnssecKeyParams) HasEnableKskAutoRollover() bool {
	if o != nil && !IsNil(o.EnableKskAutoRollover) {
		return true
	}

	return false
}

// SetEnableKskAutoRollover gets a reference to the given bool and assigns it to the EnableKskAutoRollover field.
func (o *ZoneAuthDnssecKeyParams) SetEnableKskAutoRollover(v bool) {
	o.EnableKskAutoRollover = &v
}

// GetKskAlgorithm returns the KskAlgorithm field value if set, zero value otherwise.
func (o *ZoneAuthDnssecKeyParams) GetKskAlgorithm() string {
	if o == nil || IsNil(o.KskAlgorithm) {
		var ret string
		return ret
	}
	return *o.KskAlgorithm
}

// GetKskAlgorithmOk returns a tuple with the KskAlgorithm field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ZoneAuthDnssecKeyParams) GetKskAlgorithmOk() (*string, bool) {
	if o == nil || IsNil(o.KskAlgorithm) {
		return nil, false
	}
	return o.KskAlgorithm, true
}

// HasKskAlgorithm returns a boolean if a field has been set.
func (o *ZoneAuthDnssecKeyParams) HasKskAlgorithm() bool {
	if o != nil && !IsNil(o.KskAlgorithm) {
		return true
	}

	return false
}

// SetKskAlgorithm gets a reference to the given string and assigns it to the KskAlgorithm field.
func (o *ZoneAuthDnssecKeyParams) SetKskAlgorithm(v string) {
	o.KskAlgorithm = &v
}

// GetKskAlgorithms returns the KskAlgorithms field value if set, zero value otherwise.
func (o *ZoneAuthDnssecKeyParams) GetKskAlgorithms() []ZoneauthdnsseckeyparamsKskAlgorithms {
	if o == nil || IsNil(o.KskAlgorithms) {
		var ret []ZoneauthdnsseckeyparamsKskAlgorithms
		return ret
	}
	return o.KskAlgorithms
}

// GetKskAlgorithmsOk returns a tuple with the KskAlgorithms field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ZoneAuthDnssecKeyParams) GetKskAlgorithmsOk() ([]ZoneauthdnsseckeyparamsKskAlgorithms, bool) {
	if o == nil || IsNil(o.KskAlgorithms) {
		return nil, false
	}
	return o.KskAlgorithms, true
}

// HasKskAlgorithms returns a boolean if a field has been set.
func (o *ZoneAuthDnssecKeyParams) HasKskAlgorithms() bool {
	if o != nil && !IsNil(o.KskAlgorithms) {
		return true
	}

	return false
}

// SetKskAlgorithms gets a reference to the given []ZoneauthdnsseckeyparamsKskAlgorithms and assigns it to the KskAlgorithms field.
func (o *ZoneAuthDnssecKeyParams) SetKskAlgorithms(v []ZoneauthdnsseckeyparamsKskAlgorithms) {
	o.KskAlgorithms = v
}

// GetKskRollover returns the KskRollover field value if set, zero value otherwise.
func (o *ZoneAuthDnssecKeyParams) GetKskRollover() int64 {
	if o == nil || IsNil(o.KskRollover) {
		var ret int64
		return ret
	}
	return *o.KskRollover
}

// GetKskRolloverOk returns a tuple with the KskRollover field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ZoneAuthDnssecKeyParams) GetKskRolloverOk() (*int64, bool) {
	if o == nil || IsNil(o.KskRollover) {
		return nil, false
	}
	return o.KskRollover, true
}

// HasKskRollover returns a boolean if a field has been set.
func (o *ZoneAuthDnssecKeyParams) HasKskRollover() bool {
	if o != nil && !IsNil(o.KskRollover) {
		return true
	}

	return false
}

// SetKskRollover gets a reference to the given int64 and assigns it to the KskRollover field.
func (o *ZoneAuthDnssecKeyParams) SetKskRollover(v int64) {
	o.KskRollover = &v
}

// GetKskSize returns the KskSize field value if set, zero value otherwise.
func (o *ZoneAuthDnssecKeyParams) GetKskSize() int64 {
	if o == nil || IsNil(o.KskSize) {
		var ret int64
		return ret
	}
	return *o.KskSize
}

// GetKskSizeOk returns a tuple with the KskSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ZoneAuthDnssecKeyParams) GetKskSizeOk() (*int64, bool) {
	if o == nil || IsNil(o.KskSize) {
		return nil, false
	}
	return o.KskSize, true
}

// HasKskSize returns a boolean if a field has been set.
func (o *ZoneAuthDnssecKeyParams) HasKskSize() bool {
	if o != nil && !IsNil(o.KskSize) {
		return true
	}

	return false
}

// SetKskSize gets a reference to the given int64 and assigns it to the KskSize field.
func (o *ZoneAuthDnssecKeyParams) SetKskSize(v int64) {
	o.KskSize = &v
}

// GetNextSecureType returns the NextSecureType field value if set, zero value otherwise.
func (o *ZoneAuthDnssecKeyParams) GetNextSecureType() string {
	if o == nil || IsNil(o.NextSecureType) {
		var ret string
		return ret
	}
	return *o.NextSecureType
}

// GetNextSecureTypeOk returns a tuple with the NextSecureType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ZoneAuthDnssecKeyParams) GetNextSecureTypeOk() (*string, bool) {
	if o == nil || IsNil(o.NextSecureType) {
		return nil, false
	}
	return o.NextSecureType, true
}

// HasNextSecureType returns a boolean if a field has been set.
func (o *ZoneAuthDnssecKeyParams) HasNextSecureType() bool {
	if o != nil && !IsNil(o.NextSecureType) {
		return true
	}

	return false
}

// SetNextSecureType gets a reference to the given string and assigns it to the NextSecureType field.
func (o *ZoneAuthDnssecKeyParams) SetNextSecureType(v string) {
	o.NextSecureType = &v
}

// GetKskRolloverNotificationConfig returns the KskRolloverNotificationConfig field value if set, zero value otherwise.
func (o *ZoneAuthDnssecKeyParams) GetKskRolloverNotificationConfig() string {
	if o == nil || IsNil(o.KskRolloverNotificationConfig) {
		var ret string
		return ret
	}
	return *o.KskRolloverNotificationConfig
}

// GetKskRolloverNotificationConfigOk returns a tuple with the KskRolloverNotificationConfig field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ZoneAuthDnssecKeyParams) GetKskRolloverNotificationConfigOk() (*string, bool) {
	if o == nil || IsNil(o.KskRolloverNotificationConfig) {
		return nil, false
	}
	return o.KskRolloverNotificationConfig, true
}

// HasKskRolloverNotificationConfig returns a boolean if a field has been set.
func (o *ZoneAuthDnssecKeyParams) HasKskRolloverNotificationConfig() bool {
	if o != nil && !IsNil(o.KskRolloverNotificationConfig) {
		return true
	}

	return false
}

// SetKskRolloverNotificationConfig gets a reference to the given string and assigns it to the KskRolloverNotificationConfig field.
func (o *ZoneAuthDnssecKeyParams) SetKskRolloverNotificationConfig(v string) {
	o.KskRolloverNotificationConfig = &v
}

// GetKskSnmpNotificationEnabled returns the KskSnmpNotificationEnabled field value if set, zero value otherwise.
func (o *ZoneAuthDnssecKeyParams) GetKskSnmpNotificationEnabled() bool {
	if o == nil || IsNil(o.KskSnmpNotificationEnabled) {
		var ret bool
		return ret
	}
	return *o.KskSnmpNotificationEnabled
}

// GetKskSnmpNotificationEnabledOk returns a tuple with the KskSnmpNotificationEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ZoneAuthDnssecKeyParams) GetKskSnmpNotificationEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.KskSnmpNotificationEnabled) {
		return nil, false
	}
	return o.KskSnmpNotificationEnabled, true
}

// HasKskSnmpNotificationEnabled returns a boolean if a field has been set.
func (o *ZoneAuthDnssecKeyParams) HasKskSnmpNotificationEnabled() bool {
	if o != nil && !IsNil(o.KskSnmpNotificationEnabled) {
		return true
	}

	return false
}

// SetKskSnmpNotificationEnabled gets a reference to the given bool and assigns it to the KskSnmpNotificationEnabled field.
func (o *ZoneAuthDnssecKeyParams) SetKskSnmpNotificationEnabled(v bool) {
	o.KskSnmpNotificationEnabled = &v
}

// GetKskEmailNotificationEnabled returns the KskEmailNotificationEnabled field value if set, zero value otherwise.
func (o *ZoneAuthDnssecKeyParams) GetKskEmailNotificationEnabled() bool {
	if o == nil || IsNil(o.KskEmailNotificationEnabled) {
		var ret bool
		return ret
	}
	return *o.KskEmailNotificationEnabled
}

// GetKskEmailNotificationEnabledOk returns a tuple with the KskEmailNotificationEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ZoneAuthDnssecKeyParams) GetKskEmailNotificationEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.KskEmailNotificationEnabled) {
		return nil, false
	}
	return o.KskEmailNotificationEnabled, true
}

// HasKskEmailNotificationEnabled returns a boolean if a field has been set.
func (o *ZoneAuthDnssecKeyParams) HasKskEmailNotificationEnabled() bool {
	if o != nil && !IsNil(o.KskEmailNotificationEnabled) {
		return true
	}

	return false
}

// SetKskEmailNotificationEnabled gets a reference to the given bool and assigns it to the KskEmailNotificationEnabled field.
func (o *ZoneAuthDnssecKeyParams) SetKskEmailNotificationEnabled(v bool) {
	o.KskEmailNotificationEnabled = &v
}

// GetNsec3SaltMinLength returns the Nsec3SaltMinLength field value if set, zero value otherwise.
func (o *ZoneAuthDnssecKeyParams) GetNsec3SaltMinLength() int64 {
	if o == nil || IsNil(o.Nsec3SaltMinLength) {
		var ret int64
		return ret
	}
	return *o.Nsec3SaltMinLength
}

// GetNsec3SaltMinLengthOk returns a tuple with the Nsec3SaltMinLength field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ZoneAuthDnssecKeyParams) GetNsec3SaltMinLengthOk() (*int64, bool) {
	if o == nil || IsNil(o.Nsec3SaltMinLength) {
		return nil, false
	}
	return o.Nsec3SaltMinLength, true
}

// HasNsec3SaltMinLength returns a boolean if a field has been set.
func (o *ZoneAuthDnssecKeyParams) HasNsec3SaltMinLength() bool {
	if o != nil && !IsNil(o.Nsec3SaltMinLength) {
		return true
	}

	return false
}

// SetNsec3SaltMinLength gets a reference to the given int64 and assigns it to the Nsec3SaltMinLength field.
func (o *ZoneAuthDnssecKeyParams) SetNsec3SaltMinLength(v int64) {
	o.Nsec3SaltMinLength = &v
}

// GetNsec3SaltMaxLength returns the Nsec3SaltMaxLength field value if set, zero value otherwise.
func (o *ZoneAuthDnssecKeyParams) GetNsec3SaltMaxLength() int64 {
	if o == nil || IsNil(o.Nsec3SaltMaxLength) {
		var ret int64
		return ret
	}
	return *o.Nsec3SaltMaxLength
}

// GetNsec3SaltMaxLengthOk returns a tuple with the Nsec3SaltMaxLength field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ZoneAuthDnssecKeyParams) GetNsec3SaltMaxLengthOk() (*int64, bool) {
	if o == nil || IsNil(o.Nsec3SaltMaxLength) {
		return nil, false
	}
	return o.Nsec3SaltMaxLength, true
}

// HasNsec3SaltMaxLength returns a boolean if a field has been set.
func (o *ZoneAuthDnssecKeyParams) HasNsec3SaltMaxLength() bool {
	if o != nil && !IsNil(o.Nsec3SaltMaxLength) {
		return true
	}

	return false
}

// SetNsec3SaltMaxLength gets a reference to the given int64 and assigns it to the Nsec3SaltMaxLength field.
func (o *ZoneAuthDnssecKeyParams) SetNsec3SaltMaxLength(v int64) {
	o.Nsec3SaltMaxLength = &v
}

// GetNsec3Iterations returns the Nsec3Iterations field value if set, zero value otherwise.
func (o *ZoneAuthDnssecKeyParams) GetNsec3Iterations() int64 {
	if o == nil || IsNil(o.Nsec3Iterations) {
		var ret int64
		return ret
	}
	return *o.Nsec3Iterations
}

// GetNsec3IterationsOk returns a tuple with the Nsec3Iterations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ZoneAuthDnssecKeyParams) GetNsec3IterationsOk() (*int64, bool) {
	if o == nil || IsNil(o.Nsec3Iterations) {
		return nil, false
	}
	return o.Nsec3Iterations, true
}

// HasNsec3Iterations returns a boolean if a field has been set.
func (o *ZoneAuthDnssecKeyParams) HasNsec3Iterations() bool {
	if o != nil && !IsNil(o.Nsec3Iterations) {
		return true
	}

	return false
}

// SetNsec3Iterations gets a reference to the given int64 and assigns it to the Nsec3Iterations field.
func (o *ZoneAuthDnssecKeyParams) SetNsec3Iterations(v int64) {
	o.Nsec3Iterations = &v
}

// GetSignatureExpiration returns the SignatureExpiration field value if set, zero value otherwise.
func (o *ZoneAuthDnssecKeyParams) GetSignatureExpiration() int64 {
	if o == nil || IsNil(o.SignatureExpiration) {
		var ret int64
		return ret
	}
	return *o.SignatureExpiration
}

// GetSignatureExpirationOk returns a tuple with the SignatureExpiration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ZoneAuthDnssecKeyParams) GetSignatureExpirationOk() (*int64, bool) {
	if o == nil || IsNil(o.SignatureExpiration) {
		return nil, false
	}
	return o.SignatureExpiration, true
}

// HasSignatureExpiration returns a boolean if a field has been set.
func (o *ZoneAuthDnssecKeyParams) HasSignatureExpiration() bool {
	if o != nil && !IsNil(o.SignatureExpiration) {
		return true
	}

	return false
}

// SetSignatureExpiration gets a reference to the given int64 and assigns it to the SignatureExpiration field.
func (o *ZoneAuthDnssecKeyParams) SetSignatureExpiration(v int64) {
	o.SignatureExpiration = &v
}

// GetZskAlgorithm returns the ZskAlgorithm field value if set, zero value otherwise.
func (o *ZoneAuthDnssecKeyParams) GetZskAlgorithm() string {
	if o == nil || IsNil(o.ZskAlgorithm) {
		var ret string
		return ret
	}
	return *o.ZskAlgorithm
}

// GetZskAlgorithmOk returns a tuple with the ZskAlgorithm field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ZoneAuthDnssecKeyParams) GetZskAlgorithmOk() (*string, bool) {
	if o == nil || IsNil(o.ZskAlgorithm) {
		return nil, false
	}
	return o.ZskAlgorithm, true
}

// HasZskAlgorithm returns a boolean if a field has been set.
func (o *ZoneAuthDnssecKeyParams) HasZskAlgorithm() bool {
	if o != nil && !IsNil(o.ZskAlgorithm) {
		return true
	}

	return false
}

// SetZskAlgorithm gets a reference to the given string and assigns it to the ZskAlgorithm field.
func (o *ZoneAuthDnssecKeyParams) SetZskAlgorithm(v string) {
	o.ZskAlgorithm = &v
}

// GetZskAlgorithms returns the ZskAlgorithms field value if set, zero value otherwise.
func (o *ZoneAuthDnssecKeyParams) GetZskAlgorithms() []ZoneauthdnsseckeyparamsZskAlgorithms {
	if o == nil || IsNil(o.ZskAlgorithms) {
		var ret []ZoneauthdnsseckeyparamsZskAlgorithms
		return ret
	}
	return o.ZskAlgorithms
}

// GetZskAlgorithmsOk returns a tuple with the ZskAlgorithms field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ZoneAuthDnssecKeyParams) GetZskAlgorithmsOk() ([]ZoneauthdnsseckeyparamsZskAlgorithms, bool) {
	if o == nil || IsNil(o.ZskAlgorithms) {
		return nil, false
	}
	return o.ZskAlgorithms, true
}

// HasZskAlgorithms returns a boolean if a field has been set.
func (o *ZoneAuthDnssecKeyParams) HasZskAlgorithms() bool {
	if o != nil && !IsNil(o.ZskAlgorithms) {
		return true
	}

	return false
}

// SetZskAlgorithms gets a reference to the given []ZoneauthdnsseckeyparamsZskAlgorithms and assigns it to the ZskAlgorithms field.
func (o *ZoneAuthDnssecKeyParams) SetZskAlgorithms(v []ZoneauthdnsseckeyparamsZskAlgorithms) {
	o.ZskAlgorithms = v
}

// GetZskRollover returns the ZskRollover field value if set, zero value otherwise.
func (o *ZoneAuthDnssecKeyParams) GetZskRollover() int64 {
	if o == nil || IsNil(o.ZskRollover) {
		var ret int64
		return ret
	}
	return *o.ZskRollover
}

// GetZskRolloverOk returns a tuple with the ZskRollover field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ZoneAuthDnssecKeyParams) GetZskRolloverOk() (*int64, bool) {
	if o == nil || IsNil(o.ZskRollover) {
		return nil, false
	}
	return o.ZskRollover, true
}

// HasZskRollover returns a boolean if a field has been set.
func (o *ZoneAuthDnssecKeyParams) HasZskRollover() bool {
	if o != nil && !IsNil(o.ZskRollover) {
		return true
	}

	return false
}

// SetZskRollover gets a reference to the given int64 and assigns it to the ZskRollover field.
func (o *ZoneAuthDnssecKeyParams) SetZskRollover(v int64) {
	o.ZskRollover = &v
}

// GetZskRolloverMechanism returns the ZskRolloverMechanism field value if set, zero value otherwise.
func (o *ZoneAuthDnssecKeyParams) GetZskRolloverMechanism() string {
	if o == nil || IsNil(o.ZskRolloverMechanism) {
		var ret string
		return ret
	}
	return *o.ZskRolloverMechanism
}

// GetZskRolloverMechanismOk returns a tuple with the ZskRolloverMechanism field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ZoneAuthDnssecKeyParams) GetZskRolloverMechanismOk() (*string, bool) {
	if o == nil || IsNil(o.ZskRolloverMechanism) {
		return nil, false
	}
	return o.ZskRolloverMechanism, true
}

// HasZskRolloverMechanism returns a boolean if a field has been set.
func (o *ZoneAuthDnssecKeyParams) HasZskRolloverMechanism() bool {
	if o != nil && !IsNil(o.ZskRolloverMechanism) {
		return true
	}

	return false
}

// SetZskRolloverMechanism gets a reference to the given string and assigns it to the ZskRolloverMechanism field.
func (o *ZoneAuthDnssecKeyParams) SetZskRolloverMechanism(v string) {
	o.ZskRolloverMechanism = &v
}

// GetZskSize returns the ZskSize field value if set, zero value otherwise.
func (o *ZoneAuthDnssecKeyParams) GetZskSize() int64 {
	if o == nil || IsNil(o.ZskSize) {
		var ret int64
		return ret
	}
	return *o.ZskSize
}

// GetZskSizeOk returns a tuple with the ZskSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ZoneAuthDnssecKeyParams) GetZskSizeOk() (*int64, bool) {
	if o == nil || IsNil(o.ZskSize) {
		return nil, false
	}
	return o.ZskSize, true
}

// HasZskSize returns a boolean if a field has been set.
func (o *ZoneAuthDnssecKeyParams) HasZskSize() bool {
	if o != nil && !IsNil(o.ZskSize) {
		return true
	}

	return false
}

// SetZskSize gets a reference to the given int64 and assigns it to the ZskSize field.
func (o *ZoneAuthDnssecKeyParams) SetZskSize(v int64) {
	o.ZskSize = &v
}

func (o ZoneAuthDnssecKeyParams) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ZoneAuthDnssecKeyParams) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.EnableKskAutoRollover) {
		toSerialize["enable_ksk_auto_rollover"] = o.EnableKskAutoRollover
	}
	if !IsNil(o.KskAlgorithm) {
		toSerialize["ksk_algorithm"] = o.KskAlgorithm
	}
	if !IsNil(o.KskAlgorithms) {
		toSerialize["ksk_algorithms"] = o.KskAlgorithms
	}
	if !IsNil(o.KskRollover) {
		toSerialize["ksk_rollover"] = o.KskRollover
	}
	if !IsNil(o.KskSize) {
		toSerialize["ksk_size"] = o.KskSize
	}
	if !IsNil(o.NextSecureType) {
		toSerialize["next_secure_type"] = o.NextSecureType
	}
	if !IsNil(o.KskRolloverNotificationConfig) {
		toSerialize["ksk_rollover_notification_config"] = o.KskRolloverNotificationConfig
	}
	if !IsNil(o.KskSnmpNotificationEnabled) {
		toSerialize["ksk_snmp_notification_enabled"] = o.KskSnmpNotificationEnabled
	}
	if !IsNil(o.KskEmailNotificationEnabled) {
		toSerialize["ksk_email_notification_enabled"] = o.KskEmailNotificationEnabled
	}
	if !IsNil(o.Nsec3SaltMinLength) {
		toSerialize["nsec3_salt_min_length"] = o.Nsec3SaltMinLength
	}
	if !IsNil(o.Nsec3SaltMaxLength) {
		toSerialize["nsec3_salt_max_length"] = o.Nsec3SaltMaxLength
	}
	if !IsNil(o.Nsec3Iterations) {
		toSerialize["nsec3_iterations"] = o.Nsec3Iterations
	}
	if !IsNil(o.SignatureExpiration) {
		toSerialize["signature_expiration"] = o.SignatureExpiration
	}
	if !IsNil(o.ZskAlgorithm) {
		toSerialize["zsk_algorithm"] = o.ZskAlgorithm
	}
	if !IsNil(o.ZskAlgorithms) {
		toSerialize["zsk_algorithms"] = o.ZskAlgorithms
	}
	if !IsNil(o.ZskRollover) {
		toSerialize["zsk_rollover"] = o.ZskRollover
	}
	if !IsNil(o.ZskRolloverMechanism) {
		toSerialize["zsk_rollover_mechanism"] = o.ZskRolloverMechanism
	}
	if !IsNil(o.ZskSize) {
		toSerialize["zsk_size"] = o.ZskSize
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ZoneAuthDnssecKeyParams) UnmarshalJSON(data []byte) (err error) {
	varZoneAuthDnssecKeyParams := _ZoneAuthDnssecKeyParams{}

	err = json.Unmarshal(data, &varZoneAuthDnssecKeyParams)

	if err != nil {
		return err
	}

	*o = ZoneAuthDnssecKeyParams(varZoneAuthDnssecKeyParams)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "enable_ksk_auto_rollover")
		delete(additionalProperties, "ksk_algorithm")
		delete(additionalProperties, "ksk_algorithms")
		delete(additionalProperties, "ksk_rollover")
		delete(additionalProperties, "ksk_size")
		delete(additionalProperties, "next_secure_type")
		delete(additionalProperties, "ksk_rollover_notification_config")
		delete(additionalProperties, "ksk_snmp_notification_enabled")
		delete(additionalProperties, "ksk_email_notification_enabled")
		delete(additionalProperties, "nsec3_salt_min_length")
		delete(additionalProperties, "nsec3_salt_max_length")
		delete(additionalProperties, "nsec3_iterations")
		delete(additionalProperties, "signature_expiration")
		delete(additionalProperties, "zsk_algorithm")
		delete(additionalProperties, "zsk_algorithms")
		delete(additionalProperties, "zsk_rollover")
		delete(additionalProperties, "zsk_rollover_mechanism")
		delete(additionalProperties, "zsk_size")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableZoneAuthDnssecKeyParams struct {
	value *ZoneAuthDnssecKeyParams
	isSet bool
}

func (v NullableZoneAuthDnssecKeyParams) Get() *ZoneAuthDnssecKeyParams {
	return v.value
}

func (v *NullableZoneAuthDnssecKeyParams) Set(val *ZoneAuthDnssecKeyParams) {
	v.value = val
	v.isSet = true
}

func (v NullableZoneAuthDnssecKeyParams) IsSet() bool {
	return v.isSet
}

func (v *NullableZoneAuthDnssecKeyParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableZoneAuthDnssecKeyParams(val *ZoneAuthDnssecKeyParams) *NullableZoneAuthDnssecKeyParams {
	return &NullableZoneAuthDnssecKeyParams{value: val, isSet: true}
}

func (v NullableZoneAuthDnssecKeyParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableZoneAuthDnssecKeyParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
