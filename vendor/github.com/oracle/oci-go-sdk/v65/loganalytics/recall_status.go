// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// LogAnalytics API
//
// The LogAnalytics API for the LogAnalytics service.
//

package loganalytics

import (
	"strings"
)

// RecallStatusEnum Enum with underlying type: string
type RecallStatusEnum string

// Set of constants representing the allowable values for RecallStatusEnum
const (
	RecallStatusRecalled RecallStatusEnum = "RECALLED"
	RecallStatusPending  RecallStatusEnum = "PENDING"
	RecallStatusFailed   RecallStatusEnum = "FAILED"
)

var mappingRecallStatusEnum = map[string]RecallStatusEnum{
	"RECALLED": RecallStatusRecalled,
	"PENDING":  RecallStatusPending,
	"FAILED":   RecallStatusFailed,
}

var mappingRecallStatusEnumLowerCase = map[string]RecallStatusEnum{
	"recalled": RecallStatusRecalled,
	"pending":  RecallStatusPending,
	"failed":   RecallStatusFailed,
}

// GetRecallStatusEnumValues Enumerates the set of values for RecallStatusEnum
func GetRecallStatusEnumValues() []RecallStatusEnum {
	values := make([]RecallStatusEnum, 0)
	for _, v := range mappingRecallStatusEnum {
		values = append(values, v)
	}
	return values
}

// GetRecallStatusEnumStringValues Enumerates the set of values in String for RecallStatusEnum
func GetRecallStatusEnumStringValues() []string {
	return []string{
		"RECALLED",
		"PENDING",
		"FAILED",
	}
}

// GetMappingRecallStatusEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingRecallStatusEnum(val string) (RecallStatusEnum, bool) {
	enum, ok := mappingRecallStatusEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
