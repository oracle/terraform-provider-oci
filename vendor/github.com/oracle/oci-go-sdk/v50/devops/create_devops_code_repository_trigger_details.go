// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// DevOps API
//
// Use the DevOps APIs to create a DevOps project to group the pipelines,  add reference to target deployment environments, add artifacts to deploy,  and create deployment pipelines needed to deploy your software.
//

package devops

import (
	"encoding/json"
	"github.com/oracle/oci-go-sdk/v50/common"
)

// CreateDevopsCodeRepositoryTriggerDetails The Trigger for Devops-Repository as the caller
type CreateDevopsCodeRepositoryTriggerDetails struct {

	// Project to which the Trigger will belong
	ProjectId *string `mandatory:"true" json:"projectId"`

	// The list of actions that are to be performed for this Trigger
	Actions []TriggerAction `mandatory:"true" json:"actions"`

	// Name of the Trigger
	DisplayName *string `mandatory:"false" json:"displayName"`

	// Optional description about the Trigger
	Description *string `mandatory:"false" json:"description"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.  See Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm). Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace. See Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm). Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// The Devops Code Repository Id
	RepositoryId *string `mandatory:"false" json:"repositoryId"`
}

//GetDisplayName returns DisplayName
func (m CreateDevopsCodeRepositoryTriggerDetails) GetDisplayName() *string {
	return m.DisplayName
}

//GetDescription returns Description
func (m CreateDevopsCodeRepositoryTriggerDetails) GetDescription() *string {
	return m.Description
}

//GetProjectId returns ProjectId
func (m CreateDevopsCodeRepositoryTriggerDetails) GetProjectId() *string {
	return m.ProjectId
}

//GetActions returns Actions
func (m CreateDevopsCodeRepositoryTriggerDetails) GetActions() []TriggerAction {
	return m.Actions
}

//GetFreeformTags returns FreeformTags
func (m CreateDevopsCodeRepositoryTriggerDetails) GetFreeformTags() map[string]string {
	return m.FreeformTags
}

//GetDefinedTags returns DefinedTags
func (m CreateDevopsCodeRepositoryTriggerDetails) GetDefinedTags() map[string]map[string]interface{} {
	return m.DefinedTags
}

func (m CreateDevopsCodeRepositoryTriggerDetails) String() string {
	return common.PointerString(m)
}

// MarshalJSON marshals to json representation
func (m CreateDevopsCodeRepositoryTriggerDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeCreateDevopsCodeRepositoryTriggerDetails CreateDevopsCodeRepositoryTriggerDetails
	s := struct {
		DiscriminatorParam string `json:"triggerSource"`
		MarshalTypeCreateDevopsCodeRepositoryTriggerDetails
	}{
		"DEVOPS_CODE_REPOSITORY",
		(MarshalTypeCreateDevopsCodeRepositoryTriggerDetails)(m),
	}

	return json.Marshal(&s)
}

// UnmarshalJSON unmarshals from json
func (m *CreateDevopsCodeRepositoryTriggerDetails) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		DisplayName  *string                           `json:"displayName"`
		Description  *string                           `json:"description"`
		FreeformTags map[string]string                 `json:"freeformTags"`
		DefinedTags  map[string]map[string]interface{} `json:"definedTags"`
		RepositoryId *string                           `json:"repositoryId"`
		ProjectId    *string                           `json:"projectId"`
		Actions      []triggeraction                   `json:"actions"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.DisplayName = model.DisplayName

	m.Description = model.Description

	m.FreeformTags = model.FreeformTags

	m.DefinedTags = model.DefinedTags

	m.RepositoryId = model.RepositoryId

	m.ProjectId = model.ProjectId

	m.Actions = make([]TriggerAction, len(model.Actions))
	for i, n := range model.Actions {
		nn, e = n.UnmarshalPolymorphicJSON(n.JsonData)
		if e != nil {
			return e
		}
		if nn != nil {
			m.Actions[i] = nn.(TriggerAction)
		} else {
			m.Actions[i] = nil
		}
	}

	return
}
