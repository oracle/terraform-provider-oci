// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Search Service API
//
// Search for resources in your cloud network.
//

package resourcesearch

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// QueryableFieldDescription An individual field that can be used as part of a query filter.
type QueryableFieldDescription struct {

	// The type of the field, which dictates what semantics and query constraints you can use when searching or querying.
	FieldType QueryableFieldDescriptionFieldTypeEnum `mandatory:"true" json:"fieldType"`

	// The name of the field to use when constructing the query. Field names are present for all types except `OBJECT`.
	FieldName *string `mandatory:"true" json:"fieldName"`

	// Indicates that this field is actually an array of the specified field type.
	IsArray *bool `mandatory:"false" json:"isArray"`

	// If the field type is `OBJECT`, then this property will provide all the individual properties of the object that can
	// be queried.
	ObjectProperties []QueryableFieldDescription `mandatory:"false" json:"objectProperties"`
}

func (m QueryableFieldDescription) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m QueryableFieldDescription) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingQueryableFieldDescriptionFieldTypeEnum(string(m.FieldType)); !ok && m.FieldType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for FieldType: %s. Supported values are: %s.", m.FieldType, strings.Join(GetQueryableFieldDescriptionFieldTypeEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// QueryableFieldDescriptionFieldTypeEnum Enum with underlying type: string
type QueryableFieldDescriptionFieldTypeEnum string

// Set of constants representing the allowable values for QueryableFieldDescriptionFieldTypeEnum
const (
	QueryableFieldDescriptionFieldTypeIdentifier QueryableFieldDescriptionFieldTypeEnum = "IDENTIFIER"
	QueryableFieldDescriptionFieldTypeString     QueryableFieldDescriptionFieldTypeEnum = "STRING"
	QueryableFieldDescriptionFieldTypeInteger    QueryableFieldDescriptionFieldTypeEnum = "INTEGER"
	QueryableFieldDescriptionFieldTypeRational   QueryableFieldDescriptionFieldTypeEnum = "RATIONAL"
	QueryableFieldDescriptionFieldTypeBoolean    QueryableFieldDescriptionFieldTypeEnum = "BOOLEAN"
	QueryableFieldDescriptionFieldTypeDatetime   QueryableFieldDescriptionFieldTypeEnum = "DATETIME"
	QueryableFieldDescriptionFieldTypeIp         QueryableFieldDescriptionFieldTypeEnum = "IP"
	QueryableFieldDescriptionFieldTypeObject     QueryableFieldDescriptionFieldTypeEnum = "OBJECT"
)

var mappingQueryableFieldDescriptionFieldTypeEnum = map[string]QueryableFieldDescriptionFieldTypeEnum{
	"IDENTIFIER": QueryableFieldDescriptionFieldTypeIdentifier,
	"STRING":     QueryableFieldDescriptionFieldTypeString,
	"INTEGER":    QueryableFieldDescriptionFieldTypeInteger,
	"RATIONAL":   QueryableFieldDescriptionFieldTypeRational,
	"BOOLEAN":    QueryableFieldDescriptionFieldTypeBoolean,
	"DATETIME":   QueryableFieldDescriptionFieldTypeDatetime,
	"IP":         QueryableFieldDescriptionFieldTypeIp,
	"OBJECT":     QueryableFieldDescriptionFieldTypeObject,
}

var mappingQueryableFieldDescriptionFieldTypeEnumLowerCase = map[string]QueryableFieldDescriptionFieldTypeEnum{
	"identifier": QueryableFieldDescriptionFieldTypeIdentifier,
	"string":     QueryableFieldDescriptionFieldTypeString,
	"integer":    QueryableFieldDescriptionFieldTypeInteger,
	"rational":   QueryableFieldDescriptionFieldTypeRational,
	"boolean":    QueryableFieldDescriptionFieldTypeBoolean,
	"datetime":   QueryableFieldDescriptionFieldTypeDatetime,
	"ip":         QueryableFieldDescriptionFieldTypeIp,
	"object":     QueryableFieldDescriptionFieldTypeObject,
}

// GetQueryableFieldDescriptionFieldTypeEnumValues Enumerates the set of values for QueryableFieldDescriptionFieldTypeEnum
func GetQueryableFieldDescriptionFieldTypeEnumValues() []QueryableFieldDescriptionFieldTypeEnum {
	values := make([]QueryableFieldDescriptionFieldTypeEnum, 0)
	for _, v := range mappingQueryableFieldDescriptionFieldTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetQueryableFieldDescriptionFieldTypeEnumStringValues Enumerates the set of values in String for QueryableFieldDescriptionFieldTypeEnum
func GetQueryableFieldDescriptionFieldTypeEnumStringValues() []string {
	return []string{
		"IDENTIFIER",
		"STRING",
		"INTEGER",
		"RATIONAL",
		"BOOLEAN",
		"DATETIME",
		"IP",
		"OBJECT",
	}
}

// GetMappingQueryableFieldDescriptionFieldTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingQueryableFieldDescriptionFieldTypeEnum(val string) (QueryableFieldDescriptionFieldTypeEnum, bool) {
	enum, ok := mappingQueryableFieldDescriptionFieldTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
