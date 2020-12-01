// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// LogAnalytics API
//
// The LogAnalytics API for the LogAnalytics service.
//

package loganalytics

import (
	"github.com/oracle/oci-go-sdk/v30/common"
)

// UsageStatusItem UsageStatusItem
type UsageStatusItem struct {

	// data type
	DataType *string `mandatory:"false" json:"dataType"`

	// is the field multi valued
	IsMultiValued *bool `mandatory:"false" json:"isMultiValued"`

	// current usage
	CurrentUsage *int64 `mandatory:"false" json:"currentUsage"`

	// maximum available
	MaxAvailable *int `mandatory:"false" json:"maxAvailable"`
}

func (m UsageStatusItem) String() string {
	return common.PointerString(m)
}
