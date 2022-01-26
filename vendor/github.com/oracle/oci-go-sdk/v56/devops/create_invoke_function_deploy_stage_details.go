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

// CreateInvokeFunctionDeployStageDetails Specifies Invoke Function stage.
type CreateInvokeFunctionDeployStageDetails struct {

	// The OCID of a pipeline.
	DeployPipelineId *string `mandatory:"true" json:"deployPipelineId"`

	DeployStagePredecessorCollection *DeployStagePredecessorCollection `mandatory:"true" json:"deployStagePredecessorCollection"`

	// Function environment OCID.
	FunctionDeployEnvironmentId *string `mandatory:"true" json:"functionDeployEnvironmentId"`

	// A boolean flag specifies whether this stage executes asynchronously.
	IsAsync *bool `mandatory:"true" json:"isAsync"`

	// A boolean flag specifies whether the invoked function should be validated.
	IsValidationEnabled *bool `mandatory:"true" json:"isValidationEnabled"`

	// Optional description about the deployment stage.
	Description *string `mandatory:"false" json:"description"`

	// Deployment stage display name, which can be renamed and is not necessarily unique. Avoid entering confidential information.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.  See Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm). Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace. See Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm). Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// Optional binary artifact OCID user may provide to this stage.
	DeployArtifactId *string `mandatory:"false" json:"deployArtifactId"`
}

//GetDescription returns Description
func (m CreateInvokeFunctionDeployStageDetails) GetDescription() *string {
	return m.Description
}

//GetDisplayName returns DisplayName
func (m CreateInvokeFunctionDeployStageDetails) GetDisplayName() *string {
	return m.DisplayName
}

//GetDeployPipelineId returns DeployPipelineId
func (m CreateInvokeFunctionDeployStageDetails) GetDeployPipelineId() *string {
	return m.DeployPipelineId
}

//GetDeployStagePredecessorCollection returns DeployStagePredecessorCollection
func (m CreateInvokeFunctionDeployStageDetails) GetDeployStagePredecessorCollection() *DeployStagePredecessorCollection {
	return m.DeployStagePredecessorCollection
}

//GetFreeformTags returns FreeformTags
func (m CreateInvokeFunctionDeployStageDetails) GetFreeformTags() map[string]string {
	return m.FreeformTags
}

//GetDefinedTags returns DefinedTags
func (m CreateInvokeFunctionDeployStageDetails) GetDefinedTags() map[string]map[string]interface{} {
	return m.DefinedTags
}

func (m CreateInvokeFunctionDeployStageDetails) String() string {
	return common.PointerString(m)
}

// MarshalJSON marshals to json representation
func (m CreateInvokeFunctionDeployStageDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeCreateInvokeFunctionDeployStageDetails CreateInvokeFunctionDeployStageDetails
	s := struct {
		DiscriminatorParam string `json:"deployStageType"`
		MarshalTypeCreateInvokeFunctionDeployStageDetails
	}{
		"INVOKE_FUNCTION",
		(MarshalTypeCreateInvokeFunctionDeployStageDetails)(m),
	}

	return json.Marshal(&s)
}
