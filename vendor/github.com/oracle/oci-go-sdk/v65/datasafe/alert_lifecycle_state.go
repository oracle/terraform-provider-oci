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

// AlertLifecycleStateEnum Enum with underlying type: string
type AlertLifecycleStateEnum string

// Set of constants representing the allowable values for AlertLifecycleStateEnum
const (
	AlertLifecycleStateUpdating  AlertLifecycleStateEnum = "UPDATING"
	AlertLifecycleStateSucceeded AlertLifecycleStateEnum = "SUCCEEDED"
	AlertLifecycleStateFailed    AlertLifecycleStateEnum = "FAILED"
)

var mappingAlertLifecycleStateEnum = map[string]AlertLifecycleStateEnum{
	"UPDATING":  AlertLifecycleStateUpdating,
	"SUCCEEDED": AlertLifecycleStateSucceeded,
	"FAILED":    AlertLifecycleStateFailed,
}

var mappingAlertLifecycleStateEnumLowerCase = map[string]AlertLifecycleStateEnum{
	"updating":  AlertLifecycleStateUpdating,
	"succeeded": AlertLifecycleStateSucceeded,
	"failed":    AlertLifecycleStateFailed,
}

// GetAlertLifecycleStateEnumValues Enumerates the set of values for AlertLifecycleStateEnum
func GetAlertLifecycleStateEnumValues() []AlertLifecycleStateEnum {
	values := make([]AlertLifecycleStateEnum, 0)
	for _, v := range mappingAlertLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetAlertLifecycleStateEnumStringValues Enumerates the set of values in String for AlertLifecycleStateEnum
func GetAlertLifecycleStateEnumStringValues() []string {
	return []string{
		"UPDATING",
		"SUCCEEDED",
		"FAILED",
	}
}

// GetMappingAlertLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingAlertLifecycleStateEnum(val string) (AlertLifecycleStateEnum, bool) {
	enum, ok := mappingAlertLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
