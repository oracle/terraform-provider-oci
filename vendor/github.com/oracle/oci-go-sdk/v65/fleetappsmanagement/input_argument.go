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

// InputArgument The details of the Input argument.
type InputArgument interface {

	// The name of the argument
	GetName() *string

	// The description of the argument.
	GetDescription() *string
}

type inputargument struct {
	JsonData    []byte
	Description *string `mandatory:"false" json:"description"`
	Name        *string `mandatory:"true" json:"name"`
	Type        string  `json:"type"`
}

// UnmarshalJSON unmarshals json
func (m *inputargument) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerinputargument inputargument
	s := struct {
		Model Unmarshalerinputargument
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.Name = s.Model.Name
	m.Description = s.Model.Description
	m.Type = s.Model.Type

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *inputargument) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.Type {
	case "OUTPUT_VARIABLE":
		mm := OutputVariableInputArgument{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "STRING":
		mm := StringInputArgument{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Recieved unsupported enum value for InputArgument: %s.", m.Type)
		return *m, nil
	}
}

// GetDescription returns Description
func (m inputargument) GetDescription() *string {
	return m.Description
}

// GetName returns Name
func (m inputargument) GetName() *string {
	return m.Name
}

func (m inputargument) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m inputargument) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// InputArgumentTypeEnum Enum with underlying type: string
type InputArgumentTypeEnum string

// Set of constants representing the allowable values for InputArgumentTypeEnum
const (
	InputArgumentTypeString         InputArgumentTypeEnum = "STRING"
	InputArgumentTypeOutputVariable InputArgumentTypeEnum = "OUTPUT_VARIABLE"
)

var mappingInputArgumentTypeEnum = map[string]InputArgumentTypeEnum{
	"STRING":          InputArgumentTypeString,
	"OUTPUT_VARIABLE": InputArgumentTypeOutputVariable,
}

var mappingInputArgumentTypeEnumLowerCase = map[string]InputArgumentTypeEnum{
	"string":          InputArgumentTypeString,
	"output_variable": InputArgumentTypeOutputVariable,
}

// GetInputArgumentTypeEnumValues Enumerates the set of values for InputArgumentTypeEnum
func GetInputArgumentTypeEnumValues() []InputArgumentTypeEnum {
	values := make([]InputArgumentTypeEnum, 0)
	for _, v := range mappingInputArgumentTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetInputArgumentTypeEnumStringValues Enumerates the set of values in String for InputArgumentTypeEnum
func GetInputArgumentTypeEnumStringValues() []string {
	return []string{
		"STRING",
		"OUTPUT_VARIABLE",
	}
}

// GetMappingInputArgumentTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingInputArgumentTypeEnum(val string) (InputArgumentTypeEnum, bool) {
	enum, ok := mappingInputArgumentTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
