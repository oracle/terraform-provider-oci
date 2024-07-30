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

// UserAttributesSettingsAttributeSettings User Schema Attribute Settings
type UserAttributesSettingsAttributeSettings struct {

	// Fully-qualified attribute or complex mapping Name
	// **SCIM++ Properties:**
	//  - caseExact: false
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: true
	//  - returned: default
	//  - type: string
	//  - uniqueness: none
	Name *string `mandatory:"true" json:"name"`

	// End User mutability
	// **SCIM++ Properties:**
	//  - idcsCanonicalValueSourceFilter: attrName eq "mutabilityValues" and attrValues.value eq "$(endUserMutability)"
	//  - idcsCanonicalValueSourceResourceType: AllowedValue
	//  - caseExact: false
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: true
	//  - returned: default
	//  - type: string
	//  - uniqueness: none
	EndUserMutability *string `mandatory:"true" json:"endUserMutability"`

	// Specifies the list of User mutabilities allowed.
	// **Added In:** 18.3.4
	// **SCIM++ Properties:**
	//  - idcsCanonicalValueSourceFilter: attrName eq "mutabilityValues" and attrValues.value eq "$(endUserMutability)"
	//  - idcsCanonicalValueSourceResourceType: AllowedValue
	//  - caseExact: true
	//  - multiValued: true
	//  - mutability: readWrite
	//  - required: false
	//  - returned: default
	//  - type: string
	//  - uniqueness: none
	EndUserMutabilityCanonicalValues []string `mandatory:"false" json:"endUserMutabilityCanonicalValues"`
}

func (m UserAttributesSettingsAttributeSettings) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m UserAttributesSettingsAttributeSettings) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
