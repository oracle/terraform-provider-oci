// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Operations Insights API
//
// Use the Operations Insights API to perform data extraction operations to obtain database
// resource utilization, performance statistics, and reference information. For more information,
// see About Oracle Cloud Infrastructure Operations Insights (https://docs.cloud.oracle.com/en-us/iaas/operations-insights/doc/operations-insights.html).
//

package opsi

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// QueryDataObjectResultSetColumnMetadata Metadata of a column in a data object query result set.
type QueryDataObjectResultSetColumnMetadata struct {

	// Name of the column in a data object query result set.
	Name *string `mandatory:"true" json:"name"`

	// Type of the column in a data object query result.
	DataType *string `mandatory:"false" json:"dataType"`

	// Type name of the column in a data object query result set.
	DataTypeName QueryDataObjectResultSetColumnMetadataDataTypeNameEnum `mandatory:"false" json:"dataTypeName,omitempty"`
}

func (m QueryDataObjectResultSetColumnMetadata) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m QueryDataObjectResultSetColumnMetadata) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingQueryDataObjectResultSetColumnMetadataDataTypeNameEnum(string(m.DataTypeName)); !ok && m.DataTypeName != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for DataTypeName: %s. Supported values are: %s.", m.DataTypeName, strings.Join(GetQueryDataObjectResultSetColumnMetadataDataTypeNameEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// QueryDataObjectResultSetColumnMetadataDataTypeNameEnum Enum with underlying type: string
type QueryDataObjectResultSetColumnMetadataDataTypeNameEnum string

// Set of constants representing the allowable values for QueryDataObjectResultSetColumnMetadataDataTypeNameEnum
const (
	QueryDataObjectResultSetColumnMetadataDataTypeNameNumber    QueryDataObjectResultSetColumnMetadataDataTypeNameEnum = "NUMBER"
	QueryDataObjectResultSetColumnMetadataDataTypeNameTimestamp QueryDataObjectResultSetColumnMetadataDataTypeNameEnum = "TIMESTAMP"
	QueryDataObjectResultSetColumnMetadataDataTypeNameVarchar2  QueryDataObjectResultSetColumnMetadataDataTypeNameEnum = "VARCHAR2"
	QueryDataObjectResultSetColumnMetadataDataTypeNameOther     QueryDataObjectResultSetColumnMetadataDataTypeNameEnum = "OTHER"
)

var mappingQueryDataObjectResultSetColumnMetadataDataTypeNameEnum = map[string]QueryDataObjectResultSetColumnMetadataDataTypeNameEnum{
	"NUMBER":    QueryDataObjectResultSetColumnMetadataDataTypeNameNumber,
	"TIMESTAMP": QueryDataObjectResultSetColumnMetadataDataTypeNameTimestamp,
	"VARCHAR2":  QueryDataObjectResultSetColumnMetadataDataTypeNameVarchar2,
	"OTHER":     QueryDataObjectResultSetColumnMetadataDataTypeNameOther,
}

var mappingQueryDataObjectResultSetColumnMetadataDataTypeNameEnumLowerCase = map[string]QueryDataObjectResultSetColumnMetadataDataTypeNameEnum{
	"number":    QueryDataObjectResultSetColumnMetadataDataTypeNameNumber,
	"timestamp": QueryDataObjectResultSetColumnMetadataDataTypeNameTimestamp,
	"varchar2":  QueryDataObjectResultSetColumnMetadataDataTypeNameVarchar2,
	"other":     QueryDataObjectResultSetColumnMetadataDataTypeNameOther,
}

// GetQueryDataObjectResultSetColumnMetadataDataTypeNameEnumValues Enumerates the set of values for QueryDataObjectResultSetColumnMetadataDataTypeNameEnum
func GetQueryDataObjectResultSetColumnMetadataDataTypeNameEnumValues() []QueryDataObjectResultSetColumnMetadataDataTypeNameEnum {
	values := make([]QueryDataObjectResultSetColumnMetadataDataTypeNameEnum, 0)
	for _, v := range mappingQueryDataObjectResultSetColumnMetadataDataTypeNameEnum {
		values = append(values, v)
	}
	return values
}

// GetQueryDataObjectResultSetColumnMetadataDataTypeNameEnumStringValues Enumerates the set of values in String for QueryDataObjectResultSetColumnMetadataDataTypeNameEnum
func GetQueryDataObjectResultSetColumnMetadataDataTypeNameEnumStringValues() []string {
	return []string{
		"NUMBER",
		"TIMESTAMP",
		"VARCHAR2",
		"OTHER",
	}
}

// GetMappingQueryDataObjectResultSetColumnMetadataDataTypeNameEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingQueryDataObjectResultSetColumnMetadataDataTypeNameEnum(val string) (QueryDataObjectResultSetColumnMetadataDataTypeNameEnum, bool) {
	enum, ok := mappingQueryDataObjectResultSetColumnMetadataDataTypeNameEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
