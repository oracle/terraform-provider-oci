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

// DeployStage A single node in a pipeline. It is usually associated with some action on a specific set of OCI resources such as environments. For example, updating a Function or a Kubernetes cluster.
type DeployStage interface {

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

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.  See Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm). Example: `{"bar-key": "value"}`
	GetFreeformTags() map[string]string

	// Defined tags for this resource. Each key is predefined and scoped to a namespace. See Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm). Example: `{"foo-namespace": {"bar-key": "value"}}`
	GetDefinedTags() map[string]map[string]interface{}

	// Usage of system tag keys. These predefined keys are scoped to namespaces. See Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm). Example: `{"orcl-cloud": {"free-tier-retained": "true"}}`
	GetSystemTags() map[string]map[string]interface{}
}

type deploystage struct {
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
func (m *deploystage) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerdeploystage deploystage
	s := struct {
		Model Unmarshalerdeploystage
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
func (m *deploystage) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.DeployStageType {
	case "COMPUTE_INSTANCE_GROUP_BLUE_GREEN_TRAFFIC_SHIFT":
		mm := ComputeInstanceGroupBlueGreenTrafficShiftDeployStage{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "LOAD_BALANCER_TRAFFIC_SHIFT":
		mm := LoadBalancerTrafficShiftDeployStage{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "SHELL":
		mm := ShellDeployStage{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "INVOKE_FUNCTION":
		mm := InvokeFunctionDeployStage{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "OKE_CANARY_DEPLOYMENT":
		mm := OkeCanaryDeployStage{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "OKE_DEPLOYMENT":
		mm := OkeDeployStage{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "OKE_CANARY_TRAFFIC_SHIFT":
		mm := OkeCanaryTrafficShiftDeployStage{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "DEPLOY_FUNCTION":
		mm := FunctionDeployStage{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "OKE_BLUE_GREEN_DEPLOYMENT":
		mm := OkeBlueGreenDeployStage{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "OKE_CANARY_APPROVAL":
		mm := OkeCanaryApprovalDeployStage{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "COMPUTE_INSTANCE_GROUP_CANARY_TRAFFIC_SHIFT":
		mm := ComputeInstanceGroupCanaryTrafficShiftDeployStage{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "COMPUTE_INSTANCE_GROUP_CANARY_APPROVAL":
		mm := ComputeInstanceGroupCanaryApprovalDeployStage{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "WAIT":
		mm := WaitDeployStage{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "OKE_HELM_CHART_DEPLOYMENT":
		mm := OkeHelmChartDeployStage{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "COMPUTE_INSTANCE_GROUP_BLUE_GREEN_DEPLOYMENT":
		mm := ComputeInstanceGroupBlueGreenDeployStage{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "COMPUTE_INSTANCE_GROUP_CANARY_DEPLOYMENT":
		mm := ComputeInstanceGroupCanaryDeployStage{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "OKE_BLUE_GREEN_TRAFFIC_SHIFT":
		mm := OkeBlueGreenTrafficShiftDeployStage{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "MANUAL_APPROVAL":
		mm := ManualApprovalDeployStage{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "COMPUTE_INSTANCE_GROUP_ROLLING_DEPLOYMENT":
		mm := ComputeInstanceGroupDeployStage{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Recieved unsupported enum value for DeployStage: %s.", m.DeployStageType)
		return *m, nil
	}
}

// GetDescription returns Description
func (m deploystage) GetDescription() *string {
	return m.Description
}

// GetDisplayName returns DisplayName
func (m deploystage) GetDisplayName() *string {
	return m.DisplayName
}

// GetTimeCreated returns TimeCreated
func (m deploystage) GetTimeCreated() *common.SDKTime {
	return m.TimeCreated
}

// GetTimeUpdated returns TimeUpdated
func (m deploystage) GetTimeUpdated() *common.SDKTime {
	return m.TimeUpdated
}

// GetLifecycleState returns LifecycleState
func (m deploystage) GetLifecycleState() DeployStageLifecycleStateEnum {
	return m.LifecycleState
}

// GetLifecycleDetails returns LifecycleDetails
func (m deploystage) GetLifecycleDetails() *string {
	return m.LifecycleDetails
}

// GetDeployStagePredecessorCollection returns DeployStagePredecessorCollection
func (m deploystage) GetDeployStagePredecessorCollection() *DeployStagePredecessorCollection {
	return m.DeployStagePredecessorCollection
}

// GetFreeformTags returns FreeformTags
func (m deploystage) GetFreeformTags() map[string]string {
	return m.FreeformTags
}

// GetDefinedTags returns DefinedTags
func (m deploystage) GetDefinedTags() map[string]map[string]interface{} {
	return m.DefinedTags
}

// GetSystemTags returns SystemTags
func (m deploystage) GetSystemTags() map[string]map[string]interface{} {
	return m.SystemTags
}

// GetId returns Id
func (m deploystage) GetId() *string {
	return m.Id
}

// GetProjectId returns ProjectId
func (m deploystage) GetProjectId() *string {
	return m.ProjectId
}

// GetDeployPipelineId returns DeployPipelineId
func (m deploystage) GetDeployPipelineId() *string {
	return m.DeployPipelineId
}

// GetCompartmentId returns CompartmentId
func (m deploystage) GetCompartmentId() *string {
	return m.CompartmentId
}

func (m deploystage) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m deploystage) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingDeployStageLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetDeployStageLifecycleStateEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// DeployStageLifecycleStateEnum Enum with underlying type: string
type DeployStageLifecycleStateEnum string

// Set of constants representing the allowable values for DeployStageLifecycleStateEnum
const (
	DeployStageLifecycleStateCreating DeployStageLifecycleStateEnum = "CREATING"
	DeployStageLifecycleStateUpdating DeployStageLifecycleStateEnum = "UPDATING"
	DeployStageLifecycleStateActive   DeployStageLifecycleStateEnum = "ACTIVE"
	DeployStageLifecycleStateDeleting DeployStageLifecycleStateEnum = "DELETING"
	DeployStageLifecycleStateDeleted  DeployStageLifecycleStateEnum = "DELETED"
	DeployStageLifecycleStateFailed   DeployStageLifecycleStateEnum = "FAILED"
)

var mappingDeployStageLifecycleStateEnum = map[string]DeployStageLifecycleStateEnum{
	"CREATING": DeployStageLifecycleStateCreating,
	"UPDATING": DeployStageLifecycleStateUpdating,
	"ACTIVE":   DeployStageLifecycleStateActive,
	"DELETING": DeployStageLifecycleStateDeleting,
	"DELETED":  DeployStageLifecycleStateDeleted,
	"FAILED":   DeployStageLifecycleStateFailed,
}

var mappingDeployStageLifecycleStateEnumLowerCase = map[string]DeployStageLifecycleStateEnum{
	"creating": DeployStageLifecycleStateCreating,
	"updating": DeployStageLifecycleStateUpdating,
	"active":   DeployStageLifecycleStateActive,
	"deleting": DeployStageLifecycleStateDeleting,
	"deleted":  DeployStageLifecycleStateDeleted,
	"failed":   DeployStageLifecycleStateFailed,
}

// GetDeployStageLifecycleStateEnumValues Enumerates the set of values for DeployStageLifecycleStateEnum
func GetDeployStageLifecycleStateEnumValues() []DeployStageLifecycleStateEnum {
	values := make([]DeployStageLifecycleStateEnum, 0)
	for _, v := range mappingDeployStageLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetDeployStageLifecycleStateEnumStringValues Enumerates the set of values in String for DeployStageLifecycleStateEnum
func GetDeployStageLifecycleStateEnumStringValues() []string {
	return []string{
		"CREATING",
		"UPDATING",
		"ACTIVE",
		"DELETING",
		"DELETED",
		"FAILED",
	}
}

// GetMappingDeployStageLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDeployStageLifecycleStateEnum(val string) (DeployStageLifecycleStateEnum, bool) {
	enum, ok := mappingDeployStageLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// DeployStageDeployStageTypeEnum Enum with underlying type: string
type DeployStageDeployStageTypeEnum string

// Set of constants representing the allowable values for DeployStageDeployStageTypeEnum
const (
	DeployStageDeployStageTypeWait                                      DeployStageDeployStageTypeEnum = "WAIT"
	DeployStageDeployStageTypeComputeInstanceGroupRollingDeployment     DeployStageDeployStageTypeEnum = "COMPUTE_INSTANCE_GROUP_ROLLING_DEPLOYMENT"
	DeployStageDeployStageTypeComputeInstanceGroupBlueGreenDeployment   DeployStageDeployStageTypeEnum = "COMPUTE_INSTANCE_GROUP_BLUE_GREEN_DEPLOYMENT"
	DeployStageDeployStageTypeComputeInstanceGroupBlueGreenTrafficShift DeployStageDeployStageTypeEnum = "COMPUTE_INSTANCE_GROUP_BLUE_GREEN_TRAFFIC_SHIFT"
	DeployStageDeployStageTypeComputeInstanceGroupCanaryDeployment      DeployStageDeployStageTypeEnum = "COMPUTE_INSTANCE_GROUP_CANARY_DEPLOYMENT"
	DeployStageDeployStageTypeComputeInstanceGroupCanaryTrafficShift    DeployStageDeployStageTypeEnum = "COMPUTE_INSTANCE_GROUP_CANARY_TRAFFIC_SHIFT"
	DeployStageDeployStageTypeComputeInstanceGroupCanaryApproval        DeployStageDeployStageTypeEnum = "COMPUTE_INSTANCE_GROUP_CANARY_APPROVAL"
	DeployStageDeployStageTypeOkeBlueGreenDeployment                    DeployStageDeployStageTypeEnum = "OKE_BLUE_GREEN_DEPLOYMENT"
	DeployStageDeployStageTypeOkeBlueGreenTrafficShift                  DeployStageDeployStageTypeEnum = "OKE_BLUE_GREEN_TRAFFIC_SHIFT"
	DeployStageDeployStageTypeOkeCanaryDeployment                       DeployStageDeployStageTypeEnum = "OKE_CANARY_DEPLOYMENT"
	DeployStageDeployStageTypeOkeCanaryTrafficShift                     DeployStageDeployStageTypeEnum = "OKE_CANARY_TRAFFIC_SHIFT"
	DeployStageDeployStageTypeOkeCanaryApproval                         DeployStageDeployStageTypeEnum = "OKE_CANARY_APPROVAL"
	DeployStageDeployStageTypeOkeDeployment                             DeployStageDeployStageTypeEnum = "OKE_DEPLOYMENT"
	DeployStageDeployStageTypeDeployFunction                            DeployStageDeployStageTypeEnum = "DEPLOY_FUNCTION"
	DeployStageDeployStageTypeInvokeFunction                            DeployStageDeployStageTypeEnum = "INVOKE_FUNCTION"
	DeployStageDeployStageTypeLoadBalancerTrafficShift                  DeployStageDeployStageTypeEnum = "LOAD_BALANCER_TRAFFIC_SHIFT"
	DeployStageDeployStageTypeManualApproval                            DeployStageDeployStageTypeEnum = "MANUAL_APPROVAL"
	DeployStageDeployStageTypeOkeHelmChartDeployment                    DeployStageDeployStageTypeEnum = "OKE_HELM_CHART_DEPLOYMENT"
	DeployStageDeployStageTypeShell                                     DeployStageDeployStageTypeEnum = "SHELL"
)

var mappingDeployStageDeployStageTypeEnum = map[string]DeployStageDeployStageTypeEnum{
	"WAIT": DeployStageDeployStageTypeWait,
	"COMPUTE_INSTANCE_GROUP_ROLLING_DEPLOYMENT":       DeployStageDeployStageTypeComputeInstanceGroupRollingDeployment,
	"COMPUTE_INSTANCE_GROUP_BLUE_GREEN_DEPLOYMENT":    DeployStageDeployStageTypeComputeInstanceGroupBlueGreenDeployment,
	"COMPUTE_INSTANCE_GROUP_BLUE_GREEN_TRAFFIC_SHIFT": DeployStageDeployStageTypeComputeInstanceGroupBlueGreenTrafficShift,
	"COMPUTE_INSTANCE_GROUP_CANARY_DEPLOYMENT":        DeployStageDeployStageTypeComputeInstanceGroupCanaryDeployment,
	"COMPUTE_INSTANCE_GROUP_CANARY_TRAFFIC_SHIFT":     DeployStageDeployStageTypeComputeInstanceGroupCanaryTrafficShift,
	"COMPUTE_INSTANCE_GROUP_CANARY_APPROVAL":          DeployStageDeployStageTypeComputeInstanceGroupCanaryApproval,
	"OKE_BLUE_GREEN_DEPLOYMENT":                       DeployStageDeployStageTypeOkeBlueGreenDeployment,
	"OKE_BLUE_GREEN_TRAFFIC_SHIFT":                    DeployStageDeployStageTypeOkeBlueGreenTrafficShift,
	"OKE_CANARY_DEPLOYMENT":                           DeployStageDeployStageTypeOkeCanaryDeployment,
	"OKE_CANARY_TRAFFIC_SHIFT":                        DeployStageDeployStageTypeOkeCanaryTrafficShift,
	"OKE_CANARY_APPROVAL":                             DeployStageDeployStageTypeOkeCanaryApproval,
	"OKE_DEPLOYMENT":                                  DeployStageDeployStageTypeOkeDeployment,
	"DEPLOY_FUNCTION":                                 DeployStageDeployStageTypeDeployFunction,
	"INVOKE_FUNCTION":                                 DeployStageDeployStageTypeInvokeFunction,
	"LOAD_BALANCER_TRAFFIC_SHIFT":                     DeployStageDeployStageTypeLoadBalancerTrafficShift,
	"MANUAL_APPROVAL":                                 DeployStageDeployStageTypeManualApproval,
	"OKE_HELM_CHART_DEPLOYMENT":                       DeployStageDeployStageTypeOkeHelmChartDeployment,
	"SHELL":                                           DeployStageDeployStageTypeShell,
}

var mappingDeployStageDeployStageTypeEnumLowerCase = map[string]DeployStageDeployStageTypeEnum{
	"wait": DeployStageDeployStageTypeWait,
	"compute_instance_group_rolling_deployment":       DeployStageDeployStageTypeComputeInstanceGroupRollingDeployment,
	"compute_instance_group_blue_green_deployment":    DeployStageDeployStageTypeComputeInstanceGroupBlueGreenDeployment,
	"compute_instance_group_blue_green_traffic_shift": DeployStageDeployStageTypeComputeInstanceGroupBlueGreenTrafficShift,
	"compute_instance_group_canary_deployment":        DeployStageDeployStageTypeComputeInstanceGroupCanaryDeployment,
	"compute_instance_group_canary_traffic_shift":     DeployStageDeployStageTypeComputeInstanceGroupCanaryTrafficShift,
	"compute_instance_group_canary_approval":          DeployStageDeployStageTypeComputeInstanceGroupCanaryApproval,
	"oke_blue_green_deployment":                       DeployStageDeployStageTypeOkeBlueGreenDeployment,
	"oke_blue_green_traffic_shift":                    DeployStageDeployStageTypeOkeBlueGreenTrafficShift,
	"oke_canary_deployment":                           DeployStageDeployStageTypeOkeCanaryDeployment,
	"oke_canary_traffic_shift":                        DeployStageDeployStageTypeOkeCanaryTrafficShift,
	"oke_canary_approval":                             DeployStageDeployStageTypeOkeCanaryApproval,
	"oke_deployment":                                  DeployStageDeployStageTypeOkeDeployment,
	"deploy_function":                                 DeployStageDeployStageTypeDeployFunction,
	"invoke_function":                                 DeployStageDeployStageTypeInvokeFunction,
	"load_balancer_traffic_shift":                     DeployStageDeployStageTypeLoadBalancerTrafficShift,
	"manual_approval":                                 DeployStageDeployStageTypeManualApproval,
	"oke_helm_chart_deployment":                       DeployStageDeployStageTypeOkeHelmChartDeployment,
	"shell":                                           DeployStageDeployStageTypeShell,
}

// GetDeployStageDeployStageTypeEnumValues Enumerates the set of values for DeployStageDeployStageTypeEnum
func GetDeployStageDeployStageTypeEnumValues() []DeployStageDeployStageTypeEnum {
	values := make([]DeployStageDeployStageTypeEnum, 0)
	for _, v := range mappingDeployStageDeployStageTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetDeployStageDeployStageTypeEnumStringValues Enumerates the set of values in String for DeployStageDeployStageTypeEnum
func GetDeployStageDeployStageTypeEnumStringValues() []string {
	return []string{
		"WAIT",
		"COMPUTE_INSTANCE_GROUP_ROLLING_DEPLOYMENT",
		"COMPUTE_INSTANCE_GROUP_BLUE_GREEN_DEPLOYMENT",
		"COMPUTE_INSTANCE_GROUP_BLUE_GREEN_TRAFFIC_SHIFT",
		"COMPUTE_INSTANCE_GROUP_CANARY_DEPLOYMENT",
		"COMPUTE_INSTANCE_GROUP_CANARY_TRAFFIC_SHIFT",
		"COMPUTE_INSTANCE_GROUP_CANARY_APPROVAL",
		"OKE_BLUE_GREEN_DEPLOYMENT",
		"OKE_BLUE_GREEN_TRAFFIC_SHIFT",
		"OKE_CANARY_DEPLOYMENT",
		"OKE_CANARY_TRAFFIC_SHIFT",
		"OKE_CANARY_APPROVAL",
		"OKE_DEPLOYMENT",
		"DEPLOY_FUNCTION",
		"INVOKE_FUNCTION",
		"LOAD_BALANCER_TRAFFIC_SHIFT",
		"MANUAL_APPROVAL",
		"OKE_HELM_CHART_DEPLOYMENT",
		"SHELL",
	}
}

// GetMappingDeployStageDeployStageTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDeployStageDeployStageTypeEnum(val string) (DeployStageDeployStageTypeEnum, bool) {
	enum, ok := mappingDeployStageDeployStageTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
