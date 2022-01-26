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

// ParsedContent Parsed representation of the log file.
type ParsedContent struct {

	// List of field names.
	FieldNames []string `mandatory:"false" json:"fieldNames"`

	// List of field display names.
	FieldDisplayNames []string `mandatory:"false" json:"fieldDisplayNames"`

	// Parsed field values.
	ParsedFieldValues []ParsedField `mandatory:"false" json:"parsedFieldValues"`

	// Sample log entries picked up from the given file for validation.
	LogContent *string `mandatory:"false" json:"logContent"`

	// Sample Size taken for validation.
	SampleSize *int `mandatory:"false" json:"sampleSize"`

	// Match Status.
	MatchStatus *string `mandatory:"false" json:"matchStatus"`
}

func (m ParsedContent) String() string {
	return common.PointerString(m)
}
