// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Oracle Cloud Bridge API
//
// API for Oracle Cloud Bridge service.
//

package cloudbridge

import (
	"strings"
)

// SupportedCloudRegionLifecycleStateEnum Enum with underlying type: string
type SupportedCloudRegionLifecycleStateEnum string

// Set of constants representing the allowable values for SupportedCloudRegionLifecycleStateEnum
const (
	SupportedCloudRegionLifecycleStateActive   SupportedCloudRegionLifecycleStateEnum = "ACTIVE"
	SupportedCloudRegionLifecycleStateInactive SupportedCloudRegionLifecycleStateEnum = "INACTIVE"
)

var mappingSupportedCloudRegionLifecycleStateEnum = map[string]SupportedCloudRegionLifecycleStateEnum{
	"ACTIVE":   SupportedCloudRegionLifecycleStateActive,
	"INACTIVE": SupportedCloudRegionLifecycleStateInactive,
}

var mappingSupportedCloudRegionLifecycleStateEnumLowerCase = map[string]SupportedCloudRegionLifecycleStateEnum{
	"active":   SupportedCloudRegionLifecycleStateActive,
	"inactive": SupportedCloudRegionLifecycleStateInactive,
}

// GetSupportedCloudRegionLifecycleStateEnumValues Enumerates the set of values for SupportedCloudRegionLifecycleStateEnum
func GetSupportedCloudRegionLifecycleStateEnumValues() []SupportedCloudRegionLifecycleStateEnum {
	values := make([]SupportedCloudRegionLifecycleStateEnum, 0)
	for _, v := range mappingSupportedCloudRegionLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetSupportedCloudRegionLifecycleStateEnumStringValues Enumerates the set of values in String for SupportedCloudRegionLifecycleStateEnum
func GetSupportedCloudRegionLifecycleStateEnumStringValues() []string {
	return []string{
		"ACTIVE",
		"INACTIVE",
	}
}

// GetMappingSupportedCloudRegionLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingSupportedCloudRegionLifecycleStateEnum(val string) (SupportedCloudRegionLifecycleStateEnum, bool) {
	enum, ok := mappingSupportedCloudRegionLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
