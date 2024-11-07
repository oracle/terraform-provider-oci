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

// ConditionExtensionOciconsolesignonpolicyconsentPolicy This extension defines attributes used to record consent for modification of the "Security Policy for OCI Console" sign-on policy, Rule, Condition or ConditionGroup.
type ConditionExtensionOciconsolesignonpolicyconsentPolicy struct {

	// Set to true when an identity domain administrator opts to change the Oracle security defaults for the "Security Policy for OCI Console" sign-on policy shipped by Oracle. Defaults to false.
	// **Added In:** 2405220110
	// **SCIM++ Properties:**
	//  - idcsSearchable: false
	//  - multiValued: false
	//  - mutability: writeOnly
	//  - required: false
	//  - returned: never
	//  - type: boolean
	Consent *bool `mandatory:"false" json:"consent"`

	// The justification for the change when an identity domain administrator opts to modify the Oracle security defaults for the "Security Policy for OCI Console" sign-on policy shipped by Oracle.
	// **Added In:** 2405220110
	// **SCIM++ Properties:**
	//  - idcsSearchable: false
	//  - multiValued: false
	//  - mutability: writeOnly
	//  - required: false
	//  - returned: never
	//  - type: string
	Justification *string `mandatory:"false" json:"justification"`

	// The detailed reason for the change when an identity domain administrator opts to modify the Oracle security defaults for the "Security Policy for OCI Console" sign-on policy shipped by Oracle.
	// **Added In:** 2405220110
	// **SCIM++ Properties:**
	//  - idcsSearchable: false
	//  - multiValued: false
	//  - mutability: writeOnly
	//  - required: false
	//  - returned: never
	//  - type: string
	Reason *string `mandatory:"false" json:"reason"`
}

func (m ConditionExtensionOciconsolesignonpolicyconsentPolicy) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ConditionExtensionOciconsolesignonpolicyconsentPolicy) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
