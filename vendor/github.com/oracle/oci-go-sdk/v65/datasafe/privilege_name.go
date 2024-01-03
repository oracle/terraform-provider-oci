// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Safe API
//
// APIs for using Oracle Data Safe.
//

package datasafe

import (
	"strings"
)

// PrivilegeNameEnum Enum with underlying type: string
type PrivilegeNameEnum string

// Set of constants representing the allowable values for PrivilegeNameEnum
const (
	PrivilegeNameSelect         PrivilegeNameEnum = "SELECT"
	PrivilegeNameUpdate         PrivilegeNameEnum = "UPDATE"
	PrivilegeNameInsert         PrivilegeNameEnum = "INSERT"
	PrivilegeNameDelete         PrivilegeNameEnum = "DELETE"
	PrivilegeNameRead           PrivilegeNameEnum = "READ"
	PrivilegeNameOwner          PrivilegeNameEnum = "OWNER"
	PrivilegeNameIndex          PrivilegeNameEnum = "INDEX"
	PrivilegeNameSelectAnyTable PrivilegeNameEnum = "SELECT_ANY_TABLE"
	PrivilegeNameUpdateAnyTable PrivilegeNameEnum = "UPDATE_ANY_TABLE"
	PrivilegeNameInsertAnyTable PrivilegeNameEnum = "INSERT_ANY_TABLE"
	PrivilegeNameDeleteAnyTable PrivilegeNameEnum = "DELETE_ANY_TABLE"
	PrivilegeNameReadAnyTable   PrivilegeNameEnum = "READ_ANY_TABLE"
	PrivilegeNameCreateAnyIndex PrivilegeNameEnum = "CREATE_ANY_INDEX"
)

var mappingPrivilegeNameEnum = map[string]PrivilegeNameEnum{
	"SELECT":           PrivilegeNameSelect,
	"UPDATE":           PrivilegeNameUpdate,
	"INSERT":           PrivilegeNameInsert,
	"DELETE":           PrivilegeNameDelete,
	"READ":             PrivilegeNameRead,
	"OWNER":            PrivilegeNameOwner,
	"INDEX":            PrivilegeNameIndex,
	"SELECT_ANY_TABLE": PrivilegeNameSelectAnyTable,
	"UPDATE_ANY_TABLE": PrivilegeNameUpdateAnyTable,
	"INSERT_ANY_TABLE": PrivilegeNameInsertAnyTable,
	"DELETE_ANY_TABLE": PrivilegeNameDeleteAnyTable,
	"READ_ANY_TABLE":   PrivilegeNameReadAnyTable,
	"CREATE_ANY_INDEX": PrivilegeNameCreateAnyIndex,
}

var mappingPrivilegeNameEnumLowerCase = map[string]PrivilegeNameEnum{
	"select":           PrivilegeNameSelect,
	"update":           PrivilegeNameUpdate,
	"insert":           PrivilegeNameInsert,
	"delete":           PrivilegeNameDelete,
	"read":             PrivilegeNameRead,
	"owner":            PrivilegeNameOwner,
	"index":            PrivilegeNameIndex,
	"select_any_table": PrivilegeNameSelectAnyTable,
	"update_any_table": PrivilegeNameUpdateAnyTable,
	"insert_any_table": PrivilegeNameInsertAnyTable,
	"delete_any_table": PrivilegeNameDeleteAnyTable,
	"read_any_table":   PrivilegeNameReadAnyTable,
	"create_any_index": PrivilegeNameCreateAnyIndex,
}

// GetPrivilegeNameEnumValues Enumerates the set of values for PrivilegeNameEnum
func GetPrivilegeNameEnumValues() []PrivilegeNameEnum {
	values := make([]PrivilegeNameEnum, 0)
	for _, v := range mappingPrivilegeNameEnum {
		values = append(values, v)
	}
	return values
}

// GetPrivilegeNameEnumStringValues Enumerates the set of values in String for PrivilegeNameEnum
func GetPrivilegeNameEnumStringValues() []string {
	return []string{
		"SELECT",
		"UPDATE",
		"INSERT",
		"DELETE",
		"READ",
		"OWNER",
		"INDEX",
		"SELECT_ANY_TABLE",
		"UPDATE_ANY_TABLE",
		"INSERT_ANY_TABLE",
		"DELETE_ANY_TABLE",
		"READ_ANY_TABLE",
		"CREATE_ANY_INDEX",
	}
}

// GetMappingPrivilegeNameEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingPrivilegeNameEnum(val string) (PrivilegeNameEnum, bool) {
	enum, ok := mappingPrivilegeNameEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
