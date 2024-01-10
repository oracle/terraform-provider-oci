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

// RecallCount This is the recall count statistics for a given tenant
type RecallCount struct {

	// This is the total number of recalls made so far
	RecallCount *int `mandatory:"true" json:"recallCount"`

	// This is the number of recalls that succeeded
	RecallSucceeded *int `mandatory:"true" json:"recallSucceeded"`

	// This is the number of recalls that failed
	RecallFailed *int `mandatory:"true" json:"recallFailed"`

	// This is the number of recalls in pending state
	RecallPending *int `mandatory:"true" json:"recallPending"`

	// This is the maximum number of recalls (including successful and pending recalls) allowed
	RecallLimit *int `mandatory:"true" json:"recallLimit"`
}

func (m RecallCount) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m RecallCount) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
