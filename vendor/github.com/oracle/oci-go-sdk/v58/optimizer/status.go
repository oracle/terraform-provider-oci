// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Cloud Advisor API
//
// Use the Cloud Advisor API to find potential inefficiencies in your tenancy and address them.
// Cloud Advisor can help you save money, improve performance, strengthen system resilience, and improve security.
// For more information, see Cloud Advisor (https://docs.cloud.oracle.com/Content/CloudAdvisor/Concepts/cloudadvisoroverview.htm).
//

package optimizer

import (
	"strings"
)

// StatusEnum Enum with underlying type: string
type StatusEnum string

// Set of constants representing the allowable values for StatusEnum
const (
	StatusPending     StatusEnum = "PENDING"
	StatusDismissed   StatusEnum = "DISMISSED"
	StatusPostponed   StatusEnum = "POSTPONED"
	StatusImplemented StatusEnum = "IMPLEMENTED"
)

var mappingStatusEnum = map[string]StatusEnum{
	"PENDING":     StatusPending,
	"DISMISSED":   StatusDismissed,
	"POSTPONED":   StatusPostponed,
	"IMPLEMENTED": StatusImplemented,
}

// GetStatusEnumValues Enumerates the set of values for StatusEnum
func GetStatusEnumValues() []StatusEnum {
	values := make([]StatusEnum, 0)
	for _, v := range mappingStatusEnum {
		values = append(values, v)
	}
	return values
}

// GetStatusEnumStringValues Enumerates the set of values in String for StatusEnum
func GetStatusEnumStringValues() []string {
	return []string{
		"PENDING",
		"DISMISSED",
		"POSTPONED",
		"IMPLEMENTED",
	}
}

// GetMappingStatusEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingStatusEnum(val string) (StatusEnum, bool) {
	mappingStatusEnumIgnoreCase := make(map[string]StatusEnum)
	for k, v := range mappingStatusEnum {
		mappingStatusEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingStatusEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}
