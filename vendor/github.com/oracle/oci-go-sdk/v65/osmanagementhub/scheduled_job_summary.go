// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
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

// ScheduledJobSummary Summary of the scheduled job.
type ScheduledJobSummary struct {

	// The OCID of the scheduled job.
	Id *string `mandatory:"true" json:"id"`

	// Scheduled job name.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The OCID of the compartment containing the scheduled job.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The type of scheduling this scheduled job follows.
	ScheduleType ScheduleTypesEnum `mandatory:"true" json:"scheduleType"`

	// The time this scheduled job was created. An RFC3339 formatted datetime string.
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The time this scheduled job was updated. An RFC3339 formatted datetime string.
	TimeUpdated *common.SDKTime `mandatory:"true" json:"timeUpdated"`

	// The time/date of the next scheduled execution of this scheduled job.
	TimeNextExecution *common.SDKTime `mandatory:"true" json:"timeNextExecution"`

	// The list of operations this scheduled job needs to perform (can only support one operation if the operationType is not UPDATE_PACKAGES/UPDATE_ALL/UPDATE_SECURITY/UPDATE_BUGFIX/UPDATE_ENHANCEMENT/UPDATE_OTHER/UPDATE_KSPLICE_USERSPACE/UPDATE_KSPLICE_KERNEL).
	Operations []ScheduledJobOperation `mandatory:"true" json:"operations"`

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

	// The time/date of the last execution of this scheduled job.
	TimeLastExecution *common.SDKTime `mandatory:"false" json:"timeLastExecution"`

	// The list of managed instance OCIDs this scheduled job operates on (mutually exclusive with managedInstanceGroupIds, managedCompartmentIds and lifecycleStageIds).
	ManagedInstanceIds []string `mandatory:"false" json:"managedInstanceIds"`

	// The list of managed instance group OCIDs this scheduled job operates on (mutually exclusive with managedInstances, managedCompartmentIds and lifecycleStageIds).
	ManagedInstanceGroupIds []string `mandatory:"false" json:"managedInstanceGroupIds"`

	// The list of target compartment OCIDs if this scheduled job operates on a compartment level (mutually exclusive with managedInstances, managedInstanceGroupIds and lifecycleStageIds).
	ManagedCompartmentIds []string `mandatory:"false" json:"managedCompartmentIds"`

	// The list of target lifecycle stage OCIDs if this scheduled job operates on lifecycle stages (mutually exclusive with managedInstances, managedInstanceGroupIds and managedCompartmentIds).
	LifecycleStageIds []string `mandatory:"false" json:"lifecycleStageIds"`

	// true, if the schedule job has its update/deletion capabilities restricted. (Used to track scheduled job for management station syncing).
	IsRestricted *bool `mandatory:"false" json:"isRestricted"`

	// System tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"orcl-cloud": {"free-tier-retained": "true"}}`
	SystemTags map[string]map[string]interface{} `mandatory:"false" json:"systemTags"`
}

func (m ScheduledJobSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ScheduledJobSummary) ValidateEnumValue() (bool, error) {
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
