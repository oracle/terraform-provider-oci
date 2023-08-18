// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
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

// CreateOkeBlueGreenDeployStageDetails Specifies the Container Engine for Kubernetes (OKE) cluster Blue-Green deployment stage.
type CreateOkeBlueGreenDeployStageDetails struct {

	// The OCID of a pipeline.
	DeployPipelineId *string `mandatory:"true" json:"deployPipelineId"`

	DeployStagePredecessorCollection *DeployStagePredecessorCollection `mandatory:"true" json:"deployStagePredecessorCollection"`

	// Kubernetes cluster environment OCID for deployment.
	OkeClusterDeployEnvironmentId *string `mandatory:"true" json:"okeClusterDeployEnvironmentId"`

	// List of Kubernetes manifest artifact OCIDs.
	KubernetesManifestDeployArtifactIds []string `mandatory:"true" json:"kubernetesManifestDeployArtifactIds"`

	BlueGreenStrategy OkeBlueGreenStrategy `mandatory:"true" json:"blueGreenStrategy"`

	// Optional description about the deployment stage.
	Description *string `mandatory:"false" json:"description"`

	// Deployment stage display name, which can be renamed and is not necessarily unique. Avoid entering confidential information.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.  See Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm). Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace. See Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm). Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

//GetDescription returns Description
func (m CreateOkeBlueGreenDeployStageDetails) GetDescription() *string {
	return m.Description
}

//GetDisplayName returns DisplayName
func (m CreateOkeBlueGreenDeployStageDetails) GetDisplayName() *string {
	return m.DisplayName
}

//GetDeployPipelineId returns DeployPipelineId
func (m CreateOkeBlueGreenDeployStageDetails) GetDeployPipelineId() *string {
	return m.DeployPipelineId
}

//GetDeployStagePredecessorCollection returns DeployStagePredecessorCollection
func (m CreateOkeBlueGreenDeployStageDetails) GetDeployStagePredecessorCollection() *DeployStagePredecessorCollection {
	return m.DeployStagePredecessorCollection
}

//GetFreeformTags returns FreeformTags
func (m CreateOkeBlueGreenDeployStageDetails) GetFreeformTags() map[string]string {
	return m.FreeformTags
}

//GetDefinedTags returns DefinedTags
func (m CreateOkeBlueGreenDeployStageDetails) GetDefinedTags() map[string]map[string]interface{} {
	return m.DefinedTags
}

func (m CreateOkeBlueGreenDeployStageDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CreateOkeBlueGreenDeployStageDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m CreateOkeBlueGreenDeployStageDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeCreateOkeBlueGreenDeployStageDetails CreateOkeBlueGreenDeployStageDetails
	s := struct {
		DiscriminatorParam string `json:"deployStageType"`
		MarshalTypeCreateOkeBlueGreenDeployStageDetails
	}{
		"OKE_BLUE_GREEN_DEPLOYMENT",
		(MarshalTypeCreateOkeBlueGreenDeployStageDetails)(m),
	}

	return json.Marshal(&s)
}

// UnmarshalJSON unmarshals from json
func (m *CreateOkeBlueGreenDeployStageDetails) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		Description                         *string                           `json:"description"`
		DisplayName                         *string                           `json:"displayName"`
		FreeformTags                        map[string]string                 `json:"freeformTags"`
		DefinedTags                         map[string]map[string]interface{} `json:"definedTags"`
		DeployPipelineId                    *string                           `json:"deployPipelineId"`
		DeployStagePredecessorCollection    *DeployStagePredecessorCollection `json:"deployStagePredecessorCollection"`
		OkeClusterDeployEnvironmentId       *string                           `json:"okeClusterDeployEnvironmentId"`
		KubernetesManifestDeployArtifactIds []string                          `json:"kubernetesManifestDeployArtifactIds"`
		BlueGreenStrategy                   okebluegreenstrategy              `json:"blueGreenStrategy"`
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

	m.DeployPipelineId = model.DeployPipelineId

	m.DeployStagePredecessorCollection = model.DeployStagePredecessorCollection

	m.OkeClusterDeployEnvironmentId = model.OkeClusterDeployEnvironmentId

	m.KubernetesManifestDeployArtifactIds = make([]string, len(model.KubernetesManifestDeployArtifactIds))
	copy(model.KubernetesManifestDeployArtifactIds, m.KubernetesManifestDeployArtifactIds)
	nn, e = model.BlueGreenStrategy.UnmarshalPolymorphicJSON(model.BlueGreenStrategy.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.BlueGreenStrategy = nn.(OkeBlueGreenStrategy)
	} else {
		m.BlueGreenStrategy = nil
	}

	return
}
