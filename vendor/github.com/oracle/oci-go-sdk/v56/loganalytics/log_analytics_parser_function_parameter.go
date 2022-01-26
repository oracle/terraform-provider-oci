// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// LogAnalytics API
//
// The LogAnalytics API for the LogAnalytics service.
//

package loganalytics

import (
	"github.com/oracle/oci-go-sdk/v56/common"
)

// LogAnalyticsParserFunctionParameter LogAnalyticsParserFunctionParameter
type LogAnalyticsParserFunctionParameter struct {

	// The parser function unique identifier.
	ParserFunctionId *int64 `mandatory:"false" json:"parserFunctionId"`

	// The internal name
	ParserFunctionParameterName *string `mandatory:"false" json:"parserFunctionParameterName"`

	// The parameter unique identifier.
	ParserFunctionParameterId *int64 `mandatory:"false" json:"parserFunctionParameterId"`

	// The parameter internal name.
	ParserMetaPluginParameterName *string `mandatory:"false" json:"parserMetaPluginParameterName"`

	// The parameter value.
	ParserMetaPluginParameterValue *string `mandatory:"false" json:"parserMetaPluginParameterValue"`

	// The parser internal name.
	ParserName *string `mandatory:"false" json:"parserName"`

	ParserMetaPluginParameter *LogAnalyticsParserMetaPluginParameter `mandatory:"false" json:"parserMetaPluginParameter"`
}

func (m LogAnalyticsParserFunctionParameter) String() string {
	return common.PointerString(m)
}
