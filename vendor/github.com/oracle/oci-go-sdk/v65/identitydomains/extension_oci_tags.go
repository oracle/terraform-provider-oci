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

// ExtensionOciTags OCI Tags.
type ExtensionOciTags struct {

	// OCI Freeform Tags
	// **Added In:** 2011192329
	// **SCIM++ Properties:**
	//  - idcsCompositeKey: [key, value]
	//  - idcsSearchable: true
	//  - type: complex
	//  - required: false
	//  - mutability: readWrite
	//  - returned: default
	//  - multiValued: true
	FreeformTags []FreeformTags `mandatory:"false" json:"freeformTags"`

	// OCI Defined Tags
	// **Added In:** 2011192329
	// **SCIM++ Properties:**
	//  - idcsCompositeKey: [namespace, key, value]
	//  - type: complex
	//  - idcsSearchable: true
	//  - required: false
	//  - mutability: readWrite
	//  - multiValued: true
	//  - returned: default
	DefinedTags []DefinedTags `mandatory:"false" json:"definedTags"`

	// OCI Tag slug
	// **Added In:** 2011192329
	// **SCIM++ Properties:**
	//  - type: binary
	//  - mutability: readOnly
	//  - returned: request
	TagSlug *interface{} `mandatory:"false" json:"tagSlug"`
}

func (m ExtensionOciTags) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ExtensionOciTags) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
