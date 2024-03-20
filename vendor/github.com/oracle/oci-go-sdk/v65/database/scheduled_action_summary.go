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

// ScheduledActionSummary Details of a scheduled action.
type ScheduledActionSummary struct {

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
	ActionType ScheduledActionSummaryActionTypeEnum `mandatory:"true" json:"actionType"`

	// The current state of the Scheduled Action. Valid states are CREATING, NEEDS_ATTENTION, AVAILABLE, UPDATING, FAILED, DELETING and DELETED.
	LifecycleState ScheduledActionSummaryLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

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

func (m ScheduledActionSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ScheduledActionSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingScheduledActionSummaryActionTypeEnum(string(m.ActionType)); !ok && m.ActionType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ActionType: %s. Supported values are: %s.", m.ActionType, strings.Join(GetScheduledActionSummaryActionTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingScheduledActionSummaryLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetScheduledActionSummaryLifecycleStateEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ScheduledActionSummaryActionTypeEnum Enum with underlying type: string
type ScheduledActionSummaryActionTypeEnum string

// Set of constants representing the allowable values for ScheduledActionSummaryActionTypeEnum
const (
	ScheduledActionSummaryActionTypeDbServerFullSoftwareUpdate      ScheduledActionSummaryActionTypeEnum = "DB_SERVER_FULL_SOFTWARE_UPDATE"
	ScheduledActionSummaryActionTypeStorageServerFullSoftwareUpdate ScheduledActionSummaryActionTypeEnum = "STORAGE_SERVER_FULL_SOFTWARE_UPDATE"
	ScheduledActionSummaryActionTypeNetworkSwitchFullSoftwareUpdate ScheduledActionSummaryActionTypeEnum = "NETWORK_SWITCH_FULL_SOFTWARE_UPDATE"
)

var mappingScheduledActionSummaryActionTypeEnum = map[string]ScheduledActionSummaryActionTypeEnum{
	"DB_SERVER_FULL_SOFTWARE_UPDATE":      ScheduledActionSummaryActionTypeDbServerFullSoftwareUpdate,
	"STORAGE_SERVER_FULL_SOFTWARE_UPDATE": ScheduledActionSummaryActionTypeStorageServerFullSoftwareUpdate,
	"NETWORK_SWITCH_FULL_SOFTWARE_UPDATE": ScheduledActionSummaryActionTypeNetworkSwitchFullSoftwareUpdate,
}

var mappingScheduledActionSummaryActionTypeEnumLowerCase = map[string]ScheduledActionSummaryActionTypeEnum{
	"db_server_full_software_update":      ScheduledActionSummaryActionTypeDbServerFullSoftwareUpdate,
	"storage_server_full_software_update": ScheduledActionSummaryActionTypeStorageServerFullSoftwareUpdate,
	"network_switch_full_software_update": ScheduledActionSummaryActionTypeNetworkSwitchFullSoftwareUpdate,
}

// GetScheduledActionSummaryActionTypeEnumValues Enumerates the set of values for ScheduledActionSummaryActionTypeEnum
func GetScheduledActionSummaryActionTypeEnumValues() []ScheduledActionSummaryActionTypeEnum {
	values := make([]ScheduledActionSummaryActionTypeEnum, 0)
	for _, v := range mappingScheduledActionSummaryActionTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetScheduledActionSummaryActionTypeEnumStringValues Enumerates the set of values in String for ScheduledActionSummaryActionTypeEnum
func GetScheduledActionSummaryActionTypeEnumStringValues() []string {
	return []string{
		"DB_SERVER_FULL_SOFTWARE_UPDATE",
		"STORAGE_SERVER_FULL_SOFTWARE_UPDATE",
		"NETWORK_SWITCH_FULL_SOFTWARE_UPDATE",
	}
}

// GetMappingScheduledActionSummaryActionTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingScheduledActionSummaryActionTypeEnum(val string) (ScheduledActionSummaryActionTypeEnum, bool) {
	enum, ok := mappingScheduledActionSummaryActionTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ScheduledActionSummaryLifecycleStateEnum Enum with underlying type: string
type ScheduledActionSummaryLifecycleStateEnum string

// Set of constants representing the allowable values for ScheduledActionSummaryLifecycleStateEnum
const (
	ScheduledActionSummaryLifecycleStateCreating       ScheduledActionSummaryLifecycleStateEnum = "CREATING"
	ScheduledActionSummaryLifecycleStateNeedsAttention ScheduledActionSummaryLifecycleStateEnum = "NEEDS_ATTENTION"
	ScheduledActionSummaryLifecycleStateAvailable      ScheduledActionSummaryLifecycleStateEnum = "AVAILABLE"
	ScheduledActionSummaryLifecycleStateUpdating       ScheduledActionSummaryLifecycleStateEnum = "UPDATING"
	ScheduledActionSummaryLifecycleStateFailed         ScheduledActionSummaryLifecycleStateEnum = "FAILED"
	ScheduledActionSummaryLifecycleStateDeleting       ScheduledActionSummaryLifecycleStateEnum = "DELETING"
	ScheduledActionSummaryLifecycleStateDeleted        ScheduledActionSummaryLifecycleStateEnum = "DELETED"
)

var mappingScheduledActionSummaryLifecycleStateEnum = map[string]ScheduledActionSummaryLifecycleStateEnum{
	"CREATING":        ScheduledActionSummaryLifecycleStateCreating,
	"NEEDS_ATTENTION": ScheduledActionSummaryLifecycleStateNeedsAttention,
	"AVAILABLE":       ScheduledActionSummaryLifecycleStateAvailable,
	"UPDATING":        ScheduledActionSummaryLifecycleStateUpdating,
	"FAILED":          ScheduledActionSummaryLifecycleStateFailed,
	"DELETING":        ScheduledActionSummaryLifecycleStateDeleting,
	"DELETED":         ScheduledActionSummaryLifecycleStateDeleted,
}

var mappingScheduledActionSummaryLifecycleStateEnumLowerCase = map[string]ScheduledActionSummaryLifecycleStateEnum{
	"creating":        ScheduledActionSummaryLifecycleStateCreating,
	"needs_attention": ScheduledActionSummaryLifecycleStateNeedsAttention,
	"available":       ScheduledActionSummaryLifecycleStateAvailable,
	"updating":        ScheduledActionSummaryLifecycleStateUpdating,
	"failed":          ScheduledActionSummaryLifecycleStateFailed,
	"deleting":        ScheduledActionSummaryLifecycleStateDeleting,
	"deleted":         ScheduledActionSummaryLifecycleStateDeleted,
}

// GetScheduledActionSummaryLifecycleStateEnumValues Enumerates the set of values for ScheduledActionSummaryLifecycleStateEnum
func GetScheduledActionSummaryLifecycleStateEnumValues() []ScheduledActionSummaryLifecycleStateEnum {
	values := make([]ScheduledActionSummaryLifecycleStateEnum, 0)
	for _, v := range mappingScheduledActionSummaryLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetScheduledActionSummaryLifecycleStateEnumStringValues Enumerates the set of values in String for ScheduledActionSummaryLifecycleStateEnum
func GetScheduledActionSummaryLifecycleStateEnumStringValues() []string {
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

// GetMappingScheduledActionSummaryLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingScheduledActionSummaryLifecycleStateEnum(val string) (ScheduledActionSummaryLifecycleStateEnum, bool) {
	enum, ok := mappingScheduledActionSummaryLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
