// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Fleet Application Management Service API
//
// Fleet Application Management provides a centralized platform to help you automate resource management tasks, validate patch compliance, and enhance operational efficiency across an enterprise.
//

package fleetappsmanagement

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// SchedulerJob A SchedulerJob resource.
type SchedulerJob struct {

	// The OCID of the resource.
	Id *string `mandatory:"true" json:"id"`

	// A user-friendly name. Does not have to be unique, and it's changeable.
	// Avoid entering confidential information.
	// Example: `My new resource`
	DisplayName *string `mandatory:"true" json:"displayName"`

	// Compartment OCID
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The time this resource was created. An RFC3339 formatted datetime string.
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The current state of the SchedulerJob.
	LifecycleState SchedulerJobLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"true" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"true" json:"definedTags"`

	// The time this resource was last updated. An RFC3339 formatted datetime string.
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`

	// The scheduled date and time for the Job.
	TimeScheduled *common.SDKTime `mandatory:"false" json:"timeScheduled"`

	// Actual start date and time for the Job.
	TimeStarted *common.SDKTime `mandatory:"false" json:"timeStarted"`

	// Actual end date and time for the Job.
	TimeEnded *common.SDKTime `mandatory:"false" json:"timeEnded"`

	// Count of Action Groups affected by the Schedule.
	// An action group is a group of entities grouped for a lifecycle activity.
	// Example - A Fleet will be considered an Action Group for any lifecycle activity.
	CountOfAffectedActionGroups *int `mandatory:"false" json:"countOfAffectedActionGroups"`

	// Count of resources affected by the schedule.
	CountOfAffectedResources *int `mandatory:"false" json:"countOfAffectedResources"`

	// Count of targets affected by the schedule.
	CountOfAffectedTargets *int `mandatory:"false" json:"countOfAffectedTargets"`

	// All products that are part of the schedule for a PRODUCT action group type.
	Products []string `mandatory:"false" json:"products"`

	// All LifeCycle Operations that are part of the schedule.
	LifecycleOperations []string `mandatory:"false" json:"lifecycleOperations"`

	// Action Groups associated with the Schedule.
	ActionGroups []ActionGroupDetails `mandatory:"false" json:"actionGroups"`

	SchedulerDefinition *AssociatedSchedulerDefinition `mandatory:"false" json:"schedulerDefinition"`

	// A message describing the current state in more detail. For example, can be used to provide actionable information for a resource in Failed state.
	LifecycleDetails *string `mandatory:"false" json:"lifecycleDetails"`

	// System tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"orcl-cloud": {"free-tier-retained": "true"}}`
	SystemTags map[string]map[string]interface{} `mandatory:"false" json:"systemTags"`
}

func (m SchedulerJob) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m SchedulerJob) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingSchedulerJobLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetSchedulerJobLifecycleStateEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UnmarshalJSON unmarshals from json
func (m *SchedulerJob) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		TimeUpdated                 *common.SDKTime                   `json:"timeUpdated"`
		TimeScheduled               *common.SDKTime                   `json:"timeScheduled"`
		TimeStarted                 *common.SDKTime                   `json:"timeStarted"`
		TimeEnded                   *common.SDKTime                   `json:"timeEnded"`
		CountOfAffectedActionGroups *int                              `json:"countOfAffectedActionGroups"`
		CountOfAffectedResources    *int                              `json:"countOfAffectedResources"`
		CountOfAffectedTargets      *int                              `json:"countOfAffectedTargets"`
		Products                    []string                          `json:"products"`
		LifecycleOperations         []string                          `json:"lifecycleOperations"`
		ActionGroups                []actiongroupdetails              `json:"actionGroups"`
		SchedulerDefinition         *AssociatedSchedulerDefinition    `json:"schedulerDefinition"`
		LifecycleDetails            *string                           `json:"lifecycleDetails"`
		SystemTags                  map[string]map[string]interface{} `json:"systemTags"`
		Id                          *string                           `json:"id"`
		DisplayName                 *string                           `json:"displayName"`
		CompartmentId               *string                           `json:"compartmentId"`
		TimeCreated                 *common.SDKTime                   `json:"timeCreated"`
		LifecycleState              SchedulerJobLifecycleStateEnum    `json:"lifecycleState"`
		FreeformTags                map[string]string                 `json:"freeformTags"`
		DefinedTags                 map[string]map[string]interface{} `json:"definedTags"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.TimeUpdated = model.TimeUpdated

	m.TimeScheduled = model.TimeScheduled

	m.TimeStarted = model.TimeStarted

	m.TimeEnded = model.TimeEnded

	m.CountOfAffectedActionGroups = model.CountOfAffectedActionGroups

	m.CountOfAffectedResources = model.CountOfAffectedResources

	m.CountOfAffectedTargets = model.CountOfAffectedTargets

	m.Products = make([]string, len(model.Products))
	copy(m.Products, model.Products)
	m.LifecycleOperations = make([]string, len(model.LifecycleOperations))
	copy(m.LifecycleOperations, model.LifecycleOperations)
	m.ActionGroups = make([]ActionGroupDetails, len(model.ActionGroups))
	for i, n := range model.ActionGroups {
		nn, e = n.UnmarshalPolymorphicJSON(n.JsonData)
		if e != nil {
			return e
		}
		if nn != nil {
			m.ActionGroups[i] = nn.(ActionGroupDetails)
		} else {
			m.ActionGroups[i] = nil
		}
	}
	m.SchedulerDefinition = model.SchedulerDefinition

	m.LifecycleDetails = model.LifecycleDetails

	m.SystemTags = model.SystemTags

	m.Id = model.Id

	m.DisplayName = model.DisplayName

	m.CompartmentId = model.CompartmentId

	m.TimeCreated = model.TimeCreated

	m.LifecycleState = model.LifecycleState

	m.FreeformTags = model.FreeformTags

	m.DefinedTags = model.DefinedTags

	return
}

// SchedulerJobLifecycleStateEnum Enum with underlying type: string
type SchedulerJobLifecycleStateEnum string

// Set of constants representing the allowable values for SchedulerJobLifecycleStateEnum
const (
	SchedulerJobLifecycleStateAccepted       SchedulerJobLifecycleStateEnum = "ACCEPTED"
	SchedulerJobLifecycleStateInProgress     SchedulerJobLifecycleStateEnum = "IN_PROGRESS"
	SchedulerJobLifecycleStateWaiting        SchedulerJobLifecycleStateEnum = "WAITING"
	SchedulerJobLifecycleStateFailed         SchedulerJobLifecycleStateEnum = "FAILED"
	SchedulerJobLifecycleStateSucceeded      SchedulerJobLifecycleStateEnum = "SUCCEEDED"
	SchedulerJobLifecycleStateCanceled       SchedulerJobLifecycleStateEnum = "CANCELED"
	SchedulerJobLifecycleStateNeedsAttention SchedulerJobLifecycleStateEnum = "NEEDS_ATTENTION"
)

var mappingSchedulerJobLifecycleStateEnum = map[string]SchedulerJobLifecycleStateEnum{
	"ACCEPTED":        SchedulerJobLifecycleStateAccepted,
	"IN_PROGRESS":     SchedulerJobLifecycleStateInProgress,
	"WAITING":         SchedulerJobLifecycleStateWaiting,
	"FAILED":          SchedulerJobLifecycleStateFailed,
	"SUCCEEDED":       SchedulerJobLifecycleStateSucceeded,
	"CANCELED":        SchedulerJobLifecycleStateCanceled,
	"NEEDS_ATTENTION": SchedulerJobLifecycleStateNeedsAttention,
}

var mappingSchedulerJobLifecycleStateEnumLowerCase = map[string]SchedulerJobLifecycleStateEnum{
	"accepted":        SchedulerJobLifecycleStateAccepted,
	"in_progress":     SchedulerJobLifecycleStateInProgress,
	"waiting":         SchedulerJobLifecycleStateWaiting,
	"failed":          SchedulerJobLifecycleStateFailed,
	"succeeded":       SchedulerJobLifecycleStateSucceeded,
	"canceled":        SchedulerJobLifecycleStateCanceled,
	"needs_attention": SchedulerJobLifecycleStateNeedsAttention,
}

// GetSchedulerJobLifecycleStateEnumValues Enumerates the set of values for SchedulerJobLifecycleStateEnum
func GetSchedulerJobLifecycleStateEnumValues() []SchedulerJobLifecycleStateEnum {
	values := make([]SchedulerJobLifecycleStateEnum, 0)
	for _, v := range mappingSchedulerJobLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetSchedulerJobLifecycleStateEnumStringValues Enumerates the set of values in String for SchedulerJobLifecycleStateEnum
func GetSchedulerJobLifecycleStateEnumStringValues() []string {
	return []string{
		"ACCEPTED",
		"IN_PROGRESS",
		"WAITING",
		"FAILED",
		"SUCCEEDED",
		"CANCELED",
		"NEEDS_ATTENTION",
	}
}

// GetMappingSchedulerJobLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingSchedulerJobLifecycleStateEnum(val string) (SchedulerJobLifecycleStateEnum, bool) {
	enum, ok := mappingSchedulerJobLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
