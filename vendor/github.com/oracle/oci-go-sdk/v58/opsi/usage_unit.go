// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Operations Insights API
//
// Use the Operations Insights API to perform data extraction operations to obtain database
// resource utilization, performance statistics, and reference information. For more information,
// see About Oracle Cloud Infrastructure Operations Insights (https://docs.cloud.oracle.com/en-us/iaas/operations-insights/doc/operations-insights.html).
//

package opsi

import (
	"strings"
)

// UsageUnitEnum Enum with underlying type: string
type UsageUnitEnum string

// Set of constants representing the allowable values for UsageUnitEnum
const (
	UsageUnitCores   UsageUnitEnum = "CORES"
	UsageUnitGb      UsageUnitEnum = "GB"
	UsageUnitMbps    UsageUnitEnum = "MBPS"
	UsageUnitIops    UsageUnitEnum = "IOPS"
	UsageUnitPercent UsageUnitEnum = "PERCENT"
)

var mappingUsageUnitEnum = map[string]UsageUnitEnum{
	"CORES":   UsageUnitCores,
	"GB":      UsageUnitGb,
	"MBPS":    UsageUnitMbps,
	"IOPS":    UsageUnitIops,
	"PERCENT": UsageUnitPercent,
}

// GetUsageUnitEnumValues Enumerates the set of values for UsageUnitEnum
func GetUsageUnitEnumValues() []UsageUnitEnum {
	values := make([]UsageUnitEnum, 0)
	for _, v := range mappingUsageUnitEnum {
		values = append(values, v)
	}
	return values
}

// GetUsageUnitEnumStringValues Enumerates the set of values in String for UsageUnitEnum
func GetUsageUnitEnumStringValues() []string {
	return []string{
		"CORES",
		"GB",
		"MBPS",
		"IOPS",
		"PERCENT",
	}
}

// GetMappingUsageUnitEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingUsageUnitEnum(val string) (UsageUnitEnum, bool) {
	mappingUsageUnitEnumIgnoreCase := make(map[string]UsageUnitEnum)
	for k, v := range mappingUsageUnitEnum {
		mappingUsageUnitEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingUsageUnitEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}
