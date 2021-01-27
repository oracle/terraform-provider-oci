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

// LogAnalyticsLabelDefinition LogAnalyticsLabelDefinition
type LogAnalyticsLabelDefinition struct {

	// edit version
	EditVersion *int64 `mandatory:"false" json:"editVersion"`

	// is system flag
	IsSystem *bool `mandatory:"false" json:"isSystem"`

	// source Id
	SourceId *int64 `mandatory:"false" json:"sourceId"`

	// label name
	LabelName *string `mandatory:"false" json:"labelName"`
}

func (m LogAnalyticsLabelDefinition) String() string {
	return common.PointerString(m)
}
