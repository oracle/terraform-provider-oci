// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Operations Insights API
//
// Use the Operations Insights API to perform data extraction operations to obtain database
// resource utilization, performance statistics, and reference information. For more information,
// see About Oracle Cloud Infrastructure Operations Insights (https://docs.cloud.oracle.com/en-us/iaas/operations-insights/doc/operations-insights.html).
//

package opsi

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

var mappingOperationsInsightsWarehouseUserLifecycleState = map[string]OperationsInsightsWarehouseUserLifecycleStateEnum{
	"CREATING": OperationsInsightsWarehouseUserLifecycleStateCreating,
	"UPDATING": OperationsInsightsWarehouseUserLifecycleStateUpdating,
	"ACTIVE":   OperationsInsightsWarehouseUserLifecycleStateActive,
	"DELETING": OperationsInsightsWarehouseUserLifecycleStateDeleting,
	"DELETED":  OperationsInsightsWarehouseUserLifecycleStateDeleted,
	"FAILED":   OperationsInsightsWarehouseUserLifecycleStateFailed,
}

// GetOperationsInsightsWarehouseUserLifecycleStateEnumValues Enumerates the set of values for OperationsInsightsWarehouseUserLifecycleStateEnum
func GetOperationsInsightsWarehouseUserLifecycleStateEnumValues() []OperationsInsightsWarehouseUserLifecycleStateEnum {
	values := make([]OperationsInsightsWarehouseUserLifecycleStateEnum, 0)
	for _, v := range mappingOperationsInsightsWarehouseUserLifecycleState {
		values = append(values, v)
	}
	return values
}
