// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Cloud Advisor API
//
// Use the Cloud Advisor API to find potential inefficiencies in your tenancy and address them.
// Cloud Advisor can help you save money, improve performance, strengthen system resilience, and improve security.
// For more information, see Cloud Advisor (https://docs.cloud.oracle.com/Content/CloudAdvisor/Concepts/cloudadvisoroverview.htm).
//

package optimizer

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v58/common"
	"strings"
)

// QueryableFieldSummary An individual field that can be used as part of a query filter.
type QueryableFieldSummary struct {

	// The type of the field, which dictates the semantics and query constraints that you can use when searching or querying.
	FieldType QueryableFieldSummaryFieldTypeEnum `mandatory:"true" json:"fieldType"`

	// The name of the field to use when constructing the query. Field names are present for all types except `OBJECT`.
	FieldName *string `mandatory:"true" json:"fieldName"`

	// If the field type is `OBJECT`, this property lists the individual properties of the object that can be queried.
	ObjectProperties []QueryableFieldSummary `mandatory:"false" json:"objectProperties"`
}

func (m QueryableFieldSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m QueryableFieldSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingQueryableFieldSummaryFieldTypeEnum(string(m.FieldType)); !ok && m.FieldType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for FieldType: %s. Supported values are: %s.", m.FieldType, strings.Join(GetQueryableFieldSummaryFieldTypeEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// QueryableFieldSummaryFieldTypeEnum Enum with underlying type: string
type QueryableFieldSummaryFieldTypeEnum string

// Set of constants representing the allowable values for QueryableFieldSummaryFieldTypeEnum
const (
	QueryableFieldSummaryFieldTypeString   QueryableFieldSummaryFieldTypeEnum = "STRING"
	QueryableFieldSummaryFieldTypeInteger  QueryableFieldSummaryFieldTypeEnum = "INTEGER"
	QueryableFieldSummaryFieldTypeBoolean  QueryableFieldSummaryFieldTypeEnum = "BOOLEAN"
	QueryableFieldSummaryFieldTypeDateTime QueryableFieldSummaryFieldTypeEnum = "DATE_TIME"
	QueryableFieldSummaryFieldTypeObject   QueryableFieldSummaryFieldTypeEnum = "OBJECT"
)

var mappingQueryableFieldSummaryFieldTypeEnum = map[string]QueryableFieldSummaryFieldTypeEnum{
	"STRING":    QueryableFieldSummaryFieldTypeString,
	"INTEGER":   QueryableFieldSummaryFieldTypeInteger,
	"BOOLEAN":   QueryableFieldSummaryFieldTypeBoolean,
	"DATE_TIME": QueryableFieldSummaryFieldTypeDateTime,
	"OBJECT":    QueryableFieldSummaryFieldTypeObject,
}

// GetQueryableFieldSummaryFieldTypeEnumValues Enumerates the set of values for QueryableFieldSummaryFieldTypeEnum
func GetQueryableFieldSummaryFieldTypeEnumValues() []QueryableFieldSummaryFieldTypeEnum {
	values := make([]QueryableFieldSummaryFieldTypeEnum, 0)
	for _, v := range mappingQueryableFieldSummaryFieldTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetQueryableFieldSummaryFieldTypeEnumStringValues Enumerates the set of values in String for QueryableFieldSummaryFieldTypeEnum
func GetQueryableFieldSummaryFieldTypeEnumStringValues() []string {
	return []string{
		"STRING",
		"INTEGER",
		"BOOLEAN",
		"DATE_TIME",
		"OBJECT",
	}
}

// GetMappingQueryableFieldSummaryFieldTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingQueryableFieldSummaryFieldTypeEnum(val string) (QueryableFieldSummaryFieldTypeEnum, bool) {
	mappingQueryableFieldSummaryFieldTypeEnumIgnoreCase := make(map[string]QueryableFieldSummaryFieldTypeEnum)
	for k, v := range mappingQueryableFieldSummaryFieldTypeEnum {
		mappingQueryableFieldSummaryFieldTypeEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingQueryableFieldSummaryFieldTypeEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}
