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

// MyAuthenticationFactorInitiator This schema defines the attributes of Initiator call.
type MyAuthenticationFactorInitiator struct {

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

	// Auth Factor represents the type of multi-factor authentication channel for which the request has been initiated.
	// **SCIM++ Properties:**
	//  - idcsSearchable: false
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: true
	//  - returned: default
	//  - type: string
	//  - uniqueness: none
	AuthFactor MyAuthenticationFactorInitiatorAuthFactorEnum `mandatory:"true" json:"authFactor"`

	// Enrolled Device id on which the multi factor has been initiated.
	// **SCIM++ Properties:**
	//  - caseExact: true
	//  - idcsSearchable: false
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: true
	//  - returned: default
	//  - type: string
	//  - uniqueness: none
	DeviceId *string `mandatory:"true" json:"deviceId"`

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

	// Authentication flow type either SAML / OIDC
	// **SCIM++ Properties:**
	//  - caseExact: true
	//  - idcsSearchable: false
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: false
	//  - returned: default
	//  - type: string
	//  - uniqueness: none
	Type MyAuthenticationFactorInitiatorTypeEnum `mandatory:"false" json:"type,omitempty"`

	// Unique RequestId generated for each initiator request.
	// **SCIM++ Properties:**
	//  - caseExact: true
	//  - idcsSearchable: false
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: false
	//  - returned: default
	//  - type: string
	//  - uniqueness: none
	RequestId *string `mandatory:"false" json:"requestId"`

	// Name of the user who initiates the request.
	// **SCIM++ Properties:**
	//  - caseExact: true
	//  - idcsSearchable: false
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: false
	//  - returned: default
	//  - type: string
	//  - uniqueness: none
	//  - idcsPii: true
	UserName *string `mandatory:"false" json:"userName"`

	// Specifies the scenario to initiate.
	// **SCIM++ Properties:**
	//  - type: string
	//  - multiValued: false
	//  - required: false
	//  - mutability: readWrite
	//  - returned: default
	//  - uniqueness: none
	//  - idcsSearchable: false
	Scenario MyAuthenticationFactorInitiatorScenarioEnum `mandatory:"false" json:"scenario,omitempty"`

	ThirdPartyFactor *MyAuthenticationFactorInitiatorThirdPartyFactor `mandatory:"false" json:"thirdPartyFactor"`

	// Indicates whether to user passwordless factor to be updated or mfa factor to be updated
	// **SCIM++ Properties:**
	//  - type: string
	//  - multiValued: false
	//  - required: false
	//  - mutability: writeOnly
	//  - returned: never
	//  - uniqueness: none
	PreferenceType MyAuthenticationFactorInitiatorPreferenceTypeEnum `mandatory:"false" json:"preferenceType,omitempty"`

	// Additional attributes which will be sent as part of a push notification
	// **SCIM++ Properties:**
	//  - idcsSearchable: false
	//  - multiValued: true
	//  - mutability: readWrite
	//  - required: false
	//  - returned: default
	//  - type: complex
	//  - uniqueness: none
	AdditionalAttributes []MyAuthenticationFactorInitiatorAdditionalAttributes `mandatory:"false" json:"additionalAttributes"`

	// Flag indicates whether the device is enrolled in account recovery
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
}

func (m MyAuthenticationFactorInitiator) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m MyAuthenticationFactorInitiator) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingMyAuthenticationFactorInitiatorAuthFactorEnum(string(m.AuthFactor)); !ok && m.AuthFactor != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for AuthFactor: %s. Supported values are: %s.", m.AuthFactor, strings.Join(GetMyAuthenticationFactorInitiatorAuthFactorEnumStringValues(), ",")))
	}

	for _, val := range m.IdcsPreventedOperations {
		if _, ok := GetMappingIdcsPreventedOperationsEnum(string(val)); !ok && val != "" {
			errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for IdcsPreventedOperations: %s. Supported values are: %s.", val, strings.Join(GetIdcsPreventedOperationsEnumStringValues(), ",")))
		}
	}

	if _, ok := GetMappingMyAuthenticationFactorInitiatorTypeEnum(string(m.Type)); !ok && m.Type != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Type: %s. Supported values are: %s.", m.Type, strings.Join(GetMyAuthenticationFactorInitiatorTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingMyAuthenticationFactorInitiatorScenarioEnum(string(m.Scenario)); !ok && m.Scenario != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Scenario: %s. Supported values are: %s.", m.Scenario, strings.Join(GetMyAuthenticationFactorInitiatorScenarioEnumStringValues(), ",")))
	}
	if _, ok := GetMappingMyAuthenticationFactorInitiatorPreferenceTypeEnum(string(m.PreferenceType)); !ok && m.PreferenceType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for PreferenceType: %s. Supported values are: %s.", m.PreferenceType, strings.Join(GetMyAuthenticationFactorInitiatorPreferenceTypeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MyAuthenticationFactorInitiatorAuthFactorEnum Enum with underlying type: string
type MyAuthenticationFactorInitiatorAuthFactorEnum string

// Set of constants representing the allowable values for MyAuthenticationFactorInitiatorAuthFactorEnum
const (
	MyAuthenticationFactorInitiatorAuthFactorEmail             MyAuthenticationFactorInitiatorAuthFactorEnum = "EMAIL"
	MyAuthenticationFactorInitiatorAuthFactorTotp              MyAuthenticationFactorInitiatorAuthFactorEnum = "TOTP"
	MyAuthenticationFactorInitiatorAuthFactorPush              MyAuthenticationFactorInitiatorAuthFactorEnum = "PUSH"
	MyAuthenticationFactorInitiatorAuthFactorSms               MyAuthenticationFactorInitiatorAuthFactorEnum = "SMS"
	MyAuthenticationFactorInitiatorAuthFactorVoice             MyAuthenticationFactorInitiatorAuthFactorEnum = "VOICE"
	MyAuthenticationFactorInitiatorAuthFactorPhoneCall         MyAuthenticationFactorInitiatorAuthFactorEnum = "PHONE_CALL"
	MyAuthenticationFactorInitiatorAuthFactorThirdparty        MyAuthenticationFactorInitiatorAuthFactorEnum = "THIRDPARTY"
	MyAuthenticationFactorInitiatorAuthFactorFidoAuthenticator MyAuthenticationFactorInitiatorAuthFactorEnum = "FIDO_AUTHENTICATOR"
	MyAuthenticationFactorInitiatorAuthFactorYubicoOtp         MyAuthenticationFactorInitiatorAuthFactorEnum = "YUBICO_OTP"
)

var mappingMyAuthenticationFactorInitiatorAuthFactorEnum = map[string]MyAuthenticationFactorInitiatorAuthFactorEnum{
	"EMAIL":              MyAuthenticationFactorInitiatorAuthFactorEmail,
	"TOTP":               MyAuthenticationFactorInitiatorAuthFactorTotp,
	"PUSH":               MyAuthenticationFactorInitiatorAuthFactorPush,
	"SMS":                MyAuthenticationFactorInitiatorAuthFactorSms,
	"VOICE":              MyAuthenticationFactorInitiatorAuthFactorVoice,
	"PHONE_CALL":         MyAuthenticationFactorInitiatorAuthFactorPhoneCall,
	"THIRDPARTY":         MyAuthenticationFactorInitiatorAuthFactorThirdparty,
	"FIDO_AUTHENTICATOR": MyAuthenticationFactorInitiatorAuthFactorFidoAuthenticator,
	"YUBICO_OTP":         MyAuthenticationFactorInitiatorAuthFactorYubicoOtp,
}

var mappingMyAuthenticationFactorInitiatorAuthFactorEnumLowerCase = map[string]MyAuthenticationFactorInitiatorAuthFactorEnum{
	"email":              MyAuthenticationFactorInitiatorAuthFactorEmail,
	"totp":               MyAuthenticationFactorInitiatorAuthFactorTotp,
	"push":               MyAuthenticationFactorInitiatorAuthFactorPush,
	"sms":                MyAuthenticationFactorInitiatorAuthFactorSms,
	"voice":              MyAuthenticationFactorInitiatorAuthFactorVoice,
	"phone_call":         MyAuthenticationFactorInitiatorAuthFactorPhoneCall,
	"thirdparty":         MyAuthenticationFactorInitiatorAuthFactorThirdparty,
	"fido_authenticator": MyAuthenticationFactorInitiatorAuthFactorFidoAuthenticator,
	"yubico_otp":         MyAuthenticationFactorInitiatorAuthFactorYubicoOtp,
}

// GetMyAuthenticationFactorInitiatorAuthFactorEnumValues Enumerates the set of values for MyAuthenticationFactorInitiatorAuthFactorEnum
func GetMyAuthenticationFactorInitiatorAuthFactorEnumValues() []MyAuthenticationFactorInitiatorAuthFactorEnum {
	values := make([]MyAuthenticationFactorInitiatorAuthFactorEnum, 0)
	for _, v := range mappingMyAuthenticationFactorInitiatorAuthFactorEnum {
		values = append(values, v)
	}
	return values
}

// GetMyAuthenticationFactorInitiatorAuthFactorEnumStringValues Enumerates the set of values in String for MyAuthenticationFactorInitiatorAuthFactorEnum
func GetMyAuthenticationFactorInitiatorAuthFactorEnumStringValues() []string {
	return []string{
		"EMAIL",
		"TOTP",
		"PUSH",
		"SMS",
		"VOICE",
		"PHONE_CALL",
		"THIRDPARTY",
		"FIDO_AUTHENTICATOR",
		"YUBICO_OTP",
	}
}

// GetMappingMyAuthenticationFactorInitiatorAuthFactorEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingMyAuthenticationFactorInitiatorAuthFactorEnum(val string) (MyAuthenticationFactorInitiatorAuthFactorEnum, bool) {
	enum, ok := mappingMyAuthenticationFactorInitiatorAuthFactorEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// MyAuthenticationFactorInitiatorTypeEnum Enum with underlying type: string
type MyAuthenticationFactorInitiatorTypeEnum string

// Set of constants representing the allowable values for MyAuthenticationFactorInitiatorTypeEnum
const (
	MyAuthenticationFactorInitiatorTypeSaml MyAuthenticationFactorInitiatorTypeEnum = "SAML"
	MyAuthenticationFactorInitiatorTypeOidc MyAuthenticationFactorInitiatorTypeEnum = "OIDC"
)

var mappingMyAuthenticationFactorInitiatorTypeEnum = map[string]MyAuthenticationFactorInitiatorTypeEnum{
	"SAML": MyAuthenticationFactorInitiatorTypeSaml,
	"OIDC": MyAuthenticationFactorInitiatorTypeOidc,
}

var mappingMyAuthenticationFactorInitiatorTypeEnumLowerCase = map[string]MyAuthenticationFactorInitiatorTypeEnum{
	"saml": MyAuthenticationFactorInitiatorTypeSaml,
	"oidc": MyAuthenticationFactorInitiatorTypeOidc,
}

// GetMyAuthenticationFactorInitiatorTypeEnumValues Enumerates the set of values for MyAuthenticationFactorInitiatorTypeEnum
func GetMyAuthenticationFactorInitiatorTypeEnumValues() []MyAuthenticationFactorInitiatorTypeEnum {
	values := make([]MyAuthenticationFactorInitiatorTypeEnum, 0)
	for _, v := range mappingMyAuthenticationFactorInitiatorTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetMyAuthenticationFactorInitiatorTypeEnumStringValues Enumerates the set of values in String for MyAuthenticationFactorInitiatorTypeEnum
func GetMyAuthenticationFactorInitiatorTypeEnumStringValues() []string {
	return []string{
		"SAML",
		"OIDC",
	}
}

// GetMappingMyAuthenticationFactorInitiatorTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingMyAuthenticationFactorInitiatorTypeEnum(val string) (MyAuthenticationFactorInitiatorTypeEnum, bool) {
	enum, ok := mappingMyAuthenticationFactorInitiatorTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// MyAuthenticationFactorInitiatorScenarioEnum Enum with underlying type: string
type MyAuthenticationFactorInitiatorScenarioEnum string

// Set of constants representing the allowable values for MyAuthenticationFactorInitiatorScenarioEnum
const (
	MyAuthenticationFactorInitiatorScenarioEnrollment     MyAuthenticationFactorInitiatorScenarioEnum = "ENROLLMENT"
	MyAuthenticationFactorInitiatorScenarioAuthentication MyAuthenticationFactorInitiatorScenarioEnum = "AUTHENTICATION"
)

var mappingMyAuthenticationFactorInitiatorScenarioEnum = map[string]MyAuthenticationFactorInitiatorScenarioEnum{
	"ENROLLMENT":     MyAuthenticationFactorInitiatorScenarioEnrollment,
	"AUTHENTICATION": MyAuthenticationFactorInitiatorScenarioAuthentication,
}

var mappingMyAuthenticationFactorInitiatorScenarioEnumLowerCase = map[string]MyAuthenticationFactorInitiatorScenarioEnum{
	"enrollment":     MyAuthenticationFactorInitiatorScenarioEnrollment,
	"authentication": MyAuthenticationFactorInitiatorScenarioAuthentication,
}

// GetMyAuthenticationFactorInitiatorScenarioEnumValues Enumerates the set of values for MyAuthenticationFactorInitiatorScenarioEnum
func GetMyAuthenticationFactorInitiatorScenarioEnumValues() []MyAuthenticationFactorInitiatorScenarioEnum {
	values := make([]MyAuthenticationFactorInitiatorScenarioEnum, 0)
	for _, v := range mappingMyAuthenticationFactorInitiatorScenarioEnum {
		values = append(values, v)
	}
	return values
}

// GetMyAuthenticationFactorInitiatorScenarioEnumStringValues Enumerates the set of values in String for MyAuthenticationFactorInitiatorScenarioEnum
func GetMyAuthenticationFactorInitiatorScenarioEnumStringValues() []string {
	return []string{
		"ENROLLMENT",
		"AUTHENTICATION",
	}
}

// GetMappingMyAuthenticationFactorInitiatorScenarioEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingMyAuthenticationFactorInitiatorScenarioEnum(val string) (MyAuthenticationFactorInitiatorScenarioEnum, bool) {
	enum, ok := mappingMyAuthenticationFactorInitiatorScenarioEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// MyAuthenticationFactorInitiatorPreferenceTypeEnum Enum with underlying type: string
type MyAuthenticationFactorInitiatorPreferenceTypeEnum string

// Set of constants representing the allowable values for MyAuthenticationFactorInitiatorPreferenceTypeEnum
const (
	MyAuthenticationFactorInitiatorPreferenceTypePasswordless MyAuthenticationFactorInitiatorPreferenceTypeEnum = "PASSWORDLESS"
	MyAuthenticationFactorInitiatorPreferenceTypeMfa          MyAuthenticationFactorInitiatorPreferenceTypeEnum = "MFA"
)

var mappingMyAuthenticationFactorInitiatorPreferenceTypeEnum = map[string]MyAuthenticationFactorInitiatorPreferenceTypeEnum{
	"PASSWORDLESS": MyAuthenticationFactorInitiatorPreferenceTypePasswordless,
	"MFA":          MyAuthenticationFactorInitiatorPreferenceTypeMfa,
}

var mappingMyAuthenticationFactorInitiatorPreferenceTypeEnumLowerCase = map[string]MyAuthenticationFactorInitiatorPreferenceTypeEnum{
	"passwordless": MyAuthenticationFactorInitiatorPreferenceTypePasswordless,
	"mfa":          MyAuthenticationFactorInitiatorPreferenceTypeMfa,
}

// GetMyAuthenticationFactorInitiatorPreferenceTypeEnumValues Enumerates the set of values for MyAuthenticationFactorInitiatorPreferenceTypeEnum
func GetMyAuthenticationFactorInitiatorPreferenceTypeEnumValues() []MyAuthenticationFactorInitiatorPreferenceTypeEnum {
	values := make([]MyAuthenticationFactorInitiatorPreferenceTypeEnum, 0)
	for _, v := range mappingMyAuthenticationFactorInitiatorPreferenceTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetMyAuthenticationFactorInitiatorPreferenceTypeEnumStringValues Enumerates the set of values in String for MyAuthenticationFactorInitiatorPreferenceTypeEnum
func GetMyAuthenticationFactorInitiatorPreferenceTypeEnumStringValues() []string {
	return []string{
		"PASSWORDLESS",
		"MFA",
	}
}

// GetMappingMyAuthenticationFactorInitiatorPreferenceTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingMyAuthenticationFactorInitiatorPreferenceTypeEnum(val string) (MyAuthenticationFactorInitiatorPreferenceTypeEnum, bool) {
	enum, ok := mappingMyAuthenticationFactorInitiatorPreferenceTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
