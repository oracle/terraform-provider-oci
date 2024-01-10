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

// OAuth2ClientCredential The user's OAuth2 client credentials.
type OAuth2ClientCredential struct {

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

	// Name
	// **SCIM++ Properties:**
	//  - caseExact: false
	//  - type: string
	//  - mutability: readWrite
	//  - required: true
	//  - returned: default
	Name *string `mandatory:"true" json:"name"`

	// Scopes
	// **SCIM++ Properties:**
	//  - caseExact: false
	//  - idcsCompositeKey: [audience, scope]
	//  - type: complex
	//  - mutability: readWrite
	//  - multiValued: true
	//  - required: true
	//  - returned: default
	Scopes []OAuth2ClientCredentialScopes `mandatory:"true" json:"scopes"`

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

	// Description
	// **Added In:** 2101262133
	// **SCIM++ Properties:**
	//  - caseExact: false
	//  - type: string
	//  - mutability: readWrite
	//  - required: false
	//  - returned: default
	Description *string `mandatory:"false" json:"description"`

	// The user's credential status.
	// **Added In:** 2109090424
	// **SCIM++ Properties:**
	//  - caseExact: false
	//  - idcsSearchable: false
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: false
	//  - returned: never
	//  - type: string
	//  - uniqueness: none
	Status OAuth2ClientCredentialStatusEnum `mandatory:"false" json:"status,omitempty"`

	// When the user's credentials expire.
	// **Added In:** 2109090424
	// **SCIM++ Properties:**
	//  - caseExact: false
	//  - idcsSearchable: false
	//  - multiValued: false
	//  - mutability: immutable
	//  - required: false
	//  - returned: default
	//  - type: dateTime
	//  - uniqueness: none
	ExpiresOn *string `mandatory:"false" json:"expiresOn"`

	// Specifies whether the secret must be reset.
	// **Added In:** 2109090424
	// **SCIM++ Properties:**
	//  - caseExact: false
	//  - idcsSearchable: false
	//  - multiValued: false
	//  - mutability: writeOnly
	//  - required: false
	//  - returned: default
	//  - type: boolean
	//  - uniqueness: none
	IsResetSecret *bool `mandatory:"false" json:"isResetSecret"`

	User *OAuth2ClientCredentialUser `mandatory:"false" json:"user"`

	UrnIetfParamsScimSchemasOracleIdcsExtensionSelfChangeUser *ExtensionSelfChangeUser `mandatory:"false" json:"urn:ietf:params:scim:schemas:oracle:idcs:extension:selfChange:User"`
}

func (m OAuth2ClientCredential) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m OAuth2ClientCredential) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	for _, val := range m.IdcsPreventedOperations {
		if _, ok := GetMappingIdcsPreventedOperationsEnum(string(val)); !ok && val != "" {
			errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for IdcsPreventedOperations: %s. Supported values are: %s.", val, strings.Join(GetIdcsPreventedOperationsEnumStringValues(), ",")))
		}
	}

	if _, ok := GetMappingOAuth2ClientCredentialStatusEnum(string(m.Status)); !ok && m.Status != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Status: %s. Supported values are: %s.", m.Status, strings.Join(GetOAuth2ClientCredentialStatusEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// OAuth2ClientCredentialStatusEnum Enum with underlying type: string
type OAuth2ClientCredentialStatusEnum string

// Set of constants representing the allowable values for OAuth2ClientCredentialStatusEnum
const (
	OAuth2ClientCredentialStatusActive   OAuth2ClientCredentialStatusEnum = "ACTIVE"
	OAuth2ClientCredentialStatusInactive OAuth2ClientCredentialStatusEnum = "INACTIVE"
)

var mappingOAuth2ClientCredentialStatusEnum = map[string]OAuth2ClientCredentialStatusEnum{
	"ACTIVE":   OAuth2ClientCredentialStatusActive,
	"INACTIVE": OAuth2ClientCredentialStatusInactive,
}

var mappingOAuth2ClientCredentialStatusEnumLowerCase = map[string]OAuth2ClientCredentialStatusEnum{
	"active":   OAuth2ClientCredentialStatusActive,
	"inactive": OAuth2ClientCredentialStatusInactive,
}

// GetOAuth2ClientCredentialStatusEnumValues Enumerates the set of values for OAuth2ClientCredentialStatusEnum
func GetOAuth2ClientCredentialStatusEnumValues() []OAuth2ClientCredentialStatusEnum {
	values := make([]OAuth2ClientCredentialStatusEnum, 0)
	for _, v := range mappingOAuth2ClientCredentialStatusEnum {
		values = append(values, v)
	}
	return values
}

// GetOAuth2ClientCredentialStatusEnumStringValues Enumerates the set of values in String for OAuth2ClientCredentialStatusEnum
func GetOAuth2ClientCredentialStatusEnumStringValues() []string {
	return []string{
		"ACTIVE",
		"INACTIVE",
	}
}

// GetMappingOAuth2ClientCredentialStatusEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingOAuth2ClientCredentialStatusEnum(val string) (OAuth2ClientCredentialStatusEnum, bool) {
	enum, ok := mappingOAuth2ClientCredentialStatusEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
