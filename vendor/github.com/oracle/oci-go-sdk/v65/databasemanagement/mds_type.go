// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Management API
//
// Use the Database Management API to perform tasks such as obtaining performance and resource usage metrics
// for a fleet of Managed Databases or a specific Managed Database, creating Managed Database Groups, and
// running a SQL job on a Managed Database or Managed Database Group.
//

package databasemanagement

import (
	"strings"
)

// MdsTypeEnum Enum with underlying type: string
type MdsTypeEnum string

// Set of constants representing the allowable values for MdsTypeEnum
const (
	MdsTypeHa         MdsTypeEnum = "HA"
	MdsTypeHeatwave   MdsTypeEnum = "HEATWAVE"
	MdsTypeStandalone MdsTypeEnum = "STANDALONE"
)

var mappingMdsTypeEnum = map[string]MdsTypeEnum{
	"HA":         MdsTypeHa,
	"HEATWAVE":   MdsTypeHeatwave,
	"STANDALONE": MdsTypeStandalone,
}

var mappingMdsTypeEnumLowerCase = map[string]MdsTypeEnum{
	"ha":         MdsTypeHa,
	"heatwave":   MdsTypeHeatwave,
	"standalone": MdsTypeStandalone,
}

// GetMdsTypeEnumValues Enumerates the set of values for MdsTypeEnum
func GetMdsTypeEnumValues() []MdsTypeEnum {
	values := make([]MdsTypeEnum, 0)
	for _, v := range mappingMdsTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetMdsTypeEnumStringValues Enumerates the set of values in String for MdsTypeEnum
func GetMdsTypeEnumStringValues() []string {
	return []string{
		"HA",
		"HEATWAVE",
		"STANDALONE",
	}
}

// GetMappingMdsTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingMdsTypeEnum(val string) (MdsTypeEnum, bool) {
	enum, ok := mappingMdsTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
