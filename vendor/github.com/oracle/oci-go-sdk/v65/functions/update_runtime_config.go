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

// UpdateRuntimeConfig FunctionsRuntime configuration used to update a function.
type UpdateRuntimeConfig interface {
}

type updateruntimeconfig struct {
	JsonData          []byte
	RuntimeConfigType string `json:"runtimeConfigType"`
}

// UnmarshalJSON unmarshals json
func (m *updateruntimeconfig) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerupdateruntimeconfig updateruntimeconfig
	s := struct {
		Model Unmarshalerupdateruntimeconfig
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.RuntimeConfigType = s.Model.RuntimeConfigType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *updateruntimeconfig) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.RuntimeConfigType {
	case "MANUAL":
		mm := UpdateManualRuntimeConfig{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "FUNCTION_UPDATE":
		mm := UpdateFunctionUpdateRuntimeConfig{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Received unsupported enum value for UpdateRuntimeConfig: %s.", m.RuntimeConfigType)
		return *m, nil
	}
}

func (m updateruntimeconfig) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m updateruntimeconfig) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UpdateRuntimeConfigRuntimeConfigTypeEnum Enum with underlying type: string
type UpdateRuntimeConfigRuntimeConfigTypeEnum string

// Set of constants representing the allowable values for UpdateRuntimeConfigRuntimeConfigTypeEnum
const (
	UpdateRuntimeConfigRuntimeConfigTypeFunctionUpdate UpdateRuntimeConfigRuntimeConfigTypeEnum = "FUNCTION_UPDATE"
	UpdateRuntimeConfigRuntimeConfigTypeManual         UpdateRuntimeConfigRuntimeConfigTypeEnum = "MANUAL"
)

var mappingUpdateRuntimeConfigRuntimeConfigTypeEnum = map[string]UpdateRuntimeConfigRuntimeConfigTypeEnum{
	"FUNCTION_UPDATE": UpdateRuntimeConfigRuntimeConfigTypeFunctionUpdate,
	"MANUAL":          UpdateRuntimeConfigRuntimeConfigTypeManual,
}

var mappingUpdateRuntimeConfigRuntimeConfigTypeEnumLowerCase = map[string]UpdateRuntimeConfigRuntimeConfigTypeEnum{
	"function_update": UpdateRuntimeConfigRuntimeConfigTypeFunctionUpdate,
	"manual":          UpdateRuntimeConfigRuntimeConfigTypeManual,
}

// GetUpdateRuntimeConfigRuntimeConfigTypeEnumValues Enumerates the set of values for UpdateRuntimeConfigRuntimeConfigTypeEnum
func GetUpdateRuntimeConfigRuntimeConfigTypeEnumValues() []UpdateRuntimeConfigRuntimeConfigTypeEnum {
	values := make([]UpdateRuntimeConfigRuntimeConfigTypeEnum, 0)
	for _, v := range mappingUpdateRuntimeConfigRuntimeConfigTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetUpdateRuntimeConfigRuntimeConfigTypeEnumStringValues Enumerates the set of values in String for UpdateRuntimeConfigRuntimeConfigTypeEnum
func GetUpdateRuntimeConfigRuntimeConfigTypeEnumStringValues() []string {
	return []string{
		"FUNCTION_UPDATE",
		"MANUAL",
	}
}

// GetMappingUpdateRuntimeConfigRuntimeConfigTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingUpdateRuntimeConfigRuntimeConfigTypeEnum(val string) (UpdateRuntimeConfigRuntimeConfigTypeEnum, bool) {
	enum, ok := mappingUpdateRuntimeConfigRuntimeConfigTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
