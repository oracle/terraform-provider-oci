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

// UserExtLocked A complex attribute that indicates an account is locked (blocking any new sessions).
// **SCIM++ Properties:**
//   - idcsCsvAttributeNameMappings: [[columnHeaderName:Locked, mapsTo:locked.on], [columnHeaderName:Locked Reason, mapsTo:locked.reason], [columnHeaderName:Locked Date, mapsTo:locked.lockDate]]
//   - idcsSearchable: false
//   - idcsAllowUpdatesInReadOnlyMode: true
//   - multiValued: false
//   - mutability: readWrite
//   - required: false
//   - returned: default
//   - type: complex
//   - uniqueness: none
type UserExtLocked struct {

	// Indicates the reason for locking the account. Valid values are: 0 - failed password login attempts, 1 - admin lock, 2 - failed reset password attempts, 3 - failed MFA login attempts, 4 - failed MFA login attempts for federated user, 5 - failed Database login attempts
	// **SCIM++ Properties:**
	//  - idcsSearchable: false
	//  - multiValued: false
	//  - mutability: readWrite
	//  - idcsRequiresWriteForAccessFlows: true
	//  - idcsRequiresImmediateReadAfterWriteForAccessFlows: true
	//  - required: false
	//  - returned: default
	//  - type: integer
	//  - uniqueness: none
	Reason *int `mandatory:"false" json:"reason"`

	// Indicates that the account is locked.
	// **SCIM++ Properties:**
	//  - caseExact: false
	//  - idcsSearchable: true
	//  - multiValued: false
	//  - mutability: readWrite
	//  - idcsRequiresWriteForAccessFlows: true
	//  - idcsRequiresImmediateReadAfterWriteForAccessFlows: true
	//  - required: false
	//  - returned: default
	//  - type: boolean
	//  - uniqueness: none
	On *bool `mandatory:"false" json:"on"`

	// The date and time that the current resource was locked.
	// **SCIM++ Properties:**
	//  - idcsSearchable: false
	//  - multiValued: false
	//  - mutability: readWrite
	//  - idcsRequiresWriteForAccessFlows: true
	//  - required: false
	//  - returned: default
	//  - type: dateTime
	//  - uniqueness: none
	LockDate *string `mandatory:"false" json:"lockDate"`

	// Indicates whether the user password is expired. If this value is false, password expiry is still evaluated during user login.
	// **Added In:** 20.1.3
	// **SCIM++ Properties:**
	//  - caseExact: false
	//  - idcsSearchable: false
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: false
	//  - returned: request
	//  - type: boolean
	//  - uniqueness: none
	Expired *bool `mandatory:"false" json:"expired"`
}

func (m UserExtLocked) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m UserExtLocked) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
