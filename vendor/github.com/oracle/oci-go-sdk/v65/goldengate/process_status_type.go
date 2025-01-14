// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// GoldenGate API
//
// Use the Oracle Cloud Infrastructure GoldenGate APIs to perform data replication operations.
//

package goldengate

import (
	"strings"
)

// ProcessStatusTypeEnum Enum with underlying type: string
type ProcessStatusTypeEnum string

// Set of constants representing the allowable values for ProcessStatusTypeEnum
const (
	ProcessStatusTypeStopped ProcessStatusTypeEnum = "STOPPED"
	ProcessStatusTypeRunning ProcessStatusTypeEnum = "RUNNING"
	ProcessStatusTypeError   ProcessStatusTypeEnum = "ERROR"
)

var mappingProcessStatusTypeEnum = map[string]ProcessStatusTypeEnum{
	"STOPPED": ProcessStatusTypeStopped,
	"RUNNING": ProcessStatusTypeRunning,
	"ERROR":   ProcessStatusTypeError,
}

var mappingProcessStatusTypeEnumLowerCase = map[string]ProcessStatusTypeEnum{
	"stopped": ProcessStatusTypeStopped,
	"running": ProcessStatusTypeRunning,
	"error":   ProcessStatusTypeError,
}

// GetProcessStatusTypeEnumValues Enumerates the set of values for ProcessStatusTypeEnum
func GetProcessStatusTypeEnumValues() []ProcessStatusTypeEnum {
	values := make([]ProcessStatusTypeEnum, 0)
	for _, v := range mappingProcessStatusTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetProcessStatusTypeEnumStringValues Enumerates the set of values in String for ProcessStatusTypeEnum
func GetProcessStatusTypeEnumStringValues() []string {
	return []string{
		"STOPPED",
		"RUNNING",
		"ERROR",
	}
}

// GetMappingProcessStatusTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingProcessStatusTypeEnum(val string) (ProcessStatusTypeEnum, bool) {
	enum, ok := mappingProcessStatusTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
