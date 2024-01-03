// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Science API
//
// Use the Data Science API to organize your data science work, access data and computing resources, and build, train, deploy and manage models and model deployments. For more information, see Data Science (https://docs.oracle.com/iaas/data-science/using/data-science.htm).
//

package datascience

import (
	"strings"
)

// NodeGroupLifecycleStateEnum Enum with underlying type: string
type NodeGroupLifecycleStateEnum string

// Set of constants representing the allowable values for NodeGroupLifecycleStateEnum
const (
	NodeGroupLifecycleStateCreating NodeGroupLifecycleStateEnum = "CREATING"
	NodeGroupLifecycleStateActive   NodeGroupLifecycleStateEnum = "ACTIVE"
	NodeGroupLifecycleStateDeleting NodeGroupLifecycleStateEnum = "DELETING"
	NodeGroupLifecycleStateDeleted  NodeGroupLifecycleStateEnum = "DELETED"
	NodeGroupLifecycleStateFailed   NodeGroupLifecycleStateEnum = "FAILED"
)

var mappingNodeGroupLifecycleStateEnum = map[string]NodeGroupLifecycleStateEnum{
	"CREATING": NodeGroupLifecycleStateCreating,
	"ACTIVE":   NodeGroupLifecycleStateActive,
	"DELETING": NodeGroupLifecycleStateDeleting,
	"DELETED":  NodeGroupLifecycleStateDeleted,
	"FAILED":   NodeGroupLifecycleStateFailed,
}

var mappingNodeGroupLifecycleStateEnumLowerCase = map[string]NodeGroupLifecycleStateEnum{
	"creating": NodeGroupLifecycleStateCreating,
	"active":   NodeGroupLifecycleStateActive,
	"deleting": NodeGroupLifecycleStateDeleting,
	"deleted":  NodeGroupLifecycleStateDeleted,
	"failed":   NodeGroupLifecycleStateFailed,
}

// GetNodeGroupLifecycleStateEnumValues Enumerates the set of values for NodeGroupLifecycleStateEnum
func GetNodeGroupLifecycleStateEnumValues() []NodeGroupLifecycleStateEnum {
	values := make([]NodeGroupLifecycleStateEnum, 0)
	for _, v := range mappingNodeGroupLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetNodeGroupLifecycleStateEnumStringValues Enumerates the set of values in String for NodeGroupLifecycleStateEnum
func GetNodeGroupLifecycleStateEnumStringValues() []string {
	return []string{
		"CREATING",
		"ACTIVE",
		"DELETING",
		"DELETED",
		"FAILED",
	}
}

// GetMappingNodeGroupLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingNodeGroupLifecycleStateEnum(val string) (NodeGroupLifecycleStateEnum, bool) {
	enum, ok := mappingNodeGroupLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
