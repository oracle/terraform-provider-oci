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

// OpsiConfigurationTypeEnum Enum with underlying type: string
type OpsiConfigurationTypeEnum string

// Set of constants representing the allowable values for OpsiConfigurationTypeEnum
const (
	OpsiConfigurationTypeUxConfiguration OpsiConfigurationTypeEnum = "UX_CONFIGURATION"
)

var mappingOpsiConfigurationTypeEnum = map[string]OpsiConfigurationTypeEnum{
	"UX_CONFIGURATION": OpsiConfigurationTypeUxConfiguration,
}

var mappingOpsiConfigurationTypeEnumLowerCase = map[string]OpsiConfigurationTypeEnum{
	"ux_configuration": OpsiConfigurationTypeUxConfiguration,
}

// GetOpsiConfigurationTypeEnumValues Enumerates the set of values for OpsiConfigurationTypeEnum
func GetOpsiConfigurationTypeEnumValues() []OpsiConfigurationTypeEnum {
	values := make([]OpsiConfigurationTypeEnum, 0)
	for _, v := range mappingOpsiConfigurationTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetOpsiConfigurationTypeEnumStringValues Enumerates the set of values in String for OpsiConfigurationTypeEnum
func GetOpsiConfigurationTypeEnumStringValues() []string {
	return []string{
		"UX_CONFIGURATION",
	}
}

// GetMappingOpsiConfigurationTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingOpsiConfigurationTypeEnum(val string) (OpsiConfigurationTypeEnum, bool) {
	enum, ok := mappingOpsiConfigurationTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
