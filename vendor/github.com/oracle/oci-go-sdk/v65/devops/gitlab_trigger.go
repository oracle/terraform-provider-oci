// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// DevOps API
//
// Use the DevOps API to create DevOps projects, configure code repositories,  add artifacts to deploy, build and test software applications, configure  target deployment environments, and deploy software applications.  For more information, see DevOps (https://docs.oracle.com/iaas/Content/devops/using/home.htm).
//

package devops

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// GitlabTrigger Trigger specific to GitLab.
type GitlabTrigger struct {

	// Unique identifier that is immutable on creation.
	Id *string `mandatory:"true" json:"id"`

	// The OCID of the DevOps project to which the trigger belongs to.
	ProjectId *string `mandatory:"true" json:"projectId"`

	// The OCID of the compartment that contains the trigger.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The list of actions that are to be performed for this trigger.
	Actions []TriggerAction `mandatory:"true" json:"actions"`

	// The endpoint that listens to trigger events.
	TriggerUrl *string `mandatory:"true" json:"triggerUrl"`

	// Trigger display name. Avoid entering confidential information.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// Description about the trigger.
	Description *string `mandatory:"false" json:"description"`

	// The time the trigger was created. Format defined by RFC3339 (https://datatracker.ietf.org/doc/html/rfc3339).
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`

	// The time the trigger was updated. Format defined by RFC3339 (https://datatracker.ietf.org/doc/html/rfc3339).
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`

	// A message describing the current state in more detail. For example, can be used to provide actionable information for a resource in Failed state.
	LifecycleDetails *string `mandatory:"false" json:"lifecycleDetails"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.  See Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace. See Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// Usage of system tag keys. These predefined keys are scoped to namespaces. See Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"orcl-cloud": {"free-tier-retained": "true"}}`
	SystemTags map[string]map[string]interface{} `mandatory:"false" json:"systemTags"`

	// The OCID of the connection resource used to get details for triggered events.
	ConnectionId *string `mandatory:"false" json:"connectionId"`

	// The current state of the trigger.
	LifecycleState TriggerLifecycleStateEnum `mandatory:"false" json:"lifecycleState,omitempty"`
}

// GetId returns Id
func (m GitlabTrigger) GetId() *string {
	return m.Id
}

// GetDisplayName returns DisplayName
func (m GitlabTrigger) GetDisplayName() *string {
	return m.DisplayName
}

// GetDescription returns Description
func (m GitlabTrigger) GetDescription() *string {
	return m.Description
}

// GetProjectId returns ProjectId
func (m GitlabTrigger) GetProjectId() *string {
	return m.ProjectId
}

// GetCompartmentId returns CompartmentId
func (m GitlabTrigger) GetCompartmentId() *string {
	return m.CompartmentId
}

// GetTimeCreated returns TimeCreated
func (m GitlabTrigger) GetTimeCreated() *common.SDKTime {
	return m.TimeCreated
}

// GetTimeUpdated returns TimeUpdated
func (m GitlabTrigger) GetTimeUpdated() *common.SDKTime {
	return m.TimeUpdated
}

// GetLifecycleState returns LifecycleState
func (m GitlabTrigger) GetLifecycleState() TriggerLifecycleStateEnum {
	return m.LifecycleState
}

// GetLifecycleDetails returns LifecycleDetails
func (m GitlabTrigger) GetLifecycleDetails() *string {
	return m.LifecycleDetails
}

// GetActions returns Actions
func (m GitlabTrigger) GetActions() []TriggerAction {
	return m.Actions
}

// GetFreeformTags returns FreeformTags
func (m GitlabTrigger) GetFreeformTags() map[string]string {
	return m.FreeformTags
}

// GetDefinedTags returns DefinedTags
func (m GitlabTrigger) GetDefinedTags() map[string]map[string]interface{} {
	return m.DefinedTags
}

// GetSystemTags returns SystemTags
func (m GitlabTrigger) GetSystemTags() map[string]map[string]interface{} {
	return m.SystemTags
}

func (m GitlabTrigger) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m GitlabTrigger) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingTriggerLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetTriggerLifecycleStateEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m GitlabTrigger) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeGitlabTrigger GitlabTrigger
	s := struct {
		DiscriminatorParam string `json:"triggerSource"`
		MarshalTypeGitlabTrigger
	}{
		"GITLAB",
		(MarshalTypeGitlabTrigger)(m),
	}

	return json.Marshal(&s)
}

// UnmarshalJSON unmarshals from json
func (m *GitlabTrigger) UnmarshalJSON(data []byte) (e error) {
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
		ConnectionId     *string                           `json:"connectionId"`
		Id               *string                           `json:"id"`
		ProjectId        *string                           `json:"projectId"`
		CompartmentId    *string                           `json:"compartmentId"`
		Actions          []triggeraction                   `json:"actions"`
		TriggerUrl       *string                           `json:"triggerUrl"`
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

	m.ConnectionId = model.ConnectionId

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
	m.TriggerUrl = model.TriggerUrl

	return
}
