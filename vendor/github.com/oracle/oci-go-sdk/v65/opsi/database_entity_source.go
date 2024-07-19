// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Ops Insights API
//
// Use the Ops Insights API to perform data extraction operations to obtain database
// resource utilization, performance statistics, and reference information. For more information,
// see About Oracle Cloud Infrastructure Ops Insights (https://docs.cloud.oracle.com/en-us/iaas/operations-insights/doc/operations-insights.html).
//

package opsi

import (
	"strings"
)

// DatabaseEntitySourceEnum Enum with underlying type: string
type DatabaseEntitySourceEnum string

// Set of constants representing the allowable values for DatabaseEntitySourceEnum
const (
	DatabaseEntitySourceEmManagedExternalDatabase DatabaseEntitySourceEnum = "EM_MANAGED_EXTERNAL_DATABASE"
	DatabaseEntitySourcePeComanagedDatabase       DatabaseEntitySourceEnum = "PE_COMANAGED_DATABASE"
	DatabaseEntitySourceMdsMysqlDatabaseSystem    DatabaseEntitySourceEnum = "MDS_MYSQL_DATABASE_SYSTEM"
)

var mappingDatabaseEntitySourceEnum = map[string]DatabaseEntitySourceEnum{
	"EM_MANAGED_EXTERNAL_DATABASE": DatabaseEntitySourceEmManagedExternalDatabase,
	"PE_COMANAGED_DATABASE":        DatabaseEntitySourcePeComanagedDatabase,
	"MDS_MYSQL_DATABASE_SYSTEM":    DatabaseEntitySourceMdsMysqlDatabaseSystem,
}

var mappingDatabaseEntitySourceEnumLowerCase = map[string]DatabaseEntitySourceEnum{
	"em_managed_external_database": DatabaseEntitySourceEmManagedExternalDatabase,
	"pe_comanaged_database":        DatabaseEntitySourcePeComanagedDatabase,
	"mds_mysql_database_system":    DatabaseEntitySourceMdsMysqlDatabaseSystem,
}

// GetDatabaseEntitySourceEnumValues Enumerates the set of values for DatabaseEntitySourceEnum
func GetDatabaseEntitySourceEnumValues() []DatabaseEntitySourceEnum {
	values := make([]DatabaseEntitySourceEnum, 0)
	for _, v := range mappingDatabaseEntitySourceEnum {
		values = append(values, v)
	}
	return values
}

// GetDatabaseEntitySourceEnumStringValues Enumerates the set of values in String for DatabaseEntitySourceEnum
func GetDatabaseEntitySourceEnumStringValues() []string {
	return []string{
		"EM_MANAGED_EXTERNAL_DATABASE",
		"PE_COMANAGED_DATABASE",
		"MDS_MYSQL_DATABASE_SYSTEM",
	}
}

// GetMappingDatabaseEntitySourceEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDatabaseEntitySourceEnum(val string) (DatabaseEntitySourceEnum, bool) {
	enum, ok := mappingDatabaseEntitySourceEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
