// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Ops Insights API
//
// Use the Ops Insights API to perform data extraction operations to obtain database
// resource utilization, performance statistics, and reference information. For more information,
// see About Oracle Cloud Infrastructure Ops Insights (https://docs.oracle.com/iaas/en-us/iaas/operations-insights/doc/operations-insights.html).
//

package opsi

import (
	"strings"
)

// ExadataInsightLifecycleStateEnum Enum with underlying type: string
type ExadataInsightLifecycleStateEnum string

// Set of constants representing the allowable values for ExadataInsightLifecycleStateEnum
const (
	ExadataInsightLifecycleStateCreating ExadataInsightLifecycleStateEnum = "CREATING"
	ExadataInsightLifecycleStateUpdating ExadataInsightLifecycleStateEnum = "UPDATING"
	ExadataInsightLifecycleStateActive   ExadataInsightLifecycleStateEnum = "ACTIVE"
	ExadataInsightLifecycleStateDeleting ExadataInsightLifecycleStateEnum = "DELETING"
	ExadataInsightLifecycleStateDeleted  ExadataInsightLifecycleStateEnum = "DELETED"
	ExadataInsightLifecycleStateFailed   ExadataInsightLifecycleStateEnum = "FAILED"
)

var mappingExadataInsightLifecycleStateEnum = map[string]ExadataInsightLifecycleStateEnum{
	"CREATING": ExadataInsightLifecycleStateCreating,
	"UPDATING": ExadataInsightLifecycleStateUpdating,
	"ACTIVE":   ExadataInsightLifecycleStateActive,
	"DELETING": ExadataInsightLifecycleStateDeleting,
	"DELETED":  ExadataInsightLifecycleStateDeleted,
	"FAILED":   ExadataInsightLifecycleStateFailed,
}

var mappingExadataInsightLifecycleStateEnumLowerCase = map[string]ExadataInsightLifecycleStateEnum{
	"creating": ExadataInsightLifecycleStateCreating,
	"updating": ExadataInsightLifecycleStateUpdating,
	"active":   ExadataInsightLifecycleStateActive,
	"deleting": ExadataInsightLifecycleStateDeleting,
	"deleted":  ExadataInsightLifecycleStateDeleted,
	"failed":   ExadataInsightLifecycleStateFailed,
}

// GetExadataInsightLifecycleStateEnumValues Enumerates the set of values for ExadataInsightLifecycleStateEnum
func GetExadataInsightLifecycleStateEnumValues() []ExadataInsightLifecycleStateEnum {
	values := make([]ExadataInsightLifecycleStateEnum, 0)
	for _, v := range mappingExadataInsightLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetExadataInsightLifecycleStateEnumStringValues Enumerates the set of values in String for ExadataInsightLifecycleStateEnum
func GetExadataInsightLifecycleStateEnumStringValues() []string {
	return []string{
		"CREATING",
		"UPDATING",
		"ACTIVE",
		"DELETING",
		"DELETED",
		"FAILED",
	}
}

// GetMappingExadataInsightLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingExadataInsightLifecycleStateEnum(val string) (ExadataInsightLifecycleStateEnum, bool) {
	enum, ok := mappingExadataInsightLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
