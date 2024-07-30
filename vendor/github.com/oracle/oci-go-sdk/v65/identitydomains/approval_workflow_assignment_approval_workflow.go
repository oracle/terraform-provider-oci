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

// ApprovalWorkflowAssignmentApprovalWorkflow Details of the Approval Workflow
// **SCIM++ Properties:**
//   - caseExact: true
//   - idcsSearchable: true
//   - multiValued: false
//   - mutability: readWrite
//   - required: true
//   - returned: default
//   - type: complex
//   - uniqueness: none
type ApprovalWorkflowAssignmentApprovalWorkflow struct {

	// Identifier of the approval workflow
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

	// Indicates type of the entity that is associated with this assignment (for ARM validation)
	// **SCIM++ Properties:**
	//  - caseExact: true
	//  - idcsSearchable: true
	//  - idcsDefaultValue: ApprovalWorkflow
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: true
	//  - returned: request
	//  - type: string
	//  - uniqueness: none
	Type ApprovalWorkflowAssignmentApprovalWorkflowTypeEnum `mandatory:"true" json:"type"`

	// Unique OCI Identifier of the approval workflow
	// **SCIM++ Properties:**
	//  - caseExact: true
	//  - idcsSearchable: true
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: false
	//  - returned: default
	//  - type: string
	//  - uniqueness: none
	Ocid *string `mandatory:"false" json:"ocid"`

	// Display name of the approval workflow
	// **SCIM++ Properties:**
	//  - caseExact: true
	//  - idcsSearchable: true
	//  - multiValued: false
	//  - mutability: readOnly
	//  - required: false
	//  - returned: request
	//  - type: string
	//  - uniqueness: none
	Display *string `mandatory:"false" json:"display"`

	// URI of the approval workflow
	// **SCIM++ Properties:**
	//  - idcsSearchable: false
	//  - multiValued: false
	//  - mutability: readOnly
	//  - required: false
	//  - returned: request
	//  - type: reference
	//  - uniqueness: none
	Ref *string `mandatory:"false" json:"$ref"`
}

func (m ApprovalWorkflowAssignmentApprovalWorkflow) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ApprovalWorkflowAssignmentApprovalWorkflow) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingApprovalWorkflowAssignmentApprovalWorkflowTypeEnum(string(m.Type)); !ok && m.Type != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Type: %s. Supported values are: %s.", m.Type, strings.Join(GetApprovalWorkflowAssignmentApprovalWorkflowTypeEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ApprovalWorkflowAssignmentApprovalWorkflowTypeEnum Enum with underlying type: string
type ApprovalWorkflowAssignmentApprovalWorkflowTypeEnum string

// Set of constants representing the allowable values for ApprovalWorkflowAssignmentApprovalWorkflowTypeEnum
const (
	ApprovalWorkflowAssignmentApprovalWorkflowTypeApprovalworkflow ApprovalWorkflowAssignmentApprovalWorkflowTypeEnum = "ApprovalWorkflow"
)

var mappingApprovalWorkflowAssignmentApprovalWorkflowTypeEnum = map[string]ApprovalWorkflowAssignmentApprovalWorkflowTypeEnum{
	"ApprovalWorkflow": ApprovalWorkflowAssignmentApprovalWorkflowTypeApprovalworkflow,
}

var mappingApprovalWorkflowAssignmentApprovalWorkflowTypeEnumLowerCase = map[string]ApprovalWorkflowAssignmentApprovalWorkflowTypeEnum{
	"approvalworkflow": ApprovalWorkflowAssignmentApprovalWorkflowTypeApprovalworkflow,
}

// GetApprovalWorkflowAssignmentApprovalWorkflowTypeEnumValues Enumerates the set of values for ApprovalWorkflowAssignmentApprovalWorkflowTypeEnum
func GetApprovalWorkflowAssignmentApprovalWorkflowTypeEnumValues() []ApprovalWorkflowAssignmentApprovalWorkflowTypeEnum {
	values := make([]ApprovalWorkflowAssignmentApprovalWorkflowTypeEnum, 0)
	for _, v := range mappingApprovalWorkflowAssignmentApprovalWorkflowTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetApprovalWorkflowAssignmentApprovalWorkflowTypeEnumStringValues Enumerates the set of values in String for ApprovalWorkflowAssignmentApprovalWorkflowTypeEnum
func GetApprovalWorkflowAssignmentApprovalWorkflowTypeEnumStringValues() []string {
	return []string{
		"ApprovalWorkflow",
	}
}

// GetMappingApprovalWorkflowAssignmentApprovalWorkflowTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingApprovalWorkflowAssignmentApprovalWorkflowTypeEnum(val string) (ApprovalWorkflowAssignmentApprovalWorkflowTypeEnum, bool) {
	enum, ok := mappingApprovalWorkflowAssignmentApprovalWorkflowTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
