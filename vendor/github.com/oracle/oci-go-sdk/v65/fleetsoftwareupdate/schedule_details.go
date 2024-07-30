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

// ScheduleDetails Scheduling related details for the Exadata Fleet Update Action.
// The specified time should not conflict with existing Exadata Infrastructure maintenance windows.
// Null scheduleDetails would execute the Exadata Fleet Update Action as soon as possible.
type ScheduleDetails interface {
}

type scheduledetails struct {
	JsonData []byte
	Type     string `json:"type"`
}

// UnmarshalJSON unmarshals json
func (m *scheduledetails) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerscheduledetails scheduledetails
	s := struct {
		Model Unmarshalerscheduledetails
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.Type = s.Model.Type

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *scheduledetails) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.Type {
	case "START_TIME":
		mm := StartTimeScheduleDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Recieved unsupported enum value for ScheduleDetails: %s.", m.Type)
		return *m, nil
	}
}

func (m scheduledetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m scheduledetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ScheduleDetailsTypeEnum Enum with underlying type: string
type ScheduleDetailsTypeEnum string

// Set of constants representing the allowable values for ScheduleDetailsTypeEnum
const (
	ScheduleDetailsTypeStartTime ScheduleDetailsTypeEnum = "START_TIME"
)

var mappingScheduleDetailsTypeEnum = map[string]ScheduleDetailsTypeEnum{
	"START_TIME": ScheduleDetailsTypeStartTime,
}

var mappingScheduleDetailsTypeEnumLowerCase = map[string]ScheduleDetailsTypeEnum{
	"start_time": ScheduleDetailsTypeStartTime,
}

// GetScheduleDetailsTypeEnumValues Enumerates the set of values for ScheduleDetailsTypeEnum
func GetScheduleDetailsTypeEnumValues() []ScheduleDetailsTypeEnum {
	values := make([]ScheduleDetailsTypeEnum, 0)
	for _, v := range mappingScheduleDetailsTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetScheduleDetailsTypeEnumStringValues Enumerates the set of values in String for ScheduleDetailsTypeEnum
func GetScheduleDetailsTypeEnumStringValues() []string {
	return []string{
		"START_TIME",
	}
}

// GetMappingScheduleDetailsTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingScheduleDetailsTypeEnum(val string) (ScheduleDetailsTypeEnum, bool) {
	enum, ok := mappingScheduleDetailsTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
