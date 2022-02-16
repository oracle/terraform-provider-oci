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
	"fmt"
	"github.com/oracle/oci-go-sdk/v58/common"
	"strings"
)

// SqlStatistics Performance statistics for the SQL.
type SqlStatistics struct {

	// Database Time in seconds
	DatabaseTimeInSec *float64 `mandatory:"true" json:"databaseTimeInSec"`

	// Number of executions per hour
	ExecutionsPerHour *float64 `mandatory:"true" json:"executionsPerHour"`

	// Total number of executions
	ExecutionsCount *int64 `mandatory:"true" json:"executionsCount"`

	// CPU Time in seconds
	CpuTimeInSec *float64 `mandatory:"true" json:"cpuTimeInSec"`

	// I/O Time in seconds
	IoTimeInSec *float64 `mandatory:"true" json:"ioTimeInSec"`

	// Inefficient Wait Time in seconds
	InefficientWaitTimeInSec *float64 `mandatory:"true" json:"inefficientWaitTimeInSec"`

	// Response time is the average elaspsed time per execution. It is the ratio of Total Database Time to the number of executions
	ResponseTimeInSec *float64 `mandatory:"true" json:"responseTimeInSec"`

	// Number of SQL execution plans used by the SQL
	PlanCount *int64 `mandatory:"true" json:"planCount"`

	// Variability is the ratio of the standard deviation in response time to the mean of response time of the SQL
	Variability *float64 `mandatory:"true" json:"variability"`

	// Average Active Sessions represent the average active sessions at a point in time. It is the number of sessions that are either working or waiting.
	AverageActiveSessions *float64 `mandatory:"true" json:"averageActiveSessions"`

	// Percentage of Database Time
	DatabaseTimePct *float64 `mandatory:"true" json:"databaseTimePct"`

	// Percentage of Inefficiency. It is calculated by Total Database Time divided by Total Wait Time
	InefficiencyInPct *float64 `mandatory:"true" json:"inefficiencyInPct"`

	// Percent change in CPU Time based on linear regression
	ChangeInCpuTimeInPct *float64 `mandatory:"true" json:"changeInCpuTimeInPct"`

	// Percent change in IO Time based on linear regression
	ChangeInIoTimeInPct *float64 `mandatory:"true" json:"changeInIoTimeInPct"`

	// Percent change in Inefficient Wait Time based on linear regression
	ChangeInInefficientWaitTimeInPct *float64 `mandatory:"true" json:"changeInInefficientWaitTimeInPct"`

	// Percent change in Response Time based on linear regression
	ChangeInResponseTimeInPct *float64 `mandatory:"true" json:"changeInResponseTimeInPct"`

	// Percent change in Average Active Sessions based on linear regression
	ChangeInAverageActiveSessionsInPct *float64 `mandatory:"true" json:"changeInAverageActiveSessionsInPct"`

	// Percent change in Executions per hour based on linear regression
	ChangeInExecutionsPerHourInPct *float64 `mandatory:"true" json:"changeInExecutionsPerHourInPct"`

	// Percent change in Inefficiency based on linear regression
	ChangeInInefficiencyInPct *float64 `mandatory:"true" json:"changeInInefficiencyInPct"`
}

func (m SqlStatistics) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m SqlStatistics) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
