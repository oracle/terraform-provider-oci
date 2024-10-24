// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
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

// WorkflowComponent Workflow Component Details.
type WorkflowComponent interface {
}

type workflowcomponent struct {
	JsonData []byte
	Type     string `json:"type"`
}

// UnmarshalJSON unmarshals json
func (m *workflowcomponent) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerworkflowcomponent workflowcomponent
	s := struct {
		Model Unmarshalerworkflowcomponent
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.Type = s.Model.Type

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *workflowcomponent) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.Type {
	case "PARALLEL_TASK_GROUP":
		mm := WorkflowGroupComponent{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "TASK":
		mm := WorkflowTaskComponent{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Recieved unsupported enum value for WorkflowComponent: %s.", m.Type)
		return *m, nil
	}
}

func (m workflowcomponent) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m workflowcomponent) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// WorkflowComponentTypeEnum Enum with underlying type: string
type WorkflowComponentTypeEnum string

// Set of constants representing the allowable values for WorkflowComponentTypeEnum
const (
	WorkflowComponentTypeTask              WorkflowComponentTypeEnum = "TASK"
	WorkflowComponentTypeParallelTaskGroup WorkflowComponentTypeEnum = "PARALLEL_TASK_GROUP"
)

var mappingWorkflowComponentTypeEnum = map[string]WorkflowComponentTypeEnum{
	"TASK":                WorkflowComponentTypeTask,
	"PARALLEL_TASK_GROUP": WorkflowComponentTypeParallelTaskGroup,
}

var mappingWorkflowComponentTypeEnumLowerCase = map[string]WorkflowComponentTypeEnum{
	"task":                WorkflowComponentTypeTask,
	"parallel_task_group": WorkflowComponentTypeParallelTaskGroup,
}

// GetWorkflowComponentTypeEnumValues Enumerates the set of values for WorkflowComponentTypeEnum
func GetWorkflowComponentTypeEnumValues() []WorkflowComponentTypeEnum {
	values := make([]WorkflowComponentTypeEnum, 0)
	for _, v := range mappingWorkflowComponentTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetWorkflowComponentTypeEnumStringValues Enumerates the set of values in String for WorkflowComponentTypeEnum
func GetWorkflowComponentTypeEnumStringValues() []string {
	return []string{
		"TASK",
		"PARALLEL_TASK_GROUP",
	}
}

// GetMappingWorkflowComponentTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingWorkflowComponentTypeEnum(val string) (WorkflowComponentTypeEnum, bool) {
	enum, ok := mappingWorkflowComponentTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
