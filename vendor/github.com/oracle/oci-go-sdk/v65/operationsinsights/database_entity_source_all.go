// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Operations Insights API
//
// Use the Operations Insights API to perform data extraction operations to obtain database
// resource utilization, performance statistics, and reference information. For more information,
// see About Oracle Cloud Infrastructure Operations Insights (https://docs.cloud.oracle.com/en-us/iaas/operations-insights/doc/operations-insights.html).
//

package operationsinsights

import (
	"strings"
)

// DatabaseEntitySourceAllEnum Enum with underlying type: string
type DatabaseEntitySourceAllEnum string

// Set of constants representing the allowable values for DatabaseEntitySourceAllEnum
const (
	DatabaseEntitySourceAllAutonomousDatabase          DatabaseEntitySourceAllEnum = "AUTONOMOUS_DATABASE"
	DatabaseEntitySourceAllEmManagedExternalDatabase   DatabaseEntitySourceAllEnum = "EM_MANAGED_EXTERNAL_DATABASE"
	DatabaseEntitySourceAllMacsManagedExternalDatabase DatabaseEntitySourceAllEnum = "MACS_MANAGED_EXTERNAL_DATABASE"
	DatabaseEntitySourceAllPeComanagedDatabase         DatabaseEntitySourceAllEnum = "PE_COMANAGED_DATABASE"
)

var mappingDatabaseEntitySourceAllEnum = map[string]DatabaseEntitySourceAllEnum{
	"AUTONOMOUS_DATABASE":            DatabaseEntitySourceAllAutonomousDatabase,
	"EM_MANAGED_EXTERNAL_DATABASE":   DatabaseEntitySourceAllEmManagedExternalDatabase,
	"MACS_MANAGED_EXTERNAL_DATABASE": DatabaseEntitySourceAllMacsManagedExternalDatabase,
	"PE_COMANAGED_DATABASE":          DatabaseEntitySourceAllPeComanagedDatabase,
}

var mappingDatabaseEntitySourceAllEnumLowerCase = map[string]DatabaseEntitySourceAllEnum{
	"autonomous_database":            DatabaseEntitySourceAllAutonomousDatabase,
	"em_managed_external_database":   DatabaseEntitySourceAllEmManagedExternalDatabase,
	"macs_managed_external_database": DatabaseEntitySourceAllMacsManagedExternalDatabase,
	"pe_comanaged_database":          DatabaseEntitySourceAllPeComanagedDatabase,
}

// GetDatabaseEntitySourceAllEnumValues Enumerates the set of values for DatabaseEntitySourceAllEnum
func GetDatabaseEntitySourceAllEnumValues() []DatabaseEntitySourceAllEnum {
	values := make([]DatabaseEntitySourceAllEnum, 0)
	for _, v := range mappingDatabaseEntitySourceAllEnum {
		values = append(values, v)
	}
	return values
}

// GetDatabaseEntitySourceAllEnumStringValues Enumerates the set of values in String for DatabaseEntitySourceAllEnum
func GetDatabaseEntitySourceAllEnumStringValues() []string {
	return []string{
		"AUTONOMOUS_DATABASE",
		"EM_MANAGED_EXTERNAL_DATABASE",
		"MACS_MANAGED_EXTERNAL_DATABASE",
		"PE_COMANAGED_DATABASE",
	}
}

// GetMappingDatabaseEntitySourceAllEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDatabaseEntitySourceAllEnum(val string) (DatabaseEntitySourceAllEnum, bool) {
	enum, ok := mappingDatabaseEntitySourceAllEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
