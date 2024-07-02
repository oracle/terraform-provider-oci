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

// AddmDbSummary ADDM summary for a database
type AddmDbSummary struct {
	DatabaseDetails *DatabaseDetails `mandatory:"true" json:"databaseDetails"`

	// Number of ADDM findings
	NumberOfFindings *int `mandatory:"false" json:"numberOfFindings"`

	// Number of ADDM tasks
	NumberOfAddmTasks *int `mandatory:"false" json:"numberOfAddmTasks"`

	// The start timestamp that was passed into the request.
	TimeFirstSnapshotBegin *common.SDKTime `mandatory:"false" json:"timeFirstSnapshotBegin"`

	// The end timestamp that was passed into the request.
	TimeLatestSnapshotEnd *common.SDKTime `mandatory:"false" json:"timeLatestSnapshotEnd"`

	// AWR snapshot id.
	SnapshotIntervalStart *string `mandatory:"false" json:"snapshotIntervalStart"`

	// AWR snapshot id.
	SnapshotIntervalEnd *string `mandatory:"false" json:"snapshotIntervalEnd"`

	// Maximum overall impact in terms of percentage of total activity
	MaxOverallImpact *float64 `mandatory:"false" json:"maxOverallImpact"`

	// Category name
	MostFrequentCategoryName *string `mandatory:"false" json:"mostFrequentCategoryName"`

	// Category display name
	MostFrequentCategoryDisplayName *string `mandatory:"false" json:"mostFrequentCategoryDisplayName"`
}

func (m AddmDbSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m AddmDbSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
