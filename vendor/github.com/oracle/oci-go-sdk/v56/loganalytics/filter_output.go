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

// FilterOutput Query builder api response object containing updated querystring's
type FilterOutput struct {

	// Modified user visible query string.
	DisplayQueryString *string `mandatory:"true" json:"displayQueryString"`

	// Modified localization agnostic query string.
	InternalQueryString *string `mandatory:"true" json:"internalQueryString"`

	// Operation response time.
	ResponseTimeInMs *int64 `mandatory:"false" json:"responseTimeInMs"`
}

func (m FilterOutput) String() string {
	return common.PointerString(m)
}
