// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Safe API
//
// APIs for using Oracle Data Safe.
//

package datasafe

import (
	"strings"
)

// SensitiveTypeGroupLifecycleStateEnum Enum with underlying type: string
type SensitiveTypeGroupLifecycleStateEnum string

// Set of constants representing the allowable values for SensitiveTypeGroupLifecycleStateEnum
const (
	SensitiveTypeGroupLifecycleStateCreating SensitiveTypeGroupLifecycleStateEnum = "CREATING"
	SensitiveTypeGroupLifecycleStateActive   SensitiveTypeGroupLifecycleStateEnum = "ACTIVE"
	SensitiveTypeGroupLifecycleStateUpdating SensitiveTypeGroupLifecycleStateEnum = "UPDATING"
	SensitiveTypeGroupLifecycleStateDeleting SensitiveTypeGroupLifecycleStateEnum = "DELETING"
	SensitiveTypeGroupLifecycleStateDeleted  SensitiveTypeGroupLifecycleStateEnum = "DELETED"
	SensitiveTypeGroupLifecycleStateFailed   SensitiveTypeGroupLifecycleStateEnum = "FAILED"
)

var mappingSensitiveTypeGroupLifecycleStateEnum = map[string]SensitiveTypeGroupLifecycleStateEnum{
	"CREATING": SensitiveTypeGroupLifecycleStateCreating,
	"ACTIVE":   SensitiveTypeGroupLifecycleStateActive,
	"UPDATING": SensitiveTypeGroupLifecycleStateUpdating,
	"DELETING": SensitiveTypeGroupLifecycleStateDeleting,
	"DELETED":  SensitiveTypeGroupLifecycleStateDeleted,
	"FAILED":   SensitiveTypeGroupLifecycleStateFailed,
}

var mappingSensitiveTypeGroupLifecycleStateEnumLowerCase = map[string]SensitiveTypeGroupLifecycleStateEnum{
	"creating": SensitiveTypeGroupLifecycleStateCreating,
	"active":   SensitiveTypeGroupLifecycleStateActive,
	"updating": SensitiveTypeGroupLifecycleStateUpdating,
	"deleting": SensitiveTypeGroupLifecycleStateDeleting,
	"deleted":  SensitiveTypeGroupLifecycleStateDeleted,
	"failed":   SensitiveTypeGroupLifecycleStateFailed,
}

// GetSensitiveTypeGroupLifecycleStateEnumValues Enumerates the set of values for SensitiveTypeGroupLifecycleStateEnum
func GetSensitiveTypeGroupLifecycleStateEnumValues() []SensitiveTypeGroupLifecycleStateEnum {
	values := make([]SensitiveTypeGroupLifecycleStateEnum, 0)
	for _, v := range mappingSensitiveTypeGroupLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetSensitiveTypeGroupLifecycleStateEnumStringValues Enumerates the set of values in String for SensitiveTypeGroupLifecycleStateEnum
func GetSensitiveTypeGroupLifecycleStateEnumStringValues() []string {
	return []string{
		"CREATING",
		"ACTIVE",
		"UPDATING",
		"DELETING",
		"DELETED",
		"FAILED",
	}
}

// GetMappingSensitiveTypeGroupLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingSensitiveTypeGroupLifecycleStateEnum(val string) (SensitiveTypeGroupLifecycleStateEnum, bool) {
	enum, ok := mappingSensitiveTypeGroupLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
