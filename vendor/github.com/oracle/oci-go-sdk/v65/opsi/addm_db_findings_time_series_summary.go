// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
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
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// AddmDbFindingsTimeSeriesSummary ADDM findings time series data
type AddmDbFindingsTimeSeriesSummary struct {

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Database insight.
	Id *string `mandatory:"true" json:"id"`

	// Unique ADDM task id
	TaskId *int `mandatory:"true" json:"taskId"`

	// ADDM task name
	TaskName *string `mandatory:"true" json:"taskName"`

	// Unique finding id
	FindingId *string `mandatory:"true" json:"findingId"`

	// Timestamp when finding was generated
	Timestamp *common.SDKTime `mandatory:"true" json:"timestamp"`

	// Category name
	CategoryName *string `mandatory:"true" json:"categoryName"`

	// Category display name
	CategoryDisplayName *string `mandatory:"true" json:"categoryDisplayName"`

	// Finding name
	Name *string `mandatory:"true" json:"name"`

	// Finding message
	Message *string `mandatory:"true" json:"message"`

	// Impact in terms of percentage of total activity
	ImpactPercent *float64 `mandatory:"true" json:"impactPercent"`

	// Impact in terms of average active sessions
	ImpactAvgActiveSessions *float64 `mandatory:"true" json:"impactAvgActiveSessions"`

	// Start Timestamp of snapshot
	TimeAnalysisStarted *common.SDKTime `mandatory:"false" json:"timeAnalysisStarted"`

	// End Timestamp of snapshot
	TimeAnalysisEnded *common.SDKTime `mandatory:"false" json:"timeAnalysisEnded"`

	// DB time in seconds for the snapshot
	AnalysisDbTimeInSecs *float64 `mandatory:"false" json:"analysisDbTimeInSecs"`

	// DB avg active sessions for the snapshot
	AnalysisAvgActiveSessions *float64 `mandatory:"false" json:"analysisAvgActiveSessions"`

	// Impact in seconds
	ImpactDbTimeInSecs *float64 `mandatory:"false" json:"impactDbTimeInSecs"`
}

func (m AddmDbFindingsTimeSeriesSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m AddmDbFindingsTimeSeriesSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
