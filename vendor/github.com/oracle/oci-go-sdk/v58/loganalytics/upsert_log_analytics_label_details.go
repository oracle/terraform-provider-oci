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

// UpsertLogAnalyticsLabelDetails Upsert LogAnalytics Label Details
type UpsertLogAnalyticsLabelDetails struct {

	// The alias list.
	Aliases []LogAnalyticsLabelAlias `mandatory:"false" json:"aliases"`

	// suggest type
	SuggestType *int64 `mandatory:"false" json:"suggestType"`

	// The label description.
	Description *string `mandatory:"false" json:"description"`

	// The label display name.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// The edit version.
	EditVersion *int64 `mandatory:"false" json:"editVersion"`

	// The label impact.
	Impact *string `mandatory:"false" json:"impact"`

	// The system flag.  A value of false denotes a custom, or user
	// defined label.  A value of true denotes a built in label.
	IsSystem *bool `mandatory:"false" json:"isSystem"`

	// The label name.
	Name *string `mandatory:"false" json:"name"`

	// The label priority. Valid values are (NONE, LOW, HIGH). NONE is default.
	Priority UpsertLogAnalyticsLabelDetailsPriorityEnum `mandatory:"false" json:"priority,omitempty"`

	// The label recommendation.
	Recommendation *string `mandatory:"false" json:"recommendation"`

	// The label type. Valid values are (INFO, PROBLEM). INFO is default.
	Type UpsertLogAnalyticsLabelDetailsTypeEnum `mandatory:"false" json:"type,omitempty"`
}

func (m UpsertLogAnalyticsLabelDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m UpsertLogAnalyticsLabelDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingUpsertLogAnalyticsLabelDetailsPriorityEnum(string(m.Priority)); !ok && m.Priority != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Priority: %s. Supported values are: %s.", m.Priority, strings.Join(GetUpsertLogAnalyticsLabelDetailsPriorityEnumStringValues(), ",")))
	}
	if _, ok := GetMappingUpsertLogAnalyticsLabelDetailsTypeEnum(string(m.Type)); !ok && m.Type != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Type: %s. Supported values are: %s.", m.Type, strings.Join(GetUpsertLogAnalyticsLabelDetailsTypeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UpsertLogAnalyticsLabelDetailsPriorityEnum Enum with underlying type: string
type UpsertLogAnalyticsLabelDetailsPriorityEnum string

// Set of constants representing the allowable values for UpsertLogAnalyticsLabelDetailsPriorityEnum
const (
	UpsertLogAnalyticsLabelDetailsPriorityNone   UpsertLogAnalyticsLabelDetailsPriorityEnum = "NONE"
	UpsertLogAnalyticsLabelDetailsPriorityLow    UpsertLogAnalyticsLabelDetailsPriorityEnum = "LOW"
	UpsertLogAnalyticsLabelDetailsPriorityMedium UpsertLogAnalyticsLabelDetailsPriorityEnum = "MEDIUM"
	UpsertLogAnalyticsLabelDetailsPriorityHigh   UpsertLogAnalyticsLabelDetailsPriorityEnum = "HIGH"
)

var mappingUpsertLogAnalyticsLabelDetailsPriorityEnum = map[string]UpsertLogAnalyticsLabelDetailsPriorityEnum{
	"NONE":   UpsertLogAnalyticsLabelDetailsPriorityNone,
	"LOW":    UpsertLogAnalyticsLabelDetailsPriorityLow,
	"MEDIUM": UpsertLogAnalyticsLabelDetailsPriorityMedium,
	"HIGH":   UpsertLogAnalyticsLabelDetailsPriorityHigh,
}

// GetUpsertLogAnalyticsLabelDetailsPriorityEnumValues Enumerates the set of values for UpsertLogAnalyticsLabelDetailsPriorityEnum
func GetUpsertLogAnalyticsLabelDetailsPriorityEnumValues() []UpsertLogAnalyticsLabelDetailsPriorityEnum {
	values := make([]UpsertLogAnalyticsLabelDetailsPriorityEnum, 0)
	for _, v := range mappingUpsertLogAnalyticsLabelDetailsPriorityEnum {
		values = append(values, v)
	}
	return values
}

// GetUpsertLogAnalyticsLabelDetailsPriorityEnumStringValues Enumerates the set of values in String for UpsertLogAnalyticsLabelDetailsPriorityEnum
func GetUpsertLogAnalyticsLabelDetailsPriorityEnumStringValues() []string {
	return []string{
		"NONE",
		"LOW",
		"MEDIUM",
		"HIGH",
	}
}

// GetMappingUpsertLogAnalyticsLabelDetailsPriorityEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingUpsertLogAnalyticsLabelDetailsPriorityEnum(val string) (UpsertLogAnalyticsLabelDetailsPriorityEnum, bool) {
	mappingUpsertLogAnalyticsLabelDetailsPriorityEnumIgnoreCase := make(map[string]UpsertLogAnalyticsLabelDetailsPriorityEnum)
	for k, v := range mappingUpsertLogAnalyticsLabelDetailsPriorityEnum {
		mappingUpsertLogAnalyticsLabelDetailsPriorityEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingUpsertLogAnalyticsLabelDetailsPriorityEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}

// UpsertLogAnalyticsLabelDetailsTypeEnum Enum with underlying type: string
type UpsertLogAnalyticsLabelDetailsTypeEnum string

// Set of constants representing the allowable values for UpsertLogAnalyticsLabelDetailsTypeEnum
const (
	UpsertLogAnalyticsLabelDetailsTypeInfo    UpsertLogAnalyticsLabelDetailsTypeEnum = "INFO"
	UpsertLogAnalyticsLabelDetailsTypeProblem UpsertLogAnalyticsLabelDetailsTypeEnum = "PROBLEM"
)

var mappingUpsertLogAnalyticsLabelDetailsTypeEnum = map[string]UpsertLogAnalyticsLabelDetailsTypeEnum{
	"INFO":    UpsertLogAnalyticsLabelDetailsTypeInfo,
	"PROBLEM": UpsertLogAnalyticsLabelDetailsTypeProblem,
}

// GetUpsertLogAnalyticsLabelDetailsTypeEnumValues Enumerates the set of values for UpsertLogAnalyticsLabelDetailsTypeEnum
func GetUpsertLogAnalyticsLabelDetailsTypeEnumValues() []UpsertLogAnalyticsLabelDetailsTypeEnum {
	values := make([]UpsertLogAnalyticsLabelDetailsTypeEnum, 0)
	for _, v := range mappingUpsertLogAnalyticsLabelDetailsTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetUpsertLogAnalyticsLabelDetailsTypeEnumStringValues Enumerates the set of values in String for UpsertLogAnalyticsLabelDetailsTypeEnum
func GetUpsertLogAnalyticsLabelDetailsTypeEnumStringValues() []string {
	return []string{
		"INFO",
		"PROBLEM",
	}
}

// GetMappingUpsertLogAnalyticsLabelDetailsTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingUpsertLogAnalyticsLabelDetailsTypeEnum(val string) (UpsertLogAnalyticsLabelDetailsTypeEnum, bool) {
	mappingUpsertLogAnalyticsLabelDetailsTypeEnumIgnoreCase := make(map[string]UpsertLogAnalyticsLabelDetailsTypeEnum)
	for k, v := range mappingUpsertLogAnalyticsLabelDetailsTypeEnum {
		mappingUpsertLogAnalyticsLabelDetailsTypeEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingUpsertLogAnalyticsLabelDetailsTypeEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}
