// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// LogAnalytics API
//
// The LogAnalytics API for the LogAnalytics service.
//

package loganalytics

import (
	"github.com/oracle/oci-go-sdk/v31/common"
)

// LogAnalyticsImportCustomChangeList LogAnalyticsImportCustomChangeList
type LogAnalyticsImportCustomChangeList struct {

	// createdParserNames
	CreatedParserNames []string `mandatory:"false" json:"createdParserNames"`

	// updatedParserNames
	UpdatedParserNames []string `mandatory:"false" json:"updatedParserNames"`

	// createdSourceNames
	CreatedSourceNames []string `mandatory:"false" json:"createdSourceNames"`

	// updatedSourceNames
	UpdatedSourceNames []string `mandatory:"false" json:"updatedSourceNames"`

	// createdFieldDisplayNames
	CreatedFieldDisplayNames []string `mandatory:"false" json:"createdFieldDisplayNames"`

	// updatedFieldDisplayNames
	UpdatedFieldDisplayNames []string `mandatory:"false" json:"updatedFieldDisplayNames"`

	// conflictParserNames
	ConflictParserNames []string `mandatory:"false" json:"conflictParserNames"`

	// conflictSourceNames
	ConflictSourceNames []string `mandatory:"false" json:"conflictSourceNames"`

	// conflictFieldDisplayNames
	ConflictFieldDisplayNames []string `mandatory:"false" json:"conflictFieldDisplayNames"`
}

func (m LogAnalyticsImportCustomChangeList) String() string {
	return common.PointerString(m)
}
