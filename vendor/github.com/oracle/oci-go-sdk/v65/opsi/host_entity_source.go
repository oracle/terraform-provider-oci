// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
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

// HostEntitySourceEnum Enum with underlying type: string
type HostEntitySourceEnum string

// Set of constants representing the allowable values for HostEntitySourceEnum
const (
	HostEntitySourceMacsManagedExternalHost HostEntitySourceEnum = "MACS_MANAGED_EXTERNAL_HOST"
	HostEntitySourceEmManagedExternalHost   HostEntitySourceEnum = "EM_MANAGED_EXTERNAL_HOST"
	HostEntitySourceMacsManagedCloudHost    HostEntitySourceEnum = "MACS_MANAGED_CLOUD_HOST"
	HostEntitySourcePeComanagedHost         HostEntitySourceEnum = "PE_COMANAGED_HOST"
)

var mappingHostEntitySourceEnum = map[string]HostEntitySourceEnum{
	"MACS_MANAGED_EXTERNAL_HOST": HostEntitySourceMacsManagedExternalHost,
	"EM_MANAGED_EXTERNAL_HOST":   HostEntitySourceEmManagedExternalHost,
	"MACS_MANAGED_CLOUD_HOST":    HostEntitySourceMacsManagedCloudHost,
	"PE_COMANAGED_HOST":          HostEntitySourcePeComanagedHost,
}

var mappingHostEntitySourceEnumLowerCase = map[string]HostEntitySourceEnum{
	"macs_managed_external_host": HostEntitySourceMacsManagedExternalHost,
	"em_managed_external_host":   HostEntitySourceEmManagedExternalHost,
	"macs_managed_cloud_host":    HostEntitySourceMacsManagedCloudHost,
	"pe_comanaged_host":          HostEntitySourcePeComanagedHost,
}

// GetHostEntitySourceEnumValues Enumerates the set of values for HostEntitySourceEnum
func GetHostEntitySourceEnumValues() []HostEntitySourceEnum {
	values := make([]HostEntitySourceEnum, 0)
	for _, v := range mappingHostEntitySourceEnum {
		values = append(values, v)
	}
	return values
}

// GetHostEntitySourceEnumStringValues Enumerates the set of values in String for HostEntitySourceEnum
func GetHostEntitySourceEnumStringValues() []string {
	return []string{
		"MACS_MANAGED_EXTERNAL_HOST",
		"EM_MANAGED_EXTERNAL_HOST",
		"MACS_MANAGED_CLOUD_HOST",
		"PE_COMANAGED_HOST",
	}
}

// GetMappingHostEntitySourceEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingHostEntitySourceEnum(val string) (HostEntitySourceEnum, bool) {
	enum, ok := mappingHostEntitySourceEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
