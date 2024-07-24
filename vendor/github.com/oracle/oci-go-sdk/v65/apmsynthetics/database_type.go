// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Application Performance Monitoring Synthetic Monitoring API
//
// Use the Application Performance Monitoring Synthetic Monitoring API to query synthetic scripts and monitors. For more information, see Application Performance Monitoring (https://docs.oracle.com/iaas/application-performance-monitoring/index.html).
//

package apmsynthetics

import (
	"strings"
)

// DatabaseTypeEnum Enum with underlying type: string
type DatabaseTypeEnum string

// Set of constants representing the allowable values for DatabaseTypeEnum
const (
	DatabaseTypeOracle DatabaseTypeEnum = "ORACLE"
	DatabaseTypeMysql  DatabaseTypeEnum = "MYSQL"
)

var mappingDatabaseTypeEnum = map[string]DatabaseTypeEnum{
	"ORACLE": DatabaseTypeOracle,
	"MYSQL":  DatabaseTypeMysql,
}

var mappingDatabaseTypeEnumLowerCase = map[string]DatabaseTypeEnum{
	"oracle": DatabaseTypeOracle,
	"mysql":  DatabaseTypeMysql,
}

// GetDatabaseTypeEnumValues Enumerates the set of values for DatabaseTypeEnum
func GetDatabaseTypeEnumValues() []DatabaseTypeEnum {
	values := make([]DatabaseTypeEnum, 0)
	for _, v := range mappingDatabaseTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetDatabaseTypeEnumStringValues Enumerates the set of values in String for DatabaseTypeEnum
func GetDatabaseTypeEnumStringValues() []string {
	return []string{
		"ORACLE",
		"MYSQL",
	}
}

// GetMappingDatabaseTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDatabaseTypeEnum(val string) (DatabaseTypeEnum, bool) {
	enum, ok := mappingDatabaseTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
