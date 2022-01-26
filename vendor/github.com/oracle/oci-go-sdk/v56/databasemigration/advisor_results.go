// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Migration API
//
// Use the Oracle Cloud Infrastructure Database Migration APIs to perform database migration operations.
//

package databasemigration

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

var mappingAdvisorResults = map[string]AdvisorResultsEnum{
	"FATAL":         AdvisorResultsFatal,
	"BLOCKER":       AdvisorResultsBlocker,
	"WARNING":       AdvisorResultsWarning,
	"INFORMATIONAL": AdvisorResultsInformational,
	"PASS":          AdvisorResultsPass,
}

// GetAdvisorResultsEnumValues Enumerates the set of values for AdvisorResultsEnum
func GetAdvisorResultsEnumValues() []AdvisorResultsEnum {
	values := make([]AdvisorResultsEnum, 0)
	for _, v := range mappingAdvisorResults {
		values = append(values, v)
	}
	return values
}
