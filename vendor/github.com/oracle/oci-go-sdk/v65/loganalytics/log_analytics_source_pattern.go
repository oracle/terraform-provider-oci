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

// LogAnalyticsSourcePattern LogAnalyticsSourcePattern
type LogAnalyticsSourcePattern struct {

	// The converted text.
	ConvertedText *string `mandatory:"false" json:"convertedText"`

	// The parser unique identifier.
	DbParserId *int64 `mandatory:"false" json:"dbParserId"`

	// The date time columns.
	DbPatternDateTimeColumns *string `mandatory:"false" json:"dbPatternDateTimeColumns"`

	// The date time field.
	DbPatternDateTimeField *string `mandatory:"false" json:"dbPatternDateTimeField"`

	// The sequence column.
	DbPatternSequenceColumn *string `mandatory:"false" json:"dbPatternSequenceColumn"`

	// The parser field list.
	Fields []LogAnalyticsParserField `mandatory:"false" json:"fields"`

	// A flag indicating if this is source pattern is included.
	IsInclude *bool `mandatory:"false" json:"isInclude"`

	// A flag indicating if this is the default source pattern.
	IsDefault *bool `mandatory:"false" json:"isDefault"`

	PatternFilter *LogAnalyticsPatternFilter `mandatory:"false" json:"patternFilter"`

	// The source pattern alias.
	Alias *string `mandatory:"false" json:"alias"`

	// The source pattern description.
	Description *string `mandatory:"false" json:"description"`

	// A flag inidcating whether or not the source pattern is enabled.
	IsEnabled *bool `mandatory:"false" json:"isEnabled"`

	// The source pattern unique identifier.
	PatternId *int64 `mandatory:"false" json:"patternId"`

	// The system flag.  A value of false denotes a custom, or user
	// defined object.  A value of true denotes a built in object.
	IsSystem *bool `mandatory:"false" json:"isSystem"`

	// The source unique identifier.
	SourceId *int64 `mandatory:"false" json:"sourceId"`

	// A flag indicating whether or not agent warnings are suppressed for
	// this source pattern.
	IsAgentWarningSuppressed *bool `mandatory:"false" json:"isAgentWarningSuppressed"`

	// The pattern text.
	PatternText *string `mandatory:"false" json:"patternText"`

	// The pattern type.
	PatternType *int64 `mandatory:"false" json:"patternType"`

	// The source entity type.
	EntityType []string `mandatory:"false" json:"entityType"`

	// A list of pattern properties.
	PatternProperties []LogAnalyticsProperty `mandatory:"false" json:"patternProperties"`
}

func (m LogAnalyticsSourcePattern) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m LogAnalyticsSourcePattern) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
