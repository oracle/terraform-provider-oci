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

// Task The details of the task.
type Task struct {

	// The name of the task step.
	StepName *string `mandatory:"true" json:"stepName"`

	TaskRecordDetails AssociatedTaskDetails `mandatory:"true" json:"taskRecordDetails"`

	StepProperties *ComponentProperties `mandatory:"false" json:"stepProperties"`

	// Mapping output variables of previous tasks to the input variables of the current task.
	OutputVariableMappings []OutputVariableMapping `mandatory:"false" json:"outputVariableMappings"`
}

func (m Task) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m Task) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UnmarshalJSON unmarshals from json
func (m *Task) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		StepProperties         *ComponentProperties    `json:"stepProperties"`
		OutputVariableMappings []OutputVariableMapping `json:"outputVariableMappings"`
		StepName               *string                 `json:"stepName"`
		TaskRecordDetails      associatedtaskdetails   `json:"taskRecordDetails"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.StepProperties = model.StepProperties

	m.OutputVariableMappings = make([]OutputVariableMapping, len(model.OutputVariableMappings))
	copy(m.OutputVariableMappings, model.OutputVariableMappings)
	m.StepName = model.StepName

	nn, e = model.TaskRecordDetails.UnmarshalPolymorphicJSON(model.TaskRecordDetails.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.TaskRecordDetails = nn.(AssociatedTaskDetails)
	} else {
		m.TaskRecordDetails = nil
	}

	return
}
