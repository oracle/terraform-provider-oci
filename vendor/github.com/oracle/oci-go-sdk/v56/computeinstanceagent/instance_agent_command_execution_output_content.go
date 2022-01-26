// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
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
	"github.com/oracle/oci-go-sdk/v56/common"
)

// InstanceAgentCommandExecutionOutputContent The execution output from a command.
type InstanceAgentCommandExecutionOutputContent interface {

	// The exit code for the command. Exit code `0` indicates success.
	GetExitCode() *int

	// An optional status message that Oracle Cloud Agent can populate for additional troubleshooting.
	GetMessage() *string
}

type instanceagentcommandexecutionoutputcontent struct {
	JsonData   []byte
	ExitCode   *int    `mandatory:"true" json:"exitCode"`
	Message    *string `mandatory:"false" json:"message"`
	OutputType string  `json:"outputType"`
}

// UnmarshalJSON unmarshals json
func (m *instanceagentcommandexecutionoutputcontent) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerinstanceagentcommandexecutionoutputcontent instanceagentcommandexecutionoutputcontent
	s := struct {
		Model Unmarshalerinstanceagentcommandexecutionoutputcontent
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.ExitCode = s.Model.ExitCode
	m.Message = s.Model.Message
	m.OutputType = s.Model.OutputType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *instanceagentcommandexecutionoutputcontent) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.OutputType {
	case "TEXT":
		mm := InstanceAgentCommandExecutionOutputViaTextDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "OBJECT_STORAGE_URI":
		mm := InstanceAgentCommandExecutionOutputViaObjectStorageUriDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "OBJECT_STORAGE_TUPLE":
		mm := InstanceAgentCommandExecutionOutputViaObjectStorageTupleDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		return *m, nil
	}
}

//GetExitCode returns ExitCode
func (m instanceagentcommandexecutionoutputcontent) GetExitCode() *int {
	return m.ExitCode
}

//GetMessage returns Message
func (m instanceagentcommandexecutionoutputcontent) GetMessage() *string {
	return m.Message
}

func (m instanceagentcommandexecutionoutputcontent) String() string {
	return common.PointerString(m)
}

// InstanceAgentCommandExecutionOutputContentOutputTypeEnum Enum with underlying type: string
type InstanceAgentCommandExecutionOutputContentOutputTypeEnum string

// Set of constants representing the allowable values for InstanceAgentCommandExecutionOutputContentOutputTypeEnum
const (
	InstanceAgentCommandExecutionOutputContentOutputTypeText               InstanceAgentCommandExecutionOutputContentOutputTypeEnum = "TEXT"
	InstanceAgentCommandExecutionOutputContentOutputTypeObjectStorageUri   InstanceAgentCommandExecutionOutputContentOutputTypeEnum = "OBJECT_STORAGE_URI"
	InstanceAgentCommandExecutionOutputContentOutputTypeObjectStorageTuple InstanceAgentCommandExecutionOutputContentOutputTypeEnum = "OBJECT_STORAGE_TUPLE"
)

var mappingInstanceAgentCommandExecutionOutputContentOutputType = map[string]InstanceAgentCommandExecutionOutputContentOutputTypeEnum{
	"TEXT":                 InstanceAgentCommandExecutionOutputContentOutputTypeText,
	"OBJECT_STORAGE_URI":   InstanceAgentCommandExecutionOutputContentOutputTypeObjectStorageUri,
	"OBJECT_STORAGE_TUPLE": InstanceAgentCommandExecutionOutputContentOutputTypeObjectStorageTuple,
}

// GetInstanceAgentCommandExecutionOutputContentOutputTypeEnumValues Enumerates the set of values for InstanceAgentCommandExecutionOutputContentOutputTypeEnum
func GetInstanceAgentCommandExecutionOutputContentOutputTypeEnumValues() []InstanceAgentCommandExecutionOutputContentOutputTypeEnum {
	values := make([]InstanceAgentCommandExecutionOutputContentOutputTypeEnum, 0)
	for _, v := range mappingInstanceAgentCommandExecutionOutputContentOutputType {
		values = append(values, v)
	}
	return values
}
