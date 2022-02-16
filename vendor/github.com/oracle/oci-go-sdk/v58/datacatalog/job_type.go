// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
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

// JobTypeEnum Enum with underlying type: string
type JobTypeEnum string

// Set of constants representing the allowable values for JobTypeEnum
const (
	JobTypeHarvest                    JobTypeEnum = "HARVEST"
	JobTypeProfiling                  JobTypeEnum = "PROFILING"
	JobTypeSampling                   JobTypeEnum = "SAMPLING"
	JobTypePreview                    JobTypeEnum = "PREVIEW"
	JobTypeImport                     JobTypeEnum = "IMPORT"
	JobTypeExport                     JobTypeEnum = "EXPORT"
	JobTypeImportGlossary             JobTypeEnum = "IMPORT_GLOSSARY"
	JobTypeExportGlossary             JobTypeEnum = "EXPORT_GLOSSARY"
	JobTypeInternal                   JobTypeEnum = "INTERNAL"
	JobTypePurge                      JobTypeEnum = "PURGE"
	JobTypeImmediate                  JobTypeEnum = "IMMEDIATE"
	JobTypeScheduled                  JobTypeEnum = "SCHEDULED"
	JobTypeImmediateExecution         JobTypeEnum = "IMMEDIATE_EXECUTION"
	JobTypeScheduledExecution         JobTypeEnum = "SCHEDULED_EXECUTION"
	JobTypeScheduledExecutionInstance JobTypeEnum = "SCHEDULED_EXECUTION_INSTANCE"
	JobTypeAsyncDelete                JobTypeEnum = "ASYNC_DELETE"
	JobTypeImportDataAsset            JobTypeEnum = "IMPORT_DATA_ASSET"
)

var mappingJobTypeEnum = map[string]JobTypeEnum{
	"HARVEST":                      JobTypeHarvest,
	"PROFILING":                    JobTypeProfiling,
	"SAMPLING":                     JobTypeSampling,
	"PREVIEW":                      JobTypePreview,
	"IMPORT":                       JobTypeImport,
	"EXPORT":                       JobTypeExport,
	"IMPORT_GLOSSARY":              JobTypeImportGlossary,
	"EXPORT_GLOSSARY":              JobTypeExportGlossary,
	"INTERNAL":                     JobTypeInternal,
	"PURGE":                        JobTypePurge,
	"IMMEDIATE":                    JobTypeImmediate,
	"SCHEDULED":                    JobTypeScheduled,
	"IMMEDIATE_EXECUTION":          JobTypeImmediateExecution,
	"SCHEDULED_EXECUTION":          JobTypeScheduledExecution,
	"SCHEDULED_EXECUTION_INSTANCE": JobTypeScheduledExecutionInstance,
	"ASYNC_DELETE":                 JobTypeAsyncDelete,
	"IMPORT_DATA_ASSET":            JobTypeImportDataAsset,
}

// GetJobTypeEnumValues Enumerates the set of values for JobTypeEnum
func GetJobTypeEnumValues() []JobTypeEnum {
	values := make([]JobTypeEnum, 0)
	for _, v := range mappingJobTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetJobTypeEnumStringValues Enumerates the set of values in String for JobTypeEnum
func GetJobTypeEnumStringValues() []string {
	return []string{
		"HARVEST",
		"PROFILING",
		"SAMPLING",
		"PREVIEW",
		"IMPORT",
		"EXPORT",
		"IMPORT_GLOSSARY",
		"EXPORT_GLOSSARY",
		"INTERNAL",
		"PURGE",
		"IMMEDIATE",
		"SCHEDULED",
		"IMMEDIATE_EXECUTION",
		"SCHEDULED_EXECUTION",
		"SCHEDULED_EXECUTION_INSTANCE",
		"ASYNC_DELETE",
		"IMPORT_DATA_ASSET",
	}
}

// GetMappingJobTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingJobTypeEnum(val string) (JobTypeEnum, bool) {
	mappingJobTypeEnumIgnoreCase := make(map[string]JobTypeEnum)
	for k, v := range mappingJobTypeEnum {
		mappingJobTypeEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingJobTypeEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}
