// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
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
	"github.com/oracle/oci-go-sdk/v56/common"
)

// AwrDbSnapshotSummary The AWR snapshot summary of one snapshot.
type AwrDbSnapshotSummary struct {

	// Internal ID of the database. The internal ID of the database is not the OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm).
	// It can be retrieved from the following endpoint:
	// /managedDatabases/{managedDatabaseId}/awrDbs
	AwrDbId *string `mandatory:"true" json:"awrDbId"`

	// The ID of the snapshot. The snapshot ID is not the OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm).
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
