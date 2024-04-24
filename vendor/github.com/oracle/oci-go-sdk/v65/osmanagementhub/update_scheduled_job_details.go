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

// UpdateScheduledJobDetails Provides the information used to update a scheduled job.
type UpdateScheduledJobDetails struct {

	// User-friendly name for the scheduled job. Avoid entering confidential information.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// User-specified description for the scheduled job. Avoid entering confidential information.
	Description *string `mandatory:"false" json:"description"`

	// The type of scheduling frequency for the job.
	ScheduleType ScheduleTypesEnum `mandatory:"false" json:"scheduleType,omitempty"`

	// The desired time of the next execution of this scheduled job (in RFC 3339 (https://tools.ietf.org/rfc/rfc3339) format).
	TimeNextExecution *common.SDKTime `mandatory:"false" json:"timeNextExecution"`

	// The frequency schedule for a recurring scheduled job.
	RecurringRule *string `mandatory:"false" json:"recurringRule"`

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
	Operations []ScheduledJobOperation `mandatory:"false" json:"operations"`

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
}

func (m UpdateScheduledJobDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m UpdateScheduledJobDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingScheduleTypesEnum(string(m.ScheduleType)); !ok && m.ScheduleType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ScheduleType: %s. Supported values are: %s.", m.ScheduleType, strings.Join(GetScheduleTypesEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
