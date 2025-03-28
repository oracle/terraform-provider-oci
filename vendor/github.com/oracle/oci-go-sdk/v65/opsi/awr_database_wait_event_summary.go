// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Ops Insights API
//
// Use the Ops Insights API to perform data extraction operations to obtain database
// resource utilization, performance statistics, and reference information. For more information,
// see About Oracle Cloud Infrastructure Ops Insights (https://docs.oracle.com/iaas/en-us/iaas/operations-insights/doc/operations-insights.html).
//

package opsi

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// AwrDatabaseWaitEventSummary The summary of the AWR wait event time series data for one event.
type AwrDatabaseWaitEventSummary struct {

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

	// The ID of the snapshot. The snapshot identifier is not the OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm).
	// It can be retrieved from the following endpoint:
	// /awrHubs/{awrHubId}/awrDatabaseSnapshots
	SnapshotIdentifier *int `mandatory:"false" json:"snapshotIdentifier"`
}

func (m AwrDatabaseWaitEventSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m AwrDatabaseWaitEventSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
