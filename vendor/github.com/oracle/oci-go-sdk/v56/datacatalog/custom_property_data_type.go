// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Catalog API
//
// Use the Data Catalog APIs to collect, organize, find, access, understand, enrich, and activate technical, business, and operational metadata.
// For more information, see Data Catalog (https://docs.oracle.com/iaas/data-catalog/home.htm).
//

package datacatalog

// CustomPropertyDataTypeEnum Enum with underlying type: string
type CustomPropertyDataTypeEnum string

// Set of constants representing the allowable values for CustomPropertyDataTypeEnum
const (
	CustomPropertyDataTypeText     CustomPropertyDataTypeEnum = "TEXT"
	CustomPropertyDataTypeRichText CustomPropertyDataTypeEnum = "RICH_TEXT"
	CustomPropertyDataTypeBoolean  CustomPropertyDataTypeEnum = "BOOLEAN"
	CustomPropertyDataTypeNumber   CustomPropertyDataTypeEnum = "NUMBER"
	CustomPropertyDataTypeDate     CustomPropertyDataTypeEnum = "DATE"
)

var mappingCustomPropertyDataType = map[string]CustomPropertyDataTypeEnum{
	"TEXT":      CustomPropertyDataTypeText,
	"RICH_TEXT": CustomPropertyDataTypeRichText,
	"BOOLEAN":   CustomPropertyDataTypeBoolean,
	"NUMBER":    CustomPropertyDataTypeNumber,
	"DATE":      CustomPropertyDataTypeDate,
}

// GetCustomPropertyDataTypeEnumValues Enumerates the set of values for CustomPropertyDataTypeEnum
func GetCustomPropertyDataTypeEnumValues() []CustomPropertyDataTypeEnum {
	values := make([]CustomPropertyDataTypeEnum, 0)
	for _, v := range mappingCustomPropertyDataType {
		values = append(values, v)
	}
	return values
}
