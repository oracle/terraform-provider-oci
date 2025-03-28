// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// DevOps API
//
// Use the DevOps API to create DevOps projects, configure code repositories,  add artifacts to deploy, build and test software applications, configure  target deployment environments, and deploy software applications.  For more information, see DevOps (https://docs.oracle.com/iaas/Content/devops/using/home.htm).
//

package devops

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// CreateBuildStageDetails Specifies the build stage.
type CreateBuildStageDetails struct {

	// The OCID of the build pipeline.
	BuildPipelineId *string `mandatory:"true" json:"buildPipelineId"`

	BuildPipelineStagePredecessorCollection *BuildPipelineStagePredecessorCollection `mandatory:"true" json:"buildPipelineStagePredecessorCollection"`

	BuildSourceCollection *BuildSourceCollection `mandatory:"true" json:"buildSourceCollection"`

	// Stage display name, which can be renamed and is not necessarily unique. Avoid entering confidential information.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// Optional description about the stage.
	Description *string `mandatory:"false" json:"description"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.  See Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace. See Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// The path to the build specification file for this environment. The default location of the file if not specified is build_spec.yaml.
	BuildSpecFile *string `mandatory:"false" json:"buildSpecFile"`

	// Timeout for the build stage execution. Specify value in seconds.
	StageExecutionTimeoutInSeconds *int `mandatory:"false" json:"stageExecutionTimeoutInSeconds"`

	// Name of the build source where the build_spec.yml file is located. If not specified, the first entry in the build source collection is chosen as primary build source.
	PrimaryBuildSource *string `mandatory:"false" json:"primaryBuildSource"`

	BuildRunnerShapeConfig BuildRunnerShapeConfig `mandatory:"false" json:"buildRunnerShapeConfig"`

	PrivateAccessConfig NetworkChannel `mandatory:"false" json:"privateAccessConfig"`

	// Image name for the build environment
	Image BuildStageImageEnum `mandatory:"true" json:"image"`
}

// GetDisplayName returns DisplayName
func (m CreateBuildStageDetails) GetDisplayName() *string {
	return m.DisplayName
}

// GetDescription returns Description
func (m CreateBuildStageDetails) GetDescription() *string {
	return m.Description
}

// GetBuildPipelineId returns BuildPipelineId
func (m CreateBuildStageDetails) GetBuildPipelineId() *string {
	return m.BuildPipelineId
}

// GetBuildPipelineStagePredecessorCollection returns BuildPipelineStagePredecessorCollection
func (m CreateBuildStageDetails) GetBuildPipelineStagePredecessorCollection() *BuildPipelineStagePredecessorCollection {
	return m.BuildPipelineStagePredecessorCollection
}

// GetFreeformTags returns FreeformTags
func (m CreateBuildStageDetails) GetFreeformTags() map[string]string {
	return m.FreeformTags
}

// GetDefinedTags returns DefinedTags
func (m CreateBuildStageDetails) GetDefinedTags() map[string]map[string]interface{} {
	return m.DefinedTags
}

func (m CreateBuildStageDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CreateBuildStageDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingBuildStageImageEnum(string(m.Image)); !ok && m.Image != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Image: %s. Supported values are: %s.", m.Image, strings.Join(GetBuildStageImageEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m CreateBuildStageDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeCreateBuildStageDetails CreateBuildStageDetails
	s := struct {
		DiscriminatorParam string `json:"buildPipelineStageType"`
		MarshalTypeCreateBuildStageDetails
	}{
		"BUILD",
		(MarshalTypeCreateBuildStageDetails)(m),
	}

	return json.Marshal(&s)
}

// UnmarshalJSON unmarshals from json
func (m *CreateBuildStageDetails) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		DisplayName                             *string                                  `json:"displayName"`
		Description                             *string                                  `json:"description"`
		FreeformTags                            map[string]string                        `json:"freeformTags"`
		DefinedTags                             map[string]map[string]interface{}        `json:"definedTags"`
		BuildSpecFile                           *string                                  `json:"buildSpecFile"`
		StageExecutionTimeoutInSeconds          *int                                     `json:"stageExecutionTimeoutInSeconds"`
		PrimaryBuildSource                      *string                                  `json:"primaryBuildSource"`
		BuildRunnerShapeConfig                  buildrunnershapeconfig                   `json:"buildRunnerShapeConfig"`
		PrivateAccessConfig                     networkchannel                           `json:"privateAccessConfig"`
		BuildPipelineId                         *string                                  `json:"buildPipelineId"`
		BuildPipelineStagePredecessorCollection *BuildPipelineStagePredecessorCollection `json:"buildPipelineStagePredecessorCollection"`
		Image                                   BuildStageImageEnum                      `json:"image"`
		BuildSourceCollection                   *BuildSourceCollection                   `json:"buildSourceCollection"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.DisplayName = model.DisplayName

	m.Description = model.Description

	m.FreeformTags = model.FreeformTags

	m.DefinedTags = model.DefinedTags

	m.BuildSpecFile = model.BuildSpecFile

	m.StageExecutionTimeoutInSeconds = model.StageExecutionTimeoutInSeconds

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

	m.BuildPipelineId = model.BuildPipelineId

	m.BuildPipelineStagePredecessorCollection = model.BuildPipelineStagePredecessorCollection

	m.Image = model.Image

	m.BuildSourceCollection = model.BuildSourceCollection

	return
}
