// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Science API
//
// Use the Data Science API to organize your data science work, access data and computing resources, and build, train, deploy and manage models and model deployments. For more information, see Data Science (https://docs.oracle.com/iaas/data-science/using/data-science.htm).
//

package datascience

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// ConfigurationPropertySchema Schema for single configuration property
type ConfigurationPropertySchema struct {

	// Name of key (parameter name)
	KeyName *string `mandatory:"true" json:"keyName"`

	// Type of vaule
	ValueType ConfigurationPropertySchemaValueTypeEnum `mandatory:"true" json:"valueType"`

	// Description of this configuration property
	Description *string `mandatory:"false" json:"description"`
}

func (m ConfigurationPropertySchema) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ConfigurationPropertySchema) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingConfigurationPropertySchemaValueTypeEnum(string(m.ValueType)); !ok && m.ValueType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ValueType: %s. Supported values are: %s.", m.ValueType, strings.Join(GetConfigurationPropertySchemaValueTypeEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ConfigurationPropertySchemaValueTypeEnum Enum with underlying type: string
type ConfigurationPropertySchemaValueTypeEnum string

// Set of constants representing the allowable values for ConfigurationPropertySchemaValueTypeEnum
const (
	ConfigurationPropertySchemaValueTypeString ConfigurationPropertySchemaValueTypeEnum = "STRING"
	ConfigurationPropertySchemaValueTypeSecret ConfigurationPropertySchemaValueTypeEnum = "SECRET"
)

var mappingConfigurationPropertySchemaValueTypeEnum = map[string]ConfigurationPropertySchemaValueTypeEnum{
	"STRING": ConfigurationPropertySchemaValueTypeString,
	"SECRET": ConfigurationPropertySchemaValueTypeSecret,
}

var mappingConfigurationPropertySchemaValueTypeEnumLowerCase = map[string]ConfigurationPropertySchemaValueTypeEnum{
	"string": ConfigurationPropertySchemaValueTypeString,
	"secret": ConfigurationPropertySchemaValueTypeSecret,
}

// GetConfigurationPropertySchemaValueTypeEnumValues Enumerates the set of values for ConfigurationPropertySchemaValueTypeEnum
func GetConfigurationPropertySchemaValueTypeEnumValues() []ConfigurationPropertySchemaValueTypeEnum {
	values := make([]ConfigurationPropertySchemaValueTypeEnum, 0)
	for _, v := range mappingConfigurationPropertySchemaValueTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetConfigurationPropertySchemaValueTypeEnumStringValues Enumerates the set of values in String for ConfigurationPropertySchemaValueTypeEnum
func GetConfigurationPropertySchemaValueTypeEnumStringValues() []string {
	return []string{
		"STRING",
		"SECRET",
	}
}

// GetMappingConfigurationPropertySchemaValueTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingConfigurationPropertySchemaValueTypeEnum(val string) (ConfigurationPropertySchemaValueTypeEnum, bool) {
	enum, ok := mappingConfigurationPropertySchemaValueTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
