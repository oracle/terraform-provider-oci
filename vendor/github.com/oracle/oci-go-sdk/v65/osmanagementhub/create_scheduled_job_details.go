// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// OS Management Hub API
//
// Use the OS Management Hub API to manage and monitor updates and patches for instances in OCI, your private data center, or 3rd-party clouds.
// For more information, see Overview of OS Management Hub (https://docs.cloud.oracle.com/iaas/osmh/doc/overview.htm).
//

package osmanagementhub

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// CreateScheduledJobDetails Provides the information used to create a scheduled job.
type CreateScheduledJobDetails struct {

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment that contains the scheduled job.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The type of scheduling frequency for the scheduled job.
	ScheduleType ScheduleTypesEnum `mandatory:"true" json:"scheduleType"`

	// The desired time of the next execution of this scheduled job (in RFC 3339 (https://tools.ietf.org/rfc/rfc3339) format).
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

	// User-friendly name for the scheduled job. Does not have to be unique and you can change the name later. Avoid entering confidential information.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// User-specified description of the scheduled job. Avoid entering confidential information.
	Description *string `mandatory:"false" json:"description"`

	// The list of locations this scheduled job should operate on for a job targeting on compartments. (Empty list means apply to all locations). This can only be set when managedCompartmentIds is not empty.
	Locations []ManagedInstanceLocationEnum `mandatory:"false" json:"locations"`

	// The frequency schedule for a recurring scheduled job.
	RecurringRule *string `mandatory:"false" json:"recurringRule"`

	// The managed instance OCIDs (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) that this scheduled job operates on.
	// A scheduled job can only operate on one type of target, therefore you must supply either this or
	// managedInstanceGroupIds, or managedCompartmentIds, or lifecycleStageIds.
	ManagedInstanceIds []string `mandatory:"false" json:"managedInstanceIds"`

	// The managed instance group OCIDs (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) that this scheduled job operates on.
	// A scheduled job can only operate on one type of target, therefore you must supply either this or managedInstanceIds,
	// or managedCompartmentIds, or lifecycleStageIds.
	ManagedInstanceGroupIds []string `mandatory:"false" json:"managedInstanceGroupIds"`

	// The compartment OCIDs (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) that this scheduled job operates on.
	// To apply the job to all compartments in the tenancy, set this to the tenancy OCID (root compartment) and set
	// isSubcompartmentIncluded to true. A scheduled job can only operate on one type of target, therefore you must
	// supply either this or managedInstanceIds, or managedInstanceGroupIds, or lifecycleStageIds.
	ManagedCompartmentIds []string `mandatory:"false" json:"managedCompartmentIds"`

	// The lifecycle stage OCIDs (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) that this scheduled job operates on.
	// A scheduled job can only operate on one type of target, therefore you must supply either this or managedInstanceIds,
	// or managedInstanceGroupIds, or managedCompartmentIds.
	LifecycleStageIds []string `mandatory:"false" json:"lifecycleStageIds"`

	// Indicates whether to apply the scheduled job to all compartments in the tenancy when managedCompartmentIds specifies
	// the tenancy OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) (root compartment).
	IsSubcompartmentIncluded *bool `mandatory:"false" json:"isSubcompartmentIncluded"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// The amount of time in minutes to wait until retrying the scheduled job. If set, the service will automatically
	// retry a failed scheduled job after the interval. For example, you could set the interval to [2,5,10]. If the
	// initial execution of the job fails, the service waits 2 minutes and then retries. If that fails, the service
	// waits 5 minutes and then retries. If that fails, the service waits 10 minutes and then retries.
	RetryIntervals []int `mandatory:"false" json:"retryIntervals"`

	// Indicates whether this scheduled job is managed by the Autonomous Linux service.
	IsManagedByAutonomousLinux *bool `mandatory:"false" json:"isManagedByAutonomousLinux"`
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
