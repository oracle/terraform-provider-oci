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

// LogAnalyticsParserFunction LogAnalyticsParserFunction
type LogAnalyticsParserFunction struct {
	ParserMetaPlugin *LogAnalyticsParserMetaPlugin `mandatory:"false" json:"parserMetaPlugin"`

	// The parser function unique identifier.
	ParserFunctionId *int64 `mandatory:"false" json:"parserFunctionId"`

	// The parser function internal name.
	ParserFunctionName *string `mandatory:"false" json:"parserFunctionName"`

	// A flag inidcating whether or not the parser function is enabled.
	IsEnabled *bool `mandatory:"false" json:"isEnabled"`

	// The system flag.  A value of false denotes a custom, or user
	// defined object.  A value of true denotes a built in object.
	IsSystem *bool `mandatory:"false" json:"isSystem"`

	// The associated parser unique identifier.
	ParserId *int64 `mandatory:"false" json:"parserId"`

	// The associated parser internal name.
	ParserName *string `mandatory:"false" json:"parserName"`

	// The plugin internal name.
	ParserMetaPluginName *string `mandatory:"false" json:"parserMetaPluginName"`

	// The parser function priority.
	ParserFunctionPriority *int64 `mandatory:"false" json:"parserFunctionPriority"`

	// The parser function parameter list.
	ParserFunctionParameters []LogAnalyticsParserFunctionParameter `mandatory:"false" json:"parserFunctionParameters"`
}

func (m LogAnalyticsParserFunction) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m LogAnalyticsParserFunction) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
