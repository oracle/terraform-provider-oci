// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Analytics API
//
// Analytics API.
//

package analytics

import (
	"strings"
)

// WorkRequestResourceTypeEnum Enum with underlying type: string
type WorkRequestResourceTypeEnum string

// Set of constants representing the allowable values for WorkRequestResourceTypeEnum
const (
	WorkRequestResourceTypeAnalyticsInstance WorkRequestResourceTypeEnum = "ANALYTICS_INSTANCE"
)

var mappingWorkRequestResourceTypeEnum = map[string]WorkRequestResourceTypeEnum{
	"ANALYTICS_INSTANCE": WorkRequestResourceTypeAnalyticsInstance,
}

var mappingWorkRequestResourceTypeEnumLowerCase = map[string]WorkRequestResourceTypeEnum{
	"analytics_instance": WorkRequestResourceTypeAnalyticsInstance,
}

// GetWorkRequestResourceTypeEnumValues Enumerates the set of values for WorkRequestResourceTypeEnum
func GetWorkRequestResourceTypeEnumValues() []WorkRequestResourceTypeEnum {
	values := make([]WorkRequestResourceTypeEnum, 0)
	for _, v := range mappingWorkRequestResourceTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetWorkRequestResourceTypeEnumStringValues Enumerates the set of values in String for WorkRequestResourceTypeEnum
func GetWorkRequestResourceTypeEnumStringValues() []string {
	return []string{
		"ANALYTICS_INSTANCE",
	}
}

// GetMappingWorkRequestResourceTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingWorkRequestResourceTypeEnum(val string) (WorkRequestResourceTypeEnum, bool) {
	enum, ok := mappingWorkRequestResourceTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
