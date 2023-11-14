// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
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

// ApprovalWorkflowAssignmentAssignedTo Details of resource for which Approval Workflow is assigned
// **SCIM++ Properties:**
//   - caseExact: true
//   - idcsSearchable: true
//   - multiValued: false
//   - mutability: readWrite
//   - required: true
//   - returned: default
//   - type: complex
//   - uniqueness: none
type ApprovalWorkflowAssignmentAssignedTo struct {

	// Identifier of the resource for which Approval Workflow is assigned
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

	// Type of the resource (stripe and non-stripe) for which Approval Workflow is assigned
	// **SCIM++ Properties:**
	//  - caseExact: true
	//  - idcsSearchable: true
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: true
	//  - returned: always
	//  - type: string
	//  - uniqueness: none
	Type ApprovalWorkflowAssignmentAssignedToTypeEnum `mandatory:"true" json:"type"`

	// Ocid of the resource for which Approval Workflow is assigned
	// **SCIM++ Properties:**
	//  - caseExact: true
	//  - idcsSearchable: true
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: false
	//  - returned: always
	//  - type: string
	//  - uniqueness: none
	Ocid *string `mandatory:"false" json:"ocid"`

	// Display name of the resource for which Approval Workflow is assigned
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

	// Description
	// **SCIM++ Properties:**
	//  - caseExact: true
	//  - idcsSearchable: false
	//  - multiValued: false
	//  - mutability: readOnly
	//  - required: false
	//  - returned: request
	//  - type: string
	//  - uniqueness: none
	Description *string `mandatory:"false" json:"description"`
}

func (m ApprovalWorkflowAssignmentAssignedTo) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ApprovalWorkflowAssignmentAssignedTo) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingApprovalWorkflowAssignmentAssignedToTypeEnum(string(m.Type)); !ok && m.Type != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Type: %s. Supported values are: %s.", m.Type, strings.Join(GetApprovalWorkflowAssignmentAssignedToTypeEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ApprovalWorkflowAssignmentAssignedToTypeEnum Enum with underlying type: string
type ApprovalWorkflowAssignmentAssignedToTypeEnum string

// Set of constants representing the allowable values for ApprovalWorkflowAssignmentAssignedToTypeEnum
const (
	ApprovalWorkflowAssignmentAssignedToTypeGroup ApprovalWorkflowAssignmentAssignedToTypeEnum = "Group"
)

var mappingApprovalWorkflowAssignmentAssignedToTypeEnum = map[string]ApprovalWorkflowAssignmentAssignedToTypeEnum{
	"Group": ApprovalWorkflowAssignmentAssignedToTypeGroup,
}

var mappingApprovalWorkflowAssignmentAssignedToTypeEnumLowerCase = map[string]ApprovalWorkflowAssignmentAssignedToTypeEnum{
	"group": ApprovalWorkflowAssignmentAssignedToTypeGroup,
}

// GetApprovalWorkflowAssignmentAssignedToTypeEnumValues Enumerates the set of values for ApprovalWorkflowAssignmentAssignedToTypeEnum
func GetApprovalWorkflowAssignmentAssignedToTypeEnumValues() []ApprovalWorkflowAssignmentAssignedToTypeEnum {
	values := make([]ApprovalWorkflowAssignmentAssignedToTypeEnum, 0)
	for _, v := range mappingApprovalWorkflowAssignmentAssignedToTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetApprovalWorkflowAssignmentAssignedToTypeEnumStringValues Enumerates the set of values in String for ApprovalWorkflowAssignmentAssignedToTypeEnum
func GetApprovalWorkflowAssignmentAssignedToTypeEnumStringValues() []string {
	return []string{
		"Group",
	}
}

// GetMappingApprovalWorkflowAssignmentAssignedToTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingApprovalWorkflowAssignmentAssignedToTypeEnum(val string) (ApprovalWorkflowAssignmentAssignedToTypeEnum, bool) {
	enum, ok := mappingApprovalWorkflowAssignmentAssignedToTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
