// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// DevOps API
//
// Use the DevOps APIs to create a DevOps project to group the pipelines,  add reference to target deployment environments, add artifacts to deploy,  and create deployment pipelines needed to deploy your software.
//

package devops

import (
	"encoding/json"
	"github.com/oracle/oci-go-sdk/v43/common"
)

// ComputeInstanceGroupCanaryDeployStageExecutionProgress Specifies the Instance Group Canary deployment stage.
type ComputeInstanceGroupCanaryDeployStageExecutionProgress struct {

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
func (m ComputeInstanceGroupCanaryDeployStageExecutionProgress) GetDeployStageDisplayName() *string {
	return m.DeployStageDisplayName
}

//GetDeployStageId returns DeployStageId
func (m ComputeInstanceGroupCanaryDeployStageExecutionProgress) GetDeployStageId() *string {
	return m.DeployStageId
}

//GetTimeStarted returns TimeStarted
func (m ComputeInstanceGroupCanaryDeployStageExecutionProgress) GetTimeStarted() *common.SDKTime {
	return m.TimeStarted
}

//GetTimeFinished returns TimeFinished
func (m ComputeInstanceGroupCanaryDeployStageExecutionProgress) GetTimeFinished() *common.SDKTime {
	return m.TimeFinished
}

//GetStatus returns Status
func (m ComputeInstanceGroupCanaryDeployStageExecutionProgress) GetStatus() DeployStageExecutionProgressStatusEnum {
	return m.Status
}

//GetDeployStagePredecessors returns DeployStagePredecessors
func (m ComputeInstanceGroupCanaryDeployStageExecutionProgress) GetDeployStagePredecessors() *DeployStagePredecessorCollection {
	return m.DeployStagePredecessors
}

//GetDeployStageExecutionProgressDetails returns DeployStageExecutionProgressDetails
func (m ComputeInstanceGroupCanaryDeployStageExecutionProgress) GetDeployStageExecutionProgressDetails() []DeployStageExecutionProgressDetails {
	return m.DeployStageExecutionProgressDetails
}

func (m ComputeInstanceGroupCanaryDeployStageExecutionProgress) String() string {
	return common.PointerString(m)
}

// MarshalJSON marshals to json representation
func (m ComputeInstanceGroupCanaryDeployStageExecutionProgress) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeComputeInstanceGroupCanaryDeployStageExecutionProgress ComputeInstanceGroupCanaryDeployStageExecutionProgress
	s := struct {
		DiscriminatorParam string `json:"deployStageType"`
		MarshalTypeComputeInstanceGroupCanaryDeployStageExecutionProgress
	}{
		"COMPUTE_INSTANCE_GROUP_CANARY_DEPLOYMENT",
		(MarshalTypeComputeInstanceGroupCanaryDeployStageExecutionProgress)(m),
	}

	return json.Marshal(&s)
}
