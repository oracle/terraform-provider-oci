// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// DevOps API
//
// Use the DevOps API to create DevOps projects, configure code repositories,  add artifacts to deploy, build and test software applications, configure  target deployment environments, and deploy software applications.  For more information, see DevOps (https://docs.cloud.oracle.com/Content/devops/using/home.htm).
//

package devops

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// ApprovalRuleMergeCheck The status of the approval rules.
type ApprovalRuleMergeCheck struct {

	// The name of the rule.
	RuleName *string `mandatory:"false" json:"ruleName"`

	// The number of total approvals needed.
	TotalApprovalCount *int `mandatory:"false" json:"totalApprovalCount"`

	// The current number of approvals.
	CurrentApprovalCount *int `mandatory:"false" json:"currentApprovalCount"`

	// The list of default reviewers.
	Reviewers []PrincipalDetails `mandatory:"false" json:"reviewers"`

	// The status of the approval rule.
	Status ApprovalRuleMergeCheckStatusEnum `mandatory:"false" json:"status,omitempty"`

	// The level of rule.
	Level ApprovalRuleMergeCheckLevelEnum `mandatory:"false" json:"level,omitempty"`
}

func (m ApprovalRuleMergeCheck) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ApprovalRuleMergeCheck) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingApprovalRuleMergeCheckStatusEnum(string(m.Status)); !ok && m.Status != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Status: %s. Supported values are: %s.", m.Status, strings.Join(GetApprovalRuleMergeCheckStatusEnumStringValues(), ",")))
	}
	if _, ok := GetMappingApprovalRuleMergeCheckLevelEnum(string(m.Level)); !ok && m.Level != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Level: %s. Supported values are: %s.", m.Level, strings.Join(GetApprovalRuleMergeCheckLevelEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m ApprovalRuleMergeCheck) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeApprovalRuleMergeCheck ApprovalRuleMergeCheck
	s := struct {
		DiscriminatorParam string `json:"type"`
		MarshalTypeApprovalRuleMergeCheck
	}{
		"APPROVAL_RULE",
		(MarshalTypeApprovalRuleMergeCheck)(m),
	}

	return json.Marshal(&s)
}

// ApprovalRuleMergeCheckStatusEnum Enum with underlying type: string
type ApprovalRuleMergeCheckStatusEnum string

// Set of constants representing the allowable values for ApprovalRuleMergeCheckStatusEnum
const (
	ApprovalRuleMergeCheckStatusNeedsApproval ApprovalRuleMergeCheckStatusEnum = "NEEDS_APPROVAL"
	ApprovalRuleMergeCheckStatusSucceeded     ApprovalRuleMergeCheckStatusEnum = "SUCCEEDED"
)

var mappingApprovalRuleMergeCheckStatusEnum = map[string]ApprovalRuleMergeCheckStatusEnum{
	"NEEDS_APPROVAL": ApprovalRuleMergeCheckStatusNeedsApproval,
	"SUCCEEDED":      ApprovalRuleMergeCheckStatusSucceeded,
}

var mappingApprovalRuleMergeCheckStatusEnumLowerCase = map[string]ApprovalRuleMergeCheckStatusEnum{
	"needs_approval": ApprovalRuleMergeCheckStatusNeedsApproval,
	"succeeded":      ApprovalRuleMergeCheckStatusSucceeded,
}

// GetApprovalRuleMergeCheckStatusEnumValues Enumerates the set of values for ApprovalRuleMergeCheckStatusEnum
func GetApprovalRuleMergeCheckStatusEnumValues() []ApprovalRuleMergeCheckStatusEnum {
	values := make([]ApprovalRuleMergeCheckStatusEnum, 0)
	for _, v := range mappingApprovalRuleMergeCheckStatusEnum {
		values = append(values, v)
	}
	return values
}

// GetApprovalRuleMergeCheckStatusEnumStringValues Enumerates the set of values in String for ApprovalRuleMergeCheckStatusEnum
func GetApprovalRuleMergeCheckStatusEnumStringValues() []string {
	return []string{
		"NEEDS_APPROVAL",
		"SUCCEEDED",
	}
}

// GetMappingApprovalRuleMergeCheckStatusEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingApprovalRuleMergeCheckStatusEnum(val string) (ApprovalRuleMergeCheckStatusEnum, bool) {
	enum, ok := mappingApprovalRuleMergeCheckStatusEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ApprovalRuleMergeCheckLevelEnum Enum with underlying type: string
type ApprovalRuleMergeCheckLevelEnum string

// Set of constants representing the allowable values for ApprovalRuleMergeCheckLevelEnum
const (
	ApprovalRuleMergeCheckLevelProject    ApprovalRuleMergeCheckLevelEnum = "PROJECT"
	ApprovalRuleMergeCheckLevelRepository ApprovalRuleMergeCheckLevelEnum = "REPOSITORY"
)

var mappingApprovalRuleMergeCheckLevelEnum = map[string]ApprovalRuleMergeCheckLevelEnum{
	"PROJECT":    ApprovalRuleMergeCheckLevelProject,
	"REPOSITORY": ApprovalRuleMergeCheckLevelRepository,
}

var mappingApprovalRuleMergeCheckLevelEnumLowerCase = map[string]ApprovalRuleMergeCheckLevelEnum{
	"project":    ApprovalRuleMergeCheckLevelProject,
	"repository": ApprovalRuleMergeCheckLevelRepository,
}

// GetApprovalRuleMergeCheckLevelEnumValues Enumerates the set of values for ApprovalRuleMergeCheckLevelEnum
func GetApprovalRuleMergeCheckLevelEnumValues() []ApprovalRuleMergeCheckLevelEnum {
	values := make([]ApprovalRuleMergeCheckLevelEnum, 0)
	for _, v := range mappingApprovalRuleMergeCheckLevelEnum {
		values = append(values, v)
	}
	return values
}

// GetApprovalRuleMergeCheckLevelEnumStringValues Enumerates the set of values in String for ApprovalRuleMergeCheckLevelEnum
func GetApprovalRuleMergeCheckLevelEnumStringValues() []string {
	return []string{
		"PROJECT",
		"REPOSITORY",
	}
}

// GetMappingApprovalRuleMergeCheckLevelEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingApprovalRuleMergeCheckLevelEnum(val string) (ApprovalRuleMergeCheckLevelEnum, bool) {
	enum, ok := mappingApprovalRuleMergeCheckLevelEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
