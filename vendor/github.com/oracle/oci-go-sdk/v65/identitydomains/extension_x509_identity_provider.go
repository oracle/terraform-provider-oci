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

// ExtensionX509IdentityProvider X509 Identity Provider Extension Schema
type ExtensionX509IdentityProvider struct {

	// X509 Certificate Matching Attribute
	// **Added In:** 2010242156
	// **SCIM++ Properties:**
	//  - caseExact: false
	//  - idcsSearchable: false
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: true
	//  - returned: default
	//  - type: string
	//  - uniqueness: none
	CertMatchAttribute *string `mandatory:"true" json:"certMatchAttribute"`

	// This property specifies the userstore attribute value that must match the incoming certificate attribute.
	// **Added In:** 2010242156
	// **SCIM++ Properties:**
	//  - caseExact: false
	//  - idcsSearchable: false
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: true
	//  - returned: default
	//  - type: string
	//  - uniqueness: none
	UserMatchAttribute *string `mandatory:"true" json:"userMatchAttribute"`

	// Certificate alias list to create a chain for the incoming client certificate
	// **Added In:** 2010242156
	// **SCIM++ Properties:**
	//  - caseExact: false
	//  - idcsSearchable: false
	//  - multiValued: true
	//  - mutability: readWrite
	//  - required: true
	//  - returned: default
	//  - type: string
	//  - uniqueness: none
	SigningCertificateChain []string `mandatory:"true" json:"signingCertificateChain"`

	// Check for specific conditions of other certificate attributes
	// **Added In:** 2010242156
	// **SCIM++ Properties:**
	//  - caseExact: false
	//  - idcsSearchable: false
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: false
	//  - returned: default
	//  - type: string
	//  - uniqueness: none
	OtherCertMatchAttribute *string `mandatory:"false" json:"otherCertMatchAttribute"`

	// Set to true to enable OCSP Validation
	// **Added In:** 2010242156
	// **SCIM++ Properties:**
	//  - caseExact: false
	//  - idcsSearchable: false
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: false
	//  - returned: default
	//  - type: boolean
	//  - uniqueness: none
	OcspEnabled *bool `mandatory:"false" json:"ocspEnabled"`

	// This property specifies the OCSP Server alias name
	// **Added In:** 2010242156
	// **SCIM++ Properties:**
	//  - caseExact: false
	//  - idcsSearchable: false
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: false
	//  - returned: default
	//  - type: string
	//  - uniqueness: none
	OcspServerName *string `mandatory:"false" json:"ocspServerName"`

	// This property specifies OCSP Responder URL.
	// **Added In:** 2010242156
	// **SCIM++ Properties:**
	//  - caseExact: false
	//  - idcsSearchable: false
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: false
	//  - returned: default
	//  - type: string
	//  - uniqueness: none
	OcspResponderURL *string `mandatory:"false" json:"ocspResponderURL"`

	// Allow access if OCSP response is UNKNOWN or OCSP Responder does not respond within the timeout duration
	// **Added In:** 2010242156
	// **SCIM++ Properties:**
	//  - caseExact: false
	//  - idcsSearchable: false
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: false
	//  - returned: default
	//  - type: boolean
	//  - uniqueness: none
	OcspAllowUnknownResponseStatus *bool `mandatory:"false" json:"ocspAllowUnknownResponseStatus"`

	// Revalidate OCSP status for user after X hours
	// **Added In:** 2010242156
	// **SCIM++ Properties:**
	//  - idcsMaxValue: 24
	//  - idcsMinValue: 0
	//  - idcsSearchable: false
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: false
	//  - returned: default
	//  - type: integer
	//  - uniqueness: none
	OcspRevalidateTime *int `mandatory:"false" json:"ocspRevalidateTime"`

	// Describes if the OCSP response is signed
	// **Added In:** 2010242156
	// **SCIM++ Properties:**
	//  - caseExact: false
	//  - idcsSearchable: false
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: false
	//  - returned: default
	//  - type: boolean
	//  - uniqueness: none
	OcspEnableSignedResponse *bool `mandatory:"false" json:"ocspEnableSignedResponse"`

	// OCSP Trusted Certificate Chain
	// **Added In:** 2010242156
	// **SCIM++ Properties:**
	//  - caseExact: false
	//  - idcsSearchable: false
	//  - multiValued: true
	//  - mutability: readWrite
	//  - required: false
	//  - returned: default
	//  - type: string
	//  - uniqueness: none
	OcspTrustCertChain []string `mandatory:"false" json:"ocspTrustCertChain"`

	// Set to true to enable CRL Validation
	// **Added In:** 2010242156
	// **SCIM++ Properties:**
	//  - caseExact: false
	//  - idcsSearchable: false
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: false
	//  - returned: default
	//  - type: boolean
	//  - uniqueness: none
	CrlEnabled *bool `mandatory:"false" json:"crlEnabled"`

	// Fallback on CRL Validation if OCSP fails.
	// **Added In:** 2010242156
	// **SCIM++ Properties:**
	//  - caseExact: false
	//  - idcsSearchable: false
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: false
	//  - returned: default
	//  - type: boolean
	//  - uniqueness: none
	CrlCheckOnOCSPFailureEnabled *bool `mandatory:"false" json:"crlCheckOnOCSPFailureEnabled"`

	// CRL Location URL
	// **Added In:** 2010242156
	// **SCIM++ Properties:**
	//  - caseExact: false
	//  - idcsSearchable: false
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: false
	//  - returned: default
	//  - type: string
	//  - uniqueness: none
	CrlLocation *string `mandatory:"false" json:"crlLocation"`

	// Fetch the CRL contents every X minutes
	// **Added In:** 2010242156
	// **SCIM++ Properties:**
	//  - idcsSearchable: false
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: false
	//  - returned: default
	//  - type: integer
	//  - uniqueness: none
	CrlReloadDuration *int `mandatory:"false" json:"crlReloadDuration"`

	// Set to true to enable EKU Validation
	// **Added In:** 2304270343
	// **SCIM++ Properties:**
	//  - caseExact: false
	//  - idcsSearchable: false
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: false
	//  - returned: default
	//  - type: boolean
	//  - uniqueness: none
	EkuValidationEnabled *bool `mandatory:"false" json:"ekuValidationEnabled"`

	// List of EKU which needs to be validated
	// **Added In:** 2304270343
	// **SCIM++ Properties:**
	//  - caseExact: false
	//  - idcsSearchable: false
	//  - multiValued: true
	//  - mutability: readWrite
	//  - required: false
	//  - returned: default
	//  - type: string
	//  - uniqueness: none
	EkuValues []ExtensionX509IdentityProviderEkuValuesEnum `mandatory:"false" json:"ekuValues,omitempty"`
}

func (m ExtensionX509IdentityProvider) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ExtensionX509IdentityProvider) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	for _, val := range m.EkuValues {
		if _, ok := GetMappingExtensionX509IdentityProviderEkuValuesEnum(string(val)); !ok && val != "" {
			errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for EkuValues: %s. Supported values are: %s.", val, strings.Join(GetExtensionX509IdentityProviderEkuValuesEnumStringValues(), ",")))
		}
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ExtensionX509IdentityProviderEkuValuesEnum Enum with underlying type: string
type ExtensionX509IdentityProviderEkuValuesEnum string

// Set of constants representing the allowable values for ExtensionX509IdentityProviderEkuValuesEnum
const (
	ExtensionX509IdentityProviderEkuValuesServerAuth      ExtensionX509IdentityProviderEkuValuesEnum = "SERVER_AUTH"
	ExtensionX509IdentityProviderEkuValuesClientAuth      ExtensionX509IdentityProviderEkuValuesEnum = "CLIENT_AUTH"
	ExtensionX509IdentityProviderEkuValuesCodeSigning     ExtensionX509IdentityProviderEkuValuesEnum = "CODE_SIGNING"
	ExtensionX509IdentityProviderEkuValuesEmailProtection ExtensionX509IdentityProviderEkuValuesEnum = "EMAIL_PROTECTION"
	ExtensionX509IdentityProviderEkuValuesTimeStamping    ExtensionX509IdentityProviderEkuValuesEnum = "TIME_STAMPING"
	ExtensionX509IdentityProviderEkuValuesOcspSigning     ExtensionX509IdentityProviderEkuValuesEnum = "OCSP_SIGNING"
)

var mappingExtensionX509IdentityProviderEkuValuesEnum = map[string]ExtensionX509IdentityProviderEkuValuesEnum{
	"SERVER_AUTH":      ExtensionX509IdentityProviderEkuValuesServerAuth,
	"CLIENT_AUTH":      ExtensionX509IdentityProviderEkuValuesClientAuth,
	"CODE_SIGNING":     ExtensionX509IdentityProviderEkuValuesCodeSigning,
	"EMAIL_PROTECTION": ExtensionX509IdentityProviderEkuValuesEmailProtection,
	"TIME_STAMPING":    ExtensionX509IdentityProviderEkuValuesTimeStamping,
	"OCSP_SIGNING":     ExtensionX509IdentityProviderEkuValuesOcspSigning,
}

var mappingExtensionX509IdentityProviderEkuValuesEnumLowerCase = map[string]ExtensionX509IdentityProviderEkuValuesEnum{
	"server_auth":      ExtensionX509IdentityProviderEkuValuesServerAuth,
	"client_auth":      ExtensionX509IdentityProviderEkuValuesClientAuth,
	"code_signing":     ExtensionX509IdentityProviderEkuValuesCodeSigning,
	"email_protection": ExtensionX509IdentityProviderEkuValuesEmailProtection,
	"time_stamping":    ExtensionX509IdentityProviderEkuValuesTimeStamping,
	"ocsp_signing":     ExtensionX509IdentityProviderEkuValuesOcspSigning,
}

// GetExtensionX509IdentityProviderEkuValuesEnumValues Enumerates the set of values for ExtensionX509IdentityProviderEkuValuesEnum
func GetExtensionX509IdentityProviderEkuValuesEnumValues() []ExtensionX509IdentityProviderEkuValuesEnum {
	values := make([]ExtensionX509IdentityProviderEkuValuesEnum, 0)
	for _, v := range mappingExtensionX509IdentityProviderEkuValuesEnum {
		values = append(values, v)
	}
	return values
}

// GetExtensionX509IdentityProviderEkuValuesEnumStringValues Enumerates the set of values in String for ExtensionX509IdentityProviderEkuValuesEnum
func GetExtensionX509IdentityProviderEkuValuesEnumStringValues() []string {
	return []string{
		"SERVER_AUTH",
		"CLIENT_AUTH",
		"CODE_SIGNING",
		"EMAIL_PROTECTION",
		"TIME_STAMPING",
		"OCSP_SIGNING",
	}
}

// GetMappingExtensionX509IdentityProviderEkuValuesEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingExtensionX509IdentityProviderEkuValuesEnum(val string) (ExtensionX509IdentityProviderEkuValuesEnum, bool) {
	enum, ok := mappingExtensionX509IdentityProviderEkuValuesEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
