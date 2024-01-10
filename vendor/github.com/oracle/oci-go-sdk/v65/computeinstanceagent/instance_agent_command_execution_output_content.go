// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
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

// InstanceAgentCommandExecutionOutputContent The execution output from a command.
type InstanceAgentCommandExecutionOutputContent interface {

	// The exit code for the command. Exit code `0` indicates success.
	GetExitCode() *int

	// An optional status message that Oracle Cloud Agent can populate for additional troubleshooting.
	GetMessage() *string
}

type instanceagentcommandexecutionoutputcontent struct {
	JsonData   []byte
	Message    *string `mandatory:"false" json:"message"`
	ExitCode   *int    `mandatory:"true" json:"exitCode"`
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
		common.Logf("Recieved unsupported enum value for InstanceAgentCommandExecutionOutputContent: %s.", m.OutputType)
		return *m, nil
	}
}

// GetMessage returns Message
func (m instanceagentcommandexecutionoutputcontent) GetMessage() *string {
	return m.Message
}

// GetExitCode returns ExitCode
func (m instanceagentcommandexecutionoutputcontent) GetExitCode() *int {
	return m.ExitCode
}

func (m instanceagentcommandexecutionoutputcontent) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m instanceagentcommandexecutionoutputcontent) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// InstanceAgentCommandExecutionOutputContentOutputTypeEnum Enum with underlying type: string
type InstanceAgentCommandExecutionOutputContentOutputTypeEnum string

// Set of constants representing the allowable values for InstanceAgentCommandExecutionOutputContentOutputTypeEnum
const (
	InstanceAgentCommandExecutionOutputContentOutputTypeText               InstanceAgentCommandExecutionOutputContentOutputTypeEnum = "TEXT"
	InstanceAgentCommandExecutionOutputContentOutputTypeObjectStorageUri   InstanceAgentCommandExecutionOutputContentOutputTypeEnum = "OBJECT_STORAGE_URI"
	InstanceAgentCommandExecutionOutputContentOutputTypeObjectStorageTuple InstanceAgentCommandExecutionOutputContentOutputTypeEnum = "OBJECT_STORAGE_TUPLE"
)

var mappingInstanceAgentCommandExecutionOutputContentOutputTypeEnum = map[string]InstanceAgentCommandExecutionOutputContentOutputTypeEnum{
	"TEXT":                 InstanceAgentCommandExecutionOutputContentOutputTypeText,
	"OBJECT_STORAGE_URI":   InstanceAgentCommandExecutionOutputContentOutputTypeObjectStorageUri,
	"OBJECT_STORAGE_TUPLE": InstanceAgentCommandExecutionOutputContentOutputTypeObjectStorageTuple,
}

var mappingInstanceAgentCommandExecutionOutputContentOutputTypeEnumLowerCase = map[string]InstanceAgentCommandExecutionOutputContentOutputTypeEnum{
	"text":                 InstanceAgentCommandExecutionOutputContentOutputTypeText,
	"object_storage_uri":   InstanceAgentCommandExecutionOutputContentOutputTypeObjectStorageUri,
	"object_storage_tuple": InstanceAgentCommandExecutionOutputContentOutputTypeObjectStorageTuple,
}

// GetInstanceAgentCommandExecutionOutputContentOutputTypeEnumValues Enumerates the set of values for InstanceAgentCommandExecutionOutputContentOutputTypeEnum
func GetInstanceAgentCommandExecutionOutputContentOutputTypeEnumValues() []InstanceAgentCommandExecutionOutputContentOutputTypeEnum {
	values := make([]InstanceAgentCommandExecutionOutputContentOutputTypeEnum, 0)
	for _, v := range mappingInstanceAgentCommandExecutionOutputContentOutputTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetInstanceAgentCommandExecutionOutputContentOutputTypeEnumStringValues Enumerates the set of values in String for InstanceAgentCommandExecutionOutputContentOutputTypeEnum
func GetInstanceAgentCommandExecutionOutputContentOutputTypeEnumStringValues() []string {
	return []string{
		"TEXT",
		"OBJECT_STORAGE_URI",
		"OBJECT_STORAGE_TUPLE",
	}
}

// GetMappingInstanceAgentCommandExecutionOutputContentOutputTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingInstanceAgentCommandExecutionOutputContentOutputTypeEnum(val string) (InstanceAgentCommandExecutionOutputContentOutputTypeEnum, bool) {
	enum, ok := mappingInstanceAgentCommandExecutionOutputContentOutputTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
