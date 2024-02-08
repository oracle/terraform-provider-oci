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

// OperationsInsightsPrivateEndpointLifecycleStateEnum Enum with underlying type: string
type OperationsInsightsPrivateEndpointLifecycleStateEnum string

// Set of constants representing the allowable values for OperationsInsightsPrivateEndpointLifecycleStateEnum
const (
	OperationsInsightsPrivateEndpointLifecycleStateCreating       OperationsInsightsPrivateEndpointLifecycleStateEnum = "CREATING"
	OperationsInsightsPrivateEndpointLifecycleStateUpdating       OperationsInsightsPrivateEndpointLifecycleStateEnum = "UPDATING"
	OperationsInsightsPrivateEndpointLifecycleStateActive         OperationsInsightsPrivateEndpointLifecycleStateEnum = "ACTIVE"
	OperationsInsightsPrivateEndpointLifecycleStateDeleting       OperationsInsightsPrivateEndpointLifecycleStateEnum = "DELETING"
	OperationsInsightsPrivateEndpointLifecycleStateDeleted        OperationsInsightsPrivateEndpointLifecycleStateEnum = "DELETED"
	OperationsInsightsPrivateEndpointLifecycleStateFailed         OperationsInsightsPrivateEndpointLifecycleStateEnum = "FAILED"
	OperationsInsightsPrivateEndpointLifecycleStateNeedsAttention OperationsInsightsPrivateEndpointLifecycleStateEnum = "NEEDS_ATTENTION"
)

var mappingOperationsInsightsPrivateEndpointLifecycleStateEnum = map[string]OperationsInsightsPrivateEndpointLifecycleStateEnum{
	"CREATING":        OperationsInsightsPrivateEndpointLifecycleStateCreating,
	"UPDATING":        OperationsInsightsPrivateEndpointLifecycleStateUpdating,
	"ACTIVE":          OperationsInsightsPrivateEndpointLifecycleStateActive,
	"DELETING":        OperationsInsightsPrivateEndpointLifecycleStateDeleting,
	"DELETED":         OperationsInsightsPrivateEndpointLifecycleStateDeleted,
	"FAILED":          OperationsInsightsPrivateEndpointLifecycleStateFailed,
	"NEEDS_ATTENTION": OperationsInsightsPrivateEndpointLifecycleStateNeedsAttention,
}

var mappingOperationsInsightsPrivateEndpointLifecycleStateEnumLowerCase = map[string]OperationsInsightsPrivateEndpointLifecycleStateEnum{
	"creating":        OperationsInsightsPrivateEndpointLifecycleStateCreating,
	"updating":        OperationsInsightsPrivateEndpointLifecycleStateUpdating,
	"active":          OperationsInsightsPrivateEndpointLifecycleStateActive,
	"deleting":        OperationsInsightsPrivateEndpointLifecycleStateDeleting,
	"deleted":         OperationsInsightsPrivateEndpointLifecycleStateDeleted,
	"failed":          OperationsInsightsPrivateEndpointLifecycleStateFailed,
	"needs_attention": OperationsInsightsPrivateEndpointLifecycleStateNeedsAttention,
}

// GetOperationsInsightsPrivateEndpointLifecycleStateEnumValues Enumerates the set of values for OperationsInsightsPrivateEndpointLifecycleStateEnum
func GetOperationsInsightsPrivateEndpointLifecycleStateEnumValues() []OperationsInsightsPrivateEndpointLifecycleStateEnum {
	values := make([]OperationsInsightsPrivateEndpointLifecycleStateEnum, 0)
	for _, v := range mappingOperationsInsightsPrivateEndpointLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetOperationsInsightsPrivateEndpointLifecycleStateEnumStringValues Enumerates the set of values in String for OperationsInsightsPrivateEndpointLifecycleStateEnum
func GetOperationsInsightsPrivateEndpointLifecycleStateEnumStringValues() []string {
	return []string{
		"CREATING",
		"UPDATING",
		"ACTIVE",
		"DELETING",
		"DELETED",
		"FAILED",
		"NEEDS_ATTENTION",
	}
}

// GetMappingOperationsInsightsPrivateEndpointLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingOperationsInsightsPrivateEndpointLifecycleStateEnum(val string) (OperationsInsightsPrivateEndpointLifecycleStateEnum, bool) {
	enum, ok := mappingOperationsInsightsPrivateEndpointLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
