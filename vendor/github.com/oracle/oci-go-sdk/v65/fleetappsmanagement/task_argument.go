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

// TaskArgument A Task argument that holds a value.
type TaskArgument interface {

	// Name of the input variable
	GetName() *string
}

type taskargument struct {
	JsonData []byte
	Name     *string `mandatory:"true" json:"name"`
	Kind     string  `json:"kind"`
}

// UnmarshalJSON unmarshals json
func (m *taskargument) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalertaskargument taskargument
	s := struct {
		Model Unmarshalertaskargument
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.Name = s.Model.Name
	m.Kind = s.Model.Kind

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *taskargument) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.Kind {
	case "STRING":
		mm := StringTaskArgument{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "FILE":
		mm := FileTaskArgument{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Received unsupported enum value for TaskArgument: %s.", m.Kind)
		return *m, nil
	}
}

// GetName returns Name
func (m taskargument) GetName() *string {
	return m.Name
}

func (m taskargument) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m taskargument) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// TaskArgumentKindEnum Enum with underlying type: string
type TaskArgumentKindEnum string

// Set of constants representing the allowable values for TaskArgumentKindEnum
const (
	TaskArgumentKindString TaskArgumentKindEnum = "STRING"
	TaskArgumentKindFile   TaskArgumentKindEnum = "FILE"
)

var mappingTaskArgumentKindEnum = map[string]TaskArgumentKindEnum{
	"STRING": TaskArgumentKindString,
	"FILE":   TaskArgumentKindFile,
}

var mappingTaskArgumentKindEnumLowerCase = map[string]TaskArgumentKindEnum{
	"string": TaskArgumentKindString,
	"file":   TaskArgumentKindFile,
}

// GetTaskArgumentKindEnumValues Enumerates the set of values for TaskArgumentKindEnum
func GetTaskArgumentKindEnumValues() []TaskArgumentKindEnum {
	values := make([]TaskArgumentKindEnum, 0)
	for _, v := range mappingTaskArgumentKindEnum {
		values = append(values, v)
	}
	return values
}

// GetTaskArgumentKindEnumStringValues Enumerates the set of values in String for TaskArgumentKindEnum
func GetTaskArgumentKindEnumStringValues() []string {
	return []string{
		"STRING",
		"FILE",
	}
}

// GetMappingTaskArgumentKindEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingTaskArgumentKindEnum(val string) (TaskArgumentKindEnum, bool) {
	enum, ok := mappingTaskArgumentKindEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
