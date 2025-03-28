// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
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

// UpdateExtract Parameters for Extract processes.
// If an empty object is specified, the stored Extract details will be removed.
type UpdateExtract struct {

	// Extract performance.
	PerformanceProfile ExtractPerformanceProfileEnum `mandatory:"false" json:"performanceProfile,omitempty"`

	// Length of time (in seconds) that a transaction can be open before Extract generates a warning message that the transaction is long-running.
	// If not specified, Extract will not generate a warning on long-running transactions.
	LongTransDuration *int `mandatory:"false" json:"longTransDuration"`
}

func (m UpdateExtract) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m UpdateExtract) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingExtractPerformanceProfileEnum(string(m.PerformanceProfile)); !ok && m.PerformanceProfile != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for PerformanceProfile: %s. Supported values are: %s.", m.PerformanceProfile, strings.Join(GetExtractPerformanceProfileEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
