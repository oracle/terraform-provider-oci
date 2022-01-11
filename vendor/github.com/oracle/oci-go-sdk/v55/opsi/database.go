// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Operations Insights API
//
// Use the Operations Insights API to perform data extraction operations to obtain database
// resource utilization, performance statistics, and reference information. For more information,
// see About Oracle Cloud Infrastructure Operations Insights (https://docs.cloud.oracle.com/en-us/iaas/operations-insights/doc/operations-insights.html).
//

package opsi

// DatabaseEnum Enum with underlying type: string
type DatabaseEnum string

// Set of constants representing the allowable values for DatabaseEnum
const (
	DatabaseCpu     DatabaseEnum = "CPU"
	DatabaseStorage DatabaseEnum = "STORAGE"
	DatabaseIo      DatabaseEnum = "IO"
	DatabaseMemory  DatabaseEnum = "MEMORY"
)

var mappingDatabase = map[string]DatabaseEnum{
	"CPU":     DatabaseCpu,
	"STORAGE": DatabaseStorage,
	"IO":      DatabaseIo,
	"MEMORY":  DatabaseMemory,
}

// GetDatabaseEnumValues Enumerates the set of values for DatabaseEnum
func GetDatabaseEnumValues() []DatabaseEnum {
	values := make([]DatabaseEnum, 0)
	for _, v := range mappingDatabase {
		values = append(values, v)
	}
	return values
}
