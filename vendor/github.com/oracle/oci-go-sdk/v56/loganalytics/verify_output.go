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

// VerifyOutput Verify acceleration output.
type VerifyOutput struct {

	// Acceleration task identifier.
	ScheduledTaskId *string `mandatory:"true" json:"scheduledTaskId"`

	// Response time in ms.
	ResponseTimeInMs *int64 `mandatory:"true" json:"responseTimeInMs"`

	// Total match count.
	TotalMatchedCount *int64 `mandatory:"true" json:"totalMatchedCount"`

	// Total count.
	TotalCount *int `mandatory:"true" json:"totalCount"`

	// Acceleration result columns, included if requested (shouldIncludeResults).
	Columns []ResultColumn `mandatory:"false" json:"columns"`

	// Acceleration result values, included if requested (shouldIncludeResults).
	Results []map[string]interface{} `mandatory:"false" json:"results"`
}

func (m VerifyOutput) String() string {
	return common.PointerString(m)
}
