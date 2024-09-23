// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Fleet Application Management Service API
//
// Fleet Application Management Service API. Use this API to for all FAMS related activities.
// To manage fleets,view complaince report for the Fleet,scedule patches and other lifecycle activities
//

package fleetappsmanagement

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// ScriptBasedExecutionDetails Details for script based execution
type ScriptBasedExecutionDetails struct {
	Variables *TaskVariable `mandatory:"false" json:"variables"`

	Content ContentDetails `mandatory:"false" json:"content"`

	// Optional Command to execute the content.
	Command *string `mandatory:"false" json:"command"`
}

func (m ScriptBasedExecutionDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ScriptBasedExecutionDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m ScriptBasedExecutionDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeScriptBasedExecutionDetails ScriptBasedExecutionDetails
	s := struct {
		DiscriminatorParam string `json:"executionType"`
		MarshalTypeScriptBasedExecutionDetails
	}{
		"SCRIPT",
		(MarshalTypeScriptBasedExecutionDetails)(m),
	}

	return json.Marshal(&s)
}

// UnmarshalJSON unmarshals from json
func (m *ScriptBasedExecutionDetails) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		Variables *TaskVariable  `json:"variables"`
		Content   contentdetails `json:"content"`
		Command   *string        `json:"command"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.Variables = model.Variables

	nn, e = model.Content.UnmarshalPolymorphicJSON(model.Content.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.Content = nn.(ContentDetails)
	} else {
		m.Content = nil
	}

	m.Command = model.Command

	return
}
