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

// RunOnDetails The runon conditions
type RunOnDetails interface {
}

type runondetails struct {
	JsonData []byte
	Kind     string `json:"kind"`
}

// UnmarshalJSON unmarshals json
func (m *runondetails) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerrunondetails runondetails
	s := struct {
		Model Unmarshalerrunondetails
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.Kind = s.Model.Kind

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *runondetails) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.Kind {
	case "SELF_HOSTED_INSTANCES":
		mm := SelfHostedInstanceRunOnDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "SCHEDULE_INSTANCES":
		mm := ScheduleInstanceRunOnDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "PREVIOUS_TASK_INSTANCES":
		mm := PreviousTaskInstanceRunOnDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Received unsupported enum value for RunOnDetails: %s.", m.Kind)
		return *m, nil
	}
}

func (m runondetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m runondetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// RunOnDetailsKindEnum Enum with underlying type: string
type RunOnDetailsKindEnum string

// Set of constants representing the allowable values for RunOnDetailsKindEnum
const (
	RunOnDetailsKindScheduledInstances    RunOnDetailsKindEnum = "SCHEDULED_INSTANCES"
	RunOnDetailsKindSelfHostedInstances   RunOnDetailsKindEnum = "SELF_HOSTED_INSTANCES"
	RunOnDetailsKindPreviousTaskInstances RunOnDetailsKindEnum = "PREVIOUS_TASK_INSTANCES"
)

var mappingRunOnDetailsKindEnum = map[string]RunOnDetailsKindEnum{
	"SCHEDULED_INSTANCES":     RunOnDetailsKindScheduledInstances,
	"SELF_HOSTED_INSTANCES":   RunOnDetailsKindSelfHostedInstances,
	"PREVIOUS_TASK_INSTANCES": RunOnDetailsKindPreviousTaskInstances,
}

var mappingRunOnDetailsKindEnumLowerCase = map[string]RunOnDetailsKindEnum{
	"scheduled_instances":     RunOnDetailsKindScheduledInstances,
	"self_hosted_instances":   RunOnDetailsKindSelfHostedInstances,
	"previous_task_instances": RunOnDetailsKindPreviousTaskInstances,
}

// GetRunOnDetailsKindEnumValues Enumerates the set of values for RunOnDetailsKindEnum
func GetRunOnDetailsKindEnumValues() []RunOnDetailsKindEnum {
	values := make([]RunOnDetailsKindEnum, 0)
	for _, v := range mappingRunOnDetailsKindEnum {
		values = append(values, v)
	}
	return values
}

// GetRunOnDetailsKindEnumStringValues Enumerates the set of values in String for RunOnDetailsKindEnum
func GetRunOnDetailsKindEnumStringValues() []string {
	return []string{
		"SCHEDULED_INSTANCES",
		"SELF_HOSTED_INSTANCES",
		"PREVIOUS_TASK_INSTANCES",
	}
}

// GetMappingRunOnDetailsKindEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingRunOnDetailsKindEnum(val string) (RunOnDetailsKindEnum, bool) {
	enum, ok := mappingRunOnDetailsKindEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
