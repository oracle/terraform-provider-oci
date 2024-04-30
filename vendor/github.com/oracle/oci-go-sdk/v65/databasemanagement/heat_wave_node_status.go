// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Management API
//
// Use the Database Management API to monitor and manage resources such as
// Oracle Databases, MySQL Databases, and External Database Systems.
// For more information, see Database Management (https://docs.cloud.oracle.com/iaas/database-management/home.htm).
//

package databasemanagement

import (
	"strings"
)

// HeatWaveNodeStatusEnum Enum with underlying type: string
type HeatWaveNodeStatusEnum string

// Set of constants representing the allowable values for HeatWaveNodeStatusEnum
const (
	HeatWaveNodeStatusUp      HeatWaveNodeStatusEnum = "UP"
	HeatWaveNodeStatusDown    HeatWaveNodeStatusEnum = "DOWN"
	HeatWaveNodeStatusUnknown HeatWaveNodeStatusEnum = "UNKNOWN"
)

var mappingHeatWaveNodeStatusEnum = map[string]HeatWaveNodeStatusEnum{
	"UP":      HeatWaveNodeStatusUp,
	"DOWN":    HeatWaveNodeStatusDown,
	"UNKNOWN": HeatWaveNodeStatusUnknown,
}

var mappingHeatWaveNodeStatusEnumLowerCase = map[string]HeatWaveNodeStatusEnum{
	"up":      HeatWaveNodeStatusUp,
	"down":    HeatWaveNodeStatusDown,
	"unknown": HeatWaveNodeStatusUnknown,
}

// GetHeatWaveNodeStatusEnumValues Enumerates the set of values for HeatWaveNodeStatusEnum
func GetHeatWaveNodeStatusEnumValues() []HeatWaveNodeStatusEnum {
	values := make([]HeatWaveNodeStatusEnum, 0)
	for _, v := range mappingHeatWaveNodeStatusEnum {
		values = append(values, v)
	}
	return values
}

// GetHeatWaveNodeStatusEnumStringValues Enumerates the set of values in String for HeatWaveNodeStatusEnum
func GetHeatWaveNodeStatusEnumStringValues() []string {
	return []string{
		"UP",
		"DOWN",
		"UNKNOWN",
	}
}

// GetMappingHeatWaveNodeStatusEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingHeatWaveNodeStatusEnum(val string) (HeatWaveNodeStatusEnum, bool) {
	enum, ok := mappingHeatWaveNodeStatusEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
