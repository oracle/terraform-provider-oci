// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
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

// MappedAttribute Schema for MappedAttribute resource.
type MappedAttribute struct {

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

	// The Oracle Identity Cloud Service Resource Type for which the mapping is being done
	// **SCIM++ Properties:**
	//  - caseExact: true
	//  - idcsSearchable: true
	//  - multiValued: false
	//  - mutability: immutable
	//  - required: true
	//  - returned: default
	//  - type: string
	//  - uniqueness: none
	IdcsResourceType MappedAttributeIdcsResourceTypeEnum `mandatory:"true" json:"idcsResourceType"`

	// The Reference Resource Type that holds the mapping
	// **Added In:** 17.4.2
	// **SCIM++ Properties:**
	//  - caseExact: true
	//  - idcsSearchable: true
	//  - multiValued: false
	//  - mutability: readOnly
	//  - required: true
	//  - returned: default
	//  - type: string
	//  - uniqueness: none
	RefResourceType MappedAttributeRefResourceTypeEnum `mandatory:"true" json:"refResourceType"`

	// Soft Reference to store Resource ID that holds the mapping
	// **Added In:** 17.4.6
	// **SCIM++ Properties:**
	//  - caseExact: true
	//  - idcsSearchable: true
	//  - multiValued: false
	//  - mutability: readOnly
	//  - required: true
	//  - returned: default
	//  - type: string
	//  - uniqueness: none
	RefResourceID *string `mandatory:"true" json:"refResourceID"`

	// Direction of the attribute mapping. inbound indicates mapping is from source ManagedObjectClass to Oracle Identity Cloud Service ResourceType. outbound indicates mapping is from Oracle Identity Cloud Service Resource Type to target ManagedObjectClass.
	// **SCIM++ Properties:**
	//  - caseExact: true
	//  - idcsSearchable: true
	//  - multiValued: false
	//  - mutability: immutable
	//  - required: true
	//  - returned: default
	//  - type: string
	//  - uniqueness: none
	Direction MappedAttributeDirectionEnum `mandatory:"true" json:"direction"`

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

	// A list of attribute mappings between Oracle Identity Cloud Service Resource Type and Account Object Class
	// **SCIM++ Properties:**
	//  - idcsCompositeKey: [idcsAttributeName, managedObjectAttributeName]
	//  - idcsSearchable: true
	//  - multiValued: true
	//  - mutability: readWrite
	//  - required: false
	//  - returned: default
	//  - type: complex
	//  - uniqueness: none
	AttributeMappings []MappedAttributeAttributeMappings `mandatory:"false" json:"attributeMappings"`
}

func (m MappedAttribute) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m MappedAttribute) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingMappedAttributeIdcsResourceTypeEnum(string(m.IdcsResourceType)); !ok && m.IdcsResourceType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for IdcsResourceType: %s. Supported values are: %s.", m.IdcsResourceType, strings.Join(GetMappedAttributeIdcsResourceTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingMappedAttributeRefResourceTypeEnum(string(m.RefResourceType)); !ok && m.RefResourceType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for RefResourceType: %s. Supported values are: %s.", m.RefResourceType, strings.Join(GetMappedAttributeRefResourceTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingMappedAttributeDirectionEnum(string(m.Direction)); !ok && m.Direction != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Direction: %s. Supported values are: %s.", m.Direction, strings.Join(GetMappedAttributeDirectionEnumStringValues(), ",")))
	}

	for _, val := range m.IdcsPreventedOperations {
		if _, ok := GetMappingIdcsPreventedOperationsEnum(string(val)); !ok && val != "" {
			errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for IdcsPreventedOperations: %s. Supported values are: %s.", val, strings.Join(GetIdcsPreventedOperationsEnumStringValues(), ",")))
		}
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MappedAttributeIdcsResourceTypeEnum Enum with underlying type: string
type MappedAttributeIdcsResourceTypeEnum string

// Set of constants representing the allowable values for MappedAttributeIdcsResourceTypeEnum
const (
	MappedAttributeIdcsResourceTypeUser      MappedAttributeIdcsResourceTypeEnum = "User"
	MappedAttributeIdcsResourceTypeGroup     MappedAttributeIdcsResourceTypeEnum = "Group"
	MappedAttributeIdcsResourceTypeContainer MappedAttributeIdcsResourceTypeEnum = "Container"
)

var mappingMappedAttributeIdcsResourceTypeEnum = map[string]MappedAttributeIdcsResourceTypeEnum{
	"User":      MappedAttributeIdcsResourceTypeUser,
	"Group":     MappedAttributeIdcsResourceTypeGroup,
	"Container": MappedAttributeIdcsResourceTypeContainer,
}

var mappingMappedAttributeIdcsResourceTypeEnumLowerCase = map[string]MappedAttributeIdcsResourceTypeEnum{
	"user":      MappedAttributeIdcsResourceTypeUser,
	"group":     MappedAttributeIdcsResourceTypeGroup,
	"container": MappedAttributeIdcsResourceTypeContainer,
}

// GetMappedAttributeIdcsResourceTypeEnumValues Enumerates the set of values for MappedAttributeIdcsResourceTypeEnum
func GetMappedAttributeIdcsResourceTypeEnumValues() []MappedAttributeIdcsResourceTypeEnum {
	values := make([]MappedAttributeIdcsResourceTypeEnum, 0)
	for _, v := range mappingMappedAttributeIdcsResourceTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetMappedAttributeIdcsResourceTypeEnumStringValues Enumerates the set of values in String for MappedAttributeIdcsResourceTypeEnum
func GetMappedAttributeIdcsResourceTypeEnumStringValues() []string {
	return []string{
		"User",
		"Group",
		"Container",
	}
}

// GetMappingMappedAttributeIdcsResourceTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingMappedAttributeIdcsResourceTypeEnum(val string) (MappedAttributeIdcsResourceTypeEnum, bool) {
	enum, ok := mappingMappedAttributeIdcsResourceTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// MappedAttributeRefResourceTypeEnum Enum with underlying type: string
type MappedAttributeRefResourceTypeEnum string

// Set of constants representing the allowable values for MappedAttributeRefResourceTypeEnum
const (
	MappedAttributeRefResourceTypeAccountobjectclass MappedAttributeRefResourceTypeEnum = "AccountObjectClass"
	MappedAttributeRefResourceTypeApp                MappedAttributeRefResourceTypeEnum = "App"
	MappedAttributeRefResourceTypeManagedobjectclass MappedAttributeRefResourceTypeEnum = "ManagedObjectClass"
	MappedAttributeRefResourceTypeIdentityprovider   MappedAttributeRefResourceTypeEnum = "IdentityProvider"
)

var mappingMappedAttributeRefResourceTypeEnum = map[string]MappedAttributeRefResourceTypeEnum{
	"AccountObjectClass": MappedAttributeRefResourceTypeAccountobjectclass,
	"App":                MappedAttributeRefResourceTypeApp,
	"ManagedObjectClass": MappedAttributeRefResourceTypeManagedobjectclass,
	"IdentityProvider":   MappedAttributeRefResourceTypeIdentityprovider,
}

var mappingMappedAttributeRefResourceTypeEnumLowerCase = map[string]MappedAttributeRefResourceTypeEnum{
	"accountobjectclass": MappedAttributeRefResourceTypeAccountobjectclass,
	"app":                MappedAttributeRefResourceTypeApp,
	"managedobjectclass": MappedAttributeRefResourceTypeManagedobjectclass,
	"identityprovider":   MappedAttributeRefResourceTypeIdentityprovider,
}

// GetMappedAttributeRefResourceTypeEnumValues Enumerates the set of values for MappedAttributeRefResourceTypeEnum
func GetMappedAttributeRefResourceTypeEnumValues() []MappedAttributeRefResourceTypeEnum {
	values := make([]MappedAttributeRefResourceTypeEnum, 0)
	for _, v := range mappingMappedAttributeRefResourceTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetMappedAttributeRefResourceTypeEnumStringValues Enumerates the set of values in String for MappedAttributeRefResourceTypeEnum
func GetMappedAttributeRefResourceTypeEnumStringValues() []string {
	return []string{
		"AccountObjectClass",
		"App",
		"ManagedObjectClass",
		"IdentityProvider",
	}
}

// GetMappingMappedAttributeRefResourceTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingMappedAttributeRefResourceTypeEnum(val string) (MappedAttributeRefResourceTypeEnum, bool) {
	enum, ok := mappingMappedAttributeRefResourceTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// MappedAttributeDirectionEnum Enum with underlying type: string
type MappedAttributeDirectionEnum string

// Set of constants representing the allowable values for MappedAttributeDirectionEnum
const (
	MappedAttributeDirectionInbound  MappedAttributeDirectionEnum = "inbound"
	MappedAttributeDirectionOutbound MappedAttributeDirectionEnum = "outbound"
)

var mappingMappedAttributeDirectionEnum = map[string]MappedAttributeDirectionEnum{
	"inbound":  MappedAttributeDirectionInbound,
	"outbound": MappedAttributeDirectionOutbound,
}

var mappingMappedAttributeDirectionEnumLowerCase = map[string]MappedAttributeDirectionEnum{
	"inbound":  MappedAttributeDirectionInbound,
	"outbound": MappedAttributeDirectionOutbound,
}

// GetMappedAttributeDirectionEnumValues Enumerates the set of values for MappedAttributeDirectionEnum
func GetMappedAttributeDirectionEnumValues() []MappedAttributeDirectionEnum {
	values := make([]MappedAttributeDirectionEnum, 0)
	for _, v := range mappingMappedAttributeDirectionEnum {
		values = append(values, v)
	}
	return values
}

// GetMappedAttributeDirectionEnumStringValues Enumerates the set of values in String for MappedAttributeDirectionEnum
func GetMappedAttributeDirectionEnumStringValues() []string {
	return []string{
		"inbound",
		"outbound",
	}
}

// GetMappingMappedAttributeDirectionEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingMappedAttributeDirectionEnum(val string) (MappedAttributeDirectionEnum, bool) {
	enum, ok := mappingMappedAttributeDirectionEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
