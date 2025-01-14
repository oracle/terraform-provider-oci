// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
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

// DrPlanLifecycleSubStateEnum Enum with underlying type: string
type DrPlanLifecycleSubStateEnum string

// Set of constants representing the allowable values for DrPlanLifecycleSubStateEnum
const (
	DrPlanLifecycleSubStateNeedsRefresh      DrPlanLifecycleSubStateEnum = "NEEDS_REFRESH"
	DrPlanLifecycleSubStateNeedsVerification DrPlanLifecycleSubStateEnum = "NEEDS_VERIFICATION"
	DrPlanLifecycleSubStateRefreshing        DrPlanLifecycleSubStateEnum = "REFRESHING"
	DrPlanLifecycleSubStateVerifying         DrPlanLifecycleSubStateEnum = "VERIFYING"
)

var mappingDrPlanLifecycleSubStateEnum = map[string]DrPlanLifecycleSubStateEnum{
	"NEEDS_REFRESH":      DrPlanLifecycleSubStateNeedsRefresh,
	"NEEDS_VERIFICATION": DrPlanLifecycleSubStateNeedsVerification,
	"REFRESHING":         DrPlanLifecycleSubStateRefreshing,
	"VERIFYING":          DrPlanLifecycleSubStateVerifying,
}

var mappingDrPlanLifecycleSubStateEnumLowerCase = map[string]DrPlanLifecycleSubStateEnum{
	"needs_refresh":      DrPlanLifecycleSubStateNeedsRefresh,
	"needs_verification": DrPlanLifecycleSubStateNeedsVerification,
	"refreshing":         DrPlanLifecycleSubStateRefreshing,
	"verifying":          DrPlanLifecycleSubStateVerifying,
}

// GetDrPlanLifecycleSubStateEnumValues Enumerates the set of values for DrPlanLifecycleSubStateEnum
func GetDrPlanLifecycleSubStateEnumValues() []DrPlanLifecycleSubStateEnum {
	values := make([]DrPlanLifecycleSubStateEnum, 0)
	for _, v := range mappingDrPlanLifecycleSubStateEnum {
		values = append(values, v)
	}
	return values
}

// GetDrPlanLifecycleSubStateEnumStringValues Enumerates the set of values in String for DrPlanLifecycleSubStateEnum
func GetDrPlanLifecycleSubStateEnumStringValues() []string {
	return []string{
		"NEEDS_REFRESH",
		"NEEDS_VERIFICATION",
		"REFRESHING",
		"VERIFYING",
	}
}

// GetMappingDrPlanLifecycleSubStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDrPlanLifecycleSubStateEnum(val string) (DrPlanLifecycleSubStateEnum, bool) {
	enum, ok := mappingDrPlanLifecycleSubStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
