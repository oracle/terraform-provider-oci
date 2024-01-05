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

// BuildStageSummary Specifies the build stage.
type BuildStageSummary struct {

	// Unique identifier that is immutable on creation.
	Id *string `mandatory:"true" json:"id"`

	// The OCID of the DevOps project.
	ProjectId *string `mandatory:"true" json:"projectId"`

	// The OCID of the build pipeline.
	BuildPipelineId *string `mandatory:"true" json:"buildPipelineId"`

	// The OCID of the compartment where the pipeline is created.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

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

	// The path to the build specification file for this environment. The default location of the file if not specified is build_spec.yaml.
	BuildSpecFile *string `mandatory:"false" json:"buildSpecFile"`

	// Timeout for the build stage execution. Specify value in seconds.
	StageExecutionTimeoutInSeconds *int `mandatory:"false" json:"stageExecutionTimeoutInSeconds"`

	BuildSourceCollection *BuildSourceCollection `mandatory:"false" json:"buildSourceCollection"`

	// Name of the build source where the build_spec.yml file is located. If not specified, the first entry in the build source collection is chosen as primary build source.
	PrimaryBuildSource *string `mandatory:"false" json:"primaryBuildSource"`

	BuildRunnerShapeConfig BuildRunnerShapeConfig `mandatory:"false" json:"buildRunnerShapeConfig"`

	PrivateAccessConfig NetworkChannel `mandatory:"false" json:"privateAccessConfig"`

	// The current state of the stage.
	LifecycleState BuildPipelineStageLifecycleStateEnum `mandatory:"false" json:"lifecycleState,omitempty"`

	// Image for the build environment.
	Image BuildStageImageEnum `mandatory:"true" json:"image"`
}

// GetId returns Id
func (m BuildStageSummary) GetId() *string {
	return m.Id
}

// GetDisplayName returns DisplayName
func (m BuildStageSummary) GetDisplayName() *string {
	return m.DisplayName
}

// GetProjectId returns ProjectId
func (m BuildStageSummary) GetProjectId() *string {
	return m.ProjectId
}

// GetBuildPipelineId returns BuildPipelineId
func (m BuildStageSummary) GetBuildPipelineId() *string {
	return m.BuildPipelineId
}

// GetCompartmentId returns CompartmentId
func (m BuildStageSummary) GetCompartmentId() *string {
	return m.CompartmentId
}

// GetTimeCreated returns TimeCreated
func (m BuildStageSummary) GetTimeCreated() *common.SDKTime {
	return m.TimeCreated
}

// GetTimeUpdated returns TimeUpdated
func (m BuildStageSummary) GetTimeUpdated() *common.SDKTime {
	return m.TimeUpdated
}

// GetLifecycleState returns LifecycleState
func (m BuildStageSummary) GetLifecycleState() BuildPipelineStageLifecycleStateEnum {
	return m.LifecycleState
}

// GetLifecycleDetails returns LifecycleDetails
func (m BuildStageSummary) GetLifecycleDetails() *string {
	return m.LifecycleDetails
}

// GetDescription returns Description
func (m BuildStageSummary) GetDescription() *string {
	return m.Description
}

// GetBuildPipelineStagePredecessorCollection returns BuildPipelineStagePredecessorCollection
func (m BuildStageSummary) GetBuildPipelineStagePredecessorCollection() *BuildPipelineStagePredecessorCollection {
	return m.BuildPipelineStagePredecessorCollection
}

// GetFreeformTags returns FreeformTags
func (m BuildStageSummary) GetFreeformTags() map[string]string {
	return m.FreeformTags
}

// GetDefinedTags returns DefinedTags
func (m BuildStageSummary) GetDefinedTags() map[string]map[string]interface{} {
	return m.DefinedTags
}

// GetSystemTags returns SystemTags
func (m BuildStageSummary) GetSystemTags() map[string]map[string]interface{} {
	return m.SystemTags
}

func (m BuildStageSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m BuildStageSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingBuildPipelineStageLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetBuildPipelineStageLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingBuildStageImageEnum(string(m.Image)); !ok && m.Image != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Image: %s. Supported values are: %s.", m.Image, strings.Join(GetBuildStageImageEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m BuildStageSummary) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeBuildStageSummary BuildStageSummary
	s := struct {
		DiscriminatorParam string `json:"buildPipelineStageType"`
		MarshalTypeBuildStageSummary
	}{
		"BUILD",
		(MarshalTypeBuildStageSummary)(m),
	}

	return json.Marshal(&s)
}

// UnmarshalJSON unmarshals from json
func (m *BuildStageSummary) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		DisplayName                             *string                                  `json:"displayName"`
		TimeCreated                             *common.SDKTime                          `json:"timeCreated"`
		TimeUpdated                             *common.SDKTime                          `json:"timeUpdated"`
		LifecycleState                          BuildPipelineStageLifecycleStateEnum     `json:"lifecycleState"`
		LifecycleDetails                        *string                                  `json:"lifecycleDetails"`
		Description                             *string                                  `json:"description"`
		BuildPipelineStagePredecessorCollection *BuildPipelineStagePredecessorCollection `json:"buildPipelineStagePredecessorCollection"`
		FreeformTags                            map[string]string                        `json:"freeformTags"`
		DefinedTags                             map[string]map[string]interface{}        `json:"definedTags"`
		SystemTags                              map[string]map[string]interface{}        `json:"systemTags"`
		BuildSpecFile                           *string                                  `json:"buildSpecFile"`
		StageExecutionTimeoutInSeconds          *int                                     `json:"stageExecutionTimeoutInSeconds"`
		BuildSourceCollection                   *BuildSourceCollection                   `json:"buildSourceCollection"`
		PrimaryBuildSource                      *string                                  `json:"primaryBuildSource"`
		BuildRunnerShapeConfig                  buildrunnershapeconfig                   `json:"buildRunnerShapeConfig"`
		PrivateAccessConfig                     networkchannel                           `json:"privateAccessConfig"`
		Id                                      *string                                  `json:"id"`
		ProjectId                               *string                                  `json:"projectId"`
		BuildPipelineId                         *string                                  `json:"buildPipelineId"`
		CompartmentId                           *string                                  `json:"compartmentId"`
		Image                                   BuildStageImageEnum                      `json:"image"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.DisplayName = model.DisplayName

	m.TimeCreated = model.TimeCreated

	m.TimeUpdated = model.TimeUpdated

	m.LifecycleState = model.LifecycleState

	m.LifecycleDetails = model.LifecycleDetails

	m.Description = model.Description

	m.BuildPipelineStagePredecessorCollection = model.BuildPipelineStagePredecessorCollection

	m.FreeformTags = model.FreeformTags

	m.DefinedTags = model.DefinedTags

	m.SystemTags = model.SystemTags

	m.BuildSpecFile = model.BuildSpecFile

	m.StageExecutionTimeoutInSeconds = model.StageExecutionTimeoutInSeconds

	m.BuildSourceCollection = model.BuildSourceCollection

	m.PrimaryBuildSource = model.PrimaryBuildSource

	nn, e = model.BuildRunnerShapeConfig.UnmarshalPolymorphicJSON(model.BuildRunnerShapeConfig.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.BuildRunnerShapeConfig = nn.(BuildRunnerShapeConfig)
	} else {
		m.BuildRunnerShapeConfig = nil
	}

	nn, e = model.PrivateAccessConfig.UnmarshalPolymorphicJSON(model.PrivateAccessConfig.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.PrivateAccessConfig = nn.(NetworkChannel)
	} else {
		m.PrivateAccessConfig = nil
	}

	m.Id = model.Id

	m.ProjectId = model.ProjectId

	m.BuildPipelineId = model.BuildPipelineId

	m.CompartmentId = model.CompartmentId

	m.Image = model.Image

	return
}
