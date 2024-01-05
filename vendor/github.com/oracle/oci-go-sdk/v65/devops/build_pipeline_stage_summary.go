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

// BuildPipelineStageSummary Summary of the Stage.
type BuildPipelineStageSummary interface {

	// Unique identifier that is immutable on creation.
	GetId() *string

	// The OCID of the DevOps project.
	GetProjectId() *string

	// The OCID of the build pipeline.
	GetBuildPipelineId() *string

	// The OCID of the compartment where the pipeline is created.
	GetCompartmentId() *string

	// Stage display name, which can be renamed and is not necessarily unique. Avoid entering confidential information.
	GetDisplayName() *string

	// The time the stage was created. Format defined by RFC3339 (https://datatracker.ietf.org/doc/html/rfc3339).
	GetTimeCreated() *common.SDKTime

	// The time the stage was updated. Format defined by RFC3339 (https://datatracker.ietf.org/doc/html/rfc3339).
	GetTimeUpdated() *common.SDKTime

	// The current state of the stage.
	GetLifecycleState() BuildPipelineStageLifecycleStateEnum

	// A message describing the current state in more detail. For example, can be used to provide actionable information for a resource in Failed state.
	GetLifecycleDetails() *string

	// Optional description about the build stage.
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
	Id                                      *string                                  `mandatory:"true" json:"id"`
	ProjectId                               *string                                  `mandatory:"true" json:"projectId"`
	BuildPipelineId                         *string                                  `mandatory:"true" json:"buildPipelineId"`
	CompartmentId                           *string                                  `mandatory:"true" json:"compartmentId"`
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
		common.Logf("Recieved unsupported enum value for BuildPipelineStageSummary: %s.", m.BuildPipelineStageType)
		return *m, nil
	}
}

// GetDisplayName returns DisplayName
func (m buildpipelinestagesummary) GetDisplayName() *string {
	return m.DisplayName
}

// GetTimeCreated returns TimeCreated
func (m buildpipelinestagesummary) GetTimeCreated() *common.SDKTime {
	return m.TimeCreated
}

// GetTimeUpdated returns TimeUpdated
func (m buildpipelinestagesummary) GetTimeUpdated() *common.SDKTime {
	return m.TimeUpdated
}

// GetLifecycleState returns LifecycleState
func (m buildpipelinestagesummary) GetLifecycleState() BuildPipelineStageLifecycleStateEnum {
	return m.LifecycleState
}

// GetLifecycleDetails returns LifecycleDetails
func (m buildpipelinestagesummary) GetLifecycleDetails() *string {
	return m.LifecycleDetails
}

// GetDescription returns Description
func (m buildpipelinestagesummary) GetDescription() *string {
	return m.Description
}

// GetBuildPipelineStagePredecessorCollection returns BuildPipelineStagePredecessorCollection
func (m buildpipelinestagesummary) GetBuildPipelineStagePredecessorCollection() *BuildPipelineStagePredecessorCollection {
	return m.BuildPipelineStagePredecessorCollection
}

// GetFreeformTags returns FreeformTags
func (m buildpipelinestagesummary) GetFreeformTags() map[string]string {
	return m.FreeformTags
}

// GetDefinedTags returns DefinedTags
func (m buildpipelinestagesummary) GetDefinedTags() map[string]map[string]interface{} {
	return m.DefinedTags
}

// GetSystemTags returns SystemTags
func (m buildpipelinestagesummary) GetSystemTags() map[string]map[string]interface{} {
	return m.SystemTags
}

// GetId returns Id
func (m buildpipelinestagesummary) GetId() *string {
	return m.Id
}

// GetProjectId returns ProjectId
func (m buildpipelinestagesummary) GetProjectId() *string {
	return m.ProjectId
}

// GetBuildPipelineId returns BuildPipelineId
func (m buildpipelinestagesummary) GetBuildPipelineId() *string {
	return m.BuildPipelineId
}

// GetCompartmentId returns CompartmentId
func (m buildpipelinestagesummary) GetCompartmentId() *string {
	return m.CompartmentId
}

func (m buildpipelinestagesummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m buildpipelinestagesummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingBuildPipelineStageLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetBuildPipelineStageLifecycleStateEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
