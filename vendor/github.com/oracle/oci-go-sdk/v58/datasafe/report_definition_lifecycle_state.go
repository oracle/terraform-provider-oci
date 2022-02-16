// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
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

// ReportDefinitionLifecycleStateEnum Enum with underlying type: string
type ReportDefinitionLifecycleStateEnum string

// Set of constants representing the allowable values for ReportDefinitionLifecycleStateEnum
const (
	ReportDefinitionLifecycleStateCreating ReportDefinitionLifecycleStateEnum = "CREATING"
	ReportDefinitionLifecycleStateUpdating ReportDefinitionLifecycleStateEnum = "UPDATING"
	ReportDefinitionLifecycleStateActive   ReportDefinitionLifecycleStateEnum = "ACTIVE"
	ReportDefinitionLifecycleStateDeleting ReportDefinitionLifecycleStateEnum = "DELETING"
	ReportDefinitionLifecycleStateDeleted  ReportDefinitionLifecycleStateEnum = "DELETED"
)

var mappingReportDefinitionLifecycleStateEnum = map[string]ReportDefinitionLifecycleStateEnum{
	"CREATING": ReportDefinitionLifecycleStateCreating,
	"UPDATING": ReportDefinitionLifecycleStateUpdating,
	"ACTIVE":   ReportDefinitionLifecycleStateActive,
	"DELETING": ReportDefinitionLifecycleStateDeleting,
	"DELETED":  ReportDefinitionLifecycleStateDeleted,
}

// GetReportDefinitionLifecycleStateEnumValues Enumerates the set of values for ReportDefinitionLifecycleStateEnum
func GetReportDefinitionLifecycleStateEnumValues() []ReportDefinitionLifecycleStateEnum {
	values := make([]ReportDefinitionLifecycleStateEnum, 0)
	for _, v := range mappingReportDefinitionLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetReportDefinitionLifecycleStateEnumStringValues Enumerates the set of values in String for ReportDefinitionLifecycleStateEnum
func GetReportDefinitionLifecycleStateEnumStringValues() []string {
	return []string{
		"CREATING",
		"UPDATING",
		"ACTIVE",
		"DELETING",
		"DELETED",
	}
}

// GetMappingReportDefinitionLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingReportDefinitionLifecycleStateEnum(val string) (ReportDefinitionLifecycleStateEnum, bool) {
	mappingReportDefinitionLifecycleStateEnumIgnoreCase := make(map[string]ReportDefinitionLifecycleStateEnum)
	for k, v := range mappingReportDefinitionLifecycleStateEnum {
		mappingReportDefinitionLifecycleStateEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingReportDefinitionLifecycleStateEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}
