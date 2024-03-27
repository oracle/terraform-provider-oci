// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Management API
//
// Use the Database Management API to perform tasks such as obtaining performance and resource usage metrics
// for a fleet of Managed Databases or a specific Managed Database, creating Managed Database Groups, and
// running a SQL job on a Managed Database or Managed Database Group.
//

package databasemanagement

import (
	"strings"
)

// HeatWaveClusterStatusEnum Enum with underlying type: string
type HeatWaveClusterStatusEnum string

// Set of constants representing the allowable values for HeatWaveClusterStatusEnum
const (
	HeatWaveClusterStatusUp      HeatWaveClusterStatusEnum = "UP"
	HeatWaveClusterStatusDown    HeatWaveClusterStatusEnum = "DOWN"
	HeatWaveClusterStatusUnknown HeatWaveClusterStatusEnum = "UNKNOWN"
)

var mappingHeatWaveClusterStatusEnum = map[string]HeatWaveClusterStatusEnum{
	"UP":      HeatWaveClusterStatusUp,
	"DOWN":    HeatWaveClusterStatusDown,
	"UNKNOWN": HeatWaveClusterStatusUnknown,
}

var mappingHeatWaveClusterStatusEnumLowerCase = map[string]HeatWaveClusterStatusEnum{
	"up":      HeatWaveClusterStatusUp,
	"down":    HeatWaveClusterStatusDown,
	"unknown": HeatWaveClusterStatusUnknown,
}

// GetHeatWaveClusterStatusEnumValues Enumerates the set of values for HeatWaveClusterStatusEnum
func GetHeatWaveClusterStatusEnumValues() []HeatWaveClusterStatusEnum {
	values := make([]HeatWaveClusterStatusEnum, 0)
	for _, v := range mappingHeatWaveClusterStatusEnum {
		values = append(values, v)
	}
	return values
}

// GetHeatWaveClusterStatusEnumStringValues Enumerates the set of values in String for HeatWaveClusterStatusEnum
func GetHeatWaveClusterStatusEnumStringValues() []string {
	return []string{
		"UP",
		"DOWN",
		"UNKNOWN",
	}
}

// GetMappingHeatWaveClusterStatusEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingHeatWaveClusterStatusEnum(val string) (HeatWaveClusterStatusEnum, bool) {
	enum, ok := mappingHeatWaveClusterStatusEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
