// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// LogAnalytics API
//
// The LogAnalytics API for the LogAnalytics service.
//

package loganalytics

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// LogAnalyticsSourceExtendedFieldDefinition LogAnalyticsSourceExtendedFieldDefinition
type LogAnalyticsSourceExtendedFieldDefinition struct {
	Field *LogAnalyticsField `mandatory:"false" json:"field"`

	// The regular expression.
	DisplayRegularExpression *string `mandatory:"false" json:"displayRegularExpression"`

	// An array of extended fields.
	ExtendedFields []LogAnalyticsExtendedField `mandatory:"false" json:"extendedFields"`

	// The base field internal name.
	BaseFieldName *string `mandatory:"false" json:"baseFieldName"`

	// The base field log text.
	BaseFieldLogText *string `mandatory:"false" json:"baseFieldLogText"`

	// The conditional data type.
	ConditionDataType *string `mandatory:"false" json:"conditionDataType"`

	// The onditional field.
	ConditionField *string `mandatory:"false" json:"conditionField"`

	// The conditional operator.
	ConditionOperator *string `mandatory:"false" json:"conditionOperator"`

	// The conditional value.
	ConditionValue *string `mandatory:"false" json:"conditionValue"`

	// The converted regular expression.
	ConvertedRegularExpression *string `mandatory:"false" json:"convertedRegularExpression"`

	// A flag inidcating whether or not the extended definition is enabled.
	IsEnabled *bool `mandatory:"false" json:"isEnabled"`

	// The extended field definition unique identifier.
	ExtendedFieldDefinitionId *int64 `mandatory:"false" json:"extendedFieldDefinitionId"`

	// The system flag.  A value of false denotes a custom, or user
	// defined object.  A value of true denotes a built in object.
	IsSystem *bool `mandatory:"false" json:"isSystem"`

	// The regular expression.
	RegularExpression *string `mandatory:"false" json:"regularExpression"`

	// The source unique identifier.
	SourceId *int64 `mandatory:"false" json:"sourceId"`

	// The last updated date.
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`
}

func (m LogAnalyticsSourceExtendedFieldDefinition) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m LogAnalyticsSourceExtendedFieldDefinition) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
