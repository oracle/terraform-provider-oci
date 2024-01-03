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

// MyRequestApprovalDetails Approvals created for this request.
type MyRequestApprovalDetails struct {

	// Approver Id
	// **Added In:** 2307071836
	// **SCIM++ Properties:**
	//  - idcsSearchable: false
	//  - multiValued: false
	//  - returned: default
	//  - type: string
	//  - uniqueness: none
	//  - mutability: readOnly
	ApproverId *string `mandatory:"false" json:"approverId"`

	// Approver display name
	// **Added In:** 2307071836
	// **SCIM++ Properties:**
	//  - idcsSearchable: false
	//  - multiValued: false
	//  - returned: default
	//  - type: string
	//  - uniqueness: none
	//  - mutability: readOnly
	ApproverDisplayName *string `mandatory:"false" json:"approverDisplayName"`

	// Approval Justification
	// **Added In:** 2307071836
	// **SCIM++ Properties:**
	//  - multiValued: false
	//  - idcsSearchable: false
	//  - returned: default
	//  - type: string
	//  - uniqueness: none
	//  - mutability: readOnly
	Justification *string `mandatory:"false" json:"justification"`

	// Approval Status
	// **Added In:** 2307071836
	// **SCIM++ Properties:**
	//  - idcsSearchable: false
	//  - multiValued: false
	//  - returned: default
	//  - type: string
	//  - uniqueness: none
	//  - mutability: readOnly
	Status *string `mandatory:"false" json:"status"`

	// Approval Order
	// **Added In:** 2307071836
	// **SCIM++ Properties:**
	//  - idcsSearchable: false
	//  - multiValued: false
	//  - returned: default
	//  - type: integer
	//  - uniqueness: none
	//  - mutability: readOnly
	Order *int `mandatory:"false" json:"order"`

	// Approval Type (Escalation or Regular)
	// **Added In:** 2307071836
	// **SCIM++ Properties:**
	//  - idcsSearchable: false
	//  - multiValued: false
	//  - returned: default
	//  - type: string
	//  - uniqueness: none
	//  - mutability: readOnly
	ApprovalType *string `mandatory:"false" json:"approvalType"`

	// Approval Update Time
	// **Added In:** 2307071836
	// **SCIM++ Properties:**
	//  - multiValued: false
	//  - idcsSearchable: false
	//  - returned: default
	//  - type: dateTime
	//  - uniqueness: none
	//  - mutability: readOnly
	TimeUpdated *string `mandatory:"false" json:"timeUpdated"`
}

func (m MyRequestApprovalDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m MyRequestApprovalDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
