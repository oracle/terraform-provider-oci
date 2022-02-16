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

// StatusSummary StatusSummary
type StatusSummary struct {

	// The number of chunks processed.
	ChunksProcessed *int64 `mandatory:"false" json:"chunksProcessed"`

	// The failure details, if any.
	FailureDetails *string `mandatory:"false" json:"failureDetails"`

	// The filename.
	Filename *string `mandatory:"false" json:"filename"`

	// The status.
	Status *string `mandatory:"false" json:"status"`

	// The total number of chunks.
	TotalChunks *int64 `mandatory:"false" json:"totalChunks"`
}

func (m StatusSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m StatusSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
