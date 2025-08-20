// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Object Storage Service API
//
// Use Object Storage and Archive Storage APIs to manage buckets, objects, and related resources.
// For more information, see Overview of Object Storage (https://docs.oracle.com/iaas/Content/Object/Concepts/objectstorageoverview.htm) and
// Overview of Archive Storage (https://docs.oracle.com/iaas/Content/Archive/Concepts/archivestorageoverview.htm).
//

package objectstorage

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// Report Defines output report.
type Report struct {

	// An array of fields to be included in report.
	// The valid schemaFields for 'object' rules are:
	//   - SIZE
	//   - ETAG
	//   - TIME_CREATED
	//   - MD5_CHECKSUM
	//   - TIME_MODIFIED
	//   - STORAGE_TIER
	//   - ARCHIVAL_STATE
	//   - COMPARTMENT_ID
	//   - CREATED_BY
	// The valid schemaFields for 'bucket' rules are:
	//   - ETAG
	//   - TIME_CREATED
	//   - TIME_MODIFIED
	//   - COMPARTMENT_ID
	//   - CREATED_BY
	//   - DEFINED_TAGS
	//   - FREEFORM_TAGS
	SchemaFields []ReportSchemaFieldsEnum `mandatory:"true" json:"schemaFields"`

	// The target bucket name to upload the generated report. Create the target bucket before creating the rule.
	TargetBucket *string `mandatory:"true" json:"targetBucket"`

	// The inventory report format specification
	Format ReportFormatEnum `mandatory:"false" json:"format,omitempty"`
}

func (m Report) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m Report) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	for _, val := range m.SchemaFields {
		if _, ok := GetMappingReportSchemaFieldsEnum(string(val)); !ok && val != "" {
			errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SchemaFields: %s. Supported values are: %s.", val, strings.Join(GetReportSchemaFieldsEnumStringValues(), ",")))
		}
	}

	if _, ok := GetMappingReportFormatEnum(string(m.Format)); !ok && m.Format != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Format: %s. Supported values are: %s.", m.Format, strings.Join(GetReportFormatEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ReportFormatEnum Enum with underlying type: string
type ReportFormatEnum string

// Set of constants representing the allowable values for ReportFormatEnum
const (
	ReportFormatParquet ReportFormatEnum = "PARQUET"
)

var mappingReportFormatEnum = map[string]ReportFormatEnum{
	"PARQUET": ReportFormatParquet,
}

var mappingReportFormatEnumLowerCase = map[string]ReportFormatEnum{
	"parquet": ReportFormatParquet,
}

// GetReportFormatEnumValues Enumerates the set of values for ReportFormatEnum
func GetReportFormatEnumValues() []ReportFormatEnum {
	values := make([]ReportFormatEnum, 0)
	for _, v := range mappingReportFormatEnum {
		values = append(values, v)
	}
	return values
}

// GetReportFormatEnumStringValues Enumerates the set of values in String for ReportFormatEnum
func GetReportFormatEnumStringValues() []string {
	return []string{
		"PARQUET",
	}
}

// GetMappingReportFormatEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingReportFormatEnum(val string) (ReportFormatEnum, bool) {
	enum, ok := mappingReportFormatEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ReportSchemaFieldsEnum Enum with underlying type: string
type ReportSchemaFieldsEnum string

// Set of constants representing the allowable values for ReportSchemaFieldsEnum
const (
	ReportSchemaFieldsSize          ReportSchemaFieldsEnum = "SIZE"
	ReportSchemaFieldsEtag          ReportSchemaFieldsEnum = "ETAG"
	ReportSchemaFieldsTimeCreated   ReportSchemaFieldsEnum = "TIME_CREATED"
	ReportSchemaFieldsMd5Checksum   ReportSchemaFieldsEnum = "MD5_CHECKSUM"
	ReportSchemaFieldsTimeModified  ReportSchemaFieldsEnum = "TIME_MODIFIED"
	ReportSchemaFieldsStorageTier   ReportSchemaFieldsEnum = "STORAGE_TIER"
	ReportSchemaFieldsArchivalState ReportSchemaFieldsEnum = "ARCHIVAL_STATE"
	ReportSchemaFieldsCompartmentId ReportSchemaFieldsEnum = "COMPARTMENT_ID"
	ReportSchemaFieldsCreatedBy     ReportSchemaFieldsEnum = "CREATED_BY"
	ReportSchemaFieldsDefinedTags   ReportSchemaFieldsEnum = "DEFINED_TAGS"
	ReportSchemaFieldsFreeformTags  ReportSchemaFieldsEnum = "FREEFORM_TAGS"
)

var mappingReportSchemaFieldsEnum = map[string]ReportSchemaFieldsEnum{
	"SIZE":           ReportSchemaFieldsSize,
	"ETAG":           ReportSchemaFieldsEtag,
	"TIME_CREATED":   ReportSchemaFieldsTimeCreated,
	"MD5_CHECKSUM":   ReportSchemaFieldsMd5Checksum,
	"TIME_MODIFIED":  ReportSchemaFieldsTimeModified,
	"STORAGE_TIER":   ReportSchemaFieldsStorageTier,
	"ARCHIVAL_STATE": ReportSchemaFieldsArchivalState,
	"COMPARTMENT_ID": ReportSchemaFieldsCompartmentId,
	"CREATED_BY":     ReportSchemaFieldsCreatedBy,
	"DEFINED_TAGS":   ReportSchemaFieldsDefinedTags,
	"FREEFORM_TAGS":  ReportSchemaFieldsFreeformTags,
}

var mappingReportSchemaFieldsEnumLowerCase = map[string]ReportSchemaFieldsEnum{
	"size":           ReportSchemaFieldsSize,
	"etag":           ReportSchemaFieldsEtag,
	"time_created":   ReportSchemaFieldsTimeCreated,
	"md5_checksum":   ReportSchemaFieldsMd5Checksum,
	"time_modified":  ReportSchemaFieldsTimeModified,
	"storage_tier":   ReportSchemaFieldsStorageTier,
	"archival_state": ReportSchemaFieldsArchivalState,
	"compartment_id": ReportSchemaFieldsCompartmentId,
	"created_by":     ReportSchemaFieldsCreatedBy,
	"defined_tags":   ReportSchemaFieldsDefinedTags,
	"freeform_tags":  ReportSchemaFieldsFreeformTags,
}

// GetReportSchemaFieldsEnumValues Enumerates the set of values for ReportSchemaFieldsEnum
func GetReportSchemaFieldsEnumValues() []ReportSchemaFieldsEnum {
	values := make([]ReportSchemaFieldsEnum, 0)
	for _, v := range mappingReportSchemaFieldsEnum {
		values = append(values, v)
	}
	return values
}

// GetReportSchemaFieldsEnumStringValues Enumerates the set of values in String for ReportSchemaFieldsEnum
func GetReportSchemaFieldsEnumStringValues() []string {
	return []string{
		"SIZE",
		"ETAG",
		"TIME_CREATED",
		"MD5_CHECKSUM",
		"TIME_MODIFIED",
		"STORAGE_TIER",
		"ARCHIVAL_STATE",
		"COMPARTMENT_ID",
		"CREATED_BY",
		"DEFINED_TAGS",
		"FREEFORM_TAGS",
	}
}

// GetMappingReportSchemaFieldsEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingReportSchemaFieldsEnum(val string) (ReportSchemaFieldsEnum, bool) {
	enum, ok := mappingReportSchemaFieldsEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
