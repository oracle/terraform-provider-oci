// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Fleet Application Management Service API
//
// Fleet Application Management Service API. Use this API to for all FAMS related activities.
// To manage fleets,view complaince report for the Fleet,scedule patches and other lifecycle activities
//

package fleetappsmanagement

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// UpdateMaintenanceWindowDetails Maintenance Window information to be updated.
type UpdateMaintenanceWindowDetails struct {

	// A user-friendly name. Does not have to be unique, and it's changeable.
	// Avoid entering confidential information.
	// Example: `My new resource`
	DisplayName *string `mandatory:"false" json:"displayName"`

	// A user-friendly description. To provide some insight about the resource.
	// Avoid entering confidential information.
	Description *string `mandatory:"false" json:"description"`

	// MaintenanceWindow Identifier
	MaintenanceWindowType MaintenanceWindowTypeEnum `mandatory:"false" json:"maintenanceWindowType,omitempty"`

	// Does the maintenenace window cause outage?
	IsOutage *bool `mandatory:"false" json:"isOutage"`

	// Start time of schedule
	TimeScheduleStart *common.SDKTime `mandatory:"false" json:"timeScheduleStart"`

	// Duration if schedule type is Custom
	Duration *string `mandatory:"false" json:"duration"`

	// Is this is a recurring maintenance window
	IsRecurring *bool `mandatory:"false" json:"isRecurring"`

	// Recurrence rule specification if recurring
	Recurrences *string `mandatory:"false" json:"recurrences"`

	// Task initiation cutoff
	TaskInitiationCutoff *int `mandatory:"false" json:"taskInitiationCutoff"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m UpdateMaintenanceWindowDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m UpdateMaintenanceWindowDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingMaintenanceWindowTypeEnum(string(m.MaintenanceWindowType)); !ok && m.MaintenanceWindowType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for MaintenanceWindowType: %s. Supported values are: %s.", m.MaintenanceWindowType, strings.Join(GetMaintenanceWindowTypeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
