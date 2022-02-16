// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// DevOps API
//
// Use the DevOps API to create DevOps projects, configure code repositories,  add artifacts to deploy, build and test software applications, configure  target deployment environments, and deploy software applications.  For more information, see DevOps (https://docs.cloud.oracle.com/Content/devops/using/home.htm).
//

package devops

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v58/common"
	"strings"
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

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ApproveDeploymentDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingApproveDeploymentDetailsActionEnum(string(m.Action)); !ok && m.Action != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Action: %s. Supported values are: %s.", m.Action, strings.Join(GetApproveDeploymentDetailsActionEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ApproveDeploymentDetailsActionEnum Enum with underlying type: string
type ApproveDeploymentDetailsActionEnum string

// Set of constants representing the allowable values for ApproveDeploymentDetailsActionEnum
const (
	ApproveDeploymentDetailsActionApprove ApproveDeploymentDetailsActionEnum = "APPROVE"
	ApproveDeploymentDetailsActionReject  ApproveDeploymentDetailsActionEnum = "REJECT"
)

var mappingApproveDeploymentDetailsActionEnum = map[string]ApproveDeploymentDetailsActionEnum{
	"APPROVE": ApproveDeploymentDetailsActionApprove,
	"REJECT":  ApproveDeploymentDetailsActionReject,
}

// GetApproveDeploymentDetailsActionEnumValues Enumerates the set of values for ApproveDeploymentDetailsActionEnum
func GetApproveDeploymentDetailsActionEnumValues() []ApproveDeploymentDetailsActionEnum {
	values := make([]ApproveDeploymentDetailsActionEnum, 0)
	for _, v := range mappingApproveDeploymentDetailsActionEnum {
		values = append(values, v)
	}
	return values
}

// GetApproveDeploymentDetailsActionEnumStringValues Enumerates the set of values in String for ApproveDeploymentDetailsActionEnum
func GetApproveDeploymentDetailsActionEnumStringValues() []string {
	return []string{
		"APPROVE",
		"REJECT",
	}
}

// GetMappingApproveDeploymentDetailsActionEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingApproveDeploymentDetailsActionEnum(val string) (ApproveDeploymentDetailsActionEnum, bool) {
	mappingApproveDeploymentDetailsActionEnumIgnoreCase := make(map[string]ApproveDeploymentDetailsActionEnum)
	for k, v := range mappingApproveDeploymentDetailsActionEnum {
		mappingApproveDeploymentDetailsActionEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingApproveDeploymentDetailsActionEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}
