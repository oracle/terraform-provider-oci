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

// CreateDeployStageDetails The information about new deployment stage.
type CreateDeployStageDetails interface {

	// The OCID of a pipeline.
	GetDeployPipelineId() *string

	GetDeployStagePredecessorCollection() *DeployStagePredecessorCollection

	// Optional description about the deployment stage.
	GetDescription() *string

	// Deployment stage display name, which can be renamed and is not necessarily unique. Avoid entering confidential information.
	GetDisplayName() *string

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.  See Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm). Example: `{"bar-key": "value"}`
	GetFreeformTags() map[string]string

	// Defined tags for this resource. Each key is predefined and scoped to a namespace. See Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm). Example: `{"foo-namespace": {"bar-key": "value"}}`
	GetDefinedTags() map[string]map[string]interface{}
}

type createdeploystagedetails struct {
	JsonData                         []byte
	DeployPipelineId                 *string                           `mandatory:"true" json:"deployPipelineId"`
	DeployStagePredecessorCollection *DeployStagePredecessorCollection `mandatory:"true" json:"deployStagePredecessorCollection"`
	Description                      *string                           `mandatory:"false" json:"description"`
	DisplayName                      *string                           `mandatory:"false" json:"displayName"`
	FreeformTags                     map[string]string                 `mandatory:"false" json:"freeformTags"`
	DefinedTags                      map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
	DeployStageType                  string                            `json:"deployStageType"`
}

// UnmarshalJSON unmarshals json
func (m *createdeploystagedetails) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalercreatedeploystagedetails createdeploystagedetails
	s := struct {
		Model Unmarshalercreatedeploystagedetails
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.DeployPipelineId = s.Model.DeployPipelineId
	m.DeployStagePredecessorCollection = s.Model.DeployStagePredecessorCollection
	m.Description = s.Model.Description
	m.DisplayName = s.Model.DisplayName
	m.FreeformTags = s.Model.FreeformTags
	m.DefinedTags = s.Model.DefinedTags
	m.DeployStageType = s.Model.DeployStageType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *createdeploystagedetails) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.DeployStageType {
	case "MANUAL_APPROVAL":
		mm := CreateManualApprovalDeployStageDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "WAIT":
		mm := CreateWaitDeployStageDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "OKE_DEPLOYMENT":
		mm := CreateOkeDeployStageDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "LOAD_BALANCER_TRAFFIC_SHIFT":
		mm := CreateLoadBalancerTrafficShiftDeployStageDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "COMPUTE_INSTANCE_GROUP_ROLLING_DEPLOYMENT":
		mm := CreateComputeInstanceGroupDeployStageDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "INVOKE_FUNCTION":
		mm := CreateInvokeFunctionDeployStageDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "DEPLOY_FUNCTION":
		mm := CreateFunctionDeployStageDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		return *m, nil
	}
}

//GetDeployPipelineId returns DeployPipelineId
func (m createdeploystagedetails) GetDeployPipelineId() *string {
	return m.DeployPipelineId
}

//GetDeployStagePredecessorCollection returns DeployStagePredecessorCollection
func (m createdeploystagedetails) GetDeployStagePredecessorCollection() *DeployStagePredecessorCollection {
	return m.DeployStagePredecessorCollection
}

//GetDescription returns Description
func (m createdeploystagedetails) GetDescription() *string {
	return m.Description
}

//GetDisplayName returns DisplayName
func (m createdeploystagedetails) GetDisplayName() *string {
	return m.DisplayName
}

//GetFreeformTags returns FreeformTags
func (m createdeploystagedetails) GetFreeformTags() map[string]string {
	return m.FreeformTags
}

//GetDefinedTags returns DefinedTags
func (m createdeploystagedetails) GetDefinedTags() map[string]map[string]interface{} {
	return m.DefinedTags
}

func (m createdeploystagedetails) String() string {
	return common.PointerString(m)
}
