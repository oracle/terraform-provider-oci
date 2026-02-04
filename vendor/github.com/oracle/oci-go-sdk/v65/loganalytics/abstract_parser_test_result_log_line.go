// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
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

// AbstractParserTestResultLogLine AbstractParserTestResultLogLine
type AbstractParserTestResultLogLine struct {

	// The original log line.
	OriginalLogLine *string `mandatory:"false" json:"originalLogLine"`

	// The pre-processed log line.
	PreProcessedLogLine *string `mandatory:"false" json:"preProcessedLogLine"`

	// The find start index.
	FindStartIndex *int `mandatory:"false" json:"findStartIndex"`

	// The find end index.
	FindEndIndex *int `mandatory:"false" json:"findEndIndex"`

	// The replacement string.
	ReplaceString *string `mandatory:"false" json:"replaceString"`

	// The replace start index.
	ReplaceStartIndex *int `mandatory:"false" json:"replaceStartIndex"`

	// The replace end index.
	ReplaceEndIndex *int `mandatory:"false" json:"replaceEndIndex"`

	// The group name value map.
	GrpNameValueMap map[string]NamedCaptureValue `mandatory:"false" json:"grpNameValueMap"`
}

func (m AbstractParserTestResultLogLine) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m AbstractParserTestResultLogLine) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}
