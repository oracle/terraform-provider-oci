// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Fleet Application Management Service API
//
// Fleet Application Management provides a centralized platform to help you automate resource management tasks, validate patch compliance, and enhance operational efficiency across an enterprise.
//

package fleetappsmanagement

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// SchedulerExecutionSummary Summary of the Scheduler Executions.
type SchedulerExecutionSummary struct {

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

	// The current state of the Scheduler Execution.
	LifecycleState SchedulerExecutionSummaryLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"true" json:"freeformTags"`

	// Name of the compartment in which resource exist.
	CompartmentName *string `mandatory:"false" json:"compartmentName"`

	// The time this resource was last updated. An RFC3339 formatted datetime string.
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`

	// The scheduled date and time for the Job.
	TimeScheduled *common.SDKTime `mandatory:"false" json:"timeScheduled"`

	// Action Group associated with the Schedule.
	ActivityId *string `mandatory:"false" json:"activityId"`

	// FleetId associated with the Schedule.
	ResourceId *string `mandatory:"false" json:"resourceId"`

	// SchedulerJobId associated with the Schedule.
	SchedulerJobId *string `mandatory:"false" json:"schedulerJobId"`

	// Display Name of the Fleet associated with the Schedule.
	ResourceDisplayName *string `mandatory:"false" json:"resourceDisplayName"`

	// RunbookId associated with the Schedule.
	RunbookId *string `mandatory:"false" json:"runbookId"`

	// Name of the Runbook version associated with the Schedule.
	RunbookVersionName *string `mandatory:"false" json:"runbookVersionName"`

	// Display name of Runbook associated with the Schedule.
	RunbookDisplayName *string `mandatory:"false" json:"runbookDisplayName"`

	// Latest Runbook version available.
	LatestRunbookVersionName *string `mandatory:"false" json:"latestRunbookVersionName"`

	SchedulerDefinition *AssociatedSchedulerDefinition `mandatory:"false" json:"schedulerDefinition"`

	// Actual start date and time for the Execution.
	TimeStarted *common.SDKTime `mandatory:"false" json:"timeStarted"`

	// Actual end date and time for the Execution.
	TimeEnded *common.SDKTime `mandatory:"false" json:"timeEnded"`

	// A message describing the current state in more detail. For example, can be used to provide actionable information for a resource in Failed state.
	LifecycleDetails *string `mandatory:"false" json:"lifecycleDetails"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// System tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"orcl-cloud": {"free-tier-retained": "true"}}`
	SystemTags map[string]map[string]interface{} `mandatory:"false" json:"systemTags"`
}

func (m SchedulerExecutionSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m SchedulerExecutionSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingSchedulerExecutionSummaryLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetSchedulerExecutionSummaryLifecycleStateEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// SchedulerExecutionSummaryLifecycleStateEnum Enum with underlying type: string
type SchedulerExecutionSummaryLifecycleStateEnum string

// Set of constants representing the allowable values for SchedulerExecutionSummaryLifecycleStateEnum
const (
	SchedulerExecutionSummaryLifecycleStateActive   SchedulerExecutionSummaryLifecycleStateEnum = "ACTIVE"
	SchedulerExecutionSummaryLifecycleStateDeleted  SchedulerExecutionSummaryLifecycleStateEnum = "DELETED"
	SchedulerExecutionSummaryLifecycleStateFailed   SchedulerExecutionSummaryLifecycleStateEnum = "FAILED"
	SchedulerExecutionSummaryLifecycleStateInactive SchedulerExecutionSummaryLifecycleStateEnum = "INACTIVE"
)

var mappingSchedulerExecutionSummaryLifecycleStateEnum = map[string]SchedulerExecutionSummaryLifecycleStateEnum{
	"ACTIVE":   SchedulerExecutionSummaryLifecycleStateActive,
	"DELETED":  SchedulerExecutionSummaryLifecycleStateDeleted,
	"FAILED":   SchedulerExecutionSummaryLifecycleStateFailed,
	"INACTIVE": SchedulerExecutionSummaryLifecycleStateInactive,
}

var mappingSchedulerExecutionSummaryLifecycleStateEnumLowerCase = map[string]SchedulerExecutionSummaryLifecycleStateEnum{
	"active":   SchedulerExecutionSummaryLifecycleStateActive,
	"deleted":  SchedulerExecutionSummaryLifecycleStateDeleted,
	"failed":   SchedulerExecutionSummaryLifecycleStateFailed,
	"inactive": SchedulerExecutionSummaryLifecycleStateInactive,
}

// GetSchedulerExecutionSummaryLifecycleStateEnumValues Enumerates the set of values for SchedulerExecutionSummaryLifecycleStateEnum
func GetSchedulerExecutionSummaryLifecycleStateEnumValues() []SchedulerExecutionSummaryLifecycleStateEnum {
	values := make([]SchedulerExecutionSummaryLifecycleStateEnum, 0)
	for _, v := range mappingSchedulerExecutionSummaryLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetSchedulerExecutionSummaryLifecycleStateEnumStringValues Enumerates the set of values in String for SchedulerExecutionSummaryLifecycleStateEnum
func GetSchedulerExecutionSummaryLifecycleStateEnumStringValues() []string {
	return []string{
		"ACTIVE",
		"DELETED",
		"FAILED",
		"INACTIVE",
	}
}

// GetMappingSchedulerExecutionSummaryLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingSchedulerExecutionSummaryLifecycleStateEnum(val string) (SchedulerExecutionSummaryLifecycleStateEnum, bool) {
	enum, ok := mappingSchedulerExecutionSummaryLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
