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

// DbManagementAnalyticsMetric The metric details of a Database Management resource.
type DbManagementAnalyticsMetric struct {

	// The name of the metric.
	MetricName *string `mandatory:"false" json:"metricName"`

	// The duration of the returned aggregated data in seconds.
	DurationInSeconds *int64 `mandatory:"false" json:"durationInSeconds"`

	// The additional information about the metric.
	// Example: `"unit": "bytes"`
	Metadata map[string]string `mandatory:"false" json:"metadata"`

	// The qualifiers provided in the definition of the returned metric.
	Dimensions map[string]string `mandatory:"false" json:"dimensions"`

	// The start time associated with the value of the metric.
	StartTimestampInEpochSeconds *int64 `mandatory:"false" json:"startTimestampInEpochSeconds"`

	// The mean value of the metric.
	Mean *float64 `mandatory:"false" json:"mean"`
}

func (m DbManagementAnalyticsMetric) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m DbManagementAnalyticsMetric) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
