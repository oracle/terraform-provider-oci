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

// TriggerSummary Summary of the Trigger.
type TriggerSummary interface {

	// Unique identifier that is immutable on creation
	GetId() *string

	// Project to which the Trigger belongs
	GetProjectId() *string

	// Compartment to which the Trigger belongs
	GetCompartmentId() *string

	// Name for Trigger.
	GetDisplayName() *string

	// Description about the Trigger
	GetDescription() *string

	// The time the the Trigger was created. An RFC3339 formatted datetime string
	GetTimeCreated() *common.SDKTime

	// The time the Trigger was updated. An RFC3339 formatted datetime string
	GetTimeUpdated() *common.SDKTime

	// The current state of the Trigger.
	GetLifecycleState() TriggerLifecycleStateEnum

	// A message describing the current state in more detail. For example, can be used to provide actionable information for a resource in Failed state.
	GetLifecycleDetails() *string

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.  See Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm). Example: `{"bar-key": "value"}`
	GetFreeformTags() map[string]string

	// Defined tags for this resource. Each key is predefined and scoped to a namespace. See Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm). Example: `{"foo-namespace": {"bar-key": "value"}}`
	GetDefinedTags() map[string]map[string]interface{}

	// Usage of system tag keys. These predefined keys are scoped to namespaces. See Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm). Example: `{"orcl-cloud": {"free-tier-retained": "true"}}`
	GetSystemTags() map[string]map[string]interface{}
}

type triggersummary struct {
	JsonData         []byte
	Id               *string                           `mandatory:"true" json:"id"`
	ProjectId        *string                           `mandatory:"true" json:"projectId"`
	CompartmentId    *string                           `mandatory:"true" json:"compartmentId"`
	DisplayName      *string                           `mandatory:"false" json:"displayName"`
	Description      *string                           `mandatory:"false" json:"description"`
	TimeCreated      *common.SDKTime                   `mandatory:"false" json:"timeCreated"`
	TimeUpdated      *common.SDKTime                   `mandatory:"false" json:"timeUpdated"`
	LifecycleState   TriggerLifecycleStateEnum         `mandatory:"false" json:"lifecycleState,omitempty"`
	LifecycleDetails *string                           `mandatory:"false" json:"lifecycleDetails"`
	FreeformTags     map[string]string                 `mandatory:"false" json:"freeformTags"`
	DefinedTags      map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
	SystemTags       map[string]map[string]interface{} `mandatory:"false" json:"systemTags"`
	TriggerSource    string                            `json:"triggerSource"`
}

// UnmarshalJSON unmarshals json
func (m *triggersummary) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalertriggersummary triggersummary
	s := struct {
		Model Unmarshalertriggersummary
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.Id = s.Model.Id
	m.ProjectId = s.Model.ProjectId
	m.CompartmentId = s.Model.CompartmentId
	m.DisplayName = s.Model.DisplayName
	m.Description = s.Model.Description
	m.TimeCreated = s.Model.TimeCreated
	m.TimeUpdated = s.Model.TimeUpdated
	m.LifecycleState = s.Model.LifecycleState
	m.LifecycleDetails = s.Model.LifecycleDetails
	m.FreeformTags = s.Model.FreeformTags
	m.DefinedTags = s.Model.DefinedTags
	m.SystemTags = s.Model.SystemTags
	m.TriggerSource = s.Model.TriggerSource

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *triggersummary) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.TriggerSource {
	case "GITLAB":
		mm := GitlabTriggerSummary{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "GITHUB":
		mm := GithubTriggerSummary{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "DEVOPS_CODE_REPOSITORY":
		mm := DevopsCodeRepositoryTriggerSummary{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		return *m, nil
	}
}

//GetId returns Id
func (m triggersummary) GetId() *string {
	return m.Id
}

//GetProjectId returns ProjectId
func (m triggersummary) GetProjectId() *string {
	return m.ProjectId
}

//GetCompartmentId returns CompartmentId
func (m triggersummary) GetCompartmentId() *string {
	return m.CompartmentId
}

//GetDisplayName returns DisplayName
func (m triggersummary) GetDisplayName() *string {
	return m.DisplayName
}

//GetDescription returns Description
func (m triggersummary) GetDescription() *string {
	return m.Description
}

//GetTimeCreated returns TimeCreated
func (m triggersummary) GetTimeCreated() *common.SDKTime {
	return m.TimeCreated
}

//GetTimeUpdated returns TimeUpdated
func (m triggersummary) GetTimeUpdated() *common.SDKTime {
	return m.TimeUpdated
}

//GetLifecycleState returns LifecycleState
func (m triggersummary) GetLifecycleState() TriggerLifecycleStateEnum {
	return m.LifecycleState
}

//GetLifecycleDetails returns LifecycleDetails
func (m triggersummary) GetLifecycleDetails() *string {
	return m.LifecycleDetails
}

//GetFreeformTags returns FreeformTags
func (m triggersummary) GetFreeformTags() map[string]string {
	return m.FreeformTags
}

//GetDefinedTags returns DefinedTags
func (m triggersummary) GetDefinedTags() map[string]map[string]interface{} {
	return m.DefinedTags
}

//GetSystemTags returns SystemTags
func (m triggersummary) GetSystemTags() map[string]map[string]interface{} {
	return m.SystemTags
}

func (m triggersummary) String() string {
	return common.PointerString(m)
}
