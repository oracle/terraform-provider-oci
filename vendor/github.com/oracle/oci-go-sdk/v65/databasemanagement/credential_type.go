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

// CredentialTypeEnum Enum with underlying type: string
type CredentialTypeEnum string

// Set of constants representing the allowable values for CredentialTypeEnum
const (
	CredentialTypeOracleDb CredentialTypeEnum = "ORACLE_DB"
)

var mappingCredentialTypeEnum = map[string]CredentialTypeEnum{
	"ORACLE_DB": CredentialTypeOracleDb,
}

var mappingCredentialTypeEnumLowerCase = map[string]CredentialTypeEnum{
	"oracle_db": CredentialTypeOracleDb,
}

// GetCredentialTypeEnumValues Enumerates the set of values for CredentialTypeEnum
func GetCredentialTypeEnumValues() []CredentialTypeEnum {
	values := make([]CredentialTypeEnum, 0)
	for _, v := range mappingCredentialTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetCredentialTypeEnumStringValues Enumerates the set of values in String for CredentialTypeEnum
func GetCredentialTypeEnumStringValues() []string {
	return []string{
		"ORACLE_DB",
	}
}

// GetMappingCredentialTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingCredentialTypeEnum(val string) (CredentialTypeEnum, bool) {
	enum, ok := mappingCredentialTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
