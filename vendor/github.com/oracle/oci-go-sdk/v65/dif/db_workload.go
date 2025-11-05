// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// DataIntelligences Control Plane API
//
// Use the DataIntelligences Control Plane API to manage dataIntelligences.
//

package dif

import (
	"strings"
)

// DbWorkloadEnum Enum with underlying type: string
type DbWorkloadEnum string

// Set of constants representing the allowable values for DbWorkloadEnum
const (
	DbWorkloadOltp DbWorkloadEnum = "OLTP"
	DbWorkloadDw   DbWorkloadEnum = "DW"
)

var mappingDbWorkloadEnum = map[string]DbWorkloadEnum{
	"OLTP": DbWorkloadOltp,
	"DW":   DbWorkloadDw,
}

var mappingDbWorkloadEnumLowerCase = map[string]DbWorkloadEnum{
	"oltp": DbWorkloadOltp,
	"dw":   DbWorkloadDw,
}

// GetDbWorkloadEnumValues Enumerates the set of values for DbWorkloadEnum
func GetDbWorkloadEnumValues() []DbWorkloadEnum {
	values := make([]DbWorkloadEnum, 0)
	for _, v := range mappingDbWorkloadEnum {
		values = append(values, v)
	}
	return values
}

// GetDbWorkloadEnumStringValues Enumerates the set of values in String for DbWorkloadEnum
func GetDbWorkloadEnumStringValues() []string {
	return []string{
		"OLTP",
		"DW",
	}
}

// GetMappingDbWorkloadEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDbWorkloadEnum(val string) (DbWorkloadEnum, bool) {
	enum, ok := mappingDbWorkloadEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
