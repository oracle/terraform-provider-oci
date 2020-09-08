// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Service Connector Hub API
//
// Use the Service Connector Hub API to transfer data between services in Oracle Cloud Infrastructure.
// For more information about Service Connector Hub, see
// Service Connector Hub Overview (https://docs.cloud.oracle.com/iaas/service-connector-hub/using/index.htm).
//

package sch

import (
	"encoding/json"
	"github.com/oracle/oci-go-sdk/common"
)

// TaskDetails An object that represents a task within the flow defined by the service connector.
// An example task is a filter for error logs.
// For more information about flows defined by service connectors, see
// Service Connector Hub Overview (https://docs.cloud.oracle.com/iaas/service-connector-hub/using/index.htm).
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
	case "logRule":
		mm := LogRuleTaskDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		return *m, nil
	}
}

func (m taskdetails) String() string {
	return common.PointerString(m)
}

// TaskDetailsKindEnum Enum with underlying type: string
type TaskDetailsKindEnum string

// Set of constants representing the allowable values for TaskDetailsKindEnum
const (
	TaskDetailsKindLogrule TaskDetailsKindEnum = "logRule"
)

var mappingTaskDetailsKind = map[string]TaskDetailsKindEnum{
	"logRule": TaskDetailsKindLogrule,
}

// GetTaskDetailsKindEnumValues Enumerates the set of values for TaskDetailsKindEnum
func GetTaskDetailsKindEnumValues() []TaskDetailsKindEnum {
	values := make([]TaskDetailsKindEnum, 0)
	for _, v := range mappingTaskDetailsKind {
		values = append(values, v)
	}
	return values
}
