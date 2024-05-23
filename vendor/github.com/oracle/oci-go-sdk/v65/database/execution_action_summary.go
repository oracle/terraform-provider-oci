// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Service API
//
// The API for the Database Service. Use this API to manage resources such as databases and DB Systems. For more information, see Overview of the Database Service (https://docs.cloud.oracle.com/iaas/Content/Database/Concepts/databaseoverview.htm).
//

package database

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// ExecutionActionSummary Details of an execution action.
type ExecutionActionSummary struct {

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the execution action.
	Id *string `mandatory:"true" json:"id"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the execution window resource the execution action belongs to.
	ExecutionWindowId *string `mandatory:"true" json:"executionWindowId"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the compartment.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The user-friendly name for the execution action. The name does not need to be unique.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The current state of the execution action. Valid states are SCHEDULED, IN_PROGRESS, FAILED, CANCELED,
	// UPDATING, DELETED, SUCCEEDED and PARTIAL_SUCCESS.
	LifecycleState ExecutionActionSummaryLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// The action type of the execution action being performed
	ActionType ExecutionActionSummaryActionTypeEnum `mandatory:"true" json:"actionType"`

	// Map<ParamName, ParamValue> where a key value pair describes the specific action parameter.
	// Example: `{"count": "3"}`
	ActionParams map[string]string `mandatory:"true" json:"actionParams"`

	// Description of the execution action.
	Description *string `mandatory:"false" json:"description"`

	// The current sub-state of the execution action. Valid states are DURATION_EXCEEDED, MAINTENANCE_IN_PROGRESS and WAITING.
	LifecycleSubstate ExecutionActionSummaryLifecycleSubstateEnum `mandatory:"false" json:"lifecycleSubstate,omitempty"`

	// Additional information about the current lifecycle state.
	LifecycleDetails *string `mandatory:"false" json:"lifecycleDetails"`

	// The date and time the execution action was created.
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`

	// The last date and time that the execution action was updated.
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`

	// The estimated time of the execution action in minutes.
	EstimatedTimeInMins *int `mandatory:"false" json:"estimatedTimeInMins"`

	// The total time taken by corresponding resource activity in minutes.
	TotalTimeTakenInMins *int `mandatory:"false" json:"totalTimeTakenInMins"`

	// The priority order of the execution action.
	ExecutionActionOrder *int `mandatory:"false" json:"executionActionOrder"`

	// List of action members of this execution action.
	ActionMembers []ExecutionActionMember `mandatory:"false" json:"actionMembers"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m ExecutionActionSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ExecutionActionSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingExecutionActionSummaryLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetExecutionActionSummaryLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingExecutionActionSummaryActionTypeEnum(string(m.ActionType)); !ok && m.ActionType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ActionType: %s. Supported values are: %s.", m.ActionType, strings.Join(GetExecutionActionSummaryActionTypeEnumStringValues(), ",")))
	}

	if _, ok := GetMappingExecutionActionSummaryLifecycleSubstateEnum(string(m.LifecycleSubstate)); !ok && m.LifecycleSubstate != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleSubstate: %s. Supported values are: %s.", m.LifecycleSubstate, strings.Join(GetExecutionActionSummaryLifecycleSubstateEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ExecutionActionSummaryLifecycleStateEnum Enum with underlying type: string
type ExecutionActionSummaryLifecycleStateEnum string

// Set of constants representing the allowable values for ExecutionActionSummaryLifecycleStateEnum
const (
	ExecutionActionSummaryLifecycleStateScheduled      ExecutionActionSummaryLifecycleStateEnum = "SCHEDULED"
	ExecutionActionSummaryLifecycleStateInProgress     ExecutionActionSummaryLifecycleStateEnum = "IN_PROGRESS"
	ExecutionActionSummaryLifecycleStateFailed         ExecutionActionSummaryLifecycleStateEnum = "FAILED"
	ExecutionActionSummaryLifecycleStateCanceled       ExecutionActionSummaryLifecycleStateEnum = "CANCELED"
	ExecutionActionSummaryLifecycleStateUpdating       ExecutionActionSummaryLifecycleStateEnum = "UPDATING"
	ExecutionActionSummaryLifecycleStateDeleted        ExecutionActionSummaryLifecycleStateEnum = "DELETED"
	ExecutionActionSummaryLifecycleStateSucceeded      ExecutionActionSummaryLifecycleStateEnum = "SUCCEEDED"
	ExecutionActionSummaryLifecycleStatePartialSuccess ExecutionActionSummaryLifecycleStateEnum = "PARTIAL_SUCCESS"
)

var mappingExecutionActionSummaryLifecycleStateEnum = map[string]ExecutionActionSummaryLifecycleStateEnum{
	"SCHEDULED":       ExecutionActionSummaryLifecycleStateScheduled,
	"IN_PROGRESS":     ExecutionActionSummaryLifecycleStateInProgress,
	"FAILED":          ExecutionActionSummaryLifecycleStateFailed,
	"CANCELED":        ExecutionActionSummaryLifecycleStateCanceled,
	"UPDATING":        ExecutionActionSummaryLifecycleStateUpdating,
	"DELETED":         ExecutionActionSummaryLifecycleStateDeleted,
	"SUCCEEDED":       ExecutionActionSummaryLifecycleStateSucceeded,
	"PARTIAL_SUCCESS": ExecutionActionSummaryLifecycleStatePartialSuccess,
}

var mappingExecutionActionSummaryLifecycleStateEnumLowerCase = map[string]ExecutionActionSummaryLifecycleStateEnum{
	"scheduled":       ExecutionActionSummaryLifecycleStateScheduled,
	"in_progress":     ExecutionActionSummaryLifecycleStateInProgress,
	"failed":          ExecutionActionSummaryLifecycleStateFailed,
	"canceled":        ExecutionActionSummaryLifecycleStateCanceled,
	"updating":        ExecutionActionSummaryLifecycleStateUpdating,
	"deleted":         ExecutionActionSummaryLifecycleStateDeleted,
	"succeeded":       ExecutionActionSummaryLifecycleStateSucceeded,
	"partial_success": ExecutionActionSummaryLifecycleStatePartialSuccess,
}

// GetExecutionActionSummaryLifecycleStateEnumValues Enumerates the set of values for ExecutionActionSummaryLifecycleStateEnum
func GetExecutionActionSummaryLifecycleStateEnumValues() []ExecutionActionSummaryLifecycleStateEnum {
	values := make([]ExecutionActionSummaryLifecycleStateEnum, 0)
	for _, v := range mappingExecutionActionSummaryLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetExecutionActionSummaryLifecycleStateEnumStringValues Enumerates the set of values in String for ExecutionActionSummaryLifecycleStateEnum
func GetExecutionActionSummaryLifecycleStateEnumStringValues() []string {
	return []string{
		"SCHEDULED",
		"IN_PROGRESS",
		"FAILED",
		"CANCELED",
		"UPDATING",
		"DELETED",
		"SUCCEEDED",
		"PARTIAL_SUCCESS",
	}
}

// GetMappingExecutionActionSummaryLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingExecutionActionSummaryLifecycleStateEnum(val string) (ExecutionActionSummaryLifecycleStateEnum, bool) {
	enum, ok := mappingExecutionActionSummaryLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ExecutionActionSummaryLifecycleSubstateEnum Enum with underlying type: string
type ExecutionActionSummaryLifecycleSubstateEnum string

// Set of constants representing the allowable values for ExecutionActionSummaryLifecycleSubstateEnum
const (
	ExecutionActionSummaryLifecycleSubstateDurationExceeded      ExecutionActionSummaryLifecycleSubstateEnum = "DURATION_EXCEEDED"
	ExecutionActionSummaryLifecycleSubstateMaintenanceInProgress ExecutionActionSummaryLifecycleSubstateEnum = "MAINTENANCE_IN_PROGRESS"
	ExecutionActionSummaryLifecycleSubstateWaiting               ExecutionActionSummaryLifecycleSubstateEnum = "WAITING"
	ExecutionActionSummaryLifecycleSubstateRescheduled           ExecutionActionSummaryLifecycleSubstateEnum = "RESCHEDULED"
)

var mappingExecutionActionSummaryLifecycleSubstateEnum = map[string]ExecutionActionSummaryLifecycleSubstateEnum{
	"DURATION_EXCEEDED":       ExecutionActionSummaryLifecycleSubstateDurationExceeded,
	"MAINTENANCE_IN_PROGRESS": ExecutionActionSummaryLifecycleSubstateMaintenanceInProgress,
	"WAITING":                 ExecutionActionSummaryLifecycleSubstateWaiting,
	"RESCHEDULED":             ExecutionActionSummaryLifecycleSubstateRescheduled,
}

var mappingExecutionActionSummaryLifecycleSubstateEnumLowerCase = map[string]ExecutionActionSummaryLifecycleSubstateEnum{
	"duration_exceeded":       ExecutionActionSummaryLifecycleSubstateDurationExceeded,
	"maintenance_in_progress": ExecutionActionSummaryLifecycleSubstateMaintenanceInProgress,
	"waiting":                 ExecutionActionSummaryLifecycleSubstateWaiting,
	"rescheduled":             ExecutionActionSummaryLifecycleSubstateRescheduled,
}

// GetExecutionActionSummaryLifecycleSubstateEnumValues Enumerates the set of values for ExecutionActionSummaryLifecycleSubstateEnum
func GetExecutionActionSummaryLifecycleSubstateEnumValues() []ExecutionActionSummaryLifecycleSubstateEnum {
	values := make([]ExecutionActionSummaryLifecycleSubstateEnum, 0)
	for _, v := range mappingExecutionActionSummaryLifecycleSubstateEnum {
		values = append(values, v)
	}
	return values
}

// GetExecutionActionSummaryLifecycleSubstateEnumStringValues Enumerates the set of values in String for ExecutionActionSummaryLifecycleSubstateEnum
func GetExecutionActionSummaryLifecycleSubstateEnumStringValues() []string {
	return []string{
		"DURATION_EXCEEDED",
		"MAINTENANCE_IN_PROGRESS",
		"WAITING",
		"RESCHEDULED",
	}
}

// GetMappingExecutionActionSummaryLifecycleSubstateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingExecutionActionSummaryLifecycleSubstateEnum(val string) (ExecutionActionSummaryLifecycleSubstateEnum, bool) {
	enum, ok := mappingExecutionActionSummaryLifecycleSubstateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ExecutionActionSummaryActionTypeEnum Enum with underlying type: string
type ExecutionActionSummaryActionTypeEnum string

// Set of constants representing the allowable values for ExecutionActionSummaryActionTypeEnum
const (
	ExecutionActionSummaryActionTypeDbServerFullSoftwareUpdate      ExecutionActionSummaryActionTypeEnum = "DB_SERVER_FULL_SOFTWARE_UPDATE"
	ExecutionActionSummaryActionTypeStorageServerFullSoftwareUpdate ExecutionActionSummaryActionTypeEnum = "STORAGE_SERVER_FULL_SOFTWARE_UPDATE"
	ExecutionActionSummaryActionTypeNetworkSwitchFullSoftwareUpdate ExecutionActionSummaryActionTypeEnum = "NETWORK_SWITCH_FULL_SOFTWARE_UPDATE"
)

var mappingExecutionActionSummaryActionTypeEnum = map[string]ExecutionActionSummaryActionTypeEnum{
	"DB_SERVER_FULL_SOFTWARE_UPDATE":      ExecutionActionSummaryActionTypeDbServerFullSoftwareUpdate,
	"STORAGE_SERVER_FULL_SOFTWARE_UPDATE": ExecutionActionSummaryActionTypeStorageServerFullSoftwareUpdate,
	"NETWORK_SWITCH_FULL_SOFTWARE_UPDATE": ExecutionActionSummaryActionTypeNetworkSwitchFullSoftwareUpdate,
}

var mappingExecutionActionSummaryActionTypeEnumLowerCase = map[string]ExecutionActionSummaryActionTypeEnum{
	"db_server_full_software_update":      ExecutionActionSummaryActionTypeDbServerFullSoftwareUpdate,
	"storage_server_full_software_update": ExecutionActionSummaryActionTypeStorageServerFullSoftwareUpdate,
	"network_switch_full_software_update": ExecutionActionSummaryActionTypeNetworkSwitchFullSoftwareUpdate,
}

// GetExecutionActionSummaryActionTypeEnumValues Enumerates the set of values for ExecutionActionSummaryActionTypeEnum
func GetExecutionActionSummaryActionTypeEnumValues() []ExecutionActionSummaryActionTypeEnum {
	values := make([]ExecutionActionSummaryActionTypeEnum, 0)
	for _, v := range mappingExecutionActionSummaryActionTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetExecutionActionSummaryActionTypeEnumStringValues Enumerates the set of values in String for ExecutionActionSummaryActionTypeEnum
func GetExecutionActionSummaryActionTypeEnumStringValues() []string {
	return []string{
		"DB_SERVER_FULL_SOFTWARE_UPDATE",
		"STORAGE_SERVER_FULL_SOFTWARE_UPDATE",
		"NETWORK_SWITCH_FULL_SOFTWARE_UPDATE",
	}
}

// GetMappingExecutionActionSummaryActionTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingExecutionActionSummaryActionTypeEnum(val string) (ExecutionActionSummaryActionTypeEnum, bool) {
	enum, ok := mappingExecutionActionSummaryActionTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
