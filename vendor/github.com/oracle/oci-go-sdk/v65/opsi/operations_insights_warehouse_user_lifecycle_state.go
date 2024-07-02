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

// OperationsInsightsWarehouseUserLifecycleStateEnum Enum with underlying type: string
type OperationsInsightsWarehouseUserLifecycleStateEnum string

// Set of constants representing the allowable values for OperationsInsightsWarehouseUserLifecycleStateEnum
const (
	OperationsInsightsWarehouseUserLifecycleStateCreating OperationsInsightsWarehouseUserLifecycleStateEnum = "CREATING"
	OperationsInsightsWarehouseUserLifecycleStateUpdating OperationsInsightsWarehouseUserLifecycleStateEnum = "UPDATING"
	OperationsInsightsWarehouseUserLifecycleStateActive   OperationsInsightsWarehouseUserLifecycleStateEnum = "ACTIVE"
	OperationsInsightsWarehouseUserLifecycleStateDeleting OperationsInsightsWarehouseUserLifecycleStateEnum = "DELETING"
	OperationsInsightsWarehouseUserLifecycleStateDeleted  OperationsInsightsWarehouseUserLifecycleStateEnum = "DELETED"
	OperationsInsightsWarehouseUserLifecycleStateFailed   OperationsInsightsWarehouseUserLifecycleStateEnum = "FAILED"
)

var mappingOperationsInsightsWarehouseUserLifecycleStateEnum = map[string]OperationsInsightsWarehouseUserLifecycleStateEnum{
	"CREATING": OperationsInsightsWarehouseUserLifecycleStateCreating,
	"UPDATING": OperationsInsightsWarehouseUserLifecycleStateUpdating,
	"ACTIVE":   OperationsInsightsWarehouseUserLifecycleStateActive,
	"DELETING": OperationsInsightsWarehouseUserLifecycleStateDeleting,
	"DELETED":  OperationsInsightsWarehouseUserLifecycleStateDeleted,
	"FAILED":   OperationsInsightsWarehouseUserLifecycleStateFailed,
}

var mappingOperationsInsightsWarehouseUserLifecycleStateEnumLowerCase = map[string]OperationsInsightsWarehouseUserLifecycleStateEnum{
	"creating": OperationsInsightsWarehouseUserLifecycleStateCreating,
	"updating": OperationsInsightsWarehouseUserLifecycleStateUpdating,
	"active":   OperationsInsightsWarehouseUserLifecycleStateActive,
	"deleting": OperationsInsightsWarehouseUserLifecycleStateDeleting,
	"deleted":  OperationsInsightsWarehouseUserLifecycleStateDeleted,
	"failed":   OperationsInsightsWarehouseUserLifecycleStateFailed,
}

// GetOperationsInsightsWarehouseUserLifecycleStateEnumValues Enumerates the set of values for OperationsInsightsWarehouseUserLifecycleStateEnum
func GetOperationsInsightsWarehouseUserLifecycleStateEnumValues() []OperationsInsightsWarehouseUserLifecycleStateEnum {
	values := make([]OperationsInsightsWarehouseUserLifecycleStateEnum, 0)
	for _, v := range mappingOperationsInsightsWarehouseUserLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetOperationsInsightsWarehouseUserLifecycleStateEnumStringValues Enumerates the set of values in String for OperationsInsightsWarehouseUserLifecycleStateEnum
func GetOperationsInsightsWarehouseUserLifecycleStateEnumStringValues() []string {
	return []string{
		"CREATING",
		"UPDATING",
		"ACTIVE",
		"DELETING",
		"DELETED",
		"FAILED",
	}
}

// GetMappingOperationsInsightsWarehouseUserLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingOperationsInsightsWarehouseUserLifecycleStateEnum(val string) (OperationsInsightsWarehouseUserLifecycleStateEnum, bool) {
	enum, ok := mappingOperationsInsightsWarehouseUserLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
