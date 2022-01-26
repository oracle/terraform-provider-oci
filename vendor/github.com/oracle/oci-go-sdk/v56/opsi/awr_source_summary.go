// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Operations Insights API
//
// Use the Operations Insights API to perform data extraction operations to obtain database
// resource utilization, performance statistics, and reference information. For more information,
// see About Oracle Cloud Infrastructure Operations Insights (https://docs.cloud.oracle.com/en-us/iaas/operations-insights/doc/operations-insights.html).
//

package opsi

import (
	"github.com/oracle/oci-go-sdk/v56/common"
)

// AwrSourceSummary Summary of an AwrSource.
type AwrSourceSummary struct {

	// AWR Hub OCID
	AwrHubId *string `mandatory:"true" json:"awrHubId"`

	// Database name of the Source database for which AWR Data will be uploaded to AWR Hub.
	Name *string `mandatory:"true" json:"name"`

	// DatabaseId of the Source database for which AWR Data will be uploaded to AWR Hub.
	AwrSourceDatabaseId *string `mandatory:"true" json:"awrSourceDatabaseId"`

	// Number of AWR snapshots uploaded from the Source database.
	SnapshotsUploaded *float32 `mandatory:"true" json:"snapshotsUploaded"`

	// The minimum snapshot identifier of the source database for which AWR data is uploaded to AWR Hub.
	MinSnapshotIdentifier *float32 `mandatory:"true" json:"minSnapshotIdentifier"`

	// The maximum snapshot identifier of the source database for which AWR data is uploaded to AWR Hub.
	MaxSnapshotIdentifier *float32 `mandatory:"true" json:"maxSnapshotIdentifier"`

	// The time at which the earliest snapshot was generated in the source database for which data is uploaded to AWR Hub. An RFC3339 formatted datetime string
	TimeFirstSnapshotGenerated *common.SDKTime `mandatory:"true" json:"timeFirstSnapshotGenerated"`

	// The time at which the latest snapshot was generated in the source database for which data is uploaded to AWR Hub. An RFC3339 formatted datetime string
	TimeLastSnapshotGenerated *common.SDKTime `mandatory:"true" json:"timeLastSnapshotGenerated"`

	// Number of hours since last AWR snapshots import happened from the Source database.
	HoursSinceLastImport *float64 `mandatory:"true" json:"hoursSinceLastImport"`
}

func (m AwrSourceSummary) String() string {
	return common.PointerString(m)
}
