// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Fleet Application Management Service API
//
// Fleet Application Management provides a centralized platform to help you automate resource management tasks, validate patch compliance, and enhance operational efficiency across an enterprise.
//

package fleetappsmanagement

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// Schedule Schedule Information.
type Schedule interface {

	// Start Date for the schedule. An RFC3339 formatted datetime string
	GetExecutionStartdate() *common.SDKTime
}

type schedule struct {
	JsonData           []byte
	ExecutionStartdate *common.SDKTime `mandatory:"true" json:"executionStartdate"`
	Type               string          `json:"type"`
}

// UnmarshalJSON unmarshals json
func (m *schedule) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerschedule schedule
	s := struct {
		Model Unmarshalerschedule
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.ExecutionStartdate = s.Model.ExecutionStartdate
	m.Type = s.Model.Type

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *schedule) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.Type {
	case "CUSTOM":
		mm := CustomSchedule{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "MAINTENANCE_WINDOW":
		mm := MaintenanceWindowSchedule{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Received unsupported enum value for Schedule: %s.", m.Type)
		return *m, nil
	}
}

// GetExecutionStartdate returns ExecutionStartdate
func (m schedule) GetExecutionStartdate() *common.SDKTime {
	return m.ExecutionStartdate
}

func (m schedule) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m schedule) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
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
