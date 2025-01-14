// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Management API
//
// Use the Database Management API to monitor and manage resources such as
// Oracle Databases, MySQL Databases, and External Database Systems.
// For more information, see Database Management (https://docs.cloud.oracle.com/iaas/database-management/home.htm).
//

package databasemanagement

import (
	"strings"
)

// ExternalMysqlAssociatedServiceNameEnum Enum with underlying type: string
type ExternalMysqlAssociatedServiceNameEnum string

// Set of constants representing the allowable values for ExternalMysqlAssociatedServiceNameEnum
const (
	ExternalMysqlAssociatedServiceNameOperationsInsights ExternalMysqlAssociatedServiceNameEnum = "OPERATIONS_INSIGHTS"
)

var mappingExternalMysqlAssociatedServiceNameEnum = map[string]ExternalMysqlAssociatedServiceNameEnum{
	"OPERATIONS_INSIGHTS": ExternalMysqlAssociatedServiceNameOperationsInsights,
}

var mappingExternalMysqlAssociatedServiceNameEnumLowerCase = map[string]ExternalMysqlAssociatedServiceNameEnum{
	"operations_insights": ExternalMysqlAssociatedServiceNameOperationsInsights,
}

// GetExternalMysqlAssociatedServiceNameEnumValues Enumerates the set of values for ExternalMysqlAssociatedServiceNameEnum
func GetExternalMysqlAssociatedServiceNameEnumValues() []ExternalMysqlAssociatedServiceNameEnum {
	values := make([]ExternalMysqlAssociatedServiceNameEnum, 0)
	for _, v := range mappingExternalMysqlAssociatedServiceNameEnum {
		values = append(values, v)
	}
	return values
}

// GetExternalMysqlAssociatedServiceNameEnumStringValues Enumerates the set of values in String for ExternalMysqlAssociatedServiceNameEnum
func GetExternalMysqlAssociatedServiceNameEnumStringValues() []string {
	return []string{
		"OPERATIONS_INSIGHTS",
	}
}

// GetMappingExternalMysqlAssociatedServiceNameEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingExternalMysqlAssociatedServiceNameEnum(val string) (ExternalMysqlAssociatedServiceNameEnum, bool) {
	enum, ok := mappingExternalMysqlAssociatedServiceNameEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
