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

// ExtensionEnterprise20User Enterprise User
type ExtensionEnterprise20User struct {

	// Numeric or alphanumeric identifier assigned to  a person, typically based on order of hire or association with an organization.
	// **SCIM++ Properties:**
	//  - caseExact: false
	//  - idcsCsvAttributeName: Employee Number
	//  - idcsCsvAttributeNameMappings: [[columnHeaderName:Employee Number]]
	//  - idcsPii: true
	//  - idcsSearchable: true
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: false
	//  - returned: default
	//  - type: string
	//  - uniqueness: none
	EmployeeNumber *string `mandatory:"false" json:"employeeNumber"`

	// Identifies the name of a cost center.
	// **SCIM++ Properties:**
	//  - caseExact: false
	//  - idcsCsvAttributeName: Cost Center
	//  - idcsCsvAttributeNameMappings: [[columnHeaderName:Cost Center]]
	//  - idcsPii: true
	//  - idcsSearchable: true
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: false
	//  - returned: default
	//  - type: string
	//  - uniqueness: none
	CostCenter *string `mandatory:"false" json:"costCenter"`

	// Identifies the name of an organization.
	// **SCIM++ Properties:**
	//  - caseExact: false
	//  - idcsCsvAttributeName: Organization
	//  - idcsCsvAttributeNameMappings: [[columnHeaderName:Organization Name, deprecatedColumnHeaderName:Organization]]
	//  - idcsPii: true
	//  - idcsSearchable: true
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: false
	//  - returned: default
	//  - type: string
	//  - uniqueness: none
	Organization *string `mandatory:"false" json:"organization"`

	// Identifies the name of a division.
	// **SCIM++ Properties:**
	//  - caseExact: false
	//  - idcsCsvAttributeName: Division
	//  - idcsCsvAttributeNameMappings: [[columnHeaderName:Division]]
	//  - idcsPii: true
	//  - idcsSearchable: true
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: false
	//  - returned: default
	//  - type: string
	//  - uniqueness: none
	Division *string `mandatory:"false" json:"division"`

	// Identifies the name of a department.
	// **SCIM++ Properties:**
	//  - caseExact: false
	//  - idcsCsvAttributeName: Department
	//  - idcsCsvAttributeNameMappings: [[columnHeaderName:Department]]
	//  - idcsPii: true
	//  - idcsSearchable: true
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: false
	//  - returned: default
	//  - type: string
	//  - uniqueness: none
	Department *string `mandatory:"false" json:"department"`

	Manager *UserExtManager `mandatory:"false" json:"manager"`
}

func (m ExtensionEnterprise20User) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ExtensionEnterprise20User) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
