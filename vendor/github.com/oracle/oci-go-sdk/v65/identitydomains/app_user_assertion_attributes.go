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

// AppUserAssertionAttributes Each value of this attribute describes an attribute of User that will be sent in a Security Assertion Markup Language (SAML) assertion.
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
type AppUserAssertionAttributes struct {

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

	// This attribute specifies which user attribute should be used to create the value of the SAML assertion attribute. The userstore attribute can be constructed by using attributes from the Oracle Identity Cloud Service Core Users schema. <br><b>Note</b>: Attributes from extensions to the Core User schema are not supported in v1.0.
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
	UserStoreAttributeName *string `mandatory:"true" json:"userStoreAttributeName"`

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
}

func (m AppUserAssertionAttributes) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m AppUserAssertionAttributes) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
