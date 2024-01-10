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

// InstanceAgentCommandOutputDetails The output destination for the command.
type InstanceAgentCommandOutputDetails interface {
}

type instanceagentcommandoutputdetails struct {
	JsonData   []byte
	OutputType string `json:"outputType"`
}

// UnmarshalJSON unmarshals json
func (m *instanceagentcommandoutputdetails) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerinstanceagentcommandoutputdetails instanceagentcommandoutputdetails
	s := struct {
		Model Unmarshalerinstanceagentcommandoutputdetails
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.OutputType = s.Model.OutputType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *instanceagentcommandoutputdetails) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.OutputType {
	case "OBJECT_STORAGE_URI":
		mm := InstanceAgentCommandOutputViaObjectStorageUriDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "OBJECT_STORAGE_TUPLE":
		mm := InstanceAgentCommandOutputViaObjectStorageTupleDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "TEXT":
		mm := InstanceAgentCommandOutputViaTextDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Recieved unsupported enum value for InstanceAgentCommandOutputDetails: %s.", m.OutputType)
		return *m, nil
	}
}

func (m instanceagentcommandoutputdetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m instanceagentcommandoutputdetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// InstanceAgentCommandOutputDetailsOutputTypeEnum Enum with underlying type: string
type InstanceAgentCommandOutputDetailsOutputTypeEnum string

// Set of constants representing the allowable values for InstanceAgentCommandOutputDetailsOutputTypeEnum
const (
	InstanceAgentCommandOutputDetailsOutputTypeText               InstanceAgentCommandOutputDetailsOutputTypeEnum = "TEXT"
	InstanceAgentCommandOutputDetailsOutputTypeObjectStorageUri   InstanceAgentCommandOutputDetailsOutputTypeEnum = "OBJECT_STORAGE_URI"
	InstanceAgentCommandOutputDetailsOutputTypeObjectStorageTuple InstanceAgentCommandOutputDetailsOutputTypeEnum = "OBJECT_STORAGE_TUPLE"
)

var mappingInstanceAgentCommandOutputDetailsOutputTypeEnum = map[string]InstanceAgentCommandOutputDetailsOutputTypeEnum{
	"TEXT":                 InstanceAgentCommandOutputDetailsOutputTypeText,
	"OBJECT_STORAGE_URI":   InstanceAgentCommandOutputDetailsOutputTypeObjectStorageUri,
	"OBJECT_STORAGE_TUPLE": InstanceAgentCommandOutputDetailsOutputTypeObjectStorageTuple,
}

var mappingInstanceAgentCommandOutputDetailsOutputTypeEnumLowerCase = map[string]InstanceAgentCommandOutputDetailsOutputTypeEnum{
	"text":                 InstanceAgentCommandOutputDetailsOutputTypeText,
	"object_storage_uri":   InstanceAgentCommandOutputDetailsOutputTypeObjectStorageUri,
	"object_storage_tuple": InstanceAgentCommandOutputDetailsOutputTypeObjectStorageTuple,
}

// GetInstanceAgentCommandOutputDetailsOutputTypeEnumValues Enumerates the set of values for InstanceAgentCommandOutputDetailsOutputTypeEnum
func GetInstanceAgentCommandOutputDetailsOutputTypeEnumValues() []InstanceAgentCommandOutputDetailsOutputTypeEnum {
	values := make([]InstanceAgentCommandOutputDetailsOutputTypeEnum, 0)
	for _, v := range mappingInstanceAgentCommandOutputDetailsOutputTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetInstanceAgentCommandOutputDetailsOutputTypeEnumStringValues Enumerates the set of values in String for InstanceAgentCommandOutputDetailsOutputTypeEnum
func GetInstanceAgentCommandOutputDetailsOutputTypeEnumStringValues() []string {
	return []string{
		"TEXT",
		"OBJECT_STORAGE_URI",
		"OBJECT_STORAGE_TUPLE",
	}
}

// GetMappingInstanceAgentCommandOutputDetailsOutputTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingInstanceAgentCommandOutputDetailsOutputTypeEnum(val string) (InstanceAgentCommandOutputDetailsOutputTypeEnum, bool) {
	enum, ok := mappingInstanceAgentCommandOutputDetailsOutputTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
