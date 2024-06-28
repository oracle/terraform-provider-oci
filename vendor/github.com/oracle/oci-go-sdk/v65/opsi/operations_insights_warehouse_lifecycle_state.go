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

// OperationsInsightsWarehouseLifecycleStateEnum Enum with underlying type: string
type OperationsInsightsWarehouseLifecycleStateEnum string

// Set of constants representing the allowable values for OperationsInsightsWarehouseLifecycleStateEnum
const (
	OperationsInsightsWarehouseLifecycleStateCreating OperationsInsightsWarehouseLifecycleStateEnum = "CREATING"
	OperationsInsightsWarehouseLifecycleStateUpdating OperationsInsightsWarehouseLifecycleStateEnum = "UPDATING"
	OperationsInsightsWarehouseLifecycleStateActive   OperationsInsightsWarehouseLifecycleStateEnum = "ACTIVE"
	OperationsInsightsWarehouseLifecycleStateDeleting OperationsInsightsWarehouseLifecycleStateEnum = "DELETING"
	OperationsInsightsWarehouseLifecycleStateDeleted  OperationsInsightsWarehouseLifecycleStateEnum = "DELETED"
	OperationsInsightsWarehouseLifecycleStateFailed   OperationsInsightsWarehouseLifecycleStateEnum = "FAILED"
)

var mappingOperationsInsightsWarehouseLifecycleStateEnum = map[string]OperationsInsightsWarehouseLifecycleStateEnum{
	"CREATING": OperationsInsightsWarehouseLifecycleStateCreating,
	"UPDATING": OperationsInsightsWarehouseLifecycleStateUpdating,
	"ACTIVE":   OperationsInsightsWarehouseLifecycleStateActive,
	"DELETING": OperationsInsightsWarehouseLifecycleStateDeleting,
	"DELETED":  OperationsInsightsWarehouseLifecycleStateDeleted,
	"FAILED":   OperationsInsightsWarehouseLifecycleStateFailed,
}

var mappingOperationsInsightsWarehouseLifecycleStateEnumLowerCase = map[string]OperationsInsightsWarehouseLifecycleStateEnum{
	"creating": OperationsInsightsWarehouseLifecycleStateCreating,
	"updating": OperationsInsightsWarehouseLifecycleStateUpdating,
	"active":   OperationsInsightsWarehouseLifecycleStateActive,
	"deleting": OperationsInsightsWarehouseLifecycleStateDeleting,
	"deleted":  OperationsInsightsWarehouseLifecycleStateDeleted,
	"failed":   OperationsInsightsWarehouseLifecycleStateFailed,
}

// GetOperationsInsightsWarehouseLifecycleStateEnumValues Enumerates the set of values for OperationsInsightsWarehouseLifecycleStateEnum
func GetOperationsInsightsWarehouseLifecycleStateEnumValues() []OperationsInsightsWarehouseLifecycleStateEnum {
	values := make([]OperationsInsightsWarehouseLifecycleStateEnum, 0)
	for _, v := range mappingOperationsInsightsWarehouseLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetOperationsInsightsWarehouseLifecycleStateEnumStringValues Enumerates the set of values in String for OperationsInsightsWarehouseLifecycleStateEnum
func GetOperationsInsightsWarehouseLifecycleStateEnumStringValues() []string {
	return []string{
		"CREATING",
		"UPDATING",
		"ACTIVE",
		"DELETING",
		"DELETED",
		"FAILED",
	}
}

// GetMappingOperationsInsightsWarehouseLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingOperationsInsightsWarehouseLifecycleStateEnum(val string) (OperationsInsightsWarehouseLifecycleStateEnum, bool) {
	enum, ok := mappingOperationsInsightsWarehouseLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
