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

// DependentParser A parser used by another parser.
type DependentParser struct {

	// The parser name.
	ParserName *string `mandatory:"false" json:"parserName"`

	// The parser display name.
	ParserDisplayName *string `mandatory:"false" json:"parserDisplayName"`

	// The parser unique identifier.
	ParserId *int64 `mandatory:"false" json:"parserId"`

	// The system flag.  A value of false denotes a custom, or user
	// defined object.  A value of true denotes a built in object.
	IsSystem *bool `mandatory:"false" json:"isSystem"`

	// The parser type
	ParserType DependentParserParserTypeEnum `mandatory:"false" json:"parserType,omitempty"`

	// The list of dependencies of the parser.
	Dependencies []Dependency `mandatory:"false" json:"dependencies"`
}

func (m DependentParser) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m DependentParser) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingDependentParserParserTypeEnum(string(m.ParserType)); !ok && m.ParserType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ParserType: %s. Supported values are: %s.", m.ParserType, strings.Join(GetDependentParserParserTypeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// DependentParserParserTypeEnum Enum with underlying type: string
type DependentParserParserTypeEnum string

// Set of constants representing the allowable values for DependentParserParserTypeEnum
const (
	DependentParserParserTypeXml       DependentParserParserTypeEnum = "XML"
	DependentParserParserTypeJson      DependentParserParserTypeEnum = "JSON"
	DependentParserParserTypeRegex     DependentParserParserTypeEnum = "REGEX"
	DependentParserParserTypeOdl       DependentParserParserTypeEnum = "ODL"
	DependentParserParserTypeDelimited DependentParserParserTypeEnum = "DELIMITED"
)

var mappingDependentParserParserTypeEnum = map[string]DependentParserParserTypeEnum{
	"XML":       DependentParserParserTypeXml,
	"JSON":      DependentParserParserTypeJson,
	"REGEX":     DependentParserParserTypeRegex,
	"ODL":       DependentParserParserTypeOdl,
	"DELIMITED": DependentParserParserTypeDelimited,
}

var mappingDependentParserParserTypeEnumLowerCase = map[string]DependentParserParserTypeEnum{
	"xml":       DependentParserParserTypeXml,
	"json":      DependentParserParserTypeJson,
	"regex":     DependentParserParserTypeRegex,
	"odl":       DependentParserParserTypeOdl,
	"delimited": DependentParserParserTypeDelimited,
}

// GetDependentParserParserTypeEnumValues Enumerates the set of values for DependentParserParserTypeEnum
func GetDependentParserParserTypeEnumValues() []DependentParserParserTypeEnum {
	values := make([]DependentParserParserTypeEnum, 0)
	for _, v := range mappingDependentParserParserTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetDependentParserParserTypeEnumStringValues Enumerates the set of values in String for DependentParserParserTypeEnum
func GetDependentParserParserTypeEnumStringValues() []string {
	return []string{
		"XML",
		"JSON",
		"REGEX",
		"ODL",
		"DELIMITED",
	}
}

// GetMappingDependentParserParserTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDependentParserParserTypeEnum(val string) (DependentParserParserTypeEnum, bool) {
	enum, ok := mappingDependentParserParserTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
