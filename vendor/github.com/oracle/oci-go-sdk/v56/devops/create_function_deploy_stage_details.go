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

// CreateFunctionDeployStageDetails Specifies the Function stage.
type CreateFunctionDeployStageDetails struct {

	// The OCID of a pipeline.
	DeployPipelineId *string `mandatory:"true" json:"deployPipelineId"`

	DeployStagePredecessorCollection *DeployStagePredecessorCollection `mandatory:"true" json:"deployStagePredecessorCollection"`

	// Function environment OCID.
	FunctionDeployEnvironmentId *string `mandatory:"true" json:"functionDeployEnvironmentId"`

	// A Docker image artifact OCID.
	DockerImageDeployArtifactId *string `mandatory:"true" json:"dockerImageDeployArtifactId"`

	// Optional description about the deployment stage.
	Description *string `mandatory:"false" json:"description"`

	// Deployment stage display name, which can be renamed and is not necessarily unique. Avoid entering confidential information.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.  See Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm). Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace. See Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm). Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// User provided key and value pair configuration, which is assigned through constants or parameter.
	Config map[string]string `mandatory:"false" json:"config"`

	// Maximum usable memory for the Function (in MB).
	MaxMemoryInMBs *int64 `mandatory:"false" json:"maxMemoryInMBs"`

	// Timeout for execution of the Function. Value in seconds.
	FunctionTimeoutInSeconds *int `mandatory:"false" json:"functionTimeoutInSeconds"`
}

//GetDescription returns Description
func (m CreateFunctionDeployStageDetails) GetDescription() *string {
	return m.Description
}

//GetDisplayName returns DisplayName
func (m CreateFunctionDeployStageDetails) GetDisplayName() *string {
	return m.DisplayName
}

//GetDeployPipelineId returns DeployPipelineId
func (m CreateFunctionDeployStageDetails) GetDeployPipelineId() *string {
	return m.DeployPipelineId
}

//GetDeployStagePredecessorCollection returns DeployStagePredecessorCollection
func (m CreateFunctionDeployStageDetails) GetDeployStagePredecessorCollection() *DeployStagePredecessorCollection {
	return m.DeployStagePredecessorCollection
}

//GetFreeformTags returns FreeformTags
func (m CreateFunctionDeployStageDetails) GetFreeformTags() map[string]string {
	return m.FreeformTags
}

//GetDefinedTags returns DefinedTags
func (m CreateFunctionDeployStageDetails) GetDefinedTags() map[string]map[string]interface{} {
	return m.DefinedTags
}

func (m CreateFunctionDeployStageDetails) String() string {
	return common.PointerString(m)
}

// MarshalJSON marshals to json representation
func (m CreateFunctionDeployStageDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeCreateFunctionDeployStageDetails CreateFunctionDeployStageDetails
	s := struct {
		DiscriminatorParam string `json:"deployStageType"`
		MarshalTypeCreateFunctionDeployStageDetails
	}{
		"DEPLOY_FUNCTION",
		(MarshalTypeCreateFunctionDeployStageDetails)(m),
	}

	return json.Marshal(&s)
}
