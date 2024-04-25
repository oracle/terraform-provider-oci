// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Migration API
//
// Use the Oracle Cloud Infrastructure Database Migration APIs to perform database migration operations.
//

package databasemigration

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// CreateReplicat Parameters for GoldenGate Replicat processes.
type CreateReplicat struct {

	// Replicat performance.
	PerformanceProfile ReplicatPerformanceProfileEnum `mandatory:"false" json:"performanceProfile,omitempty"`

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

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CreateReplicat) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingReplicatPerformanceProfileEnum(string(m.PerformanceProfile)); !ok && m.PerformanceProfile != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for PerformanceProfile: %s. Supported values are: %s.", m.PerformanceProfile, strings.Join(GetReplicatPerformanceProfileEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
