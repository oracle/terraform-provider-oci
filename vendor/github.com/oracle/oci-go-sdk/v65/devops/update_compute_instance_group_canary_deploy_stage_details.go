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

// UpdateComputeInstanceGroupCanaryDeployStageDetails Specifies the Instance Group Canary deployment stage.
type UpdateComputeInstanceGroupCanaryDeployStageDetails struct {

	// Optional description about the deployment stage.
	Description *string `mandatory:"false" json:"description"`

	// Deployment stage display name, which can be renamed and is not necessarily unique. Avoid entering confidential information.
	DisplayName *string `mandatory:"false" json:"displayName"`

	DeployStagePredecessorCollection *DeployStagePredecessorCollection `mandatory:"false" json:"deployStagePredecessorCollection"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.  See Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace. See Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// The OCID of the artifact that contains the deployment specification.
	DeploymentSpecDeployArtifactId *string `mandatory:"false" json:"deploymentSpecDeployArtifactId"`

	// The list of file artifact OCIDs to deploy.
	DeployArtifactIds []string `mandatory:"false" json:"deployArtifactIds"`

	RolloutPolicy ComputeInstanceGroupRolloutPolicy `mandatory:"false" json:"rolloutPolicy"`

	TestLoadBalancerConfig *LoadBalancerConfig `mandatory:"false" json:"testLoadBalancerConfig"`
}

// GetDescription returns Description
func (m UpdateComputeInstanceGroupCanaryDeployStageDetails) GetDescription() *string {
	return m.Description
}

// GetDisplayName returns DisplayName
func (m UpdateComputeInstanceGroupCanaryDeployStageDetails) GetDisplayName() *string {
	return m.DisplayName
}

// GetDeployStagePredecessorCollection returns DeployStagePredecessorCollection
func (m UpdateComputeInstanceGroupCanaryDeployStageDetails) GetDeployStagePredecessorCollection() *DeployStagePredecessorCollection {
	return m.DeployStagePredecessorCollection
}

// GetFreeformTags returns FreeformTags
func (m UpdateComputeInstanceGroupCanaryDeployStageDetails) GetFreeformTags() map[string]string {
	return m.FreeformTags
}

// GetDefinedTags returns DefinedTags
func (m UpdateComputeInstanceGroupCanaryDeployStageDetails) GetDefinedTags() map[string]map[string]interface{} {
	return m.DefinedTags
}

func (m UpdateComputeInstanceGroupCanaryDeployStageDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m UpdateComputeInstanceGroupCanaryDeployStageDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m UpdateComputeInstanceGroupCanaryDeployStageDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeUpdateComputeInstanceGroupCanaryDeployStageDetails UpdateComputeInstanceGroupCanaryDeployStageDetails
	s := struct {
		DiscriminatorParam string `json:"deployStageType"`
		MarshalTypeUpdateComputeInstanceGroupCanaryDeployStageDetails
	}{
		"COMPUTE_INSTANCE_GROUP_CANARY_DEPLOYMENT",
		(MarshalTypeUpdateComputeInstanceGroupCanaryDeployStageDetails)(m),
	}

	return json.Marshal(&s)
}

// UnmarshalJSON unmarshals from json
func (m *UpdateComputeInstanceGroupCanaryDeployStageDetails) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		Description                      *string                           `json:"description"`
		DisplayName                      *string                           `json:"displayName"`
		DeployStagePredecessorCollection *DeployStagePredecessorCollection `json:"deployStagePredecessorCollection"`
		FreeformTags                     map[string]string                 `json:"freeformTags"`
		DefinedTags                      map[string]map[string]interface{} `json:"definedTags"`
		DeploymentSpecDeployArtifactId   *string                           `json:"deploymentSpecDeployArtifactId"`
		DeployArtifactIds                []string                          `json:"deployArtifactIds"`
		RolloutPolicy                    computeinstancegrouprolloutpolicy `json:"rolloutPolicy"`
		TestLoadBalancerConfig           *LoadBalancerConfig               `json:"testLoadBalancerConfig"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.Description = model.Description

	m.DisplayName = model.DisplayName

	m.DeployStagePredecessorCollection = model.DeployStagePredecessorCollection

	m.FreeformTags = model.FreeformTags

	m.DefinedTags = model.DefinedTags

	m.DeploymentSpecDeployArtifactId = model.DeploymentSpecDeployArtifactId

	m.DeployArtifactIds = make([]string, len(model.DeployArtifactIds))
	copy(m.DeployArtifactIds, model.DeployArtifactIds)
	nn, e = model.RolloutPolicy.UnmarshalPolymorphicJSON(model.RolloutPolicy.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.RolloutPolicy = nn.(ComputeInstanceGroupRolloutPolicy)
	} else {
		m.RolloutPolicy = nil
	}

	m.TestLoadBalancerConfig = model.TestLoadBalancerConfig

	return
}
