// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// OS Management Hub API
//
// Use the OS Management Hub API to manage and monitor updates and patches for instances in OCI, your private data center, or 3rd-party clouds.
// For more information, see Overview of OS Management Hub (https://docs.oracle.com/iaas/osmh/doc/overview.htm).
//

package osmanagementhub

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// ScheduledJob The object that defines a scheduled job. For more information about jobs, see Managing Jobs (https://docs.oracle.com/iaas/osmh/doc/jobs.htm).
type ScheduledJob struct {

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the scheduled job.
	Id *string `mandatory:"true" json:"id"`

	// User-friendly name for the scheduled job.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment that contains the scheduled job.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The type of scheduling frequency for the job.
	ScheduleType ScheduleTypesEnum `mandatory:"true" json:"scheduleType"`

	// The time of the next execution of this scheduled job (in RFC 3339 (https://tools.ietf.org/rfc/rfc3339) format).
	TimeNextExecution *common.SDKTime `mandatory:"true" json:"timeNextExecution"`

	// The list of operations this scheduled job needs to perform.
	// A scheduled job supports only one operation type, unless it is one of the following:
	// * UPDATE_PACKAGES
	// * UPDATE_ALL
	// * UPDATE_SECURITY
	// * UPDATE_BUGFIX
	// * UPDATE_ENHANCEMENT
	// * UPDATE_OTHER
	// * UPDATE_KSPLICE_USERSPACE
	// * UPDATE_KSPLICE_KERNEL
	Operations []ScheduledJobOperation `mandatory:"true" json:"operations"`

	// The time this scheduled job was created (in RFC 3339 (https://tools.ietf.org/rfc/rfc3339) format).
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The time this scheduled job was updated (in RFC 3339 (https://tools.ietf.org/rfc/rfc3339) format).
	TimeUpdated *common.SDKTime `mandatory:"true" json:"timeUpdated"`

	// The current state of the scheduled job.
	LifecycleState ScheduledJobLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace.
	// For more information, see Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"true" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// For more information, see Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"true" json:"definedTags"`

	// User-specified description for the scheduled job.
	Description *string `mandatory:"false" json:"description"`

	// The list of locations this scheduled job should operate on for a job targeting on compartments. (Empty list means apply to all locations). This can only be set when managedCompartmentIds is not empty.
	Locations []ManagedInstanceLocationEnum `mandatory:"false" json:"locations"`

	// The time of the last execution of this scheduled job (in RFC 3339 (https://tools.ietf.org/rfc/rfc3339) format).
	TimeLastExecution *common.SDKTime `mandatory:"false" json:"timeLastExecution"`

	// The frequency schedule for a recurring scheduled job.
	RecurringRule *string `mandatory:"false" json:"recurringRule"`

	// The managed instance OCIDs (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) that this scheduled job operates on.
	// A scheduled job can only operate on one type of target, therefore this parameter is mutually exclusive with
	// managedInstanceGroupIds, managedCompartmentIds, and lifecycleStageIds.
	ManagedInstanceIds []string `mandatory:"false" json:"managedInstanceIds"`

	// The managed instance group OCIDs (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) that this scheduled job operates on. A scheduled job can only operate on one type of target, therefore this parameter is mutually exclusive with managedInstanceIds, managedCompartmentIds, and lifecycleStageIds.
	ManagedInstanceGroupIds []string `mandatory:"false" json:"managedInstanceGroupIds"`

	// The compartment OCIDs (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) that this scheduled job operates on. A scheduled job can only operate on one type of target, therefore this parameter is mutually exclusive with managedInstanceIds, managedInstanceGroupIds, and lifecycleStageIds.
	ManagedCompartmentIds []string `mandatory:"false" json:"managedCompartmentIds"`

	// The lifecycle stage OCIDs (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) that this scheduled job operates on.
	// A scheduled job can only operate on one type of target, therefore this parameter is mutually exclusive with
	// managedInstanceIds, managedInstanceGroupIds, and managedCompartmentIds.
	LifecycleStageIds []string `mandatory:"false" json:"lifecycleStageIds"`

	// Indicates whether to apply the scheduled job to all compartments in the tenancy when managedCompartmentIds specifies the tenancy OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) (root compartment).
	IsSubcompartmentIncluded *bool `mandatory:"false" json:"isSubcompartmentIncluded"`

	// The list of work request OCIDs (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) associated with this scheduled job.
	WorkRequestIds []string `mandatory:"false" json:"workRequestIds"`

	// Indicates whether this scheduled job is managed by the Autonomous Linux service.
	IsManagedByAutonomousLinux *bool `mandatory:"false" json:"isManagedByAutonomousLinux"`

	// System tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"orcl-cloud": {"free-tier-retained": "true"}}`
	SystemTags map[string]map[string]interface{} `mandatory:"false" json:"systemTags"`

	// Indicates if the schedule job has restricted update and deletion capabilities. For restricted scheduled jobs,
	// you can update only the timeNextExecution, recurringRule, and tags.
	IsRestricted *bool `mandatory:"false" json:"isRestricted"`

	// The amount of time in minutes to wait until retrying the scheduled job. If set, the service will automatically retry
	// a failed scheduled job after the interval. For example, you could set the interval to [2,5,10]. If the initial
	// execution of the job fails, the service waits 2 minutes and then retries. If that fails, the service waits 5 minutes
	// and then retries. If that fails, the service waits 10 minutes and then retries.
	RetryIntervals []int `mandatory:"false" json:"retryIntervals"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) for the work request that will be rerun.
	WorkRequestId *string `mandatory:"false" json:"workRequestId"`
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
