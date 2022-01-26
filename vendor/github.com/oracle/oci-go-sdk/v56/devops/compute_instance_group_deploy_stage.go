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

// ComputeInstanceGroupDeployStage Specifies the instance group rolling deployment stage.
type ComputeInstanceGroupDeployStage struct {

	// Unique identifier that is immutable on creation.
	Id *string `mandatory:"true" json:"id"`

	// The OCID of a project.
	ProjectId *string `mandatory:"true" json:"projectId"`

	// The OCID of a pipeline.
	DeployPipelineId *string `mandatory:"true" json:"deployPipelineId"`

	// The OCID of a compartment.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// A compute instance group environment OCID for rolling deployment.
	ComputeInstanceGroupDeployEnvironmentId *string `mandatory:"true" json:"computeInstanceGroupDeployEnvironmentId"`

	// The OCID of the artifact that contains the deployment specification.
	DeploymentSpecDeployArtifactId *string `mandatory:"true" json:"deploymentSpecDeployArtifactId"`

	RolloutPolicy ComputeInstanceGroupRolloutPolicy `mandatory:"true" json:"rolloutPolicy"`

	// Optional description about the deployment stage.
	Description *string `mandatory:"false" json:"description"`

	// Deployment stage display name, which can be renamed and is not necessarily unique. Avoid entering confidential information.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// Time the deployment stage was created. Format defined by RFC3339 (https://datatracker.ietf.org/doc/html/rfc3339).
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`

	// Time the deployment stage was updated. Format defined by RFC3339 (https://datatracker.ietf.org/doc/html/rfc3339).
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`

	// A message describing the current state in more detail. For example, can be used to provide actionable information for a resource in Failed state.
	LifecycleDetails *string `mandatory:"false" json:"lifecycleDetails"`

	DeployStagePredecessorCollection *DeployStagePredecessorCollection `mandatory:"false" json:"deployStagePredecessorCollection"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.  See Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm). Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace. See Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm). Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// Usage of system tag keys. These predefined keys are scoped to namespaces. See Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm). Example: `{"orcl-cloud": {"free-tier-retained": "true"}}`
	SystemTags map[string]map[string]interface{} `mandatory:"false" json:"systemTags"`

	// Additional file artifact OCIDs.
	DeployArtifactIds []string `mandatory:"false" json:"deployArtifactIds"`

	RollbackPolicy DeployStageRollbackPolicy `mandatory:"false" json:"rollbackPolicy"`

	FailurePolicy ComputeInstanceGroupFailurePolicy `mandatory:"false" json:"failurePolicy"`

	LoadBalancerConfig *LoadBalancerConfig `mandatory:"false" json:"loadBalancerConfig"`

	// The current state of the deployment stage.
	LifecycleState DeployStageLifecycleStateEnum `mandatory:"false" json:"lifecycleState,omitempty"`
}

//GetId returns Id
func (m ComputeInstanceGroupDeployStage) GetId() *string {
	return m.Id
}

//GetDescription returns Description
func (m ComputeInstanceGroupDeployStage) GetDescription() *string {
	return m.Description
}

//GetDisplayName returns DisplayName
func (m ComputeInstanceGroupDeployStage) GetDisplayName() *string {
	return m.DisplayName
}

//GetProjectId returns ProjectId
func (m ComputeInstanceGroupDeployStage) GetProjectId() *string {
	return m.ProjectId
}

//GetDeployPipelineId returns DeployPipelineId
func (m ComputeInstanceGroupDeployStage) GetDeployPipelineId() *string {
	return m.DeployPipelineId
}

//GetCompartmentId returns CompartmentId
func (m ComputeInstanceGroupDeployStage) GetCompartmentId() *string {
	return m.CompartmentId
}

//GetTimeCreated returns TimeCreated
func (m ComputeInstanceGroupDeployStage) GetTimeCreated() *common.SDKTime {
	return m.TimeCreated
}

//GetTimeUpdated returns TimeUpdated
func (m ComputeInstanceGroupDeployStage) GetTimeUpdated() *common.SDKTime {
	return m.TimeUpdated
}

//GetLifecycleState returns LifecycleState
func (m ComputeInstanceGroupDeployStage) GetLifecycleState() DeployStageLifecycleStateEnum {
	return m.LifecycleState
}

//GetLifecycleDetails returns LifecycleDetails
func (m ComputeInstanceGroupDeployStage) GetLifecycleDetails() *string {
	return m.LifecycleDetails
}

//GetDeployStagePredecessorCollection returns DeployStagePredecessorCollection
func (m ComputeInstanceGroupDeployStage) GetDeployStagePredecessorCollection() *DeployStagePredecessorCollection {
	return m.DeployStagePredecessorCollection
}

//GetFreeformTags returns FreeformTags
func (m ComputeInstanceGroupDeployStage) GetFreeformTags() map[string]string {
	return m.FreeformTags
}

//GetDefinedTags returns DefinedTags
func (m ComputeInstanceGroupDeployStage) GetDefinedTags() map[string]map[string]interface{} {
	return m.DefinedTags
}

//GetSystemTags returns SystemTags
func (m ComputeInstanceGroupDeployStage) GetSystemTags() map[string]map[string]interface{} {
	return m.SystemTags
}

func (m ComputeInstanceGroupDeployStage) String() string {
	return common.PointerString(m)
}

// MarshalJSON marshals to json representation
func (m ComputeInstanceGroupDeployStage) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeComputeInstanceGroupDeployStage ComputeInstanceGroupDeployStage
	s := struct {
		DiscriminatorParam string `json:"deployStageType"`
		MarshalTypeComputeInstanceGroupDeployStage
	}{
		"COMPUTE_INSTANCE_GROUP_ROLLING_DEPLOYMENT",
		(MarshalTypeComputeInstanceGroupDeployStage)(m),
	}

	return json.Marshal(&s)
}

// UnmarshalJSON unmarshals from json
func (m *ComputeInstanceGroupDeployStage) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		Description                             *string                           `json:"description"`
		DisplayName                             *string                           `json:"displayName"`
		TimeCreated                             *common.SDKTime                   `json:"timeCreated"`
		TimeUpdated                             *common.SDKTime                   `json:"timeUpdated"`
		LifecycleState                          DeployStageLifecycleStateEnum     `json:"lifecycleState"`
		LifecycleDetails                        *string                           `json:"lifecycleDetails"`
		DeployStagePredecessorCollection        *DeployStagePredecessorCollection `json:"deployStagePredecessorCollection"`
		FreeformTags                            map[string]string                 `json:"freeformTags"`
		DefinedTags                             map[string]map[string]interface{} `json:"definedTags"`
		SystemTags                              map[string]map[string]interface{} `json:"systemTags"`
		DeployArtifactIds                       []string                          `json:"deployArtifactIds"`
		RollbackPolicy                          deploystagerollbackpolicy         `json:"rollbackPolicy"`
		FailurePolicy                           computeinstancegroupfailurepolicy `json:"failurePolicy"`
		LoadBalancerConfig                      *LoadBalancerConfig               `json:"loadBalancerConfig"`
		Id                                      *string                           `json:"id"`
		ProjectId                               *string                           `json:"projectId"`
		DeployPipelineId                        *string                           `json:"deployPipelineId"`
		CompartmentId                           *string                           `json:"compartmentId"`
		ComputeInstanceGroupDeployEnvironmentId *string                           `json:"computeInstanceGroupDeployEnvironmentId"`
		DeploymentSpecDeployArtifactId          *string                           `json:"deploymentSpecDeployArtifactId"`
		RolloutPolicy                           computeinstancegrouprolloutpolicy `json:"rolloutPolicy"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.Description = model.Description

	m.DisplayName = model.DisplayName

	m.TimeCreated = model.TimeCreated

	m.TimeUpdated = model.TimeUpdated

	m.LifecycleState = model.LifecycleState

	m.LifecycleDetails = model.LifecycleDetails

	m.DeployStagePredecessorCollection = model.DeployStagePredecessorCollection

	m.FreeformTags = model.FreeformTags

	m.DefinedTags = model.DefinedTags

	m.SystemTags = model.SystemTags

	m.DeployArtifactIds = make([]string, len(model.DeployArtifactIds))
	for i, n := range model.DeployArtifactIds {
		m.DeployArtifactIds[i] = n
	}

	nn, e = model.RollbackPolicy.UnmarshalPolymorphicJSON(model.RollbackPolicy.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.RollbackPolicy = nn.(DeployStageRollbackPolicy)
	} else {
		m.RollbackPolicy = nil
	}

	nn, e = model.FailurePolicy.UnmarshalPolymorphicJSON(model.FailurePolicy.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.FailurePolicy = nn.(ComputeInstanceGroupFailurePolicy)
	} else {
		m.FailurePolicy = nil
	}

	m.LoadBalancerConfig = model.LoadBalancerConfig

	m.Id = model.Id

	m.ProjectId = model.ProjectId

	m.DeployPipelineId = model.DeployPipelineId

	m.CompartmentId = model.CompartmentId

	m.ComputeInstanceGroupDeployEnvironmentId = model.ComputeInstanceGroupDeployEnvironmentId

	m.DeploymentSpecDeployArtifactId = model.DeploymentSpecDeployArtifactId

	nn, e = model.RolloutPolicy.UnmarshalPolymorphicJSON(model.RolloutPolicy.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.RolloutPolicy = nn.(ComputeInstanceGroupRolloutPolicy)
	} else {
		m.RolloutPolicy = nil
	}

	return
}
