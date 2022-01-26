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

// Deployment A single execution or run of a pipeline.
type Deployment interface {

	// Unique identifier that is immutable on creation.
	GetId() *string

	// The OCID of a project.
	GetProjectId() *string

	// The OCID of a pipeline.
	GetDeployPipelineId() *string

	// The OCID of a compartment.
	GetCompartmentId() *string

	GetDeployPipelineArtifacts() *DeployPipelineArtifactCollection

	GetDeployPipelineEnvironments() *DeployPipelineEnvironmentCollection

	// Deployment identifier which can be renamed and is not necessarily unique. Avoid entering confidential information.
	GetDisplayName() *string

	// Time the deployment was created. Format defined by RFC3339 (https://datatracker.ietf.org/doc/html/rfc3339).
	GetTimeCreated() *common.SDKTime

	// Time the deployment was updated. Format defined by RFC3339 (https://datatracker.ietf.org/doc/html/rfc3339).
	GetTimeUpdated() *common.SDKTime

	// The current state of the deployment.
	GetLifecycleState() DeploymentLifecycleStateEnum

	// A message describing the current state in more detail. For example, can be used to provide actionable information for a resource in Failed state.
	GetLifecycleDetails() *string

	GetDeploymentArguments() *DeploymentArgumentCollection

	GetDeployArtifactOverrideArguments() *DeployArtifactOverrideArgumentCollection

	GetDeploymentExecutionProgress() *DeploymentExecutionProgress

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.  See Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm). Example: `{"bar-key": "value"}`
	GetFreeformTags() map[string]string

	// Defined tags for this resource. Each key is predefined and scoped to a namespace. See Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm). Example: `{"foo-namespace": {"bar-key": "value"}}`
	GetDefinedTags() map[string]map[string]interface{}

	// Usage of system tag keys. These predefined keys are scoped to namespaces. See Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm). Example: `{"orcl-cloud": {"free-tier-retained": "true"}}`
	GetSystemTags() map[string]map[string]interface{}
}

type deployment struct {
	JsonData                        []byte
	Id                              *string                                   `mandatory:"true" json:"id"`
	ProjectId                       *string                                   `mandatory:"true" json:"projectId"`
	DeployPipelineId                *string                                   `mandatory:"true" json:"deployPipelineId"`
	CompartmentId                   *string                                   `mandatory:"true" json:"compartmentId"`
	DeployPipelineArtifacts         *DeployPipelineArtifactCollection         `mandatory:"false" json:"deployPipelineArtifacts"`
	DeployPipelineEnvironments      *DeployPipelineEnvironmentCollection      `mandatory:"false" json:"deployPipelineEnvironments"`
	DisplayName                     *string                                   `mandatory:"false" json:"displayName"`
	TimeCreated                     *common.SDKTime                           `mandatory:"false" json:"timeCreated"`
	TimeUpdated                     *common.SDKTime                           `mandatory:"false" json:"timeUpdated"`
	LifecycleState                  DeploymentLifecycleStateEnum              `mandatory:"false" json:"lifecycleState,omitempty"`
	LifecycleDetails                *string                                   `mandatory:"false" json:"lifecycleDetails"`
	DeploymentArguments             *DeploymentArgumentCollection             `mandatory:"false" json:"deploymentArguments"`
	DeployArtifactOverrideArguments *DeployArtifactOverrideArgumentCollection `mandatory:"false" json:"deployArtifactOverrideArguments"`
	DeploymentExecutionProgress     *DeploymentExecutionProgress              `mandatory:"false" json:"deploymentExecutionProgress"`
	FreeformTags                    map[string]string                         `mandatory:"false" json:"freeformTags"`
	DefinedTags                     map[string]map[string]interface{}         `mandatory:"false" json:"definedTags"`
	SystemTags                      map[string]map[string]interface{}         `mandatory:"false" json:"systemTags"`
	DeploymentType                  string                                    `json:"deploymentType"`
}

// UnmarshalJSON unmarshals json
func (m *deployment) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerdeployment deployment
	s := struct {
		Model Unmarshalerdeployment
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.Id = s.Model.Id
	m.ProjectId = s.Model.ProjectId
	m.DeployPipelineId = s.Model.DeployPipelineId
	m.CompartmentId = s.Model.CompartmentId
	m.DeployPipelineArtifacts = s.Model.DeployPipelineArtifacts
	m.DeployPipelineEnvironments = s.Model.DeployPipelineEnvironments
	m.DisplayName = s.Model.DisplayName
	m.TimeCreated = s.Model.TimeCreated
	m.TimeUpdated = s.Model.TimeUpdated
	m.LifecycleState = s.Model.LifecycleState
	m.LifecycleDetails = s.Model.LifecycleDetails
	m.DeploymentArguments = s.Model.DeploymentArguments
	m.DeployArtifactOverrideArguments = s.Model.DeployArtifactOverrideArguments
	m.DeploymentExecutionProgress = s.Model.DeploymentExecutionProgress
	m.FreeformTags = s.Model.FreeformTags
	m.DefinedTags = s.Model.DefinedTags
	m.SystemTags = s.Model.SystemTags
	m.DeploymentType = s.Model.DeploymentType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *deployment) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.DeploymentType {
	case "PIPELINE_DEPLOYMENT":
		mm := DeployPipelineDeployment{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "PIPELINE_REDEPLOYMENT":
		mm := DeployPipelineRedeployment{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "SINGLE_STAGE_DEPLOYMENT":
		mm := SingleDeployStageDeployment{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		return *m, nil
	}
}

//GetId returns Id
func (m deployment) GetId() *string {
	return m.Id
}

//GetProjectId returns ProjectId
func (m deployment) GetProjectId() *string {
	return m.ProjectId
}

//GetDeployPipelineId returns DeployPipelineId
func (m deployment) GetDeployPipelineId() *string {
	return m.DeployPipelineId
}

//GetCompartmentId returns CompartmentId
func (m deployment) GetCompartmentId() *string {
	return m.CompartmentId
}

//GetDeployPipelineArtifacts returns DeployPipelineArtifacts
func (m deployment) GetDeployPipelineArtifacts() *DeployPipelineArtifactCollection {
	return m.DeployPipelineArtifacts
}

//GetDeployPipelineEnvironments returns DeployPipelineEnvironments
func (m deployment) GetDeployPipelineEnvironments() *DeployPipelineEnvironmentCollection {
	return m.DeployPipelineEnvironments
}

//GetDisplayName returns DisplayName
func (m deployment) GetDisplayName() *string {
	return m.DisplayName
}

//GetTimeCreated returns TimeCreated
func (m deployment) GetTimeCreated() *common.SDKTime {
	return m.TimeCreated
}

//GetTimeUpdated returns TimeUpdated
func (m deployment) GetTimeUpdated() *common.SDKTime {
	return m.TimeUpdated
}

//GetLifecycleState returns LifecycleState
func (m deployment) GetLifecycleState() DeploymentLifecycleStateEnum {
	return m.LifecycleState
}

//GetLifecycleDetails returns LifecycleDetails
func (m deployment) GetLifecycleDetails() *string {
	return m.LifecycleDetails
}

//GetDeploymentArguments returns DeploymentArguments
func (m deployment) GetDeploymentArguments() *DeploymentArgumentCollection {
	return m.DeploymentArguments
}

//GetDeployArtifactOverrideArguments returns DeployArtifactOverrideArguments
func (m deployment) GetDeployArtifactOverrideArguments() *DeployArtifactOverrideArgumentCollection {
	return m.DeployArtifactOverrideArguments
}

//GetDeploymentExecutionProgress returns DeploymentExecutionProgress
func (m deployment) GetDeploymentExecutionProgress() *DeploymentExecutionProgress {
	return m.DeploymentExecutionProgress
}

//GetFreeformTags returns FreeformTags
func (m deployment) GetFreeformTags() map[string]string {
	return m.FreeformTags
}

//GetDefinedTags returns DefinedTags
func (m deployment) GetDefinedTags() map[string]map[string]interface{} {
	return m.DefinedTags
}

//GetSystemTags returns SystemTags
func (m deployment) GetSystemTags() map[string]map[string]interface{} {
	return m.SystemTags
}

func (m deployment) String() string {
	return common.PointerString(m)
}

// DeploymentLifecycleStateEnum Enum with underlying type: string
type DeploymentLifecycleStateEnum string

// Set of constants representing the allowable values for DeploymentLifecycleStateEnum
const (
	DeploymentLifecycleStateAccepted   DeploymentLifecycleStateEnum = "ACCEPTED"
	DeploymentLifecycleStateInProgress DeploymentLifecycleStateEnum = "IN_PROGRESS"
	DeploymentLifecycleStateFailed     DeploymentLifecycleStateEnum = "FAILED"
	DeploymentLifecycleStateSucceeded  DeploymentLifecycleStateEnum = "SUCCEEDED"
	DeploymentLifecycleStateCanceling  DeploymentLifecycleStateEnum = "CANCELING"
	DeploymentLifecycleStateCanceled   DeploymentLifecycleStateEnum = "CANCELED"
)

var mappingDeploymentLifecycleState = map[string]DeploymentLifecycleStateEnum{
	"ACCEPTED":    DeploymentLifecycleStateAccepted,
	"IN_PROGRESS": DeploymentLifecycleStateInProgress,
	"FAILED":      DeploymentLifecycleStateFailed,
	"SUCCEEDED":   DeploymentLifecycleStateSucceeded,
	"CANCELING":   DeploymentLifecycleStateCanceling,
	"CANCELED":    DeploymentLifecycleStateCanceled,
}

// GetDeploymentLifecycleStateEnumValues Enumerates the set of values for DeploymentLifecycleStateEnum
func GetDeploymentLifecycleStateEnumValues() []DeploymentLifecycleStateEnum {
	values := make([]DeploymentLifecycleStateEnum, 0)
	for _, v := range mappingDeploymentLifecycleState {
		values = append(values, v)
	}
	return values
}

// DeploymentDeploymentTypeEnum Enum with underlying type: string
type DeploymentDeploymentTypeEnum string

// Set of constants representing the allowable values for DeploymentDeploymentTypeEnum
const (
	DeploymentDeploymentTypePipelineDeployment    DeploymentDeploymentTypeEnum = "PIPELINE_DEPLOYMENT"
	DeploymentDeploymentTypePipelineRedeployment  DeploymentDeploymentTypeEnum = "PIPELINE_REDEPLOYMENT"
	DeploymentDeploymentTypeSingleStageDeployment DeploymentDeploymentTypeEnum = "SINGLE_STAGE_DEPLOYMENT"
)

var mappingDeploymentDeploymentType = map[string]DeploymentDeploymentTypeEnum{
	"PIPELINE_DEPLOYMENT":     DeploymentDeploymentTypePipelineDeployment,
	"PIPELINE_REDEPLOYMENT":   DeploymentDeploymentTypePipelineRedeployment,
	"SINGLE_STAGE_DEPLOYMENT": DeploymentDeploymentTypeSingleStageDeployment,
}

// GetDeploymentDeploymentTypeEnumValues Enumerates the set of values for DeploymentDeploymentTypeEnum
func GetDeploymentDeploymentTypeEnumValues() []DeploymentDeploymentTypeEnum {
	values := make([]DeploymentDeploymentTypeEnum, 0)
	for _, v := range mappingDeploymentDeploymentType {
		values = append(values, v)
	}
	return values
}
