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

// IdentitySettingsMyProfile Whether to allow users to update their own profile.
// **Added In:** 2207040824
// **SCIM++ Properties:**
//   - caseExact: false
//   - multiValued: false
//   - required: false
//   - type: complex
//   - uniqueness: none
type IdentitySettingsMyProfile struct {

	// Whether to allow users to change their own password.
	// **Added In:** 2207040824
	// **SCIM++ Properties:**
	//  - caseExact: false
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: false
	//  - returned: default
	//  - type: boolean
	//  - uniqueness: none
	AllowEndUsersToChangeTheirPassword *bool `mandatory:"false" json:"allowEndUsersToChangeTheirPassword"`

	// Whether to allow users to link or unlink their support accounts.
	// **Added In:** 2207040824
	// **SCIM++ Properties:**
	//  - caseExact: false
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: false
	//  - returned: default
	//  - type: boolean
	//  - uniqueness: none
	AllowEndUsersToLinkTheirSupportAccount *bool `mandatory:"false" json:"allowEndUsersToLinkTheirSupportAccount"`

	// Whether to allow users to update their security settings.
	// **Added In:** 2207040824
	// **SCIM++ Properties:**
	//  - caseExact: false
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: false
	//  - returned: default
	//  - type: boolean
	//  - uniqueness: none
	AllowEndUsersToUpdateTheirSecuritySettings *bool `mandatory:"false" json:"allowEndUsersToUpdateTheirSecuritySettings"`

	// Whether to allow users to update their capabilities.
	// **Added In:** 2207040824
	// **SCIM++ Properties:**
	//  - caseExact: false
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: false
	//  - returned: default
	//  - type: boolean
	//  - uniqueness: none
	AllowEndUsersToManageTheirCapabilities *bool `mandatory:"false" json:"allowEndUsersToManageTheirCapabilities"`
}

func (m IdentitySettingsMyProfile) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m IdentitySettingsMyProfile) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
