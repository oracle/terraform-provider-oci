// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Fleet Application Management Service API
//
// Fleet Application Management provides a centralized platform to help you automate resource management tasks, validate patch compliance, and enhance operational efficiency across an enterprise.
//

package fleetappsmanagement

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// MaintenanceWindow Details of Maintenance Window in Fleet Application Management.
type MaintenanceWindow struct {

	// The OCID of the resource.
	Id *string `mandatory:"true" json:"id"`

	// Tenancy OCID
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// A user-friendly name. Does not have to be unique, and it's changeable.
	// Avoid entering confidential information.
	// Example: `My new resource`
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The time this resource was created. An RFC3339 formatted datetime string.
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// Does the maintenenace window cause outage?
	// An outage indicates whether a maintenance window can consider operations that require downtime.
	// It means a period when the application is not accessible.
	IsOutage *bool `mandatory:"true" json:"isOutage"`

	// Duration of the maintenance window.
	// Specify how long the maintenance window remains open.
	Duration *string `mandatory:"true" json:"duration"`

	// Is this a recurring maintenance window?
	IsRecurring *bool `mandatory:"true" json:"isRecurring"`

	// Task initiation cutoff time for the maintenance window.
	TaskInitiationCutoff *int `mandatory:"true" json:"taskInitiationCutoff"`

	// The current state of the MaintenanceWindow.
	LifecycleState MaintenanceWindowLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"true" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"true" json:"definedTags"`

	// Associated region
	ResourceRegion *string `mandatory:"false" json:"resourceRegion"`

	// A user-friendly description. To provide some insight about the resource.
	// Avoid entering confidential information.
	Description *string `mandatory:"false" json:"description"`

	// The time this resource was last updated. An RFC3339 formatted datetime string.
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`

	// Type of the MaintenanceWindow.
	MaintenanceWindowType MaintenanceWindowTypeEnum `mandatory:"false" json:"maintenanceWindowType,omitempty"`

	// Specify the date and time of the day that the maintenance window starts.
	TimeScheduleStart *common.SDKTime `mandatory:"false" json:"timeScheduleStart"`

	// Recurrence rule specification if maintenance window recurring.
	// Specify the frequency of running the maintenance window.
	Recurrences *string `mandatory:"false" json:"recurrences"`

	// A message describing the current state in more detail. For example, can be used to provide actionable information for a resource in Failed state.
	LifecycleDetails *string `mandatory:"false" json:"lifecycleDetails"`

	// System tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"orcl-cloud": {"free-tier-retained": "true"}}`
	SystemTags map[string]map[string]interface{} `mandatory:"false" json:"systemTags"`
}

func (m MaintenanceWindow) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m MaintenanceWindow) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingMaintenanceWindowLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetMaintenanceWindowLifecycleStateEnumStringValues(), ",")))
	}

	if _, ok := GetMappingMaintenanceWindowTypeEnum(string(m.MaintenanceWindowType)); !ok && m.MaintenanceWindowType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for MaintenanceWindowType: %s. Supported values are: %s.", m.MaintenanceWindowType, strings.Join(GetMaintenanceWindowTypeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MaintenanceWindowLifecycleStateEnum Enum with underlying type: string
type MaintenanceWindowLifecycleStateEnum string

// Set of constants representing the allowable values for MaintenanceWindowLifecycleStateEnum
const (
	MaintenanceWindowLifecycleStateActive         MaintenanceWindowLifecycleStateEnum = "ACTIVE"
	MaintenanceWindowLifecycleStateDeleted        MaintenanceWindowLifecycleStateEnum = "DELETED"
	MaintenanceWindowLifecycleStateFailed         MaintenanceWindowLifecycleStateEnum = "FAILED"
	MaintenanceWindowLifecycleStateDeleting       MaintenanceWindowLifecycleStateEnum = "DELETING"
	MaintenanceWindowLifecycleStateUpdating       MaintenanceWindowLifecycleStateEnum = "UPDATING"
	MaintenanceWindowLifecycleStateNeedsAttention MaintenanceWindowLifecycleStateEnum = "NEEDS_ATTENTION"
)

var mappingMaintenanceWindowLifecycleStateEnum = map[string]MaintenanceWindowLifecycleStateEnum{
	"ACTIVE":          MaintenanceWindowLifecycleStateActive,
	"DELETED":         MaintenanceWindowLifecycleStateDeleted,
	"FAILED":          MaintenanceWindowLifecycleStateFailed,
	"DELETING":        MaintenanceWindowLifecycleStateDeleting,
	"UPDATING":        MaintenanceWindowLifecycleStateUpdating,
	"NEEDS_ATTENTION": MaintenanceWindowLifecycleStateNeedsAttention,
}

var mappingMaintenanceWindowLifecycleStateEnumLowerCase = map[string]MaintenanceWindowLifecycleStateEnum{
	"active":          MaintenanceWindowLifecycleStateActive,
	"deleted":         MaintenanceWindowLifecycleStateDeleted,
	"failed":          MaintenanceWindowLifecycleStateFailed,
	"deleting":        MaintenanceWindowLifecycleStateDeleting,
	"updating":        MaintenanceWindowLifecycleStateUpdating,
	"needs_attention": MaintenanceWindowLifecycleStateNeedsAttention,
}

// GetMaintenanceWindowLifecycleStateEnumValues Enumerates the set of values for MaintenanceWindowLifecycleStateEnum
func GetMaintenanceWindowLifecycleStateEnumValues() []MaintenanceWindowLifecycleStateEnum {
	values := make([]MaintenanceWindowLifecycleStateEnum, 0)
	for _, v := range mappingMaintenanceWindowLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetMaintenanceWindowLifecycleStateEnumStringValues Enumerates the set of values in String for MaintenanceWindowLifecycleStateEnum
func GetMaintenanceWindowLifecycleStateEnumStringValues() []string {
	return []string{
		"ACTIVE",
		"DELETED",
		"FAILED",
		"DELETING",
		"UPDATING",
		"NEEDS_ATTENTION",
	}
}

// GetMappingMaintenanceWindowLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingMaintenanceWindowLifecycleStateEnum(val string) (MaintenanceWindowLifecycleStateEnum, bool) {
	enum, ok := mappingMaintenanceWindowLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
