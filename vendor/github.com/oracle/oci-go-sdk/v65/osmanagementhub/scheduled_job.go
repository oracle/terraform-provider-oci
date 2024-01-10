// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// OS Management Hub API
//
// Use the OS Management Hub API to manage and monitor updates and patches for the operating system environments in your private data centers through a single management console. For more information, see Overview of OS Management Hub (https://docs.cloud.oracle.com/iaas/osmh/doc/overview.htm).
//

package osmanagementhub

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// ScheduledJob Detailed information about a scheduled job.
type ScheduledJob struct {

	// The OCID of the scheduled job.
	Id *string `mandatory:"true" json:"id"`

	// Scheduled job name.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The OCID of the compartment containing the scheduled job.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The type of scheduling this scheduled job follows.
	ScheduleType ScheduleTypesEnum `mandatory:"true" json:"scheduleType"`

	// The time of the next execution of this scheduled job.
	TimeNextExecution *common.SDKTime `mandatory:"true" json:"timeNextExecution"`

	// The list of operations this scheduled job needs to perform (can only support one operation if the operationType is not UPDATE_PACKAGES/UPDATE_ALL/UPDATE_SECURITY/UPDATE_BUGFIX/UPDATE_ENHANCEMENT/UPDATE_OTHER/UPDATE_KSPLICE_USERSPACE/UPDATE_KSPLICE_KERNEL).
	Operations []ScheduledJobOperation `mandatory:"true" json:"operations"`

	// The time this scheduled job was created. An RFC3339 formatted datetime string.
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The time this scheduled job was updated. An RFC3339 formatted datetime string.
	TimeUpdated *common.SDKTime `mandatory:"true" json:"timeUpdated"`

	// The current state of the scheduled job.
	LifecycleState ScheduledJobLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"true" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"true" json:"definedTags"`

	// Details describing the scheduled job.
	Description *string `mandatory:"false" json:"description"`

	// The time of the last execution of this scheduled job.
	TimeLastExecution *common.SDKTime `mandatory:"false" json:"timeLastExecution"`

	// The recurring rule for a RECURRING scheduled job.
	RecurringRule *string `mandatory:"false" json:"recurringRule"`

	// The list of managed instance OCIDs this scheduled job operates on (mutually exclusive with managedInstanceGroupIds, managedCompartmentIds and lifecycleStageIds).
	ManagedInstanceIds []string `mandatory:"false" json:"managedInstanceIds"`

	// The list of managed instance group OCIDs this scheduled job operates on (mutually exclusive with managedInstances, managedCompartmentIds and lifecycleStageIds).
	ManagedInstanceGroupIds []string `mandatory:"false" json:"managedInstanceGroupIds"`

	// The list of target compartment OCIDs if this scheduled job operates on a compartment level (mutually exclusive with managedInstances, managedInstanceGroupIds and lifecycleStageIds).
	ManagedCompartmentIds []string `mandatory:"false" json:"managedCompartmentIds"`

	// The list of target lifecycle stage OCIDs if this scheduled job operates on lifecycle stages (mutually exclusive with managedInstances, managedInstanceGroupIds and managedCompartmentIds).
	LifecycleStageIds []string `mandatory:"false" json:"lifecycleStageIds"`

	// Whether to create jobs for all compartments in the tenancy when managedCompartmentIds specifies the tenancy OCID.
	IsSubcompartmentIncluded *bool `mandatory:"false" json:"isSubcompartmentIncluded"`

	// The list of work request OCIDs associated with this scheduled job.
	WorkRequestIds []string `mandatory:"false" json:"workRequestIds"`

	// System tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"orcl-cloud": {"free-tier-retained": "true"}}`
	SystemTags map[string]map[string]interface{} `mandatory:"false" json:"systemTags"`

	// true, if the schedule job has its update/deletion capabilities restricted. (Used to track scheduled job for management station syncing).
	IsRestricted *bool `mandatory:"false" json:"isRestricted"`
}

func (m ScheduledJob) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ScheduledJob) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingScheduleTypesEnum(string(m.ScheduleType)); !ok && m.ScheduleType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ScheduleType: %s. Supported values are: %s.", m.ScheduleType, strings.Join(GetScheduleTypesEnumStringValues(), ",")))
	}
	if _, ok := GetMappingScheduledJobLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetScheduledJobLifecycleStateEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ScheduledJobLifecycleStateEnum Enum with underlying type: string
type ScheduledJobLifecycleStateEnum string

// Set of constants representing the allowable values for ScheduledJobLifecycleStateEnum
const (
	ScheduledJobLifecycleStateCreating ScheduledJobLifecycleStateEnum = "CREATING"
	ScheduledJobLifecycleStateUpdating ScheduledJobLifecycleStateEnum = "UPDATING"
	ScheduledJobLifecycleStateActive   ScheduledJobLifecycleStateEnum = "ACTIVE"
	ScheduledJobLifecycleStateInactive ScheduledJobLifecycleStateEnum = "INACTIVE"
	ScheduledJobLifecycleStateDeleting ScheduledJobLifecycleStateEnum = "DELETING"
	ScheduledJobLifecycleStateDeleted  ScheduledJobLifecycleStateEnum = "DELETED"
	ScheduledJobLifecycleStateFailed   ScheduledJobLifecycleStateEnum = "FAILED"
)

var mappingScheduledJobLifecycleStateEnum = map[string]ScheduledJobLifecycleStateEnum{
	"CREATING": ScheduledJobLifecycleStateCreating,
	"UPDATING": ScheduledJobLifecycleStateUpdating,
	"ACTIVE":   ScheduledJobLifecycleStateActive,
	"INACTIVE": ScheduledJobLifecycleStateInactive,
	"DELETING": ScheduledJobLifecycleStateDeleting,
	"DELETED":  ScheduledJobLifecycleStateDeleted,
	"FAILED":   ScheduledJobLifecycleStateFailed,
}

var mappingScheduledJobLifecycleStateEnumLowerCase = map[string]ScheduledJobLifecycleStateEnum{
	"creating": ScheduledJobLifecycleStateCreating,
	"updating": ScheduledJobLifecycleStateUpdating,
	"active":   ScheduledJobLifecycleStateActive,
	"inactive": ScheduledJobLifecycleStateInactive,
	"deleting": ScheduledJobLifecycleStateDeleting,
	"deleted":  ScheduledJobLifecycleStateDeleted,
	"failed":   ScheduledJobLifecycleStateFailed,
}

// GetScheduledJobLifecycleStateEnumValues Enumerates the set of values for ScheduledJobLifecycleStateEnum
func GetScheduledJobLifecycleStateEnumValues() []ScheduledJobLifecycleStateEnum {
	values := make([]ScheduledJobLifecycleStateEnum, 0)
	for _, v := range mappingScheduledJobLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetScheduledJobLifecycleStateEnumStringValues Enumerates the set of values in String for ScheduledJobLifecycleStateEnum
func GetScheduledJobLifecycleStateEnumStringValues() []string {
	return []string{
		"CREATING",
		"UPDATING",
		"ACTIVE",
		"INACTIVE",
		"DELETING",
		"DELETED",
		"FAILED",
	}
}

// GetMappingScheduledJobLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingScheduledJobLifecycleStateEnum(val string) (ScheduledJobLifecycleStateEnum, bool) {
	enum, ok := mappingScheduledJobLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
