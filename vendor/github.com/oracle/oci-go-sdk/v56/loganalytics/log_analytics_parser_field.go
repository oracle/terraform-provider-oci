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

// LogAnalyticsParserField LogAnalyticsParserField
type LogAnalyticsParserField struct {
	Field *LogAnalyticsField `mandatory:"false" json:"field"`

	// The parser field unique identifier.
	ParserFieldId *int64 `mandatory:"false" json:"parserFieldId"`

	// the parser field expression.
	ParserFieldExpression *string `mandatory:"false" json:"parserFieldExpression"`

	// The parser field internal name.
	ParserFieldName *string `mandatory:"false" json:"parserFieldName"`

	// The storage field name.
	StorageFieldName *string `mandatory:"false" json:"storageFieldName"`

	// The integrator name.
	ParserFieldIntegratorName *string `mandatory:"false" json:"parserFieldIntegratorName"`

	// The parser internal name.
	ParserName *string `mandatory:"false" json:"parserName"`

	// The parser field sequence.
	ParserFieldSequence *int64 `mandatory:"false" json:"parserFieldSequence"`

	Parser *LogAnalyticsParser `mandatory:"false" json:"parser"`

	// The structured column information.
	StructuredColumnInfo *string `mandatory:"false" json:"structuredColumnInfo"`
}

func (m LogAnalyticsParserField) String() string {
	return common.PointerString(m)
}
