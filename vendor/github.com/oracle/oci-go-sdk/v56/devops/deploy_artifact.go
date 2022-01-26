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

// DeployArtifact Artifacts are deployment manifests that are referenced in a pipeline stage for automated deployment to the target environment.  DevOps artifacts can be an OCI Container image repository, Kubernetes manifest, an Artifact Registry artifact, or defined inline.
type DeployArtifact struct {

	// Unique identifier that is immutable on creation.
	Id *string `mandatory:"true" json:"id"`

	// The OCID of a project.
	ProjectId *string `mandatory:"true" json:"projectId"`

	// The OCID of a compartment.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// Type of the deployment artifact.
	DeployArtifactType DeployArtifactDeployArtifactTypeEnum `mandatory:"true" json:"deployArtifactType"`

	// Mode for artifact parameter substitution.
	ArgumentSubstitutionMode DeployArtifactArgumentSubstitutionModeEnum `mandatory:"true" json:"argumentSubstitutionMode"`

	DeployArtifactSource DeployArtifactSource `mandatory:"true" json:"deployArtifactSource"`

	// Optional description about the artifact to be deployed.
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

func (m DeployArtifact) String() string {
	return common.PointerString(m)
}

// UnmarshalJSON unmarshals from json
func (m *DeployArtifact) UnmarshalJSON(data []byte) (e error) {
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
		ArgumentSubstitutionMode DeployArtifactArgumentSubstitutionModeEnum `json:"argumentSubstitutionMode"`
		DeployArtifactSource     deployartifactsource                       `json:"deployArtifactSource"`
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

	m.ArgumentSubstitutionMode = model.ArgumentSubstitutionMode

	nn, e = model.DeployArtifactSource.UnmarshalPolymorphicJSON(model.DeployArtifactSource.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.DeployArtifactSource = nn.(DeployArtifactSource)
	} else {
		m.DeployArtifactSource = nil
	}

	return
}

// DeployArtifactDeployArtifactTypeEnum Enum with underlying type: string
type DeployArtifactDeployArtifactTypeEnum string

// Set of constants representing the allowable values for DeployArtifactDeployArtifactTypeEnum
const (
	DeployArtifactDeployArtifactTypeDeploymentSpec     DeployArtifactDeployArtifactTypeEnum = "DEPLOYMENT_SPEC"
	DeployArtifactDeployArtifactTypeJobSpec            DeployArtifactDeployArtifactTypeEnum = "JOB_SPEC"
	DeployArtifactDeployArtifactTypeKubernetesManifest DeployArtifactDeployArtifactTypeEnum = "KUBERNETES_MANIFEST"
	DeployArtifactDeployArtifactTypeGenericFile        DeployArtifactDeployArtifactTypeEnum = "GENERIC_FILE"
	DeployArtifactDeployArtifactTypeDockerImage        DeployArtifactDeployArtifactTypeEnum = "DOCKER_IMAGE"
)

var mappingDeployArtifactDeployArtifactType = map[string]DeployArtifactDeployArtifactTypeEnum{
	"DEPLOYMENT_SPEC":     DeployArtifactDeployArtifactTypeDeploymentSpec,
	"JOB_SPEC":            DeployArtifactDeployArtifactTypeJobSpec,
	"KUBERNETES_MANIFEST": DeployArtifactDeployArtifactTypeKubernetesManifest,
	"GENERIC_FILE":        DeployArtifactDeployArtifactTypeGenericFile,
	"DOCKER_IMAGE":        DeployArtifactDeployArtifactTypeDockerImage,
}

// GetDeployArtifactDeployArtifactTypeEnumValues Enumerates the set of values for DeployArtifactDeployArtifactTypeEnum
func GetDeployArtifactDeployArtifactTypeEnumValues() []DeployArtifactDeployArtifactTypeEnum {
	values := make([]DeployArtifactDeployArtifactTypeEnum, 0)
	for _, v := range mappingDeployArtifactDeployArtifactType {
		values = append(values, v)
	}
	return values
}

// DeployArtifactArgumentSubstitutionModeEnum Enum with underlying type: string
type DeployArtifactArgumentSubstitutionModeEnum string

// Set of constants representing the allowable values for DeployArtifactArgumentSubstitutionModeEnum
const (
	DeployArtifactArgumentSubstitutionModeNone                   DeployArtifactArgumentSubstitutionModeEnum = "NONE"
	DeployArtifactArgumentSubstitutionModeSubstitutePlaceholders DeployArtifactArgumentSubstitutionModeEnum = "SUBSTITUTE_PLACEHOLDERS"
)

var mappingDeployArtifactArgumentSubstitutionMode = map[string]DeployArtifactArgumentSubstitutionModeEnum{
	"NONE":                    DeployArtifactArgumentSubstitutionModeNone,
	"SUBSTITUTE_PLACEHOLDERS": DeployArtifactArgumentSubstitutionModeSubstitutePlaceholders,
}

// GetDeployArtifactArgumentSubstitutionModeEnumValues Enumerates the set of values for DeployArtifactArgumentSubstitutionModeEnum
func GetDeployArtifactArgumentSubstitutionModeEnumValues() []DeployArtifactArgumentSubstitutionModeEnum {
	values := make([]DeployArtifactArgumentSubstitutionModeEnum, 0)
	for _, v := range mappingDeployArtifactArgumentSubstitutionMode {
		values = append(values, v)
	}
	return values
}

// DeployArtifactLifecycleStateEnum Enum with underlying type: string
type DeployArtifactLifecycleStateEnum string

// Set of constants representing the allowable values for DeployArtifactLifecycleStateEnum
const (
	DeployArtifactLifecycleStateCreating DeployArtifactLifecycleStateEnum = "CREATING"
	DeployArtifactLifecycleStateUpdating DeployArtifactLifecycleStateEnum = "UPDATING"
	DeployArtifactLifecycleStateActive   DeployArtifactLifecycleStateEnum = "ACTIVE"
	DeployArtifactLifecycleStateDeleting DeployArtifactLifecycleStateEnum = "DELETING"
	DeployArtifactLifecycleStateDeleted  DeployArtifactLifecycleStateEnum = "DELETED"
	DeployArtifactLifecycleStateFailed   DeployArtifactLifecycleStateEnum = "FAILED"
)

var mappingDeployArtifactLifecycleState = map[string]DeployArtifactLifecycleStateEnum{
	"CREATING": DeployArtifactLifecycleStateCreating,
	"UPDATING": DeployArtifactLifecycleStateUpdating,
	"ACTIVE":   DeployArtifactLifecycleStateActive,
	"DELETING": DeployArtifactLifecycleStateDeleting,
	"DELETED":  DeployArtifactLifecycleStateDeleted,
	"FAILED":   DeployArtifactLifecycleStateFailed,
}

// GetDeployArtifactLifecycleStateEnumValues Enumerates the set of values for DeployArtifactLifecycleStateEnum
func GetDeployArtifactLifecycleStateEnumValues() []DeployArtifactLifecycleStateEnum {
	values := make([]DeployArtifactLifecycleStateEnum, 0)
	for _, v := range mappingDeployArtifactLifecycleState {
		values = append(values, v)
	}
	return values
}
