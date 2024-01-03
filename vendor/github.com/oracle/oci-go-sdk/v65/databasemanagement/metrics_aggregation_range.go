// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
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

// MetricsAggregationRange The set of aggregated data returned for a metric.
type MetricsAggregationRange struct {
	Header *DbManagementAnalyticsMetric `mandatory:"false" json:"header"`

	// The list of metrics returned for the specified request. Each of the metrics
	// has a `metricName` and additional properties like `metadata`, `dimensions`.
	// If a property is not set, then use the value from `header`.
	// Suppose `m` be an item in the `metrics` array:
	// - If `m.metricName` is not set, use `header.metricName` instead
	// - If `m.durationInSeconds` is not set, use `header.durationInSeconds` instead
	// - If `m.dimensions` is not set, use `header.dimensions` instead
	// - If `m.metadata` is not set, use `header.metadata` instead
	Metrics []DbManagementAnalyticsMetric `mandatory:"false" json:"metrics"`

	// The beginning of the time range (inclusive) of the returned metric data.
	RangeStartTimeInEpochSeconds *int64 `mandatory:"false" json:"rangeStartTimeInEpochSeconds"`

	// The end of the time range (exclusive) of the returned metric data.
	RangeEndTimeInEpochSeconds *int64 `mandatory:"false" json:"rangeEndTimeInEpochSeconds"`
}

func (m MetricsAggregationRange) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m MetricsAggregationRange) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
