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

// LogAnalyticsResourceCategoryCollection A collection of resources and their category assignments.
type LogAnalyticsResourceCategoryCollection struct {

	// An array of categories. The array contents include detailed information about
	// the distinct set of categories assigned to all the listed resources under items.
	Categories []LogAnalyticsCategory `mandatory:"false" json:"categories"`

	// A list of resources and their category assignments
	Items []LogAnalyticsResourceCategory `mandatory:"false" json:"items"`
}

func (m LogAnalyticsResourceCategoryCollection) String() string {
	return common.PointerString(m)
}
