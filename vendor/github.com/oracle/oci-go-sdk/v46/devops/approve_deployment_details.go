// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// DevOps API
//
// Use the DevOps APIs to create a DevOps project to group the pipelines,  add reference to target deployment environments, add artifacts to deploy,  and create deployment pipelines needed to deploy your software.
//

package devops

import (
	"github.com/oracle/oci-go-sdk/v46/common"
)

// ApproveDeploymentDetails The stage information for submitting for approval.
type ApproveDeploymentDetails struct {

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the stage which is marked for approval.
	DeployStageId *string `mandatory:"true" json:"deployStageId"`

	// The action of Approve or Reject.
	Action ApproveDeploymentDetailsActionEnum `mandatory:"true" json:"action"`

	// The reason for approving or rejecting the deployment.
	Reason *string `mandatory:"false" json:"reason"`
}

func (m ApproveDeploymentDetails) String() string {
	return common.PointerString(m)
}

// ApproveDeploymentDetailsActionEnum Enum with underlying type: string
type ApproveDeploymentDetailsActionEnum string

// Set of constants representing the allowable values for ApproveDeploymentDetailsActionEnum
const (
	ApproveDeploymentDetailsActionApprove ApproveDeploymentDetailsActionEnum = "APPROVE"
	ApproveDeploymentDetailsActionReject  ApproveDeploymentDetailsActionEnum = "REJECT"
)

var mappingApproveDeploymentDetailsAction = map[string]ApproveDeploymentDetailsActionEnum{
	"APPROVE": ApproveDeploymentDetailsActionApprove,
	"REJECT":  ApproveDeploymentDetailsActionReject,
}

// GetApproveDeploymentDetailsActionEnumValues Enumerates the set of values for ApproveDeploymentDetailsActionEnum
func GetApproveDeploymentDetailsActionEnumValues() []ApproveDeploymentDetailsActionEnum {
	values := make([]ApproveDeploymentDetailsActionEnum, 0)
	for _, v := range mappingApproveDeploymentDetailsAction {
		values = append(values, v)
	}
	return values
}
