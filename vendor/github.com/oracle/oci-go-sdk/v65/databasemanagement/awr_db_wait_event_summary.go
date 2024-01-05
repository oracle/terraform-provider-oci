// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Management API
//
// Use the Database Management API to perform tasks such as obtaining performance and resource usage metrics
// for a fleet of Managed Databases or a specific Managed Database, creating Managed Database Groups, and
// running a SQL job on a Managed Database or Managed Database Group.
//

package databasemanagement

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// AwrDbWaitEventSummary The summary of the AWR wait event time series data for one event.
type AwrDbWaitEventSummary struct {

	// The name of the event.
	Name *string `mandatory:"true" json:"name"`

	// The begin time of the wait event.
	TimeBegin *common.SDKTime `mandatory:"false" json:"timeBegin"`

	// The end time of the wait event.
	TimeEnd *common.SDKTime `mandatory:"false" json:"timeEnd"`

	// The wait count per second.
	WaitsPerSec *float64 `mandatory:"false" json:"waitsPerSec"`

	// The average wait time per second.
	AvgWaitTimePerSec *float64 `mandatory:"false" json:"avgWaitTimePerSec"`

	// The average wait time in milliseconds per wait.
	AvgWaitTimePerWait *float64 `mandatory:"false" json:"avgWaitTimePerWait"`

	// The ID of the snapshot. The snapshot ID is not the OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm).
	// It can be retrieved from the following endpoint:
	// /managedDatabases/{managedDatabaseId}/awrDbs/{awrDbId}/awrDbSnapshots
	SnapshotId *int `mandatory:"false" json:"snapshotId"`
}

func (m AwrDbWaitEventSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m AwrDbWaitEventSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
