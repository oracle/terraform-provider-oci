// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// LogAnalytics API
//
// The LogAnalytics API for the LogAnalytics service.
//

package loganalytics

import (
	"github.com/oracle/oci-go-sdk/v33/common"
)

// UpsertLogAnalyticsSourceDetails UpsertLogAnalyticsSourceDetails
type UpsertLogAnalyticsSourceDetails struct {

	// source label conditions
	LabelConditions []LogAnalyticsSourceLabelCondition `mandatory:"false" json:"labelConditions"`

	// data filter definitions
	DataFilterDefinitions []LogAnalyticsSourceDataFilter `mandatory:"false" json:"dataFilterDefinitions"`

	// DB credential name
	DatabaseCredential *string `mandatory:"false" json:"databaseCredential"`

	// extended field definition
	ExtendedFieldDefinitions []LogAnalyticsSourceExtendedFieldDefinition `mandatory:"false" json:"extendedFieldDefinitions"`

	// is for cloud flag
	IsForCloud *bool `mandatory:"false" json:"isForCloud"`

	// labels
	Labels []LogAnalyticsLabelView `mandatory:"false" json:"labels"`

	// metric definitions
	MetricDefinitions []LogAnalyticsMetric `mandatory:"false" json:"metricDefinitions"`

	// metric source map
	Metrics []LogAnalyticsSourceMetric `mandatory:"false" json:"metrics"`

	// out-of-the-box source parser list
	OobParsers []LogAnalyticsParser `mandatory:"false" json:"oobParsers"`

	// parameters
	Parameters []LogAnalyticsParameter `mandatory:"false" json:"parameters"`

	// patterns
	Patterns []LogAnalyticsSourcePattern `mandatory:"false" json:"patterns"`

	// description
	Description *string `mandatory:"false" json:"description"`

	// display name
	DisplayName *string `mandatory:"false" json:"displayName"`

	// source edit version
	EditVersion *int64 `mandatory:"false" json:"editVersion"`

	// source functions
	Functions []LogAnalyticsSourceFunction `mandatory:"false" json:"functions"`

	// source Id
	SourceId *int64 `mandatory:"false" json:"sourceId"`

	// source internal name
	Name *string `mandatory:"false" json:"name"`

	// is secure content flag
	IsSecureContent *bool `mandatory:"false" json:"isSecureContent"`

	// is system flag
	IsSystem *bool `mandatory:"false" json:"isSystem"`

	// parser list
	Parsers []LogAnalyticsParser `mandatory:"false" json:"parsers"`

	// rule Id
	RuleId *int64 `mandatory:"false" json:"ruleId"`

	// source type internal name
	TypeName *string `mandatory:"false" json:"typeName"`

	// source warning configuration
	WarningConfig *int64 `mandatory:"false" json:"warningConfig"`

	// source metadata fields
	MetadataFields []LogAnalyticsSourceMetadataField `mandatory:"false" json:"metadataFields"`

	// tags
	LabelDefinitions []LogAnalyticsLabelDefinition `mandatory:"false" json:"labelDefinitions"`

	// entity types
	EntityTypes []LogAnalyticsSourceEntityType `mandatory:"false" json:"entityTypes"`

	// time zone override
	IsTimezoneOverride *bool `mandatory:"false" json:"isTimezoneOverride"`

	// source parser list
	UserParsers []LogAnalyticsParser `mandatory:"false" json:"userParsers"`
}

func (m UpsertLogAnalyticsSourceDetails) String() string {
	return common.PointerString(m)
}
