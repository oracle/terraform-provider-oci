// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// OS Management Hub API
//
// Use the OS Management Hub API to manage and monitor updates and patches for instances in OCI, your private data center, or 3rd-party clouds.
// For more information, see Overview of OS Management Hub (https://docs.oracle.com/iaas/osmh/doc/overview.htm).
//

package osmanagementhub

import (
	"strings"
)

// HealthStateEnum Enum with underlying type: string
type HealthStateEnum string

// Set of constants representing the allowable values for HealthStateEnum
const (
	HealthStateHealthy     HealthStateEnum = "HEALTHY"
	HealthStateUnhealthy   HealthStateEnum = "UNHEALTHY"
	HealthStateUnavailable HealthStateEnum = "UNAVAILABLE"
)

var mappingHealthStateEnum = map[string]HealthStateEnum{
	"HEALTHY":     HealthStateHealthy,
	"UNHEALTHY":   HealthStateUnhealthy,
	"UNAVAILABLE": HealthStateUnavailable,
}

var mappingHealthStateEnumLowerCase = map[string]HealthStateEnum{
	"healthy":     HealthStateHealthy,
	"unhealthy":   HealthStateUnhealthy,
	"unavailable": HealthStateUnavailable,
}

// GetHealthStateEnumValues Enumerates the set of values for HealthStateEnum
func GetHealthStateEnumValues() []HealthStateEnum {
	values := make([]HealthStateEnum, 0)
	for _, v := range mappingHealthStateEnum {
		values = append(values, v)
	}
	return values
}

// GetHealthStateEnumStringValues Enumerates the set of values in String for HealthStateEnum
func GetHealthStateEnumStringValues() []string {
	return []string{
		"HEALTHY",
		"UNHEALTHY",
		"UNAVAILABLE",
	}
}

// GetMappingHealthStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingHealthStateEnum(val string) (HealthStateEnum, bool) {
	enum, ok := mappingHealthStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
