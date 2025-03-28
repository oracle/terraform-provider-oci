// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Catalog API
//
// Use the Data Catalog APIs to collect, organize, find, access, understand, enrich, and activate technical, business, and operational metadata.
// For more information, see Data Catalog (https://docs.oracle.com/iaas/data-catalog/home.htm).
//

package datacatalog

import (
	"strings"
)

// JobScheduleTypeEnum Enum with underlying type: string
type JobScheduleTypeEnum string

// Set of constants representing the allowable values for JobScheduleTypeEnum
const (
	JobScheduleTypeScheduled JobScheduleTypeEnum = "SCHEDULED"
	JobScheduleTypeImmediate JobScheduleTypeEnum = "IMMEDIATE"
)

var mappingJobScheduleTypeEnum = map[string]JobScheduleTypeEnum{
	"SCHEDULED": JobScheduleTypeScheduled,
	"IMMEDIATE": JobScheduleTypeImmediate,
}

var mappingJobScheduleTypeEnumLowerCase = map[string]JobScheduleTypeEnum{
	"scheduled": JobScheduleTypeScheduled,
	"immediate": JobScheduleTypeImmediate,
}

// GetJobScheduleTypeEnumValues Enumerates the set of values for JobScheduleTypeEnum
func GetJobScheduleTypeEnumValues() []JobScheduleTypeEnum {
	values := make([]JobScheduleTypeEnum, 0)
	for _, v := range mappingJobScheduleTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetJobScheduleTypeEnumStringValues Enumerates the set of values in String for JobScheduleTypeEnum
func GetJobScheduleTypeEnumStringValues() []string {
	return []string{
		"SCHEDULED",
		"IMMEDIATE",
	}
}

// GetMappingJobScheduleTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingJobScheduleTypeEnum(val string) (JobScheduleTypeEnum, bool) {
	enum, ok := mappingJobScheduleTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
