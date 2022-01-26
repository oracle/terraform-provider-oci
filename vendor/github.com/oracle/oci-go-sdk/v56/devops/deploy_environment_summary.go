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

// DeployEnvironmentSummary Summary of the deployment environment.
type DeployEnvironmentSummary interface {

	// Unique identifier that is immutable on creation.
	GetId() *string

	// The OCID of a project.
	GetProjectId() *string

	// The OCID of a compartment.
	GetCompartmentId() *string

	// Optional description about the deployment environment.
	GetDescription() *string

	// Deployment environment display name, which can be renamed and is not necessarily unique.
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

type deployenvironmentsummary struct {
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
func (m *deployenvironmentsummary) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerdeployenvironmentsummary deployenvironmentsummary
	s := struct {
		Model Unmarshalerdeployenvironmentsummary
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
func (m *deployenvironmentsummary) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.DeployEnvironmentType {
	case "FUNCTION":
		mm := FunctionDeployEnvironmentSummary{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "COMPUTE_INSTANCE_GROUP":
		mm := ComputeInstanceGroupDeployEnvironmentSummary{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "OKE_CLUSTER":
		mm := OkeClusterDeployEnvironmentSummary{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		return *m, nil
	}
}

//GetId returns Id
func (m deployenvironmentsummary) GetId() *string {
	return m.Id
}

//GetProjectId returns ProjectId
func (m deployenvironmentsummary) GetProjectId() *string {
	return m.ProjectId
}

//GetCompartmentId returns CompartmentId
func (m deployenvironmentsummary) GetCompartmentId() *string {
	return m.CompartmentId
}

//GetDescription returns Description
func (m deployenvironmentsummary) GetDescription() *string {
	return m.Description
}

//GetDisplayName returns DisplayName
func (m deployenvironmentsummary) GetDisplayName() *string {
	return m.DisplayName
}

//GetTimeCreated returns TimeCreated
func (m deployenvironmentsummary) GetTimeCreated() *common.SDKTime {
	return m.TimeCreated
}

//GetTimeUpdated returns TimeUpdated
func (m deployenvironmentsummary) GetTimeUpdated() *common.SDKTime {
	return m.TimeUpdated
}

//GetLifecycleState returns LifecycleState
func (m deployenvironmentsummary) GetLifecycleState() DeployEnvironmentLifecycleStateEnum {
	return m.LifecycleState
}

//GetLifecycleDetails returns LifecycleDetails
func (m deployenvironmentsummary) GetLifecycleDetails() *string {
	return m.LifecycleDetails
}

//GetFreeformTags returns FreeformTags
func (m deployenvironmentsummary) GetFreeformTags() map[string]string {
	return m.FreeformTags
}

//GetDefinedTags returns DefinedTags
func (m deployenvironmentsummary) GetDefinedTags() map[string]map[string]interface{} {
	return m.DefinedTags
}

//GetSystemTags returns SystemTags
func (m deployenvironmentsummary) GetSystemTags() map[string]map[string]interface{} {
	return m.SystemTags
}

func (m deployenvironmentsummary) String() string {
	return common.PointerString(m)
}
