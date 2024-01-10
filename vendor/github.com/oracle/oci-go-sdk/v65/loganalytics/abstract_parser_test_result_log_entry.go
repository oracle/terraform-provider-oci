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

// AbstractParserTestResultLogEntry AbstractParserTestResultLogEntry
type AbstractParserTestResultLogEntry struct {

	// Extra information attributes.
	ExtraInfoAttributes map[string]string `mandatory:"false" json:"extraInfoAttributes"`

	// The field name value map.
	FieldNameValueMap map[string]string `mandatory:"false" json:"fieldNameValueMap"`

	// The field position value map.
	FieldPositionValueMap map[string]string `mandatory:"false" json:"fieldPositionValueMap"`

	// The parser fields.
	Fields map[string]string `mandatory:"false" json:"fields"`

	// The log entry.
	LogEntry *string `mandatory:"false" json:"logEntry"`

	// The match status.
	MatchStatus *string `mandatory:"false" json:"matchStatus"`

	// The match status description.
	MatchStatusDescription *string `mandatory:"false" json:"matchStatusDescription"`
}

func (m AbstractParserTestResultLogEntry) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m AbstractParserTestResultLogEntry) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
