// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Safe API
//
// APIs for using Oracle Data Safe.
//

package datasafe

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// ColumnsInfo columnsInfo object has details of column group with schema details.
type ColumnsInfo struct {

	// The schema name.
	SchemaName *string `mandatory:"true" json:"schemaName"`

	// The type of the database object that contains the sensitive column.
	ObjectType ColumnsInfoObjectTypeEnum `mandatory:"true" json:"objectType"`

	// The database object that contains the columns.
	ObjectName *string `mandatory:"true" json:"objectName"`

	// The application name.
	AppName *string `mandatory:"true" json:"appName"`

	// Group of columns in referential relation. Order needs to be maintained in the elements of the parent/child array listing.
	ColumnGroup []string `mandatory:"true" json:"columnGroup"`

	// Sensitive type ocids of each column groups. Order needs to be maintained with the parent column group.
	// For the DB defined referential relations identified during SDM creation, we cannot add sensitive types.
	// Instead use the sensitiveColumn POST API to mark the columns sensitive.
	SensitiveTypeIds []string `mandatory:"false" json:"sensitiveTypeIds"`
}

func (m ColumnsInfo) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ColumnsInfo) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingColumnsInfoObjectTypeEnum(string(m.ObjectType)); !ok && m.ObjectType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ObjectType: %s. Supported values are: %s.", m.ObjectType, strings.Join(GetColumnsInfoObjectTypeEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ColumnsInfoObjectTypeEnum Enum with underlying type: string
type ColumnsInfoObjectTypeEnum string

// Set of constants representing the allowable values for ColumnsInfoObjectTypeEnum
const (
	ColumnsInfoObjectTypeTable ColumnsInfoObjectTypeEnum = "TABLE"
)

var mappingColumnsInfoObjectTypeEnum = map[string]ColumnsInfoObjectTypeEnum{
	"TABLE": ColumnsInfoObjectTypeTable,
}

var mappingColumnsInfoObjectTypeEnumLowerCase = map[string]ColumnsInfoObjectTypeEnum{
	"table": ColumnsInfoObjectTypeTable,
}

// GetColumnsInfoObjectTypeEnumValues Enumerates the set of values for ColumnsInfoObjectTypeEnum
func GetColumnsInfoObjectTypeEnumValues() []ColumnsInfoObjectTypeEnum {
	values := make([]ColumnsInfoObjectTypeEnum, 0)
	for _, v := range mappingColumnsInfoObjectTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetColumnsInfoObjectTypeEnumStringValues Enumerates the set of values in String for ColumnsInfoObjectTypeEnum
func GetColumnsInfoObjectTypeEnumStringValues() []string {
	return []string{
		"TABLE",
	}
}

// GetMappingColumnsInfoObjectTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingColumnsInfoObjectTypeEnum(val string) (ColumnsInfoObjectTypeEnum, bool) {
	enum, ok := mappingColumnsInfoObjectTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
