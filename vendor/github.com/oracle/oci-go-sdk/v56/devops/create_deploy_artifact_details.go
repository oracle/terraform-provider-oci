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

// CreateDeployArtifactDetails Information about a new deployment artifact.
type CreateDeployArtifactDetails struct {

	// Type of the deployment artifact.
	DeployArtifactType DeployArtifactDeployArtifactTypeEnum `mandatory:"true" json:"deployArtifactType"`

	DeployArtifactSource DeployArtifactSource `mandatory:"true" json:"deployArtifactSource"`

	// Mode for artifact parameter substitution.
	ArgumentSubstitutionMode DeployArtifactArgumentSubstitutionModeEnum `mandatory:"true" json:"argumentSubstitutionMode"`

	// The OCID of a project.
	ProjectId *string `mandatory:"true" json:"projectId"`

	// Optional description about the deployment artifact.
	Description *string `mandatory:"false" json:"description"`

	// Deployment artifact display name. Avoid entering confidential information.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.  See Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm). Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace. See Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm). Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m CreateDeployArtifactDetails) String() string {
	return common.PointerString(m)
}

// UnmarshalJSON unmarshals from json
func (m *CreateDeployArtifactDetails) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		Description              *string                                    `json:"description"`
		DisplayName              *string                                    `json:"displayName"`
		FreeformTags             map[string]string                          `json:"freeformTags"`
		DefinedTags              map[string]map[string]interface{}          `json:"definedTags"`
		DeployArtifactType       DeployArtifactDeployArtifactTypeEnum       `json:"deployArtifactType"`
		DeployArtifactSource     deployartifactsource                       `json:"deployArtifactSource"`
		ArgumentSubstitutionMode DeployArtifactArgumentSubstitutionModeEnum `json:"argumentSubstitutionMode"`
		ProjectId                *string                                    `json:"projectId"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.Description = model.Description

	m.DisplayName = model.DisplayName

	m.FreeformTags = model.FreeformTags

	m.DefinedTags = model.DefinedTags

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

	m.ProjectId = model.ProjectId

	return
}
