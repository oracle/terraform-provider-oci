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

// MyDeviceUser Device member
// **Deprecated Since: 17.3.4**
// **SCIM++ Properties:**
//   - caseExact: false
//   - idcsSearchable: true
//   - multiValued: false
//   - mutability: immutable
//   - required: true
//   - returned: default
//   - type: complex
//   - uniqueness: none
type MyDeviceUser struct {

	// The identifier of the user
	// **Deprecated Since: 17.3.4**
	// **SCIM++ Properties:**
	//  - caseExact: true
	//  - idcsSearchable: true
	//  - multiValued: false
	//  - mutability: immutable
	//  - required: true
	//  - returned: always
	//  - type: string
	//  - uniqueness: none
	//  - idcsRequiresWriteForAccessFlows: true
	//  - idcsRequiresImmediateReadAfterWriteForAccessFlows: true
	Value *string `mandatory:"true" json:"value"`

	// The URI that corresponds to the member Resource of this device
	// **Deprecated Since: 17.3.4**
	// **SCIM++ Properties:**
	//  - idcsSearchable: false
	//  - multiValued: false
	//  - mutability: readOnly
	//  - required: false
	//  - returned: default
	//  - type: reference
	//  - uniqueness: none
	//  - idcsRequiresWriteForAccessFlows: true
	//  - idcsRequiresImmediateReadAfterWriteForAccessFlows: true
	Ref *string `mandatory:"false" json:"$ref"`

	// User display name
	// **Deprecated Since: 17.3.4**
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

	// The OCID of the user
	// **Added In:** 2105091740
	// **SCIM++ Properties:**
	//  - caseExact: true
	//  - idcsSearchable: true
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: false
	//  - returned: always
	//  - type: string
	//  - uniqueness: none
	//  - idcsRequiresWriteForAccessFlows: true
	//  - idcsRequiresImmediateReadAfterWriteForAccessFlows: true
	Ocid *string `mandatory:"false" json:"ocid"`
}

func (m MyDeviceUser) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m MyDeviceUser) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
