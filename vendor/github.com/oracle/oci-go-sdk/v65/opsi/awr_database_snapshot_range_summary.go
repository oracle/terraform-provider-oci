// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Ops Insights API
//
// Use the Ops Insights API to perform data extraction operations to obtain database
// resource utilization, performance statistics, and reference information. For more information,
// see About Oracle Cloud Infrastructure Ops Insights (https://docs.cloud.oracle.com/en-us/iaas/operations-insights/doc/operations-insights.html).
//

package opsi

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// AwrDatabaseSnapshotRangeSummary The summary data for a range of AWR snapshots.
type AwrDatabaseSnapshotRangeSummary struct {

	// The internal ID of the database. The internal ID of the database is not the OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm).
	// It can be retrieved from the following endpoint:
	// /awrHubs/{awrHubId}/awrDatabases
	AwrSourceDatabaseIdentifier *string `mandatory:"true" json:"awrSourceDatabaseIdentifier"`

	// The name of the database.
	DbName *string `mandatory:"true" json:"dbName"`

	// The database instance numbers.
	InstanceList []int `mandatory:"false" json:"instanceList"`

	// The timestamp of the database startup.
	TimeDbStartup *common.SDKTime `mandatory:"false" json:"timeDbStartup"`

	// The start time of the earliest snapshot.
	TimeFirstSnapshotBegin *common.SDKTime `mandatory:"false" json:"timeFirstSnapshotBegin"`

	// The end time of the latest snapshot.
	TimeLatestSnapshotEnd *common.SDKTime `mandatory:"false" json:"timeLatestSnapshotEnd"`

	// The ID of the earliest snapshot. The snapshot identifier is not the OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm).
	// It can be retrieved from the following endpoint:
	// /awrHubs/{awrHubId}/awrDatabaseSnapshots
	FirstSnapshotIdentifier *int `mandatory:"false" json:"firstSnapshotIdentifier"`

	// The ID of the latest snapshot. The snapshot identifier is not the OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm).
	// It can be retrieved from the following endpoint:
	// /awrHubs/{awrHubId}/awrDatabaseSnapshots
	LatestSnapshotIdentifier *int `mandatory:"false" json:"latestSnapshotIdentifier"`

	// The total number of snapshots.
	SnapshotCount *int64 `mandatory:"false" json:"snapshotCount"`

	// The interval time between snapshots (in minutes).
	SnapshotIntervalInMin *int `mandatory:"false" json:"snapshotIntervalInMin"`

	// The version of the database.
	DbVersion *string `mandatory:"false" json:"dbVersion"`

	// The time zone of the snapshot. sample -  snapshotTimezone=+0 00:00:00
	SnapshotTimezone *string `mandatory:"false" json:"snapshotTimezone"`
}

func (m AwrDatabaseSnapshotRangeSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m AwrDatabaseSnapshotRangeSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
