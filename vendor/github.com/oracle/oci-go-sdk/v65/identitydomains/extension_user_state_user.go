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

// ExtensionUserStateUser This extension defines the attributes used to manage account passwords within a service provider. The extension is typically applied to a User resource, but MAY be applied to other resources that use passwords.
type ExtensionUserStateUser struct {

	// The last successful login date.
	// **SCIM++ Properties:**
	//  - idcsSearchable: true
	//  - idcsAllowUpdatesInReadOnlyMode: true
	//  - multiValued: false
	//  - mutability: readOnly
	//  - idcsRequiresWriteForAccessFlows: true
	//  - required: false
	//  - returned: request
	//  - type: dateTime
	//  - uniqueness: none
	LastSuccessfulLoginDate *string `mandatory:"false" json:"lastSuccessfulLoginDate"`

	// The previous successful login date.
	// **SCIM++ Properties:**
	//  - caseExact: false
	//  - idcsSearchable: false
	//  - multiValued: false
	//  - mutability: readOnly
	//  - idcsRequiresWriteForAccessFlows: true
	//  - required: false
	//  - returned: request
	//  - type: dateTime
	//  - uniqueness: none
	PreviousSuccessfulLoginDate *string `mandatory:"false" json:"previousSuccessfulLoginDate"`

	// The last failed login date.
	// **SCIM++ Properties:**
	//  - idcsSearchable: false
	//  - idcsAllowUpdatesInReadOnlyMode: true
	//  - multiValued: false
	//  - mutability: readOnly
	//  - idcsRequiresWriteForAccessFlows: true
	//  - required: false
	//  - returned: request
	//  - type: dateTime
	//  - uniqueness: none
	LastFailedLoginDate *string `mandatory:"false" json:"lastFailedLoginDate"`

	// The number of failed login attempts. The value is reset to 0 after a successful login.
	// **SCIM++ Properties:**
	//  - idcsSearchable: false
	//  - idcsAllowUpdatesInReadOnlyMode: true
	//  - multiValued: false
	//  - mutability: readOnly
	//  - idcsRequiresWriteForAccessFlows: true
	//  - idcsRequiresImmediateReadAfterWriteForAccessFlows: true
	//  - required: false
	//  - returned: request
	//  - type: integer
	//  - uniqueness: none
	LoginAttempts *int `mandatory:"false" json:"loginAttempts"`

	// The number of failed recovery attempts. The value is reset to 0 after a successful login.
	// **Added In:** 19.1.4
	// **SCIM++ Properties:**
	//  - idcsSearchable: false
	//  - multiValued: false
	//  - mutability: readOnly
	//  - idcsRequiresWriteForAccessFlows: true
	//  - required: false
	//  - returned: request
	//  - type: integer
	//  - uniqueness: none
	RecoveryAttempts *int `mandatory:"false" json:"recoveryAttempts"`

	// The number of failed account recovery enrollment attempts.
	// **Added In:** 19.1.4
	// **SCIM++ Properties:**
	//  - idcsSearchable: false
	//  - multiValued: false
	//  - mutability: readOnly
	//  - required: false
	//  - returned: request
	//  - type: integer
	//  - uniqueness: none
	RecoveryEnrollAttempts *int `mandatory:"false" json:"recoveryEnrollAttempts"`

	// The maximum number of concurrent sessions for a user.
	// **Added In:** 20.1.3
	// **SCIM++ Properties:**
	//  - caseExact: false
	//  - idcsMaxValue: 999
	//  - idcsMinValue: 1
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: false
	//  - returned: default
	//  - type: integer
	//  - uniqueness: none
	MaxConcurrentSessions *int `mandatory:"false" json:"maxConcurrentSessions"`

	RecoveryLocked *UserExtRecoveryLocked `mandatory:"false" json:"recoveryLocked"`

	Locked *UserExtLocked `mandatory:"false" json:"locked"`
}

func (m ExtensionUserStateUser) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ExtensionUserStateUser) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
