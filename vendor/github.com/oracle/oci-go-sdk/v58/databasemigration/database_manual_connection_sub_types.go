// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Migration API
//
// Use the Oracle Cloud Infrastructure Database Migration APIs to perform database migration operations.
//

package databasemigration

// DatabaseManualConnectionSubTypesEnum Enum with underlying type: string
type DatabaseManualConnectionSubTypesEnum string

// Set of constants representing the allowable values for DatabaseManualConnectionSubTypesEnum
const (
	DatabaseManualConnectionSubTypesOracle    DatabaseManualConnectionSubTypesEnum = "ORACLE"
	DatabaseManualConnectionSubTypesRdsOracle DatabaseManualConnectionSubTypesEnum = "RDS_ORACLE"
)

var mappingDatabaseManualConnectionSubTypesEnum = map[string]DatabaseManualConnectionSubTypesEnum{
	"ORACLE":     DatabaseManualConnectionSubTypesOracle,
	"RDS_ORACLE": DatabaseManualConnectionSubTypesRdsOracle,
}

// GetDatabaseManualConnectionSubTypesEnumValues Enumerates the set of values for DatabaseManualConnectionSubTypesEnum
func GetDatabaseManualConnectionSubTypesEnumValues() []DatabaseManualConnectionSubTypesEnum {
	values := make([]DatabaseManualConnectionSubTypesEnum, 0)
	for _, v := range mappingDatabaseManualConnectionSubTypesEnum {
		values = append(values, v)
	}
	return values
}

// GetDatabaseManualConnectionSubTypesEnumStringValues Enumerates the set of values in String for DatabaseManualConnectionSubTypesEnum
func GetDatabaseManualConnectionSubTypesEnumStringValues() []string {
	return []string{
		"ORACLE",
		"RDS_ORACLE",
	}
}
