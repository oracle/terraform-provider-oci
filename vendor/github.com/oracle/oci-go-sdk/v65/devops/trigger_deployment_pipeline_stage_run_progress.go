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

// TriggerDeploymentPipelineStageRunProgress Specifies Trigger Deployment Pipleline stage specific run details.
type TriggerDeploymentPipelineStageRunProgress struct {

	// Build Run display name, which can be renamed and is not necessarily unique. Avoid entering confidential information.
	StageDisplayName *string `mandatory:"false" json:"stageDisplayName"`

	// The stage OCID.
	BuildPipelineStageId *string `mandatory:"false" json:"buildPipelineStageId"`

	// The time the stage started executing. Format defined by RFC3339 (https://datatracker.ietf.org/doc/html/rfc3339).
	TimeStarted *common.SDKTime `mandatory:"false" json:"timeStarted"`

	// The time the stage finished executing. Format defined by RFC3339 (https://datatracker.ietf.org/doc/html/rfc3339).
	TimeFinished *common.SDKTime `mandatory:"false" json:"timeFinished"`

	BuildPipelineStagePredecessors *BuildPipelineStagePredecessorCollection `mandatory:"false" json:"buildPipelineStagePredecessors"`

	ExportedVariables *ExportedVariableCollection `mandatory:"false" json:"exportedVariables"`

	ArtifactOverrideParameters *DeployArtifactOverrideArgumentCollection `mandatory:"false" json:"artifactOverrideParameters"`

	// Identifier of the deployment triggered.
	DeploymentId *string `mandatory:"false" json:"deploymentId"`

	// The current status of the stage.
	Status BuildPipelineStageRunProgressStatusEnum `mandatory:"false" json:"status,omitempty"`
}

// GetStageDisplayName returns StageDisplayName
func (m TriggerDeploymentPipelineStageRunProgress) GetStageDisplayName() *string {
	return m.StageDisplayName
}

// GetBuildPipelineStageId returns BuildPipelineStageId
func (m TriggerDeploymentPipelineStageRunProgress) GetBuildPipelineStageId() *string {
	return m.BuildPipelineStageId
}

// GetTimeStarted returns TimeStarted
func (m TriggerDeploymentPipelineStageRunProgress) GetTimeStarted() *common.SDKTime {
	return m.TimeStarted
}

// GetTimeFinished returns TimeFinished
func (m TriggerDeploymentPipelineStageRunProgress) GetTimeFinished() *common.SDKTime {
	return m.TimeFinished
}

// GetStatus returns Status
func (m TriggerDeploymentPipelineStageRunProgress) GetStatus() BuildPipelineStageRunProgressStatusEnum {
	return m.Status
}

// GetBuildPipelineStagePredecessors returns BuildPipelineStagePredecessors
func (m TriggerDeploymentPipelineStageRunProgress) GetBuildPipelineStagePredecessors() *BuildPipelineStagePredecessorCollection {
	return m.BuildPipelineStagePredecessors
}

func (m TriggerDeploymentPipelineStageRunProgress) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m TriggerDeploymentPipelineStageRunProgress) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingBuildPipelineStageRunProgressStatusEnum(string(m.Status)); !ok && m.Status != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Status: %s. Supported values are: %s.", m.Status, strings.Join(GetBuildPipelineStageRunProgressStatusEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m TriggerDeploymentPipelineStageRunProgress) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeTriggerDeploymentPipelineStageRunProgress TriggerDeploymentPipelineStageRunProgress
	s := struct {
		DiscriminatorParam string `json:"buildPipelineStageType"`
		MarshalTypeTriggerDeploymentPipelineStageRunProgress
	}{
		"TRIGGER_DEPLOYMENT_PIPELINE",
		(MarshalTypeTriggerDeploymentPipelineStageRunProgress)(m),
	}

	return json.Marshal(&s)
}
