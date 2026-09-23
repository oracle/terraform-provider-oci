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

// RuntimeConfig FunctionsRuntime configuration for a function.
type RuntimeConfig interface {
}

type runtimeconfig struct {
	JsonData          []byte
	RuntimeConfigType string `json:"runtimeConfigType"`
}

// UnmarshalJSON unmarshals json
func (m *runtimeconfig) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerruntimeconfig runtimeconfig
	s := struct {
		Model Unmarshalerruntimeconfig
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.RuntimeConfigType = s.Model.RuntimeConfigType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *runtimeconfig) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.RuntimeConfigType {
	case "FUNCTION_UPDATE":
		mm := FunctionUpdateRuntimeConfig{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "MANUAL":
		mm := ManualRuntimeConfig{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Received unsupported enum value for RuntimeConfig: %s.", m.RuntimeConfigType)
		return *m, nil
	}
}

func (m runtimeconfig) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m runtimeconfig) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// RuntimeConfigRuntimeConfigTypeEnum Enum with underlying type: string
type RuntimeConfigRuntimeConfigTypeEnum string

// Set of constants representing the allowable values for RuntimeConfigRuntimeConfigTypeEnum
const (
	RuntimeConfigRuntimeConfigTypeFunctionUpdate RuntimeConfigRuntimeConfigTypeEnum = "FUNCTION_UPDATE"
	RuntimeConfigRuntimeConfigTypeManual         RuntimeConfigRuntimeConfigTypeEnum = "MANUAL"
)

var mappingRuntimeConfigRuntimeConfigTypeEnum = map[string]RuntimeConfigRuntimeConfigTypeEnum{
	"FUNCTION_UPDATE": RuntimeConfigRuntimeConfigTypeFunctionUpdate,
	"MANUAL":          RuntimeConfigRuntimeConfigTypeManual,
}

var mappingRuntimeConfigRuntimeConfigTypeEnumLowerCase = map[string]RuntimeConfigRuntimeConfigTypeEnum{
	"function_update": RuntimeConfigRuntimeConfigTypeFunctionUpdate,
	"manual":          RuntimeConfigRuntimeConfigTypeManual,
}

// GetRuntimeConfigRuntimeConfigTypeEnumValues Enumerates the set of values for RuntimeConfigRuntimeConfigTypeEnum
func GetRuntimeConfigRuntimeConfigTypeEnumValues() []RuntimeConfigRuntimeConfigTypeEnum {
	values := make([]RuntimeConfigRuntimeConfigTypeEnum, 0)
	for _, v := range mappingRuntimeConfigRuntimeConfigTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetRuntimeConfigRuntimeConfigTypeEnumStringValues Enumerates the set of values in String for RuntimeConfigRuntimeConfigTypeEnum
func GetRuntimeConfigRuntimeConfigTypeEnumStringValues() []string {
	return []string{
		"FUNCTION_UPDATE",
		"MANUAL",
	}
}

// GetMappingRuntimeConfigRuntimeConfigTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingRuntimeConfigRuntimeConfigTypeEnum(val string) (RuntimeConfigRuntimeConfigTypeEnum, bool) {
	enum, ok := mappingRuntimeConfigRuntimeConfigTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
