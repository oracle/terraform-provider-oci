// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
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

// ExadataTypeEnum Enum with underlying type: string
type ExadataTypeEnum string

// Set of constants representing the allowable values for ExadataTypeEnum
const (
	ExadataTypeDbmachine ExadataTypeEnum = "DBMACHINE"
	ExadataTypeExacs     ExadataTypeEnum = "EXACS"
	ExadataTypeExacc     ExadataTypeEnum = "EXACC"
)

var mappingExadataTypeEnum = map[string]ExadataTypeEnum{
	"DBMACHINE": ExadataTypeDbmachine,
	"EXACS":     ExadataTypeExacs,
	"EXACC":     ExadataTypeExacc,
}

var mappingExadataTypeEnumLowerCase = map[string]ExadataTypeEnum{
	"dbmachine": ExadataTypeDbmachine,
	"exacs":     ExadataTypeExacs,
	"exacc":     ExadataTypeExacc,
}

// GetExadataTypeEnumValues Enumerates the set of values for ExadataTypeEnum
func GetExadataTypeEnumValues() []ExadataTypeEnum {
	values := make([]ExadataTypeEnum, 0)
	for _, v := range mappingExadataTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetExadataTypeEnumStringValues Enumerates the set of values in String for ExadataTypeEnum
func GetExadataTypeEnumStringValues() []string {
	return []string{
		"DBMACHINE",
		"EXACS",
		"EXACC",
	}
}

// GetMappingExadataTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingExadataTypeEnum(val string) (ExadataTypeEnum, bool) {
	enum, ok := mappingExadataTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
