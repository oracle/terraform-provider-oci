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

// AwrDbParameterChangeSummary A summary of the changes made to a single AWR database parameter.
type AwrDbParameterChangeSummary struct {

	// The ID of the snapshot with the parameter value changed. The snapshot ID is not the OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm).
	// It can be retrieved from the following endpoint:
	// /managedDatabases/{managedDatabaseId}/awrDbs/{awrDbId}/awrDbSnapshots
	SnapshotId *int `mandatory:"true" json:"snapshotId"`

	// The start time of the interval.
	TimeBegin *common.SDKTime `mandatory:"false" json:"timeBegin"`

	// The end time of the interval.
	TimeEnd *common.SDKTime `mandatory:"false" json:"timeEnd"`

	// The database instance number.
	InstanceNumber *int `mandatory:"false" json:"instanceNumber"`

	// The previous value of the database parameter.
	PreviousValue *string `mandatory:"false" json:"previousValue"`

	// The current value of the database parameter.
	Value *string `mandatory:"false" json:"value"`

	// Indicates whether the parameter has been modified after instance startup:
	//  - MODIFIED - Parameter has been modified with ALTER SESSION
	//  - SYSTEM_MOD - Parameter has been modified with ALTER SYSTEM (which causes all the currently logged in sessions’ values to be modified)
	//  - FALSE - Parameter has not been modified after instance startup
	ValueModified *string `mandatory:"false" json:"valueModified"`

	// Indicates whether the parameter value in the end snapshot is the default.
	IsDefault *bool `mandatory:"false" json:"isDefault"`
}

func (m AwrDbParameterChangeSummary) String() string {
	return common.PointerString(m)
}
