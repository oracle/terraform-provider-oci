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

// AppGroupAssertionAttributes Each value of this attribute describes an attribute of Group that will be sent in a Security Assertion Markup Language (SAML) assertion.
// **Deprecated Since: 18.2.2**
// **SCIM++ Properties:**
//   - caseExact: false
//   - idcsCompositeKey: [name]
//   - idcsSearchable: false
//   - idcsValuePersistedInOtherAttribute: true
//   - multiValued: true
//   - mutability: readWrite
//   - required: false
//   - returned: default
//   - type: complex
//   - uniqueness: none
type AppGroupAssertionAttributes struct {

	// The attribute represents the name of the attribute that will be used in the Security Assertion Markup Language (SAML) assertion
	// **Deprecated Since: 18.2.2**
	// **SCIM++ Properties:**
	//  - caseExact: false
	//  - idcsSearchable: false
	//  - idcsValuePersistedInOtherAttribute: true
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: true
	//  - returned: default
	//  - type: string
	//  - uniqueness: none
	Name *string `mandatory:"true" json:"name"`

	// Indicates the format of the assertion attribute.
	// **Deprecated Since: 18.2.2**
	// **SCIM++ Properties:**
	//  - caseExact: true
	//  - idcsSearchable: false
	//  - idcsValuePersistedInOtherAttribute: true
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: false
	//  - returned: default
	//  - type: string
	//  - uniqueness: none
	Format *string `mandatory:"false" json:"format"`

	// Indicates the filter types that are supported for the Group assertion attributes.
	// **Deprecated Since: 18.2.2**
	// **SCIM++ Properties:**
	//  - caseExact: false
	//  - idcsSearchable: false
	//  - idcsValuePersistedInOtherAttribute: true
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: false
	//  - returned: default
	//  - type: string
	//  - uniqueness: none
	Condition *string `mandatory:"false" json:"condition"`

	// Indicates the group name that are supported for the group assertion attributes.
	// **Deprecated Since: 18.2.2**
	// **SCIM++ Properties:**
	//  - caseExact: true
	//  - idcsSearchable: false
	//  - idcsValuePersistedInOtherAttribute: true
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: false
	//  - returned: default
	//  - type: string
	//  - uniqueness: none
	GroupName *string `mandatory:"false" json:"groupName"`
}

func (m AppGroupAssertionAttributes) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m AppGroupAssertionAttributes) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
