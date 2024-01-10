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

// ApprovalWorkflowMaxDuration Max duration of the ApprovalWorkflow must be acted at all levels.
// **SCIM++ Properties:**
//   - caseExact: true
//   - idcsSearchable: false
//   - multiValued: false
//   - mutability: readWrite
//   - required: true
//   - returned: default
//   - type: complex
//   - uniqueness: none
type ApprovalWorkflowMaxDuration struct {

	// The value of the max duration.
	// **SCIM++ Properties:**
	//  - caseExact: true
	//  - idcsSearchable: false
	//  - idcsMaxValue: 1488
	//  - idcsMinValue: 1
	//  - idcsDefaultValue: 14
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: true
	//  - returned: default
	//  - type: integer
	//  - uniqueness: none
	Value *int `mandatory:"true" json:"value"`

	// The unit of the max duration.
	// **SCIM++ Properties:**
	//  - caseExact: true
	//  - idcsSearchable: false
	//  - idcsDefaultValue: DAY
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: true
	//  - returned: default
	//  - type: string
	//  - uniqueness: none
	Unit ApprovalWorkflowMaxDurationUnitEnum `mandatory:"true" json:"unit"`
}

func (m ApprovalWorkflowMaxDuration) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ApprovalWorkflowMaxDuration) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingApprovalWorkflowMaxDurationUnitEnum(string(m.Unit)); !ok && m.Unit != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Unit: %s. Supported values are: %s.", m.Unit, strings.Join(GetApprovalWorkflowMaxDurationUnitEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ApprovalWorkflowMaxDurationUnitEnum Enum with underlying type: string
type ApprovalWorkflowMaxDurationUnitEnum string

// Set of constants representing the allowable values for ApprovalWorkflowMaxDurationUnitEnum
const (
	ApprovalWorkflowMaxDurationUnitMonth ApprovalWorkflowMaxDurationUnitEnum = "MONTH"
	ApprovalWorkflowMaxDurationUnitWeek  ApprovalWorkflowMaxDurationUnitEnum = "WEEK"
	ApprovalWorkflowMaxDurationUnitDay   ApprovalWorkflowMaxDurationUnitEnum = "DAY"
	ApprovalWorkflowMaxDurationUnitHour  ApprovalWorkflowMaxDurationUnitEnum = "HOUR"
)

var mappingApprovalWorkflowMaxDurationUnitEnum = map[string]ApprovalWorkflowMaxDurationUnitEnum{
	"MONTH": ApprovalWorkflowMaxDurationUnitMonth,
	"WEEK":  ApprovalWorkflowMaxDurationUnitWeek,
	"DAY":   ApprovalWorkflowMaxDurationUnitDay,
	"HOUR":  ApprovalWorkflowMaxDurationUnitHour,
}

var mappingApprovalWorkflowMaxDurationUnitEnumLowerCase = map[string]ApprovalWorkflowMaxDurationUnitEnum{
	"month": ApprovalWorkflowMaxDurationUnitMonth,
	"week":  ApprovalWorkflowMaxDurationUnitWeek,
	"day":   ApprovalWorkflowMaxDurationUnitDay,
	"hour":  ApprovalWorkflowMaxDurationUnitHour,
}

// GetApprovalWorkflowMaxDurationUnitEnumValues Enumerates the set of values for ApprovalWorkflowMaxDurationUnitEnum
func GetApprovalWorkflowMaxDurationUnitEnumValues() []ApprovalWorkflowMaxDurationUnitEnum {
	values := make([]ApprovalWorkflowMaxDurationUnitEnum, 0)
	for _, v := range mappingApprovalWorkflowMaxDurationUnitEnum {
		values = append(values, v)
	}
	return values
}

// GetApprovalWorkflowMaxDurationUnitEnumStringValues Enumerates the set of values in String for ApprovalWorkflowMaxDurationUnitEnum
func GetApprovalWorkflowMaxDurationUnitEnumStringValues() []string {
	return []string{
		"MONTH",
		"WEEK",
		"DAY",
		"HOUR",
	}
}

// GetMappingApprovalWorkflowMaxDurationUnitEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingApprovalWorkflowMaxDurationUnitEnum(val string) (ApprovalWorkflowMaxDurationUnitEnum, bool) {
	enum, ok := mappingApprovalWorkflowMaxDurationUnitEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
