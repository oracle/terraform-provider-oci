// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Object Storage Service API
//
// Use Object Storage and Archive Storage APIs to manage buckets, objects, and related resources.
// For more information, see Overview of Object Storage (https://docs.cloud.oracle.com/Content/Object/Concepts/objectstorageoverview.htm) and
// Overview of Archive Storage (https://docs.cloud.oracle.com/Content/Archive/Concepts/archivestorageoverview.htm).
//

package objectstorage

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// ReadObjectSchemaDetails The parameters required by Object Storage to process a read schema request on an object.
// To use any of the API operations, you must be authorized in an IAM policy. If you are not authorized,
// talk to an administrator. If you are an administrator who needs to write policies to give users access, see
// Getting Started with Policies (https://docs.cloud.oracle.com/Content/Identity/Concepts/policygetstarted.htm).
type ReadObjectSchemaDetails struct {

	// field to specify objectName for which we get the schema.
	ObjectName *string `mandatory:"true" json:"objectName"`

	// Optional field to specify the version ID of the object.
	VersionId *string `mandatory:"false" json:"versionId"`

	// Optional field to specify the data format of the target object. By default, this is autodetected.
	DataFormat ReadObjectSchemaDetailsDataFormatEnum `mandatory:"false" json:"dataFormat,omitempty"`
}

func (m ReadObjectSchemaDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ReadObjectSchemaDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingReadObjectSchemaDetailsDataFormatEnum(string(m.DataFormat)); !ok && m.DataFormat != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for DataFormat: %s. Supported values are: %s.", m.DataFormat, strings.Join(GetReadObjectSchemaDetailsDataFormatEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ReadObjectSchemaDetailsDataFormatEnum Enum with underlying type: string
type ReadObjectSchemaDetailsDataFormatEnum string

// Set of constants representing the allowable values for ReadObjectSchemaDetailsDataFormatEnum
const (
	ReadObjectSchemaDetailsDataFormatCsv     ReadObjectSchemaDetailsDataFormatEnum = "CSV"
	ReadObjectSchemaDetailsDataFormatJson    ReadObjectSchemaDetailsDataFormatEnum = "JSON"
	ReadObjectSchemaDetailsDataFormatParquet ReadObjectSchemaDetailsDataFormatEnum = "PARQUET"
	ReadObjectSchemaDetailsDataFormatOrc     ReadObjectSchemaDetailsDataFormatEnum = "ORC"
)

var mappingReadObjectSchemaDetailsDataFormatEnum = map[string]ReadObjectSchemaDetailsDataFormatEnum{
	"CSV":     ReadObjectSchemaDetailsDataFormatCsv,
	"JSON":    ReadObjectSchemaDetailsDataFormatJson,
	"PARQUET": ReadObjectSchemaDetailsDataFormatParquet,
	"ORC":     ReadObjectSchemaDetailsDataFormatOrc,
}

var mappingReadObjectSchemaDetailsDataFormatEnumLowerCase = map[string]ReadObjectSchemaDetailsDataFormatEnum{
	"csv":     ReadObjectSchemaDetailsDataFormatCsv,
	"json":    ReadObjectSchemaDetailsDataFormatJson,
	"parquet": ReadObjectSchemaDetailsDataFormatParquet,
	"orc":     ReadObjectSchemaDetailsDataFormatOrc,
}

// GetReadObjectSchemaDetailsDataFormatEnumValues Enumerates the set of values for ReadObjectSchemaDetailsDataFormatEnum
func GetReadObjectSchemaDetailsDataFormatEnumValues() []ReadObjectSchemaDetailsDataFormatEnum {
	values := make([]ReadObjectSchemaDetailsDataFormatEnum, 0)
	for _, v := range mappingReadObjectSchemaDetailsDataFormatEnum {
		values = append(values, v)
	}
	return values
}

// GetReadObjectSchemaDetailsDataFormatEnumStringValues Enumerates the set of values in String for ReadObjectSchemaDetailsDataFormatEnum
func GetReadObjectSchemaDetailsDataFormatEnumStringValues() []string {
	return []string{
		"CSV",
		"JSON",
		"PARQUET",
		"ORC",
	}
}

// GetMappingReadObjectSchemaDetailsDataFormatEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingReadObjectSchemaDetailsDataFormatEnum(val string) (ReadObjectSchemaDetailsDataFormatEnum, bool) {
	enum, ok := mappingReadObjectSchemaDetailsDataFormatEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
