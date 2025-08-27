// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Science API
//
// Use the Data Science API to organize your data science work, access data and computing resources, and build, train, deploy and manage models and model deployments. For more information, see Data Science (https://docs.oracle.com/iaas/data-science/using/data-science.htm).
//

package datascience

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// ScheduleAction The schedule action
type ScheduleAction interface {
}

type scheduleaction struct {
	JsonData   []byte
	ActionType string `json:"actionType"`
}

// UnmarshalJSON unmarshals json
func (m *scheduleaction) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerscheduleaction scheduleaction
	s := struct {
		Model Unmarshalerscheduleaction
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.ActionType = s.Model.ActionType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *scheduleaction) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.ActionType {
	case "HTTP":
		mm := ScheduleHttpAction{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Received unsupported enum value for ScheduleAction: %s.", m.ActionType)
		return *m, nil
	}
}

func (m scheduleaction) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m scheduleaction) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ScheduleActionActionTypeEnum Enum with underlying type: string
type ScheduleActionActionTypeEnum string

// Set of constants representing the allowable values for ScheduleActionActionTypeEnum
const (
	ScheduleActionActionTypeHttp ScheduleActionActionTypeEnum = "HTTP"
)

var mappingScheduleActionActionTypeEnum = map[string]ScheduleActionActionTypeEnum{
	"HTTP": ScheduleActionActionTypeHttp,
}

var mappingScheduleActionActionTypeEnumLowerCase = map[string]ScheduleActionActionTypeEnum{
	"http": ScheduleActionActionTypeHttp,
}

// GetScheduleActionActionTypeEnumValues Enumerates the set of values for ScheduleActionActionTypeEnum
func GetScheduleActionActionTypeEnumValues() []ScheduleActionActionTypeEnum {
	values := make([]ScheduleActionActionTypeEnum, 0)
	for _, v := range mappingScheduleActionActionTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetScheduleActionActionTypeEnumStringValues Enumerates the set of values in String for ScheduleActionActionTypeEnum
func GetScheduleActionActionTypeEnumStringValues() []string {
	return []string{
		"HTTP",
	}
}

// GetMappingScheduleActionActionTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingScheduleActionActionTypeEnum(val string) (ScheduleActionActionTypeEnum, bool) {
	enum, ok := mappingScheduleActionActionTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
