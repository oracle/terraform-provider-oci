// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// LogAnalytics API
//
// The LogAnalytics API for the LogAnalytics service.
//

package loganalytics

import (
	"github.com/oracle/oci-go-sdk/v30/common"
)

// LogAnalyticsLabelAlias Label alias mapping view
type LogAnalyticsLabelAlias struct {

	// alias
	Alias *string `mandatory:"false" json:"alias"`

	// alias display name
	AliasDisplayName *string `mandatory:"false" json:"aliasDisplayName"`

	// is system flag
	IsSystem *bool `mandatory:"false" json:"isSystem"`

	// label display name
	DisplayName *string `mandatory:"false" json:"displayName"`

	// label name
	Name *string `mandatory:"false" json:"name"`

	// priority
	Priority LogAnalyticsLabelAliasPriorityEnum `mandatory:"false" json:"priority,omitempty"`
}

func (m LogAnalyticsLabelAlias) String() string {
	return common.PointerString(m)
}

// LogAnalyticsLabelAliasPriorityEnum Enum with underlying type: string
type LogAnalyticsLabelAliasPriorityEnum string

// Set of constants representing the allowable values for LogAnalyticsLabelAliasPriorityEnum
const (
	LogAnalyticsLabelAliasPriorityNone   LogAnalyticsLabelAliasPriorityEnum = "NONE"
	LogAnalyticsLabelAliasPriorityLow    LogAnalyticsLabelAliasPriorityEnum = "LOW"
	LogAnalyticsLabelAliasPriorityMedium LogAnalyticsLabelAliasPriorityEnum = "MEDIUM"
	LogAnalyticsLabelAliasPriorityHigh   LogAnalyticsLabelAliasPriorityEnum = "HIGH"
)

var mappingLogAnalyticsLabelAliasPriority = map[string]LogAnalyticsLabelAliasPriorityEnum{
	"NONE":   LogAnalyticsLabelAliasPriorityNone,
	"LOW":    LogAnalyticsLabelAliasPriorityLow,
	"MEDIUM": LogAnalyticsLabelAliasPriorityMedium,
	"HIGH":   LogAnalyticsLabelAliasPriorityHigh,
}

// GetLogAnalyticsLabelAliasPriorityEnumValues Enumerates the set of values for LogAnalyticsLabelAliasPriorityEnum
func GetLogAnalyticsLabelAliasPriorityEnumValues() []LogAnalyticsLabelAliasPriorityEnum {
	values := make([]LogAnalyticsLabelAliasPriorityEnum, 0)
	for _, v := range mappingLogAnalyticsLabelAliasPriority {
		values = append(values, v)
	}
	return values
}
