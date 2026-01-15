// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
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

// SensitiveTypesExportLifecycleStateEnum Enum with underlying type: string
type SensitiveTypesExportLifecycleStateEnum string

// Set of constants representing the allowable values for SensitiveTypesExportLifecycleStateEnum
const (
	SensitiveTypesExportLifecycleStateCreating SensitiveTypesExportLifecycleStateEnum = "CREATING"
	SensitiveTypesExportLifecycleStateActive   SensitiveTypesExportLifecycleStateEnum = "ACTIVE"
	SensitiveTypesExportLifecycleStateUpdating SensitiveTypesExportLifecycleStateEnum = "UPDATING"
	SensitiveTypesExportLifecycleStateDeleting SensitiveTypesExportLifecycleStateEnum = "DELETING"
	SensitiveTypesExportLifecycleStateDeleted  SensitiveTypesExportLifecycleStateEnum = "DELETED"
	SensitiveTypesExportLifecycleStateFailed   SensitiveTypesExportLifecycleStateEnum = "FAILED"
)

var mappingSensitiveTypesExportLifecycleStateEnum = map[string]SensitiveTypesExportLifecycleStateEnum{
	"CREATING": SensitiveTypesExportLifecycleStateCreating,
	"ACTIVE":   SensitiveTypesExportLifecycleStateActive,
	"UPDATING": SensitiveTypesExportLifecycleStateUpdating,
	"DELETING": SensitiveTypesExportLifecycleStateDeleting,
	"DELETED":  SensitiveTypesExportLifecycleStateDeleted,
	"FAILED":   SensitiveTypesExportLifecycleStateFailed,
}

var mappingSensitiveTypesExportLifecycleStateEnumLowerCase = map[string]SensitiveTypesExportLifecycleStateEnum{
	"creating": SensitiveTypesExportLifecycleStateCreating,
	"active":   SensitiveTypesExportLifecycleStateActive,
	"updating": SensitiveTypesExportLifecycleStateUpdating,
	"deleting": SensitiveTypesExportLifecycleStateDeleting,
	"deleted":  SensitiveTypesExportLifecycleStateDeleted,
	"failed":   SensitiveTypesExportLifecycleStateFailed,
}

// GetSensitiveTypesExportLifecycleStateEnumValues Enumerates the set of values for SensitiveTypesExportLifecycleStateEnum
func GetSensitiveTypesExportLifecycleStateEnumValues() []SensitiveTypesExportLifecycleStateEnum {
	values := make([]SensitiveTypesExportLifecycleStateEnum, 0)
	for _, v := range mappingSensitiveTypesExportLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetSensitiveTypesExportLifecycleStateEnumStringValues Enumerates the set of values in String for SensitiveTypesExportLifecycleStateEnum
func GetSensitiveTypesExportLifecycleStateEnumStringValues() []string {
	return []string{
		"CREATING",
		"ACTIVE",
		"UPDATING",
		"DELETING",
		"DELETED",
		"FAILED",
	}
}

// GetMappingSensitiveTypesExportLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingSensitiveTypesExportLifecycleStateEnum(val string) (SensitiveTypesExportLifecycleStateEnum, bool) {
	enum, ok := mappingSensitiveTypesExportLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
