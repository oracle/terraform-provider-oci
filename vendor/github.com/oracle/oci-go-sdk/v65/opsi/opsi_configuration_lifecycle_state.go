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

// OpsiConfigurationLifecycleStateEnum Enum with underlying type: string
type OpsiConfigurationLifecycleStateEnum string

// Set of constants representing the allowable values for OpsiConfigurationLifecycleStateEnum
const (
	OpsiConfigurationLifecycleStateCreating OpsiConfigurationLifecycleStateEnum = "CREATING"
	OpsiConfigurationLifecycleStateUpdating OpsiConfigurationLifecycleStateEnum = "UPDATING"
	OpsiConfigurationLifecycleStateActive   OpsiConfigurationLifecycleStateEnum = "ACTIVE"
	OpsiConfigurationLifecycleStateDeleting OpsiConfigurationLifecycleStateEnum = "DELETING"
	OpsiConfigurationLifecycleStateDeleted  OpsiConfigurationLifecycleStateEnum = "DELETED"
	OpsiConfigurationLifecycleStateFailed   OpsiConfigurationLifecycleStateEnum = "FAILED"
)

var mappingOpsiConfigurationLifecycleStateEnum = map[string]OpsiConfigurationLifecycleStateEnum{
	"CREATING": OpsiConfigurationLifecycleStateCreating,
	"UPDATING": OpsiConfigurationLifecycleStateUpdating,
	"ACTIVE":   OpsiConfigurationLifecycleStateActive,
	"DELETING": OpsiConfigurationLifecycleStateDeleting,
	"DELETED":  OpsiConfigurationLifecycleStateDeleted,
	"FAILED":   OpsiConfigurationLifecycleStateFailed,
}

var mappingOpsiConfigurationLifecycleStateEnumLowerCase = map[string]OpsiConfigurationLifecycleStateEnum{
	"creating": OpsiConfigurationLifecycleStateCreating,
	"updating": OpsiConfigurationLifecycleStateUpdating,
	"active":   OpsiConfigurationLifecycleStateActive,
	"deleting": OpsiConfigurationLifecycleStateDeleting,
	"deleted":  OpsiConfigurationLifecycleStateDeleted,
	"failed":   OpsiConfigurationLifecycleStateFailed,
}

// GetOpsiConfigurationLifecycleStateEnumValues Enumerates the set of values for OpsiConfigurationLifecycleStateEnum
func GetOpsiConfigurationLifecycleStateEnumValues() []OpsiConfigurationLifecycleStateEnum {
	values := make([]OpsiConfigurationLifecycleStateEnum, 0)
	for _, v := range mappingOpsiConfigurationLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetOpsiConfigurationLifecycleStateEnumStringValues Enumerates the set of values in String for OpsiConfigurationLifecycleStateEnum
func GetOpsiConfigurationLifecycleStateEnumStringValues() []string {
	return []string{
		"CREATING",
		"UPDATING",
		"ACTIVE",
		"DELETING",
		"DELETED",
		"FAILED",
	}
}

// GetMappingOpsiConfigurationLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingOpsiConfigurationLifecycleStateEnum(val string) (OpsiConfigurationLifecycleStateEnum, bool) {
	enum, ok := mappingOpsiConfigurationLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
