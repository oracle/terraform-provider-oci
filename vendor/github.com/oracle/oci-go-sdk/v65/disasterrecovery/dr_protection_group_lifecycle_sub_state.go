// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Full Stack Disaster Recovery API
//
// Use the Full Stack Disaster Recovery (DR) API to manage disaster recovery for business applications.
// Full Stack DR is an OCI disaster recovery orchestration and management service that provides comprehensive disaster
// recovery capabilities for all layers of an application stack, including infrastructure, middleware, database,
// and application.
//

package disasterrecovery

import (
	"strings"
)

// DrProtectionGroupLifecycleSubStateEnum Enum with underlying type: string
type DrProtectionGroupLifecycleSubStateEnum string

// Set of constants representing the allowable values for DrProtectionGroupLifecycleSubStateEnum
const (
	DrProtectionGroupLifecycleSubStateDrDrillInProgress DrProtectionGroupLifecycleSubStateEnum = "DR_DRILL_IN_PROGRESS"
)

var mappingDrProtectionGroupLifecycleSubStateEnum = map[string]DrProtectionGroupLifecycleSubStateEnum{
	"DR_DRILL_IN_PROGRESS": DrProtectionGroupLifecycleSubStateDrDrillInProgress,
}

var mappingDrProtectionGroupLifecycleSubStateEnumLowerCase = map[string]DrProtectionGroupLifecycleSubStateEnum{
	"dr_drill_in_progress": DrProtectionGroupLifecycleSubStateDrDrillInProgress,
}

// GetDrProtectionGroupLifecycleSubStateEnumValues Enumerates the set of values for DrProtectionGroupLifecycleSubStateEnum
func GetDrProtectionGroupLifecycleSubStateEnumValues() []DrProtectionGroupLifecycleSubStateEnum {
	values := make([]DrProtectionGroupLifecycleSubStateEnum, 0)
	for _, v := range mappingDrProtectionGroupLifecycleSubStateEnum {
		values = append(values, v)
	}
	return values
}

// GetDrProtectionGroupLifecycleSubStateEnumStringValues Enumerates the set of values in String for DrProtectionGroupLifecycleSubStateEnum
func GetDrProtectionGroupLifecycleSubStateEnumStringValues() []string {
	return []string{
		"DR_DRILL_IN_PROGRESS",
	}
}

// GetMappingDrProtectionGroupLifecycleSubStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDrProtectionGroupLifecycleSubStateEnum(val string) (DrProtectionGroupLifecycleSubStateEnum, bool) {
	enum, ok := mappingDrProtectionGroupLifecycleSubStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
