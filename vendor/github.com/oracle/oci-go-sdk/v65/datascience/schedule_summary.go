// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
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

// ScheduleSummary Summary representation of a schedule.
type ScheduleSummary struct {

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the schedule.
	Id *string `mandatory:"true" json:"id"`

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment that contains the schedule.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the project associated with the schedule.
	ProjectId *string `mandatory:"true" json:"projectId"`

	// The date and time the schedule was created.
	// Format is defined by RFC3339.
	// Example: `2022-08-05T01:02:29.600Z`
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the user who created the schedule.
	CreatedBy *string `mandatory:"true" json:"createdBy"`

	// The current state of the schedule.
	// Example: `ACTIVE`
	LifecycleState ScheduleLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// The name of the schedule.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// The date and time the schedule was updated.
	// Format is defined by RFC3339.
	// Example: `2022-08-05T01:02:29.600Z`
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`

	Trigger ScheduleTrigger `mandatory:"false" json:"trigger"`

	// A message describing the current state in more detail. For example, can be used to provide actionable information for a resource in Failed state.
	LifecycleDetails *string `mandatory:"false" json:"lifecycleDetails"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. See Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace. See Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// Usage of system tag keys. These predefined keys are scoped to namespaces.
	// Example: `{"orcl-cloud": {"free-tier-retained": "true"}}`
	SystemTags map[string]map[string]interface{} `mandatory:"false" json:"systemTags"`
}

func (m ScheduleSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ScheduleSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingScheduleLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetScheduleLifecycleStateEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UnmarshalJSON unmarshals from json
func (m *ScheduleSummary) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		DisplayName      *string                           `json:"displayName"`
		TimeUpdated      *common.SDKTime                   `json:"timeUpdated"`
		Trigger          scheduletrigger                   `json:"trigger"`
		LifecycleDetails *string                           `json:"lifecycleDetails"`
		FreeformTags     map[string]string                 `json:"freeformTags"`
		DefinedTags      map[string]map[string]interface{} `json:"definedTags"`
		SystemTags       map[string]map[string]interface{} `json:"systemTags"`
		Id               *string                           `json:"id"`
		CompartmentId    *string                           `json:"compartmentId"`
		ProjectId        *string                           `json:"projectId"`
		TimeCreated      *common.SDKTime                   `json:"timeCreated"`
		CreatedBy        *string                           `json:"createdBy"`
		LifecycleState   ScheduleLifecycleStateEnum        `json:"lifecycleState"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.DisplayName = model.DisplayName

	m.TimeUpdated = model.TimeUpdated

	nn, e = model.Trigger.UnmarshalPolymorphicJSON(model.Trigger.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.Trigger = nn.(ScheduleTrigger)
	} else {
		m.Trigger = nil
	}

	m.LifecycleDetails = model.LifecycleDetails

	m.FreeformTags = model.FreeformTags

	m.DefinedTags = model.DefinedTags

	m.SystemTags = model.SystemTags

	m.Id = model.Id

	m.CompartmentId = model.CompartmentId

	m.ProjectId = model.ProjectId

	m.TimeCreated = model.TimeCreated

	m.CreatedBy = model.CreatedBy

	m.LifecycleState = model.LifecycleState

	return
}
