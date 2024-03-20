// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Java Management Service Download API
//
// The APIs for the <a href="https://docs.oracle.com/en-us/iaas/jms/doc/java-download.html">Java Download</a> feature of Java Management Service.
//

package jmsjavadownloads

import (
	"strings"
)

// AggregationSortByEnum Enum with underlying type: string
type AggregationSortByEnum string

// Set of constants representing the allowable values for AggregationSortByEnum
const (
	AggregationSortByFamilyVersion AggregationSortByEnum = "FAMILY_VERSION"
	AggregationSortByDownloadCount AggregationSortByEnum = "DOWNLOAD_COUNT"
)

var mappingAggregationSortByEnum = map[string]AggregationSortByEnum{
	"FAMILY_VERSION": AggregationSortByFamilyVersion,
	"DOWNLOAD_COUNT": AggregationSortByDownloadCount,
}

var mappingAggregationSortByEnumLowerCase = map[string]AggregationSortByEnum{
	"family_version": AggregationSortByFamilyVersion,
	"download_count": AggregationSortByDownloadCount,
}

// GetAggregationSortByEnumValues Enumerates the set of values for AggregationSortByEnum
func GetAggregationSortByEnumValues() []AggregationSortByEnum {
	values := make([]AggregationSortByEnum, 0)
	for _, v := range mappingAggregationSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetAggregationSortByEnumStringValues Enumerates the set of values in String for AggregationSortByEnum
func GetAggregationSortByEnumStringValues() []string {
	return []string{
		"FAMILY_VERSION",
		"DOWNLOAD_COUNT",
	}
}

// GetMappingAggregationSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingAggregationSortByEnum(val string) (AggregationSortByEnum, bool) {
	enum, ok := mappingAggregationSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
