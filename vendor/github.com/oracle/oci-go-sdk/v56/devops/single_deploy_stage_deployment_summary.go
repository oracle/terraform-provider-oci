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

// SingleDeployStageDeploymentSummary Summary of single stage deployment.
type SingleDeployStageDeploymentSummary struct {

	// Unique identifier that is immutable on creation.
	Id *string `mandatory:"true" json:"id"`

	// The OCID of a project.
	ProjectId *string `mandatory:"true" json:"projectId"`

	// The OCID of a pipeline.
	DeployPipelineId *string `mandatory:"true" json:"deployPipelineId"`

	// The OCID of a compartment.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// Deployment identifier which can be renamed and is not necessarily unique. Avoid entering confidential information.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// Time the deployment was created. Format defined by RFC3339 (https://datatracker.ietf.org/doc/html/rfc3339).
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`

	// Time the deployment was updated. Format defined by RFC3339 (https://datatracker.ietf.org/doc/html/rfc3339).
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`

	DeploymentArguments *DeploymentArgumentCollection `mandatory:"false" json:"deploymentArguments"`

	DeployArtifactOverrideArguments *DeployArtifactOverrideArgumentCollection `mandatory:"false" json:"deployArtifactOverrideArguments"`

	// A message describing the current state in more detail. For example, can be used to provide actionable information for a resource in Failed state.
	LifecycleDetails *string `mandatory:"false" json:"lifecycleDetails"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.  See Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm). Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace. See Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm). Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// Usage of system tag keys. These predefined keys are scoped to namespaces. See Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm). Example: `{"orcl-cloud": {"free-tier-retained": "true"}}`
	SystemTags map[string]map[string]interface{} `mandatory:"false" json:"systemTags"`

	// Specifies the OCID of the stage to be deployed.
	DeployStageId *string `mandatory:"false" json:"deployStageId"`

	// The current state of the deployment.
	LifecycleState DeploymentLifecycleStateEnum `mandatory:"false" json:"lifecycleState,omitempty"`
}

//GetId returns Id
func (m SingleDeployStageDeploymentSummary) GetId() *string {
	return m.Id
}

//GetDisplayName returns DisplayName
func (m SingleDeployStageDeploymentSummary) GetDisplayName() *string {
	return m.DisplayName
}

//GetProjectId returns ProjectId
func (m SingleDeployStageDeploymentSummary) GetProjectId() *string {
	return m.ProjectId
}

//GetDeployPipelineId returns DeployPipelineId
func (m SingleDeployStageDeploymentSummary) GetDeployPipelineId() *string {
	return m.DeployPipelineId
}

//GetCompartmentId returns CompartmentId
func (m SingleDeployStageDeploymentSummary) GetCompartmentId() *string {
	return m.CompartmentId
}

//GetTimeCreated returns TimeCreated
func (m SingleDeployStageDeploymentSummary) GetTimeCreated() *common.SDKTime {
	return m.TimeCreated
}

//GetTimeUpdated returns TimeUpdated
func (m SingleDeployStageDeploymentSummary) GetTimeUpdated() *common.SDKTime {
	return m.TimeUpdated
}

//GetLifecycleState returns LifecycleState
func (m SingleDeployStageDeploymentSummary) GetLifecycleState() DeploymentLifecycleStateEnum {
	return m.LifecycleState
}

//GetDeploymentArguments returns DeploymentArguments
func (m SingleDeployStageDeploymentSummary) GetDeploymentArguments() *DeploymentArgumentCollection {
	return m.DeploymentArguments
}

//GetDeployArtifactOverrideArguments returns DeployArtifactOverrideArguments
func (m SingleDeployStageDeploymentSummary) GetDeployArtifactOverrideArguments() *DeployArtifactOverrideArgumentCollection {
	return m.DeployArtifactOverrideArguments
}

//GetLifecycleDetails returns LifecycleDetails
func (m SingleDeployStageDeploymentSummary) GetLifecycleDetails() *string {
	return m.LifecycleDetails
}

//GetFreeformTags returns FreeformTags
func (m SingleDeployStageDeploymentSummary) GetFreeformTags() map[string]string {
	return m.FreeformTags
}

//GetDefinedTags returns DefinedTags
func (m SingleDeployStageDeploymentSummary) GetDefinedTags() map[string]map[string]interface{} {
	return m.DefinedTags
}

//GetSystemTags returns SystemTags
func (m SingleDeployStageDeploymentSummary) GetSystemTags() map[string]map[string]interface{} {
	return m.SystemTags
}

func (m SingleDeployStageDeploymentSummary) String() string {
	return common.PointerString(m)
}

// MarshalJSON marshals to json representation
func (m SingleDeployStageDeploymentSummary) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeSingleDeployStageDeploymentSummary SingleDeployStageDeploymentSummary
	s := struct {
		DiscriminatorParam string `json:"deploymentType"`
		MarshalTypeSingleDeployStageDeploymentSummary
	}{
		"SINGLE_STAGE_DEPLOYMENT",
		(MarshalTypeSingleDeployStageDeploymentSummary)(m),
	}

	return json.Marshal(&s)
}
