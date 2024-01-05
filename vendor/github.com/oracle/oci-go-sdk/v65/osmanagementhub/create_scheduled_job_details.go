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

// CreateScheduledJobDetails Information for creating a scheduled job.
type CreateScheduledJobDetails struct {

	// The OCID of the compartment containing the scheduled job.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The type of scheduling this scheduled job follows.
	ScheduleType ScheduleTypesEnum `mandatory:"true" json:"scheduleType"`

	// The desired time for the next execution of this scheduled job.
	TimeNextExecution *common.SDKTime `mandatory:"true" json:"timeNextExecution"`

	// The list of operations this scheduled job needs to perform (can only support one operation if the operationType is not UPDATE_PACKAGES/UPDATE_ALL/UPDATE_SECURITY/UPDATE_BUGFIX/UPDATE_ENHANCEMENT/UPDATE_OTHER/UPDATE_KSPLICE_USERSPACE/UPDATE_KSPLICE_KERNEL).
	Operations []ScheduledJobOperation `mandatory:"true" json:"operations"`

	// Scheduled job name.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// Details describing the scheduled job.
	Description *string `mandatory:"false" json:"description"`

	// The recurring rule for a recurring scheduled job.
	RecurringRule *string `mandatory:"false" json:"recurringRule"`

	// The list of managed instance OCIDs this scheduled job operates on. Either this or
	// managedInstanceGroupIds, or managedCompartmentIds, or lifecycleStageIds must be supplied.
	ManagedInstanceIds []string `mandatory:"false" json:"managedInstanceIds"`

	// The list of managed instance group OCIDs this scheduled job operates on. Either this or
	// managedInstanceIds, or managedCompartmentIds, or lifecycleStageIds must be supplied.
	ManagedInstanceGroupIds []string `mandatory:"false" json:"managedInstanceGroupIds"`

	// The list of target compartment OCIDs if this scheduled job operates on a compartment level.
	// Either this or managedInstanceIds, or managedInstanceGroupIds, or lifecycleStageIds must be supplied.
	ManagedCompartmentIds []string `mandatory:"false" json:"managedCompartmentIds"`

	// The list of lifecycle stage OCIDs this scheduled job operates on. Either this or
	// managedInstanceIds, or managedInstanceGroupIds, or managedCompartmentIds must be supplied.
	LifecycleStageIds []string `mandatory:"false" json:"lifecycleStageIds"`

	// Whether to create jobs for all compartments in the tenancy when managedCompartmentIds specifies the tenancy OCID.
	IsSubcompartmentIncluded *bool `mandatory:"false" json:"isSubcompartmentIncluded"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m CreateScheduledJobDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CreateScheduledJobDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingScheduleTypesEnum(string(m.ScheduleType)); !ok && m.ScheduleType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ScheduleType: %s. Supported values are: %s.", m.ScheduleType, strings.Join(GetScheduleTypesEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
