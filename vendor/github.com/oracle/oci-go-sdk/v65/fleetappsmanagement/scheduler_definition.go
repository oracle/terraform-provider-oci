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

// SchedulerDefinition Definition of a Schedule.
type SchedulerDefinition struct {

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

	// The current state of the SchedulerDefinition.
	LifecycleState SchedulerDefinitionLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// A user-friendly description. To provide some insight about the resource.
	// Avoid entering confidential information.
	Description *string `mandatory:"false" json:"description"`

	// Associated region
	ResourceRegion *string `mandatory:"false" json:"resourceRegion"`

	// The time this resource was last updated. An RFC3339 formatted datetime string.
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`

	// The scheduled date for the next run of the Job.
	TimeOfNextRun *common.SDKTime `mandatory:"false" json:"timeOfNextRun"`

	Schedule Schedule `mandatory:"false" json:"schedule"`

	// Count of Action Groups affected by the Schedule.
	CountOfAffectedActionGroups *int `mandatory:"false" json:"countOfAffectedActionGroups"`

	// Count of Resources affected by the Schedule.
	CountOfAffectedResources *int `mandatory:"false" json:"countOfAffectedResources"`

	// Count of Targets affected by the Schedule.
	CountOfAffectedTargets *int `mandatory:"false" json:"countOfAffectedTargets"`

	// All products that are part of the schedule for PRODUCT ActionGroup Type.
	Products []string `mandatory:"false" json:"products"`

	// All LifeCycle Operations that are part of the schedule.
	LifecycleOperations []string `mandatory:"false" json:"lifecycleOperations"`

	// Action Groups associated with the Schedule.
	ActionGroups []ActionGroup `mandatory:"false" json:"actionGroups"`

	// Runbooks.
	RunBooks []OperationRunbook `mandatory:"false" json:"runBooks"`

	// A message describing the current state in more detail. For example, can be used to provide actionable information for a resource in Failed state.
	LifecycleDetails *string `mandatory:"false" json:"lifecycleDetails"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// System tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"orcl-cloud": {"free-tier-retained": "true"}}`
	SystemTags map[string]map[string]interface{} `mandatory:"false" json:"systemTags"`
}

func (m SchedulerDefinition) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m SchedulerDefinition) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingSchedulerDefinitionLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetSchedulerDefinitionLifecycleStateEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UnmarshalJSON unmarshals from json
func (m *SchedulerDefinition) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		Description                 *string                               `json:"description"`
		ResourceRegion              *string                               `json:"resourceRegion"`
		TimeUpdated                 *common.SDKTime                       `json:"timeUpdated"`
		TimeOfNextRun               *common.SDKTime                       `json:"timeOfNextRun"`
		Schedule                    schedule                              `json:"schedule"`
		CountOfAffectedActionGroups *int                                  `json:"countOfAffectedActionGroups"`
		CountOfAffectedResources    *int                                  `json:"countOfAffectedResources"`
		CountOfAffectedTargets      *int                                  `json:"countOfAffectedTargets"`
		Products                    []string                              `json:"products"`
		LifecycleOperations         []string                              `json:"lifecycleOperations"`
		ActionGroups                []actiongroup                         `json:"actionGroups"`
		RunBooks                    []OperationRunbook                    `json:"runBooks"`
		LifecycleDetails            *string                               `json:"lifecycleDetails"`
		FreeformTags                map[string]string                     `json:"freeformTags"`
		DefinedTags                 map[string]map[string]interface{}     `json:"definedTags"`
		SystemTags                  map[string]map[string]interface{}     `json:"systemTags"`
		Id                          *string                               `json:"id"`
		DisplayName                 *string                               `json:"displayName"`
		CompartmentId               *string                               `json:"compartmentId"`
		TimeCreated                 *common.SDKTime                       `json:"timeCreated"`
		LifecycleState              SchedulerDefinitionLifecycleStateEnum `json:"lifecycleState"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.Description = model.Description

	m.ResourceRegion = model.ResourceRegion

	m.TimeUpdated = model.TimeUpdated

	m.TimeOfNextRun = model.TimeOfNextRun

	nn, e = model.Schedule.UnmarshalPolymorphicJSON(model.Schedule.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.Schedule = nn.(Schedule)
	} else {
		m.Schedule = nil
	}

	m.CountOfAffectedActionGroups = model.CountOfAffectedActionGroups

	m.CountOfAffectedResources = model.CountOfAffectedResources

	m.CountOfAffectedTargets = model.CountOfAffectedTargets

	m.Products = make([]string, len(model.Products))
	copy(m.Products, model.Products)
	m.LifecycleOperations = make([]string, len(model.LifecycleOperations))
	copy(m.LifecycleOperations, model.LifecycleOperations)
	m.ActionGroups = make([]ActionGroup, len(model.ActionGroups))
	for i, n := range model.ActionGroups {
		nn, e = n.UnmarshalPolymorphicJSON(n.JsonData)
		if e != nil {
			return e
		}
		if nn != nil {
			m.ActionGroups[i] = nn.(ActionGroup)
		} else {
			m.ActionGroups[i] = nil
		}
	}
	m.RunBooks = make([]OperationRunbook, len(model.RunBooks))
	copy(m.RunBooks, model.RunBooks)
	m.LifecycleDetails = model.LifecycleDetails

	m.FreeformTags = model.FreeformTags

	m.DefinedTags = model.DefinedTags

	m.SystemTags = model.SystemTags

	m.Id = model.Id

	m.DisplayName = model.DisplayName

	m.CompartmentId = model.CompartmentId

	m.TimeCreated = model.TimeCreated

	m.LifecycleState = model.LifecycleState

	return
}

// SchedulerDefinitionLifecycleStateEnum Enum with underlying type: string
type SchedulerDefinitionLifecycleStateEnum string

// Set of constants representing the allowable values for SchedulerDefinitionLifecycleStateEnum
const (
	SchedulerDefinitionLifecycleStateCreating SchedulerDefinitionLifecycleStateEnum = "CREATING"
	SchedulerDefinitionLifecycleStateUpdating SchedulerDefinitionLifecycleStateEnum = "UPDATING"
	SchedulerDefinitionLifecycleStateActive   SchedulerDefinitionLifecycleStateEnum = "ACTIVE"
	SchedulerDefinitionLifecycleStateDeleting SchedulerDefinitionLifecycleStateEnum = "DELETING"
	SchedulerDefinitionLifecycleStateDeleted  SchedulerDefinitionLifecycleStateEnum = "DELETED"
	SchedulerDefinitionLifecycleStateFailed   SchedulerDefinitionLifecycleStateEnum = "FAILED"
	SchedulerDefinitionLifecycleStateInactive SchedulerDefinitionLifecycleStateEnum = "INACTIVE"
)

var mappingSchedulerDefinitionLifecycleStateEnum = map[string]SchedulerDefinitionLifecycleStateEnum{
	"CREATING": SchedulerDefinitionLifecycleStateCreating,
	"UPDATING": SchedulerDefinitionLifecycleStateUpdating,
	"ACTIVE":   SchedulerDefinitionLifecycleStateActive,
	"DELETING": SchedulerDefinitionLifecycleStateDeleting,
	"DELETED":  SchedulerDefinitionLifecycleStateDeleted,
	"FAILED":   SchedulerDefinitionLifecycleStateFailed,
	"INACTIVE": SchedulerDefinitionLifecycleStateInactive,
}

var mappingSchedulerDefinitionLifecycleStateEnumLowerCase = map[string]SchedulerDefinitionLifecycleStateEnum{
	"creating": SchedulerDefinitionLifecycleStateCreating,
	"updating": SchedulerDefinitionLifecycleStateUpdating,
	"active":   SchedulerDefinitionLifecycleStateActive,
	"deleting": SchedulerDefinitionLifecycleStateDeleting,
	"deleted":  SchedulerDefinitionLifecycleStateDeleted,
	"failed":   SchedulerDefinitionLifecycleStateFailed,
	"inactive": SchedulerDefinitionLifecycleStateInactive,
}

// GetSchedulerDefinitionLifecycleStateEnumValues Enumerates the set of values for SchedulerDefinitionLifecycleStateEnum
func GetSchedulerDefinitionLifecycleStateEnumValues() []SchedulerDefinitionLifecycleStateEnum {
	values := make([]SchedulerDefinitionLifecycleStateEnum, 0)
	for _, v := range mappingSchedulerDefinitionLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetSchedulerDefinitionLifecycleStateEnumStringValues Enumerates the set of values in String for SchedulerDefinitionLifecycleStateEnum
func GetSchedulerDefinitionLifecycleStateEnumStringValues() []string {
	return []string{
		"CREATING",
		"UPDATING",
		"ACTIVE",
		"DELETING",
		"DELETED",
		"FAILED",
		"INACTIVE",
	}
}

// GetMappingSchedulerDefinitionLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingSchedulerDefinitionLifecycleStateEnum(val string) (SchedulerDefinitionLifecycleStateEnum, bool) {
	enum, ok := mappingSchedulerDefinitionLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
