// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Science API
//
// Use the Data Science API to organize your data science work, access data and computing resources, and build, train, deploy and manage models and model deployments. For more information, see Data Science (https://docs.oracle.com/iaas/data-science/using/data-science.htm).
//

package datascience

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// CreateScheduleDetails Creation details for a new schedule.
type CreateScheduleDetails struct {

	// A user-friendly name. Avoid entering confidential information.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the project to associate the schedule with.
	ProjectId *string `mandatory:"true" json:"projectId"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment that contains the schedule.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	Trigger ScheduleTrigger `mandatory:"true" json:"trigger"`

	Action ScheduleAction `mandatory:"true" json:"action"`

	// A short description of the schedule.
	Description *string `mandatory:"false" json:"description"`

	LogDetails *ScheduleLogDetails `mandatory:"false" json:"logDetails"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. See Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace. See Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m CreateScheduleDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CreateScheduleDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UnmarshalJSON unmarshals from json
func (m *CreateScheduleDetails) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		Description   *string                           `json:"description"`
		LogDetails    *ScheduleLogDetails               `json:"logDetails"`
		FreeformTags  map[string]string                 `json:"freeformTags"`
		DefinedTags   map[string]map[string]interface{} `json:"definedTags"`
		DisplayName   *string                           `json:"displayName"`
		ProjectId     *string                           `json:"projectId"`
		CompartmentId *string                           `json:"compartmentId"`
		Trigger       scheduletrigger                   `json:"trigger"`
		Action        scheduleaction                    `json:"action"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.Description = model.Description

	m.LogDetails = model.LogDetails

	m.FreeformTags = model.FreeformTags

	m.DefinedTags = model.DefinedTags

	m.DisplayName = model.DisplayName

	m.ProjectId = model.ProjectId

	m.CompartmentId = model.CompartmentId

	nn, e = model.Trigger.UnmarshalPolymorphicJSON(model.Trigger.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.Trigger = nn.(ScheduleTrigger)
	} else {
		m.Trigger = nil
	}

	nn, e = model.Action.UnmarshalPolymorphicJSON(model.Action.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.Action = nn.(ScheduleAction)
	} else {
		m.Action = nil
	}

	return
}
