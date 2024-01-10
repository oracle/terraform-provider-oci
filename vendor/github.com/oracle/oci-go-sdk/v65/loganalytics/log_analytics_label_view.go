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

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m LogAnalyticsLabelView) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingLogAnalyticsLabelViewPriorityEnum(string(m.Priority)); !ok && m.Priority != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Priority: %s. Supported values are: %s.", m.Priority, strings.Join(GetLogAnalyticsLabelViewPriorityEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
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

var mappingLogAnalyticsLabelViewPriorityEnum = map[string]LogAnalyticsLabelViewPriorityEnum{
	"NONE":   LogAnalyticsLabelViewPriorityNone,
	"LOW":    LogAnalyticsLabelViewPriorityLow,
	"MEDIUM": LogAnalyticsLabelViewPriorityMedium,
	"HIGH":   LogAnalyticsLabelViewPriorityHigh,
}

var mappingLogAnalyticsLabelViewPriorityEnumLowerCase = map[string]LogAnalyticsLabelViewPriorityEnum{
	"none":   LogAnalyticsLabelViewPriorityNone,
	"low":    LogAnalyticsLabelViewPriorityLow,
	"medium": LogAnalyticsLabelViewPriorityMedium,
	"high":   LogAnalyticsLabelViewPriorityHigh,
}

// GetLogAnalyticsLabelViewPriorityEnumValues Enumerates the set of values for LogAnalyticsLabelViewPriorityEnum
func GetLogAnalyticsLabelViewPriorityEnumValues() []LogAnalyticsLabelViewPriorityEnum {
	values := make([]LogAnalyticsLabelViewPriorityEnum, 0)
	for _, v := range mappingLogAnalyticsLabelViewPriorityEnum {
		values = append(values, v)
	}
	return values
}

// GetLogAnalyticsLabelViewPriorityEnumStringValues Enumerates the set of values in String for LogAnalyticsLabelViewPriorityEnum
func GetLogAnalyticsLabelViewPriorityEnumStringValues() []string {
	return []string{
		"NONE",
		"LOW",
		"MEDIUM",
		"HIGH",
	}
}

// GetMappingLogAnalyticsLabelViewPriorityEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingLogAnalyticsLabelViewPriorityEnum(val string) (LogAnalyticsLabelViewPriorityEnum, bool) {
	enum, ok := mappingLogAnalyticsLabelViewPriorityEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
