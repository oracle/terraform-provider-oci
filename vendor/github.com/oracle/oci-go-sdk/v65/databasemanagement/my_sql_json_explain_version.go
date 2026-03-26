// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
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

// MySqlJsonExplainVersionEnum Enum with underlying type: string
type MySqlJsonExplainVersionEnum string

// Set of constants representing the allowable values for MySqlJsonExplainVersionEnum
const (
	MySqlJsonExplainVersionV1 MySqlJsonExplainVersionEnum = "V1"
	MySqlJsonExplainVersionV2 MySqlJsonExplainVersionEnum = "V2"
)

var mappingMySqlJsonExplainVersionEnum = map[string]MySqlJsonExplainVersionEnum{
	"V1": MySqlJsonExplainVersionV1,
	"V2": MySqlJsonExplainVersionV2,
}

var mappingMySqlJsonExplainVersionEnumLowerCase = map[string]MySqlJsonExplainVersionEnum{
	"v1": MySqlJsonExplainVersionV1,
	"v2": MySqlJsonExplainVersionV2,
}

// GetMySqlJsonExplainVersionEnumValues Enumerates the set of values for MySqlJsonExplainVersionEnum
func GetMySqlJsonExplainVersionEnumValues() []MySqlJsonExplainVersionEnum {
	values := make([]MySqlJsonExplainVersionEnum, 0)
	for _, v := range mappingMySqlJsonExplainVersionEnum {
		values = append(values, v)
	}
	return values
}

// GetMySqlJsonExplainVersionEnumStringValues Enumerates the set of values in String for MySqlJsonExplainVersionEnum
func GetMySqlJsonExplainVersionEnumStringValues() []string {
	return []string{
		"V1",
		"V2",
	}
}

// GetMappingMySqlJsonExplainVersionEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingMySqlJsonExplainVersionEnum(val string) (MySqlJsonExplainVersionEnum, bool) {
	enum, ok := mappingMySqlJsonExplainVersionEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
