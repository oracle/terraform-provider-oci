// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Operations Insights API
//
// Use the Operations Insights API to perform data extraction operations to obtain database
// resource utilization, performance statistics, and reference information. For more information,
// see About Oracle Cloud Infrastructure Operations Insights (https://docs.cloud.oracle.com/en-us/iaas/operations-insights/doc/operations-insights.html).
//

package operationsinsights

// DatabaseEntitySourceEnum Enum with underlying type: string
type DatabaseEntitySourceEnum string

// Set of constants representing the allowable values for DatabaseEntitySourceEnum
const (
	DatabaseEntitySourceEmManagedExternalDatabase DatabaseEntitySourceEnum = "EM_MANAGED_EXTERNAL_DATABASE"
	DatabaseEntitySourcePeComanagedDatabase       DatabaseEntitySourceEnum = "PE_COMANAGED_DATABASE"
)

var mappingDatabaseEntitySourceEnum = map[string]DatabaseEntitySourceEnum{
	"EM_MANAGED_EXTERNAL_DATABASE": DatabaseEntitySourceEmManagedExternalDatabase,
	"PE_COMANAGED_DATABASE":        DatabaseEntitySourcePeComanagedDatabase,
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
	}
}
