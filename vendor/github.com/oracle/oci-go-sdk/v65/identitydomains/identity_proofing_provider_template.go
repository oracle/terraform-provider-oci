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

// IdentityProofingProviderTemplate Maintains a registered Collection of Identity Proofing vendors list. This REST API is SCIM compliant.
type IdentityProofingProviderTemplate struct {

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

	// Name of the Identity Proofing Provider.
	// **SCIM++ Properties:**
	//  - idcsSearchable: true
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: true
	//  - returned: always
	//  - type: string
	//  - uniqueness: global
	Provider *string `mandatory:"true" json:"provider"`

	// verification URL.
	// **SCIM++ Properties:**
	//  - idcsSearchable: false
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: true
	//  - returned: default
	//  - type: string
	//  - uniqueness: none
	VerificationUrl *string `mandatory:"true" json:"verificationUrl"`

	// Service type supported by the Identity Proofing Provider.
	// **Added In:** 2505161800
	// **SCIM++ Properties:**
	//  - idcsSearchable: false
	//  - multiValued: true
	//  - mutability: readWrite
	//  - required: true
	//  - returned: default
	//  - type: string
	//  - uniqueness: none
	ServiceType []string `mandatory:"true" json:"serviceType"`

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
	IdcsLockedOperations []IdentityProofingProviderTemplateIdcsLockedOperationsEnum `mandatory:"false" json:"idcsLockedOperations,omitempty"`

	IdcsLockedBy *IdentityProofingProviderTemplateIdcsLockedBy `mandatory:"false" json:"idcsLockedBy"`

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

	// Proofing Provider Configuration Template.
	// **SCIM++ Properties:**
	//  - idcsCompositeKey: [name]
	//  - multiValued: true
	//  - mutability: readWrite
	//  - required: false
	//  - returned: default
	//  - type: complex
	Configuration []IdentityProofingProviderTemplateConfiguration `mandatory:"false" json:"configuration"`
}

func (m IdentityProofingProviderTemplate) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m IdentityProofingProviderTemplate) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	for _, val := range m.IdcsPreventedOperations {
		if _, ok := GetMappingIdcsPreventedOperationsEnum(string(val)); !ok && val != "" {
			errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for IdcsPreventedOperations: %s. Supported values are: %s.", val, strings.Join(GetIdcsPreventedOperationsEnumStringValues(), ",")))
		}
	}

	for _, val := range m.IdcsLockedOperations {
		if _, ok := GetMappingIdentityProofingProviderTemplateIdcsLockedOperationsEnum(string(val)); !ok && val != "" {
			errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for IdcsLockedOperations: %s. Supported values are: %s.", val, strings.Join(GetIdentityProofingProviderTemplateIdcsLockedOperationsEnumStringValues(), ",")))
		}
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// IdentityProofingProviderTemplateIdcsLockedOperationsEnum Enum with underlying type: string
type IdentityProofingProviderTemplateIdcsLockedOperationsEnum string

// Set of constants representing the allowable values for IdentityProofingProviderTemplateIdcsLockedOperationsEnum
const (
	IdentityProofingProviderTemplateIdcsLockedOperationsReplace IdentityProofingProviderTemplateIdcsLockedOperationsEnum = "replace"
	IdentityProofingProviderTemplateIdcsLockedOperationsUpdate  IdentityProofingProviderTemplateIdcsLockedOperationsEnum = "update"
	IdentityProofingProviderTemplateIdcsLockedOperationsDelete  IdentityProofingProviderTemplateIdcsLockedOperationsEnum = "delete"
)

var mappingIdentityProofingProviderTemplateIdcsLockedOperationsEnum = map[string]IdentityProofingProviderTemplateIdcsLockedOperationsEnum{
	"replace": IdentityProofingProviderTemplateIdcsLockedOperationsReplace,
	"update":  IdentityProofingProviderTemplateIdcsLockedOperationsUpdate,
	"delete":  IdentityProofingProviderTemplateIdcsLockedOperationsDelete,
}

var mappingIdentityProofingProviderTemplateIdcsLockedOperationsEnumLowerCase = map[string]IdentityProofingProviderTemplateIdcsLockedOperationsEnum{
	"replace": IdentityProofingProviderTemplateIdcsLockedOperationsReplace,
	"update":  IdentityProofingProviderTemplateIdcsLockedOperationsUpdate,
	"delete":  IdentityProofingProviderTemplateIdcsLockedOperationsDelete,
}

// GetIdentityProofingProviderTemplateIdcsLockedOperationsEnumValues Enumerates the set of values for IdentityProofingProviderTemplateIdcsLockedOperationsEnum
func GetIdentityProofingProviderTemplateIdcsLockedOperationsEnumValues() []IdentityProofingProviderTemplateIdcsLockedOperationsEnum {
	values := make([]IdentityProofingProviderTemplateIdcsLockedOperationsEnum, 0)
	for _, v := range mappingIdentityProofingProviderTemplateIdcsLockedOperationsEnum {
		values = append(values, v)
	}
	return values
}

// GetIdentityProofingProviderTemplateIdcsLockedOperationsEnumStringValues Enumerates the set of values in String for IdentityProofingProviderTemplateIdcsLockedOperationsEnum
func GetIdentityProofingProviderTemplateIdcsLockedOperationsEnumStringValues() []string {
	return []string{
		"replace",
		"update",
		"delete",
	}
}

// GetMappingIdentityProofingProviderTemplateIdcsLockedOperationsEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingIdentityProofingProviderTemplateIdcsLockedOperationsEnum(val string) (IdentityProofingProviderTemplateIdcsLockedOperationsEnum, bool) {
	enum, ok := mappingIdentityProofingProviderTemplateIdcsLockedOperationsEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
