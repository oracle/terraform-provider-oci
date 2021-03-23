// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// LogAnalytics API
//
// The LogAnalytics API for the LogAnalytics service.
//

package loganalytics

import (
	"github.com/oracle/oci-go-sdk/v37/common"
)

// LogAnalyticsLabelView LogAnalyticsLabelView
type LogAnalyticsLabelView struct {

	// alias list
	Aliases []LogAnalyticsLabelAlias `mandatory:"false" json:"aliases"`

	// alert rule usage count
	CountUsageInAlertRule *int64 `mandatory:"false" json:"countUsageInAlertRule"`

	// source usage count
	CountUsageInSource *int64 `mandatory:"false" json:"countUsageInSource"`

	// id
	Id *interface{} `mandatory:"false" json:"id"`

	// suggest type
	SuggestType *int64 `mandatory:"false" json:"suggestType"`

	// label description
	Description *string `mandatory:"false" json:"description"`

	// label display name
	DisplayName *string `mandatory:"false" json:"displayName"`

	// tag edit version
	EditVersion *int64 `mandatory:"false" json:"editVersion"`

	// label impact
	Impact *string `mandatory:"false" json:"impact"`

	// is system flag
	IsSystem *bool `mandatory:"false" json:"isSystem"`

	// label name
	Name *string `mandatory:"false" json:"name"`

	// priority
	Priority LogAnalyticsLabelViewPriorityEnum `mandatory:"false" json:"priority,omitempty"`

	// recommendation
	Recommendation *string `mandatory:"false" json:"recommendation"`

	// type
	Type *int64 `mandatory:"false" json:"type"`

	// user deleted flag
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
