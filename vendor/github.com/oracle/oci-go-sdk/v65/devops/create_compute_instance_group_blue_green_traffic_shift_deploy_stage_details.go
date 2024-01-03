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

// CreateComputeInstanceGroupBlueGreenTrafficShiftDeployStageDetails Specifies the instance group blue-green deployment load balancer traffic shift stage.
type CreateComputeInstanceGroupBlueGreenTrafficShiftDeployStageDetails struct {

	// The OCID of a pipeline.
	DeployPipelineId *string `mandatory:"true" json:"deployPipelineId"`

	DeployStagePredecessorCollection *DeployStagePredecessorCollection `mandatory:"true" json:"deployStagePredecessorCollection"`

	// The OCID of the upstream compute instance group blue-green deployment stage in this pipeline.
	ComputeInstanceGroupBlueGreenDeploymentDeployStageId *string `mandatory:"true" json:"computeInstanceGroupBlueGreenDeploymentDeployStageId"`

	// Optional description about the deployment stage.
	Description *string `mandatory:"false" json:"description"`

	// Deployment stage display name, which can be renamed and is not necessarily unique. Avoid entering confidential information.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.  See Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm). Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace. See Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm). Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

// GetDescription returns Description
func (m CreateComputeInstanceGroupBlueGreenTrafficShiftDeployStageDetails) GetDescription() *string {
	return m.Description
}

// GetDisplayName returns DisplayName
func (m CreateComputeInstanceGroupBlueGreenTrafficShiftDeployStageDetails) GetDisplayName() *string {
	return m.DisplayName
}

// GetDeployPipelineId returns DeployPipelineId
func (m CreateComputeInstanceGroupBlueGreenTrafficShiftDeployStageDetails) GetDeployPipelineId() *string {
	return m.DeployPipelineId
}

// GetDeployStagePredecessorCollection returns DeployStagePredecessorCollection
func (m CreateComputeInstanceGroupBlueGreenTrafficShiftDeployStageDetails) GetDeployStagePredecessorCollection() *DeployStagePredecessorCollection {
	return m.DeployStagePredecessorCollection
}

// GetFreeformTags returns FreeformTags
func (m CreateComputeInstanceGroupBlueGreenTrafficShiftDeployStageDetails) GetFreeformTags() map[string]string {
	return m.FreeformTags
}

// GetDefinedTags returns DefinedTags
func (m CreateComputeInstanceGroupBlueGreenTrafficShiftDeployStageDetails) GetDefinedTags() map[string]map[string]interface{} {
	return m.DefinedTags
}

func (m CreateComputeInstanceGroupBlueGreenTrafficShiftDeployStageDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CreateComputeInstanceGroupBlueGreenTrafficShiftDeployStageDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m CreateComputeInstanceGroupBlueGreenTrafficShiftDeployStageDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeCreateComputeInstanceGroupBlueGreenTrafficShiftDeployStageDetails CreateComputeInstanceGroupBlueGreenTrafficShiftDeployStageDetails
	s := struct {
		DiscriminatorParam string `json:"deployStageType"`
		MarshalTypeCreateComputeInstanceGroupBlueGreenTrafficShiftDeployStageDetails
	}{
		"COMPUTE_INSTANCE_GROUP_BLUE_GREEN_TRAFFIC_SHIFT",
		(MarshalTypeCreateComputeInstanceGroupBlueGreenTrafficShiftDeployStageDetails)(m),
	}

	return json.Marshal(&s)
}
