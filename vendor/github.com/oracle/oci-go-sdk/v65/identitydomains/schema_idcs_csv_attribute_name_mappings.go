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

// SchemaIdcsCsvAttributeNameMappings Csv meta data for those resource type attributes which can be imported / exported from / to csv.
type SchemaIdcsCsvAttributeNameMappings struct {

	// The CSV column header name that maps to this attribute.
	// **SCIM++ Properties:**
	//  - caseExact: false
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: false
	//  - returned: default
	//  - type: string
	//  - uniqueness: none
	ColumnHeaderName *string `mandatory:"false" json:"columnHeaderName"`

	// The deprecated CSV column header name that maps to this attribute.
	// **SCIM++ Properties:**
	//  - caseExact: false
	//  - multiValued: false
	//  - mutability: readOnly
	//  - required: false
	//  - returned: default
	//  - type: string
	//  - uniqueness: none
	DeprecatedColumnHeaderName *string `mandatory:"false" json:"deprecatedColumnHeaderName"`

	// The attribute path that the CSV column header name maps to for complex multi-valued attributes.
	// **SCIM++ Properties:**
	//  - caseExact: false
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: false
	//  - returned: default
	//  - type: string
	//  - uniqueness: none
	MapsTo *string `mandatory:"false" json:"mapsTo"`

	// The default value to be used during import processing in case the CSV column header is not present or value is not given in the import CSV.
	// **SCIM++ Properties:**
	//  - caseExact: false
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: false
	//  - returned: default
	//  - type: string
	//  - uniqueness: none
	DefaultValue *string `mandatory:"false" json:"defaultValue"`

	// If values of the CSV column header name can contain multiple values, this attribute specifies the delimiter to be used. For example, Group's \"User Members\" CSV column header is multi-valued and it's delimiter is a semi-colon (\";\").
	// **SCIM++ Properties:**
	//  - caseExact: false
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: false
	//  - returned: default
	//  - type: string
	//  - uniqueness: none
	MultiValueDelimiter *string `mandatory:"false" json:"multiValueDelimiter"`

	// This specifies the Csv Header for resolving Resource Type for this Column Header
	// **SCIM++ Properties:**
	//  - caseExact: false
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: false
	//  - returned: default
	//  - type: string
	//  - uniqueness: none
	CsvColumnForResolvingResourceType *string `mandatory:"false" json:"csvColumnForResolvingResourceType"`

	// This attribute gives a maps for resolving Resource Type after reading it's value from \"csvColumnForResolvingResourceType\" attribute
	// **SCIM++ Properties:**
	//  - caseExact: false
	//  - multiValued: true
	//  - mutability: readWrite
	//  - required: false
	//  - returned: default
	//  - type: complex
	//  - uniqueness: none
	ResolveValueUsingResourceType []SchemaResolveValueUsingResourceType `mandatory:"false" json:"resolveValueUsingResourceType"`

	// This attribute specifies the mapping of \"uniqueAttributeNameForDisplay\" attributes(s) of the referenced resource with the columnHeaderName(s). This attribute should be given in the idcsCsvAttributeNameMappings when uniqueAttributeNameForDisplay contains more than one attribute.
	// **SCIM++ Properties:**
	//  - caseExact: false
	//  - multiValued: true
	//  - mutability: readWrite
	//  - required: false
	//  - returned: default
	//  - type: complex
	//  - uniqueness: none
	ReferencedResourceTypeUniqueAttributeNameMappings []SchemaReferencedResourceTypeUniqueAttributeNameMappings `mandatory:"false" json:"referencedResourceTypeUniqueAttributeNameMappings"`
}

func (m SchemaIdcsCsvAttributeNameMappings) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m SchemaIdcsCsvAttributeNameMappings) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
