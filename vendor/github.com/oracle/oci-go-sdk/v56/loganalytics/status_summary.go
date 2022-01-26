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
