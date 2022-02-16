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

// DeploymentSummary Summary of the deployment.
type DeploymentSummary interface {

	// Unique identifier that is immutable on creation.
	GetId() *string

	// The OCID of a project.
	GetProjectId() *string

	// The OCID of a pipeline.
	GetDeployPipelineId() *string

	// The OCID of a compartment.
	GetCompartmentId() *string

	// Deployment identifier which can be renamed and is not necessarily unique. Avoid entering confidential information.
	GetDisplayName() *string

	// Time the deployment was created. Format defined by RFC3339 (https://datatracker.ietf.org/doc/html/rfc3339).
	GetTimeCreated() *common.SDKTime

	// Time the deployment was updated. Format defined by RFC3339 (https://datatracker.ietf.org/doc/html/rfc3339).
	GetTimeUpdated() *common.SDKTime

	// The current state of the deployment.
	GetLifecycleState() DeploymentLifecycleStateEnum

	GetDeploymentArguments() *DeploymentArgumentCollection

	GetDeployArtifactOverrideArguments() *DeployArtifactOverrideArgumentCollection

	// A message describing the current state in more detail. For example, can be used to provide actionable information for a resource in Failed state.
	GetLifecycleDetails() *string

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.  See Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm). Example: `{"bar-key": "value"}`
	GetFreeformTags() map[string]string

	// Defined tags for this resource. Each key is predefined and scoped to a namespace. See Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm). Example: `{"foo-namespace": {"bar-key": "value"}}`
	GetDefinedTags() map[string]map[string]interface{}

	// Usage of system tag keys. These predefined keys are scoped to namespaces. See Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm). Example: `{"orcl-cloud": {"free-tier-retained": "true"}}`
	GetSystemTags() map[string]map[string]interface{}
}

type deploymentsummary struct {
	JsonData                        []byte
	Id                              *string                                   `mandatory:"true" json:"id"`
	ProjectId                       *string                                   `mandatory:"true" json:"projectId"`
	DeployPipelineId                *string                                   `mandatory:"true" json:"deployPipelineId"`
	CompartmentId                   *string                                   `mandatory:"true" json:"compartmentId"`
	DisplayName                     *string                                   `mandatory:"false" json:"displayName"`
	TimeCreated                     *common.SDKTime                           `mandatory:"false" json:"timeCreated"`
	TimeUpdated                     *common.SDKTime                           `mandatory:"false" json:"timeUpdated"`
	LifecycleState                  DeploymentLifecycleStateEnum              `mandatory:"false" json:"lifecycleState,omitempty"`
	DeploymentArguments             *DeploymentArgumentCollection             `mandatory:"false" json:"deploymentArguments"`
	DeployArtifactOverrideArguments *DeployArtifactOverrideArgumentCollection `mandatory:"false" json:"deployArtifactOverrideArguments"`
	LifecycleDetails                *string                                   `mandatory:"false" json:"lifecycleDetails"`
	FreeformTags                    map[string]string                         `mandatory:"false" json:"freeformTags"`
	DefinedTags                     map[string]map[string]interface{}         `mandatory:"false" json:"definedTags"`
	SystemTags                      map[string]map[string]interface{}         `mandatory:"false" json:"systemTags"`
	DeploymentType                  string                                    `json:"deploymentType"`
}

// UnmarshalJSON unmarshals json
func (m *deploymentsummary) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerdeploymentsummary deploymentsummary
	s := struct {
		Model Unmarshalerdeploymentsummary
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.Id = s.Model.Id
	m.ProjectId = s.Model.ProjectId
	m.DeployPipelineId = s.Model.DeployPipelineId
	m.CompartmentId = s.Model.CompartmentId
	m.DisplayName = s.Model.DisplayName
	m.TimeCreated = s.Model.TimeCreated
	m.TimeUpdated = s.Model.TimeUpdated
	m.LifecycleState = s.Model.LifecycleState
	m.DeploymentArguments = s.Model.DeploymentArguments
	m.DeployArtifactOverrideArguments = s.Model.DeployArtifactOverrideArguments
	m.LifecycleDetails = s.Model.LifecycleDetails
	m.FreeformTags = s.Model.FreeformTags
	m.DefinedTags = s.Model.DefinedTags
	m.SystemTags = s.Model.SystemTags
	m.DeploymentType = s.Model.DeploymentType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *deploymentsummary) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.DeploymentType {
	case "SINGLE_STAGE_DEPLOYMENT":
		mm := SingleDeployStageDeploymentSummary{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "PIPELINE_REDEPLOYMENT":
		mm := DeployPipelineRedeploymentSummary{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "PIPELINE_DEPLOYMENT":
		mm := DeployPipelineDeploymentSummary{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		return *m, nil
	}
}

//GetId returns Id
func (m deploymentsummary) GetId() *string {
	return m.Id
}

//GetProjectId returns ProjectId
func (m deploymentsummary) GetProjectId() *string {
	return m.ProjectId
}

//GetDeployPipelineId returns DeployPipelineId
func (m deploymentsummary) GetDeployPipelineId() *string {
	return m.DeployPipelineId
}

//GetCompartmentId returns CompartmentId
func (m deploymentsummary) GetCompartmentId() *string {
	return m.CompartmentId
}

//GetDisplayName returns DisplayName
func (m deploymentsummary) GetDisplayName() *string {
	return m.DisplayName
}

//GetTimeCreated returns TimeCreated
func (m deploymentsummary) GetTimeCreated() *common.SDKTime {
	return m.TimeCreated
}

//GetTimeUpdated returns TimeUpdated
func (m deploymentsummary) GetTimeUpdated() *common.SDKTime {
	return m.TimeUpdated
}

//GetLifecycleState returns LifecycleState
func (m deploymentsummary) GetLifecycleState() DeploymentLifecycleStateEnum {
	return m.LifecycleState
}

//GetDeploymentArguments returns DeploymentArguments
func (m deploymentsummary) GetDeploymentArguments() *DeploymentArgumentCollection {
	return m.DeploymentArguments
}

//GetDeployArtifactOverrideArguments returns DeployArtifactOverrideArguments
func (m deploymentsummary) GetDeployArtifactOverrideArguments() *DeployArtifactOverrideArgumentCollection {
	return m.DeployArtifactOverrideArguments
}

//GetLifecycleDetails returns LifecycleDetails
func (m deploymentsummary) GetLifecycleDetails() *string {
	return m.LifecycleDetails
}

//GetFreeformTags returns FreeformTags
func (m deploymentsummary) GetFreeformTags() map[string]string {
	return m.FreeformTags
}

//GetDefinedTags returns DefinedTags
func (m deploymentsummary) GetDefinedTags() map[string]map[string]interface{} {
	return m.DefinedTags
}

//GetSystemTags returns SystemTags
func (m deploymentsummary) GetSystemTags() map[string]map[string]interface{} {
	return m.SystemTags
}

func (m deploymentsummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m deploymentsummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingDeploymentLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetDeploymentLifecycleStateEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
