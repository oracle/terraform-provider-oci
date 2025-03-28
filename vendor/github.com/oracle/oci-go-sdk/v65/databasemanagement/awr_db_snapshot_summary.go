// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Management API
//
// Use the Database Management API to monitor and manage resources such as
// Oracle Databases, MySQL Databases, and External Database Systems.
// For more information, see Database Management (https://docs.oracle.com/iaas/database-management/home.htm).
//

package databasemanagement

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// AwrDbSnapshotSummary The AWR snapshot summary of one snapshot.
type AwrDbSnapshotSummary struct {

	// Internal ID of the database. The internal ID of the database is not the OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm).
	// It can be retrieved from the following endpoint:
	// /managedDatabases/{managedDatabaseId}/awrDbs
	AwrDbId *string `mandatory:"true" json:"awrDbId"`

	// The ID of the snapshot. The snapshot ID is not the OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm).
	// It can be retrieved from the following endpoint:
	// /managedDatabases/{managedDatabaseId}/awrDbs/{awrDbId}/awrDbSnapshots
	SnapshotId *int `mandatory:"true" json:"snapshotId"`

	// The database instance number.
	InstanceNumber *int `mandatory:"false" json:"instanceNumber"`

	// The timestamp of the database startup.
	TimeDbStartup *common.SDKTime `mandatory:"false" json:"timeDbStartup"`

	// The start time of the snapshot.
	TimeBegin *common.SDKTime `mandatory:"false" json:"timeBegin"`

	// The end time of the snapshot.
	TimeEnd *common.SDKTime `mandatory:"false" json:"timeEnd"`

	// The total number of errors.
	ErrorCount *int64 `mandatory:"false" json:"errorCount"`
}

func (m AwrDbSnapshotSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m AwrDbSnapshotSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
