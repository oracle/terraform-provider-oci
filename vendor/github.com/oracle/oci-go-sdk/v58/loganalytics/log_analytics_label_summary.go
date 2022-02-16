// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// LogAnalytics API
//
// The LogAnalytics API for the LogAnalytics service.
//

package loganalytics

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v58/common"
	"strings"
)

// LogAnalyticsLabelSummary LogAnalytics label
type LogAnalyticsLabelSummary struct {

	// The alias list.
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
	Priority LogAnalyticsLabelSummaryPriorityEnum `mandatory:"false" json:"priority,omitempty"`

	// The label recommendation.
	Recommendation *string `mandatory:"false" json:"recommendation"`

	// The label type.  Valid values are (INFO, PROBLEM). INFO is default.
	Type LogAnalyticsLabelSummaryTypeEnum `mandatory:"false" json:"type,omitempty"`

	// A flag indicating whether or not the label has been deleted.
	IsUserDeleted *bool `mandatory:"false" json:"isUserDeleted"`
}

func (m LogAnalyticsLabelSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m LogAnalyticsLabelSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingLogAnalyticsLabelSummaryPriorityEnum(string(m.Priority)); !ok && m.Priority != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Priority: %s. Supported values are: %s.", m.Priority, strings.Join(GetLogAnalyticsLabelSummaryPriorityEnumStringValues(), ",")))
	}
	if _, ok := GetMappingLogAnalyticsLabelSummaryTypeEnum(string(m.Type)); !ok && m.Type != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Type: %s. Supported values are: %s.", m.Type, strings.Join(GetLogAnalyticsLabelSummaryTypeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
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

var mappingLogAnalyticsLabelSummaryPriorityEnum = map[string]LogAnalyticsLabelSummaryPriorityEnum{
	"NONE":   LogAnalyticsLabelSummaryPriorityNone,
	"LOW":    LogAnalyticsLabelSummaryPriorityLow,
	"MEDIUM": LogAnalyticsLabelSummaryPriorityMedium,
	"HIGH":   LogAnalyticsLabelSummaryPriorityHigh,
}

// GetLogAnalyticsLabelSummaryPriorityEnumValues Enumerates the set of values for LogAnalyticsLabelSummaryPriorityEnum
func GetLogAnalyticsLabelSummaryPriorityEnumValues() []LogAnalyticsLabelSummaryPriorityEnum {
	values := make([]LogAnalyticsLabelSummaryPriorityEnum, 0)
	for _, v := range mappingLogAnalyticsLabelSummaryPriorityEnum {
		values = append(values, v)
	}
	return values
}

// GetLogAnalyticsLabelSummaryPriorityEnumStringValues Enumerates the set of values in String for LogAnalyticsLabelSummaryPriorityEnum
func GetLogAnalyticsLabelSummaryPriorityEnumStringValues() []string {
	return []string{
		"NONE",
		"LOW",
		"MEDIUM",
		"HIGH",
	}
}

// GetMappingLogAnalyticsLabelSummaryPriorityEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingLogAnalyticsLabelSummaryPriorityEnum(val string) (LogAnalyticsLabelSummaryPriorityEnum, bool) {
	mappingLogAnalyticsLabelSummaryPriorityEnumIgnoreCase := make(map[string]LogAnalyticsLabelSummaryPriorityEnum)
	for k, v := range mappingLogAnalyticsLabelSummaryPriorityEnum {
		mappingLogAnalyticsLabelSummaryPriorityEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingLogAnalyticsLabelSummaryPriorityEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}

// LogAnalyticsLabelSummaryTypeEnum Enum with underlying type: string
type LogAnalyticsLabelSummaryTypeEnum string

// Set of constants representing the allowable values for LogAnalyticsLabelSummaryTypeEnum
const (
	LogAnalyticsLabelSummaryTypeInfo    LogAnalyticsLabelSummaryTypeEnum = "INFO"
	LogAnalyticsLabelSummaryTypeProblem LogAnalyticsLabelSummaryTypeEnum = "PROBLEM"
)

var mappingLogAnalyticsLabelSummaryTypeEnum = map[string]LogAnalyticsLabelSummaryTypeEnum{
	"INFO":    LogAnalyticsLabelSummaryTypeInfo,
	"PROBLEM": LogAnalyticsLabelSummaryTypeProblem,
}

// GetLogAnalyticsLabelSummaryTypeEnumValues Enumerates the set of values for LogAnalyticsLabelSummaryTypeEnum
func GetLogAnalyticsLabelSummaryTypeEnumValues() []LogAnalyticsLabelSummaryTypeEnum {
	values := make([]LogAnalyticsLabelSummaryTypeEnum, 0)
	for _, v := range mappingLogAnalyticsLabelSummaryTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetLogAnalyticsLabelSummaryTypeEnumStringValues Enumerates the set of values in String for LogAnalyticsLabelSummaryTypeEnum
func GetLogAnalyticsLabelSummaryTypeEnumStringValues() []string {
	return []string{
		"INFO",
		"PROBLEM",
	}
}

// GetMappingLogAnalyticsLabelSummaryTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingLogAnalyticsLabelSummaryTypeEnum(val string) (LogAnalyticsLabelSummaryTypeEnum, bool) {
	mappingLogAnalyticsLabelSummaryTypeEnumIgnoreCase := make(map[string]LogAnalyticsLabelSummaryTypeEnum)
	for k, v := range mappingLogAnalyticsLabelSummaryTypeEnum {
		mappingLogAnalyticsLabelSummaryTypeEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingLogAnalyticsLabelSummaryTypeEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}
