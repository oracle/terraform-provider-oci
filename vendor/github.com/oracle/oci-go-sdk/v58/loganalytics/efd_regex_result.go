// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// LogAnalytics API
//
// The LogAnalytics API for the LogAnalytics service.
//

package loganalytics

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v58/common"
	"strings"
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

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m EfdRegexResult) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
