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
	"github.com/oracle/oci-go-sdk/v56/common"
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

var mappingQueryableFieldSummaryFieldType = map[string]QueryableFieldSummaryFieldTypeEnum{
	"STRING":    QueryableFieldSummaryFieldTypeString,
	"INTEGER":   QueryableFieldSummaryFieldTypeInteger,
	"BOOLEAN":   QueryableFieldSummaryFieldTypeBoolean,
	"DATE_TIME": QueryableFieldSummaryFieldTypeDateTime,
	"OBJECT":    QueryableFieldSummaryFieldTypeObject,
}

// GetQueryableFieldSummaryFieldTypeEnumValues Enumerates the set of values for QueryableFieldSummaryFieldTypeEnum
func GetQueryableFieldSummaryFieldTypeEnumValues() []QueryableFieldSummaryFieldTypeEnum {
	values := make([]QueryableFieldSummaryFieldTypeEnum, 0)
	for _, v := range mappingQueryableFieldSummaryFieldType {
		values = append(values, v)
	}
	return values
}
