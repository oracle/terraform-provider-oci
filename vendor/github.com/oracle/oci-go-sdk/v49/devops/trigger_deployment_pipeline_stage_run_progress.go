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
	"github.com/oracle/oci-go-sdk/v49/common"
)

// TriggerDeploymentPipelineStageRunProgress Specifies Trigger Deployment Pipleline stage specific exeution details.
type TriggerDeploymentPipelineStageRunProgress struct {

	// BuildRun identifier which can be renamed and is not necessarily unique
	StageDisplayName *string `mandatory:"false" json:"stageDisplayName"`

	// Stage id
	BuildPipelineStageId *string `mandatory:"false" json:"buildPipelineStageId"`

	// The time the Stage was started executing. An RFC3339 formatted datetime string
	TimeStarted *common.SDKTime `mandatory:"false" json:"timeStarted"`

	// The time the Stage was finished executing. An RFC3339 formatted datetime string
	TimeFinished *common.SDKTime `mandatory:"false" json:"timeFinished"`

	BuildPipelineStagePredecessors *BuildPipelineStagePredecessorCollection `mandatory:"false" json:"buildPipelineStagePredecessors"`

	ExportedVariables *ExportedVariableCollection `mandatory:"false" json:"exportedVariables"`

	ArtifactOverrideParameters *DeployArtifactOverrideArgumentCollection `mandatory:"false" json:"artifactOverrideParameters"`

	// Identifier of the Deployment Trigerred.
	DeploymentId *string `mandatory:"false" json:"deploymentId"`

	// The current status of the Stage.
	Status BuildPipelineStageRunProgressStatusEnum `mandatory:"false" json:"status,omitempty"`
}

//GetStageDisplayName returns StageDisplayName
func (m TriggerDeploymentPipelineStageRunProgress) GetStageDisplayName() *string {
	return m.StageDisplayName
}

//GetBuildPipelineStageId returns BuildPipelineStageId
func (m TriggerDeploymentPipelineStageRunProgress) GetBuildPipelineStageId() *string {
	return m.BuildPipelineStageId
}

//GetTimeStarted returns TimeStarted
func (m TriggerDeploymentPipelineStageRunProgress) GetTimeStarted() *common.SDKTime {
	return m.TimeStarted
}

//GetTimeFinished returns TimeFinished
func (m TriggerDeploymentPipelineStageRunProgress) GetTimeFinished() *common.SDKTime {
	return m.TimeFinished
}

//GetStatus returns Status
func (m TriggerDeploymentPipelineStageRunProgress) GetStatus() BuildPipelineStageRunProgressStatusEnum {
	return m.Status
}

//GetBuildPipelineStagePredecessors returns BuildPipelineStagePredecessors
func (m TriggerDeploymentPipelineStageRunProgress) GetBuildPipelineStagePredecessors() *BuildPipelineStagePredecessorCollection {
	return m.BuildPipelineStagePredecessors
}

func (m TriggerDeploymentPipelineStageRunProgress) String() string {
	return common.PointerString(m)
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
