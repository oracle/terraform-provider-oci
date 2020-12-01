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

// LogAnalyticsSourcePattern LogAnalyticsSourcePattern
type LogAnalyticsSourcePattern struct {

	// converted text
	ConvertedText *string `mandatory:"false" json:"convertedText"`

	// parser Id
	DbParserId *int64 `mandatory:"false" json:"dbParserId"`

	// date time columns
	DbPatternDateTimeColumns *string `mandatory:"false" json:"dbPatternDateTimeColumns"`

	// date time field
	DbPatternDateTimeField *string `mandatory:"false" json:"dbPatternDateTimeField"`

	// sequence column
	DbPatternSequenceColumn *string `mandatory:"false" json:"dbPatternSequenceColumn"`

	// field list
	Fields []LogAnalyticsParserField `mandatory:"false" json:"fields"`

	// is include flag
	IsInclude *bool `mandatory:"false" json:"isInclude"`

	// is default flag
	IsDefault *bool `mandatory:"false" json:"isDefault"`

	PatternFilter *LogAnalyticsPatternFilter `mandatory:"false" json:"patternFilter"`

	// alias
	Alias *string `mandatory:"false" json:"alias"`

	// description
	Description *string `mandatory:"false" json:"description"`

	// is enabled flag
	IsEnabled *bool `mandatory:"false" json:"isEnabled"`

	// pattern Id
	PatternId *int64 `mandatory:"false" json:"patternId"`

	// is system flag
	IsSystem *bool `mandatory:"false" json:"isSystem"`

	// source Id
	SourceId *int64 `mandatory:"false" json:"sourceId"`

	// suppress agent warning
	IsAgentWarningSuppressed *bool `mandatory:"false" json:"isAgentWarningSuppressed"`

	// pattern text
	PatternText *string `mandatory:"false" json:"patternText"`

	// pattern type
	PatternType *int64 `mandatory:"false" json:"patternType"`

	// source entity types
	EntityType []string `mandatory:"false" json:"entityType"`
}

func (m LogAnalyticsSourcePattern) String() string {
	return common.PointerString(m)
}
