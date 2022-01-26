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

// TriggerDeploymentStage Specifies the Trigger Deployment stage, which runs another pipeline of the application.
type TriggerDeploymentStage struct {

	// Unique identifier that is immutable on creation.
	Id *string `mandatory:"true" json:"id"`

	// The OCID of the DevOps project.
	ProjectId *string `mandatory:"true" json:"projectId"`

	// The OCID of the build pipeline.
	BuildPipelineId *string `mandatory:"true" json:"buildPipelineId"`

	// The OCID of the compartment where the pipeline is created.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// A target deployment pipeline OCID that will run in this stage.
	DeployPipelineId *string `mandatory:"true" json:"deployPipelineId"`

	// A boolean flag that specifies whether all the parameters must be passed when the deployment is triggered.
	IsPassAllParametersEnabled *bool `mandatory:"true" json:"isPassAllParametersEnabled"`

	// Stage display name, which can be renamed and is not necessarily unique. Avoid entering confidential information.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// Optional description about the build stage.
	Description *string `mandatory:"false" json:"description"`

	// The time the stage was created. Format defined by RFC3339 (https://datatracker.ietf.org/doc/html/rfc3339).
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`

	// The time the stage was updated. Format defined by RFC3339 (https://datatracker.ietf.org/doc/html/rfc3339).
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`

	// A message describing the current state in more detail. For example, can be used to provide actionable information for a resource in Failed state.
	LifecycleDetails *string `mandatory:"false" json:"lifecycleDetails"`

	BuildPipelineStagePredecessorCollection *BuildPipelineStagePredecessorCollection `mandatory:"false" json:"buildPipelineStagePredecessorCollection"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.  See Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm). Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace. See Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm). Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// Usage of system tag keys. These predefined keys are scoped to namespaces. See Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm). Example: `{"orcl-cloud": {"free-tier-retained": "true"}}`
	SystemTags map[string]map[string]interface{} `mandatory:"false" json:"systemTags"`

	// The current state of the stage.
	LifecycleState BuildPipelineStageLifecycleStateEnum `mandatory:"false" json:"lifecycleState,omitempty"`
}

//GetId returns Id
func (m TriggerDeploymentStage) GetId() *string {
	return m.Id
}

//GetDisplayName returns DisplayName
func (m TriggerDeploymentStage) GetDisplayName() *string {
	return m.DisplayName
}

//GetDescription returns Description
func (m TriggerDeploymentStage) GetDescription() *string {
	return m.Description
}

//GetProjectId returns ProjectId
func (m TriggerDeploymentStage) GetProjectId() *string {
	return m.ProjectId
}

//GetBuildPipelineId returns BuildPipelineId
func (m TriggerDeploymentStage) GetBuildPipelineId() *string {
	return m.BuildPipelineId
}

//GetCompartmentId returns CompartmentId
func (m TriggerDeploymentStage) GetCompartmentId() *string {
	return m.CompartmentId
}

//GetTimeCreated returns TimeCreated
func (m TriggerDeploymentStage) GetTimeCreated() *common.SDKTime {
	return m.TimeCreated
}

//GetTimeUpdated returns TimeUpdated
func (m TriggerDeploymentStage) GetTimeUpdated() *common.SDKTime {
	return m.TimeUpdated
}

//GetLifecycleState returns LifecycleState
func (m TriggerDeploymentStage) GetLifecycleState() BuildPipelineStageLifecycleStateEnum {
	return m.LifecycleState
}

//GetLifecycleDetails returns LifecycleDetails
func (m TriggerDeploymentStage) GetLifecycleDetails() *string {
	return m.LifecycleDetails
}

//GetBuildPipelineStagePredecessorCollection returns BuildPipelineStagePredecessorCollection
func (m TriggerDeploymentStage) GetBuildPipelineStagePredecessorCollection() *BuildPipelineStagePredecessorCollection {
	return m.BuildPipelineStagePredecessorCollection
}

//GetFreeformTags returns FreeformTags
func (m TriggerDeploymentStage) GetFreeformTags() map[string]string {
	return m.FreeformTags
}

//GetDefinedTags returns DefinedTags
func (m TriggerDeploymentStage) GetDefinedTags() map[string]map[string]interface{} {
	return m.DefinedTags
}

//GetSystemTags returns SystemTags
func (m TriggerDeploymentStage) GetSystemTags() map[string]map[string]interface{} {
	return m.SystemTags
}

func (m TriggerDeploymentStage) String() string {
	return common.PointerString(m)
}

// MarshalJSON marshals to json representation
func (m TriggerDeploymentStage) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeTriggerDeploymentStage TriggerDeploymentStage
	s := struct {
		DiscriminatorParam string `json:"buildPipelineStageType"`
		MarshalTypeTriggerDeploymentStage
	}{
		"TRIGGER_DEPLOYMENT_PIPELINE",
		(MarshalTypeTriggerDeploymentStage)(m),
	}

	return json.Marshal(&s)
}
