// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Management API
//
// Use the Database Management API to monitor and manage resources such as
// Oracle Databases, MySQL Databases, and External Database Systems.
// For more information, see Database Management (https://docs.oracle.com/iaas/database-management/home.htm).
//

package databasemanagement

import (
	"strings"
)

// MySqlDeploymentTypeEnum Enum with underlying type: string
type MySqlDeploymentTypeEnum string

// Set of constants representing the allowable values for MySqlDeploymentTypeEnum
const (
	MySqlDeploymentTypeOnpremise MySqlDeploymentTypeEnum = "ONPREMISE"
	MySqlDeploymentTypeMds       MySqlDeploymentTypeEnum = "MDS"
)

var mappingMySqlDeploymentTypeEnum = map[string]MySqlDeploymentTypeEnum{
	"ONPREMISE": MySqlDeploymentTypeOnpremise,
	"MDS":       MySqlDeploymentTypeMds,
}

var mappingMySqlDeploymentTypeEnumLowerCase = map[string]MySqlDeploymentTypeEnum{
	"onpremise": MySqlDeploymentTypeOnpremise,
	"mds":       MySqlDeploymentTypeMds,
}

// GetMySqlDeploymentTypeEnumValues Enumerates the set of values for MySqlDeploymentTypeEnum
func GetMySqlDeploymentTypeEnumValues() []MySqlDeploymentTypeEnum {
	values := make([]MySqlDeploymentTypeEnum, 0)
	for _, v := range mappingMySqlDeploymentTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetMySqlDeploymentTypeEnumStringValues Enumerates the set of values in String for MySqlDeploymentTypeEnum
func GetMySqlDeploymentTypeEnumStringValues() []string {
	return []string{
		"ONPREMISE",
		"MDS",
	}
}

// GetMappingMySqlDeploymentTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingMySqlDeploymentTypeEnum(val string) (MySqlDeploymentTypeEnum, bool) {
	enum, ok := mappingMySqlDeploymentTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
