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

// LifecycleStateEnum Enum with underlying type: string
type LifecycleStateEnum string

// Set of constants representing the allowable values for LifecycleStateEnum
const (
	LifecycleStateCreating       LifecycleStateEnum = "CREATING"
	LifecycleStateUpdating       LifecycleStateEnum = "UPDATING"
	LifecycleStateActive         LifecycleStateEnum = "ACTIVE"
	LifecycleStateDeleting       LifecycleStateEnum = "DELETING"
	LifecycleStateDeleted        LifecycleStateEnum = "DELETED"
	LifecycleStateFailed         LifecycleStateEnum = "FAILED"
	LifecycleStateNeedsAttention LifecycleStateEnum = "NEEDS_ATTENTION"
)

var mappingLifecycleStateEnum = map[string]LifecycleStateEnum{
	"CREATING":        LifecycleStateCreating,
	"UPDATING":        LifecycleStateUpdating,
	"ACTIVE":          LifecycleStateActive,
	"DELETING":        LifecycleStateDeleting,
	"DELETED":         LifecycleStateDeleted,
	"FAILED":          LifecycleStateFailed,
	"NEEDS_ATTENTION": LifecycleStateNeedsAttention,
}

var mappingLifecycleStateEnumLowerCase = map[string]LifecycleStateEnum{
	"creating":        LifecycleStateCreating,
	"updating":        LifecycleStateUpdating,
	"active":          LifecycleStateActive,
	"deleting":        LifecycleStateDeleting,
	"deleted":         LifecycleStateDeleted,
	"failed":          LifecycleStateFailed,
	"needs_attention": LifecycleStateNeedsAttention,
}

// GetLifecycleStateEnumValues Enumerates the set of values for LifecycleStateEnum
func GetLifecycleStateEnumValues() []LifecycleStateEnum {
	values := make([]LifecycleStateEnum, 0)
	for _, v := range mappingLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetLifecycleStateEnumStringValues Enumerates the set of values in String for LifecycleStateEnum
func GetLifecycleStateEnumStringValues() []string {
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

// GetMappingLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingLifecycleStateEnum(val string) (LifecycleStateEnum, bool) {
	enum, ok := mappingLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
