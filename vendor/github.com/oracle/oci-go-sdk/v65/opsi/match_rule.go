// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Ops Insights API
//
// Use the Ops Insights API to perform data extraction operations to obtain database
// resource utilization, performance statistics, and reference information. For more information,
// see About Oracle Cloud Infrastructure Ops Insights (https://docs.oracle.com/iaas/en-us/iaas/operations-insights/doc/operations-insights.html).
//

package opsi

import (
	"strings"
)

// MatchRuleEnum Enum with underlying type: string
type MatchRuleEnum string

// Set of constants representing the allowable values for MatchRuleEnum
const (
	MatchRuleMatchAny  MatchRuleEnum = "MATCH_ANY"
	MatchRuleMatchAll  MatchRuleEnum = "MATCH_ALL"
	MatchRuleMatchNone MatchRuleEnum = "MATCH_NONE"
)

var mappingMatchRuleEnum = map[string]MatchRuleEnum{
	"MATCH_ANY":  MatchRuleMatchAny,
	"MATCH_ALL":  MatchRuleMatchAll,
	"MATCH_NONE": MatchRuleMatchNone,
}

var mappingMatchRuleEnumLowerCase = map[string]MatchRuleEnum{
	"match_any":  MatchRuleMatchAny,
	"match_all":  MatchRuleMatchAll,
	"match_none": MatchRuleMatchNone,
}

// GetMatchRuleEnumValues Enumerates the set of values for MatchRuleEnum
func GetMatchRuleEnumValues() []MatchRuleEnum {
	values := make([]MatchRuleEnum, 0)
	for _, v := range mappingMatchRuleEnum {
		values = append(values, v)
	}
	return values
}

// GetMatchRuleEnumStringValues Enumerates the set of values in String for MatchRuleEnum
func GetMatchRuleEnumStringValues() []string {
	return []string{
		"MATCH_ANY",
		"MATCH_ALL",
		"MATCH_NONE",
	}
}

// GetMappingMatchRuleEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingMatchRuleEnum(val string) (MatchRuleEnum, bool) {
	enum, ok := mappingMatchRuleEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
