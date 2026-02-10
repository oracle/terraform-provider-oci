// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Migration API
//
// Use the Oracle Cloud Infrastructure Database Migration APIs to perform database migration operations.
//

package databasemigration

import (
	"strings"
)

// AssessorActionNamesEnum Enum with underlying type: string
type AssessorActionNamesEnum string

// Set of constants representing the allowable values for AssessorActionNamesEnum
const (
	AssessorActionNamesRun                AssessorActionNamesEnum = "RUN"
	AssessorActionNamesConfigure          AssessorActionNamesEnum = "CONFIGURE"
	AssessorActionNamesConfigureMigration AssessorActionNamesEnum = "CONFIGURE_MIGRATION"
	AssessorActionNamesCreateMigration    AssessorActionNamesEnum = "CREATE_MIGRATION"
	AssessorActionNamesUpdateMigration    AssessorActionNamesEnum = "UPDATE_MIGRATION"
	AssessorActionNamesRecheck            AssessorActionNamesEnum = "RECHECK"
	AssessorActionNamesRunSql             AssessorActionNamesEnum = "RUN_SQL"
	AssessorActionNamesDownloadSql        AssessorActionNamesEnum = "DOWNLOAD_SQL"
	AssessorActionNamesDownloadLog        AssessorActionNamesEnum = "DOWNLOAD_LOG"
	AssessorActionNamesConfirm            AssessorActionNamesEnum = "CONFIRM"
	AssessorActionNamesManage             AssessorActionNamesEnum = "MANAGE"
)

var mappingAssessorActionNamesEnum = map[string]AssessorActionNamesEnum{
	"RUN":                 AssessorActionNamesRun,
	"CONFIGURE":           AssessorActionNamesConfigure,
	"CONFIGURE_MIGRATION": AssessorActionNamesConfigureMigration,
	"CREATE_MIGRATION":    AssessorActionNamesCreateMigration,
	"UPDATE_MIGRATION":    AssessorActionNamesUpdateMigration,
	"RECHECK":             AssessorActionNamesRecheck,
	"RUN_SQL":             AssessorActionNamesRunSql,
	"DOWNLOAD_SQL":        AssessorActionNamesDownloadSql,
	"DOWNLOAD_LOG":        AssessorActionNamesDownloadLog,
	"CONFIRM":             AssessorActionNamesConfirm,
	"MANAGE":              AssessorActionNamesManage,
}

var mappingAssessorActionNamesEnumLowerCase = map[string]AssessorActionNamesEnum{
	"run":                 AssessorActionNamesRun,
	"configure":           AssessorActionNamesConfigure,
	"configure_migration": AssessorActionNamesConfigureMigration,
	"create_migration":    AssessorActionNamesCreateMigration,
	"update_migration":    AssessorActionNamesUpdateMigration,
	"recheck":             AssessorActionNamesRecheck,
	"run_sql":             AssessorActionNamesRunSql,
	"download_sql":        AssessorActionNamesDownloadSql,
	"download_log":        AssessorActionNamesDownloadLog,
	"confirm":             AssessorActionNamesConfirm,
	"manage":              AssessorActionNamesManage,
}

// GetAssessorActionNamesEnumValues Enumerates the set of values for AssessorActionNamesEnum
func GetAssessorActionNamesEnumValues() []AssessorActionNamesEnum {
	values := make([]AssessorActionNamesEnum, 0)
	for _, v := range mappingAssessorActionNamesEnum {
		values = append(values, v)
	}
	return values
}

// GetAssessorActionNamesEnumStringValues Enumerates the set of values in String for AssessorActionNamesEnum
func GetAssessorActionNamesEnumStringValues() []string {
	return []string{
		"RUN",
		"CONFIGURE",
		"CONFIGURE_MIGRATION",
		"CREATE_MIGRATION",
		"UPDATE_MIGRATION",
		"RECHECK",
		"RUN_SQL",
		"DOWNLOAD_SQL",
		"DOWNLOAD_LOG",
		"CONFIRM",
		"MANAGE",
	}
}

// GetMappingAssessorActionNamesEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingAssessorActionNamesEnum(val string) (AssessorActionNamesEnum, bool) {
	enum, ok := mappingAssessorActionNamesEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
