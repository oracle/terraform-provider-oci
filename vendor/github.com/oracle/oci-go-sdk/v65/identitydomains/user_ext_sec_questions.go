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

// UserExtSecQuestions The schema used to mnage security question and answers provided by a user for account recovery and/or MFA. While setting up security questions, a user can also provide a hint for the answer.
// **SCIM++ Properties:**
//   - idcsCompositeKey: [value]
//   - multiValued: true
//   - mutability: readWrite
//   - required: false
//   - returned: request
//   - type: complex
//   - uniqueness: none
type UserExtSecQuestions struct {

	// The identifier of the question selected by the user when setting up a security question.
	// **SCIM++ Properties:**
	//  - caseExact: true
	//  - idcsSearchable: true
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: true
	//  - returned: always
	//  - type: string
	//  - uniqueness: none
	Value *string `mandatory:"true" json:"value"`

	// The answer provided by a user for a security question.
	// **SCIM++ Properties:**
	//  - idcsCsvAttributeName: Answer
	//  - idcsSearchable: false
	//  - idcsSensitive: hash
	//  - multiValued: false
	//  - mutability: writeOnly
	//  - required: true
	//  - returned: never
	//  - type: string
	//  - uniqueness: none
	//  - idcsPii: true
	Answer *string `mandatory:"true" json:"answer"`

	// The URI of the corresponding Security Question resource.
	// **SCIM++ Properties:**
	//  - idcsSearchable: false
	//  - multiValued: false
	//  - mutability: readOnly
	//  - required: false
	//  - returned: default
	//  - type: reference
	//  - uniqueness: none
	Ref *string `mandatory:"false" json:"$ref"`

	// The hint for an answer that's given by user when setting up a security question.
	// **SCIM++ Properties:**
	//  - caseExact: false
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: false
	//  - returned: default
	//  - type: string
	//  - uniqueness: none
	HintText *string `mandatory:"false" json:"hintText"`
}

func (m UserExtSecQuestions) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m UserExtSecQuestions) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
