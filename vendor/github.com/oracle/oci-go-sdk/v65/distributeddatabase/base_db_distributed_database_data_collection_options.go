// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Globally Distributed Database
//
// Use the Globally Distributed Database service APIs to create and manage the Globally distributed databases.
//

package distributeddatabase

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// BaseDbDistributedDatabaseDataCollectionOptions Indicates user preferences for the various diagnostic collection options for the base database.
type BaseDbDistributedDatabaseDataCollectionOptions struct {

	// Indicates whether diagnostic collection is enabled for the base database.
	IsDiagnosticsEventsEnabled *bool `mandatory:"false" json:"isDiagnosticsEventsEnabled"`

	// Indicates whether health monitoring is enabled for the the base database.
	IsHealthMonitoringEnabled *bool `mandatory:"false" json:"isHealthMonitoringEnabled"`

	// Indicates whether incident logs and trace collection are enabled for the base database.
	IsIncidentLogsEnabled *bool `mandatory:"false" json:"isIncidentLogsEnabled"`
}

func (m BaseDbDistributedDatabaseDataCollectionOptions) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m BaseDbDistributedDatabaseDataCollectionOptions) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}
