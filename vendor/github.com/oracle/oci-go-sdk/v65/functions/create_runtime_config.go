// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Functions Service API
//
// API for the Functions service.
//

package functions

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// CreateRuntimeConfig FunctionsRuntime configuration used to create a function.
type CreateRuntimeConfig interface {
}

type createruntimeconfig struct {
	JsonData          []byte
	RuntimeConfigType string `json:"runtimeConfigType"`
}

// UnmarshalJSON unmarshals json
func (m *createruntimeconfig) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalercreateruntimeconfig createruntimeconfig
	s := struct {
		Model Unmarshalercreateruntimeconfig
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.RuntimeConfigType = s.Model.RuntimeConfigType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *createruntimeconfig) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.RuntimeConfigType {
	case "FUNCTION_UPDATE":
		mm := CreateFunctionUpdateRuntimeConfig{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "MANUAL":
		mm := CreateManualRuntimeConfig{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Received unsupported enum value for CreateRuntimeConfig: %s.", m.RuntimeConfigType)
		return *m, nil
	}
}

func (m createruntimeconfig) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m createruntimeconfig) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// CreateRuntimeConfigRuntimeConfigTypeEnum Enum with underlying type: string
type CreateRuntimeConfigRuntimeConfigTypeEnum string

// Set of constants representing the allowable values for CreateRuntimeConfigRuntimeConfigTypeEnum
const (
	CreateRuntimeConfigRuntimeConfigTypeFunctionUpdate CreateRuntimeConfigRuntimeConfigTypeEnum = "FUNCTION_UPDATE"
	CreateRuntimeConfigRuntimeConfigTypeManual         CreateRuntimeConfigRuntimeConfigTypeEnum = "MANUAL"
)

var mappingCreateRuntimeConfigRuntimeConfigTypeEnum = map[string]CreateRuntimeConfigRuntimeConfigTypeEnum{
	"FUNCTION_UPDATE": CreateRuntimeConfigRuntimeConfigTypeFunctionUpdate,
	"MANUAL":          CreateRuntimeConfigRuntimeConfigTypeManual,
}

var mappingCreateRuntimeConfigRuntimeConfigTypeEnumLowerCase = map[string]CreateRuntimeConfigRuntimeConfigTypeEnum{
	"function_update": CreateRuntimeConfigRuntimeConfigTypeFunctionUpdate,
	"manual":          CreateRuntimeConfigRuntimeConfigTypeManual,
}

// GetCreateRuntimeConfigRuntimeConfigTypeEnumValues Enumerates the set of values for CreateRuntimeConfigRuntimeConfigTypeEnum
func GetCreateRuntimeConfigRuntimeConfigTypeEnumValues() []CreateRuntimeConfigRuntimeConfigTypeEnum {
	values := make([]CreateRuntimeConfigRuntimeConfigTypeEnum, 0)
	for _, v := range mappingCreateRuntimeConfigRuntimeConfigTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetCreateRuntimeConfigRuntimeConfigTypeEnumStringValues Enumerates the set of values in String for CreateRuntimeConfigRuntimeConfigTypeEnum
func GetCreateRuntimeConfigRuntimeConfigTypeEnumStringValues() []string {
	return []string{
		"FUNCTION_UPDATE",
		"MANUAL",
	}
}

// GetMappingCreateRuntimeConfigRuntimeConfigTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingCreateRuntimeConfigRuntimeConfigTypeEnum(val string) (CreateRuntimeConfigRuntimeConfigTypeEnum, bool) {
	enum, ok := mappingCreateRuntimeConfigRuntimeConfigTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
