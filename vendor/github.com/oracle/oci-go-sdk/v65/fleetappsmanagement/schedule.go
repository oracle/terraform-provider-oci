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

// Schedule Schedule Information.
type Schedule struct {

	// Schedule Type
	Type ScheduleTypeEnum `mandatory:"true" json:"type"`

	// Start Date for the schedule. An RFC3339 formatted datetime string
	ExecutionStartdate *common.SDKTime `mandatory:"true" json:"executionStartdate"`

	// Provide MaintenanceWindowId if Schedule Type is Maintenance Window
	MaintenanceWindowId *string `mandatory:"false" json:"maintenanceWindowId"`

	// Recurrence rule specification if Schedule Type is Custom and Recurring
	Recurrences *string `mandatory:"false" json:"recurrences"`

	// Duration if schedule type is Custom
	Duration *string `mandatory:"false" json:"duration"`
}

func (m Schedule) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m Schedule) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingScheduleTypeEnum(string(m.Type)); !ok && m.Type != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Type: %s. Supported values are: %s.", m.Type, strings.Join(GetScheduleTypeEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ScheduleTypeEnum Enum with underlying type: string
type ScheduleTypeEnum string

// Set of constants representing the allowable values for ScheduleTypeEnum
const (
	ScheduleTypeCustom            ScheduleTypeEnum = "CUSTOM"
	ScheduleTypeMaintenanceWindow ScheduleTypeEnum = "MAINTENANCE_WINDOW"
)

var mappingScheduleTypeEnum = map[string]ScheduleTypeEnum{
	"CUSTOM":             ScheduleTypeCustom,
	"MAINTENANCE_WINDOW": ScheduleTypeMaintenanceWindow,
}

var mappingScheduleTypeEnumLowerCase = map[string]ScheduleTypeEnum{
	"custom":             ScheduleTypeCustom,
	"maintenance_window": ScheduleTypeMaintenanceWindow,
}

// GetScheduleTypeEnumValues Enumerates the set of values for ScheduleTypeEnum
func GetScheduleTypeEnumValues() []ScheduleTypeEnum {
	values := make([]ScheduleTypeEnum, 0)
	for _, v := range mappingScheduleTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetScheduleTypeEnumStringValues Enumerates the set of values in String for ScheduleTypeEnum
func GetScheduleTypeEnumStringValues() []string {
	return []string{
		"CUSTOM",
		"MAINTENANCE_WINDOW",
	}
}

// GetMappingScheduleTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingScheduleTypeEnum(val string) (ScheduleTypeEnum, bool) {
	enum, ok := mappingScheduleTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
