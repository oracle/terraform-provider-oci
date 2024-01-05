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

// UpdateInvokeFunctionDeployStageDetails Specifies Invoke Function stage.
type UpdateInvokeFunctionDeployStageDetails struct {

	// Optional description about the deployment stage.
	Description *string `mandatory:"false" json:"description"`

	// Deployment stage display name, which can be renamed and is not necessarily unique. Avoid entering confidential information.
	DisplayName *string `mandatory:"false" json:"displayName"`

	DeployStagePredecessorCollection *DeployStagePredecessorCollection `mandatory:"false" json:"deployStagePredecessorCollection"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.  See Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm). Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace. See Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm). Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// Function environment OCID.
	FunctionDeployEnvironmentId *string `mandatory:"false" json:"functionDeployEnvironmentId"`

	// Optional artifact OCID. The artifact will be included in the body for the function invocation during the stage's execution.
	// If the DeployArtifact.argumentSubstituitionMode is set to SUBSTITUTE_PLACEHOLDERS, then the pipeline parameter values will be used to replace the placeholders in the artifact content.
	DeployArtifactId *string `mandatory:"false" json:"deployArtifactId"`

	// A boolean flag specifies whether this stage executes asynchronously.
	IsAsync *bool `mandatory:"false" json:"isAsync"`

	// A boolean flag specifies whether the invoked function must be validated.
	IsValidationEnabled *bool `mandatory:"false" json:"isValidationEnabled"`
}

// GetDescription returns Description
func (m UpdateInvokeFunctionDeployStageDetails) GetDescription() *string {
	return m.Description
}

// GetDisplayName returns DisplayName
func (m UpdateInvokeFunctionDeployStageDetails) GetDisplayName() *string {
	return m.DisplayName
}

// GetDeployStagePredecessorCollection returns DeployStagePredecessorCollection
func (m UpdateInvokeFunctionDeployStageDetails) GetDeployStagePredecessorCollection() *DeployStagePredecessorCollection {
	return m.DeployStagePredecessorCollection
}

// GetFreeformTags returns FreeformTags
func (m UpdateInvokeFunctionDeployStageDetails) GetFreeformTags() map[string]string {
	return m.FreeformTags
}

// GetDefinedTags returns DefinedTags
func (m UpdateInvokeFunctionDeployStageDetails) GetDefinedTags() map[string]map[string]interface{} {
	return m.DefinedTags
}

func (m UpdateInvokeFunctionDeployStageDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m UpdateInvokeFunctionDeployStageDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m UpdateInvokeFunctionDeployStageDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeUpdateInvokeFunctionDeployStageDetails UpdateInvokeFunctionDeployStageDetails
	s := struct {
		DiscriminatorParam string `json:"deployStageType"`
		MarshalTypeUpdateInvokeFunctionDeployStageDetails
	}{
		"INVOKE_FUNCTION",
		(MarshalTypeUpdateInvokeFunctionDeployStageDetails)(m),
	}

	return json.Marshal(&s)
}
