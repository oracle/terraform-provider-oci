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

// CreateShellDeployStageDetails Specifies the shell stage.
type CreateShellDeployStageDetails struct {

	// The OCID of a pipeline.
	DeployPipelineId *string `mandatory:"true" json:"deployPipelineId"`

	DeployStagePredecessorCollection *DeployStagePredecessorCollection `mandatory:"true" json:"deployStagePredecessorCollection"`

	ContainerConfig ContainerConfig `mandatory:"true" json:"containerConfig"`

	// The OCID of the artifact that contains the command specification.
	CommandSpecDeployArtifactId *string `mandatory:"true" json:"commandSpecDeployArtifactId"`

	// Optional description about the deployment stage.
	Description *string `mandatory:"false" json:"description"`

	// Deployment stage display name, which can be renamed and is not necessarily unique. Avoid entering confidential information.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.  See Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace. See Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// Time to wait for execution of a shell stage. Defaults to 36000 seconds.
	TimeoutInSeconds *int `mandatory:"false" json:"timeoutInSeconds"`
}

// GetDescription returns Description
func (m CreateShellDeployStageDetails) GetDescription() *string {
	return m.Description
}

// GetDisplayName returns DisplayName
func (m CreateShellDeployStageDetails) GetDisplayName() *string {
	return m.DisplayName
}

// GetDeployPipelineId returns DeployPipelineId
func (m CreateShellDeployStageDetails) GetDeployPipelineId() *string {
	return m.DeployPipelineId
}

// GetDeployStagePredecessorCollection returns DeployStagePredecessorCollection
func (m CreateShellDeployStageDetails) GetDeployStagePredecessorCollection() *DeployStagePredecessorCollection {
	return m.DeployStagePredecessorCollection
}

// GetFreeformTags returns FreeformTags
func (m CreateShellDeployStageDetails) GetFreeformTags() map[string]string {
	return m.FreeformTags
}

// GetDefinedTags returns DefinedTags
func (m CreateShellDeployStageDetails) GetDefinedTags() map[string]map[string]interface{} {
	return m.DefinedTags
}

func (m CreateShellDeployStageDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CreateShellDeployStageDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m CreateShellDeployStageDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeCreateShellDeployStageDetails CreateShellDeployStageDetails
	s := struct {
		DiscriminatorParam string `json:"deployStageType"`
		MarshalTypeCreateShellDeployStageDetails
	}{
		"SHELL",
		(MarshalTypeCreateShellDeployStageDetails)(m),
	}

	return json.Marshal(&s)
}

// UnmarshalJSON unmarshals from json
func (m *CreateShellDeployStageDetails) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		Description                      *string                           `json:"description"`
		DisplayName                      *string                           `json:"displayName"`
		FreeformTags                     map[string]string                 `json:"freeformTags"`
		DefinedTags                      map[string]map[string]interface{} `json:"definedTags"`
		TimeoutInSeconds                 *int                              `json:"timeoutInSeconds"`
		DeployPipelineId                 *string                           `json:"deployPipelineId"`
		DeployStagePredecessorCollection *DeployStagePredecessorCollection `json:"deployStagePredecessorCollection"`
		ContainerConfig                  containerconfig                   `json:"containerConfig"`
		CommandSpecDeployArtifactId      *string                           `json:"commandSpecDeployArtifactId"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.Description = model.Description

	m.DisplayName = model.DisplayName

	m.FreeformTags = model.FreeformTags

	m.DefinedTags = model.DefinedTags

	m.TimeoutInSeconds = model.TimeoutInSeconds

	m.DeployPipelineId = model.DeployPipelineId

	m.DeployStagePredecessorCollection = model.DeployStagePredecessorCollection

	nn, e = model.ContainerConfig.UnmarshalPolymorphicJSON(model.ContainerConfig.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.ContainerConfig = nn.(ContainerConfig)
	} else {
		m.ContainerConfig = nil
	}

	m.CommandSpecDeployArtifactId = model.CommandSpecDeployArtifactId

	return
}
