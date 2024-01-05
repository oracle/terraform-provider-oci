// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// LogAnalytics API
//
// The LogAnalytics API for the LogAnalytics service.
//

package loganalytics

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// IngestTimeRule An ingest time rule object.
type IngestTimeRule struct {

	// The log analytics entity OCID. This ID is a reference used by log analytics features and it represents
	// a resource that is provisioned and managed by the customer on their premises or on the cloud.
	Id *string `mandatory:"true" json:"id"`

	// Compartment Identifier OCID  (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm).
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The ingest time rule display name.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// Description for this resource.
	Description *string `mandatory:"false" json:"description"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// The date and time the resource was created, in the format defined by RFC3339.
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`

	// The date and time the resource was last updated, in the format defined by RFC3339.
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`

	// The current state of the ingest time rule.
	LifecycleState ConfigLifecycleStateEnum `mandatory:"false" json:"lifecycleState,omitempty"`

	// A flag indicating whether or not the ingest time rule is enabled.
	IsEnabled *bool `mandatory:"false" json:"isEnabled"`

	Conditions IngestTimeRuleCondition `mandatory:"false" json:"conditions"`

	// The action(s) to be performed if the ingest time rule condition(s) are satisfied.
	Actions []IngestTimeRuleAction `mandatory:"false" json:"actions"`
}

func (m IngestTimeRule) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m IngestTimeRule) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingConfigLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetConfigLifecycleStateEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UnmarshalJSON unmarshals from json
func (m *IngestTimeRule) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		Description    *string                           `json:"description"`
		FreeformTags   map[string]string                 `json:"freeformTags"`
		DefinedTags    map[string]map[string]interface{} `json:"definedTags"`
		TimeCreated    *common.SDKTime                   `json:"timeCreated"`
		TimeUpdated    *common.SDKTime                   `json:"timeUpdated"`
		LifecycleState ConfigLifecycleStateEnum          `json:"lifecycleState"`
		IsEnabled      *bool                             `json:"isEnabled"`
		Conditions     ingesttimerulecondition           `json:"conditions"`
		Actions        []ingesttimeruleaction            `json:"actions"`
		Id             *string                           `json:"id"`
		CompartmentId  *string                           `json:"compartmentId"`
		DisplayName    *string                           `json:"displayName"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.Description = model.Description

	m.FreeformTags = model.FreeformTags

	m.DefinedTags = model.DefinedTags

	m.TimeCreated = model.TimeCreated

	m.TimeUpdated = model.TimeUpdated

	m.LifecycleState = model.LifecycleState

	m.IsEnabled = model.IsEnabled

	nn, e = model.Conditions.UnmarshalPolymorphicJSON(model.Conditions.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.Conditions = nn.(IngestTimeRuleCondition)
	} else {
		m.Conditions = nil
	}

	m.Actions = make([]IngestTimeRuleAction, len(model.Actions))
	for i, n := range model.Actions {
		nn, e = n.UnmarshalPolymorphicJSON(n.JsonData)
		if e != nil {
			return e
		}
		if nn != nil {
			m.Actions[i] = nn.(IngestTimeRuleAction)
		} else {
			m.Actions[i] = nil
		}
	}
	m.Id = model.Id

	m.CompartmentId = model.CompartmentId

	m.DisplayName = model.DisplayName

	return
}
