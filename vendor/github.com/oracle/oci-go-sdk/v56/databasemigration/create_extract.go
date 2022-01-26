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

// CreateExtract Parameters for GoldenGate Extract processes.
type CreateExtract struct {

	// Extract performance.
	PerformanceProfile ExtractPerformanceProfileEnum `mandatory:"false" json:"performanceProfile,omitempty"`

	// Length of time (in seconds) that a transaction can be open before Extract generates a warning message that the transaction is long-running.
	// If not specified, Extract will not generate a warning on long-running transactions.
	LongTransDuration *int `mandatory:"false" json:"longTransDuration"`
}

func (m CreateExtract) String() string {
	return common.PointerString(m)
}
