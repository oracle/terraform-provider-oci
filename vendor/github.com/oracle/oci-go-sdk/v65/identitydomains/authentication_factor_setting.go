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

// AuthenticationFactorSetting Multi Factor Authentication Settings for Tenant
type AuthenticationFactorSetting struct {

	// REQUIRED. The schemas attribute is an array of Strings which allows introspection of the supported schema version for a SCIM representation as well any schema extensions supported by that representation. Each String value must be a unique URI. This specification defines URIs for User, Group, and a standard \"enterprise\" extension. All representations of SCIM schema MUST include a non-zero value array with value(s) of the URIs supported by that representation. Duplicate values MUST NOT be included. Value order is not specified and MUST not impact behavior.
	// **SCIM++ Properties:**
	//  - caseExact: false
	//  - idcsSearchable: false
	//  - multiValued: true
	//  - mutability: readWrite
	//  - required: true
	//  - returned: default
	//  - type: string
	//  - uniqueness: none
	Schemas []string `mandatory:"true" json:"schemas"`

	// If true, indicates that the Short Message Service (SMS) channel is enabled for authentication
	// **SCIM++ Properties:**
	//  - idcsSearchable: false
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: true
	//  - returned: default
	//  - type: boolean
	//  - uniqueness: none
	SmsEnabled *bool `mandatory:"true" json:"smsEnabled"`

	// If true, indicates that the Mobile App One Time Passcode channel is enabled for authentication
	// **SCIM++ Properties:**
	//  - idcsSearchable: false
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: true
	//  - returned: default
	//  - type: boolean
	//  - uniqueness: none
	TotpEnabled *bool `mandatory:"true" json:"totpEnabled"`

	// If true, indicates that the Mobile App Push Notification channel is enabled for authentication
	// **SCIM++ Properties:**
	//  - idcsSearchable: false
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: true
	//  - returned: default
	//  - type: boolean
	//  - uniqueness: none
	PushEnabled *bool `mandatory:"true" json:"pushEnabled"`

	// If true, indicates that Bypass Code is enabled for authentication
	// **SCIM++ Properties:**
	//  - idcsSearchable: false
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: true
	//  - returned: default
	//  - type: boolean
	//  - uniqueness: none
	BypassCodeEnabled *bool `mandatory:"true" json:"bypassCodeEnabled"`

	// If true, indicates that Security Questions are enabled for authentication
	// **SCIM++ Properties:**
	//  - idcsSearchable: false
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: true
	//  - returned: default
	//  - type: boolean
	//  - uniqueness: none
	SecurityQuestionsEnabled *bool `mandatory:"true" json:"securityQuestionsEnabled"`

	// Specifies if Multi-Factor Authentication enrollment is mandatory or optional for a user
	// **Deprecated Since: 18.1.2**
	// **SCIM++ Properties:**
	//  - idcsCanonicalValueSourceFilter: attrName eq "mfaEnrollmentType" and attrValues.value eq "$(mfaEnrollmentType)"
	//  - idcsCanonicalValueSourceResourceType: AllowedValue
	//  - idcsSearchable: false
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: true
	//  - returned: default
	//  - type: string
	//  - uniqueness: none
	MfaEnrollmentType *string `mandatory:"true" json:"mfaEnrollmentType"`

	NotificationSettings *AuthenticationFactorSettingsNotificationSettings `mandatory:"true" json:"notificationSettings"`

	BypassCodeSettings *AuthenticationFactorSettingsBypassCodeSettings `mandatory:"true" json:"bypassCodeSettings"`

	ClientAppSettings *AuthenticationFactorSettingsClientAppSettings `mandatory:"true" json:"clientAppSettings"`

	EndpointRestrictions *AuthenticationFactorSettingsEndpointRestrictions `mandatory:"true" json:"endpointRestrictions"`

	// Compliance Policy that defines actions to be taken when a condition is violated
	// **SCIM++ Properties:**
	//  - idcsCompositeKey: [name]
	//  - idcsSearchable: false
	//  - multiValued: true
	//  - mutability: readWrite
	//  - required: true
	//  - returned: default
	//  - type: complex
	//  - uniqueness: none
	CompliancePolicy []AuthenticationFactorSettingsCompliancePolicy `mandatory:"true" json:"compliancePolicy"`

	TotpSettings *AuthenticationFactorSettingsTotpSettings `mandatory:"true" json:"totpSettings"`

	// Unique identifier for the SCIM Resource as defined by the Service Provider. Each representation of the Resource MUST include a non-empty id value. This identifier MUST be unique across the Service Provider's entire set of Resources. It MUST be a stable, non-reassignable identifier that does not change when the same Resource is returned in subsequent requests. The value of the id attribute is always issued by the Service Provider and MUST never be specified by the Service Consumer. bulkId: is a reserved keyword and MUST NOT be used in the unique identifier.
	// **SCIM++ Properties:**
	//  - caseExact: false
	//  - idcsSearchable: true
	//  - multiValued: false
	//  - mutability: readOnly
	//  - required: false
	//  - returned: always
	//  - type: string
	//  - uniqueness: global
	Id *string `mandatory:"false" json:"id"`

	// Unique OCI identifier for the SCIM Resource.
	// **SCIM++ Properties:**
	//  - caseExact: true
	//  - idcsSearchable: true
	//  - multiValued: false
	//  - mutability: immutable
	//  - required: false
	//  - returned: default
	//  - type: string
	//  - uniqueness: global
	Ocid *string `mandatory:"false" json:"ocid"`

	Meta *Meta `mandatory:"false" json:"meta"`

	IdcsCreatedBy *IdcsCreatedBy `mandatory:"false" json:"idcsCreatedBy"`

	IdcsLastModifiedBy *IdcsLastModifiedBy `mandatory:"false" json:"idcsLastModifiedBy"`

	// Each value of this attribute specifies an operation that only an internal client may perform on this particular resource.
	// **SCIM++ Properties:**
	//  - idcsSearchable: false
	//  - multiValued: true
	//  - mutability: readOnly
	//  - required: false
	//  - returned: request
	//  - type: string
	//  - uniqueness: none
	IdcsPreventedOperations []IdcsPreventedOperationsEnum `mandatory:"false" json:"idcsPreventedOperations,omitempty"`

	// A list of tags on this resource.
	// **SCIM++ Properties:**
	//  - idcsCompositeKey: [key, value]
	//  - idcsSearchable: true
	//  - multiValued: true
	//  - mutability: readWrite
	//  - required: false
	//  - returned: request
	//  - type: complex
	//  - uniqueness: none
	Tags []Tags `mandatory:"false" json:"tags"`

	// A boolean flag indicating this resource in the process of being deleted. Usually set to true when synchronous deletion of the resource would take too long.
	// **SCIM++ Properties:**
	//  - caseExact: false
	//  - idcsSearchable: true
	//  - multiValued: false
	//  - mutability: readOnly
	//  - required: false
	//  - returned: default
	//  - type: boolean
	//  - uniqueness: none
	DeleteInProgress *bool `mandatory:"false" json:"deleteInProgress"`

	// The release number when the resource was upgraded.
	// **SCIM++ Properties:**
	//  - caseExact: false
	//  - idcsSearchable: false
	//  - multiValued: false
	//  - mutability: readOnly
	//  - required: false
	//  - returned: request
	//  - type: string
	//  - uniqueness: none
	IdcsLastUpgradedInRelease *string `mandatory:"false" json:"idcsLastUpgradedInRelease"`

	// OCI Domain Id (ocid) in which the resource lives.
	// **SCIM++ Properties:**
	//  - caseExact: false
	//  - idcsSearchable: false
	//  - multiValued: false
	//  - mutability: readOnly
	//  - required: false
	//  - returned: default
	//  - type: string
	//  - uniqueness: none
	DomainOcid *string `mandatory:"false" json:"domainOcid"`

	// OCI Compartment Id (ocid) in which the resource lives.
	// **SCIM++ Properties:**
	//  - caseExact: false
	//  - idcsSearchable: false
	//  - multiValued: false
	//  - mutability: readOnly
	//  - required: false
	//  - returned: default
	//  - type: string
	//  - uniqueness: none
	CompartmentOcid *string `mandatory:"false" json:"compartmentOcid"`

	// OCI Tenant Id (ocid) in which the resource lives.
	// **SCIM++ Properties:**
	//  - caseExact: false
	//  - idcsSearchable: false
	//  - multiValued: false
	//  - mutability: readOnly
	//  - required: false
	//  - returned: default
	//  - type: string
	//  - uniqueness: none
	TenancyOcid *string `mandatory:"false" json:"tenancyOcid"`

	// If true, indicates that the EMAIL channel is enabled for authentication
	// **Added In:** 18.1.2
	// **SCIM++ Properties:**
	//  - idcsSearchable: false
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: false
	//  - returned: default
	//  - type: boolean
	//  - uniqueness: none
	EmailEnabled *bool `mandatory:"false" json:"emailEnabled"`

	// If true, indicates that the phone (PHONE_CALL) channel is enabled for authentication
	// **Added In:** 20.1.3
	// **SCIM++ Properties:**
	//  - idcsSearchable: false
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: false
	//  - returned: default
	//  - type: boolean
	//  - uniqueness: none
	PhoneCallEnabled *bool `mandatory:"false" json:"phoneCallEnabled"`

	// If true, indicates that the Fido Authenticator channels are enabled for authentication
	// **Added In:** 2009232244
	// **SCIM++ Properties:**
	//  - idcsSearchable: false
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: false
	//  - returned: default
	//  - type: boolean
	//  - uniqueness: none
	FidoAuthenticatorEnabled *bool `mandatory:"false" json:"fidoAuthenticatorEnabled"`

	// If true, indicates that the Yubico OTP is enabled for authentication
	// **Added In:** 2109090424
	// **SCIM++ Properties:**
	//  - idcsSearchable: false
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: false
	//  - returned: default
	//  - type: boolean
	//  - uniqueness: none
	YubicoOtpEnabled *bool `mandatory:"false" json:"yubicoOtpEnabled"`

	// Specifies the category of people for whom Multi-Factor Authentication is enabled. This is a readOnly attribute which reflects the value of mfaEnabledCategory attribute in SsoSettings
	// **Deprecated Since: 18.1.2**
	// **SCIM++ Properties:**
	//  - idcsSearchable: false
	//  - multiValued: false
	//  - mutability: readOnly
	//  - required: false
	//  - returned: default
	//  - type: string
	//  - uniqueness: none
	MfaEnabledCategory *string `mandatory:"false" json:"mfaEnabledCategory"`

	// If true, indicates that 'Show backup factor(s)' button will be hidden during authentication
	// **Added In:** 19.3.3
	// **SCIM++ Properties:**
	//  - idcsSearchable: false
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: false
	//  - returned: default
	//  - type: boolean
	//  - uniqueness: none
	HideBackupFactorEnabled *bool `mandatory:"false" json:"hideBackupFactorEnabled"`

	// If true, indicates that email will not be enrolled as a MFA factor automatically if it a account recovery factor
	// **Added In:** 2011192329
	// **SCIM++ Properties:**
	//  - idcsSearchable: false
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: false
	//  - returned: default
	//  - type: boolean
	//  - uniqueness: none
	AutoEnrollEmailFactorDisabled *bool `mandatory:"false" json:"autoEnrollEmailFactorDisabled"`

	// Factors for which enrollment should be blocked for End User
	// **Added In:** 2012271618
	// **SCIM++ Properties:**
	//  - idcsSearchable: false
	//  - multiValued: true
	//  - mutability: readWrite
	//  - required: false
	//  - returned: default
	//  - type: string
	//  - uniqueness: none
	UserEnrollmentDisabledFactors []AuthenticationFactorSettingUserEnrollmentDisabledFactorsEnum `mandatory:"false" json:"userEnrollmentDisabledFactors,omitempty"`

	EmailSettings *AuthenticationFactorSettingsEmailSettings `mandatory:"false" json:"emailSettings"`

	ThirdPartyFactor *AuthenticationFactorSettingsThirdPartyFactor `mandatory:"false" json:"thirdPartyFactor"`

	IdentityStoreSettings *AuthenticationFactorSettingsIdentityStoreSettings `mandatory:"false" json:"identityStoreSettings"`

	UrnietfparamsscimschemasoracleidcsextensionthirdPartyAuthenticationFactorSettings *ExtensionThirdPartyAuthenticationFactorSettings `mandatory:"false" json:"urn:ietf:params:scim:schemas:oracle:idcs:extension:thirdParty:AuthenticationFactorSettings"`

	UrnietfparamsscimschemasoracleidcsextensionfidoAuthenticationFactorSettings *ExtensionFidoAuthenticationFactorSettings `mandatory:"false" json:"urn:ietf:params:scim:schemas:oracle:idcs:extension:fido:AuthenticationFactorSettings"`
}

func (m AuthenticationFactorSetting) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m AuthenticationFactorSetting) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	for _, val := range m.IdcsPreventedOperations {
		if _, ok := GetMappingIdcsPreventedOperationsEnum(string(val)); !ok && val != "" {
			errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for IdcsPreventedOperations: %s. Supported values are: %s.", val, strings.Join(GetIdcsPreventedOperationsEnumStringValues(), ",")))
		}
	}

	for _, val := range m.UserEnrollmentDisabledFactors {
		if _, ok := GetMappingAuthenticationFactorSettingUserEnrollmentDisabledFactorsEnum(string(val)); !ok && val != "" {
			errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for UserEnrollmentDisabledFactors: %s. Supported values are: %s.", val, strings.Join(GetAuthenticationFactorSettingUserEnrollmentDisabledFactorsEnumStringValues(), ",")))
		}
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// AuthenticationFactorSettingUserEnrollmentDisabledFactorsEnum Enum with underlying type: string
type AuthenticationFactorSettingUserEnrollmentDisabledFactorsEnum string

// Set of constants representing the allowable values for AuthenticationFactorSettingUserEnrollmentDisabledFactorsEnum
const (
	AuthenticationFactorSettingUserEnrollmentDisabledFactorsEmail             AuthenticationFactorSettingUserEnrollmentDisabledFactorsEnum = "EMAIL"
	AuthenticationFactorSettingUserEnrollmentDisabledFactorsSms               AuthenticationFactorSettingUserEnrollmentDisabledFactorsEnum = "SMS"
	AuthenticationFactorSettingUserEnrollmentDisabledFactorsTotp              AuthenticationFactorSettingUserEnrollmentDisabledFactorsEnum = "TOTP"
	AuthenticationFactorSettingUserEnrollmentDisabledFactorsPush              AuthenticationFactorSettingUserEnrollmentDisabledFactorsEnum = "PUSH"
	AuthenticationFactorSettingUserEnrollmentDisabledFactorsOfflinetotp       AuthenticationFactorSettingUserEnrollmentDisabledFactorsEnum = "OFFLINETOTP"
	AuthenticationFactorSettingUserEnrollmentDisabledFactorsVoice             AuthenticationFactorSettingUserEnrollmentDisabledFactorsEnum = "VOICE"
	AuthenticationFactorSettingUserEnrollmentDisabledFactorsPhoneCall         AuthenticationFactorSettingUserEnrollmentDisabledFactorsEnum = "PHONE_CALL"
	AuthenticationFactorSettingUserEnrollmentDisabledFactorsThirdparty        AuthenticationFactorSettingUserEnrollmentDisabledFactorsEnum = "THIRDPARTY"
	AuthenticationFactorSettingUserEnrollmentDisabledFactorsFidoAuthenticator AuthenticationFactorSettingUserEnrollmentDisabledFactorsEnum = "FIDO_AUTHENTICATOR"
	AuthenticationFactorSettingUserEnrollmentDisabledFactorsYubicoOtp         AuthenticationFactorSettingUserEnrollmentDisabledFactorsEnum = "YUBICO_OTP"
)

var mappingAuthenticationFactorSettingUserEnrollmentDisabledFactorsEnum = map[string]AuthenticationFactorSettingUserEnrollmentDisabledFactorsEnum{
	"EMAIL":              AuthenticationFactorSettingUserEnrollmentDisabledFactorsEmail,
	"SMS":                AuthenticationFactorSettingUserEnrollmentDisabledFactorsSms,
	"TOTP":               AuthenticationFactorSettingUserEnrollmentDisabledFactorsTotp,
	"PUSH":               AuthenticationFactorSettingUserEnrollmentDisabledFactorsPush,
	"OFFLINETOTP":        AuthenticationFactorSettingUserEnrollmentDisabledFactorsOfflinetotp,
	"VOICE":              AuthenticationFactorSettingUserEnrollmentDisabledFactorsVoice,
	"PHONE_CALL":         AuthenticationFactorSettingUserEnrollmentDisabledFactorsPhoneCall,
	"THIRDPARTY":         AuthenticationFactorSettingUserEnrollmentDisabledFactorsThirdparty,
	"FIDO_AUTHENTICATOR": AuthenticationFactorSettingUserEnrollmentDisabledFactorsFidoAuthenticator,
	"YUBICO_OTP":         AuthenticationFactorSettingUserEnrollmentDisabledFactorsYubicoOtp,
}

var mappingAuthenticationFactorSettingUserEnrollmentDisabledFactorsEnumLowerCase = map[string]AuthenticationFactorSettingUserEnrollmentDisabledFactorsEnum{
	"email":              AuthenticationFactorSettingUserEnrollmentDisabledFactorsEmail,
	"sms":                AuthenticationFactorSettingUserEnrollmentDisabledFactorsSms,
	"totp":               AuthenticationFactorSettingUserEnrollmentDisabledFactorsTotp,
	"push":               AuthenticationFactorSettingUserEnrollmentDisabledFactorsPush,
	"offlinetotp":        AuthenticationFactorSettingUserEnrollmentDisabledFactorsOfflinetotp,
	"voice":              AuthenticationFactorSettingUserEnrollmentDisabledFactorsVoice,
	"phone_call":         AuthenticationFactorSettingUserEnrollmentDisabledFactorsPhoneCall,
	"thirdparty":         AuthenticationFactorSettingUserEnrollmentDisabledFactorsThirdparty,
	"fido_authenticator": AuthenticationFactorSettingUserEnrollmentDisabledFactorsFidoAuthenticator,
	"yubico_otp":         AuthenticationFactorSettingUserEnrollmentDisabledFactorsYubicoOtp,
}

// GetAuthenticationFactorSettingUserEnrollmentDisabledFactorsEnumValues Enumerates the set of values for AuthenticationFactorSettingUserEnrollmentDisabledFactorsEnum
func GetAuthenticationFactorSettingUserEnrollmentDisabledFactorsEnumValues() []AuthenticationFactorSettingUserEnrollmentDisabledFactorsEnum {
	values := make([]AuthenticationFactorSettingUserEnrollmentDisabledFactorsEnum, 0)
	for _, v := range mappingAuthenticationFactorSettingUserEnrollmentDisabledFactorsEnum {
		values = append(values, v)
	}
	return values
}

// GetAuthenticationFactorSettingUserEnrollmentDisabledFactorsEnumStringValues Enumerates the set of values in String for AuthenticationFactorSettingUserEnrollmentDisabledFactorsEnum
func GetAuthenticationFactorSettingUserEnrollmentDisabledFactorsEnumStringValues() []string {
	return []string{
		"EMAIL",
		"SMS",
		"TOTP",
		"PUSH",
		"OFFLINETOTP",
		"VOICE",
		"PHONE_CALL",
		"THIRDPARTY",
		"FIDO_AUTHENTICATOR",
		"YUBICO_OTP",
	}
}

// GetMappingAuthenticationFactorSettingUserEnrollmentDisabledFactorsEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingAuthenticationFactorSettingUserEnrollmentDisabledFactorsEnum(val string) (AuthenticationFactorSettingUserEnrollmentDisabledFactorsEnum, bool) {
	enum, ok := mappingAuthenticationFactorSettingUserEnrollmentDisabledFactorsEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
