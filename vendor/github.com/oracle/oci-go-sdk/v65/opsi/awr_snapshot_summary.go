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

// AwrSnapshotSummary The AWR snapshot summary of one snapshot.
type AwrSnapshotSummary struct {

	// DatabaseId of the Source database for which AWR Data will be uploaded to AWR Hub.
	AwrSourceDatabaseId *string `mandatory:"true" json:"awrSourceDatabaseId"`

	// The identifier of the snapshot.
	SnapshotIdentifier *int `mandatory:"true" json:"snapshotIdentifier"`

	// The database instance number.
	InstanceNumber *int `mandatory:"false" json:"instanceNumber"`

	// The timestamp of the database startup.
	TimeDbStartup *common.SDKTime `mandatory:"false" json:"timeDbStartup"`

	// The start time of the snapshot.
	TimeSnapshotBegin *common.SDKTime `mandatory:"false" json:"timeSnapshotBegin"`

	// The end time of the snapshot.
	TimeSnapshotEnd *common.SDKTime `mandatory:"false" json:"timeSnapshotEnd"`

	// The total number of errors.
	ErrorCount *int64 `mandatory:"false" json:"errorCount"`
}

func (m AwrSnapshotSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m AwrSnapshotSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
