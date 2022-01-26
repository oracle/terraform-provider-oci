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

// Trigger Trigger the deployment pipeline to deploy the artifact.
type Trigger interface {

	// Unique identifier that is immutable on creation.
	GetId() *string

	// The OCID of the DevOps project to which the trigger belongs to.
	GetProjectId() *string

	// The OCID of the compartment that contains the trigger.
	GetCompartmentId() *string

	// The list of actions that are to be performed for this trigger.
	GetActions() []TriggerAction

	// Trigger display name. Avoid entering confidential information.
	GetDisplayName() *string

	// Description about the trigger.
	GetDescription() *string

	// The time the trigger was created. Format defined by RFC3339 (https://datatracker.ietf.org/doc/html/rfc3339).
	GetTimeCreated() *common.SDKTime

	// The time the trigger was updated. Format defined by RFC3339 (https://datatracker.ietf.org/doc/html/rfc3339).
	GetTimeUpdated() *common.SDKTime

	// The current state of the trigger.
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

type trigger struct {
	JsonData         []byte
	Id               *string                           `mandatory:"true" json:"id"`
	ProjectId        *string                           `mandatory:"true" json:"projectId"`
	CompartmentId    *string                           `mandatory:"true" json:"compartmentId"`
	Actions          json.RawMessage                   `mandatory:"true" json:"actions"`
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
func (m *trigger) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalertrigger trigger
	s := struct {
		Model Unmarshalertrigger
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.Id = s.Model.Id
	m.ProjectId = s.Model.ProjectId
	m.CompartmentId = s.Model.CompartmentId
	m.Actions = s.Model.Actions
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
func (m *trigger) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.TriggerSource {
	case "GITLAB":
		mm := GitlabTrigger{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "GITHUB":
		mm := GithubTrigger{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "DEVOPS_CODE_REPOSITORY":
		mm := DevopsCodeRepositoryTrigger{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		return *m, nil
	}
}

//GetId returns Id
func (m trigger) GetId() *string {
	return m.Id
}

//GetProjectId returns ProjectId
func (m trigger) GetProjectId() *string {
	return m.ProjectId
}

//GetCompartmentId returns CompartmentId
func (m trigger) GetCompartmentId() *string {
	return m.CompartmentId
}

//GetActions returns Actions
func (m trigger) GetActions() json.RawMessage {
	return m.Actions
}

//GetDisplayName returns DisplayName
func (m trigger) GetDisplayName() *string {
	return m.DisplayName
}

//GetDescription returns Description
func (m trigger) GetDescription() *string {
	return m.Description
}

//GetTimeCreated returns TimeCreated
func (m trigger) GetTimeCreated() *common.SDKTime {
	return m.TimeCreated
}

//GetTimeUpdated returns TimeUpdated
func (m trigger) GetTimeUpdated() *common.SDKTime {
	return m.TimeUpdated
}

//GetLifecycleState returns LifecycleState
func (m trigger) GetLifecycleState() TriggerLifecycleStateEnum {
	return m.LifecycleState
}

//GetLifecycleDetails returns LifecycleDetails
func (m trigger) GetLifecycleDetails() *string {
	return m.LifecycleDetails
}

//GetFreeformTags returns FreeformTags
func (m trigger) GetFreeformTags() map[string]string {
	return m.FreeformTags
}

//GetDefinedTags returns DefinedTags
func (m trigger) GetDefinedTags() map[string]map[string]interface{} {
	return m.DefinedTags
}

//GetSystemTags returns SystemTags
func (m trigger) GetSystemTags() map[string]map[string]interface{} {
	return m.SystemTags
}

func (m trigger) String() string {
	return common.PointerString(m)
}

// TriggerLifecycleStateEnum Enum with underlying type: string
type TriggerLifecycleStateEnum string

// Set of constants representing the allowable values for TriggerLifecycleStateEnum
const (
	TriggerLifecycleStateActive TriggerLifecycleStateEnum = "ACTIVE"
)

var mappingTriggerLifecycleState = map[string]TriggerLifecycleStateEnum{
	"ACTIVE": TriggerLifecycleStateActive,
}

// GetTriggerLifecycleStateEnumValues Enumerates the set of values for TriggerLifecycleStateEnum
func GetTriggerLifecycleStateEnumValues() []TriggerLifecycleStateEnum {
	values := make([]TriggerLifecycleStateEnum, 0)
	for _, v := range mappingTriggerLifecycleState {
		values = append(values, v)
	}
	return values
}

// TriggerTriggerSourceEnum Enum with underlying type: string
type TriggerTriggerSourceEnum string

// Set of constants representing the allowable values for TriggerTriggerSourceEnum
const (
	TriggerTriggerSourceGithub               TriggerTriggerSourceEnum = "GITHUB"
	TriggerTriggerSourceGitlab               TriggerTriggerSourceEnum = "GITLAB"
	TriggerTriggerSourceDevopsCodeRepository TriggerTriggerSourceEnum = "DEVOPS_CODE_REPOSITORY"
)

var mappingTriggerTriggerSource = map[string]TriggerTriggerSourceEnum{
	"GITHUB":                 TriggerTriggerSourceGithub,
	"GITLAB":                 TriggerTriggerSourceGitlab,
	"DEVOPS_CODE_REPOSITORY": TriggerTriggerSourceDevopsCodeRepository,
}

// GetTriggerTriggerSourceEnumValues Enumerates the set of values for TriggerTriggerSourceEnum
func GetTriggerTriggerSourceEnumValues() []TriggerTriggerSourceEnum {
	values := make([]TriggerTriggerSourceEnum, 0)
	for _, v := range mappingTriggerTriggerSource {
		values = append(values, v)
	}
	return values
}
