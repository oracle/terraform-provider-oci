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

// DeployStageSummary Summary of the deployment stage.
type DeployStageSummary interface {

	// Unique identifier that is immutable on creation.
	GetId() *string

	// The OCID of a project.
	GetProjectId() *string

	// The OCID of a pipeline.
	GetDeployPipelineId() *string

	// The OCID of a compartment.
	GetCompartmentId() *string

	// Optional description about the deployment stage.
	GetDescription() *string

	// Deployment stage display name, which can be renamed and is not necessarily unique. Avoid entering confidential information.
	GetDisplayName() *string

	// Time the deployment stage was created. Format defined by RFC3339 (https://datatracker.ietf.org/doc/html/rfc3339).
	GetTimeCreated() *common.SDKTime

	// Time the deployment stage was updated. Format defined by RFC3339 (https://datatracker.ietf.org/doc/html/rfc3339).
	GetTimeUpdated() *common.SDKTime

	// The current state of the deployment stage.
	GetLifecycleState() DeployStageLifecycleStateEnum

	// A message describing the current state in more detail. For example, can be used to provide actionable information for a resource in Failed state.
	GetLifecycleDetails() *string

	GetDeployStagePredecessorCollection() *DeployStagePredecessorCollection

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.  See Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"bar-key": "value"}`
	GetFreeformTags() map[string]string

	// Defined tags for this resource. Each key is predefined and scoped to a namespace. See Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"foo-namespace": {"bar-key": "value"}}`
	GetDefinedTags() map[string]map[string]interface{}

	// Usage of system tag keys. These predefined keys are scoped to namespaces. See Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"orcl-cloud": {"free-tier-retained": "true"}}`
	GetSystemTags() map[string]map[string]interface{}
}

type deploystagesummary struct {
	JsonData                         []byte
	Description                      *string                           `mandatory:"false" json:"description"`
	DisplayName                      *string                           `mandatory:"false" json:"displayName"`
	TimeCreated                      *common.SDKTime                   `mandatory:"false" json:"timeCreated"`
	TimeUpdated                      *common.SDKTime                   `mandatory:"false" json:"timeUpdated"`
	LifecycleState                   DeployStageLifecycleStateEnum     `mandatory:"false" json:"lifecycleState,omitempty"`
	LifecycleDetails                 *string                           `mandatory:"false" json:"lifecycleDetails"`
	DeployStagePredecessorCollection *DeployStagePredecessorCollection `mandatory:"false" json:"deployStagePredecessorCollection"`
	FreeformTags                     map[string]string                 `mandatory:"false" json:"freeformTags"`
	DefinedTags                      map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
	SystemTags                       map[string]map[string]interface{} `mandatory:"false" json:"systemTags"`
	Id                               *string                           `mandatory:"true" json:"id"`
	ProjectId                        *string                           `mandatory:"true" json:"projectId"`
	DeployPipelineId                 *string                           `mandatory:"true" json:"deployPipelineId"`
	CompartmentId                    *string                           `mandatory:"true" json:"compartmentId"`
	DeployStageType                  string                            `json:"deployStageType"`
}

// UnmarshalJSON unmarshals json
func (m *deploystagesummary) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerdeploystagesummary deploystagesummary
	s := struct {
		Model Unmarshalerdeploystagesummary
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.Id = s.Model.Id
	m.ProjectId = s.Model.ProjectId
	m.DeployPipelineId = s.Model.DeployPipelineId
	m.CompartmentId = s.Model.CompartmentId
	m.Description = s.Model.Description
	m.DisplayName = s.Model.DisplayName
	m.TimeCreated = s.Model.TimeCreated
	m.TimeUpdated = s.Model.TimeUpdated
	m.LifecycleState = s.Model.LifecycleState
	m.LifecycleDetails = s.Model.LifecycleDetails
	m.DeployStagePredecessorCollection = s.Model.DeployStagePredecessorCollection
	m.FreeformTags = s.Model.FreeformTags
	m.DefinedTags = s.Model.DefinedTags
	m.SystemTags = s.Model.SystemTags
	m.DeployStageType = s.Model.DeployStageType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *deploystagesummary) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.DeployStageType {
	case "COMPUTE_INSTANCE_GROUP_BLUE_GREEN_DEPLOYMENT":
		mm := ComputeInstanceGroupBlueGreenDeployStageSummary{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "COMPUTE_INSTANCE_GROUP_BLUE_GREEN_TRAFFIC_SHIFT":
		mm := ComputeInstanceGroupBlueGreenTrafficShiftDeployStageSummary{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "OKE_BLUE_GREEN_DEPLOYMENT":
		mm := OkeBlueGreenDeployStageSummary{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "WAIT":
		mm := WaitDeployStageSummary{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "OKE_DEPLOYMENT":
		mm := OkeDeployStageSummary{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "COMPUTE_INSTANCE_GROUP_CANARY_APPROVAL":
		mm := ComputeInstanceGroupCanaryApprovalDeployStageSummary{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "INVOKE_FUNCTION":
		mm := InvokeFunctionDeployStageSummary{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "OKE_HELM_CHART_DEPLOYMENT":
		mm := OkeHelmChartDeployStageSummary{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "SHELL":
		mm := ShellDeployStageSummary{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "OKE_CANARY_TRAFFIC_SHIFT":
		mm := OkeCanaryTrafficShiftDeployStageSummary{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "COMPUTE_INSTANCE_GROUP_ROLLING_DEPLOYMENT":
		mm := ComputeInstanceGroupDeployStageSummary{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "COMPUTE_INSTANCE_GROUP_CANARY_DEPLOYMENT":
		mm := ComputeInstanceGroupCanaryDeployStageSummary{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "OKE_CANARY_APPROVAL":
		mm := OkeCanaryApprovalDeployStageSummary{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "OKE_CANARY_DEPLOYMENT":
		mm := OkeCanaryDeployStageSummary{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "LOAD_BALANCER_TRAFFIC_SHIFT":
		mm := LoadBalancerTrafficShiftDeployStageSummary{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "MANUAL_APPROVAL":
		mm := ManualApprovalDeployStageSummary{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "OKE_BLUE_GREEN_TRAFFIC_SHIFT":
		mm := OkeBlueGreenTrafficShiftDeployStageSummary{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "DEPLOY_FUNCTION":
		mm := FunctionDeployStageSummary{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "COMPUTE_INSTANCE_GROUP_CANARY_TRAFFIC_SHIFT":
		mm := ComputeInstanceGroupCanaryTrafficShiftDeployStageSummary{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Received unsupported enum value for DeployStageSummary: %s.", m.DeployStageType)
		return *m, nil
	}
}

// GetDescription returns Description
func (m deploystagesummary) GetDescription() *string {
	return m.Description
}

// GetDisplayName returns DisplayName
func (m deploystagesummary) GetDisplayName() *string {
	return m.DisplayName
}

// GetTimeCreated returns TimeCreated
func (m deploystagesummary) GetTimeCreated() *common.SDKTime {
	return m.TimeCreated
}

// GetTimeUpdated returns TimeUpdated
func (m deploystagesummary) GetTimeUpdated() *common.SDKTime {
	return m.TimeUpdated
}

// GetLifecycleState returns LifecycleState
func (m deploystagesummary) GetLifecycleState() DeployStageLifecycleStateEnum {
	return m.LifecycleState
}

// GetLifecycleDetails returns LifecycleDetails
func (m deploystagesummary) GetLifecycleDetails() *string {
	return m.LifecycleDetails
}

// GetDeployStagePredecessorCollection returns DeployStagePredecessorCollection
func (m deploystagesummary) GetDeployStagePredecessorCollection() *DeployStagePredecessorCollection {
	return m.DeployStagePredecessorCollection
}

// GetFreeformTags returns FreeformTags
func (m deploystagesummary) GetFreeformTags() map[string]string {
	return m.FreeformTags
}

// GetDefinedTags returns DefinedTags
func (m deploystagesummary) GetDefinedTags() map[string]map[string]interface{} {
	return m.DefinedTags
}

// GetSystemTags returns SystemTags
func (m deploystagesummary) GetSystemTags() map[string]map[string]interface{} {
	return m.SystemTags
}

// GetId returns Id
func (m deploystagesummary) GetId() *string {
	return m.Id
}

// GetProjectId returns ProjectId
func (m deploystagesummary) GetProjectId() *string {
	return m.ProjectId
}

// GetDeployPipelineId returns DeployPipelineId
func (m deploystagesummary) GetDeployPipelineId() *string {
	return m.DeployPipelineId
}

// GetCompartmentId returns CompartmentId
func (m deploystagesummary) GetCompartmentId() *string {
	return m.CompartmentId
}

func (m deploystagesummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m deploystagesummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingDeployStageLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetDeployStageLifecycleStateEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
