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

// AppOutboundAssertionAttributes Use to construct the outgoing SAML attributes
// **Added In:** 18.2.6
// **SCIM++ Properties:**
//   - caseExact: false
//   - idcsCompositeKey: [value]
//   - idcsSearchable: true
//   - multiValued: false
//   - mutability: readOnly
//   - required: false
//   - returned: default
//   - type: complex
//   - uniqueness: none
type AppOutboundAssertionAttributes struct {

	// Mapped Attribute identifier
	// **Added In:** 18.2.6
	// **SCIM++ Properties:**
	//  - caseExact: true
	//  - idcsSearchable: true
	//  - mutability: readOnly
	//  - required: true
	//  - returned: default
	//  - type: string
	//  - uniqueness: none
	Value *string `mandatory:"true" json:"value"`

	// Mapped Attribute URI
	// **Added In:** 18.2.6
	// **SCIM++ Properties:**
	//  - idcsSearchable: false
	//  - mutability: readOnly
	//  - required: false
	//  - returned: default
	//  - type: reference
	//  - uniqueness: none
	Ref *string `mandatory:"false" json:"$ref"`

	// Mapped Attribute Direction
	// **Added In:** 18.2.6
	// **SCIM++ Properties:**
	//  - idcsSearchable: true
	//  - multiValued: false
	//  - mutability: readOnly
	//  - required: false
	//  - returned: default
	//  - type: string
	//  - uniqueness: none
	Direction *string `mandatory:"false" json:"direction"`
}

func (m AppOutboundAssertionAttributes) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m AppOutboundAssertionAttributes) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
