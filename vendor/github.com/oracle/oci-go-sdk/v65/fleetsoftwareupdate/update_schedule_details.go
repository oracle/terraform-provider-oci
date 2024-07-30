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

// UpdateScheduleDetails Scheduling related details for the Exadata Fleet Update Action.
// The specified time should not conflict with existing Exadata Infrastructure maintenance windows.
// 'NONE' type scheduleDetails for UpdateAction would execute the Exadata Fleet Update Action as soon as possible.
type UpdateScheduleDetails interface {
}

type updatescheduledetails struct {
	JsonData []byte
	Type     string `json:"type"`
}

// UnmarshalJSON unmarshals json
func (m *updatescheduledetails) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerupdatescheduledetails updatescheduledetails
	s := struct {
		Model Unmarshalerupdatescheduledetails
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.Type = s.Model.Type

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *updatescheduledetails) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.Type {
	case "NONE":
		mm := NoneScheduleDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "START_TIME":
		mm := UpdateStartTimeScheduleDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Recieved unsupported enum value for UpdateScheduleDetails: %s.", m.Type)
		return *m, nil
	}
}

func (m updatescheduledetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m updatescheduledetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UpdateScheduleDetailsTypeEnum Enum with underlying type: string
type UpdateScheduleDetailsTypeEnum string

// Set of constants representing the allowable values for UpdateScheduleDetailsTypeEnum
const (
	UpdateScheduleDetailsTypeStartTime UpdateScheduleDetailsTypeEnum = "START_TIME"
	UpdateScheduleDetailsTypeNone      UpdateScheduleDetailsTypeEnum = "NONE"
)

var mappingUpdateScheduleDetailsTypeEnum = map[string]UpdateScheduleDetailsTypeEnum{
	"START_TIME": UpdateScheduleDetailsTypeStartTime,
	"NONE":       UpdateScheduleDetailsTypeNone,
}

var mappingUpdateScheduleDetailsTypeEnumLowerCase = map[string]UpdateScheduleDetailsTypeEnum{
	"start_time": UpdateScheduleDetailsTypeStartTime,
	"none":       UpdateScheduleDetailsTypeNone,
}

// GetUpdateScheduleDetailsTypeEnumValues Enumerates the set of values for UpdateScheduleDetailsTypeEnum
func GetUpdateScheduleDetailsTypeEnumValues() []UpdateScheduleDetailsTypeEnum {
	values := make([]UpdateScheduleDetailsTypeEnum, 0)
	for _, v := range mappingUpdateScheduleDetailsTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetUpdateScheduleDetailsTypeEnumStringValues Enumerates the set of values in String for UpdateScheduleDetailsTypeEnum
func GetUpdateScheduleDetailsTypeEnumStringValues() []string {
	return []string{
		"START_TIME",
		"NONE",
	}
}

// GetMappingUpdateScheduleDetailsTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingUpdateScheduleDetailsTypeEnum(val string) (UpdateScheduleDetailsTypeEnum, bool) {
	enum, ok := mappingUpdateScheduleDetailsTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
