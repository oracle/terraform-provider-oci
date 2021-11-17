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
	"github.com/oracle/oci-go-sdk/v52/common"
)

// CreateBuildStageDetails Specifies the Build Stage.
type CreateBuildStageDetails struct {

	// buildPipeline Identifier
	BuildPipelineId *string `mandatory:"true" json:"buildPipelineId"`

	BuildPipelineStagePredecessorCollection *BuildPipelineStagePredecessorCollection `mandatory:"true" json:"buildPipelineStagePredecessorCollection"`

	BuildSourceCollection *BuildSourceCollection `mandatory:"true" json:"buildSourceCollection"`

	// Stage identifier which can be renamed and is not necessarily unique
	DisplayName *string `mandatory:"false" json:"displayName"`

	// Optional description about the Stage
	Description *string `mandatory:"false" json:"description"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.  See Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm). Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace. See Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm). Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// The path to the build specification file for this Environment. The default location if not specified is build_spec.yaml
	BuildSpecFile *string `mandatory:"false" json:"buildSpecFile"`

	// Timeout for the Build Stage Execution. Value in seconds.
	StageExecutionTimeoutInSeconds *int `mandatory:"false" json:"stageExecutionTimeoutInSeconds"`

	// Name of the BuildSource in which the build_spec.yml file need to be located. If not specified, the 1st entry in the BuildSource collection will be chosen as Primary.
	PrimaryBuildSource *string `mandatory:"false" json:"primaryBuildSource"`

	// Image name for the Build Environment
	Image BuildStageImageEnum `mandatory:"true" json:"image"`
}

//GetDisplayName returns DisplayName
func (m CreateBuildStageDetails) GetDisplayName() *string {
	return m.DisplayName
}

//GetDescription returns Description
func (m CreateBuildStageDetails) GetDescription() *string {
	return m.Description
}

//GetBuildPipelineId returns BuildPipelineId
func (m CreateBuildStageDetails) GetBuildPipelineId() *string {
	return m.BuildPipelineId
}

//GetBuildPipelineStagePredecessorCollection returns BuildPipelineStagePredecessorCollection
func (m CreateBuildStageDetails) GetBuildPipelineStagePredecessorCollection() *BuildPipelineStagePredecessorCollection {
	return m.BuildPipelineStagePredecessorCollection
}

//GetFreeformTags returns FreeformTags
func (m CreateBuildStageDetails) GetFreeformTags() map[string]string {
	return m.FreeformTags
}

//GetDefinedTags returns DefinedTags
func (m CreateBuildStageDetails) GetDefinedTags() map[string]map[string]interface{} {
	return m.DefinedTags
}

func (m CreateBuildStageDetails) String() string {
	return common.PointerString(m)
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
