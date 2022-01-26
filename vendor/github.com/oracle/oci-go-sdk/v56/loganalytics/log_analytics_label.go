// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// LogAnalytics API
//
// The LogAnalytics API for the LogAnalytics service.
//

package loganalytics

import (
	"github.com/oracle/oci-go-sdk/v56/common"
)

// LogAnalyticsLabel LogAnalytics label
type LogAnalyticsLabel struct {

	// An array of label aliases.
	Aliases []LogAnalyticsLabelAlias `mandatory:"false" json:"aliases"`

	// The source usage count for this label.
	CountUsageInSource *int64 `mandatory:"false" json:"countUsageInSource"`

	// The type of suggestion for label usage.
	SuggestType *int64 `mandatory:"false" json:"suggestType"`

	// The label description.
	Description *string `mandatory:"false" json:"description"`

	// The label display name.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// The label edit version.
	EditVersion *int64 `mandatory:"false" json:"editVersion"`

	// The label impact.
	Impact *string `mandatory:"false" json:"impact"`

	// The system flag.  A value of false denotes a custom, or user
	// defined label.  A value of true denotes a built in label.
	IsSystem *bool `mandatory:"false" json:"isSystem"`

	// The label name.
	Name *string `mandatory:"false" json:"name"`

	// The label priority. Valid values are (NONE, LOW, HIGH). NONE is default.
	Priority LogAnalyticsLabelPriorityEnum `mandatory:"false" json:"priority,omitempty"`

	// The label recommendation.
	Recommendation *string `mandatory:"false" json:"recommendation"`

	// The label type. Valid values are (INFO, PROBLEM). INFO is default.
	Type LogAnalyticsLabelTypeEnum `mandatory:"false" json:"type,omitempty"`

	// A flag indicating whether or not the label has been deleted.
	IsUserDeleted *bool `mandatory:"false" json:"isUserDeleted"`
}

func (m LogAnalyticsLabel) String() string {
	return common.PointerString(m)
}

// LogAnalyticsLabelPriorityEnum Enum with underlying type: string
type LogAnalyticsLabelPriorityEnum string

// Set of constants representing the allowable values for LogAnalyticsLabelPriorityEnum
const (
	LogAnalyticsLabelPriorityNone   LogAnalyticsLabelPriorityEnum = "NONE"
	LogAnalyticsLabelPriorityLow    LogAnalyticsLabelPriorityEnum = "LOW"
	LogAnalyticsLabelPriorityMedium LogAnalyticsLabelPriorityEnum = "MEDIUM"
	LogAnalyticsLabelPriorityHigh   LogAnalyticsLabelPriorityEnum = "HIGH"
)

var mappingLogAnalyticsLabelPriority = map[string]LogAnalyticsLabelPriorityEnum{
	"NONE":   LogAnalyticsLabelPriorityNone,
	"LOW":    LogAnalyticsLabelPriorityLow,
	"MEDIUM": LogAnalyticsLabelPriorityMedium,
	"HIGH":   LogAnalyticsLabelPriorityHigh,
}

// GetLogAnalyticsLabelPriorityEnumValues Enumerates the set of values for LogAnalyticsLabelPriorityEnum
func GetLogAnalyticsLabelPriorityEnumValues() []LogAnalyticsLabelPriorityEnum {
	values := make([]LogAnalyticsLabelPriorityEnum, 0)
	for _, v := range mappingLogAnalyticsLabelPriority {
		values = append(values, v)
	}
	return values
}

// LogAnalyticsLabelTypeEnum Enum with underlying type: string
type LogAnalyticsLabelTypeEnum string

// Set of constants representing the allowable values for LogAnalyticsLabelTypeEnum
const (
	LogAnalyticsLabelTypeInfo    LogAnalyticsLabelTypeEnum = "INFO"
	LogAnalyticsLabelTypeProblem LogAnalyticsLabelTypeEnum = "PROBLEM"
)

var mappingLogAnalyticsLabelType = map[string]LogAnalyticsLabelTypeEnum{
	"INFO":    LogAnalyticsLabelTypeInfo,
	"PROBLEM": LogAnalyticsLabelTypeProblem,
}

// GetLogAnalyticsLabelTypeEnumValues Enumerates the set of values for LogAnalyticsLabelTypeEnum
func GetLogAnalyticsLabelTypeEnumValues() []LogAnalyticsLabelTypeEnum {
	values := make([]LogAnalyticsLabelTypeEnum, 0)
	for _, v := range mappingLogAnalyticsLabelType {
		values = append(values, v)
	}
	return values
}
