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

var mappingAwrHubLifecycleState = map[string]AwrHubLifecycleStateEnum{
	"CREATING": AwrHubLifecycleStateCreating,
	"UPDATING": AwrHubLifecycleStateUpdating,
	"ACTIVE":   AwrHubLifecycleStateActive,
	"DELETING": AwrHubLifecycleStateDeleting,
	"DELETED":  AwrHubLifecycleStateDeleted,
	"FAILED":   AwrHubLifecycleStateFailed,
}

// GetAwrHubLifecycleStateEnumValues Enumerates the set of values for AwrHubLifecycleStateEnum
func GetAwrHubLifecycleStateEnumValues() []AwrHubLifecycleStateEnum {
	values := make([]AwrHubLifecycleStateEnum, 0)
	for _, v := range mappingAwrHubLifecycleState {
		values = append(values, v)
	}
	return values
}
