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

// IdentityPropagationTrust Schema used for Identity Propagation Trust.
type IdentityPropagationTrust struct {

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

	// The name of the the Identity Propagation Trust.
	// **SCIM++ Properties:**
	//  - type: string
	//  - caseExact: false
	//  - idcsSearchable: true
	//  - multiValued: false
	//  - required: true
	//  - mutability: immutable
	//  - returned: default
	//  - uniqueness: none
	Name *string `mandatory:"true" json:"name"`

	// The type of the inbound token from the Identity cloud provider.
	// **SCIM++ Properties:**
	//  - caseExact: true
	//  - idcsSearchable: false
	//  - required: true
	//  - mutability: readWrite
	//  - returned: default
	//  - type: string
	//  - multiValued: false
	//  - uniqueness: none
	Type IdentityPropagationTrustTypeEnum `mandatory:"true" json:"type"`

	// The issuer claim of the Identity provider.
	// **SCIM++ Properties:**
	//  - type: string
	//  - multiValued: false
	//  - required: true
	//  - mutability: readWrite
	//  - returned: always
	//  - caseExact: true
	//  - idcsSearchable: true
	//  - uniqueness: server
	Issuer *string `mandatory:"true" json:"issuer"`

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

	// The description of the Identity Propagation Trust.
	// **SCIM++ Properties:**
	//  - type: string
	//  - multiValued: false
	//  - required: false
	//  - mutability: readWrite
	//  - returned: default
	//  - uniqueness: none
	//  - caseExact: false
	//  - idcsSearchable: false
	Description *string `mandatory:"false" json:"description"`

	// The Identity cloud provider service identifier, for example, the Azure Tenancy ID, AWS Account ID, or GCP Project ID.
	// **SCIM++ Properties:**
	//  - type: string
	//  - multiValued: false
	//  - required: false
	//  - mutability: readWrite
	//  - returned: default
	//  - caseExact: true
	//  - idcsSearchable: true
	//  - uniqueness: none
	AccountId *string `mandatory:"false" json:"accountId"`

	// Used for locating the subject claim from the incoming token.
	// **SCIM++ Properties:**
	//  - type: string
	//  - multiValued: false
	//  - required: false
	//  - mutability: readWrite
	//  - returned: default
	//  - uniqueness: none
	//  - caseExact: true
	//  - idcsSearchable: false
	SubjectClaimName *string `mandatory:"false" json:"subjectClaimName"`

	// Subject Mapping Attribute to which the value from subject claim name value would be used for identity lookup.
	// **SCIM++ Properties:**
	//  - type: string
	//  - multiValued: false
	//  - idcsSearchable: false
	//  - required: false
	//  - mutability: readWrite
	//  - returned: default
	//  - uniqueness: none
	SubjectMappingAttribute *string `mandatory:"false" json:"subjectMappingAttribute"`

	// The type of the resource against which lookup will be made in the identity domain in IAM for the incoming subject claim value.
	// **SCIM++ Properties:**
	//  - idcsSearchable: false
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: false
	//  - returned: default
	//  - type: string
	//  - uniqueness: none
	SubjectType IdentityPropagationTrustSubjectTypeEnum `mandatory:"false" json:"subjectType,omitempty"`

	// The claim name that identifies to whom the JWT/SAML token is issued. If AWS, then \"aud\" or \"client_id\". If Azure, then \"appid\". If GCP, then \"aud\".
	// **SCIM++ Properties:**
	//  - type: string
	//  - multiValued: false
	//  - required: false
	//  - mutability: readWrite
	//  - returned: default
	//  - uniqueness: none
	//  - idcsSearchable: false
	ClientClaimName *string `mandatory:"false" json:"clientClaimName"`

	// The value that corresponds to the client claim name used to identify to whom the token is issued.
	// **SCIM++ Properties:**
	//  - type: string
	//  - multiValued: true
	//  - required: false
	//  - mutability: readWrite
	//  - returned: default
	//  - uniqueness: none
	//  - caseExact: true
	//  - idcsSearchable: false
	ClientClaimValues []string `mandatory:"false" json:"clientClaimValues"`

	// If true, specifies that this Identity Propagation Trust is in an enabled state. The default value is false.
	// **SCIM++ Properties:**
	//  - type: boolean
	//  - multiValued: false
	//  - required: false
	//  - mutability: readWrite
	//  - returned: default
	//  - uniqueness: none
	//  - idcsSearchable: true
	Active *bool `mandatory:"false" json:"active"`

	// The cloud provider's public key API of SAML and OIDC providers for signature validation.
	// **SCIM++ Properties:**
	//  - type: string
	//  - multiValued: false
	//  - required: false
	//  - mutability: readWrite
	//  - returned: default
	//  - uniqueness: none
	//  - caseExact: false
	//  - idcsSearchable: false
	PublicKeyEndpoint *string `mandatory:"false" json:"publicKeyEndpoint"`

	// Store the public key if public key cert.
	// **SCIM++ Properties:**
	//  - type: string
	//  - multiValued: false
	//  - required: false
	//  - mutability: readWrite
	//  - returned: default
	//  - uniqueness: none
	//  - idcsSearchable: false
	PublicCertificate *string `mandatory:"false" json:"publicCertificate"`

	// The value of all the authorized OAuth Clients.
	// **SCIM++ Properties:**
	//  - idcsSearchable: false
	//  - multiValued: true
	//  - mutability: readWrite
	//  - required: false
	//  - returned: default
	//  - type: string
	//  - uniqueness: none
	OauthClients []string `mandatory:"false" json:"oauthClients"`

	// Allow customers to define whether the resulting token should contain the authenticated user as the subject or whether the token should impersonate another Application Principal in IAM.
	// **SCIM++ Properties:**
	//  - type: boolean
	//  - multiValued: false
	//  - required: false
	//  - mutability: readWrite
	//  - returned: default
	//  - uniqueness: none
	//  - idcsSearchable: false
	AllowImpersonation *bool `mandatory:"false" json:"allowImpersonation"`

	// The clock skew (in secs) that's allowed for the token issue and expiry time.
	// **Added In:** 2308181911
	// **SCIM++ Properties:**
	//  - caseExact: false
	//  - idcsSearchable: false
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: false
	//  - returned: default
	//  - type: integer
	//  - uniqueness: none
	ClockSkewSeconds *int `mandatory:"false" json:"clockSkewSeconds"`

	// The Impersonating Principal.
	// **SCIM++ Properties:**
	//  - idcsCompositeKey: [rule, value]
	//  - idcsSearchable: false
	//  - multiValued: true
	//  - mutability: readWrite
	//  - required: false
	//  - returned: request
	//  - type: complex
	//  - uniqueness: none
	ImpersonationServiceUsers []IdentityPropagationTrustImpersonationServiceUsers `mandatory:"false" json:"impersonationServiceUsers"`

	Keytab *IdentityPropagationTrustKeytab `mandatory:"false" json:"keytab"`
}

func (m IdentityPropagationTrust) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m IdentityPropagationTrust) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingIdentityPropagationTrustTypeEnum(string(m.Type)); !ok && m.Type != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Type: %s. Supported values are: %s.", m.Type, strings.Join(GetIdentityPropagationTrustTypeEnumStringValues(), ",")))
	}

	for _, val := range m.IdcsPreventedOperations {
		if _, ok := GetMappingIdcsPreventedOperationsEnum(string(val)); !ok && val != "" {
			errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for IdcsPreventedOperations: %s. Supported values are: %s.", val, strings.Join(GetIdcsPreventedOperationsEnumStringValues(), ",")))
		}
	}

	if _, ok := GetMappingIdentityPropagationTrustSubjectTypeEnum(string(m.SubjectType)); !ok && m.SubjectType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SubjectType: %s. Supported values are: %s.", m.SubjectType, strings.Join(GetIdentityPropagationTrustSubjectTypeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// IdentityPropagationTrustTypeEnum Enum with underlying type: string
type IdentityPropagationTrustTypeEnum string

// Set of constants representing the allowable values for IdentityPropagationTrustTypeEnum
const (
	IdentityPropagationTrustTypeJwt    IdentityPropagationTrustTypeEnum = "JWT"
	IdentityPropagationTrustTypeSaml   IdentityPropagationTrustTypeEnum = "SAML"
	IdentityPropagationTrustTypeSpnego IdentityPropagationTrustTypeEnum = "SPNEGO"
	IdentityPropagationTrustTypeAws    IdentityPropagationTrustTypeEnum = "AWS"
)

var mappingIdentityPropagationTrustTypeEnum = map[string]IdentityPropagationTrustTypeEnum{
	"JWT":    IdentityPropagationTrustTypeJwt,
	"SAML":   IdentityPropagationTrustTypeSaml,
	"SPNEGO": IdentityPropagationTrustTypeSpnego,
	"AWS":    IdentityPropagationTrustTypeAws,
}

var mappingIdentityPropagationTrustTypeEnumLowerCase = map[string]IdentityPropagationTrustTypeEnum{
	"jwt":    IdentityPropagationTrustTypeJwt,
	"saml":   IdentityPropagationTrustTypeSaml,
	"spnego": IdentityPropagationTrustTypeSpnego,
	"aws":    IdentityPropagationTrustTypeAws,
}

// GetIdentityPropagationTrustTypeEnumValues Enumerates the set of values for IdentityPropagationTrustTypeEnum
func GetIdentityPropagationTrustTypeEnumValues() []IdentityPropagationTrustTypeEnum {
	values := make([]IdentityPropagationTrustTypeEnum, 0)
	for _, v := range mappingIdentityPropagationTrustTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetIdentityPropagationTrustTypeEnumStringValues Enumerates the set of values in String for IdentityPropagationTrustTypeEnum
func GetIdentityPropagationTrustTypeEnumStringValues() []string {
	return []string{
		"JWT",
		"SAML",
		"SPNEGO",
		"AWS",
	}
}

// GetMappingIdentityPropagationTrustTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingIdentityPropagationTrustTypeEnum(val string) (IdentityPropagationTrustTypeEnum, bool) {
	enum, ok := mappingIdentityPropagationTrustTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// IdentityPropagationTrustSubjectTypeEnum Enum with underlying type: string
type IdentityPropagationTrustSubjectTypeEnum string

// Set of constants representing the allowable values for IdentityPropagationTrustSubjectTypeEnum
const (
	IdentityPropagationTrustSubjectTypeUser IdentityPropagationTrustSubjectTypeEnum = "User"
	IdentityPropagationTrustSubjectTypeApp  IdentityPropagationTrustSubjectTypeEnum = "App"
)

var mappingIdentityPropagationTrustSubjectTypeEnum = map[string]IdentityPropagationTrustSubjectTypeEnum{
	"User": IdentityPropagationTrustSubjectTypeUser,
	"App":  IdentityPropagationTrustSubjectTypeApp,
}

var mappingIdentityPropagationTrustSubjectTypeEnumLowerCase = map[string]IdentityPropagationTrustSubjectTypeEnum{
	"user": IdentityPropagationTrustSubjectTypeUser,
	"app":  IdentityPropagationTrustSubjectTypeApp,
}

// GetIdentityPropagationTrustSubjectTypeEnumValues Enumerates the set of values for IdentityPropagationTrustSubjectTypeEnum
func GetIdentityPropagationTrustSubjectTypeEnumValues() []IdentityPropagationTrustSubjectTypeEnum {
	values := make([]IdentityPropagationTrustSubjectTypeEnum, 0)
	for _, v := range mappingIdentityPropagationTrustSubjectTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetIdentityPropagationTrustSubjectTypeEnumStringValues Enumerates the set of values in String for IdentityPropagationTrustSubjectTypeEnum
func GetIdentityPropagationTrustSubjectTypeEnumStringValues() []string {
	return []string{
		"User",
		"App",
	}
}

// GetMappingIdentityPropagationTrustSubjectTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingIdentityPropagationTrustSubjectTypeEnum(val string) (IdentityPropagationTrustSubjectTypeEnum, bool) {
	enum, ok := mappingIdentityPropagationTrustSubjectTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
