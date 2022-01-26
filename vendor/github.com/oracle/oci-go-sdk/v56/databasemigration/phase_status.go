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

// PhaseStatus Job phase status details.
type PhaseStatus struct {

	// Phase name
	Name OdmsJobPhasesEnum `mandatory:"true" json:"name"`

	// Phase status
	Status JobPhaseStatusEnum `mandatory:"true" json:"status"`

	// Duration of the phase in milliseconds
	DurationInMs *int `mandatory:"true" json:"durationInMs"`

	// True if a Pre-Migration Advisor report is available for this phase. False or null if no report is available.
	IsAdvisorReportAvailable *bool `mandatory:"false" json:"isAdvisorReportAvailable"`

	// Summary of phase status results.
	Extract []PhaseExtractEntry `mandatory:"false" json:"extract"`

	LogLocation *LogLocationBucketDetails `mandatory:"false" json:"logLocation"`

	// Percent progress of job phase.
	Progress *int `mandatory:"false" json:"progress"`
}

func (m PhaseStatus) String() string {
	return common.PointerString(m)
}
