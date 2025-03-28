// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// DevOps API
//
// Use the DevOps API to create DevOps projects, configure code repositories,  add artifacts to deploy, build and test software applications, configure  target deployment environments, and deploy software applications.  For more information, see DevOps (https://docs.oracle.com/iaas/Content/devops/using/home.htm).
//

package devops

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// ComputeInstanceGroupBlueGreenTrafficShiftDeployStageExecutionProgress Specifies the Instance Group Blue-Green deployment load balancer traffic shift stage.
type ComputeInstanceGroupBlueGreenTrafficShiftDeployStageExecutionProgress struct {

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

	// The OCID of the environment where traffic is going.
	EnvironmentId *string `mandatory:"false" json:"environmentId"`

	// The current state of the stage.
	Status DeployStageExecutionProgressStatusEnum `mandatory:"false" json:"status,omitempty"`
}

// GetDeployStageDisplayName returns DeployStageDisplayName
func (m ComputeInstanceGroupBlueGreenTrafficShiftDeployStageExecutionProgress) GetDeployStageDisplayName() *string {
	return m.DeployStageDisplayName
}

// GetDeployStageId returns DeployStageId
func (m ComputeInstanceGroupBlueGreenTrafficShiftDeployStageExecutionProgress) GetDeployStageId() *string {
	return m.DeployStageId
}

// GetTimeStarted returns TimeStarted
func (m ComputeInstanceGroupBlueGreenTrafficShiftDeployStageExecutionProgress) GetTimeStarted() *common.SDKTime {
	return m.TimeStarted
}

// GetTimeFinished returns TimeFinished
func (m ComputeInstanceGroupBlueGreenTrafficShiftDeployStageExecutionProgress) GetTimeFinished() *common.SDKTime {
	return m.TimeFinished
}

// GetStatus returns Status
func (m ComputeInstanceGroupBlueGreenTrafficShiftDeployStageExecutionProgress) GetStatus() DeployStageExecutionProgressStatusEnum {
	return m.Status
}

// GetDeployStagePredecessors returns DeployStagePredecessors
func (m ComputeInstanceGroupBlueGreenTrafficShiftDeployStageExecutionProgress) GetDeployStagePredecessors() *DeployStagePredecessorCollection {
	return m.DeployStagePredecessors
}

// GetDeployStageExecutionProgressDetails returns DeployStageExecutionProgressDetails
func (m ComputeInstanceGroupBlueGreenTrafficShiftDeployStageExecutionProgress) GetDeployStageExecutionProgressDetails() []DeployStageExecutionProgressDetails {
	return m.DeployStageExecutionProgressDetails
}

func (m ComputeInstanceGroupBlueGreenTrafficShiftDeployStageExecutionProgress) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ComputeInstanceGroupBlueGreenTrafficShiftDeployStageExecutionProgress) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingDeployStageExecutionProgressStatusEnum(string(m.Status)); !ok && m.Status != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Status: %s. Supported values are: %s.", m.Status, strings.Join(GetDeployStageExecutionProgressStatusEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m ComputeInstanceGroupBlueGreenTrafficShiftDeployStageExecutionProgress) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeComputeInstanceGroupBlueGreenTrafficShiftDeployStageExecutionProgress ComputeInstanceGroupBlueGreenTrafficShiftDeployStageExecutionProgress
	s := struct {
		DiscriminatorParam string `json:"deployStageType"`
		MarshalTypeComputeInstanceGroupBlueGreenTrafficShiftDeployStageExecutionProgress
	}{
		"COMPUTE_INSTANCE_GROUP_BLUE_GREEN_TRAFFIC_SHIFT",
		(MarshalTypeComputeInstanceGroupBlueGreenTrafficShiftDeployStageExecutionProgress)(m),
	}

	return json.Marshal(&s)
}
