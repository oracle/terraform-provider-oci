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

// DeployArtifactSummary Summary of the deployment artifact.
type DeployArtifactSummary struct {

	// Unique identifier that is immutable on creation.
	Id *string `mandatory:"true" json:"id"`

	// The OCID of a project.
	ProjectId *string `mandatory:"true" json:"projectId"`

	// The OCID of a compartment.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// Type of the deployment artifact.
	DeployArtifactType DeployArtifactDeployArtifactTypeEnum `mandatory:"true" json:"deployArtifactType"`

	DeployArtifactSource DeployArtifactSource `mandatory:"true" json:"deployArtifactSource"`

	// Mode for artifact parameter substitution.
	ArgumentSubstitutionMode DeployArtifactArgumentSubstitutionModeEnum `mandatory:"true" json:"argumentSubstitutionMode"`

	// Optional description about the deployment artifact.
	Description *string `mandatory:"false" json:"description"`

	// Deployment artifact identifier, which can be renamed and is not necessarily unique. Avoid entering confidential information.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// Time the deployment artifact was created. Format defined by RFC3339 (https://datatracker.ietf.org/doc/html/rfc3339).
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`

	// Time the deployment artifact was updated. Format defined by RFC3339 (https://datatracker.ietf.org/doc/html/rfc3339).
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`

	// Current state of the deployment artifact.
	LifecycleState DeployArtifactLifecycleStateEnum `mandatory:"false" json:"lifecycleState,omitempty"`

	// A detailed message describing the current state. For example, can be used to provide actionable information for a resource in Failed state.
	LifecycleDetails *string `mandatory:"false" json:"lifecycleDetails"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.  See Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm). Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace. See Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm). Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// Usage of system tag keys. These predefined keys are scoped to namespaces. See Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm). Example: `{"orcl-cloud": {"free-tier-retained": "true"}}`
	SystemTags map[string]map[string]interface{} `mandatory:"false" json:"systemTags"`
}

func (m DeployArtifactSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m DeployArtifactSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingDeployArtifactDeployArtifactTypeEnum(string(m.DeployArtifactType)); !ok && m.DeployArtifactType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for DeployArtifactType: %s. Supported values are: %s.", m.DeployArtifactType, strings.Join(GetDeployArtifactDeployArtifactTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingDeployArtifactArgumentSubstitutionModeEnum(string(m.ArgumentSubstitutionMode)); !ok && m.ArgumentSubstitutionMode != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ArgumentSubstitutionMode: %s. Supported values are: %s.", m.ArgumentSubstitutionMode, strings.Join(GetDeployArtifactArgumentSubstitutionModeEnumStringValues(), ",")))
	}

	if _, ok := GetMappingDeployArtifactLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetDeployArtifactLifecycleStateEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UnmarshalJSON unmarshals from json
func (m *DeployArtifactSummary) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		Description              *string                                    `json:"description"`
		DisplayName              *string                                    `json:"displayName"`
		TimeCreated              *common.SDKTime                            `json:"timeCreated"`
		TimeUpdated              *common.SDKTime                            `json:"timeUpdated"`
		LifecycleState           DeployArtifactLifecycleStateEnum           `json:"lifecycleState"`
		LifecycleDetails         *string                                    `json:"lifecycleDetails"`
		FreeformTags             map[string]string                          `json:"freeformTags"`
		DefinedTags              map[string]map[string]interface{}          `json:"definedTags"`
		SystemTags               map[string]map[string]interface{}          `json:"systemTags"`
		Id                       *string                                    `json:"id"`
		ProjectId                *string                                    `json:"projectId"`
		CompartmentId            *string                                    `json:"compartmentId"`
		DeployArtifactType       DeployArtifactDeployArtifactTypeEnum       `json:"deployArtifactType"`
		DeployArtifactSource     deployartifactsource                       `json:"deployArtifactSource"`
		ArgumentSubstitutionMode DeployArtifactArgumentSubstitutionModeEnum `json:"argumentSubstitutionMode"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.Description = model.Description

	m.DisplayName = model.DisplayName

	m.TimeCreated = model.TimeCreated

	m.TimeUpdated = model.TimeUpdated

	m.LifecycleState = model.LifecycleState

	m.LifecycleDetails = model.LifecycleDetails

	m.FreeformTags = model.FreeformTags

	m.DefinedTags = model.DefinedTags

	m.SystemTags = model.SystemTags

	m.Id = model.Id

	m.ProjectId = model.ProjectId

	m.CompartmentId = model.CompartmentId

	m.DeployArtifactType = model.DeployArtifactType

	nn, e = model.DeployArtifactSource.UnmarshalPolymorphicJSON(model.DeployArtifactSource.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.DeployArtifactSource = nn.(DeployArtifactSource)
	} else {
		m.DeployArtifactSource = nil
	}

	m.ArgumentSubstitutionMode = model.ArgumentSubstitutionMode

	return
}
