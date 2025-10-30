// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Oracle Cloud Agent API
//
// API for the Oracle Cloud Agent software running on compute instances. Oracle Cloud Agent
// is a lightweight process that monitors and manages compute instances.
//

package computeinstanceagent

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// InstanceAgentCommandContent The contents of the command.
type InstanceAgentCommandContent struct {

	// The source of the command.
	Source InstanceAgentCommandSourceDetails `mandatory:"true" json:"source"`

	// The output destination for the command.
	Output InstanceAgentCommandOutputDetails `mandatory:"false" json:"output"`

	// Command String is a fully formed command that runcommand executes.
	// Example: main.sh is stored in object storage and user provides the following command with parameters to execute
	// /bin/sh main.sh abc 10 foo.sh
	CommandString *string `mandatory:"false" json:"commandString"`
}

func (m InstanceAgentCommandContent) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m InstanceAgentCommandContent) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UnmarshalJSON unmarshals from json
func (m *InstanceAgentCommandContent) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		Output        instanceagentcommandoutputdetails `json:"output"`
		CommandString *string                           `json:"commandString"`
		Source        instanceagentcommandsourcedetails `json:"source"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	nn, e = model.Output.UnmarshalPolymorphicJSON(model.Output.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.Output = nn.(InstanceAgentCommandOutputDetails)
	} else {
		m.Output = nil
	}

	m.CommandString = model.CommandString

	nn, e = model.Source.UnmarshalPolymorphicJSON(model.Source.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.Source = nn.(InstanceAgentCommandSourceDetails)
	} else {
		m.Source = nil
	}

	return
}
