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

// RuleKindEnum Enum with underlying type: string
type RuleKindEnum string

// Set of constants representing the allowable values for RuleKindEnum
const (
	RuleKindIngestTime  RuleKindEnum = "INGEST_TIME"
	RuleKindSavedSearch RuleKindEnum = "SAVED_SEARCH"
)

var mappingRuleKindEnum = map[string]RuleKindEnum{
	"INGEST_TIME":  RuleKindIngestTime,
	"SAVED_SEARCH": RuleKindSavedSearch,
}

var mappingRuleKindEnumLowerCase = map[string]RuleKindEnum{
	"ingest_time":  RuleKindIngestTime,
	"saved_search": RuleKindSavedSearch,
}

// GetRuleKindEnumValues Enumerates the set of values for RuleKindEnum
func GetRuleKindEnumValues() []RuleKindEnum {
	values := make([]RuleKindEnum, 0)
	for _, v := range mappingRuleKindEnum {
		values = append(values, v)
	}
	return values
}

// GetRuleKindEnumStringValues Enumerates the set of values in String for RuleKindEnum
func GetRuleKindEnumStringValues() []string {
	return []string{
		"INGEST_TIME",
		"SAVED_SEARCH",
	}
}

// GetMappingRuleKindEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingRuleKindEnum(val string) (RuleKindEnum, bool) {
	enum, ok := mappingRuleKindEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
