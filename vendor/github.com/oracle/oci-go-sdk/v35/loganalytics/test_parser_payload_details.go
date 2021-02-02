// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// LogAnalytics API
//
// The LogAnalytics API for the LogAnalytics service.
//

package loganalytics

import (
	"github.com/oracle/oci-go-sdk/v35/common"
)

// TestParserPayloadDetails TestParserPayloadDetails
type TestParserPayloadDetails struct {

	// content
	Content *string `mandatory:"false" json:"content"`

	// description
	Description *string `mandatory:"false" json:"description"`

	// Display name
	DisplayName *string `mandatory:"false" json:"displayName"`

	// encoding
	Encoding *string `mandatory:"false" json:"encoding"`

	// exampleContent
	ExampleContent *string `mandatory:"false" json:"exampleContent"`

	// fieldMaps
	FieldMaps []LogAnalyticsParserField `mandatory:"false" json:"fieldMaps"`

	// footerRegex
	FooterContent *string `mandatory:"false" json:"footerContent"`

	// headerContent
	HeaderContent *string `mandatory:"false" json:"headerContent"`

	// name
	Name *string `mandatory:"false" json:"name"`

	// isDefault
	IsDefault *bool `mandatory:"false" json:"isDefault"`

	// isSingleLineContent
	IsSingleLineContent *bool `mandatory:"false" json:"isSingleLineContent"`

	// isSystem
	IsSystem *bool `mandatory:"false" json:"isSystem"`

	// language
	Language *string `mandatory:"false" json:"language"`

	// lastUpdatedDate
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`

	// logTypeTestRequestVersion
	LogTypeTestRequestVersion *int `mandatory:"false" json:"logTypeTestRequestVersion"`

	Metadata *UiParserTestMetadata `mandatory:"false" json:"metadata"`

	// parser ignore linechars
	ParserIgnorelineCharacters *string `mandatory:"false" json:"parserIgnorelineCharacters"`

	// parser is hidden
	IsHidden *int64 `mandatory:"false" json:"isHidden"`

	// parser seq
	ParserSequence *int `mandatory:"false" json:"parserSequence"`

	// parser timezone
	ParserTimezone *string `mandatory:"false" json:"parserTimezone"`

	// isParserWrittenOnce
	IsParserWrittenOnce *bool `mandatory:"false" json:"isParserWrittenOnce"`

	// plugin instance list
	ParserFunctions []LogAnalyticsParserFunction `mandatory:"false" json:"parserFunctions"`

	// tokenize original text
	ShouldTokenizeOriginalText *bool `mandatory:"false" json:"shouldTokenizeOriginalText"`

	// type
	Type TestParserPayloadDetailsTypeEnum `mandatory:"false" json:"type,omitempty"`
}

func (m TestParserPayloadDetails) String() string {
	return common.PointerString(m)
}

// TestParserPayloadDetailsTypeEnum Enum with underlying type: string
type TestParserPayloadDetailsTypeEnum string

// Set of constants representing the allowable values for TestParserPayloadDetailsTypeEnum
const (
	TestParserPayloadDetailsTypeXml   TestParserPayloadDetailsTypeEnum = "XML"
	TestParserPayloadDetailsTypeJson  TestParserPayloadDetailsTypeEnum = "JSON"
	TestParserPayloadDetailsTypeRegex TestParserPayloadDetailsTypeEnum = "REGEX"
	TestParserPayloadDetailsTypeOdl   TestParserPayloadDetailsTypeEnum = "ODL"
)

var mappingTestParserPayloadDetailsType = map[string]TestParserPayloadDetailsTypeEnum{
	"XML":   TestParserPayloadDetailsTypeXml,
	"JSON":  TestParserPayloadDetailsTypeJson,
	"REGEX": TestParserPayloadDetailsTypeRegex,
	"ODL":   TestParserPayloadDetailsTypeOdl,
}

// GetTestParserPayloadDetailsTypeEnumValues Enumerates the set of values for TestParserPayloadDetailsTypeEnum
func GetTestParserPayloadDetailsTypeEnumValues() []TestParserPayloadDetailsTypeEnum {
	values := make([]TestParserPayloadDetailsTypeEnum, 0)
	for _, v := range mappingTestParserPayloadDetailsType {
		values = append(values, v)
	}
	return values
}
