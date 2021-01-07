// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// LogAnalytics API
//
// The LogAnalytics API for the LogAnalytics service.
//

package loganalytics

import (
	"github.com/oracle/oci-go-sdk/v31/common"
)

// UpsertLogAnalyticsParserDetails UpsertLogAnalyticsParserDetails
type UpsertLogAnalyticsParserDetails struct {

	// content
	Content *string `mandatory:"false" json:"content"`

	// description
	Description *string `mandatory:"false" json:"description"`

	// display name
	DisplayName *string `mandatory:"false" json:"displayName"`

	// edit version
	EditVersion *int64 `mandatory:"false" json:"editVersion"`

	// encoding
	Encoding *string `mandatory:"false" json:"encoding"`

	// example content
	ExampleContent *string `mandatory:"false" json:"exampleContent"`

	// fields Maps
	FieldMaps []LogAnalyticsParserField `mandatory:"false" json:"fieldMaps"`

	// footer regular expression
	FooterContent *string `mandatory:"false" json:"footerContent"`

	// header content
	HeaderContent *string `mandatory:"false" json:"headerContent"`

	// Name
	Name *string `mandatory:"false" json:"name"`

	// is default flag
	IsDefault *bool `mandatory:"false" json:"isDefault"`

	// is single line content
	IsSingleLineContent *bool `mandatory:"false" json:"isSingleLineContent"`

	// is system flag
	IsSystem *bool `mandatory:"false" json:"isSystem"`

	// language
	Language *string `mandatory:"false" json:"language"`

	// log type test request version
	LogTypeTestRequestVersion *int `mandatory:"false" json:"logTypeTestRequestVersion"`

	// parser ignore line characters
	ParserIgnorelineCharacters *string `mandatory:"false" json:"parserIgnorelineCharacters"`

	// sequence
	ParserSequence *int `mandatory:"false" json:"parserSequence"`

	// time zone
	ParserTimezone *string `mandatory:"false" json:"parserTimezone"`

	// write once
	IsParserWrittenOnce *bool `mandatory:"false" json:"isParserWrittenOnce"`

	// plugin instance list
	ParserFunctions []LogAnalyticsParserFunction `mandatory:"false" json:"parserFunctions"`

	// tokenize original text
	ShouldTokenizeOriginalText *bool `mandatory:"false" json:"shouldTokenizeOriginalText"`

	// type
	Type UpsertLogAnalyticsParserDetailsTypeEnum `mandatory:"false" json:"type,omitempty"`
}

func (m UpsertLogAnalyticsParserDetails) String() string {
	return common.PointerString(m)
}

// UpsertLogAnalyticsParserDetailsTypeEnum Enum with underlying type: string
type UpsertLogAnalyticsParserDetailsTypeEnum string

// Set of constants representing the allowable values for UpsertLogAnalyticsParserDetailsTypeEnum
const (
	UpsertLogAnalyticsParserDetailsTypeXml   UpsertLogAnalyticsParserDetailsTypeEnum = "XML"
	UpsertLogAnalyticsParserDetailsTypeJson  UpsertLogAnalyticsParserDetailsTypeEnum = "JSON"
	UpsertLogAnalyticsParserDetailsTypeRegex UpsertLogAnalyticsParserDetailsTypeEnum = "REGEX"
	UpsertLogAnalyticsParserDetailsTypeOdl   UpsertLogAnalyticsParserDetailsTypeEnum = "ODL"
)

var mappingUpsertLogAnalyticsParserDetailsType = map[string]UpsertLogAnalyticsParserDetailsTypeEnum{
	"XML":   UpsertLogAnalyticsParserDetailsTypeXml,
	"JSON":  UpsertLogAnalyticsParserDetailsTypeJson,
	"REGEX": UpsertLogAnalyticsParserDetailsTypeRegex,
	"ODL":   UpsertLogAnalyticsParserDetailsTypeOdl,
}

// GetUpsertLogAnalyticsParserDetailsTypeEnumValues Enumerates the set of values for UpsertLogAnalyticsParserDetailsTypeEnum
func GetUpsertLogAnalyticsParserDetailsTypeEnumValues() []UpsertLogAnalyticsParserDetailsTypeEnum {
	values := make([]UpsertLogAnalyticsParserDetailsTypeEnum, 0)
	for _, v := range mappingUpsertLogAnalyticsParserDetailsType {
		values = append(values, v)
	}
	return values
}
