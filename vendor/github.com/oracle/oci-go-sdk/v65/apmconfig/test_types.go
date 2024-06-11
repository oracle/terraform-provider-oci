// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Application Performance Monitoring Configuration API
//
// Use the Application Performance Monitoring Configuration API to query and set Application Performance Monitoring
// configuration. For more information, see Application Performance Monitoring (https://docs.oracle.com/iaas/application-performance-monitoring/index.html).
//

package apmconfig

import (
	"strings"
)

// TestTypesEnum Enum with underlying type: string
type TestTypesEnum string

// Set of constants representing the allowable values for TestTypesEnum
const (
	TestTypesSpanEnrichment TestTypesEnum = "SPAN_ENRICHMENT"
)

var mappingTestTypesEnum = map[string]TestTypesEnum{
	"SPAN_ENRICHMENT": TestTypesSpanEnrichment,
}

var mappingTestTypesEnumLowerCase = map[string]TestTypesEnum{
	"span_enrichment": TestTypesSpanEnrichment,
}

// GetTestTypesEnumValues Enumerates the set of values for TestTypesEnum
func GetTestTypesEnumValues() []TestTypesEnum {
	values := make([]TestTypesEnum, 0)
	for _, v := range mappingTestTypesEnum {
		values = append(values, v)
	}
	return values
}

// GetTestTypesEnumStringValues Enumerates the set of values in String for TestTypesEnum
func GetTestTypesEnumStringValues() []string {
	return []string{
		"SPAN_ENRICHMENT",
	}
}

// GetMappingTestTypesEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingTestTypesEnum(val string) (TestTypesEnum, bool) {
	enum, ok := mappingTestTypesEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
