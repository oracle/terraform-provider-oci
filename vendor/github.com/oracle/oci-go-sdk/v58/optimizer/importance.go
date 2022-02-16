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

// ImportanceEnum Enum with underlying type: string
type ImportanceEnum string

// Set of constants representing the allowable values for ImportanceEnum
const (
	ImportanceCritical ImportanceEnum = "CRITICAL"
	ImportanceHigh     ImportanceEnum = "HIGH"
	ImportanceModerate ImportanceEnum = "MODERATE"
	ImportanceLow      ImportanceEnum = "LOW"
	ImportanceMinor    ImportanceEnum = "MINOR"
)

var mappingImportanceEnum = map[string]ImportanceEnum{
	"CRITICAL": ImportanceCritical,
	"HIGH":     ImportanceHigh,
	"MODERATE": ImportanceModerate,
	"LOW":      ImportanceLow,
	"MINOR":    ImportanceMinor,
}

// GetImportanceEnumValues Enumerates the set of values for ImportanceEnum
func GetImportanceEnumValues() []ImportanceEnum {
	values := make([]ImportanceEnum, 0)
	for _, v := range mappingImportanceEnum {
		values = append(values, v)
	}
	return values
}

// GetImportanceEnumStringValues Enumerates the set of values in String for ImportanceEnum
func GetImportanceEnumStringValues() []string {
	return []string{
		"CRITICAL",
		"HIGH",
		"MODERATE",
		"LOW",
		"MINOR",
	}
}

// GetMappingImportanceEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingImportanceEnum(val string) (ImportanceEnum, bool) {
	mappingImportanceEnumIgnoreCase := make(map[string]ImportanceEnum)
	for k, v := range mappingImportanceEnum {
		mappingImportanceEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingImportanceEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}
