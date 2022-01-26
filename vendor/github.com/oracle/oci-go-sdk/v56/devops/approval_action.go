// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// DevOps API
//
// Use the DevOps API to create DevOps projects, configure code repositories,  add artifacts to deploy, build and test software applications, configure  target deployment environments, and deploy software applications.  For more information, see DevOps (https://docs.cloud.oracle.com/Content/devops/using/home.htm).
//

package devops

import (
	"github.com/oracle/oci-go-sdk/v56/common"
)

// ApprovalAction Information about the approval action of DevOps deployment stages.
type ApprovalAction struct {

	// The subject ID of the user who approves or disapproves a DevOps deployment stage.
	SubjectId *string `mandatory:"true" json:"subjectId"`

	// The action of the user on the DevOps deployment stage.
	Action ApprovalActionActionEnum `mandatory:"true" json:"action"`
}

func (m ApprovalAction) String() string {
	return common.PointerString(m)
}

// ApprovalActionActionEnum Enum with underlying type: string
type ApprovalActionActionEnum string

// Set of constants representing the allowable values for ApprovalActionActionEnum
const (
	ApprovalActionActionApprove ApprovalActionActionEnum = "APPROVE"
	ApprovalActionActionReject  ApprovalActionActionEnum = "REJECT"
)

var mappingApprovalActionAction = map[string]ApprovalActionActionEnum{
	"APPROVE": ApprovalActionActionApprove,
	"REJECT":  ApprovalActionActionReject,
}

// GetApprovalActionActionEnumValues Enumerates the set of values for ApprovalActionActionEnum
func GetApprovalActionActionEnumValues() []ApprovalActionActionEnum {
	values := make([]ApprovalActionActionEnum, 0)
	for _, v := range mappingApprovalActionAction {
		values = append(values, v)
	}
	return values
}
