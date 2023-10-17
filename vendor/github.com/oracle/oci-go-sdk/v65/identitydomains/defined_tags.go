// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
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

// DefinedTags OCI Defined Tags
// **Added In:** 2011192329
// **SCIM++ Properties:**
//   - idcsCompositeKey: [namespace, key, value]
//   - type: complex
//   - idcsSearchable: true
//   - required: false
//   - mutability: readWrite
//   - multiValued: true
//   - returned: default
type DefinedTags struct {

	// OCI Tag namespace
	// **Added In:** 2011192329
	// **SCIM++ Properties:**
	//  - caseExact: false
	//  - type: string
	//  - required: true
	//  - mutability: readWrite
	//  - returned: default
	//  - idcsSearchable: true
	//  - uniqueness: none
	Namespace *string `mandatory:"true" json:"namespace"`

	// OCI Tag key
	// **Added In:** 2011192329
	// **SCIM++ Properties:**
	//  - caseExact: false
	//  - type: string
	//  - required: true
	//  - mutability: readWrite
	//  - returned: default
	//  - idcsSearchable: true
	//  - uniqueness: none
	Key *string `mandatory:"true" json:"key"`

	// OCI Tag value
	// **Added In:** 2011192329
	// **SCIM++ Properties:**
	//  - caseExact: false
	//  - required: true
	//  - idcsReturnEmptyWhenNull: true
	//  - mutability: readWrite
	//  - returned: default
	//  - type: string
	//  - idcsSearchable: true
	//  - uniqueness: none
	Value *string `mandatory:"true" json:"value"`
}

func (m DefinedTags) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m DefinedTags) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
