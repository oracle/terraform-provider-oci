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

// CreateTriggerDeploymentStageDetails Specifies Trigger Deployment Pipleline stage which runs another pipeline of the application.
type CreateTriggerDeploymentStageDetails struct {

	// buildPipeline Identifier
	BuildPipelineId *string `mandatory:"true" json:"buildPipelineId"`

	BuildPipelineStagePredecessorCollection *BuildPipelineStagePredecessorCollection `mandatory:"true" json:"buildPipelineStagePredecessorCollection"`

	// A target Pipeline ocid that will be run in this stage.
	DeployPipelineId *string `mandatory:"true" json:"deployPipelineId"`

	// A boolean flag specifies whether the parameters should be passed during the deployment trigger.
	IsPassAllParametersEnabled *bool `mandatory:"true" json:"isPassAllParametersEnabled"`

	// Stage identifier which can be renamed and is not necessarily unique
	DisplayName *string `mandatory:"false" json:"displayName"`

	// Optional description about the Stage
	Description *string `mandatory:"false" json:"description"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.  See Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm). Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace. See Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm). Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

//GetDisplayName returns DisplayName
func (m CreateTriggerDeploymentStageDetails) GetDisplayName() *string {
	return m.DisplayName
}

//GetDescription returns Description
func (m CreateTriggerDeploymentStageDetails) GetDescription() *string {
	return m.Description
}

//GetBuildPipelineId returns BuildPipelineId
func (m CreateTriggerDeploymentStageDetails) GetBuildPipelineId() *string {
	return m.BuildPipelineId
}

//GetBuildPipelineStagePredecessorCollection returns BuildPipelineStagePredecessorCollection
func (m CreateTriggerDeploymentStageDetails) GetBuildPipelineStagePredecessorCollection() *BuildPipelineStagePredecessorCollection {
	return m.BuildPipelineStagePredecessorCollection
}

//GetFreeformTags returns FreeformTags
func (m CreateTriggerDeploymentStageDetails) GetFreeformTags() map[string]string {
	return m.FreeformTags
}

//GetDefinedTags returns DefinedTags
func (m CreateTriggerDeploymentStageDetails) GetDefinedTags() map[string]map[string]interface{} {
	return m.DefinedTags
}

func (m CreateTriggerDeploymentStageDetails) String() string {
	return common.PointerString(m)
}

// MarshalJSON marshals to json representation
func (m CreateTriggerDeploymentStageDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeCreateTriggerDeploymentStageDetails CreateTriggerDeploymentStageDetails
	s := struct {
		DiscriminatorParam string `json:"buildPipelineStageType"`
		MarshalTypeCreateTriggerDeploymentStageDetails
	}{
		"TRIGGER_DEPLOYMENT_PIPELINE",
		(MarshalTypeCreateTriggerDeploymentStageDetails)(m),
	}

	return json.Marshal(&s)
}
