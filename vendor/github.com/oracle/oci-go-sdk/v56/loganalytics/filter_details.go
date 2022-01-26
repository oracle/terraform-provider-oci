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

// FilterDetails Query builder edit request details.
type FilterDetails struct {

	// Query to apply edits to.
	QueryString *string `mandatory:"true" json:"queryString"`

	// Default subsystem to qualify fields with in the queryString if not specified.
	SubSystem SubSystemNameEnum `mandatory:"true" json:"subSystem"`

	// List of edit operations to be applied in the specified order to the specified queryString.
	Filters []Filter `mandatory:"false" json:"filters"`
}

func (m FilterDetails) String() string {
	return common.PointerString(m)
}
