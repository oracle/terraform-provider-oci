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
	"github.com/oracle/oci-go-sdk/v52/common"
)

// DevopsCodeRepositoryTriggerCreateResult Trigger Create response specific to Gitlab
type DevopsCodeRepositoryTriggerCreateResult struct {

	// Unique identifier that is immutable on creation
	Id *string `mandatory:"true" json:"id"`

	// Project to which the Trigger belongs
	ProjectId *string `mandatory:"true" json:"projectId"`

	// Compartment to which the Trigger belongs
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The list of actions that are to be performed for this Trigger
	Actions []TriggerAction `mandatory:"true" json:"actions"`

	// The id of the OCI Devops Repository
	RepositoryId *string `mandatory:"true" json:"repositoryId"`

	// Name for Trigger.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// Description about the Trigger
	Description *string `mandatory:"false" json:"description"`

	// The time the the Trigger was created. An RFC3339 formatted datetime string
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`

	// The time the Trigger was updated. An RFC3339 formatted datetime string
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`

	// A message describing the current state in more detail. For example, can be used to provide actionable information for a resource in Failed state.
	LifecycleDetails *string `mandatory:"false" json:"lifecycleDetails"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.  See Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm). Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace. See Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm). Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// Usage of system tag keys. These predefined keys are scoped to namespaces. See Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm). Example: `{"orcl-cloud": {"free-tier-retained": "true"}}`
	SystemTags map[string]map[string]interface{} `mandatory:"false" json:"systemTags"`

	// The current state of the Trigger.
	LifecycleState TriggerLifecycleStateEnum `mandatory:"false" json:"lifecycleState,omitempty"`
}

//GetId returns Id
func (m DevopsCodeRepositoryTriggerCreateResult) GetId() *string {
	return m.Id
}

//GetDisplayName returns DisplayName
func (m DevopsCodeRepositoryTriggerCreateResult) GetDisplayName() *string {
	return m.DisplayName
}

//GetDescription returns Description
func (m DevopsCodeRepositoryTriggerCreateResult) GetDescription() *string {
	return m.Description
}

//GetProjectId returns ProjectId
func (m DevopsCodeRepositoryTriggerCreateResult) GetProjectId() *string {
	return m.ProjectId
}

//GetCompartmentId returns CompartmentId
func (m DevopsCodeRepositoryTriggerCreateResult) GetCompartmentId() *string {
	return m.CompartmentId
}

//GetTimeCreated returns TimeCreated
func (m DevopsCodeRepositoryTriggerCreateResult) GetTimeCreated() *common.SDKTime {
	return m.TimeCreated
}

//GetTimeUpdated returns TimeUpdated
func (m DevopsCodeRepositoryTriggerCreateResult) GetTimeUpdated() *common.SDKTime {
	return m.TimeUpdated
}

//GetLifecycleState returns LifecycleState
func (m DevopsCodeRepositoryTriggerCreateResult) GetLifecycleState() TriggerLifecycleStateEnum {
	return m.LifecycleState
}

//GetLifecycleDetails returns LifecycleDetails
func (m DevopsCodeRepositoryTriggerCreateResult) GetLifecycleDetails() *string {
	return m.LifecycleDetails
}

//GetActions returns Actions
func (m DevopsCodeRepositoryTriggerCreateResult) GetActions() []TriggerAction {
	return m.Actions
}

//GetFreeformTags returns FreeformTags
func (m DevopsCodeRepositoryTriggerCreateResult) GetFreeformTags() map[string]string {
	return m.FreeformTags
}

//GetDefinedTags returns DefinedTags
func (m DevopsCodeRepositoryTriggerCreateResult) GetDefinedTags() map[string]map[string]interface{} {
	return m.DefinedTags
}

//GetSystemTags returns SystemTags
func (m DevopsCodeRepositoryTriggerCreateResult) GetSystemTags() map[string]map[string]interface{} {
	return m.SystemTags
}

func (m DevopsCodeRepositoryTriggerCreateResult) String() string {
	return common.PointerString(m)
}

// MarshalJSON marshals to json representation
func (m DevopsCodeRepositoryTriggerCreateResult) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeDevopsCodeRepositoryTriggerCreateResult DevopsCodeRepositoryTriggerCreateResult
	s := struct {
		DiscriminatorParam string `json:"triggerSource"`
		MarshalTypeDevopsCodeRepositoryTriggerCreateResult
	}{
		"DEVOPS_CODE_REPOSITORY",
		(MarshalTypeDevopsCodeRepositoryTriggerCreateResult)(m),
	}

	return json.Marshal(&s)
}

// UnmarshalJSON unmarshals from json
func (m *DevopsCodeRepositoryTriggerCreateResult) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		DisplayName      *string                           `json:"displayName"`
		Description      *string                           `json:"description"`
		TimeCreated      *common.SDKTime                   `json:"timeCreated"`
		TimeUpdated      *common.SDKTime                   `json:"timeUpdated"`
		LifecycleState   TriggerLifecycleStateEnum         `json:"lifecycleState"`
		LifecycleDetails *string                           `json:"lifecycleDetails"`
		FreeformTags     map[string]string                 `json:"freeformTags"`
		DefinedTags      map[string]map[string]interface{} `json:"definedTags"`
		SystemTags       map[string]map[string]interface{} `json:"systemTags"`
		Id               *string                           `json:"id"`
		ProjectId        *string                           `json:"projectId"`
		CompartmentId    *string                           `json:"compartmentId"`
		Actions          []triggeraction                   `json:"actions"`
		RepositoryId     *string                           `json:"repositoryId"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.DisplayName = model.DisplayName

	m.Description = model.Description

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

	m.RepositoryId = model.RepositoryId

	return
}
