// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Exadata Fleet Update service API
//
// Use the Exadata Fleet Update service to patch large collections of components directly,
// as a single entity, orchestrating the maintenance actions to update all chosen components in the stack in a single cycle.
//

package fleetsoftwareupdate

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// CreateScheduleDetails Scheduling related details for the Exadata Fleet Update Action during create operations.
// The specified time should not conflict with existing Exadata Infrastructure maintenance windows.
// Null scheduleDetails for Stage and Apply Actions in Exadata Fleet Update Cycle creation would not create Actions.
// Null scheduleDetails for CreateAction would execute the Exadata Fleet Update Action as soon as possible.
type CreateScheduleDetails interface {
}

type createscheduledetails struct {
	JsonData []byte
	Type     string `json:"type"`
}

// UnmarshalJSON unmarshals json
func (m *createscheduledetails) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalercreatescheduledetails createscheduledetails
	s := struct {
		Model Unmarshalercreatescheduledetails
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.Type = s.Model.Type

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *createscheduledetails) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.Type {
	case "START_TIME":
		mm := CreateStartTimeScheduleDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Recieved unsupported enum value for CreateScheduleDetails: %s.", m.Type)
		return *m, nil
	}
}

func (m createscheduledetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m createscheduledetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// CreateScheduleDetailsTypeEnum Enum with underlying type: string
type CreateScheduleDetailsTypeEnum string

// Set of constants representing the allowable values for CreateScheduleDetailsTypeEnum
const (
	CreateScheduleDetailsTypeStartTime CreateScheduleDetailsTypeEnum = "START_TIME"
)

var mappingCreateScheduleDetailsTypeEnum = map[string]CreateScheduleDetailsTypeEnum{
	"START_TIME": CreateScheduleDetailsTypeStartTime,
}

var mappingCreateScheduleDetailsTypeEnumLowerCase = map[string]CreateScheduleDetailsTypeEnum{
	"start_time": CreateScheduleDetailsTypeStartTime,
}

// GetCreateScheduleDetailsTypeEnumValues Enumerates the set of values for CreateScheduleDetailsTypeEnum
func GetCreateScheduleDetailsTypeEnumValues() []CreateScheduleDetailsTypeEnum {
	values := make([]CreateScheduleDetailsTypeEnum, 0)
	for _, v := range mappingCreateScheduleDetailsTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetCreateScheduleDetailsTypeEnumStringValues Enumerates the set of values in String for CreateScheduleDetailsTypeEnum
func GetCreateScheduleDetailsTypeEnumStringValues() []string {
	return []string{
		"START_TIME",
	}
}

// GetMappingCreateScheduleDetailsTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingCreateScheduleDetailsTypeEnum(val string) (CreateScheduleDetailsTypeEnum, bool) {
	enum, ok := mappingCreateScheduleDetailsTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
