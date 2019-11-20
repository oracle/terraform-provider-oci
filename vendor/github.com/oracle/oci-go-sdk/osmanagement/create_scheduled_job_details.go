// Copyright (c) 2016, 2018, 2019, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// OSMS
//
// OS Management as a Service API definition
//

package osmanagement

import (
	"github.com/oracle/oci-go-sdk/common"
)

// CreateScheduledJobDetails Information for creating a Scheduled Job
type CreateScheduledJobDetails struct {

	// OCID for the Compartment
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// Scheduled Job name
	DisplayName *string `mandatory:"true" json:"displayName"`

	// the type of scheduling this Scheduled Job follows
	ScheduleType ScheduleTypesEnum `mandatory:"true" json:"scheduleType"`

	// the desired time for the next execution of this Scheduled Job
	TimeNextExecution *common.SDKTime `mandatory:"true" json:"timeNextExecution"`

	// the type of operation this Scheduled Job performs
	OperationType OperationTypesEnum `mandatory:"true" json:"operationType"`

	// Details describing the Scheduled Job.
	Description *string `mandatory:"false" json:"description"`

	// the interval period for a recurring Scheduled Job (only if schedule type is RECURRING)
	IntervalType IntervalTypesEnum `mandatory:"false" json:"intervalType,omitempty"`

	// the value for the interval period for a recurring Scheduled Job (only if schedule type is RECURRING)
	IntervalValue *string `mandatory:"false" json:"intervalValue"`

	// The list of managed instances this scheduled job operates on
	// (mutually exclusive with managedInstanceGroups). Either this or the
	// managedInstanceGroups must be supplied.
	ManagedInstances []Id `mandatory:"false" json:"managedInstances"`

	// The list of managed instance groups this scheduled job operates on
	// (mutually exclusive with managedInstances). Either this or
	// managedInstances must be supplied.
	ManagedInstanceGroups []Id `mandatory:"false" json:"managedInstanceGroups"`

	// Type of the update (only if operation type is UPDATE_ALL_PACKAGES)
	UpdateType PackageUpdateTypesEnum `mandatory:"false" json:"updateType,omitempty"`

	// the id of the package (only if operation type is INSTALL/UPDATE/REMOVE_PACKAGE)
	PackageNames []PackageName `mandatory:"false" json:"packageNames"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m CreateScheduledJobDetails) String() string {
	return common.PointerString(m)
}
