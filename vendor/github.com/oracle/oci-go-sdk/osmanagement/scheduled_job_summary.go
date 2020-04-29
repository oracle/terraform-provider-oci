// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// OS Management API
//
// API for the OS Management service. Use these API operations for working
// with Managed instances and Managed instance groups.
//

package osmanagement

import (
	"github.com/oracle/oci-go-sdk/common"
)

// ScheduledJobSummary Basic information about a Scheduled Job
type ScheduledJobSummary struct {

	// OCID for the Scheduled Job
	Id *string `mandatory:"true" json:"id"`

	// Scheduled Job name
	DisplayName *string `mandatory:"true" json:"displayName"`

	// OCID for the Compartment
	CompartmentId *string `mandatory:"false" json:"compartmentId"`

	// the type of scheduling this Scheduled Job follows
	ScheduleType ScheduleTypesEnum `mandatory:"false" json:"scheduleType,omitempty"`

	// the time/date of the next scheduled execution of this Scheduled Job
	TimeNextExecution *common.SDKTime `mandatory:"false" json:"timeNextExecution"`

	// the time/date of the last execution of this Scheduled Job
	TimeLastExecution *common.SDKTime `mandatory:"false" json:"timeLastExecution"`

	// the list of managed instances this scheduled job operates on (mutually exclusive with managedInstanceGroups)
	ManagedInstances []Id `mandatory:"false" json:"managedInstances"`

	// the list of managed instance groups this scheduled job operates on (mutually exclusive with managedInstances)
	ManagedInstanceGroups []Id `mandatory:"false" json:"managedInstanceGroups"`

	// the type of operation this Scheduled Job performs
	OperationType OperationTypesEnum `mandatory:"false" json:"operationType,omitempty"`

	// The current state of the Scheduled Job.
	LifecycleState LifecycleStatesEnum `mandatory:"false" json:"lifecycleState,omitempty"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// The Operating System type of the managed instance.
	OsFamily OsFamiliesEnum `mandatory:"false" json:"osFamily,omitempty"`
}

func (m ScheduledJobSummary) String() string {
	return common.PointerString(m)
}
