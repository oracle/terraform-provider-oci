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

// CompareLineResult The result of a comparison of two lines in the two content input strings.
type CompareLineResult struct {

	// A line from the content on the left. This may be empty if there is no matching line on
	// the left for a line in the right content.
	LeftContent *string `mandatory:"false" json:"leftContent"`

	// A line from the content on the right. This may be empty if there is no matching line on
	// the right for a line in the left content.
	RightContent *string `mandatory:"false" json:"rightContent"`

	// The result of the line comparison. An empty string means the lines being compared are the
	// same. A pipe, |, means the lines are different, and a caret, > or <, means the
	// line is only found either on the right or the left.
	DiffType *string `mandatory:"false" json:"diffType"`

	// A comma delimited set of indices that identify which characters are different from those
	// in the right string.
	LeftIndices *string `mandatory:"false" json:"leftIndices"`

	// A comma delimited set of indices that identify which characters are different from those
	// in the left string.
	RightIndices *string `mandatory:"false" json:"rightIndices"`
}

func (m CompareLineResult) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CompareLineResult) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
