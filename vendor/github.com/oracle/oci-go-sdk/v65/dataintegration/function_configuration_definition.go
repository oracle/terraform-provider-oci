// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Integration API
//
// Use the Data Integration API to organize your data integration projects, create data flows, pipelines and tasks, and then publish, schedule, and run tasks that extract, transform, and load data. For more information, see Data Integration (https://docs.oracle.com/iaas/data-integration/home.htm).
//

package dataintegration

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// FunctionConfigurationDefinition The configuration details of a configurable object. This contains one or more config param definitions.
type FunctionConfigurationDefinition struct {

	// The key of the object.
	Key *string `mandatory:"false" json:"key"`

	// The type of the object.
	ModelType FunctionConfigurationDefinitionModelTypeEnum `mandatory:"false" json:"modelType,omitempty"`

	// The model version of an object.
	ModelVersion *string `mandatory:"false" json:"modelVersion"`

	ParentRef *ParentReference `mandatory:"false" json:"parentRef"`

	// Specifies whether the configuration is contained or not.
	IsContained *bool `mandatory:"false" json:"isContained"`

	// The parameter configuration details.
	ConfigParamDefs map[string]ConfigParameterDefinition `mandatory:"false" json:"configParamDefs"`
}

func (m FunctionConfigurationDefinition) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m FunctionConfigurationDefinition) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingFunctionConfigurationDefinitionModelTypeEnum(string(m.ModelType)); !ok && m.ModelType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ModelType: %s. Supported values are: %s.", m.ModelType, strings.Join(GetFunctionConfigurationDefinitionModelTypeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// FunctionConfigurationDefinitionModelTypeEnum Enum with underlying type: string
type FunctionConfigurationDefinitionModelTypeEnum string

// Set of constants representing the allowable values for FunctionConfigurationDefinitionModelTypeEnum
const (
	FunctionConfigurationDefinitionModelTypeConfigDefinition FunctionConfigurationDefinitionModelTypeEnum = "CONFIG_DEFINITION"
)

var mappingFunctionConfigurationDefinitionModelTypeEnum = map[string]FunctionConfigurationDefinitionModelTypeEnum{
	"CONFIG_DEFINITION": FunctionConfigurationDefinitionModelTypeConfigDefinition,
}

var mappingFunctionConfigurationDefinitionModelTypeEnumLowerCase = map[string]FunctionConfigurationDefinitionModelTypeEnum{
	"config_definition": FunctionConfigurationDefinitionModelTypeConfigDefinition,
}

// GetFunctionConfigurationDefinitionModelTypeEnumValues Enumerates the set of values for FunctionConfigurationDefinitionModelTypeEnum
func GetFunctionConfigurationDefinitionModelTypeEnumValues() []FunctionConfigurationDefinitionModelTypeEnum {
	values := make([]FunctionConfigurationDefinitionModelTypeEnum, 0)
	for _, v := range mappingFunctionConfigurationDefinitionModelTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetFunctionConfigurationDefinitionModelTypeEnumStringValues Enumerates the set of values in String for FunctionConfigurationDefinitionModelTypeEnum
func GetFunctionConfigurationDefinitionModelTypeEnumStringValues() []string {
	return []string{
		"CONFIG_DEFINITION",
	}
}

// GetMappingFunctionConfigurationDefinitionModelTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingFunctionConfigurationDefinitionModelTypeEnum(val string) (FunctionConfigurationDefinitionModelTypeEnum, bool) {
	enum, ok := mappingFunctionConfigurationDefinitionModelTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
