// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Safe API
//
// APIs for using Oracle Data Safe.
//

package datasafe

import (
	"strings"
)

// ReportLifecycleStateEnum Enum with underlying type: string
type ReportLifecycleStateEnum string

// Set of constants representing the allowable values for ReportLifecycleStateEnum
const (
	ReportLifecycleStateUpdating ReportLifecycleStateEnum = "UPDATING"
	ReportLifecycleStateActive   ReportLifecycleStateEnum = "ACTIVE"
)

var mappingReportLifecycleStateEnum = map[string]ReportLifecycleStateEnum{
	"UPDATING": ReportLifecycleStateUpdating,
	"ACTIVE":   ReportLifecycleStateActive,
}

var mappingReportLifecycleStateEnumLowerCase = map[string]ReportLifecycleStateEnum{
	"updating": ReportLifecycleStateUpdating,
	"active":   ReportLifecycleStateActive,
}

// GetReportLifecycleStateEnumValues Enumerates the set of values for ReportLifecycleStateEnum
func GetReportLifecycleStateEnumValues() []ReportLifecycleStateEnum {
	values := make([]ReportLifecycleStateEnum, 0)
	for _, v := range mappingReportLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetReportLifecycleStateEnumStringValues Enumerates the set of values in String for ReportLifecycleStateEnum
func GetReportLifecycleStateEnumStringValues() []string {
	return []string{
		"UPDATING",
		"ACTIVE",
	}
}

// GetMappingReportLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingReportLifecycleStateEnum(val string) (ReportLifecycleStateEnum, bool) {
	enum, ok := mappingReportLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
