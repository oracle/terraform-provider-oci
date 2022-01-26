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

// UpsertLogAnalyticsLabelDetailsPriorityEnum Enum with underlying type: string
type UpsertLogAnalyticsLabelDetailsPriorityEnum string

// Set of constants representing the allowable values for UpsertLogAnalyticsLabelDetailsPriorityEnum
const (
	UpsertLogAnalyticsLabelDetailsPriorityNone   UpsertLogAnalyticsLabelDetailsPriorityEnum = "NONE"
	UpsertLogAnalyticsLabelDetailsPriorityLow    UpsertLogAnalyticsLabelDetailsPriorityEnum = "LOW"
	UpsertLogAnalyticsLabelDetailsPriorityMedium UpsertLogAnalyticsLabelDetailsPriorityEnum = "MEDIUM"
	UpsertLogAnalyticsLabelDetailsPriorityHigh   UpsertLogAnalyticsLabelDetailsPriorityEnum = "HIGH"
)

var mappingUpsertLogAnalyticsLabelDetailsPriority = map[string]UpsertLogAnalyticsLabelDetailsPriorityEnum{
	"NONE":   UpsertLogAnalyticsLabelDetailsPriorityNone,
	"LOW":    UpsertLogAnalyticsLabelDetailsPriorityLow,
	"MEDIUM": UpsertLogAnalyticsLabelDetailsPriorityMedium,
	"HIGH":   UpsertLogAnalyticsLabelDetailsPriorityHigh,
}

// GetUpsertLogAnalyticsLabelDetailsPriorityEnumValues Enumerates the set of values for UpsertLogAnalyticsLabelDetailsPriorityEnum
func GetUpsertLogAnalyticsLabelDetailsPriorityEnumValues() []UpsertLogAnalyticsLabelDetailsPriorityEnum {
	values := make([]UpsertLogAnalyticsLabelDetailsPriorityEnum, 0)
	for _, v := range mappingUpsertLogAnalyticsLabelDetailsPriority {
		values = append(values, v)
	}
	return values
}

// UpsertLogAnalyticsLabelDetailsTypeEnum Enum with underlying type: string
type UpsertLogAnalyticsLabelDetailsTypeEnum string

// Set of constants representing the allowable values for UpsertLogAnalyticsLabelDetailsTypeEnum
const (
	UpsertLogAnalyticsLabelDetailsTypeInfo    UpsertLogAnalyticsLabelDetailsTypeEnum = "INFO"
	UpsertLogAnalyticsLabelDetailsTypeProblem UpsertLogAnalyticsLabelDetailsTypeEnum = "PROBLEM"
)

var mappingUpsertLogAnalyticsLabelDetailsType = map[string]UpsertLogAnalyticsLabelDetailsTypeEnum{
	"INFO":    UpsertLogAnalyticsLabelDetailsTypeInfo,
	"PROBLEM": UpsertLogAnalyticsLabelDetailsTypeProblem,
}

// GetUpsertLogAnalyticsLabelDetailsTypeEnumValues Enumerates the set of values for UpsertLogAnalyticsLabelDetailsTypeEnum
func GetUpsertLogAnalyticsLabelDetailsTypeEnumValues() []UpsertLogAnalyticsLabelDetailsTypeEnum {
	values := make([]UpsertLogAnalyticsLabelDetailsTypeEnum, 0)
	for _, v := range mappingUpsertLogAnalyticsLabelDetailsType {
		values = append(values, v)
	}
	return values
}
