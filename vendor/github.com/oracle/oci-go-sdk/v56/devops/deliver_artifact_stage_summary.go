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

// DeliverArtifactStageSummary Specifies the Deliver Artifacts stage.
type DeliverArtifactStageSummary struct {

	// Unique identifier that is immutable on creation.
	Id *string `mandatory:"true" json:"id"`

	// The OCID of the DevOps project.
	ProjectId *string `mandatory:"true" json:"projectId"`

	// The OCID of the build pipeline.
	BuildPipelineId *string `mandatory:"true" json:"buildPipelineId"`

	// The OCID of the compartment where the pipeline is created.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	DeliverArtifactCollection *DeliverArtifactCollection `mandatory:"true" json:"deliverArtifactCollection"`

	// Stage display name, which can be renamed and is not necessarily unique. Avoid entering confidential information.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// The time the stage was created. Format defined by RFC3339 (https://datatracker.ietf.org/doc/html/rfc3339).
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`

	// The time the stage was updated. Format defined by RFC3339 (https://datatracker.ietf.org/doc/html/rfc3339).
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`

	// A message describing the current state in more detail. For example, can be used to provide actionable information for a resource in Failed state.
	LifecycleDetails *string `mandatory:"false" json:"lifecycleDetails"`

	// Optional description about the build stage.
	Description *string `mandatory:"false" json:"description"`

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
func (m DeliverArtifactStageSummary) GetId() *string {
	return m.Id
}

//GetDisplayName returns DisplayName
func (m DeliverArtifactStageSummary) GetDisplayName() *string {
	return m.DisplayName
}

//GetProjectId returns ProjectId
func (m DeliverArtifactStageSummary) GetProjectId() *string {
	return m.ProjectId
}

//GetBuildPipelineId returns BuildPipelineId
func (m DeliverArtifactStageSummary) GetBuildPipelineId() *string {
	return m.BuildPipelineId
}

//GetCompartmentId returns CompartmentId
func (m DeliverArtifactStageSummary) GetCompartmentId() *string {
	return m.CompartmentId
}

//GetTimeCreated returns TimeCreated
func (m DeliverArtifactStageSummary) GetTimeCreated() *common.SDKTime {
	return m.TimeCreated
}

//GetTimeUpdated returns TimeUpdated
func (m DeliverArtifactStageSummary) GetTimeUpdated() *common.SDKTime {
	return m.TimeUpdated
}

//GetLifecycleState returns LifecycleState
func (m DeliverArtifactStageSummary) GetLifecycleState() BuildPipelineStageLifecycleStateEnum {
	return m.LifecycleState
}

//GetLifecycleDetails returns LifecycleDetails
func (m DeliverArtifactStageSummary) GetLifecycleDetails() *string {
	return m.LifecycleDetails
}

//GetDescription returns Description
func (m DeliverArtifactStageSummary) GetDescription() *string {
	return m.Description
}

//GetBuildPipelineStagePredecessorCollection returns BuildPipelineStagePredecessorCollection
func (m DeliverArtifactStageSummary) GetBuildPipelineStagePredecessorCollection() *BuildPipelineStagePredecessorCollection {
	return m.BuildPipelineStagePredecessorCollection
}

//GetFreeformTags returns FreeformTags
func (m DeliverArtifactStageSummary) GetFreeformTags() map[string]string {
	return m.FreeformTags
}

//GetDefinedTags returns DefinedTags
func (m DeliverArtifactStageSummary) GetDefinedTags() map[string]map[string]interface{} {
	return m.DefinedTags
}

//GetSystemTags returns SystemTags
func (m DeliverArtifactStageSummary) GetSystemTags() map[string]map[string]interface{} {
	return m.SystemTags
}

func (m DeliverArtifactStageSummary) String() string {
	return common.PointerString(m)
}

// MarshalJSON marshals to json representation
func (m DeliverArtifactStageSummary) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeDeliverArtifactStageSummary DeliverArtifactStageSummary
	s := struct {
		DiscriminatorParam string `json:"buildPipelineStageType"`
		MarshalTypeDeliverArtifactStageSummary
	}{
		"DELIVER_ARTIFACT",
		(MarshalTypeDeliverArtifactStageSummary)(m),
	}

	return json.Marshal(&s)
}
