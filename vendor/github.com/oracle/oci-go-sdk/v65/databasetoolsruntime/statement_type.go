// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Tools Runtime API
//
// Use the Database Tools Runtime API to connect to databases through Database Tools Connections.
//

package databasetoolsruntime

import (
	"strings"
)

// StatementTypeEnum Enum with underlying type: string
type StatementTypeEnum string

// Set of constants representing the allowable values for StatementTypeEnum
const (
	StatementTypeQuery              StatementTypeEnum = "QUERY"
	StatementTypeDml                StatementTypeEnum = "DML"
	StatementTypeDdl                StatementTypeEnum = "DDL"
	StatementTypePlsql              StatementTypeEnum = "PLSQL"
	StatementTypeSqlplus            StatementTypeEnum = "SQLPLUS"
	StatementTypeIgnore             StatementTypeEnum = "IGNORE"
	StatementTypeTransactionControl StatementTypeEnum = "TRANSACTION_CONTROL"
	StatementTypeSessionControl     StatementTypeEnum = "SESSION_CONTROL"
	StatementTypeSystemControl      StatementTypeEnum = "SYSTEM_CONTROL"
	StatementTypeJdbc               StatementTypeEnum = "JDBC"
	StatementTypeOther              StatementTypeEnum = "OTHER"
)

var mappingStatementTypeEnum = map[string]StatementTypeEnum{
	"QUERY":               StatementTypeQuery,
	"DML":                 StatementTypeDml,
	"DDL":                 StatementTypeDdl,
	"PLSQL":               StatementTypePlsql,
	"SQLPLUS":             StatementTypeSqlplus,
	"IGNORE":              StatementTypeIgnore,
	"TRANSACTION_CONTROL": StatementTypeTransactionControl,
	"SESSION_CONTROL":     StatementTypeSessionControl,
	"SYSTEM_CONTROL":      StatementTypeSystemControl,
	"JDBC":                StatementTypeJdbc,
	"OTHER":               StatementTypeOther,
}

var mappingStatementTypeEnumLowerCase = map[string]StatementTypeEnum{
	"query":               StatementTypeQuery,
	"dml":                 StatementTypeDml,
	"ddl":                 StatementTypeDdl,
	"plsql":               StatementTypePlsql,
	"sqlplus":             StatementTypeSqlplus,
	"ignore":              StatementTypeIgnore,
	"transaction_control": StatementTypeTransactionControl,
	"session_control":     StatementTypeSessionControl,
	"system_control":      StatementTypeSystemControl,
	"jdbc":                StatementTypeJdbc,
	"other":               StatementTypeOther,
}

// GetStatementTypeEnumValues Enumerates the set of values for StatementTypeEnum
func GetStatementTypeEnumValues() []StatementTypeEnum {
	values := make([]StatementTypeEnum, 0)
	for _, v := range mappingStatementTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetStatementTypeEnumStringValues Enumerates the set of values in String for StatementTypeEnum
func GetStatementTypeEnumStringValues() []string {
	return []string{
		"QUERY",
		"DML",
		"DDL",
		"PLSQL",
		"SQLPLUS",
		"IGNORE",
		"TRANSACTION_CONTROL",
		"SESSION_CONTROL",
		"SYSTEM_CONTROL",
		"JDBC",
		"OTHER",
	}
}

// GetMappingStatementTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingStatementTypeEnum(val string) (StatementTypeEnum, bool) {
	enum, ok := mappingStatementTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
