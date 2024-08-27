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

// ScheduledAction Details of a Scheduled Action.
type ScheduledAction struct {

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the Scheduled Action.
	Id *string `mandatory:"true" json:"id"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the Scheduling Plan.
	SchedulingPlanId *string `mandatory:"true" json:"schedulingPlanId"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the compartment.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The display name of the Scheduled Action.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The order of the scheduled action.
	ActionOrder *int `mandatory:"true" json:"actionOrder"`

	// The type of the scheduled action being performed
	ActionType ScheduledActionActionTypeEnum `mandatory:"true" json:"actionType"`

	// The current state of the Scheduled Action. Valid states are CREATING, NEEDS_ATTENTION, AVAILABLE, UPDATING, FAILED, DELETING and DELETED.
	LifecycleState ScheduledActionLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// The date and time the Scheduled Action Resource was created.
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the Scheduling Window.
	SchedulingWindowId *string `mandatory:"false" json:"schedulingWindowId"`

	// The estimated patching time for the scheduled action.
	EstimatedTimeInMins *int `mandatory:"false" json:"estimatedTimeInMins"`

	// Map<ParamName, ParamValue> where a key value pair describes the specific action parameter.
	// Example: `{"count": "3"}`
	ActionParams map[string]string `mandatory:"false" json:"actionParams"`

	// The list of action members in a scheduled action.
	ActionMembers []ActionMember `mandatory:"false" json:"actionMembers"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// System tags for this resource. Each key is predefined and scoped to a namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	SystemTags map[string]map[string]interface{} `mandatory:"false" json:"systemTags"`

	// The date and time the Scheduled Action Resource was updated.
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`
}

func (m ScheduledAction) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ScheduledAction) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingScheduledActionActionTypeEnum(string(m.ActionType)); !ok && m.ActionType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ActionType: %s. Supported values are: %s.", m.ActionType, strings.Join(GetScheduledActionActionTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingScheduledActionLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetScheduledActionLifecycleStateEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ScheduledActionActionTypeEnum Enum with underlying type: string
type ScheduledActionActionTypeEnum string

// Set of constants representing the allowable values for ScheduledActionActionTypeEnum
const (
	ScheduledActionActionTypeDbServerFullSoftwareUpdate      ScheduledActionActionTypeEnum = "DB_SERVER_FULL_SOFTWARE_UPDATE"
	ScheduledActionActionTypeStorageServerFullSoftwareUpdate ScheduledActionActionTypeEnum = "STORAGE_SERVER_FULL_SOFTWARE_UPDATE"
	ScheduledActionActionTypeNetworkSwitchFullSoftwareUpdate ScheduledActionActionTypeEnum = "NETWORK_SWITCH_FULL_SOFTWARE_UPDATE"
)

var mappingScheduledActionActionTypeEnum = map[string]ScheduledActionActionTypeEnum{
	"DB_SERVER_FULL_SOFTWARE_UPDATE":      ScheduledActionActionTypeDbServerFullSoftwareUpdate,
	"STORAGE_SERVER_FULL_SOFTWARE_UPDATE": ScheduledActionActionTypeStorageServerFullSoftwareUpdate,
	"NETWORK_SWITCH_FULL_SOFTWARE_UPDATE": ScheduledActionActionTypeNetworkSwitchFullSoftwareUpdate,
}

var mappingScheduledActionActionTypeEnumLowerCase = map[string]ScheduledActionActionTypeEnum{
	"db_server_full_software_update":      ScheduledActionActionTypeDbServerFullSoftwareUpdate,
	"storage_server_full_software_update": ScheduledActionActionTypeStorageServerFullSoftwareUpdate,
	"network_switch_full_software_update": ScheduledActionActionTypeNetworkSwitchFullSoftwareUpdate,
}

// GetScheduledActionActionTypeEnumValues Enumerates the set of values for ScheduledActionActionTypeEnum
func GetScheduledActionActionTypeEnumValues() []ScheduledActionActionTypeEnum {
	values := make([]ScheduledActionActionTypeEnum, 0)
	for _, v := range mappingScheduledActionActionTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetScheduledActionActionTypeEnumStringValues Enumerates the set of values in String for ScheduledActionActionTypeEnum
func GetScheduledActionActionTypeEnumStringValues() []string {
	return []string{
		"DB_SERVER_FULL_SOFTWARE_UPDATE",
		"STORAGE_SERVER_FULL_SOFTWARE_UPDATE",
		"NETWORK_SWITCH_FULL_SOFTWARE_UPDATE",
	}
}

// GetMappingScheduledActionActionTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingScheduledActionActionTypeEnum(val string) (ScheduledActionActionTypeEnum, bool) {
	enum, ok := mappingScheduledActionActionTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ScheduledActionLifecycleStateEnum Enum with underlying type: string
type ScheduledActionLifecycleStateEnum string

// Set of constants representing the allowable values for ScheduledActionLifecycleStateEnum
const (
	ScheduledActionLifecycleStateCreating       ScheduledActionLifecycleStateEnum = "CREATING"
	ScheduledActionLifecycleStateNeedsAttention ScheduledActionLifecycleStateEnum = "NEEDS_ATTENTION"
	ScheduledActionLifecycleStateAvailable      ScheduledActionLifecycleStateEnum = "AVAILABLE"
	ScheduledActionLifecycleStateUpdating       ScheduledActionLifecycleStateEnum = "UPDATING"
	ScheduledActionLifecycleStateFailed         ScheduledActionLifecycleStateEnum = "FAILED"
	ScheduledActionLifecycleStateDeleting       ScheduledActionLifecycleStateEnum = "DELETING"
	ScheduledActionLifecycleStateDeleted        ScheduledActionLifecycleStateEnum = "DELETED"
)

var mappingScheduledActionLifecycleStateEnum = map[string]ScheduledActionLifecycleStateEnum{
	"CREATING":        ScheduledActionLifecycleStateCreating,
	"NEEDS_ATTENTION": ScheduledActionLifecycleStateNeedsAttention,
	"AVAILABLE":       ScheduledActionLifecycleStateAvailable,
	"UPDATING":        ScheduledActionLifecycleStateUpdating,
	"FAILED":          ScheduledActionLifecycleStateFailed,
	"DELETING":        ScheduledActionLifecycleStateDeleting,
	"DELETED":         ScheduledActionLifecycleStateDeleted,
}

var mappingScheduledActionLifecycleStateEnumLowerCase = map[string]ScheduledActionLifecycleStateEnum{
	"creating":        ScheduledActionLifecycleStateCreating,
	"needs_attention": ScheduledActionLifecycleStateNeedsAttention,
	"available":       ScheduledActionLifecycleStateAvailable,
	"updating":        ScheduledActionLifecycleStateUpdating,
	"failed":          ScheduledActionLifecycleStateFailed,
	"deleting":        ScheduledActionLifecycleStateDeleting,
	"deleted":         ScheduledActionLifecycleStateDeleted,
}

// GetScheduledActionLifecycleStateEnumValues Enumerates the set of values for ScheduledActionLifecycleStateEnum
func GetScheduledActionLifecycleStateEnumValues() []ScheduledActionLifecycleStateEnum {
	values := make([]ScheduledActionLifecycleStateEnum, 0)
	for _, v := range mappingScheduledActionLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetScheduledActionLifecycleStateEnumStringValues Enumerates the set of values in String for ScheduledActionLifecycleStateEnum
func GetScheduledActionLifecycleStateEnumStringValues() []string {
	return []string{
		"CREATING",
		"NEEDS_ATTENTION",
		"AVAILABLE",
		"UPDATING",
		"FAILED",
		"DELETING",
		"DELETED",
	}
}

// GetMappingScheduledActionLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingScheduledActionLifecycleStateEnum(val string) (ScheduledActionLifecycleStateEnum, bool) {
	enum, ok := mappingScheduledActionLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
