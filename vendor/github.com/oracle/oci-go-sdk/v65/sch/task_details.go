// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Connector Hub API
//
// Use the Connector Hub API to transfer data between services in Oracle Cloud Infrastructure.
// For more information about Connector Hub, see
// the Connector Hub documentation (https://docs.oracle.com/iaas/Content/connector-hub/home.htm).
// Connector Hub is formerly known as Service Connector Hub.
//

package sch

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// TaskDetails An object that represents a task within the flow defined by the connector.
// An example task is a filter for error logs.
// For more information about flows defined by connectors, see
// Overview of Connector Hub (https://docs.oracle.com/iaas/Content/connector-hub/overview.htm).
// For configuration instructions, see
// Creating a Connector (https://docs.oracle.com/iaas/Content/connector-hub/create-service-connector.htm).
type TaskDetails interface {
}

type taskdetails struct {
	JsonData []byte
	Kind     string `json:"kind"`
}

// UnmarshalJSON unmarshals json
func (m *taskdetails) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalertaskdetails taskdetails
	s := struct {
		Model Unmarshalertaskdetails
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.Kind = s.Model.Kind

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *taskdetails) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.Kind {
	case "function":
		mm := FunctionTaskDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "logRule":
		mm := LogRuleTaskDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Received unsupported enum value for TaskDetails: %s.", m.Kind)
		return *m, nil
	}
}

func (m taskdetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m taskdetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// TaskDetailsKindEnum Enum with underlying type: string
type TaskDetailsKindEnum string

// Set of constants representing the allowable values for TaskDetailsKindEnum
const (
	TaskDetailsKindFunction TaskDetailsKindEnum = "function"
	TaskDetailsKindLogrule  TaskDetailsKindEnum = "logRule"
)

var mappingTaskDetailsKindEnum = map[string]TaskDetailsKindEnum{
	"function": TaskDetailsKindFunction,
	"logRule":  TaskDetailsKindLogrule,
}

var mappingTaskDetailsKindEnumLowerCase = map[string]TaskDetailsKindEnum{
	"function": TaskDetailsKindFunction,
	"logrule":  TaskDetailsKindLogrule,
}

// GetTaskDetailsKindEnumValues Enumerates the set of values for TaskDetailsKindEnum
func GetTaskDetailsKindEnumValues() []TaskDetailsKindEnum {
	values := make([]TaskDetailsKindEnum, 0)
	for _, v := range mappingTaskDetailsKindEnum {
		values = append(values, v)
	}
	return values
}

// GetTaskDetailsKindEnumStringValues Enumerates the set of values in String for TaskDetailsKindEnum
func GetTaskDetailsKindEnumStringValues() []string {
	return []string{
		"function",
		"logRule",
	}
}

// GetMappingTaskDetailsKindEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingTaskDetailsKindEnum(val string) (TaskDetailsKindEnum, bool) {
	enum, ok := mappingTaskDetailsKindEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
