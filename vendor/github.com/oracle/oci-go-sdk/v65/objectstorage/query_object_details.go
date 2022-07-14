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

// QueryObjectDetails The parameters required by Object Storage to process a query request on an object.
// To use any of the API operations, you must be authorized in an IAM policy. If you are not authorized,
// talk to an administrator. If you are an administrator who needs to write policies to give users access, see
// Getting Started with Policies (https://docs.cloud.oracle.com/Content/Identity/Concepts/policygetstarted.htm).
type QueryObjectDetails struct {

	// SQL like query to execute on the file.
	Expression *string `mandatory:"true" json:"expression"`

	// Field to specify the data format of the results. The default ResultFormat is CSV.
	ResultFormat QueryObjectDetailsResultFormatEnum `mandatory:"true" json:"resultFormat"`

	// Optional field to specify the version ID of the object.
	VersionId *string `mandatory:"false" json:"versionId"`

	// Optional field to specify the data format of the target object. By default, this is autodetected.
	DataFormat QueryObjectDetailsDataFormatEnum `mandatory:"false" json:"dataFormat,omitempty"`
}

func (m QueryObjectDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m QueryObjectDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingQueryObjectDetailsResultFormatEnum(string(m.ResultFormat)); !ok && m.ResultFormat != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ResultFormat: %s. Supported values are: %s.", m.ResultFormat, strings.Join(GetQueryObjectDetailsResultFormatEnumStringValues(), ",")))
	}

	if _, ok := GetMappingQueryObjectDetailsDataFormatEnum(string(m.DataFormat)); !ok && m.DataFormat != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for DataFormat: %s. Supported values are: %s.", m.DataFormat, strings.Join(GetQueryObjectDetailsDataFormatEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// QueryObjectDetailsDataFormatEnum Enum with underlying type: string
type QueryObjectDetailsDataFormatEnum string

// Set of constants representing the allowable values for QueryObjectDetailsDataFormatEnum
const (
	QueryObjectDetailsDataFormatCsv     QueryObjectDetailsDataFormatEnum = "CSV"
	QueryObjectDetailsDataFormatJson    QueryObjectDetailsDataFormatEnum = "JSON"
	QueryObjectDetailsDataFormatParquet QueryObjectDetailsDataFormatEnum = "PARQUET"
	QueryObjectDetailsDataFormatOrc     QueryObjectDetailsDataFormatEnum = "ORC"
)

var mappingQueryObjectDetailsDataFormatEnum = map[string]QueryObjectDetailsDataFormatEnum{
	"CSV":     QueryObjectDetailsDataFormatCsv,
	"JSON":    QueryObjectDetailsDataFormatJson,
	"PARQUET": QueryObjectDetailsDataFormatParquet,
	"ORC":     QueryObjectDetailsDataFormatOrc,
}

var mappingQueryObjectDetailsDataFormatEnumLowerCase = map[string]QueryObjectDetailsDataFormatEnum{
	"csv":     QueryObjectDetailsDataFormatCsv,
	"json":    QueryObjectDetailsDataFormatJson,
	"parquet": QueryObjectDetailsDataFormatParquet,
	"orc":     QueryObjectDetailsDataFormatOrc,
}

// GetQueryObjectDetailsDataFormatEnumValues Enumerates the set of values for QueryObjectDetailsDataFormatEnum
func GetQueryObjectDetailsDataFormatEnumValues() []QueryObjectDetailsDataFormatEnum {
	values := make([]QueryObjectDetailsDataFormatEnum, 0)
	for _, v := range mappingQueryObjectDetailsDataFormatEnum {
		values = append(values, v)
	}
	return values
}

// GetQueryObjectDetailsDataFormatEnumStringValues Enumerates the set of values in String for QueryObjectDetailsDataFormatEnum
func GetQueryObjectDetailsDataFormatEnumStringValues() []string {
	return []string{
		"CSV",
		"JSON",
		"PARQUET",
		"ORC",
	}
}

// GetMappingQueryObjectDetailsDataFormatEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingQueryObjectDetailsDataFormatEnum(val string) (QueryObjectDetailsDataFormatEnum, bool) {
	enum, ok := mappingQueryObjectDetailsDataFormatEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// QueryObjectDetailsResultFormatEnum Enum with underlying type: string
type QueryObjectDetailsResultFormatEnum string

// Set of constants representing the allowable values for QueryObjectDetailsResultFormatEnum
const (
	QueryObjectDetailsResultFormatCsv  QueryObjectDetailsResultFormatEnum = "CSV"
	QueryObjectDetailsResultFormatJson QueryObjectDetailsResultFormatEnum = "JSON"
)

var mappingQueryObjectDetailsResultFormatEnum = map[string]QueryObjectDetailsResultFormatEnum{
	"CSV":  QueryObjectDetailsResultFormatCsv,
	"JSON": QueryObjectDetailsResultFormatJson,
}

var mappingQueryObjectDetailsResultFormatEnumLowerCase = map[string]QueryObjectDetailsResultFormatEnum{
	"csv":  QueryObjectDetailsResultFormatCsv,
	"json": QueryObjectDetailsResultFormatJson,
}

// GetQueryObjectDetailsResultFormatEnumValues Enumerates the set of values for QueryObjectDetailsResultFormatEnum
func GetQueryObjectDetailsResultFormatEnumValues() []QueryObjectDetailsResultFormatEnum {
	values := make([]QueryObjectDetailsResultFormatEnum, 0)
	for _, v := range mappingQueryObjectDetailsResultFormatEnum {
		values = append(values, v)
	}
	return values
}

// GetQueryObjectDetailsResultFormatEnumStringValues Enumerates the set of values in String for QueryObjectDetailsResultFormatEnum
func GetQueryObjectDetailsResultFormatEnumStringValues() []string {
	return []string{
		"CSV",
		"JSON",
	}
}

// GetMappingQueryObjectDetailsResultFormatEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingQueryObjectDetailsResultFormatEnum(val string) (QueryObjectDetailsResultFormatEnum, bool) {
	enum, ok := mappingQueryObjectDetailsResultFormatEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
