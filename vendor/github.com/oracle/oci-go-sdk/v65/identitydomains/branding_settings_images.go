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

// BrandingSettingsImages References to various images
type BrandingSettingsImages struct {

	// Image URI
	// **SCIM++ Properties:**
	//  - multiValued: false
	//  - mutability: readOnly
	//  - required: true
	//  - returned: default
	//  - type: reference
	Value *string `mandatory:"true" json:"value"`

	// Indicates the image type
	// **SCIM++ Properties:**
	//  - multiValued: false
	//  - mutability: readOnly
	//  - required: true
	//  - returned: default
	//  - type: string
	Type *string `mandatory:"true" json:"type"`

	// A human-readable name, primarily used for display purposes
	// **SCIM++ Properties:**
	//  - multiValued: false
	//  - mutability: readOnly
	//  - required: false
	//  - returned: default
	//  - type: string
	Display *string `mandatory:"false" json:"display"`
}

func (m BrandingSettingsImages) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m BrandingSettingsImages) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
