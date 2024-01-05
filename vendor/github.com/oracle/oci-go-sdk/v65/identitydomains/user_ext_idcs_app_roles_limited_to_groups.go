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

// UserExtIdcsAppRolesLimitedToGroups Description:
// **Added In:** 19.2.1
// **SCIM++ Properties:**
//   - idcsCompositeKey: [value, idcsAppRoleId]
//   - idcsSearchable: true
//   - multiValued: true
//   - mutability: readOnly
//   - required: false
//   - returned: request
//   - type: complex
type UserExtIdcsAppRolesLimitedToGroups struct {

	// The id of the Oracle Identity Cloud Service AppRole grant limited to one or more Groups.
	// **Added In:** 19.2.1
	// **SCIM++ Properties:**
	//  - idcsCsvAttributeName: IDCS AppRole Name
	//  - idcsSearchable: true
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: true
	//  - returned: default
	//  - type: string
	//  - uniqueness: none
	IdcsAppRoleId *string `mandatory:"true" json:"idcsAppRoleId"`

	// The id of a Group the AppRole Grant is limited to
	// **Added In:** 19.2.1
	// **SCIM++ Properties:**
	//  - idcsCsvAttributeName: Group Name
	//  - idcsSearchable: true
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: true
	//  - returned: default
	//  - type: string
	//  - uniqueness: none
	Value *string `mandatory:"true" json:"value"`

	// The URI of the SCIM resource representing the Group manager.  RECOMMENDED.
	// **Added In:** 19.2.1
	// **SCIM++ Properties:**
	//  - idcsSearchable: false
	//  - multiValued: false
	//  - mutability: readOnly
	//  - required: false
	//  - returned: default
	//  - type: reference
	//  - uniqueness: none
	Ref *string `mandatory:"false" json:"$ref"`

	// Group display name
	// **Added In:** 19.2.1
	// **SCIM++ Properties:**
	//  - caseExact: false
	//  - idcsSearchable: false
	//  - multiValued: false
	//  - mutability: readOnly
	//  - required: false
	//  - returned: default
	//  - type: string
	//  - uniqueness: none
	Display *string `mandatory:"false" json:"display"`

	// The ocid of a Group the AppRole Grant is limited to
	// **Added In:** 2202230830
	// **SCIM++ Properties:**
	//  - idcsCsvAttributeName: Group Ocid
	//  - idcsSearchable: true
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: false
	//  - returned: default
	//  - type: string
	//  - uniqueness: none
	Ocid *string `mandatory:"false" json:"ocid"`
}

func (m UserExtIdcsAppRolesLimitedToGroups) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m UserExtIdcsAppRolesLimitedToGroups) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
