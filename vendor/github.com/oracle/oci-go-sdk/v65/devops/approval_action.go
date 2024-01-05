// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// DevOps API
//
// Use the DevOps API to create DevOps projects, configure code repositories,  add artifacts to deploy, build and test software applications, configure  target deployment environments, and deploy software applications.  For more information, see DevOps (https://docs.cloud.oracle.com/Content/devops/using/home.htm).
//

package devops

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// ApprovalAction Information about the approval action of DevOps deployment stages.
type ApprovalAction struct {

	// The subject ID of the user who approves or disapproves a DevOps deployment stage.
	SubjectId *string `mandatory:"true" json:"subjectId"`

	// The action of the user on the DevOps deployment stage.
	Action ApprovalActionActionEnum `mandatory:"true" json:"action"`

	// The reason for approving or rejecting the deployment.
	Reason *string `mandatory:"false" json:"reason"`
}

func (m ApprovalAction) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ApprovalAction) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingApprovalActionActionEnum(string(m.Action)); !ok && m.Action != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Action: %s. Supported values are: %s.", m.Action, strings.Join(GetApprovalActionActionEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ApprovalActionActionEnum Enum with underlying type: string
type ApprovalActionActionEnum string

// Set of constants representing the allowable values for ApprovalActionActionEnum
const (
	ApprovalActionActionApprove ApprovalActionActionEnum = "APPROVE"
	ApprovalActionActionReject  ApprovalActionActionEnum = "REJECT"
)

var mappingApprovalActionActionEnum = map[string]ApprovalActionActionEnum{
	"APPROVE": ApprovalActionActionApprove,
	"REJECT":  ApprovalActionActionReject,
}

var mappingApprovalActionActionEnumLowerCase = map[string]ApprovalActionActionEnum{
	"approve": ApprovalActionActionApprove,
	"reject":  ApprovalActionActionReject,
}

// GetApprovalActionActionEnumValues Enumerates the set of values for ApprovalActionActionEnum
func GetApprovalActionActionEnumValues() []ApprovalActionActionEnum {
	values := make([]ApprovalActionActionEnum, 0)
	for _, v := range mappingApprovalActionActionEnum {
		values = append(values, v)
	}
	return values
}

// GetApprovalActionActionEnumStringValues Enumerates the set of values in String for ApprovalActionActionEnum
func GetApprovalActionActionEnumStringValues() []string {
	return []string{
		"APPROVE",
		"REJECT",
	}
}

// GetMappingApprovalActionActionEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingApprovalActionActionEnum(val string) (ApprovalActionActionEnum, bool) {
	enum, ok := mappingApprovalActionActionEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
