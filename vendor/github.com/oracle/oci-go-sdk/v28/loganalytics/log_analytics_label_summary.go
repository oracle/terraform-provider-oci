// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// LogAnalytics API
//
// The LogAnalytics API for the LogAnalytics service.
//

package loganalytics

import (
	"github.com/oracle/oci-go-sdk/v28/common"
)

// LogAnalyticsLabelSummary LogAnalytics label
type LogAnalyticsLabelSummary struct {

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
	Priority LogAnalyticsLabelSummaryPriorityEnum `mandatory:"false" json:"priority,omitempty"`

	// tag recommendation
	Recommendation *string `mandatory:"false" json:"recommendation"`

	// Valid values are (INFO, PROBLEM). INFO is default.
	Type LogAnalyticsLabelSummaryTypeEnum `mandatory:"false" json:"type,omitempty"`

	// user deleted flag
	IsUserDeleted *bool `mandatory:"false" json:"isUserDeleted"`
}

func (m LogAnalyticsLabelSummary) String() string {
	return common.PointerString(m)
}

// LogAnalyticsLabelSummaryPriorityEnum Enum with underlying type: string
type LogAnalyticsLabelSummaryPriorityEnum string

// Set of constants representing the allowable values for LogAnalyticsLabelSummaryPriorityEnum
const (
	LogAnalyticsLabelSummaryPriorityNone   LogAnalyticsLabelSummaryPriorityEnum = "NONE"
	LogAnalyticsLabelSummaryPriorityLow    LogAnalyticsLabelSummaryPriorityEnum = "LOW"
	LogAnalyticsLabelSummaryPriorityMedium LogAnalyticsLabelSummaryPriorityEnum = "MEDIUM"
	LogAnalyticsLabelSummaryPriorityHigh   LogAnalyticsLabelSummaryPriorityEnum = "HIGH"
)

var mappingLogAnalyticsLabelSummaryPriority = map[string]LogAnalyticsLabelSummaryPriorityEnum{
	"NONE":   LogAnalyticsLabelSummaryPriorityNone,
	"LOW":    LogAnalyticsLabelSummaryPriorityLow,
	"MEDIUM": LogAnalyticsLabelSummaryPriorityMedium,
	"HIGH":   LogAnalyticsLabelSummaryPriorityHigh,
}

// GetLogAnalyticsLabelSummaryPriorityEnumValues Enumerates the set of values for LogAnalyticsLabelSummaryPriorityEnum
func GetLogAnalyticsLabelSummaryPriorityEnumValues() []LogAnalyticsLabelSummaryPriorityEnum {
	values := make([]LogAnalyticsLabelSummaryPriorityEnum, 0)
	for _, v := range mappingLogAnalyticsLabelSummaryPriority {
		values = append(values, v)
	}
	return values
}

// LogAnalyticsLabelSummaryTypeEnum Enum with underlying type: string
type LogAnalyticsLabelSummaryTypeEnum string

// Set of constants representing the allowable values for LogAnalyticsLabelSummaryTypeEnum
const (
	LogAnalyticsLabelSummaryTypeInfo    LogAnalyticsLabelSummaryTypeEnum = "INFO"
	LogAnalyticsLabelSummaryTypeProblem LogAnalyticsLabelSummaryTypeEnum = "PROBLEM"
)

var mappingLogAnalyticsLabelSummaryType = map[string]LogAnalyticsLabelSummaryTypeEnum{
	"INFO":    LogAnalyticsLabelSummaryTypeInfo,
	"PROBLEM": LogAnalyticsLabelSummaryTypeProblem,
}

// GetLogAnalyticsLabelSummaryTypeEnumValues Enumerates the set of values for LogAnalyticsLabelSummaryTypeEnum
func GetLogAnalyticsLabelSummaryTypeEnumValues() []LogAnalyticsLabelSummaryTypeEnum {
	values := make([]LogAnalyticsLabelSummaryTypeEnum, 0)
	for _, v := range mappingLogAnalyticsLabelSummaryType {
		values = append(values, v)
	}
	return values
}
