// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Catalog API
//
// Use the Data Catalog APIs to collect, organize, find, access, understand, enrich, and activate technical, business, and operational metadata.
//

package datacatalog

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
)

var mappingJobType = map[string]JobTypeEnum{
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
}

// GetJobTypeEnumValues Enumerates the set of values for JobTypeEnum
func GetJobTypeEnumValues() []JobTypeEnum {
	values := make([]JobTypeEnum, 0)
	for _, v := range mappingJobType {
		values = append(values, v)
	}
	return values
}
