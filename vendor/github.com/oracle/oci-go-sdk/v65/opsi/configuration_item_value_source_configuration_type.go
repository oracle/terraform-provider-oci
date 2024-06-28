// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Ops Insights API
//
// Use the Ops Insights API to perform data extraction operations to obtain database
// resource utilization, performance statistics, and reference information. For more information,
// see About Oracle Cloud Infrastructure Ops Insights (https://docs.cloud.oracle.com/en-us/iaas/operations-insights/doc/operations-insights.html).
//

package opsi

import (
	"strings"
)

// ConfigurationItemValueSourceConfigurationTypeEnum Enum with underlying type: string
type ConfigurationItemValueSourceConfigurationTypeEnum string

// Set of constants representing the allowable values for ConfigurationItemValueSourceConfigurationTypeEnum
const (
	ConfigurationItemValueSourceConfigurationTypeDefault     ConfigurationItemValueSourceConfigurationTypeEnum = "DEFAULT"
	ConfigurationItemValueSourceConfigurationTypeTenant      ConfigurationItemValueSourceConfigurationTypeEnum = "TENANT"
	ConfigurationItemValueSourceConfigurationTypeCompartment ConfigurationItemValueSourceConfigurationTypeEnum = "COMPARTMENT"
)

var mappingConfigurationItemValueSourceConfigurationTypeEnum = map[string]ConfigurationItemValueSourceConfigurationTypeEnum{
	"DEFAULT":     ConfigurationItemValueSourceConfigurationTypeDefault,
	"TENANT":      ConfigurationItemValueSourceConfigurationTypeTenant,
	"COMPARTMENT": ConfigurationItemValueSourceConfigurationTypeCompartment,
}

var mappingConfigurationItemValueSourceConfigurationTypeEnumLowerCase = map[string]ConfigurationItemValueSourceConfigurationTypeEnum{
	"default":     ConfigurationItemValueSourceConfigurationTypeDefault,
	"tenant":      ConfigurationItemValueSourceConfigurationTypeTenant,
	"compartment": ConfigurationItemValueSourceConfigurationTypeCompartment,
}

// GetConfigurationItemValueSourceConfigurationTypeEnumValues Enumerates the set of values for ConfigurationItemValueSourceConfigurationTypeEnum
func GetConfigurationItemValueSourceConfigurationTypeEnumValues() []ConfigurationItemValueSourceConfigurationTypeEnum {
	values := make([]ConfigurationItemValueSourceConfigurationTypeEnum, 0)
	for _, v := range mappingConfigurationItemValueSourceConfigurationTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetConfigurationItemValueSourceConfigurationTypeEnumStringValues Enumerates the set of values in String for ConfigurationItemValueSourceConfigurationTypeEnum
func GetConfigurationItemValueSourceConfigurationTypeEnumStringValues() []string {
	return []string{
		"DEFAULT",
		"TENANT",
		"COMPARTMENT",
	}
}

// GetMappingConfigurationItemValueSourceConfigurationTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingConfigurationItemValueSourceConfigurationTypeEnum(val string) (ConfigurationItemValueSourceConfigurationTypeEnum, bool) {
	enum, ok := mappingConfigurationItemValueSourceConfigurationTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
