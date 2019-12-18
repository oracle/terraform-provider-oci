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

// UpdateScheduledJobDetails Information for updating a Scheduled Job
type UpdateScheduledJobDetails struct {

	// Scheduled Job name
	DisplayName *string `mandatory:"false" json:"displayName"`

	// Details describing the Scheduled Job.
	Description *string `mandatory:"false" json:"description"`

	// the type of scheduling this Scheduled Job follows
	ScheduleType ScheduleTypesEnum `mandatory:"false" json:"scheduleType,omitempty"`

	// the desired time for the next execution of this Scheduled Job
	TimeNextExecution *common.SDKTime `mandatory:"false" json:"timeNextExecution"`

	// the interval period for a recurring Scheduled Job (only if schedule type is RECURRING)
	IntervalType IntervalTypesEnum `mandatory:"false" json:"intervalType,omitempty"`

	// the value for the interval period for a recurring Scheduled Job (only if schedule type is RECURRING)
	IntervalValue *string `mandatory:"false" json:"intervalValue"`

	// the type of operation this Scheduled Job performs
	OperationType OperationTypesEnum `mandatory:"false" json:"operationType,omitempty"`

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

func (m UpdateScheduledJobDetails) String() string {
	return common.PointerString(m)
}
