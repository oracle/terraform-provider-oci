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

// ExtensionPasswordStateUser This extension defines attributes used to manage account passwords within a Service Provider. The extension is typically applied to a User resource, but MAY be applied to other resources that use passwords.
type ExtensionPasswordStateUser struct {

	// A DateTime that specifies the date and time when the current password was set
	// **SCIM++ Properties:**
	//  - idcsSearchable: false
	//  - multiValued: false
	//  - mutability: readOnly
	//  - required: false
	//  - returned: request
	//  - type: dateTime
	//  - uniqueness: none
	LastSuccessfulSetDate *string `mandatory:"false" json:"lastSuccessfulSetDate"`

	// Indicates that the current password MAY NOT be changed and all other password expiry settings SHALL be ignored
	// **SCIM++ Properties:**
	//  - idcsSearchable: false
	//  - multiValued: false
	//  - mutability: readOnly
	//  - required: false
	//  - returned: request
	//  - type: boolean
	//  - uniqueness: none
	CantChange *bool `mandatory:"false" json:"cantChange"`

	// Indicates that the password expiry policy will not be applied for the current Resource
	// **SCIM++ Properties:**
	//  - idcsSearchable: false
	//  - multiValued: false
	//  - mutability: readOnly
	//  - required: false
	//  - returned: request
	//  - type: boolean
	//  - uniqueness: none
	CantExpire *bool `mandatory:"false" json:"cantExpire"`

	// Indicates that the subject password value MUST change on next login. If not changed, typically the account is locked. The value may be set indirectly when the subject's current password expires or directly set by an administrator.
	// **SCIM++ Properties:**
	//  - idcsSearchable: false
	//  - multiValued: false
	//  - mutability: readOnly
	//  - required: false
	//  - returned: request
	//  - type: boolean
	//  - uniqueness: none
	MustChange *bool `mandatory:"false" json:"mustChange"`

	// Indicates that the password has expired
	// **SCIM++ Properties:**
	//  - idcsSearchable: false
	//  - multiValued: false
	//  - mutability: readOnly
	//  - required: false
	//  - returned: request
	//  - type: boolean
	//  - uniqueness: none
	Expired *bool `mandatory:"false" json:"expired"`

	// A DateTime that specifies the date and time when last successful password validation was set
	// **Added In:** 2011192329
	// **SCIM++ Properties:**
	//  - idcsSearchable: true
	//  - multiValued: false
	//  - mutability: readOnly
	//  - required: false
	//  - returned: request
	//  - type: dateTime
	//  - uniqueness: none
	LastSuccessfulValidationDate *string `mandatory:"false" json:"lastSuccessfulValidationDate"`

	// A DateTime that specifies the date and time when last failed password validation was set
	// **Added In:** 2011192329
	// **SCIM++ Properties:**
	//  - idcsSearchable: true
	//  - multiValued: false
	//  - mutability: readOnly
	//  - required: false
	//  - returned: request
	//  - type: dateTime
	//  - uniqueness: none
	LastFailedValidationDate *string `mandatory:"false" json:"lastFailedValidationDate"`

	ApplicablePasswordPolicy *UserExtApplicablePasswordPolicy `mandatory:"false" json:"applicablePasswordPolicy"`
}

func (m ExtensionPasswordStateUser) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ExtensionPasswordStateUser) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
