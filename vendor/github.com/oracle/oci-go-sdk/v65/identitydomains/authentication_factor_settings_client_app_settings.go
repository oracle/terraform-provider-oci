// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Identity Domains API
//
// Use the Identity Domains API to manage resources within an identity domain, for example, users, dynamic resource groups, groups, and identity providers. For information about managing resources within identity domains, see Identity and Access Management (with identity domains) (https://docs.oracle.com/iaas/Content/Identity/home.htm). This REST API is SCIM compliant.
// Use the table of contents and search tool to explore the Identity Domains API.
//

package identitydomains

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// AuthenticationFactorSettingsClientAppSettings Settings related to compliance, Personal Identification Number (PIN) policy, and so on
// **SCIM++ Properties:**
//   - idcsSearchable: false
//   - multiValued: false
//   - mutability: readWrite
//   - required: true
//   - returned: default
//   - type: complex
//   - uniqueness: none
type AuthenticationFactorSettingsClientAppSettings struct {

	// Minimum length of the Personal Identification Number (PIN)
	// **SCIM++ Properties:**
	//  - idcsMaxValue: 10
	//  - idcsMinValue: 6
	//  - idcsSearchable: false
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: true
	//  - returned: default
	//  - type: integer
	//  - uniqueness: none
	MinPinLength *int `mandatory:"true" json:"minPinLength"`

	// The maximum number of login failures that the system will allow before raising a warning and sending an alert via email
	// **SCIM++ Properties:**
	//  - idcsMaxValue: 10
	//  - idcsMinValue: 0
	//  - idcsSearchable: false
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: true
	//  - returned: default
	//  - type: integer
	//  - uniqueness: none
	MaxFailuresBeforeWarning *int `mandatory:"true" json:"maxFailuresBeforeWarning"`

	// The maximum number of times that a particular user can fail to login before the system locks that user out of the service
	// **SCIM++ Properties:**
	//  - idcsMaxValue: 10
	//  - idcsMinValue: 5
	//  - idcsSearchable: false
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: true
	//  - returned: default
	//  - type: integer
	//  - uniqueness: none
	MaxFailuresBeforeLockout *int `mandatory:"true" json:"maxFailuresBeforeLockout"`

	// The period of time in seconds that the system will lock a user out of the service after that user exceeds the maximum number of login failures
	// **SCIM++ Properties:**
	//  - idcsMaxValue: 86400
	//  - idcsMinValue: 30
	//  - idcsSearchable: false
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: true
	//  - returned: default
	//  - type: integer
	//  - uniqueness: none
	InitialLockoutPeriodInSecs *int `mandatory:"true" json:"initialLockoutPeriodInSecs"`

	// The pattern of escalation that the system follows, in locking a particular user out of the service.
	// **SCIM++ Properties:**
	//  - idcsCanonicalValueSourceFilter: attrName eq "lockoutEscalationPattern" and attrValues.value eq "$(lockoutEscalationPattern)"
	//  - idcsCanonicalValueSourceResourceType: AllowedValue
	//  - idcsSearchable: false
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: true
	//  - returned: default
	//  - type: string
	//  - uniqueness: none
	LockoutEscalationPattern *string `mandatory:"true" json:"lockoutEscalationPattern"`

	// The maximum period of time that the system will lock a particular user out of the service regardless of what the configured pattern of escalation would otherwise dictate
	// **SCIM++ Properties:**
	//  - idcsMaxValue: 86400
	//  - idcsMinValue: 30
	//  - idcsSearchable: false
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: true
	//  - returned: default
	//  - type: integer
	//  - uniqueness: none
	MaxLockoutIntervalInSecs *int `mandatory:"true" json:"maxLockoutIntervalInSecs"`

	// Indicates which algorithm the system will use to sign requests
	// **SCIM++ Properties:**
	//  - idcsSearchable: false
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: true
	//  - returned: default
	//  - type: string
	//  - uniqueness: none
	RequestSigningAlgo AuthenticationFactorSettingsClientAppSettingsRequestSigningAlgoEnum `mandatory:"true" json:"requestSigningAlgo"`

	// The period of time in days after which a client should refresh its policy by re-reading that policy from the server
	// **SCIM++ Properties:**
	//  - idcsMaxValue: 999
	//  - idcsMinValue: 1
	//  - idcsSearchable: false
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: true
	//  - returned: default
	//  - type: integer
	//  - uniqueness: none
	PolicyUpdateFreqInDays *int `mandatory:"true" json:"policyUpdateFreqInDays"`

	// The size of the key that the system uses to generate the public-private key pair
	// **SCIM++ Properties:**
	//  - idcsMaxValue: 4000
	//  - idcsMinValue: 32
	//  - idcsSearchable: false
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: true
	//  - returned: default
	//  - type: integer
	//  - uniqueness: none
	KeyPairLength *int `mandatory:"true" json:"keyPairLength"`

	// Indicates what protection policy that the system applies on a device. By default, the value is NONE, which indicates that the system applies no protection policy. A value of APP_PIN indicates that the system requires a Personal Identification Number (PIN). A value of DEVICE_BIOMETRIC_OR_APP_PIN indicates that either a PIN or a biometric authentication factor is required.
	// **SCIM++ Properties:**
	//  - idcsCanonicalValueSourceFilter: attrName eq "deviceProtectionPolicy" and attrValues.value eq "$(deviceProtectionPolicy)"
	//  - idcsCanonicalValueSourceResourceType: AllowedValue
	//  - idcsSearchable: false
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: true
	//  - returned: default
	//  - type: string
	//  - uniqueness: none
	DeviceProtectionPolicy *string `mandatory:"true" json:"deviceProtectionPolicy"`

	// If true, indicates that the system should require the user to unlock the client app for each request. In order to unlock the App, the user must supply a Personal Identification Number (PIN) or a biometric authentication-factor.
	// **SCIM++ Properties:**
	//  - idcsSearchable: false
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: true
	//  - returned: default
	//  - type: boolean
	//  - uniqueness: none
	UnlockAppForEachRequestEnabled *bool `mandatory:"true" json:"unlockAppForEachRequestEnabled"`

	// If true, indicates that the system should require the user to unlock the client App whenever the App is started. In order to unlock the App, the user must supply a Personal Identification Number (PIN) or a biometric authentication-factor.
	// **SCIM++ Properties:**
	//  - idcsSearchable: false
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: true
	//  - returned: default
	//  - type: boolean
	//  - uniqueness: none
	UnlockOnAppStartEnabled *bool `mandatory:"true" json:"unlockOnAppStartEnabled"`

	// Specifies the period of time in seconds after which the client App should require the user to unlock the App. In order to unlock the App, the user must supply a Personal Identification Number (PIN) or a biometric authentication-factor. A value of zero means that it is disabled.
	// **SCIM++ Properties:**
	//  - idcsMaxValue: 9999999
	//  - idcsMinValue: 0
	//  - idcsSearchable: false
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: true
	//  - returned: default
	//  - type: integer
	//  - uniqueness: none
	UnlockAppIntervalInSecs *int `mandatory:"true" json:"unlockAppIntervalInSecs"`

	// Indicates the type of encoding that the system should use to generate a shared secret
	// **SCIM++ Properties:**
	//  - idcsSearchable: false
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: true
	//  - returned: default
	//  - type: string
	//  - uniqueness: none
	SharedSecretEncoding AuthenticationFactorSettingsClientAppSettingsSharedSecretEncodingEnum `mandatory:"true" json:"sharedSecretEncoding"`

	// If true, indicates that the system should require the user to unlock the client App, when the client App comes to the foreground in the display of the device. In order to unlock the App, the user must supply a Personal Identification Number (PIN) or a biometric authentication-factor.
	// **SCIM++ Properties:**
	//  - idcsSearchable: false
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: true
	//  - returned: default
	//  - type: boolean
	//  - uniqueness: none
	UnlockOnAppForegroundEnabled *bool `mandatory:"true" json:"unlockOnAppForegroundEnabled"`
}

func (m AuthenticationFactorSettingsClientAppSettings) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m AuthenticationFactorSettingsClientAppSettings) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingAuthenticationFactorSettingsClientAppSettingsRequestSigningAlgoEnum(string(m.RequestSigningAlgo)); !ok && m.RequestSigningAlgo != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for RequestSigningAlgo: %s. Supported values are: %s.", m.RequestSigningAlgo, strings.Join(GetAuthenticationFactorSettingsClientAppSettingsRequestSigningAlgoEnumStringValues(), ",")))
	}
	if _, ok := GetMappingAuthenticationFactorSettingsClientAppSettingsSharedSecretEncodingEnum(string(m.SharedSecretEncoding)); !ok && m.SharedSecretEncoding != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SharedSecretEncoding: %s. Supported values are: %s.", m.SharedSecretEncoding, strings.Join(GetAuthenticationFactorSettingsClientAppSettingsSharedSecretEncodingEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// AuthenticationFactorSettingsClientAppSettingsRequestSigningAlgoEnum Enum with underlying type: string
type AuthenticationFactorSettingsClientAppSettingsRequestSigningAlgoEnum string

// Set of constants representing the allowable values for AuthenticationFactorSettingsClientAppSettingsRequestSigningAlgoEnum
const (
	AuthenticationFactorSettingsClientAppSettingsRequestSigningAlgoSha256withrsa AuthenticationFactorSettingsClientAppSettingsRequestSigningAlgoEnum = "SHA256withRSA"
	AuthenticationFactorSettingsClientAppSettingsRequestSigningAlgoSha384withrsa AuthenticationFactorSettingsClientAppSettingsRequestSigningAlgoEnum = "SHA384withRSA"
	AuthenticationFactorSettingsClientAppSettingsRequestSigningAlgoSha512withrsa AuthenticationFactorSettingsClientAppSettingsRequestSigningAlgoEnum = "SHA512withRSA"
)

var mappingAuthenticationFactorSettingsClientAppSettingsRequestSigningAlgoEnum = map[string]AuthenticationFactorSettingsClientAppSettingsRequestSigningAlgoEnum{
	"SHA256withRSA": AuthenticationFactorSettingsClientAppSettingsRequestSigningAlgoSha256withrsa,
	"SHA384withRSA": AuthenticationFactorSettingsClientAppSettingsRequestSigningAlgoSha384withrsa,
	"SHA512withRSA": AuthenticationFactorSettingsClientAppSettingsRequestSigningAlgoSha512withrsa,
}

var mappingAuthenticationFactorSettingsClientAppSettingsRequestSigningAlgoEnumLowerCase = map[string]AuthenticationFactorSettingsClientAppSettingsRequestSigningAlgoEnum{
	"sha256withrsa": AuthenticationFactorSettingsClientAppSettingsRequestSigningAlgoSha256withrsa,
	"sha384withrsa": AuthenticationFactorSettingsClientAppSettingsRequestSigningAlgoSha384withrsa,
	"sha512withrsa": AuthenticationFactorSettingsClientAppSettingsRequestSigningAlgoSha512withrsa,
}

// GetAuthenticationFactorSettingsClientAppSettingsRequestSigningAlgoEnumValues Enumerates the set of values for AuthenticationFactorSettingsClientAppSettingsRequestSigningAlgoEnum
func GetAuthenticationFactorSettingsClientAppSettingsRequestSigningAlgoEnumValues() []AuthenticationFactorSettingsClientAppSettingsRequestSigningAlgoEnum {
	values := make([]AuthenticationFactorSettingsClientAppSettingsRequestSigningAlgoEnum, 0)
	for _, v := range mappingAuthenticationFactorSettingsClientAppSettingsRequestSigningAlgoEnum {
		values = append(values, v)
	}
	return values
}

// GetAuthenticationFactorSettingsClientAppSettingsRequestSigningAlgoEnumStringValues Enumerates the set of values in String for AuthenticationFactorSettingsClientAppSettingsRequestSigningAlgoEnum
func GetAuthenticationFactorSettingsClientAppSettingsRequestSigningAlgoEnumStringValues() []string {
	return []string{
		"SHA256withRSA",
		"SHA384withRSA",
		"SHA512withRSA",
	}
}

// GetMappingAuthenticationFactorSettingsClientAppSettingsRequestSigningAlgoEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingAuthenticationFactorSettingsClientAppSettingsRequestSigningAlgoEnum(val string) (AuthenticationFactorSettingsClientAppSettingsRequestSigningAlgoEnum, bool) {
	enum, ok := mappingAuthenticationFactorSettingsClientAppSettingsRequestSigningAlgoEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// AuthenticationFactorSettingsClientAppSettingsSharedSecretEncodingEnum Enum with underlying type: string
type AuthenticationFactorSettingsClientAppSettingsSharedSecretEncodingEnum string

// Set of constants representing the allowable values for AuthenticationFactorSettingsClientAppSettingsSharedSecretEncodingEnum
const (
	AuthenticationFactorSettingsClientAppSettingsSharedSecretEncodingBase32 AuthenticationFactorSettingsClientAppSettingsSharedSecretEncodingEnum = "Base32"
	AuthenticationFactorSettingsClientAppSettingsSharedSecretEncodingBase64 AuthenticationFactorSettingsClientAppSettingsSharedSecretEncodingEnum = "Base64"
)

var mappingAuthenticationFactorSettingsClientAppSettingsSharedSecretEncodingEnum = map[string]AuthenticationFactorSettingsClientAppSettingsSharedSecretEncodingEnum{
	"Base32": AuthenticationFactorSettingsClientAppSettingsSharedSecretEncodingBase32,
	"Base64": AuthenticationFactorSettingsClientAppSettingsSharedSecretEncodingBase64,
}

var mappingAuthenticationFactorSettingsClientAppSettingsSharedSecretEncodingEnumLowerCase = map[string]AuthenticationFactorSettingsClientAppSettingsSharedSecretEncodingEnum{
	"base32": AuthenticationFactorSettingsClientAppSettingsSharedSecretEncodingBase32,
	"base64": AuthenticationFactorSettingsClientAppSettingsSharedSecretEncodingBase64,
}

// GetAuthenticationFactorSettingsClientAppSettingsSharedSecretEncodingEnumValues Enumerates the set of values for AuthenticationFactorSettingsClientAppSettingsSharedSecretEncodingEnum
func GetAuthenticationFactorSettingsClientAppSettingsSharedSecretEncodingEnumValues() []AuthenticationFactorSettingsClientAppSettingsSharedSecretEncodingEnum {
	values := make([]AuthenticationFactorSettingsClientAppSettingsSharedSecretEncodingEnum, 0)
	for _, v := range mappingAuthenticationFactorSettingsClientAppSettingsSharedSecretEncodingEnum {
		values = append(values, v)
	}
	return values
}

// GetAuthenticationFactorSettingsClientAppSettingsSharedSecretEncodingEnumStringValues Enumerates the set of values in String for AuthenticationFactorSettingsClientAppSettingsSharedSecretEncodingEnum
func GetAuthenticationFactorSettingsClientAppSettingsSharedSecretEncodingEnumStringValues() []string {
	return []string{
		"Base32",
		"Base64",
	}
}

// GetMappingAuthenticationFactorSettingsClientAppSettingsSharedSecretEncodingEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingAuthenticationFactorSettingsClientAppSettingsSharedSecretEncodingEnum(val string) (AuthenticationFactorSettingsClientAppSettingsSharedSecretEncodingEnum, bool) {
	enum, ok := mappingAuthenticationFactorSettingsClientAppSettingsSharedSecretEncodingEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
