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
	"github.com/oracle/oci-go-sdk/v50/common"
)

// BuildPipelineStageSummary Summary of the Stage.
type BuildPipelineStageSummary interface {

	// Unique identifier that is immutable on creation
	GetId() *string

	// Build Project Identifier
	GetProjectId() *string

	// Build Pipeline Identifier
	GetBuildPipelineId() *string

	// Compartment Identifier
	GetCompartmentId() *string

	// Stage identifier which can be renamed and is not necessarily unique
	GetDisplayName() *string

	// The time at which the Stage was created. An RFC3339 formatted datetime string
	GetTimeCreated() *common.SDKTime

	// The time at which the Stage was updated. An RFC3339 formatted datetime string
	GetTimeUpdated() *common.SDKTime

	// The current state of the Stage.
	GetLifecycleState() BuildPipelineStageLifecycleStateEnum

	// A message describing the current state in more detail. For example, can be used to provide actionable information for a resource in Failed state.
	GetLifecycleDetails() *string

	// Optional description about the BuildStage
	GetDescription() *string

	GetBuildPipelineStagePredecessorCollection() *BuildPipelineStagePredecessorCollection

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.  See Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm). Example: `{"bar-key": "value"}`
	GetFreeformTags() map[string]string

	// Defined tags for this resource. Each key is predefined and scoped to a namespace. See Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm). Example: `{"foo-namespace": {"bar-key": "value"}}`
	GetDefinedTags() map[string]map[string]interface{}

	// Usage of system tag keys. These predefined keys are scoped to namespaces. See Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm). Example: `{"orcl-cloud": {"free-tier-retained": "true"}}`
	GetSystemTags() map[string]map[string]interface{}
}

type buildpipelinestagesummary struct {
	JsonData                                []byte
	Id                                      *string                                  `mandatory:"true" json:"id"`
	ProjectId                               *string                                  `mandatory:"true" json:"projectId"`
	BuildPipelineId                         *string                                  `mandatory:"true" json:"buildPipelineId"`
	CompartmentId                           *string                                  `mandatory:"true" json:"compartmentId"`
	DisplayName                             *string                                  `mandatory:"false" json:"displayName"`
	TimeCreated                             *common.SDKTime                          `mandatory:"false" json:"timeCreated"`
	TimeUpdated                             *common.SDKTime                          `mandatory:"false" json:"timeUpdated"`
	LifecycleState                          BuildPipelineStageLifecycleStateEnum     `mandatory:"false" json:"lifecycleState,omitempty"`
	LifecycleDetails                        *string                                  `mandatory:"false" json:"lifecycleDetails"`
	Description                             *string                                  `mandatory:"false" json:"description"`
	BuildPipelineStagePredecessorCollection *BuildPipelineStagePredecessorCollection `mandatory:"false" json:"buildPipelineStagePredecessorCollection"`
	FreeformTags                            map[string]string                        `mandatory:"false" json:"freeformTags"`
	DefinedTags                             map[string]map[string]interface{}        `mandatory:"false" json:"definedTags"`
	SystemTags                              map[string]map[string]interface{}        `mandatory:"false" json:"systemTags"`
	BuildPipelineStageType                  string                                   `json:"buildPipelineStageType"`
}

// UnmarshalJSON unmarshals json
func (m *buildpipelinestagesummary) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerbuildpipelinestagesummary buildpipelinestagesummary
	s := struct {
		Model Unmarshalerbuildpipelinestagesummary
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.Id = s.Model.Id
	m.ProjectId = s.Model.ProjectId
	m.BuildPipelineId = s.Model.BuildPipelineId
	m.CompartmentId = s.Model.CompartmentId
	m.DisplayName = s.Model.DisplayName
	m.TimeCreated = s.Model.TimeCreated
	m.TimeUpdated = s.Model.TimeUpdated
	m.LifecycleState = s.Model.LifecycleState
	m.LifecycleDetails = s.Model.LifecycleDetails
	m.Description = s.Model.Description
	m.BuildPipelineStagePredecessorCollection = s.Model.BuildPipelineStagePredecessorCollection
	m.FreeformTags = s.Model.FreeformTags
	m.DefinedTags = s.Model.DefinedTags
	m.SystemTags = s.Model.SystemTags
	m.BuildPipelineStageType = s.Model.BuildPipelineStageType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *buildpipelinestagesummary) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.BuildPipelineStageType {
	case "WAIT":
		mm := WaitStageSummary{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "BUILD":
		mm := BuildStageSummary{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "DELIVER_ARTIFACT":
		mm := DeliverArtifactStageSummary{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "TRIGGER_DEPLOYMENT_PIPELINE":
		mm := TriggerDeploymentStageSummary{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		return *m, nil
	}
}

//GetId returns Id
func (m buildpipelinestagesummary) GetId() *string {
	return m.Id
}

//GetProjectId returns ProjectId
func (m buildpipelinestagesummary) GetProjectId() *string {
	return m.ProjectId
}

//GetBuildPipelineId returns BuildPipelineId
func (m buildpipelinestagesummary) GetBuildPipelineId() *string {
	return m.BuildPipelineId
}

//GetCompartmentId returns CompartmentId
func (m buildpipelinestagesummary) GetCompartmentId() *string {
	return m.CompartmentId
}

//GetDisplayName returns DisplayName
func (m buildpipelinestagesummary) GetDisplayName() *string {
	return m.DisplayName
}

//GetTimeCreated returns TimeCreated
func (m buildpipelinestagesummary) GetTimeCreated() *common.SDKTime {
	return m.TimeCreated
}

//GetTimeUpdated returns TimeUpdated
func (m buildpipelinestagesummary) GetTimeUpdated() *common.SDKTime {
	return m.TimeUpdated
}

//GetLifecycleState returns LifecycleState
func (m buildpipelinestagesummary) GetLifecycleState() BuildPipelineStageLifecycleStateEnum {
	return m.LifecycleState
}

//GetLifecycleDetails returns LifecycleDetails
func (m buildpipelinestagesummary) GetLifecycleDetails() *string {
	return m.LifecycleDetails
}

//GetDescription returns Description
func (m buildpipelinestagesummary) GetDescription() *string {
	return m.Description
}

//GetBuildPipelineStagePredecessorCollection returns BuildPipelineStagePredecessorCollection
func (m buildpipelinestagesummary) GetBuildPipelineStagePredecessorCollection() *BuildPipelineStagePredecessorCollection {
	return m.BuildPipelineStagePredecessorCollection
}

//GetFreeformTags returns FreeformTags
func (m buildpipelinestagesummary) GetFreeformTags() map[string]string {
	return m.FreeformTags
}

//GetDefinedTags returns DefinedTags
func (m buildpipelinestagesummary) GetDefinedTags() map[string]map[string]interface{} {
	return m.DefinedTags
}

//GetSystemTags returns SystemTags
func (m buildpipelinestagesummary) GetSystemTags() map[string]map[string]interface{} {
	return m.SystemTags
}

func (m buildpipelinestagesummary) String() string {
	return common.PointerString(m)
}
