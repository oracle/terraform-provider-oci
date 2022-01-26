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

// LogAnalyticsLabelView LogAnalyticsLabelView
type LogAnalyticsLabelView struct {

	// An arrya of label aliases.
	Aliases []LogAnalyticsLabelAlias `mandatory:"false" json:"aliases"`

	// The label alert rule usage count.
	CountUsageInAlertRule *int64 `mandatory:"false" json:"countUsageInAlertRule"`

	// The label source usage count.
	CountUsageInSource *int64 `mandatory:"false" json:"countUsageInSource"`

	// The label unique identifier.
	Id *interface{} `mandatory:"false" json:"id"`

	// The label suggestion type.
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
	// defined object.  A value of true denotes a built in object.
	IsSystem *bool `mandatory:"false" json:"isSystem"`

	// The label name.
	Name *string `mandatory:"false" json:"name"`

	// The label priority.  Default value is NONE.
	Priority LogAnalyticsLabelViewPriorityEnum `mandatory:"false" json:"priority,omitempty"`

	// The label recommendation.
	Recommendation *string `mandatory:"false" json:"recommendation"`

	// The label type.
	Type *int64 `mandatory:"false" json:"type"`

	// A flag indicating whether or not the label has been deleted.
	IsUserDeleted *bool `mandatory:"false" json:"isUserDeleted"`
}

func (m LogAnalyticsLabelView) String() string {
	return common.PointerString(m)
}

// LogAnalyticsLabelViewPriorityEnum Enum with underlying type: string
type LogAnalyticsLabelViewPriorityEnum string

// Set of constants representing the allowable values for LogAnalyticsLabelViewPriorityEnum
const (
	LogAnalyticsLabelViewPriorityNone   LogAnalyticsLabelViewPriorityEnum = "NONE"
	LogAnalyticsLabelViewPriorityLow    LogAnalyticsLabelViewPriorityEnum = "LOW"
	LogAnalyticsLabelViewPriorityMedium LogAnalyticsLabelViewPriorityEnum = "MEDIUM"
	LogAnalyticsLabelViewPriorityHigh   LogAnalyticsLabelViewPriorityEnum = "HIGH"
)

var mappingLogAnalyticsLabelViewPriority = map[string]LogAnalyticsLabelViewPriorityEnum{
	"NONE":   LogAnalyticsLabelViewPriorityNone,
	"LOW":    LogAnalyticsLabelViewPriorityLow,
	"MEDIUM": LogAnalyticsLabelViewPriorityMedium,
	"HIGH":   LogAnalyticsLabelViewPriorityHigh,
}

// GetLogAnalyticsLabelViewPriorityEnumValues Enumerates the set of values for LogAnalyticsLabelViewPriorityEnum
func GetLogAnalyticsLabelViewPriorityEnumValues() []LogAnalyticsLabelViewPriorityEnum {
	values := make([]LogAnalyticsLabelViewPriorityEnum, 0)
	for _, v := range mappingLogAnalyticsLabelViewPriority {
		values = append(values, v)
	}
	return values
}
