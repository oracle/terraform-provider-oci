// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
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

// SensitiveColumnLifecycleStateEnum Enum with underlying type: string
type SensitiveColumnLifecycleStateEnum string

// Set of constants representing the allowable values for SensitiveColumnLifecycleStateEnum
const (
	SensitiveColumnLifecycleStateCreating SensitiveColumnLifecycleStateEnum = "CREATING"
	SensitiveColumnLifecycleStateActive   SensitiveColumnLifecycleStateEnum = "ACTIVE"
	SensitiveColumnLifecycleStateUpdating SensitiveColumnLifecycleStateEnum = "UPDATING"
	SensitiveColumnLifecycleStateDeleting SensitiveColumnLifecycleStateEnum = "DELETING"
	SensitiveColumnLifecycleStateFailed   SensitiveColumnLifecycleStateEnum = "FAILED"
)

var mappingSensitiveColumnLifecycleStateEnum = map[string]SensitiveColumnLifecycleStateEnum{
	"CREATING": SensitiveColumnLifecycleStateCreating,
	"ACTIVE":   SensitiveColumnLifecycleStateActive,
	"UPDATING": SensitiveColumnLifecycleStateUpdating,
	"DELETING": SensitiveColumnLifecycleStateDeleting,
	"FAILED":   SensitiveColumnLifecycleStateFailed,
}

var mappingSensitiveColumnLifecycleStateEnumLowerCase = map[string]SensitiveColumnLifecycleStateEnum{
	"creating": SensitiveColumnLifecycleStateCreating,
	"active":   SensitiveColumnLifecycleStateActive,
	"updating": SensitiveColumnLifecycleStateUpdating,
	"deleting": SensitiveColumnLifecycleStateDeleting,
	"failed":   SensitiveColumnLifecycleStateFailed,
}

// GetSensitiveColumnLifecycleStateEnumValues Enumerates the set of values for SensitiveColumnLifecycleStateEnum
func GetSensitiveColumnLifecycleStateEnumValues() []SensitiveColumnLifecycleStateEnum {
	values := make([]SensitiveColumnLifecycleStateEnum, 0)
	for _, v := range mappingSensitiveColumnLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetSensitiveColumnLifecycleStateEnumStringValues Enumerates the set of values in String for SensitiveColumnLifecycleStateEnum
func GetSensitiveColumnLifecycleStateEnumStringValues() []string {
	return []string{
		"CREATING",
		"ACTIVE",
		"UPDATING",
		"DELETING",
		"FAILED",
	}
}

// GetMappingSensitiveColumnLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingSensitiveColumnLifecycleStateEnum(val string) (SensitiveColumnLifecycleStateEnum, bool) {
	enum, ok := mappingSensitiveColumnLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
