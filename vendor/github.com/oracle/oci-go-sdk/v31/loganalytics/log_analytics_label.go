// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// LogAnalytics API
//
// The LogAnalytics API for the LogAnalytics service.
//

package loganalytics

import (
	"github.com/oracle/oci-go-sdk/v31/common"
)

// LogAnalyticsLabel LogAnalytics label
type LogAnalyticsLabel struct {

	// alias list
	Aliases []LogAnalyticsLabelAlias `mandatory:"false" json:"aliases"`

	// count usage in source
	CountUsageInSource *int64 `mandatory:"false" json:"countUsageInSource"`

	// suggest type
	SuggestType *int64 `mandatory:"false" json:"suggestType"`

	// description
	Description *string `mandatory:"false" json:"description"`

	// display name
	DisplayName *string `mandatory:"false" json:"displayName"`

	// edit version
	EditVersion *int64 `mandatory:"false" json:"editVersion"`

	// impact
	Impact *string `mandatory:"false" json:"impact"`

	// is system flag
	IsSystem *bool `mandatory:"false" json:"isSystem"`

	// label identifier
	Name *string `mandatory:"false" json:"name"`

	// Valid values are (NONE, LOW, HIGH). NONE is default.
	Priority LogAnalyticsLabelPriorityEnum `mandatory:"false" json:"priority,omitempty"`

	// tag recommendation
	Recommendation *string `mandatory:"false" json:"recommendation"`

	// Valid values are (INFO, PROBLEM). INFO is default.
	Type LogAnalyticsLabelTypeEnum `mandatory:"false" json:"type,omitempty"`

	// user deleted flag
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
