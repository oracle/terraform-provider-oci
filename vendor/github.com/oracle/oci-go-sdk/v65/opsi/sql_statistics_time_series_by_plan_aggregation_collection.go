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

// SqlStatisticsTimeSeriesByPlanAggregationCollection SQL performance statistics by plan over the selected time window.
type SqlStatisticsTimeSeriesByPlanAggregationCollection struct {

	// Unique SQL_ID for a SQL Statement.
	SqlIdentifier *string `mandatory:"true" json:"sqlIdentifier"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the database insight resource.
	Id *string `mandatory:"true" json:"id"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the database.
	DatabaseId *string `mandatory:"true" json:"databaseId"`

	// The start timestamp that was passed into the request.
	TimeIntervalStart *common.SDKTime `mandatory:"true" json:"timeIntervalStart"`

	// The end timestamp that was passed into the request.
	TimeIntervalEnd *common.SDKTime `mandatory:"true" json:"timeIntervalEnd"`

	// Time duration in milliseconds between data points (one hour or one day).
	ItemDurationInMs *int64 `mandatory:"true" json:"itemDurationInMs"`

	// Array comprising of all the sampling period end timestamps in RFC 3339 format.
	EndTimestamps []common.SDKTime `mandatory:"true" json:"endTimestamps"`

	// array of SQL performance statistics by plans
	Items []SqlStatisticsTimeSeriesByPlanAggregation `mandatory:"true" json:"items"`
}

func (m SqlStatisticsTimeSeriesByPlanAggregationCollection) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m SqlStatisticsTimeSeriesByPlanAggregationCollection) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
