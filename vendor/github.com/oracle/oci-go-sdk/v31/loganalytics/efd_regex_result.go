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

// EfdRegexResult EfdRegexResult
type EfdRegexResult struct {

	// baseFieldName
	BaseFieldName *string `mandatory:"false" json:"baseFieldName"`

	// id
	Id *int64 `mandatory:"false" json:"id"`

	MatchResult *RegexMatchResult `mandatory:"false" json:"matchResult"`

	// parsedFieldCount
	ParsedFieldCount *int `mandatory:"false" json:"parsedFieldCount"`

	// parsedFields
	ParsedFields map[string]string `mandatory:"false" json:"parsedFields"`

	// regex
	Regex *string `mandatory:"false" json:"regex"`

	// status
	Status *string `mandatory:"false" json:"status"`

	// statusDescription
	StatusDescription *string `mandatory:"false" json:"statusDescription"`

	// isValidRegexSyntax
	IsValidRegexSyntax *bool `mandatory:"false" json:"isValidRegexSyntax"`

	// violations
	Violations []Violation `mandatory:"false" json:"violations"`
}

func (m EfdRegexResult) String() string {
	return common.PointerString(m)
}
