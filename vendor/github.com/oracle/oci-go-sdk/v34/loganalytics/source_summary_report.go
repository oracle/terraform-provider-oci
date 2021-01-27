// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// LogAnalytics API
//
// The LogAnalytics API for the LogAnalytics service.
//

package loganalytics

import (
	"github.com/oracle/oci-go-sdk/v34/common"
)

// SourceSummaryReport SourceSummaryReport
type SourceSummaryReport struct {

	// non out-of-the-box count
	NonOobCount *int `mandatory:"false" json:"nonOobCount"`

	// count of sources set to auto-associate
	AutoAssociationSourceCount *int `mandatory:"false" json:"autoAssociationSourceCount"`

	// out-of-the-box count
	OobCount *int `mandatory:"false" json:"oobCount"`
}

func (m SourceSummaryReport) String() string {
	return common.PointerString(m)
}
