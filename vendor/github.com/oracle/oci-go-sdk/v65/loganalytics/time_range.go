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

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m TimeRange) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
