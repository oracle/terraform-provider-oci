// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// LogAnalytics API
//
// The LogAnalytics API for the LogAnalytics service.
//

package loganalytics

import (
	"github.com/oracle/oci-go-sdk/v37/common"
)

// LogAnalyticsParserFunction LogAnalyticsParserFunction
type LogAnalyticsParserFunction struct {
	ParserMetaPlugin *LogAnalyticsParserMetaPlugin `mandatory:"false" json:"parserMetaPlugin"`

	// plugin instance Id
	ParserFunctionId *int64 `mandatory:"false" json:"parserFunctionId"`

	// plugin instance internal name
	ParserFunctionName *string `mandatory:"false" json:"parserFunctionName"`

	// is enabled flag
	IsEnabled *bool `mandatory:"false" json:"isEnabled"`

	// is system flag
	IsSystem *bool `mandatory:"false" json:"isSystem"`

	// parser Id
	ParserId *int64 `mandatory:"false" json:"parserId"`

	// parser internal name
	ParserName *string `mandatory:"false" json:"parserName"`

	// plugin type internal name
	ParserMetaPluginName *string `mandatory:"false" json:"parserMetaPluginName"`

	// priority
	ParserFunctionPriority *int64 `mandatory:"false" json:"parserFunctionPriority"`

	// parameter map list
	ParserFunctionParameters []LogAnalyticsParserFunctionParameter `mandatory:"false" json:"parserFunctionParameters"`
}

func (m LogAnalyticsParserFunction) String() string {
	return common.PointerString(m)
}
