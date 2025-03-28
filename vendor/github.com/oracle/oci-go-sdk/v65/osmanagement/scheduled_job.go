// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// OS Management API
//
// API for the OS Management service. Use these API operations for working
// with Managed instances and Managed instance groups.
//

package osmanagement

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// ScheduledJob Detailed information about a Scheduled Job
type ScheduledJob struct {

	// OCID for the Scheduled Job
	Id *string `mandatory:"true" json:"id"`

	// Scheduled Job name
	DisplayName *string `mandatory:"true" json:"displayName"`

	// OCID for the Compartment
	CompartmentId *string `mandatory:"false" json:"compartmentId"`

	// Details describing the Scheduled Job.
	Description *string `mandatory:"false" json:"description"`

	// the type of scheduling this Scheduled Job follows
	ScheduleType ScheduleTypesEnum `mandatory:"false" json:"scheduleType,omitempty"`

	// the time of the next execution of this Scheduled Job
	TimeNextExecution *common.SDKTime `mandatory:"false" json:"timeNextExecution"`

	// the time of the last execution of this Scheduled Job
	TimeLastExecution *common.SDKTime `mandatory:"false" json:"timeLastExecution"`

	// the interval period for a recurring Scheduled Job (only if schedule type is RECURRING)
	IntervalType IntervalTypesEnum `mandatory:"false" json:"intervalType,omitempty"`

	// the value for the interval period for a recurring Scheduled Job (only if schedule type is RECURRING)
	IntervalValue *string `mandatory:"false" json:"intervalValue"`

	// the list of managed instances this scheduled job operates on (mutually exclusive with managedInstanceGroups)
	ManagedInstances []Id `mandatory:"false" json:"managedInstances"`

	// the list of managed instance groups this scheduled job operates on (mutually exclusive with managedInstances)
	ManagedInstanceGroups []Id `mandatory:"false" json:"managedInstanceGroups"`

	// the type of operation this Scheduled Job performs
	OperationType OperationTypesEnum `mandatory:"false" json:"operationType,omitempty"`

	// Type of the update (only if operation type is UPDATEALL)
	UpdateType PackageUpdateTypesEnum `mandatory:"false" json:"updateType,omitempty"`

	// the names of the updates (only if operation type is INSTALL/UPDATE/REMOVE)
	PackageNames []PackageName `mandatory:"false" json:"packageNames"`

	// list of Work Requests associated with this Scheduled Job
	WorkRequests []Id `mandatory:"false" json:"workRequests"`

	// The current state of the Scheduled Job.
	LifecycleState LifecycleStatesEnum `mandatory:"false" json:"lifecycleState,omitempty"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// The unique names of the Windows Updates (only if operation type is INSTALL).
	// This is only applicable when the osFamily is for Windows managed instances.
	UpdateNames []string `mandatory:"false" json:"updateNames"`

	// The Operating System type of the managed instance.
	OsFamily OsFamiliesEnum `mandatory:"false" json:"osFamily,omitempty"`

	// true, if the schedule job has its update capabilities restricted. (Used to track Autonomous Scheduled Job)
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
	if _, ok := GetMappingIntervalTypesEnum(string(m.IntervalType)); !ok && m.IntervalType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for IntervalType: %s. Supported values are: %s.", m.IntervalType, strings.Join(GetIntervalTypesEnumStringValues(), ",")))
	}
	if _, ok := GetMappingOperationTypesEnum(string(m.OperationType)); !ok && m.OperationType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for OperationType: %s. Supported values are: %s.", m.OperationType, strings.Join(GetOperationTypesEnumStringValues(), ",")))
	}
	if _, ok := GetMappingPackageUpdateTypesEnum(string(m.UpdateType)); !ok && m.UpdateType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for UpdateType: %s. Supported values are: %s.", m.UpdateType, strings.Join(GetPackageUpdateTypesEnumStringValues(), ",")))
	}
	if _, ok := GetMappingLifecycleStatesEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetLifecycleStatesEnumStringValues(), ",")))
	}
	if _, ok := GetMappingOsFamiliesEnum(string(m.OsFamily)); !ok && m.OsFamily != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for OsFamily: %s. Supported values are: %s.", m.OsFamily, strings.Join(GetOsFamiliesEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
