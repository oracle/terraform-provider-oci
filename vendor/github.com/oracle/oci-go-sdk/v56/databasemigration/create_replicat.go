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

// CreateReplicat Parameters for GoldenGate Replicat processes.
type CreateReplicat struct {

	// Number of threads used to read trail files (valid for Parallel Replicat)
	MapParallelism *int `mandatory:"false" json:"mapParallelism"`

	// Defines the range in which the Replicat automatically adjusts its apply parallelism (valid for Parallel Replicat)
	MinApplyParallelism *int `mandatory:"false" json:"minApplyParallelism"`

	// Defines the range in which the Replicat automatically adjusts its apply parallelism (valid for Parallel Replicat)
	MaxApplyParallelism *int `mandatory:"false" json:"maxApplyParallelism"`
}

func (m CreateReplicat) String() string {
	return common.PointerString(m)
}
