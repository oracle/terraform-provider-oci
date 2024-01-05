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

// MyAuthenticationFactorValidator Validate any given Authentication Factor
type MyAuthenticationFactorValidator struct {

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

	// Authentication Factor which is being validated
	// **SCIM++ Properties:**
	//  - type: string
	//  - multiValued: false
	//  - required: true
	//  - mutability: readWrite
	//  - returned: default
	//  - uniqueness: none
	//  - idcsSearchable: false
	AuthFactor MyAuthenticationFactorValidatorAuthFactorEnum `mandatory:"true" json:"authFactor"`

	// Specifies whether the service is being used to enroll or validate a factor
	// **SCIM++ Properties:**
	//  - type: string
	//  - multiValued: false
	//  - required: true
	//  - mutability: readWrite
	//  - returned: default
	//  - uniqueness: none
	//  - idcsSearchable: false
	Scenario MyAuthenticationFactorValidatorScenarioEnum `mandatory:"true" json:"scenario"`

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

	// Request ID which is being validated
	// **SCIM++ Properties:**
	//  - type: string
	//  - multiValued: false
	//  - required: false
	//  - mutability: readWrite
	//  - returned: default
	//  - uniqueness: none
	//  - idcsSearchable: false
	RequestId *string `mandatory:"false" json:"requestId"`

	// The One Time Passcode which needs to be validated
	// **SCIM++ Properties:**
	//  - type: string
	//  - multiValued: false
	//  - required: false
	//  - mutability: readWrite
	//  - returned: never
	//  - uniqueness: none
	//  - idcsSensitive: encrypt
	//  - idcsSearchable: false
	OtpCode *string `mandatory:"false" json:"otpCode"`

	// Device id whose factor is being validated
	// **SCIM++ Properties:**
	//  - type: string
	//  - multiValued: false
	//  - required: false
	//  - mutability: readWrite
	//  - returned: default
	//  - uniqueness: none
	//  - idcsSearchable: false
	DeviceId *string `mandatory:"false" json:"deviceId"`

	// Validation status returned in the response
	// **SCIM++ Properties:**
	//  - type: string
	//  - multiValued: false
	//  - required: false
	//  - mutability: readOnly
	//  - returned: default
	//  - uniqueness: none
	//  - idcsSearchable: false
	Status MyAuthenticationFactorValidatorStatusEnum `mandatory:"false" json:"status,omitempty"`

	// User guid for whom the validation has initiated. Optional.
	// **SCIM++ Properties:**
	//  - type: string
	//  - multiValued: false
	//  - required: false
	//  - mutability: readWrite
	//  - returned: default
	//  - uniqueness: none
	//  - idcsSearchable: false
	UserId *string `mandatory:"false" json:"userId"`

	// User name for whom the validation has initiated
	// **SCIM++ Properties:**
	//  - type: string
	//  - multiValued: false
	//  - required: false
	//  - mutability: readWrite
	//  - returned: default
	//  - uniqueness: none
	//  - idcsSearchable: false
	//  - idcsPii: true
	UserName *string `mandatory:"false" json:"userName"`

	// Display name of the verified device
	// **SCIM++ Properties:**
	//  - type: string
	//  - multiValued: false
	//  - required: false
	//  - mutability: readWrite
	//  - returned: default
	//  - uniqueness: none
	//  - idcsSearchable: false
	DisplayName *string `mandatory:"false" json:"displayName"`

	// Validator message which is passed by the client. When it is a PUSH notification, it can be a rejection message.
	// **SCIM++ Properties:**
	//  - type: string
	//  - multiValued: false
	//  - required: false
	//  - mutability: readWrite
	//  - returned: default
	//  - uniqueness: none
	//  - idcsSearchable: false
	Message *string `mandatory:"false" json:"message"`

	// type indicating whether the flow is OIDC, SAML etc.,
	// **SCIM++ Properties:**
	//  - type: string
	//  - multiValued: false
	//  - required: false
	//  - mutability: writeOnly
	//  - returned: default
	//  - uniqueness: none
	Type MyAuthenticationFactorValidatorTypeEnum `mandatory:"false" json:"type,omitempty"`

	// Indicates whether to update user preferred mfa factor or not
	// **SCIM++ Properties:**
	//  - type: boolean
	//  - multiValued: false
	//  - required: false
	//  - mutability: writeOnly
	//  - returned: never
	//  - uniqueness: none
	UpdateUserPreference *bool `mandatory:"false" json:"updateUserPreference"`

	// Indicates whether to user passwordless factor to be updated or mfa factor to be updated
	// **SCIM++ Properties:**
	//  - type: string
	//  - multiValued: false
	//  - required: false
	//  - mutability: writeOnly
	//  - returned: never
	//  - uniqueness: none
	PreferenceType MyAuthenticationFactorValidatorPreferenceTypeEnum `mandatory:"false" json:"preferenceType,omitempty"`

	// List of security questions the user has submitted to get authenticated.
	// **SCIM++ Properties:**
	//  - type: complex
	//  - multiValued: true
	//  - required: false
	//  - mutability: writeOnly
	//  - returned: never
	//  - uniqueness: none
	//  - idcsSearchable: false
	SecurityQuestions []MyAuthenticationFactorValidatorSecurityQuestions `mandatory:"false" json:"securityQuestions"`

	// Name of the client to be trusted
	// **SCIM++ Properties:**
	//  - type: string
	//  - multiValued: false
	//  - required: false
	//  - mutability: writeOnly
	//  - returned: never
	//  - uniqueness: none
	Name *string `mandatory:"false" json:"name"`

	// Platform of the client to be trusted
	// **SCIM++ Properties:**
	//  - type: string
	//  - multiValued: false
	//  - required: false
	//  - mutability: writeOnly
	//  - returned: never
	//  - uniqueness: none
	Platform *string `mandatory:"false" json:"platform"`

	// Location of the trusted client.
	// **SCIM++ Properties:**
	//  - type: string
	//  - multiValued: false
	//  - required: false
	//  - mutability: writeOnly
	//  - returned: never
	//  - uniqueness: none
	Location *string `mandatory:"false" json:"location"`

	// Trusted token resource identifier.
	// **SCIM++ Properties:**
	//  - type: string
	//  - multiValued: false
	//  - required: false
	//  - mutability: writeOnly
	//  - returned: never
	//  - uniqueness: none
	TrustedTokenId *string `mandatory:"false" json:"trustedTokenId"`

	// KMSI token resource identifier.
	// **SCIM++ Properties:**
	//  - type: string
	//  - multiValued: false
	//  - required: false
	//  - mutability: writeOnly
	//  - returned: never
	//  - uniqueness: none
	KmsiTokenId *string `mandatory:"false" json:"kmsiTokenId"`

	// Sign-On Policy dictated allowed second factors.
	// **SCIM++ Properties:**
	//  - type: string
	//  - multiValued: true
	//  - required: false
	//  - mutability: writeOnly
	//  - returned: never
	//  - uniqueness: none
	PolicyEnabledSecondFactors []string `mandatory:"false" json:"policyEnabledSecondFactors"`

	// Indicates to create trust token.
	// **SCIM++ Properties:**
	//  - type: boolean
	//  - multiValued: false
	//  - required: false
	//  - mutability: writeOnly
	//  - returned: never
	//  - uniqueness: none
	CreateTrustedAgent *bool `mandatory:"false" json:"createTrustedAgent"`

	// Indicates to create kmsi token.
	// **SCIM++ Properties:**
	//  - type: boolean
	//  - multiValued: false
	//  - required: false
	//  - mutability: writeOnly
	//  - returned: never
	//  - uniqueness: none
	CreateKmsiToken *bool `mandatory:"false" json:"createKmsiToken"`

	// Flag indicates whether the factor is enrolled in account recovery. If the value is not provided or false, then it will be treated as MFA factor validation.
	// **SCIM++ Properties:**
	//  - caseExact: true
	//  - idcsSearchable: true
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: false
	//  - returned: default
	//  - type: boolean
	//  - uniqueness: none
	IsAccRecEnabled *bool `mandatory:"false" json:"isAccRecEnabled"`

	// Sign-On Policy dictated validity duration for trusted client in Minutes.
	// **SCIM++ Properties:**
	//  - type: integer
	//  - multiValued: false
	//  - required: false
	//  - mutability: writeOnly
	//  - returned: never
	//  - uniqueness: none
	PolicyTrustedFrequencyMins *int `mandatory:"false" json:"policyTrustedFrequencyMins"`

	ThirdPartyFactor *MyAuthenticationFactorValidatorThirdPartyFactor `mandatory:"false" json:"thirdPartyFactor"`

	// Additional attributes which will be sent as part of a push notification
	// **SCIM++ Properties:**
	//  - idcsSearchable: false
	//  - multiValued: true
	//  - mutability: readWrite
	//  - required: false
	//  - returned: default
	//  - type: complex
	//  - uniqueness: none
	AdditionalAttributes []MyAuthenticationFactorValidatorAdditionalAttributes `mandatory:"false" json:"additionalAttributes"`
}

func (m MyAuthenticationFactorValidator) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m MyAuthenticationFactorValidator) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingMyAuthenticationFactorValidatorAuthFactorEnum(string(m.AuthFactor)); !ok && m.AuthFactor != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for AuthFactor: %s. Supported values are: %s.", m.AuthFactor, strings.Join(GetMyAuthenticationFactorValidatorAuthFactorEnumStringValues(), ",")))
	}
	if _, ok := GetMappingMyAuthenticationFactorValidatorScenarioEnum(string(m.Scenario)); !ok && m.Scenario != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Scenario: %s. Supported values are: %s.", m.Scenario, strings.Join(GetMyAuthenticationFactorValidatorScenarioEnumStringValues(), ",")))
	}

	for _, val := range m.IdcsPreventedOperations {
		if _, ok := GetMappingIdcsPreventedOperationsEnum(string(val)); !ok && val != "" {
			errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for IdcsPreventedOperations: %s. Supported values are: %s.", val, strings.Join(GetIdcsPreventedOperationsEnumStringValues(), ",")))
		}
	}

	if _, ok := GetMappingMyAuthenticationFactorValidatorStatusEnum(string(m.Status)); !ok && m.Status != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Status: %s. Supported values are: %s.", m.Status, strings.Join(GetMyAuthenticationFactorValidatorStatusEnumStringValues(), ",")))
	}
	if _, ok := GetMappingMyAuthenticationFactorValidatorTypeEnum(string(m.Type)); !ok && m.Type != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Type: %s. Supported values are: %s.", m.Type, strings.Join(GetMyAuthenticationFactorValidatorTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingMyAuthenticationFactorValidatorPreferenceTypeEnum(string(m.PreferenceType)); !ok && m.PreferenceType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for PreferenceType: %s. Supported values are: %s.", m.PreferenceType, strings.Join(GetMyAuthenticationFactorValidatorPreferenceTypeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MyAuthenticationFactorValidatorAuthFactorEnum Enum with underlying type: string
type MyAuthenticationFactorValidatorAuthFactorEnum string

// Set of constants representing the allowable values for MyAuthenticationFactorValidatorAuthFactorEnum
const (
	MyAuthenticationFactorValidatorAuthFactorUsernamePassword  MyAuthenticationFactorValidatorAuthFactorEnum = "USERNAME_PASSWORD"
	MyAuthenticationFactorValidatorAuthFactorPush              MyAuthenticationFactorValidatorAuthFactorEnum = "PUSH"
	MyAuthenticationFactorValidatorAuthFactorTotp              MyAuthenticationFactorValidatorAuthFactorEnum = "TOTP"
	MyAuthenticationFactorValidatorAuthFactorEmail             MyAuthenticationFactorValidatorAuthFactorEnum = "EMAIL"
	MyAuthenticationFactorValidatorAuthFactorSms               MyAuthenticationFactorValidatorAuthFactorEnum = "SMS"
	MyAuthenticationFactorValidatorAuthFactorVoice             MyAuthenticationFactorValidatorAuthFactorEnum = "VOICE"
	MyAuthenticationFactorValidatorAuthFactorBypasscode        MyAuthenticationFactorValidatorAuthFactorEnum = "BYPASSCODE"
	MyAuthenticationFactorValidatorAuthFactorSecurityQuestions MyAuthenticationFactorValidatorAuthFactorEnum = "SECURITY_QUESTIONS"
	MyAuthenticationFactorValidatorAuthFactorTrustToken        MyAuthenticationFactorValidatorAuthFactorEnum = "TRUST_TOKEN"
	MyAuthenticationFactorValidatorAuthFactorPhoneCall         MyAuthenticationFactorValidatorAuthFactorEnum = "PHONE_CALL"
	MyAuthenticationFactorValidatorAuthFactorThirdparty        MyAuthenticationFactorValidatorAuthFactorEnum = "THIRDPARTY"
	MyAuthenticationFactorValidatorAuthFactorFidoAuthenticator MyAuthenticationFactorValidatorAuthFactorEnum = "FIDO_AUTHENTICATOR"
	MyAuthenticationFactorValidatorAuthFactorYubicoOtp         MyAuthenticationFactorValidatorAuthFactorEnum = "YUBICO_OTP"
	MyAuthenticationFactorValidatorAuthFactorKmsiToken         MyAuthenticationFactorValidatorAuthFactorEnum = "KMSI_TOKEN"
)

var mappingMyAuthenticationFactorValidatorAuthFactorEnum = map[string]MyAuthenticationFactorValidatorAuthFactorEnum{
	"USERNAME_PASSWORD":  MyAuthenticationFactorValidatorAuthFactorUsernamePassword,
	"PUSH":               MyAuthenticationFactorValidatorAuthFactorPush,
	"TOTP":               MyAuthenticationFactorValidatorAuthFactorTotp,
	"EMAIL":              MyAuthenticationFactorValidatorAuthFactorEmail,
	"SMS":                MyAuthenticationFactorValidatorAuthFactorSms,
	"VOICE":              MyAuthenticationFactorValidatorAuthFactorVoice,
	"BYPASSCODE":         MyAuthenticationFactorValidatorAuthFactorBypasscode,
	"SECURITY_QUESTIONS": MyAuthenticationFactorValidatorAuthFactorSecurityQuestions,
	"TRUST_TOKEN":        MyAuthenticationFactorValidatorAuthFactorTrustToken,
	"PHONE_CALL":         MyAuthenticationFactorValidatorAuthFactorPhoneCall,
	"THIRDPARTY":         MyAuthenticationFactorValidatorAuthFactorThirdparty,
	"FIDO_AUTHENTICATOR": MyAuthenticationFactorValidatorAuthFactorFidoAuthenticator,
	"YUBICO_OTP":         MyAuthenticationFactorValidatorAuthFactorYubicoOtp,
	"KMSI_TOKEN":         MyAuthenticationFactorValidatorAuthFactorKmsiToken,
}

var mappingMyAuthenticationFactorValidatorAuthFactorEnumLowerCase = map[string]MyAuthenticationFactorValidatorAuthFactorEnum{
	"username_password":  MyAuthenticationFactorValidatorAuthFactorUsernamePassword,
	"push":               MyAuthenticationFactorValidatorAuthFactorPush,
	"totp":               MyAuthenticationFactorValidatorAuthFactorTotp,
	"email":              MyAuthenticationFactorValidatorAuthFactorEmail,
	"sms":                MyAuthenticationFactorValidatorAuthFactorSms,
	"voice":              MyAuthenticationFactorValidatorAuthFactorVoice,
	"bypasscode":         MyAuthenticationFactorValidatorAuthFactorBypasscode,
	"security_questions": MyAuthenticationFactorValidatorAuthFactorSecurityQuestions,
	"trust_token":        MyAuthenticationFactorValidatorAuthFactorTrustToken,
	"phone_call":         MyAuthenticationFactorValidatorAuthFactorPhoneCall,
	"thirdparty":         MyAuthenticationFactorValidatorAuthFactorThirdparty,
	"fido_authenticator": MyAuthenticationFactorValidatorAuthFactorFidoAuthenticator,
	"yubico_otp":         MyAuthenticationFactorValidatorAuthFactorYubicoOtp,
	"kmsi_token":         MyAuthenticationFactorValidatorAuthFactorKmsiToken,
}

// GetMyAuthenticationFactorValidatorAuthFactorEnumValues Enumerates the set of values for MyAuthenticationFactorValidatorAuthFactorEnum
func GetMyAuthenticationFactorValidatorAuthFactorEnumValues() []MyAuthenticationFactorValidatorAuthFactorEnum {
	values := make([]MyAuthenticationFactorValidatorAuthFactorEnum, 0)
	for _, v := range mappingMyAuthenticationFactorValidatorAuthFactorEnum {
		values = append(values, v)
	}
	return values
}

// GetMyAuthenticationFactorValidatorAuthFactorEnumStringValues Enumerates the set of values in String for MyAuthenticationFactorValidatorAuthFactorEnum
func GetMyAuthenticationFactorValidatorAuthFactorEnumStringValues() []string {
	return []string{
		"USERNAME_PASSWORD",
		"PUSH",
		"TOTP",
		"EMAIL",
		"SMS",
		"VOICE",
		"BYPASSCODE",
		"SECURITY_QUESTIONS",
		"TRUST_TOKEN",
		"PHONE_CALL",
		"THIRDPARTY",
		"FIDO_AUTHENTICATOR",
		"YUBICO_OTP",
		"KMSI_TOKEN",
	}
}

// GetMappingMyAuthenticationFactorValidatorAuthFactorEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingMyAuthenticationFactorValidatorAuthFactorEnum(val string) (MyAuthenticationFactorValidatorAuthFactorEnum, bool) {
	enum, ok := mappingMyAuthenticationFactorValidatorAuthFactorEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// MyAuthenticationFactorValidatorScenarioEnum Enum with underlying type: string
type MyAuthenticationFactorValidatorScenarioEnum string

// Set of constants representing the allowable values for MyAuthenticationFactorValidatorScenarioEnum
const (
	MyAuthenticationFactorValidatorScenarioEnrollment     MyAuthenticationFactorValidatorScenarioEnum = "ENROLLMENT"
	MyAuthenticationFactorValidatorScenarioAuthentication MyAuthenticationFactorValidatorScenarioEnum = "AUTHENTICATION"
)

var mappingMyAuthenticationFactorValidatorScenarioEnum = map[string]MyAuthenticationFactorValidatorScenarioEnum{
	"ENROLLMENT":     MyAuthenticationFactorValidatorScenarioEnrollment,
	"AUTHENTICATION": MyAuthenticationFactorValidatorScenarioAuthentication,
}

var mappingMyAuthenticationFactorValidatorScenarioEnumLowerCase = map[string]MyAuthenticationFactorValidatorScenarioEnum{
	"enrollment":     MyAuthenticationFactorValidatorScenarioEnrollment,
	"authentication": MyAuthenticationFactorValidatorScenarioAuthentication,
}

// GetMyAuthenticationFactorValidatorScenarioEnumValues Enumerates the set of values for MyAuthenticationFactorValidatorScenarioEnum
func GetMyAuthenticationFactorValidatorScenarioEnumValues() []MyAuthenticationFactorValidatorScenarioEnum {
	values := make([]MyAuthenticationFactorValidatorScenarioEnum, 0)
	for _, v := range mappingMyAuthenticationFactorValidatorScenarioEnum {
		values = append(values, v)
	}
	return values
}

// GetMyAuthenticationFactorValidatorScenarioEnumStringValues Enumerates the set of values in String for MyAuthenticationFactorValidatorScenarioEnum
func GetMyAuthenticationFactorValidatorScenarioEnumStringValues() []string {
	return []string{
		"ENROLLMENT",
		"AUTHENTICATION",
	}
}

// GetMappingMyAuthenticationFactorValidatorScenarioEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingMyAuthenticationFactorValidatorScenarioEnum(val string) (MyAuthenticationFactorValidatorScenarioEnum, bool) {
	enum, ok := mappingMyAuthenticationFactorValidatorScenarioEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// MyAuthenticationFactorValidatorStatusEnum Enum with underlying type: string
type MyAuthenticationFactorValidatorStatusEnum string

// Set of constants representing the allowable values for MyAuthenticationFactorValidatorStatusEnum
const (
	MyAuthenticationFactorValidatorStatusSuccess MyAuthenticationFactorValidatorStatusEnum = "SUCCESS"
	MyAuthenticationFactorValidatorStatusFailure MyAuthenticationFactorValidatorStatusEnum = "FAILURE"
)

var mappingMyAuthenticationFactorValidatorStatusEnum = map[string]MyAuthenticationFactorValidatorStatusEnum{
	"SUCCESS": MyAuthenticationFactorValidatorStatusSuccess,
	"FAILURE": MyAuthenticationFactorValidatorStatusFailure,
}

var mappingMyAuthenticationFactorValidatorStatusEnumLowerCase = map[string]MyAuthenticationFactorValidatorStatusEnum{
	"success": MyAuthenticationFactorValidatorStatusSuccess,
	"failure": MyAuthenticationFactorValidatorStatusFailure,
}

// GetMyAuthenticationFactorValidatorStatusEnumValues Enumerates the set of values for MyAuthenticationFactorValidatorStatusEnum
func GetMyAuthenticationFactorValidatorStatusEnumValues() []MyAuthenticationFactorValidatorStatusEnum {
	values := make([]MyAuthenticationFactorValidatorStatusEnum, 0)
	for _, v := range mappingMyAuthenticationFactorValidatorStatusEnum {
		values = append(values, v)
	}
	return values
}

// GetMyAuthenticationFactorValidatorStatusEnumStringValues Enumerates the set of values in String for MyAuthenticationFactorValidatorStatusEnum
func GetMyAuthenticationFactorValidatorStatusEnumStringValues() []string {
	return []string{
		"SUCCESS",
		"FAILURE",
	}
}

// GetMappingMyAuthenticationFactorValidatorStatusEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingMyAuthenticationFactorValidatorStatusEnum(val string) (MyAuthenticationFactorValidatorStatusEnum, bool) {
	enum, ok := mappingMyAuthenticationFactorValidatorStatusEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// MyAuthenticationFactorValidatorTypeEnum Enum with underlying type: string
type MyAuthenticationFactorValidatorTypeEnum string

// Set of constants representing the allowable values for MyAuthenticationFactorValidatorTypeEnum
const (
	MyAuthenticationFactorValidatorTypeSaml MyAuthenticationFactorValidatorTypeEnum = "SAML"
	MyAuthenticationFactorValidatorTypeOidc MyAuthenticationFactorValidatorTypeEnum = "OIDC"
)

var mappingMyAuthenticationFactorValidatorTypeEnum = map[string]MyAuthenticationFactorValidatorTypeEnum{
	"SAML": MyAuthenticationFactorValidatorTypeSaml,
	"OIDC": MyAuthenticationFactorValidatorTypeOidc,
}

var mappingMyAuthenticationFactorValidatorTypeEnumLowerCase = map[string]MyAuthenticationFactorValidatorTypeEnum{
	"saml": MyAuthenticationFactorValidatorTypeSaml,
	"oidc": MyAuthenticationFactorValidatorTypeOidc,
}

// GetMyAuthenticationFactorValidatorTypeEnumValues Enumerates the set of values for MyAuthenticationFactorValidatorTypeEnum
func GetMyAuthenticationFactorValidatorTypeEnumValues() []MyAuthenticationFactorValidatorTypeEnum {
	values := make([]MyAuthenticationFactorValidatorTypeEnum, 0)
	for _, v := range mappingMyAuthenticationFactorValidatorTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetMyAuthenticationFactorValidatorTypeEnumStringValues Enumerates the set of values in String for MyAuthenticationFactorValidatorTypeEnum
func GetMyAuthenticationFactorValidatorTypeEnumStringValues() []string {
	return []string{
		"SAML",
		"OIDC",
	}
}

// GetMappingMyAuthenticationFactorValidatorTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingMyAuthenticationFactorValidatorTypeEnum(val string) (MyAuthenticationFactorValidatorTypeEnum, bool) {
	enum, ok := mappingMyAuthenticationFactorValidatorTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// MyAuthenticationFactorValidatorPreferenceTypeEnum Enum with underlying type: string
type MyAuthenticationFactorValidatorPreferenceTypeEnum string

// Set of constants representing the allowable values for MyAuthenticationFactorValidatorPreferenceTypeEnum
const (
	MyAuthenticationFactorValidatorPreferenceTypePasswordless MyAuthenticationFactorValidatorPreferenceTypeEnum = "PASSWORDLESS"
	MyAuthenticationFactorValidatorPreferenceTypeMfa          MyAuthenticationFactorValidatorPreferenceTypeEnum = "MFA"
)

var mappingMyAuthenticationFactorValidatorPreferenceTypeEnum = map[string]MyAuthenticationFactorValidatorPreferenceTypeEnum{
	"PASSWORDLESS": MyAuthenticationFactorValidatorPreferenceTypePasswordless,
	"MFA":          MyAuthenticationFactorValidatorPreferenceTypeMfa,
}

var mappingMyAuthenticationFactorValidatorPreferenceTypeEnumLowerCase = map[string]MyAuthenticationFactorValidatorPreferenceTypeEnum{
	"passwordless": MyAuthenticationFactorValidatorPreferenceTypePasswordless,
	"mfa":          MyAuthenticationFactorValidatorPreferenceTypeMfa,
}

// GetMyAuthenticationFactorValidatorPreferenceTypeEnumValues Enumerates the set of values for MyAuthenticationFactorValidatorPreferenceTypeEnum
func GetMyAuthenticationFactorValidatorPreferenceTypeEnumValues() []MyAuthenticationFactorValidatorPreferenceTypeEnum {
	values := make([]MyAuthenticationFactorValidatorPreferenceTypeEnum, 0)
	for _, v := range mappingMyAuthenticationFactorValidatorPreferenceTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetMyAuthenticationFactorValidatorPreferenceTypeEnumStringValues Enumerates the set of values in String for MyAuthenticationFactorValidatorPreferenceTypeEnum
func GetMyAuthenticationFactorValidatorPreferenceTypeEnumStringValues() []string {
	return []string{
		"PASSWORDLESS",
		"MFA",
	}
}

// GetMappingMyAuthenticationFactorValidatorPreferenceTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingMyAuthenticationFactorValidatorPreferenceTypeEnum(val string) (MyAuthenticationFactorValidatorPreferenceTypeEnum, bool) {
	enum, ok := mappingMyAuthenticationFactorValidatorPreferenceTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
