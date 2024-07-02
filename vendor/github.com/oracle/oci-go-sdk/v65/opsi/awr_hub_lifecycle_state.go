// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Ops Insights API
//
// Use the Ops Insights API to perform data extraction operations to obtain database
// resource utilization, performance statistics, and reference information. For more information,
// see About Oracle Cloud Infrastructure Ops Insights (https://docs.cloud.oracle.com/en-us/iaas/operations-insights/doc/operations-insights.html).
//

package opsi

import (
	"strings"
)

// AwrHubLifecycleStateEnum Enum with underlying type: string
type AwrHubLifecycleStateEnum string

// Set of constants representing the allowable values for AwrHubLifecycleStateEnum
const (
	AwrHubLifecycleStateCreating AwrHubLifecycleStateEnum = "CREATING"
	AwrHubLifecycleStateUpdating AwrHubLifecycleStateEnum = "UPDATING"
	AwrHubLifecycleStateActive   AwrHubLifecycleStateEnum = "ACTIVE"
	AwrHubLifecycleStateDeleting AwrHubLifecycleStateEnum = "DELETING"
	AwrHubLifecycleStateDeleted  AwrHubLifecycleStateEnum = "DELETED"
	AwrHubLifecycleStateFailed   AwrHubLifecycleStateEnum = "FAILED"
)

var mappingAwrHubLifecycleStateEnum = map[string]AwrHubLifecycleStateEnum{
	"CREATING": AwrHubLifecycleStateCreating,
	"UPDATING": AwrHubLifecycleStateUpdating,
	"ACTIVE":   AwrHubLifecycleStateActive,
	"DELETING": AwrHubLifecycleStateDeleting,
	"DELETED":  AwrHubLifecycleStateDeleted,
	"FAILED":   AwrHubLifecycleStateFailed,
}

var mappingAwrHubLifecycleStateEnumLowerCase = map[string]AwrHubLifecycleStateEnum{
	"creating": AwrHubLifecycleStateCreating,
	"updating": AwrHubLifecycleStateUpdating,
	"active":   AwrHubLifecycleStateActive,
	"deleting": AwrHubLifecycleStateDeleting,
	"deleted":  AwrHubLifecycleStateDeleted,
	"failed":   AwrHubLifecycleStateFailed,
}

// GetAwrHubLifecycleStateEnumValues Enumerates the set of values for AwrHubLifecycleStateEnum
func GetAwrHubLifecycleStateEnumValues() []AwrHubLifecycleStateEnum {
	values := make([]AwrHubLifecycleStateEnum, 0)
	for _, v := range mappingAwrHubLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetAwrHubLifecycleStateEnumStringValues Enumerates the set of values in String for AwrHubLifecycleStateEnum
func GetAwrHubLifecycleStateEnumStringValues() []string {
	return []string{
		"CREATING",
		"UPDATING",
		"ACTIVE",
		"DELETING",
		"DELETED",
		"FAILED",
	}
}

// GetMappingAwrHubLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingAwrHubLifecycleStateEnum(val string) (AwrHubLifecycleStateEnum, bool) {
	enum, ok := mappingAwrHubLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
