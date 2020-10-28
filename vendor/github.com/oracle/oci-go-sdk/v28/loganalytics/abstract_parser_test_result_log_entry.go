// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// LogAnalytics API
//
// The LogAnalytics API for the LogAnalytics service.
//

package loganalytics

import (
	"github.com/oracle/oci-go-sdk/v28/common"
)

// AbstractParserTestResultLogEntry AbstractParserTestResultLogEntry
type AbstractParserTestResultLogEntry struct {

	// extra info attributes
	ExtraInfoAttributes map[string]string `mandatory:"false" json:"extraInfoAttributes"`

	// field name value map
	FieldNameValueMap map[string]string `mandatory:"false" json:"fieldNameValueMap"`

	// field position value map
	FieldPositionValueMap map[string]string `mandatory:"false" json:"fieldPositionValueMap"`

	// fields
	Fields map[string]string `mandatory:"false" json:"fields"`

	// log entry
	LogEntry *string `mandatory:"false" json:"logEntry"`

	// match status
	MatchStatus *string `mandatory:"false" json:"matchStatus"`

	// match status description
	MatchStatusDescription *string `mandatory:"false" json:"matchStatusDescription"`
}

func (m AbstractParserTestResultLogEntry) String() string {
	return common.PointerString(m)
}
