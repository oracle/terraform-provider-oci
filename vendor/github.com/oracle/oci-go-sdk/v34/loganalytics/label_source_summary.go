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

// LabelSourceSummary source summary
type LabelSourceSummary struct {

	// display name
	SourceDisplayName *string `mandatory:"false" json:"sourceDisplayName"`

	// source internal name
	SourceName *string `mandatory:"false" json:"sourceName"`

	// source Id
	SourceId *int64 `mandatory:"false" json:"sourceId"`

	// label Operator
	LabelOperatorName *string `mandatory:"false" json:"labelOperatorName"`

	// label Condition
	LabelCondition *string `mandatory:"false" json:"labelCondition"`

	// label Field Display Name
	LabelFieldDisplayname *string `mandatory:"false" json:"labelFieldDisplayname"`

	// label Field name
	LabelFieldName *string `mandatory:"false" json:"labelFieldName"`
}

func (m LabelSourceSummary) String() string {
	return common.PointerString(m)
}
