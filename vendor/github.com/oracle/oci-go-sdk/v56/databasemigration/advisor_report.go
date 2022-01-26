// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Migration API
//
// Use the Oracle Cloud Infrastructure Database Migration APIs to perform database migration operations.
//

package databasemigration

import (
	"github.com/oracle/oci-go-sdk/v56/common"
)

// AdvisorReport Pre-Migration advisor report details.
type AdvisorReport struct {

	// Pre-Migration advisor result.
	Result AdvisorResultsEnum `mandatory:"true" json:"result"`

	// Number of Fatal results in the advisor report.
	NumberOfFatal *int `mandatory:"true" json:"numberOfFatal"`

	// Number of Fatal Blocker results in the advisor report.
	NumberOfFatalBlockers *int `mandatory:"true" json:"numberOfFatalBlockers"`

	// Number of Warning results in the advisor report.
	NumberOfWarnings *int `mandatory:"true" json:"numberOfWarnings"`

	// Number of Informational results in the advisor report.
	NumberOfInformationalResults *int `mandatory:"true" json:"numberOfInformationalResults"`

	ReportLocationDetails *AdvisorReportLocationDetails `mandatory:"false" json:"reportLocationDetails"`
}

func (m AdvisorReport) String() string {
	return common.PointerString(m)
}
