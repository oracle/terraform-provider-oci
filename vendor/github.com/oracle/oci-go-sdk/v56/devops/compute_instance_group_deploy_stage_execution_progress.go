// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// DevOps API
//
// Use the DevOps API to create DevOps projects, configure code repositories,  add artifacts to deploy, build and test software applications, configure  target deployment environments, and deploy software applications.  For more information, see DevOps (https://docs.cloud.oracle.com/Content/devops/using/home.htm).
//

package devops

import (
	"encoding/json"
	"github.com/oracle/oci-go-sdk/v56/common"
)

// ComputeInstanceGroupDeployStageExecutionProgress Specifies the execution details for the instance group rolling deployment stage.
type ComputeInstanceGroupDeployStageExecutionProgress struct {

	// Stage display name. Avoid entering confidential information.
	DeployStageDisplayName *string `mandatory:"false" json:"deployStageDisplayName"`

	// The OCID of the stage.
	DeployStageId *string `mandatory:"false" json:"deployStageId"`

	// Time the stage started executing. Format defined by RFC3339 (https://datatracker.ietf.org/doc/html/rfc3339).
	TimeStarted *common.SDKTime `mandatory:"false" json:"timeStarted"`

	// Time the stage finished executing. Format defined by RFC3339 (https://datatracker.ietf.org/doc/html/rfc3339).
	TimeFinished *common.SDKTime `mandatory:"false" json:"timeFinished"`

	DeployStagePredecessors *DeployStagePredecessorCollection `mandatory:"false" json:"deployStagePredecessors"`

	// Details about stage execution for all the target environments.
	DeployStageExecutionProgressDetails []DeployStageExecutionProgressDetails `mandatory:"false" json:"deployStageExecutionProgressDetails"`

	// The current state of the stage.
	Status DeployStageExecutionProgressStatusEnum `mandatory:"false" json:"status,omitempty"`
}

//GetDeployStageDisplayName returns DeployStageDisplayName
func (m ComputeInstanceGroupDeployStageExecutionProgress) GetDeployStageDisplayName() *string {
	return m.DeployStageDisplayName
}

//GetDeployStageId returns DeployStageId
func (m ComputeInstanceGroupDeployStageExecutionProgress) GetDeployStageId() *string {
	return m.DeployStageId
}

//GetTimeStarted returns TimeStarted
func (m ComputeInstanceGroupDeployStageExecutionProgress) GetTimeStarted() *common.SDKTime {
	return m.TimeStarted
}

//GetTimeFinished returns TimeFinished
func (m ComputeInstanceGroupDeployStageExecutionProgress) GetTimeFinished() *common.SDKTime {
	return m.TimeFinished
}

//GetStatus returns Status
func (m ComputeInstanceGroupDeployStageExecutionProgress) GetStatus() DeployStageExecutionProgressStatusEnum {
	return m.Status
}

//GetDeployStagePredecessors returns DeployStagePredecessors
func (m ComputeInstanceGroupDeployStageExecutionProgress) GetDeployStagePredecessors() *DeployStagePredecessorCollection {
	return m.DeployStagePredecessors
}

//GetDeployStageExecutionProgressDetails returns DeployStageExecutionProgressDetails
func (m ComputeInstanceGroupDeployStageExecutionProgress) GetDeployStageExecutionProgressDetails() []DeployStageExecutionProgressDetails {
	return m.DeployStageExecutionProgressDetails
}

func (m ComputeInstanceGroupDeployStageExecutionProgress) String() string {
	return common.PointerString(m)
}

// MarshalJSON marshals to json representation
func (m ComputeInstanceGroupDeployStageExecutionProgress) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeComputeInstanceGroupDeployStageExecutionProgress ComputeInstanceGroupDeployStageExecutionProgress
	s := struct {
		DiscriminatorParam string `json:"deployStageType"`
		MarshalTypeComputeInstanceGroupDeployStageExecutionProgress
	}{
		"COMPUTE_INSTANCE_GROUP_ROLLING_DEPLOYMENT",
		(MarshalTypeComputeInstanceGroupDeployStageExecutionProgress)(m),
	}

	return json.Marshal(&s)
}
