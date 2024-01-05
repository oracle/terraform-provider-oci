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

// AuthenticationFactorSettingsTotpSettings Settings related to Time-Based One-Time Passcodes (TOTP), such as hashing algo, totp time step, passcode length, and so on
// **SCIM++ Properties:**
//   - idcsSearchable: false
//   - multiValued: false
//   - mutability: readWrite
//   - required: true
//   - returned: default
//   - type: complex
//   - uniqueness: none
type AuthenticationFactorSettingsTotpSettings struct {

	// The hashing algorithm to be used to calculate a One-Time Passcode. By default, the system uses SHA1.
	// **SCIM++ Properties:**
	//  - idcsSearchable: false
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: true
	//  - returned: default
	//  - type: string
	//  - uniqueness: none
	HashingAlgorithm AuthenticationFactorSettingsTotpSettingsHashingAlgorithmEnum `mandatory:"true" json:"hashingAlgorithm"`

	// Exact length of the One-Time Passcode that the system should generate
	// **SCIM++ Properties:**
	//  - idcsMaxValue: 10
	//  - idcsMinValue: 4
	//  - idcsSearchable: false
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: true
	//  - returned: default
	//  - type: integer
	//  - uniqueness: none
	PasscodeLength *int `mandatory:"true" json:"passcodeLength"`

	// The duration of time (in days) after which the shared secret has to be refreshed
	// **SCIM++ Properties:**
	//  - idcsMaxValue: 999
	//  - idcsMinValue: 30
	//  - idcsSearchable: false
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: true
	//  - returned: default
	//  - type: integer
	//  - uniqueness: none
	KeyRefreshIntervalInDays *int `mandatory:"true" json:"keyRefreshIntervalInDays"`

	// Time (in secs) to be used as the time step
	// **SCIM++ Properties:**
	//  - idcsMaxValue: 300
	//  - idcsMinValue: 30
	//  - idcsSearchable: false
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: true
	//  - returned: default
	//  - type: integer
	//  - uniqueness: none
	TimeStepInSecs *int `mandatory:"true" json:"timeStepInSecs"`

	// The tolerance/step-size that the system should use when validating a One-Time Passcode
	// **SCIM++ Properties:**
	//  - idcsMaxValue: 3
	//  - idcsMinValue: 2
	//  - idcsSearchable: false
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: true
	//  - returned: default
	//  - type: integer
	//  - uniqueness: none
	TimeStepTolerance *int `mandatory:"true" json:"timeStepTolerance"`

	// The period of time (in minutes) for which a One-Time Passcode that the system sends by Short Message Service (SMS) or by voice remains valid
	// **SCIM++ Properties:**
	//  - idcsMaxValue: 60
	//  - idcsMinValue: 2
	//  - idcsSearchable: false
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: true
	//  - returned: default
	//  - type: integer
	//  - uniqueness: none
	SmsOtpValidityDurationInMins *int `mandatory:"true" json:"smsOtpValidityDurationInMins"`

	// The period of time (in seconds) that a JSON Web Token (JWT) is valid
	// **SCIM++ Properties:**
	//  - idcsMaxValue: 99999
	//  - idcsMinValue: 30
	//  - idcsSearchable: false
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: true
	//  - returned: default
	//  - type: integer
	//  - uniqueness: none
	JwtValidityDurationInSecs *int `mandatory:"true" json:"jwtValidityDurationInSecs"`

	// Exact length of the Short Message Service (SMS) One-Time Passcode
	// **SCIM++ Properties:**
	//  - idcsMaxValue: 10
	//  - idcsMinValue: 4
	//  - idcsSearchable: false
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: true
	//  - returned: default
	//  - type: integer
	//  - uniqueness: none
	SmsPasscodeLength *int `mandatory:"true" json:"smsPasscodeLength"`

	// The period of time (in minutes) that a one-time passcode remains valid that the system sends by email.
	// **Added In:** 18.1.2
	// **SCIM++ Properties:**
	//  - idcsMaxValue: 60
	//  - idcsMinValue: 2
	//  - idcsSearchable: false
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: true
	//  - returned: default
	//  - type: integer
	//  - uniqueness: none
	EmailOtpValidityDurationInMins *int `mandatory:"true" json:"emailOtpValidityDurationInMins"`

	// Exact length of the email one-time passcode.
	// **Added In:** 18.1.2
	// **SCIM++ Properties:**
	//  - idcsMaxValue: 10
	//  - idcsMinValue: 4
	//  - idcsSearchable: false
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: true
	//  - returned: default
	//  - type: integer
	//  - uniqueness: none
	EmailPasscodeLength *int `mandatory:"true" json:"emailPasscodeLength"`
}

func (m AuthenticationFactorSettingsTotpSettings) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m AuthenticationFactorSettingsTotpSettings) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingAuthenticationFactorSettingsTotpSettingsHashingAlgorithmEnum(string(m.HashingAlgorithm)); !ok && m.HashingAlgorithm != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for HashingAlgorithm: %s. Supported values are: %s.", m.HashingAlgorithm, strings.Join(GetAuthenticationFactorSettingsTotpSettingsHashingAlgorithmEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// AuthenticationFactorSettingsTotpSettingsHashingAlgorithmEnum Enum with underlying type: string
type AuthenticationFactorSettingsTotpSettingsHashingAlgorithmEnum string

// Set of constants representing the allowable values for AuthenticationFactorSettingsTotpSettingsHashingAlgorithmEnum
const (
	AuthenticationFactorSettingsTotpSettingsHashingAlgorithmSha1   AuthenticationFactorSettingsTotpSettingsHashingAlgorithmEnum = "SHA1"
	AuthenticationFactorSettingsTotpSettingsHashingAlgorithmSha256 AuthenticationFactorSettingsTotpSettingsHashingAlgorithmEnum = "SHA256"
	AuthenticationFactorSettingsTotpSettingsHashingAlgorithmSha384 AuthenticationFactorSettingsTotpSettingsHashingAlgorithmEnum = "SHA384"
	AuthenticationFactorSettingsTotpSettingsHashingAlgorithmSha512 AuthenticationFactorSettingsTotpSettingsHashingAlgorithmEnum = "SHA512"
	AuthenticationFactorSettingsTotpSettingsHashingAlgorithmMd5    AuthenticationFactorSettingsTotpSettingsHashingAlgorithmEnum = "MD5"
)

var mappingAuthenticationFactorSettingsTotpSettingsHashingAlgorithmEnum = map[string]AuthenticationFactorSettingsTotpSettingsHashingAlgorithmEnum{
	"SHA1":   AuthenticationFactorSettingsTotpSettingsHashingAlgorithmSha1,
	"SHA256": AuthenticationFactorSettingsTotpSettingsHashingAlgorithmSha256,
	"SHA384": AuthenticationFactorSettingsTotpSettingsHashingAlgorithmSha384,
	"SHA512": AuthenticationFactorSettingsTotpSettingsHashingAlgorithmSha512,
	"MD5":    AuthenticationFactorSettingsTotpSettingsHashingAlgorithmMd5,
}

var mappingAuthenticationFactorSettingsTotpSettingsHashingAlgorithmEnumLowerCase = map[string]AuthenticationFactorSettingsTotpSettingsHashingAlgorithmEnum{
	"sha1":   AuthenticationFactorSettingsTotpSettingsHashingAlgorithmSha1,
	"sha256": AuthenticationFactorSettingsTotpSettingsHashingAlgorithmSha256,
	"sha384": AuthenticationFactorSettingsTotpSettingsHashingAlgorithmSha384,
	"sha512": AuthenticationFactorSettingsTotpSettingsHashingAlgorithmSha512,
	"md5":    AuthenticationFactorSettingsTotpSettingsHashingAlgorithmMd5,
}

// GetAuthenticationFactorSettingsTotpSettingsHashingAlgorithmEnumValues Enumerates the set of values for AuthenticationFactorSettingsTotpSettingsHashingAlgorithmEnum
func GetAuthenticationFactorSettingsTotpSettingsHashingAlgorithmEnumValues() []AuthenticationFactorSettingsTotpSettingsHashingAlgorithmEnum {
	values := make([]AuthenticationFactorSettingsTotpSettingsHashingAlgorithmEnum, 0)
	for _, v := range mappingAuthenticationFactorSettingsTotpSettingsHashingAlgorithmEnum {
		values = append(values, v)
	}
	return values
}

// GetAuthenticationFactorSettingsTotpSettingsHashingAlgorithmEnumStringValues Enumerates the set of values in String for AuthenticationFactorSettingsTotpSettingsHashingAlgorithmEnum
func GetAuthenticationFactorSettingsTotpSettingsHashingAlgorithmEnumStringValues() []string {
	return []string{
		"SHA1",
		"SHA256",
		"SHA384",
		"SHA512",
		"MD5",
	}
}

// GetMappingAuthenticationFactorSettingsTotpSettingsHashingAlgorithmEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingAuthenticationFactorSettingsTotpSettingsHashingAlgorithmEnum(val string) (AuthenticationFactorSettingsTotpSettingsHashingAlgorithmEnum, bool) {
	enum, ok := mappingAuthenticationFactorSettingsTotpSettingsHashingAlgorithmEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
