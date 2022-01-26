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

// EfdRegexResult EfdRegexResult
type EfdRegexResult struct {

	// The base field name.
	BaseFieldName *string `mandatory:"false" json:"baseFieldName"`

	// the unique identifier.
	Id *int64 `mandatory:"false" json:"id"`

	MatchResult *RegexMatchResult `mandatory:"false" json:"matchResult"`

	// The parsed field count.
	ParsedFieldCount *int `mandatory:"false" json:"parsedFieldCount"`

	// The parsed fields.
	ParsedFields map[string]string `mandatory:"false" json:"parsedFields"`

	// The regular expression.
	Regex *string `mandatory:"false" json:"regex"`

	// The status.
	Status *string `mandatory:"false" json:"status"`

	// The Status description.
	StatusDescription *string `mandatory:"false" json:"statusDescription"`

	// A flag indicating whether or not the regular expression is valid.
	IsValidRegexSyntax *bool `mandatory:"false" json:"isValidRegexSyntax"`

	// The list of violations (if any).
	Violations []Violation `mandatory:"false" json:"violations"`
}

func (m EfdRegexResult) String() string {
	return common.PointerString(m)
}
