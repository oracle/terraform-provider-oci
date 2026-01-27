// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// MySQL Database Service API
//
// The API for the MySQL Database Service
//

package mysql

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// ExportOptions The additional options used while exporting the DB system backup.
type ExportOptions struct {
	CompressionOptions *CompressionOptions `mandatory:"false" json:"compressionOptions"`

	// The format used for storing data.
	DataFormat ExportOptionsDataFormatEnum `mandatory:"false" json:"dataFormat,omitempty"`

	// The name of the folder in the Object Storage bucket where the dump files will be stored.
	// A folder with the same name must not exist in the bucket. The folder will be created in the export process.
	FolderName *string `mandatory:"false" json:"folderName"`
}

func (m ExportOptions) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ExportOptions) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingExportOptionsDataFormatEnum(string(m.DataFormat)); !ok && m.DataFormat != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for DataFormat: %s. Supported values are: %s.", m.DataFormat, strings.Join(GetExportOptionsDataFormatEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ExportOptionsDataFormatEnum Enum with underlying type: string
type ExportOptionsDataFormatEnum string

// Set of constants representing the allowable values for ExportOptionsDataFormatEnum
const (
	ExportOptionsDataFormatCsv ExportOptionsDataFormatEnum = "CSV"
	ExportOptionsDataFormatTsv ExportOptionsDataFormatEnum = "TSV"
)

var mappingExportOptionsDataFormatEnum = map[string]ExportOptionsDataFormatEnum{
	"CSV": ExportOptionsDataFormatCsv,
	"TSV": ExportOptionsDataFormatTsv,
}

var mappingExportOptionsDataFormatEnumLowerCase = map[string]ExportOptionsDataFormatEnum{
	"csv": ExportOptionsDataFormatCsv,
	"tsv": ExportOptionsDataFormatTsv,
}

// GetExportOptionsDataFormatEnumValues Enumerates the set of values for ExportOptionsDataFormatEnum
func GetExportOptionsDataFormatEnumValues() []ExportOptionsDataFormatEnum {
	values := make([]ExportOptionsDataFormatEnum, 0)
	for _, v := range mappingExportOptionsDataFormatEnum {
		values = append(values, v)
	}
	return values
}

// GetExportOptionsDataFormatEnumStringValues Enumerates the set of values in String for ExportOptionsDataFormatEnum
func GetExportOptionsDataFormatEnumStringValues() []string {
	return []string{
		"CSV",
		"TSV",
	}
}

// GetMappingExportOptionsDataFormatEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingExportOptionsDataFormatEnum(val string) (ExportOptionsDataFormatEnum, bool) {
	enum, ok := mappingExportOptionsDataFormatEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
