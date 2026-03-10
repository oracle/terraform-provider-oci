// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
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

// IdentityProofingProvider Manage Identity Proofing Providers. Oracle Identity Cloud Service provides cross-domain SSO capabilities via the OASIS SAML 2.0 SSO protocol and implements two modes of operation for Federation SSO--as an IdP where the user is authenticated on behalf of remote Service Providers (SP), and as an SP where Oracle Identity Cloud Service delegates authentication to a remote IdP. As an IdP, Oracle Identity Cloud Service can integrate with multiple SPs at the same time. As an SP, This REST API is SCIM compliant.
type IdentityProofingProvider struct {

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

	// The Identity Proofing Provider.
	// **SCIM++ Properties:**
	//  - idcsSearchable: true
	//  - multiValued: false
	//  - mutability: readWrite
	//  - idcsCanonicalValueSourceFilter: provider eq "$(provider)"
	//  - idcsCanonicalValueSourceResourceType: IdentityProofingProviderTemplate
	//  - caseExact: true
	//  - required: true
	//  - returned: always
	//  - type: string
	//  - uniqueness: none
	Provider *string `mandatory:"true" json:"provider"`

	// Name of the Identity Proofing Provider.
	// **SCIM++ Properties:**
	//  - idcsSearchable: false
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: true
	//  - returned: always
	//  - type: string
	//  - uniqueness: global
	Name *string `mandatory:"true" json:"name"`

	// Configure the verification provider.
	// **SCIM++ Properties:**
	//  - idcsSearchable: false
	//  - multiValued: true
	//  - mutability: immutable
	//  - required: true
	//  - returned: always
	//  - idcsCompositeKey: [name]
	//  - type: complex
	//  - uniqueness: server
	Configuration []IdentityProofingProviderConfiguration `mandatory:"true" json:"configuration"`

	// Configure the verification claims and IAM Domain user attribute mapping.
	// **SCIM++ Properties:**
	//  - idcsSearchable: false
	//  - multiValued: true
	//  - mutability: immutable
	//  - required: true
	//  - returned: always
	//  - idcsCompositeKey: [verifiableClaim]
	//  - type: complex
	//  - uniqueness: server
	ClaimMapping []IdentityProofingProviderClaimMapping `mandatory:"true" json:"claimMapping"`

	// Unique identifier for the SCIM Resource as defined by the Service Provider. Each representation of the Resource MUST include a non-empty id value. This identifier MUST be unique across the Service Provider's entire set of Resources. It MUST be a stable, non-reassignable identifier that does not change when the same Resource is returned in subsequent requests. The value of the id attribute is always issued by the Service Provider and MUST never be specified by the Service Consumer. bulkId: is a reserved keyword and MUST NOT be used in the unique identifier.
	// **SCIM++ Properties:**
	//  - caseExact: true
	//  - idcsSearchable: true
	//  - multiValued: false
	//  - mutability: readOnly
	//  - required: false
	//  - returned: always
	//  - type: string
	//  - uniqueness: global
	Id *string `mandatory:"false" json:"id"`

	// Unique OCI identifier (OCID) for the SCIM Resource.
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
	//  - idcsCsvAttributeNameMappings: [[columnHeaderName:Tag Key, mapsTo:tags.key], [columnHeaderName:Tag Value, mapsTo:tags.value]]
	//  - idcsSearchable: true
	//  - multiValued: true
	//  - mutability: readWrite
	//  - required: false
	//  - returned: request
	//  - type: complex
	//  - uniqueness: none
	Tags []Tags `mandatory:"false" json:"tags"`

	// Operations that are locked on the resource.
	// **SCIM++ Properties:**
	//  - idcsSearchable: false
	//  - idcsCsvAttributeNameMappings: [[columnHeaderName:Locked Operations, multiValueDelimiter:;]]
	//  - multiValued: true
	//  - mutability: readWrite
	//  - required: false
	//  - returned: default
	//  - type: string
	//  - uniqueness: none
	IdcsLockedOperations []IdentityProofingProviderIdcsLockedOperationsEnum `mandatory:"false" json:"idcsLockedOperations,omitempty"`

	IdcsLockedBy *IdentityProofingProviderIdcsLockedBy `mandatory:"false" json:"idcsLockedBy"`

	// The most recent DateTime the resource was locked.
	// **SCIM++ Properties:**
	//  - caseExact: false
	//  - idcsSearchable: false
	//  - multiValued: false
	//  - mutability: readOnly
	//  - required: false
	//  - returned: default
	//  - type: dateTime
	//  - uniqueness: none
	IdcsLockedOn *string `mandatory:"false" json:"idcsLockedOn"`

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

	// OCI Domain Id (OCID) in which the resource lives.
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

	// OCI Compartment Id (OCID) in which the resource lives.
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

	// OCI Tenant Id (OCID) in which the resource lives.
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

	// Description field to add comments and additional information about the Identity Proofing Provider.
	// **SCIM++ Properties:**
	//  - idcsSearchable: false
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: false
	//  - returned: default
	//  - type: string
	//  - uniqueness: none
	Description *string `mandatory:"false" json:"description"`

	// Status of the Identity Proofing Provider.
	// **SCIM++ Properties:**
	//  - idcsSearchable: true
	//  - multiValued: false
	//  - caseExact: true
	//  - mutability: readWrite
	//  - required: false
	//  - returned: default
	//  - type: string
	//  - uniqueness: none
	Status IdentityProofingProviderStatusEnum `mandatory:"false" json:"status,omitempty"`

	// Identity Proofing runtime custom data.
	// **Added In:** 2505161800
	// **SCIM++ Properties:**
	//  - idcsCompositeKey: [attrName]
	//  - multiValued: true
	//  - mutability: readWrite
	//  - required: false
	//  - returned: default
	//  - type: complex
	RuntimeData []IdentityProofingProviderRuntimeData `mandatory:"false" json:"runtimeData"`
}

func (m IdentityProofingProvider) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m IdentityProofingProvider) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	for _, val := range m.IdcsPreventedOperations {
		if _, ok := GetMappingIdcsPreventedOperationsEnum(string(val)); !ok && val != "" {
			errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for IdcsPreventedOperations: %s. Supported values are: %s.", val, strings.Join(GetIdcsPreventedOperationsEnumStringValues(), ",")))
		}
	}

	for _, val := range m.IdcsLockedOperations {
		if _, ok := GetMappingIdentityProofingProviderIdcsLockedOperationsEnum(string(val)); !ok && val != "" {
			errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for IdcsLockedOperations: %s. Supported values are: %s.", val, strings.Join(GetIdentityProofingProviderIdcsLockedOperationsEnumStringValues(), ",")))
		}
	}

	if _, ok := GetMappingIdentityProofingProviderStatusEnum(string(m.Status)); !ok && m.Status != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Status: %s. Supported values are: %s.", m.Status, strings.Join(GetIdentityProofingProviderStatusEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// IdentityProofingProviderIdcsLockedOperationsEnum Enum with underlying type: string
type IdentityProofingProviderIdcsLockedOperationsEnum string

// Set of constants representing the allowable values for IdentityProofingProviderIdcsLockedOperationsEnum
const (
	IdentityProofingProviderIdcsLockedOperationsReplace IdentityProofingProviderIdcsLockedOperationsEnum = "replace"
	IdentityProofingProviderIdcsLockedOperationsUpdate  IdentityProofingProviderIdcsLockedOperationsEnum = "update"
	IdentityProofingProviderIdcsLockedOperationsDelete  IdentityProofingProviderIdcsLockedOperationsEnum = "delete"
)

var mappingIdentityProofingProviderIdcsLockedOperationsEnum = map[string]IdentityProofingProviderIdcsLockedOperationsEnum{
	"replace": IdentityProofingProviderIdcsLockedOperationsReplace,
	"update":  IdentityProofingProviderIdcsLockedOperationsUpdate,
	"delete":  IdentityProofingProviderIdcsLockedOperationsDelete,
}

var mappingIdentityProofingProviderIdcsLockedOperationsEnumLowerCase = map[string]IdentityProofingProviderIdcsLockedOperationsEnum{
	"replace": IdentityProofingProviderIdcsLockedOperationsReplace,
	"update":  IdentityProofingProviderIdcsLockedOperationsUpdate,
	"delete":  IdentityProofingProviderIdcsLockedOperationsDelete,
}

// GetIdentityProofingProviderIdcsLockedOperationsEnumValues Enumerates the set of values for IdentityProofingProviderIdcsLockedOperationsEnum
func GetIdentityProofingProviderIdcsLockedOperationsEnumValues() []IdentityProofingProviderIdcsLockedOperationsEnum {
	values := make([]IdentityProofingProviderIdcsLockedOperationsEnum, 0)
	for _, v := range mappingIdentityProofingProviderIdcsLockedOperationsEnum {
		values = append(values, v)
	}
	return values
}

// GetIdentityProofingProviderIdcsLockedOperationsEnumStringValues Enumerates the set of values in String for IdentityProofingProviderIdcsLockedOperationsEnum
func GetIdentityProofingProviderIdcsLockedOperationsEnumStringValues() []string {
	return []string{
		"replace",
		"update",
		"delete",
	}
}

// GetMappingIdentityProofingProviderIdcsLockedOperationsEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingIdentityProofingProviderIdcsLockedOperationsEnum(val string) (IdentityProofingProviderIdcsLockedOperationsEnum, bool) {
	enum, ok := mappingIdentityProofingProviderIdcsLockedOperationsEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// IdentityProofingProviderStatusEnum Enum with underlying type: string
type IdentityProofingProviderStatusEnum string

// Set of constants representing the allowable values for IdentityProofingProviderStatusEnum
const (
	IdentityProofingProviderStatusInactive IdentityProofingProviderStatusEnum = "INACTIVE"
	IdentityProofingProviderStatusActive   IdentityProofingProviderStatusEnum = "ACTIVE"
)

var mappingIdentityProofingProviderStatusEnum = map[string]IdentityProofingProviderStatusEnum{
	"INACTIVE": IdentityProofingProviderStatusInactive,
	"ACTIVE":   IdentityProofingProviderStatusActive,
}

var mappingIdentityProofingProviderStatusEnumLowerCase = map[string]IdentityProofingProviderStatusEnum{
	"inactive": IdentityProofingProviderStatusInactive,
	"active":   IdentityProofingProviderStatusActive,
}

// GetIdentityProofingProviderStatusEnumValues Enumerates the set of values for IdentityProofingProviderStatusEnum
func GetIdentityProofingProviderStatusEnumValues() []IdentityProofingProviderStatusEnum {
	values := make([]IdentityProofingProviderStatusEnum, 0)
	for _, v := range mappingIdentityProofingProviderStatusEnum {
		values = append(values, v)
	}
	return values
}

// GetIdentityProofingProviderStatusEnumStringValues Enumerates the set of values in String for IdentityProofingProviderStatusEnum
func GetIdentityProofingProviderStatusEnumStringValues() []string {
	return []string{
		"INACTIVE",
		"ACTIVE",
	}
}

// GetMappingIdentityProofingProviderStatusEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingIdentityProofingProviderStatusEnum(val string) (IdentityProofingProviderStatusEnum, bool) {
	enum, ok := mappingIdentityProofingProviderStatusEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
