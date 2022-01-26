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

// DeployEnvironment The target OCI resources, such as Compute instances, Container Engine for Kubernetes(OKE) clusters, or Function, where artifacts will be deployed.
type DeployEnvironment interface {

	// Unique identifier that is immutable on creation.
	GetId() *string

	// The OCID of a project.
	GetProjectId() *string

	// The OCID of a compartment.
	GetCompartmentId() *string

	// Optional description about the deployment environment.
	GetDescription() *string

	// Deployment environment display name, which can be renamed and is not necessarily unique. Avoid entering confidential information.
	GetDisplayName() *string

	// Time the deployment environment was created. Format defined by RFC3339 (https://datatracker.ietf.org/doc/html/rfc3339).
	GetTimeCreated() *common.SDKTime

	// Time the deployment environment was updated. Format defined by RFC3339 (https://datatracker.ietf.org/doc/html/rfc3339).
	GetTimeUpdated() *common.SDKTime

	// The current state of the deployment environment.
	GetLifecycleState() DeployEnvironmentLifecycleStateEnum

	// A message describing the current state in more detail. For example, can be used to provide actionable information for a resource in Failed state.
	GetLifecycleDetails() *string

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.  See Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm). Example: `{"bar-key": "value"}`
	GetFreeformTags() map[string]string

	// Defined tags for this resource. Each key is predefined and scoped to a namespace. See Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm). Example: `{"foo-namespace": {"bar-key": "value"}}`
	GetDefinedTags() map[string]map[string]interface{}

	// Usage of system tag keys. These predefined keys are scoped to namespaces. See Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm). Example: `{"orcl-cloud": {"free-tier-retained": "true"}}`
	GetSystemTags() map[string]map[string]interface{}
}

type deployenvironment struct {
	JsonData              []byte
	Id                    *string                             `mandatory:"true" json:"id"`
	ProjectId             *string                             `mandatory:"true" json:"projectId"`
	CompartmentId         *string                             `mandatory:"true" json:"compartmentId"`
	Description           *string                             `mandatory:"false" json:"description"`
	DisplayName           *string                             `mandatory:"false" json:"displayName"`
	TimeCreated           *common.SDKTime                     `mandatory:"false" json:"timeCreated"`
	TimeUpdated           *common.SDKTime                     `mandatory:"false" json:"timeUpdated"`
	LifecycleState        DeployEnvironmentLifecycleStateEnum `mandatory:"false" json:"lifecycleState,omitempty"`
	LifecycleDetails      *string                             `mandatory:"false" json:"lifecycleDetails"`
	FreeformTags          map[string]string                   `mandatory:"false" json:"freeformTags"`
	DefinedTags           map[string]map[string]interface{}   `mandatory:"false" json:"definedTags"`
	SystemTags            map[string]map[string]interface{}   `mandatory:"false" json:"systemTags"`
	DeployEnvironmentType string                              `json:"deployEnvironmentType"`
}

// UnmarshalJSON unmarshals json
func (m *deployenvironment) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerdeployenvironment deployenvironment
	s := struct {
		Model Unmarshalerdeployenvironment
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.Id = s.Model.Id
	m.ProjectId = s.Model.ProjectId
	m.CompartmentId = s.Model.CompartmentId
	m.Description = s.Model.Description
	m.DisplayName = s.Model.DisplayName
	m.TimeCreated = s.Model.TimeCreated
	m.TimeUpdated = s.Model.TimeUpdated
	m.LifecycleState = s.Model.LifecycleState
	m.LifecycleDetails = s.Model.LifecycleDetails
	m.FreeformTags = s.Model.FreeformTags
	m.DefinedTags = s.Model.DefinedTags
	m.SystemTags = s.Model.SystemTags
	m.DeployEnvironmentType = s.Model.DeployEnvironmentType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *deployenvironment) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.DeployEnvironmentType {
	case "COMPUTE_INSTANCE_GROUP":
		mm := ComputeInstanceGroupDeployEnvironment{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "OKE_CLUSTER":
		mm := OkeClusterDeployEnvironment{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "FUNCTION":
		mm := FunctionDeployEnvironment{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		return *m, nil
	}
}

//GetId returns Id
func (m deployenvironment) GetId() *string {
	return m.Id
}

//GetProjectId returns ProjectId
func (m deployenvironment) GetProjectId() *string {
	return m.ProjectId
}

//GetCompartmentId returns CompartmentId
func (m deployenvironment) GetCompartmentId() *string {
	return m.CompartmentId
}

//GetDescription returns Description
func (m deployenvironment) GetDescription() *string {
	return m.Description
}

//GetDisplayName returns DisplayName
func (m deployenvironment) GetDisplayName() *string {
	return m.DisplayName
}

//GetTimeCreated returns TimeCreated
func (m deployenvironment) GetTimeCreated() *common.SDKTime {
	return m.TimeCreated
}

//GetTimeUpdated returns TimeUpdated
func (m deployenvironment) GetTimeUpdated() *common.SDKTime {
	return m.TimeUpdated
}

//GetLifecycleState returns LifecycleState
func (m deployenvironment) GetLifecycleState() DeployEnvironmentLifecycleStateEnum {
	return m.LifecycleState
}

//GetLifecycleDetails returns LifecycleDetails
func (m deployenvironment) GetLifecycleDetails() *string {
	return m.LifecycleDetails
}

//GetFreeformTags returns FreeformTags
func (m deployenvironment) GetFreeformTags() map[string]string {
	return m.FreeformTags
}

//GetDefinedTags returns DefinedTags
func (m deployenvironment) GetDefinedTags() map[string]map[string]interface{} {
	return m.DefinedTags
}

//GetSystemTags returns SystemTags
func (m deployenvironment) GetSystemTags() map[string]map[string]interface{} {
	return m.SystemTags
}

func (m deployenvironment) String() string {
	return common.PointerString(m)
}

// DeployEnvironmentLifecycleStateEnum Enum with underlying type: string
type DeployEnvironmentLifecycleStateEnum string

// Set of constants representing the allowable values for DeployEnvironmentLifecycleStateEnum
const (
	DeployEnvironmentLifecycleStateCreating DeployEnvironmentLifecycleStateEnum = "CREATING"
	DeployEnvironmentLifecycleStateUpdating DeployEnvironmentLifecycleStateEnum = "UPDATING"
	DeployEnvironmentLifecycleStateActive   DeployEnvironmentLifecycleStateEnum = "ACTIVE"
	DeployEnvironmentLifecycleStateDeleting DeployEnvironmentLifecycleStateEnum = "DELETING"
	DeployEnvironmentLifecycleStateDeleted  DeployEnvironmentLifecycleStateEnum = "DELETED"
	DeployEnvironmentLifecycleStateFailed   DeployEnvironmentLifecycleStateEnum = "FAILED"
)

var mappingDeployEnvironmentLifecycleState = map[string]DeployEnvironmentLifecycleStateEnum{
	"CREATING": DeployEnvironmentLifecycleStateCreating,
	"UPDATING": DeployEnvironmentLifecycleStateUpdating,
	"ACTIVE":   DeployEnvironmentLifecycleStateActive,
	"DELETING": DeployEnvironmentLifecycleStateDeleting,
	"DELETED":  DeployEnvironmentLifecycleStateDeleted,
	"FAILED":   DeployEnvironmentLifecycleStateFailed,
}

// GetDeployEnvironmentLifecycleStateEnumValues Enumerates the set of values for DeployEnvironmentLifecycleStateEnum
func GetDeployEnvironmentLifecycleStateEnumValues() []DeployEnvironmentLifecycleStateEnum {
	values := make([]DeployEnvironmentLifecycleStateEnum, 0)
	for _, v := range mappingDeployEnvironmentLifecycleState {
		values = append(values, v)
	}
	return values
}

// DeployEnvironmentDeployEnvironmentTypeEnum Enum with underlying type: string
type DeployEnvironmentDeployEnvironmentTypeEnum string

// Set of constants representing the allowable values for DeployEnvironmentDeployEnvironmentTypeEnum
const (
	DeployEnvironmentDeployEnvironmentTypeOkeCluster           DeployEnvironmentDeployEnvironmentTypeEnum = "OKE_CLUSTER"
	DeployEnvironmentDeployEnvironmentTypeComputeInstanceGroup DeployEnvironmentDeployEnvironmentTypeEnum = "COMPUTE_INSTANCE_GROUP"
	DeployEnvironmentDeployEnvironmentTypeFunction             DeployEnvironmentDeployEnvironmentTypeEnum = "FUNCTION"
)

var mappingDeployEnvironmentDeployEnvironmentType = map[string]DeployEnvironmentDeployEnvironmentTypeEnum{
	"OKE_CLUSTER":            DeployEnvironmentDeployEnvironmentTypeOkeCluster,
	"COMPUTE_INSTANCE_GROUP": DeployEnvironmentDeployEnvironmentTypeComputeInstanceGroup,
	"FUNCTION":               DeployEnvironmentDeployEnvironmentTypeFunction,
}

// GetDeployEnvironmentDeployEnvironmentTypeEnumValues Enumerates the set of values for DeployEnvironmentDeployEnvironmentTypeEnum
func GetDeployEnvironmentDeployEnvironmentTypeEnumValues() []DeployEnvironmentDeployEnvironmentTypeEnum {
	values := make([]DeployEnvironmentDeployEnvironmentTypeEnum, 0)
	for _, v := range mappingDeployEnvironmentDeployEnvironmentType {
		values = append(values, v)
	}
	return values
}
