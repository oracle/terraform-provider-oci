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

// MappedAttributeAttributeMappings A list of attribute mappings between Oracle Identity Cloud Service Resource Type and Account Object Class
type MappedAttributeAttributeMappings struct {

	// The name or expression of an attribute defined in the schema of the Managed Object Class. This is the SCIM compliant Oracle Identity Cloud Service Name of the attribute that maps to the \"idcsName\" attribute in the schema of an Managed Object Class.
	// **SCIM++ Properties:**
	//  - idcsSearchable: true
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: true
	//  - returned: default
	//  - type: string
	//  - uniqueness: none
	ManagedObjectAttributeName *string `mandatory:"true" json:"managedObjectAttributeName"`

	// The name or expression of an attribute that corresponds to the Oracle Identity Cloud Service Resource referred in the \"idcsResourceType\" attribute.
	// **SCIM++ Properties:**
	//  - idcsSearchable: false
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: true
	//  - returned: default
	//  - type: string
	//  - uniqueness: none
	IdcsAttributeName *string `mandatory:"true" json:"idcsAttributeName"`

	// If true, indicates that this attribute must have a value. This attribute maps to the \"required\" sub-attribute in the schema of an Managed Object Class.
	// **SCIM++ Properties:**
	//  - idcsSearchable: false
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: false
	//  - returned: default
	//  - type: boolean
	//  - uniqueness: none
	Required *bool `mandatory:"false" json:"required"`

	// If specified, indicates a subset of mappedActions to which this attribute-mapping applies.If not specified, this attribute-mapping applies to all mappedActions that use mappedAttributes
	// **SCIM++ Properties:**
	//  - caseExact: true
	//  - idcsSearchable: false
	//  - multiValued: true
	//  - mutability: readWrite
	//  - required: false
	//  - returned: default
	//  - type: string
	//  - uniqueness: none
	AppliesToActions []MappedAttributeAttributeMappingsAppliesToActionsEnum `mandatory:"false" json:"appliesToActions,omitempty"`

	// Indicates the format of the assertion attribute. Also stores AttributeNamespace for WSFed1.1.
	// **Added In:** 17.4.2
	// **SCIM++ Properties:**
	//  - caseExact: true
	//  - idcsSearchable: false
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: false
	//  - returned: default
	//  - type: string
	//  - uniqueness: none
	SamlFormat *string `mandatory:"false" json:"samlFormat"`
}

func (m MappedAttributeAttributeMappings) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m MappedAttributeAttributeMappings) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	for _, val := range m.AppliesToActions {
		if _, ok := GetMappingMappedAttributeAttributeMappingsAppliesToActionsEnum(string(val)); !ok && val != "" {
			errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for AppliesToActions: %s. Supported values are: %s.", val, strings.Join(GetMappedAttributeAttributeMappingsAppliesToActionsEnumStringValues(), ",")))
		}
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MappedAttributeAttributeMappingsAppliesToActionsEnum Enum with underlying type: string
type MappedAttributeAttributeMappingsAppliesToActionsEnum string

// Set of constants representing the allowable values for MappedAttributeAttributeMappingsAppliesToActionsEnum
const (
	MappedAttributeAttributeMappingsAppliesToActionsCreate MappedAttributeAttributeMappingsAppliesToActionsEnum = "create"
	MappedAttributeAttributeMappingsAppliesToActionsUpdate MappedAttributeAttributeMappingsAppliesToActionsEnum = "update"
)

var mappingMappedAttributeAttributeMappingsAppliesToActionsEnum = map[string]MappedAttributeAttributeMappingsAppliesToActionsEnum{
	"create": MappedAttributeAttributeMappingsAppliesToActionsCreate,
	"update": MappedAttributeAttributeMappingsAppliesToActionsUpdate,
}

var mappingMappedAttributeAttributeMappingsAppliesToActionsEnumLowerCase = map[string]MappedAttributeAttributeMappingsAppliesToActionsEnum{
	"create": MappedAttributeAttributeMappingsAppliesToActionsCreate,
	"update": MappedAttributeAttributeMappingsAppliesToActionsUpdate,
}

// GetMappedAttributeAttributeMappingsAppliesToActionsEnumValues Enumerates the set of values for MappedAttributeAttributeMappingsAppliesToActionsEnum
func GetMappedAttributeAttributeMappingsAppliesToActionsEnumValues() []MappedAttributeAttributeMappingsAppliesToActionsEnum {
	values := make([]MappedAttributeAttributeMappingsAppliesToActionsEnum, 0)
	for _, v := range mappingMappedAttributeAttributeMappingsAppliesToActionsEnum {
		values = append(values, v)
	}
	return values
}

// GetMappedAttributeAttributeMappingsAppliesToActionsEnumStringValues Enumerates the set of values in String for MappedAttributeAttributeMappingsAppliesToActionsEnum
func GetMappedAttributeAttributeMappingsAppliesToActionsEnumStringValues() []string {
	return []string{
		"create",
		"update",
	}
}

// GetMappingMappedAttributeAttributeMappingsAppliesToActionsEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingMappedAttributeAttributeMappingsAppliesToActionsEnum(val string) (MappedAttributeAttributeMappingsAppliesToActionsEnum, bool) {
	enum, ok := mappingMappedAttributeAttributeMappingsAppliesToActionsEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
