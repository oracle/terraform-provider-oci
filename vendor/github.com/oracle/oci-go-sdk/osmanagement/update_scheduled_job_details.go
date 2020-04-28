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

	// Type of the update (only if operation type is UPDATEALL)
	UpdateType PackageUpdateTypesEnum `mandatory:"false" json:"updateType,omitempty"`

	// the id of the package (only if operation type is INSTALL/UPDATE/REMOVE)
	PackageNames []PackageName `mandatory:"false" json:"packageNames"`

	// The unique names of the Windows Updates (only if operation type is INSTALL).
	// This is only applicable when the osFamily is for Windows managed instances.
	UpdateNames []string `mandatory:"false" json:"updateNames"`

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
