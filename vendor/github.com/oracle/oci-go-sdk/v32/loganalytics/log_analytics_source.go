// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// LogAnalytics API
//
// The LogAnalytics API for the LogAnalytics service.
//

package loganalytics

import (
	"github.com/oracle/oci-go-sdk/v32/common"
)

// LogAnalyticsSource LogAnalyticsSource
type LogAnalyticsSource struct {

	// alert conditions
	LabelConditions []LogAnalyticsSourceLabelCondition `mandatory:"false" json:"labelConditions"`

	// association count
	AssociationCount *int `mandatory:"false" json:"associationCount"`

	// association entity
	AssociationEntity []LogAnalyticsAssociation `mandatory:"false" json:"associationEntity"`

	// data filter definitions
	DataFilterDefinitions []LogAnalyticsSourceDataFilter `mandatory:"false" json:"dataFilterDefinitions"`

	// DB credential
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

	// pattern count
	PatternCount *int `mandatory:"false" json:"patternCount"`

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

	// rule auto association enabled flag
	IsAutoAssociationEnabled *bool `mandatory:"false" json:"isAutoAssociationEnabled"`

	// rule auto association override
	IsAutoAssociationOverride *bool `mandatory:"false" json:"isAutoAssociationOverride"`

	// rule Id
	RuleId *int64 `mandatory:"false" json:"ruleId"`

	// source type internal name
	TypeName *string `mandatory:"false" json:"typeName"`

	// source type name
	TypeDisplayName *string `mandatory:"false" json:"typeDisplayName"`

	// source warning configuration
	WarningConfig *int64 `mandatory:"false" json:"warningConfig"`

	// source metadata fields
	MetadataFields []LogAnalyticsSourceMetadataField `mandatory:"false" json:"metadataFields"`

	// tags
	LabelDefinitions []LogAnalyticsLabelDefinition `mandatory:"false" json:"labelDefinitions"`

	// Entity types
	EntityTypes []LogAnalyticsSourceEntityType `mandatory:"false" json:"entityTypes"`

	// time zone override
	IsTimezoneOverride *bool `mandatory:"false" json:"isTimezoneOverride"`

	// source parser list
	UserParsers []LogAnalyticsParser `mandatory:"false" json:"userParsers"`

	// timeUpdated
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`
}

func (m LogAnalyticsSource) String() string {
	return common.PointerString(m)
}
