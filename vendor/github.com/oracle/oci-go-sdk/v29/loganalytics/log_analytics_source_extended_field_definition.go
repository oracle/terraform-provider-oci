// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// LogAnalytics API
//
// The LogAnalytics API for the LogAnalytics service.
//

package loganalytics

import (
	"github.com/oracle/oci-go-sdk/v29/common"
)

// LogAnalyticsSourceExtendedFieldDefinition LogAnalyticsSourceExtendedFieldDefinition
type LogAnalyticsSourceExtendedFieldDefinition struct {
	Field *LogAnalyticsField `mandatory:"false" json:"field"`

	// display regular expression
	DisplayRegularExpression *string `mandatory:"false" json:"displayRegularExpression"`

	// extended fields
	ExtendedFields []LogAnalyticsExtendedField `mandatory:"false" json:"extendedFields"`

	// base field internal name
	BaseFieldName *string `mandatory:"false" json:"baseFieldName"`

	// base field log text
	BaseFieldLogText *string `mandatory:"false" json:"baseFieldLogText"`

	// conditional data type
	ConditionDataType *string `mandatory:"false" json:"conditionDataType"`

	// conditional field
	ConditionField *string `mandatory:"false" json:"conditionField"`

	// conditional operator
	ConditionOperator *string `mandatory:"false" json:"conditionOperator"`

	// conditional value
	ConditionValue *string `mandatory:"false" json:"conditionValue"`

	// converted regular expression
	ConvertedRegularExpression *string `mandatory:"false" json:"convertedRegularExpression"`

	// enabled
	IsEnabled *bool `mandatory:"false" json:"isEnabled"`

	// id
	ExtendedFieldDefinitionId *int64 `mandatory:"false" json:"extendedFieldDefinitionId"`

	// is system flag
	IsSystem *bool `mandatory:"false" json:"isSystem"`

	// regular expression
	RegularExpression *string `mandatory:"false" json:"regularExpression"`

	// source Id
	SourceId *int64 `mandatory:"false" json:"sourceId"`

	// last updated date
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`
}

func (m LogAnalyticsSourceExtendedFieldDefinition) String() string {
	return common.PointerString(m)
}
