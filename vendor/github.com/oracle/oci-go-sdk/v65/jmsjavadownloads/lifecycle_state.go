// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Java Management Service Download API
//
// The APIs for the <a href="https://docs.oracle.com/en-us/iaas/jms/doc/java-download.html">Java Download</a> feature of Java Management Service.
//

package jmsjavadownloads

import (
	"strings"
)

// LifecycleStateEnum Enum with underlying type: string
type LifecycleStateEnum string

// Set of constants representing the allowable values for LifecycleStateEnum
const (
	LifecycleStateActive         LifecycleStateEnum = "ACTIVE"
	LifecycleStateCreating       LifecycleStateEnum = "CREATING"
	LifecycleStateDeleted        LifecycleStateEnum = "DELETED"
	LifecycleStateDeleting       LifecycleStateEnum = "DELETING"
	LifecycleStateFailed         LifecycleStateEnum = "FAILED"
	LifecycleStateNeedsAttention LifecycleStateEnum = "NEEDS_ATTENTION"
	LifecycleStateUpdating       LifecycleStateEnum = "UPDATING"
)

var mappingLifecycleStateEnum = map[string]LifecycleStateEnum{
	"ACTIVE":          LifecycleStateActive,
	"CREATING":        LifecycleStateCreating,
	"DELETED":         LifecycleStateDeleted,
	"DELETING":        LifecycleStateDeleting,
	"FAILED":          LifecycleStateFailed,
	"NEEDS_ATTENTION": LifecycleStateNeedsAttention,
	"UPDATING":        LifecycleStateUpdating,
}

var mappingLifecycleStateEnumLowerCase = map[string]LifecycleStateEnum{
	"active":          LifecycleStateActive,
	"creating":        LifecycleStateCreating,
	"deleted":         LifecycleStateDeleted,
	"deleting":        LifecycleStateDeleting,
	"failed":          LifecycleStateFailed,
	"needs_attention": LifecycleStateNeedsAttention,
	"updating":        LifecycleStateUpdating,
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
		"ACTIVE",
		"CREATING",
		"DELETED",
		"DELETING",
		"FAILED",
		"NEEDS_ATTENTION",
		"UPDATING",
	}
}

// GetMappingLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingLifecycleStateEnum(val string) (LifecycleStateEnum, bool) {
	enum, ok := mappingLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
