// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Operations Insights API
//
// Use the Operations Insights API to perform data extraction operations to obtain database
// resource utilization, performance statistics, and reference information. For more information,
// see About Oracle Cloud Infrastructure Operations Insights (https://docs.cloud.oracle.com/en-us/iaas/operations-insights/doc/operations-insights.html).
//

package opsi

import (
	"strings"
)

// NewsFrequencyEnum Enum with underlying type: string
type NewsFrequencyEnum string

// Set of constants representing the allowable values for NewsFrequencyEnum
const (
	NewsFrequencyWeekly NewsFrequencyEnum = "WEEKLY"
)

var mappingNewsFrequencyEnum = map[string]NewsFrequencyEnum{
	"WEEKLY": NewsFrequencyWeekly,
}

var mappingNewsFrequencyEnumLowerCase = map[string]NewsFrequencyEnum{
	"weekly": NewsFrequencyWeekly,
}

// GetNewsFrequencyEnumValues Enumerates the set of values for NewsFrequencyEnum
func GetNewsFrequencyEnumValues() []NewsFrequencyEnum {
	values := make([]NewsFrequencyEnum, 0)
	for _, v := range mappingNewsFrequencyEnum {
		values = append(values, v)
	}
	return values
}

// GetNewsFrequencyEnumStringValues Enumerates the set of values in String for NewsFrequencyEnum
func GetNewsFrequencyEnumStringValues() []string {
	return []string{
		"WEEKLY",
	}
}

// GetMappingNewsFrequencyEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingNewsFrequencyEnum(val string) (NewsFrequencyEnum, bool) {
	enum, ok := mappingNewsFrequencyEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
