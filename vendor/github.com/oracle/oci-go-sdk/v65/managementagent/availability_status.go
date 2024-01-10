// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Management Agent API
//
// Use the Management Agent API to manage your infrastructure's management agents, including their plugins and install keys.
// For more information, see Management Agent (https://docs.cloud.oracle.com/iaas/management-agents/index.html).
//

package managementagent

import (
	"strings"
)

// AvailabilityStatusEnum Enum with underlying type: string
type AvailabilityStatusEnum string

// Set of constants representing the allowable values for AvailabilityStatusEnum
const (
	AvailabilityStatusActive       AvailabilityStatusEnum = "ACTIVE"
	AvailabilityStatusSilent       AvailabilityStatusEnum = "SILENT"
	AvailabilityStatusNotAvailable AvailabilityStatusEnum = "NOT_AVAILABLE"
)

var mappingAvailabilityStatusEnum = map[string]AvailabilityStatusEnum{
	"ACTIVE":        AvailabilityStatusActive,
	"SILENT":        AvailabilityStatusSilent,
	"NOT_AVAILABLE": AvailabilityStatusNotAvailable,
}

var mappingAvailabilityStatusEnumLowerCase = map[string]AvailabilityStatusEnum{
	"active":        AvailabilityStatusActive,
	"silent":        AvailabilityStatusSilent,
	"not_available": AvailabilityStatusNotAvailable,
}

// GetAvailabilityStatusEnumValues Enumerates the set of values for AvailabilityStatusEnum
func GetAvailabilityStatusEnumValues() []AvailabilityStatusEnum {
	values := make([]AvailabilityStatusEnum, 0)
	for _, v := range mappingAvailabilityStatusEnum {
		values = append(values, v)
	}
	return values
}

// GetAvailabilityStatusEnumStringValues Enumerates the set of values in String for AvailabilityStatusEnum
func GetAvailabilityStatusEnumStringValues() []string {
	return []string{
		"ACTIVE",
		"SILENT",
		"NOT_AVAILABLE",
	}
}

// GetMappingAvailabilityStatusEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingAvailabilityStatusEnum(val string) (AvailabilityStatusEnum, bool) {
	enum, ok := mappingAvailabilityStatusEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
