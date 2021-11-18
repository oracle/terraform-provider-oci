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

// UpdateTriggerDeploymentStageDetails Specifies trigger Deployment Pipleline stage which runs another pipeline of the application.
type UpdateTriggerDeploymentStageDetails struct {

	// Stage identifier which can be renamed and is not necessarily unique
	DisplayName *string `mandatory:"false" json:"displayName"`

	// Optional description about the BuildStage
	Description *string `mandatory:"false" json:"description"`

	BuildPipelineStagePredecessorCollection *BuildPipelineStagePredecessorCollection `mandatory:"false" json:"buildPipelineStagePredecessorCollection"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.  See Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm). Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace. See Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm). Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// A target Pipeline ocid that will be run in this stage.
	DeployPipelineId *string `mandatory:"false" json:"deployPipelineId"`

	// A boolean flag specifies whether the parameters should be passed during the deployment trigger.
	IsPassAllParametersEnabled *bool `mandatory:"false" json:"isPassAllParametersEnabled"`
}

//GetDisplayName returns DisplayName
func (m UpdateTriggerDeploymentStageDetails) GetDisplayName() *string {
	return m.DisplayName
}

//GetDescription returns Description
func (m UpdateTriggerDeploymentStageDetails) GetDescription() *string {
	return m.Description
}

//GetBuildPipelineStagePredecessorCollection returns BuildPipelineStagePredecessorCollection
func (m UpdateTriggerDeploymentStageDetails) GetBuildPipelineStagePredecessorCollection() *BuildPipelineStagePredecessorCollection {
	return m.BuildPipelineStagePredecessorCollection
}

//GetFreeformTags returns FreeformTags
func (m UpdateTriggerDeploymentStageDetails) GetFreeformTags() map[string]string {
	return m.FreeformTags
}

//GetDefinedTags returns DefinedTags
func (m UpdateTriggerDeploymentStageDetails) GetDefinedTags() map[string]map[string]interface{} {
	return m.DefinedTags
}

func (m UpdateTriggerDeploymentStageDetails) String() string {
	return common.PointerString(m)
}

// MarshalJSON marshals to json representation
func (m UpdateTriggerDeploymentStageDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeUpdateTriggerDeploymentStageDetails UpdateTriggerDeploymentStageDetails
	s := struct {
		DiscriminatorParam string `json:"buildPipelineStageType"`
		MarshalTypeUpdateTriggerDeploymentStageDetails
	}{
		"TRIGGER_DEPLOYMENT_PIPELINE",
		(MarshalTypeUpdateTriggerDeploymentStageDetails)(m),
	}

	return json.Marshal(&s)
}
