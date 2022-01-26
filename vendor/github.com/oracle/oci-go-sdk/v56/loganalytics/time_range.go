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

// TimeRange Specify time range. This paramter can be overwritten if time criteria is specified in the query string. If no time criteria are found in query string this time range is used.
type TimeRange struct {

	// Time for query to start matching results from. Start time must be less than end time otherwise it will result in error.
	TimeStart *common.SDKTime `mandatory:"true" json:"timeStart"`

	// Time for query to stop matching results to. End Time must be greater than or equal to start time otherwise it will result in error.
	TimeEnd *common.SDKTime `mandatory:"true" json:"timeEnd"`

	// Time zone for query.
	TimeZone *string `mandatory:"false" json:"timeZone"`
}

func (m TimeRange) String() string {
	return common.PointerString(m)
}
