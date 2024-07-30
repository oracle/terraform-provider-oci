// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Identity Domains API
//
// Use the Identity Domains API to manage resources within an identity domain, for example, users, dynamic resource groups, groups, and identity providers. For information about managing resources within identity domains, see Identity and Access Management (with identity domains) (https://docs.oracle.com/iaas/Content/Identity/home.htm).
// Use this pattern to construct endpoints for identity domains: `https://<domainURL>/admin/v1/`. See Finding an Identity Domain URL (https://docs.oracle.com/en-us/iaas/Content/Identity/api-getstarted/locate-identity-domain-url.htm) to locate the domain URL you need.
// Use the table of contents and search tool to explore the Identity Domains API.
//

package identitydomains

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// ExtensionMfaUser This extension defines attributes used to manage Multi-Factor Authentication within a service provider. The extension is typically applied to a User resource, but MAY be applied to other resources that use MFA.
type ExtensionMfaUser struct {

	// The preferred authentication factor type.
	// **Added In:** 18.3.6
	// **SCIM++ Properties:**
	//  - caseExact: true
	//  - idcsSearchable: true
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: false
	//  - returned: default
	//  - type: string
	//  - uniqueness: none
	PreferredAuthenticationFactor ExtensionMfaUserPreferredAuthenticationFactorEnum `mandatory:"false" json:"preferredAuthenticationFactor,omitempty"`

	// The user opted for MFA.
	// **Added In:** 18.3.6
	// **SCIM++ Properties:**
	//  - caseExact: true
	//  - idcsSearchable: true
	//  - multiValued: false
	//  - mutability: readOnly
	//  - idcsRequiresWriteForAccessFlows: true
	//  - required: false
	//  - returned: default
	//  - type: string
	//  - uniqueness: none
	MfaStatus ExtensionMfaUserMfaStatusEnum `mandatory:"false" json:"mfaStatus,omitempty"`

	// The preferred third-party vendor name.
	// **Added In:** 19.2.1
	// **SCIM++ Properties:**
	//  - caseExact: true
	//  - idcsSearchable: true
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: false
	//  - returned: default
	//  - type: string
	//  - uniqueness: none
	PreferredThirdPartyVendor *string `mandatory:"false" json:"preferredThirdPartyVendor"`

	// The preferred authentication method.
	// **Added In:** 2009232244
	// **SCIM++ Properties:**
	//  - caseExact: true
	//  - idcsSearchable: false
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: false
	//  - returned: default
	//  - type: string
	//  - uniqueness: none
	PreferredAuthenticationMethod *string `mandatory:"false" json:"preferredAuthenticationMethod"`

	// The number of incorrect multi factor authentication sign in attempts made by this user. The user is  locked if this reaches the threshold specified in the maxIncorrectAttempts attribute in AuthenticationFactorSettings.
	// **Added In:** 18.3.6
	// **SCIM++ Properties:**
	//  - idcsSearchable: false
	//  - multiValued: false
	//  - mutability: readOnly
	//  - idcsRequiresWriteForAccessFlows: true
	//  - idcsRequiresImmediateReadAfterWriteForAccessFlows: true
	//  - required: false
	//  - returned: default
	//  - type: integer
	//  - uniqueness: none
	LoginAttempts *int `mandatory:"false" json:"loginAttempts"`

	// The date when the user enrolled in multi factor authentication. This will be set to null, when the user resets their factors.
	// **Added In:** 18.3.6
	// **SCIM++ Properties:**
	//  - idcsSearchable: false
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: false
	//  - returned: request
	//  - type: dateTime
	//  - uniqueness: none
	MfaEnabledOn *string `mandatory:"false" json:"mfaEnabledOn"`

	// User MFA Ignored Apps Identifiers
	// **Added In:** 19.2.1
	// **SCIM++ Properties:**
	//  - caseExact: true
	//  - idcsSearchable: true
	//  - multiValued: true
	//  - mutability: readWrite
	//  - required: false
	//  - returned: default
	//  - type: string
	//  - uniqueness: none
	MfaIgnoredApps []string `mandatory:"false" json:"mfaIgnoredApps"`

	PreferredDevice *UserExtPreferredDevice `mandatory:"false" json:"preferredDevice"`

	// A list of devices enrolled by the user.
	// **Added In:** 18.3.6
	// **SCIM++ Properties:**
	//  - idcsCompositeKey: [value]
	//  - multiValued: true
	//  - mutability: readOnly
	//  - required: false
	//  - returned: request
	//  - type: complex
	//  - uniqueness: none
	Devices []UserExtDevices `mandatory:"false" json:"devices"`

	// A list of bypass codes that belongs to the user.
	// **Added In:** 18.3.6
	// **SCIM++ Properties:**
	//  - idcsCompositeKey: [value]
	//  - multiValued: true
	//  - mutability: readOnly
	//  - required: false
	//  - returned: request
	//  - type: complex
	//  - uniqueness: none
	BypassCodes []UserExtBypassCodes `mandatory:"false" json:"bypassCodes"`

	// A list of trusted User Agents owned by this user. Multi-Factored Authentication uses Trusted User Agents to authenticate users.  A User Agent is software application that a user uses to issue requests. For example, a User Agent could be a particular browser (possibly one of several executing on a desktop or laptop) or a particular mobile application (again, oneof several executing on a particular mobile device). A User Agent is trusted once the Multi-Factor Authentication has verified it in some way.
	// **Added In:** 18.3.6
	// **SCIM++ Properties:**
	//  - idcsCompositeKey: [value]
	//  - multiValued: true
	//  - mutability: readWrite
	//  - required: false
	//  - returned: request
	//  - type: complex
	//  - uniqueness: none
	TrustedUserAgents []UserExtTrustedUserAgents `mandatory:"false" json:"trustedUserAgents"`
}

func (m ExtensionMfaUser) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ExtensionMfaUser) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingExtensionMfaUserPreferredAuthenticationFactorEnum(string(m.PreferredAuthenticationFactor)); !ok && m.PreferredAuthenticationFactor != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for PreferredAuthenticationFactor: %s. Supported values are: %s.", m.PreferredAuthenticationFactor, strings.Join(GetExtensionMfaUserPreferredAuthenticationFactorEnumStringValues(), ",")))
	}
	if _, ok := GetMappingExtensionMfaUserMfaStatusEnum(string(m.MfaStatus)); !ok && m.MfaStatus != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for MfaStatus: %s. Supported values are: %s.", m.MfaStatus, strings.Join(GetExtensionMfaUserMfaStatusEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ExtensionMfaUserPreferredAuthenticationFactorEnum Enum with underlying type: string
type ExtensionMfaUserPreferredAuthenticationFactorEnum string

// Set of constants representing the allowable values for ExtensionMfaUserPreferredAuthenticationFactorEnum
const (
	ExtensionMfaUserPreferredAuthenticationFactorEmail             ExtensionMfaUserPreferredAuthenticationFactorEnum = "EMAIL"
	ExtensionMfaUserPreferredAuthenticationFactorSms               ExtensionMfaUserPreferredAuthenticationFactorEnum = "SMS"
	ExtensionMfaUserPreferredAuthenticationFactorTotp              ExtensionMfaUserPreferredAuthenticationFactorEnum = "TOTP"
	ExtensionMfaUserPreferredAuthenticationFactorPush              ExtensionMfaUserPreferredAuthenticationFactorEnum = "PUSH"
	ExtensionMfaUserPreferredAuthenticationFactorOfflinetotp       ExtensionMfaUserPreferredAuthenticationFactorEnum = "OFFLINETOTP"
	ExtensionMfaUserPreferredAuthenticationFactorUsernamePassword  ExtensionMfaUserPreferredAuthenticationFactorEnum = "USERNAME_PASSWORD"
	ExtensionMfaUserPreferredAuthenticationFactorSecurityQuestions ExtensionMfaUserPreferredAuthenticationFactorEnum = "SECURITY_QUESTIONS"
	ExtensionMfaUserPreferredAuthenticationFactorVoice             ExtensionMfaUserPreferredAuthenticationFactorEnum = "VOICE"
	ExtensionMfaUserPreferredAuthenticationFactorPhoneCall         ExtensionMfaUserPreferredAuthenticationFactorEnum = "PHONE_CALL"
	ExtensionMfaUserPreferredAuthenticationFactorThirdparty        ExtensionMfaUserPreferredAuthenticationFactorEnum = "THIRDPARTY"
	ExtensionMfaUserPreferredAuthenticationFactorFidoAuthenticator ExtensionMfaUserPreferredAuthenticationFactorEnum = "FIDO_AUTHENTICATOR"
	ExtensionMfaUserPreferredAuthenticationFactorYubicoOtp         ExtensionMfaUserPreferredAuthenticationFactorEnum = "YUBICO_OTP"
)

var mappingExtensionMfaUserPreferredAuthenticationFactorEnum = map[string]ExtensionMfaUserPreferredAuthenticationFactorEnum{
	"EMAIL":              ExtensionMfaUserPreferredAuthenticationFactorEmail,
	"SMS":                ExtensionMfaUserPreferredAuthenticationFactorSms,
	"TOTP":               ExtensionMfaUserPreferredAuthenticationFactorTotp,
	"PUSH":               ExtensionMfaUserPreferredAuthenticationFactorPush,
	"OFFLINETOTP":        ExtensionMfaUserPreferredAuthenticationFactorOfflinetotp,
	"USERNAME_PASSWORD":  ExtensionMfaUserPreferredAuthenticationFactorUsernamePassword,
	"SECURITY_QUESTIONS": ExtensionMfaUserPreferredAuthenticationFactorSecurityQuestions,
	"VOICE":              ExtensionMfaUserPreferredAuthenticationFactorVoice,
	"PHONE_CALL":         ExtensionMfaUserPreferredAuthenticationFactorPhoneCall,
	"THIRDPARTY":         ExtensionMfaUserPreferredAuthenticationFactorThirdparty,
	"FIDO_AUTHENTICATOR": ExtensionMfaUserPreferredAuthenticationFactorFidoAuthenticator,
	"YUBICO_OTP":         ExtensionMfaUserPreferredAuthenticationFactorYubicoOtp,
}

var mappingExtensionMfaUserPreferredAuthenticationFactorEnumLowerCase = map[string]ExtensionMfaUserPreferredAuthenticationFactorEnum{
	"email":              ExtensionMfaUserPreferredAuthenticationFactorEmail,
	"sms":                ExtensionMfaUserPreferredAuthenticationFactorSms,
	"totp":               ExtensionMfaUserPreferredAuthenticationFactorTotp,
	"push":               ExtensionMfaUserPreferredAuthenticationFactorPush,
	"offlinetotp":        ExtensionMfaUserPreferredAuthenticationFactorOfflinetotp,
	"username_password":  ExtensionMfaUserPreferredAuthenticationFactorUsernamePassword,
	"security_questions": ExtensionMfaUserPreferredAuthenticationFactorSecurityQuestions,
	"voice":              ExtensionMfaUserPreferredAuthenticationFactorVoice,
	"phone_call":         ExtensionMfaUserPreferredAuthenticationFactorPhoneCall,
	"thirdparty":         ExtensionMfaUserPreferredAuthenticationFactorThirdparty,
	"fido_authenticator": ExtensionMfaUserPreferredAuthenticationFactorFidoAuthenticator,
	"yubico_otp":         ExtensionMfaUserPreferredAuthenticationFactorYubicoOtp,
}

// GetExtensionMfaUserPreferredAuthenticationFactorEnumValues Enumerates the set of values for ExtensionMfaUserPreferredAuthenticationFactorEnum
func GetExtensionMfaUserPreferredAuthenticationFactorEnumValues() []ExtensionMfaUserPreferredAuthenticationFactorEnum {
	values := make([]ExtensionMfaUserPreferredAuthenticationFactorEnum, 0)
	for _, v := range mappingExtensionMfaUserPreferredAuthenticationFactorEnum {
		values = append(values, v)
	}
	return values
}

// GetExtensionMfaUserPreferredAuthenticationFactorEnumStringValues Enumerates the set of values in String for ExtensionMfaUserPreferredAuthenticationFactorEnum
func GetExtensionMfaUserPreferredAuthenticationFactorEnumStringValues() []string {
	return []string{
		"EMAIL",
		"SMS",
		"TOTP",
		"PUSH",
		"OFFLINETOTP",
		"USERNAME_PASSWORD",
		"SECURITY_QUESTIONS",
		"VOICE",
		"PHONE_CALL",
		"THIRDPARTY",
		"FIDO_AUTHENTICATOR",
		"YUBICO_OTP",
	}
}

// GetMappingExtensionMfaUserPreferredAuthenticationFactorEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingExtensionMfaUserPreferredAuthenticationFactorEnum(val string) (ExtensionMfaUserPreferredAuthenticationFactorEnum, bool) {
	enum, ok := mappingExtensionMfaUserPreferredAuthenticationFactorEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ExtensionMfaUserMfaStatusEnum Enum with underlying type: string
type ExtensionMfaUserMfaStatusEnum string

// Set of constants representing the allowable values for ExtensionMfaUserMfaStatusEnum
const (
	ExtensionMfaUserMfaStatusEnrolled   ExtensionMfaUserMfaStatusEnum = "ENROLLED"
	ExtensionMfaUserMfaStatusIgnored    ExtensionMfaUserMfaStatusEnum = "IGNORED"
	ExtensionMfaUserMfaStatusUnEnrolled ExtensionMfaUserMfaStatusEnum = "UN_ENROLLED"
	ExtensionMfaUserMfaStatusDisabled   ExtensionMfaUserMfaStatusEnum = "DISABLED"
)

var mappingExtensionMfaUserMfaStatusEnum = map[string]ExtensionMfaUserMfaStatusEnum{
	"ENROLLED":    ExtensionMfaUserMfaStatusEnrolled,
	"IGNORED":     ExtensionMfaUserMfaStatusIgnored,
	"UN_ENROLLED": ExtensionMfaUserMfaStatusUnEnrolled,
	"DISABLED":    ExtensionMfaUserMfaStatusDisabled,
}

var mappingExtensionMfaUserMfaStatusEnumLowerCase = map[string]ExtensionMfaUserMfaStatusEnum{
	"enrolled":    ExtensionMfaUserMfaStatusEnrolled,
	"ignored":     ExtensionMfaUserMfaStatusIgnored,
	"un_enrolled": ExtensionMfaUserMfaStatusUnEnrolled,
	"disabled":    ExtensionMfaUserMfaStatusDisabled,
}

// GetExtensionMfaUserMfaStatusEnumValues Enumerates the set of values for ExtensionMfaUserMfaStatusEnum
func GetExtensionMfaUserMfaStatusEnumValues() []ExtensionMfaUserMfaStatusEnum {
	values := make([]ExtensionMfaUserMfaStatusEnum, 0)
	for _, v := range mappingExtensionMfaUserMfaStatusEnum {
		values = append(values, v)
	}
	return values
}

// GetExtensionMfaUserMfaStatusEnumStringValues Enumerates the set of values in String for ExtensionMfaUserMfaStatusEnum
func GetExtensionMfaUserMfaStatusEnumStringValues() []string {
	return []string{
		"ENROLLED",
		"IGNORED",
		"UN_ENROLLED",
		"DISABLED",
	}
}

// GetMappingExtensionMfaUserMfaStatusEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingExtensionMfaUserMfaStatusEnum(val string) (ExtensionMfaUserMfaStatusEnum, bool) {
	enum, ok := mappingExtensionMfaUserMfaStatusEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
