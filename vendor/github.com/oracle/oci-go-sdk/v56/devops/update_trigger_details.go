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

// UpdateTriggerDetails The information to be updated.
type UpdateTriggerDetails interface {

	// Trigger display name. Avoid entering confidential information.
	GetDisplayName() *string

	// Optional description about the trigger.
	GetDescription() *string

	// The list of actions that are to be performed for this trigger.
	GetActions() []TriggerAction

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.  See Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm). Example: `{"bar-key": "value"}`
	GetFreeformTags() map[string]string

	// Defined tags for this resource. Each key is predefined and scoped to a namespace. See Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm). Example: `{"foo-namespace": {"bar-key": "value"}}`
	GetDefinedTags() map[string]map[string]interface{}
}

type updatetriggerdetails struct {
	JsonData      []byte
	DisplayName   *string                           `mandatory:"false" json:"displayName"`
	Description   *string                           `mandatory:"false" json:"description"`
	Actions       json.RawMessage                   `mandatory:"false" json:"actions"`
	FreeformTags  map[string]string                 `mandatory:"false" json:"freeformTags"`
	DefinedTags   map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
	TriggerSource string                            `json:"triggerSource"`
}

// UnmarshalJSON unmarshals json
func (m *updatetriggerdetails) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerupdatetriggerdetails updatetriggerdetails
	s := struct {
		Model Unmarshalerupdatetriggerdetails
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.DisplayName = s.Model.DisplayName
	m.Description = s.Model.Description
	m.Actions = s.Model.Actions
	m.FreeformTags = s.Model.FreeformTags
	m.DefinedTags = s.Model.DefinedTags
	m.TriggerSource = s.Model.TriggerSource

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *updatetriggerdetails) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.TriggerSource {
	case "DEVOPS_CODE_REPOSITORY":
		mm := UpdateDevopsCodeRepositoryTriggerDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "GITHUB":
		mm := UpdateGithubTriggerDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "GITLAB":
		mm := UpdateGitlabTriggerDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		return *m, nil
	}
}

//GetDisplayName returns DisplayName
func (m updatetriggerdetails) GetDisplayName() *string {
	return m.DisplayName
}

//GetDescription returns Description
func (m updatetriggerdetails) GetDescription() *string {
	return m.Description
}

//GetActions returns Actions
func (m updatetriggerdetails) GetActions() json.RawMessage {
	return m.Actions
}

//GetFreeformTags returns FreeformTags
func (m updatetriggerdetails) GetFreeformTags() map[string]string {
	return m.FreeformTags
}

//GetDefinedTags returns DefinedTags
func (m updatetriggerdetails) GetDefinedTags() map[string]map[string]interface{} {
	return m.DefinedTags
}

func (m updatetriggerdetails) String() string {
	return common.PointerString(m)
}
