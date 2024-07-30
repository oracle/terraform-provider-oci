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

// AppExtensionSamlServiceProviderApp This extension defines attributes related to the Service Providers configuration.
type AppExtensionSamlServiceProviderApp struct {

	// This attribute represents the metadata of a Security Provider in the Security Assertion Markup Language protocol.
	// **SCIM++ Properties:**
	//  - idcsSearchable: false
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: false
	//  - returned: default
	//  - type: string
	//  - uniqueness: none
	Metadata *string `mandatory:"false" json:"metadata"`

	// The ID of the Provider. This value corresponds to the entityID from the Service Provider metadata.
	// **SCIM++ Properties:**
	//  - caseExact: false
	//  - idcsSearchable: true
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: false
	//  - returned: default
	//  - type: string
	//  - uniqueness: none
	PartnerProviderId *string `mandatory:"false" json:"partnerProviderId"`

	// The pattern of the Provider. This value corresponds to the entityID from the Service Provider metadata.
	// **Added In:** 2202230830
	// **SCIM++ Properties:**
	//  - caseExact: false
	//  - idcsSearchable: false
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: false
	//  - returned: request
	//  - type: string
	//  - uniqueness: none
	PartnerProviderPattern *string `mandatory:"false" json:"partnerProviderPattern"`

	// The alternate Provider ID to be used as the Oracle Identity Cloud Service providerID (instead of the one in SamlSettings) when interacting with this SP.
	// **Added In:** 19.2.1
	// **SCIM++ Properties:**
	//  - caseExact: true
	//  - idcsSearchable: false
	//  - multiValued: false
	//  - mutability: readOnly
	//  - required: false
	//  - returned: default
	//  - type: string
	//  - uniqueness: none
	TenantProviderId *string `mandatory:"false" json:"tenantProviderId"`

	// This attribute represents the Succinct ID.
	// **SCIM++ Properties:**
	//  - caseExact: true
	//  - idcsSearchable: true
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: false
	//  - returned: default
	//  - type: string
	//  - uniqueness: server
	SuccinctId *string `mandatory:"false" json:"succinctId"`

	// The attribute represents the URL to which the SAML Assertions will be sent by the SAML IdP.
	// **SCIM++ Properties:**
	//  - caseExact: false
	//  - idcsSearchable: false
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: false
	//  - returned: default
	//  - type: string
	//  - uniqueness: none
	AssertionConsumerUrl *string `mandatory:"false" json:"assertionConsumerUrl"`

	// The URL to which the partner sends the logout request.
	// **SCIM++ Properties:**
	//  - caseExact: false
	//  - idcsSearchable: false
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: false
	//  - returned: default
	//  - type: string
	//  - uniqueness: none
	LogoutRequestUrl *string `mandatory:"false" json:"logoutRequestUrl"`

	// The URL to which the partner sends the logout response.
	// **SCIM++ Properties:**
	//  - caseExact: false
	//  - idcsSearchable: false
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: false
	//  - returned: default
	//  - type: string
	//  - uniqueness: none
	LogoutResponseUrl *string `mandatory:"false" json:"logoutResponseUrl"`

	// This can be any string, but there are a set of standard nameIdFormats. If a nameIdFormat other than the standard list is chosen, it will be considered a custom nameidformat. The standard nameidformats include: saml-x509, saml-emailaddress, saml-windowsnamequalifier, saml-kerberos, saml-persistent, saml-transient, saml-unspecified, saml-none, and saml-persistent-opaque.
	// **SCIM++ Properties:**
	//  - caseExact: false
	//  - idcsSearchable: true
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: false
	//  - returned: default
	//  - type: string
	//  - uniqueness: none
	NameIdFormat *string `mandatory:"false" json:"nameIdFormat"`

	// This attribute represents the signing certificate that an App uses to verify the signed authentication request.
	// **SCIM++ Properties:**
	//  - idcsSearchable: false
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: false
	//  - returned: default
	//  - type: string
	//  - uniqueness: none
	SigningCertificate *string `mandatory:"false" json:"signingCertificate"`

	// This attribute represents the encryption certificate that an App uses to encrypt the Security Assertion Markup Language (SAML) assertion.
	// **SCIM++ Properties:**
	//  - idcsSearchable: false
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: false
	//  - returned: default
	//  - type: string
	//  - uniqueness: none
	EncryptionCertificate *string `mandatory:"false" json:"encryptionCertificate"`

	// This attribute indicates the encryption algorithm used to encrypt the SAML assertion.
	// **SCIM++ Properties:**
	//  - caseExact: true
	//  - idcsSearchable: false
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: false
	//  - returned: default
	//  - type: string
	//  - uniqueness: none
	EncryptionAlgorithm AppExtensionSamlServiceProviderAppEncryptionAlgorithmEnum `mandatory:"false" json:"encryptionAlgorithm,omitempty"`

	// This attribute indicates the key encryption algorithm.
	// **Added In:** 18.4.2
	// **SCIM++ Properties:**
	//  - caseExact: true
	//  - idcsSearchable: false
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: false
	//  - returned: default
	//  - type: string
	//  - uniqueness: none
	KeyEncryptionAlgorithm AppExtensionSamlServiceProviderAppKeyEncryptionAlgorithmEnum `mandatory:"false" json:"keyEncryptionAlgorithm,omitempty"`

	// If true, indicates that the system must encrypt the Security Assertion Markup Language (SAML) assertion.
	// **SCIM++ Properties:**
	//  - idcsSearchable: false
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: false
	//  - returned: default
	//  - type: boolean
	//  - uniqueness: none
	EncryptAssertion *bool `mandatory:"false" json:"encryptAssertion"`

	// Indicates which part of the response should be signed.  A value of \"Assertion\" indicates that the Assertion should be signed.  A value of \"Response\" indicates that the SSO Response should be signed. A value of \"AssertionAndResponse\" indicates that both the Assertion and the SSO Response should be signed.
	// **SCIM++ Properties:**
	//  - caseExact: true
	//  - idcsSearchable: false
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: false
	//  - returned: default
	//  - type: string
	//  - uniqueness: none
	SignResponseOrAssertion AppExtensionSamlServiceProviderAppSignResponseOrAssertionEnum `mandatory:"false" json:"signResponseOrAssertion,omitempty"`

	// If true, then the signing certificate is included in the signature.
	// **SCIM++ Properties:**
	//  - idcsSearchable: false
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: false
	//  - returned: default
	//  - type: boolean
	//  - uniqueness: none
	IncludeSigningCertInSignature *bool `mandatory:"false" json:"includeSigningCertInSignature"`

	// This attribute represents the HTTP binding that would be used while logout.
	// **SCIM++ Properties:**
	//  - caseExact: true
	//  - idcsSearchable: false
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: false
	//  - returned: default
	//  - type: string
	//  - uniqueness: none
	LogoutBinding AppExtensionSamlServiceProviderAppLogoutBindingEnum `mandatory:"false" json:"logoutBinding,omitempty"`

	// If true, then the SAML Service supports logout for this App.
	// **SCIM++ Properties:**
	//  - idcsSearchable: false
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: false
	//  - returned: default
	//  - type: boolean
	//  - uniqueness: none
	LogoutEnabled *bool `mandatory:"false" json:"logoutEnabled"`

	// This attribute represents the algorithm used to hash the signature.
	// **SCIM++ Properties:**
	//  - caseExact: true
	//  - idcsSearchable: false
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: false
	//  - returned: default
	//  - type: string
	//  - uniqueness: none
	SignatureHashAlgorithm AppExtensionSamlServiceProviderAppSignatureHashAlgorithmEnum `mandatory:"false" json:"signatureHashAlgorithm,omitempty"`

	// Specifies the preferred federation protocol (SAML2.0 or WS-Fed1.1).
	// **Added In:** 18.2.6
	// **SCIM++ Properties:**
	//  - caseExact: true
	//  - idcsDefaultValue: SAML2.0
	//  - idcsSearchable: true
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: false
	//  - returned: default
	//  - type: string
	//  - uniqueness: none
	FederationProtocol AppExtensionSamlServiceProviderAppFederationProtocolEnum `mandatory:"false" json:"federationProtocol,omitempty"`

	// **Deprecated Since: 18.2.2**
	// **SCIM++ Properties:**
	// - caseExact: false
	// - idcsSearchable: true
	// - idcsValuePersistedInOtherAttribute: true
	// - multiValued: false
	// - mutability: readWrite
	// - required: false
	// - returned: default
	// - type: string
	// - uniqueness: none
	// This property specifies which user attribute is used as the NameID value in the SAML assertion. This attribute can be constructed by using attributes from the Oracle Identity Cloud Service Core Users schema.
	NameIdUserstoreAttribute *string `mandatory:"false" json:"nameIdUserstoreAttribute"`

	// If enabled, then the SAML Service supports Hok for this App.
	// **Added In:** 2101262133
	// **SCIM++ Properties:**
	//  - idcsSearchable: false
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: false
	//  - returned: default
	//  - type: boolean
	//  - uniqueness: none
	HokRequired *bool `mandatory:"false" json:"hokRequired"`

	// Hok Assertion Consumer Service Url
	// **Added In:** 2101262133
	// **SCIM++ Properties:**
	//  - caseExact: false
	//  - idcsSearchable: false
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: false
	//  - returned: default
	//  - type: string
	//  - uniqueness: none
	HokAcsUrl *string `mandatory:"false" json:"hokAcsUrl"`

	// Records the notification timestamp for the SP whose signing certificate is about to expire.
	// **Added In:** 2302092332
	// **SCIM++ Properties:**
	//  - caseExact: false
	//  - idcsSearchable: false
	//  - multiValued: false
	//  - mutability: readOnly
	//  - required: false
	//  - returned: default
	//  - type: dateTime
	//  - uniqueness: none
	LastNotificationSentTime *string `mandatory:"false" json:"lastNotificationSentTime"`

	OutboundAssertionAttributes *AppOutboundAssertionAttributes `mandatory:"false" json:"outboundAssertionAttributes"`

	// Each value of this attribute describes an attribute of User that will be sent in a Security Assertion Markup Language (SAML) assertion.
	// **Deprecated Since: 18.2.2**
	// **SCIM++ Properties:**
	//  - caseExact: false
	//  - idcsCompositeKey: [name]
	//  - idcsSearchable: false
	//  - idcsValuePersistedInOtherAttribute: true
	//  - multiValued: true
	//  - mutability: readWrite
	//  - required: false
	//  - returned: default
	//  - type: complex
	//  - uniqueness: none
	UserAssertionAttributes []AppUserAssertionAttributes `mandatory:"false" json:"userAssertionAttributes"`

	// Each value of this attribute describes an attribute of Group that will be sent in a Security Assertion Markup Language (SAML) assertion.
	// **Deprecated Since: 18.2.2**
	// **SCIM++ Properties:**
	//  - caseExact: false
	//  - idcsCompositeKey: [name]
	//  - idcsSearchable: false
	//  - idcsValuePersistedInOtherAttribute: true
	//  - multiValued: true
	//  - mutability: readWrite
	//  - required: false
	//  - returned: default
	//  - type: complex
	//  - uniqueness: none
	GroupAssertionAttributes []AppGroupAssertionAttributes `mandatory:"false" json:"groupAssertionAttributes"`
}

func (m AppExtensionSamlServiceProviderApp) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m AppExtensionSamlServiceProviderApp) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingAppExtensionSamlServiceProviderAppEncryptionAlgorithmEnum(string(m.EncryptionAlgorithm)); !ok && m.EncryptionAlgorithm != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for EncryptionAlgorithm: %s. Supported values are: %s.", m.EncryptionAlgorithm, strings.Join(GetAppExtensionSamlServiceProviderAppEncryptionAlgorithmEnumStringValues(), ",")))
	}
	if _, ok := GetMappingAppExtensionSamlServiceProviderAppKeyEncryptionAlgorithmEnum(string(m.KeyEncryptionAlgorithm)); !ok && m.KeyEncryptionAlgorithm != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for KeyEncryptionAlgorithm: %s. Supported values are: %s.", m.KeyEncryptionAlgorithm, strings.Join(GetAppExtensionSamlServiceProviderAppKeyEncryptionAlgorithmEnumStringValues(), ",")))
	}
	if _, ok := GetMappingAppExtensionSamlServiceProviderAppSignResponseOrAssertionEnum(string(m.SignResponseOrAssertion)); !ok && m.SignResponseOrAssertion != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SignResponseOrAssertion: %s. Supported values are: %s.", m.SignResponseOrAssertion, strings.Join(GetAppExtensionSamlServiceProviderAppSignResponseOrAssertionEnumStringValues(), ",")))
	}
	if _, ok := GetMappingAppExtensionSamlServiceProviderAppLogoutBindingEnum(string(m.LogoutBinding)); !ok && m.LogoutBinding != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LogoutBinding: %s. Supported values are: %s.", m.LogoutBinding, strings.Join(GetAppExtensionSamlServiceProviderAppLogoutBindingEnumStringValues(), ",")))
	}
	if _, ok := GetMappingAppExtensionSamlServiceProviderAppSignatureHashAlgorithmEnum(string(m.SignatureHashAlgorithm)); !ok && m.SignatureHashAlgorithm != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SignatureHashAlgorithm: %s. Supported values are: %s.", m.SignatureHashAlgorithm, strings.Join(GetAppExtensionSamlServiceProviderAppSignatureHashAlgorithmEnumStringValues(), ",")))
	}
	if _, ok := GetMappingAppExtensionSamlServiceProviderAppFederationProtocolEnum(string(m.FederationProtocol)); !ok && m.FederationProtocol != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for FederationProtocol: %s. Supported values are: %s.", m.FederationProtocol, strings.Join(GetAppExtensionSamlServiceProviderAppFederationProtocolEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// AppExtensionSamlServiceProviderAppEncryptionAlgorithmEnum Enum with underlying type: string
type AppExtensionSamlServiceProviderAppEncryptionAlgorithmEnum string

// Set of constants representing the allowable values for AppExtensionSamlServiceProviderAppEncryptionAlgorithmEnum
const (
	AppExtensionSamlServiceProviderAppEncryptionAlgorithm3des      AppExtensionSamlServiceProviderAppEncryptionAlgorithmEnum = "3DES"
	AppExtensionSamlServiceProviderAppEncryptionAlgorithmAes128    AppExtensionSamlServiceProviderAppEncryptionAlgorithmEnum = "AES-128"
	AppExtensionSamlServiceProviderAppEncryptionAlgorithmAes256    AppExtensionSamlServiceProviderAppEncryptionAlgorithmEnum = "AES-256"
	AppExtensionSamlServiceProviderAppEncryptionAlgorithmAes192    AppExtensionSamlServiceProviderAppEncryptionAlgorithmEnum = "AES-192"
	AppExtensionSamlServiceProviderAppEncryptionAlgorithmAes128Gcm AppExtensionSamlServiceProviderAppEncryptionAlgorithmEnum = "AES-128-GCM"
	AppExtensionSamlServiceProviderAppEncryptionAlgorithmAes256Gcm AppExtensionSamlServiceProviderAppEncryptionAlgorithmEnum = "AES-256-GCM"
	AppExtensionSamlServiceProviderAppEncryptionAlgorithmAes192Gcm AppExtensionSamlServiceProviderAppEncryptionAlgorithmEnum = "AES-192-GCM"
)

var mappingAppExtensionSamlServiceProviderAppEncryptionAlgorithmEnum = map[string]AppExtensionSamlServiceProviderAppEncryptionAlgorithmEnum{
	"3DES":        AppExtensionSamlServiceProviderAppEncryptionAlgorithm3des,
	"AES-128":     AppExtensionSamlServiceProviderAppEncryptionAlgorithmAes128,
	"AES-256":     AppExtensionSamlServiceProviderAppEncryptionAlgorithmAes256,
	"AES-192":     AppExtensionSamlServiceProviderAppEncryptionAlgorithmAes192,
	"AES-128-GCM": AppExtensionSamlServiceProviderAppEncryptionAlgorithmAes128Gcm,
	"AES-256-GCM": AppExtensionSamlServiceProviderAppEncryptionAlgorithmAes256Gcm,
	"AES-192-GCM": AppExtensionSamlServiceProviderAppEncryptionAlgorithmAes192Gcm,
}

var mappingAppExtensionSamlServiceProviderAppEncryptionAlgorithmEnumLowerCase = map[string]AppExtensionSamlServiceProviderAppEncryptionAlgorithmEnum{
	"3des":        AppExtensionSamlServiceProviderAppEncryptionAlgorithm3des,
	"aes-128":     AppExtensionSamlServiceProviderAppEncryptionAlgorithmAes128,
	"aes-256":     AppExtensionSamlServiceProviderAppEncryptionAlgorithmAes256,
	"aes-192":     AppExtensionSamlServiceProviderAppEncryptionAlgorithmAes192,
	"aes-128-gcm": AppExtensionSamlServiceProviderAppEncryptionAlgorithmAes128Gcm,
	"aes-256-gcm": AppExtensionSamlServiceProviderAppEncryptionAlgorithmAes256Gcm,
	"aes-192-gcm": AppExtensionSamlServiceProviderAppEncryptionAlgorithmAes192Gcm,
}

// GetAppExtensionSamlServiceProviderAppEncryptionAlgorithmEnumValues Enumerates the set of values for AppExtensionSamlServiceProviderAppEncryptionAlgorithmEnum
func GetAppExtensionSamlServiceProviderAppEncryptionAlgorithmEnumValues() []AppExtensionSamlServiceProviderAppEncryptionAlgorithmEnum {
	values := make([]AppExtensionSamlServiceProviderAppEncryptionAlgorithmEnum, 0)
	for _, v := range mappingAppExtensionSamlServiceProviderAppEncryptionAlgorithmEnum {
		values = append(values, v)
	}
	return values
}

// GetAppExtensionSamlServiceProviderAppEncryptionAlgorithmEnumStringValues Enumerates the set of values in String for AppExtensionSamlServiceProviderAppEncryptionAlgorithmEnum
func GetAppExtensionSamlServiceProviderAppEncryptionAlgorithmEnumStringValues() []string {
	return []string{
		"3DES",
		"AES-128",
		"AES-256",
		"AES-192",
		"AES-128-GCM",
		"AES-256-GCM",
		"AES-192-GCM",
	}
}

// GetMappingAppExtensionSamlServiceProviderAppEncryptionAlgorithmEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingAppExtensionSamlServiceProviderAppEncryptionAlgorithmEnum(val string) (AppExtensionSamlServiceProviderAppEncryptionAlgorithmEnum, bool) {
	enum, ok := mappingAppExtensionSamlServiceProviderAppEncryptionAlgorithmEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// AppExtensionSamlServiceProviderAppKeyEncryptionAlgorithmEnum Enum with underlying type: string
type AppExtensionSamlServiceProviderAppKeyEncryptionAlgorithmEnum string

// Set of constants representing the allowable values for AppExtensionSamlServiceProviderAppKeyEncryptionAlgorithmEnum
const (
	AppExtensionSamlServiceProviderAppKeyEncryptionAlgorithmV15  AppExtensionSamlServiceProviderAppKeyEncryptionAlgorithmEnum = "RSA-v1.5"
	AppExtensionSamlServiceProviderAppKeyEncryptionAlgorithmOaep AppExtensionSamlServiceProviderAppKeyEncryptionAlgorithmEnum = "RSA-OAEP"
)

var mappingAppExtensionSamlServiceProviderAppKeyEncryptionAlgorithmEnum = map[string]AppExtensionSamlServiceProviderAppKeyEncryptionAlgorithmEnum{
	"RSA-v1.5": AppExtensionSamlServiceProviderAppKeyEncryptionAlgorithmV15,
	"RSA-OAEP": AppExtensionSamlServiceProviderAppKeyEncryptionAlgorithmOaep,
}

var mappingAppExtensionSamlServiceProviderAppKeyEncryptionAlgorithmEnumLowerCase = map[string]AppExtensionSamlServiceProviderAppKeyEncryptionAlgorithmEnum{
	"rsa-v1.5": AppExtensionSamlServiceProviderAppKeyEncryptionAlgorithmV15,
	"rsa-oaep": AppExtensionSamlServiceProviderAppKeyEncryptionAlgorithmOaep,
}

// GetAppExtensionSamlServiceProviderAppKeyEncryptionAlgorithmEnumValues Enumerates the set of values for AppExtensionSamlServiceProviderAppKeyEncryptionAlgorithmEnum
func GetAppExtensionSamlServiceProviderAppKeyEncryptionAlgorithmEnumValues() []AppExtensionSamlServiceProviderAppKeyEncryptionAlgorithmEnum {
	values := make([]AppExtensionSamlServiceProviderAppKeyEncryptionAlgorithmEnum, 0)
	for _, v := range mappingAppExtensionSamlServiceProviderAppKeyEncryptionAlgorithmEnum {
		values = append(values, v)
	}
	return values
}

// GetAppExtensionSamlServiceProviderAppKeyEncryptionAlgorithmEnumStringValues Enumerates the set of values in String for AppExtensionSamlServiceProviderAppKeyEncryptionAlgorithmEnum
func GetAppExtensionSamlServiceProviderAppKeyEncryptionAlgorithmEnumStringValues() []string {
	return []string{
		"RSA-v1.5",
		"RSA-OAEP",
	}
}

// GetMappingAppExtensionSamlServiceProviderAppKeyEncryptionAlgorithmEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingAppExtensionSamlServiceProviderAppKeyEncryptionAlgorithmEnum(val string) (AppExtensionSamlServiceProviderAppKeyEncryptionAlgorithmEnum, bool) {
	enum, ok := mappingAppExtensionSamlServiceProviderAppKeyEncryptionAlgorithmEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// AppExtensionSamlServiceProviderAppSignResponseOrAssertionEnum Enum with underlying type: string
type AppExtensionSamlServiceProviderAppSignResponseOrAssertionEnum string

// Set of constants representing the allowable values for AppExtensionSamlServiceProviderAppSignResponseOrAssertionEnum
const (
	AppExtensionSamlServiceProviderAppSignResponseOrAssertionAssertion            AppExtensionSamlServiceProviderAppSignResponseOrAssertionEnum = "Assertion"
	AppExtensionSamlServiceProviderAppSignResponseOrAssertionResponse             AppExtensionSamlServiceProviderAppSignResponseOrAssertionEnum = "Response"
	AppExtensionSamlServiceProviderAppSignResponseOrAssertionAssertionandresponse AppExtensionSamlServiceProviderAppSignResponseOrAssertionEnum = "AssertionAndResponse"
)

var mappingAppExtensionSamlServiceProviderAppSignResponseOrAssertionEnum = map[string]AppExtensionSamlServiceProviderAppSignResponseOrAssertionEnum{
	"Assertion":            AppExtensionSamlServiceProviderAppSignResponseOrAssertionAssertion,
	"Response":             AppExtensionSamlServiceProviderAppSignResponseOrAssertionResponse,
	"AssertionAndResponse": AppExtensionSamlServiceProviderAppSignResponseOrAssertionAssertionandresponse,
}

var mappingAppExtensionSamlServiceProviderAppSignResponseOrAssertionEnumLowerCase = map[string]AppExtensionSamlServiceProviderAppSignResponseOrAssertionEnum{
	"assertion":            AppExtensionSamlServiceProviderAppSignResponseOrAssertionAssertion,
	"response":             AppExtensionSamlServiceProviderAppSignResponseOrAssertionResponse,
	"assertionandresponse": AppExtensionSamlServiceProviderAppSignResponseOrAssertionAssertionandresponse,
}

// GetAppExtensionSamlServiceProviderAppSignResponseOrAssertionEnumValues Enumerates the set of values for AppExtensionSamlServiceProviderAppSignResponseOrAssertionEnum
func GetAppExtensionSamlServiceProviderAppSignResponseOrAssertionEnumValues() []AppExtensionSamlServiceProviderAppSignResponseOrAssertionEnum {
	values := make([]AppExtensionSamlServiceProviderAppSignResponseOrAssertionEnum, 0)
	for _, v := range mappingAppExtensionSamlServiceProviderAppSignResponseOrAssertionEnum {
		values = append(values, v)
	}
	return values
}

// GetAppExtensionSamlServiceProviderAppSignResponseOrAssertionEnumStringValues Enumerates the set of values in String for AppExtensionSamlServiceProviderAppSignResponseOrAssertionEnum
func GetAppExtensionSamlServiceProviderAppSignResponseOrAssertionEnumStringValues() []string {
	return []string{
		"Assertion",
		"Response",
		"AssertionAndResponse",
	}
}

// GetMappingAppExtensionSamlServiceProviderAppSignResponseOrAssertionEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingAppExtensionSamlServiceProviderAppSignResponseOrAssertionEnum(val string) (AppExtensionSamlServiceProviderAppSignResponseOrAssertionEnum, bool) {
	enum, ok := mappingAppExtensionSamlServiceProviderAppSignResponseOrAssertionEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// AppExtensionSamlServiceProviderAppLogoutBindingEnum Enum with underlying type: string
type AppExtensionSamlServiceProviderAppLogoutBindingEnum string

// Set of constants representing the allowable values for AppExtensionSamlServiceProviderAppLogoutBindingEnum
const (
	AppExtensionSamlServiceProviderAppLogoutBindingRedirect AppExtensionSamlServiceProviderAppLogoutBindingEnum = "Redirect"
	AppExtensionSamlServiceProviderAppLogoutBindingPost     AppExtensionSamlServiceProviderAppLogoutBindingEnum = "Post"
)

var mappingAppExtensionSamlServiceProviderAppLogoutBindingEnum = map[string]AppExtensionSamlServiceProviderAppLogoutBindingEnum{
	"Redirect": AppExtensionSamlServiceProviderAppLogoutBindingRedirect,
	"Post":     AppExtensionSamlServiceProviderAppLogoutBindingPost,
}

var mappingAppExtensionSamlServiceProviderAppLogoutBindingEnumLowerCase = map[string]AppExtensionSamlServiceProviderAppLogoutBindingEnum{
	"redirect": AppExtensionSamlServiceProviderAppLogoutBindingRedirect,
	"post":     AppExtensionSamlServiceProviderAppLogoutBindingPost,
}

// GetAppExtensionSamlServiceProviderAppLogoutBindingEnumValues Enumerates the set of values for AppExtensionSamlServiceProviderAppLogoutBindingEnum
func GetAppExtensionSamlServiceProviderAppLogoutBindingEnumValues() []AppExtensionSamlServiceProviderAppLogoutBindingEnum {
	values := make([]AppExtensionSamlServiceProviderAppLogoutBindingEnum, 0)
	for _, v := range mappingAppExtensionSamlServiceProviderAppLogoutBindingEnum {
		values = append(values, v)
	}
	return values
}

// GetAppExtensionSamlServiceProviderAppLogoutBindingEnumStringValues Enumerates the set of values in String for AppExtensionSamlServiceProviderAppLogoutBindingEnum
func GetAppExtensionSamlServiceProviderAppLogoutBindingEnumStringValues() []string {
	return []string{
		"Redirect",
		"Post",
	}
}

// GetMappingAppExtensionSamlServiceProviderAppLogoutBindingEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingAppExtensionSamlServiceProviderAppLogoutBindingEnum(val string) (AppExtensionSamlServiceProviderAppLogoutBindingEnum, bool) {
	enum, ok := mappingAppExtensionSamlServiceProviderAppLogoutBindingEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// AppExtensionSamlServiceProviderAppSignatureHashAlgorithmEnum Enum with underlying type: string
type AppExtensionSamlServiceProviderAppSignatureHashAlgorithmEnum string

// Set of constants representing the allowable values for AppExtensionSamlServiceProviderAppSignatureHashAlgorithmEnum
const (
	AppExtensionSamlServiceProviderAppSignatureHashAlgorithm1   AppExtensionSamlServiceProviderAppSignatureHashAlgorithmEnum = "SHA-1"
	AppExtensionSamlServiceProviderAppSignatureHashAlgorithm256 AppExtensionSamlServiceProviderAppSignatureHashAlgorithmEnum = "SHA-256"
)

var mappingAppExtensionSamlServiceProviderAppSignatureHashAlgorithmEnum = map[string]AppExtensionSamlServiceProviderAppSignatureHashAlgorithmEnum{
	"SHA-1":   AppExtensionSamlServiceProviderAppSignatureHashAlgorithm1,
	"SHA-256": AppExtensionSamlServiceProviderAppSignatureHashAlgorithm256,
}

var mappingAppExtensionSamlServiceProviderAppSignatureHashAlgorithmEnumLowerCase = map[string]AppExtensionSamlServiceProviderAppSignatureHashAlgorithmEnum{
	"sha-1":   AppExtensionSamlServiceProviderAppSignatureHashAlgorithm1,
	"sha-256": AppExtensionSamlServiceProviderAppSignatureHashAlgorithm256,
}

// GetAppExtensionSamlServiceProviderAppSignatureHashAlgorithmEnumValues Enumerates the set of values for AppExtensionSamlServiceProviderAppSignatureHashAlgorithmEnum
func GetAppExtensionSamlServiceProviderAppSignatureHashAlgorithmEnumValues() []AppExtensionSamlServiceProviderAppSignatureHashAlgorithmEnum {
	values := make([]AppExtensionSamlServiceProviderAppSignatureHashAlgorithmEnum, 0)
	for _, v := range mappingAppExtensionSamlServiceProviderAppSignatureHashAlgorithmEnum {
		values = append(values, v)
	}
	return values
}

// GetAppExtensionSamlServiceProviderAppSignatureHashAlgorithmEnumStringValues Enumerates the set of values in String for AppExtensionSamlServiceProviderAppSignatureHashAlgorithmEnum
func GetAppExtensionSamlServiceProviderAppSignatureHashAlgorithmEnumStringValues() []string {
	return []string{
		"SHA-1",
		"SHA-256",
	}
}

// GetMappingAppExtensionSamlServiceProviderAppSignatureHashAlgorithmEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingAppExtensionSamlServiceProviderAppSignatureHashAlgorithmEnum(val string) (AppExtensionSamlServiceProviderAppSignatureHashAlgorithmEnum, bool) {
	enum, ok := mappingAppExtensionSamlServiceProviderAppSignatureHashAlgorithmEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// AppExtensionSamlServiceProviderAppFederationProtocolEnum Enum with underlying type: string
type AppExtensionSamlServiceProviderAppFederationProtocolEnum string

// Set of constants representing the allowable values for AppExtensionSamlServiceProviderAppFederationProtocolEnum
const (
	AppExtensionSamlServiceProviderAppFederationProtocolSaml20  AppExtensionSamlServiceProviderAppFederationProtocolEnum = "SAML2.0"
	AppExtensionSamlServiceProviderAppFederationProtocolWsFed11 AppExtensionSamlServiceProviderAppFederationProtocolEnum = "WS-Fed1.1"
)

var mappingAppExtensionSamlServiceProviderAppFederationProtocolEnum = map[string]AppExtensionSamlServiceProviderAppFederationProtocolEnum{
	"SAML2.0":   AppExtensionSamlServiceProviderAppFederationProtocolSaml20,
	"WS-Fed1.1": AppExtensionSamlServiceProviderAppFederationProtocolWsFed11,
}

var mappingAppExtensionSamlServiceProviderAppFederationProtocolEnumLowerCase = map[string]AppExtensionSamlServiceProviderAppFederationProtocolEnum{
	"saml2.0":   AppExtensionSamlServiceProviderAppFederationProtocolSaml20,
	"ws-fed1.1": AppExtensionSamlServiceProviderAppFederationProtocolWsFed11,
}

// GetAppExtensionSamlServiceProviderAppFederationProtocolEnumValues Enumerates the set of values for AppExtensionSamlServiceProviderAppFederationProtocolEnum
func GetAppExtensionSamlServiceProviderAppFederationProtocolEnumValues() []AppExtensionSamlServiceProviderAppFederationProtocolEnum {
	values := make([]AppExtensionSamlServiceProviderAppFederationProtocolEnum, 0)
	for _, v := range mappingAppExtensionSamlServiceProviderAppFederationProtocolEnum {
		values = append(values, v)
	}
	return values
}

// GetAppExtensionSamlServiceProviderAppFederationProtocolEnumStringValues Enumerates the set of values in String for AppExtensionSamlServiceProviderAppFederationProtocolEnum
func GetAppExtensionSamlServiceProviderAppFederationProtocolEnumStringValues() []string {
	return []string{
		"SAML2.0",
		"WS-Fed1.1",
	}
}

// GetMappingAppExtensionSamlServiceProviderAppFederationProtocolEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingAppExtensionSamlServiceProviderAppFederationProtocolEnum(val string) (AppExtensionSamlServiceProviderAppFederationProtocolEnum, bool) {
	enum, ok := mappingAppExtensionSamlServiceProviderAppFederationProtocolEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
