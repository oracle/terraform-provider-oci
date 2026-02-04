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

// ParserTestResultFieldValue A parser test result field value object.
type ParserTestResultFieldValue struct {

	// The field position.
	Position *int64 `mandatory:"false" json:"position"`

	// The field name.
	FieldName *string `mandatory:"false" json:"fieldName"`

	// The field value.
	Value *string `mandatory:"false" json:"value"`

	// The sub parser name.
	ParserName *string `mandatory:"false" json:"parserName"`

	SubParserResult *ParserTestResult `mandatory:"false" json:"subParserResult"`
}

func (m ParserTestResultFieldValue) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ParserTestResultFieldValue) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}
