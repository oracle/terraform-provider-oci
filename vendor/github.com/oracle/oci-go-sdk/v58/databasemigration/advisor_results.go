// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Migration API
//
// Use the Oracle Cloud Infrastructure Database Migration APIs to perform database migration operations.
//

package databasemigration

import (
	"strings"
)

// AdvisorResultsEnum Enum with underlying type: string
type AdvisorResultsEnum string

// Set of constants representing the allowable values for AdvisorResultsEnum
const (
	AdvisorResultsFatal         AdvisorResultsEnum = "FATAL"
	AdvisorResultsBlocker       AdvisorResultsEnum = "BLOCKER"
	AdvisorResultsWarning       AdvisorResultsEnum = "WARNING"
	AdvisorResultsInformational AdvisorResultsEnum = "INFORMATIONAL"
	AdvisorResultsPass          AdvisorResultsEnum = "PASS"
)

var mappingAdvisorResultsEnum = map[string]AdvisorResultsEnum{
	"FATAL":         AdvisorResultsFatal,
	"BLOCKER":       AdvisorResultsBlocker,
	"WARNING":       AdvisorResultsWarning,
	"INFORMATIONAL": AdvisorResultsInformational,
	"PASS":          AdvisorResultsPass,
}

// GetAdvisorResultsEnumValues Enumerates the set of values for AdvisorResultsEnum
func GetAdvisorResultsEnumValues() []AdvisorResultsEnum {
	values := make([]AdvisorResultsEnum, 0)
	for _, v := range mappingAdvisorResultsEnum {
		values = append(values, v)
	}
	return values
}

// GetAdvisorResultsEnumStringValues Enumerates the set of values in String for AdvisorResultsEnum
func GetAdvisorResultsEnumStringValues() []string {
	return []string{
		"FATAL",
		"BLOCKER",
		"WARNING",
		"INFORMATIONAL",
		"PASS",
	}
}

// GetMappingAdvisorResultsEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingAdvisorResultsEnum(val string) (AdvisorResultsEnum, bool) {
	mappingAdvisorResultsEnumIgnoreCase := make(map[string]AdvisorResultsEnum)
	for k, v := range mappingAdvisorResultsEnum {
		mappingAdvisorResultsEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingAdvisorResultsEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}
