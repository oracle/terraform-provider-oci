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

// ExtensionFidoAuthenticationFactorSettings This extension defines attributes used to manage Multi-Factor Authentication settings of fido authentication
type ExtensionFidoAuthenticationFactorSettings struct {

	// Attribute used to define the type of attestation required.
	// **Added In:** 2009232244
	// **SCIM++ Properties:**
	//  - idcsSearchable: false
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: true
	//  - returned: default
	//  - type: string
	//  - uniqueness: none
	Attestation ExtensionFidoAuthenticationFactorSettingsAttestationEnum `mandatory:"true" json:"attestation"`

	// Attribute used to define authenticator selection attachment.
	// **Added In:** 2009232244
	// **SCIM++ Properties:**
	//  - idcsSearchable: false
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: true
	//  - returned: default
	//  - type: string
	//  - uniqueness: none
	AuthenticatorSelectionAttachment ExtensionFidoAuthenticationFactorSettingsAuthenticatorSelectionAttachmentEnum `mandatory:"true" json:"authenticatorSelectionAttachment"`

	// Attribute used to define authenticator selection verification.
	// **Added In:** 2009232244
	// **SCIM++ Properties:**
	//  - idcsSearchable: false
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: true
	//  - returned: default
	//  - type: string
	//  - uniqueness: none
	AuthenticatorSelectionUserVerification ExtensionFidoAuthenticationFactorSettingsAuthenticatorSelectionUserVerificationEnum `mandatory:"true" json:"authenticatorSelectionUserVerification"`

	// Attribute used to define authenticator selection resident key requirement.
	// **Added In:** 2009232244
	// **SCIM++ Properties:**
	//  - idcsSearchable: false
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: true
	//  - returned: default
	//  - type: string
	//  - uniqueness: none
	AuthenticatorSelectionResidentKey ExtensionFidoAuthenticationFactorSettingsAuthenticatorSelectionResidentKeyEnum `mandatory:"true" json:"authenticatorSelectionResidentKey"`

	// Timeout for the fido authentication to complete
	// **Added In:** 2009232244
	// **SCIM++ Properties:**
	//  - idcsSearchable: false
	//  - multiValued: false
	//  - mutability: readWrite
	//  - idcsMaxValue: 600000
	//  - idcsMinValue: 10000
	//  - required: true
	//  - returned: default
	//  - type: integer
	//  - uniqueness: none
	Timeout *int `mandatory:"true" json:"timeout"`

	// Flag used to indicate authenticator selection is required or not
	// **Added In:** 2009232244
	// **SCIM++ Properties:**
	//  - idcsSearchable: false
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: true
	//  - returned: default
	//  - type: boolean
	//  - uniqueness: none
	AuthenticatorSelectionRequireResidentKey *bool `mandatory:"true" json:"authenticatorSelectionRequireResidentKey"`

	// List of server supported public key algorithms
	// **Added In:** 2009232244
	// **SCIM++ Properties:**
	//  - idcsSearchable: false
	//  - multiValued: true
	//  - mutability: readWrite
	//  - required: true
	//  - returned: default
	//  - type: string
	//  - uniqueness: none
	PublicKeyTypes []ExtensionFidoAuthenticationFactorSettingsPublicKeyTypesEnum `mandatory:"true" json:"publicKeyTypes"`

	// Flag used to indicate whether we need to restrict creation of multiple credentials in same authenticator
	// **Added In:** 2009232244
	// **SCIM++ Properties:**
	//  - idcsSearchable: false
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: true
	//  - returned: default
	//  - type: boolean
	//  - uniqueness: none
	ExcludeCredentials *bool `mandatory:"true" json:"excludeCredentials"`

	// Number of domain levels Oracle Identity Cloud Service should use for origin comparision
	// **Added In:** 2109020413
	// **SCIM++ Properties:**
	//  - idcsSearchable: false
	//  - multiValued: false
	//  - mutability: readWrite
	//  - idcsMaxValue: 2
	//  - idcsMinValue: 0
	//  - required: false
	//  - returned: default
	//  - type: integer
	//  - uniqueness: none
	DomainValidationLevel *int `mandatory:"false" json:"domainValidationLevel"`
}

func (m ExtensionFidoAuthenticationFactorSettings) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ExtensionFidoAuthenticationFactorSettings) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingExtensionFidoAuthenticationFactorSettingsAttestationEnum(string(m.Attestation)); !ok && m.Attestation != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Attestation: %s. Supported values are: %s.", m.Attestation, strings.Join(GetExtensionFidoAuthenticationFactorSettingsAttestationEnumStringValues(), ",")))
	}
	if _, ok := GetMappingExtensionFidoAuthenticationFactorSettingsAuthenticatorSelectionAttachmentEnum(string(m.AuthenticatorSelectionAttachment)); !ok && m.AuthenticatorSelectionAttachment != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for AuthenticatorSelectionAttachment: %s. Supported values are: %s.", m.AuthenticatorSelectionAttachment, strings.Join(GetExtensionFidoAuthenticationFactorSettingsAuthenticatorSelectionAttachmentEnumStringValues(), ",")))
	}
	if _, ok := GetMappingExtensionFidoAuthenticationFactorSettingsAuthenticatorSelectionUserVerificationEnum(string(m.AuthenticatorSelectionUserVerification)); !ok && m.AuthenticatorSelectionUserVerification != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for AuthenticatorSelectionUserVerification: %s. Supported values are: %s.", m.AuthenticatorSelectionUserVerification, strings.Join(GetExtensionFidoAuthenticationFactorSettingsAuthenticatorSelectionUserVerificationEnumStringValues(), ",")))
	}
	if _, ok := GetMappingExtensionFidoAuthenticationFactorSettingsAuthenticatorSelectionResidentKeyEnum(string(m.AuthenticatorSelectionResidentKey)); !ok && m.AuthenticatorSelectionResidentKey != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for AuthenticatorSelectionResidentKey: %s. Supported values are: %s.", m.AuthenticatorSelectionResidentKey, strings.Join(GetExtensionFidoAuthenticationFactorSettingsAuthenticatorSelectionResidentKeyEnumStringValues(), ",")))
	}
	for _, val := range m.PublicKeyTypes {
		if _, ok := GetMappingExtensionFidoAuthenticationFactorSettingsPublicKeyTypesEnum(string(val)); !ok && val != "" {
			errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for PublicKeyTypes: %s. Supported values are: %s.", val, strings.Join(GetExtensionFidoAuthenticationFactorSettingsPublicKeyTypesEnumStringValues(), ",")))
		}
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ExtensionFidoAuthenticationFactorSettingsAttestationEnum Enum with underlying type: string
type ExtensionFidoAuthenticationFactorSettingsAttestationEnum string

// Set of constants representing the allowable values for ExtensionFidoAuthenticationFactorSettingsAttestationEnum
const (
	ExtensionFidoAuthenticationFactorSettingsAttestationNone     ExtensionFidoAuthenticationFactorSettingsAttestationEnum = "NONE"
	ExtensionFidoAuthenticationFactorSettingsAttestationDirect   ExtensionFidoAuthenticationFactorSettingsAttestationEnum = "DIRECT"
	ExtensionFidoAuthenticationFactorSettingsAttestationIndirect ExtensionFidoAuthenticationFactorSettingsAttestationEnum = "INDIRECT"
)

var mappingExtensionFidoAuthenticationFactorSettingsAttestationEnum = map[string]ExtensionFidoAuthenticationFactorSettingsAttestationEnum{
	"NONE":     ExtensionFidoAuthenticationFactorSettingsAttestationNone,
	"DIRECT":   ExtensionFidoAuthenticationFactorSettingsAttestationDirect,
	"INDIRECT": ExtensionFidoAuthenticationFactorSettingsAttestationIndirect,
}

var mappingExtensionFidoAuthenticationFactorSettingsAttestationEnumLowerCase = map[string]ExtensionFidoAuthenticationFactorSettingsAttestationEnum{
	"none":     ExtensionFidoAuthenticationFactorSettingsAttestationNone,
	"direct":   ExtensionFidoAuthenticationFactorSettingsAttestationDirect,
	"indirect": ExtensionFidoAuthenticationFactorSettingsAttestationIndirect,
}

// GetExtensionFidoAuthenticationFactorSettingsAttestationEnumValues Enumerates the set of values for ExtensionFidoAuthenticationFactorSettingsAttestationEnum
func GetExtensionFidoAuthenticationFactorSettingsAttestationEnumValues() []ExtensionFidoAuthenticationFactorSettingsAttestationEnum {
	values := make([]ExtensionFidoAuthenticationFactorSettingsAttestationEnum, 0)
	for _, v := range mappingExtensionFidoAuthenticationFactorSettingsAttestationEnum {
		values = append(values, v)
	}
	return values
}

// GetExtensionFidoAuthenticationFactorSettingsAttestationEnumStringValues Enumerates the set of values in String for ExtensionFidoAuthenticationFactorSettingsAttestationEnum
func GetExtensionFidoAuthenticationFactorSettingsAttestationEnumStringValues() []string {
	return []string{
		"NONE",
		"DIRECT",
		"INDIRECT",
	}
}

// GetMappingExtensionFidoAuthenticationFactorSettingsAttestationEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingExtensionFidoAuthenticationFactorSettingsAttestationEnum(val string) (ExtensionFidoAuthenticationFactorSettingsAttestationEnum, bool) {
	enum, ok := mappingExtensionFidoAuthenticationFactorSettingsAttestationEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ExtensionFidoAuthenticationFactorSettingsAuthenticatorSelectionAttachmentEnum Enum with underlying type: string
type ExtensionFidoAuthenticationFactorSettingsAuthenticatorSelectionAttachmentEnum string

// Set of constants representing the allowable values for ExtensionFidoAuthenticationFactorSettingsAuthenticatorSelectionAttachmentEnum
const (
	ExtensionFidoAuthenticationFactorSettingsAuthenticatorSelectionAttachmentPlatform      ExtensionFidoAuthenticationFactorSettingsAuthenticatorSelectionAttachmentEnum = "PLATFORM"
	ExtensionFidoAuthenticationFactorSettingsAuthenticatorSelectionAttachmentCrossPlatform ExtensionFidoAuthenticationFactorSettingsAuthenticatorSelectionAttachmentEnum = "CROSS-PLATFORM"
	ExtensionFidoAuthenticationFactorSettingsAuthenticatorSelectionAttachmentBoth          ExtensionFidoAuthenticationFactorSettingsAuthenticatorSelectionAttachmentEnum = "BOTH"
)

var mappingExtensionFidoAuthenticationFactorSettingsAuthenticatorSelectionAttachmentEnum = map[string]ExtensionFidoAuthenticationFactorSettingsAuthenticatorSelectionAttachmentEnum{
	"PLATFORM":       ExtensionFidoAuthenticationFactorSettingsAuthenticatorSelectionAttachmentPlatform,
	"CROSS-PLATFORM": ExtensionFidoAuthenticationFactorSettingsAuthenticatorSelectionAttachmentCrossPlatform,
	"BOTH":           ExtensionFidoAuthenticationFactorSettingsAuthenticatorSelectionAttachmentBoth,
}

var mappingExtensionFidoAuthenticationFactorSettingsAuthenticatorSelectionAttachmentEnumLowerCase = map[string]ExtensionFidoAuthenticationFactorSettingsAuthenticatorSelectionAttachmentEnum{
	"platform":       ExtensionFidoAuthenticationFactorSettingsAuthenticatorSelectionAttachmentPlatform,
	"cross-platform": ExtensionFidoAuthenticationFactorSettingsAuthenticatorSelectionAttachmentCrossPlatform,
	"both":           ExtensionFidoAuthenticationFactorSettingsAuthenticatorSelectionAttachmentBoth,
}

// GetExtensionFidoAuthenticationFactorSettingsAuthenticatorSelectionAttachmentEnumValues Enumerates the set of values for ExtensionFidoAuthenticationFactorSettingsAuthenticatorSelectionAttachmentEnum
func GetExtensionFidoAuthenticationFactorSettingsAuthenticatorSelectionAttachmentEnumValues() []ExtensionFidoAuthenticationFactorSettingsAuthenticatorSelectionAttachmentEnum {
	values := make([]ExtensionFidoAuthenticationFactorSettingsAuthenticatorSelectionAttachmentEnum, 0)
	for _, v := range mappingExtensionFidoAuthenticationFactorSettingsAuthenticatorSelectionAttachmentEnum {
		values = append(values, v)
	}
	return values
}

// GetExtensionFidoAuthenticationFactorSettingsAuthenticatorSelectionAttachmentEnumStringValues Enumerates the set of values in String for ExtensionFidoAuthenticationFactorSettingsAuthenticatorSelectionAttachmentEnum
func GetExtensionFidoAuthenticationFactorSettingsAuthenticatorSelectionAttachmentEnumStringValues() []string {
	return []string{
		"PLATFORM",
		"CROSS-PLATFORM",
		"BOTH",
	}
}

// GetMappingExtensionFidoAuthenticationFactorSettingsAuthenticatorSelectionAttachmentEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingExtensionFidoAuthenticationFactorSettingsAuthenticatorSelectionAttachmentEnum(val string) (ExtensionFidoAuthenticationFactorSettingsAuthenticatorSelectionAttachmentEnum, bool) {
	enum, ok := mappingExtensionFidoAuthenticationFactorSettingsAuthenticatorSelectionAttachmentEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ExtensionFidoAuthenticationFactorSettingsAuthenticatorSelectionUserVerificationEnum Enum with underlying type: string
type ExtensionFidoAuthenticationFactorSettingsAuthenticatorSelectionUserVerificationEnum string

// Set of constants representing the allowable values for ExtensionFidoAuthenticationFactorSettingsAuthenticatorSelectionUserVerificationEnum
const (
	ExtensionFidoAuthenticationFactorSettingsAuthenticatorSelectionUserVerificationRequired    ExtensionFidoAuthenticationFactorSettingsAuthenticatorSelectionUserVerificationEnum = "REQUIRED"
	ExtensionFidoAuthenticationFactorSettingsAuthenticatorSelectionUserVerificationPreferred   ExtensionFidoAuthenticationFactorSettingsAuthenticatorSelectionUserVerificationEnum = "PREFERRED"
	ExtensionFidoAuthenticationFactorSettingsAuthenticatorSelectionUserVerificationDiscouraged ExtensionFidoAuthenticationFactorSettingsAuthenticatorSelectionUserVerificationEnum = "DISCOURAGED"
)

var mappingExtensionFidoAuthenticationFactorSettingsAuthenticatorSelectionUserVerificationEnum = map[string]ExtensionFidoAuthenticationFactorSettingsAuthenticatorSelectionUserVerificationEnum{
	"REQUIRED":    ExtensionFidoAuthenticationFactorSettingsAuthenticatorSelectionUserVerificationRequired,
	"PREFERRED":   ExtensionFidoAuthenticationFactorSettingsAuthenticatorSelectionUserVerificationPreferred,
	"DISCOURAGED": ExtensionFidoAuthenticationFactorSettingsAuthenticatorSelectionUserVerificationDiscouraged,
}

var mappingExtensionFidoAuthenticationFactorSettingsAuthenticatorSelectionUserVerificationEnumLowerCase = map[string]ExtensionFidoAuthenticationFactorSettingsAuthenticatorSelectionUserVerificationEnum{
	"required":    ExtensionFidoAuthenticationFactorSettingsAuthenticatorSelectionUserVerificationRequired,
	"preferred":   ExtensionFidoAuthenticationFactorSettingsAuthenticatorSelectionUserVerificationPreferred,
	"discouraged": ExtensionFidoAuthenticationFactorSettingsAuthenticatorSelectionUserVerificationDiscouraged,
}

// GetExtensionFidoAuthenticationFactorSettingsAuthenticatorSelectionUserVerificationEnumValues Enumerates the set of values for ExtensionFidoAuthenticationFactorSettingsAuthenticatorSelectionUserVerificationEnum
func GetExtensionFidoAuthenticationFactorSettingsAuthenticatorSelectionUserVerificationEnumValues() []ExtensionFidoAuthenticationFactorSettingsAuthenticatorSelectionUserVerificationEnum {
	values := make([]ExtensionFidoAuthenticationFactorSettingsAuthenticatorSelectionUserVerificationEnum, 0)
	for _, v := range mappingExtensionFidoAuthenticationFactorSettingsAuthenticatorSelectionUserVerificationEnum {
		values = append(values, v)
	}
	return values
}

// GetExtensionFidoAuthenticationFactorSettingsAuthenticatorSelectionUserVerificationEnumStringValues Enumerates the set of values in String for ExtensionFidoAuthenticationFactorSettingsAuthenticatorSelectionUserVerificationEnum
func GetExtensionFidoAuthenticationFactorSettingsAuthenticatorSelectionUserVerificationEnumStringValues() []string {
	return []string{
		"REQUIRED",
		"PREFERRED",
		"DISCOURAGED",
	}
}

// GetMappingExtensionFidoAuthenticationFactorSettingsAuthenticatorSelectionUserVerificationEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingExtensionFidoAuthenticationFactorSettingsAuthenticatorSelectionUserVerificationEnum(val string) (ExtensionFidoAuthenticationFactorSettingsAuthenticatorSelectionUserVerificationEnum, bool) {
	enum, ok := mappingExtensionFidoAuthenticationFactorSettingsAuthenticatorSelectionUserVerificationEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ExtensionFidoAuthenticationFactorSettingsAuthenticatorSelectionResidentKeyEnum Enum with underlying type: string
type ExtensionFidoAuthenticationFactorSettingsAuthenticatorSelectionResidentKeyEnum string

// Set of constants representing the allowable values for ExtensionFidoAuthenticationFactorSettingsAuthenticatorSelectionResidentKeyEnum
const (
	ExtensionFidoAuthenticationFactorSettingsAuthenticatorSelectionResidentKeyRequired    ExtensionFidoAuthenticationFactorSettingsAuthenticatorSelectionResidentKeyEnum = "REQUIRED"
	ExtensionFidoAuthenticationFactorSettingsAuthenticatorSelectionResidentKeyPreferred   ExtensionFidoAuthenticationFactorSettingsAuthenticatorSelectionResidentKeyEnum = "PREFERRED"
	ExtensionFidoAuthenticationFactorSettingsAuthenticatorSelectionResidentKeyDiscouraged ExtensionFidoAuthenticationFactorSettingsAuthenticatorSelectionResidentKeyEnum = "DISCOURAGED"
	ExtensionFidoAuthenticationFactorSettingsAuthenticatorSelectionResidentKeyNone        ExtensionFidoAuthenticationFactorSettingsAuthenticatorSelectionResidentKeyEnum = "NONE"
)

var mappingExtensionFidoAuthenticationFactorSettingsAuthenticatorSelectionResidentKeyEnum = map[string]ExtensionFidoAuthenticationFactorSettingsAuthenticatorSelectionResidentKeyEnum{
	"REQUIRED":    ExtensionFidoAuthenticationFactorSettingsAuthenticatorSelectionResidentKeyRequired,
	"PREFERRED":   ExtensionFidoAuthenticationFactorSettingsAuthenticatorSelectionResidentKeyPreferred,
	"DISCOURAGED": ExtensionFidoAuthenticationFactorSettingsAuthenticatorSelectionResidentKeyDiscouraged,
	"NONE":        ExtensionFidoAuthenticationFactorSettingsAuthenticatorSelectionResidentKeyNone,
}

var mappingExtensionFidoAuthenticationFactorSettingsAuthenticatorSelectionResidentKeyEnumLowerCase = map[string]ExtensionFidoAuthenticationFactorSettingsAuthenticatorSelectionResidentKeyEnum{
	"required":    ExtensionFidoAuthenticationFactorSettingsAuthenticatorSelectionResidentKeyRequired,
	"preferred":   ExtensionFidoAuthenticationFactorSettingsAuthenticatorSelectionResidentKeyPreferred,
	"discouraged": ExtensionFidoAuthenticationFactorSettingsAuthenticatorSelectionResidentKeyDiscouraged,
	"none":        ExtensionFidoAuthenticationFactorSettingsAuthenticatorSelectionResidentKeyNone,
}

// GetExtensionFidoAuthenticationFactorSettingsAuthenticatorSelectionResidentKeyEnumValues Enumerates the set of values for ExtensionFidoAuthenticationFactorSettingsAuthenticatorSelectionResidentKeyEnum
func GetExtensionFidoAuthenticationFactorSettingsAuthenticatorSelectionResidentKeyEnumValues() []ExtensionFidoAuthenticationFactorSettingsAuthenticatorSelectionResidentKeyEnum {
	values := make([]ExtensionFidoAuthenticationFactorSettingsAuthenticatorSelectionResidentKeyEnum, 0)
	for _, v := range mappingExtensionFidoAuthenticationFactorSettingsAuthenticatorSelectionResidentKeyEnum {
		values = append(values, v)
	}
	return values
}

// GetExtensionFidoAuthenticationFactorSettingsAuthenticatorSelectionResidentKeyEnumStringValues Enumerates the set of values in String for ExtensionFidoAuthenticationFactorSettingsAuthenticatorSelectionResidentKeyEnum
func GetExtensionFidoAuthenticationFactorSettingsAuthenticatorSelectionResidentKeyEnumStringValues() []string {
	return []string{
		"REQUIRED",
		"PREFERRED",
		"DISCOURAGED",
		"NONE",
	}
}

// GetMappingExtensionFidoAuthenticationFactorSettingsAuthenticatorSelectionResidentKeyEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingExtensionFidoAuthenticationFactorSettingsAuthenticatorSelectionResidentKeyEnum(val string) (ExtensionFidoAuthenticationFactorSettingsAuthenticatorSelectionResidentKeyEnum, bool) {
	enum, ok := mappingExtensionFidoAuthenticationFactorSettingsAuthenticatorSelectionResidentKeyEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ExtensionFidoAuthenticationFactorSettingsPublicKeyTypesEnum Enum with underlying type: string
type ExtensionFidoAuthenticationFactorSettingsPublicKeyTypesEnum string

// Set of constants representing the allowable values for ExtensionFidoAuthenticationFactorSettingsPublicKeyTypesEnum
const (
	ExtensionFidoAuthenticationFactorSettingsPublicKeyTypesRs1   ExtensionFidoAuthenticationFactorSettingsPublicKeyTypesEnum = "RS1"
	ExtensionFidoAuthenticationFactorSettingsPublicKeyTypesRs256 ExtensionFidoAuthenticationFactorSettingsPublicKeyTypesEnum = "RS256"
	ExtensionFidoAuthenticationFactorSettingsPublicKeyTypesEs256 ExtensionFidoAuthenticationFactorSettingsPublicKeyTypesEnum = "ES256"
)

var mappingExtensionFidoAuthenticationFactorSettingsPublicKeyTypesEnum = map[string]ExtensionFidoAuthenticationFactorSettingsPublicKeyTypesEnum{
	"RS1":   ExtensionFidoAuthenticationFactorSettingsPublicKeyTypesRs1,
	"RS256": ExtensionFidoAuthenticationFactorSettingsPublicKeyTypesRs256,
	"ES256": ExtensionFidoAuthenticationFactorSettingsPublicKeyTypesEs256,
}

var mappingExtensionFidoAuthenticationFactorSettingsPublicKeyTypesEnumLowerCase = map[string]ExtensionFidoAuthenticationFactorSettingsPublicKeyTypesEnum{
	"rs1":   ExtensionFidoAuthenticationFactorSettingsPublicKeyTypesRs1,
	"rs256": ExtensionFidoAuthenticationFactorSettingsPublicKeyTypesRs256,
	"es256": ExtensionFidoAuthenticationFactorSettingsPublicKeyTypesEs256,
}

// GetExtensionFidoAuthenticationFactorSettingsPublicKeyTypesEnumValues Enumerates the set of values for ExtensionFidoAuthenticationFactorSettingsPublicKeyTypesEnum
func GetExtensionFidoAuthenticationFactorSettingsPublicKeyTypesEnumValues() []ExtensionFidoAuthenticationFactorSettingsPublicKeyTypesEnum {
	values := make([]ExtensionFidoAuthenticationFactorSettingsPublicKeyTypesEnum, 0)
	for _, v := range mappingExtensionFidoAuthenticationFactorSettingsPublicKeyTypesEnum {
		values = append(values, v)
	}
	return values
}

// GetExtensionFidoAuthenticationFactorSettingsPublicKeyTypesEnumStringValues Enumerates the set of values in String for ExtensionFidoAuthenticationFactorSettingsPublicKeyTypesEnum
func GetExtensionFidoAuthenticationFactorSettingsPublicKeyTypesEnumStringValues() []string {
	return []string{
		"RS1",
		"RS256",
		"ES256",
	}
}

// GetMappingExtensionFidoAuthenticationFactorSettingsPublicKeyTypesEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingExtensionFidoAuthenticationFactorSettingsPublicKeyTypesEnum(val string) (ExtensionFidoAuthenticationFactorSettingsPublicKeyTypesEnum, bool) {
	enum, ok := mappingExtensionFidoAuthenticationFactorSettingsPublicKeyTypesEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
