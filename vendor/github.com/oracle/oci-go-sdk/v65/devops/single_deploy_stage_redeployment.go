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

// SingleDeployStageRedeployment Redeployment of a single stage of a previous deployment.
type SingleDeployStageRedeployment struct {

	// Unique identifier that is immutable on creation.
	Id *string `mandatory:"true" json:"id"`

	// The OCID of a project.
	ProjectId *string `mandatory:"true" json:"projectId"`

	// The OCID of a pipeline.
	DeployPipelineId *string `mandatory:"true" json:"deployPipelineId"`

	// The OCID of a compartment.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// Specifies the OCID of the stage to be redeployed.
	DeployStageId *string `mandatory:"true" json:"deployStageId"`

	DeployPipelineArtifacts *DeployPipelineArtifactCollection `mandatory:"false" json:"deployPipelineArtifacts"`

	DeployPipelineEnvironments *DeployPipelineEnvironmentCollection `mandatory:"false" json:"deployPipelineEnvironments"`

	// Deployment identifier which can be renamed and is not necessarily unique. Avoid entering confidential information.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// Time the deployment was created. Format defined by RFC3339 (https://datatracker.ietf.org/doc/html/rfc3339).
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`

	// Time the deployment was updated. Format defined by RFC3339 (https://datatracker.ietf.org/doc/html/rfc3339).
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`

	// A message describing the current state in more detail. For example, can be used to provide actionable information for a resource in Failed state.
	LifecycleDetails *string `mandatory:"false" json:"lifecycleDetails"`

	DeploymentArguments *DeploymentArgumentCollection `mandatory:"false" json:"deploymentArguments"`

	DeployStageOverrideArguments *DeployStageOverrideArgumentCollection `mandatory:"false" json:"deployStageOverrideArguments"`

	DeployArtifactOverrideArguments *DeployArtifactOverrideArgumentCollection `mandatory:"false" json:"deployArtifactOverrideArguments"`

	DeploymentExecutionProgress *DeploymentExecutionProgress `mandatory:"false" json:"deploymentExecutionProgress"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.  See Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm). Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace. See Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm). Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// Usage of system tag keys. These predefined keys are scoped to namespaces. See Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm). Example: `{"orcl-cloud": {"free-tier-retained": "true"}}`
	SystemTags map[string]map[string]interface{} `mandatory:"false" json:"systemTags"`

	// Specifies the OCID of the previous deployment to be redeployed.
	PreviousDeploymentId *string `mandatory:"false" json:"previousDeploymentId"`

	// The current state of the deployment.
	LifecycleState DeploymentLifecycleStateEnum `mandatory:"false" json:"lifecycleState,omitempty"`
}

// GetDeployPipelineArtifacts returns DeployPipelineArtifacts
func (m SingleDeployStageRedeployment) GetDeployPipelineArtifacts() *DeployPipelineArtifactCollection {
	return m.DeployPipelineArtifacts
}

// GetDeployPipelineEnvironments returns DeployPipelineEnvironments
func (m SingleDeployStageRedeployment) GetDeployPipelineEnvironments() *DeployPipelineEnvironmentCollection {
	return m.DeployPipelineEnvironments
}

// GetId returns Id
func (m SingleDeployStageRedeployment) GetId() *string {
	return m.Id
}

// GetDisplayName returns DisplayName
func (m SingleDeployStageRedeployment) GetDisplayName() *string {
	return m.DisplayName
}

// GetProjectId returns ProjectId
func (m SingleDeployStageRedeployment) GetProjectId() *string {
	return m.ProjectId
}

// GetDeployPipelineId returns DeployPipelineId
func (m SingleDeployStageRedeployment) GetDeployPipelineId() *string {
	return m.DeployPipelineId
}

// GetCompartmentId returns CompartmentId
func (m SingleDeployStageRedeployment) GetCompartmentId() *string {
	return m.CompartmentId
}

// GetTimeCreated returns TimeCreated
func (m SingleDeployStageRedeployment) GetTimeCreated() *common.SDKTime {
	return m.TimeCreated
}

// GetTimeUpdated returns TimeUpdated
func (m SingleDeployStageRedeployment) GetTimeUpdated() *common.SDKTime {
	return m.TimeUpdated
}

// GetLifecycleState returns LifecycleState
func (m SingleDeployStageRedeployment) GetLifecycleState() DeploymentLifecycleStateEnum {
	return m.LifecycleState
}

// GetLifecycleDetails returns LifecycleDetails
func (m SingleDeployStageRedeployment) GetLifecycleDetails() *string {
	return m.LifecycleDetails
}

// GetDeploymentArguments returns DeploymentArguments
func (m SingleDeployStageRedeployment) GetDeploymentArguments() *DeploymentArgumentCollection {
	return m.DeploymentArguments
}

// GetDeployStageOverrideArguments returns DeployStageOverrideArguments
func (m SingleDeployStageRedeployment) GetDeployStageOverrideArguments() *DeployStageOverrideArgumentCollection {
	return m.DeployStageOverrideArguments
}

// GetDeployArtifactOverrideArguments returns DeployArtifactOverrideArguments
func (m SingleDeployStageRedeployment) GetDeployArtifactOverrideArguments() *DeployArtifactOverrideArgumentCollection {
	return m.DeployArtifactOverrideArguments
}

// GetDeploymentExecutionProgress returns DeploymentExecutionProgress
func (m SingleDeployStageRedeployment) GetDeploymentExecutionProgress() *DeploymentExecutionProgress {
	return m.DeploymentExecutionProgress
}

// GetFreeformTags returns FreeformTags
func (m SingleDeployStageRedeployment) GetFreeformTags() map[string]string {
	return m.FreeformTags
}

// GetDefinedTags returns DefinedTags
func (m SingleDeployStageRedeployment) GetDefinedTags() map[string]map[string]interface{} {
	return m.DefinedTags
}

// GetSystemTags returns SystemTags
func (m SingleDeployStageRedeployment) GetSystemTags() map[string]map[string]interface{} {
	return m.SystemTags
}

func (m SingleDeployStageRedeployment) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m SingleDeployStageRedeployment) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingDeploymentLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetDeploymentLifecycleStateEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m SingleDeployStageRedeployment) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeSingleDeployStageRedeployment SingleDeployStageRedeployment
	s := struct {
		DiscriminatorParam string `json:"deploymentType"`
		MarshalTypeSingleDeployStageRedeployment
	}{
		"SINGLE_STAGE_REDEPLOYMENT",
		(MarshalTypeSingleDeployStageRedeployment)(m),
	}

	return json.Marshal(&s)
}
