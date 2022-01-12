// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Management API
//
// Use the Database Management API to perform tasks such as obtaining performance and resource usage metrics
// for a fleet of Managed Databases or a specific Managed Database, creating Managed Database Groups, and
// running a SQL job on a Managed Database or Managed Database Group.
//

package databasemanagement

// CompareTypeEnum Enum with underlying type: string
type CompareTypeEnum string

// Set of constants representing the allowable values for CompareTypeEnum
const (
	CompareTypeHour CompareTypeEnum = "HOUR"
	CompareTypeDay  CompareTypeEnum = "DAY"
)

var mappingCompareType = map[string]CompareTypeEnum{
	"HOUR": CompareTypeHour,
	"DAY":  CompareTypeDay,
}

// GetCompareTypeEnumValues Enumerates the set of values for CompareTypeEnum
func GetCompareTypeEnumValues() []CompareTypeEnum {
	values := make([]CompareTypeEnum, 0)
	for _, v := range mappingCompareType {
		values = append(values, v)
	}
	return values
}
