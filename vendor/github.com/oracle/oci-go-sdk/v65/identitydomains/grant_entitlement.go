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

// GrantEntitlement The entitlement or privilege that is being granted
// **SCIM++ Properties:**
//   - idcsCsvAttributeNameMappings: [[columnHeaderName:Entitlement Value, csvColumnForResolvingResourceType:Entitlement Name, mapsTo:entitlement.attributeValue, referencedResourceTypeUniqueAttributeNameMappings:[[mapsFromColumnName:Entitlement Value, resourceTypeAttributeName:displayName], [mapsFromColumnName:App Name, resourceTypeAttributeName:app.display]], resolveValueUsingResourceType:[[resolveBy:AppRole, valueToBeResolved:appRoles]]], [columnHeaderName:Entitlement Name, defaultValue:appRoles, mapsTo:entitlement.attributeName]]
//   - idcsSearchable: true
//   - multiValued: false
//   - mutability: immutable
//   - required: false
//   - returned: default
//   - type: complex
type GrantEntitlement struct {

	// The name of the attribute whose value (specified by attributeValue) confers privilege within the service-instance (specified by app).
	// **SCIM++ Properties:**
	//  - caseExact: false
	//  - idcsSearchable: true
	//  - multiValued: false
	//  - mutability: immutable
	//  - required: true
	//  - returned: default
	//  - type: string
	//  - uniqueness: none
	AttributeName *string `mandatory:"true" json:"attributeName"`

	// The value of the attribute (specified by attributeName) that confers privilege within the service-instance (specified by app).  If attributeName is 'appRoles', then attributeValue is the ID of the AppRole.
	// **SCIM++ Properties:**
	//  - caseExact: true
	//  - idcsCsvAttributeName: Display Name
	//  - idcsSearchable: true
	//  - multiValued: false
	//  - mutability: immutable
	//  - required: true
	//  - returned: default
	//  - type: string
	//  - uniqueness: none
	AttributeValue *string `mandatory:"true" json:"attributeValue"`
}

func (m GrantEntitlement) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m GrantEntitlement) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
