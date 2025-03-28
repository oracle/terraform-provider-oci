// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Java Management Service Fleets API
//
// The APIs for the Fleet Management (https://docs.oracle.com/en-us/iaas/jms/doc/fleet-management.html) feature of Java Management Service to monitor and manage the usage of Java in your enterprise. Use these APIs to manage fleets, configure managed instances to report to fleets, and gain insights into the Java workloads running on these instances by carrying out basic and advanced features.
//

package jms

import (
	"strings"
)

// GlobalLoggingLevelEnum Enum with underlying type: string
type GlobalLoggingLevelEnum string

// Set of constants representing the allowable values for GlobalLoggingLevelEnum
const (
	GlobalLoggingLevelAll     GlobalLoggingLevelEnum = "ALL"
	GlobalLoggingLevelSevere  GlobalLoggingLevelEnum = "SEVERE"
	GlobalLoggingLevelWarning GlobalLoggingLevelEnum = "WARNING"
	GlobalLoggingLevelInfo    GlobalLoggingLevelEnum = "INFO"
	GlobalLoggingLevelConfig  GlobalLoggingLevelEnum = "CONFIG"
	GlobalLoggingLevelFine    GlobalLoggingLevelEnum = "FINE"
	GlobalLoggingLevelFiner   GlobalLoggingLevelEnum = "FINER"
	GlobalLoggingLevelFinest  GlobalLoggingLevelEnum = "FINEST"
	GlobalLoggingLevelOff     GlobalLoggingLevelEnum = "OFF"
)

var mappingGlobalLoggingLevelEnum = map[string]GlobalLoggingLevelEnum{
	"ALL":     GlobalLoggingLevelAll,
	"SEVERE":  GlobalLoggingLevelSevere,
	"WARNING": GlobalLoggingLevelWarning,
	"INFO":    GlobalLoggingLevelInfo,
	"CONFIG":  GlobalLoggingLevelConfig,
	"FINE":    GlobalLoggingLevelFine,
	"FINER":   GlobalLoggingLevelFiner,
	"FINEST":  GlobalLoggingLevelFinest,
	"OFF":     GlobalLoggingLevelOff,
}

var mappingGlobalLoggingLevelEnumLowerCase = map[string]GlobalLoggingLevelEnum{
	"all":     GlobalLoggingLevelAll,
	"severe":  GlobalLoggingLevelSevere,
	"warning": GlobalLoggingLevelWarning,
	"info":    GlobalLoggingLevelInfo,
	"config":  GlobalLoggingLevelConfig,
	"fine":    GlobalLoggingLevelFine,
	"finer":   GlobalLoggingLevelFiner,
	"finest":  GlobalLoggingLevelFinest,
	"off":     GlobalLoggingLevelOff,
}

// GetGlobalLoggingLevelEnumValues Enumerates the set of values for GlobalLoggingLevelEnum
func GetGlobalLoggingLevelEnumValues() []GlobalLoggingLevelEnum {
	values := make([]GlobalLoggingLevelEnum, 0)
	for _, v := range mappingGlobalLoggingLevelEnum {
		values = append(values, v)
	}
	return values
}

// GetGlobalLoggingLevelEnumStringValues Enumerates the set of values in String for GlobalLoggingLevelEnum
func GetGlobalLoggingLevelEnumStringValues() []string {
	return []string{
		"ALL",
		"SEVERE",
		"WARNING",
		"INFO",
		"CONFIG",
		"FINE",
		"FINER",
		"FINEST",
		"OFF",
	}
}

// GetMappingGlobalLoggingLevelEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingGlobalLoggingLevelEnum(val string) (GlobalLoggingLevelEnum, bool) {
	enum, ok := mappingGlobalLoggingLevelEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
