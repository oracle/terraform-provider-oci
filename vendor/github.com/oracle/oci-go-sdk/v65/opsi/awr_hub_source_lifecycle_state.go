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

// AwrHubSourceLifecycleStateEnum Enum with underlying type: string
type AwrHubSourceLifecycleStateEnum string

// Set of constants representing the allowable values for AwrHubSourceLifecycleStateEnum
const (
	AwrHubSourceLifecycleStateCreating AwrHubSourceLifecycleStateEnum = "CREATING"
	AwrHubSourceLifecycleStateUpdating AwrHubSourceLifecycleStateEnum = "UPDATING"
	AwrHubSourceLifecycleStateActive   AwrHubSourceLifecycleStateEnum = "ACTIVE"
	AwrHubSourceLifecycleStateDeleting AwrHubSourceLifecycleStateEnum = "DELETING"
	AwrHubSourceLifecycleStateDeleted  AwrHubSourceLifecycleStateEnum = "DELETED"
	AwrHubSourceLifecycleStateFailed   AwrHubSourceLifecycleStateEnum = "FAILED"
)

var mappingAwrHubSourceLifecycleStateEnum = map[string]AwrHubSourceLifecycleStateEnum{
	"CREATING": AwrHubSourceLifecycleStateCreating,
	"UPDATING": AwrHubSourceLifecycleStateUpdating,
	"ACTIVE":   AwrHubSourceLifecycleStateActive,
	"DELETING": AwrHubSourceLifecycleStateDeleting,
	"DELETED":  AwrHubSourceLifecycleStateDeleted,
	"FAILED":   AwrHubSourceLifecycleStateFailed,
}

var mappingAwrHubSourceLifecycleStateEnumLowerCase = map[string]AwrHubSourceLifecycleStateEnum{
	"creating": AwrHubSourceLifecycleStateCreating,
	"updating": AwrHubSourceLifecycleStateUpdating,
	"active":   AwrHubSourceLifecycleStateActive,
	"deleting": AwrHubSourceLifecycleStateDeleting,
	"deleted":  AwrHubSourceLifecycleStateDeleted,
	"failed":   AwrHubSourceLifecycleStateFailed,
}

// GetAwrHubSourceLifecycleStateEnumValues Enumerates the set of values for AwrHubSourceLifecycleStateEnum
func GetAwrHubSourceLifecycleStateEnumValues() []AwrHubSourceLifecycleStateEnum {
	values := make([]AwrHubSourceLifecycleStateEnum, 0)
	for _, v := range mappingAwrHubSourceLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetAwrHubSourceLifecycleStateEnumStringValues Enumerates the set of values in String for AwrHubSourceLifecycleStateEnum
func GetAwrHubSourceLifecycleStateEnumStringValues() []string {
	return []string{
		"CREATING",
		"UPDATING",
		"ACTIVE",
		"DELETING",
		"DELETED",
		"FAILED",
	}
}

// GetMappingAwrHubSourceLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingAwrHubSourceLifecycleStateEnum(val string) (AwrHubSourceLifecycleStateEnum, bool) {
	enum, ok := mappingAwrHubSourceLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
