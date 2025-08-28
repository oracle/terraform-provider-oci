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

// Schedule A repeating action. Examples:
// * Invoke a ML Pipeline Run once an hour.
// * Call ML Job Run every night at midnight.
type Schedule struct {

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the schedule.
	Id *string `mandatory:"true" json:"id"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment where you want to create the schedule.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the project associated with the schedule.
	ProjectId *string `mandatory:"true" json:"projectId"`

	// A user-friendly display name for the resource. Avoid entering confidential information.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The date and time the schedule was created.
	// Format is defined by RFC3339.
	// Example: `2022-08-05T01:02:29.600Z`
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the user who created the schedule.
	CreatedBy *string `mandatory:"true" json:"createdBy"`

	// The current state of the schedule.
	// Example: `ACTIVE`
	LifecycleState ScheduleLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	Trigger ScheduleTrigger `mandatory:"true" json:"trigger"`

	Action ScheduleAction `mandatory:"true" json:"action"`

	// A short description of the schedule.
	Description *string `mandatory:"false" json:"description"`

	// The last schedule execution time.
	// Format is defined by RFC3339.
	// Example: `2022-08-05T01:02:29.600Z`
	TimeLastScheduleRun *common.SDKTime `mandatory:"false" json:"timeLastScheduleRun"`

	// The next scheduled execution time for the schedule.
	// Format is defined by RFC3339.
	// Example: `2022-08-05T01:02:29.600Z`
	TimeNextScheduledRun *common.SDKTime `mandatory:"false" json:"timeNextScheduledRun"`

	// The date and time the schedule was updated.
	// Format is defined by RFC3339.
	// Example: `2022-09-05T01:02:29.600Z`
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`

	// A message describing the current state in more detail.
	LifecycleDetails *string `mandatory:"false" json:"lifecycleDetails"`

	// Details about the action performed by the last schedule execution.
	// Example: `Invoked ML Application trigger.`
	LastScheduleRunDetails *string `mandatory:"false" json:"lastScheduleRunDetails"`

	LogDetails *ScheduleLogDetails `mandatory:"false" json:"logDetails"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. See Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace. See Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// Usage of system tag keys. These predefined keys are scoped to namespaces.
	// Example: `{"orcl-cloud": {"free-tier-retained": "true"}}`
	SystemTags map[string]map[string]interface{} `mandatory:"false" json:"systemTags"`
}

func (m Schedule) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m Schedule) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingScheduleLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetScheduleLifecycleStateEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UnmarshalJSON unmarshals from json
func (m *Schedule) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		Description            *string                           `json:"description"`
		TimeLastScheduleRun    *common.SDKTime                   `json:"timeLastScheduleRun"`
		TimeNextScheduledRun   *common.SDKTime                   `json:"timeNextScheduledRun"`
		TimeUpdated            *common.SDKTime                   `json:"timeUpdated"`
		LifecycleDetails       *string                           `json:"lifecycleDetails"`
		LastScheduleRunDetails *string                           `json:"lastScheduleRunDetails"`
		LogDetails             *ScheduleLogDetails               `json:"logDetails"`
		FreeformTags           map[string]string                 `json:"freeformTags"`
		DefinedTags            map[string]map[string]interface{} `json:"definedTags"`
		SystemTags             map[string]map[string]interface{} `json:"systemTags"`
		Id                     *string                           `json:"id"`
		CompartmentId          *string                           `json:"compartmentId"`
		ProjectId              *string                           `json:"projectId"`
		DisplayName            *string                           `json:"displayName"`
		TimeCreated            *common.SDKTime                   `json:"timeCreated"`
		CreatedBy              *string                           `json:"createdBy"`
		LifecycleState         ScheduleLifecycleStateEnum        `json:"lifecycleState"`
		Trigger                scheduletrigger                   `json:"trigger"`
		Action                 scheduleaction                    `json:"action"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.Description = model.Description

	m.TimeLastScheduleRun = model.TimeLastScheduleRun

	m.TimeNextScheduledRun = model.TimeNextScheduledRun

	m.TimeUpdated = model.TimeUpdated

	m.LifecycleDetails = model.LifecycleDetails

	m.LastScheduleRunDetails = model.LastScheduleRunDetails

	m.LogDetails = model.LogDetails

	m.FreeformTags = model.FreeformTags

	m.DefinedTags = model.DefinedTags

	m.SystemTags = model.SystemTags

	m.Id = model.Id

	m.CompartmentId = model.CompartmentId

	m.ProjectId = model.ProjectId

	m.DisplayName = model.DisplayName

	m.TimeCreated = model.TimeCreated

	m.CreatedBy = model.CreatedBy

	m.LifecycleState = model.LifecycleState

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
