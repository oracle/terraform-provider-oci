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

// ExecutionAction Details of an execution action.
type ExecutionAction struct {

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
	LifecycleState ExecutionActionLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// The action type of the execution action being performed
	ActionType ExecutionActionActionTypeEnum `mandatory:"true" json:"actionType"`

	// Map<ParamName, ParamValue> where a key value pair describes the specific action parameter.
	// Example: `{"count": "3"}`
	ActionParams map[string]string `mandatory:"true" json:"actionParams"`

	// Description of the execution action.
	Description *string `mandatory:"false" json:"description"`

	// The current sub-state of the execution action. Valid states are DURATION_EXCEEDED, MAINTENANCE_IN_PROGRESS and WAITING.
	LifecycleSubstate ExecutionActionLifecycleSubstateEnum `mandatory:"false" json:"lifecycleSubstate,omitempty"`

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

func (m ExecutionAction) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ExecutionAction) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingExecutionActionLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetExecutionActionLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingExecutionActionActionTypeEnum(string(m.ActionType)); !ok && m.ActionType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ActionType: %s. Supported values are: %s.", m.ActionType, strings.Join(GetExecutionActionActionTypeEnumStringValues(), ",")))
	}

	if _, ok := GetMappingExecutionActionLifecycleSubstateEnum(string(m.LifecycleSubstate)); !ok && m.LifecycleSubstate != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleSubstate: %s. Supported values are: %s.", m.LifecycleSubstate, strings.Join(GetExecutionActionLifecycleSubstateEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ExecutionActionLifecycleStateEnum Enum with underlying type: string
type ExecutionActionLifecycleStateEnum string

// Set of constants representing the allowable values for ExecutionActionLifecycleStateEnum
const (
	ExecutionActionLifecycleStateScheduled      ExecutionActionLifecycleStateEnum = "SCHEDULED"
	ExecutionActionLifecycleStateInProgress     ExecutionActionLifecycleStateEnum = "IN_PROGRESS"
	ExecutionActionLifecycleStateFailed         ExecutionActionLifecycleStateEnum = "FAILED"
	ExecutionActionLifecycleStateCanceled       ExecutionActionLifecycleStateEnum = "CANCELED"
	ExecutionActionLifecycleStateUpdating       ExecutionActionLifecycleStateEnum = "UPDATING"
	ExecutionActionLifecycleStateDeleted        ExecutionActionLifecycleStateEnum = "DELETED"
	ExecutionActionLifecycleStateSucceeded      ExecutionActionLifecycleStateEnum = "SUCCEEDED"
	ExecutionActionLifecycleStatePartialSuccess ExecutionActionLifecycleStateEnum = "PARTIAL_SUCCESS"
)

var mappingExecutionActionLifecycleStateEnum = map[string]ExecutionActionLifecycleStateEnum{
	"SCHEDULED":       ExecutionActionLifecycleStateScheduled,
	"IN_PROGRESS":     ExecutionActionLifecycleStateInProgress,
	"FAILED":          ExecutionActionLifecycleStateFailed,
	"CANCELED":        ExecutionActionLifecycleStateCanceled,
	"UPDATING":        ExecutionActionLifecycleStateUpdating,
	"DELETED":         ExecutionActionLifecycleStateDeleted,
	"SUCCEEDED":       ExecutionActionLifecycleStateSucceeded,
	"PARTIAL_SUCCESS": ExecutionActionLifecycleStatePartialSuccess,
}

var mappingExecutionActionLifecycleStateEnumLowerCase = map[string]ExecutionActionLifecycleStateEnum{
	"scheduled":       ExecutionActionLifecycleStateScheduled,
	"in_progress":     ExecutionActionLifecycleStateInProgress,
	"failed":          ExecutionActionLifecycleStateFailed,
	"canceled":        ExecutionActionLifecycleStateCanceled,
	"updating":        ExecutionActionLifecycleStateUpdating,
	"deleted":         ExecutionActionLifecycleStateDeleted,
	"succeeded":       ExecutionActionLifecycleStateSucceeded,
	"partial_success": ExecutionActionLifecycleStatePartialSuccess,
}

// GetExecutionActionLifecycleStateEnumValues Enumerates the set of values for ExecutionActionLifecycleStateEnum
func GetExecutionActionLifecycleStateEnumValues() []ExecutionActionLifecycleStateEnum {
	values := make([]ExecutionActionLifecycleStateEnum, 0)
	for _, v := range mappingExecutionActionLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetExecutionActionLifecycleStateEnumStringValues Enumerates the set of values in String for ExecutionActionLifecycleStateEnum
func GetExecutionActionLifecycleStateEnumStringValues() []string {
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

// GetMappingExecutionActionLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingExecutionActionLifecycleStateEnum(val string) (ExecutionActionLifecycleStateEnum, bool) {
	enum, ok := mappingExecutionActionLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ExecutionActionLifecycleSubstateEnum Enum with underlying type: string
type ExecutionActionLifecycleSubstateEnum string

// Set of constants representing the allowable values for ExecutionActionLifecycleSubstateEnum
const (
	ExecutionActionLifecycleSubstateDurationExceeded      ExecutionActionLifecycleSubstateEnum = "DURATION_EXCEEDED"
	ExecutionActionLifecycleSubstateMaintenanceInProgress ExecutionActionLifecycleSubstateEnum = "MAINTENANCE_IN_PROGRESS"
	ExecutionActionLifecycleSubstateWaiting               ExecutionActionLifecycleSubstateEnum = "WAITING"
	ExecutionActionLifecycleSubstateRescheduled           ExecutionActionLifecycleSubstateEnum = "RESCHEDULED"
)

var mappingExecutionActionLifecycleSubstateEnum = map[string]ExecutionActionLifecycleSubstateEnum{
	"DURATION_EXCEEDED":       ExecutionActionLifecycleSubstateDurationExceeded,
	"MAINTENANCE_IN_PROGRESS": ExecutionActionLifecycleSubstateMaintenanceInProgress,
	"WAITING":                 ExecutionActionLifecycleSubstateWaiting,
	"RESCHEDULED":             ExecutionActionLifecycleSubstateRescheduled,
}

var mappingExecutionActionLifecycleSubstateEnumLowerCase = map[string]ExecutionActionLifecycleSubstateEnum{
	"duration_exceeded":       ExecutionActionLifecycleSubstateDurationExceeded,
	"maintenance_in_progress": ExecutionActionLifecycleSubstateMaintenanceInProgress,
	"waiting":                 ExecutionActionLifecycleSubstateWaiting,
	"rescheduled":             ExecutionActionLifecycleSubstateRescheduled,
}

// GetExecutionActionLifecycleSubstateEnumValues Enumerates the set of values for ExecutionActionLifecycleSubstateEnum
func GetExecutionActionLifecycleSubstateEnumValues() []ExecutionActionLifecycleSubstateEnum {
	values := make([]ExecutionActionLifecycleSubstateEnum, 0)
	for _, v := range mappingExecutionActionLifecycleSubstateEnum {
		values = append(values, v)
	}
	return values
}

// GetExecutionActionLifecycleSubstateEnumStringValues Enumerates the set of values in String for ExecutionActionLifecycleSubstateEnum
func GetExecutionActionLifecycleSubstateEnumStringValues() []string {
	return []string{
		"DURATION_EXCEEDED",
		"MAINTENANCE_IN_PROGRESS",
		"WAITING",
		"RESCHEDULED",
	}
}

// GetMappingExecutionActionLifecycleSubstateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingExecutionActionLifecycleSubstateEnum(val string) (ExecutionActionLifecycleSubstateEnum, bool) {
	enum, ok := mappingExecutionActionLifecycleSubstateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ExecutionActionActionTypeEnum Enum with underlying type: string
type ExecutionActionActionTypeEnum string

// Set of constants representing the allowable values for ExecutionActionActionTypeEnum
const (
	ExecutionActionActionTypeDbServerFullSoftwareUpdate      ExecutionActionActionTypeEnum = "DB_SERVER_FULL_SOFTWARE_UPDATE"
	ExecutionActionActionTypeStorageServerFullSoftwareUpdate ExecutionActionActionTypeEnum = "STORAGE_SERVER_FULL_SOFTWARE_UPDATE"
	ExecutionActionActionTypeNetworkSwitchFullSoftwareUpdate ExecutionActionActionTypeEnum = "NETWORK_SWITCH_FULL_SOFTWARE_UPDATE"
)

var mappingExecutionActionActionTypeEnum = map[string]ExecutionActionActionTypeEnum{
	"DB_SERVER_FULL_SOFTWARE_UPDATE":      ExecutionActionActionTypeDbServerFullSoftwareUpdate,
	"STORAGE_SERVER_FULL_SOFTWARE_UPDATE": ExecutionActionActionTypeStorageServerFullSoftwareUpdate,
	"NETWORK_SWITCH_FULL_SOFTWARE_UPDATE": ExecutionActionActionTypeNetworkSwitchFullSoftwareUpdate,
}

var mappingExecutionActionActionTypeEnumLowerCase = map[string]ExecutionActionActionTypeEnum{
	"db_server_full_software_update":      ExecutionActionActionTypeDbServerFullSoftwareUpdate,
	"storage_server_full_software_update": ExecutionActionActionTypeStorageServerFullSoftwareUpdate,
	"network_switch_full_software_update": ExecutionActionActionTypeNetworkSwitchFullSoftwareUpdate,
}

// GetExecutionActionActionTypeEnumValues Enumerates the set of values for ExecutionActionActionTypeEnum
func GetExecutionActionActionTypeEnumValues() []ExecutionActionActionTypeEnum {
	values := make([]ExecutionActionActionTypeEnum, 0)
	for _, v := range mappingExecutionActionActionTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetExecutionActionActionTypeEnumStringValues Enumerates the set of values in String for ExecutionActionActionTypeEnum
func GetExecutionActionActionTypeEnumStringValues() []string {
	return []string{
		"DB_SERVER_FULL_SOFTWARE_UPDATE",
		"STORAGE_SERVER_FULL_SOFTWARE_UPDATE",
		"NETWORK_SWITCH_FULL_SOFTWARE_UPDATE",
	}
}

// GetMappingExecutionActionActionTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingExecutionActionActionTypeEnum(val string) (ExecutionActionActionTypeEnum, bool) {
	enum, ok := mappingExecutionActionActionTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
