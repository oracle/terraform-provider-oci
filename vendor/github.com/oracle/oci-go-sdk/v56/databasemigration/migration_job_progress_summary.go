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

// MigrationJobProgressSummary Summary of the progress of a Migration Job.
type MigrationJobProgressSummary struct {

	// Current phase of the job.
	CurrentPhase OdmsJobPhasesEnum `mandatory:"true" json:"currentPhase"`

	// Current status of the job.
	CurrentStatus JobPhaseStatusEnum `mandatory:"true" json:"currentStatus"`

	// Job progress percentage (0 - 100)
	JobProgress *int `mandatory:"true" json:"jobProgress"`
}

func (m MigrationJobProgressSummary) String() string {
	return common.PointerString(m)
}
