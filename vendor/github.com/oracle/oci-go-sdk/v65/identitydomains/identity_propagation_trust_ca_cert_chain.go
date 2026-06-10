// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
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

// IdentityPropagationTrustCaCertChain Certificate trust store. This is required if identity propagation type is X509.
// **Added In:** 2508041610
// **SCIM++ Properties:**
//   - caseExact: true
//   - type: complex
//   - multiValued: false
//   - required: false
//   - mutability: readWrite
//   - returned: default
//   - uniqueness: none
type IdentityPropagationTrustCaCertChain struct {

	// A list of PEM-encoded root CA certificates.
	// **Added In:** 2508041610
	// **SCIM++ Properties:**
	//  - caseExact: true
	//  - type: string
	//  - multiValued: true
	//  - required: true
	//  - mutability: readWrite
	//  - returned: default
	//  - uniqueness: none
	RootCAs []string `mandatory:"true" json:"rootCAs"`

	// A list of PEM-encoded intermediate CA certificates.
	// **Added In:** 2508041610
	// **SCIM++ Properties:**
	//  - caseExact: true
	//  - type: string
	//  - multiValued: true
	//  - required: false
	//  - mutability: readWrite
	//  - returned: default
	//  - uniqueness: none
	IntermediateCAs []string `mandatory:"false" json:"intermediateCAs"`
}

func (m IdentityPropagationTrustCaCertChain) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m IdentityPropagationTrustCaCertChain) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}
