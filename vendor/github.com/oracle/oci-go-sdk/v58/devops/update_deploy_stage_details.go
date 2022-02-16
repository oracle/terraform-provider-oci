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
	"fmt"
	"github.com/oracle/oci-go-sdk/v58/common"
	"strings"
)

// UpdateDeployStageDetails The information to be updated.
type UpdateDeployStageDetails interface {

	// Optional description about the deployment stage.
	GetDescription() *string

	// Deployment stage display name, which can be renamed and is not necessarily unique. Avoid entering confidential information.
	GetDisplayName() *string

	GetDeployStagePredecessorCollection() *DeployStagePredecessorCollection

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.  See Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm). Example: `{"bar-key": "value"}`
	GetFreeformTags() map[string]string

	// Defined tags for this resource. Each key is predefined and scoped to a namespace. See Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm). Example: `{"foo-namespace": {"bar-key": "value"}}`
	GetDefinedTags() map[string]map[string]interface{}
}

type updatedeploystagedetails struct {
	JsonData                         []byte
	Description                      *string                           `mandatory:"false" json:"description"`
	DisplayName                      *string                           `mandatory:"false" json:"displayName"`
	DeployStagePredecessorCollection *DeployStagePredecessorCollection `mandatory:"false" json:"deployStagePredecessorCollection"`
	FreeformTags                     map[string]string                 `mandatory:"false" json:"freeformTags"`
	DefinedTags                      map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
	DeployStageType                  string                            `json:"deployStageType"`
}

// UnmarshalJSON unmarshals json
func (m *updatedeploystagedetails) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerupdatedeploystagedetails updatedeploystagedetails
	s := struct {
		Model Unmarshalerupdatedeploystagedetails
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.Description = s.Model.Description
	m.DisplayName = s.Model.DisplayName
	m.DeployStagePredecessorCollection = s.Model.DeployStagePredecessorCollection
	m.FreeformTags = s.Model.FreeformTags
	m.DefinedTags = s.Model.DefinedTags
	m.DeployStageType = s.Model.DeployStageType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *updatedeploystagedetails) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.DeployStageType {
	case "OKE_DEPLOYMENT":
		mm := UpdateOkeDeployStageDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "LOAD_BALANCER_TRAFFIC_SHIFT":
		mm := UpdateLoadBalancerTrafficShiftDeployStageDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "COMPUTE_INSTANCE_GROUP_ROLLING_DEPLOYMENT":
		mm := UpdateComputeInstanceGroupDeployStageDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "WAIT":
		mm := UpdateWaitDeployStageDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "MANUAL_APPROVAL":
		mm := UpdateManualApprovalDeployStageDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "DEPLOY_FUNCTION":
		mm := UpdateFunctionDeployStageDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "INVOKE_FUNCTION":
		mm := UpdateInvokeFunctionDeployStageDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		return *m, nil
	}
}

//GetDescription returns Description
func (m updatedeploystagedetails) GetDescription() *string {
	return m.Description
}

//GetDisplayName returns DisplayName
func (m updatedeploystagedetails) GetDisplayName() *string {
	return m.DisplayName
}

//GetDeployStagePredecessorCollection returns DeployStagePredecessorCollection
func (m updatedeploystagedetails) GetDeployStagePredecessorCollection() *DeployStagePredecessorCollection {
	return m.DeployStagePredecessorCollection
}

//GetFreeformTags returns FreeformTags
func (m updatedeploystagedetails) GetFreeformTags() map[string]string {
	return m.FreeformTags
}

//GetDefinedTags returns DefinedTags
func (m updatedeploystagedetails) GetDefinedTags() map[string]map[string]interface{} {
	return m.DefinedTags
}

func (m updatedeploystagedetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m updatedeploystagedetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
