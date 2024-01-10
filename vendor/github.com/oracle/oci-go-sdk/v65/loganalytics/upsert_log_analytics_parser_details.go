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

// UpsertLogAnalyticsParserDetails UpsertLogAnalyticsParserDetails
type UpsertLogAnalyticsParserDetails struct {

	// The content.
	Content *string `mandatory:"false" json:"content"`

	// The parser description.
	Description *string `mandatory:"false" json:"description"`

	// The parser display name.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// The parser edit version.
	EditVersion *int64 `mandatory:"false" json:"editVersion"`

	// The encoding.
	Encoding *string `mandatory:"false" json:"encoding"`

	// Example content.
	ExampleContent *string `mandatory:"false" json:"exampleContent"`

	// The parser fields.
	FieldMaps []LogAnalyticsParserField `mandatory:"false" json:"fieldMaps"`

	// The footer regular expression.
	FooterContent *string `mandatory:"false" json:"footerContent"`

	// The header content.
	HeaderContent *string `mandatory:"false" json:"headerContent"`

	// The parser internal name.
	Name *string `mandatory:"false" json:"name"`

	// A flag indicating if this is a default parser.
	IsDefault *bool `mandatory:"false" json:"isDefault"`

	// A flag indicating if this is a single line content parser.
	IsSingleLineContent *bool `mandatory:"false" json:"isSingleLineContent"`

	// The system flag.  A value of false denotes a custom, or user
	// defined object.  A value of true denotes a built in object.
	IsSystem *bool `mandatory:"false" json:"isSystem"`

	// The language.
	Language *string `mandatory:"false" json:"language"`

	// The log type test request version.
	LogTypeTestRequestVersion *int `mandatory:"false" json:"logTypeTestRequestVersion"`

	// The line characters for the parser to ignore.
	ParserIgnorelineCharacters *string `mandatory:"false" json:"parserIgnorelineCharacters"`

	// The parser sequence.
	ParserSequence *int `mandatory:"false" json:"parserSequence"`

	// The time zone.
	ParserTimezone *string `mandatory:"false" json:"parserTimezone"`

	// A flag indicating whther or not the parser is write once.
	IsParserWrittenOnce *bool `mandatory:"false" json:"isParserWrittenOnce"`

	// The parser function list.
	ParserFunctions []LogAnalyticsParserFunction `mandatory:"false" json:"parserFunctions"`

	// A flag indicating whether or not to tokenize the original text.
	ShouldTokenizeOriginalText *bool `mandatory:"false" json:"shouldTokenizeOriginalText"`

	// The parser field delimiter.
	FieldDelimiter *string `mandatory:"false" json:"fieldDelimiter"`

	// The parser field qualifier.
	FieldQualifier *string `mandatory:"false" json:"fieldQualifier"`

	// The parser type.  Default value is REGEX.
	Type UpsertLogAnalyticsParserDetailsTypeEnum `mandatory:"false" json:"type,omitempty"`

	// A flag indicating whether the XML parser should consider the namespace(s) while processing the log data.
	IsNamespaceAware *bool `mandatory:"false" json:"isNamespaceAware"`

	// An array of categories to assign to the parser. Specifying the name attribute for each category would suffice.
	// Oracle-defined category assignments cannot be removed.
	Categories []LogAnalyticsCategory `mandatory:"false" json:"categories"`
}

func (m UpsertLogAnalyticsParserDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m UpsertLogAnalyticsParserDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingUpsertLogAnalyticsParserDetailsTypeEnum(string(m.Type)); !ok && m.Type != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Type: %s. Supported values are: %s.", m.Type, strings.Join(GetUpsertLogAnalyticsParserDetailsTypeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UpsertLogAnalyticsParserDetailsTypeEnum Enum with underlying type: string
type UpsertLogAnalyticsParserDetailsTypeEnum string

// Set of constants representing the allowable values for UpsertLogAnalyticsParserDetailsTypeEnum
const (
	UpsertLogAnalyticsParserDetailsTypeXml       UpsertLogAnalyticsParserDetailsTypeEnum = "XML"
	UpsertLogAnalyticsParserDetailsTypeJson      UpsertLogAnalyticsParserDetailsTypeEnum = "JSON"
	UpsertLogAnalyticsParserDetailsTypeRegex     UpsertLogAnalyticsParserDetailsTypeEnum = "REGEX"
	UpsertLogAnalyticsParserDetailsTypeOdl       UpsertLogAnalyticsParserDetailsTypeEnum = "ODL"
	UpsertLogAnalyticsParserDetailsTypeDelimited UpsertLogAnalyticsParserDetailsTypeEnum = "DELIMITED"
)

var mappingUpsertLogAnalyticsParserDetailsTypeEnum = map[string]UpsertLogAnalyticsParserDetailsTypeEnum{
	"XML":       UpsertLogAnalyticsParserDetailsTypeXml,
	"JSON":      UpsertLogAnalyticsParserDetailsTypeJson,
	"REGEX":     UpsertLogAnalyticsParserDetailsTypeRegex,
	"ODL":       UpsertLogAnalyticsParserDetailsTypeOdl,
	"DELIMITED": UpsertLogAnalyticsParserDetailsTypeDelimited,
}

var mappingUpsertLogAnalyticsParserDetailsTypeEnumLowerCase = map[string]UpsertLogAnalyticsParserDetailsTypeEnum{
	"xml":       UpsertLogAnalyticsParserDetailsTypeXml,
	"json":      UpsertLogAnalyticsParserDetailsTypeJson,
	"regex":     UpsertLogAnalyticsParserDetailsTypeRegex,
	"odl":       UpsertLogAnalyticsParserDetailsTypeOdl,
	"delimited": UpsertLogAnalyticsParserDetailsTypeDelimited,
}

// GetUpsertLogAnalyticsParserDetailsTypeEnumValues Enumerates the set of values for UpsertLogAnalyticsParserDetailsTypeEnum
func GetUpsertLogAnalyticsParserDetailsTypeEnumValues() []UpsertLogAnalyticsParserDetailsTypeEnum {
	values := make([]UpsertLogAnalyticsParserDetailsTypeEnum, 0)
	for _, v := range mappingUpsertLogAnalyticsParserDetailsTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetUpsertLogAnalyticsParserDetailsTypeEnumStringValues Enumerates the set of values in String for UpsertLogAnalyticsParserDetailsTypeEnum
func GetUpsertLogAnalyticsParserDetailsTypeEnumStringValues() []string {
	return []string{
		"XML",
		"JSON",
		"REGEX",
		"ODL",
		"DELIMITED",
	}
}

// GetMappingUpsertLogAnalyticsParserDetailsTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingUpsertLogAnalyticsParserDetailsTypeEnum(val string) (UpsertLogAnalyticsParserDetailsTypeEnum, bool) {
	enum, ok := mappingUpsertLogAnalyticsParserDetailsTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
