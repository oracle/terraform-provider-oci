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

// TaskVariable The variable of the task.
// At least one of the dynamicArguments or output needs to be provided.
type TaskVariable struct {

	// The input variables for the task.
	InputVariables []InputArgument `mandatory:"false" json:"inputVariables"`

	// The list of output variables.
	OutputVariables []string `mandatory:"false" json:"outputVariables"`
}

func (m TaskVariable) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m TaskVariable) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UnmarshalJSON unmarshals from json
func (m *TaskVariable) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		InputVariables  []inputargument `json:"inputVariables"`
		OutputVariables []string        `json:"outputVariables"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.InputVariables = make([]InputArgument, len(model.InputVariables))
	for i, n := range model.InputVariables {
		nn, e = n.UnmarshalPolymorphicJSON(n.JsonData)
		if e != nil {
			return e
		}
		if nn != nil {
			m.InputVariables[i] = nn.(InputArgument)
		} else {
			m.InputVariables[i] = nil
		}
	}
	m.OutputVariables = make([]string, len(model.OutputVariables))
	copy(m.OutputVariables, model.OutputVariables)
	return
}
