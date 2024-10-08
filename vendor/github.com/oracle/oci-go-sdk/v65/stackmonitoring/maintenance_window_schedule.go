// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Stack Monitoring API
//
// Stack Monitoring API.
//

package stackmonitoring

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// MaintenanceWindowSchedule Schedule information of the Maintenance Window
type MaintenanceWindowSchedule interface {
}

type maintenancewindowschedule struct {
	JsonData     []byte
	ScheduleType string `json:"scheduleType"`
}

// UnmarshalJSON unmarshals json
func (m *maintenancewindowschedule) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalermaintenancewindowschedule maintenancewindowschedule
	s := struct {
		Model Unmarshalermaintenancewindowschedule
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.ScheduleType = s.Model.ScheduleType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *maintenancewindowschedule) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.ScheduleType {
	case "RECURRENT":
		mm := RecurrentMaintenanceWindowSchedule{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "ONE_TIME":
		mm := OneTimeMaintenanceWindowSchedule{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Recieved unsupported enum value for MaintenanceWindowSchedule: %s.", m.ScheduleType)
		return *m, nil
	}
}

func (m maintenancewindowschedule) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m maintenancewindowschedule) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MaintenanceWindowScheduleScheduleTypeEnum Enum with underlying type: string
type MaintenanceWindowScheduleScheduleTypeEnum string

// Set of constants representing the allowable values for MaintenanceWindowScheduleScheduleTypeEnum
const (
	MaintenanceWindowScheduleScheduleTypeOneTime   MaintenanceWindowScheduleScheduleTypeEnum = "ONE_TIME"
	MaintenanceWindowScheduleScheduleTypeRecurrent MaintenanceWindowScheduleScheduleTypeEnum = "RECURRENT"
)

var mappingMaintenanceWindowScheduleScheduleTypeEnum = map[string]MaintenanceWindowScheduleScheduleTypeEnum{
	"ONE_TIME":  MaintenanceWindowScheduleScheduleTypeOneTime,
	"RECURRENT": MaintenanceWindowScheduleScheduleTypeRecurrent,
}

var mappingMaintenanceWindowScheduleScheduleTypeEnumLowerCase = map[string]MaintenanceWindowScheduleScheduleTypeEnum{
	"one_time":  MaintenanceWindowScheduleScheduleTypeOneTime,
	"recurrent": MaintenanceWindowScheduleScheduleTypeRecurrent,
}

// GetMaintenanceWindowScheduleScheduleTypeEnumValues Enumerates the set of values for MaintenanceWindowScheduleScheduleTypeEnum
func GetMaintenanceWindowScheduleScheduleTypeEnumValues() []MaintenanceWindowScheduleScheduleTypeEnum {
	values := make([]MaintenanceWindowScheduleScheduleTypeEnum, 0)
	for _, v := range mappingMaintenanceWindowScheduleScheduleTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetMaintenanceWindowScheduleScheduleTypeEnumStringValues Enumerates the set of values in String for MaintenanceWindowScheduleScheduleTypeEnum
func GetMaintenanceWindowScheduleScheduleTypeEnumStringValues() []string {
	return []string{
		"ONE_TIME",
		"RECURRENT",
	}
}

// GetMappingMaintenanceWindowScheduleScheduleTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingMaintenanceWindowScheduleScheduleTypeEnum(val string) (MaintenanceWindowScheduleScheduleTypeEnum, bool) {
	enum, ok := mappingMaintenanceWindowScheduleScheduleTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
