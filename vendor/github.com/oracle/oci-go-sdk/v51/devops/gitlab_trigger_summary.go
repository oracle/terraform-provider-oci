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
	"github.com/oracle/oci-go-sdk/v51/common"
)

// GitlabTriggerSummary Summary of the Gitlab Trigger.
type GitlabTriggerSummary struct {

	// Unique identifier that is immutable on creation
	Id *string `mandatory:"true" json:"id"`

	// Project to which the Trigger belongs
	ProjectId *string `mandatory:"true" json:"projectId"`

	// Compartment to which the Trigger belongs
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

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
func (m GitlabTriggerSummary) GetId() *string {
	return m.Id
}

//GetDisplayName returns DisplayName
func (m GitlabTriggerSummary) GetDisplayName() *string {
	return m.DisplayName
}

//GetDescription returns Description
func (m GitlabTriggerSummary) GetDescription() *string {
	return m.Description
}

//GetProjectId returns ProjectId
func (m GitlabTriggerSummary) GetProjectId() *string {
	return m.ProjectId
}

//GetCompartmentId returns CompartmentId
func (m GitlabTriggerSummary) GetCompartmentId() *string {
	return m.CompartmentId
}

//GetTimeCreated returns TimeCreated
func (m GitlabTriggerSummary) GetTimeCreated() *common.SDKTime {
	return m.TimeCreated
}

//GetTimeUpdated returns TimeUpdated
func (m GitlabTriggerSummary) GetTimeUpdated() *common.SDKTime {
	return m.TimeUpdated
}

//GetLifecycleState returns LifecycleState
func (m GitlabTriggerSummary) GetLifecycleState() TriggerLifecycleStateEnum {
	return m.LifecycleState
}

//GetLifecycleDetails returns LifecycleDetails
func (m GitlabTriggerSummary) GetLifecycleDetails() *string {
	return m.LifecycleDetails
}

//GetFreeformTags returns FreeformTags
func (m GitlabTriggerSummary) GetFreeformTags() map[string]string {
	return m.FreeformTags
}

//GetDefinedTags returns DefinedTags
func (m GitlabTriggerSummary) GetDefinedTags() map[string]map[string]interface{} {
	return m.DefinedTags
}

//GetSystemTags returns SystemTags
func (m GitlabTriggerSummary) GetSystemTags() map[string]map[string]interface{} {
	return m.SystemTags
}

func (m GitlabTriggerSummary) String() string {
	return common.PointerString(m)
}

// MarshalJSON marshals to json representation
func (m GitlabTriggerSummary) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeGitlabTriggerSummary GitlabTriggerSummary
	s := struct {
		DiscriminatorParam string `json:"triggerSource"`
		MarshalTypeGitlabTriggerSummary
	}{
		"GITLAB",
		(MarshalTypeGitlabTriggerSummary)(m),
	}

	return json.Marshal(&s)
}
