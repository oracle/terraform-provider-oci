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
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// HistoricAddmResult The result of the creation and execution of the historic ADDM report, which contains the name of the ADDM task and the report.
type HistoricAddmResult struct {

	// The ID of the historic ADDM task.
	TaskId *int64 `mandatory:"true" json:"taskId"`

	// The name of the historic ADDM task.
	TaskName *string `mandatory:"false" json:"taskName"`

	// The owner of the historic ADDM task.
	DbUser *string `mandatory:"false" json:"dbUser"`

	// The timestamp of the beginning AWR snapshot used in the ADDM report as defined by date-time RFC3339 format.
	StartSnapShotTime *common.SDKTime `mandatory:"false" json:"startSnapShotTime"`

	// The timestamp of the ending AWR snapshot used in the ADDM report as defined by date-time RFC3339 format.
	EndSnapshotTime *common.SDKTime `mandatory:"false" json:"endSnapshotTime"`

	// The ID number of the beginning AWR snapshot.
	BeginSnaphotId *int64 `mandatory:"false" json:"beginSnaphotId"`

	// The ID number of the ending AWR snapshot.
	EndSnapshotId *int64 `mandatory:"false" json:"endSnapshotId"`
}

func (m HistoricAddmResult) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m HistoricAddmResult) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
