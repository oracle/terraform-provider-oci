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

// ConfigurationItemTypeEnum Enum with underlying type: string
type ConfigurationItemTypeEnum string

// Set of constants representing the allowable values for ConfigurationItemTypeEnum
const (
	ConfigurationItemTypeBasic ConfigurationItemTypeEnum = "BASIC"
)

var mappingConfigurationItemTypeEnum = map[string]ConfigurationItemTypeEnum{
	"BASIC": ConfigurationItemTypeBasic,
}

var mappingConfigurationItemTypeEnumLowerCase = map[string]ConfigurationItemTypeEnum{
	"basic": ConfigurationItemTypeBasic,
}

// GetConfigurationItemTypeEnumValues Enumerates the set of values for ConfigurationItemTypeEnum
func GetConfigurationItemTypeEnumValues() []ConfigurationItemTypeEnum {
	values := make([]ConfigurationItemTypeEnum, 0)
	for _, v := range mappingConfigurationItemTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetConfigurationItemTypeEnumStringValues Enumerates the set of values in String for ConfigurationItemTypeEnum
func GetConfigurationItemTypeEnumStringValues() []string {
	return []string{
		"BASIC",
	}
}

// GetMappingConfigurationItemTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingConfigurationItemTypeEnum(val string) (ConfigurationItemTypeEnum, bool) {
	enum, ok := mappingConfigurationItemTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
